/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 11:20:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 11:20:00
 * @FilePath: \go-config\integrated_manager.go
 * @Description: 集成配置管理器，整合所有配置热更新功能
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
	"time"

	"github.com/kamalyes/go-logger"
	"github.com/spf13/viper"
)

// ManagerBuilder 配置管理器构建器 - 支持链式调用和泛型
// 提供灵活的配置管理器构建方式，支持多种配置发现模式
type ManagerBuilder[T any] struct {
	config          *T                 // 配置结构体指针
	configPath      string             // 直接指定的配置文件路径
	searchPath      string             // 配置文件搜索路径
	environment     EnvironmentType    // 运行环境类型
	configPrefix    string             // 配置文件名前缀
	pattern         string             // 文件匹配模式
	hotReloadConfig *HotReloadConfig   // 热重载配置
	contextOptions  *ContextKeyOptions // 上下文选项
	autoDiscovery   bool               // 是否启用自动发现
	usePattern      bool               // 是否使用模式匹配
	useCustomPrefix bool               // 是否使用自定义前缀
}

// IntegratedConfigManager 集成配置管理器
// 统一管理配置文件、环境变量、热重载和上下文的核心组件
type IntegratedConfigManager struct {
	mu              sync.RWMutex     // 读写锁，保护并发访问
	environment     *Environment     // 环境管理器
	hotReloader     HotReloader      // 热重载管理器
	contextManager  *ContextManager  // 上下文管理器
	viper           *viper.Viper     // Viper配置解析器
	config          interface{}      // 当前配置对象
	configPath      string           // 配置文件路径
	hotReloadConfig *HotReloadConfig // 热重载配置
	running         bool             // 运行状态标识
}

// IntegratedConfigOptions 集成配置管理器选项
type IntegratedConfigOptions struct {
	ConfigPath      string             // 配置文件路径
	Environment     EnvironmentType    // 初始环境
	HotReloadConfig *HotReloadConfig   // 热更新配置
	ContextOptions  *ContextKeyOptions // 上下文选项
}

// DefaultIntegratedConfigOptions 默认集成配置管理器选项
func DefaultIntegratedConfigOptions() *IntegratedConfigOptions {
	return &IntegratedConfigOptions{
		ConfigPath:      "",
		Environment:     DefaultEnv,
		HotReloadConfig: DefaultHotReloadConfig(),
		ContextOptions:  &ContextKeyOptions{Value: DefaultEnv},
	}
}

// NewManager 创建新的配置管理器构建器 - 链式调用API入口
// 这是创建配置管理器的推荐方式，支持泛型和流畅的链式调用
// 使用示例:
//
//	type MyConfig struct {
//	    Database string `yaml:"database"`
//	    Port     int    `yaml:"port"`
//	}
//	var config MyConfig
//	manager, err := NewManager(&config).
//	    WithSearchPath("./configs").
//	    WithPrefix("app").
//	    WithEnvironment(EnvDevelopment).
//	    WithHotReload(nil).
//	    BuildAndStart()
func NewManager[T any](config *T) *ManagerBuilder[T] {
	return &ManagerBuilder[T]{
		config:      config,
		environment: GetEnvironment(),
	}
}

// WithConfigPath 设置配置文件路径
// 直接指定配置文件的完整路径，优先级最高
// path: 配置文件的绝对路径或相对路径
func (b *ManagerBuilder[T]) WithConfigPath(path string) *ManagerBuilder[T] {
	b.configPath = path
	return b
}

// WithSearchPath 设置配置文件搜索路径
// 启用自动发现模式，在指定目录中查找配置文件
// path: 搜索配置文件的目录路径
func (b *ManagerBuilder[T]) WithSearchPath(path string) *ManagerBuilder[T] {
	b.searchPath = path
	b.autoDiscovery = true
	return b
}

// WithEnvironment 设置运行环境
// 指定当前应用的运行环境，影响配置文件的选择和环境变量的读取
// env: 环境类型 (EnvDevelopment, EnvTest, EnvStaging, EnvProduction)
func (b *ManagerBuilder[T]) WithEnvironment(env EnvironmentType) *ManagerBuilder[T] {
	b.environment = env
	return b
}

// WithPrefix 设置配置文件名前缀
// 用于匹配特定前缀的配置文件，结合环境后缀使用
// 例如: prefix="app" 可匹配 "app-dev.yaml", "app-prod.json" 等
// prefix: 配置文件名的前缀字符串
func (b *ManagerBuilder[T]) WithPrefix(prefix string) *ManagerBuilder[T] {
	b.configPrefix = prefix
	b.useCustomPrefix = true
	return b
}

// WithPattern 设置文件匹配模式
// 使用glob模式匹配配置文件，支持通配符
// 例如: "*.yaml" 匹配所有yaml文件, "config-*.json" 匹配以config-开头的json文件
// pattern: glob匹配模式字符串
func (b *ManagerBuilder[T]) WithPattern(pattern string) *ManagerBuilder[T] {
	b.pattern = pattern
	b.usePattern = true
	return b
}

// WithHotReload 启用配置热重载功能
// 当配置文件发生变化时自动重新加载配置
// config: 热重载配置，传nil使用默认配置
func (b *ManagerBuilder[T]) WithHotReload(config *HotReloadConfig) *ManagerBuilder[T] {
	if config == nil {
		config = DefaultHotReloadConfig()
	}
	b.hotReloadConfig = config
	return b
}

// WithContext 设置上下文配置选项
// 配置上下文管理器的行为和键值设置
// options: 上下文键选项配置
func (b *ManagerBuilder[T]) WithContext(options *ContextKeyOptions) *ManagerBuilder[T] {
	b.contextOptions = options
	return b
}

// Build 构建配置管理器
// 根据设置的选项构建管理器实例，但不启动热重载等服务
// 返回: 配置管理器实例和可能的错误
func (b *ManagerBuilder[T]) Build() (*IntegratedConfigManager, error) {
	configPath, err := b.resolveConfigPath()
	if err != nil {
		return nil, fmt.Errorf("解析配置路径失败: %w", err)
	}

	return CreateIntegratedManager(b.config, configPath, b.environment)
}

// BuildAndStart 构建并启动配置管理器
// 这是推荐的使用方式，一步完成管理器的创建和启动
// ctx: 可选的上下文，用于控制启动过程的超时和取消
// 返回: 已启动的配置管理器实例和可能的错误
func (b *ManagerBuilder[T]) BuildAndStart(ctx ...context.Context) (*IntegratedConfigManager, error) {
	manager, err := b.Build()
	if err != nil {
		return nil, err
	}

	// 使用提供的上下文或创建默认上下文
	var startCtx context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		startCtx = ctx[0]
	} else {
		var cancel context.CancelFunc
		startCtx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
	}

	if err := manager.Start(startCtx); err != nil {
		return nil, fmt.Errorf("启动管理器失败: %w", err)
	}

	return manager, nil
}

// MustBuildAndStart 构建并启动配置管理器
// 功能同BuildAndStart，但失败时会panic，适用于启动阶段
// ctx: 可选的上下文
// 返回: 已启动的配置管理器实例
func (b *ManagerBuilder[T]) MustBuildAndStart(ctx ...context.Context) *IntegratedConfigManager {
	manager, err := b.BuildAndStart(ctx...)
	if err != nil {
		panic(fmt.Sprintf("构建并启动配置管理器失败: %v", err))
	}
	return manager
}

// resolveConfigPath 解析配置文件路径
// 根据设置的选项和优先级顺序解析最终的配置文件路径
// 返回: 解析出的配置文件路径和可能的错误
func (b *ManagerBuilder[T]) resolveConfigPath() (string, error) {
	discovery := GetGlobalConfigDiscovery()

	switch {
	case b.usePattern:
		// 使用模式匹配
		configFiles, err := discovery.FindConfigFileByPattern(b.searchPath, b.pattern, b.environment)
		if err != nil {
			return "", fmt.Errorf("按模式查找配置文件失败: %w", err)
		}
		if len(configFiles) == 0 {
			return "", fmt.Errorf("未找到匹配模式 '%s' 的配置文件", b.pattern)
		}
		logger.GetGlobalLogger().Info("🔍 模式匹配找到配置文件: %s", configFiles[0].Path)
		return configFiles[0].Path, nil

	case b.useCustomPrefix:
		// 使用自定义前缀发现
		return b.discoverWithPrefix()

	case b.autoDiscovery:
		// 自动发现
		return b.autoDiscover()

	case b.configPath != "":
		// 直接使用指定路径
		logger.GetGlobalLogger().Info("📁 使用指定配置文件: %s", b.configPath)
		return b.configPath, nil

	default:
		return "", fmt.Errorf("未指定配置路径或搜索选项")
	}
}

// discoverWithPrefix 使用自定义前缀发现配置文件
// 根据指定的前缀和环境类型查找匹配的配置文件
// 返回: 发现的配置文件路径和可能的错误
func (b *ManagerBuilder[T]) discoverWithPrefix() (string, error) {
	discovery := &ConfigDiscovery{
		SupportedExtensions: []string{".yaml", ".yml", ".json", ".toml", ".properties"},
		DefaultNames:        []string{b.configPrefix},
		EnvPrefixes: map[EnvironmentType][]string{
			EnvDevelopment: {"dev", "development", "local"},
			EnvTest:        {"test", "testing"},
			EnvStaging:     {"staging", "stage", "pre", "preprod"},
			EnvProduction:  {"prod", "production", "release"},
		},
	}

	configFiles, err := discovery.DiscoverConfigFiles(b.searchPath, b.environment)
	if err != nil {
		return "", fmt.Errorf("发现配置文件失败: %w", err)
	}

	for _, file := range configFiles {
		if file.Exists {
			logger.GetGlobalLogger().Info("🎯 前缀匹配找到配置文件: %s (前缀: %s)", file.Path, b.configPrefix)
			return file.Path, nil
		}
	}

	return "", fmt.Errorf("未找到前缀为 '%s' 的配置文件", b.configPrefix)
}

// autoDiscover 自动发现配置文件
// 在指定路径中自动查找适合的配置文件，优先选择匹配环境的文件
// 返回: 发现的配置文件路径和可能的错误
func (b *ManagerBuilder[T]) autoDiscover() (string, error) {
	discovery := GetGlobalConfigDiscovery()

	configFiles, err := discovery.DiscoverConfigFiles(b.searchPath, b.environment)
	if err != nil {
		return "", fmt.Errorf("自动发现配置文件失败: %w", err)
	}

	for _, file := range configFiles {
		if file.Exists {
			logger.GetGlobalLogger().Info("🔍 自动发现配置文件: %s", file.Path)
			return file.Path, nil
		}
	}

	return "", fmt.Errorf("在路径 '%s' 中未找到有效配置文件", b.searchPath)
}

// NewIntegratedConfigManager 创建集成配置管理器
// 使用指定的配置对象和选项创建一个完整的配置管理器实例
// 这是底层创建函数，一般推荐使用NewManager进行链式调用
// config: 配置结构体指针
// options: 集成配置管理器选项
// 返回: 配置管理器实例和可能的错误
func NewIntegratedConfigManager(config interface{}, options *IntegratedConfigOptions) (*IntegratedConfigManager, error) {
	if options == nil {
		options = DefaultIntegratedConfigOptions()
	}

	// 创建环境管理器
	env := NewEnvironment()
	if options.Environment != "" {
		env.SetEnvironment(options.Environment)
	}

	// 设置上下文键
	if options.ContextOptions != nil {
		SetContextKey(options.ContextOptions)
	}

	// 创建Viper实例
	v := viper.New()

	// 配置Viper
	if options.ConfigPath != "" {
		if info, err := os.Stat(options.ConfigPath); err == nil && !info.IsDir() {
			// 是文件，设置配置文件路径
			v.SetConfigFile(options.ConfigPath)
			ext := filepath.Ext(options.ConfigPath)
			if len(ext) > 1 {
				v.SetConfigType(ext[1:]) // 去掉点号
			}
		} else {
			// 是目录或不存在，让Viper在该目录中查找配置文件
			v.AddConfigPath(options.ConfigPath)
			v.SetConfigName("config") // 默认配置文件名
		}
	}

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置到结构体
	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	// 创建热更新器
	hotReloader, err := NewHotReloader(config, v, options.ConfigPath, options.HotReloadConfig)
	if err != nil {
		return nil, fmt.Errorf("创建热更新器失败: %w", err)
	}

	// 创建上下文管理器
	contextManager := NewContextManager(env, hotReloader)

	// 初始化全局上下文管理器
	InitializeContextManager(env, hotReloader)

	manager := &IntegratedConfigManager{
		environment:     env,
		hotReloader:     hotReloader,
		contextManager:  contextManager,
		viper:           v,
		config:          config,
		configPath:      options.ConfigPath,
		hotReloadConfig: options.HotReloadConfig,
		running:         false,
	}

	// 注册内部回调
	manager.registerInternalCallbacks()

	logger.GetGlobalLogger().Info("✅ 集成配置管理器创建完成，配置文件: %s", options.ConfigPath)
	return manager, nil
}

// registerInternalCallbacks 注册内部回调函数
// 设置管理器内部的事件监听和响应机制
func (icm *IntegratedConfigManager) registerInternalCallbacks() {
	// 注册配置变更回调
	icm.hotReloader.RegisterCallback(icm.onConfigReloaded, CallbackOptions{
		ID:       "integrated_manager_config",
		Types:    []CallbackType{CallbackTypeConfigChanged, CallbackTypeReloaded},
		Priority: -1000, // 高优先级
		Async:    false,
		Timeout:  5 * time.Second,
	})

	// 注册环境变更回调
	icm.environment.RegisterCallback("integrated_manager_env", icm.onEnvironmentChanged, -1000, false)

	// 注册错误回调
	icm.hotReloader.RegisterCallback(icm.onError, CallbackOptions{
		ID:       "integrated_manager_error",
		Types:    []CallbackType{CallbackTypeError},
		Priority: -1000,
		Async:    true,
		Timeout:  10 * time.Second,
	})
}

// onConfigReloaded 处理配置重新加载事件
// 当配置文件发生变化并成功重新加载时触发
func (icm *IntegratedConfigManager) onConfigReloaded(ctx context.Context, event CallbackEvent) error {
	icm.mu.Lock()
	defer icm.mu.Unlock()

	logger.GetGlobalLogger().Info("🔄 集成管理器: 配置已重新加载，来源: %s", event.Source)

	// 自动记录美化的配置变更日志
	if isAutoLogEnabled() {
		LogConfigChange(event, event.NewValue)
	}

	// 更新本地配置引用
	icm.config = event.NewValue

	// 更新上下文管理器中的配置
	icm.contextManager.UpdateConfig(event.NewValue)

	return nil
}

// onEnvironmentChanged 处理环境变更事件
// 当应用环境发生变化时触发，记录日志并执行相关操作
func (icm *IntegratedConfigManager) onEnvironmentChanged(oldEnv, newEnv EnvironmentType) error {
	logger.GetGlobalLogger().Info("🌍 集成管理器: 环境已变更: %s -> %s", oldEnv, newEnv)

	// 自动记录美化的环境变更日志
	if isAutoLogEnabled() {
		LogEnvChange(oldEnv, newEnv)
	}

	return nil
}

// onError 处理错误事件
// 当配置管理过程中发生错误时触发，统一记录和处理错误
func (icm *IntegratedConfigManager) onError(ctx context.Context, event CallbackEvent) error {
	logger.GetGlobalLogger().Error("❌ 集成管理器: 发生错误: %s, 来源: %s", event.Error, event.Source)

	// 自动记录美化的错误日志
	if isAutoLogEnabled() {
		LogConfigError(event)
	}

	return nil
}

// Start 启动配置管理器
// 启动热重载服务和相关的监控机制
// ctx: 用于控制启动过程的上下文
// 返回: 启动成功返回nil，否则返回错误
func (icm *IntegratedConfigManager) Start(ctx context.Context) error {
	icm.mu.Lock()
	defer icm.mu.Unlock()

	if icm.running {
		return fmt.Errorf("集成配置管理器已在运行")
	}

	// 启动热更新器
	if err := icm.hotReloader.Start(ctx); err != nil {
		return fmt.Errorf("启动热更新器失败: %w", err)
	}

	// 更新上下文管理器配置
	icm.contextManager.UpdateConfig(icm.config)

	icm.running = true
	logger.GetGlobalLogger().Info("🚀 集成配置管理器启动成功")
	return nil
}

// Stop 停止配置管理器
// 停止热重载服务和所有监控机制，释放相关资源
// 返回: 停止成功返回nil，否则返回错误
func (icm *IntegratedConfigManager) Stop() error {
	icm.mu.Lock()
	defer icm.mu.Unlock()

	if !icm.running {
		return fmt.Errorf("集成配置管理器未运行")
	}

	// 停止热更新器
	if err := icm.hotReloader.Stop(); err != nil {
		return fmt.Errorf("停止热更新器失败: %w", err)
	}

	// 停止环境监控
	icm.environment.StopWatch()

	icm.running = false
	logger.GetGlobalLogger().Info("⏹️ 集成配置管理器已停止")
	return nil
}

// IsRunning 检查管理器是否正在运行
// 返回: true表示正在运行，false表示已停止
func (icm *IntegratedConfigManager) IsRunning() bool {
	icm.mu.RLock()
	defer icm.mu.RUnlock()
	return icm.running
}

// GetConfig 获取当前配置
// 返回当前加载的配置对象，线程安全
// 返回: 配置对象接口
func (icm *IntegratedConfigManager) GetConfig() interface{} {
	icm.mu.RLock()
	defer icm.mu.RUnlock()
	return icm.config
}

// GetConfigAs 获取指定类型的配置
// 泛型函数，安全地将配置转换为指定类型
// T: 目标配置类型
// icm: 配置管理器实例
// 返回: 类型安全的配置指针和可能的类型转换错误
func GetConfigAs[T any](icm *IntegratedConfigManager) (*T, error) {
	config := icm.GetConfig()
	if typedConfig, ok := config.(*T); ok {
		return typedConfig, nil
	}
	return nil, fmt.Errorf("配置类型不匹配: 期望 %T, 实际 %T", new(T), config)
}

// GetEnvironment 获取当前环境
// 返回: 当前应用的运行环境类型
func (icm *IntegratedConfigManager) GetEnvironment() EnvironmentType {
	return icm.environment.Value
}

// GetViper 获取Viper实例
// 返回内部使用的Viper配置解析器实例以便高级操作
// 返回: Viper实例指针
func (icm *IntegratedConfigManager) GetViper() *viper.Viper {
	return icm.viper
}

// GetHotReloader 获取热重载器
// 返回内部使用的热重载器实例以便直接操作
// 返回: 热重载器接口
func (icm *IntegratedConfigManager) GetHotReloader() HotReloader {
	return icm.hotReloader
}

// GetContextManager 获取上下文管理器
// 返回内部使用的上下文管理器实例
// 返回: 上下文管理器指针
func (icm *IntegratedConfigManager) GetContextManager() *ContextManager {
	return icm.contextManager
}

// GetEnvironmentManager 获取环境管理器
// 返回内部使用的环境管理器实例
// 返回: 环境管理器指针
func (icm *IntegratedConfigManager) GetEnvironmentManager() *Environment {
	return icm.environment
}

// WithContext 将配置信息注入到上下文中
// 返回包含配置信息的新上下文，便于跨组件传递配置
// ctx: 原始上下文
// 返回: 包含配置的新上下文
func (icm *IntegratedConfigManager) WithContext(ctx context.Context) context.Context {
	return icm.contextManager.WithConfig(ctx)
}

// RegisterConfigCallback 注册配置变更回调
// 当配置发生变化时会触发指定的回调函数
// callback: 回调函数
// options: 回调选项，包括ID、优先级、异步等设置
// 返回: 注册成功返回nil，否则返回错误
func (icm *IntegratedConfigManager) RegisterConfigCallback(callback CallbackFunc, options CallbackOptions) error {
	return icm.hotReloader.RegisterCallback(callback, options)
}

// RegisterEnvironmentCallback 注册环境变更回调
// 当应用环境发生变化时会触发指定的回调函数
// id: 回调的唯一标识符
// callback: 环境变更回调函数
// priority: 回调优先级，数值越小优先级越高
// async: 是否异步执行
// 返回: 注册成功返回nil，否则返回错误
func (icm *IntegratedConfigManager) RegisterEnvironmentCallback(id string, callback EnvironmentCallback, priority int, async bool) error {
	return icm.environment.RegisterCallback(id, callback, priority, async)
}

// UnregisterConfigCallback 取消配置变更回调
// 根据ID移除指定的配置变更回调
// id: 回调的唯一标识符
// 返回: 取消成功返回nil，否则返回错误
func (icm *IntegratedConfigManager) UnregisterConfigCallback(id string) error {
	return icm.hotReloader.UnregisterCallback(id)
}

// UnregisterEnvironmentCallback 取消环境变更回调
// 根据ID移除指定的环境变更回调
// id: 回调的唯一标识符
// 返回: 取消成功返回nil，否则返回错误
func (icm *IntegratedConfigManager) UnregisterEnvironmentCallback(id string) error {
	return icm.environment.UnregisterCallback(id)
}

// ReloadConfig 手动重新加载配置
// 立即从配置文件重新读取配置，触发相关回调
// ctx: 用于控制重载过程的上下文
// 返回: 重载成功返回nil，否则返回错误
func (icm *IntegratedConfigManager) ReloadConfig(ctx context.Context) error {
	return icm.hotReloader.Reload(ctx)
}

// SetEnvironment 设置应用环境
// 更新当前应用的运行环境，会触发环境变更回调
// env: 新的环境类型
// 返回: 设置成功返回nil，否则返回错误
func (icm *IntegratedConfigManager) SetEnvironment(env EnvironmentType) error {
	icm.environment.SetEnvironment(env)
	return nil
}

// ValidateConfig 验证配置有效性
// 检查当前加载的配置是否有效
// 返回: 配置有效返回nil，否则返回错误
func (icm *IntegratedConfigManager) ValidateConfig() error {
	if icm.config == nil {
		return fmt.Errorf("配置为空")
	}

	logger.GetGlobalLogger().Info("✅ 配置验证通过")
	return nil
}

// GetConfigMetadata 获取配置元数据
// 返回配置管理器和配置文件的详细信息
// 返回: 包含元数据的字典
func (icm *IntegratedConfigManager) GetConfigMetadata() map[string]interface{} {
	metadata := make(map[string]interface{})

	metadata["config_path"] = icm.configPath
	metadata["environment"] = icm.GetEnvironment()
	metadata["running"] = icm.IsRunning()
	metadata["hot_reload_enabled"] = icm.hotReloadConfig.Enabled
	metadata["created_at"] = icm.contextManager.GetConfigContext().CreatedAt
	metadata["updated_at"] = icm.contextManager.GetConfigContext().UpdatedAt

	return metadata
}

// MustStart 必须成功启动配置管理器
// 功能同Start，但失败时会panic，适用于必须成功的场景
// ctx: 用于控制启动过程的上下文
func (icm *IntegratedConfigManager) MustStart(ctx context.Context) {
	if err := icm.Start(ctx); err != nil {
		panic(fmt.Sprintf("启动集成配置管理器失败: %v", err))
	}
}

// MustGetConfigAs 必须成功获取指定类型的配置
// 功能同GetConfigAs，但失败时会panic，适用于确定类型正确的场景
// T: 目标配置类型
// icm: 配置管理器实例
// 返回: 类型安全的配置指针
func MustGetConfigAs[T any](icm *IntegratedConfigManager) *T {
	config, err := GetConfigAs[T](icm)
	if err != nil {
		panic(fmt.Sprintf("获取配置失败: %v", err))
	}
	return config
}

// CreateIntegratedManager 创建集成配置管理器的便捷函数
// 使用默认选项快速创建配置管理器，适合简单场景
// config: 配置结构体指针
// configPath: 配置文件路径
// env: 运行环境
// 返回: 配置管理器实例和可能的错误
func CreateIntegratedManager(config interface{}, configPath string, env EnvironmentType) (*IntegratedConfigManager, error) {
	options := &IntegratedConfigOptions{
		ConfigPath:      configPath,
		Environment:     env,
		HotReloadConfig: DefaultHotReloadConfig(),
		ContextOptions: &ContextKeyOptions{
			Value: env,
		},
	}

	return NewIntegratedConfigManager(config, options)
}

// ScanAndDisplayConfigs 扫描并显示可用的配置文件
// 用于调试和排错，显示指定目录中的所有配置文件
// searchPath: 搜索目录路径
// env: 目标环境类型
// 返回: 配置文件信息列表和可能的错误
func ScanAndDisplayConfigs(searchPath string, env EnvironmentType) ([]*ConfigFileInfo, error) {
	discovery := GetGlobalConfigDiscovery()

	// 发现所有配置文件
	allConfigs, err := discovery.DiscoverConfigFiles(searchPath, env)
	if err != nil {
		return nil, err
	}

	// 扫描目录中实际存在的配置文件
	existingConfigs, err := discovery.ScanDirectory(searchPath)
	if err != nil {
		logger.GetGlobalLogger().Error("扫描目录失败: %v", err)
	}

	logger.GetGlobalLogger().Info("\n📋 配置文件发现报告:")
	logger.GetGlobalLogger().Info("🔍 搜索路径: %s", searchPath)
	logger.GetGlobalLogger().Info("🌍 目标环境: %s", env)

	if len(existingConfigs) > 0 {
		logger.GetGlobalLogger().Info("\n✅ 发现的现有配置文件:")
		for i, info := range existingConfigs {
			if i < 5 { // 只显示前5个
				logger.GetGlobalLogger().Info("   %d. %s (环境: %s, 优先级: %d)",
					i+1, info.Name, info.Environment, info.Priority)
			}
		}
		if len(existingConfigs) > 5 {
			logger.GetGlobalLogger().Info("   ... 还有 %d 个文件", len(existingConfigs)-5)
		}
	}

	// 显示推荐的配置文件
	logger.GetGlobalLogger().Info("\n💡 推荐的配置文件候选:")
	shown := 0
	for _, info := range allConfigs {
		if shown >= 3 {
			break
		}
		status := "❌ 不存在"
		if info.Exists {
			status = "✅ 存在"
		}
		logger.GetGlobalLogger().Info("   %d. %s (%s, 优先级: %d)",
			shown+1, info.Name, status, info.Priority)
		shown++
	}

	return allConfigs, nil
}
