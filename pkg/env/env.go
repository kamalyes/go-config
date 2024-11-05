/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-05 17:55:56
 * @FilePath: \go-config\pkg\env\env.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package env

import (
	"context"
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
	Dev        EnvironmentType = "dev"
	Sit        EnvironmentType = "sit"
	Fat        EnvironmentType = "fat"
	Uat        EnvironmentType = "uat"
	Pro        EnvironmentType = "pro"
	DefaultEnv EnvironmentType = Dev
	AppEnvKey  string          = "APP_ENV"
)

// 环境配置
type Environment struct {
	value EnvironmentType
	mu    sync.RWMutex
	quit  chan struct{}
}

// 自定义上下文键类型
type EnvContextKey struct{}

// 使用自定义上下文键
var envKey = EnvContextKey{}

// 环境映射
var environments = map[EnvironmentType]*Environment{
	Dev: {value: Dev, quit: make(chan struct{})},
	Sit: {value: Sit, quit: make(chan struct{})},
	Fat: {value: Fat, quit: make(chan struct{})},
	Uat: {value: Uat, quit: make(chan struct{})},
	Pro: {value: Pro, quit: make(chan struct{})},
}

// NewEnvironment 创建一个新的 Environment 实例
func NewEnvironment(value EnvironmentType) *Environment {
	setEnv(AppEnvKey, string(value))
	envInstance := &Environment{
		value: value,
		quit:  make(chan struct{}),
	}
	go envInstance.watchEnv() // 启动监控环境变量的 goroutine
	return envInstance
}

// setEnv 设置环境变量并记录日志
func setEnv(key, value string) {
	if err := os.Setenv(key, value); err != nil {
		log.Printf("设置环境变量 %s 失败: %v", key, err)
	}
	log.Printf("环境变量 %s 设置为 %s", key, value)
}

// Switch 方法用于切换环境变量
func (e *Environment) Switch(newEnv EnvironmentType) {
	e.mu.Lock()
	defer e.mu.Unlock()

	setEnv(AppEnvKey, string(newEnv))
	e.value = newEnv
}

// ClearEnv 清除环境变量
func ClearEnv() {
	os.Unsetenv(AppEnvKey)
}

// FromContext 从上下文获取当前配置的环境
func FromContext(ctx context.Context) *Environment {
	if env, ok := ctx.Value(envKey).(*Environment); ok {
		return env
	}
	return Active()
}

// Active 获取当前配置的环境
func Active() *Environment {
	env := os.Getenv(AppEnvKey)
	if env == "" {
		log.Printf("未设置 APP_ENV 环境变量，使用默认环境 %v", DefaultEnv)
		return environments[DefaultEnv]
	}
	return &Environment{value: EnvironmentType(strings.ToLower(strings.TrimSpace(env))), quit: make(chan struct{})}
}

// NewEnv 创建新的环境并返回包含该环境的上下文
func NewEnv(ctx context.Context, env EnvironmentType) context.Context {
	newEnv := NewEnvironment(env)
	return context.WithValue(ctx, envKey, newEnv)
}

// Value 返回当前环境的值
func (e *Environment) Value() EnvironmentType {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.value
}

// IsEnvironment 检查当前环境是否与给定的环境匹配
func (e *Environment) IsEnvironment(env EnvironmentType) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.value == env
}

// String 返回环境的字符串表示
func (e *Environment) String() string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return string(e.value)
}

// watchEnv 监控环境变量的变化
func (e *Environment) watchEnv() {
	lastEnv := e.Value()
	for {
		select {
		case <-time.After(1 * time.Second):
			currentEnv := os.Getenv(AppEnvKey)
			if strings.ToLower(strings.TrimSpace(currentEnv)) != string(lastEnv) {
				e.mu.Lock()
				e.value = EnvironmentType(strings.ToLower(strings.TrimSpace(currentEnv)))
				e.mu.Unlock()
				lastEnv = e.Value()
				log.Printf("环境变量 %s 已更新为 %s", AppEnvKey, lastEnv)
			}
		case <-e.quit:
			return // 退出监控
		}
	}
}

// Stop 停止监控环境变量
func (e *Environment) Stop() {
	close(e.quit)
}
