/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 14:27:35
 * @FilePath: \go-config\config.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"fmt"
	"reflect"

	"github.com/kamalyes/go-config/captcha"
	"github.com/kamalyes/go-config/consul"
	"github.com/kamalyes/go-config/cors"
	"github.com/kamalyes/go-config/database"
	"github.com/kamalyes/go-config/email"
	"github.com/kamalyes/go-config/ftp"
	"github.com/kamalyes/go-config/jwt"
	"github.com/kamalyes/go-config/mqtt"
	"github.com/kamalyes/go-config/oss"
	"github.com/kamalyes/go-config/pay"
	"github.com/kamalyes/go-config/redis"
	"github.com/kamalyes/go-config/server"
	"github.com/kamalyes/go-config/sms"
	"github.com/kamalyes/go-config/sts"
	"github.com/kamalyes/go-config/youzan"
	"github.com/kamalyes/go-config/zap"
	"github.com/spf13/viper"
)

// Configurable 接口
type Configurable interface {
	GetModuleName() string
}

// Config 项目统一配置
type Config struct {
	Server     []server.Server       `mapstructure:"server" json:"server" yaml:"server"`
	Cors       []cors.Cors           `mapstructure:"cors" json:"cors" yaml:"cors"`
	Consul     []consul.Consul       `mapstructure:"consul" json:"consul" yaml:"consul"`
	Captcha    []captcha.Captcha     `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	MySQL      []database.MySQL      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PostgreSQL []database.PostgreSQL `mapstructure:"postgre" json:"postgre" yaml:"postgre"`
	SQLite     []database.SQLite     `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Redis      []redis.Redis         `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email      []email.Email         `mapstructure:"email" json:"email" yaml:"email"`
	Ftp        []ftp.Ftp             `mapstructure:"ftp" json:"ftp" yaml:"ftp"`
	JWT        []jwt.JWT             `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Minio      []oss.Minio           `mapstructure:"minio" json:"minio" yaml:"minio"`
	Mqtt       []mqtt.Mqtt           `mapstructure:"mqtt" json:"mqtt" yaml:"mqtt"`
	Zap        []zap.Zap             `mapstructure:"zap" json:"zap" yaml:"zap"`
	AliPay     []pay.Alipay          `mapstructure:"alipay" json:"alipay" yaml:"alipay"`
	Wechat     []pay.Wechat          `mapstructure:"wechat" json:"wechat" yaml:"wechat"`
	AliyunSms  []sms.AliyunSms       `mapstructure:"aliyunsms" json:"aliyunsms" yaml:"aliyunsms"`
	AliyunSts  []sts.AliyunSts       `mapstructure:"aliyunsts" json:"aliyunsts" yaml:"aliyunsts"`
	Youzan     []youzan.YouZan       `mapstructure:"youzan" json:"youzan" yaml:"youzan"`
	Viper      *viper.Viper          `mapstructure:"-" json:"-" yaml:"-"`
}

// GetModules 返回指定模块的所有实例
func GetModules[T any](config *Config) ([]T, error) {
	moduleType := reflect.TypeOf((*T)(nil)).Elem()
	configValue := reflect.ValueOf(config).Elem()

	for i := 0; i < configValue.NumField(); i++ {
		field := configValue.Type().Field(i)
		fieldValue := configValue.Field(i)

		if fieldValue.Kind() == reflect.Slice && field.Type.Elem() == moduleType {
			modules := fieldValue.Interface().([]T)
			return modules, nil
		}
	}

	return nil, fmt.Errorf("no modules of type %T found", new(T))
}

// GetConfigByModuleName 根据模块名称返回包含该模块的配置
func GetConfigByModuleName(config *Config, moduleName string) (*Config, error) {
	result := &Config{}
	configValue := reflect.ValueOf(config).Elem()

	for i := 0; i < configValue.NumField(); i++ {
		fieldValue := configValue.Field(i)

		if fieldValue.Kind() == reflect.Slice {
			matchingModules := reflect.MakeSlice(fieldValue.Type(), 0, 0)

			for j := 0; j < fieldValue.Len(); j++ {
				module := fieldValue.Index(j)
				moduleNameValue := module.FieldByName("ModuleName").String()

				if moduleNameValue == moduleName {
					matchingModules = reflect.Append(matchingModules, module)
				}
			}

			if matchingModules.Len() > 0 {
				resultValue := reflect.ValueOf(result).Elem()
				resultValue.Field(i).Set(matchingModules)
			}
		}
	}

	if reflect.DeepEqual(result, &Config{}) {
		return nil, fmt.Errorf("no modules found with name %s", moduleName)
	}

	return result, nil
}

// GetModuleByName 使用泛型来获取模块
func GetModuleByName[T any](modules []T, moduleName string) (T, error) {
	// 使用反射获取模块的字段
	moduleType := reflect.TypeOf((*T)(nil)).Elem()
	moduleField, ok := moduleType.FieldByName("ModuleName")
	if !ok {
		return *new(T), fmt.Errorf("type %T does not have a ModuleName field", new(T))
	}

	// 检查模块切片的长度
	switch len(modules) {
	case 0:
		return *new(T), fmt.Errorf("no modules available")
	default:
		// 长度大于1时，遍历模块
		for _, module := range modules {
			// 通过反射获取ModuleName的值
			moduleValue := reflect.ValueOf(module).FieldByName(moduleField.Name).String()
			if moduleValue == moduleName {
				return module, nil
			}
		}
	}

	// 如果没有找到，返回错误
	return *new(T), fmt.Errorf("module %s not found", moduleName)
}

// LoadConfig 从指定的配置文件加载配置
func LoadConfig(filePath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	config.Viper = v // 将 Viper 实例分配到配置中
	return &config, nil
}
