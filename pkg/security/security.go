/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-13 00:00:57
 * @FilePath: \go-config\pkg\security\security.go
 * @Description: 统一安全配置模块 - 管理所有安全相关功能
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package security

import "github.com/kamalyes/go-config/internal"

// Security 统一安全配置 - 直接使用Security而不是SecurityConfig
type Security struct {
	ModuleName string      `mapstructure:"module_name" yaml:"module_name" json:"module_name"` // 模块名称
	Enabled    bool        `mapstructure:"enabled" yaml:"enabled" json:"enabled"`             // 是否启用安全功能
	JWT        *JWT        `mapstructure:"jwt" yaml:"jwt" json:"jwt"`                         // JWT配置
	Auth       *Auth       `mapstructure:"auth" yaml:"auth" json:"auth"`                      // 通用认证配置
	Protection *Protection `mapstructure:"protection" yaml:"protection" json:"protection"`    // 服务保护配置
}

// JWT JWT配置
type JWT struct {
	Enabled   bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`       // 是否启用JWT
	Secret    string `mapstructure:"secret" yaml:"secret" json:"secret"`          // JWT密钥
	Expiry    int    `mapstructure:"expiry" yaml:"expiry" json:"expiry"`          // 过期时间(小时)
	Issuer    string `mapstructure:"issuer" yaml:"issuer" json:"issuer"`          // 签发者
	Algorithm string `mapstructure:"algorithm" yaml:"algorithm" json:"algorithm"` // 加密算法
}

// Auth 通用认证配置 - 支持多种认证方式
type Auth struct {
	Enabled     bool        `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                // 是否启用认证
	Type        string      `mapstructure:"type" yaml:"type" json:"type"`                         // 认证类型 (basic, bearer, custom, apikey)
	HeaderName  string      `mapstructure:"header_name" yaml:"header_name" json:"header_name"`    // 认证头名称
	TokenPrefix string      `mapstructure:"token_prefix" yaml:"token_prefix" json:"token_prefix"` // 令牌前缀
	Basic       *BasicAuth  `mapstructure:"basic" yaml:"basic" json:"basic"`                      // Basic认证
	Bearer      *BearerAuth `mapstructure:"bearer" yaml:"bearer" json:"bearer"`                   // Bearer认证
	APIKey      *APIKeyAuth `mapstructure:"apikey" yaml:"apikey" json:"apikey"`                   // APIKey认证
	Custom      *CustomAuth `mapstructure:"custom" yaml:"custom" json:"custom"`                   // 自定义认证
}

// BasicAuth Basic认证配置
type BasicAuth struct {
	Users []User `mapstructure:"users" yaml:"users" json:"users"` // 用户列表
}

// BearerAuth Bearer认证配置
type BearerAuth struct {
	Tokens []string `mapstructure:"tokens" yaml:"tokens" json:"tokens"` // 有效令牌列表
}

// APIKeyAuth API Key认证配置
type APIKeyAuth struct {
	Keys       []string `mapstructure:"keys" yaml:"keys" json:"keys"`                      // 有效API Key列表
	HeaderName string   `mapstructure:"header_name" yaml:"header_name" json:"header_name"` // API Key头名称
	QueryParam string   `mapstructure:"query_param" yaml:"query_param" json:"query_param"` // API Key查询参数名
}

// CustomAuth 自定义认证配置
type CustomAuth struct {
	HeaderName    string            `mapstructure:"header_name" yaml:"header_name" json:"header_name"`          // 自定义头名称
	ExpectedValue string            `mapstructure:"expected_value" yaml:"expected_value" json:"expected_value"` // 期望值
	Headers       map[string]string `mapstructure:"headers" yaml:"headers" json:"headers"`                      // 自定义头部验证
}

// User 用户信息
type User struct {
	Username    string   `mapstructure:"username" yaml:"username" json:"username"`          // 用户名
	Password    string   `mapstructure:"password" yaml:"password" json:"password"`          // 密码(建议加密存储)
	Role        string   `mapstructure:"role" yaml:"role" json:"role"`                      // 角色
	Permissions []string `mapstructure:"permissions" yaml:"permissions" json:"permissions"` // 权限列表
}

// Protection 服务保护配置 - 针对不同服务的安全设置
type Protection struct {
	Swagger *ServiceProtection `mapstructure:"swagger" yaml:"swagger" json:"swagger"` // Swagger保护
	PProf   *ServiceProtection `mapstructure:"pprof" yaml:"pprof" json:"pprof"`       // PProf保护
	Metrics *ServiceProtection `mapstructure:"metrics" yaml:"metrics" json:"metrics"` // Metrics保护
	Health  *ServiceProtection `mapstructure:"health" yaml:"health" json:"health"`    // Health保护
	API     *ServiceProtection `mapstructure:"api" yaml:"api" json:"api"`             // API保护
}

// Pprof 安全配置
type Pprof struct {
	Enabled      bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                   // 是否启用安全认证
	Username     string   `mapstructure:"username" yaml:"username" json:"username"`                // 用户名
	Password     string   `mapstructure:"password" yaml:"password" json:"password"`                // 密码
	AllowedIPs   []string `mapstructure:"allowed_ips" yaml:"allowed_ips" json:"allowed_ips"`       // 允许的IP地址
	RequireHTTPS bool     `mapstructure:"require_https" yaml:"require_https" json:"require_https"` // 是否要求HTTPS
}

// ServiceProtection 服务保护配置 - 统一的服务安全配置
type ServiceProtection struct {
	Enabled      bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                   // 是否启用保护
	AuthRequired bool     `mapstructure:"auth_required" yaml:"auth_required" json:"auth_required"` // 是否需要认证
	AuthType     string   `mapstructure:"auth_type" yaml:"auth_type" json:"auth_type"`             // 认证类型
	IPWhitelist  []string `mapstructure:"ip_whitelist" yaml:"ip_whitelist" json:"ip_whitelist"`    // IP白名单
	RequireHTTPS bool     `mapstructure:"require_https" yaml:"require_https" json:"require_https"` // 是否要求HTTPS
	Username     string   `mapstructure:"username" yaml:"username" json:"username"`                // 用户名
	Password     string   `mapstructure:"password" yaml:"password" json:"password"`                // 密码
}

// Default 创建默认安全配置
func Default() *Security {
	return &Security{
		ModuleName: "security",
		Enabled:    true,
		JWT: &JWT{
			Enabled:   false,
			Secret:    "jwt_secret_key_please_change_in_production",
			Expiry:    24,
			Issuer:    "go-rpc-gateway",
			Algorithm: "HS256",
		},
		Auth: &Auth{
			Enabled:     false,
			Type:        "bearer",
			HeaderName:  "Authorization",
			TokenPrefix: "Bearer ",
			Basic:       &BasicAuth{Users: []User{}},
			Bearer:      &BearerAuth{Tokens: []string{}},
			APIKey: &APIKeyAuth{
				Keys:       []string{},
				HeaderName: "X-API-Key",
				QueryParam: "api_key",
			},
			Custom: &CustomAuth{
				Headers: make(map[string]string),
			},
		},
		Protection: &Protection{
			Swagger: &ServiceProtection{
				Enabled:      false,
				AuthRequired: false,
				AuthType:     "basic",
				IPWhitelist:  []string{},
				RequireHTTPS: false,
			},
			PProf: &ServiceProtection{
				Enabled:      false,
				AuthRequired: true,
				AuthType:     "basic",
				IPWhitelist:  []string{"127.0.0.1", "::1"},
				RequireHTTPS: true,
				Username:     "admin",
				Password:     "",
			},
			Metrics: &ServiceProtection{
				Enabled:      false,
				AuthRequired: false,
				IPWhitelist:  []string{},
			},
			Health: &ServiceProtection{
				Enabled:      false,
				AuthRequired: false,
				IPWhitelist:  []string{},
			},
			API: &ServiceProtection{
				Enabled:      true,
				AuthRequired: false,
				IPWhitelist:  []string{},
			},
		},
	}
}

// Get 返回配置接口
func (s *Security) Get() interface{} {
	return s
}

// Set 设置配置数据
func (s *Security) Set(data interface{}) {
	if cfg, ok := data.(*Security); ok {
		*s = *cfg
	}
}

// Clone 返回配置的副本
func (s *Security) Clone() internal.Configurable {
	cloned := Default() // 创建默认配置作为基础
	cloned.ModuleName = s.ModuleName
	cloned.Enabled = s.Enabled
	// 简化处理，实际可以进行深度克隆
	return cloned
}

// Validate 验证配置
func (s *Security) Validate() error {
	return internal.ValidateStruct(s)
}

// WithModuleName 设置模块名称
func (s *Security) WithModuleName(moduleName string) *Security {
	s.ModuleName = moduleName
	return s
}

// WithEnabled 设置是否启用安全功能
func (s *Security) WithEnabled(enabled bool) *Security {
	s.Enabled = enabled
	return s
}

// WithJWT 设置JWT配置
func (s *Security) WithJWT(enabled bool, secret string, expiry int, issuer, algorithm string) *Security {
	if s.JWT == nil {
		s.JWT = &JWT{}
	}
	s.JWT.Enabled = enabled
	s.JWT.Secret = secret
	s.JWT.Expiry = expiry
	s.JWT.Issuer = issuer
	s.JWT.Algorithm = algorithm
	return s
}

// WithAuth 设置认证配置
func (s *Security) WithAuth(enabled bool, authType, headerName, tokenPrefix string) *Security {
	if s.Auth == nil {
		s.Auth = &Auth{}
	}
	s.Auth.Enabled = enabled
	s.Auth.Type = authType
	s.Auth.HeaderName = headerName
	s.Auth.TokenPrefix = tokenPrefix
	return s
}

// EnableJWT 启用JWT
func (s *Security) EnableJWT() *Security {
	if s.JWT == nil {
		s.JWT = &JWT{}
	}
	s.JWT.Enabled = true
	return s
}

// EnableAuth 启用认证
func (s *Security) EnableAuth() *Security {
	if s.Auth == nil {
		s.Auth = &Auth{}
	}
	s.Auth.Enabled = true
	return s
}

// Enable 启用安全功能
func (s *Security) Enable() *Security {
	s.Enabled = true
	return s
}

// Disable 禁用安全功能
func (s *Security) Disable() *Security {
	s.Enabled = false
	return s
}

// IsEnabled 检查是否启用
func (s *Security) IsEnabled() bool {
	return s.Enabled
}
