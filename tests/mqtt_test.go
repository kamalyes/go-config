/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 11:25:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 02:04:15
 * @FilePath: \go-config\tests\mqtt_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/queue"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Mqtt 配置参数
func generateMqttTestParams() *queue.Mqtt {
	return &queue.Mqtt{
		ModuleName:           random.RandString(10, random.CAPITAL),
		Endpoint:             fmt.Sprintf("tcp://%s:%d", random.RandString(5, random.CAPITAL), random.FRandInt(1883, 8883)), // 随机生成 URL
		Username:             random.RandString(8, random.CAPITAL),                                                          // 随机生成用户名
		Password:             random.RandString(16, random.CAPITAL),                                                         // 随机生成密码
		ProtocolVersion:      uint(random.FRandInt(3, 5)),                                                                   // 随机生成协议版本（3 或 4）
		CleanSession:         random.FRandBool(),                                                                            // 随机生成是否清除会话
		AutoReconnect:        random.FRandBool(),                                                                            // 随机生成是否自动重连
		KeepAlive:            random.FRandInt(10, 120),                                                                      // 随机生成保持连接时间（10到120秒）
		MaxReconnectInterval: random.FRandInt(1, 30),                                                                        // 随机生成最大重连间隔（1到30秒）
		PingTimeout:          random.FRandInt(1, 10),                                                                        // 随机生成 ping 超时（1到10秒）
		WriteTimeout:         random.FRandInt(1, 10),                                                                        // 随机生成写超时（1到10秒）
		ConnectTimeout:       random.FRandInt(1, 10),                                                                        // 随机生成连接超时（1到10秒）
		WillTopic:            random.RandString(10, random.CAPITAL),                                                         // 随机生成遗言发送的 topic
	}
}

func TestMqttClone(t *testing.T) {
	params := generateMqttTestParams()
	original := queue.NewMqtt(params)
	cloned := original.Clone().(*queue.Mqtt)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestMqttSet(t *testing.T) {
	oldParams := generateMqttTestParams()
	newParams := generateMqttTestParams()

	mqttInstance := queue.NewMqtt(oldParams)
	newConfig := queue.NewMqtt(newParams)

	mqttInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, mqttInstance.ModuleName)
	assert.Equal(t, newParams.Endpoint, mqttInstance.Endpoint)
	assert.Equal(t, newParams.Username, mqttInstance.Username)
	assert.Equal(t, newParams.Password, mqttInstance.Password)
	assert.Equal(t, newParams.ProtocolVersion, mqttInstance.ProtocolVersion)
	assert.Equal(t, newParams.CleanSession, mqttInstance.CleanSession)
	assert.Equal(t, newParams.AutoReconnect, mqttInstance.AutoReconnect)
	assert.Equal(t, newParams.KeepAlive, mqttInstance.KeepAlive)
	assert.Equal(t, newParams.MaxReconnectInterval, mqttInstance.MaxReconnectInterval)
	assert.Equal(t, newParams.PingTimeout, mqttInstance.PingTimeout)
	assert.Equal(t, newParams.WriteTimeout, mqttInstance.WriteTimeout)
	assert.Equal(t, newParams.ConnectTimeout, mqttInstance.ConnectTimeout)
	assert.Equal(t, newParams.WillTopic, mqttInstance.WillTopic)
}

// TestMqttDefault 测试默认配置
func TestMqttDefault(t *testing.T) {
	defaultConfig := queue.DefaultMqtt()
	
	// 检查默认值
	assert.Equal(t, "mqtt", defaultConfig.ModuleName)
	assert.Equal(t, "tcp://127.0.0.1:1883", defaultConfig.Endpoint)
	assert.Equal(t, uint(4), defaultConfig.ProtocolVersion)
	assert.Equal(t, 60, defaultConfig.KeepAlive)
	assert.Equal(t, 300, defaultConfig.MaxReconnectInterval)
	assert.Equal(t, 10, defaultConfig.PingTimeout)
	assert.Equal(t, 10, defaultConfig.WriteTimeout)
	assert.Equal(t, 30, defaultConfig.ConnectTimeout)
	assert.Equal(t, "", defaultConfig.Username)
	assert.Equal(t, "", defaultConfig.Password)
	assert.Equal(t, true, defaultConfig.CleanSession)
	assert.Equal(t, true, defaultConfig.AutoReconnect)
	assert.Equal(t, "", defaultConfig.WillTopic)
}

// TestMqttDefaultPointer 测试默认配置指针
func TestMqttDefaultPointer(t *testing.T) {
	config := queue.Default()
	
	assert.NotNil(t, config)
	assert.Equal(t, "mqtt", config.ModuleName)
	assert.Equal(t, "tcp://127.0.0.1:1883", config.Endpoint)
}

// TestMqttChainMethods 测试链式方法
func TestMqttChainMethods(t *testing.T) {
	config := queue.Default().
		WithModuleName("message-queue").
		WithEndpoint("tcp://192.168.1.100:1883").
		WithUsername("testuser").
		WithPassword("testpass").
		WithProtocolVersion(3).
		WithKeepAlive(120).
		WithMaxReconnectInterval(600).
		WithPingTimeout(15).
		WithWriteTimeout(20).
		WithConnectTimeout(45).
		WithCleanSession(false).
		WithAutoReconnect(false).
		WithWillTopic("will/topic")
	
	assert.Equal(t, "message-queue", config.ModuleName)
	assert.Equal(t, "tcp://192.168.1.100:1883", config.Endpoint)
	assert.Equal(t, "testuser", config.Username)
	assert.Equal(t, "testpass", config.Password)
	assert.Equal(t, uint(3), config.ProtocolVersion)
	assert.Equal(t, 120, config.KeepAlive)
	assert.Equal(t, 600, config.MaxReconnectInterval)
	assert.Equal(t, 15, config.PingTimeout)
	assert.Equal(t, 20, config.WriteTimeout)
	assert.Equal(t, 45, config.ConnectTimeout)
	assert.Equal(t, false, config.CleanSession)
	assert.Equal(t, false, config.AutoReconnect)
	assert.Equal(t, "will/topic", config.WillTopic)
}

// TestMqttChainMethodsReturnPointer 测试链式方法返回指针
func TestMqttChainMethodsReturnPointer(t *testing.T) {
	config1 := queue.Default()
	config2 := config1.WithEndpoint("tcp://localhost:1883")
	
	// 应该返回同一个实例
	assert.Same(t, config1, config2)
		assert.Equal(t, "tcp://localhost:1883", config1.Endpoint)
}
