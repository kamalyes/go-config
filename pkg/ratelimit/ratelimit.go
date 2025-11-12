/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 15:00:00
 * @FilePath: \go-config\pkg\ratelimit\ratelimit.go
 * @Description: 限流中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package ratelimit

import (
	"time"

	"github.com/kamalyes/go-config/internal"
)

// RateLimit 限流中间件配置
type RateLimit struct {
	ModuleName        string        `mapstructure:"module_name" yaml:"module-name" json:"module_name"`                      // 模块名称
	Enabled           bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                 // 是否启用限流
	RequestsPerSecond int           `mapstructure:"requests_per_second" yaml:"requests-per-second" json:"requests_per_second"` // 每秒允许的请求数
	BurstSize         int           `mapstructure:"burst_size" yaml:"burst-size" json:"burst_size"`                         // 突发请求数
	CleanupInterval   time.Duration `mapstructure:"cleanup_interval" yaml:"cleanup-interval" json:"cleanup_interval"`      // 清理间隔
	WindowSize        time.Duration `mapstructure:"window_size" yaml:"window-size" json:"window_size"`                     // 时间窗口大小
	KeyPrefix         string        `mapstructure:"key_prefix" yaml:"key-prefix" json:"key_prefix"`                        // Redis key 前缀
}

// Default 创建默认限流配置
func Default() *RateLimit {
	return &RateLimit{
		ModuleName:        "ratelimit",
		Enabled:           true,
		RequestsPerSecond: 100,
		BurstSize:         200,
		CleanupInterval:   time.Minute,
		WindowSize:        time.Minute,
		KeyPrefix:         "rate_limit",
	}
}

// Get 返回配置接口
func (r *RateLimit) Get() interface{} {
	return r
}

// Set 设置配置数据
func (r *RateLimit) Set(data interface{}) {
	if cfg, ok := data.(*RateLimit); ok {
		*r = *cfg
	}
}

// Clone 返回配置的副本
func (r *RateLimit) Clone() internal.Configurable {
	return &RateLimit{
		ModuleName:        r.ModuleName,
		Enabled:           r.Enabled,
		RequestsPerSecond: r.RequestsPerSecond,
		BurstSize:         r.BurstSize,
		CleanupInterval:   r.CleanupInterval,
		WindowSize:        r.WindowSize,
		KeyPrefix:         r.KeyPrefix,
	}
}

// Validate 验证配置
func (r *RateLimit) Validate() error {
	return internal.ValidateStruct(r)
}

// WithRequestsPerSecond 设置每秒请求数
func (r *RateLimit) WithRequestsPerSecond(requestsPerSecond int) *RateLimit {
	r.RequestsPerSecond = requestsPerSecond
	return r
}

// WithBurstSize 设置突发大小
func (r *RateLimit) WithBurstSize(burstSize int) *RateLimit {
	r.BurstSize = burstSize
	return r
}

// WithCleanupInterval 设置清理间隔
func (r *RateLimit) WithCleanupInterval(cleanupInterval time.Duration) *RateLimit {
	r.CleanupInterval = cleanupInterval
	return r
}

// WithWindowSize 设置时间窗口大小
func (r *RateLimit) WithWindowSize(windowSize time.Duration) *RateLimit {
	r.WindowSize = windowSize
	return r
}

// WithKeyPrefix 设置Redis key前缀
func (r *RateLimit) WithKeyPrefix(keyPrefix string) *RateLimit {
	r.KeyPrefix = keyPrefix
	return r
}

// Enable 启用限流中间件
func (r *RateLimit) Enable() *RateLimit {
	r.Enabled = true
	return r
}

// Disable 禁用限流中间件
func (r *RateLimit) Disable() *RateLimit {
	r.Enabled = false
	return r
}

// IsEnabled 检查是否启用
func (r *RateLimit) IsEnabled() bool {
	return r.Enabled
}