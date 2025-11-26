/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 23:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-16 19:03:40
 * @FilePath: \go-config\pkg\database\sqlite.go
 * @Description: SQLite数据库配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package database

import (
	"github.com/kamalyes/go-config/internal"
)

// SQLite SQLite数据库配置
type SQLite struct {
	ModuleName      string `mapstructure:"module_name" yaml:"module_name" json:"module_name"`                                       // 模块名称
	Config          string `mapstructure:"config" yaml:"config" json:"config"          validate:"required"`                         // 后缀配置
	LogLevel        string `mapstructure:"log_level" yaml:"log-level" json:"log_level"       validate:"required"`                   // SQL 日志等级
	DbPath          string `mapstructure:"db_path" yaml:"db-path" json:"db_path"`                                                   // SQLite 文件存放位置
	Vacuum          bool   `mapstructure:"vacuum" yaml:"vacuum" json:"vacuum"`                                                      // 是否执行清除命令
	MaxIdleConns    int    `mapstructure:"max_idle_conns" yaml:"max-idle-conns" json:"max_idle_conns"  validate:"min=0"`            // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"max_open_conns" yaml:"max-open-conns" json:"max_open_conns"  validate:"min=0"`            // 最大连接数
	ConnMaxIdleTime int    `mapstructure:"conn_max_idle_time" yaml:"conn-max-idle-time" json:"conn_max_idle_time" validate:"min=0"` // 连接最大空闲时间 单位：秒
	ConnMaxLifeTime int    `mapstructure:"conn_max_life_time" yaml:"conn-max-life-time" json:"conn_max_life_time" validate:"min=0"` // 连接最大生命周期 单位：秒
}

// 为SQLite配置实现DatabaseProvider接口
func (s *SQLite) GetDBType() DBType                        { return DBTypeSQLite }
func (s *SQLite) GetHost() string                          { return "" }       // SQLite无需host
func (s *SQLite) GetPort() string                          { return "" }       // SQLite无需port
func (s *SQLite) GetDBName() string                        { return s.DbPath } // SQLite使用文件路径
func (s *SQLite) GetUsername() string                      { return "" }       // SQLite无需用户名
func (s *SQLite) GetPassword() string                      { return "" }       // SQLite无需密码
func (s *SQLite) GetConfig() string                        { return s.Config }
func (s *SQLite) GetModuleName() string                    { return s.ModuleName }
func (s *SQLite) SetCredentials(username, password string) {} // SQLite不支持凭证
func (s *SQLite) SetHost(host string)                      {} // SQLite不支持host
func (s *SQLite) SetPort(port string)                      {} // SQLite不支持port
func (s *SQLite) SetDBName(dbName string)                  { s.DbPath = dbName }

func (s *SQLite) Clone() internal.Configurable {
	return &SQLite{
		ModuleName:      s.ModuleName,
		Config:          s.Config,
		LogLevel:        s.LogLevel,
		DbPath:          s.DbPath,
		Vacuum:          s.Vacuum,
		MaxIdleConns:    s.MaxIdleConns,
		MaxOpenConns:    s.MaxOpenConns,
		ConnMaxIdleTime: s.ConnMaxIdleTime,
		ConnMaxLifeTime: s.ConnMaxLifeTime,
	}
}
func (s *SQLite) Get() interface{} { return s }
func (s *SQLite) Set(data interface{}) {
	if cfg, ok := data.(*SQLite); ok {
		*s = *cfg
	}
}
func (s *SQLite) Validate() error { return internal.ValidateStruct(s) }

// DefaultSQLite 创建默认SQLite配置
func DefaultSQLite() *SQLite {
	return &SQLite{
		ModuleName:      "sqlite",
		DbPath:          "./data.db",
		Config:          "_foreign_keys=on",
		LogLevel:        "info",
		Vacuum:          false,
		MaxIdleConns:    10,
		MaxOpenConns:    1, // SQLite 通常使用单连接
		ConnMaxIdleTime: 300,
		ConnMaxLifeTime: 3600,
	}
}

// NewSQLiteConfig 创建一个新的SQLite配置实例
func NewSQLiteConfig(opt *SQLite) *SQLite {
	var sqliteInstance *SQLite
	internal.LockFunc(func() {
		sqliteInstance = opt
	})
	return sqliteInstance
}

// WithModuleName 设置模块名称
func (s *SQLite) WithModuleName(moduleName string) *SQLite {
	s.ModuleName = moduleName
	return s
}

// WithConfig 设置数据库配置参数
func (s *SQLite) WithConfig(config string) *SQLite {
	s.Config = config
	return s
}

// WithLogLevel 设置日志级别
func (s *SQLite) WithLogLevel(logLevel string) *SQLite {
	s.LogLevel = logLevel
	return s
}

// WithDbPath 设置SQLite数据库文件路径
func (s *SQLite) WithDbPath(dbPath string) *SQLite {
	s.DbPath = dbPath
	return s
}

// WithVacuum 设置是否执行清除命令
func (s *SQLite) WithVacuum(vacuum bool) *SQLite {
	s.Vacuum = vacuum
	return s
}

// WithMaxIdleConns 设置最大空闲连接数
func (s *SQLite) WithMaxIdleConns(maxIdleConns int) *SQLite {
	s.MaxIdleConns = maxIdleConns
	return s
}

// WithMaxOpenConns 设置最大连接数
func (s *SQLite) WithMaxOpenConns(maxOpenConns int) *SQLite {
	s.MaxOpenConns = maxOpenConns
	return s
}

// WithConnMaxIdleTime 设置连接最大空闲时间
func (s *SQLite) WithConnMaxIdleTime(connMaxIdleTime int) *SQLite {
	s.ConnMaxIdleTime = connMaxIdleTime
	return s
}

// WithConnMaxLifeTime 设置连接最大生命周期
func (s *SQLite) WithConnMaxLifeTime(connMaxLifeTime int) *SQLite {
	s.ConnMaxLifeTime = connMaxLifeTime
	return s
}
