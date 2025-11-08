/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:20:10
 * @FilePath: \go-config\pkg\database\sqlite.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package database

import (
	"github.com/kamalyes/go-config/internal"
)

// SQLite 数据库配置结构体
type SQLite struct {
	DbPath          string `mapstructure:"db-path"                  yaml:"db-path"             json:"db_path"          validate:"required"` // SQLite 文件存放位置
	MaxIdleConns    int    `mapstructure:"max-idle-conns"           yaml:"max-idle-conns"      json:"max_idle_conns"   validate:"min=0"`    // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"max-open-conns"           yaml:"max-open-conns"      json:"max_open_conns"   validate:"min=0"`    // 最大连接数
	LogLevel        string `mapstructure:"log-level"                yaml:"log-level"           json:"log_level"        validate:"required"` // SQL 日志等级
	ConnMaxIdleTime int    `mapstructure:"conn-max-idle-time"       yaml:"conn-max-idle-time"  json:"conn_max_idle_time" validate:"min=0"`  // 连接最大空闲时间（秒）
	ConnMaxLifeTime int    `mapstructure:"conn-max-life-time"       yaml:"conn-max-life-time"  json:"conn_max_life_time" validate:"min=0"`  // 连接最大生命周期（秒）
	Vacuum          bool   `mapstructure:"vacuum"                   yaml:"vacuum"              json:"vacuum"`                               // 是否执行清除命令
	ModuleName      string `mapstructure:"modulename"               yaml:"modulename"          json:"module_name"`                          // 模块名称
}

// NewSQLite 创建一个新的 SQLite 实例
func NewSQLite(opt *SQLite) *SQLite {
	var sqliteInstance *SQLite

	internal.LockFunc(func() {
		sqliteInstance = opt
	})
	return sqliteInstance
}

// Clone 返回 SQLite 配置的副本
func (s *SQLite) Clone() internal.Configurable {
	return &SQLite{
		ModuleName:      s.ModuleName,
		DbPath:          s.DbPath,
		MaxIdleConns:    s.MaxIdleConns,
		MaxOpenConns:    s.MaxOpenConns,
		LogLevel:        s.LogLevel,
		Vacuum:          s.Vacuum,
		ConnMaxIdleTime: s.ConnMaxIdleTime,
		ConnMaxLifeTime: s.ConnMaxLifeTime,
	}
}

// Get 返回 SQLite 配置的所有字段
func (s *SQLite) Get() interface{} {
	return s
}

// Set 更新 SQLite 配置的字段
func (s *SQLite) Set(data interface{}) {
	if configData, ok := data.(*SQLite); ok {
		s.ModuleName = configData.ModuleName
		s.DbPath = configData.DbPath
		s.MaxIdleConns = configData.MaxIdleConns
		s.MaxOpenConns = configData.MaxOpenConns
		s.LogLevel = configData.LogLevel
		s.Vacuum = configData.Vacuum
		s.ConnMaxIdleTime = configData.ConnMaxIdleTime
		s.ConnMaxLifeTime = configData.ConnMaxLifeTime
	}
}

// Validate 检查 SQLite 配置的有效性
func (s *SQLite) Validate() error {
	return internal.ValidateStruct(s)
}

func (s SQLite) GetCommonConfig() *DBConfig {
	return &DBConfig{
		DbPath:          s.DbPath,
		MaxIdleConns:    s.MaxIdleConns,
		MaxOpenConns:    s.MaxOpenConns,
		ConnMaxIdleTime: s.ConnMaxIdleTime,
		ConnMaxLifeTime: s.ConnMaxLifeTime,
		LogLevel:        s.LogLevel,
		Vacuum:          s.Vacuum,
	}
}

// DefaultSQLite 返回默认SQLite配置
func DefaultSQLite() SQLite {
	return SQLite{
		ModuleName:      "sqlite",
		DbPath:          "./data.db",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		LogLevel:        "silent",
		ConnMaxIdleTime: 3600,  // 1小时
		ConnMaxLifeTime: 7200,  // 2小时
		Vacuum:          false,
	}
}

// DefaultSQLiteConfig 返回默认SQLite配置的指针，支持链式调用
func DefaultSQLiteConfig() *SQLite {
	config := DefaultSQLite()
	return &config
}

// WithModuleName 设置模块名称
func (s *SQLite) WithModuleName(moduleName string) *SQLite {
	s.ModuleName = moduleName
	return s
}

// WithDbPath 设置数据库文件路径
func (s *SQLite) WithDbPath(dbPath string) *SQLite {
	s.DbPath = dbPath
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

// WithLogLevel 设置日志级别
func (s *SQLite) WithLogLevel(logLevel string) *SQLite {
	s.LogLevel = logLevel
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

// WithVacuum 设置是否执行清除命令
func (s *SQLite) WithVacuum(vacuum bool) *SQLite {
	s.Vacuum = vacuum
	return s
}
