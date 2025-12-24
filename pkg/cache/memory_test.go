/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-24 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-24 00:00:00
 * @FilePath: \go-config\pkg\cache\memory_test.go
 * @Description: Memory缓存配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemory_Clone(t *testing.T) {
	memory := &Memory{
		ModuleName:  "memory",
		Capacity:    1000,
		DefaultTTL:  600,
		CleanupSize: 100,
		MaxSize:     10000,
	}

	cloned := memory.Clone()

	assert.NotNil(t, cloned)
	clonedMemory, ok := cloned.(*Memory)
	assert.True(t, ok)
	assert.Equal(t, memory.Capacity, clonedMemory.Capacity)
	assert.Equal(t, memory.DefaultTTL, clonedMemory.DefaultTTL)
	assert.Equal(t, memory.CleanupSize, clonedMemory.CleanupSize)
	assert.Equal(t, memory.MaxSize, clonedMemory.MaxSize)

	// 验证是独立副本
	clonedMemory.Capacity = 2000
	assert.NotEqual(t, memory.Capacity, clonedMemory.Capacity)
}
