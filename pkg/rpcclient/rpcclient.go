/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-15 00:00:00
 * @FilePath: \go-config\pkg\rpcclient\rpcclient.go
 * @Description: RPC客户端配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package rpcclient

import (
	"github.com/kamalyes/go-config/internal"
)

// RpcClient 结构体表示 RPC 客户端配置
type RpcClient struct {
	ModuleName     string            `mapstructure:"module_name" yaml:"module-name" json:"module_name"`             // 模块名称
	Enabled        bool              `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                         // 是否启用
	Endpoints      []string          `mapstructure:"endpoints" yaml:"endpoints" json:"endpoints"`                   // RPC 服务端点
	Target         string            `mapstructure:"target" yaml:"target" json:"target"`                            // 目标服务
	App            string            `mapstructure:"app" yaml:"app" json:"app"`                                     // 应用名称
	Token          string            `mapstructure:"token" yaml:"token" json:"token"`                               // 认证Token
	NonBlock       bool              `mapstructure:"non_block" yaml:"non-block" json:"non_block"`                   // 是否非阻塞
	Timeout        int               `mapstructure:"timeout" yaml:"timeout" json:"timeout"`                         // 超时时间(毫秒)
	KeepaliveTime  int               `mapstructure:"keepalive_time" yaml:"keepalive-time" json:"keepalive_time"`    // Keepalive时间(秒)
	DialTimeout    int               `mapstructure:"dial_timeout" yaml:"dial-timeout" json:"dial_timeout"`          // 连接超时(秒)
	MaxRetries     int               `mapstructure:"max_retries" yaml:"max-retries" json:"max_retries"`             // 最大重试次数
	RetryInterval  int               `mapstructure:"retry_interval" yaml:"retry-interval" json:"retry_interval"`    // 重试间隔(毫秒)
	LoadBalance    string            `mapstructure:"load_balance" yaml:"load-balance" json:"load_balance"`          // 负载均衡策略 (round_robin, random, weighted)
	Compression    string            `mapstructure:"compression" yaml:"compression" json:"compression"`             // 压缩算法 (gzip, deflate)
	Headers        map[string]string `mapstructure:"headers" yaml:"headers" json:"headers"`                         // 自定义头部
	TLS            *TLSConfig        `mapstructure:"tls" yaml:"tls" json:"tls"`                                     // TLS配置
	CircuitBreaker *CircuitBreaker   `mapstructure:"circuit_breaker" yaml:"circuit-breaker" json:"circuit_breaker"` // 熔断器配置
}

// TLSConfig TLS配置
type TLSConfig struct {
	Enabled            bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                        // 是否启用TLS
	CertFile           string `mapstructure:"cert_file" yaml:"cert-file" json:"cert_file"`                                  // 证书文件
	KeyFile            string `mapstructure:"key_file" yaml:"key-file" json:"key_file"`                                     // 私钥文件
	CACertFile         string `mapstructure:"ca_cert_file" yaml:"ca-cert-file" json:"ca_cert_file"`                         // CA证书文件
	ServerName         string `mapstructure:"server_name" yaml:"server-name" json:"server_name"`                            // 服务器名称
	InsecureSkipVerify bool   `mapstructure:"insecure_skip_verify" yaml:"insecure-skip-verify" json:"insecure_skip_verify"` // 跳过证书验证
}

// CircuitBreaker 熔断器配置
type CircuitBreaker struct {
	Enabled          bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                     // 是否启用熔断器
	MaxRequests      int  `mapstructure:"max_requests" yaml:"max-requests" json:"max_requests"`                      // 最大请求数
	Interval         int  `mapstructure:"interval" yaml:"interval" json:"interval"`                                  // 检测间隔(秒)
	Timeout          int  `mapstructure:"timeout" yaml:"timeout" json:"timeout"`                                     // 熔断超时(秒)
	FailureThreshold int  `mapstructure:"failure_threshold" yaml:"failure-threshold" json:"failure_threshold"`       // 失败阈值
	SuccessThreshold int  `mapstructure:"success_threshold" yaml:"success-threshold" json:"success_threshold"`       // 成功阈值
	HalfOpenMaxCalls int  `mapstructure:"half_open_max_calls" yaml:"half-open-max-calls" json:"half_open_max_calls"` // 半开状态最大调用数
}

// NewRpcClient 创建一个新的 RpcClient 实例
func NewRpcClient(opt *RpcClient) *RpcClient {
	var clientInstance *RpcClient

	internal.LockFunc(func() {
		clientInstance = opt
	})
	return clientInstance
}

// Clone 返回 RpcClient 配置的副本
func (r *RpcClient) Clone() internal.Configurable {
	var tls *TLSConfig
	var cb *CircuitBreaker

	if r.TLS != nil {
		tls = &TLSConfig{
			Enabled:            r.TLS.Enabled,
			CertFile:           r.TLS.CertFile,
			KeyFile:            r.TLS.KeyFile,
			CACertFile:         r.TLS.CACertFile,
			ServerName:         r.TLS.ServerName,
			InsecureSkipVerify: r.TLS.InsecureSkipVerify,
		}
	}

	if r.CircuitBreaker != nil {
		cb = &CircuitBreaker{
			Enabled:          r.CircuitBreaker.Enabled,
			MaxRequests:      r.CircuitBreaker.MaxRequests,
			Interval:         r.CircuitBreaker.Interval,
			Timeout:          r.CircuitBreaker.Timeout,
			FailureThreshold: r.CircuitBreaker.FailureThreshold,
			SuccessThreshold: r.CircuitBreaker.SuccessThreshold,
			HalfOpenMaxCalls: r.CircuitBreaker.HalfOpenMaxCalls,
		}
	}

	headers := make(map[string]string)
	for k, v := range r.Headers {
		headers[k] = v
	}

	return &RpcClient{
		ModuleName:     r.ModuleName,
		Enabled:        r.Enabled,
		Endpoints:      append([]string(nil), r.Endpoints...),
		Target:         r.Target,
		App:            r.App,
		Token:          r.Token,
		NonBlock:       r.NonBlock,
		Timeout:        r.Timeout,
		KeepaliveTime:  r.KeepaliveTime,
		DialTimeout:    r.DialTimeout,
		MaxRetries:     r.MaxRetries,
		RetryInterval:  r.RetryInterval,
		LoadBalance:    r.LoadBalance,
		Compression:    r.Compression,
		Headers:        headers,
		TLS:            tls,
		CircuitBreaker: cb,
	}
}

// Get 返回 RpcClient 配置的所有字段
func (r *RpcClient) Get() interface{} {
	return r
}

// Set 更新 RpcClient 配置的字段
func (r *RpcClient) Set(data interface{}) {
	if configData, ok := data.(*RpcClient); ok {
		*r = *configData
	}
}

// Validate 验证 RpcClient 配置的有效性
func (r *RpcClient) Validate() error {
	// 设置默认值
	if r.ModuleName == "" {
		r.ModuleName = "rpc-client"
	}
	if r.Timeout <= 0 {
		r.Timeout = 5000 // 5秒
	}
	if r.DialTimeout <= 0 {
		r.DialTimeout = 3
	}
	if r.KeepaliveTime <= 0 {
		r.KeepaliveTime = 30
	}
	if r.MaxRetries <= 0 {
		r.MaxRetries = 3
	}
	if r.RetryInterval <= 0 {
		r.RetryInterval = 1000
	}
	if r.LoadBalance == "" {
		r.LoadBalance = "round_robin"
	}
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}

	return internal.ValidateStruct(r)
}

// DefaultRpcClient 返回默认RpcClient配置
func DefaultRpcClient() RpcClient {
	return RpcClient{
		ModuleName:    "rpc-client",
		Enabled:       false,
		Endpoints:     []string{"localhost:8080"},
		Target:        "demo.service",
		App:           "demo-app",
		Token:         "demo_token",
		NonBlock:      false,
		Timeout:       5000,
		KeepaliveTime: 30,
		DialTimeout:   3,
		MaxRetries:    3,
		RetryInterval: 1000,
		LoadBalance:   "round_robin",
		Compression:   "gzip",
		Headers:       make(map[string]string),
		TLS: &TLSConfig{
			Enabled:            false,
			InsecureSkipVerify: true,
		},
		CircuitBreaker: &CircuitBreaker{
			Enabled:          false,
			MaxRequests:      100,
			Interval:         60,
			Timeout:          10,
			FailureThreshold: 5,
			SuccessThreshold: 2,
			HalfOpenMaxCalls: 3,
		},
	}
}

// Default 返回默认RpcClient配置的指针，支持链式调用
func Default() *RpcClient {
	config := DefaultRpcClient()
	return &config
}

// WithModuleName 设置模块名称
func (r *RpcClient) WithModuleName(moduleName string) *RpcClient {
	r.ModuleName = moduleName
	return r
}

// WithEndpoints 设置服务端点
func (r *RpcClient) WithEndpoints(endpoints []string) *RpcClient {
	r.Endpoints = endpoints
	return r
}

// WithTarget 设置目标服务
func (r *RpcClient) WithTarget(target string) *RpcClient {
	r.Target = target
	return r
}

// WithApp 设置应用名称
func (r *RpcClient) WithApp(app string) *RpcClient {
	r.App = app
	return r
}

// WithAuth 设置认证Token
func (r *RpcClient) WithAuth(token string) *RpcClient {
	r.Token = token
	return r
}

// WithTimeout 设置超时时间
func (r *RpcClient) WithTimeout(timeout int) *RpcClient {
	r.Timeout = timeout
	return r
}

// WithRetry 设置重试配置
func (r *RpcClient) WithRetry(maxRetries, retryInterval int) *RpcClient {
	r.MaxRetries = maxRetries
	r.RetryInterval = retryInterval
	return r
}

// WithLoadBalance 设置负载均衡策略
func (r *RpcClient) WithLoadBalance(strategy string) *RpcClient {
	r.LoadBalance = strategy
	return r
}

// WithTLS 设置TLS配置
func (r *RpcClient) WithTLS(enabled bool, certFile, keyFile, caCertFile string) *RpcClient {
	if r.TLS == nil {
		r.TLS = &TLSConfig{}
	}
	r.TLS.Enabled = enabled
	r.TLS.CertFile = certFile
	r.TLS.KeyFile = keyFile
	r.TLS.CACertFile = caCertFile
	return r
}

// EnableCircuitBreaker 启用熔断器
func (r *RpcClient) EnableCircuitBreaker() *RpcClient {
	if r.CircuitBreaker == nil {
		r.CircuitBreaker = &CircuitBreaker{}
	}
	r.CircuitBreaker.Enabled = true
	return r
}

// Enable 启用RPC客户端
func (r *RpcClient) Enable() *RpcClient {
	r.Enabled = true
	return r
}

// Disable 禁用RPC客户端
func (r *RpcClient) Disable() *RpcClient {
	r.Enabled = false
	return r
}
