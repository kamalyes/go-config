/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 02:51:01
 * @FilePath: \go-config\pkg\register\jaeger.go
 * @Description:
 *
 * Copyright (j) 2024 by kamalyes, All Rights Reserved.
 */

package register

import (
	"github.com/kamalyes/go-config/internal"
)

// Jaeger 结构体用于配置 Jaeger 注册中心相关参数
type Jaeger struct {
	Type               string `mapstructure:"type"                     yaml:"type"                      json:"type" validate:"required"`                  // Jaeger 类型
	Param              int    `mapstructure:"param"                    yaml:"param"                     json:"param" validate:"required"`                 // 参数
	LogSpans           bool   `mapstructure:"log-spans"                yaml:"log-spans"                 json:"log_spans"`                                 // 是否记录跨度
	Endpoint         string `mapstructure:"endpoint"                yaml:"endpoint"          json:"endpoint"          validate:"required,url"` // 本地代理地址和端口
	Service            string `mapstructure:"service"                  yaml:"service"                   json:"service" validate:"required"`               // 服务名称
	ModuleName         string `mapstructure:"modulename"               yaml:"modulename"                json:"module_name"`                               // 模块名称
}

// NewJaeger 创建一个新的 Jaeger 实例
func NewJaeger(opt *Jaeger) *Jaeger {
	var jaegerInstance *Jaeger

	internal.LockFunc(func() {
		jaegerInstance = opt
	})
	return jaegerInstance
}

// Clone 返回 Jaeger 配置的副本
func (j *Jaeger) Clone() internal.Configurable {
	return &Jaeger{
		ModuleName:         j.ModuleName,
		Type:               j.Type,
		Param:              j.Param,
		LogSpans:           j.LogSpans,
		Endpoint: j.Endpoint,
		Service:            j.Service,
	}
}

// Get 返回 Jaeger 配置的所有字段
func (j *Jaeger) Get() interface{} {
	return j
}

// Set 更新 Jaeger 配置的字段
func (j *Jaeger) Set(data interface{}) {
	if configData, ok := data.(*Jaeger); ok {
		j.ModuleName = configData.ModuleName
		j.Type = configData.Type
		j.Param = configData.Param
		j.LogSpans = configData.LogSpans
		j.Endpoint = configData.Endpoint
		j.Service = configData.Service
	}
}

// Validate 检查 Jaeger 配置的有效性
func (j *Jaeger) Validate() error {
	return internal.ValidateStruct(j)
}

// DefaultJaeger 返回默认Jaeger配置
func DefaultJaeger() Jaeger {
	return Jaeger{
		ModuleName:         "jaeger",
		Type:               "const",
		Param:              1,
		LogSpans:           true,
		Endpoint: "127.0.0.1:6831",
		Service:            "go-config-service",
	}
}

// DefaultJaegerConfig 返回默认Jaeger配置的指针，支持链式调用
func DefaultJaegerConfig() *Jaeger {
	config := DefaultJaeger()
	return &config
}

// WithModuleName 设置模块名称
func (j *Jaeger) WithModuleName(moduleName string) *Jaeger {
	j.ModuleName = moduleName
	return j
}

// WithType 设置Jaeger类型
func (j *Jaeger) WithType(jaegerType string) *Jaeger {
	j.Type = jaegerType
	return j
}

// WithParam 设置参数
func (j *Jaeger) WithParam(param int) *Jaeger {
	j.Param = param
	return j
}

// WithLogSpans 设置是否记录跨度
func (j *Jaeger) WithLogSpans(logSpans bool) *Jaeger {
	j.LogSpans = logSpans
	return j
}

// WithEndpoint 设置本地代理地址和端口
func (j *Jaeger) WithEndpoint(endpoint string) *Jaeger {
	j.Endpoint = endpoint
	return j
}

// WithService 设置服务名称
func (j *Jaeger) WithService(service string) *Jaeger {
	j.Service = service
	return j
}
