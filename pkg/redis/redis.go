/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:53:50
 * @FilePath: \go-config\redis\redis.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package redis

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

type Redis struct {

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** redis 数据服务器ip和端口 */
	Addr string `mapstructure:"addr"              json:"addr"             yaml:"addr"`

	/** 指定连接的数据库 默认连数据库0 */
	DB int `mapstructure:"db"                json:"db"               yaml:"db"`

	/** 连接密码 */
	Password string `mapstructure:"password"          json:"password"         yaml:"password"`

	/** 最大重试次数 */
	MaxRetries int `mapstructure:"max-retries"       json:"maxRetries"       yaml:"max-retries"`

	/** 最大空闲连接数 */
	MinIdleConns int `mapstructure:"min-idle-conns"    json:"minIdleConns"     yaml:"min-idle-conns"`

	/** 连接池大小 */
	PoolSize int `mapstructure:"pool-size"         json:"poolSize"         yaml:"pool-size"`
}

// NewRedis 创建一个新的 Redis 实例
func NewRedis(moduleName, addr string, db int, password string, maxRetries, minIdleConns, poolSize int) *Redis {
	var redisInstance *Redis

	internal.LockFunc(func() {
		redisInstance = &Redis{
			ModuleName:   moduleName,
			Addr:         addr,
			DB:           db,
			Password:     password,
			MaxRetries:   maxRetries,
			MinIdleConns: minIdleConns,
			PoolSize:     poolSize,
		}
	})
	return redisInstance
}

// ToMap 将配置转换为映射
func (r *Redis) ToMap() map[string]interface{} {
	return internal.ToMap(r)
}

// FromMap 从映射中填充配置
func (r *Redis) FromMap(data map[string]interface{}) {
	internal.FromMap(r, data)
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
