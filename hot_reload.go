/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 11:10:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 12:00:00
 * @FilePath: \go-config\hot_reload.go
 * @Description: 配置热更新和回调监听功能模块 - 重构优化版本
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// CallbackType 定义回调类型
type CallbackType string

const (
	CallbackTypeConfigChanged CallbackType = "config_changed" // 配置文件变更
	CallbackTypeEnvChanged    CallbackType = "env_changed"    // 环境变量变更
	CallbackTypeReloaded      CallbackType = "reloaded"       // 重新加载完成
	CallbackTypeError         CallbackType = "error"          // 错误回调
	CallbackTypeValidation    CallbackType = "validation"     // 配置验证回调
	CallbackTypeStartup       CallbackType = "startup"        // 启动回调
	CallbackTypeShutdown      CallbackType = "shutdown"       // 关闭回调
)

// CallbackPriority 定义回调优先级常量
const (
	CallbackPriorityHighest = -1000 // 最高优先级
	CallbackPriorityHigh    = -100  // 高优先级
	CallbackPriorityNormal  = 0     // 普通优先级
	CallbackPriorityLow     = 100   // 低优先级
	CallbackPriorityLowest  = 1000  // 最低优先级
)

// CallbackEvent 回调事件
type CallbackEvent struct {
	Type        CallbackType           `json:"type"`        // 事件类型
	Timestamp   time.Time              `json:"timestamp"`   // 事件时间
	Source      string                 `json:"source"`      // 事件源（文件路径、环境变量名等）
	OldValue    interface{}            `json:"old_value"`   // 旧值
	NewValue    interface{}            `json:"new_value"`   // 新值
	Environment EnvironmentType        `json:"environment"` // 当前环境
	Error       error                  `json:"error"`       // 错误信息（仅错误回调）
	Metadata    map[string]interface{} `json:"metadata"`    // 附加元数据
	ConfigPath  string                 `json:"config_path"` // 配置文件路径
	Duration    time.Duration          `json:"duration"`    // 事件处理耗时
}

// CallbackFunc 回调函数类型
type CallbackFunc func(ctx context.Context, event CallbackEvent) error

// CallbackOptions 回调选项
type CallbackOptions struct {
	ID       string                 `json:"id"`       // 回调唯一标识
	Types    []CallbackType         `json:"types"`    // 监听的事件类型
	Priority int                    `json:"priority"` // 优先级（数字越小优先级越高）
	Async    bool                   `json:"async"`    // 是否异步执行
	Timeout  time.Duration          `json:"timeout"`  // 超时时间
	Retry    int                    `json:"retry"`    // 重试次数
	Metadata map[string]interface{} `json:"metadata"` // 附加元数据
}

// HotReloadConfig 热更新配置
type HotReloadConfig struct {
	Enabled         bool          `yaml:"enabled" json:"enabled"`                   // 是否启用热更新
	WatchInterval   time.Duration `yaml:"watch-interval" json:"watch_interval"`     // 监控间隔
	DebounceDelay   time.Duration `yaml:"debounce-delay" json:"debounce_delay"`     // 防抖延迟
	MaxRetries      int           `yaml:"max-retries" json:"max_retries"`           // 最大重试次数
	CallbackTimeout time.Duration `yaml:"callback-timeout" json:"callback_timeout"` // 回调超时
	EnableEnvWatch  bool          `yaml:"enable-env-watch" json:"enable_env_watch"` // 是否监控环境变量
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

// CallbackManager 回调管理器接口
type CallbackManager interface {
	// RegisterCallback 注册回调
	RegisterCallback(callback CallbackFunc, options CallbackOptions) error
	// UnregisterCallback 注销回调
	UnregisterCallback(id string) error
	// TriggerCallbacks 触发回调
	TriggerCallbacks(ctx context.Context, event CallbackEvent) error
	// ListCallbacks 列出所有回调
	ListCallbacks() []CallbackOptions
	// ClearCallbacks 清除所有回调
	ClearCallbacks()
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

// callbackInfo 回调信息
type callbackInfo struct {
	Func    CallbackFunc    `json:"func"`
	Options CallbackOptions `json:"options"`
}

// hotReloadManager 热更新管理器实现
type hotReloadManager struct {
	mu            sync.RWMutex             // 读写锁
	config        interface{}              // 当前配置对象
	viper         *viper.Viper             // Viper实例
	callbacks     map[string]*callbackInfo // 回调映射
	hotConfig     *HotReloadConfig         // 热更新配置
	watcher       *fsnotify.Watcher        // 文件监控器
	envManager    *EnvironmentManager      // 环境管理器
	environment   *Environment             // 环境实例
	running       bool                     // 是否运行中
	cancel        context.CancelFunc       // 取消函数
	configPath    string                   // 配置文件路径
	debounceTimer *time.Timer              // 防抖定时器
	lastModTime   time.Time                // 最后修改时间
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
		config:      config,
		viper:       viper,
		callbacks:   make(map[string]*callbackInfo),
		hotConfig:   options,
		watcher:     watcher,
		envManager:  defaultEnvManager,
		environment: NewEnvironment(),
		configPath:  configPath,
	}

	return manager, nil
}

// RegisterCallback 注册回调
func (h *hotReloadManager) RegisterCallback(callback CallbackFunc, options CallbackOptions) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if options.ID == "" {
		return fmt.Errorf("回调ID不能为空")
	}

	if _, exists := h.callbacks[options.ID]; exists {
		return fmt.Errorf("回调ID %s 已存在", options.ID)
	}

	// 设置默认值
	if options.Timeout == 0 {
		options.Timeout = h.hotConfig.CallbackTimeout
	}
	if len(options.Types) == 0 {
		options.Types = []CallbackType{CallbackTypeConfigChanged, CallbackTypeEnvChanged}
	}

	h.callbacks[options.ID] = &callbackInfo{
		Func:    callback,
		Options: options,
	}

	log.Printf("注册回调 %s, 类型: %v, 优先级: %d", options.ID, options.Types, options.Priority)
	return nil
}

// UnregisterCallback 注销回调
func (h *hotReloadManager) UnregisterCallback(id string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.callbacks[id]; !exists {
		return fmt.Errorf("回调ID %s 不存在", id)
	}

	delete(h.callbacks, id)
	log.Printf("注销回调 %s", id)
	return nil
}

// TriggerCallbacks 触发回调
func (h *hotReloadManager) TriggerCallbacks(ctx context.Context, event CallbackEvent) error {
	h.mu.RLock()
	callbacks := make([]*callbackInfo, 0)
	for _, cb := range h.callbacks {
		// 检查是否监听此事件类型
		for _, eventType := range cb.Options.Types {
			if eventType == event.Type {
				callbacks = append(callbacks, cb)
				break
			}
		}
	}
	h.mu.RUnlock()

	// 按优先级排序
	for i := 0; i < len(callbacks)-1; i++ {
		for j := i + 1; j < len(callbacks); j++ {
			if callbacks[i].Options.Priority > callbacks[j].Options.Priority {
				callbacks[i], callbacks[j] = callbacks[j], callbacks[i]
			}
		}
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(callbacks))

	// 执行回调
	for _, cb := range callbacks {
		callback := cb
		if callback.Options.Async {
			wg.Add(1)
			go func() {
				defer wg.Done()
				h.executeCallback(ctx, callback, event, errChan)
			}()
		} else {
			h.executeCallback(ctx, callback, event, errChan)
		}
	}

	if len(callbacks) > 0 {
		// 等待异步回调完成
		go func() {
			wg.Wait()
			close(errChan)
		}()

		// 收集错误
		var errors []error
		for err := range errChan {
			if err != nil {
				errors = append(errors, err)
			}
		}

		if len(errors) > 0 {
			return fmt.Errorf("回调执行错误: %v", errors)
		}
	}

	return nil
}

// executeCallback 执行单个回调
func (h *hotReloadManager) executeCallback(ctx context.Context, callback *callbackInfo, event CallbackEvent, errChan chan<- error) {
	defer func() {
		if r := recover(); r != nil {
			errChan <- fmt.Errorf("回调 %s panic: %v", callback.Options.ID, r)
		}
	}()

	// 设置超时上下文
	timeoutCtx, cancel := context.WithTimeout(ctx, callback.Options.Timeout)
	defer cancel()

	var lastErr error
	maxRetries := callback.Options.Retry
	if maxRetries <= 0 {
		maxRetries = 1
	}

	// 重试机制
	for i := 0; i < maxRetries; i++ {
		err := callback.Func(timeoutCtx, event)
		if err == nil {
			errChan <- nil
			return
		}
		lastErr = err
		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * time.Second) // 指数退避
		}
	}

	errChan <- fmt.Errorf("回调 %s 执行失败 (重试%d次): %w", callback.Options.ID, maxRetries, lastErr)
}

// ListCallbacks 列出所有回调
func (h *hotReloadManager) ListCallbacks() []CallbackOptions {
	h.mu.RLock()
	defer h.mu.RUnlock()

	options := make([]CallbackOptions, 0, len(h.callbacks))
	for _, cb := range h.callbacks {
		options = append(options, cb.Options)
	}
	return options
}

// ClearCallbacks 清除所有回调
func (h *hotReloadManager) ClearCallbacks() {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.callbacks = make(map[string]*callbackInfo)
	log.Printf("已清除所有回调")
}

// Start 启动热更新监控
func (h *hotReloadManager) Start(ctx context.Context) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.running {
		return fmt.Errorf("热更新器已经在运行")
	}

	if !h.hotConfig.Enabled {
		log.Printf("热更新功能已禁用")
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
		log.Printf("开始监控配置文件: %s", h.configPath)
	}

	h.running = true

	// 启动监控协程
	go h.watchLoop(runCtx)

	// 启动环境变量监控
	if h.hotConfig.EnableEnvWatch {
		go h.watchEnvironment(runCtx)
	}

	log.Printf("热更新器启动成功")
	return nil
}

// Stop 停止热更新监控
func (h *hotReloadManager) Stop() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if !h.running {
		return fmt.Errorf("热更新器未运行")
	}

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
	log.Printf("热更新器已停止")
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
	event := CallbackEvent{
		Type:        CallbackTypeReloaded,
		Timestamp:   time.Now(),
		Source:      "manual",
		OldValue:    oldConfig,
		NewValue:    config,
		Environment: GetEnvironment(),
		Metadata:    map[string]interface{}{"manual": true},
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), h.hotConfig.CallbackTimeout)
		defer cancel()
		h.TriggerCallbacks(ctx, event)
	}()

	return nil
}

// Reload 手动重新加载配置
func (h *hotReloadManager) Reload(ctx context.Context) error {
	log.Printf("手动重新加载配置")
	return h.reloadConfig(ctx, "manual_reload")
}

// watchLoop 监控文件变化
func (h *hotReloadManager) watchLoop(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("文件监控协程发生panic: %v", r)
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
			log.Printf("文件监控错误: %v", err)
			h.triggerErrorCallback(ctx, err, "file_watcher")
		}
	}
}

// handleFileEvent 处理文件事件
func (h *hotReloadManager) handleFileEvent(ctx context.Context, event fsnotify.Event) {
	if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
		log.Printf("配置文件发生变化: %s", event.Name)

		// 防抖处理
		if h.debounceTimer != nil {
			h.debounceTimer.Stop()
		}

		h.debounceTimer = time.AfterFunc(h.hotConfig.DebounceDelay, func() {
			h.reloadConfig(ctx, event.Name)
		})
	}
}

// reloadConfig 重新加载配置
func (h *hotReloadManager) reloadConfig(ctx context.Context, source string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	oldConfig := h.config

	// 重新读取配置文件
	if err := h.viper.ReadInConfig(); err != nil {
		log.Printf("重新读取配置文件失败: %v", err)
		h.triggerErrorCallback(ctx, err, source)
		return err
	}

	// 解析到配置结构
	newConfig := reflect.New(reflect.TypeOf(h.config).Elem()).Interface()
	if err := h.viper.Unmarshal(newConfig); err != nil {
		log.Printf("解析配置文件失败: %v", err)
		h.triggerErrorCallback(ctx, err, source)
		return err
	}

	h.config = newConfig

	// 触发配置变更回调
	event := CallbackEvent{
		Type:        CallbackTypeConfigChanged,
		Timestamp:   time.Now(),
		Source:      source,
		OldValue:    oldConfig,
		NewValue:    newConfig,
		Environment: GetEnvironment(),
		Metadata:    map[string]interface{}{"config_path": h.configPath},
	}

	go func() {
		if err := h.TriggerCallbacks(ctx, event); err != nil {
			log.Printf("触发配置变更回调失败: %v", err)
		}
	}()

	log.Printf("配置重新加载完成")
	return nil
}

// watchEnvironment 监控环境变量变化
func (h *hotReloadManager) watchEnvironment(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("环境监控协程发生panic: %v", r)
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
				log.Printf("环境变量发生变化: %s -> %s", lastEnv, currentEnv)

				event := CallbackEvent{
					Type:        CallbackTypeEnvChanged,
					Timestamp:   time.Now(),
					Source:      string(GetContextKey()),
					OldValue:    lastEnv,
					NewValue:    currentEnv,
					Environment: currentEnv,
					Metadata:    map[string]interface{}{"context_key": GetContextKey()},
				}

				go func() {
					if err := h.TriggerCallbacks(ctx, event); err != nil {
						log.Printf("触发环境变更回调失败: %v", err)
					}
				}()

				lastEnv = currentEnv
			}
		}
	}
}

// triggerErrorCallback 触发错误回调
func (h *hotReloadManager) triggerErrorCallback(ctx context.Context, err error, source string) {
	event := CallbackEvent{
		Type:        CallbackTypeError,
		Timestamp:   time.Now(),
		Source:      source,
		Error:       err,
		Environment: GetEnvironment(),
		Metadata:    map[string]interface{}{"error_source": source},
	}

	go func() {
		if triggerErr := h.TriggerCallbacks(ctx, event); triggerErr != nil {
			log.Printf("触发错误回调失败: %v", triggerErr)
		}
	}()
}
