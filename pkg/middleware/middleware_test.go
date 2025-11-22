/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\middleware\middleware_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package middleware

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddleware_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "middleware", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.NotNil(t, config.Logging)
	assert.NotNil(t, config.Recovery)
	assert.NotNil(t, config.Tracing)
	assert.NotNil(t, config.Metrics)
	assert.NotNil(t, config.RequestID)
	assert.NotNil(t, config.I18N)
	assert.NotNil(t, config.PProf)
}

func TestMiddleware_WithModuleName(t *testing.T) {
	config := Default().WithModuleName("custom-middleware")
	assert.Equal(t, "custom-middleware", config.ModuleName)
}

func TestMiddleware_WithEnabled(t *testing.T) {
	config := Default().WithEnabled(false)
	assert.False(t, config.Enabled)
}

func TestMiddleware_Enable(t *testing.T) {
	config := Default().Disable().Enable()
	assert.True(t, config.Enabled)
}

func TestMiddleware_Disable(t *testing.T) {
	config := Default().Disable()
	assert.False(t, config.Enabled)
}

func TestMiddleware_IsEnabled(t *testing.T) {
	config := Default()
	assert.True(t, config.IsEnabled())
	config.Disable()
	assert.False(t, config.IsEnabled())
}

func TestMiddleware_EnableLogging(t *testing.T) {
	config := Default()
	config.Logging.Disable()
	config.EnableLogging()
	assert.True(t, config.Logging.IsEnabled())
}

func TestMiddleware_EnableRecovery(t *testing.T) {
	config := Default()
	config.Recovery.Disable()
	config.EnableRecovery()
	assert.True(t, config.Recovery.IsEnabled())
}

func TestMiddleware_EnableTracing(t *testing.T) {
	config := Default()
	config.Tracing.Disable()
	config.EnableTracing()
	assert.True(t, config.Tracing.IsEnabled())
}

func TestMiddleware_EnableMetrics(t *testing.T) {
	config := Default()
	config.Metrics.Disable()
	config.EnableMetrics()
	assert.True(t, config.Metrics.IsEnabled())
}

func TestMiddleware_EnableRequestID(t *testing.T) {
	config := Default()
	config.RequestID.Disable()
	config.EnableRequestID()
	assert.True(t, config.RequestID.IsEnabled())
}

func TestMiddleware_EnableI18N(t *testing.T) {
	config := Default()
	config.I18N.Disable()
	config.EnableI18N()
	assert.True(t, config.I18N.IsEnabled())
}

func TestMiddleware_EnablePProf(t *testing.T) {
	config := Default()
	config.PProf.Disable()
	config.EnablePProf()
	assert.True(t, config.PProf.IsEnabled())
}

func TestMiddleware_Clone(t *testing.T) {
	original := Default().
		WithModuleName("test-middleware").
		Enable().
		EnableLogging().
		EnableRecovery().
		EnableTracing().
		EnableMetrics()

	cloned := original.Clone().(*Middleware)

	// 验证值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)

	// 验证嵌套配置独立性
	cloned.Metrics.WithNamespace("new-namespace")
	assert.NotEqual(t, original.Metrics.Namespace, cloned.Metrics.Namespace)
}

func TestMiddleware_Get(t *testing.T) {
	config := Default().WithEnabled(false)
	got := config.Get()
	assert.NotNil(t, got)
	middlewareConfig, ok := got.(*Middleware)
	assert.True(t, ok)
	assert.False(t, middlewareConfig.Enabled)
}

func TestMiddleware_Set(t *testing.T) {
	config := Default()
	newConfig := &Middleware{
		ModuleName: "new-middleware",
		Enabled:    false,
	}

	config.Set(newConfig)
	assert.Equal(t, "new-middleware", config.ModuleName)
	assert.False(t, config.Enabled)
}

func TestMiddleware_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestMiddleware_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("chain-middleware").
		Enable().
		EnableLogging().
		EnableRecovery().
		EnableTracing().
		EnableMetrics().
		EnableRequestID().
		EnableI18N().
		EnablePProf()

	assert.Equal(t, "chain-middleware", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.True(t, config.Logging.IsEnabled())
	assert.True(t, config.Recovery.IsEnabled())
	assert.True(t, config.Tracing.IsEnabled())
	assert.True(t, config.Metrics.IsEnabled())
	assert.True(t, config.RequestID.IsEnabled())
	assert.True(t, config.I18N.IsEnabled())
	assert.True(t, config.PProf.IsEnabled())
}

func TestMiddleware_WithSubConfigs(t *testing.T) {
	config := Default()

	// Test WithRecovery
	config.WithRecovery(config.Recovery)
	assert.NotNil(t, config.Recovery)

	// Test WithTracing
	config.WithTracing(config.Tracing)
	assert.NotNil(t, config.Tracing)

	// Test WithMetrics
	config.WithMetrics(config.Metrics.WithNamespace("custom"))
	assert.Equal(t, "custom", config.Metrics.Namespace)

	// Test WithRequestID
	config.WithRequestID(config.RequestID)
	assert.NotNil(t, config.RequestID)

	// Test WithI18N
	config.WithI18N(config.I18N)
	assert.NotNil(t, config.I18N)

	// Test WithPProf
	config.WithPProf(config.PProf)
	assert.NotNil(t, config.PProf)
}

func TestMiddleware_NilSubConfigs(t *testing.T) {
	config := Default()

	// Set sub-configs to nil
	config.Logging = nil
	config.Recovery = nil
	config.Tracing = nil
	config.Metrics = nil
	config.RequestID = nil
	config.I18N = nil
	config.PProf = nil

	// These should not panic
	config.EnableLogging()
	config.EnableRecovery()
	config.EnableTracing()
	config.EnableMetrics()
	config.EnableRequestID()
	config.EnableI18N()
	config.EnablePProf()

	// Verify no panic occurred
	assert.True(t, true)
}
