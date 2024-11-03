/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-03 20:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 21:15:56
 * @FilePath: \go-config\pkg\database\common.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package database

import "fmt"

// 数据库类型常量
const (
	DBTypeMySQL    = "mysql"
	DBTypePostgres = "postgres"
	DBTypeSQLite   = "sqlite"
)

// DBConfig 数据库配置结构体
type DBConfig struct {
	ModuleName      string `mapstructure:"modulename"          yaml:"modulename"          json:"module_name"     validate:"required"`  // 模块名称
	Host            string `mapstructure:"host"                yaml:"host"                json:"host"            validate:"required"`  // 数据库 IP 地址
	Port            string `mapstructure:"port"                yaml:"port"                json:"port"            validate:"required"`  // 端口
	Config          string `mapstructure:"config"              yaml:"config"              json:"config"          validate:"required"`  // 后缀配置 默认配置 charset=utf8mb4&parseTime=True&loc=Local
	LogLevel        string `mapstructure:"log-level"           yaml:"log-level"           json:"log_level"       validate:"required"`  // SQL 日志等级
	Dbname          string `mapstructure:"db-name"             yaml:"db-name"             json:"db_name"         validate:"required"`  // 数据库名称
	Username        string `mapstructure:"username"            yaml:"username"            json:"username"        validate:"required"`  // 数据库用户名
	Password        string `mapstructure:"password"            yaml:"password"            json:"password"        validate:"required"`  // 数据库密码
	MaxIdleConns    int    `mapstructure:"max-idle-conns"      yaml:"max-idle-conns"      json:"max_idle_conns"  validate:"min=0"`     // 最大空闲连接数
	MaxOpenConns    int    `mapstructure:"max-open-conns"      yaml:"max-open-conns"      json:"max_open_conns"  validate:"min=0"`     // 最大连接数
	ConnMaxIdleTime int    `mapstructure:"conn-max-idle-time"  yaml:"conn-max-idle-time"  json:"conn_max_idle_time" validate:"min=0"`  // 连接最大空闲时间 单位：秒
	ConnMaxLifeTime int    `mapstructure:"conn-max-life-time"  yaml:"conn-max-life-time"  json:"conn_max_life_time" validate:"min=0"`  // 连接最大生命周期 单位：秒
	DbPath          string `mapstructure:"db-path"             yaml:"db-path"             json:"db_path"          validate:"required"` // SQLite 文件存放位置
	Vacuum          bool   `mapstructure:"vacuum"              yaml:"vacuum"              json:"vacuum"`                               // 是否执行清除命令
}

// DBConfigInterface 定义一个接口以统一不同数据库的配置
type DBConfigInterface interface {
	GetCommonConfig() *DBConfig // 修改为返回指针
}

// NewDBConfig 创建数据库配置实例
func NewDBConfig(dbType string, configMap DBConfigInterface) (*DBConfig, error) {
	if configMap == nil {
		return nil, fmt.Errorf("configMap cannot be nil")
	}

	switch dbType {
	case DBTypeMySQL:
		mysqlConfig, ok := configMap.(MySQL)
		if !ok {
			return nil, fmt.Errorf("invalid config for MySQL")
		}
		return mysqlConfig.GetCommonConfig(), nil // 直接返回指针

	case DBTypePostgres:
		postgresConfig, ok := configMap.(PostgreSQL)
		if !ok {
			return nil, fmt.Errorf("invalid config for PostgreSQL")
		}
		return postgresConfig.GetCommonConfig(), nil // 直接返回指针

	case DBTypeSQLite:
		sqliteConfig, ok := configMap.(SQLite)
		if !ok {
			return nil, fmt.Errorf("invalid config for SQLite")
		}
		return sqliteConfig.GetCommonConfig(), nil // 直接返回指针

	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
