/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 18:00:00
 * @FilePath: \go-config\pkg\metrics\metrics.go
 * @Description: 指标中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package metrics

import "github.com/kamalyes/go-config/internal"

// Metrics 指标中间件配置
type Metrics struct {
	ModuleName   string    `mapstructure:"module_name" yaml:"module-name" json:"module_name"`       // 模块名称
	Enabled      bool      `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                   // 是否启用指标
	Path         string    `mapstructure:"path" yaml:"path" json:"path"`                            // 指标路径
	Subsystem    string    `mapstructure:"subsystem" yaml:"subsystem" json:"subsystem"`             // 子系统名称
	SkipPaths    []string  `mapstructure:"skip_paths" yaml:"skip-paths" json:"skip_paths"`          // 跳过的路径
	Buckets      []float64 `mapstructure:"buckets" yaml:"buckets" json:"buckets"`                   // 直方图桶
	RequestCount bool      `mapstructure:"request_count" yaml:"request-count" json:"request_count"` // 请求计数
	Duration     bool      `mapstructure:"duration" yaml:"duration" json:"duration"`                // 请求时长
	RequestSize  bool      `mapstructure:"request_size" yaml:"request-size" json:"request_size"`    // 请求大小
	ResponseSize bool      `mapstructure:"response_size" yaml:"response-size" json:"response_size"` // 响应大小
}

// Default 创建默认指标配置
func Default() *Metrics {
	return &Metrics{
		ModuleName:   "metrics",
		Enabled:      false,
		Path:         "/metrics",
		Subsystem:    "gateway",
		SkipPaths:    []string{"/health"},
		Buckets:      []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0, 10.0},
		RequestCount: true,
		Duration:     true,
		RequestSize:  true,
		ResponseSize: true,
	}
}

// Get 返回配置接口
func (m *Metrics) Get() interface{} {
	return m
}

// Set 设置配置数据
func (m *Metrics) Set(data interface{}) {
	if cfg, ok := data.(*Metrics); ok {
		*m = *cfg
	}
}

// Clone 返回配置的副本
func (m *Metrics) Clone() internal.Configurable {
	clone := &Metrics{
		ModuleName:   m.ModuleName,
		Enabled:      m.Enabled,
		Path:         m.Path,
		Subsystem:    m.Subsystem,
		RequestCount: m.RequestCount,
		Duration:     m.Duration,
		RequestSize:  m.RequestSize,
		ResponseSize: m.ResponseSize,
	}
	clone.SkipPaths = append([]string(nil), m.SkipPaths...)
	clone.Buckets = append([]float64(nil), m.Buckets...)
	return clone
}

// Validate 验证配置
func (m *Metrics) Validate() error {
	return internal.ValidateStruct(m)
}

// WithPath 设置指标路径
func (m *Metrics) WithPath(path string) *Metrics {
	m.Path = path
	return m
}

// WithSubsystem 设置子系统名称
func (m *Metrics) WithSubsystem(subsystem string) *Metrics {
	m.Subsystem = subsystem
	return m
}

// WithSkipPaths 设置跳过的路径
func (m *Metrics) WithSkipPaths(skipPaths []string) *Metrics {
	m.SkipPaths = skipPaths
	return m
}

// WithBuckets 设置直方图桶
func (m *Metrics) WithBuckets(buckets []float64) *Metrics {
	m.Buckets = buckets
	return m
}

// Enable 启用指标中间件
func (m *Metrics) Enable() *Metrics {
	m.Enabled = true
	return m
}

// Disable 禁用指标中间件
func (m *Metrics) Disable() *Metrics {
	m.Enabled = false
	return m
}

// IsEnabled 检查是否启用
func (m *Metrics) IsEnabled() bool {
	return m.Enabled
}
