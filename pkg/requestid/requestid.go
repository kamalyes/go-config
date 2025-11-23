/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 18:00:00
 * @FilePath: \go-config\pkg\requestid\requestid.go
 * @Description: 请求ID中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package requestid

import "github.com/kamalyes/go-config/internal"

// RequestID 请求ID中间件配置
type RequestID struct {
	ModuleName string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"` // 模块名称
	Enabled    bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`            // 是否启用请求ID
	HeaderName string `mapstructure:"header-name" yaml:"header-name" json:"headerName"` // 头部名称
	Generator  string `mapstructure:"generator" yaml:"generator" json:"generator"`      // 生成器类型 (uuid, nanoid)
}

// Default 创建默认请求ID配置
func Default() *RequestID {
	return &RequestID{
		ModuleName: "requestid",
		Enabled:    true,
		HeaderName: "X-Request-ID",
		Generator:  "uuid",
	}
}

// Get 返回配置接口
func (r *RequestID) Get() interface{} {
	return r
}

// Set 设置配置数据
func (r *RequestID) Set(data interface{}) {
	if cfg, ok := data.(*RequestID); ok {
		*r = *cfg
	}
}

// Clone 返回配置的副本
func (r *RequestID) Clone() internal.Configurable {
	return &RequestID{
		ModuleName: r.ModuleName,
		Enabled:    r.Enabled,
		HeaderName: r.HeaderName,
		Generator:  r.Generator,
	}
}

// Validate 验证配置
func (r *RequestID) Validate() error {
	return internal.ValidateStruct(r)
}

// WithHeaderName 设置头部名称
func (r *RequestID) WithHeaderName(headerName string) *RequestID {
	r.HeaderName = headerName
	return r
}

// WithGenerator 设置生成器类型
func (r *RequestID) WithGenerator(generator string) *RequestID {
	r.Generator = generator
	return r
}

// Enable 启用请求ID中间件
func (r *RequestID) Enable() *RequestID {
	r.Enabled = true
	return r
}

// Disable 禁用请求ID中间件
func (r *RequestID) Disable() *RequestID {
	r.Enabled = false
	return r
}

// IsEnabled 检查是否启用
func (r *RequestID) IsEnabled() bool {
	return r.Enabled
}
