/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 23:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 23:30:00
 * @FilePath: \go-config\pkg\timeout\timeout.go
 * @Description: 超时中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package timeout

import (
	"github.com/kamalyes/go-config/internal"
	"time"
)

// Timeout 超时中间件配置
type Timeout struct {
	ModuleName string        `mapstructure:"module-name" yaml:"module-name" json:"moduleName"` // 模块名称
	Enabled    bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`            // 是否启用超时
	Duration   time.Duration `mapstructure:"duration" yaml:"duration" json:"duration"`         // 超时时长
	Message    string        `mapstructure:"message" yaml:"message" json:"message"`            // 超时消息
}

// Default 创建默认超时配置
func Default() *Timeout {
	return &Timeout{
		ModuleName: "timeout",
		Enabled:    false,
		Duration:   30 * time.Second,
		Message:    "请求超时",
	}
}

// WithModuleName 设置模块名称
func (t *Timeout) WithModuleName(moduleName string) *Timeout {
	t.ModuleName = moduleName
	return t
}

// WithEnabled 设置是否启用
func (t *Timeout) WithEnabled(enabled bool) *Timeout {
	t.Enabled = enabled
	return t
}

// WithDuration 设置超时时长
func (t *Timeout) WithDuration(duration time.Duration) *Timeout {
	t.Duration = duration
	return t
}

// WithMessage 设置超时消息
func (t *Timeout) WithMessage(message string) *Timeout {
	t.Message = message
	return t
}

// Get 返回配置接口
func (t *Timeout) Get() interface{} {
	return t
}

// Set 设置配置数据
func (t *Timeout) Set(data interface{}) {
	if cfg, ok := data.(*Timeout); ok {
		*t = *cfg
	}
}

// Clone 返回配置的副本
func (t *Timeout) Clone() internal.Configurable {
	return &Timeout{
		ModuleName: t.ModuleName,
		Enabled:    t.Enabled,
		Duration:   t.Duration,
		Message:    t.Message,
	}
}

// Validate 验证配置
func (t *Timeout) Validate() error {
	return internal.ValidateStruct(t)
}

// Enable 启用超时
func (t *Timeout) Enable() *Timeout {
	t.Enabled = true
	return t
}

// Disable 禁用超时
func (t *Timeout) Disable() *Timeout {
	t.Enabled = false
	return t
}

// IsEnabled 检查是否启用
func (t *Timeout) IsEnabled() bool {
	return t.Enabled
}
