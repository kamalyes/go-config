/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 11:52:03
 * @FilePath: \go-config\pkg\jwt\jwt.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package jwt

import (
	"github.com/kamalyes/go-config/internal"
)

// JWT 结构体用于配置 JSON Web Token 的相关参数
type JWT struct {
	ModuleName       string                 `mapstructure:"modulename"        yaml:"modulename"      json:"module_name"`                                // 模块名称
	SigningKey       string                 `mapstructure:"signing-key"       yaml:"signing_key"     json:"signing_key"      validate:"required"`       // jwt 签名
	ExpiresTime      int64                  `mapstructure:"expires-time"      yaml:"expires_time"    json:"expires_time"     validate:"required,min=1"` // 过期时间（单位：秒）
	BufferTime       int64                  `mapstructure:"buffer-time"       yaml:"buffer_time"     json:"buffer_time"      validate:"required,min=0"` // 缓冲时间（单位：秒）
	UseMultipoint    bool                   `mapstructure:"use-multipoint"    yaml:"use_multipoint"  json:"use_multipoint"`                             // 多地登录拦截，true 表示拦截，false 表示不拦截
	Issuer           string                 `mapstructure:"issuer"            yaml:"issuer"          json:"issuer"`                                     // JWT 发行者
	Audience         string                 `mapstructure:"audience"          yaml:"audience"        json:"audience"`                                   // JWT 接收者
	Algorithm        string                 `mapstructure:"algorithm"         yaml:"algorithm"       json:"algorithm"      validate:"required"`         // 签名算法，例如 HMAC, RSA, etc.
	EnableRefresh    bool                   `mapstructure:"enable-refresh"    yaml:"enable_refresh"  json:"enable_refresh"`                             // 是否启用刷新 token
	RefreshTokenLife int64                  `mapstructure:"refresh-token-life" yaml:"refresh_token_life" json:"refresh_token_life" validate:"min=1"`    // 刷新 token 生命周期（单位：秒）
	Subject          string                 `mapstructure:"subject"           yaml:"subject"         json:"subject"`                                    // JWT 主题
	CustomClaims     map[string]interface{} `mapstructure:"custom-claims" yaml:"custom_claims" json:"custom_claims"`                                    // 自定义声明
}

// NewJWT 创建一个新的 JWT 实例
func NewJWT(opt *JWT) *JWT {
	var jwtInstance *JWT

	internal.LockFunc(func() {
		jwtInstance = opt
	})

	return jwtInstance
}

// Clone 返回 JWT 配置的副本
func (j *JWT) Clone() internal.Configurable {
	return &JWT{
		ModuleName:       j.ModuleName,
		SigningKey:       j.SigningKey,
		ExpiresTime:      j.ExpiresTime,
		BufferTime:       j.BufferTime,
		UseMultipoint:    j.UseMultipoint,
		Issuer:           j.Issuer,
		Audience:         j.Audience,
		Algorithm:        j.Algorithm,
		EnableRefresh:    j.EnableRefresh,
		RefreshTokenLife: j.RefreshTokenLife,
		Subject:          j.Subject,
		CustomClaims:     j.CustomClaims,
	}
}

// Get 返回 JWT 配置的所有字段
func (j *JWT) Get() interface{} {
	return j
}

// Set 更新 JWT 配置的字段
func (j *JWT) Set(data interface{}) {
	if configData, ok := data.(*JWT); ok {
		j.ModuleName = configData.ModuleName
		j.SigningKey = configData.SigningKey
		j.ExpiresTime = configData.ExpiresTime
		j.BufferTime = configData.BufferTime
		j.UseMultipoint = configData.UseMultipoint
		j.Issuer = configData.Issuer
		j.Audience = configData.Audience
		j.Algorithm = configData.Algorithm
		j.EnableRefresh = configData.EnableRefresh
		j.RefreshTokenLife = configData.RefreshTokenLife
		j.Subject = configData.Subject
		j.CustomClaims = configData.CustomClaims
	}
}

// Validate 检查 JWT 配置的有效性
func (j *JWT) Validate() error {
	return internal.ValidateStruct(j)
}

// DefaultJWT 返回默认JWT配置
func DefaultJWT() JWT {
	return JWT{
		ModuleName:       "jwt",
		SigningKey:       "go-config-default-key",
		ExpiresTime:      3600 * 24 * 7, // 7天，单位：秒
		BufferTime:       3600,          // 1小时，单位：秒
		UseMultipoint:    false,
		Issuer:           "go-config",
		Audience:         "go-config-audience",
		Algorithm:        "HS256",
		EnableRefresh:    true,
		RefreshTokenLife: 3600 * 24 * 30, // 30天，单位：秒
		Subject:          "user-token",
		CustomClaims:     make(map[string]interface{}),
	}
}

// Default 返回默认JWT配置的指针，支持链式调用
func Default() *JWT {
	config := DefaultJWT()
	return &config
}

// WithModuleName 设置模块名称
func (j *JWT) WithModuleName(moduleName string) *JWT {
	j.ModuleName = moduleName
	return j
}

// WithSigningKey 设置签名密钥
func (j *JWT) WithSigningKey(signingKey string) *JWT {
	j.SigningKey = signingKey
	return j
}

// WithExpiresTime 设置过期时间（单位：秒）
func (j *JWT) WithExpiresTime(expiresTime int64) *JWT {
	j.ExpiresTime = expiresTime
	return j
}

// WithBufferTime 设置缓冲时间（单位：秒）
func (j *JWT) WithBufferTime(bufferTime int64) *JWT {
	j.BufferTime = bufferTime
	return j
}

// WithUseMultipoint 设置是否启用多地登录拦截
func (j *JWT) WithUseMultipoint(useMultipoint bool) *JWT {
	j.UseMultipoint = useMultipoint
	return j
}

// WithIssuer 设置发行者
func (j *JWT) WithIssuer(issuer string) *JWT {
	j.Issuer = issuer
	return j
}

// WithAudience 设置接收者
func (j *JWT) WithAudience(audience string) *JWT {
	j.Audience = audience
	return j
}

// WithAlgorithm 设置签名算法
func (j *JWT) WithAlgorithm(algorithm string) *JWT {
	j.Algorithm = algorithm
	return j
}

// WithEnableRefresh 设置是否启用刷新 token
func (j *JWT) WithEnableRefresh(enableRefresh bool) *JWT {
	j.EnableRefresh = enableRefresh
	return j
}

// WithRefreshTokenLife 设置刷新 token 生命周期（单位：秒）
func (j *JWT) WithRefreshTokenLife(refreshTokenLife int64) *JWT {
	j.RefreshTokenLife = refreshTokenLife
	return j
}

// WithSubject 设置主题
func (j *JWT) WithSubject(subject string) *JWT {
	j.Subject = subject
	return j
}

// WithCustomClaims 设置自定义声明
func (j *JWT) WithCustomClaims(customClaims map[string]interface{}) *JWT {
	j.CustomClaims = customClaims
	return j
}

// Load 从另一个 JWT 实例重新加载配置
func (j *JWT) Reload(data interface{}) {
	if configData, ok := data.(*JWT); ok {
		j.ModuleName = configData.ModuleName
		j.SigningKey = configData.SigningKey
		j.ExpiresTime = configData.ExpiresTime
		j.BufferTime = configData.BufferTime
		j.UseMultipoint = configData.UseMultipoint
		j.Issuer = configData.Issuer
		j.Audience = configData.Audience
		j.Algorithm = configData.Algorithm
		j.EnableRefresh = configData.EnableRefresh
		j.RefreshTokenLife = configData.RefreshTokenLife
		j.Subject = configData.Subject
		j.CustomClaims = configData.CustomClaims
	}
}
