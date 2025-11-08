/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 00:16:03
 * @FilePath: \go-config\pkg\cache\memory.go
 * @Description: 内存缓存配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import (
	"time"

	"github.com/kamalyes/go-config/internal"
)

// Memory 内存缓存配置
type Memory struct {
	ModuleName  string        `mapstructure:"module_name" yaml:"module_name" json:"module_name"`                               // 模块名
	Capacity    int           `mapstructure:"capacity" yaml:"capacity" json:"capacity" validate:"min=1"`                       // 缓存容量
	DefaultTTL  time.Duration `mapstructure:"default_ttl" yaml:"default_ttl" json:"default_ttl"`                            // 默认过期时间
	CleanupSize int           `mapstructure:"cleanup_size" yaml:"cleanup_size" json:"cleanup_size" validate:"min=0"`         // 清理大小
	MaxSize     int           `mapstructure:"max_size" yaml:"max_size" json:"max_size" validate:"min=0"`                     // 最大大小
}

// NewMemory 创建一个新的 Memory 实例
func NewMemory(opt *Memory) *Memory {
	var memoryInstance *Memory
	internal.LockFunc(func() {
		memoryInstance = opt
	})
	return memoryInstance
}

// Clone 返回 Memory 配置的副本
func (m *Memory) Clone() internal.Configurable {
	return &Memory{
		ModuleName:  m.ModuleName,
		Capacity:    m.Capacity,
		DefaultTTL:  m.DefaultTTL,
		CleanupSize: m.CleanupSize,
		MaxSize:     m.MaxSize,
	}
}

// Get 返回 Memory 配置的所有字段
func (m *Memory) Get() interface{} {
	return m
}

// Set 更新 Memory 配置的字段
func (m *Memory) Set(data interface{}) {
	if configData, ok := data.(*Memory); ok {
		m.ModuleName = configData.ModuleName
		m.Capacity = configData.Capacity
		m.DefaultTTL = configData.DefaultTTL
		m.CleanupSize = configData.CleanupSize
		m.MaxSize = configData.MaxSize
	}
}

// DefaultMemoryConfig 返回默认内存缓存配置
func DefaultMemoryConfig() Memory {
	return Memory{
		ModuleName:  "memory",
		Capacity:    1000,
		DefaultTTL:  30 * time.Minute,
		CleanupSize: 100,
		MaxSize:     0, // 0表示无限制
	}
}

// DefaultMemory 返回默认内存缓存配置的指针，支持链式调用
func DefaultMemory() *Memory {
	config := DefaultMemoryConfig()
	return &config
}

// WithModuleName 设置模块名称
func (m *Memory) WithModuleName(moduleName string) *Memory {
	m.ModuleName = moduleName
	return m
}

// WithCapacity 设置缓存容量
func (m *Memory) WithCapacity(capacity int) *Memory {
	m.Capacity = capacity
	return m
}

// WithDefaultTTL 设置默认过期时间
func (m *Memory) WithDefaultTTL(defaultTTL time.Duration) *Memory {
	m.DefaultTTL = defaultTTL
	return m
}

// WithCleanupSize 设置清理大小
func (m *Memory) WithCleanupSize(cleanupSize int) *Memory {
	m.CleanupSize = cleanupSize
	return m
}

// WithMaxSize 设置最大大小
func (m *Memory) WithMaxSize(maxSize int) *Memory {
	m.MaxSize = maxSize
	return m
}

// Validate 验证内存缓存配置
func (m *Memory) Validate() error {
	if m.Capacity <= 0 {
		m.Capacity = 1000
	}
	if m.DefaultTTL <= 0 {
		m.DefaultTTL = 30 * time.Minute
	}
	if m.CleanupSize < 0 {
		m.CleanupSize = 100
	}
	return internal.ValidateStruct(m)
}