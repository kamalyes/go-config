/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 11:30:00
 * @FilePath: \go-config\pkg\cache\ristretto_test.go
 * @Description: Ristretto缓存配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRistretto_Clone(t *testing.T) {
	original := &Ristretto{
		ModuleName:         "test-ristretto",
		NumCounters:        1000000,
		MaxCost:            10485760,
		BufferItems:        64,
		Metrics:            true,
		IgnoreInternalCost: false,
		KeyToHash:          true,
		Cost:               1,
	}

	cloned := original.Clone().(*Ristretto)

	// 验证克隆后的值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.NumCounters, cloned.NumCounters)
	assert.Equal(t, original.MaxCost, cloned.MaxCost)
	assert.Equal(t, original.BufferItems, cloned.BufferItems)
	assert.Equal(t, original.Metrics, cloned.Metrics)
	assert.Equal(t, original.IgnoreInternalCost, cloned.IgnoreInternalCost)
	assert.Equal(t, original.KeyToHash, cloned.KeyToHash)
	assert.Equal(t, original.Cost, cloned.Cost)

	// 修改原始对象不应影响克隆对象
	original.NumCounters = 2000000
	original.MaxCost = 20971520
	original.Metrics = false

	assert.NotEqual(t, original.NumCounters, cloned.NumCounters)
	assert.NotEqual(t, original.MaxCost, cloned.MaxCost)
	assert.NotEqual(t, original.Metrics, cloned.Metrics)
}
