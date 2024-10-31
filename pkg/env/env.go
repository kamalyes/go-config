/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 13:06:18
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

	"github.com/kamalyes/go-config/internal"
)

// EnvironmentType 定义环境类型
type EnvironmentType string

// 定义有效的环境常量
const (
	Dev        EnvironmentType = "dev"
	Fat        EnvironmentType = "fat"
	Uat        EnvironmentType = "uat"
	Pro        EnvironmentType = "pro"
	DefaultEnv EnvironmentType = Dev
	AppEnvKey  string          = "APP_ENV"
)

// 环境配置
type Environment struct {
	value EnvironmentType
}

// 自定义上下文键类型
type EnvContextKey struct{}

// 使用自定义上下文键
var envKey = EnvContextKey{}

// 环境映射
var environments = map[EnvironmentType]*Environment{
	Dev: {value: Dev},
	Fat: {value: Fat},
	Uat: {value: Uat},
	Pro: {value: Pro},
}

// NewEnvironment 创建一个新的 Environment 实例
func NewEnvironment(value EnvironmentType) *Environment {
	var envInstance *Environment

	internal.LockFunc(func() {
		envInstance = &Environment{value: value}
	})

	return envInstance
}

func ClearEnv() {
	os.Unsetenv(AppEnvKey)
}

// 从上下文获取当前配置的环境
func FromContext(ctx context.Context) *Environment {
	if env, ok := ctx.Value(envKey).(*Environment); ok {
		return env
	}
	// 如果上下文中没有环境，则调用 Active() 获取环境变量
	return Active()
}

// Active 获取当前配置的环境
func Active() *Environment {
	env := os.Getenv(AppEnvKey) // 从环境变量中获取当前环境
	if env == "" {
		log.Printf("未设置 APP_ENV 环境变量，使用默认环境 %v", DefaultEnv)
		return environments[DefaultEnv] // 使用映射中的实例
	}

	// 如果环境在映射中不存在，使用自定义环境
	return &Environment{value: EnvironmentType(strings.ToLower(strings.TrimSpace(env)))}
}

// NewEnv 创建新的环境并返回包含该环境的上下文
func NewEnv(ctx context.Context, env EnvironmentType) context.Context {
	newEnv := &Environment{value: env}
	return context.WithValue(ctx, envKey, newEnv)
}

// Value 返回当前环境的值
func (e *Environment) Value() EnvironmentType {
	return e.value
}

// IsEnvironment 检查当前环境是否与给定的环境匹配
func (e *Environment) IsEnvironment(env EnvironmentType) bool {
	return e.value == env
}

// String 返回环境的字符串表示
func (e *Environment) String() string {
	return string(e.value)
}
