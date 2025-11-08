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

// DefaultMySQL 返回默认MySQL配置
func DefaultMySQL() MySQL {
	return MySQL{
		ModuleName:      "mysql",
		Host:            "127.0.0.1",
		Port:            "3306",
		Config:          "charset=utf8mb4&parseTime=True&loc=Local",
		LogLevel:        "silent",
		Dbname:          "",
		Username:        "",
		Password:        "",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxIdleTime: 3600,  // 1小时
		ConnMaxLifeTime: 7200,  // 2小时
	}
}

// Default 返回默认MySQL配置的指针，支持链式调用
func Default() *MySQL {
	config := DefaultMySQL()
	return &config
}

// WithModuleName 设置模块名称
func (m *MySQL) WithModuleName(moduleName string) *MySQL {
	m.ModuleName = moduleName
	return m
}

// WithHost 设置主机地址
func (m *MySQL) WithHost(host string) *MySQL {
	m.Host = host
	return m
}

// WithPort 设置端口
func (m *MySQL) WithPort(port string) *MySQL {
	m.Port = port
	return m
}

// WithConfig 设置配置字符串
func (m *MySQL) WithConfig(config string) *MySQL {
	m.Config = config
	return m
}

// WithLogLevel 设置日志级别
func (m *MySQL) WithLogLevel(logLevel string) *MySQL {
	m.LogLevel = logLevel
	return m
}

// WithDbname 设置数据库名称
func (m *MySQL) WithDbname(dbname string) *MySQL {
	m.Dbname = dbname
	return m
}

// WithUsername 设置用户名
func (m *MySQL) WithUsername(username string) *MySQL {
	m.Username = username
	return m
}

// WithPassword 设置密码
func (m *MySQL) WithPassword(password string) *MySQL {
	m.Password = password
	return m
}

// WithMaxIdleConns 设置最大空闲连接数
func (m *MySQL) WithMaxIdleConns(maxIdleConns int) *MySQL {
	m.MaxIdleConns = maxIdleConns
	return m
}

// WithMaxOpenConns 设置最大连接数
func (m *MySQL) WithMaxOpenConns(maxOpenConns int) *MySQL {
	m.MaxOpenConns = maxOpenConns
	return m
}

// WithConnMaxIdleTime 设置连接最大空闲时间
func (m *MySQL) WithConnMaxIdleTime(connMaxIdleTime int) *MySQL {
	m.ConnMaxIdleTime = connMaxIdleTime
	return m
}

// WithConnMaxLifeTime 设置连接最大生命周期
func (m *MySQL) WithConnMaxLifeTime(connMaxLifeTime int) *MySQL {
	m.ConnMaxLifeTime = connMaxLifeTime
	return m
}
