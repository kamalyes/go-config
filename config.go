/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 08:58:09
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 22:11:22
 * @FilePath: \go-config\config.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/kamalyes/go-config/pkg/captcha"
	"github.com/kamalyes/go-config/pkg/cors"
	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-config/pkg/email"
	"github.com/kamalyes/go-config/pkg/ftp"
	"github.com/kamalyes/go-config/pkg/jwt"
	"github.com/kamalyes/go-config/pkg/oss"
	"github.com/kamalyes/go-config/pkg/pay"
	"github.com/kamalyes/go-config/pkg/queue"
	"github.com/kamalyes/go-config/pkg/redis"
	"github.com/kamalyes/go-config/pkg/register"
	"github.com/kamalyes/go-config/pkg/sms"
	"github.com/kamalyes/go-config/pkg/sts"
	"github.com/kamalyes/go-config/pkg/youzan"
	"github.com/kamalyes/go-config/pkg/zap"
	"github.com/kamalyes/go-config/pkg/zero"
	"github.com/spf13/viper"
)

// MultiConfig 表示多个配置
type MultiConfig struct {
	Server        []register.Server     `mapstructure:"server"       yaml:"server"       json:"server"`
	Cors          []cors.Cors           `mapstructure:"cors"         yaml:"cors"         json:"cors"`
	Consul        []register.Consul     `mapstructure:"consul"       yaml:"consul"       json:"consul"`
	Captcha       []captcha.Captcha     `mapstructure:"captcha"      yaml:"captcha"      json:"captcha"`
	MySQL         []database.MySQL      `mapstructure:"mysql"        yaml:"mysql"        json:"mysql"`
	PostgreSQL    []database.PostgreSQL `mapstructure:"postgre"      yaml:"postgre"      json:"postgre"`
	SQLite        []database.SQLite     `mapstructure:"sqlite"       yaml:"sqlite"       json:"sqlite"`
	Redis         []redis.Redis         `mapstructure:"redis"        yaml:"redis"        json:"redis"`
	Email         []email.Email         `mapstructure:"email"        yaml:"email"        json:"email"`
	Ftp           []ftp.Ftp             `mapstructure:"ftp"          yaml:"ftp"          json:"ftp"`
	JWT           []jwt.JWT             `mapstructure:"jwt"          yaml:"jwt"          json:"jwt"`
	Minio         []oss.Minio           `mapstructure:"minio"        yaml:"minio"        json:"minio"`
	Mqtt          []queue.Mqtt          `mapstructure:"mqtt"         yaml:"mqtt"         json:"mqtt"`
	Zap           []zap.Zap             `mapstructure:"zap"          yaml:"zap"          json:"zap"`
	AliPay        []pay.Alipay          `mapstructure:"alipay"       yaml:"alipay"       json:"alipay"`
	Wechat        []pay.Wechat          `mapstructure:"wechat"       yaml:"wechat"       json:"wechat"`
	AliyunSms     []sms.AliyunSms       `mapstructure:"aliyun-sms"   yaml:"aliyun-sms"    json:"aliyun_sms"`
	AliyunSts     []sts.AliyunSts       `mapstructure:"aliyun-sts"   yaml:"aliyun-sts"    json:"aliyun_sts"`
	Youzan        []youzan.YouZan       `mapstructure:"youzan"       yaml:"youzan"       json:"youzan"`
	ZeroRpcServer []zero.RpcServer      `mapstructure:"zero-rpc-server" yaml:"zero-rpc-server" json:"zero_rpc_server"`
	ZeroRpcClient []zero.RpcClient      `mapstructure:"zero-rpc-client" yaml:"zero-rpc-client" json:"zero_rpc_client"`
	Viper         *viper.Viper          `mapstructure:"-"            yaml:"-"            json:"-"`
}

// SingleConfig 单一配置
type SingleConfig struct {
	Server        register.Server     `mapstructure:"server"       yaml:"server"       json:"server"`
	Cors          cors.Cors           `mapstructure:"cors"         yaml:"cors"         json:"cors"`
	Consul        register.Consul     `mapstructure:"consul"       yaml:"consul"       json:"consul"`
	Captcha       captcha.Captcha     `mapstructure:"captcha"      yaml:"captcha"      json:"captcha"`
	MySQL         database.MySQL      `mapstructure:"mysql"        yaml:"mysql"        json:"mysql"`
	PostgreSQL    database.PostgreSQL `mapstructure:"postgre"      yaml:"postgre"      json:"postgre"`
	SQLite        database.SQLite     `mapstructure:"sqlite"       yaml:"sqlite"       json:"sqlite"`
	Redis         redis.Redis         `mapstructure:"redis"        yaml:"redis"        json:"redis"`
	Email         email.Email         `mapstructure:"email"        yaml:"email"        json:"email"`
	Ftp           ftp.Ftp             `mapstructure:"ftp"          yaml:"ftp"          json:"ftp"`
	JWT           jwt.JWT             `mapstructure:"jwt"          yaml:"jwt"          json:"jwt"`
	Minio         oss.Minio           `mapstructure:"minio"        yaml:"minio"        json:"minio"`
	Mqtt          queue.Mqtt          `mapstructure:"mqtt"         yaml:"mqtt"         json:"mqtt"`
	Zap           zap.Zap             `mapstructure:"zap"          yaml:"zap"          json:"zap"`
	AliPay        pay.Alipay          `mapstructure:"alipay"       yaml:"alipay"       json:"alipay"`
	Wechat        pay.Wechat          `mapstructure:"wechat"       yaml:"wechat"       json:"wechat"`
	AliyunSms     sms.AliyunSms       `mapstructure:"aliyun-sms"   yaml:"aliyun-sms"    json:"aliyun_sms"`
	AliyunSts     sts.AliyunSts       `mapstructure:"aliyun-sts"   yaml:"aliyun-sts"    json:"aliyun_sts"`
	Youzan        youzan.YouZan       `mapstructure:"youzan"       yaml:"youzan"       json:"youzan"`
	ZeroRpcServer zero.RpcServer      `mapstructure:"zero-rpc-server" yaml:"zero-rpc-server" json:"zero_rpc_server"`
	ZeroRpcClient zero.RpcClient      `mapstructure:"zero-rpc-client" yaml:"zero-rpc-client" json:"zero_rpc_client"`
	Viper         *viper.Viper        `mapstructure:"-"            yaml:"-"            json:"-"`
}

// GetModules 返回指定模块的所有实例
func GetModules[T any](config *MultiConfig) ([]T, error) {
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
func GetConfigByModuleName(config *MultiConfig, moduleName string) (*MultiConfig, error) {
	result := &MultiConfig{}
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

	if reflect.DeepEqual(result, &MultiConfig{}) {
		return nil, fmt.Errorf("no modules found with name %s", moduleName)
	}

	return result, nil
}

// GetModuleByName 使用泛型来获取模块
func GetModuleByName[T any](modules []T, moduleName string) (T, error) {
	var zeroValue T // 用于返回的零值

	// 检查模块切片的长度
	if len(modules) == 0 {
		return zeroValue, errors.New("no modules available")
	}

	// 检查第一个模块是否包含 ModuleName 字段
	moduleValue, err := getModuleName(modules[0])
	if err != nil {
		return zeroValue, err
	}

	// 如果模块切片长度为1，直接检查该模块
	if len(modules) == 1 {
		if moduleValue == "" || moduleValue == moduleName {
			return modules[0], nil
		}
		return zeroValue, fmt.Errorf("module %s not found", moduleName)
	}

	// 长度大于1时，遍历模块
	for _, module := range modules {
		moduleValue, err := getModuleName(module)
		if err != nil {
			return zeroValue, err
		}
		if moduleValue == moduleName {
			return module, nil
		}
	}

	return zeroValue, fmt.Errorf("module %s not found", moduleName)
}

// getModuleName 使用反射获取模块的 ModuleName 字段
func getModuleName[T any](module T) (string, error) {
	moduleType := reflect.TypeOf(module)
	if field, ok := moduleType.FieldByName("ModuleName"); ok {
		moduleValue := reflect.ValueOf(module).FieldByName(field.Name)
		if moduleValue.IsValid() {
			return moduleValue.String(), nil
		}
	}
	return "", fmt.Errorf("type %T does not have a ModuleName field", module)
}
