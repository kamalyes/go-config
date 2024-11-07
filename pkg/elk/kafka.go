/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:16:55
 * @FilePath: \go-config\pkg\elk\kafka.go
 * @Description:
 *
 * Copyright (j) 2024 by kamalyes, All Rights Reserved.
 */
package elk

import "github.com/kamalyes/go-config/internal"

// Kafka 结构体
type Kafka struct {
	Brokers     string `mapstructure:"brokers"        yaml:"brokers"        json:"brokers"           validate:"required"`       // Kafka brokers
	TryTimes    int    `mapstructure:"try-times"      yaml:"try-times"      json:"try_times"         validate:"required,min=1"` // 尝试次数
	SyncESTopic string `mapstructure:"sync-es-topic"  yaml:"sync-es-topic"  json:"sync_es_topic"     validate:"required"`       // 同步到 ES 的主题
	ModuleName  string `mapstructure:"modulename"     yaml:"modulename"     json:"module_name"`                                 // 模块名称
}

// NewKafka 创建一个新的 Kafka 实例
func NewKafka(opt *Kafka) *Kafka {
	var kafkaInstance *Kafka

	internal.LockFunc(func() {
		kafkaInstance = opt
	})
	return kafkaInstance
}

// Clone 返回 Kafka 配置的副本
func (k *Kafka) Clone() internal.Configurable {
	return &Kafka{
		Brokers:     k.Brokers,
		TryTimes:    k.TryTimes,
		SyncESTopic: k.SyncESTopic,
		ModuleName:  k.ModuleName,
	}
}

// Get 返回 Kafka 配置的所有字段
func (k *Kafka) Get() interface{} {
	return k
}

// Set 更新 Kafka 配置的字段
func (k *Kafka) Set(data interface{}) {
	if configData, ok := data.(*Kafka); ok {
		k.Brokers = configData.Brokers
		k.TryTimes = configData.TryTimes
		k.SyncESTopic = configData.SyncESTopic
		k.ModuleName = configData.ModuleName
	}
}

// Validate 检查 Kafka 配置的有效性
func (k *Kafka) Validate() error {
	return internal.ValidateStruct(k)
}
