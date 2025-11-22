/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 16:50:26
 * @FilePath: \go-config\pkg\kafka\kafka_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKafka_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "kafka", config.ModuleName)
	assert.Equal(t, "127.0.0.1:9092", config.Brokers)
	assert.Equal(t, "default-topic", config.Topic)
	assert.Equal(t, "default-group", config.GroupID)
	assert.Equal(t, "kafka_user", config.Username)
	assert.Equal(t, "kafka_password", config.Password)
	assert.Equal(t, 0, config.Partition)
	assert.Equal(t, "latest", config.Offset)
	assert.Equal(t, 3, config.TryTimes)
	assert.Equal(t, "sync-es-topic", config.SyncESTopic)
}

func TestKafka_DefaultKafka(t *testing.T) {
	config := DefaultKafka()
	assert.Equal(t, "kafka", config.ModuleName)
	assert.Equal(t, "127.0.0.1:9092", config.Brokers)
	assert.Equal(t, "default-topic", config.Topic)
	assert.Equal(t, "default-group", config.GroupID)
	assert.Equal(t, 3, config.TryTimes)
}

func TestKafka_WithModuleName(t *testing.T) {
	config := Default().WithModuleName("custom-kafka")
	assert.Equal(t, "custom-kafka", config.ModuleName)
}

func TestKafka_WithBrokers(t *testing.T) {
	config := Default().WithBrokers("kafka1:9092,kafka2:9092")
	assert.Equal(t, "kafka1:9092,kafka2:9092", config.Brokers)
}

func TestKafka_WithTryTimes(t *testing.T) {
	config := Default().WithTryTimes(5)
	assert.Equal(t, 5, config.TryTimes)
}

func TestKafka_WithSyncESTopic(t *testing.T) {
	config := Default().WithSyncESTopic("custom-sync-topic")
	assert.Equal(t, "custom-sync-topic", config.SyncESTopic)
}

func TestKafka_WithTopic(t *testing.T) {
	config := Default().WithTopic("test-topic")
	assert.Equal(t, "test-topic", config.Topic)
}

func TestKafka_WithGroupID(t *testing.T) {
	config := Default().WithGroupID("test-group")
	assert.Equal(t, "test-group", config.GroupID)
}

func TestKafka_WithUsername(t *testing.T) {
	config := Default().WithUsername("testuser")
	assert.Equal(t, "testuser", config.Username)
}

func TestKafka_WithPassword(t *testing.T) {
	config := Default().WithPassword("testpass")
	assert.Equal(t, "testpass", config.Password)
}

func TestKafka_WithPartition(t *testing.T) {
	config := Default().WithPartition(5)
	assert.Equal(t, 5, config.Partition)
}

func TestKafka_WithOffset(t *testing.T) {
	config := Default().WithOffset("earliest")
	assert.Equal(t, "earliest", config.Offset)
}

func TestKafka_Clone(t *testing.T) {
	original := Default().
		WithModuleName("test-kafka").
		WithBrokers("broker1:9092,broker2:9092").
		WithTopic("test-topic").
		WithGroupID("test-group").
		WithUsername("user").
		WithPassword("pass").
		WithPartition(3).
		WithOffset("earliest").
		WithTryTimes(5).
		WithSyncESTopic("sync-topic")

	cloned := original.Clone().(*Kafka)

	// 验证值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Brokers, cloned.Brokers)
	assert.Equal(t, original.Topic, cloned.Topic)
	assert.Equal(t, original.GroupID, cloned.GroupID)
	assert.Equal(t, original.Username, cloned.Username)
	assert.Equal(t, original.Password, cloned.Password)
	assert.Equal(t, original.Partition, cloned.Partition)
	assert.Equal(t, original.Offset, cloned.Offset)
	assert.Equal(t, original.TryTimes, cloned.TryTimes)
	assert.Equal(t, original.SyncESTopic, cloned.SyncESTopic)

	// 验证独立性
	cloned.ModuleName = "modified-kafka"
	cloned.Topic = "modified-topic"
	assert.NotEqual(t, original.ModuleName, cloned.ModuleName)
	assert.NotEqual(t, original.Topic, cloned.Topic)
}

func TestKafka_Get(t *testing.T) {
	config := Default().WithTopic("test-topic")
	got := config.Get()
	assert.NotNil(t, got)
	kafkaConfig, ok := got.(*Kafka)
	assert.True(t, ok)
	assert.Equal(t, "test-topic", kafkaConfig.Topic)
}

func TestKafka_Set(t *testing.T) {
	config := Default()
	newConfig := &Kafka{
		ModuleName:  "new-kafka",
		Brokers:     "new-broker:9092",
		Topic:       "new-topic",
		GroupID:     "new-group",
		Username:    "newuser",
		Password:    "newpass",
		Partition:   2,
		Offset:      "earliest",
		TryTimes:    10,
		SyncESTopic: "new-sync-topic",
	}

	config.Set(newConfig)
	assert.Equal(t, "new-kafka", config.ModuleName)
	assert.Equal(t, "new-broker:9092", config.Brokers)
	assert.Equal(t, "new-topic", config.Topic)
	assert.Equal(t, "new-group", config.GroupID)
	assert.Equal(t, "newuser", config.Username)
	assert.Equal(t, "newpass", config.Password)
	assert.Equal(t, 2, config.Partition)
	assert.Equal(t, "earliest", config.Offset)
	assert.Equal(t, 10, config.TryTimes)
	assert.Equal(t, "new-sync-topic", config.SyncESTopic)
}

func TestKafka_Validate(t *testing.T) {
	// 有效配置
	validConfig := Default()
	err := validConfig.Validate()
	assert.NoError(t, err)

	// 无效配置 - 缺少必需字段
	invalidConfig := &Kafka{
		ModuleName: "test",
		Brokers:    "", // required field
		Topic:      "", // required field
	}
	err = invalidConfig.Validate()
	assert.Error(t, err)

	// 无效配置 - TryTimes小于1
	invalidConfig2 := &Kafka{
		ModuleName:  "test",
		Brokers:     "broker:9092",
		Topic:       "topic",
		TryTimes:    0, // must be >= 1
		SyncESTopic: "sync-topic",
	}
	err = invalidConfig2.Validate()
	assert.Error(t, err)
}

func TestKafka_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("chain-kafka").
		WithBrokers("chain-broker:9092").
		WithTopic("chain-topic").
		WithGroupID("chain-group").
		WithUsername("chainuser").
		WithPassword("chainpass").
		WithPartition(10).
		WithOffset("earliest").
		WithTryTimes(7).
		WithSyncESTopic("chain-sync-topic")

	assert.Equal(t, "chain-kafka", config.ModuleName)
	assert.Equal(t, "chain-broker:9092", config.Brokers)
	assert.Equal(t, "chain-topic", config.Topic)
	assert.Equal(t, "chain-group", config.GroupID)
	assert.Equal(t, "chainuser", config.Username)
	assert.Equal(t, "chainpass", config.Password)
	assert.Equal(t, 10, config.Partition)
	assert.Equal(t, "earliest", config.Offset)
	assert.Equal(t, 7, config.TryTimes)
	assert.Equal(t, "chain-sync-topic", config.SyncESTopic)
}

func TestNewKafka(t *testing.T) {
	opt := &Kafka{
		ModuleName:  "new-instance",
		Brokers:     "new-broker:9092",
		Topic:       "new-topic",
		GroupID:     "new-group",
		Username:    "newuser",
		Password:    "newpass",
		Partition:   1,
		Offset:      "latest",
		TryTimes:    5,
		SyncESTopic: "new-sync-topic",
	}

	instance := NewKafka(opt)
	assert.NotNil(t, instance)
	assert.Equal(t, "new-instance", instance.ModuleName)
	assert.Equal(t, "new-broker:9092", instance.Brokers)
	assert.Equal(t, "new-topic", instance.Topic)
	assert.Equal(t, "new-group", instance.GroupID)
	assert.Equal(t, "newuser", instance.Username)
	assert.Equal(t, "newpass", instance.Password)
	assert.Equal(t, 1, instance.Partition)
	assert.Equal(t, "latest", instance.Offset)
	assert.Equal(t, 5, instance.TryTimes)
	assert.Equal(t, "new-sync-topic", instance.SyncESTopic)
}
