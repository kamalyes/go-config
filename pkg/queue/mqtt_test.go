/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 11:30:00
 * @FilePath: \go-config\pkg\queue\mqtt_test.go
 * @Description: MQTT配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMqtt_Clone(t *testing.T) {
	original := &Mqtt{
		ModuleName:           "test-mqtt",
		Endpoint:             "tcp://localhost:1883",
		ClientID:             "test-client",
		ProtocolVersion:      4,
		KeepAlive:            60,
		MaxReconnectInterval: 300,
		PingTimeout:          10,
		WriteTimeout:         10,
		ConnectTimeout:       30,
		Username:             "user",
		Password:             "pass",
		CleanSession:         true,
		AutoReconnect:        true,
		WillTopic:            "/test/will",
	}

	cloned := original.Clone().(*Mqtt)

	// 验证克隆后的值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Endpoint, cloned.Endpoint)
	assert.Equal(t, original.ClientID, cloned.ClientID)
	assert.Equal(t, original.ProtocolVersion, cloned.ProtocolVersion)
	assert.Equal(t, original.KeepAlive, cloned.KeepAlive)
	assert.Equal(t, original.Username, cloned.Username)
	assert.Equal(t, original.Password, cloned.Password)
	assert.Equal(t, original.CleanSession, cloned.CleanSession)
	assert.Equal(t, original.AutoReconnect, cloned.AutoReconnect)
	assert.Equal(t, original.WillTopic, cloned.WillTopic)

	// 修改原始对象不应影响克隆对象
	original.ClientID = "new-client"
	original.KeepAlive = 120
	assert.NotEqual(t, original.ClientID, cloned.ClientID)
	assert.NotEqual(t, original.KeepAlive, cloned.KeepAlive)
}
