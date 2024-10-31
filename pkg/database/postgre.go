/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 17:20:57
 * @FilePath: \go-config\pkg\database\postgre.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package database

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// PostgreSQL 配置结构体
type PostgreSQL struct {
	ModuleName      string `mapstructure:"MODULE_NAME"              yaml:"modulename"`         // 模块名称
	Host            string `mapstructure:"HOST"                     yaml:"host"`               // 数据库IP地址
	Port            string `mapstructure:"PORT"                     yaml:"port"`               // 端口
	Config          string `mapstructure:"CONFIG"                   yaml:"config"`             // 后缀配置
	LogLevel        string `mapstructure:"LOG_LEVEL"                yaml:"log-level"`          // SQL日志等级
	Dbname          string `mapstructure:"DB_NAME"                  yaml:"db-name"`            // 数据库名称
	Username        string `mapstructure:"USERNAME"                 yaml:"username"`           // 数据库用户名
	Password        string `mapstructure:"PASSWORD"                 yaml:"password"`           // 数据库密码
	MaxIdleConns    int    `mapstructure:"MAX_IDLE_CONNS"           yaml:"max-idle-conns"`     // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"MAX_OPEN_CONNS"           yaml:"max-open-conns"`     // 最大连接数
	ConnMaxIdleTime int    `mapstructure:"CONN_MAX_IDLE_TIME"       yaml:"conn-max-idle-time"` // 连接最大空闲时间（秒）
	ConnMaxLifeTime int    `mapstructure:"CONN_MAX_LIFE_TIME"       yaml:"conn-max-life-time"` // 连接最大生命周期（秒）
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
	if p.Host == "" {
		return errors.New("host cannot be empty")
	}
	if p.Port == "" {
		return errors.New("port cannot be empty")
	}
	if p.Dbname == "" {
		return errors.New("dbname cannot be empty")
	}
	if p.Username == "" {
		return errors.New("username cannot be empty")
	}
	if p.MaxIdleConns < 0 {
		return errors.New("MaxIdleConns cannot be negative")
	}
	if p.MaxOpenConns < 0 {
		return errors.New("MaxOpenConns cannot be negative")
	}
	if p.ConnMaxIdleTime < 0 {
		return errors.New("ConnMaxIdleTime cannot be negative")
	}
	if p.ConnMaxLifeTime < 0 {
		return errors.New("ConnMaxLifeTime cannot be negative")
	}
	return nil
}
