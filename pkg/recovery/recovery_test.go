/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\recovery\recovery_test.go
 * @Description: 恢复中间件配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package recovery

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecovery_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "recovery", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.True(t, config.EnableStack)
	assert.Equal(t, 4096, config.StackSize)
	assert.False(t, config.EnableDebug)
	assert.Equal(t, "服务器内部错误", config.ErrorMessage)
	assert.Equal(t, "error", config.LogLevel)
	assert.False(t, config.EnableNotify)
	assert.Nil(t, config.RecoveryHandler)
	assert.True(t, config.PrintStack)
}

func TestRecovery_WithPrintStack(t *testing.T) {
	config := Default()
	result := config.WithPrintStack(false)
	assert.False(t, result.PrintStack)
	assert.False(t, result.EnableStack)
	assert.Equal(t, config, result)
}

func TestRecovery_WithEnableStack(t *testing.T) {
	config := Default()
	result := config.WithEnableStack(false)
	assert.False(t, result.EnableStack)
	assert.Equal(t, config, result)
}

func TestRecovery_WithStackSize(t *testing.T) {
	config := Default()
	result := config.WithStackSize(8192)
	assert.Equal(t, 8192, result.StackSize)
	assert.Equal(t, config, result)
}

func TestRecovery_WithEnableDebug(t *testing.T) {
	config := Default()
	result := config.WithEnableDebug(true)
	assert.True(t, result.EnableDebug)
	assert.Equal(t, config, result)
}

func TestRecovery_WithErrorMessage(t *testing.T) {
	config := Default()
	result := config.WithErrorMessage("Custom error")
	assert.Equal(t, "Custom error", result.ErrorMessage)
	assert.Equal(t, config, result)
}

func TestRecovery_WithRecoveryHandler(t *testing.T) {
	config := Default()
	handler := func(w http.ResponseWriter, r *http.Request, err interface{}) {}
	result := config.WithRecoveryHandler(handler)
	assert.NotNil(t, result.RecoveryHandler)
	assert.Equal(t, config, result)
}

func TestRecovery_WithLogLevel(t *testing.T) {
	config := Default()
	result := config.WithLogLevel("warn")
	assert.Equal(t, "warn", result.LogLevel)
	assert.Equal(t, config, result)
}

func TestRecovery_WithNotify(t *testing.T) {
	config := Default()
	result := config.WithNotify(true)
	assert.True(t, result.EnableNotify)
	assert.Equal(t, config, result)
}

func TestRecovery_Enable(t *testing.T) {
	config := Default()
	config.Enabled = false
	result := config.Enable()
	assert.True(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestRecovery_Disable(t *testing.T) {
	config := Default()
	result := config.Disable()
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestRecovery_IsEnabled(t *testing.T) {
	config := Default()
	assert.True(t, config.IsEnabled())
	
	config.Enabled = false
	assert.False(t, config.IsEnabled())
}

func TestRecovery_Clone(t *testing.T) {
	config := Default()
	config.WithStackSize(8192).WithEnableDebug(true).WithNotify(true)
	
	clone := config.Clone()
	assert.NotNil(t, clone)
	
	clonedConfig, ok := clone.(*Recovery)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.Enabled, clonedConfig.Enabled)
	assert.Equal(t, config.EnableStack, clonedConfig.EnableStack)
	assert.Equal(t, config.StackSize, clonedConfig.StackSize)
	assert.Equal(t, config.EnableDebug, clonedConfig.EnableDebug)
	assert.Equal(t, config.ErrorMessage, clonedConfig.ErrorMessage)
	assert.Equal(t, config.LogLevel, clonedConfig.LogLevel)
	assert.Equal(t, config.EnableNotify, clonedConfig.EnableNotify)
	assert.Equal(t, config.PrintStack, clonedConfig.PrintStack)
}

func TestRecovery_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestRecovery_Set(t *testing.T) {
	config := Default()
	newConfig := &Recovery{
		ModuleName:   "custom-recovery",
		Enabled:      false,
		EnableStack:  false,
		StackSize:    2048,
		EnableDebug:  true,
		ErrorMessage: "Custom error",
		LogLevel:     "debug",
		EnableNotify: true,
		PrintStack:   false,
	}
	
	config.Set(newConfig)
	assert.Equal(t, "custom-recovery", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.False(t, config.EnableStack)
	assert.Equal(t, 2048, config.StackSize)
	assert.True(t, config.EnableDebug)
	assert.Equal(t, "Custom error", config.ErrorMessage)
	assert.Equal(t, "debug", config.LogLevel)
	assert.True(t, config.EnableNotify)
	assert.False(t, config.PrintStack)
}

func TestRecovery_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestRecovery_ChainedCalls(t *testing.T) {
	config := Default().
		WithStackSize(16384).
		WithEnableDebug(true).
		WithErrorMessage("Panic occurred").
		WithLogLevel("fatal").
		WithNotify(true).
		WithPrintStack(false).
		Enable()
	
	assert.Equal(t, 16384, config.StackSize)
	assert.True(t, config.EnableDebug)
	assert.Equal(t, "Panic occurred", config.ErrorMessage)
	assert.Equal(t, "fatal", config.LogLevel)
	assert.True(t, config.EnableNotify)
	assert.False(t, config.PrintStack)
	assert.False(t, config.EnableStack)
	assert.True(t, config.Enabled)
}
