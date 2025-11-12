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

// IntegratedConfigManager 集成配置管理器
type IntegratedConfigManager struct {
	mu              sync.RWMutex
	environment     *Environment
	hotReloader     HotReloader
	contextManager  *ContextManager
	viper           *viper.Viper
	config          interface{}
	configPath      string
	hotReloadConfig *HotReloadConfig
	running         bool
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
		ContextOptions:  nil,
	}
}

// NewIntegratedConfigManager 创建新的集成配置管理器
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
		// 检查是否是文件还是目录
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

	logger.GetGlobalLogger().Info("集成配置管理器创建完成，配置文件: %s", options.ConfigPath)
	return manager, nil
}

// registerInternalCallbacks 注册内部回调
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

// onConfigReloaded 配置重新加载回调
func (icm *IntegratedConfigManager) onConfigReloaded(ctx context.Context, event CallbackEvent) error {
	icm.mu.Lock()
	defer icm.mu.Unlock()

	logger.GetGlobalLogger().Info("集成管理器: 配置已重新加载，来源: %s", event.Source)

	// 🆕 自动记录美化的配置变更日志
	if isAutoLogEnabled() {
		LogConfigChange(event, event.NewValue)
	}

	// 更新本地配置引用
	icm.config = event.NewValue

	// 更新上下文管理器中的配置
	icm.contextManager.UpdateConfig(event.NewValue)

	return nil
}

// onEnvironmentChanged 环境变更回调
func (icm *IntegratedConfigManager) onEnvironmentChanged(oldEnv, newEnv EnvironmentType) error {
	logger.GetGlobalLogger().Info("集成管理器: 环境已变更: %s -> %s", oldEnv, newEnv)

	// 自动记录美化的环境变更日志
	if isAutoLogEnabled() {
		LogEnvChange(oldEnv, newEnv)
	}

	// 如果需要，可以在这里实现环境变更时的特殊逻辑
	// 例如：重新加载对应环境的配置文件

	return nil
}

// onError 错误回调
func (icm *IntegratedConfigManager) onError(ctx context.Context, event CallbackEvent) error {
	logger.GetGlobalLogger().Error("集成管理器: 发生错误: %s, 来源: %s", event.Error, event.Source)

	// 🆕 自动记录美化的错误日志
	if isAutoLogEnabled() {
		LogConfigError(event)
	}

	// 在这里可以实现错误处理逻辑，如发送告警等

	return nil
}

// Start 启动集成配置管理器
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
	logger.GetGlobalLogger().Info("集成配置管理器启动成功")
	return nil
}

// Stop 停止集成配置管理器
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
	logger.GetGlobalLogger().Info("集成配置管理器已停止")
	return nil
}

// IsRunning 检查是否正在运行
func (icm *IntegratedConfigManager) IsRunning() bool {
	icm.mu.RLock()
	defer icm.mu.RUnlock()
	return icm.running
}

// GetConfig 获取当前配置
func (icm *IntegratedConfigManager) GetConfig() interface{} {
	icm.mu.RLock()
	defer icm.mu.RUnlock()
	return icm.config
}

// GetConfigWithType 获取指定类型的配置
func GetConfigWithType[T any](icm *IntegratedConfigManager) (*T, error) {
	config := icm.GetConfig()
	if typedConfig, ok := config.(*T); ok {
		return typedConfig, nil
	}
	return nil, fmt.Errorf("配置类型不匹配")
}

// GetEnvironment 获取当前环境
func (icm *IntegratedConfigManager) GetEnvironment() EnvironmentType {
	return icm.environment.Value
}

// GetViper 获取Viper实例
func (icm *IntegratedConfigManager) GetViper() *viper.Viper {
	return icm.viper
}

// GetHotReloader 获取热更新器
func (icm *IntegratedConfigManager) GetHotReloader() HotReloader {
	return icm.hotReloader
}

// GetContextManager 获取上下文管理器
func (icm *IntegratedConfigManager) GetContextManager() *ContextManager {
	return icm.contextManager
}

// GetEnvironmentManager 获取环境管理器
func (icm *IntegratedConfigManager) GetEnvironmentManager() *Environment {
	return icm.environment
}

// WithContext 将配置信息添加到上下文中
func (icm *IntegratedConfigManager) WithContext(ctx context.Context) context.Context {
	return icm.contextManager.WithConfig(ctx)
}

// RegisterConfigCallback 注册配置变更回调
func (icm *IntegratedConfigManager) RegisterConfigCallback(callback CallbackFunc, options CallbackOptions) error {
	return icm.hotReloader.RegisterCallback(callback, options)
}

// RegisterEnvironmentCallback 注册环境变更回调
func (icm *IntegratedConfigManager) RegisterEnvironmentCallback(id string, callback EnvironmentCallback, priority int, async bool) error {
	return icm.environment.RegisterCallback(id, callback, priority, async)
}

// UnregisterConfigCallback 注销配置变更回调
func (icm *IntegratedConfigManager) UnregisterConfigCallback(id string) error {
	return icm.hotReloader.UnregisterCallback(id)
}

// UnregisterEnvironmentCallback 注销环境变更回调
func (icm *IntegratedConfigManager) UnregisterEnvironmentCallback(id string) error {
	return icm.environment.UnregisterCallback(id)
}

// ReloadConfig 手动重新加载配置
func (icm *IntegratedConfigManager) ReloadConfig(ctx context.Context) error {
	return icm.hotReloader.Reload(ctx)
}

// SetEnvironment 设置环境
func (icm *IntegratedConfigManager) SetEnvironment(env EnvironmentType) error {
	icm.environment.SetEnvironment(env)
	return nil
}

// ValidateConfig 验证配置
func (icm *IntegratedConfigManager) ValidateConfig() error {
	// 这里可以添加配置验证逻辑
	// 例如检查必需字段、验证格式等

	if icm.config == nil {
		return fmt.Errorf("配置为空")
	}

	logger.GetGlobalLogger().Info("配置验证通过")
	return nil
}

// GetConfigMetadata 获取配置元数据
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

// 便捷函数

// MustStart 启动集成配置管理器，失败时panic
func (icm *IntegratedConfigManager) MustStart(ctx context.Context) {
	if err := icm.Start(ctx); err != nil {
		panic(fmt.Sprintf("启动集成配置管理器失败: %v", err))
	}
}

// MustGetConfig 获取配置，失败时panic
func MustGetConfig[T any](icm *IntegratedConfigManager) *T {
	config, err := GetConfigWithType[T](icm)
	if err != nil {
		panic(fmt.Sprintf("获取配置失败: %v", err))
	}
	return config
}

// CreateIntegratedManager 创建集成配置管理器的便捷函数
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

// CreateAndStartIntegratedManager 创建并启动集成配置管理器的便捷函数
func CreateAndStartIntegratedManager(config interface{}, configPath string, env EnvironmentType) (*IntegratedConfigManager, error) {
	manager, err := CreateIntegratedManager(config, configPath, env)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := manager.Start(ctx); err != nil {
		return nil, fmt.Errorf("启动管理器失败: %w", err)
	}

	return manager, nil
}

// CreateIntegratedManagerWithAutoDiscovery 使用自动发现创建集成配置管理器
func CreateIntegratedManagerWithAutoDiscovery(config interface{}, searchPath string, env EnvironmentType, configType string) (*IntegratedConfigManager, error) {
	discovery := GetGlobalConfigDiscovery()

	// 尝试找到最佳配置文件
	configInfo, err := discovery.FindBestConfigFile(searchPath, env)
	if err != nil {
		// 如果没有找到，尝试创建默认配置文件
		logger.GetGlobalLogger().Warn("未找到配置文件，尝试创建默认配置: %v", err)
		configInfo, err = discovery.CreateDefaultConfigFile(searchPath, env, configType)
		if err != nil {
			return nil, fmt.Errorf("创建默认配置文件失败: %w", err)
		}
		logger.GetGlobalLogger().Info("已创建默认配置文件: %s", configInfo.Path)
	} else {
		logger.GetGlobalLogger().Info("找到配置文件: %s (环境: %s, 优先级: %d)", configInfo.Path, configInfo.Environment, configInfo.Priority)
	}

	return CreateIntegratedManager(config, configInfo.Path, env)
}

// CreateAndStartIntegratedManagerWithAutoDiscovery 使用自动发现创建并启动集成配置管理器
func CreateAndStartIntegratedManagerWithAutoDiscovery(config interface{}, searchPath string, env EnvironmentType, configType string) (*IntegratedConfigManager, error) {
	manager, err := CreateIntegratedManagerWithAutoDiscovery(config, searchPath, env, configType)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := manager.Start(ctx); err != nil {
		return nil, fmt.Errorf("启动管理器失败: %w", err)
	}

	return manager, nil
}

// CreateIntegratedManagerWithPattern 使用模式匹配创建集成配置管理器
func CreateIntegratedManagerWithPattern(config interface{}, searchPath, pattern string, env EnvironmentType) (*IntegratedConfigManager, error) {
	discovery := GetGlobalConfigDiscovery()

	// 按模式查找配置文件
	configFiles, err := discovery.FindConfigFileByPattern(searchPath, pattern, env)
	if err != nil {
		return nil, fmt.Errorf("按模式查找配置文件失败: %w", err)
	}

	if len(configFiles) == 0 {
		return nil, fmt.Errorf("未找到匹配模式 '%s' 的配置文件", pattern)
	}

	// 使用第一个匹配的文件（已按优先级排序）
	configInfo := configFiles[0]
	logger.GetGlobalLogger().Info("找到匹配的配置文件: %s (模式: %s)", configInfo.Path, pattern)

	return CreateIntegratedManager(config, configInfo.Path, env)
}

// ScanAndDisplayConfigs 扫描并显示可用的配置文件
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
