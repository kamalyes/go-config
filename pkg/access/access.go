/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 15:00:00
 * @FilePath: \go-config\pkg\access\access.go
 * @Description: 访问记录中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package access

import "github.com/kamalyes/go-config/internal"

// Access 访问记录中间件配置
type Access struct {
	ModuleName      string   `mapstructure:"module_name" yaml:"module-name" json:"module_name"`                      // 模块名称
	Enabled         bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                 // 是否启用访问记录
	ServiceName     string   `mapstructure:"service_name" yaml:"service-name" json:"service_name"`                   // 服务名称
	RetentionDays   int      `mapstructure:"retention_days" yaml:"retention-days" json:"retention_days"`             // 保留天数
	IncludeBody     bool     `mapstructure:"include_body" yaml:"include-body" json:"include_body"`                   // 是否记录请求体
	IncludeResponse bool     `mapstructure:"include_response" yaml:"include-response" json:"include_response"`       // 是否记录响应体
	IncludeHeaders  []string `mapstructure:"include_headers" yaml:"include-headers" json:"include_headers"`          // 要记录的头部
	ExcludePaths    []string `mapstructure:"exclude_paths" yaml:"exclude-paths" json:"exclude_paths"`                // 排除的路径
	MaxBodySize     int64    `mapstructure:"max_body_size" yaml:"max-body-size" json:"max_body_size"`                // 最大请求体大小
	MaxResponseSize int64    `mapstructure:"max_response_size" yaml:"max-response-size" json:"max_response_size"`    // 最大响应体大小
}

// Default 创建默认访问记录配置
func Default() *Access {
	return &Access{
		ModuleName:      "access",
		Enabled:         true,
		ServiceName:     "rpc-gateway",
		RetentionDays:   60,
		IncludeBody:     true,
		IncludeResponse: true,
		IncludeHeaders:  []string{"User-Agent", "X-Request-ID", "X-Trace-ID", "Authorization", "Content-Type"},
		ExcludePaths:    []string{"/health", "/metrics", "/debug"},
		MaxBodySize:     1024 * 1024,     // 1MB
		MaxResponseSize: 1024 * 1024 * 5, // 5MB
	}
}

// Get 返回配置接口
func (a *Access) Get() interface{} {
	return a
}

// Set 设置配置数据
func (a *Access) Set(data interface{}) {
	if cfg, ok := data.(*Access); ok {
		*a = *cfg
	}
}

// Clone 返回配置的副本
func (a *Access) Clone() internal.Configurable {
	clone := &Access{
		ModuleName:      a.ModuleName,
		Enabled:         a.Enabled,
		ServiceName:     a.ServiceName,
		RetentionDays:   a.RetentionDays,
		IncludeBody:     a.IncludeBody,
		IncludeResponse: a.IncludeResponse,
		MaxBodySize:     a.MaxBodySize,
		MaxResponseSize: a.MaxResponseSize,
	}
	clone.IncludeHeaders = append([]string(nil), a.IncludeHeaders...)
	clone.ExcludePaths = append([]string(nil), a.ExcludePaths...)
	return clone
}

// Validate 验证配置
func (a *Access) Validate() error {
	return internal.ValidateStruct(a)
}

// WithServiceName 设置服务名称
func (a *Access) WithServiceName(serviceName string) *Access {
	a.ServiceName = serviceName
	return a
}

// WithRetentionDays 设置保留天数
func (a *Access) WithRetentionDays(retentionDays int) *Access {
	a.RetentionDays = retentionDays
	return a
}

// WithIncludeBody 设置是否包含请求体
func (a *Access) WithIncludeBody(includeBody bool) *Access {
	a.IncludeBody = includeBody
	return a
}

// WithIncludeResponse 设置是否包含响应体
func (a *Access) WithIncludeResponse(includeResponse bool) *Access {
	a.IncludeResponse = includeResponse
	return a
}

// WithIncludeHeaders 设置要包含的头部
func (a *Access) WithIncludeHeaders(includeHeaders []string) *Access {
	a.IncludeHeaders = includeHeaders
	return a
}

// WithExcludePaths 设置排除的路径
func (a *Access) WithExcludePaths(excludePaths []string) *Access {
	a.ExcludePaths = excludePaths
	return a
}

// WithMaxBodySize 设置最大请求体大小
func (a *Access) WithMaxBodySize(maxBodySize int64) *Access {
	a.MaxBodySize = maxBodySize
	return a
}

// WithMaxResponseSize 设置最大响应体大小
func (a *Access) WithMaxResponseSize(maxResponseSize int64) *Access {
	a.MaxResponseSize = maxResponseSize
	return a
}

// Enable 启用访问记录中间件
func (a *Access) Enable() *Access {
	a.Enabled = true
	return a
}

// Disable 禁用访问记录中间件
func (a *Access) Disable() *Access {
	a.Enabled = false
	return a
}

// IsEnabled 检查是否启用
func (a *Access) IsEnabled() bool {
	return a.Enabled
}