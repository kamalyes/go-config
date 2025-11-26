/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 23:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 16:54:04
 * @FilePath: \go-config\pkg\database\mysql.go
 * @Description: MySQL数据库配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package database

import (
	"github.com/kamalyes/go-config/internal"
)

// MySQL MySQL数据库配置
type MySQL struct {
	ModuleName      string `mapstructure:"module_name" yaml:"module_name" json:"module_name"`                                       // 模块名称
	Host            string `mapstructure:"host" yaml:"host" json:"host"            validate:"required"`                             // 数据库 IP 地址
	Port            string `mapstructure:"port" yaml:"port" json:"port"            validate:"required"`                             // 端口
	Config          string `mapstructure:"config" yaml:"config" json:"config"          validate:"required"`                         // 后缀配置 默认配置 charset=utf8mb4&parseTime=True&loc=Local
	LogLevel        string `mapstructure:"log_level" yaml:"log-level" json:"log_level"       validate:"required"`                   // SQL 日志等级
	Dbname          string `mapstructure:"db_name" yaml:"db-name" json:"db_name"         validate:"required"`                       // 数据库名称
	Username        string `mapstructure:"username" yaml:"username" json:"username"        validate:"required"`                     // 数据库用户名
	Password        string `mapstructure:"password" yaml:"password" json:"password"        validate:"required"`                     // 数据库密码
	MaxIdleConns    int    `mapstructure:"max_idle_conns" yaml:"max-idle-conns" json:"max_idle_conns"  validate:"min=0"`            // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"max_open_conns" yaml:"max-open-conns" json:"max_open_conns"  validate:"min=0"`            // 最大连接数
	ConnMaxIdleTime int    `mapstructure:"conn_max_idle_time" yaml:"conn-max-idle-time" json:"conn_max_idle_time" validate:"min=0"` // 连接最大空闲时间 单位：秒
	ConnMaxLifeTime int    `mapstructure:"conn_max_life_time" yaml:"conn-max-life-time" json:"conn_max_life_time" validate:"min=0"` // 连接最大生命周期 单位：秒
}

// 为MySQL配置实现DatabaseProvider接口
func (m *MySQL) GetDBType() DBType     { return DBTypeMySQL }
func (m *MySQL) GetHost() string       { return m.Host }
func (m *MySQL) GetPort() string       { return m.Port }
func (m *MySQL) GetDBName() string     { return m.Dbname }
func (m *MySQL) GetUsername() string   { return m.Username }
func (m *MySQL) GetPassword() string   { return m.Password }
func (m *MySQL) GetConfig() string     { return m.Config }
func (m *MySQL) GetModuleName() string { return m.ModuleName }
func (m *MySQL) SetCredentials(username, password string) {
	m.Username, m.Password = username, password
}
func (m *MySQL) SetHost(host string)     { m.Host = host }
func (m *MySQL) SetPort(port string)     { m.Port = port }
func (m *MySQL) SetDBName(dbName string) { m.Dbname = dbName }

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

func (m *MySQL) Get() interface{} { return m }
func (m *MySQL) Set(data interface{}) {
	if cfg, ok := data.(*MySQL); ok {
		*m = *cfg
	}
}
func (m *MySQL) Validate() error { return internal.ValidateStruct(m) }

// DefaultMySQL 创建默认MySQL配置
func DefaultMySQL() *MySQL {
	return &MySQL{
		ModuleName:      "mysql",
		Host:            "localhost",
		Port:            "3306",
		Config:          "charset=utf8mb4&parseTime=True&loc=Local",
		LogLevel:        "info",
		Dbname:          "test",
		Username:        "root",
		Password:        "mysql_password",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxIdleTime: 300,  // 5分钟
		ConnMaxLifeTime: 3600, // 1小时
	}
}

// NewMySQLConfig 创建一个新的MySQL配置实例
func NewMySQLConfig(opt *MySQL) *MySQL {
	var mysqlInstance *MySQL
	internal.LockFunc(func() {
		mysqlInstance = opt
	})
	return mysqlInstance
}

// WithModuleName 设置模块名称
func (m *MySQL) WithModuleName(moduleName string) *MySQL {
	m.ModuleName = moduleName
	return m
}

// WithHost 设置数据库主机地址
func (m *MySQL) WithHost(host string) *MySQL {
	m.Host = host
	return m
}

// WithPort 设置数据库端口
func (m *MySQL) WithPort(port string) *MySQL {
	m.Port = port
	return m
}

// WithConfig 设置数据库配置参数
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

// WithUsername 设置数据库用户名
func (m *MySQL) WithUsername(username string) *MySQL {
	m.Username = username
	return m
}

// WithPassword 设置数据库密码
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
