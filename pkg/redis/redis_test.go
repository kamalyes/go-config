/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 13:55:15
 * @FilePath: \go-config\pkg\redis\redis_test.go
 * @Description: Redis配置测试（别名包）
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedis_Clone(t *testing.T) {
	// Redis 是 cache.Redis 的别名，测试别名是否正常工作
	original := &Redis{
		ModuleName: "test-redis",
		Addr:       "localhost:6379",
		Password:   "test-pass",
		DB:         0,
		PoolSize:   10,
	}

	cloned := original.Clone().(*Redis)

	// 验证克隆后的值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Addr, cloned.Addr)
	assert.Equal(t, original.Password, cloned.Password)
	assert.Equal(t, original.DB, cloned.DB)
	assert.Equal(t, original.PoolSize, cloned.PoolSize)

	// 修改原始对象不应影响克隆对象
	original.Addr = "localhost:6380"
	original.DB = 1
	assert.NotEqual(t, original.Addr, cloned.Addr)
	assert.NotEqual(t, original.DB, cloned.DB)
}
