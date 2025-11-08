/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:51:41
 * @FilePath: \go-config\pkg\redis\redis.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package redis

import (
	"github.com/kamalyes/go-config/pkg/cache"
)

// Redis 结构体用于配置 Redis 相关参数（已废弃，请使用 cache.Redis）
// Deprecated: 使用 cache.Redis 替代
type Redis = cache.Redis

// NewRedis 创建一个新的 Redis 实例（已废弃，请使用 cache.NewRedis）
// Deprecated: 使用 cache.NewRedis 替代
func NewRedis(opt *Redis) *Redis {
	return cache.NewRedis(opt)
}
