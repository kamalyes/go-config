/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 11:30:00
 * @FilePath: \go-config\pkg\logging\logging_test.go
 * @Description: 日志中间件配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging_Clone(t *testing.T) {
	original := &Logging{
		ModuleName:           "test-logging",
		Enabled:              true,
		Level:                "debug",
		Format:               "json",
		Output:               "file",
		FilePath:             "/tmp/app.log",
		MaxSize:              100,
		MaxBackups:           3,
		MaxAge:               28,
		Compress:             true,
		SkipPaths:            []string{"/health", "/metrics"},
		EnableRequest:        true,
		EnableResponse:       false,
		MaxBodySize:          2048,
		SensitiveMask:        "***",
		SensitiveKeys:        []string{"password", "token"},
		SlowHTTPThreshold:    1000,
		SlowGRPCThreshold:    1000,
		SlowStreamThreshold:  5000,
		LoggableContentTypes: []string{"application/json", "text/plain"},
	}

	cloned := original.Clone().(*Logging)

	// 验证克隆后的值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.Level, cloned.Level)
	assert.Equal(t, original.Format, cloned.Format)
	assert.Equal(t, original.Output, cloned.Output)
	assert.Equal(t, original.FilePath, cloned.FilePath)
	assert.Equal(t, original.MaxSize, cloned.MaxSize)
	assert.Equal(t, original.MaxBackups, cloned.MaxBackups)
	assert.Equal(t, original.MaxAge, cloned.MaxAge)
	assert.Equal(t, original.Compress, cloned.Compress)
	assert.Equal(t, original.EnableRequest, cloned.EnableRequest)
	assert.Equal(t, original.EnableResponse, cloned.EnableResponse)
	assert.Equal(t, original.MaxBodySize, cloned.MaxBodySize)
	assert.Equal(t, original.SensitiveMask, cloned.SensitiveMask)
	assert.Equal(t, original.SlowHTTPThreshold, cloned.SlowHTTPThreshold)
	assert.Equal(t, original.SlowGRPCThreshold, cloned.SlowGRPCThreshold)
	assert.Equal(t, original.SlowStreamThreshold, cloned.SlowStreamThreshold)

	// 验证slice深拷贝
	assert.Equal(t, original.SkipPaths, cloned.SkipPaths)
	assert.Equal(t, original.SensitiveKeys, cloned.SensitiveKeys)
	assert.Equal(t, original.LoggableContentTypes, cloned.LoggableContentTypes)

	// 修改原始对象不应影响克隆对象
	original.Level = "error"
	original.SkipPaths[0] = "/ping"
	original.SensitiveKeys[0] = "secret"

	assert.NotEqual(t, original.Level, cloned.Level)
	assert.NotEqual(t, original.SkipPaths[0], cloned.SkipPaths[0])
	assert.NotEqual(t, original.SensitiveKeys[0], cloned.SensitiveKeys[0])
}
