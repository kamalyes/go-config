/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 08:58:09
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 03:01:26
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
	"github.com/kamalyes/go-config/pkg/elk"
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

// MultiConfig 多配置
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
	AliyunOss     []oss.AliyunOss       `mapstructure:"aliyunoss"    yaml:"aliyunoss"    json:"aliyunoss"`
	S3            []oss.S3              `mapstructure:"s3"           yaml:"s3"           json:"s3"`
	Mqtt          []queue.Mqtt          `mapstructure:"mqtt"         yaml:"mqtt"         json:"mqtt"`
	Zap           []zap.Zap             `mapstructure:"zap"          yaml:"zap"          json:"zap"`
	AliPay        []pay.AliPay          `mapstructure:"alipay"       yaml:"alipay"       json:"alipay"`
	WechatPay     []pay.WechatPay       `mapstructure:"wechatpay"    yaml:"wechatpay"    json:"wechatpay"`
	AliyunSms     []sms.AliyunSms       `mapstructure:"aliyunsms"    yaml:"aliyunsms"    json:"aliyunsms"`
	AliyunSts     []sts.AliyunSts       `mapstructure:"aliyunsts"    yaml:"aliyunsts"    json:"aliyunsts"`
	Youzan        []youzan.YouZan       `mapstructure:"youzan"       yaml:"youzan"       json:"youzan"`
	ZeroServer    []zero.RpcServer      `mapstructure:"zeroserver"   yaml:"zeroserver"   json:"zeroserver"`
	ZeroClient    []zero.RpcClient      `mapstructure:"zeroclient"   yaml:"zeroclient"   json:"zeroclient"`
	ZeroRestful   []zero.Restful        `mapstructure:"zerorestful"  yaml:"zerorestful"  json:"zerorestful"`
	Kafka         []elk.Kafka           `mapstructure:"kafka"        yaml:"kafka"        json:"kafka"`
	Elasticsearch []elk.Elasticsearch   `mapstructure:"elasticsearch"  yaml:"elasticsearch"  json:"elasticsearch"`
	Jaeger        []register.Jaeger     `mapstructure:"jaeger"       yaml:"jaeger"       json:"jaeger"`
	Pprof        []register.PProf       `mapstructure:"pprof"       yaml:"pprof"       json:"pprof"`
	Viper         *viper.Viper          `mapstructure:"-"            yaml:"-"            json:"-"`
	ExternalVipers map[string]*viper.Viper `mapstructure:"-" yaml:"-" json:"-"`
	DynamicConfigs map[string]interface{}  `mapstructure:",remain" yaml:",inline" json:",inline"`
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
	AliyunOss     oss.AliyunOss       `mapstructure:"aliyunoss"    yaml:"aliyunoss"    json:"aliyunoss"`
	S3            oss.S3              `mapstructure:"s3"           yaml:"s3"           json:"s3"`
	Mqtt          queue.Mqtt          `mapstructure:"mqtt"         yaml:"mqtt"         json:"mqtt"`
	Zap           zap.Zap             `mapstructure:"zap"          yaml:"zap"          json:"zap"`
	AliPay        pay.AliPay          `mapstructure:"alipay"       yaml:"alipay"       json:"alipay"`
	WechatPay     pay.WechatPay       `mapstructure:"wechatpay"    yaml:"wechatpay"    json:"wechatpay"`
	AliyunSms     sms.AliyunSms       `mapstructure:"aliyunsms"    yaml:"aliyunsms"    json:"aliyunsms"`
	AliyunSts     sts.AliyunSts       `mapstructure:"aliyunsts"    yaml:"aliyunsts"    json:"aliyunsts"`
	Youzan        youzan.YouZan       `mapstructure:"youzan"       yaml:"youzan"       json:"youzan"`
	ZeroServer    zero.RpcServer      `mapstructure:"zeroserver"   yaml:"zeroserver"   json:"zeroserver"`
	ZeroClient    zero.RpcClient      `mapstructure:"zeroclient"   yaml:"zeroclient"   json:"zeroclient"`
	ZeroRestful   zero.Restful        `mapstructure:"zerorestful"  yaml:"zerorestful"  json:"zerorestful"`
	Kafka         elk.Kafka           `mapstructure:"kafka"        yaml:"kafka"        json:"kafka"`
	Elasticsearch elk.Elasticsearch   `mapstructure:"elasticsearch"  yaml:"elasticsearch"  json:"elasticsearch"`
	Jaeger        register.Jaeger     `mapstructure:"jaeger"       yaml:"jaeger"       json:"jaeger"`
	Pprof         register.PProf      `mapstructure:"pprof"        yaml:"pprof"        json:"pprof"`
	Viper         *viper.Viper        `mapstructure:"-"            yaml:"-"            json:"-"`
	
	// 新增：存储无关系的Viper实例和动态配置
	ExternalVipers map[string]*viper.Viper `mapstructure:"-" yaml:"-" json:"-"`
	DynamicConfigs map[string]interface{}  `mapstructure:",remain" yaml:",inline" json:",inline"`
}

// GetSingleConfigByModuleName 根据提供的模块名称从 MultiConfig 中获取对应的 SingleConfig
func GetSingleConfigByModuleName(multiConfig MultiConfig, moduleName string) (*SingleConfig, error) {
	var singleConfig SingleConfig
	var found bool

	// 获取 MultiConfig 的反射值
	val := reflect.ValueOf(multiConfig)

	// 遍历 MultiConfig 的所有字段
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// 只处理切片类型的字段
		if field.Kind() == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				item := field.Index(j)

				// 获取每个切片元素的 ModuleName 字段
				moduleNameField := item.FieldByName("ModuleName")
				if moduleNameField.IsValid() && moduleNameField.String() == moduleName {
					// 获取类型名称
					typeName := field.Type().Elem().Name()

					// 获取 SingleConfig 的反射值
					singleConfigVal := reflect.ValueOf(&singleConfig).Elem()

					// 根据类型名称找到对应的字段并赋值
					configField := singleConfigVal.FieldByName(typeName)
					if configField.IsValid() && configField.CanSet() {
						configField.Set(item)
						found = true
						break
					}
				}
			}
		}
	}

	if !found {
		return nil, fmt.Errorf("未找到模块名称: %s", moduleName)
	}

	return &singleConfig, nil
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

// InitializeExternalVipers 初始化外部Viper实例映射
func (mc *MultiConfig) InitializeExternalVipers() {
	if mc.ExternalVipers == nil {
		mc.ExternalVipers = make(map[string]*viper.Viper)
	}
	if mc.DynamicConfigs == nil {
		mc.DynamicConfigs = make(map[string]interface{})
	}
}

// InitializeExternalVipers 初始化外部Viper实例映射
func (sc *SingleConfig) InitializeExternalVipers() {
	if sc.ExternalVipers == nil {
		sc.ExternalVipers = make(map[string]*viper.Viper)
	}
	if sc.DynamicConfigs == nil {
		sc.DynamicConfigs = make(map[string]interface{})
	}
}

// AddExternalViper 添加外部Viper实例
func (mc *MultiConfig) AddExternalViper(key string, v *viper.Viper) {
	mc.InitializeExternalVipers()
	mc.ExternalVipers[key] = v
}

// AddExternalViper 添加外部Viper实例
func (sc *SingleConfig) AddExternalViper(key string, v *viper.Viper) {
	sc.InitializeExternalVipers()
	sc.ExternalVipers[key] = v
}

// GetExternalViper 获取外部Viper实例
func (mc *MultiConfig) GetExternalViper(key string) (*viper.Viper, bool) {
	if mc.ExternalVipers == nil {
		return nil, false
	}
	v, exists := mc.ExternalVipers[key]
	return v, exists
}

// GetExternalViper 获取外部Viper实例
func (sc *SingleConfig) GetExternalViper(key string) (*viper.Viper, bool) {
	if sc.ExternalVipers == nil {
		return nil, false
	}
	v, exists := sc.ExternalVipers[key]
	return v, exists
}

// SetDynamicConfig 设置动态配置
func (mc *MultiConfig) SetDynamicConfig(key string, value interface{}) {
	mc.InitializeExternalVipers()
	mc.DynamicConfigs[key] = value
}

// SetDynamicConfig 设置动态配置
func (sc *SingleConfig) SetDynamicConfig(key string, value interface{}) {
	sc.InitializeExternalVipers()
	sc.DynamicConfigs[key] = value
}

// GetDynamicConfig 获取动态配置
func (mc *MultiConfig) GetDynamicConfig(key string) (interface{}, bool) {
	if mc.DynamicConfigs == nil {
		return nil, false
	}
	value, exists := mc.DynamicConfigs[key]
	return value, exists
}

// GetDynamicConfig 获取动态配置
func (sc *SingleConfig) GetDynamicConfig(key string) (interface{}, bool) {
	if sc.DynamicConfigs == nil {
		return nil, false
	}
	value, exists := sc.DynamicConfigs[key]
	return value, exists
}

// GetAllExternalViperKeys 获取所有外部Viper键
func (mc *MultiConfig) GetAllExternalViperKeys() []string {
	if mc.ExternalVipers == nil {
		return []string{}
	}
	keys := make([]string, 0, len(mc.ExternalVipers))
	for key := range mc.ExternalVipers {
		keys = append(keys, key)
	}
	return keys
}

// GetAllExternalViperKeys 获取所有外部Viper键
func (sc *SingleConfig) GetAllExternalViperKeys() []string {
	if sc.ExternalVipers == nil {
		return []string{}
	}
	keys := make([]string, 0, len(sc.ExternalVipers))
	for key := range sc.ExternalVipers {
		keys = append(keys, key)
	}
	return keys
}

// GetAllDynamicConfigKeys 获取所有动态配置键
func (mc *MultiConfig) GetAllDynamicConfigKeys() []string {
	if mc.DynamicConfigs == nil {
		return []string{}
	}
	keys := make([]string, 0, len(mc.DynamicConfigs))
	for key := range mc.DynamicConfigs {
		keys = append(keys, key)
	}
	return keys
}

// GetAllDynamicConfigKeys 获取所有动态配置键
func (sc *SingleConfig) GetAllDynamicConfigKeys() []string {
	if sc.DynamicConfigs == nil {
		return []string{}
	}
	keys := make([]string, 0, len(sc.DynamicConfigs))
	for key := range sc.DynamicConfigs {
		keys = append(keys, key)
	}
	return keys
}

// RemoveExternalViper 移除外部Viper实例
func (mc *MultiConfig) RemoveExternalViper(key string) bool {
	if mc.ExternalVipers == nil {
		return false
	}
	_, exists := mc.ExternalVipers[key]
	if exists {
		delete(mc.ExternalVipers, key)
	}
	return exists
}

// RemoveExternalViper 移除外部Viper实例
func (sc *SingleConfig) RemoveExternalViper(key string) bool {
	if sc.ExternalVipers == nil {
		return false
	}
	_, exists := sc.ExternalVipers[key]
	if exists {
		delete(sc.ExternalVipers, key)
	}
	return exists
}

// RemoveDynamicConfig 移除动态配置
func (mc *MultiConfig) RemoveDynamicConfig(key string) bool {
	if mc.DynamicConfigs == nil {
		return false
	}
	_, exists := mc.DynamicConfigs[key]
	if exists {
		delete(mc.DynamicConfigs, key)
	}
	return exists
}

// RemoveDynamicConfig 移除动态配置
func (sc *SingleConfig) RemoveDynamicConfig(key string) bool {
	if sc.DynamicConfigs == nil {
		return false
	}
	_, exists := sc.DynamicConfigs[key]
	if exists {
		delete(sc.DynamicConfigs, key)
	}
	return exists
}

// UnmarshalFromExternalViper 从外部Viper实例解析配置到指定结构体
func (mc *MultiConfig) UnmarshalFromExternalViper(key string, target interface{}) error {
	viper, exists := mc.GetExternalViper(key)
	if !exists {
		return fmt.Errorf("external viper with key '%s' not found", key)
	}
	return viper.Unmarshal(target)
}

// UnmarshalFromExternalViper 从外部Viper实例解析配置到指定结构体
func (sc *SingleConfig) UnmarshalFromExternalViper(key string, target interface{}) error {
	viper, exists := sc.GetExternalViper(key)
	if !exists {
		return fmt.Errorf("external viper with key '%s' not found", key)
	}
	return viper.Unmarshal(target)
}

// UnmarshalSubFromExternalViper 从外部Viper实例的子配置解析到指定结构体
func (mc *MultiConfig) UnmarshalSubFromExternalViper(viperKey, subKey string, target interface{}) error {
	viper, exists := mc.GetExternalViper(viperKey)
	if !exists {
		return fmt.Errorf("external viper with key '%s' not found", viperKey)
	}
	sub := viper.Sub(subKey)
	if sub == nil {
		return fmt.Errorf("sub config '%s' not found in viper '%s'", subKey, viperKey)
	}
	return sub.Unmarshal(target)
}

// UnmarshalSubFromExternalViper 从外部Viper实例的子配置解析到指定结构体
func (sc *SingleConfig) UnmarshalSubFromExternalViper(viperKey, subKey string, target interface{}) error {
	viper, exists := sc.GetExternalViper(viperKey)
	if !exists {
		return fmt.Errorf("external viper with key '%s' not found", viperKey)
	}
	sub := viper.Sub(subKey)
	if sub == nil {
		return fmt.Errorf("sub config '%s' not found in viper '%s'", subKey, viperKey)
	}
	return sub.Unmarshal(target)
}
