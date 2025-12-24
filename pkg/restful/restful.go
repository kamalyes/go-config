/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-15 00:00:00
 * @FilePath: \go-config\pkg\restful\restful.go
 * @Description: RESTful API配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package restful

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-toolbox/pkg/syncx"
)

// Restful 结构体表示 RESTful API 配置
type Restful struct {
	ModuleName   string            `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`       // 模块名称
	Enabled      bool              `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                  // 是否启用
	Name         string            `mapstructure:"name" yaml:"name" json:"name"`                           // 服务名称
	Host         string            `mapstructure:"host" yaml:"host" json:"host"`                           // 主机地址
	Port         int               `mapstructure:"port" yaml:"port" json:"port"`                           // 端口
	Mode         string            `mapstructure:"mode" yaml:"mode" json:"mode"`                           // 运行模式 (dev, test, prod)
	MaxConns     int               `mapstructure:"max-conns" yaml:"max-conns" json:"maxConns"`             // 最大连接数
	MaxBytes     int64             `mapstructure:"max-bytes" yaml:"max-bytes" json:"maxBytes"`             // 最大请求大小
	Timeout      int               `mapstructure:"timeout" yaml:"timeout" json:"timeout"`                  // 超时时间(秒)
	CpuThreshold int64             `mapstructure:"cpu-threshold" yaml:"cpu-threshold" json:"cpuThreshold"` // CPU阈值
	Signature    *Signature        `mapstructure:"signature" yaml:"signature" json:"signature"`            // 签名配置
	Auth         bool              `mapstructure:"auth" yaml:"auth" json:"auth"`                           // 是否启用认证
	PrintRoutes  bool              `mapstructure:"print-routes" yaml:"print-routes" json:"printRoutes"`    // 是否打印路由
	StrictSlash  bool              `mapstructure:"strict-slash" yaml:"strict-slash" json:"strictSlash"`    // 是否严格斜杠
	Headers      map[string]string `mapstructure:"headers" yaml:"headers" json:"headers"`                  // 自定义头部
	Middlewares  []string          `mapstructure:"middlewares" yaml:"middlewares" json:"middlewares"`      // 中间件列表
	CORS         *CORS             `mapstructure:"cors" yaml:"cors" json:"cors"`                           // CORS配置
	TLS          *TLS              `mapstructure:"tls" yaml:"tls" json:"tls"`                              // TLS配置
	RateLimit    *RateLimit        `mapstructure:"rate-limit" yaml:"rate-limit" json:"rateLimit"`          // 限流配置
	Compression  *Compression      `mapstructure:"compression" yaml:"compression" json:"compression"`      // 压缩配置
	Static       *Static           `mapstructure:"static" yaml:"static" json:"static"`                     // 静态文件配置
}

// Signature 签名配置
type Signature struct {
	Enabled     bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`               // 是否启用签名
	PrivateKeys []string `mapstructure:"private-keys" yaml:"private-keys" json:"privateKeys"` // 私钥列表
	Strict      bool     `mapstructure:"strict" yaml:"strict" json:"strict"`                  // 严格模式
	Expiry      int      `mapstructure:"expiry" yaml:"expiry" json:"expiry"`                  // 过期时间(秒)
}

// CORS 跨域配置
type CORS struct {
	Enabled          bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                              // 是否启用CORS
	AllowOrigins     []string `mapstructure:"allow-origins" yaml:"allow-origins" json:"allowOrigins"`             // 允许的来源
	AllowMethods     []string `mapstructure:"allow-methods" yaml:"allow-methods" json:"allowMethods"`             // 允许的方法
	AllowHeaders     []string `mapstructure:"allow-headers" yaml:"allow-headers" json:"allowHeaders"`             // 允许的头部
	ExposeHeaders    []string `mapstructure:"expose-headers" yaml:"expose-headers" json:"exposeHeaders"`          // 暴露的头部
	AllowCredentials bool     `mapstructure:"allow-credentials" yaml:"allow-credentials" json:"allowCredentials"` // 是否允许凭证
	MaxAge           int      `mapstructure:"max-age" yaml:"max-age" json:"maxAge"`                               // 预检请求缓存时间
}

// TLS TLS配置
type TLS struct {
	Enabled    bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`              // 是否启用TLS
	CertFile   string `mapstructure:"cert-file" yaml:"cert-file" json:"certFile"`         // 证书文件
	KeyFile    string `mapstructure:"key-file" yaml:"key-file" json:"keyFile"`            // 私钥文件
	CACertFile string `mapstructure:"ca-cert-file" yaml:"ca-cert-file" json:"caCertFile"` // CA证书文件
}

// RateLimit 限流配置
type RateLimit struct {
	Enabled bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"` // 是否启用限流
	Seconds int  `mapstructure:"seconds" yaml:"seconds" json:"seconds"` // 时间窗口(秒)
	Quota   int  `mapstructure:"quota" yaml:"quota" json:"quota"`       // 配额
}

// Compression 压缩配置
type Compression struct {
	Enabled   bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`         // 是否启用压缩
	Level     int      `mapstructure:"level" yaml:"level" json:"level"`               // 压缩级别
	MinLength int      `mapstructure:"min-length" yaml:"min-length" json:"minLength"` // 最小压缩长度
	Types     []string `mapstructure:"types" yaml:"types" json:"types"`               // 压缩的MIME类型
}

// Static 静态文件配置
type Static struct {
	Enabled bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"` // 是否启用静态文件服务
	Root    string `mapstructure:"root" yaml:"root" json:"root"`          // 静态文件根目录
	Prefix  string `mapstructure:"prefix" yaml:"prefix" json:"prefix"`    // URL前缀
	Index   string `mapstructure:"index" yaml:"index" json:"index"`       // 默认首页
	Browse  bool   `mapstructure:"browse" yaml:"browse" json:"browse"`    // 是否允许目录浏览
}

// NewRestful 创建一个新的 Restful 实例
func NewRestful(opt *Restful) *Restful {
	var restfulInstance *Restful

	internal.LockFunc(func() {
		restfulInstance = opt
	})
	return restfulInstance
}

// Clone 返回 Restful 配置的副本
func (r *Restful) Clone() internal.Configurable {
	var cloned Restful
	if err := syncx.DeepCopy(&cloned, r); err != nil {
		// 如果深拷贝失败，返回空配置
		return &Restful{}
	}
	return &cloned
}

// Get 返回 Restful 配置的所有字段
func (r *Restful) Get() interface{} {
	return r
}

// Set 更新 Restful 配置的字段
func (r *Restful) Set(data interface{}) {
	if configData, ok := data.(*Restful); ok {
		*r = *configData
	}
}

// Validate 验证 Restful 配置的有效性
func (r *Restful) Validate() error {
	// 设置默认值
	if r.ModuleName == "" {
		r.ModuleName = "restful"
	}
	if r.Host == "" {
		r.Host = "127.0.0.1"
	}
	if r.Port <= 0 {
		r.Port = 8080
	}
	if r.Mode == "" {
		r.Mode = "dev"
	}
	if r.MaxConns <= 0 {
		r.MaxConns = 1000
	}
	if r.MaxBytes <= 0 {
		r.MaxBytes = 32 * 1024 * 1024 // 32MB
	}
	if r.Timeout <= 0 {
		r.Timeout = 30
	}
	if r.CpuThreshold <= 0 {
		r.CpuThreshold = 900 // 90%
	}
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	if r.Middlewares == nil {
		r.Middlewares = []string{}
	}

	return internal.ValidateStruct(r)
}

// DefaultRestful 返回默认Restful配置
func DefaultRestful() Restful {
	return Restful{
		ModuleName:   "restful",
		Enabled:      false,
		Name:         "restful-api",
		Host:         "127.0.0.1",
		Port:         8080,
		Mode:         "dev",
		MaxConns:     1000,
		MaxBytes:     32 * 1024 * 1024,
		Timeout:      30,
		CpuThreshold: 900,
		Auth:         false,
		PrintRoutes:  false,
		StrictSlash:  false,
		Headers:      make(map[string]string),
		Middlewares:  []string{},
		Signature: &Signature{
			Enabled: false,
			Strict:  false,
			Expiry:  300,
		},
		CORS: &CORS{
			Enabled:          false,
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: false,
			MaxAge:           3600,
		},
		TLS: &TLS{
			Enabled: false,
		},
		RateLimit: &RateLimit{
			Enabled: false,
			Seconds: 1,
			Quota:   1000,
		},
		Compression: &Compression{
			Enabled:   false,
			Level:     6,
			MinLength: 1024,
			Types:     []string{"application/json", "text/html", "text/css", "text/javascript"},
		},
		Static: &Static{
			Enabled: false,
			Root:    "./static",
			Prefix:  "/static",
			Index:   "index.html",
			Browse:  false,
		},
	}
}

// Default 返回默认Restful配置的指针，支持链式调用
func Default() *Restful {
	config := DefaultRestful()
	return &config
}

// WithModuleName 设置模块名称
func (r *Restful) WithModuleName(moduleName string) *Restful {
	r.ModuleName = moduleName
	return r
}

// WithName 设置服务名称
func (r *Restful) WithName(name string) *Restful {
	r.Name = name
	return r
}

// WithHost 设置主机地址
func (r *Restful) WithHost(host string) *Restful {
	r.Host = host
	return r
}

// WithPort 设置端口
func (r *Restful) WithPort(port int) *Restful {
	r.Port = port
	return r
}

// WithMode 设置运行模式
func (r *Restful) WithMode(mode string) *Restful {
	r.Mode = mode
	return r
}

// WithTimeout 设置超时时间
func (r *Restful) WithTimeout(timeout int) *Restful {
	r.Timeout = timeout
	return r
}

// EnableCORS 启用CORS
func (r *Restful) EnableCORS() *Restful {
	if r.CORS != nil {
		r.CORS.Enabled = true
	}
	return r
}

// EnableTLS 启用TLS
func (r *Restful) EnableTLS(certFile, keyFile string) *Restful {
	if r.TLS == nil {
		r.TLS = &TLS{}
	}
	r.TLS.Enabled = true
	r.TLS.CertFile = certFile
	r.TLS.KeyFile = keyFile
	return r
}

// EnableStatic 启用静态文件服务
func (r *Restful) EnableStatic(root, prefix string) *Restful {
	if r.Static == nil {
		r.Static = &Static{}
	}
	r.Static.Enabled = true
	r.Static.Root = root
	r.Static.Prefix = prefix
	return r
}

// Enable 启用RESTful API
func (r *Restful) Enable() *Restful {
	r.Enabled = true
	return r
}

// Disable 禁用RESTful API
func (r *Restful) Disable() *Restful {
	r.Enabled = false
	return r
}
