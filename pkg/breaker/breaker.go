/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12
 * @Description: CircuitBreaker 断路器配置模块
 * 提供统一的断路器配置管理，支持全局和路径级别的配置
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package breaker

import (
	"time"

	"github.com/kamalyes/go-config/internal"
)

// CircuitBreaker 断路器配置
type CircuitBreaker struct {
	ModuleName          string   `mapstructure:"module_name" yaml:"module_name" json:"module_name"`                               // 模块名称
	Enabled             bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                           // 是否启用断路器
	FailureThreshold    int      `mapstructure:"failure_threshold" yaml:"failure-threshold" json:"failure_threshold"`             // 失败阈值
	SuccessThreshold    int      `mapstructure:"success_threshold" yaml:"success-threshold" json:"success_threshold"`             // 成功阈值
	Timeout             int64    `mapstructure:"timeout" yaml:"timeout" json:"timeout"`                                           // 熔断后恢复时间
	VolumeThreshold     int      `mapstructure:"volume_threshold" yaml:"volume-threshold" json:"volume_threshold"`                // 最小请求量阈值
	SlidingWindowSize   int      `mapstructure:"sliding_window_size" yaml:"sliding-window-size" json:"sliding_window_size"`       // 滑动窗口大小
	SlidingWindowBucket int64    `mapstructure:"sliding_window_bucket" yaml:"sliding-window-bucket" json:"sliding_window_bucket"` // 滑动窗口桶大小
	PreventionPaths     []string `mapstructure:"prevention_paths" yaml:"prevention-paths" json:"prevention_paths"`                // 需要保护的路径
	ExcludePaths        []string `mapstructure:"exclude_paths" yaml:"exclude-paths" json:"exclude_paths"`                         // 排除的路径
}

// WebSocketBreaker WebSocket 专用断路器配置
type WebSocketBreaker struct {
	ModuleName          string  `mapstructure:"module_name" yaml:"module_name" json:"module_name"`                               // 模块名称
	Enabled             bool    `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                           // 是否启用
	FailureThreshold    int     `mapstructure:"failure_threshold" yaml:"failure-threshold" json:"failure_threshold"`             // 失败阈值
	SuccessThreshold    int     `mapstructure:"success_threshold" yaml:"success-threshold" json:"success_threshold"`             // 成功阈值
	Timeout             int64   `mapstructure:"timeout" yaml:"timeout" json:"timeout"`                                           // 熔断恢复时间
	MaxRetries          int     `mapstructure:"max_retries" yaml:"max-retries" json:"max_retries"`                               // 最大重试次数
	RetryBackoffFactor  float64 `mapstructure:"retry_backoff_factor" yaml:"retry-backoff-factor" json:"retry_backoff_factor"`    // 重试退避因子
	HealthCheckInterval int64   `mapstructure:"health_check_interval" yaml:"health-check-interval" json:"health_check_interval"` // 健康检查间隔
	MessageQueueSize    int     `mapstructure:"message_queue_size" yaml:"message-queue-size" json:"message_queue_size"`          // 消息队列大小
}

// Default 创建默认断路器配置
func Default() *CircuitBreaker {
	return &CircuitBreaker{
		ModuleName:          "circuit_breaker",
		Enabled:             true,
		FailureThreshold:    5,
		SuccessThreshold:    2,
		Timeout:             int64(30 * time.Second),
		VolumeThreshold:     10,
		SlidingWindowSize:   100,
		SlidingWindowBucket: int64(1 * time.Second),
		PreventionPaths:     []string{"/api/"},
		ExcludePaths:        []string{"/health", "/metrics", "/debug"},
	}
}

// DefaultWebSocketBreaker 创建默认 WebSocket 断路器配置
func DefaultWebSocketBreaker() *WebSocketBreaker {
	return &WebSocketBreaker{
		ModuleName:          "websocket_breaker",
		Enabled:             true,
		FailureThreshold:    5,
		SuccessThreshold:    2,
		Timeout:             int64(30 * time.Second),
		MaxRetries:          3,
		RetryBackoffFactor:  2.0,
		HealthCheckInterval: int64(10 * time.Second),
		MessageQueueSize:    1000,
	}
}

// GetLabel 获取配置标签
func (cb *CircuitBreaker) GetLabel() string {
	return cb.ModuleName
}

// Init 初始化配置
func (cb *CircuitBreaker) Init() error {
	return internal.ValidateStruct(cb)
}

// Validate 验证配置
func (cb *CircuitBreaker) Validate() error {
	return internal.ValidateStruct(cb)
}

// GetLabel 获取配置标签
func (wb *WebSocketBreaker) GetLabel() string {
	return wb.ModuleName
}

// Init 初始化配置
func (wb *WebSocketBreaker) Init() error {
	return internal.ValidateStruct(wb)
}

// Validate 验证配置
func (wb *WebSocketBreaker) Validate() error {
	return internal.ValidateStruct(wb)
}
