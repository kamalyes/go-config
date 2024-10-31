/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 11:15:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 11:29:01
 * @FilePath: \go-config\tests\env_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"context"
	"os"
	"testing"

	"github.com/kamalyes/go-config/pkg/env"
	"github.com/stretchr/testify/assert"
)

func TestActiveDefaultEnv(t *testing.T) {
	defer env.ClearEnv()

	// 获取当前活动环境
	currentEnv := env.Active()

	// 断言当前环境为默认环境
	assert.Equal(t, env.DefaultEnv, currentEnv.Value())
}

func TestActiveCustomEnv(t *testing.T) {
	defer env.ClearEnv()

	// 设置 APP_ENV 环境变量
	os.Setenv("APP_ENV", "fat")

	// 获取当前活动环境
	currentEnv := env.Active()

	// 断言当前环境为 FAT
	assert.Equal(t, env.Fat, currentEnv.Value())
}

func TestFromContext(t *testing.T) {
	defer env.ClearEnv()

	// 创建一个上下文并设置环境为 UAT
	ctx := env.NewEnv(context.Background(), env.Uat)

	// 从上下文获取环境
	currentEnv := env.FromContext(ctx)

	// 断言当前环境为 UAT
	assert.Equal(t, env.Uat, currentEnv.Value())
}

func TestIsEnvironment(t *testing.T) {
	defer env.ClearEnv()

	// 创建一个上下文并设置环境为 PRO
	ctx := env.NewEnv(context.Background(), env.Pro)

	// 从上下文获取环境
	currentEnv := env.FromContext(ctx)

	// 断言当前环境是 PRO
	assert.True(t, currentEnv.IsEnvironment(env.Pro))
	assert.False(t, currentEnv.IsEnvironment(env.Dev))
}

func TestStringMethod(t *testing.T) {
	defer env.ClearEnv()

	// 使用构造函数创建环境实例
	envInstance := env.NewEnvironment(env.Uat)

	// 断言字符串表示
	assert.Equal(t, "uat", envInstance.String())
}
