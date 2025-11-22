/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-21 23:59:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-21 23:59:00
 * @FilePath: \go-config\pkg\breaker\breaker_test.go
 * @Description: 断路器配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package breaker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCircuitBreaker_Default(t *testing.T) {
	breaker := Default()

	assert.NotNil(t, breaker)
	assert.Equal(t, "circuit_breaker", breaker.ModuleName)
	assert.True(t, breaker.Enabled)
	assert.Equal(t, 5, breaker.FailureThreshold)
	assert.Equal(t, 2, breaker.SuccessThreshold)
	assert.Equal(t, int64(30*time.Second), breaker.Timeout)
	assert.Equal(t, 10, breaker.VolumeThreshold)
	assert.Equal(t, 100, breaker.SlidingWindowSize)
	assert.Equal(t, int64(1*time.Second), breaker.SlidingWindowBucket)
	assert.Equal(t, []string{"/api/"}, breaker.PreventionPaths)
	assert.Equal(t, []string{"/health", "/metrics", "/debug"}, breaker.ExcludePaths)
}

func TestCircuitBreaker_GetLabel(t *testing.T) {
	breaker := Default()
	assert.Equal(t, "circuit_breaker", breaker.GetLabel())

	breaker.ModuleName = "custom_breaker"
	assert.Equal(t, "custom_breaker", breaker.GetLabel())
}

func TestCircuitBreaker_Init(t *testing.T) {
	breaker := Default()
	err := breaker.Init()
	assert.NoError(t, err)
}

func TestCircuitBreaker_Validate(t *testing.T) {
	breaker := Default()
	err := breaker.Validate()
	assert.NoError(t, err)
}

func TestCircuitBreaker_CustomValues(t *testing.T) {
	breaker := &CircuitBreaker{
		ModuleName:          "custom",
		Enabled:             false,
		FailureThreshold:    10,
		SuccessThreshold:    5,
		Timeout:             int64(60 * time.Second),
		VolumeThreshold:     20,
		SlidingWindowSize:   200,
		SlidingWindowBucket: int64(5 * time.Second),
		PreventionPaths:     []string{"/v1/", "/v2/"},
		ExcludePaths:        []string{"/status"},
	}

	assert.Equal(t, "custom", breaker.ModuleName)
	assert.False(t, breaker.Enabled)
	assert.Equal(t, 10, breaker.FailureThreshold)
	assert.Equal(t, 5, breaker.SuccessThreshold)
	assert.Equal(t, int64(60*time.Second), breaker.Timeout)
	assert.Equal(t, 20, breaker.VolumeThreshold)
	assert.Equal(t, 200, breaker.SlidingWindowSize)
	assert.Equal(t, int64(5*time.Second), breaker.SlidingWindowBucket)
	assert.Equal(t, []string{"/v1/", "/v2/"}, breaker.PreventionPaths)
	assert.Equal(t, []string{"/status"}, breaker.ExcludePaths)
}

func TestWebSocketBreaker_Default(t *testing.T) {
	breaker := DefaultWebSocketBreaker()

	assert.NotNil(t, breaker)
	assert.Equal(t, "websocket_breaker", breaker.ModuleName)
	assert.True(t, breaker.Enabled)
	assert.Equal(t, 5, breaker.FailureThreshold)
	assert.Equal(t, 2, breaker.SuccessThreshold)
	assert.Equal(t, int64(30*time.Second), breaker.Timeout)
	assert.Equal(t, 3, breaker.MaxRetries)
	assert.Equal(t, 2.0, breaker.RetryBackoffFactor)
	assert.Equal(t, int64(10*time.Second), breaker.HealthCheckInterval)
	assert.Equal(t, 1000, breaker.MessageQueueSize)
}

func TestWebSocketBreaker_GetLabel(t *testing.T) {
	breaker := DefaultWebSocketBreaker()
	assert.Equal(t, "websocket_breaker", breaker.GetLabel())

	breaker.ModuleName = "custom_ws_breaker"
	assert.Equal(t, "custom_ws_breaker", breaker.GetLabel())
}

func TestWebSocketBreaker_Init(t *testing.T) {
	breaker := DefaultWebSocketBreaker()
	err := breaker.Init()
	assert.NoError(t, err)
}

func TestWebSocketBreaker_Validate(t *testing.T) {
	breaker := DefaultWebSocketBreaker()
	err := breaker.Validate()
	assert.NoError(t, err)
}

func TestWebSocketBreaker_CustomValues(t *testing.T) {
	breaker := &WebSocketBreaker{
		ModuleName:          "custom_ws",
		Enabled:             false,
		FailureThreshold:    10,
		SuccessThreshold:    3,
		Timeout:             int64(45 * time.Second),
		MaxRetries:          5,
		RetryBackoffFactor:  1.5,
		HealthCheckInterval: int64(15 * time.Second),
		MessageQueueSize:    2000,
	}

	assert.Equal(t, "custom_ws", breaker.ModuleName)
	assert.False(t, breaker.Enabled)
	assert.Equal(t, 10, breaker.FailureThreshold)
	assert.Equal(t, 3, breaker.SuccessThreshold)
	assert.Equal(t, int64(45*time.Second), breaker.Timeout)
	assert.Equal(t, 5, breaker.MaxRetries)
	assert.Equal(t, 1.5, breaker.RetryBackoffFactor)
	assert.Equal(t, int64(15*time.Second), breaker.HealthCheckInterval)
	assert.Equal(t, 2000, breaker.MessageQueueSize)
}

func TestCircuitBreaker_PreventionPaths(t *testing.T) {
	breaker := Default()

	// 测试默认路径
	assert.Contains(t, breaker.PreventionPaths, "/api/")

	// 测试自定义路径
	breaker.PreventionPaths = []string{"/custom/", "/protected/"}
	assert.Len(t, breaker.PreventionPaths, 2)
	assert.Contains(t, breaker.PreventionPaths, "/custom/")
	assert.Contains(t, breaker.PreventionPaths, "/protected/")
}

func TestCircuitBreaker_ExcludePaths(t *testing.T) {
	breaker := Default()

	// 测试默认排除路径
	assert.Contains(t, breaker.ExcludePaths, "/health")
	assert.Contains(t, breaker.ExcludePaths, "/metrics")
	assert.Contains(t, breaker.ExcludePaths, "/debug")

	// 测试自定义排除路径
	breaker.ExcludePaths = []string{"/status", "/ping"}
	assert.Len(t, breaker.ExcludePaths, 2)
	assert.Contains(t, breaker.ExcludePaths, "/status")
	assert.Contains(t, breaker.ExcludePaths, "/ping")
}

func TestCircuitBreaker_Thresholds(t *testing.T) {
	breaker := Default()

	// 测试失败阈值
	assert.Equal(t, 5, breaker.FailureThreshold)
	breaker.FailureThreshold = 10
	assert.Equal(t, 10, breaker.FailureThreshold)

	// 测试成功阈值
	assert.Equal(t, 2, breaker.SuccessThreshold)
	breaker.SuccessThreshold = 3
	assert.Equal(t, 3, breaker.SuccessThreshold)

	// 测试容量阈值
	assert.Equal(t, 10, breaker.VolumeThreshold)
	breaker.VolumeThreshold = 20
	assert.Equal(t, 20, breaker.VolumeThreshold)
}
