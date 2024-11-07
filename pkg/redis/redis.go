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
	"github.com/kamalyes/go-config/internal"
)

// Redis 结构体用于配置 Redis 相关参数
type Redis struct {
	Addr         string `mapstructure:"addr"                     yaml:"addr"             json:"addr"             validate:"required,url"`   // Redis 数据服务器 IP 和端口
	DB           int    `mapstructure:"db"                       yaml:"db"               json:"db"               validate:"required,min=0"` // 指定连接的数据库，默认连数据库 0
	MaxRetries   int    `mapstructure:"max-retries"              yaml:"max-retries"      json:"max_retries"      validate:"min=0"`          // 最大重试次数，最小值为 0
	MinIdleConns int    `mapstructure:"min-idle-conns"           yaml:"min-idle-conns"   json:"min_idle_conns"   validate:"min=0"`          // 最大空闲连接数，最小值为 0
	PoolSize     int    `mapstructure:"pool-size"                yaml:"pool-size"        json:"pool_size"        validate:"min=1"`          // 连接池大小，最小值为 1
	Password     string `mapstructure:"password"                 yaml:"password"         json:"password"`                                   // 连接密码
	ModuleName   string `mapstructure:"modulename"               yaml:"modulename"       json:"module_name"`                                // 模块名称
}

// NewRedis 创建一个新的 Redis 实例
func NewRedis(opt *Redis) *Redis {
	var redisInstance *Redis

	internal.LockFunc(func() {
		redisInstance = opt
	})
	return redisInstance
}

// Clone 返回 Redis 配置的副本
func (r *Redis) Clone() internal.Configurable {
	return &Redis{
		ModuleName:   r.ModuleName,
		Addr:         r.Addr,
		DB:           r.DB,
		Password:     r.Password,
		MaxRetries:   r.MaxRetries,
		MinIdleConns: r.MinIdleConns,
		PoolSize:     r.PoolSize,
	}
}

// Get 返回 Redis 配置的所有字段
func (r *Redis) Get() interface{} {
	return r
}

// Set 更新 Redis 配置的字段
func (r *Redis) Set(data interface{}) {
	if configData, ok := data.(*Redis); ok {
		r.ModuleName = configData.ModuleName
		r.Addr = configData.Addr
		r.DB = configData.DB
		r.Password = configData.Password
		r.MaxRetries = configData.MaxRetries
		r.MinIdleConns = configData.MinIdleConns
		r.PoolSize = configData.PoolSize
	}
}

// Validate 验证 Redis 配置的有效性
func (r *Redis) Validate() error {
	return internal.ValidateStruct(r)
}
