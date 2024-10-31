/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:30:25
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
	ModuleName      string `mapstructure:"modulename"              yaml:"modulename"          json:"module_name"      validate:"required"`  // 模块名称
	DbPath          string `mapstructure:"db-path"                  yaml:"db-path"             json:"db_path"          validate:"required"` // SQLite 文件存放位置
	MaxIdleConns    int    `mapstructure:"max-idle-conns"           yaml:"max-idle-conns"      json:"max_idle_conns"   validate:"min=0"`    // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"max-open-conns"           yaml:"max-open-conns"      json:"max_open_conns"   validate:"min=0"`    // 最大连接数
	LogLevel        string `mapstructure:"log-level"                yaml:"log-level"           json:"log_level"        validate:"required"` // SQL 日志等级
	ConnMaxIdleTime int    `mapstructure:"conn-max-idle-time"       yaml:"conn-max-idle-time"  json:"conn_max_idle_time" validate:"min=0"`  // 连接最大空闲时间（秒）
	ConnMaxLifeTime int    `mapstructure:"conn-max-life-time"       yaml:"conn-max-life-time"  json:"conn_max_life_time" validate:"min=0"`  // 连接最大生命周期（秒）
	Vacuum          bool   `mapstructure:"vacuum"                   yaml:"vacuum"              json:"vacuum"`                               // 是否执行清除命令
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
