/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-10-11 21:36:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-10-11 21:36:00
 * @FilePath: \go-config\pkg\wsc\message_broker.go
 * @Description: 消息总线配置 —— MQ 入口与内部传输的统一投递通道
 *
 * 一刀切 MQ 路线（PUBSUB_REFACTOR v3）：所有生产者统一 publish，
 * 投递等级分流到 NATS 通道（express→core subject / 可靠档→JetStream 分片 /
 * 广播→core fanout）。本配置只承载投递总线编排参数（driver/topic/group/shard/consumer）
 *
 * Driver 取值：
 *   - nats   生产选型（core + JetStream 分级）
 *   - memory 契约测试与本地开发（进程内闭环，无外部依赖）
 *   - 空     不启用，broker 域零初始化（直调模式）
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package wsc

// MessageBroker 消息总线驱动类型
const (
	BrokerDriverNats   = "nats"   // NATS（core + JetStream 分级，生产选型）
	BrokerDriverMemory = "memory" // 进程内实现（契约测试与本地开发）
)

// BrokerConsumerParams 消费端批量拉取参数
type BrokerConsumerParams struct {
	BatchSize     int `mapstructure:"batch-size" yaml:"batch-size" json:"batchSize"`               // 单次批量拉取上限（摊薄 RTT）
	MaxAckPending int `mapstructure:"max-ack-pending" yaml:"max-ack-pending" json:"maxAckPending"` // 在途未 ack 上限（背压闸门，超限暂停拉取）
}

// MessageBrokerConfig 消息总线配置（投递指令的发布/订阅通道编排参数）
//
// NATS 连接配置不在本结构：宿主从 queue.Nats 模块取用，装配期传给
// natsadapter.BuildConn 建连后注入适配器（NewEventBus(conn)）
type MessageBrokerConfig struct {
	Enabled     bool                  `mapstructure:"enabled" yaml:"enabled" json:"enabled"`               // 总开关（false = 直调模式，broker 域零初始化零开销）
	Driver      string                `mapstructure:"driver" yaml:"driver" json:"driver"`                  // 驱动：nats | memory（空 = 不启用）
	TopicPrefix string                `mapstructure:"topic-prefix" yaml:"topic-prefix" json:"topicPrefix"` // topic 前缀（wsc.{app}.deliver / .broadcast / .x）
	Group       string                `mapstructure:"group" yaml:"group" json:"group"`                     // consumer group（广播语义自动按 nodeID 派生独立 group）
	ShardCount  int                   `mapstructure:"shard-count" yaml:"shard-count" json:"shardCount"`    // JetStream deliver 分片数（顺序并行度，分片内串行保序）
	Consumer    *BrokerConsumerParams `mapstructure:"consumer" yaml:"consumer" json:"consumer"`            // 消费端批量拉取参数
}

// DefaultMessageBrokerConfig 返回默认消息总线配置（禁用态，直调模式）
func DefaultMessageBrokerConfig() *MessageBrokerConfig {
	return &MessageBrokerConfig{
		Enabled:     false,
		Driver:      BrokerDriverNats,
		TopicPrefix: "wsc",
		Group:       "wsc-deliver",
		ShardCount:  64,
		Consumer: &BrokerConsumerParams{
			BatchSize:     256,
			MaxAckPending: 1024,
		},
	}
}

// IsEnabled 消息总线启用判定（nil 安全）
func (m *MessageBrokerConfig) IsEnabled() bool {
	return m != nil && m.Enabled && m.Driver != ""
}

// IsNats 是否 NATS 驱动（装配侧据此取用 Nats 连接配置）
func (m *MessageBrokerConfig) IsNats() bool {
	return m != nil && m.Driver == BrokerDriverNats
}

// ResolveShardCount 分片数兜底（0 = 默认 64）
func (m *MessageBrokerConfig) ResolveShardCount() int {
	if m == nil || m.ShardCount <= 0 {
		return 64
	}
	return m.ShardCount
}

// ResolveConsumer 消费参数兜底（nil 或零值 = 默认）
func (m *MessageBrokerConfig) ResolveConsumer() *BrokerConsumerParams {
	if m == nil || m.Consumer == nil {
		return &BrokerConsumerParams{BatchSize: 256, MaxAckPending: 1024}
	}
	if m.Consumer.BatchSize <= 0 {
		m.Consumer.BatchSize = 256
	}
	if m.Consumer.MaxAckPending <= 0 {
		m.Consumer.MaxAckPending = 1024
	}
	return m.Consumer
}

// WithEnabled 设置总开关
func (m *MessageBrokerConfig) WithEnabled(enabled bool) *MessageBrokerConfig {
	m.Enabled = enabled
	return m
}

// WithDriver 设置驱动类型（nats | memory）
func (m *MessageBrokerConfig) WithDriver(driver string) *MessageBrokerConfig {
	m.Driver = driver
	return m
}

// WithTopicPrefix 设置 topic 前缀
func (m *MessageBrokerConfig) WithTopicPrefix(prefix string) *MessageBrokerConfig {
	m.TopicPrefix = prefix
	return m
}

// WithGroup 设置 consumer group
func (m *MessageBrokerConfig) WithGroup(group string) *MessageBrokerConfig {
	m.Group = group
	return m
}

// WithShardCount 设置 JetStream deliver 分片数
func (m *MessageBrokerConfig) WithShardCount(count int) *MessageBrokerConfig {
	m.ShardCount = count
	return m
}

// WithConsumer 设置消费端批量拉取参数
func (m *MessageBrokerConfig) WithConsumer(params *BrokerConsumerParams) *MessageBrokerConfig {
	m.Consumer = params
	return m
}
