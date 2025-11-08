/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 00:38:17
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

// PProf pprof配置
// 用于控制pprof的启用、路径、认证、日志等参数
// 可用于多种配置格式（json/yaml/toml）
type PProf struct {
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
	ModuleName         string `mapstructure:"modulename"               yaml:"modulename"                json:"module_name"`                               // 模块名称
}


// NewPProf 创建一个新的 PProf 实例
func NewPProf(opt *PProf) *PProf {
	var pprofInstance *PProf

	internal.LockFunc(func() {
		if opt == nil {
			opt = &PProf{}
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

// Clone 返回 PProf 配置的副本
func (p *PProf) Clone() internal.Configurable {
	// 复制自定义处理器
	customHandlers := make(map[string]http.HandlerFunc)
	for k, v := range p.CustomHandlers {
		customHandlers[k] = v
	}

	// 复制允许的IP列表
	allowedIPs := make([]string, len(p.AllowedIPs))
	copy(allowedIPs, p.AllowedIPs)

	return &PProf{
		ModuleName:    p.ModuleName,
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

// Get 返回 PProf 配置的所有字段
func (p *PProf) Get() interface{} {
	return p
}

// Set 更新 PProf 配置的字段
func (p *PProf) Set(data interface{}) {
	if configData, ok := data.(*PProf); ok {
		p.ModuleName = configData.ModuleName
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

// Validate 检查 PProf 配置的有效性
func (p *PProf) Validate() error {
	return internal.ValidateStruct(p)
}

// DefaultPProf 返回默认PProf配置
func DefaultPProf() PProf {
	return PProf{
		ModuleName:     "pprof",
		Enabled:        false,
		PathPrefix:     "/debug/pprof",
		AllowedIPs:     []string{},
		RequireAuth:    false,
		AuthToken:      "",
		EnableLogging:  false,
		Timeout:        30,
		CustomHandlers: make(map[string]http.HandlerFunc),
	}
}

// DefaultPProfConfig 返回默认PProf配置的指针，支持链式调用
func DefaultPProfConfig() *PProf {
	config := DefaultPProf()
	return &config
}

// WithModuleName 设置模块名称
func (p *PProf) WithModuleName(moduleName string) *PProf {
	p.ModuleName = moduleName
	return p
}

// WithEnabled 设置是否启用pprof
func (p *PProf) WithEnabled(enabled bool) *PProf {
	p.Enabled = enabled
	return p
}

// WithPathPrefix 设置pprof路径前缀
func (p *PProf) WithPathPrefix(pathPrefix string) *PProf {
	p.PathPrefix = pathPrefix
	return p
}

// WithAllowedIPs 设置允许访问的IP地址列表
func (p *PProf) WithAllowedIPs(allowedIPs []string) *PProf {
	p.AllowedIPs = allowedIPs
	return p
}

// WithRequireAuth 设置是否需要认证
func (p *PProf) WithRequireAuth(requireAuth bool) *PProf {
	p.RequireAuth = requireAuth
	return p
}

// WithAuthToken 设置认证令牌
func (p *PProf) WithAuthToken(authToken string) *PProf {
	p.AuthToken = authToken
	return p
}

// WithEnableLogging 设置是否启用请求日志
func (p *PProf) WithEnableLogging(enableLogging bool) *PProf {
	p.EnableLogging = enableLogging
	return p
}

// WithTimeout 设置超时时间（秒）
func (p *PProf) WithTimeout(timeout int) *PProf {
	p.Timeout = timeout
	return p
}

// WithCustomHandlers 设置自定义处理器映射
func (p *PProf) WithCustomHandlers(customHandlers map[string]http.HandlerFunc) *PProf {
	p.CustomHandlers = customHandlers
	return p
}
