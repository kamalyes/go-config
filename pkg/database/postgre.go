/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 10:55:05
 * @FilePath: \go-config\database\postgre.go
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
	ModuleName      string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`            // 模块名称
	Host            string `mapstructure:"host"                     json:"host"                     yaml:"host"`               // 数据库IP地址
	Port            string `mapstructure:"port"                     json:"port"                     yaml:"port"`               // 端口
	Config          string `mapstructure:"config"                   json:"config"                   yaml:"config"`             // 后缀配置
	LogLevel        string `mapstructure:"log-level"                json:"logLevel"                 yaml:"log-level"`          // SQL日志等级
	Dbname          string `mapstructure:"db-name"                  json:"dbname"                   yaml:"db-name"`            // 数据库名称
	Username        string `mapstructure:"username"                 json:"username"                 yaml:"username"`           // 数据库用户名
	Password        string `mapstructure:"password"                 json:"password"                 yaml:"password"`           // 数据库密码
	MaxIdleConns    int    `mapstructure:"max-idle-conns"           json:"maxIdleConns"             yaml:"max-idle-conns"`     // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"max-open-conns"           json:"maxOpenConns"             yaml:"max-open-conns"`     // 最大连接数
	ConnMaxIdleTime int    `mapstructure:"conn-max-idle-time"       json:"connMaxIdleTime"          yaml:"conn-max-idle-time"` // 连接最大空闲时间（秒）
	ConnMaxLifeTime int    `mapstructure:"conn-max-life-time"       json:"connMaxLifeTime"          yaml:"conn-max-life-time"` // 连接最大生命周期（秒）
}

// NewPostgreSQL 创建一个新的 PostgreSQL 实例
func NewPostgreSQL(moduleName, host, port, dbname, username, password string, maxIdleConns, maxOpenConns, connMaxIdleTime, connMaxLifeTime int) *PostgreSQL {
	var postgreInstance *PostgreSQL

	internal.LockFunc(func() {
		postgreInstance = &PostgreSQL{
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
	return postgreInstance
}

// ToMap 将配置转换为映射
func (p *PostgreSQL) ToMap() map[string]interface{} {
	return internal.ToMap(p)
}

// FromMap 从映射中填充配置
func (p *PostgreSQL) FromMap(data map[string]interface{}) {
	internal.FromMap(p, data)
}

// Clone 返回 PostgreSQL 配置的副本
func (p *PostgreSQL) Clone() internal.Configurable {
	return &PostgreSQL{
		ModuleName:      p.ModuleName,
		Host:            p.Host,
		Port:            p.Port,
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
