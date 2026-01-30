/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\requestid\requestid_test.go
 * @Description: 请求ID中间件配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package requestid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestID_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "requestid", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "X-Request-ID", config.HeaderName)
	assert.Equal(t, "uuid", config.Generator)
}

func TestRequestID_WithHeaderName(t *testing.T) {
	config := Default()
	result := config.WithHeaderName("X-Custom-Request-ID")
	assert.Equal(t, "X-Custom-Request-ID", result.HeaderName)
	assert.Equal(t, config, result)
}

func TestRequestID_WithGenerator(t *testing.T) {
	config := Default()
	result := config.WithGenerator("nanoid")
	assert.Equal(t, "nanoid", result.Generator)
	assert.Equal(t, config, result)
}

func TestRequestID_Enable(t *testing.T) {
	config := Default()
	config.Enabled = false
	result := config.Enable()
	assert.True(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestRequestID_Disable(t *testing.T) {
	config := Default()
	result := config.Disable()
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestRequestID_IsEnabled(t *testing.T) {
	config := Default()
	assert.False(t, config.IsEnabled())

	config.Enabled = false
	assert.False(t, config.IsEnabled())
}

func TestRequestID_Clone(t *testing.T) {
	config := Default()
	config.WithHeaderName("X-Custom-ID").WithGenerator("nanoid")

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*RequestID)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.Enabled, clonedConfig.Enabled)
	assert.Equal(t, config.HeaderName, clonedConfig.HeaderName)
	assert.Equal(t, config.Generator, clonedConfig.Generator)

	// 验证深拷贝
	clonedConfig.HeaderName = "Modified"
	assert.NotEqual(t, config.HeaderName, clonedConfig.HeaderName)
}

func TestRequestID_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestRequestID_Set(t *testing.T) {
	config := Default()
	newConfig := &RequestID{
		ModuleName: "custom-requestid",
		Enabled:    false,
		HeaderName: "X-Custom-Header",
		Generator:  "nanoid",
	}

	config.Set(newConfig)
	assert.Equal(t, "custom-requestid", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "X-Custom-Header", config.HeaderName)
	assert.Equal(t, "nanoid", config.Generator)
}

func TestRequestID_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestRequestID_ChainedCalls(t *testing.T) {
	config := Default().
		WithHeaderName("X-Trace-Request-ID").
		WithGenerator("nanoid").
		Disable()

	assert.Equal(t, "X-Trace-Request-ID", config.HeaderName)
	assert.Equal(t, "nanoid", config.Generator)
	assert.False(t, config.Enabled)
}
