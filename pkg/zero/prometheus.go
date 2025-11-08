/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-07 15:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 18:06:52
 * @FilePath: \go-config\pkg\zero\prometheus.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

import "github.com/kamalyes/go-config/internal"

// Prometheus 表示 Prometheus 配置
type Prometheus struct {
	Enabled bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"` // 是否启用 Prometheus
}

// NewPrometheus 创建一个新的 Prometheus 实例
func NewPrometheus(opt *Prometheus) *Prometheus {
	var etcdInstance *Prometheus

	internal.LockFunc(func() {
		etcdInstance = opt
	})
	return etcdInstance
}

// Clone 返回 RpcServer 配置的副本
func (e *Prometheus) Clone() internal.Configurable {
	return &Prometheus{
		Enabled: e.Enabled,
	}
}

// Get 返回 RpcServer 配置的所有字段
func (e *Prometheus) Get() interface{} {
	return e
}

// Set 更新 RpcServer 配置的字段
func (e *Prometheus) Set(data interface{}) {
	if configData, ok := data.(*Prometheus); ok {
		e.Enabled = configData.Enabled
	}
}

// Validate 验证 Prometheus 配置的有效性
func (e *Prometheus) Validate() error {
	return internal.ValidateStruct(e)
}
