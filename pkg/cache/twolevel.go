/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 00:00:00
 * @FilePath: \go-config\pkg\cache\twolevel.go
 * @Description: 二级缓存配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import (
	"time"

	"github.com/kamalyes/go-config/internal"
)

// TwoLevel 二级缓存配置
type TwoLevel struct {
	ModuleName       string        `mapstructure:"module_name" yaml:"module_name" json:"module_name"`                           // 模块名
	L1Type           CacheType     `mapstructure:"l1_type" yaml:"l1_type" json:"l1_type"`                               // L1缓存类型
	L2Type           CacheType     `mapstructure:"l2_type" yaml:"l2_type" json:"l2_type"`                               // L2缓存类型
	L1TTL            time.Duration `mapstructure:"l1_ttl" yaml:"l1_ttl" json:"l1_ttl"`                                  // L1缓存TTL
	L2TTL            time.Duration `mapstructure:"l2_ttl" yaml:"l2_ttl" json:"l2_ttl"`                                  // L2缓存TTL
	SyncStrategy     string        `mapstructure:"sync_strategy" yaml:"sync_strategy" json:"sync_strategy"`             // 同步策略: write_through, write_back, write_around
	L1Size           int           `mapstructure:"l1_size" yaml:"l1_size" json:"l1_size"`                               // L1缓存大小
	L2Size           int           `mapstructure:"l2_size" yaml:"l2_size" json:"l2_size"`                               // L2缓存大小
	PromoteThreshold int           `mapstructure:"promote_threshold" yaml:"promote_threshold" json:"promote_threshold"` // 提升阈值
}

// NewTwoLevel 创建一个新的 TwoLevel 实例
func NewTwoLevel(opt *TwoLevel) *TwoLevel {
	var twoLevelInstance *TwoLevel
	internal.LockFunc(func() {
		twoLevelInstance = opt
	})
	return twoLevelInstance
}

// Clone 返回 TwoLevel 配置的副本
func (t *TwoLevel) Clone() internal.Configurable {
	return &TwoLevel{
		ModuleName:       t.ModuleName,
		L1Type:           t.L1Type,
		L2Type:           t.L2Type,
		L1TTL:            t.L1TTL,
		L2TTL:            t.L2TTL,
		SyncStrategy:     t.SyncStrategy,
		L1Size:           t.L1Size,
		L2Size:           t.L2Size,
		PromoteThreshold: t.PromoteThreshold,
	}
}

// Get 返回 TwoLevel 配置的所有字段
func (t *TwoLevel) Get() interface{} {
	return t
}

// Set 更新 TwoLevel 配置的字段
func (t *TwoLevel) Set(data interface{}) {
	if configData, ok := data.(*TwoLevel); ok {
		t.ModuleName = configData.ModuleName
		t.L1Type = configData.L1Type
		t.L2Type = configData.L2Type
		t.L1TTL = configData.L1TTL
		t.L2TTL = configData.L2TTL
		t.SyncStrategy = configData.SyncStrategy
		t.L1Size = configData.L1Size
		t.L2Size = configData.L2Size
		t.PromoteThreshold = configData.PromoteThreshold
	}
}

// DefaultTwoLevelConfig 返回默认二级缓存配置
func DefaultTwoLevelConfig() TwoLevel {
	return TwoLevel{
		ModuleName:       "twolevel",
		L1Type:           TypeMemory,
		L2Type:           TypeRedis,
		L1TTL:            5 * time.Minute,
		L2TTL:            30 * time.Minute,
		SyncStrategy:     "write_through",
		L1Size:           1000,
		L2Size:           10000,
		PromoteThreshold: 2,
	}
}

// Validate 验证二级缓存配置
func (t *TwoLevel) Validate() error {
	if t.L1Type == "" {
		t.L1Type = TypeMemory
	}
	if t.L2Type == "" {
		t.L2Type = TypeRedis
	}
	if t.L1TTL <= 0 {
		t.L1TTL = 5 * time.Minute
	}
	if t.L2TTL <= 0 {
		t.L2TTL = 30 * time.Minute
	}
	if t.SyncStrategy == "" {
		t.SyncStrategy = "write_through"
	}
	if t.L1Size <= 0 {
		t.L1Size = 1000
	}
	if t.L2Size <= 0 {
		t.L2Size = 10000
	}
	if t.PromoteThreshold <= 0 {
		t.PromoteThreshold = 2
	}
	return internal.ValidateStruct(t)
}