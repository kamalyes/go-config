/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:27:06
 * @FilePath: \go-config\tests\redis_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/redis"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Redis 配置参数
func generateRedisTestParams() *redis.Redis {
	return &redis.Redis{
		ModuleName:   random.RandString(10, random.CAPITAL),
		Addr:         fmt.Sprintf("%s:%d", random.RandString(5, random.CAPITAL), random.FRandInt(1, 65535)), // 随机生成地址
		DB:           random.FRandInt(0, 15),                                                                // Redis 默认数据库范围是 0-15
		Password:     random.RandString(16, random.CAPITAL),                                                 // 随机生成密码
		MaxRetries:   random.FRandInt(1, 10),                                                                // 随机重试次数
		MinIdleConns: random.FRandInt(1, 100),                                                               // 随机最大空闲连接数
		PoolSize:     random.FRandInt(1, 100),                                                               // 随机连接池大小
	}
}

func TestRedisClone(t *testing.T) {
	params := generateRedisTestParams()
	original := redis.NewRedis(params)
	cloned := original.Clone().(*redis.Redis)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestRedisSet(t *testing.T) {
	oldParams := generateRedisTestParams()
	newParams := generateRedisTestParams()

	redisInstance := redis.NewRedis(oldParams)
	newConfig := redis.NewRedis(newParams)

	redisInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, redisInstance.ModuleName)
	assert.Equal(t, newParams.Addr, redisInstance.Addr)
	assert.Equal(t, newParams.DB, redisInstance.DB)
	assert.Equal(t, newParams.Password, redisInstance.Password)
	assert.Equal(t, newParams.MaxRetries, redisInstance.MaxRetries)
	assert.Equal(t, newParams.MinIdleConns, redisInstance.MinIdleConns)
	assert.Equal(t, newParams.PoolSize, redisInstance.PoolSize)
}
