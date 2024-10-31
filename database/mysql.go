/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 10:59:24
 * @FilePath: \go-config\database\mysql.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package database

// MySQL 配置
type MySQL struct {

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** 数据库IP地址 */
	Host string `mapstructure:"host"                     json:"host"                     yaml:"host"`

	/** 端口 */
	Port string `mapstructure:"port"                     json:"port"                     yaml:"port"`

	/** 后缀配置  默认配置 charset=utf8mb4&parseTime=True&loc=Local */
	Config string `mapstructure:"config"                   json:"config"                   yaml:"config"`

	/** sql日志等级 */
	LogLevel string `mapstructure:"log-level"                json:"logLevel"                 yaml:"log-level"`

	/** 数据库名称 */
	Dbname string `mapstructure:"db-name"                  json:"dbname"                   yaml:"db-name"`

	/** 数据库用户名 */
	Username string `mapstructure:"username"                 json:"username"                 yaml:"username"`

	/** 数据库密码 */
	Password string `mapstructure:"password"                 json:"password"                 yaml:"password"`

	/** 最大空闲连接数 */
	MaxIdleConns int `mapstructure:"max-idle-conns"           json:"maxIdleConns"             yaml:"max-idle-conns"`

	/** 最大连接数 */
	MaxOpenConns int `mapstructure:"max-open-conns"           json:"maxOpenConns"             yaml:"max-open-conns"`

	/** 连接最大空闲时间  单位：秒 */
	ConnMaxIdleTime int `mapstructure:"conn-max-idle-time"       json:"connMaxIdleTime"          yaml:"conn-max-idle-time"`

	/** 连接最大生命周期  单位：秒 */
	ConnMaxLifetime int `mapstructure:"conn-max-life-time"       json:"connMaxLifetime"          yaml:"conn-max-life-time"`
}

// 实现 Configurable 接口
func (m MySQL) GetModuleName() string {
	return "mysql"
}
