/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 10:40:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 10:40:00
 * @FilePath: \engine-im-push-service\go-config\pkg\database\utils.go
 * @Description: Database utility functions
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package database

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// GetDSN 根据数据库类型生成连接字符串
func (d *Database) GetDSN() string {
	switch strings.ToLower(d.DBType) {
	case "mysql":
		return d.getMySQLDSN()
	case "postgres", "postgresql":
		return d.getPostgresDSN()
	case "sqlite":
		return d.getSQLiteDSN()
	default:
		return ""
	}
}

// getMySQLDSN 生成MySQL连接字符串
func (d *Database) getMySQLDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Dbname,
		d.Config,
	)
}

// getPostgresDSN 生成PostgreSQL连接字符串
func (d *Database) getPostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s",
		d.Host,
		d.Port,
		d.Username,
		d.Password,
		d.Dbname,
		d.Config,
	)
}

// getSQLiteDSN 生成SQLite连接字符串
func (d *Database) getSQLiteDSN() string {
	if d.Config != "" {
		return fmt.Sprintf("%s?%s", d.DbPath, d.Config)
	}
	return d.DbPath
}

// GetConnMaxIdleTimeDuration 获取连接最大空闲时间的 Duration
func (d *Database) GetConnMaxIdleTimeDuration() time.Duration {
	return time.Duration(d.ConnMaxIdleTime) * time.Second
}

// GetConnMaxLifeTimeDuration 获取连接最大生命周期的 Duration
func (d *Database) GetConnMaxLifeTimeDuration() time.Duration {
	return time.Duration(d.ConnMaxLifeTime) * time.Second
}

// IsMySQL 检查是否是MySQL数据库
func (d *Database) IsMySQL() bool {
	return strings.ToLower(d.DBType) == "mysql"
}

// IsPostgreSQL 检查是否是PostgreSQL数据库
func (d *Database) IsPostgreSQL() bool {
	dbType := strings.ToLower(d.DBType)
	return dbType == "postgres" || dbType == "postgresql"
}

// IsSQLite 检查是否是SQLite数据库
func (d *Database) IsSQLite() bool {
	return strings.ToLower(d.DBType) == "sqlite"
}

// GetPortAsInt 获取端口号的整数值
func (d *Database) GetPortAsInt() (int, error) {
	return strconv.Atoi(d.Port)
}

// ValidateConnections 验证连接参数的合理性
func (d *Database) ValidateConnections() error {
	if d.MaxIdleConns < 0 {
		return fmt.Errorf("max idle connections cannot be negative")
	}

	if d.MaxOpenConns < 0 {
		return fmt.Errorf("max open connections cannot be negative")
	}

	if d.MaxIdleConns > d.MaxOpenConns && d.MaxOpenConns > 0 {
		return fmt.Errorf("max idle connections (%d) cannot be greater than max open connections (%d)",
			d.MaxIdleConns, d.MaxOpenConns)
	}

	if d.ConnMaxIdleTime < 0 {
		return fmt.Errorf("connection max idle time cannot be negative")
	}

	if d.ConnMaxLifeTime < 0 {
		return fmt.Errorf("connection max life time cannot be negative")
	}

	return nil
}

// GetConnectionInfo 获取连接信息摘要
func (d *Database) GetConnectionInfo() map[string]interface{} {
	info := map[string]interface{}{
		"db_type":            d.DBType,
		"module_name":        d.ModuleName,
		"log_level":          d.LogLevel,
		"max_idle_conns":     d.MaxIdleConns,
		"max_open_conns":     d.MaxOpenConns,
		"conn_max_idle_time": d.ConnMaxIdleTime,
		"conn_max_life_time": d.ConnMaxLifeTime,
	}

	if !d.IsSQLite() {
		info["host"] = d.Host
		info["port"] = d.Port
		info["database"] = d.Dbname
		info["username"] = d.Username
		// 注意：不包含密码信息
	} else {
		info["db_path"] = d.DbPath
		info["vacuum"] = d.Vacuum
	}

	return info
}

// DefaultMySQL 返回默认MySQL配置
func DefaultMySQL() *Database {
	return Default().WithDBType("mysql").WithModuleName("mysql")
}

// DefaultPostgreSQL 返回默认PostgreSQL配置
func DefaultPostgreSQL() *Database {
	return Default().
		WithDBType("postgres").
		WithPort("5432").
		WithUsername("postgres").
		WithConfig("sslmode=disable&TimeZone=Asia/Shanghai").
		WithModuleName("postgres")
}

// DefaultSQLite 返回默认SQLite配置
func DefaultSQLite() *Database {
	return Default().
		WithDBType("sqlite").
		WithDbPath("./data/app.db").
		WithConfig("cache=shared&mode=rwc&_journal_mode=WAL").
		WithVacuum(false).
		WithModuleName("sqlite")
}
