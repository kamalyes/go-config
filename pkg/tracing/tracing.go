/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 18:00:00
 * @FilePath: \go-config\pkg\tracing\tracing.go
 * @Description: 追踪中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package tracing

import "github.com/kamalyes/go-config/internal"

// Tracing 追踪中间件配置
type Tracing struct {
	ModuleName         string            `mapstructure:"module_name" yaml:"module-name" json:"module_name"`                         // 模块名称
	Enabled            bool              `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                     // 是否启用追踪
	ServiceName        string            `mapstructure:"service_name" yaml:"service-name" json:"service_name"`                      // 服务名称
	ServiceVersion     string            `mapstructure:"service_version" yaml:"service-version" json:"service_version"`             // 服务版本
	Environment        string            `mapstructure:"environment" yaml:"environment" json:"environment"`                         // 环境
	Endpoint           string            `mapstructure:"endpoint" yaml:"endpoint" json:"endpoint"`                                  // 追踪端点
	ExporterType       string            `mapstructure:"exporter_type" yaml:"exporter-type" json:"exporter_type"`                   // 导出器类型 (zipkin, otlp, console, noop)
	ExporterEndpoint   string            `mapstructure:"exporter_endpoint" yaml:"exporter-endpoint" json:"exporter_endpoint"`       // 导出器端点
	SampleRate         float64           `mapstructure:"sample_rate" yaml:"sample-rate" json:"sample_rate"`                         // 采样率
	SamplerType        string            `mapstructure:"sampler_type" yaml:"sampler-type" json:"sampler_type"`                      // 采样器类型 (always, never, probability, parentBased)
	SamplerProbability float64           `mapstructure:"sampler_probability" yaml:"sampler-probability" json:"sampler_probability"` // 采样概率
	SamplerRate        float64           `mapstructure:"sampler_rate" yaml:"sampler-rate" json:"sampler_rate"`                      // 采样速率
	Headers            []string          `mapstructure:"headers" yaml:"headers" json:"headers"`                                     // 需要记录的头部
	Attributes         map[string]string `mapstructure:"attributes" yaml:"attributes" json:"attributes"`                            // 自定义属性
}

// Default 创建默认追踪配置
func Default() *Tracing {
	return &Tracing{
		ModuleName:         "tracing",
		Enabled:            false,
		ServiceName:        "go-rpc-gateway",
		ServiceVersion:     "1.0.0",
		Environment:        "development",
		Endpoint:           "http://localhost:9411/api/v2/spans",
		ExporterType:       "zipkin",
		ExporterEndpoint:   "http://localhost:9411/api/v2/spans",
		SampleRate:         0.1,
		SamplerType:        "probability",
		SamplerProbability: 0.1,
		SamplerRate:        0.1,
		Headers:            []string{"Authorization", "User-Agent"},
		Attributes:         make(map[string]string),
	}
}

// Get 返回配置接口
func (t *Tracing) Get() interface{} {
	return t
}

// Set 设置配置数据
func (t *Tracing) Set(data interface{}) {
	if cfg, ok := data.(*Tracing); ok {
		*t = *cfg
	}
}

// Clone 返回配置的副本
func (t *Tracing) Clone() internal.Configurable {
	clone := &Tracing{
		ModuleName:         t.ModuleName,
		Enabled:            t.Enabled,
		ServiceName:        t.ServiceName,
		ServiceVersion:     t.ServiceVersion,
		Environment:        t.Environment,
		Endpoint:           t.Endpoint,
		ExporterType:       t.ExporterType,
		ExporterEndpoint:   t.ExporterEndpoint,
		SampleRate:         t.SampleRate,
		SamplerType:        t.SamplerType,
		SamplerProbability: t.SamplerProbability,
		SamplerRate:        t.SamplerRate,
	}
	clone.Headers = append([]string(nil), t.Headers...)
	clone.Attributes = make(map[string]string)
	for k, v := range t.Attributes {
		clone.Attributes[k] = v
	}
	return clone
}

// Validate 验证配置
func (t *Tracing) Validate() error {
	return internal.ValidateStruct(t)
}

// WithServiceName 设置服务名称
func (t *Tracing) WithServiceName(serviceName string) *Tracing {
	t.ServiceName = serviceName
	return t
}

// WithEndpoint 设置追踪端点
func (t *Tracing) WithEndpoint(endpoint string) *Tracing {
	t.Endpoint = endpoint
	return t
}

// WithSampleRate 设置采样率
func (t *Tracing) WithSampleRate(sampleRate float64) *Tracing {
	t.SampleRate = sampleRate
	return t
}

// WithHeaders 设置需要记录的头部
func (t *Tracing) WithHeaders(headers []string) *Tracing {
	t.Headers = headers
	return t
}

// Enable 启用追踪中间件
func (t *Tracing) Enable() *Tracing {
	t.Enabled = true
	return t
}

// Disable 禁用追踪中间件
func (t *Tracing) Disable() *Tracing {
	t.Enabled = false
	return t
}

// IsEnabled 检查是否启用
func (t *Tracing) IsEnabled() bool {
	return t.Enabled
}
