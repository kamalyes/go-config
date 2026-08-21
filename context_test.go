/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-26 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-08-21 20:09:56
 * @FilePath: \go-config\context_test.go
 * @Description: 上下文测试
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- ContextManager 取值方法 ---

func TestContextManager_GetCurrentConfigAndEnvironment(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()
	cm := NewContextManager(env, nil)

	// 初始配置为空，环境为 env.Value
	assert.Nil(t, cm.GetCurrentConfig())
	assert.Equal(t, env.Value, cm.GetCurrentEnvironment())

	cm.UpdateConfig("my-config")
	assert.Equal(t, "my-config", cm.GetCurrentConfig())
}

// --- 从上下文取值（含热重载器/元数据） ---

func TestGetHotReloaderFromContext(t *testing.T) {
	// 空上下文
	reloader, ok := GetHotReloaderFromContext(context.Background())
	assert.False(t, ok)
	assert.Nil(t, reloader)

	// 带热重载器的上下文
	configFile := createTempConfigFile(t, "name: test")
	defer os.Remove(configFile)
	config := &TestConfig{}
	v, err := createViper(configFile)
	require.NoError(t, err)
	require.NoError(t, v.Unmarshal(config))
	hr, err := NewHotReloader(config, v, configFile, DefaultHotReloadConfig())
	require.NoError(t, err)

	env := NewEnvironment()
	defer env.StopWatch()
	cm := NewContextManager(env, hr)
	ctx := cm.WithConfig(context.Background())

	got, ok := GetHotReloaderFromContext(ctx)
	assert.True(t, ok)
	assert.NotNil(t, got)
}

func TestGetMetadataFromContext(t *testing.T) {
	// 空上下文
	meta, ok := GetMetadataFromContext(context.Background())
	assert.False(t, ok)
	assert.Nil(t, meta)

	// 带元数据的上下文
	env := NewEnvironment()
	defer env.StopWatch()
	cm := NewContextManager(env, nil)
	cm.SetMetadata("key", "value")
	ctx := cm.WithConfig(context.Background())

	meta, ok = GetMetadataFromContext(ctx)
	assert.True(t, ok)
	assert.Contains(t, meta, "key")
}

// --- ContextKeyHelper 截止时间/超时 ---

func TestContextKeyHelper_NewContextWithDeadline(t *testing.T) {
	// 初始化全局上下文管理器
	env := NewEnvironment()
	defer env.StopWatch()
	InitializeContextManager(env, nil)
	defer ClearGlobalContextManager()

	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := ContextHelper.NewContextWithDeadline(deadline)
	defer cancel()

	assert.NotNil(t, ctx)
	d, ok := ctx.Deadline()
	assert.True(t, ok)
	assert.True(t, d.Equal(deadline) || d.Before(deadline))
}

func TestContextKeyHelper_NewContextWithTimeout_OK(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()
	InitializeContextManager(env, nil)
	defer ClearGlobalContextManager()

	ctx, cancel := ContextHelper.NewContextWithTimeout(2 * time.Second)
	defer cancel()
	assert.NotNil(t, ctx)
	_, ok := ctx.Deadline()
	assert.True(t, ok)
}

// --- MustGet* panic 路径 ---

func TestContextKeyHelper_MustGetConfig_Panic(t *testing.T) {
	// 没有配置的上下文应 panic
	assert.Panics(t, func() {
		ContextHelper.MustGetConfig(context.Background())
	})
}

func TestContextKeyHelper_MustGetConfig_OK(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()
	InitializeContextManager(env, nil)
	defer ClearGlobalContextManager()

	cm := GetGlobalContextManager()
	require.NotNil(t, cm)
	cm.UpdateConfig("must-config-value")

	ctx := ContextHelper.NewConfigContext()
	assert.Equal(t, "must-config-value", ContextHelper.MustGetConfig(ctx))
}

func TestContextKeyHelper_MustGetEnvironment_Panic(t *testing.T) {
	assert.Panics(t, func() {
		ContextHelper.MustGetEnvironment(context.Background())
	})
}

func TestContextKeyHelper_MustGetEnvironment_OK(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()
	InitializeContextManager(env, nil)
	defer ClearGlobalContextManager()

	ctx := ContextHelper.NewConfigContext()
	// NewConfigContext 通过 WithGlobalConfig 注入了环境
	envType := ContextHelper.MustGetEnvironment(ctx)
	assert.NotEmpty(t, envType)
}

func TestContextKeyHelper_IsEnvironment(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()
	InitializeContextManager(env, nil)
	defer ClearGlobalContextManager()

	ctx := ContextHelper.NewConfigContext()
	assert.True(t, ContextHelper.IsEnvironment(ctx, env.Value))
	assert.False(t, ContextHelper.IsEnvironment(ctx, EnvironmentType("never-matches")))
}
