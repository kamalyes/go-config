/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:06:09
 * @FilePath: \go-config\pkg\cache\cache.go
 * @Description: 缓存总配置定义
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import (
	"time"

	"github.com/kamalyes/go-config/internal"
)

// CacheType 缓存类型
type CacheType string

const (
	CacheTypeMemory    CacheType = "memory"    // 内存缓存(LRU)
	CacheTypeRistretto CacheType = "ristretto" // 高性能缓存
	CacheTypeRedis     CacheType = "redis"     // Redis缓存
	CacheTypeSharded   CacheType = "sharded"   // 分片缓存
	CacheTypeTwoLevel  CacheType = "two-level" // 二级缓存
	CacheTypeExpiring  CacheType = "expiring"  // 过期缓存
)

// Cache 总缓存配置
type Cache struct {
	ModuleName string    `mapstructure:"module_name" yaml:"module-name" json:"module_name"`
	Type       CacheType `mapstructure:"type" yaml:"type" json:"type"`
	Enabled    bool      `mapstructure:"enabled" yaml:"enabled" json:"enabled"`

	// 基础配置
	DefaultTTL time.Duration `mapstructure:"default_ttl" yaml:"default-ttl" json:"default_ttl"`
	KeyPrefix  string        `mapstructure:"key_prefix" yaml:"key-prefix" json:"key_prefix"`
	Serializer string        `mapstructure:"serializer" yaml:"serializer" json:"serializer"` // json, gob, msgpack

	// 各种缓存实现的配置（按需使用）
	Memory    Memory    `mapstructure:"memory" yaml:"memory" json:"memory"`
	Ristretto Ristretto `mapstructure:"ristretto" yaml:"ristretto" json:"ristretto"`
	Redis     Redis     `mapstructure:"redis" yaml:"redis" json:"redis"`
	Sharded   Sharded   `mapstructure:"sharded" yaml:"sharded" json:"sharded"`
	TwoLevel  TwoLevel  `mapstructure:"two_level" yaml:"two-level" json:"two_level"`
	Expiring  Expiring  `mapstructure:"expiring" yaml:"expiring" json:"expiring"`
}

// NewCache 创建一个新的 Cache 实例
func NewCache(opt *Cache) *Cache {
	var cacheInstance *Cache

	internal.LockFunc(func() {
		cacheInstance = opt
	})
	return cacheInstance
}

// Clone 返回 Cache 配置的副本
func (c *Cache) Clone() internal.Configurable {
	return &Cache{
		ModuleName: c.ModuleName,
		Type:       c.Type,
		Enabled:    c.Enabled,
		Memory:     c.Memory,
		Ristretto:  c.Ristretto,
		Redis:      c.Redis,
		Sharded:    c.Sharded,
		TwoLevel:   c.TwoLevel,
		Expiring:   c.Expiring,
		DefaultTTL: c.DefaultTTL,
		KeyPrefix:  c.KeyPrefix,
		Serializer: c.Serializer,
	}
}

// Get 返回 Cache 配置的所有字段
func (c *Cache) Get() interface{} {
	return c
}

// Set 更新 Cache 配置的字段
func (c *Cache) Set(data interface{}) {
	if configData, ok := data.(*Cache); ok {
		c.ModuleName = configData.ModuleName
		c.Type = configData.Type
		c.Enabled = configData.Enabled
		c.Memory = configData.Memory
		c.Ristretto = configData.Ristretto
		c.Redis = configData.Redis
		c.Sharded = configData.Sharded
		c.TwoLevel = configData.TwoLevel
		c.Expiring = configData.Expiring
		c.DefaultTTL = configData.DefaultTTL
		c.KeyPrefix = configData.KeyPrefix
		c.Serializer = configData.Serializer
	}
}

// GetModuleName 返回模块名称
func (c *Cache) GetModuleName() string {
	return c.ModuleName
}

// IsEnabled 检查缓存是否启用
func (c *Cache) IsEnabled() bool {
	return c.Enabled
}

// GetType 获取缓存类型
func (c *Cache) GetType() CacheType {
	return c.Type
}

// GetKeyPrefix 获取键前缀
func (c *Cache) GetKeyPrefix() string {
	if c.KeyPrefix == "" {
		return "cache:"
	}
	return c.KeyPrefix
}

// GetSerializer 获取序列化器类型
func (c *Cache) GetSerializer() string {
	if c.Serializer == "" {
		return "json"
	}
	return c.Serializer
}

// GetDefaultTTL 获取默认TTL
func (c *Cache) GetDefaultTTL() time.Duration {
	if c.DefaultTTL == 0 {
		return 30 * time.Minute
	}
	return c.DefaultTTL
}

// Validate 验证配置是否有效
func (c *Cache) Validate() error {
	if !c.Enabled {
		return nil // 未启用时不需要验证
	}

	// 验证具体的缓存配置
	switch c.Type {
	case CacheTypeMemory:
		return c.Memory.Validate()
	case CacheTypeRistretto:
		return c.Ristretto.Validate()
	case CacheTypeRedis:
		return c.Redis.Validate()
	case CacheTypeSharded:
		return c.Sharded.Validate()
	case CacheTypeTwoLevel:
		return c.TwoLevel.Validate()
	case CacheTypeExpiring:
		return c.Expiring.Validate()
	}

	return internal.ValidateStruct(c)
}

// DefaultCache 返回默认缓存配置
func DefaultCache() Cache {
	return Cache{
		ModuleName: "default",
		Type:       CacheTypeMemory,
		Enabled:    true,
		DefaultTTL: 30 * time.Minute,
		KeyPrefix:  "cache:",
		Serializer: "json",
		Memory:     DefaultMemoryConfig(),
		Ristretto:  DefaultRistrettoConfig(),
		Redis:      DefaultRedisConfig(),
		Sharded:    DefaultShardedConfig(),
		TwoLevel:   DefaultTwoLevelConfig(),
		Expiring:   DefaultExpiringConfig(),
	}
}

// Default 返回默认缓存配置的指针，支持链式调用
func Default() *Cache {
	config := DefaultCache()
	return &config
}

// WithModuleName 设置模块名称
func (c *Cache) WithModuleName(moduleName string) *Cache {
	c.ModuleName = moduleName
	return c
}

// WithType 设置缓存类型
func (c *Cache) WithType(cacheType CacheType) *Cache {
	c.Type = cacheType
	return c
}

// WithEnabled 设置是否启用缓存
func (c *Cache) WithEnabled(enabled bool) *Cache {
	c.Enabled = enabled
	return c
}

// WithDefaultTTL 设置默认TTL
func (c *Cache) WithDefaultTTL(ttl time.Duration) *Cache {
	c.DefaultTTL = ttl
	return c
}

// WithKeyPrefix 设置键前缀
func (c *Cache) WithKeyPrefix(prefix string) *Cache {
	c.KeyPrefix = prefix
	return c
}

// WithSerializer 设置序列化器
func (c *Cache) WithSerializer(serializer string) *Cache {
	c.Serializer = serializer
	return c
}

// WithMemory 设置内存缓存配置
func (c *Cache) WithMemory(memory Memory) *Cache {
	c.Memory = memory
	return c
}

// WithRistretto 设置Ristretto缓存配置
func (c *Cache) WithRistretto(ristretto Ristretto) *Cache {
	c.Ristretto = ristretto
	return c
}

// WithRedis 设置Redis缓存配置
func (c *Cache) WithRedis(redis Redis) *Cache {
	c.Redis = redis
	return c
}

// WithSharded 设置分片缓存配置
func (c *Cache) WithSharded(sharded Sharded) *Cache {
	c.Sharded = sharded
	return c
}

// WithTwoLevel 设置二级缓存配置
func (c *Cache) WithTwoLevel(twoLevel TwoLevel) *Cache {
	c.TwoLevel = twoLevel
	return c
}

// WithExpiring 设置过期缓存配置
func (c *Cache) WithExpiring(expiring Expiring) *Cache {
	c.Expiring = expiring
	return c
}
