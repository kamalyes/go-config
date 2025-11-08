/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:30:50
 * @FilePath: \go-config\pkg\database\postgre.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package database

import (
	"github.com/kamalyes/go-config/internal"
)

// PostgreSQL 数据库配置结构体
type PostgreSQL struct {
	Host            string `mapstructure:"host"                     yaml:"host"                json:"host"            validate:"required"` // 数据库IP地址
	Port            string `mapstructure:"port"                     yaml:"port"                json:"port"            validate:"required"` // 端口
	Config          string `mapstructure:"config"                   yaml:"config"              json:"config"          validate:"required"` // 后缀配置
	LogLevel        string `mapstructure:"log-level"                yaml:"log-level"           json:"log_level"       validate:"required"` // SQL日志等级
	Dbname          string `mapstructure:"db-name"                  yaml:"db-name"             json:"db_name"         validate:"required"` // 数据库名称
	Username        string `mapstructure:"username"                 yaml:"username"            json:"username"        validate:"required"` // 数据库用户名
	Password        string `mapstructure:"password"                 yaml:"password"            json:"password"        validate:"required"` // 数据库密码
	MaxIdleConns    int    `mapstructure:"max-idle-conns"           yaml:"max-idle-conns"      json:"max_idle_conns"  validate:"min=0"`    // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"max-open-conns"           yaml:"max-open-conns"      json:"max_open_conns"  validate:"min=0"`    // 最大连接数
	ConnMaxIdleTime int    `mapstructure:"conn-max-idle-time"       yaml:"conn-max-idle-time"  json:"conn_max_idle_time" validate:"min=0"` // 连接最大空闲时间（秒）
	ConnMaxLifeTime int    `mapstructure:"conn-max-life-time"       yaml:"conn-max-life-time"  json:"conn_max_life_time" validate:"min=0"` // 连接最大生命周期（秒）
	ModuleName      string `mapstructure:"modulename"               yaml:"modulename"          json:"module_name"`                         // 模块名称
}

// NewPostgreSQL 创建一个新的 PostgreSQL 实例
func NewPostgreSQL(opt *PostgreSQL) *PostgreSQL {
	var postgreInstance *PostgreSQL

	internal.LockFunc(func() {
		postgreInstance = opt
	})
	return postgreInstance
}

// Clone 返回 PostgreSQL 配置的副本
func (p *PostgreSQL) Clone() internal.Configurable {
	return &PostgreSQL{
		ModuleName:      p.ModuleName,
		Host:            p.Host,
		Port:            p.Port,
		Config:          p.Config,
		LogLevel:        p.LogLevel,
		Dbname:          p.Dbname,
		Username:        p.Username,
		Password:        p.Password,
		MaxIdleConns:    p.MaxIdleConns,
		MaxOpenConns:    p.MaxOpenConns,
		ConnMaxIdleTime: p.ConnMaxIdleTime,
		ConnMaxLifeTime: p.ConnMaxLifeTime,
	}
}

// Get 返回 PostgreSQL 配置的所有字段
func (p *PostgreSQL) Get() interface{} {
	return p
}

// Set 更新 PostgreSQL 配置的字段
func (p *PostgreSQL) Set(data interface{}) {
	if configData, ok := data.(*PostgreSQL); ok {
		p.ModuleName = configData.ModuleName
		p.Host = configData.Host
		p.Port = configData.Port
		p.Config = configData.Config
		p.LogLevel = configData.LogLevel
		p.Dbname = configData.Dbname
		p.Username = configData.Username
		p.Password = configData.Password
		p.MaxIdleConns = configData.MaxIdleConns
		p.MaxOpenConns = configData.MaxOpenConns
		p.ConnMaxIdleTime = configData.ConnMaxIdleTime
		p.ConnMaxLifeTime = configData.ConnMaxLifeTime
	}
}

// Validate 检查 PostgreSQL 配置的有效性
func (p *PostgreSQL) Validate() error {
	return internal.ValidateStruct(p)
}

func (p PostgreSQL) GetCommonConfig() *DBConfig {
	return &DBConfig{
		Host:            p.Host,
		Username:        p.Username,
		Password:        p.Password,
		Dbname:          p.Dbname,
		Port:            p.Port,
		Config:          p.Config,
		MaxIdleConns:    p.MaxIdleConns,
		MaxOpenConns:    p.MaxOpenConns,
		ConnMaxIdleTime: p.ConnMaxIdleTime,
		ConnMaxLifeTime: p.ConnMaxLifeTime,
		LogLevel:        p.LogLevel,
	}
}

// DefaultPostgreSQL 返回默认PostgreSQL配置
func DefaultPostgreSQL() PostgreSQL {
	return PostgreSQL{
		ModuleName:      "postgresql",
		Host:            "127.0.0.1",
		Port:            "5432",
		Config:          "sslmode=disable",
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

// DefaultPostgreSQLConfig 返回默认PostgreSQL配置的指针，支持链式调用
func DefaultPostgreSQLConfig() *PostgreSQL {
	config := DefaultPostgreSQL()
	return &config
}

// WithModuleName 设置模块名称
func (p *PostgreSQL) WithModuleName(moduleName string) *PostgreSQL {
	p.ModuleName = moduleName
	return p
}

// WithHost 设置主机地址
func (p *PostgreSQL) WithHost(host string) *PostgreSQL {
	p.Host = host
	return p
}

// WithPort 设置端口
func (p *PostgreSQL) WithPort(port string) *PostgreSQL {
	p.Port = port
	return p
}

// WithConfig 设置配置字符串
func (p *PostgreSQL) WithConfig(config string) *PostgreSQL {
	p.Config = config
	return p
}

// WithLogLevel 设置日志级别
func (p *PostgreSQL) WithLogLevel(logLevel string) *PostgreSQL {
	p.LogLevel = logLevel
	return p
}

// WithDbname 设置数据库名称
func (p *PostgreSQL) WithDbname(dbname string) *PostgreSQL {
	p.Dbname = dbname
	return p
}

// WithUsername 设置用户名
func (p *PostgreSQL) WithUsername(username string) *PostgreSQL {
	p.Username = username
	return p
}

// WithPassword 设置密码
func (p *PostgreSQL) WithPassword(password string) *PostgreSQL {
	p.Password = password
	return p
}

// WithMaxIdleConns 设置最大空闲连接数
func (p *PostgreSQL) WithMaxIdleConns(maxIdleConns int) *PostgreSQL {
	p.MaxIdleConns = maxIdleConns
	return p
}

// WithMaxOpenConns 设置最大连接数
func (p *PostgreSQL) WithMaxOpenConns(maxOpenConns int) *PostgreSQL {
	p.MaxOpenConns = maxOpenConns
	return p
}

// WithConnMaxIdleTime 设置连接最大空闲时间
func (p *PostgreSQL) WithConnMaxIdleTime(connMaxIdleTime int) *PostgreSQL {
	p.ConnMaxIdleTime = connMaxIdleTime
	return p
}

// WithConnMaxLifeTime 设置连接最大生命周期
func (p *PostgreSQL) WithConnMaxLifeTime(connMaxLifeTime int) *PostgreSQL {
	p.ConnMaxLifeTime = connMaxLifeTime
	return p
}
