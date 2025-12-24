/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-24 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-24 00:00:00
 * @FilePath: \go-config\pkg\zap\zap_test.go
 * @Description: Zap日志配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package zap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZap_Clone(t *testing.T) {
	zap := &Zap{
		ModuleName:    "zap",
		Level:         "info",
		Format:        "json",
		Prefix:        "app",
		Director:      "/var/log/app",
		MaxSize:       100,
		MaxAge:        30,
		MaxBackups:    10,
		Compress:      true,
		LinkName:      "latest.log",
		ShowLine:      true,
		EncodeLevel:   "LowercaseColorLevelEncoder",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
		Development:   false,
	}

	cloned := zap.Clone()

	assert.NotNil(t, cloned)
	clonedZap, ok := cloned.(*Zap)
	assert.True(t, ok)
	assert.Equal(t, zap.Level, clonedZap.Level)
	assert.Equal(t, zap.Format, clonedZap.Format)
	assert.Equal(t, zap.Director, clonedZap.Director)
	assert.Equal(t, zap.MaxSize, clonedZap.MaxSize)
	assert.Equal(t, zap.Compress, clonedZap.Compress)

	// 验证是独立副本
	clonedZap.Level = "debug"
	assert.NotEqual(t, zap.Level, clonedZap.Level)
}

func TestZap_Get(t *testing.T) {
	zap := &Zap{
		Level:  "info",
		Format: "json",
	}

	result := zap.Get()
	assert.Equal(t, zap, result)
}

func TestZap_Set(t *testing.T) {
	zap := &Zap{}
	newZap := &Zap{
		Level:    "debug",
		Format:   "console",
		Director: "/tmp/logs",
	}

	zap.Set(newZap)
	assert.Equal(t, newZap.Level, zap.Level)
	assert.Equal(t, newZap.Format, zap.Format)
	assert.Equal(t, newZap.Director, zap.Director)
}
