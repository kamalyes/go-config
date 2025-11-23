/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 15:05:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 15:05:00
 * @FilePath: \go-config\pkg\signature\signature.go
 * @Description: 签名验证中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package signature

import (
	"github.com/kamalyes/go-config/internal"
	"time"
)

// Signature 签名验证中间件配置
type Signature struct {
	ModuleName      string        `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                // 模块名称
	Enabled         bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                           // 是否启用签名验证
	SecretKey       string        `mapstructure:"secret-key" yaml:"secret-key" json:"secretKey"`                   // 签名密钥
	SignatureHeader string        `mapstructure:"signature-header" yaml:"signature-header" json:"signatureHeader"` // 签名请求头
	TimestampHeader string        `mapstructure:"timestamp-header" yaml:"timestamp-header" json:"timestampHeader"` // 时间戳请求头
	NonceHeader     string        `mapstructure:"nonce-header" yaml:"nonce-header" json:"nonceHeader"`             // 随机数请求头
	Algorithm       string        `mapstructure:"algorithm" yaml:"algorithm" json:"algorithm"`                     // 签名算法 (md5, sha1, sha256)
	TimeoutWindow   time.Duration `mapstructure:"timeout-window" yaml:"timeout-window" json:"timeoutWindow"`       // 请求时间窗口
	IgnorePaths     []string      `mapstructure:"ignore-paths" yaml:"ignore-paths" json:"ignorePaths"`             // 忽略签名验证的路径
	RequiredHeaders []string      `mapstructure:"required-headers" yaml:"required-headers" json:"requiredHeaders"` // 必需的请求头参与签名
	SkipQuery       bool          `mapstructure:"skip-query" yaml:"skip-query" json:"skipQuery"`                   // 是否跳过查询参数
	SkipBody        bool          `mapstructure:"skip-body" yaml:"skip-body" json:"skipBody"`                      // 是否跳过请求体
}

// Default 创建默认签名验证配置
func Default() *Signature {
	return &Signature{
		ModuleName:      "signature",
		Enabled:         false,
		SecretKey:       "default-secret-key-change-in-production",
		SignatureHeader: "X-Signature",
		TimestampHeader: "X-Timestamp",
		NonceHeader:     "X-Nonce",
		Algorithm:       "sha256",
		TimeoutWindow:   time.Minute * 5,
		IgnorePaths: []string{
			"/health",
			"/metrics",
			"/ping",
		},
		RequiredHeaders: []string{
			"X-Timestamp",
			"X-Nonce",
		},
		SkipQuery: false,
		SkipBody:  false,
	}
}

// Get 返回配置接口
func (s *Signature) Get() interface{} {
	return s
}

// Set 设置配置数据
func (s *Signature) Set(data interface{}) {
	if cfg, ok := data.(*Signature); ok {
		*s = *cfg
	}
}

// Clone 返回配置的副本
func (s *Signature) Clone() internal.Configurable {
	ignorePaths := make([]string, len(s.IgnorePaths))
	copy(ignorePaths, s.IgnorePaths)

	requiredHeaders := make([]string, len(s.RequiredHeaders))
	copy(requiredHeaders, s.RequiredHeaders)

	return &Signature{
		ModuleName:      s.ModuleName,
		Enabled:         s.Enabled,
		SecretKey:       s.SecretKey,
		SignatureHeader: s.SignatureHeader,
		TimestampHeader: s.TimestampHeader,
		NonceHeader:     s.NonceHeader,
		Algorithm:       s.Algorithm,
		TimeoutWindow:   s.TimeoutWindow,
		IgnorePaths:     ignorePaths,
		RequiredHeaders: requiredHeaders,
		SkipQuery:       s.SkipQuery,
		SkipBody:        s.SkipBody,
	}
}

// Validate 验证配置
func (s *Signature) Validate() error {
	return internal.ValidateStruct(s)
}

// WithSecretKey 设置签名密钥
func (s *Signature) WithSecretKey(secretKey string) *Signature {
	s.SecretKey = secretKey
	return s
}

// WithAlgorithm 设置签名算法
func (s *Signature) WithAlgorithm(algorithm string) *Signature {
	s.Algorithm = algorithm
	return s
}

// WithTimeoutWindow 设置时间窗口
func (s *Signature) WithTimeoutWindow(timeout time.Duration) *Signature {
	s.TimeoutWindow = timeout
	return s
}

// WithSignatureHeader 设置签名请求头
func (s *Signature) WithSignatureHeader(header string) *Signature {
	s.SignatureHeader = header
	return s
}

// WithTimestampHeader 设置时间戳请求头
func (s *Signature) WithTimestampHeader(header string) *Signature {
	s.TimestampHeader = header
	return s
}

// WithNonceHeader 设置随机数请求头
func (s *Signature) WithNonceHeader(header string) *Signature {
	s.NonceHeader = header
	return s
}

// AddIgnorePath 添加忽略路径
func (s *Signature) AddIgnorePath(path string) *Signature {
	s.IgnorePaths = append(s.IgnorePaths, path)
	return s
}

// AddRequiredHeader 添加必需参与签名的请求头
func (s *Signature) AddRequiredHeader(header string) *Signature {
	s.RequiredHeaders = append(s.RequiredHeaders, header)
	return s
}

// WithSkipQuery 设置是否跳过查询参数
func (s *Signature) WithSkipQuery(skip bool) *Signature {
	s.SkipQuery = skip
	return s
}

// WithSkipBody 设置是否跳过请求体
func (s *Signature) WithSkipBody(skip bool) *Signature {
	s.SkipBody = skip
	return s
}

// Enable 启用签名验证中间件
func (s *Signature) Enable() *Signature {
	s.Enabled = true
	return s
}

// Disable 禁用签名验证中间件
func (s *Signature) Disable() *Signature {
	s.Enabled = false
	return s
}

// IsEnabled 检查是否启用
func (s *Signature) IsEnabled() bool {
	return s.Enabled
}
