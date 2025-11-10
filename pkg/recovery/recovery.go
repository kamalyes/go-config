/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 18:00:00
 * @FilePath: \go-config\pkg\recovery\recovery.go
 * @Description: 恢复中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package recovery

import "github.com/kamalyes/go-config/internal"

// Recovery 恢复中间件配置
type Recovery struct {
	ModuleName   string `mapstructure:"module_name" yaml:"module-name" json:"module_name"`       // 模块名称
	Enabled      bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                   // 是否启用恢复
	PrintStack   bool   `mapstructure:"print_stack" yaml:"print-stack" json:"print_stack"`       // 是否打印堆栈
	LogLevel     string `mapstructure:"log_level" yaml:"log-level" json:"log_level"`             // 日志级别
	EnableNotify bool   `mapstructure:"enable_notify" yaml:"enable-notify" json:"enable_notify"` // 是否启用通知
}

// Default 创建默认恢复配置
func Default() *Recovery {
	return &Recovery{
		ModuleName:   "recovery",
		Enabled:      true,
		PrintStack:   true,
		LogLevel:     "error",
		EnableNotify: false,
	}
}

// Get 返回配置接口
func (r *Recovery) Get() interface{} {
	return r
}

// Set 设置配置数据
func (r *Recovery) Set(data interface{}) {
	if cfg, ok := data.(*Recovery); ok {
		*r = *cfg
	}
}

// Clone 返回配置的副本
func (r *Recovery) Clone() internal.Configurable {
	return &Recovery{
		ModuleName:   r.ModuleName,
		Enabled:      r.Enabled,
		PrintStack:   r.PrintStack,
		LogLevel:     r.LogLevel,
		EnableNotify: r.EnableNotify,
	}
}

// Validate 验证配置
func (r *Recovery) Validate() error {
	return internal.ValidateStruct(r)
}

// WithPrintStack 设置是否打印堆栈
func (r *Recovery) WithPrintStack(printStack bool) *Recovery {
	r.PrintStack = printStack
	return r
}

// WithLogLevel 设置日志级别
func (r *Recovery) WithLogLevel(logLevel string) *Recovery {
	r.LogLevel = logLevel
	return r
}

// WithNotify 设置是否启用通知
func (r *Recovery) WithNotify(notify bool) *Recovery {
	r.EnableNotify = notify
	return r
}

// Enable 启用恢复中间件
func (r *Recovery) Enable() *Recovery {
	r.Enabled = true
	return r
}

// Disable 禁用恢复中间件
func (r *Recovery) Disable() *Recovery {
	r.Enabled = false
	return r
}

// IsEnabled 检查是否启用
func (r *Recovery) IsEnabled() bool {
	return r.Enabled
}
