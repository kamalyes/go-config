/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 00:41:11
 * @FilePath: \go-config\pkg\register\pprof.go
 * @Description:
 *
 * Copyright (j) 2024 by kamalyes, All Rights Reserved.
 */
package register

import (
	"net/http"

	"github.com/kamalyes/go-config/internal"
)

// PProfConfig pprof配置
// 用于控制pprof的启用、路径、认证、日志等参数
// 可用于多种配置格式（json/yaml/toml）
type PProfConfig struct {
	// 是否启用pprof
	Enabled bool `json:"enabled" yaml:"enabled" toml:"enabled"`

	// pprof路径前缀
	PathPrefix string `json:"path_prefix" yaml:"path_prefix" toml:"path_prefix"`

	// 允许访问的IP地址列表，为空表示允许所有
	AllowedIPs []string `json:"allowed_ips" yaml:"allowed_ips" toml:"allowed_ips"`

	// 是否需要认证
	RequireAuth bool `json:"require_auth" yaml:"require_auth" toml:"require_auth"`

	// 认证令牌
	AuthToken string `json:"auth_token" yaml:"auth_token" toml:"auth_token"`

	// 是否启用请求日志
	EnableLogging bool `json:"enable_logging" yaml:"enable_logging" toml:"enable_logging"`

	// 超时设置（秒）
	Timeout int `json:"timeout" yaml:"timeout" toml:"timeout"`

	// 自定义处理器映射
	CustomHandlers map[string]http.HandlerFunc `json:"-" yaml:"-" toml:"-"`
}


// NewPProfConfig 创建一个新的 PProfConfig 实例
func NewPProfConfig(opt *PProfConfig) *PProfConfig {
	var pprofInstance *PProfConfig

	internal.LockFunc(func() {
		if opt == nil {
			opt = &PProfConfig{}
		}
		
		// 设置默认值
		if opt.PathPrefix == "" {
			opt.PathPrefix = "/debug/pprof"
		}
		if opt.Timeout == 0 {
			opt.Timeout = 30
		}
		
		pprofInstance = opt
	})
	return pprofInstance
}

// Clone 返回 PProfConfig 配置的副本
func (p *PProfConfig) Clone() internal.Configurable {
	// 复制自定义处理器
	customHandlers := make(map[string]http.HandlerFunc)
	for k, v := range p.CustomHandlers {
		customHandlers[k] = v
	}

	// 复制允许的IP列表
	allowedIPs := make([]string, len(p.AllowedIPs))
	copy(allowedIPs, p.AllowedIPs)

	return &PProfConfig{
		Enabled:        p.Enabled,
		PathPrefix:     p.PathPrefix,
		AllowedIPs:     allowedIPs,
		RequireAuth:    p.RequireAuth,
		AuthToken:      p.AuthToken,
		EnableLogging:  p.EnableLogging,
		Timeout:        p.Timeout,
		CustomHandlers: customHandlers,
	}
}

// Get 返回 PProfConfig 配置的所有字段
func (p *PProfConfig) Get() interface{} {
	return p
}

// Set 更新 PProfConfig 配置的字段
func (p *PProfConfig) Set(data interface{}) {
	if configData, ok := data.(*PProfConfig); ok {
		p.Enabled = configData.Enabled
		p.PathPrefix = configData.PathPrefix
		p.AllowedIPs = configData.AllowedIPs
		p.RequireAuth = configData.RequireAuth
		p.AuthToken = configData.AuthToken
		p.EnableLogging = configData.EnableLogging
		p.Timeout = configData.Timeout
		p.CustomHandlers = configData.CustomHandlers
	}
}

// Validate 检查 PProfConfig 配置的有效性
func (p *PProfConfig) Validate() error {
	return internal.ValidateStruct(p)
}
