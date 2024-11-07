/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:56:59
 * @FilePath: \go-config\pkg\database\mysql.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package database

import (
	"github.com/kamalyes/go-config/internal"
)

// MySQL 数据库配置
type MySQL struct {
	Host            string `mapstructure:"host"                yaml:"host"                json:"host"            validate:"required"` // 数据库 IP 地址
	Port            string `mapstructure:"port"                yaml:"port"                json:"port"            validate:"required"` // 端口
	Config          string `mapstructure:"config"              yaml:"config"              json:"config"          validate:"required"` // 后缀配置 默认配置 charset=utf8mb4&parseTime=True&loc=Local
	LogLevel        string `mapstructure:"log-level"           yaml:"log-level"           json:"log_level"       validate:"required"` // SQL 日志等级
	Dbname          string `mapstructure:"db-name"             yaml:"db-name"             json:"db_name"         validate:"required"` // 数据库名称
	Username        string `mapstructure:"username"            yaml:"username"            json:"username"        validate:"required"` // 数据库用户名
	Password        string `mapstructure:"password"            yaml:"password"            json:"password"        validate:"required"` // 数据库密码
	MaxIdleConns    int    `mapstructure:"max-idle-conns"      yaml:"max-idle-conns"      json:"max_idle_conns"  validate:"min=0"`    // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"max-open-conns"      yaml:"max-open-conns"      json:"max_open_conns"  validate:"min=0"`    // 最大连接数
	ConnMaxIdleTime int    `mapstructure:"conn-max-idle-time"  yaml:"conn-max-idle-time"  json:"conn_max_idle_time" validate:"min=0"` // 连接最大空闲时间 单位：秒
	ConnMaxLifeTime int    `mapstructure:"conn-max-life-time"  yaml:"conn-max-life-time"  json:"conn_max_life_time" validate:"min=0"` // 连接最大生命周期 单位：秒
	ModuleName      string `mapstructure:"modulename"          yaml:"modulename"          json:"module_name"`                         // 模块名称
}

// NewMySQL 创建一个新的 MySQL 实例
func NewMySQL(opt *MySQL) *MySQL {
	var mysqlInstance *MySQL

	internal.LockFunc(func() {
		mysqlInstance = opt
	})
	return mysqlInstance
}

// Clone 返回 MySQL 配置的副本
func (m *MySQL) Clone() internal.Configurable {
	return &MySQL{
		ModuleName:      m.ModuleName,
		Host:            m.Host,
		Port:            m.Port,
		Config:          m.Config,
		LogLevel:        m.LogLevel,
		Dbname:          m.Dbname,
		Username:        m.Username,
		Password:        m.Password,
		MaxIdleConns:    m.MaxIdleConns,
		MaxOpenConns:    m.MaxOpenConns,
		ConnMaxIdleTime: m.ConnMaxIdleTime,
		ConnMaxLifeTime: m.ConnMaxLifeTime,
	}
}

// Get 返回 MySQL 配置的所有字段
func (m *MySQL) Get() interface{} {
	return m
}

// Set 更新 MySQL 配置的字段
func (m *MySQL) Set(data interface{}) {
	if configData, ok := data.(*MySQL); ok {
		m.ModuleName = configData.ModuleName
		m.Host = configData.Host
		m.Port = configData.Port
		m.Config = configData.Config
		m.LogLevel = configData.LogLevel
		m.Dbname = configData.Dbname
		m.Username = configData.Username
		m.Password = configData.Password
		m.MaxIdleConns = configData.MaxIdleConns
		m.MaxOpenConns = configData.MaxOpenConns
		m.ConnMaxIdleTime = configData.ConnMaxIdleTime
		m.ConnMaxLifeTime = configData.ConnMaxLifeTime
	}
}

// Validate 检查 MySQL 配置的有效性
func (m *MySQL) Validate() error {
	return internal.ValidateStruct(m)
}

// GetCommonConfig 返回公共配置
func (m MySQL) GetCommonConfig() *DBConfig {
	return &DBConfig{
		Host:            m.Host,
		Username:        m.Username,
		Password:        m.Password,
		Dbname:          m.Dbname,
		Port:            m.Port,
		Config:          m.Config,
		MaxIdleConns:    m.MaxIdleConns,
		MaxOpenConns:    m.MaxOpenConns,
		ConnMaxIdleTime: m.ConnMaxIdleTime,
		ConnMaxLifeTime: m.ConnMaxLifeTime,
		LogLevel:        m.LogLevel,
	}
}
