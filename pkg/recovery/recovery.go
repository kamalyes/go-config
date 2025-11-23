/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-24 14:43:53
 * @FilePath: \engine-im-service\go-config\pkg\recovery\recovery.go
 * @Description: 恢复中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package recovery

import (
	"net/http"

	"github.com/kamalyes/go-config/internal"
)

// Recovery 恢复中间件配置
type Recovery struct {
	ModuleName      string                                                `mapstructure:"module_name" yaml:"module_name" json:"module_name"`       // 模块名称
	Enabled         bool                                                  `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                   // 是否启用恢复
	EnableStack     bool                                                  `mapstructure:"enable_stack" yaml:"enable_stack" json:"enable_stack"`    // 是否启用堆栈跟踪
	StackSize       int                                                   `mapstructure:"stack_size" yaml:"stack_size" json:"stack_size"`          // 堆栈大小
	EnableDebug     bool                                                  `mapstructure:"enable_debug" yaml:"enable_debug" json:"enable_debug"`    // 是否启用调试模式
	ErrorMessage    string                                                `mapstructure:"error_message" yaml:"error_message" json:"error_message"` // 默认错误消息
	LogLevel        string                                                `mapstructure:"log_level" yaml:"log_level" json:"log_level"`             // 日志级别
	EnableNotify    bool                                                  `mapstructure:"enable_notify" yaml:"enable_notify" json:"enable_notify"` // 是否启用通知
	RecoveryHandler func(http.ResponseWriter, *http.Request, interface{}) `mapstructure:"-" yaml:"-" json:"-"`                                     // 自定义恢复处理器
	PrintStack      bool                                                  `mapstructure:"print_stack" yaml:"print_stack" json:"print_stack"`       // 是否打印堆栈(兼容旧版)
}

// Default 创建默认恢复配置
func Default() *Recovery {
	return &Recovery{
		ModuleName:      "recovery",
		Enabled:         true,
		EnableStack:     true,
		StackSize:       4096,
		EnableDebug:     false,
		ErrorMessage:    "服务器内部错误",
		LogLevel:        "error",
		EnableNotify:    false,
		RecoveryHandler: nil,
		PrintStack:      true, // 兼容旧版
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
		ModuleName:      r.ModuleName,
		Enabled:         r.Enabled,
		EnableStack:     r.EnableStack,
		StackSize:       r.StackSize,
		EnableDebug:     r.EnableDebug,
		ErrorMessage:    r.ErrorMessage,
		LogLevel:        r.LogLevel,
		EnableNotify:    r.EnableNotify,
		RecoveryHandler: r.RecoveryHandler,
		PrintStack:      r.PrintStack,
	}
}

// Validate 验证配置
func (r *Recovery) Validate() error {
	return internal.ValidateStruct(r)
}

// WithPrintStack 设置是否打印堆栈
func (r *Recovery) WithPrintStack(printStack bool) *Recovery {
	r.PrintStack = printStack
	r.EnableStack = printStack // 同步设置新字段
	return r
}

// WithEnableStack 设置是否启用堆栈跟踪
func (r *Recovery) WithEnableStack(enableStack bool) *Recovery {
	r.EnableStack = enableStack
	return r
}

// WithStackSize 设置堆栈大小
func (r *Recovery) WithStackSize(stackSize int) *Recovery {
	r.StackSize = stackSize
	return r
}

// WithEnableDebug 设置是否启用调试模式
func (r *Recovery) WithEnableDebug(enableDebug bool) *Recovery {
	r.EnableDebug = enableDebug
	return r
}

// WithErrorMessage 设置默认错误消息
func (r *Recovery) WithErrorMessage(errorMessage string) *Recovery {
	r.ErrorMessage = errorMessage
	return r
}

// WithRecoveryHandler 设置自定义恢复处理器
func (r *Recovery) WithRecoveryHandler(handler func(http.ResponseWriter, *http.Request, interface{})) *Recovery {
	r.RecoveryHandler = handler
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
