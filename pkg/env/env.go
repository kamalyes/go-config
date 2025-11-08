/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 16:11:10
 * @FilePath: \go-config\pkg\env\env.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package env

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// EnvironmentType 定义环境类型
type EnvironmentType string

// 定义有效的环境常量
const (
	Dev        EnvironmentType = "dev"
	Sit        EnvironmentType = "sit"
	Fat        EnvironmentType = "fat"
	Uat        EnvironmentType = "uat"
	Prod       EnvironmentType = "prod"
	DefaultEnv EnvironmentType = Dev
)

// 定义上下文键类型
type ContextKey string

var (
	envContextKey   ContextKey = "APP_ENV" // 默认上下文键
	contextKeyMutex sync.Mutex             // 互斥锁，用于保护 envContextKey
)

// ContextKeyOptions 定义上下文键的选项
type ContextKeyOptions struct {
	Key   ContextKey
	Value EnvironmentType
}

// 环境配置
type Environment struct {
	CheckFrequency time.Duration   // 检查频率
	mu             sync.RWMutex    // 读写互斥锁
	quit           chan struct{}   // 用于停止监控的信道
	Value          EnvironmentType // 当前环境值
}

// NewEnvironment 创建一个新的 Environment 实例，并将环境变量写入上下文
func NewEnvironment() *Environment {
	envInstance := &Environment{
		Value:          GetEnvironment(),
		CheckFrequency: 2 * time.Second, // 默认检查频率
		quit:           make(chan struct{}),
	}
	if err := setEnv(envContextKey, envInstance.Value); err != nil {
		log.Fatalf("初始化环境失败: %v", err)
	}

	go envInstance.watchEnv() // 启动监控环境变量的 goroutine

	return envInstance
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
	e.mu.RUnlock() // 释放读锁

	if currentOsEnv == "" {
		log.Printf("环境变量 %s 当前为空", envContextKey)
		return
	}

	newEnv := EnvironmentType(currentOsEnv)
	if newEnv != e.Value {
		e.SetEnvironment(newEnv) // 更新环境
		log.Printf("环境变量 %s 已更新为 %s", envContextKey, newEnv)
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

// GetEnvironment 从上下文中获取环境变量
func GetEnvironment() EnvironmentType {
	// 尝试获取环境变量
	currentOsEnv, err := getEnv(envContextKey) // 使用普通赋值而非短变量声明
	if err != nil {
		log.Printf("获取环境变量 %s 失败: %v", envContextKey, err)
		return DefaultEnv
	}
	return currentOsEnv // 确保返回获取到的环境变量值
}

// String ContextKey 转 String 方法
func (c ContextKey) String() string {
	return string(c)
}

// String EnvironmentType 转 String 方法
func (e EnvironmentType) String() string {
	return string(e)
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
	}
}
