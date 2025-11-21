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

	"github.com/mitchellh/mapstructure"
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

// TestDetectConfigTag_YAML 测试 YAML 文件类型检测和 tag 映射
func TestDetectConfigTag_YAML(t *testing.T) {
	// 创建临时 YAML 配置文件（使用 kebab-case）
	configContent := `
app:
  name: "yaml-test-app"
  version: "2.0.0"
  debug: false

server:
  host: "0.0.0.0"
  port: 9090

database:
  max-idle-conns: 15
  max-open-conns: 150
  timeout: 45s
`
	configFile := createTempConfigFile(t, configContent)

	// 创建 Viper 实例
	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	require.NoError(t, err)

	// 创建配置结构（包含 kebab-case 字段）
	type DBConfig struct {
		MaxIdleConns int           `yaml:"max-idle-conns" json:"maxIdleConns"`
		MaxOpenConns int           `yaml:"max-open-conns" json:"maxOpenConns"`
		Timeout      time.Duration `yaml:"timeout" json:"timeout"`
	}

	type YAMLTestConfig struct {
		App struct {
			Name    string `yaml:"name" json:"name"`
			Version string `yaml:"version" json:"version"`
			Debug   bool   `yaml:"debug" json:"debug"`
		} `yaml:"app" json:"app"`
		Server struct {
			Host string `yaml:"host" json:"host"`
			Port int    `yaml:"port" json:"port"`
		} `yaml:"server" json:"server"`
		Database DBConfig `yaml:"database" json:"database"`
	}

	config := &YAMLTestConfig{}
	reloader, err := NewHotReloader(config, v, configFile, &HotReloadConfig{
		Enabled: false,
	})
	require.NoError(t, err)

	manager := reloader.(*hotReloadManager)

	// 测试 1: 检测文件类型为 YAML
	tagName := manager.detectConfigTag()
	assert.Equal(t, "yaml", tagName, "应该检测到 YAML 配置文件")

	// 测试 2: 使用 yaml tag 解析配置
	err = v.Unmarshal(config, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = tagName
		dc.WeaklyTypedInput = true
	})
	require.NoError(t, err)

	// 测试 3: 验证 kebab-case 字段正确映射
	assert.Equal(t, "yaml-test-app", config.App.Name)
	assert.Equal(t, "2.0.0", config.App.Version)
	assert.False(t, config.App.Debug)
	assert.Equal(t, "0.0.0.0", config.Server.Host)
	assert.Equal(t, 9090, config.Server.Port)
	assert.Equal(t, 15, config.Database.MaxIdleConns, "max-idle-conns 应该映射到 MaxIdleConns")
	assert.Equal(t, 150, config.Database.MaxOpenConns, "max-open-conns 应该映射到 MaxOpenConns")
	assert.Equal(t, 45*time.Second, config.Database.Timeout)

	t.Logf("✅ YAML 文件检测和 kebab-case 映射测试通过")
}

// TestDetectConfigTag_JSON 测试 JSON 文件类型检测和 tag 映射
func TestDetectConfigTag_JSON(t *testing.T) {
	// 创建临时 JSON 配置文件（使用 camelCase）
	tmpDir := t.TempDir()
	jsonPath := filepath.Join(tmpDir, "test.json")

	jsonContent := `{
  "app": {
    "name": "json-test-app",
    "version": "3.0.0",
    "debug": true
  },
  "server": {
    "host": "127.0.0.1",
    "port": 8888
  },
  "database": {
    "maxIdleConns": 25,
    "maxOpenConns": 250,
    "timeout": "60s"
  }
}`
	err := os.WriteFile(jsonPath, []byte(jsonContent), 0644)
	require.NoError(t, err)

	// 创建 Viper 实例
	v := viper.New()
	v.SetConfigFile(jsonPath)
	err = v.ReadInConfig()
	require.NoError(t, err)

	// 创建配置结构
	type DBConfig struct {
		MaxIdleConns int           `yaml:"max-idle-conns" json:"maxIdleConns"`
		MaxOpenConns int           `yaml:"max-open-conns" json:"maxOpenConns"`
		Timeout      time.Duration `yaml:"timeout" json:"timeout"`
	}

	type JSONTestConfig struct {
		App struct {
			Name    string `yaml:"name" json:"name"`
			Version string `yaml:"version" json:"version"`
			Debug   bool   `yaml:"debug" json:"debug"`
		} `yaml:"app" json:"app"`
		Server struct {
			Host string `yaml:"host" json:"host"`
			Port int    `yaml:"port" json:"port"`
		} `yaml:"server" json:"server"`
		Database DBConfig `yaml:"database" json:"database"`
	}

	config := &JSONTestConfig{}
	reloader, err := NewHotReloader(config, v, jsonPath, &HotReloadConfig{
		Enabled: false,
	})
	require.NoError(t, err)

	manager := reloader.(*hotReloadManager)

	// 测试 1: 检测文件类型为 JSON
	tagName := manager.detectConfigTag()
	assert.Equal(t, "json", tagName, "应该检测到 JSON 配置文件")

	// 测试 2: 使用 json tag 解析配置
	err = v.Unmarshal(config, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = tagName
		dc.WeaklyTypedInput = true
	})
	require.NoError(t, err)

	// 测试 3: 验证 camelCase 字段正确映射
	assert.Equal(t, "json-test-app", config.App.Name)
	assert.Equal(t, "3.0.0", config.App.Version)
	assert.True(t, config.App.Debug)
	assert.Equal(t, "127.0.0.1", config.Server.Host)
	assert.Equal(t, 8888, config.Server.Port)
	assert.Equal(t, 25, config.Database.MaxIdleConns, "maxIdleConns 应该映射到 MaxIdleConns")
	assert.Equal(t, 250, config.Database.MaxOpenConns, "maxOpenConns 应该映射到 MaxOpenConns")
	assert.Equal(t, 60*time.Second, config.Database.Timeout)

	t.Logf("✅ JSON 文件检测和 camelCase 映射测试通过")
}

// TestDetectConfigTag_YML 测试 .yml 扩展名识别
func TestDetectConfigTag_YML(t *testing.T) {
	tmpDir := t.TempDir()
	ymlPath := filepath.Join(tmpDir, "test.yml")

	ymlContent := `
app:
  name: "yml-test"
server:
  host: "localhost"
  port: 3000
`
	err := os.WriteFile(ymlPath, []byte(ymlContent), 0644)
	require.NoError(t, err)

	v := viper.New()
	v.SetConfigFile(ymlPath)
	err = v.ReadInConfig()
	require.NoError(t, err)

	config := &TestConfig{}
	reloader, err := NewHotReloader(config, v, ymlPath, nil)
	require.NoError(t, err)

	manager := reloader.(*hotReloadManager)
	tagName := manager.detectConfigTag()

	assert.Equal(t, "yaml", tagName, ".yml 扩展名应该被识别为 YAML 文件")
	t.Logf("✅ .yml 扩展名识别测试通过")
}

// TestDetectConfigTag_TOML 测试 TOML 文件类型检测
func TestDetectConfigTag_TOML(t *testing.T) {
	tmpDir := t.TempDir()
	tomlPath := filepath.Join(tmpDir, "test.toml")

	tomlContent := `
[app]
name = "toml-test"

[server]
host = "localhost"
port = 8080
`
	err := os.WriteFile(tomlPath, []byte(tomlContent), 0644)
	require.NoError(t, err)

	v := viper.New()
	v.SetConfigFile(tomlPath)
	err = v.ReadInConfig()
	require.NoError(t, err)

	config := &TestConfig{}
	reloader, err := NewHotReloader(config, v, tomlPath, nil)
	require.NoError(t, err)

	manager := reloader.(*hotReloadManager)
	tagName := manager.detectConfigTag()

	assert.Equal(t, "toml", tagName, "应该检测到 TOML 配置文件")
	t.Logf("✅ TOML 文件类型检测测试通过")
}

// TestDetectConfigTag_Unknown 测试未知文件类型默认处理
func TestDetectConfigTag_Unknown(t *testing.T) {
	tmpDir := t.TempDir()
	unknownPath := filepath.Join(tmpDir, "test.conf")

	err := os.WriteFile(unknownPath, []byte("test content"), 0644)
	require.NoError(t, err)

	v := viper.New()
	config := &TestConfig{}

	reloader, err := NewHotReloader(config, v, unknownPath, nil)
	require.NoError(t, err)

	manager := reloader.(*hotReloadManager)
	tagName := manager.detectConfigTag()

	assert.Equal(t, "yaml", tagName, "未知文件类型应该默认使用 yaml tag")
	t.Logf("✅ 未知文件类型默认处理测试通过")
}

// TestReloadConfig_AutoDetectYAML 测试 YAML 配置自动重新加载
func TestReloadConfig_AutoDetectYAML(t *testing.T) {
	tmpDir := t.TempDir()
	yamlPath := filepath.Join(tmpDir, "reload.yaml")

	// 初始配置（kebab-case）
	initialContent := `
app:
  name: "initial-app"
  version: "1.0.0"

server:
  host: "localhost"
  port: 8080

database:
  max-idle-conns: 10
  max-open-conns: 100
`
	err := os.WriteFile(yamlPath, []byte(initialContent), 0644)
	require.NoError(t, err)

	// 创建 Viper 和配置
	v := viper.New()
	v.SetConfigFile(yamlPath)
	err = v.ReadInConfig()
	require.NoError(t, err)

	type DBConfig struct {
		MaxIdleConns int `yaml:"max-idle-conns" json:"maxIdleConns"`
		MaxOpenConns int `yaml:"max-open-conns" json:"maxOpenConns"`
	}

	type ReloadTestConfig struct {
		App struct {
			Name    string `yaml:"name" json:"name"`
			Version string `yaml:"version" json:"version"`
		} `yaml:"app" json:"app"`
		Server struct {
			Host string `yaml:"host" json:"host"`
			Port int    `yaml:"port" json:"port"`
		} `yaml:"server" json:"server"`
		Database DBConfig `yaml:"database" json:"database"`
	}

	config := &ReloadTestConfig{}
	err = v.Unmarshal(config, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "yaml"
		dc.WeaklyTypedInput = true
	})
	require.NoError(t, err)

	// 创建热更新器
	reloader, err := NewHotReloader(config, v, yamlPath, &HotReloadConfig{
		Enabled: false,
	})
	require.NoError(t, err)

	// 验证初始值
	currentConfig := reloader.GetConfig().(*ReloadTestConfig)
	assert.Equal(t, "initial-app", currentConfig.App.Name)
	assert.Equal(t, "1.0.0", currentConfig.App.Version)
	assert.Equal(t, "localhost", currentConfig.Server.Host)
	assert.Equal(t, 8080, currentConfig.Server.Port)
	assert.Equal(t, 10, currentConfig.Database.MaxIdleConns)
	assert.Equal(t, 100, currentConfig.Database.MaxOpenConns)

	// 修改配置文件
	updatedContent := `
app:
  name: "updated-app"
  version: "2.0.0"

server:
  host: "0.0.0.0"
  port: 9090

database:
  max-idle-conns: 20
  max-open-conns: 200
`
	err = os.WriteFile(yamlPath, []byte(updatedContent), 0644)
	require.NoError(t, err)

	// 手动触发重新加载
	ctx := context.Background()
	err = reloader.Reload(ctx)
	require.NoError(t, err)

	// 验证配置已更新
	updatedConfig := reloader.GetConfig().(*ReloadTestConfig)
	assert.Equal(t, "updated-app", updatedConfig.App.Name, "App.Name 应该更新")
	assert.Equal(t, "2.0.0", updatedConfig.App.Version, "App.Version 应该更新")
	assert.Equal(t, "0.0.0.0", updatedConfig.Server.Host, "Server.Host 应该更新")
	assert.Equal(t, 9090, updatedConfig.Server.Port, "Server.Port 应该更新")
	assert.Equal(t, 20, updatedConfig.Database.MaxIdleConns, "MaxIdleConns 应该更新")
	assert.Equal(t, 200, updatedConfig.Database.MaxOpenConns, "MaxOpenConns 应该更新")

	t.Logf("✅ YAML 配置自动重新加载测试通过")
}

// TestReloadConfig_AutoDetectJSON 测试 JSON 配置自动重新加载
func TestReloadConfig_AutoDetectJSON(t *testing.T) {
	tmpDir := t.TempDir()
	jsonPath := filepath.Join(tmpDir, "reload.json")

	// 初始配置（camelCase）
	initialContent := `{
  "app": {
    "name": "initial-json-app",
    "version": "1.0.0"
  },
  "server": {
    "host": "localhost",
    "port": 8080
  },
  "database": {
    "maxIdleConns": 10,
    "maxOpenConns": 100
  }
}`
	err := os.WriteFile(jsonPath, []byte(initialContent), 0644)
	require.NoError(t, err)

	v := viper.New()
	v.SetConfigFile(jsonPath)
	err = v.ReadInConfig()
	require.NoError(t, err)

	type DBConfig struct {
		MaxIdleConns int `yaml:"max-idle-conns" json:"maxIdleConns"`
		MaxOpenConns int `yaml:"max-open-conns" json:"maxOpenConns"`
	}

	type ReloadTestConfig struct {
		App struct {
			Name    string `yaml:"name" json:"name"`
			Version string `yaml:"version" json:"version"`
		} `yaml:"app" json:"app"`
		Server struct {
			Host string `yaml:"host" json:"host"`
			Port int    `yaml:"port" json:"port"`
		} `yaml:"server" json:"server"`
		Database DBConfig `yaml:"database" json:"database"`
	}

	config := &ReloadTestConfig{}
	err = v.Unmarshal(config, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "json"
		dc.WeaklyTypedInput = true
	})
	require.NoError(t, err)

	reloader, err := NewHotReloader(config, v, jsonPath, &HotReloadConfig{
		Enabled: false,
	})
	require.NoError(t, err)

	// 验证初始值
	currentConfig := reloader.GetConfig().(*ReloadTestConfig)
	assert.Equal(t, "initial-json-app", currentConfig.App.Name)
	assert.Equal(t, 8080, currentConfig.Server.Port)
	assert.Equal(t, 10, currentConfig.Database.MaxIdleConns)

	// 修改配置
	updatedContent := `{
  "app": {
    "name": "updated-json-app",
    "version": "2.0.0"
  },
  "server": {
    "host": "0.0.0.0",
    "port": 9090
  },
  "database": {
    "maxIdleConns": 25,
    "maxOpenConns": 250
  }
}`
	err = os.WriteFile(jsonPath, []byte(updatedContent), 0644)
	require.NoError(t, err)

	// 重新加载
	ctx := context.Background()
	err = reloader.Reload(ctx)
	require.NoError(t, err)

	// 验证更新
	updatedConfig := reloader.GetConfig().(*ReloadTestConfig)
	assert.Equal(t, "updated-json-app", updatedConfig.App.Name)
	assert.Equal(t, "2.0.0", updatedConfig.App.Version)
	assert.Equal(t, "0.0.0.0", updatedConfig.Server.Host)
	assert.Equal(t, 9090, updatedConfig.Server.Port)
	assert.Equal(t, 25, updatedConfig.Database.MaxIdleConns)
	assert.Equal(t, 250, updatedConfig.Database.MaxOpenConns)

	t.Logf("✅ JSON 配置自动重新加载测试通过")
}
