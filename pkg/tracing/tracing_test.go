/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\tracing\tracing_test.go
 * @Description: 追踪中间件配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package tracing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTracing_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "tracing", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "go-rpc-gateway", config.ServiceName)
	assert.Equal(t, "1.0.0", config.ServiceVersion)
	assert.Equal(t, "development", config.Environment)
	assert.Equal(t, "http://localhost:9411/api/v2/spans", config.Endpoint)
	assert.Equal(t, "zipkin", config.ExporterType)
	assert.Equal(t, "http://localhost:9411/api/v2/spans", config.ExporterEndpoint)
	assert.Equal(t, 0.1, config.SampleRate)
	assert.Equal(t, "probability", config.SamplerType)
	assert.Equal(t, 0.1, config.SamplerProbability)
	assert.Equal(t, 0.1, config.SamplerRate)
	assert.NotNil(t, config.Headers)
	assert.NotNil(t, config.Attributes)
}

func TestTracing_WithServiceName(t *testing.T) {
	config := Default()
	result := config.WithServiceName("my-service")
	assert.Equal(t, "my-service", result.ServiceName)
	assert.Equal(t, config, result)
}

func TestTracing_WithEndpoint(t *testing.T) {
	config := Default()
	result := config.WithEndpoint("http://jaeger:14268/api/traces")
	assert.Equal(t, "http://jaeger:14268/api/traces", result.Endpoint)
	assert.Equal(t, config, result)
}

func TestTracing_WithSampleRate(t *testing.T) {
	config := Default()
	result := config.WithSampleRate(0.5)
	assert.Equal(t, 0.5, result.SampleRate)
	assert.Equal(t, config, result)
}

func TestTracing_WithHeaders(t *testing.T) {
	config := Default()
	headers := []string{"X-Custom-Header", "X-Request-ID"}
	result := config.WithHeaders(headers)
	assert.Equal(t, headers, result.Headers)
	assert.Equal(t, config, result)
}

func TestTracing_Enable(t *testing.T) {
	config := Default()
	result := config.Enable()
	assert.True(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestTracing_Disable(t *testing.T) {
	config := Default()
	config.Enabled = true
	result := config.Disable()
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestTracing_IsEnabled(t *testing.T) {
	config := Default()
	assert.False(t, config.IsEnabled())

	config.Enabled = true
	assert.True(t, config.IsEnabled())
}

func TestTracing_Clone(t *testing.T) {
	config := Default()
	config.WithServiceName("test-service").
		WithSampleRate(0.8).
		WithHeaders([]string{"X-Test"})
	config.Attributes["key1"] = "value1"
	config.Attributes["key2"] = "value2"

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*Tracing)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.ServiceName, clonedConfig.ServiceName)
	assert.Equal(t, config.SampleRate, clonedConfig.SampleRate)

	// 验证深拷贝 - 切片
	clonedConfig.Headers = append(clonedConfig.Headers, "X-Extra")
	assert.NotEqual(t, len(config.Headers), len(clonedConfig.Headers))

	// 验证深拷贝 - map
	clonedConfig.Attributes["key3"] = "value3"
	_, exists := config.Attributes["key3"]
	assert.False(t, exists)
}

func TestTracing_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestTracing_Set(t *testing.T) {
	config := Default()
	newConfig := &Tracing{
		ModuleName:         "new-tracing",
		Enabled:            true,
		ServiceName:        "new-service",
		ServiceVersion:     "2.0.0",
		Environment:        "production",
		Endpoint:           "http://new-endpoint.com",
		ExporterType:       "otlp",
		ExporterEndpoint:   "http://otlp:4318",
		SampleRate:         0.2,
		SamplerType:        "always",
		SamplerProbability: 1.0,
		SamplerRate:        1.0,
		Headers:            []string{"X-New-Header"},
		Attributes:         map[string]string{"env": "prod"},
	}

	config.Set(newConfig)
	assert.Equal(t, "new-tracing", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "new-service", config.ServiceName)
	assert.Equal(t, "2.0.0", config.ServiceVersion)
	assert.Equal(t, "production", config.Environment)
	assert.Equal(t, "http://new-endpoint.com", config.Endpoint)
	assert.Equal(t, "otlp", config.ExporterType)
	assert.Equal(t, 0.2, config.SampleRate)
}

func TestTracing_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestTracing_ChainedCalls(t *testing.T) {
	config := Default().
		WithServiceName("my-microservice").
		WithEndpoint("http://jaeger:14268/api/traces").
		WithSampleRate(0.75).
		WithHeaders([]string{"X-Trace-ID", "X-Span-ID", "X-Parent-ID"}).
		Enable()

	assert.Equal(t, "my-microservice", config.ServiceName)
	assert.Equal(t, "http://jaeger:14268/api/traces", config.Endpoint)
	assert.Equal(t, 0.75, config.SampleRate)
	assert.Contains(t, config.Headers, "X-Trace-ID")
	assert.Contains(t, config.Headers, "X-Span-ID")
	assert.Contains(t, config.Headers, "X-Parent-ID")
	assert.True(t, config.Enabled)
}
