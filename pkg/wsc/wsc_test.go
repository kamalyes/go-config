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

	"github.com/kamalyes/go-config/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestWSC_Clone(t *testing.T) {
	original := &WSC{
		Enabled:               true,
		Network:               "tcp",
		NodeIP:                "127.0.0.1",
		NodePort:              8080,
		Path:                  "/ws",
		HeartbeatInterval:     30 * time.Second,
		ClientTimeout:         300 * time.Second,
		MessageBufferSize:     1024,
		MaxPendingQueueSize:   10000,
		WebSocketOrigins:      []string{"http://localhost"},
		ReadTimeout:           60 * time.Second,
		WriteTimeout:          60 * time.Second,
		IdleTimeout:           300 * time.Second,
		MaxMessageSize:        1048576,
		MinRecTime:            2 * time.Second,
		MaxRecTime:            30 * time.Second,
		RecFactor:             1.5,
		AutoReconnect:         true,
		AckTimeout:            5 * time.Second,
		AckMaxRetries:         3,
		EnableAck:             true,
		MessageRecordTTL:      24 * time.Hour,
		RecordCleanupInterval: 30 * time.Minute,
		RetryPolicy: &RetryPolicy{
			MaxRetries:    3,
			BaseDelay:     1 * time.Second,
			MaxDelay:      30 * time.Second,
			BackoffFactor: 2.0,
			Jitter:        true,
		},
		SSEHeartbeat:     30,
		SSETimeout:       300,
		SSEMessageBuffer: 100,
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

// TestDefaultClientAttributesHasNamespaceGroup 验证默认配置包含 NamespaceSources/GroupIDSources
func TestDefaultClientAttributesHasNamespaceGroup(t *testing.T) {
	ca := DefaultClientAttributes()
	assert.NotEmpty(t, ca.NamespaceSources, "默认配置应包含 NamespaceSources")
	assert.NotEmpty(t, ca.GroupIDSources, "默认配置应包含 GroupIDSources")

	// 验证默认来源包含 query 和 header
	hasNamespaceQuery, hasNamespaceHeader := false, false
	for _, src := range ca.NamespaceSources {
		if src.Type == common.SourceTypeQuery && src.Key == "namespace" {
			hasNamespaceQuery = true
		}
		if src.Type == common.SourceTypeHeader && src.Key == "X-Namespace" {
			hasNamespaceHeader = true
		}
	}
	assert.True(t, hasNamespaceQuery, "NamespaceSources 应包含 query:namespace")
	assert.True(t, hasNamespaceHeader, "NamespaceSources 应包含 header:X-Namespace")

	hasGroupQuery, hasGroupHeader := false, false
	for _, src := range ca.GroupIDSources {
		if src.Type == common.SourceTypeQuery && src.Key == "group_id" {
			hasGroupQuery = true
		}
		if src.Type == common.SourceTypeHeader && src.Key == "X-Group-ID" {
			hasGroupHeader = true
		}
	}
	assert.True(t, hasGroupQuery, "GroupIDSources 应包含 query:group_id")
	assert.True(t, hasGroupHeader, "GroupIDSources 应包含 header:X-Group-ID")
}

// TestClientAttributesValidateNamespaceGroup 验证 NamespaceSources/GroupIDSources 的 Validate
func TestClientAttributesValidateNamespaceGroup(t *testing.T) {
	// 合法配置
	ca := &ClientAttributes{
		NamespaceSources: []common.AttributeSource{
			{Type: common.SourceTypeQuery, Key: "namespace"},
		},
		GroupIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Group-ID"},
		},
	}
	assert.NoError(t, ca.Validate())

	// 非法 Namespace 来源（空 Key）
	invalid := &ClientAttributes{
		NamespaceSources: []common.AttributeSource{
			{Type: common.SourceTypeQuery, Key: ""},
		},
	}
	err := invalid.Validate()
	assert.Error(t, err, "空 Key 的 NamespaceSources 应验证失败")

	// 非法 GroupID 来源（空 Key）
	invalid2 := &ClientAttributes{
		GroupIDSources: []common.AttributeSource{
			{Type: common.SourceTypeQuery, Key: ""},
		},
	}
	err2 := invalid2.Validate()
	assert.Error(t, err2, "空 Key 的 GroupIDSources 应验证失败")
}
