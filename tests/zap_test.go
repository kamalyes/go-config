/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 12:55:00
 * @FilePath: \go-config\zap\zap_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/zap"
	"github.com/stretchr/testify/assert"
)

// generateZapTestParams 生成测试参数
func generateZapTestParams() zap.Zap {
	return zap.Zap{
		ModuleName:    "test-module",
		Level:         "info",
		Format:        "json",
		Prefix:        "[TEST]",
		Director:      "/var/log",
		MaxSize:       10,
		MaxAge:        30,
		MaxBackups:    5,
		Compress:      true,
		LinkName:      "latest.log",
		ShowLine:      true,
		EncodeLevel:   "color",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
	}
}

func TestNewZap(t *testing.T) {
	params := generateZapTestParams()
	zapConfig := zap.NewZap(params.ModuleName, params.Level, params.Format, params.Prefix, params.Director, params.MaxSize, params.MaxAge, params.MaxBackups, params.Compress, params.LinkName, params.ShowLine, params.EncodeLevel, params.StacktraceKey, params.LogInConsole)

	assert.NotNil(t, zapConfig)
	assert.Equal(t, params.ModuleName, zapConfig.ModuleName)
	assert.Equal(t, params.Level, zapConfig.Level)
	assert.Equal(t, params.Format, zapConfig.Format)
	assert.Equal(t, params.Prefix, zapConfig.Prefix)
	assert.Equal(t, params.Director, zapConfig.Director)
	assert.Equal(t, params.MaxSize, zapConfig.MaxSize)
	assert.Equal(t, params.MaxAge, zapConfig.MaxAge)
	assert.Equal(t, params.MaxBackups, zapConfig.MaxBackups)
	assert.Equal(t, params.Compress, zapConfig.Compress)
	assert.Equal(t, params.LinkName, zapConfig.LinkName)
	assert.Equal(t, params.ShowLine, zapConfig.ShowLine)
	assert.Equal(t, params.EncodeLevel, zapConfig.EncodeLevel)
	assert.Equal(t, params.StacktraceKey, zapConfig.StacktraceKey)
	assert.Equal(t, params.LogInConsole, zapConfig.LogInConsole)
}

func TestZapValidate(t *testing.T) {
	validParams := generateZapTestParams()
	validZap := zap.NewZap(validParams.ModuleName, validParams.Level, validParams.Format, validParams.Prefix, validParams.Director, validParams.MaxSize, validParams.MaxAge, validParams.MaxBackups, validParams.Compress, validParams.LinkName, validParams.ShowLine, validParams.EncodeLevel, validParams.StacktraceKey, validParams.LogInConsole)
	assert.NoError(t, validZap.Validate())

	invalidZap := zap.NewZap("", "", "", "", "", -1, -1, -1, false, "", false, "", "", false)
	assert.Error(t, invalidZap.Validate())
	assert.EqualError(t, invalidZap.Validate(), "module name cannot be empty")
}

func TestZapClone(t *testing.T) {
	params := generateZapTestParams()
	original := zap.NewZap(params.ModuleName, params.Level, params.Format, params.Prefix, params.Director, params.MaxSize, params.MaxAge, params.MaxBackups, params.Compress, params.LinkName, params.ShowLine, params.EncodeLevel, params.StacktraceKey, params.LogInConsole)
	cloned := original.Clone().(*zap.Zap)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestZapToMap(t *testing.T) {
	params := generateZapTestParams()
	zapConfig := zap.NewZap(params.ModuleName, params.Level, params.Format, params.Prefix, params.Director, params.MaxSize, params.MaxAge, params.MaxBackups, params.Compress, params.LinkName, params.ShowLine, params.EncodeLevel, params.StacktraceKey, params.LogInConsole)
	expectedMap := map[string]interface{}{
		"moduleName":    params.ModuleName,
		"level":         params.Level,
		"format":        params.Format,
		"prefix":        params.Prefix,
		"director":      params.Director,
		"maxSize":       params.MaxSize,
		"maxAge":        params.MaxAge,
		"maxBackups":    params.MaxBackups,
		"compress":      params.Compress,
		"linkName":      params.LinkName,
		"showLine":      params.ShowLine,
		"encodeLevel":   params.EncodeLevel,
		"stacktraceKey": params.StacktraceKey,
		"logInConsole":  params.LogInConsole,
	}
	assert.Equal(t, expectedMap, zapConfig.ToMap())
}

func TestZapFromMap(t *testing.T) {
	zapConfig := zap.NewZap("", "", "", "", "", 0, 0, 0, false, "", false, "", "", false)
	data := map[string]interface{}{
		"moduleName":    "new-module",
		"level":         "debug",
		"format":        "text",
		"prefix":        "[NEW]",
		"director":      "/var/log/new",
		"maxSize":       20,
		"maxAge":        15,
		"maxBackups":    10,
		"compress":      false,
		"linkName":      "newest.log",
		"showLine":      false,
		"encodeLevel":   "json",
		"stacktraceKey": "new-stacktrace",
		"logInConsole":  false,
	}
	zapConfig.FromMap(data)

	// 验证填充后的数据是否正确
	assert.Equal(t, "new-module", zapConfig.ModuleName)
	assert.Equal(t, "debug", zapConfig.Level)
	assert.Equal(t, "text", zapConfig.Format)
	assert.Equal(t, "[NEW]", zapConfig.Prefix)
	assert.Equal(t, "/var/log/new", zapConfig.Director)
	assert.Equal(t, 20, zapConfig.MaxSize)
	assert.Equal(t, 15, zapConfig.MaxAge)
	assert.Equal(t, 10, zapConfig.MaxBackups)
	assert.Equal(t, false, zapConfig.Compress)
	assert.Equal(t, "newest.log", zapConfig.LinkName)
	assert.Equal(t, false, zapConfig.ShowLine)
	assert.Equal(t, "json", zapConfig.EncodeLevel)
	assert.Equal(t, "new-stacktrace", zapConfig.StacktraceKey)
	assert.Equal(t, false, zapConfig.LogInConsole)
}

func TestZapSet(t *testing.T) {
	oldParams := generateZapTestParams()
	newParams := zap.Zap{
		ModuleName:    "new-module",
		Level:         "debug",
		Format:        "text",
		Prefix:        "[NEW]",
		Director:      "/var/log/new",
		MaxSize:       20,
		MaxAge:        15,
		MaxBackups:    10,
		Compress:      false,
		LinkName:      "newest.log",
		ShowLine:      false,
		EncodeLevel:   "json",
		StacktraceKey: "new-stacktrace",
		LogInConsole:  false,
	}

	zapConfig := zap.NewZap(oldParams.ModuleName, oldParams.Level, oldParams.Format, oldParams.Prefix, oldParams.Director, oldParams.MaxSize, oldParams.MaxAge, oldParams.MaxBackups, oldParams.Compress, oldParams.LinkName, oldParams.ShowLine, oldParams.EncodeLevel, oldParams.StacktraceKey, oldParams.LogInConsole)
	newConfig := zap.NewZap(newParams.ModuleName, newParams.Level, newParams.Format, newParams.Prefix, newParams.Director, newParams.MaxSize, newParams.MaxAge, newParams.MaxBackups, newParams.Compress, newParams.LinkName, newParams.ShowLine, newParams.EncodeLevel, newParams.StacktraceKey, newParams.LogInConsole)

	zapConfig.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, zapConfig.ModuleName)
	assert.Equal(t, newParams.Level, zapConfig.Level)
	assert.Equal(t, newParams.Format, zapConfig.Format)
	assert.Equal(t, newParams.Prefix, zapConfig.Prefix)
	assert.Equal(t, newParams.Director, zapConfig.Director)
	assert.Equal(t, newParams.MaxSize, zapConfig.MaxSize)
	assert.Equal(t, newParams.MaxAge, zapConfig.MaxAge)
	assert.Equal(t, newParams.MaxBackups, zapConfig.MaxBackups)
	assert.Equal(t, newParams.Compress, zapConfig.Compress)
	assert.Equal(t, newParams.LinkName, zapConfig.LinkName)
	assert.Equal(t, newParams.ShowLine, zapConfig.ShowLine)
	assert.Equal(t, newParams.EncodeLevel, zapConfig.EncodeLevel)
	assert.Equal(t, newParams.StacktraceKey, zapConfig.StacktraceKey)
	assert.Equal(t, newParams.LogInConsole, zapConfig.LogInConsole)
}
