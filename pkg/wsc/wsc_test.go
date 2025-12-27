/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 11:30:00
 * @FilePath: \go-config\pkg\wsc\wsc_test.go
 * @Description: WebSocket配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package wsc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWSC_Clone(t *testing.T) {
	original := &WSC{
		Enabled:             true,
		Network:             "tcp",
		NodeIP:              "127.0.0.1",
		NodePort:            8080,
		Path:                "/ws",
		HeartbeatInterval:   30,
		ClientTimeout:       300,
		MessageBufferSize:   1024,
		MaxPendingQueueSize: 10000,
		WebSocketOrigins:    []string{"http://localhost"},
		ReadTimeout:         60 * time.Second,
		WriteTimeout:        60 * time.Second,
		IdleTimeout:         300 * time.Second,
		MaxMessageSize:      1048576,
		MinRecTime:          2 * time.Second,
		MaxRecTime:          30 * time.Second,
		RecFactor:           1.5,
		AutoReconnect:       true,
		MaxRetries:          3,
		BaseDelay:           1 * time.Second,
		MaxDelay:            30 * time.Second,
		AckTimeout:          5 * time.Second,
		AckMaxRetries:       3,
		EnableAck:           true,
		BackoffFactor:       2.0,
		Jitter:              true,
		SSEHeartbeat:        30,
		SSETimeout:          300,
		SSEMessageBuffer:    100,
	}

	cloned := original.Clone().(*WSC)

	// 验证克隆后的值相等
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.Network, cloned.Network)
	assert.Equal(t, original.NodeIP, cloned.NodeIP)
	assert.Equal(t, original.NodePort, cloned.NodePort)
	assert.Equal(t, original.Path, cloned.Path)
	assert.Equal(t, original.HeartbeatInterval, cloned.HeartbeatInterval)
	assert.Equal(t, original.ClientTimeout, cloned.ClientTimeout)
	assert.Equal(t, original.MessageBufferSize, cloned.MessageBufferSize)
	assert.Equal(t, original.AutoReconnect, cloned.AutoReconnect)
	assert.Equal(t, original.EnableAck, cloned.EnableAck)

	// 验证 slice 深拷贝
	assert.Equal(t, original.WebSocketOrigins, cloned.WebSocketOrigins)

	// 修改原始对象不应影响克隆对象
	original.NodePort = 9090
	original.Path = "/websocket"
	original.WebSocketOrigins[0] = "http://newhost"

	assert.NotEqual(t, original.NodePort, cloned.NodePort)
	assert.NotEqual(t, original.Path, cloned.Path)
	assert.NotEqual(t, original.WebSocketOrigins[0], cloned.WebSocketOrigins[0])
}
