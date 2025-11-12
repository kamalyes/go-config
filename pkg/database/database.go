/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-03 20:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 10:43:32
 * @FilePath: \engine-im-push-service\go-config\pkg\database\database.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package database

import (
	"github.com/kamalyes/go-config/internal"
)

// BaseConfig 数据库基础配置结构体
type BaseConfig struct {
	Host            string `mapstructure:"host"                yaml:"host"                json:"host"            validate:"required"` // 数据库 IP 地址
	Port            string `mapstructure:"port"                yaml:"port"                json:"port"            validate:"required"` // 端口
	Config          string `mapstructure:"config"              yaml:"config"              json:"config"          validate:"required"` // 后缀配置 默认配置 charset=utf8mb4&parseTime=True&loc=Local
	LogLevel        string `mapstructure:"log-level"           yaml:"log-level"           json:"log_level"       validate:"required"` // SQL 日志等级
	Dbname          string `mapstructure:"db-name"             yaml:"db-name"             json:"db_name"         validate:"required"` // 数据库名称
	Username        string `mapstructure:"username"            yaml:"username"            json:"username"        validate:"required"` // 数据库用户名
	Password        string `mapstructure:"password"            yaml:"password"            json:"password"        validate:"required"` // 数据库密码
	MaxIdleConns    int    `mapstructure:"max-idle-conns"      yaml:"max-idle-conns"      json:"max_idle_conns"  validate:"min=0"`    // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"max-open-conns"      yaml:"max-open-conns"      json:"max_open_conns"  validate:"min=0"`    // 最大连接数
	ConnMaxIdleTime int    `mapstructure:"conn-max-idle-time"  yaml:"conn-max-idle-time"  json:"conn_max_idle_time" validate:"min=0"` // 连接最大空闲时间 单位：秒
	ConnMaxLifeTime int    `mapstructure:"conn-max-life-time"  yaml:"conn-max-life-time"  json:"conn_max_life_time" validate:"min=0"` // 连接最大生命周期 单位：秒
	DbPath          string `mapstructure:"db-path"             yaml:"db-path"             json:"db_path"`                             // SQLite 文件存放位置
	Vacuum          bool   `mapstructure:"vacuum"              yaml:"vacuum"              json:"vacuum"`                              // 是否执行清除命令
	ModuleName      string `mapstructure:"modulename"          yaml:"modulename"          json:"module_name"`                         // 模块名称
}

// Database 数据库配置
type Database struct {
	*BaseConfig
	DBType string `mapstructure:"db-type"  yaml:"db-type"  json:"db_type"  validate:"required,oneof=mysql postgres sqlite"` // 数据库类型
}

// NewDatabase 创建一个新的 Database 实例
func NewDatabase(opt *Database) *Database {
	var dbInstance *Database

	internal.LockFunc(func() {
		dbInstance = opt
	})
	return dbInstance
}

// Clone 返回 Database 配置的副本
func (d *Database) Clone() internal.Configurable {
	return &Database{
		BaseConfig: &BaseConfig{
			Host:            d.Host,
			Port:            d.Port,
			Config:          d.Config,
			LogLevel:        d.LogLevel,
			Dbname:          d.Dbname,
			Username:        d.Username,
			Password:        d.Password,
			MaxIdleConns:    d.MaxIdleConns,
			MaxOpenConns:    d.MaxOpenConns,
			ConnMaxIdleTime: d.ConnMaxIdleTime,
			ConnMaxLifeTime: d.ConnMaxLifeTime,
			DbPath:          d.DbPath,
			Vacuum:          d.Vacuum,
			ModuleName:      d.ModuleName,
		},
		DBType: d.DBType,
	}
}

// Get 返回 Database 配置的所有字段
func (d *Database) Get() interface{} {
	return d
}

// Set 更新 Database 配置的字段
func (d *Database) Set(data interface{}) {
	if configData, ok := data.(*Database); ok {
		d.Host = configData.Host
		d.Port = configData.Port
		d.Config = configData.Config
		d.LogLevel = configData.LogLevel
		d.Dbname = configData.Dbname
		d.Username = configData.Username
		d.Password = configData.Password
		d.MaxIdleConns = configData.MaxIdleConns
		d.MaxOpenConns = configData.MaxOpenConns
		d.ConnMaxIdleTime = configData.ConnMaxIdleTime
		d.ConnMaxLifeTime = configData.ConnMaxLifeTime
		d.DbPath = configData.DbPath
		d.Vacuum = configData.Vacuum
		d.ModuleName = configData.ModuleName
		d.DBType = configData.DBType
	}
}

// Validate 检查 Database 配置的有效性
func (d *Database) Validate() error {
	return internal.ValidateStruct(d)
}

// DefaultDatabase 返回默认Database配置
func DefaultDatabase() Database {
	return Database{
		BaseConfig: &BaseConfig{
			Host:            "localhost",
			Port:            "3306",
			Config:          "charset=utf8mb4&parseTime=True&loc=Local",
			LogLevel:        "info",
			Dbname:          "test",
			Username:        "root",
			Password:        "",
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			ConnMaxIdleTime: 600,
			ConnMaxLifeTime: 3600,
			DbPath:          "./data/app.db",
			Vacuum:          false,
			ModuleName:      "database",
		},
		DBType: "mysql",
	}
}

// Default 返回默认Database配置的指针，支持链式调用
func Default() *Database {
	config := DefaultDatabase()
	return &config
}

// WithModuleName 设置模块名称
func (d *Database) WithModuleName(moduleName string) *Database {
	d.ModuleName = moduleName
	return d
}

// WithDBType 设置数据库类型
func (d *Database) WithDBType(dbType string) *Database {
	d.DBType = dbType
	return d
}

// WithHost 设置数据库主机地址
func (d *Database) WithHost(host string) *Database {
	d.Host = host
	return d
}

// WithPort 设置数据库端口
func (d *Database) WithPort(port string) *Database {
	d.Port = port
	return d
}

// WithConfig 设置数据库配置参数
func (d *Database) WithConfig(config string) *Database {
	d.Config = config
	return d
}

// WithLogLevel 设置日志级别
func (d *Database) WithLogLevel(logLevel string) *Database {
	d.LogLevel = logLevel
	return d
}

// WithDbname 设置数据库名称
func (d *Database) WithDbname(dbname string) *Database {
	d.Dbname = dbname
	return d
}

// WithUsername 设置数据库用户名
func (d *Database) WithUsername(username string) *Database {
	d.Username = username
	return d
}

// WithPassword 设置数据库密码
func (d *Database) WithPassword(password string) *Database {
	d.Password = password
	return d
}

// WithMaxIdleConns 设置最大空闲连接数
func (d *Database) WithMaxIdleConns(maxIdleConns int) *Database {
	d.MaxIdleConns = maxIdleConns
	return d
}

// WithMaxOpenConns 设置最大连接数
func (d *Database) WithMaxOpenConns(maxOpenConns int) *Database {
	d.MaxOpenConns = maxOpenConns
	return d
}

// WithConnMaxIdleTime 设置连接最大空闲时间
func (d *Database) WithConnMaxIdleTime(connMaxIdleTime int) *Database {
	d.ConnMaxIdleTime = connMaxIdleTime
	return d
}

// WithConnMaxLifeTime 设置连接最大生命周期
func (d *Database) WithConnMaxLifeTime(connMaxLifeTime int) *Database {
	d.ConnMaxLifeTime = connMaxLifeTime
	return d
}

// WithDbPath 设置SQLite数据库文件路径
func (d *Database) WithDbPath(dbPath string) *Database {
	d.DbPath = dbPath
	return d
}

// WithVacuum 设置是否执行清除命令
func (d *Database) WithVacuum(vacuum bool) *Database {
	d.Vacuum = vacuum
	return d
}
