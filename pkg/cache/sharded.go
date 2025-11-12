/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 11:56:24
 * @FilePath: \go-config\pkg\cache\sharded.go
 * @Description: 分片缓存配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import "github.com/kamalyes/go-config/internal"

// Sharded 分片缓存配置
type Sharded struct {
	ModuleName   string    `mapstructure:"module_name" yaml:"module_name" json:"module_name"`                  // 模块名
	ShardCount   int       `mapstructure:"shard_count" yaml:"shard_count" json:"shard_count" validate:"min=1"` // 分片数量
	BaseType     CacheType `mapstructure:"base_type" yaml:"base_type" json:"base_type"`                        // 基础缓存类型
	HashFunc     string    `mapstructure:"hash_func" yaml:"hash_func" json:"hash_func"`                        // 哈希函数: fnv, crc32, md5
	LoadBalancer string    `mapstructure:"load_balancer" yaml:"load_balancer" json:"load_balancer"`            // 负载均衡策略: consistent_hash, round_robin
}

// NewSharded 创建一个新的 Sharded 实例
func NewSharded(opt *Sharded) *Sharded {
	var shardedInstance *Sharded
	internal.LockFunc(func() {
		shardedInstance = opt
	})
	return shardedInstance
}

// Clone 返回 Sharded 配置的副本
func (s *Sharded) Clone() internal.Configurable {
	return &Sharded{
		ModuleName:   s.ModuleName,
		ShardCount:   s.ShardCount,
		BaseType:     s.BaseType,
		HashFunc:     s.HashFunc,
		LoadBalancer: s.LoadBalancer,
	}
}

// Get 返回 Sharded 配置的所有字段
func (s *Sharded) Get() interface{} {
	return s
}

// Set 更新 Sharded 配置的字段
func (s *Sharded) Set(data interface{}) {
	if configData, ok := data.(*Sharded); ok {
		s.ModuleName = configData.ModuleName
		s.ShardCount = configData.ShardCount
		s.BaseType = configData.BaseType
		s.HashFunc = configData.HashFunc
		s.LoadBalancer = configData.LoadBalancer
	}
}

// DefaultShardedConfig 返回默认分片缓存配置
func DefaultShardedConfig() Sharded {
	return Sharded{
		ModuleName:   "sharded",
		ShardCount:   32,
		BaseType:     CacheTypeMemory,
		HashFunc:     "fnv",
		LoadBalancer: "consistent_hash",
	}
}

// Validate 验证分片缓存配置
func (s *Sharded) Validate() error {
	if s.ShardCount <= 0 {
		s.ShardCount = 32
	}
	if s.BaseType == "" {
		s.BaseType = CacheTypeMemory
	}
	if s.HashFunc == "" {
		s.HashFunc = "fnv"
	}
	if s.LoadBalancer == "" {
		s.LoadBalancer = "consistent_hash"
	}
	return internal.ValidateStruct(s)
}
