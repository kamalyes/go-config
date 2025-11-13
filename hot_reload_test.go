/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 11:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-13 10:14:20
 * @FilePath: \go-config\hot_reload_test.go
 * @Description: 配置热更新功能测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConfig 测试配置结构
type TestConfig struct {
	App struct {
		Name    string `yaml:"name" json:"name"`
		Version string `yaml:"version" json:"version"`
		Debug   bool   `yaml:"debug" json:"debug"`
	} `yaml:"app" json:"app"`

	Server struct {
		Host string `yaml:"host" json:"host"`
		Port int    `yaml:"port" json:"port"`
	} `yaml:"server" json:"server"`
}

func TestHotReloadManager_Basic(t *testing.T) {
	// 创建临时配置文件
	configFile := createTempConfigFile(t, `
app:
  name: "test-app"
  version: "1.0.0"
  debug: true

server:
  host: "localhost"
  port: 8080
`)
	defer os.Remove(configFile)

	// 创建配置实例
	config := &TestConfig{}

	// 创建热更新器
	options := DefaultHotReloadConfig()
	options.DebounceDelay = 100 * time.Millisecond // 缩短防抖时间用于测试

	v, err := createViper(configFile)
	require.NoError(t, err)
	require.NoError(t, v.Unmarshal(config))

	hotReloader, err := NewHotReloader(config, v, configFile, options)
	require.NoError(t, err)

	// 启动热更新器
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = hotReloader.Start(ctx)
	require.NoError(t, err)
	defer hotReloader.Stop()

	// 验证初始配置
	assert.Equal(t, "test-app", config.App.Name)
	assert.Equal(t, "1.0.0", config.App.Version)
	assert.True(t, config.App.Debug)
	assert.Equal(t, "localhost", config.Server.Host)
	assert.Equal(t, 8080, config.Server.Port)

	// 注册回调
	var callbackTriggered bool
	var callbackMutex sync.Mutex

	err = hotReloader.RegisterCallback(func(ctx context.Context, event CallbackEvent) error {
		callbackMutex.Lock()
		defer callbackMutex.Unlock()
		callbackTriggered = true
		return nil
	}, CallbackOptions{
		ID:    "test_callback",
		Types: []CallbackType{CallbackTypeConfigChanged},
	})
	require.NoError(t, err)

	// 修改配置文件
	newContent := `
app:
  name: "updated-app"
  version: "2.0.0"
  debug: false

server:
  host: "0.0.0.0"
  port: 9090
`
	err = os.WriteFile(configFile, []byte(newContent), 0644)
	require.NoError(t, err)

	// 等待热更新触发
	time.Sleep(500 * time.Millisecond)

	// 检查回调是否被触发
	callbackMutex.Lock()
	triggered := callbackTriggered
	callbackMutex.Unlock()

	assert.True(t, triggered, "回调应该被触发")

	// 检查配置是否更新
	updatedConfig := hotReloader.GetConfig().(*TestConfig)
	assert.Equal(t, "updated-app", updatedConfig.App.Name)
	assert.Equal(t, "2.0.0", updatedConfig.App.Version)
	assert.False(t, updatedConfig.App.Debug)
	assert.Equal(t, "0.0.0.0", updatedConfig.Server.Host)
	assert.Equal(t, 9090, updatedConfig.Server.Port)
}

func TestEnvironmentManager_Callbacks(t *testing.T) {
	// 设置环境变量
	originalEnv := os.Getenv("APP_ENV")
	defer func() {
		if originalEnv != "" {
			os.Setenv("APP_ENV", originalEnv)
		} else {
			os.Unsetenv("APP_ENV")
		}
	}()

	os.Setenv("APP_ENV", "development")

	// 创建环境管理器
	env := NewEnvironment()
	defer env.StopWatch()

	// 注册回调
	var callbackTriggered bool
	var oldEnvReceived, newEnvReceived EnvironmentType
	var callbackMutex sync.Mutex

	err := env.RegisterCallback("test_env_callback", func(oldEnv, newEnv EnvironmentType) error {
		callbackMutex.Lock()
		defer callbackMutex.Unlock()
		callbackTriggered = true
		oldEnvReceived = oldEnv
		newEnvReceived = newEnv
		return nil
	}, 1, false)
	require.NoError(t, err)

	// 修改环境变量
	os.Setenv("APP_ENV", "production")

	// 手动触发环境检查
	env.CheckAndUpdateEnv()

	// 等待处理
	time.Sleep(100 * time.Millisecond)

	// 检查回调是否被触发
	callbackMutex.Lock()
	triggered := callbackTriggered
	oldEnv := oldEnvReceived
	newEnv := newEnvReceived
	callbackMutex.Unlock()

	assert.True(t, triggered, "环境变更回调应该被触发")
	assert.Equal(t, EnvDevelopment, oldEnv)
	assert.Equal(t, EnvProduction, newEnv)
}

func TestIntegratedConfigManager(t *testing.T) {
	// 创建临时配置文件
	configFile := createTempConfigFile(t, `
app:
  name: "integrated-test"
  version: "1.0.0"
  debug: true

server:
  host: "localhost"
  port: 3000
`)
	defer os.Remove(configFile)

	// 设置环境变量
	originalEnv := os.Getenv("APP_ENV")
	defer func() {
		if originalEnv != "" {
			os.Setenv("APP_ENV", originalEnv)
		} else {
			os.Unsetenv("APP_ENV")
		}
	}()

	os.Setenv("APP_ENV", "test")

	// 创建配置实例
	config := &TestConfig{}

	// 创建集成配置管理器
	manager, err := NewManager(&config).WithConfigPath(configFile).WithEnvironment(EnvTest).BuildAndStart()
	require.NoError(t, err)
	defer manager.Stop()

	// 验证初始状态
	assert.True(t, manager.IsRunning())
	assert.Equal(t, EnvTest, manager.GetEnvironment())

	// 获取配置
	currentConfig := manager.GetConfig().(*TestConfig)
	assert.Equal(t, "integrated-test", currentConfig.App.Name)
	assert.Equal(t, "1.0.0", currentConfig.App.Version)
	assert.True(t, currentConfig.App.Debug)

	// 测试上下文功能
	ctx := manager.WithContext(context.Background())

	// 从上下文获取环境
	env, ok := GetEnvironmentFromContext(ctx)
	assert.True(t, ok)
	assert.Equal(t, EnvTest, env)

	// 从上下文获取配置
	contextConfig, ok := GetConfigFromContext(ctx)
	assert.True(t, ok)
	assert.NotNil(t, contextConfig)

	// 注册回调并测试配置更新
	var callbackCount int
	var callbackMutex sync.Mutex

	err = manager.RegisterConfigCallback(func(ctx context.Context, event CallbackEvent) error {
		callbackMutex.Lock()
		defer callbackMutex.Unlock()
		callbackCount++
		return nil
	}, CallbackOptions{
		ID:    "integrated_test_callback",
		Types: []CallbackType{CallbackTypeConfigChanged},
	})
	require.NoError(t, err)

	// 手动重新加载配置
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = manager.ReloadConfig(ctx)
	require.NoError(t, err)

	// 等待回调触发
	time.Sleep(200 * time.Millisecond)

	// 检查回调计数
	callbackMutex.Lock()
	count := callbackCount
	callbackMutex.Unlock()

	assert.Greater(t, count, 0, "应该有回调被触发")
}

func TestContextManager(t *testing.T) {
	// 创建环境管理器
	env := NewEnvironment()
	defer env.StopWatch()

	// 创建上下文管理器
	contextManager := NewContextManager(env, nil)

	// 测试配置上下文
	configCtx := contextManager.GetConfigContext()
	assert.NotNil(t, configCtx)
	assert.NotNil(t, configCtx.Metadata)

	// 设置元数据
	contextManager.SetMetadata("test_key", "test_value")

	// 获取元数据
	value, exists := contextManager.GetMetadata("test_key")
	assert.True(t, exists)
	assert.Equal(t, "test_value", value)

	// 测试上下文创建
	ctx := contextManager.WithConfig(context.Background())
	assert.NotNil(t, ctx)

	// 从上下文获取环境
	contextEnv, ok := GetEnvironmentFromContext(ctx)
	assert.True(t, ok)
	assert.Equal(t, env.Value, contextEnv)
}

func TestContextHelpers(t *testing.T) {
	// 清理全局状态
	ClearGlobalContextManager()

	// 设置环境变量
	originalEnv := os.Getenv("APP_ENV")
	defer func() {
		if originalEnv != "" {
			os.Setenv("APP_ENV", originalEnv)
		} else {
			os.Unsetenv("APP_ENV")
		}
		// 测试结束后清理全局状态
		ClearGlobalContextManager()
	}()

	// 先设置测试环境
	os.Setenv("APP_ENV", "development")

	// 创建新的环境实例来确保正确读取环境变量
	testEnv := NewEnvironment()
	defer testEnv.StopWatch()

	// 验证环境被正确设置
	assert.Equal(t, EnvDevelopment, testEnv.Value, "环境应该是development")

	InitializeContextManager(testEnv, nil)

	// 测试上下文辅助工具
	ctx := ContextHelper.NewConfigContext()
	assert.NotNil(t, ctx)

	// 测试环境检查 - 应该检查上下文中的环境
	contextEnv, ok := GetEnvironmentFromContext(ctx)
	assert.True(t, ok, "应该能从上下文中获取环境信息")
	assert.Equal(t, EnvDevelopment, contextEnv, "上下文中的环境应该是development")

	// 检查环境是否匹配
	isDevEnv := ContextHelper.IsEnvironment(ctx, contextEnv)
	assert.True(t, isDevEnv, "上下文环境应该匹配")

	// 也可以直接检查是否是development环境
	isDev := ContextHelper.IsEnvironment(ctx, EnvDevelopment)
	assert.True(t, isDev, "应该是development环境，实际环境: %v", contextEnv)

	// 测试带超时的上下文
	timeoutCtx, cancel := ContextHelper.NewContextWithTimeout(5 * time.Second)
	defer cancel()
	assert.NotNil(t, timeoutCtx)
} // createTempConfigFile 创建临时配置文件
func createTempConfigFile(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	err := os.WriteFile(configFile, []byte(content), 0644)
	require.NoError(t, err)

	return configFile
}

// createViper 创建Viper实例
func createViper(configFile string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(configFile)
	return v, v.ReadInConfig()
}

// 基准测试

func BenchmarkHotReloader_TriggerCallbacks(b *testing.B) {
	// 创建临时配置文件
	configFile := createTempConfigFileB(b, `
app:
  name: "benchmark-test"
  version: "1.0.0"

server:
  host: "localhost"
  port: 8080
`)
	defer os.Remove(configFile)

	config := &TestConfig{}
	v, _ := createViper(configFile)
	v.Unmarshal(config)

	hotReloader, _ := NewHotReloader(config, v, configFile, DefaultHotReloadConfig())

	// 注册多个回调
	for i := 0; i < 10; i++ {
		hotReloader.RegisterCallback(func(ctx context.Context, event CallbackEvent) error {
			return nil
		}, CallbackOptions{
			ID:    fmt.Sprintf("benchmark_callback_%d", i),
			Types: []CallbackType{CallbackTypeConfigChanged},
		})
	}

	event := CallbackEvent{
		Type:        CallbackTypeConfigChanged,
		Timestamp:   time.Now(),
		Source:      "benchmark",
		Environment: EnvTest,
	}

	ctx := context.Background()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hotReloader.TriggerCallbacks(ctx, event)
		}
	})
}

// createTempConfigFileB for benchmark (without testing.T)
func createTempConfigFileB(b *testing.B, content string) string {
	tmpDir := b.TempDir()
	configFile := filepath.Join(tmpDir, "benchmark_config.yaml")

	err := os.WriteFile(configFile, []byte(content), 0644)
	if err != nil {
		b.Fatal(err)
	}

	return configFile
}
