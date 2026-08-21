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

	// 验证配置已更新（reloadConfig 创建新结构体替换，须读管理器权威配置）
	assert.Equal(t, 9090, manager.GetConfig().(*AppConfig).Server.Port)
}

// TestIntegratedConfigManager_DefaultOptions 测试默认选项
func TestIntegratedConfigManager_DefaultOptions(t *testing.T) {
	options := DefaultIntegratedConfigOptions()

	assert.NotNil(t, options)
	assert.Equal(t, DefaultEnv, options.Environment)
	assert.NotNil(t, options.HotReloadConfig)
	assert.NotNil(t, options.ContextOptions)
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

// createTestManager 创建用于测试的集成配置管理器
func createTestManager(t *testing.T) *IntegratedConfigManager {
	t.Helper()
	configContent := `
server:
  host: localhost
  port: 8080
`
	configPath := createTestConfigFile(t, configContent)
	config := &AppConfig{}
	manager, err := NewIntegratedConfigManager(config, &IntegratedConfigOptions{
		ConfigPath:  configPath,
		Environment: EnvDevelopment,
	})
	require.NoError(t, err)
	return manager
}

// --- GetConfigAs ---

func TestIntegratedConfigManager_GetConfigAs_OK(t *testing.T) {
	manager := createTestManager(t)
	defer manager.Stop()

	cfg, err := GetConfigAs[AppConfig](manager)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, 8080, cfg.Server.Port)
}

func TestIntegratedConfigManager_GetConfigAs_Mismatch(t *testing.T) {
	manager := createTestManager(t)
	defer manager.Stop()

	// 类型不匹配
	cfg, err := GetConfigAs[ServerConfig](manager)
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

// --- GetViper / GetContextManager ---

func TestIntegratedConfigManager_GetViperAndContextManager(t *testing.T) {
	manager := createTestManager(t)
	defer manager.Stop()

	assert.NotNil(t, manager.GetViper())
	assert.NotNil(t, manager.GetContextManager())
}

// --- 回调注册/注销 ---

func TestIntegratedConfigManager_RegisterUnregisterCallbacks(t *testing.T) {
	manager := createTestManager(t)
	defer manager.Stop()

	// 配置回调
	cfgCb := func(ctx context.Context, event CallbackEvent) error { return nil }
	require.NoError(t, manager.RegisterConfigCallback(cfgCb, CallbackOptions{ID: "icm-cfg-cb"}))
	assert.NoError(t, manager.UnregisterConfigCallback("icm-cfg-cb")) // 首次注销成功
	// 重复注销返回错误
	assert.Error(t, manager.UnregisterConfigCallback("icm-cfg-cb"))

	// 环境回调
	envCb := func(old, newEnv EnvironmentType) error { return nil }
	require.NoError(t, manager.RegisterEnvironmentCallback("icm-env-cb", envCb, 1, false))
	require.NoError(t, manager.UnregisterEnvironmentCallback("icm-env-cb"))
	assert.Error(t, manager.UnregisterEnvironmentCallback("icm-env-cb"))
}

// --- ValidateConfig ---

func TestIntegratedConfigManager_ValidateConfig_OK(t *testing.T) {
	manager := createTestManager(t)
	defer manager.Stop()

	assert.NoError(t, manager.ValidateConfig())
}

func TestIntegratedConfigManager_ValidateConfig_Empty(t *testing.T) {
	manager := createTestManager(t)
	defer manager.Stop()

	// 同包直接置空 config 触发 ErrConfigEmpty 分支
	manager.mu.Lock()
	origConfig := manager.config
	manager.config = nil
	manager.mu.Unlock()

	assert.ErrorIs(t, manager.ValidateConfig(), ErrConfigEmpty)

	// 恢复
	manager.mu.Lock()
	manager.config = origConfig
	manager.mu.Unlock()
}

// --- GetConfigMetadata ---

func TestIntegratedConfigManager_GetConfigMetadata(t *testing.T) {
	manager := createTestManager(t)
	defer manager.Stop()

	meta := manager.GetConfigMetadata()
	assert.NotNil(t, meta)
	assert.Contains(t, meta, "config_path")
	assert.Contains(t, meta, "environment")
	assert.Contains(t, meta, "running")
	assert.Contains(t, meta, "hot_reload_enabled")
	assert.Contains(t, meta, "created_at")
	assert.Contains(t, meta, "updated_at")
	assert.False(t, meta["running"].(bool))
}

// --- MustStart ---

func TestIntegratedConfigManager_MustStart_OK(t *testing.T) {
	manager := createTestManager(t)
	// 禁用热重载避免后台 goroutine
	manager.hotReloadConfig.Enabled = false
	assert.NotPanics(t, func() {
		manager.MustStart(context.Background())
	})
	defer manager.Stop()
}

func TestIntegratedConfigManager_MustStart_Panic(t *testing.T) {
	manager := createTestManager(t)
	manager.hotReloadConfig.Enabled = false
	require.NoError(t, manager.Start(context.Background()))
	defer manager.Stop()

	// 已运行再次 MustStart 应 panic
	assert.Panics(t, func() {
		manager.MustStart(context.Background())
	})
}

// --- MustGetConfigAs ---

func TestIntegratedConfigManager_MustGetConfigAs_OK(t *testing.T) {
	manager := createTestManager(t)
	defer manager.Stop()

	cfg := MustGetConfigAs[AppConfig](manager)
	require.NotNil(t, cfg)
	assert.Equal(t, 8080, cfg.Server.Port)
}

func TestIntegratedConfigManager_MustGetConfigAs_Panic(t *testing.T) {
	manager := createTestManager(t)
	defer manager.Stop()

	assert.Panics(t, func() {
		MustGetConfigAs[ServerConfig](manager)
	})
}

// --- CreateIntegratedManager ---

func TestCreateIntegratedManager(t *testing.T) {
	configContent := `
server:
  host: localhost
  port: 9090
`
	configPath := createTestConfigFile(t, configContent)
	config := &AppConfig{}

	manager, err := CreateIntegratedManager(config, configPath, EnvDevelopment)
	require.NoError(t, err)
	require.NotNil(t, manager)
	defer manager.Stop()
	assert.Equal(t, 9090, config.Server.Port)
}

// --- onError 错误回调 ---

func TestIntegratedConfigManager_OnError(t *testing.T) {
	manager := createTestManager(t)
	defer manager.Stop()

	triggered := false
	errCb := func(ctx context.Context, event CallbackEvent) error {
		if event.Type == CallbackTypeError {
			triggered = true
		}
		return nil
	}
	require.NoError(t, manager.RegisterConfigCallback(errCb, CallbackOptions{
		ID:    "onerr-test-cb",
		Types: []CallbackType{CallbackTypeError},
	}))

	// 直接触发错误事件以执行内部 onError 回调
	errEvent := CreateErrorEvent("test_source", assertSimpleErr("boom"))
	require.NoError(t, manager.GetHotReloader().TriggerCallbacks(context.Background(), errEvent))
	assert.True(t, triggered)
}

// --- ScanAndDisplayConfigs ---

func TestScanAndDisplayConfigs(t *testing.T) {
	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(tmpDir, "config.yaml"), []byte("k: v"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(tmpDir, "config-dev.yaml"), []byte("k: v"), 0644))

	files, err := ScanAndDisplayConfigs(tmpDir, EnvDevelopment)
	require.NoError(t, err)
	assert.NotEmpty(t, files)
}

func TestScanAndDisplayConfigs_NotExist(t *testing.T) {
	files, err := ScanAndDisplayConfigs("/non/existent/path/xyz", EnvDevelopment)
	assert.Error(t, err)
	assert.Nil(t, files)
}

// assertSimpleErr 简单 error 类型
type assertSimpleErr string

func (e assertSimpleErr) Error() string { return string(e) }
