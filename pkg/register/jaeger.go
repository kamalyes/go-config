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
	LocalAgentHostPort string `mapstructure:"local-agent-host-port"    yaml:"local-agent-host-port"     json:"local_agent_host_port" validate:"required"` // 本地代理地址和端口
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
		LocalAgentHostPort: j.LocalAgentHostPort,
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
		j.LocalAgentHostPort = configData.LocalAgentHostPort
		j.Service = configData.Service
	}
}

// Validate 检查 Jaeger 配置的有效性
func (j *Jaeger) Validate() error {
	return internal.ValidateStruct(j)
}
