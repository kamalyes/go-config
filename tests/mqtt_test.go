/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 11:25:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:56:01
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
		Url:                  fmt.Sprintf("tcp://%s:%d", random.RandString(5, random.CAPITAL), random.FRandInt(1883, 8883)), // 随机生成 URL
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
	assert.Equal(t, newParams.Url, mqttInstance.Url)
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
