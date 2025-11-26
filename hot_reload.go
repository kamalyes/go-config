/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 15:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 17:08:23
 * @FilePath: \go-config\hot_reload.go
 * @Description: 重构后的配置热更新模块，使用通用回调管理器
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package goconfig

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kamalyes/go-logger"
	"github.com/spf13/viper"
	"reflect"
	"sync"
	"time"
)

// HotReloadConfig 热更新配置
type HotReloadConfig struct {
	Enabled         bool          `yaml:"enabled" json:"enabled"`                   // 是否启用热更新
	WatchInterval   time.Duration `yaml:"watch_interval" json:"watch_interval"`     // 监控间隔
	DebounceDelay   time.Duration `yaml:"debounce_delay" json:"debounce_delay"`     // 防抖延迟
	MaxRetries      int           `yaml:"max_retries" json:"max_retries"`           // 最大重试次数
	CallbackTimeout time.Duration `yaml:"callback_timeout" json:"callback_timeout"` // 回调超时
	EnableEnvWatch  bool          `yaml:"enable_env_watch" json:"enable_env_watch"` // 是否监控环境变量
}

// DefaultHotReloadConfig 默认热更新配置
func DefaultHotReloadConfig() *HotReloadConfig {
	return &HotReloadConfig{
		Enabled:         true,
		WatchInterval:   500 * time.Millisecond,
		DebounceDelay:   1 * time.Second,
		MaxRetries:      3,
		CallbackTimeout: 30 * time.Second,
		EnableEnvWatch:  true,
	}
}

// HotReloader 热更新器接口
type HotReloader interface {
	CallbackManager
	// Start 启动热更新监控
	Start(ctx context.Context) error
	// Stop 停止热更新监控
	Stop() error
	// Reload 手动重新加载配置
	Reload(ctx context.Context) error
	// IsRunning 检查是否正在运行
	IsRunning() bool
	// GetConfig 获取当前配置
	GetConfig() interface{}
	// SetConfig 设置配置
	SetConfig(config interface{}) error
}

// hotReloadManager 热更新管理器实现
type hotReloadManager struct {
	mu              sync.RWMutex       // 读写锁
	config          interface{}        // 当前配置对象
	viper           *viper.Viper       // Viper实例
	callbackManager CallbackManager    // 回调管理器
	hotConfig       *HotReloadConfig   // 热更新配置
	watcher         *fsnotify.Watcher  // 文件监控器
	environment     *Environment       // 环境实例
	running         bool               // 是否运行中
	cancel          context.CancelFunc // 取消函数
	configPath      string             // 配置文件路径
	debounceTimer   *time.Timer        // 防抖定时器
}

// NewHotReloader 创建新的热更新器
func NewHotReloader(config interface{}, viper *viper.Viper, configPath string, options *HotReloadConfig) (HotReloader, error) {
	if options == nil {
		options = DefaultHotReloadConfig()
	}

	// 创建文件监控器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("创建文件监控器失败: %w", err)
	}

	manager := &hotReloadManager{
		config:          config,
		viper:           viper,
		callbackManager: NewCallbackManager(),
		hotConfig:       options,
		watcher:         watcher,
		environment:     NewEnvironment(),
		configPath:      configPath,
	}

	return manager, nil
}

// 实现 CallbackManager 接口的代理方法
func (h *hotReloadManager) RegisterCallback(callback CallbackFunc, options CallbackOptions) error {
	// 设置默认值
	if options.Timeout == 0 {
		options.Timeout = h.hotConfig.CallbackTimeout
	}
	if len(options.Types) == 0 {
		options.Types = []CallbackType{CallbackTypeConfigChanged, CallbackTypeFileChanged}
	}

	return h.callbackManager.RegisterCallback(callback, options)
}

func (h *hotReloadManager) UnregisterCallback(id string) error {
	return h.callbackManager.UnregisterCallback(id)
}

func (h *hotReloadManager) TriggerCallbacks(ctx context.Context, event CallbackEvent) error {
	return h.callbackManager.TriggerCallbacks(ctx, event)
}

func (h *hotReloadManager) ListCallbacks() []string {
	return h.callbackManager.ListCallbacks()
}

func (h *hotReloadManager) ClearCallbacks() {
	h.callbackManager.ClearCallbacks()
}

func (h *hotReloadManager) HasCallback(id string) bool {
	return h.callbackManager.HasCallback(id)
}

// Start 启动热更新监控
func (h *hotReloadManager) Start(ctx context.Context) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.running {
		return fmt.Errorf("热更新器已经在运行")
	}

	if !h.hotConfig.Enabled {
		logger.GetGlobalLogger().Info("热更新功能已禁用")
		return nil
	}

	// 创建上下文
	runCtx, cancel := context.WithCancel(ctx)
	h.cancel = cancel

	// 添加配置文件监控
	if h.configPath != "" {
		err := h.watcher.Add(h.configPath)
		if err != nil {
			return fmt.Errorf("添加配置文件监控失败: %w", err)
		}
		logger.GetGlobalLogger().Info("开始监控配置文件: %s", h.configPath)
	}

	h.running = true

	// 启动监控协程
	go h.watchLoop(runCtx)

	// 启动环境变量监控
	if h.hotConfig.EnableEnvWatch {
		go h.watchEnvironment(runCtx)
	}

	logger.GetGlobalLogger().Info("🚀 热更新器启动成功")

	// 触发启动事件
	startEvent := CreateEvent(CallbackTypeStarted, "hot_reloader", nil, h.config)
	go h.TriggerCallbacks(ctx, startEvent)

	return nil
}

// Stop 停止热更新监控
func (h *hotReloadManager) Stop() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if !h.running {
		return fmt.Errorf("热更新器未运行")
	}

	// 触发停止事件
	stopEvent := CreateEvent(CallbackTypeStopped, "hot_reloader", h.config, nil)
	ctx, cancel := context.WithTimeout(context.Background(), h.hotConfig.CallbackTimeout)
	go func() {
		defer cancel()
		h.TriggerCallbacks(ctx, stopEvent)
	}()

	if h.cancel != nil {
		h.cancel()
	}

	if h.watcher != nil {
		h.watcher.Close()
	}

	if h.debounceTimer != nil {
		h.debounceTimer.Stop()
	}

	h.running = false
	logger.GetGlobalLogger().Info("⏹️ 热更新器已停止")
	return nil
}

// IsRunning 检查是否正在运行
func (h *hotReloadManager) IsRunning() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.running
}

// GetConfig 获取当前配置
func (h *hotReloadManager) GetConfig() interface{} {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.config
}

// SetConfig 设置配置
func (h *hotReloadManager) SetConfig(config interface{}) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	oldConfig := h.config
	h.config = config

	// 触发配置变更回调
	event := CreateEvent(CallbackTypeReloaded, "manual", oldConfig, config)
	event.WithMetadata("manual", true)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), h.hotConfig.CallbackTimeout)
		defer cancel()
		h.TriggerCallbacks(ctx, event)
	}()

	logger.GetGlobalLogger().Info("📝 配置已手动更新")
	return nil
}

// Reload 手动重新加载配置
func (h *hotReloadManager) Reload(ctx context.Context) error {
	logger.GetGlobalLogger().Info("🔄 手动重新加载配置")
	return h.reloadConfig(ctx, "manual_reload")
}

// watchLoop 监控文件变化
func (h *hotReloadManager) watchLoop(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.GetGlobalLogger().Error("文件监控协程发生panic: %v", r)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-h.watcher.Events:
			if !ok {
				return
			}
			h.handleFileEvent(ctx, event)
		case err, ok := <-h.watcher.Errors:
			if !ok {
				return
			}
			logger.GetGlobalLogger().Error("文件监控错误: %v", err)
			h.triggerErrorCallback(ctx, err, "file_watcher")
		}
	}
}

// handleFileEvent 处理文件事件
func (h *hotReloadManager) handleFileEvent(ctx context.Context, event fsnotify.Event) {
	if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
		logger.GetGlobalLogger().Info("📁 配置文件发生变化: %s", event.Name)

		// 使用锁保护防抖处理
		h.mu.Lock()
		if h.debounceTimer != nil {
			h.debounceTimer.Stop()
		}

		h.debounceTimer = time.AfterFunc(h.hotConfig.DebounceDelay, func() {
			if err := h.reloadConfig(ctx, event.Name); err != nil {
				logger.GetGlobalLogger().Error("重新加载配置失败: %v", err)
			}
		})
		h.mu.Unlock()
	}
}

// reloadConfig 重新加载配置
func (h *hotReloadManager) reloadConfig(ctx context.Context, source string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	oldConfig := h.config
	start := time.Now()

	// 重新读取配置文件
	if err := h.viper.ReadInConfig(); err != nil {
		logger.GetGlobalLogger().Error("重新读取配置文件失败: %v", err)
		h.triggerErrorCallback(ctx, err, source)
		return err
	}

	// 解析到配置结构
	newConfig := reflect.New(reflect.TypeOf(h.config).Elem()).Interface()

	// 使用灵活的命名匹配策略进行反序列化
	// 支持 kebab-case、camelCase、PascalCase 到 snake_case 的自动转换
	if err := UnmarshalWithFlexibleNaming(h.viper, newConfig); err != nil {
		logger.GetGlobalLogger().Error("解析配置文件失败: %v", err)
		h.triggerErrorCallback(ctx, err, source)
		return err
	}

	h.config = newConfig
	duration := time.Since(start)

	// 触发配置变更回调
	event := CreateEvent(CallbackTypeConfigChanged, source, oldConfig, newConfig)
	event.WithMetadata("config_path", h.configPath)
	event.WithMetadata("duration", duration)

	go func() {
		if err := h.TriggerCallbacks(ctx, event); err != nil {
			logger.GetGlobalLogger().Error("触发配置变更回调失败: %v", err)
		}
	}()

	logger.GetGlobalLogger().Info("✅ 配置重新加载完成 (耗时: %v)", duration)
	return nil
}

// watchEnvironment 监控环境变量变化
func (h *hotReloadManager) watchEnvironment(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.GetGlobalLogger().Error("环境监控协程发生panic: %v", r)
		}
	}()

	ticker := time.NewTicker(h.hotConfig.WatchInterval)
	defer ticker.Stop()

	lastEnv := GetEnvironment()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			currentEnv := GetEnvironment()
			if currentEnv != lastEnv {
				logger.GetGlobalLogger().Info("🌍 环境变量发生变化: %s -> %s", lastEnv, currentEnv)

				event := CreateEvent(CallbackTypeEnvVarChanged, string(GetContextKey()), lastEnv, currentEnv)
				event.WithMetadata("context_key", GetContextKey())

				go func() {
					if err := h.TriggerCallbacks(ctx, event); err != nil {
						logger.GetGlobalLogger().Error("触发环境变更回调失败: %v", err)
					}
				}()

				lastEnv = currentEnv
			}
		}
	}
}

// triggerErrorCallback 触发错误回调
func (h *hotReloadManager) triggerErrorCallback(ctx context.Context, err error, source string) {
	event := CreateErrorEvent(source, err)
	event.WithMetadata("error_source", source)

	go func() {
		if triggerErr := h.TriggerCallbacks(ctx, event); triggerErr != nil {
			logger.GetGlobalLogger().Error("触发错误回调失败: %v", triggerErr)
		}
	}()
}
