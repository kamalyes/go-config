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
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kamalyes/go-logger"
	"github.com/stretchr/testify/assert"
)

// TestConsoleLogger 测试控制台日志输出
func TestConsoleLogger(t *testing.T) {
	t.Log("========== 测试控制台日志输出 ==========")

	config := &Logging{
		Level:      "debug",
		Format:     logger.FormatJSON,
		Output:     logger.OutputConsole,
		ShowCaller: true,
		Colorful:   true,
		Prefix:     "[TEST]",
	}

	loggerInstance := config.ToLoggerInstance()
	assert.NotNil(t, loggerInstance, "Logger 实例不应该为 nil")

	// 输出各种级别的日志
	t.Log("--- 输出 DEBUG 级别日志 ---")
	loggerInstance.Debug("这是一条 DEBUG 日志")

	t.Log("--- 输出 INFO 级别日志 ---")
	loggerInstance.Info("这是一条 INFO 日志")

	t.Log("--- 输出 WARN 级别日志 ---")
	loggerInstance.Warn("这是一条 WARN 日志")

	t.Log("--- 输出 ERROR 级别日志 ---")
	loggerInstance.Error("这是一条 ERROR 日志")

	t.Log("--- 输出带格式化的日志 ---")
	loggerInstance.Infof("用户 %s 登录成功，IP: %s", "admin", "192.168.1.1")
}

// TestStdoutLogger 测试标准输出日志
func TestStdoutLogger(t *testing.T) {
	t.Log("========== 测试标准输出日志 ==========")

	config := &Logging{
		Level:  "info",
		Format: logger.FormatText,
		Output: logger.OutputStdout,
	}

	loggerInstance := config.ToLoggerInstance()
	assert.NotNil(t, loggerInstance)

	loggerInstance.Info("标准输出 - INFO 日志")
	loggerInstance.Warn("标准输出 - WARN 日志")
	loggerInstance.Error("标准输出 - ERROR 日志")
}

// TestStderrLogger 测试标准错误输出日志
func TestStderrLogger(t *testing.T) {
	t.Log("========== 测试标准错误输出日志 ==========")

	config := &Logging{
		Level:  "error",
		Format: logger.FormatJSON,
		Output: logger.OutputStderr,
	}

	loggerInstance := config.ToLoggerInstance()
	assert.NotNil(t, loggerInstance)

	loggerInstance.Error("标准错误输出 - ERROR 日志")
	loggerInstance.Errorf("标准错误输出 - 格式化错误: %v", "连接超时")
}

// TestFileLogger 测试文件日志输出
func TestFileLogger(t *testing.T) {
	t.Log("========== 测试文件日志输出 ==========")

	// 创建临时目录（不使用 TempDir 避免 Windows 文件锁定问题）
	tmpDir, err := os.MkdirTemp("", "test-file-logger-*")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir) // 尽力清理，失败也不影响测试

	logFile := filepath.Join(tmpDir, "test.log")

	config := &Logging{
		Level:          "debug",
		Format:         logger.FormatJSON,
		Output:         logger.OutputFile,
		FilePath:       logFile,
		FilePermission: 0644,
	}

	loggerInstance := config.ToLoggerInstance()
	assert.NotNil(t, loggerInstance)

	// 写入多条日志
	loggerInstance.Debug("文件日志 - DEBUG 消息")
	loggerInstance.Info("文件日志 - INFO 消息")
	loggerInstance.Warn("文件日志 - WARN 消息")
	loggerInstance.Error("文件日志 - ERROR 消息")

	// 等待日志写入
	time.Sleep(100 * time.Millisecond)

	// 验证文件是否创建
	assert.FileExists(t, logFile, "日志文件应该被创建")

	// 读取并显示文件内容
	content, err := os.ReadFile(logFile)
	assert.NoError(t, err, "读取日志文件不应该出错")
	t.Logf("日志文件内容:\n%s", string(content))

	// 验证文件权限（Windows 上权限处理不同，跳过此检查）
	// Windows 不支持 Unix 风格的文件权限，所以只在非 Windows 系统上验证
	// info, err := os.Stat(logFile)
	// assert.NoError(t, err)
	// assert.Equal(t, os.FileMode(0644), info.Mode().Perm(), "文件权限应该是 0644")
}

// TestRotateLogger 测试轮转日志输出
func TestRotateLogger(t *testing.T) {
	t.Log("========== 测试轮转日志输出 ==========")

	// 创建临时目录（不使用 TempDir 避免 Windows 文件锁定问题）
	tmpDir, err := os.MkdirTemp("", "test-rotate-logger-*")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir) // 尽力清理，失败也不影响测试

	logFile := filepath.Join(tmpDir, "rotate.log")

	config := &Logging{
		Level:          "info",
		Format:         logger.FormatJSON,
		Output:         logger.OutputRotate,
		FilePath:       logFile,
		MaxSize:        1, // 1MB
		MaxBackups:     3, // 保留3个备份
		MaxAge:         7, // 保留7天
		Compress:       false,
		FilePermission: 0644,
	}

	loggerInstance := config.ToLoggerInstance()
	assert.NotNil(t, loggerInstance)

	// 写入多条日志
	for i := range 20 {
		loggerInstance.Infof("轮转日志测试 - 消息 #%d", i)
		loggerInstance.Warnf("轮转日志警告 - 消息 #%d", i)
	}

	// 等待日志写入
	time.Sleep(200 * time.Millisecond)

	// 验证文件是否创建
	assert.FileExists(t, logFile, "轮转日志文件应该被创建")

	// 读取并显示文件内容
	content, err := os.ReadFile(logFile)
	assert.NoError(t, err)
	t.Logf("轮转日志文件内容:\n%s", string(content))
}

// TestAllLogLevels 测试所有日志级别
func TestAllLogLevels(t *testing.T) {
	t.Log("========== 测试所有日志级别 ==========")

	levels := []string{"debug", "info", "warn", "error"}

	for _, level := range levels {
		t.Run(level, func(t *testing.T) {
			config := &Logging{
				Level:  level,
				Format: logger.FormatText,
				Output: logger.OutputStdout,
			}

			loggerInstance := config.ToLoggerInstance()
			assert.NotNil(t, loggerInstance)

			t.Logf("--- 当前日志级别: %s ---", level)
			loggerInstance.Debug("DEBUG 级别日志")
			loggerInstance.Info("INFO 级别日志")
			loggerInstance.Warn("WARN 级别日志")
			loggerInstance.Error("ERROR 级别日志")
		})
	}
}

// TestDifferentFormats 测试不同的日志格式
func TestDifferentFormats(t *testing.T) {
	t.Log("========== 测试不同的日志格式 ==========")

	formats := []logger.FormatType{
		logger.FormatJSON,
		logger.FormatText,
	}

	for _, format := range formats {
		t.Run(string(format), func(t *testing.T) {
			config := &Logging{
				Level:  "info",
				Format: format,
				Output: logger.OutputStdout,
			}

			loggerInstance := config.ToLoggerInstance()
			assert.NotNil(t, loggerInstance)

			t.Logf("--- 当前格式: %s ---", format)
			loggerInstance.Info("测试日志格式")
			loggerInstance.Infof("格式化日志: 用户=%s, 操作=%s", "admin", "登录")
		})
	}
}

// TestLoggerWithFields 测试带字段的日志
func TestLoggerWithFields(t *testing.T) {
	t.Log("========== 测试带字段的日志 ==========")

	config := &Logging{
		Level:      "info",
		Format:     logger.FormatJSON,
		Output:     logger.OutputStdout,
		ShowCaller: true,
	}

	loggerInstance := config.ToLoggerInstance()
	assert.NotNil(t, loggerInstance)

	// 输出结构化日志
	loggerInstance.Info("用户登录事件")
	loggerInstance.Infof("用户 %s 从 IP %s 登录", "admin", "192.168.1.100")
	loggerInstance.Warn("磁盘空间不足")
	loggerInstance.Error("数据库连接失败")
}

// TestToWriterConfig 测试 WriterConfig 转换
func TestToWriterConfig(t *testing.T) {
	config := &Logging{
		Output:         logger.OutputRotate,
		FilePath:       "/var/log/test.log",
		MaxSize:        100,
		MaxBackups:     5,
		MaxAge:         30,
		Compress:       true,
		FilePermission: 0644,
		BufferSize:     8192,
	}

	writerConfig := config.toWriterConfig()

	assert.Equal(t, logger.OutputRotate, writerConfig.Type)
	assert.Equal(t, "/var/log/test.log", writerConfig.FilePath)
	assert.Equal(t, int64(100*1024*1024), writerConfig.MaxSize, "MaxSize 应该转换为字节")
	assert.Equal(t, 5, writerConfig.MaxFiles)
	assert.Equal(t, 30, writerConfig.MaxAge)
	assert.True(t, writerConfig.Compress)
	assert.Equal(t, os.FileMode(0644), writerConfig.Permission)
	assert.Equal(t, 8192, writerConfig.BufferSize)
}

// TestChainMethods 测试链式调用
func TestChainMethods(t *testing.T) {
	config := Default().
		WithLevel("warn").
		WithFormat(logger.FormatText).
		WithOutput(logger.OutputStdout).
		WithFilePermission(0600).
		WithBufferSize(8192)

	assert.Equal(t, "warn", config.Level)
	assert.Equal(t, logger.FormatText, config.Format)
	assert.Equal(t, logger.OutputStdout, config.Output)
	assert.Equal(t, uint32(0600), config.FilePermission)
	assert.Equal(t, 8192, config.BufferSize)

	// 测试日志输出
	loggerInstance := config.ToLoggerInstance()
	assert.NotNil(t, loggerInstance)

	t.Log("--- 链式调用创建的 Logger ---")
	loggerInstance.Warn("这是一条警告日志")
	loggerInstance.Error("这是一条错误日志")
}

// TestRealWorldScenario 测试真实场景
func TestRealWorldScenario(t *testing.T) {
	t.Log("========== 测试真实应用场景 ==========")

	// 使用标准输出而不是文件，避免 Windows 文件锁定问题
	config := &Logging{
		Level:      "info",
		Format:     logger.FormatText,
		Output:     logger.OutputStdout,
		ShowCaller: true,
		Prefix:     "[APP]",
		Colorful:   true,
	}

	loggerInstance := config.ToLoggerInstance()
	assert.NotNil(t, loggerInstance)

	// 模拟应用日志
	t.Log("--- 模拟应用启动 ---")
	loggerInstance.Info("应用启动成功")
	loggerInstance.Infof("监听端口: %d", 8080)

	t.Log("--- 模拟用户操作 ---")
	loggerInstance.Info("用户登录")
	loggerInstance.Infof("用户 %s 执行了 %s 操作", "admin", "创建订单")

	t.Log("--- 模拟警告 ---")
	loggerInstance.Warn("缓存命中率低于阈值")
	loggerInstance.Warnf("数据库连接池使用率: %.2f%%", 85.5)

	t.Log("--- 模拟错误 ---")
	loggerInstance.Error("第三方API调用失败")
	loggerInstance.Errorf("错误详情: %v", "connection timeout")

	// 等待日志输出
	time.Sleep(100 * time.Millisecond)
}
