/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 00:18:20
 * @FilePath: \go-config\pkg\cache\redis.go
 * @Description: Redis 缓存配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import (
	"github.com/kamalyes/go-config/internal"
	"time"
)

// Redis 结构体用于配置 Redis 相关参数（增强版配置）
type Redis struct {
	ModuleName string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"` // 模块名
	// 兼容原有配置
	Addr string `mapstructure:"addr" yaml:"addr" json:"addr" validate:"url"` // Redis 数据服务器 IP 和端口（兼容旧版）
	// 新增增强配置
	Addrs        []string      `mapstructure:"addrs" yaml:"addrs" json:"addrs"`                                           // Redis服务器地址列表（集群模式）
	Username     string        `mapstructure:"username" yaml:"username" json:"username"`                                  // 用户名
	Password     string        `mapstructure:"password" yaml:"password" json:"password"`                                  // 连接密码
	DB           int           `mapstructure:"db" yaml:"db" json:"db" validate:"min=0"`                                   // 指定连接的数据库，默认连数据库 0
	MaxRetries   int           `mapstructure:"max-retries" yaml:"max-retries" json:"maxRetries" validate:"min=0"`         // 最大重试次数，最小值为 0
	PoolSize     int           `mapstructure:"pool-size" yaml:"pool-size" json:"poolSize" validate:"min=1"`               // 连接池大小，最小值为 1
	MinIdleConns int           `mapstructure:"min-idle-conns" yaml:"min-idle-conns" json:"minIdleConns" validate:"min=0"` // 最小空闲连接数，最小值为 0
	MaxConnAge   time.Duration `mapstructure:"max-conn-age" yaml:"max-conn-age" json:"maxConnAge"`                        // 连接最大存活时间
	PoolTimeout  time.Duration `mapstructure:"pool-timeout" yaml:"pool-timeout" json:"poolTimeout"`                       // 连接池超时
	IdleTimeout  time.Duration `mapstructure:"idle-timeout" yaml:"idle-timeout" json:"idleTimeout"`                       // 空闲超时
	ReadTimeout  time.Duration `mapstructure:"read-timeout" yaml:"read-timeout" json:"readTimeout"`                       // 读取超时
	WriteTimeout time.Duration `mapstructure:"write-timeout" yaml:"write-timeout" json:"writeTimeout"`                    // 写入超时
	ClusterMode  bool          `mapstructure:"cluster-mode" yaml:"cluster-mode" json:"clusterMode"`                       // 是否集群模式
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
		Addrs:        append([]string(nil), r.Addrs...),
		Username:     r.Username,
		Password:     r.Password,
		DB:           r.DB,
		MaxRetries:   r.MaxRetries,
		PoolSize:     r.PoolSize,
		MinIdleConns: r.MinIdleConns,
		MaxConnAge:   r.MaxConnAge,
		PoolTimeout:  r.PoolTimeout,
		IdleTimeout:  r.IdleTimeout,
		ReadTimeout:  r.ReadTimeout,
		WriteTimeout: r.WriteTimeout,
		ClusterMode:  r.ClusterMode,
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
		r.Addrs = configData.Addrs
		r.Username = configData.Username
		r.Password = configData.Password
		r.DB = configData.DB
		r.MaxRetries = configData.MaxRetries
		r.PoolSize = configData.PoolSize
		r.MinIdleConns = configData.MinIdleConns
		r.MaxConnAge = configData.MaxConnAge
		r.PoolTimeout = configData.PoolTimeout
		r.IdleTimeout = configData.IdleTimeout
		r.ReadTimeout = configData.ReadTimeout
		r.WriteTimeout = configData.WriteTimeout
		r.ClusterMode = configData.ClusterMode
	}
}

// Validate 验证 Redis 配置的有效性
func (r *Redis) Validate() error {
	// 处理地址列表
	if len(r.Addrs) == 0 && r.Addr != "" {
		r.Addrs = []string{r.Addr}
	}
	if len(r.Addrs) == 0 {
		r.Addrs = []string{"127.0.0.1:6379"}
	}

	// 设置默认值
	if r.MaxRetries < 0 {
		r.MaxRetries = 3
	}
	if r.PoolSize <= 0 {
		r.PoolSize = 10
	}
	if r.MinIdleConns < 0 {
		r.MinIdleConns = 0
	}
	if r.MaxConnAge <= 0 {
		r.MaxConnAge = 30 * time.Minute
	}
	if r.PoolTimeout <= 0 {
		r.PoolTimeout = 4 * time.Second
	}
	if r.IdleTimeout <= 0 {
		r.IdleTimeout = 5 * time.Minute
	}
	if r.ReadTimeout <= 0 {
		r.ReadTimeout = 3 * time.Second
	}
	if r.WriteTimeout <= 0 {
		r.WriteTimeout = 3 * time.Second
	}

	return internal.ValidateStruct(r)
}

// DefaultRedisConfig 返回默认Redis配置
func DefaultRedisConfig() Redis {
	return Redis{
		ModuleName:   "redis",
		Addr:         "127.0.0.1:6379",
		Addrs:        []string{"127.0.0.1:6379"},
		Username:     "default",
		Password:     "redis123456",
		DB:           0,
		MaxRetries:   3,
		PoolSize:     10,
		MinIdleConns: 0,
		MaxConnAge:   30 * time.Minute,
		PoolTimeout:  4 * time.Second,
		IdleTimeout:  5 * time.Minute,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		ClusterMode:  false,
	}
}

// DefaultRedisConfig 返回默认Redis配置的指针，支持链式调用
func DefaultRedis() *Redis {
	config := DefaultRedisConfig()
	return &config
}

// WithModuleName 设置模块名称
func (r *Redis) WithModuleName(moduleName string) *Redis {
	r.ModuleName = moduleName
	return r
}

// WithAddr 设置Redis地址（单实例）
func (r *Redis) WithAddr(addr string) *Redis {
	r.Addr = addr
	if r.Addrs == nil {
		r.Addrs = []string{addr}
	} else if len(r.Addrs) == 1 {
		r.Addrs[0] = addr
	} else {
		r.Addrs = []string{addr}
	}
	return r
}

// WithAddrs 设置Redis地址列表（集群模式）
func (r *Redis) WithAddrs(addrs []string) *Redis {
	r.Addrs = addrs
	if len(addrs) > 0 {
		r.Addr = addrs[0]
	}
	return r
}

// WithUsername 设置用户名
func (r *Redis) WithUsername(username string) *Redis {
	r.Username = username
	return r
}

// WithPassword 设置密码
func (r *Redis) WithPassword(password string) *Redis {
	r.Password = password
	return r
}

// WithDB 设置数据库编号
func (r *Redis) WithDB(db int) *Redis {
	r.DB = db
	return r
}

// WithMaxRetries 设置最大重试次数
func (r *Redis) WithMaxRetries(maxRetries int) *Redis {
	r.MaxRetries = maxRetries
	return r
}

// WithPoolSize 设置连接池大小
func (r *Redis) WithPoolSize(poolSize int) *Redis {
	r.PoolSize = poolSize
	return r
}

// WithMinIdleConns 设置最小空闲连接数
func (r *Redis) WithMinIdleConns(minIdleConns int) *Redis {
	r.MinIdleConns = minIdleConns
	return r
}

// WithMaxConnAge 设置连接最大存活时间
func (r *Redis) WithMaxConnAge(maxConnAge time.Duration) *Redis {
	r.MaxConnAge = maxConnAge
	return r
}

// WithPoolTimeout 设置连接池超时
func (r *Redis) WithPoolTimeout(poolTimeout time.Duration) *Redis {
	r.PoolTimeout = poolTimeout
	return r
}

// WithIdleTimeout 设置空闲超时
func (r *Redis) WithIdleTimeout(idleTimeout time.Duration) *Redis {
	r.IdleTimeout = idleTimeout
	return r
}

// WithReadTimeout 设置读取超时
func (r *Redis) WithReadTimeout(readTimeout time.Duration) *Redis {
	r.ReadTimeout = readTimeout
	return r
}

// WithWriteTimeout 设置写入超时
func (r *Redis) WithWriteTimeout(writeTimeout time.Duration) *Redis {
	r.WriteTimeout = writeTimeout
	return r
}

// WithClusterMode 设置是否集群模式
func (r *Redis) WithClusterMode(clusterMode bool) *Redis {
	r.ClusterMode = clusterMode
	return r
}
