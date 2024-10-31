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
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// SQLite 配置结构体
type SQLite struct {
	ModuleName      string `mapstructure:"MODULE_NAME"              yaml:"modulename"`         // 模块名称
	DbPath          string `mapstructure:"DB_PATH"                  yaml:"db-path"`            // SQLite 文件存放位置
	MaxIdleConns    int    `mapstructure:"MAX_IDLE_CONNS"           yaml:"max-idle-conns"`     // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"MAX_OPEN_CONNS"           yaml:"max-open-conns"`     // 最大连接数
	LogLevel        string `mapstructure:"LOG_LEVEL"                yaml:"log-level"`          // SQL日志等级
	Vacuum          bool   `mapstructure:"VACUUM"                   yaml:"vacuum"`             // 是否执行清除命令
	ConnMaxIdleTime int    `mapstructure:"CONN_MAX_IDLE_TIME"       yaml:"conn-max-idle-time"` // 连接最大空闲时间（秒）
	ConnMaxLifeTime int    `mapstructure:"CONN_MAX_LIFE_TIME"       yaml:"conn-max-life-time"` // 连接最大生命周期（秒）
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
	if s.DbPath == "" {
		return errors.New("DbPath cannot be empty")
	}
	if s.MaxIdleConns < 0 {
		return errors.New("MaxIdleConns cannot be negative")
	}
	if s.MaxOpenConns < 0 {
		return errors.New("MaxOpenConns cannot be negative")
	}
	if s.ConnMaxIdleTime < 0 {
		return errors.New("ConnMaxIdleTime cannot be negative")
	}
	if s.ConnMaxLifeTime < 0 {
		return errors.New("ConnMaxLifeTime cannot be negative")
	}
	return nil
}
