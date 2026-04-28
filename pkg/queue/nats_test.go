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
		ModuleName:     "test-nats",
		URL:            "nats://localhost:4222",
		Name:           "test-client",
		Username:       "user",
		Password:       "pass",
		Token:          "",
		JetStream:      true,
		StreamName:     "TEST_STREAM",
		ConnectTimeout: 5,
		ReconnectWait:  1,
		MaxReconnects:  20,
		ChannelPrefix:  "casbin.policy",
		Source:         "node-1",
	}

	cloned := original.Clone().(*Nats)

	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.URL, cloned.URL)
	assert.Equal(t, original.Name, cloned.Name)
	assert.Equal(t, original.JetStream, cloned.JetStream)
	assert.Equal(t, original.StreamName, cloned.StreamName)
	assert.Equal(t, original.ChannelPrefix, cloned.ChannelPrefix)
	assert.Equal(t, original.Source, cloned.Source)

	// 修改原始对象不应影响克隆对象
	original.URL = "nats://other:4222"
	original.MaxReconnects = 99
	assert.NotEqual(t, original.URL, cloned.URL)
	assert.NotEqual(t, original.MaxReconnects, cloned.MaxReconnects)
}

func TestNats_Default(t *testing.T) {
	cfg := DefaultNats()
	assert.Equal(t, "nats", cfg.ModuleName)
	assert.Equal(t, "nats://127.0.0.1:4222", cfg.URL)
	assert.False(t, cfg.JetStream)
	assert.Equal(t, 10, cfg.MaxReconnects)
}

func TestNats_Builders(t *testing.T) {
	cfg := DefaultNatsPtr().
		WithURL("nats://broker:4222").
		WithJetStream(true).
		WithStreamName("CASBIN_POLICY").
		WithChannelPrefix("casbin.policy").
		WithSource("node-xyz")

	assert.Equal(t, "nats://broker:4222", cfg.URL)
	assert.True(t, cfg.JetStream)
	assert.Equal(t, "CASBIN_POLICY", cfg.StreamName)
	assert.Equal(t, "casbin.policy", cfg.ChannelPrefix)
	assert.Equal(t, "node-xyz", cfg.Source)
}

func TestNats_Validate(t *testing.T) {
	cfg := DefaultNats()
	assert.NoError(t, cfg.Validate())

	bad := cfg
	bad.URL = ""
	assert.Error(t, bad.Validate())
}
