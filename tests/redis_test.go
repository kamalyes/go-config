/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 09:24:50
 * @FilePath: \go-config\redis\redis.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/redis"
	"github.com/stretchr/testify/assert"
)

// getValidRedisConfig 返回有效的 Redis 配置
func getValidRedisConfig() *redis.Redis {
	return redis.NewRedis("test-module", "127.0.0.1:6379", 0, "password", 3, 10, 20)
}

// getInvalidRedisConfig 返回无效的 Redis 配置
func getInvalidRedisConfig() *redis.Redis {
	return redis.NewRedis("", "", -1, "", -1, -1, 0)
}

func TestNewRedis(t *testing.T) {
	// 测试创建 Redis 实例
	redis := getValidRedisConfig()

	assert.NotNil(t, redis)
	assert.Equal(t, "test-module", redis.ModuleName)
	assert.Equal(t, "127.0.0.1:6379", redis.Addr)
	assert.Equal(t, 0, redis.DB)
	assert.Equal(t, "password", redis.Password)
	assert.Equal(t, 3, redis.MaxRetries)
	assert.Equal(t, 10, redis.MinIdleConns)
	assert.Equal(t, 20, redis.PoolSize)
}

func TestRedis_Validate(t *testing.T) {
	// 测试有效配置
	validRedis := getValidRedisConfig()
	assert.NoError(t, validRedis.Validate())

	// 测试无效配置
	invalidRedis := getInvalidRedisConfig()
	assert.Error(t, invalidRedis.Validate())
}

func TestRedis_ToMap(t *testing.T) {
	// 测试将配置转换为映射
	r := getValidRedisConfig()
	rMap := r.ToMap()

	assert.Equal(t, "test-module", rMap["moduleName"])
	assert.Equal(t, "127.0.0.1:6379", rMap["addr"])
	assert.Equal(t, 0, rMap["db"])
	assert.Equal(t, "password", rMap["password"])
	assert.Equal(t, 3, rMap["maxRetries"])
	assert.Equal(t, 10, rMap["minIdleConns"])
	assert.Equal(t, 20, rMap["poolSize"])
}

func TestRedis_FromMap(t *testing.T) {
	// 测试从映射填充配置
	data := map[string]interface{}{
		"moduleName":   "test-module",
		"addr":         "127.0.0.1:6379",
		"db":           0,
		"password":     "password",
		"maxRetries":   3,
		"minIdleConns": 10,
		"poolSize":     20,
	}

	r := &redis.Redis{}
	r.FromMap(data)

	assert.Equal(t, "test-module", r.ModuleName)
	assert.Equal(t, "127.0.0.1:6379", r.Addr)
	assert.Equal(t, 0, r.DB)
	assert.Equal(t, "password", r.Password)
	assert.Equal(t, 3, r.MaxRetries)
	assert.Equal(t, 10, r.MinIdleConns)
	assert.Equal(t, 20, r.PoolSize)
}
