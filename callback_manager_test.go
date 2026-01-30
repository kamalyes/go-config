/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-01-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-01-30 00:00:00
 * @FilePath: \go-config\callback_manager_test.go
 * @Description: 回调管理器测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewCallbackManager 测试创建回调管理器
func TestNewCallbackManager(t *testing.T) {
	manager := NewCallbackManager()
	assert.NotNil(t, manager)
	assert.Empty(t, manager.ListCallbacks())
}

// TestRegisterCallback 测试注册回调
func TestRegisterCallback(t *testing.T) {
	manager := NewCallbackManager()
	ctx := context.Background()

	called := false
	callback := func(ctx context.Context, event CallbackEvent) error {
		called = true
		return nil
	}

	options := DefaultCallbackOptions()
	options.ID = "test-callback"
	options.Types = []CallbackType{CallbackTypeConfigChanged}

	err := manager.RegisterCallback(callback, options)
	require.NoError(t, err)

	// 验证回调已注册
	callbacks := manager.ListCallbacks()
	assert.Contains(t, callbacks, "test-callback")

	// 触发回调
	event := CallbackEvent{
		Type:      CallbackTypeConfigChanged,
		Timestamp: time.Now(),
		Source:    "test",
	}
	err = manager.TriggerCallbacks(ctx, event)
	require.NoError(t, err)

	// 等待异步回调执行
	time.Sleep(100 * time.Millisecond)
	assert.True(t, called)
}

// TestRegisterCallback_DuplicateID 测试注册重复ID的回调
func TestRegisterCallback_DuplicateID(t *testing.T) {
	manager := NewCallbackManager()

	callback := func(ctx context.Context, event CallbackEvent) error {
		return nil
	}

	options := DefaultCallbackOptions()
	options.ID = "duplicate-id"

	err := manager.RegisterCallback(callback, options)
	require.NoError(t, err)

	// 尝试注册相同ID
	err = manager.RegisterCallback(callback, options)
	assert.Error(t, err)
}

// TestUnregisterCallback 测试注销回调
func TestUnregisterCallback(t *testing.T) {
	manager := NewCallbackManager()

	callback := func(ctx context.Context, event CallbackEvent) error {
		return nil
	}

	options := DefaultCallbackOptions()
	options.ID = "test-callback"

	err := manager.RegisterCallback(callback, options)
	require.NoError(t, err)

	// 验证回调已注册
	assert.True(t, manager.HasCallback("test-callback"))

	// 注销回调
	err = manager.UnregisterCallback("test-callback")
	require.NoError(t, err)

	// 验证回调已注销
	assert.False(t, manager.HasCallback("test-callback"))
}

// TestUnregisterCallback_NotFound 测试注销不存在的回调
func TestUnregisterCallback_NotFound(t *testing.T) {
	manager := NewCallbackManager()

	err := manager.UnregisterCallback("non-existent")
	assert.Error(t, err)
}

// TestTriggerCallbacks_MultipleCallbacks 测试触发多个回调
func TestTriggerCallbacks_MultipleCallbacks(t *testing.T) {
	manager := NewCallbackManager()
	ctx := context.Background()

	callCount := 0
	mu := &sync.Mutex{}

	// 注册多个回调
	for i := 0; i < 3; i++ {
		callback := func(ctx context.Context, event CallbackEvent) error {
			mu.Lock()
			callCount++
			mu.Unlock()
			return nil
		}

		options := DefaultCallbackOptions()
		options.ID = fmt.Sprintf("callback-%d", i)
		options.Types = []CallbackType{CallbackTypeConfigChanged}

		err := manager.RegisterCallback(callback, options)
		require.NoError(t, err)
	}

	// 触发回调
	event := CallbackEvent{
		Type:      CallbackTypeConfigChanged,
		Timestamp: time.Now(),
	}
	err := manager.TriggerCallbacks(ctx, event)
	require.NoError(t, err)

	// 等待异步回调执行
	time.Sleep(200 * time.Millisecond)

	mu.Lock()
	assert.Equal(t, 3, callCount)
	mu.Unlock()
}

// TestTriggerCallbacks_Priority 测试回调优先级
func TestTriggerCallbacks_Priority(t *testing.T) {
	manager := NewCallbackManager()
	ctx := context.Background()

	executionOrder := []int{}
	mu := &sync.Mutex{}

	// 注册不同优先级的回调
	priorities := []int{10, 5, 1}
	for _, priority := range priorities {
		p := priority
		callback := func(ctx context.Context, event CallbackEvent) error {
			mu.Lock()
			executionOrder = append(executionOrder, p)
			mu.Unlock()
			return nil
		}

		options := DefaultCallbackOptions()
		options.ID = fmt.Sprintf("callback-priority-%d", p)
		options.Priority = p
		options.Async = false // 同步执行以保证顺序
		options.Types = []CallbackType{CallbackTypeConfigChanged}

		err := manager.RegisterCallback(callback, options)
		require.NoError(t, err)
	}

	// 触发回调
	event := CallbackEvent{
		Type:      CallbackTypeConfigChanged,
		Timestamp: time.Now(),
	}
	err := manager.TriggerCallbacks(ctx, event)
	require.NoError(t, err)

	// 验证执行顺序（优先级从小到大）
	assert.Equal(t, []int{1, 5, 10}, executionOrder)
}

// TestTriggerCallbacks_ErrorHandling 测试回调错误处理
func TestTriggerCallbacks_ErrorHandling(t *testing.T) {
	manager := NewCallbackManager()
	ctx := context.Background()

	errorCallback := func(ctx context.Context, event CallbackEvent) error {
		return errors.New("callback error")
	}

	options := DefaultCallbackOptions()
	options.ID = "error-callback"
	options.Async = false
	options.Types = []CallbackType{CallbackTypeConfigChanged}

	err := manager.RegisterCallback(errorCallback, options)
	require.NoError(t, err)

	// 触发回调
	event := CallbackEvent{
		Type:      CallbackTypeConfigChanged,
		Timestamp: time.Now(),
	}
	err = manager.TriggerCallbacks(ctx, event)
	assert.Error(t, err)
}

// TestTriggerCallbacks_Filter 测试回调过滤器
func TestTriggerCallbacks_Filter(t *testing.T) {
	manager := NewCallbackManager()
	ctx := context.Background()

	called := false
	callback := func(ctx context.Context, event CallbackEvent) error {
		called = true
		return nil
	}

	options := DefaultCallbackOptions()
	options.ID = "filtered-callback"
	options.Async = false
	options.Types = []CallbackType{CallbackTypeConfigChanged}
	options.Filter = func(event CallbackEvent) bool {
		// 只处理来自特定源的事件
		return event.Source == "allowed-source"
	}

	err := manager.RegisterCallback(callback, options)
	require.NoError(t, err)

	// 触发不匹配过滤器的事件
	event := CallbackEvent{
		Type:      CallbackTypeConfigChanged,
		Timestamp: time.Now(),
		Source:    "other-source",
	}
	err = manager.TriggerCallbacks(ctx, event)
	require.NoError(t, err)
	assert.False(t, called)

	// 触发匹配过滤器的事件
	event.Source = "allowed-source"
	err = manager.TriggerCallbacks(ctx, event)
	require.NoError(t, err)
	assert.True(t, called)
}

// TestTriggerCallbacks_Timeout 测试回调超时
func TestTriggerCallbacks_Timeout(t *testing.T) {
	manager := NewCallbackManager()
	ctx := context.Background()

	slowCallback := func(ctx context.Context, event CallbackEvent) error {
		time.Sleep(2 * time.Second)
		return nil
	}

	options := DefaultCallbackOptions()
	options.ID = "slow-callback"
	options.Async = false
	options.Timeout = 100 * time.Millisecond
	options.Types = []CallbackType{CallbackTypeConfigChanged}

	err := manager.RegisterCallback(slowCallback, options)
	require.NoError(t, err)

	// 触发回调
	event := CallbackEvent{
		Type:      CallbackTypeConfigChanged,
		Timestamp: time.Now(),
	}
	err = manager.TriggerCallbacks(ctx, event)
	assert.Error(t, err)
}

// TestClearCallbacks 测试清空所有回调
func TestClearCallbacks(t *testing.T) {
	manager := NewCallbackManager()

	callback := func(ctx context.Context, event CallbackEvent) error {
		return nil
	}

	// 注册多个回调
	for i := 0; i < 3; i++ {
		options := DefaultCallbackOptions()
		options.ID = fmt.Sprintf("callback-%d", i)
		err := manager.RegisterCallback(callback, options)
		require.NoError(t, err)
	}

	assert.Len(t, manager.ListCallbacks(), 3)

	// 清空所有回调
	manager.ClearCallbacks()
	assert.Empty(t, manager.ListCallbacks())
}

// TestCallbackEvent_Creation 测试回调事件创建
func TestCallbackEvent_Creation(t *testing.T) {
	event := CallbackEvent{
		Type:        CallbackTypeConfigChanged,
		Timestamp:   time.Now(),
		Source:      "config.yaml",
		OldValue:    "old",
		NewValue:    "new",
		Environment: EnvDevelopment,
		Metadata: map[string]interface{}{
			"key": "value",
		},
	}

	assert.Equal(t, CallbackTypeConfigChanged, event.Type)
	assert.NotZero(t, event.Timestamp)
	assert.Equal(t, "config.yaml", event.Source)
	assert.Equal(t, "old", event.OldValue)
	assert.Equal(t, "new", event.NewValue)
	assert.Equal(t, EnvDevelopment, event.Environment)
	assert.Equal(t, "value", event.Metadata["key"])
}

// TestDefaultCallbackOptions 测试默认回调选项
func TestDefaultCallbackOptions(t *testing.T) {
	options := DefaultCallbackOptions()

	assert.Equal(t, 0, options.Priority)
	assert.True(t, options.Async)
	assert.Equal(t, 30*time.Second, options.Timeout)
	assert.Equal(t, 3, options.Retry)
	assert.Contains(t, options.Types, CallbackTypeConfigChanged)
}
