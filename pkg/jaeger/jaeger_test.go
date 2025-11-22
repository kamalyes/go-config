/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 16:50:55
 * @FilePath: \go-config\pkg\jaeger\jaeger_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package jaeger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJaeger_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "jaeger", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "http://localhost:14268/api/traces", config.Endpoint)
	assert.Equal(t, "go-rpc-gateway", config.ServiceName)
	assert.Equal(t, 0.1, config.SampleRate)
	assert.NotNil(t, config.Agent)
	assert.Equal(t, "localhost", config.Agent.Host)
	assert.Equal(t, 6832, config.Agent.Port)
	assert.NotNil(t, config.Collector)
	assert.NotNil(t, config.Sampling)
	assert.Equal(t, "probabilistic", config.Sampling.Type)
	assert.Equal(t, 0.1, config.Sampling.Param)
	assert.Equal(t, 100, config.Sampling.MaxTracesPerSecond)
}

func TestJaeger_WithModuleName(t *testing.T) {
	config := Default().WithModuleName("custom-jaeger")
	assert.Equal(t, "custom-jaeger", config.ModuleName)
}

func TestJaeger_WithEnabled(t *testing.T) {
	config := Default().WithEnabled(true)
	assert.True(t, config.Enabled)
}

func TestJaeger_WithEndpoint(t *testing.T) {
	config := Default().WithEndpoint("http://jaeger:14268/api/traces")
	assert.Equal(t, "http://jaeger:14268/api/traces", config.Endpoint)
}

func TestJaeger_WithServiceName(t *testing.T) {
	config := Default().WithServiceName("my-service")
	assert.Equal(t, "my-service", config.ServiceName)
}

func TestJaeger_WithSampleRate(t *testing.T) {
	config := Default().WithSampleRate(0.5)
	assert.Equal(t, 0.5, config.SampleRate)
}

func TestJaeger_WithAgent(t *testing.T) {
	config := Default().WithAgent("jaeger-agent", 6831)
	assert.Equal(t, "jaeger-agent", config.Agent.Host)
	assert.Equal(t, 6831, config.Agent.Port)
}

func TestJaeger_WithCollector(t *testing.T) {
	config := Default().WithCollector("http://collector:14268", "user", "pass")
	assert.Equal(t, "http://collector:14268", config.Collector.Endpoint)
	assert.Equal(t, "user", config.Collector.Username)
	assert.Equal(t, "pass", config.Collector.Password)
}

func TestJaeger_WithSampling(t *testing.T) {
	config := Default().WithSampling("const", 1.0, 200)
	assert.Equal(t, "const", config.Sampling.Type)
	assert.Equal(t, 1.0, config.Sampling.Param)
	assert.Equal(t, 200, config.Sampling.MaxTracesPerSecond)
}

func TestJaeger_AddOperationSampling(t *testing.T) {
	config := Default().AddOperationSampling("GET /api/users", 50, 0.5)
	assert.Len(t, config.Sampling.OperationSampling, 1)
	assert.Equal(t, "GET /api/users", config.Sampling.OperationSampling[0].Operation)
	assert.Equal(t, 50, config.Sampling.OperationSampling[0].MaxTracesPerSecond)
	assert.Equal(t, 0.5, config.Sampling.OperationSampling[0].ProbabilisticSampling)
}

func TestJaeger_AddTag(t *testing.T) {
	config := Default().
		AddTag("environment", "production").
		AddTag("version", "1.0.0")
	assert.Len(t, config.Tags, 2)
	assert.Equal(t, "production", config.Tags["environment"])
	assert.Equal(t, "1.0.0", config.Tags["version"])
}

func TestJaeger_Enable(t *testing.T) {
	config := Default().Enable()
	assert.True(t, config.Enabled)
}

func TestJaeger_Disable(t *testing.T) {
	config := Default().Enable().Disable()
	assert.False(t, config.Enabled)
}

func TestJaeger_IsEnabled(t *testing.T) {
	config := Default()
	assert.False(t, config.IsEnabled())
	config.Enable()
	assert.True(t, config.IsEnabled())
}

func TestJaeger_Clone(t *testing.T) {
	original := Default().
		WithModuleName("test-jaeger").
		WithEnabled(true).
		WithEndpoint("http://test:14268").
		WithServiceName("test-service").
		WithSampleRate(0.8).
		AddOperationSampling("test-op", 100, 0.9).
		AddTag("env", "test")

	cloned := original.Clone().(*Jaeger)

	// 验证值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.Endpoint, cloned.Endpoint)
	assert.Equal(t, original.ServiceName, cloned.ServiceName)
	assert.Equal(t, original.SampleRate, cloned.SampleRate)

	// 验证嵌套结构
	assert.Equal(t, original.Agent.Host, cloned.Agent.Host)
	assert.Equal(t, original.Agent.Port, cloned.Agent.Port)

	// 验证切片独立性
	cloned.Sampling.OperationSampling = append(cloned.Sampling.OperationSampling, OperationSampling{
		Operation: "new-op",
	})
	assert.NotEqual(t, len(original.Sampling.OperationSampling), len(cloned.Sampling.OperationSampling))

	// 验证map独立性
	cloned.Tags["new-key"] = "new-value"
	assert.NotEqual(t, len(original.Tags), len(cloned.Tags))
}

func TestJaeger_Get(t *testing.T) {
	config := Default().WithEnabled(true)
	got := config.Get()
	assert.NotNil(t, got)
	jaegerConfig, ok := got.(*Jaeger)
	assert.True(t, ok)
	assert.True(t, jaegerConfig.Enabled)
}

func TestJaeger_Set(t *testing.T) {
	config := Default()
	newConfig := &Jaeger{
		ModuleName:  "new-jaeger",
		Enabled:     true,
		Endpoint:    "http://new:14268",
		ServiceName: "new-service",
		SampleRate:  0.7,
	}

	config.Set(newConfig)
	assert.Equal(t, "new-jaeger", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "http://new:14268", config.Endpoint)
	assert.Equal(t, "new-service", config.ServiceName)
	assert.Equal(t, 0.7, config.SampleRate)
}

func TestJaeger_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestJaeger_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("chain-jaeger").
		Enable().
		WithEndpoint("http://chain:14268").
		WithServiceName("chain-service").
		WithSampleRate(0.6).
		WithAgent("chain-agent", 6831).
		WithCollector("http://chain-collector:14268", "user", "pass").
		WithSampling("const", 1.0, 300).
		AddOperationSampling("chain-op", 150, 0.8).
		AddTag("env", "chain")

	assert.Equal(t, "chain-jaeger", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "http://chain:14268", config.Endpoint)
	assert.Equal(t, "chain-service", config.ServiceName)
	assert.Equal(t, 0.6, config.SampleRate)
	assert.Equal(t, "chain-agent", config.Agent.Host)
	assert.Equal(t, 6831, config.Agent.Port)
	assert.Equal(t, "const", config.Sampling.Type)
	assert.Len(t, config.Sampling.OperationSampling, 1)
	assert.Len(t, config.Tags, 1)
}
