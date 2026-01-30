/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-01-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-01-30 00:00:00
 * @FilePath: \go-config\config_builder_test.go
 * @Description: 配置构建器测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewConfigBuilder 测试创建配置构建器
func TestNewConfigBuilder(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	assert.NotNil(t, builder)
}

// TestConfigBuilder_WithConfigPath 测试设置配置文件路径
func TestConfigBuilder_WithConfigPath(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	path := "config.yaml"
	result := builder.WithConfigPath(path)

	assert.NotNil(t, result)
	assert.Equal(t, builder, result) // 验证链式调用
}

// TestConfigBuilder_WithSearchPath 测试设置搜索路径
func TestConfigBuilder_WithSearchPath(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	searchPath := "./configs"
	result := builder.WithSearchPath(searchPath)

	assert.NotNil(t, result)
	assert.Equal(t, builder, result)
}

// TestConfigBuilder_WithEnvironment 测试设置环境
func TestConfigBuilder_WithEnvironment(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	result := builder.WithEnvironment(EnvDevelopment)

	assert.NotNil(t, result)
	assert.Equal(t, builder, result)
}

// TestConfigBuilder_WithPrefix 测试设置配置文件前缀
func TestConfigBuilder_WithPrefix(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	prefix := "app"
	result := builder.WithPrefix(prefix)

	assert.NotNil(t, result)
	assert.Equal(t, builder, result)
}

// TestConfigBuilder_WithPattern 测试设置文件匹配模式
func TestConfigBuilder_WithPattern(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	pattern := "*.yaml"
	result := builder.WithPattern(pattern)

	assert.NotNil(t, result)
	assert.Equal(t, builder, result)
}

// TestConfigBuilder_WithHotReload 测试启用热重载
func TestConfigBuilder_WithHotReload(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	hotReloadConfig := &HotReloadConfig{
		Enabled:       true,
		WatchInterval: 5 * time.Second,
	}
	result := builder.WithHotReload(hotReloadConfig)

	assert.NotNil(t, result)
	assert.Equal(t, builder, result)
}

// TestConfigBuilder_WithContext 测试设置上下文选项
func TestConfigBuilder_WithContext(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	contextOptions := &ContextKeyOptions{
		Value: EnvDevelopment,
	}
	result := builder.WithContext(contextOptions)

	assert.NotNil(t, result)
	assert.Equal(t, builder, result)
}

// TestConfigBuilder_WithErrorHandler 测试设置错误处理器
func TestConfigBuilder_WithErrorHandler(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	handler := NewErrorHandler()
	result := builder.WithErrorHandler(handler)

	assert.NotNil(t, result)
	assert.Equal(t, builder, result)
}

// TestConfigBuilder_Build 测试构建配置管理器
func TestConfigBuilder_Build(t *testing.T) {
	// 创建临时配置文件
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	configContent := `
name: test-app
port: 8080
enabled: true
`
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	require.NoError(t, err)

	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	manager, err := builder.
		WithConfigPath(configPath).
		WithEnvironment(EnvDevelopment).
		Build()

	require.NoError(t, err)
	assert.NotNil(t, manager)
}

// TestConfigBuilder_BuildAndStart 测试构建并启动配置管理器
func TestConfigBuilder_BuildAndStart(t *testing.T) {
	// 创建临时配置文件
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	configContent := `
name: test-app
port: 8080
enabled: true
`
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	require.NoError(t, err)

	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	ctx := context.Background()
	manager, err := builder.
		WithConfigPath(configPath).
		WithEnvironment(EnvDevelopment).
		BuildAndStart(ctx)

	require.NoError(t, err)
	assert.NotNil(t, manager)

	// 清理
	if manager != nil {
		manager.Stop()
	}
}

// TestConfigBuilder_ChainedCalls 测试链式调用
func TestConfigBuilder_ChainedCalls(t *testing.T) {
	// 创建临时配置文件
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	configContent := `
name: test-app
port: 8080
enabled: true
`
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	require.NoError(t, err)

	config := &TestConfig{}

	// 测试完整的链式调用（移除 WithPrefix 避免触发自动发现）
	manager, err := NewConfigBuilder(config).
		WithConfigPath(configPath).
		WithEnvironment(EnvDevelopment).
		WithHotReload(&HotReloadConfig{
			Enabled:       false,
			WatchInterval: 5 * time.Second,
		}).
		WithContext(&ContextKeyOptions{
			Value: EnvDevelopment,
		}).
		WithErrorHandler(NewErrorHandler()).
		Build()

	require.NoError(t, err)
	assert.NotNil(t, manager)
}

// TestConfigBuilder_InvalidConfigPath 测试无效的配置文件路径
func TestConfigBuilder_InvalidConfigPath(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	manager, err := builder.
		WithConfigPath("/non/existent/path/config.yaml").
		Build()

	assert.Error(t, err)
	assert.Nil(t, manager)
}

// TestConfigBuilder_MultipleEnvironments 测试多环境配置
func TestConfigBuilder_MultipleEnvironments(t *testing.T) {
	// 创建临时配置文件
	tmpDir := t.TempDir()

	environments := []EnvironmentType{
		EnvDevelopment,
		EnvTest,
		EnvProduction,
	}

	for _, env := range environments {
		configPath := filepath.Join(tmpDir, string(env)+".yaml")
		configContent := `
name: test-app
port: 8080
enabled: true
`
		err := os.WriteFile(configPath, []byte(configContent), 0644)
		require.NoError(t, err)

		config := &TestConfig{}
		manager, err := NewConfigBuilder(config).
			WithConfigPath(configPath).
			WithEnvironment(env).
			Build()

		require.NoError(t, err)
		assert.NotNil(t, manager)
	}
}

// TestConfigBuilder_WithSearchPathAutoDiscovery 测试自动发现模式
func TestConfigBuilder_WithSearchPathAutoDiscovery(t *testing.T) {
	// 创建临时目录和配置文件
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	configContent := `
name: test-app
port: 8080
enabled: true
`
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	require.NoError(t, err)

	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	manager, err := builder.
		WithSearchPath(tmpDir).
		WithEnvironment(EnvDevelopment).
		Build()

	require.NoError(t, err)
	assert.NotNil(t, manager)
}

// TestConfigBuilder_WithHotReloadEnabled 测试启用热重载功能
func TestConfigBuilder_WithHotReloadEnabled(t *testing.T) {
	// 创建临时配置文件
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	configContent := `
name: test-app
port: 8080
enabled: true
`
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	require.NoError(t, err)

	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	hotReloadConfig := &HotReloadConfig{
		Enabled:       true,
		WatchInterval: 1 * time.Second,
	}

	ctx := context.Background()
	manager, err := builder.
		WithConfigPath(configPath).
		WithEnvironment(EnvDevelopment).
		WithHotReload(hotReloadConfig).
		BuildAndStart(ctx)

	require.NoError(t, err)
	assert.NotNil(t, manager)

	// 清理
	if manager != nil {
		manager.Stop()
	}
}

// TestConfigBuilder_NilConfig 测试空配置指针
func TestConfigBuilder_NilConfig(t *testing.T) {
	var config *TestConfig
	builder := NewConfigBuilder(config)

	assert.NotNil(t, builder)
}

// TestConfigBuilder_EmptyBuilder 测试空构建器
func TestConfigBuilder_EmptyBuilder(t *testing.T) {
	config := &TestConfig{}
	builder := NewConfigBuilder(config)

	// 不设置任何选项直接构建
	manager, err := builder.Build()

	// 应该返回错误，因为没有配置文件路径
	assert.Error(t, err)
	assert.Nil(t, manager)
}
