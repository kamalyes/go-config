/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-15 00:00:00
 * @FilePath: \go-config\pkg\rpcserver\rpcserver.go
 * @Description: RPC服务器配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package rpcserver

import (
	"github.com/kamalyes/go-config/internal"
)

// RpcServer 结构体表示 RPC 服务器配置
type RpcServer struct {
	ModuleName    string            `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`          // 模块名称
	Enabled       bool              `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                     // 是否启用
	Name          string            `mapstructure:"name" yaml:"name" json:"name"`                              // 服务名称
	ListenOn      string            `mapstructure:"listen-on" yaml:"listen-on" json:"listenOn"`                // 监听地址
	Mode          string            `mapstructure:"mode" yaml:"mode" json:"mode"`                              // 运行模式 (dev, test, prod)
	MaxConns      int               `mapstructure:"max-conns" yaml:"max-conns" json:"maxConns"`                // 最大连接数
	MaxMsgSize    int               `mapstructure:"max-msg-size" yaml:"max-msg-size" json:"maxMsgSize"`        // 最大消息大小
	Timeout       int               `mapstructure:"timeout" yaml:"timeout" json:"timeout"`                     // 超时时间(秒)
	CpuThreshold  int64             `mapstructure:"cpu-threshold" yaml:"cpu-threshold" json:"cpuThreshold"`    // CPU阈值
	Health        bool              `mapstructure:"health" yaml:"health" json:"health"`                        // 是否启用健康检查
	Auth          bool              `mapstructure:"auth" yaml:"auth" json:"auth"`                              // 是否启用认证
	StrictControl bool              `mapstructure:"strict-control" yaml:"strict-control" json:"strictControl"` // 是否启用严格控制
	MetricsUrl    string            `mapstructure:"metrics-url" yaml:"metrics-url" json:"metricsUrl"`          // 指标 URL
	Headers       map[string]string `mapstructure:"headers" yaml:"headers" json:"headers"`                     // 自定义头部
	Middlewares   []string          `mapstructure:"middlewares" yaml:"middlewares" json:"middlewares"`         // 中间件列表
	TLS           *TLSConfig        `mapstructure:"tls" yaml:"tls" json:"tls"`                                 // TLS配置
	RateLimit     *RateLimit        `mapstructure:"rate-limit" yaml:"rate-limit" json:"rateLimit"`             // 限流配置
	Recovery      *Recovery         `mapstructure:"recovery" yaml:"recovery" json:"recovery"`                  // 恢复配置
	Tracing       *Tracing          `mapstructure:"tracing" yaml:"tracing" json:"tracing"`                     // 链路追踪配置
}

// TLSConfig TLS配置
type TLSConfig struct {
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

// Recovery 恢复配置
type Recovery struct {
	Enabled    bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"`            // 是否启用恢复
	StackTrace bool `mapstructure:"stack-trace" yaml:"stack-trace" json:"stackTrace"` // 是否打印堆栈
	LogErrors  bool `mapstructure:"log-errors" yaml:"log-errors" json:"logErrors"`    // 是否记录错误
}

// Tracing 链路追踪配置
type Tracing struct {
	Enabled     bool    `mapstructure:"enabled" yaml:"enabled" json:"enabled"`               // 是否启用链路追踪
	Endpoint    string  `mapstructure:"endpoint" yaml:"endpoint" json:"endpoint"`            // 追踪端点
	Sampler     float64 `mapstructure:"sampler" yaml:"sampler" json:"sampler"`               // 采样率
	ServiceName string  `mapstructure:"service-name" yaml:"service-name" json:"serviceName"` // 服务名称
}

// NewRpcServer 创建一个新的 RpcServer 实例
func NewRpcServer(opt *RpcServer) *RpcServer {
	var serverInstance *RpcServer

	internal.LockFunc(func() {
		serverInstance = opt
	})
	return serverInstance
}

// Clone 返回 RpcServer 配置的副本
func (r *RpcServer) Clone() internal.Configurable {
	var tls *TLSConfig
	var rateLimit *RateLimit
	var recovery *Recovery
	var tracing *Tracing

	if r.TLS != nil {
		tls = &TLSConfig{
			Enabled:    r.TLS.Enabled,
			CertFile:   r.TLS.CertFile,
			KeyFile:    r.TLS.KeyFile,
			CACertFile: r.TLS.CACertFile,
		}
	}

	if r.RateLimit != nil {
		rateLimit = &RateLimit{
			Enabled: r.RateLimit.Enabled,
			Seconds: r.RateLimit.Seconds,
			Quota:   r.RateLimit.Quota,
		}
	}

	if r.Recovery != nil {
		recovery = &Recovery{
			Enabled:    r.Recovery.Enabled,
			StackTrace: r.Recovery.StackTrace,
			LogErrors:  r.Recovery.LogErrors,
		}
	}

	if r.Tracing != nil {
		tracing = &Tracing{
			Enabled:     r.Tracing.Enabled,
			Endpoint:    r.Tracing.Endpoint,
			Sampler:     r.Tracing.Sampler,
			ServiceName: r.Tracing.ServiceName,
		}
	}

	headers := make(map[string]string)
	for k, v := range r.Headers {
		headers[k] = v
	}

	return &RpcServer{
		ModuleName:    r.ModuleName,
		Enabled:       r.Enabled,
		Name:          r.Name,
		ListenOn:      r.ListenOn,
		Mode:          r.Mode,
		MaxConns:      r.MaxConns,
		MaxMsgSize:    r.MaxMsgSize,
		Timeout:       r.Timeout,
		CpuThreshold:  r.CpuThreshold,
		Health:        r.Health,
		Auth:          r.Auth,
		StrictControl: r.StrictControl,
		MetricsUrl:    r.MetricsUrl,
		Headers:       headers,
		Middlewares:   append([]string(nil), r.Middlewares...),
		TLS:           tls,
		RateLimit:     rateLimit,
		Recovery:      recovery,
		Tracing:       tracing,
	}
}

// Get 返回 RpcServer 配置的所有字段
func (r *RpcServer) Get() interface{} {
	return r
}

// Set 更新 RpcServer 配置的字段
func (r *RpcServer) Set(data interface{}) {
	if configData, ok := data.(*RpcServer); ok {
		*r = *configData
	}
}

// Validate 验证 RpcServer 配置的有效性
func (r *RpcServer) Validate() error {
	// 设置默认值
	if r.ModuleName == "" {
		r.ModuleName = "rpc-server"
	}
	if r.ListenOn == "" {
		r.ListenOn = "127.0.0.1:8080"
	}
	if r.Mode == "" {
		r.Mode = "dev"
	}
	if r.MaxConns <= 0 {
		r.MaxConns = 1000
	}
	if r.MaxMsgSize <= 0 {
		r.MaxMsgSize = 4 * 1024 * 1024 // 4MB
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

// DefaultRpcServer 返回默认RpcServer配置
func DefaultRpcServer() RpcServer {
	return RpcServer{
		ModuleName:    "rpc-server",
		Enabled:       false,
		Name:          "rpc-service",
		ListenOn:      "127.0.0.1:8080",
		Mode:          "dev",
		MaxConns:      1000,
		MaxMsgSize:    4 * 1024 * 1024,
		Timeout:       30,
		CpuThreshold:  900,
		Health:        true,
		Auth:          false,
		StrictControl: false,
		MetricsUrl:    "/metrics",
		Headers:       make(map[string]string),
		Middlewares:   []string{},
		TLS: &TLSConfig{
			Enabled: false,
		},
		RateLimit: &RateLimit{
			Enabled: false,
			Seconds: 1,
			Quota:   1000,
		},
		Recovery: &Recovery{
			Enabled:    true,
			StackTrace: false,
			LogErrors:  true,
		},
		Tracing: &Tracing{
			Enabled:     false,
			Sampler:     0.1,
			ServiceName: "rpc-service",
		},
	}
}

// Default 返回默认RpcServer配置的指针，支持链式调用
func Default() *RpcServer {
	config := DefaultRpcServer()
	return &config
}

// WithModuleName 设置模块名称
func (r *RpcServer) WithModuleName(moduleName string) *RpcServer {
	r.ModuleName = moduleName
	return r
}

// WithName 设置服务名称
func (r *RpcServer) WithName(name string) *RpcServer {
	r.Name = name
	return r
}

// WithListenOn 设置监听地址
func (r *RpcServer) WithListenOn(listenOn string) *RpcServer {
	r.ListenOn = listenOn
	return r
}

// WithMode 设置运行模式
func (r *RpcServer) WithMode(mode string) *RpcServer {
	r.Mode = mode
	return r
}

// WithMaxConns 设置最大连接数
func (r *RpcServer) WithMaxConns(maxConns int) *RpcServer {
	r.MaxConns = maxConns
	return r
}

// WithTimeout 设置超时时间
func (r *RpcServer) WithTimeout(timeout int) *RpcServer {
	r.Timeout = timeout
	return r
}

// WithTLS 设置TLS配置
func (r *RpcServer) WithTLS(enabled bool, certFile, keyFile, caCertFile string) *RpcServer {
	if r.TLS == nil {
		r.TLS = &TLSConfig{}
	}
	r.TLS.Enabled = enabled
	r.TLS.CertFile = certFile
	r.TLS.KeyFile = keyFile
	r.TLS.CACertFile = caCertFile
	return r
}

// WithRateLimit 设置限流配置
func (r *RpcServer) WithRateLimit(enabled bool, seconds, quota int) *RpcServer {
	if r.RateLimit == nil {
		r.RateLimit = &RateLimit{}
	}
	r.RateLimit.Enabled = enabled
	r.RateLimit.Seconds = seconds
	r.RateLimit.Quota = quota
	return r
}

// EnableAuth 启用认证
func (r *RpcServer) EnableAuth() *RpcServer {
	r.Auth = true
	return r
}

// EnableHealth 启用健康检查
func (r *RpcServer) EnableHealth() *RpcServer {
	r.Health = true
	return r
}

// EnableTracing 启用链路追踪
func (r *RpcServer) EnableTracing(endpoint, serviceName string, sampler float64) *RpcServer {
	if r.Tracing == nil {
		r.Tracing = &Tracing{}
	}
	r.Tracing.Enabled = true
	r.Tracing.Endpoint = endpoint
	r.Tracing.ServiceName = serviceName
	r.Tracing.Sampler = sampler
	return r
}

// Enable 启用RPC服务器
func (r *RpcServer) Enable() *RpcServer {
	r.Enabled = true
	return r
}

// Disable 禁用RPC服务器
func (r *RpcServer) Disable() *RpcServer {
	r.Enabled = false
	return r
}
