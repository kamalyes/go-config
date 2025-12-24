/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-24 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-24 00:00:00
 * @FilePath: \go-config\pkg\cache\expiring_test.go
 * @Description: Expiring缓存配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpiring_Clone(t *testing.T) {
	expiring := &Expiring{
		ModuleName:       "expiring",
		CleanupInterval:  60,
		DefaultTTL:       300,
		MaxSize:          1000,
		EvictionPolicy:   "LRU",
		EnableLazyExpiry: true,
		MaxMemoryUsage:   1024,
	}

	cloned := expiring.Clone()

	assert.NotNil(t, cloned)
	clonedExpiring, ok := cloned.(*Expiring)
	assert.True(t, ok)
	assert.Equal(t, expiring.CleanupInterval, clonedExpiring.CleanupInterval)
	assert.Equal(t, expiring.DefaultTTL, clonedExpiring.DefaultTTL)
	assert.Equal(t, expiring.MaxSize, clonedExpiring.MaxSize)
	assert.Equal(t, expiring.EvictionPolicy, clonedExpiring.EvictionPolicy)
	assert.Equal(t, expiring.EnableLazyExpiry, clonedExpiring.EnableLazyExpiry)

	// 验证是独立副本
	clonedExpiring.CleanupInterval = 120
	assert.NotEqual(t, expiring.CleanupInterval, clonedExpiring.CleanupInterval)
}
