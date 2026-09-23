/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-04-23 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-04-23 00:00:00
 * @FilePath: \go-config\pkg\queue\nats_test.go
 * @Description: NATS 配置测试
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNats_Clone(t *testing.T) {
	original := &Nats{
		ModuleName:       "test-nats",
		Enabled:          true,
		URL:              "nats://localhost:4222",
		Name:             "test-client",
		Username:         "user",
		Password:         "pass",
		Token:            "",
		JetStream:        true,
		Replicas:         3,
		Storage:          NatsStorageFile,
		Retention:        NatsRetentionLimits,
		MaxAge:           3600,
		MaxBytes:         1024,
		DuplicatesWindow: 60,
		ConnectTimeout:   5,
		ReconnectWait:    1,
		MaxReconnects:    20,
	}

	cloned := original.Clone().(*Nats)

	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.URL, cloned.URL)
	assert.Equal(t, original.Name, cloned.Name)
	assert.Equal(t, original.JetStream, cloned.JetStream)
	assert.Equal(t, original.Replicas, cloned.Replicas)
	assert.Equal(t, original.Storage, cloned.Storage)
	assert.Equal(t, original.Retention, cloned.Retention)
	assert.Equal(t, original.MaxAge, cloned.MaxAge)
	assert.Equal(t, original.MaxBytes, cloned.MaxBytes)
	assert.Equal(t, original.DuplicatesWindow, cloned.DuplicatesWindow)

	// 修改原始对象不应影响克隆对象
	original.URL = "nats://other:4222"
	original.MaxReconnects = 99
	assert.NotEqual(t, original.URL, cloned.URL)
	assert.NotEqual(t, original.MaxReconnects, cloned.MaxReconnects)
}

func TestNats_Default(t *testing.T) {
	cfg := DefaultNats()
	assert.Equal(t, "nats", cfg.ModuleName)
	assert.False(t, cfg.Enabled)
	assert.False(t, cfg.IsEnabled())
	assert.Equal(t, "nats://127.0.0.1:4222", cfg.URL)
	assert.False(t, cfg.JetStream)
	assert.Equal(t, 10, cfg.MaxReconnects)
}

func TestNats_Builders(t *testing.T) {
	cfg := DefaultNatsPtr().
		WithEnabled(true).
		WithURL("nats://broker:4222").
		WithJetStream(true).
		WithReplicas(3).
		WithStorage(NatsStorageMemory).
		WithRetention(NatsRetentionInterest).
		WithMaxAge(7200).
		WithMaxBytes(2048).
		WithDuplicatesWindow(30).
		WithPingInterval(30).
		WithMaxPingsOut(5)

	assert.Equal(t, "nats://broker:4222", cfg.URL)
	assert.True(t, cfg.Enabled)
	assert.True(t, cfg.IsEnabled())
	assert.True(t, cfg.JetStream)
	assert.Equal(t, 3, cfg.Replicas)
	assert.Equal(t, NatsStorageMemory, cfg.Storage)
	assert.Equal(t, NatsRetentionInterest, cfg.Retention)
	assert.Equal(t, int64(7200), cfg.MaxAge)
	assert.Equal(t, int64(2048), cfg.MaxBytes)
	assert.Equal(t, int64(30), cfg.DuplicatesWindow)
	assert.Equal(t, 30, cfg.PingInterval)
	assert.Equal(t, 5, cfg.MaxPingsOut)
}

func TestNats_Validate(t *testing.T) {
	cfg := DefaultNats()
	assert.NoError(t, cfg.Validate())

	bad := cfg
	bad.URL = ""
	assert.Error(t, bad.Validate())
}

func TestNats_SetCopiesEnabled(t *testing.T) {
	cfg := DefaultNatsPtr()
	next := DefaultNatsPtr().WithEnabled(true)

	cfg.Set(next)

	assert.True(t, cfg.Enabled)
	assert.True(t, cfg.IsEnabled())
}

func TestNats_IsEnabledNil(t *testing.T) {
	var cfg *Nats
	assert.False(t, cfg.IsEnabled())
}
