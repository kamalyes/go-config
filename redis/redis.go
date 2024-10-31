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

package redis

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

// 实现 Configurable 接口
func (r Redis) GetModuleName() string {
	return "redis"
}
