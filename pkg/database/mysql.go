/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 10:15:55
 * @FilePath: \go-config\database\mysql.go
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
	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** 数据库IP地址 */
	Host string `mapstructure:"host"                     json:"host"                     yaml:"host"`

	/** 端口 */
	Port string `mapstructure:"port"                     json:"port"                     yaml:"port"`

	/** 后缀配置  默认配置 charset=utf8mb4&parseTime=True&loc=Local */
	Config string `mapstructure:"config"                   json:"config"                   yaml:"config"`

	/** sql日志等级 */
	LogLevel string `mapstructure:"log-level"                json:"logLevel"                 yaml:"log-level"`

	/** 数据库名称 */
	Dbname string `mapstructure:"db-name"                  json:"dbname"                   yaml:"db-name"`

	/** 数据库用户名 */
	Username string `mapstructure:"username"                 json:"username"                 yaml:"username"`

	/** 数据库密码 */
	Password string `mapstructure:"password"                 json:"password"                 yaml:"password"`

	/** 最大空闲连接数 */
	MaxIdleConns int `mapstructure:"max-idle-conns"           json:"maxIdleConns"             yaml:"max-idle-conns"`

	/** 最大连接数 */
	MaxOpenConns int `mapstructure:"max-open-conns"           json:"maxOpenConns"             yaml:"max-open-conns"`

	/** 连接最大空闲时间  单位：秒 */
	ConnMaxIdleTime int `mapstructure:"conn-max-idle-time"       json:"connMaxIdleTime"          yaml:"conn-max-idle-time"`

	/** 连接最大生命周期  单位：秒 */
	ConnMaxLifeTime int `mapstructure:"conn-max-life-time"       json:"connMaxLifeTime"          yaml:"conn-max-life-time"`
}

// NewMySQL 创建一个新的 MySQL 实例
func NewMySQL(moduleName, host, port, dbname, username, password string, maxIdleConns, maxOpenConns, connMaxIdleTime, connMaxLifeTime int) *MySQL {
	var mysqlInstance *MySQL

	internal.LockFunc(func() {
		mysqlInstance = &MySQL{
			ModuleName:      moduleName,
			Host:            host,
			Port:            port,
			Dbname:          dbname,
			Username:        username,
			Password:        password,
			MaxIdleConns:    maxIdleConns,
			MaxOpenConns:    maxOpenConns,
			ConnMaxIdleTime: connMaxIdleTime,
			ConnMaxLifeTime: connMaxLifeTime,
		}
	})
	return mysqlInstance
}

// ToMap 将配置转换为映射
func (m *MySQL) ToMap() map[string]interface{} {
	return internal.ToMap(m)
}

// FromMap 从映射中填充配置
func (m *MySQL) FromMap(data map[string]interface{}) {
	internal.FromMap(m, data)
}

// Clone 返回 MySQL 配置的副本
func (m *MySQL) Clone() internal.Configurable {
	return &MySQL{
		ModuleName:      m.ModuleName,
		Host:            m.Host,
		Port:            m.Port,
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
