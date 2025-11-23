/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 23:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 16:53:36
 * @FilePath: \go-config\pkg\database\postgresql.go
 * @Description: PostgreSQL数据库配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package database

import (
	"github.com/kamalyes/go-config/internal"
)

// PostgreSQL PostgreSQL数据库配置
type PostgreSQL struct {
	ModuleName      string `mapstructure:"modulename"          yaml:"modulename"          json:"module_name"`                         // 模块名称
	Host            string `mapstructure:"host"                yaml:"host"                json:"host"            validate:"required"` // 数据库 IP 地址
	Port            string `mapstructure:"port"                yaml:"port"                json:"port"            validate:"required"` // 端口
	Config          string `mapstructure:"config"              yaml:"config"              json:"config"          validate:"required"` // 后缀配置
	LogLevel        string `mapstructure:"log-level"           yaml:"log_level"           json:"log_level"       validate:"required"` // SQL 日志等级
	Dbname          string `mapstructure:"db-name"             yaml:"db_name"             json:"db_name"         validate:"required"` // 数据库名称
	Username        string `mapstructure:"username"            yaml:"username"            json:"username"        validate:"required"` // 数据库用户名
	Password        string `mapstructure:"password"            yaml:"password"            json:"password"        validate:"required"` // 数据库密码
	MaxIdleConns    int    `mapstructure:"max-idle-conns"      yaml:"max_idle_conns"      json:"max_idle_conns"  validate:"min=0"`    // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"max-open-conns"      yaml:"max_open_conns"      json:"max_open_conns"  validate:"min=0"`    // 最大连接数
	ConnMaxIdleTime int    `mapstructure:"conn-max-idle-time"  yaml:"conn_max_idle_time"  json:"conn_max_idle_time" validate:"min=0"` // 连接最大空闲时间 单位：秒
	ConnMaxLifeTime int    `mapstructure:"conn-max-life-time"  yaml:"conn_max_life_time"  json:"conn_max_life_time" validate:"min=0"` // 连接最大生命周期 单位：秒
}

// 为PostgreSQL配置实现DatabaseProvider接口
func (p *PostgreSQL) GetDBType() DBType     { return DBTypePostgreSQL }
func (p *PostgreSQL) GetHost() string       { return p.Host }
func (p *PostgreSQL) GetPort() string       { return p.Port }
func (p *PostgreSQL) GetDBName() string     { return p.Dbname }
func (p *PostgreSQL) GetUsername() string   { return p.Username }
func (p *PostgreSQL) GetPassword() string   { return p.Password }
func (p *PostgreSQL) GetConfig() string     { return p.Config }
func (p *PostgreSQL) GetModuleName() string { return p.ModuleName }
func (p *PostgreSQL) SetCredentials(username, password string) {
	p.Username, p.Password = username, password
}
func (p *PostgreSQL) SetHost(host string)     { p.Host = host }
func (p *PostgreSQL) SetPort(port string)     { p.Port = port }
func (p *PostgreSQL) SetDBName(dbName string) { p.Dbname = dbName }

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
func (p *PostgreSQL) Get() interface{} { return p }
func (p *PostgreSQL) Set(data interface{}) {
	if cfg, ok := data.(*PostgreSQL); ok {
		*p = *cfg
	}
}
func (p *PostgreSQL) Validate() error { return internal.ValidateStruct(p) }

// DefaultPostgreSQL 创建默认PostgreSQL配置
func DefaultPostgreSQL() *PostgreSQL {
	return &PostgreSQL{
		ModuleName:      "postgresql",
		Host:            "localhost",
		Port:            "5432",
		Config:          "sslmode=disable",
		LogLevel:        "info",
		Dbname:          "postgres",
		Username:        "postgres",
		Password:        "postgres_password",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxIdleTime: 300,  // 5分钟
		ConnMaxLifeTime: 3600, // 1小时
	}
}

// NewPostgreSQLConfig 创建一个新的PostgreSQL配置实例
func NewPostgreSQLConfig(opt *PostgreSQL) *PostgreSQL {
	var postgresqlInstance *PostgreSQL
	internal.LockFunc(func() {
		postgresqlInstance = opt
	})
	return postgresqlInstance
}

// WithModuleName 设置模块名称
func (p *PostgreSQL) WithModuleName(moduleName string) *PostgreSQL {
	p.ModuleName = moduleName
	return p
}

// WithHost 设置数据库主机地址
func (p *PostgreSQL) WithHost(host string) *PostgreSQL {
	p.Host = host
	return p
}

// WithPort 设置数据库端口
func (p *PostgreSQL) WithPort(port string) *PostgreSQL {
	p.Port = port
	return p
}

// WithConfig 设置数据库配置参数
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

// WithUsername 设置数据库用户名
func (p *PostgreSQL) WithUsername(username string) *PostgreSQL {
	p.Username = username
	return p
}

// WithPassword 设置数据库密码
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
