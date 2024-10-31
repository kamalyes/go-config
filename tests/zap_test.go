/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:58:14
 * @FilePath: \go-config\tests\zap_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/zap"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Zap 配置参数
func generateZapTestParams() *zap.Zap {
	return &zap.Zap{
		ModuleName:    random.RandString(10, random.CAPITAL),
		Level:         random.RandString(5, random.CAPITAL),                                 // 随机生成日志级别
		Format:        random.RandString(5, random.CAPITAL),                                 // 随机生成日志格式
		Prefix:        random.RandString(5, random.CAPITAL),                                 // 随机生成日志前缀
		Director:      fmt.Sprintf("/var/log/%s", random.RandString(5, random.CAPITAL)),     // 随机生成日志目录
		MaxSize:       random.RandInt(1, 100),                                               // 随机生成日志文件的最大大小（以MB为单位）
		MaxAge:        random.RandInt(1, 30),                                                // 随机生成日志最大保留时间（单位：天）
		MaxBackups:    random.RandInt(1, 10),                                                // 随机生成保留旧文件的最大个数
		Compress:      random.FRandBool(),                                                   // 随机生成是否压缩
		LinkName:      fmt.Sprintf("/var/log/%s.log", random.RandString(5, random.CAPITAL)), // 随机生成日志软连接文件
		ShowLine:      random.FRandBool(),                                                   // 随机生成是否在日志中输出源码所在的行
		EncodeLevel:   random.RandString(5, random.CAPITAL),                                 // 随机生成日志编码等级
		StacktraceKey: random.RandString(10, random.CAPITAL),                                // 随机生成堆栈捕捉标识
		LogInConsole:  random.FRandBool(),                                                   // 随机生成是否在控制台打印日志
	}
}

// 将 Zap 的参数转换为 map
func zapToMap(zap *zap.Zap) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME":    zap.ModuleName,
		"LEVEL":          zap.Level,
		"FORMAT":         zap.Format,
		"PREFIX":         zap.Prefix,
		"DIRECTOR":       zap.Director,
		"MAX_SIZE":       zap.MaxSize,
		"MAX_AGE":        zap.MaxAge,
		"MAX_BACKUPS":    zap.MaxBackups,
		"COMPRESS":       zap.Compress,
		"LINK_NAME":      zap.LinkName,
		"SHOW_LINE":      zap.ShowLine,
		"ENCODE_LEVEL":   zap.EncodeLevel,
		"STACKTRACE_KEY": zap.StacktraceKey,
		"LOG_IN_CONSOLE": zap.LogInConsole,
	}
}

// 验证 Zap 的字段与期望的映射是否相等
func assertZapFields(t *testing.T, zap *zap.Zap, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], zap.ModuleName)
	assert.Equal(t, expected["LEVEL"], zap.Level)
	assert.Equal(t, expected["FORMAT"], zap.Format)
	assert.Equal(t, expected["PREFIX"], zap.Prefix)
	assert.Equal(t, expected["DIRECTOR"], zap.Director)
	assert.Equal(t, expected["MAX_SIZE"], zap.MaxSize)
	assert.Equal(t, expected["MAX_AGE"], zap.MaxAge)
	assert.Equal(t, expected["MAX_BACKUPS"], zap.MaxBackups)
	assert.Equal(t, expected["COMPRESS"], zap.Compress)
	assert.Equal(t, expected["LINK_NAME"], zap.LinkName)
	assert.Equal(t, expected["SHOW_LINE"], zap.ShowLine)
	assert.Equal(t, expected["ENCODE_LEVEL"], zap.EncodeLevel)
	assert.Equal(t, expected["STACKTRACE_KEY"], zap.StacktraceKey)
	assert.Equal(t, expected["LOG_IN_CONSOLE"], zap.LogInConsole)
}

func TestZapClone(t *testing.T) {
	params := generateZapTestParams()
	original := zap.NewZap(params)
	cloned := original.Clone().(*zap.Zap)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestZapSet(t *testing.T) {
	oldParams := generateZapTestParams()
	newParams := generateZapTestParams()

	zapInstance := zap.NewZap(oldParams)
	newConfig := zap.NewZap(newParams)

	zapInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, zapInstance.ModuleName)
	assert.Equal(t, newParams.Level, zapInstance.Level)
	assert.Equal(t, newParams.Format, zapInstance.Format)
	assert.Equal(t, newParams.Prefix, zapInstance.Prefix)
	assert.Equal(t, newParams.Director, zapInstance.Director)
	assert.Equal(t, newParams.MaxSize, zapInstance.MaxSize)
	assert.Equal(t, newParams.MaxAge, zapInstance.MaxAge)
	assert.Equal(t, newParams.MaxBackups, zapInstance.MaxBackups)
	assert.Equal(t, newParams.Compress, zapInstance.Compress)
	assert.Equal(t, newParams.LinkName, zapInstance.LinkName)
	assert.Equal(t, newParams.ShowLine, zapInstance.ShowLine)
	assert.Equal(t, newParams.EncodeLevel, zapInstance.EncodeLevel)
	assert.Equal(t, newParams.StacktraceKey, zapInstance.StacktraceKey)
	assert.Equal(t, newParams.LogInConsole, zapInstance.LogInConsole)
}
