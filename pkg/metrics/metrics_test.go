/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\metrics\metrics_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package metrics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMetrics_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "metrics", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "/metrics", config.Path)
	assert.Equal(t, "gateway", config.Namespace)
	assert.Equal(t, "http", config.Subsystem)
	assert.Equal(t, []string{"/health"}, config.SkipPaths)
	assert.True(t, config.RequestCount)
	assert.True(t, config.Duration)
	assert.True(t, config.RequestSize)
	assert.True(t, config.ResponseSize)
	assert.NotEmpty(t, config.Buckets)
}

func TestMetrics_WithPath(t *testing.T) {
	config := Default().WithPath("/prometheus/metrics")
	assert.Equal(t, "/prometheus/metrics", config.Path)
}

func TestMetrics_WithNamespace(t *testing.T) {
	config := Default().WithNamespace("myapp")
	assert.Equal(t, "myapp", config.Namespace)
}

func TestMetrics_WithSubsystem(t *testing.T) {
	config := Default().WithSubsystem("api")
	assert.Equal(t, "api", config.Subsystem)
}

func TestMetrics_WithSkipPaths(t *testing.T) {
	skipPaths := []string{"/health", "/ready", "/alive"}
	config := Default().WithSkipPaths(skipPaths)
	assert.Equal(t, skipPaths, config.SkipPaths)
}

func TestMetrics_WithBuckets(t *testing.T) {
	buckets := []float64{0.1, 0.5, 1.0, 5.0, 10.0}
	config := Default().WithBuckets(buckets)
	assert.Equal(t, buckets, config.Buckets)
}

func TestMetrics_Enable(t *testing.T) {
	config := Default().Enable()
	assert.True(t, config.Enabled)
}

func TestMetrics_Disable(t *testing.T) {
	config := Default().Enable().Disable()
	assert.False(t, config.Enabled)
}

func TestMetrics_IsEnabled(t *testing.T) {
	config := Default()
	assert.False(t, config.IsEnabled())
	config.Enable()
	assert.True(t, config.IsEnabled())
}

func TestMetrics_Clone(t *testing.T) {
	original := Default().
		Enable().
		WithPath("/custom/metrics").
		WithNamespace("custom").
		WithSubsystem("service").
		WithSkipPaths([]string{"/skip1", "/skip2"}).
		WithBuckets([]float64{0.5, 1.0, 2.0})

	cloned := original.Clone().(*Metrics)

	// 验证值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.Path, cloned.Path)
	assert.Equal(t, original.Namespace, cloned.Namespace)
	assert.Equal(t, original.Subsystem, cloned.Subsystem)
	assert.Equal(t, original.RequestCount, cloned.RequestCount)
	assert.Equal(t, original.Duration, cloned.Duration)
	assert.Equal(t, original.RequestSize, cloned.RequestSize)
	assert.Equal(t, original.ResponseSize, cloned.ResponseSize)

	// 验证切片独立性
	cloned.SkipPaths = append(cloned.SkipPaths, "/skip3")
	assert.NotEqual(t, len(original.SkipPaths), len(cloned.SkipPaths))

	cloned.Buckets = append(cloned.Buckets, 5.0)
	assert.NotEqual(t, len(original.Buckets), len(cloned.Buckets))
}

func TestMetrics_Get(t *testing.T) {
	config := Default().WithNamespace("test")
	got := config.Get()
	assert.NotNil(t, got)
	metricsConfig, ok := got.(*Metrics)
	assert.True(t, ok)
	assert.Equal(t, "test", metricsConfig.Namespace)
}

func TestMetrics_Set(t *testing.T) {
	config := Default()
	newConfig := &Metrics{
		ModuleName:   "new-metrics",
		Enabled:      true,
		Path:         "/new/metrics",
		Namespace:    "new",
		Subsystem:    "new-sub",
		SkipPaths:    []string{"/new/skip"},
		Buckets:      []float64{1.0, 2.0},
		RequestCount: false,
		Duration:     false,
		RequestSize:  false,
		ResponseSize: false,
	}

	config.Set(newConfig)
	assert.Equal(t, "new-metrics", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "/new/metrics", config.Path)
	assert.Equal(t, "new", config.Namespace)
	assert.Equal(t, "new-sub", config.Subsystem)
	assert.False(t, config.RequestCount)
	assert.False(t, config.Duration)
}

func TestMetrics_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestMetrics_ChainedCalls(t *testing.T) {
	config := Default().
		Enable().
		WithPath("/chain/metrics").
		WithNamespace("chain-app").
		WithSubsystem("chain-service").
		WithSkipPaths([]string{"/chain/health"}).
		WithBuckets([]float64{0.1, 1.0, 10.0})

	assert.True(t, config.Enabled)
	assert.Equal(t, "/chain/metrics", config.Path)
	assert.Equal(t, "chain-app", config.Namespace)
	assert.Equal(t, "chain-service", config.Subsystem)
	assert.Len(t, config.SkipPaths, 1)
	assert.Len(t, config.Buckets, 3)
}

func TestMetrics_DefaultBuckets(t *testing.T) {
	config := Default()
	expectedBuckets := []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0, 10.0}
	assert.Equal(t, expectedBuckets, config.Buckets)
}

func TestMetrics_MetricFlags(t *testing.T) {
	config := Default()
	assert.True(t, config.RequestCount)
	assert.True(t, config.Duration)
	assert.True(t, config.RequestSize)
	assert.True(t, config.ResponseSize)
}
