/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 00:00:00
 * @FilePath: \go-config\pkg\cache\expiring.go
 * @Description: 过期缓存配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import (
	"github.com/kamalyes/go-config/internal"
	"time"
)

// Expiring 过期缓存配置
type Expiring struct {
	ModuleName       string        `mapstructure:"module_name" yaml:"module-name" json:"module_name"`                      // 模块名
	CleanupInterval  time.Duration `mapstructure:"cleanup_interval" yaml:"cleanup-interval" json:"cleanup_interval"`       // 清理间隔
	DefaultTTL       time.Duration `mapstructure:"default_ttl" yaml:"default-ttl" json:"default_ttl"`                      // 默认TTL
	MaxSize          int           `mapstructure:"max_size" yaml:"max-size" json:"max_size"`                               // 最大大小
	EvictionPolicy   string        `mapstructure:"eviction_policy" yaml:"eviction-policy" json:"eviction_policy"`          // 驱逐策略: lru, lfu, fifo
	EnableLazyExpiry bool          `mapstructure:"enable_lazy_expiry" yaml:"enable-lazy-expiry" json:"enable_lazy_expiry"` // 启用懒惰过期
	MaxMemoryUsage   int64         `mapstructure:"max_memory_usage" yaml:"max-memory-usage" json:"max_memory_usage"`       // 最大内存使用量(字节)
}

// NewExpiring 创建一个新的 Expiring 实例
func NewExpiring(opt *Expiring) *Expiring {
	var expiringInstance *Expiring
	internal.LockFunc(func() {
		expiringInstance = opt
	})
	return expiringInstance
}

// Clone 返回 Expiring 配置的副本
func (e *Expiring) Clone() internal.Configurable {
	return &Expiring{
		ModuleName:       e.ModuleName,
		CleanupInterval:  e.CleanupInterval,
		DefaultTTL:       e.DefaultTTL,
		MaxSize:          e.MaxSize,
		EvictionPolicy:   e.EvictionPolicy,
		EnableLazyExpiry: e.EnableLazyExpiry,
		MaxMemoryUsage:   e.MaxMemoryUsage,
	}
}

// Get 返回 Expiring 配置的所有字段
func (e *Expiring) Get() interface{} {
	return e
}

// Set 更新 Expiring 配置的字段
func (e *Expiring) Set(data interface{}) {
	if configData, ok := data.(*Expiring); ok {
		e.ModuleName = configData.ModuleName
		e.CleanupInterval = configData.CleanupInterval
		e.DefaultTTL = configData.DefaultTTL
		e.MaxSize = configData.MaxSize
		e.EvictionPolicy = configData.EvictionPolicy
		e.EnableLazyExpiry = configData.EnableLazyExpiry
		e.MaxMemoryUsage = configData.MaxMemoryUsage
	}
}

// DefaultExpiringConfig 返回默认过期缓存配置
func DefaultExpiringConfig() Expiring {
	return Expiring{
		ModuleName:       "expiring",
		CleanupInterval:  5 * time.Minute,
		DefaultTTL:       30 * time.Minute,
		MaxSize:          10000,
		EvictionPolicy:   "lru",
		EnableLazyExpiry: true,
		MaxMemoryUsage:   100 * 1024 * 1024, // 100MB
	}
}

// DefaultExpiring 返回默认过期缓存配置的指针，支持链式调用
func DefaultExpiring() *Expiring {
	config := DefaultExpiringConfig()
	return &config
}

// Validate 验证过期缓存配置
func (e *Expiring) Validate() error {
	if e.CleanupInterval <= 0 {
		e.CleanupInterval = 5 * time.Minute
	}
	if e.DefaultTTL <= 0 {
		e.DefaultTTL = 30 * time.Minute
	}
	if e.MaxSize <= 0 {
		e.MaxSize = 10000
	}
	if e.EvictionPolicy == "" {
		e.EvictionPolicy = "lru"
	}
	if e.MaxMemoryUsage <= 0 {
		e.MaxMemoryUsage = 100 * 1024 * 1024 // 100MB
	}
	return internal.ValidateStruct(e)
}
