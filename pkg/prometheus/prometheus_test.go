/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\prometheus\prometheus_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package prometheus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrometheus_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "prometheus", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "/metrics", config.Path)
	assert.Equal(t, 9090, config.Port)
	assert.Equal(t, "http://localhost:9090", config.Endpoint)
	assert.NotNil(t, config.PushGateway)
	assert.False(t, config.PushGateway.Enabled)
	assert.Equal(t, "go-rpc-gateway", config.PushGateway.JobName)
	assert.NotNil(t, config.Scraping)
	assert.Equal(t, "15s", config.Scraping.Interval)
	assert.Equal(t, "10s", config.Scraping.Timeout)
	assert.Equal(t, "/metrics", config.Scraping.MetricsPath)
}

func TestPrometheus_WithModuleName(t *testing.T) {
	config := Default().WithModuleName("custom-prometheus")
	assert.Equal(t, "custom-prometheus", config.ModuleName)
}

func TestPrometheus_WithEnabled(t *testing.T) {
	config := Default().WithEnabled(true)
	assert.True(t, config.Enabled)
}

func TestPrometheus_WithPath(t *testing.T) {
	config := Default().WithPath("/prometheus/metrics")
	assert.Equal(t, "/prometheus/metrics", config.Path)
}

func TestPrometheus_WithPort(t *testing.T) {
	config := Default().WithPort(8080)
	assert.Equal(t, 8080, config.Port)
}

func TestPrometheus_WithEndpoint(t *testing.T) {
	config := Default().WithEndpoint("http://prometheus:9090")
	assert.Equal(t, "http://prometheus:9090", config.Endpoint)
}

func TestPrometheus_WithPushGateway(t *testing.T) {
	config := Default().WithPushGateway(true, "http://pushgateway:9091", "my-job")
	assert.True(t, config.PushGateway.Enabled)
	assert.Equal(t, "http://pushgateway:9091", config.PushGateway.Endpoint)
	assert.Equal(t, "my-job", config.PushGateway.JobName)
}

func TestPrometheus_WithScraping(t *testing.T) {
	config := Default().WithScraping("30s", "20s", "/custom/metrics")
	assert.Equal(t, "30s", config.Scraping.Interval)
	assert.Equal(t, "20s", config.Scraping.Timeout)
	assert.Equal(t, "/custom/metrics", config.Scraping.MetricsPath)
}

func TestPrometheus_EnablePushGateway(t *testing.T) {
	config := Default().EnablePushGateway()
	assert.True(t, config.PushGateway.Enabled)
}

func TestPrometheus_Enable(t *testing.T) {
	config := Default().Enable()
	assert.True(t, config.Enabled)
}

func TestPrometheus_Disable(t *testing.T) {
	config := Default().Enable().Disable()
	assert.False(t, config.Enabled)
}

func TestPrometheus_IsEnabled(t *testing.T) {
	config := Default()
	assert.False(t, config.IsEnabled())
	config.Enable()
	assert.True(t, config.IsEnabled())
}

func TestPrometheus_Clone(t *testing.T) {
	original := Default().
		WithModuleName("test-prometheus").
		Enable().
		WithPath("/test/metrics").
		WithPort(8888).
		WithEndpoint("http://test:9090").
		WithPushGateway(true, "http://test-push:9091", "test-job").
		WithScraping("20s", "15s", "/test/path")

	cloned := original.Clone().(*Prometheus)

	// 验证值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.Path, cloned.Path)
	assert.Equal(t, original.Port, cloned.Port)
	assert.Equal(t, original.Endpoint, cloned.Endpoint)

	// 验证嵌套结构
	assert.Equal(t, original.PushGateway.Enabled, cloned.PushGateway.Enabled)
	assert.Equal(t, original.PushGateway.Endpoint, cloned.PushGateway.Endpoint)
	assert.Equal(t, original.PushGateway.JobName, cloned.PushGateway.JobName)
	assert.Equal(t, original.Scraping.Interval, cloned.Scraping.Interval)
	assert.Equal(t, original.Scraping.Timeout, cloned.Scraping.Timeout)
	assert.Equal(t, original.Scraping.MetricsPath, cloned.Scraping.MetricsPath)

	// 验证独立性
	cloned.ModuleName = "modified-prometheus"
	cloned.PushGateway.JobName = "modified-job"
	assert.NotEqual(t, original.ModuleName, cloned.ModuleName)
	assert.NotEqual(t, original.PushGateway.JobName, cloned.PushGateway.JobName)
}

func TestPrometheus_Get(t *testing.T) {
	config := Default().WithEnabled(true)
	got := config.Get()
	assert.NotNil(t, got)
	prometheusConfig, ok := got.(*Prometheus)
	assert.True(t, ok)
	assert.True(t, prometheusConfig.Enabled)
}

func TestPrometheus_Set(t *testing.T) {
	config := Default()
	newConfig := &Prometheus{
		ModuleName: "new-prometheus",
		Enabled:    true,
		Path:       "/new/metrics",
		Port:       7070,
		Endpoint:   "http://new:9090",
	}

	config.Set(newConfig)
	assert.Equal(t, "new-prometheus", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "/new/metrics", config.Path)
	assert.Equal(t, 7070, config.Port)
	assert.Equal(t, "http://new:9090", config.Endpoint)
}

func TestPrometheus_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestPrometheus_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("chain-prometheus").
		Enable().
		WithPath("/chain/metrics").
		WithPort(9999).
		WithEndpoint("http://chain:9090").
		WithPushGateway(true, "http://chain-push:9091", "chain-job").
		WithScraping("25s", "18s", "/chain/path").
		EnablePushGateway()

	assert.Equal(t, "chain-prometheus", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "/chain/metrics", config.Path)
	assert.Equal(t, 9999, config.Port)
	assert.Equal(t, "http://chain:9090", config.Endpoint)
	assert.True(t, config.PushGateway.Enabled)
	assert.Equal(t, "http://chain-push:9091", config.PushGateway.Endpoint)
	assert.Equal(t, "chain-job", config.PushGateway.JobName)
	assert.Equal(t, "25s", config.Scraping.Interval)
	assert.Equal(t, "18s", config.Scraping.Timeout)
	assert.Equal(t, "/chain/path", config.Scraping.MetricsPath)
}

func TestPrometheus_NilSubConfigs(t *testing.T) {
	config := Default()
	config.PushGateway = nil
	config.Scraping = nil

	// These should not panic
	config.WithPushGateway(true, "http://test:9091", "test")
	config.WithScraping("15s", "10s", "/metrics")
	config.EnablePushGateway()

	assert.NotNil(t, config.PushGateway)
	assert.NotNil(t, config.Scraping)
}
