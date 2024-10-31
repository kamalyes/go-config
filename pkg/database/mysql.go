/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 17:31:09
 * @FilePath: \go-config\pkg\database\mysql.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package database

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// MySQL 配置
type MySQL struct {
	ModuleName      string `mapstructure:"MODULE_NAME"              yaml:"modulename"`         // 模块名称
	Host            string `mapstructure:"HOST"                     yaml:"host"`               // 数据库IP地址
	Port            string `mapstructure:"PORT"                     yaml:"port"`               // 端口
	Config          string `mapstructure:"CONFIG"                   yaml:"config"`             // 后缀配置  默认配置 charset=utf8mb4&parseTime=True&loc=Local
	LogLevel        string `mapstructure:"LOG_LEVEL"                yaml:"log-level"`          // sql日志等级
	Dbname          string `mapstructure:"DB_NAME"                  yaml:"db-name"`            // 数据库名称
	Username        string `mapstructure:"USERNAME"                 yaml:"username"`           // 数据库用户名
	Password        string `mapstructure:"PASSWORD"                 yaml:"password"`           // 数据库密码
	MaxIdleConns    int    `mapstructure:"MAX_IDLE_CONNS"           yaml:"max-idle-conns"`     // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"MAX_OPEN_CONNS"           yaml:"max-open-conns"`     // 最大连接数
	ConnMaxIdleTime int    `mapstructure:"CONN_MAX_IDLE_TIME"       yaml:"conn-max-idle-time"` // 连接最大空闲时间  单位：秒
	ConnMaxLifeTime int    `mapstructure:"CONN_MAX_LIFE_TIME"       yaml:"conn-max-life-time"` // 连接最大生命周期  单位：秒
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
	if m.Host == "" {
		return errors.New("Host cannot be empty")
	}
	if m.Port == "" {
		return errors.New("Port cannot be empty")
	}
	if m.Dbname == "" {
		return errors.New("Dbname cannot be empty")
	}
	if m.Username == "" {
		return errors.New("Username cannot be empty")
	}
	if m.MaxIdleConns < 0 {
		return errors.New("MaxIdleConns cannot be negative")
	}
	if m.MaxOpenConns < 0 {
		return errors.New("MaxOpenConns cannot be negative")
	}
	if m.ConnMaxIdleTime < 0 {
		return errors.New("ConnMaxIdleTime cannot be negative")
	}
	if m.ConnMaxLifeTime < 0 {
		return errors.New("ConnMaxLifeTime cannot be negative")
	}
	return nil
}
