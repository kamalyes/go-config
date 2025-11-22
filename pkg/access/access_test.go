/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-21 23:58:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-21 23:58:00
 * @FilePath: \go-config\pkg\access\access_test.go
 * @Description: 访问记录配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package access

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccess_Default(t *testing.T) {
	access := Default()

	assert.NotNil(t, access)
	assert.Equal(t, "access", access.ModuleName)
	assert.True(t, access.Enabled)
	assert.Equal(t, "rpc-gateway", access.ServiceName)
	assert.Equal(t, 60, access.RetentionDays)
	assert.True(t, access.IncludeBody)
	assert.True(t, access.IncludeResponse)
	assert.Equal(t, []string{"User-Agent", "X-Request-ID", "X-Trace-ID", "Authorization", "Content-Type"}, access.IncludeHeaders)
	assert.Equal(t, []string{"/health", "/metrics", "/debug"}, access.ExcludePaths)
	assert.Equal(t, int64(1024*1024), access.MaxBodySize)
	assert.Equal(t, int64(1024*1024*5), access.MaxResponseSize)
}

func TestAccess_WithServiceName(t *testing.T) {
	access := Default()
	result := access.WithServiceName("custom-service")

	assert.Equal(t, "custom-service", result.ServiceName)
	assert.Equal(t, access, result)
}

func TestAccess_WithRetentionDays(t *testing.T) {
	access := Default()
	result := access.WithRetentionDays(90)

	assert.Equal(t, 90, result.RetentionDays)
	assert.Equal(t, access, result)
}

func TestAccess_WithIncludeBody(t *testing.T) {
	access := Default()
	result := access.WithIncludeBody(false)

	assert.False(t, result.IncludeBody)
	assert.Equal(t, access, result)
}

func TestAccess_WithIncludeResponse(t *testing.T) {
	access := Default()
	result := access.WithIncludeResponse(false)

	assert.False(t, result.IncludeResponse)
	assert.Equal(t, access, result)
}

func TestAccess_WithIncludeHeaders(t *testing.T) {
	access := Default()
	customHeaders := []string{"X-Custom-Header", "X-API-Key"}
	result := access.WithIncludeHeaders(customHeaders)

	assert.Equal(t, customHeaders, result.IncludeHeaders)
	assert.Equal(t, access, result)
}

func TestAccess_WithExcludePaths(t *testing.T) {
	access := Default()
	customPaths := []string{"/api/health", "/api/status"}
	result := access.WithExcludePaths(customPaths)

	assert.Equal(t, customPaths, result.ExcludePaths)
	assert.Equal(t, access, result)
}

func TestAccess_WithMaxBodySize(t *testing.T) {
	access := Default()
	result := access.WithMaxBodySize(2048)

	assert.Equal(t, int64(2048), result.MaxBodySize)
	assert.Equal(t, access, result)
}

func TestAccess_WithMaxResponseSize(t *testing.T) {
	access := Default()
	result := access.WithMaxResponseSize(4096)

	assert.Equal(t, int64(4096), result.MaxResponseSize)
	assert.Equal(t, access, result)
}

func TestAccess_Enable(t *testing.T) {
	access := Default()
	access.Enabled = false
	result := access.Enable()

	assert.True(t, result.Enabled)
	assert.True(t, result.IsEnabled())
	assert.Equal(t, access, result)
}

func TestAccess_Disable(t *testing.T) {
	access := Default()
	result := access.Disable()

	assert.False(t, result.Enabled)
	assert.False(t, result.IsEnabled())
	assert.Equal(t, access, result)
}

func TestAccess_IsEnabled(t *testing.T) {
	access := Default()
	assert.True(t, access.IsEnabled())

	access.Enabled = false
	assert.False(t, access.IsEnabled())
}

func TestAccess_Get(t *testing.T) {
	access := Default()
	result := access.Get()

	assert.NotNil(t, result)
	assert.Equal(t, access, result)
}

func TestAccess_Set(t *testing.T) {
	access := Default()
	newAccess := &Access{
		ModuleName:      "custom-access",
		Enabled:         false,
		ServiceName:     "new-service",
		RetentionDays:   30,
		IncludeBody:     false,
		IncludeResponse: false,
		IncludeHeaders:  []string{"Custom-Header"},
		ExcludePaths:    []string{"/custom/path"},
		MaxBodySize:     512,
		MaxResponseSize: 1024,
	}

	access.Set(newAccess)

	assert.Equal(t, "custom-access", access.ModuleName)
	assert.False(t, access.Enabled)
	assert.Equal(t, "new-service", access.ServiceName)
	assert.Equal(t, 30, access.RetentionDays)
	assert.False(t, access.IncludeBody)
	assert.False(t, access.IncludeResponse)
	assert.Equal(t, []string{"Custom-Header"}, access.IncludeHeaders)
	assert.Equal(t, []string{"/custom/path"}, access.ExcludePaths)
	assert.Equal(t, int64(512), access.MaxBodySize)
	assert.Equal(t, int64(1024), access.MaxResponseSize)
}

func TestAccess_Set_InvalidType(t *testing.T) {
	access := Default()
	originalServiceName := access.ServiceName

	access.Set("invalid type")

	assert.Equal(t, originalServiceName, access.ServiceName)
}

func TestAccess_Clone(t *testing.T) {
	access := Default()
	access.ServiceName = "original-service"
	access.IncludeHeaders = []string{"X-Original"}

	cloned := access.Clone()

	assert.NotNil(t, cloned)
	clonedAccess, ok := cloned.(*Access)
	assert.True(t, ok)
	assert.Equal(t, access.ServiceName, clonedAccess.ServiceName)
	assert.Equal(t, access.IncludeHeaders, clonedAccess.IncludeHeaders)

	// 验证是独立副本 - 修改切片不影响原对象
	clonedAccess.IncludeHeaders[0] = "X-Modified"
	assert.NotEqual(t, access.IncludeHeaders[0], clonedAccess.IncludeHeaders[0])

	clonedAccess.ServiceName = "modified-service"
	assert.NotEqual(t, access.ServiceName, clonedAccess.ServiceName)
}

func TestAccess_Validate(t *testing.T) {
	access := Default()
	err := access.Validate()
	assert.NoError(t, err)
}

func TestAccess_ChainedCalls(t *testing.T) {
	access := Default()
	result := access.
		WithServiceName("chained-service").
		WithRetentionDays(120).
		WithIncludeBody(false).
		WithIncludeResponse(false).
		WithMaxBodySize(2048).
		WithMaxResponseSize(8192).
		Enable()

	assert.Equal(t, "chained-service", result.ServiceName)
	assert.Equal(t, 120, result.RetentionDays)
	assert.False(t, result.IncludeBody)
	assert.False(t, result.IncludeResponse)
	assert.Equal(t, int64(2048), result.MaxBodySize)
	assert.Equal(t, int64(8192), result.MaxResponseSize)
	assert.True(t, result.Enabled)
}
