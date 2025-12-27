/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 13:40:00
 * @FilePath: \go-config\pkg\cache\sharded_test.go
 * @Description: Sharded缓存配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharded_Clone(t *testing.T) {
	original := &Sharded{
		ModuleName:   "test-sharded",
		ShardCount:   16,
		BaseType:     CacheTypeMemory,
		HashFunc:     "fnv",
		LoadBalancer: "consistent_hash",
	}

	cloned := original.Clone().(*Sharded)

	// 验证克隆后的值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.ShardCount, cloned.ShardCount)
	assert.Equal(t, original.BaseType, cloned.BaseType)
	assert.Equal(t, original.HashFunc, cloned.HashFunc)
	assert.Equal(t, original.LoadBalancer, cloned.LoadBalancer)

	// 修改原始对象不应影响克隆对象
	original.ShardCount = 32
	original.HashFunc = "md5"
	original.LoadBalancer = "round_robin"

	assert.NotEqual(t, original.ShardCount, cloned.ShardCount)
	assert.NotEqual(t, original.HashFunc, cloned.HashFunc)
	assert.NotEqual(t, original.LoadBalancer, cloned.LoadBalancer)
}
