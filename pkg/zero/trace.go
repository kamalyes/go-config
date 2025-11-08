/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-07 15:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 18:07:55
 * @FilePath: \go-config\pkg\zero\trace.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

import "github.com/kamalyes/go-config/internal"

// Telemetry 表示追踪配置
type Telemetry struct {
	Enabled bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"` // 是否启用追踪
}

// NewTelemetry 创建一个新的 Telemetry 实例
func NewTelemetry(opt *Telemetry) *Telemetry {
	var etcdInstance *Telemetry

	internal.LockFunc(func() {
		etcdInstance = opt
	})
	return etcdInstance
}

// Clone 返回 RpcServer 配置的副本
func (e *Telemetry) Clone() internal.Configurable {
	return &Telemetry{
		Enabled: e.Enabled,
	}
}

// Get 返回 RpcServer 配置的所有字段
func (e *Telemetry) Get() interface{} {
	return e
}

// Set 更新 RpcServer 配置的字段
func (e *Telemetry) Set(data interface{}) {
	if configData, ok := data.(*Telemetry); ok {
		e.Enabled = configData.Enabled
	}
}

// Validate 验证 Telemetry 配置的有效性
func (e *Telemetry) Validate() error {
	return internal.ValidateStruct(e)
}

// DefaultTelemetry 返回默认的 Telemetry 指针，支持链式调用
func DefaultTelemetry() *Telemetry {
	config := DefaultTelemetryConfig()
	return &config
}

// DefaultTelemetryConfig 返回默认的 Telemetry 值
func DefaultTelemetryConfig() Telemetry {
	return Telemetry{
		Enabled: false,
	}
}

// WithEnabled 设置是否启用追踪
func (e *Telemetry) WithEnabled(enabled bool) *Telemetry {
	e.Enabled = enabled
	return e
}
