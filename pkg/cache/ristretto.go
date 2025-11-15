/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 00:18:34
 * @FilePath: \go-config\pkg\cache\ristretto.go
 * @Description: Ristretto 高性能缓存配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import "github.com/kamalyes/go-config/internal"

// Ristretto缓存配置
type Ristretto struct {
	ModuleName         string `mapstructure:"module_name" yaml:"module-name" json:"module_name"`                            // 模块名
	NumCounters        int64  `mapstructure:"num_counters" yaml:"num-counters" json:"num_counters" validate:"min=1"`        // 计数器数量，用于跟踪访问频率
	MaxCost            int64  `mapstructure:"max_cost" yaml:"max-cost" json:"max_cost" validate:"min=1"`                    // 最大缓存成本
	BufferItems        int64  `mapstructure:"buffer_items" yaml:"buffer-items" json:"buffer_items" validate:"min=1"`        // Get 缓存的大小
	Metrics            bool   `mapstructure:"metrics" yaml:"metrics" json:"metrics"`                                        // 是否启用缓存统计
	IgnoreInternalCost bool   `mapstructure:"ignore_internal_cost" yaml:"ignore-internal-cost" json:"ignore_internal_cost"` // 是否忽略内部成本计算
	KeyToHash          bool   `mapstructure:"key_to_hash" yaml:"key-to-hash" json:"key_to_hash"`                            // 是否将键转换为哈希
	Cost               int64  `mapstructure:"cost" yaml:"cost" json:"cost" validate:"min=0"`                                // 每个条目的默认成本
}

// NewRistretto 创建一个新的 Ristretto 实例
func NewRistretto(opt *Ristretto) *Ristretto {
	var ristrettoInstance *Ristretto
	internal.LockFunc(func() {
		ristrettoInstance = opt
	})
	return ristrettoInstance
}

// Clone 返回 Ristretto 配置的副本
func (r *Ristretto) Clone() internal.Configurable {
	return &Ristretto{
		ModuleName:         r.ModuleName,
		NumCounters:        r.NumCounters,
		MaxCost:            r.MaxCost,
		BufferItems:        r.BufferItems,
		Metrics:            r.Metrics,
		IgnoreInternalCost: r.IgnoreInternalCost,
		KeyToHash:          r.KeyToHash,
		Cost:               r.Cost,
	}
}

// Get 返回 Ristretto 配置的所有字段
func (r *Ristretto) Get() interface{} {
	return r
}

// Set 更新 Ristretto 配置的字段
func (r *Ristretto) Set(data interface{}) {
	if configData, ok := data.(*Ristretto); ok {
		r.ModuleName = configData.ModuleName
		r.NumCounters = configData.NumCounters
		r.MaxCost = configData.MaxCost
		r.BufferItems = configData.BufferItems
		r.Metrics = configData.Metrics
		r.IgnoreInternalCost = configData.IgnoreInternalCost
		r.KeyToHash = configData.KeyToHash
		r.Cost = configData.Cost
	}
}

// DefaultRistrettoConfig 返回默认 Ristretto 配置
func DefaultRistrettoConfig() Ristretto {
	return Ristretto{
		ModuleName:         "ristretto",
		NumCounters:        1e7,     // 10M
		MaxCost:            1 << 30, // 1GB
		BufferItems:        64,
		Metrics:            true,
		IgnoreInternalCost: false,
		KeyToHash:          true,
		Cost:               1,
	}
}

// Validate 验证 Ristretto 配置
func (r *Ristretto) Validate() error {
	if r.NumCounters <= 0 {
		r.NumCounters = 1e7
	}
	if r.MaxCost <= 0 {
		r.MaxCost = 1 << 30
	}
	if r.BufferItems <= 0 {
		r.BufferItems = 64
	}
	if r.Cost < 0 {
		r.Cost = 1
	}
	return internal.ValidateStruct(r)
}

// GetNumCounters 获取计数器数量
func (r *Ristretto) GetNumCounters() int64 {
	if r.NumCounters <= 0 {
		return 1e7
	}
	return r.NumCounters
}

// GetMaxCost 获取最大成本
func (r *Ristretto) GetMaxCost() int64 {
	if r.MaxCost <= 0 {
		return 1 << 30
	}
	return r.MaxCost
}

// GetBufferItems 获取缓冲项数量
func (r *Ristretto) GetBufferItems() int64 {
	if r.BufferItems <= 0 {
		return 64
	}
	return r.BufferItems
}

// IsMetricsEnabled 检查是否启用指标
func (r *Ristretto) IsMetricsEnabled() bool {
	return r.Metrics
}

// IsIgnoreInternalCost 检查是否忽略内部成本
func (r *Ristretto) IsIgnoreInternalCost() bool {
	return r.IgnoreInternalCost
}

// IsKeyToHash 检查是否将键转换为哈希
func (r *Ristretto) IsKeyToHash() bool {
	return r.KeyToHash
}

// GetCost 获取默认成本
func (r *Ristretto) GetCost() int64 {
	if r.Cost < 0 {
		return 1
	}
	return r.Cost
}
