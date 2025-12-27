/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-04 13:20:28
 * @FilePath: \go-config\pkg\logging\logging.go
 * @Description: 日志中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package logging

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-toolbox/pkg/syncx"
)

// Logging 日志中间件配置
type Logging struct {
	ModuleName           string   `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                                 // 模块名称
	Enabled              bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                            // 是否启用日志
	Level                string   `mapstructure:"level" yaml:"level" json:"level"`                                                  // 日志级别
	Format               string   `mapstructure:"format" yaml:"format" json:"format"`                                               // 日志格式 (json, text)
	Output               string   `mapstructure:"output" yaml:"output" json:"output"`                                               // 输出目标 (stdout, file)
	FilePath             string   `mapstructure:"file-path" yaml:"file-path" json:"filePath"`                                       // 日志文件路径
	MaxSize              int      `mapstructure:"max-size" yaml:"max-size" json:"maxSize"`                                          // 最大文件大小(MB)
	MaxBackups           int      `mapstructure:"max-backups" yaml:"max-backups" json:"maxBackups"`                                 // 最大备份文件数
	MaxAge               int      `mapstructure:"max-age" yaml:"max-age" json:"maxAge"`                                             // 最大保存天数
	Compress             bool     `mapstructure:"compress" yaml:"compress" json:"compress"`                                         // 是否压缩
	SkipPaths            []string `mapstructure:"skip-paths" yaml:"skip-paths" json:"skipPaths"`                                    // 跳过的路径
	EnableRequest        bool     `mapstructure:"enable-request" yaml:"enable-request" json:"enableRequest"`                        // 是否记录请求
	EnableResponse       bool     `mapstructure:"enable-response" yaml:"enable-response" json:"enableResponse"`                     // 是否记录响应
	MaxBodySize          int      `mapstructure:"max-body-size" yaml:"max-body-size" json:"maxBodySize"`                            // 最大日志体大小(字节)
	SensitiveMask        string   `mapstructure:"sensitive-mask" yaml:"sensitive-mask" json:"sensitiveMask"`                        // 敏感数据掩码
	SensitiveKeys        []string `mapstructure:"sensitive-keys" yaml:"sensitive-keys" json:"sensitiveKeys"`                        // 敏感字段关键词
	SlowHTTPThreshold    int64    `mapstructure:"slow-http-threshold" yaml:"slow-http-threshold" json:"slowHttpThreshold"`          // HTTP慢请求阈值(毫秒)
	SlowGRPCThreshold    int64    `mapstructure:"slow-grpc-threshold" yaml:"slow-grpc-threshold" json:"slowGrpcThreshold"`          // GRPC慢请求阈值(毫秒)
	SlowStreamThreshold  int64    `mapstructure:"slow-stream-threshold" yaml:"slow-stream-threshold" json:"slowStreamThreshold"`    // 流式请求慢请求阈值(毫秒)
	LoggableContentTypes []string `mapstructure:"loggable-content-types" yaml:"loggable-content-types" json:"loggableContentTypes"` // 可记录的 Content-Type
}

// Default 创建默认日志配置
func Default() *Logging {
	return &Logging{
		ModuleName:     "logging",
		Enabled:        true,
		Level:          "info",
		Format:         "json",
		Output:         "stdout",
		FilePath:       "/var/log/app.log",
		MaxSize:        100,
		MaxBackups:     3,
		MaxAge:         28,
		Compress:       true,
		SkipPaths:      []string{"/health", "/metrics", "/favicon.ico", "/ping", "/readiness", "/liveness"},
		EnableRequest:  true,
		EnableResponse: false,
		MaxBodySize:    2048,
		SensitiveMask:  "***REDACTED***",
		SensitiveKeys: []string{
			"password", "passwd", "token", "access_token", "refresh_token",
			"secret", "authorization", "api_key", "apikey",
			"mobile", "phone", "id_card", "credit_card",
		},
		SlowHTTPThreshold:    1000,
		SlowGRPCThreshold:    1000,
		SlowStreamThreshold:  5000,
		LoggableContentTypes: []string{"application/json", "application/xml", "application/x-www-form-urlencoded", "text/"},
	}
}

// Get 返回配置接口
func (l *Logging) Get() interface{} {
	return l
}

// Set 设置配置数据
func (l *Logging) Set(data interface{}) {
	if cfg, ok := data.(*Logging); ok {
		*l = *cfg
	}
}

// Clone 返回配置的副本
func (l *Logging) Clone() internal.Configurable {
	var cloned Logging
	if err := syncx.DeepCopy(&cloned, l); err != nil {
		// 如果深拷贝失败，返回空配置
		return &Logging{}
	}
	return &cloned
}

// ========== Logging 链式调用方法 ==========

// WithModuleName 设置模块名称
func (l *Logging) WithModuleName(moduleName string) *Logging {
	l.ModuleName = moduleName
	return l
}

// WithEnabled 设置是否启用日志
func (l *Logging) WithEnabled(enabled bool) *Logging {
	l.Enabled = enabled
	return l
}

// EnableLogging 启用日志
func (l *Logging) EnableLogging() *Logging {
	l.Enabled = true
	return l
}

// WithLevel 设置日志级别
func (l *Logging) WithLevel(level string) *Logging {
	l.Level = level
	return l
}

// WithFormat 设置日志格式
func (l *Logging) WithFormat(format string) *Logging {
	l.Format = format
	return l
}

// WithOutput 设置输出目标
func (l *Logging) WithOutput(output string) *Logging {
	l.Output = output
	return l
}

// WithFilePath 设置日志文件路径
func (l *Logging) WithFilePath(filePath string) *Logging {
	l.FilePath = filePath
	return l
}

// WithMaxSize 设置最大文件大小
func (l *Logging) WithMaxSize(maxSize int) *Logging {
	l.MaxSize = maxSize
	return l
}

// WithMaxBackups 设置最大备份文件数
func (l *Logging) WithMaxBackups(maxBackups int) *Logging {
	l.MaxBackups = maxBackups
	return l
}

// WithMaxAge 设置最大保存天数
func (l *Logging) WithMaxAge(maxAge int) *Logging {
	l.MaxAge = maxAge
	return l
}

// WithCompress 设置是否压缩
func (l *Logging) WithCompress(compress bool) *Logging {
	l.Compress = compress
	return l
}

// WithSkipPaths 设置跳过的路径
func (l *Logging) WithSkipPaths(skipPaths []string) *Logging {
	l.SkipPaths = skipPaths
	return l
}

// AddSkipPath 添加跳过的路径
func (l *Logging) AddSkipPath(path string) *Logging {
	l.SkipPaths = append(l.SkipPaths, path)
	return l
}

// WithEnableRequest 设置是否记录请求
func (l *Logging) WithEnableRequest(enableRequest bool) *Logging {
	l.EnableRequest = enableRequest
	return l
}

// WithEnableResponse 设置是否记录响应
func (l *Logging) WithEnableResponse(enableResponse bool) *Logging {
	l.EnableResponse = enableResponse
	return l
}

// EnableRequestResponse 同时启用请求和响应记录
func (l *Logging) EnableRequestResponse() *Logging {
	l.EnableRequest = true
	l.EnableResponse = true
	return l
}

// Validate 验证配置
func (l *Logging) Validate() error {
	return internal.ValidateStruct(l)
}

// Enable 启用日志中间件
func (l *Logging) Enable() *Logging {
	return l.EnableLogging()
}

// Disable 禁用日志中间件
func (l *Logging) Disable() *Logging {
	l.Enabled = false
	return l
}

// IsEnabled 检查是否启用
func (l *Logging) IsEnabled() bool {
	return l.Enabled
}
