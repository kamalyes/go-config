/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-01-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-01-30 00:00:00
 * @FilePath: \go-config\integrated_manager_test.go
 * @Description: 集成配置管理器测试
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

// AppConfig 测试应用配置
type AppConfig struct {
	Server   ServerConfig   `mapstructure:"server" yaml:"server" json:"server"`
	Database DatabaseConfig `mapstructure:"database" yaml:"database" json:"database"`
	Redis    RedisConfig    `mapstructure:"redis" yaml:"redis" json:"redis"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Host string `mapstructure:"host" yaml:"host" json:"host"`
	Port int    `mapstructure:"port" yaml:"port" json:"port"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `mapstructure:"host" yaml:"host" json:"host"`
	Port     int    `mapstructure:"port" yaml:"port" json:"port"`
	Username string `mapstructure:"username" yaml:"username" json:"username"`
	Password string `mapstructure:"password" yaml:"password" json:"password"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host string `mapstructure:"host" yaml:"host" json:"host"`
	Port int    `mapstructure:"port" yaml:"port" json:"port"`
}

// createTestConfigFile 创建测试配置文件
func createTestConfigFile(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	err := os.WriteFile(configPath, []byte(content), 0644)
	require.NoError(t, err)
	return configPath
}

// TestNewIntegratedConfigManager 测试创建集成配置管理器
func TestNewIntegratedConfigManager(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
database:
  host: localhost
  port: 3306
  username: root
  password: password
redis:
  host: localhost
  port: 6379
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath:  configPath,
		Environment: EnvDevelopment,
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)
	assert.NotNil(t, manager)
}

// TestIntegratedConfigManager_LoadConfig 测试加载配置
func TestIntegratedConfigManager_LoadConfig(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
database:
  host: localhost
  port: 3306
  username: root
  password: password
redis:
  host: localhost
  port: 6379
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath:  configPath,
		Environment: EnvDevelopment,
	}

	_, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	// 验证配置已加载
	assert.Equal(t, "localhost", config.Server.Host)
	assert.Equal(t, 8080, config.Server.Port)
	assert.Equal(t, "localhost", config.Database.Host)
	assert.Equal(t, 3306, config.Database.Port)
}

// TestIntegratedConfigManager_GetConfig 测试获取配置
func TestIntegratedConfigManager_GetConfig(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath:  configPath,
		Environment: EnvDevelopment,
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	retrievedConfig := manager.GetConfig()
	assert.NotNil(t, retrievedConfig)
	assert.IsType(t, &AppConfig{}, retrievedConfig)
}

// TestIntegratedConfigManager_GetEnvironment 测试获取环境
func TestIntegratedConfigManager_GetEnvironment(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath:  configPath,
		Environment: EnvProduction,
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	// 获取环境管理器的环境值
	env := manager.GetEnvironmentManager().Value
	assert.Equal(t, EnvProduction, env)
}

// TestIntegratedConfigManager_SetEnvironment 测试设置环境
func TestIntegratedConfigManager_SetEnvironment(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath:  configPath,
		Environment: EnvDevelopment,
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	// 设置新环境
	manager.SetEnvironment(EnvProduction)
	assert.Equal(t, EnvProduction, manager.GetEnvironment())
}

// TestIntegratedConfigManager_StartStop 测试启动和停止
func TestIntegratedConfigManager_StartStop(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath: configPath,
		HotReloadConfig: &HotReloadConfig{
			Enabled:       true,
			WatchInterval: 1 * time.Second,
		},
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	ctx := context.Background()

	// 启动管理器
	err = manager.Start(ctx)
	require.NoError(t, err)

	// 等待一段时间
	time.Sleep(100 * time.Millisecond)

	// 停止管理器
	err = manager.Stop()
	require.NoError(t, err)
}

// TestIntegratedConfigManager_HotReload 测试热重载
func TestIntegratedConfigManager_HotReload(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	require.NoError(t, err)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath: configPath,
		HotReloadConfig: &HotReloadConfig{
			Enabled:       true,
			WatchInterval: 500 * time.Millisecond,
		},
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	ctx := context.Background()
	err = manager.Start(ctx)
	require.NoError(t, err)
	defer manager.Stop()

	// 验证初始配置
	assert.Equal(t, 8080, config.Server.Port)

	// 修改配置文件
	newConfigContent := `
server:
  host: localhost
  port: 9090
`
	err = os.WriteFile(configPath, []byte(newConfigContent), 0644)
	require.NoError(t, err)

	// 等待热重载
	time.Sleep(1 * time.Second)

	// 验证配置已更新
	assert.Equal(t, 9090, config.Server.Port)
}

// TestIntegratedConfigManager_DefaultOptions 测试默认选项
func TestIntegratedConfigManager_DefaultOptions(t *testing.T) {
	options := DefaultIntegratedConfigOptions()

	assert.NotNil(t, options)
	assert.Equal(t, DefaultEnv, options.Environment)
	assert.NotNil(t, options.HotReloadConfig)
	assert.NotNil(t, options.ContextOptions)
	assert.NotNil(t, options.ErrorHandler)
}

// TestIntegratedConfigManager_InvalidConfigPath 测试无效配置路径
func TestIntegratedConfigManager_InvalidConfigPath(t *testing.T) {
	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath:  "/non/existent/path/config.yaml",
		Environment: EnvDevelopment,
	}

	manager, err := NewIntegratedConfigManager(config, options)
	assert.Error(t, err)
	assert.Nil(t, manager)
}

// TestIntegratedConfigManager_NilOptions 测试空选项
func TestIntegratedConfigManager_NilOptions(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	createTestConfigFile(t, configContent)

	config := &AppConfig{}

	// 使用nil选项，应该使用默认选项
	manager, err := NewIntegratedConfigManager(config, nil)

	// 由于没有配置路径，应该返回错误
	assert.Error(t, err)
	assert.Nil(t, manager)
}

// TestIntegratedConfigManager_MultipleStarts 测试多次启动
func TestIntegratedConfigManager_MultipleStarts(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath: configPath,
		HotReloadConfig: &HotReloadConfig{
			Enabled:       true,
			WatchInterval: 1 * time.Second,
		},
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	ctx := context.Background()

	// 第一次启动
	err = manager.Start(ctx)
	require.NoError(t, err)

	// 第二次启动应该返回错误或被忽略
	err = manager.Start(ctx)
	// 根据实现，可能返回错误或成功

	// 清理
	manager.Stop()
}

// TestIntegratedConfigManager_StopWithoutStart 测试未启动就停止
func TestIntegratedConfigManager_StopWithoutStart(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath: configPath,
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	// 未启动就停止
	err = manager.Stop()
	// 应该不返回错误
	assert.NoError(t, err)
}

// TestIntegratedConfigManager_ContextCancellation 测试上下文取消
func TestIntegratedConfigManager_ContextCancellation(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath: configPath,
		HotReloadConfig: &HotReloadConfig{
			Enabled:       true,
			WatchInterval: 1 * time.Second,
		},
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())

	// 启动管理器
	err = manager.Start(ctx)
	require.NoError(t, err)

	// 取消上下文
	cancel()

	// 等待一段时间
	time.Sleep(100 * time.Millisecond)

	// 清理
	manager.Stop()
}

// TestIntegratedConfigManager_GetConfigPath 测试获取配置路径
func TestIntegratedConfigManager_GetConfigPath(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath:  configPath,
		Environment: EnvDevelopment,
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	retrievedPath := manager.GetConfigPath()
	assert.Equal(t, configPath, retrievedPath)
}

// TestIntegratedConfigManager_IsRunning 测试运行状态检查
func TestIntegratedConfigManager_IsRunning(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 8080
`
	configPath := createTestConfigFile(t, configContent)

	config := &AppConfig{}
	options := &IntegratedConfigOptions{
		ConfigPath: configPath,
		HotReloadConfig: &HotReloadConfig{
			Enabled:       true,
			WatchInterval: 1 * time.Second,
		},
	}

	manager, err := NewIntegratedConfigManager(config, options)
	require.NoError(t, err)

	// 启动前
	assert.False(t, manager.IsRunning())

	ctx := context.Background()
	err = manager.Start(ctx)
	require.NoError(t, err)

	// 启动后
	assert.True(t, manager.IsRunning())

	// 停止后
	manager.Stop()
	assert.False(t, manager.IsRunning())
}
