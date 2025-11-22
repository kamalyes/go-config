/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\timeout\timeout_test.go
 * @Description: 超时中间件配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package timeout

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeout_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "timeout", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, 30*time.Second, config.Duration)
	assert.Equal(t, "请求超时", config.Message)
}

func TestTimeout_WithModuleName(t *testing.T) {
	config := Default()
	result := config.WithModuleName("custom-timeout")
	assert.Equal(t, "custom-timeout", result.ModuleName)
	assert.Equal(t, config, result)
}

func TestTimeout_WithEnabled(t *testing.T) {
	config := Default()
	result := config.WithEnabled(false)
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestTimeout_WithDuration(t *testing.T) {
	config := Default()
	result := config.WithDuration(60 * time.Second)
	assert.Equal(t, 60*time.Second, result.Duration)
	assert.Equal(t, config, result)
}

func TestTimeout_WithMessage(t *testing.T) {
	config := Default()
	result := config.WithMessage("Request timeout occurred")
	assert.Equal(t, "Request timeout occurred", result.Message)
	assert.Equal(t, config, result)
}

func TestTimeout_Enable(t *testing.T) {
	config := Default()
	config.Enabled = false
	result := config.Enable()
	assert.True(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestTimeout_Disable(t *testing.T) {
	config := Default()
	result := config.Disable()
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestTimeout_IsEnabled(t *testing.T) {
	config := Default()
	assert.False(t, config.IsEnabled())

	config.Enabled = true
	assert.True(t, config.IsEnabled())
}

func TestTimeout_Clone(t *testing.T) {
	config := Default()
	config.WithDuration(45 * time.Second).WithMessage("Custom timeout")

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*Timeout)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.Enabled, clonedConfig.Enabled)
	assert.Equal(t, config.Duration, clonedConfig.Duration)
	assert.Equal(t, config.Message, clonedConfig.Message)

	// 验证深拷贝
	clonedConfig.Duration = 90 * time.Second
	assert.NotEqual(t, config.Duration, clonedConfig.Duration)
}

func TestTimeout_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestTimeout_Set(t *testing.T) {
	config := Default()
	newConfig := &Timeout{
		ModuleName: "new-timeout",
		Enabled:    false,
		Duration:   120 * time.Second,
		Message:    "New timeout message",
	}

	config.Set(newConfig)
	assert.Equal(t, "new-timeout", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, 120*time.Second, config.Duration)
	assert.Equal(t, "New timeout message", config.Message)
}

func TestTimeout_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestTimeout_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("api-timeout").
		WithDuration(15 * time.Second).
		WithMessage("API request timed out").
		WithEnabled(true).
		Enable()

	assert.Equal(t, "api-timeout", config.ModuleName)
	assert.Equal(t, 15*time.Second, config.Duration)
	assert.Equal(t, "API request timed out", config.Message)
	assert.True(t, config.Enabled)
}
