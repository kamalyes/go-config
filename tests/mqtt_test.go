/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 11:25:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:25:55
 * @FilePath: \go-config\tests\mqtt_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/queue"
	"github.com/stretchr/testify/assert"
)

// 公共测试数据
var (
	validMqttData   = queue.NewMqtt("test-module", "tcp://127.0.0.1:1883", "user", "pass", 4, true, true, 60, 30, 10, 5, 5, "will/topic")
	invalidMqttData = queue.NewMqtt("", "", "", "", 5, true, true, 0, -1, -1, -1, -1, "")
	mqttMapData     = map[string]interface{}{
		"moduleName":           "test-module",
		"url":                  "tcp://127.0.0.1:1883",
		"username":             "user",
		"password":             "pass",
		"protocolVersion":      uint(4),
		"cleanSession":         true,
		"autoReconnect":        true,
		"keepAlive":            60,
		"maxReconnectInterval": 30,
		"pingTimeout":          10,
		"writeTimeout":         5,
		"connectTimeout":       5,
		"willTopic":            "will/topic",
	}
)

func TestNewMqtt(t *testing.T) {
	mqtt := validMqttData

	assert.NotNil(t, mqtt)
	assert.Equal(t, "test-module", mqtt.ModuleName)
	assert.Equal(t, "tcp://127.0.0.1:1883", mqtt.Url)
}

func TestMqttValidate(t *testing.T) {
	assert.NoError(t, validMqttData.Validate())
	assert.Error(t, invalidMqttData.Validate())
}

func TestMqttToMap(t *testing.T) {
	mqtt := validMqttData
	mqttMap := mqtt.ToMap()

	assert.Equal(t, "test-module", mqttMap["moduleName"])
	assert.Equal(t, "tcp://127.0.0.1:1883", mqttMap["url"])
}

func TestMqttFromMap(t *testing.T) {
	mqtt := &queue.Mqtt{}
	mqtt.FromMap(mqttMapData)

	assert.Equal(t, "test-module", mqtt.ModuleName)
	assert.Equal(t, "tcp://127.0.0.1:1883", mqtt.Url)
}

func TestMqttClone(t *testing.T) {
	mqtt := validMqttData
	clone := mqtt.Clone()

	assert.Equal(t, mqtt, clone)
	assert.NotSame(t, mqtt, clone) // 确保是不同的实例
}

func TestMqttSet(t *testing.T) {
	mqtt := queue.NewMqtt("old-module", "tcp://127.0.0.1:1883", "user", "pass", 4, true, true, 60, 30, 10, 5, 5, "will/topic")
	newData := queue.NewMqtt("new-module", "tcp://127.0.0.1:1884", "new-user", "new-pass", 3, false, false, 120, 60, 20, 10, 10, "new/will/topic")
	mqtt.Set(newData)

	assert.Equal(t, "new-module", mqtt.ModuleName)
	assert.Equal(t, "tcp://127.0.0.1:1884", mqtt.Url)
	assert.Equal(t, "new-user", mqtt.Username)
	assert.Equal(t, "new-pass", mqtt.Password)
}
