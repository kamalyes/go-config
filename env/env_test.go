/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 11:15:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 11:29:01
 * @FilePath: \go-config\env\env_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package env

import (
	"context"
	"math/rand"
	"os"
	"testing"
)

func TestActive(t *testing.T) {
	// 测试未设置 APP_ENV 环境变量时的默认行为
	os.Unsetenv("APP_ENV")
	defaultEnv := Active()
	if defaultEnv.Value() != Dev {
		t.Errorf("期望默认环境为 'dev', 但得到 '%s'", defaultEnv.Value())
	}

	// 测试设置 APP_ENV 环境变量
	os.Setenv("APP_ENV", "dev")
	devEnv := Active()
	if devEnv.Value() != Dev {
		t.Errorf("期望环境为 'dev', 但得到 '%s'", devEnv.Value())
	}
}

func TestNewEnv(t *testing.T) {
	// 创建一个背景上下文
	ctx := context.Background()

	// 使用 NewEnv 创建一个新的环境并存储在上下文中
	ctx = NewEnv(ctx, Uat)

	// 从上下文获取当前环境
	currentEnv := FromContext(ctx)
	if currentEnv.Value() != Uat {
		t.Errorf("期望环境为 'uat', 但得到 '%s'", currentEnv.Value())
	}
}

func TestFromContext(t *testing.T) {
	// 创建一个背景上下文
	ctx := context.Background()

	// 测试从上下文中获取环境
	defaultEnv := FromContext(ctx)
	if defaultEnv.Value() != Dev {
		t.Errorf("期望默认环境为 'dev', 但得到 '%s'", defaultEnv.Value())
	}

	// 使用 NewEnv 创建一个新的环境并存储在上下文中
	ctx = NewEnv(ctx, Pro)

	// 从上下文获取当前环境
	currentEnv := FromContext(ctx)
	if currentEnv.Value() != Pro {
		t.Errorf("期望环境为 'pro', 但得到 '%s'", currentEnv.Value())
	}
}

func TestIsEnvironment(t *testing.T) {
	// 创建一个背景上下文并设置环境
	ctx := NewEnv(context.Background(), Uat)

	// 从上下文获取当前环境
	currentEnv := FromContext(ctx)

	// 测试 IsEnvironment 方法
	if !currentEnv.IsEnvironment(Uat) {
		t.Errorf("期望环境为 'uat', 但得到 '%s'", currentEnv.Value())
	}
}

// 随机生成环境类型
func randomEnvironment() EnvironmentType {
	envs := []EnvironmentType{Dev, Fat, Uat, Pro}
	return envs[rand.Intn(len(envs))]
}

// 基准测试：随机设置、获取和校验环境配置
func BenchmarkRandomSetGetValidate(b *testing.B) {
	// 创建一个背景上下文
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		// 随机选择环境
		env := randomEnvironment()
		// 设置环境
		ctx = NewEnv(ctx, env)
		// 校验环境
		currentEnv := FromContext(ctx)
		// 测试 IsEnvironment 方法
		if !currentEnv.IsEnvironment(env) {
			b.Errorf("期望环境为 '%s', 但得到 '%s'", env, currentEnv.Value())
		}
	}
}
