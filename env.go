/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-30 08:22:08
 * @FilePath: \go-config\env.go
 * @Description: 重构后的环境管理器，提供线程安全的环境管理和别名支持
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

// EnvironmentType 定义环境类型
type EnvironmentType string

// 定义有效的环境常量
const (
	EnvDevelopment EnvironmentType = "development" // 开发环境
	EnvTest        EnvironmentType = "test"        // 测试环境
	EnvStaging     EnvironmentType = "staging"     // 预发布环境
	EnvUAT         EnvironmentType = "uat"         // 用户验收测试环境
	EnvProduction  EnvironmentType = "production"  // 生产环境
	EnvLocal       EnvironmentType = "local"       // 本地环境
	EnvDebug       EnvironmentType = "debug"       // 调试环境
	EnvDemo        EnvironmentType = "demo"        // 演示环境
	EnvIntegration EnvironmentType = "integration" // 集成环境

	// 向后兼容的别名
	Dev  = EnvDevelopment
	Sit  = EnvTest
	Fat  = EnvStaging
	Uat  = EnvUAT
	Prod = EnvProduction

	DefaultEnv EnvironmentType = EnvDevelopment
)

// DefaultEnvPrefixes 全局环境前缀映射（统一定义，避免重复）
// 覆盖常见的环境命名约定，支持多种风格：短名、全名、连字符等
var DefaultEnvPrefixes = map[EnvironmentType][]string{
	EnvDevelopment: {"dev", "develop", "development"},
	EnvLocal:       {"local", "localhost"},
	EnvTest:        {"test", "testing", "qa", "sit"},
	EnvStaging:     {"staging", "stage", "stg", "pre", "preprod", "pre-prod", "fat", "gray", "grey", "canary"},
	EnvProduction:  {"prod", "production", "prd", "release", "live", "online", "master", "main"},
	EnvDebug:       {"debug", "debugging", "dbg"},
	EnvDemo:        {"demo", "demonstration", "showcase", "preview", "sandbox"},
	EnvUAT:         {"uat", "acceptance", "user-acceptance", "beta"},
	EnvIntegration: {"integration", "int", "ci", "integration-test", "integ"},
}

// DefaultSupportedExtensions 默认支持的配置文件扩展名
var DefaultSupportedExtensions = []string{".yaml", ".yml", ".json", ".toml", ".properties"}

// DefaultConfigNames 默认配置文件名
var DefaultConfigNames = []string{"config", "application", "app", "gateway", "service"}

// envPrefixesMutex 保护 DefaultEnvPrefixes 的并发访问
var envPrefixesMutex sync.RWMutex

// RegisterEnvPrefixes 注册自定义环境前缀（包级别便捷函数）
// 内部调用 GetGlobalEnvManager().RegisterEnvironment()，保持逻辑统一
// 示例:
//
//	func init() {
//	    goconfig.RegisterEnvPrefixes("custom", "custom", "my-env", "myenv")
//	}
func RegisterEnvPrefixes(env EnvironmentType, prefixes ...string) {
	GetGlobalEnvManager().RegisterEnvironment(env, prefixes...)
}

// GetEnvPrefixes 获取指定环境的前缀列表
func GetEnvPrefixes(env EnvironmentType) []string {
	envPrefixesMutex.RLock()
	defer envPrefixesMutex.RUnlock()

	if prefixes, ok := DefaultEnvPrefixes[env]; ok {
		// 返回副本，避免外部修改
		result := make([]string, len(prefixes))
		copy(result, prefixes)
		return result
	}
	return nil
}

// ListAllEnvPrefixes 列出所有已注册的环境前缀
func ListAllEnvPrefixes() map[EnvironmentType][]string {
	envPrefixesMutex.RLock()
	defer envPrefixesMutex.RUnlock()

	result := make(map[EnvironmentType][]string)
	for env, prefixes := range DefaultEnvPrefixes {
		copied := make([]string, len(prefixes))
		copy(copied, prefixes)
		result[env] = copied
	}
	return result
}

// String 返回环境类型的字符串表示
func (e EnvironmentType) String() string {
	return string(e)
}

// IsValid 检查环境类型是否有效
func (e EnvironmentType) IsValid() bool {
	return defaultEnvManager.IsRegistered(e)
}

// 定义上下文键类型
type ContextKey string

// String 返回上下文键的字符串表示
func (c ContextKey) String() string {
	return string(c)
}

var (
	envContextKey   ContextKey = "APP_ENV" // 默认上下文键
	contextKeyMutex sync.Mutex             // 互斥锁，用于保护 envContextKey
)

// ContextKeyOptions 定义上下文键的选项
type ContextKeyOptions struct {
	Key   ContextKey
	Value EnvironmentType
}

// EnvironmentManager 环境管理器
type EnvironmentManager struct {
	environments map[EnvironmentType][]string // 环境类型到别名的映射
	mu           sync.RWMutex                 // 读写锁
}

// NewEnvironmentManager 创建新的环境管理器
func NewEnvironmentManager() *EnvironmentManager {
	manager := &EnvironmentManager{
		environments: make(map[EnvironmentType][]string),
	}

	// 注册默认环境和它们的别名
	manager.registerDefaultEnvironments()
	return manager
}

// registerDefaultEnvironments 注册默认环境
// 直接从 DefaultEnvPrefixes 读取，保持单一数据源
func (em *EnvironmentManager) registerDefaultEnvironments() {
	for envType, prefixes := range DefaultEnvPrefixes {
		em.registerEnvironmentInternal(envType, prefixes...)
	}
}

// registerEnvironmentInternal 内部注册方法，不会反向更新 DefaultEnvPrefixes
func (em *EnvironmentManager) registerEnvironmentInternal(envType EnvironmentType, aliases ...string) {
	em.mu.Lock()
	defer em.mu.Unlock()

	// 添加环境类型本身作为别名
	allAliases := []string{string(envType)}
	allAliases = append(allAliases, aliases...)

	// 标准化别名（转换为小写并去除空格）
	for i, alias := range allAliases {
		allAliases[i] = strings.ToLower(strings.TrimSpace(alias))
	}

	em.environments[envType] = allAliases
	log.Printf("注册环境类型: %s, 别名: %v", envType, allAliases)
}

// RegisterEnvironment 注册环境类型及其别名
// 同时会更新 DefaultEnvPrefixes，保持环境别名和配置文件前缀的一致性
func (em *EnvironmentManager) RegisterEnvironment(envType EnvironmentType, aliases ...string) {
	// 先更新 DefaultEnvPrefixes
	envPrefixesMutex.Lock()
	DefaultEnvPrefixes[envType] = aliases
	envPrefixesMutex.Unlock()

	// 再注册到环境管理器
	em.registerEnvironmentInternal(envType, aliases...)
}

// IsRegistered 检查环境类型是否已注册
func (em *EnvironmentManager) IsRegistered(envType EnvironmentType) bool {
	em.mu.RLock()
	defer em.mu.RUnlock()

	_, exists := em.environments[envType]
	return exists
}

// IsEnvironment 检查给定的环境字符串是否属于指定的环境类型
func (em *EnvironmentManager) IsEnvironment(envString string, envType EnvironmentType) bool {
	em.mu.RLock()
	defer em.mu.RUnlock()

	envString = strings.ToLower(strings.TrimSpace(envString))

	if aliases, exists := em.environments[envType]; exists {
		for _, alias := range aliases {
			if alias == envString {
				return true
			}
		}
	}

	return false
}

// DetectEnvironmentType 检测环境字符串对应的环境类型
func (em *EnvironmentManager) DetectEnvironmentType(envString string) (EnvironmentType, bool) {
	em.mu.RLock()
	defer em.mu.RUnlock()

	envString = strings.ToLower(strings.TrimSpace(envString))

	for envType, aliases := range em.environments {
		for _, alias := range aliases {
			if alias == envString {
				return envType, true
			}
		}
	}

	return "", false
}

// 全局环境管理器实例
var defaultEnvManager = NewEnvironmentManager()

// init 包初始化时自动初始化全局环境实例
func init() {
	// 自动初始化全局环境实例
	initGlobalEnvironment()
	log.Printf("go-config包已自动初始化，当前环境: %s", GetCurrentEnvironment())
}

// GetGlobalEnvManager 获取全局环境管理器
// 可用于注册自定义环境类型及其别名
func GetGlobalEnvManager() *EnvironmentManager {
	return defaultEnvManager
}

// EnvironmentCallback 环境变更回调函数类型
type EnvironmentCallback func(oldEnv, newEnv EnvironmentType) error

// EnvironmentCallbackInfo 环境回调信息
type EnvironmentCallbackInfo struct {
	ID       string              // 回调唯一标识
	Callback EnvironmentCallback // 回调函数
	Priority int                 // 优先级（数字越小优先级越高）
	Async    bool                // 是否异步执行
}

// 环境配置
type Environment struct {
	CheckFrequency time.Duration                       // 检查频率
	mu             sync.RWMutex                        // 读写互斥锁
	quit           chan struct{}                       // 用于停止监控的信道
	Value          EnvironmentType                     // 当前环境值
	registeredEnvs []EnvironmentType                   // 注册的环境类型
	callbacks      map[string]*EnvironmentCallbackInfo // 环境变更回调
}

// NewEnvironment 创建一个新的 Environment 实例，并将环境变量写入上下文
func NewEnvironment() *Environment {
	envInstance := &Environment{
		Value:          GetEnvironment(),
		CheckFrequency: 2 * time.Second, // 默认检查频率
		quit:           make(chan struct{}),
		registeredEnvs: []EnvironmentType{Dev, Sit, Fat, Uat, Prod}, // 默认注册的环境类型
		callbacks:      make(map[string]*EnvironmentCallbackInfo),   // 初始化回调映射
	}
	if err := setEnv(envContextKey, envInstance.Value); err != nil {
		log.Fatalf("初始化环境失败: %v", err)
	}

	go envInstance.watchEnv() // 启动监控环境变量的 goroutine

	return envInstance
}

// RegisterCallback 注册环境变更回调
func (e *Environment) RegisterCallback(id string, callback EnvironmentCallback, priority int, async bool) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if id == "" {
		return fmt.Errorf("回调ID不能为空")
	}

	if _, exists := e.callbacks[id]; exists {
		return fmt.Errorf("回调ID %s 已存在", id)
	}

	e.callbacks[id] = &EnvironmentCallbackInfo{
		ID:       id,
		Callback: callback,
		Priority: priority,
		Async:    async,
	}

	log.Printf("注册环境变更回调: %s, 优先级: %d, 异步: %v", id, priority, async)
	return nil
}

// UnregisterCallback 注销环境变更回调
func (e *Environment) UnregisterCallback(id string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, exists := e.callbacks[id]; !exists {
		return fmt.Errorf("回调ID %s 不存在", id)
	}

	delete(e.callbacks, id)
	log.Printf("注销环境变更回调: %s", id)
	return nil
}

// triggerCallbacks 触发环境变更回调
func (e *Environment) triggerCallbacks(oldEnv, newEnv EnvironmentType) {
	e.mu.RLock()
	// 复制回调列表避免死锁
	callbacks := make([]*EnvironmentCallbackInfo, 0, len(e.callbacks))
	for _, cb := range e.callbacks {
		callbacks = append(callbacks, cb)
	}
	e.mu.RUnlock()

	if len(callbacks) == 0 {
		return
	}

	// 按优先级排序
	for i := 0; i < len(callbacks)-1; i++ {
		for j := i + 1; j < len(callbacks); j++ {
			if callbacks[i].Priority > callbacks[j].Priority {
				callbacks[i], callbacks[j] = callbacks[j], callbacks[i]
			}
		}
	}

	log.Printf("触发 %d 个环境变更回调: %s -> %s", len(callbacks), oldEnv, newEnv)

	// 执行回调
	for _, cb := range callbacks {
		if cb.Async {
			go e.executeCallback(cb, oldEnv, newEnv)
		} else {
			e.executeCallback(cb, oldEnv, newEnv)
		}
	}
}

// executeCallback 执行单个回调
func (e *Environment) executeCallback(cb *EnvironmentCallbackInfo, oldEnv, newEnv EnvironmentType) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("环境变更回调 %s 执行时发生panic: %v", cb.ID, r)
		}
	}()

	if err := cb.Callback(oldEnv, newEnv); err != nil {
		log.Printf("环境变更回调 %s 执行失败: %v", cb.ID, err)
	} else {
		log.Printf("环境变更回调 %s 执行成功", cb.ID)
	}
}

// ListCallbacks 列出所有环境变更回调
func (e *Environment) ListCallbacks() []string {
	e.mu.RLock()
	defer e.mu.RUnlock()

	ids := make([]string, 0, len(e.callbacks))
	for id := range e.callbacks {
		ids = append(ids, id)
	}
	return ids
}

// ClearCallbacks 清除所有环境变更回调
func (e *Environment) ClearCallbacks() {
	e.mu.Lock()
	defer e.mu.Unlock()

	count := len(e.callbacks)
	e.callbacks = make(map[string]*EnvironmentCallbackInfo)
	log.Printf("已清除 %d 个环境变更回调", count)
}

// RegisterEnvironment 注册新的环境类型
func (e *Environment) RegisterEnvironment(env EnvironmentType) error {
	e.mu.Lock()         // 获取写锁
	defer e.mu.Unlock() // 确保在函数结束时释放锁

	// 检查是否已经存在该环境类型
	for _, registeredEnv := range e.registeredEnvs {
		if registeredEnv == env {
			return fmt.Errorf("环境类型 %s 已注册", env)
		}
	}

	// 注册新的环境类型
	e.registeredEnvs = append(e.registeredEnvs, env)
	log.Printf("环境类型 %s 注册成功", env)
	return nil
}

// SetContextKey 设置上下文键
// options: 指定的上下文键选项，如果为 nil，将使用默认值
func SetContextKey(options *ContextKeyOptions) {
	contextKeyMutex.Lock()         // 获取锁
	defer contextKeyMutex.Unlock() // 确保在函数结束时释放锁

	// 如果没有提供 options，则创建一个默认的 ContextKeyOptions
	if options == nil {
		options = &ContextKeyOptions{}
	}

	// 如果没有提供 Key，则使用默认值
	if options.Key == "" {
		options.Key = envContextKey // 使用当前的全局上下文键
	}

	// 如果没有提供 Value 则使用默认环境值
	if options.Value == "" {
		options.Value = DefaultEnv // 使用默认环境值
	}

	envContextKey = options.Key

	// 设置环境变量
	if err := setEnv(envContextKey, options.Value); err != nil {
		log.Printf("设置环境变量失败: %v", err)
	}
}

// GetContextKey 获取当前的上下文键
func GetContextKey() ContextKey {
	contextKeyMutex.Lock()         // 获取锁
	defer contextKeyMutex.Unlock() // 确保在函数结束时释放锁

	return envContextKey // 返回当前的上下文键
}

// SetEnvironment 设置环境变量
func (e *Environment) SetEnvironment(value EnvironmentType) *Environment {
	e.mu.Lock()         // 获取写锁
	defer e.mu.Unlock() // 确保在函数结束时释放锁

	if err := setEnv(envContextKey, value); err != nil {
		log.Printf("设置环境 %s 失败: %v", value, err)
	}
	return e // 返回当前环境实例以支持链式调用
}

// CheckAndUpdateEnv 检查并更新环境变量
func (e *Environment) CheckAndUpdateEnv() {
	e.mu.RLock() // 获取读锁
	currentOsEnv := os.Getenv(envContextKey.String())
	oldValue := e.Value
	e.mu.RUnlock() // 释放读锁

	if currentOsEnv == "" {
		log.Printf("环境变量 %s 当前为空", envContextKey)
		return
	}

	newEnv := EnvironmentType(currentOsEnv)
	if newEnv != oldValue {
		e.mu.Lock()
		e.Value = newEnv // 更新环境值
		e.mu.Unlock()

		e.SetEnvironment(newEnv) // 更新环境
		log.Printf("环境变量 %s 已更新为 %s", envContextKey, newEnv)

		// 触发环境变更回调
		e.triggerCallbacks(oldValue, newEnv)
	}
}

// watchEnv 监控环境变量的变化
func (e *Environment) watchEnv() {
	for {
		e.mu.RLock()                         // 获取读锁
		currentFrequency := e.CheckFrequency // 读取当前频率
		e.mu.RUnlock()                       // 释放读锁

		ticker := time.NewTicker(currentFrequency) // 使用当前的检查频率
		defer ticker.Stop()

		select {
		case <-ticker.C:
			e.CheckAndUpdateEnv() // 检查并更新环境变量
		case <-e.quit:
			log.Println("手动终止取消监控环境变量")
			return // 退出监控
		}
	}
}

// SetCheckFrequency 设置检查频率并支持链式调用
func (e *Environment) SetCheckFrequency(frequency time.Duration) *Environment {
	e.mu.Lock()         // 获取写锁
	defer e.mu.Unlock() // 确保在函数结束时释放锁

	e.CheckFrequency = frequency // 更新检查频率
	return e                     // 返回当前环境实例以支持链式调用
}

// StopWatch 停止监控环境变量
func (e *Environment) StopWatch() {
	close(e.quit) // 关闭 quit 信道以停止监控
}

// GetEnvironment 从上下文中获取环境变量并解析为正确的环境类型
func GetEnvironment() EnvironmentType {
	// 尝试获取环境变量
	currentOsEnv, err := getEnv(envContextKey)
	if err != nil {
		log.Printf("获取环境变量 %s 失败: %v", envContextKey, err)
		return DefaultEnv
	}

	// 尝试检测环境类型
	if detectedEnv, found := defaultEnvManager.DetectEnvironmentType(currentOsEnv.String()); found {
		return detectedEnv
	}

	// 如果检测失败，返回原始值
	log.Printf("警告: 环境变量 %s='%s' 无法映射到已知环境类型，使用原始值", envContextKey, currentOsEnv)
	return currentOsEnv
}

// setEnv 设置环境变量并记录日志
func setEnv(key ContextKey, value EnvironmentType) error {
	if err := os.Setenv(key.String(), value.String()); err != nil {
		return err // 返回错误
	}
	log.Printf("环境变量 %s 设置为 %s", key, value)
	return nil
}

// getEnv 获取环境变量并记录日志
func getEnv(key ContextKey) (EnvironmentType, error) {
	value := os.Getenv(key.String())
	if value == "" {
		log.Printf("环境变量 %s 未设置", key)
		return "", fmt.Errorf("环境变量 %s 未设置", key)
	}
	log.Printf("环境变量 %s 的值为 %s", key, value)
	return EnvironmentType(value), nil
}

// ClearEnv 清除环境变量并等待锁
func (e *Environment) ClearEnv() {
	e.mu.Lock()         // 获取写锁
	defer e.mu.Unlock() // 确保在函数结束时释放锁

	if err := os.Unsetenv(envContextKey.String()); err != nil {
		log.Printf("清除环境变量 %s 失败: %v", envContextKey.String(), err)
	} else {
		log.Printf("环境变量 %s 已清除", envContextKey.String())
	}
}

// 定义一个内部接口，隐藏实现细节
type EnvironmentInterface interface {
	SetEnvironment(value EnvironmentType) *Environment
	CheckAndUpdateEnv()
	SetCheckFrequency(frequency time.Duration) *Environment
	StopWatch()
	ClearEnv()
	RegisterCallback(id string, callback EnvironmentCallback, priority int, async bool) error
	UnregisterCallback(id string) error
	ListCallbacks() []string
	ClearCallbacks()
}

// Ensure Environment 实现了 EnvironmentInterface
var _ EnvironmentInterface = (*Environment)(nil)

// Default 返回默认的 Environment 指针，支持链式调用
func Default() *Environment {
	config := DefaultEnvironment()
	return &config
}

// DefaultEnvironment 返回默认的 Environment 值
func DefaultEnvironment() Environment {
	return Environment{
		CheckFrequency: 2 * time.Second,
		Value:          DefaultEnv,
		quit:           make(chan struct{}),
		callbacks:      make(map[string]*EnvironmentCallbackInfo),
	}
}

// 全局环境实例，自动初始化
var globalEnvironment *Environment
var globalEnvOnce sync.Once

// initGlobalEnvironment 初始化全局环境实例
func initGlobalEnvironment() {
	globalEnvOnce.Do(func() {
		globalEnvironment = NewEnvironment()
		log.Printf("全局环境实例已自动初始化，当前环境: %s", globalEnvironment.Value)
	})
}

// GetGlobalEnvironment 获取全局环境实例，如果未初始化则自动初始化
func GetGlobalEnvironment() *Environment {
	initGlobalEnvironment()
	return globalEnvironment
}

// GetCurrentEnvironment 获取当前环境类型（便捷函数）
func GetCurrentEnvironment() EnvironmentType {
	globalEnv := GetGlobalEnvironment()
	globalEnv.mu.RLock()
	defer globalEnv.mu.RUnlock()
	return globalEnv.Value
}

// SetCurrentEnvironment 设置当前环境类型（便捷函数）
func SetCurrentEnvironment(env EnvironmentType) {
	globalEnv := GetGlobalEnvironment()
	globalEnv.SetEnvironment(env)
	// 立即更新全局环境实例的值
	globalEnv.mu.Lock()
	globalEnv.Value = env
	globalEnv.mu.Unlock()
}

// 环境类型判断函数

// IsDev 判断当前环境是否为开发环境
func IsDev() bool {
	env := GetCurrentEnvironment()
	return defaultEnvManager.IsEnvironment(env.String(), EnvDevelopment)
}

// IsTest 判断当前环境是否为测试环境
func IsTest() bool {
	env := GetCurrentEnvironment()
	return defaultEnvManager.IsEnvironment(env.String(), EnvTest)
}

// IsStaging 判断当前环境是否为预发布环境
func IsStaging() bool {
	env := GetCurrentEnvironment()
	return defaultEnvManager.IsEnvironment(env.String(), EnvStaging)
}

// IsUAT 判断当前环境是否为用户验收测试环境
func IsUAT() bool {
	env := GetCurrentEnvironment()
	return defaultEnvManager.IsEnvironment(env.String(), EnvUAT)
}

// IsProduction 判断当前环境是否为生产环境
func IsProduction() bool {
	env := GetCurrentEnvironment()
	return defaultEnvManager.IsEnvironment(env.String(), EnvProduction)
}

// IsLocal 判断当前环境是否为本地环境
func IsLocal() bool {
	env := GetCurrentEnvironment()
	return defaultEnvManager.IsEnvironment(env.String(), EnvLocal)
}

// IsDebug 判断当前环境是否为调试环境
func IsDebug() bool {
	env := GetCurrentEnvironment()
	return defaultEnvManager.IsEnvironment(env.String(), EnvDebug)
}

// IsDemo 判断当前环境是否为演示环境
func IsDemo() bool {
	env := GetCurrentEnvironment()
	return defaultEnvManager.IsEnvironment(env.String(), EnvDemo)
}

// IsIntegration 判断当前环境是否为集成环境
func IsIntegration() bool {
	env := GetCurrentEnvironment()
	return defaultEnvManager.IsEnvironment(env.String(), EnvIntegration)
}

// IsEnvironment 通用环境判断函数
// env: 要检查的环境类型
func IsEnvironment(env EnvironmentType) bool {
	currentEnv := GetCurrentEnvironment()
	return defaultEnvManager.IsEnvironment(currentEnv.String(), env)
}

// IsAnyOf 判断当前环境是否为指定环境类型中的任意一种
// envs: 要检查的环境类型列表
func IsAnyOf(envs ...EnvironmentType) bool {
	currentEnv := GetCurrentEnvironment()
	for _, env := range envs {
		if defaultEnvManager.IsEnvironment(currentEnv.String(), env) {
			return true
		}
	}
	return false
}

// GetEnvironmentLevel 获取环境级别（用于比较环境的重要程度）
// 数字越小表示环境级别越高（生产环境最高）
func GetEnvironmentLevel(env EnvironmentType) int {
	switch env {
	case EnvLocal, EnvDebug:
		return 1
	case EnvDevelopment:
		return 2
	case EnvTest:
		return 3
	case EnvIntegration:
		return 4
	case EnvStaging, EnvUAT:
		return 5
	case EnvDemo:
		return 6
	case EnvProduction:
		return 10
	default:
		return 0
	}
}

// GetCurrentEnvironmentLevel 获取当前环境级别
func GetCurrentEnvironmentLevel() int {
	return GetEnvironmentLevel(GetCurrentEnvironment())
}

// IsProductionLevel 判断当前环境是否为生产级别
func IsProductionLevel() bool {
	return GetCurrentEnvironmentLevel() >= 10
}

// IsTestingLevel 判断当前环境是否为测试级别
func IsTestingLevel() bool {
	level := GetCurrentEnvironmentLevel()
	return level >= 3 && level < 10
}

// IsDevelopmentLevel 判断当前环境是否为开发级别
func IsDevelopmentLevel() bool {
	return GetCurrentEnvironmentLevel() <= 2
}
