/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:53:50
 * @FilePath: \go-config\pkg\redis\redis.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package redis

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Redis 结构体用于配置 Redis 相关参数
type Redis struct {
	ModuleName   string `mapstructure:"MODULE_NAME"              yaml:"modulename"`     // 模块名称
	Addr         string `mapstructure:"ADDR"                     yaml:"addr"`           // Redis 数据服务器 IP 和端口
	DB           int    `mapstructure:"DB"                       yaml:"db"`             // 指定连接的数据库，默认连数据库 0
	Password     string `mapstructure:"PASSWORD"                 yaml:"password"`       // 连接密码
	MaxRetries   int    `mapstructure:"MAX_RETRIES"              yaml:"max-retries"`    // 最大重试次数
	MinIdleConns int    `mapstructure:"MIN_IDLE_CONNS"           yaml:"min-idle-conns"` // 最大空闲连接数
	PoolSize     int    `mapstructure:"POOL_SIZE"                yaml:"pool-size"`      // 连接池大小
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
	if r.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if r.Addr == "" {
		return errors.New("address cannot be empty")
	}
	if r.DB < 0 {
		return errors.New("database number cannot be negative")
	}
	if r.MaxRetries < 0 {
		return errors.New("max retries cannot be negative")
	}
	if r.MinIdleConns < 0 {
		return errors.New("min idle connections cannot be negative")
	}
	if r.PoolSize <= 0 {
		return errors.New("pool size must be greater than zero")
	}
	return nil
}
