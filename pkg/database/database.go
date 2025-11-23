/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-03 20:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 22:50:17
 * @FilePath: \go-config\pkg\database\database.go
 * @Description: 数据库统一配置管理
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package database

import (
	"fmt"
	"strings"

	"github.com/kamalyes/go-config/internal"
)

// DBType 定义数据库类型
type DBType string

const (
	DBTypeMySQL      DBType = "mysql"
	DBTypePostgreSQL DBType = "postgres"
	DBTypeSQLite     DBType = "sqlite"
)

// DatabaseProvider 数据库提供商接口，统一不同数据库的操作
type DatabaseProvider interface {
	internal.Configurable

	// GetDBType 获取数据库类型
	GetDBType() DBType

	// GetHost 获取主机地址
	GetHost() string

	// GetPort 获取端口
	GetPort() string

	// GetDBName 获取数据库名称
	GetDBName() string

	// GetUsername 获取用户名
	GetUsername() string

	// GetPassword 获取密码
	GetPassword() string

	// GetConfig 获取额外配置
	GetConfig() string

	// GetModuleName 获取模块名称
	GetModuleName() string

	// SetCredentials 设置凭证
	SetCredentials(username, password string)

	// SetHost 设置主机地址
	SetHost(host string)

	// SetPort 设置端口
	SetPort(port string)

	// SetDBName 设置数据库名称
	SetDBName(dbName string)
}

// Database 数据库统一配置结构
type Database struct {
	Type       DBType      `mapstructure:"type" yaml:"type" json:"type"`                   // 数据库类型
	Enabled    bool        `mapstructure:"enabled" yaml:"enabled" json:"enabled"`          // 是否启用
	Default    string      `mapstructure:"default" yaml:"default" json:"default"`          // 默认使用的数据库
	MySQL      *MySQL      `mapstructure:"mysql" yaml:"mysql" json:"mysql"`                // MySQL配置
	PostgreSQL *PostgreSQL `mapstructure:"postgresql" yaml:"postgresql" json:"postgresql"` // PostgreSQL配置
	SQLite     *SQLite     `mapstructure:"sqlite" yaml:"sqlite" json:"sqlite"`             // SQLite配置
}

// NewDatabase 创建新的数据库配置管理器
func NewDatabase() *Database {
	return &Database{
		Type:       DBTypeMySQL,
		Enabled:    true,
		Default:    string(DBTypeMySQL),
		MySQL:      DefaultMySQL(),
		PostgreSQL: DefaultPostgreSQL(),
		SQLite:     DefaultSQLite(),
	}
}

// GetProvider 获取指定类型的数据库提供商
func (c *Database) GetProvider(dbType DBType) (DatabaseProvider, error) {
	switch dbType {
	case DBTypeMySQL:
		if c.MySQL == nil {
			return nil, fmt.Errorf("mysql config not found")
		}
		return c.MySQL, nil
	case DBTypePostgreSQL:
		if c.PostgreSQL == nil {
			return nil, fmt.Errorf("postgresql config not found")
		}
		return c.PostgreSQL, nil
	case DBTypeSQLite:
		if c.SQLite == nil {
			return nil, fmt.Errorf("sqlite config not found")
		}
		return c.SQLite, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}

// GetDefaultProvider 获取默认的数据库提供商
func (c *Database) GetDefaultProvider() (DatabaseProvider, error) {
	defaultType := DBType(c.Default)
	return c.GetProvider(defaultType)
}

// SetDefaultProvider 设置默认数据库提供商
func (c *Database) SetDefaultProvider(dbType DBType) {
	c.Default = string(dbType)
	c.Type = dbType
}

// ListAvailableProviders 列出所有可用的数据库提供商
func (c *Database) ListAvailableProviders() []DBType {
	var providers []DBType

	if c.MySQL != nil {
		providers = append(providers, DBTypeMySQL)
	}
	if c.PostgreSQL != nil {
		providers = append(providers, DBTypePostgreSQL)
	}
	if c.SQLite != nil {
		providers = append(providers, DBTypeSQLite)
	}

	return providers
}

// ValidateProvider 验证指定提供商的配置
func (c *Database) ValidateProvider(dbType DBType) error {
	provider, err := c.GetProvider(dbType)
	if err != nil {
		return err
	}

	return provider.Validate()
}

// ValidateAll 验证所有配置的提供商
func (c *Database) ValidateAll() error {
	providers := c.ListAvailableProviders()

	for _, providerType := range providers {
		if err := c.ValidateProvider(providerType); err != nil {
			return fmt.Errorf("validation failed for %s: %w", providerType, err)
		}
	}

	return nil
}

// Clone 返回Database配置的副本
func (c *Database) Clone() internal.Configurable {
	newConfig := &Database{
		Type:    c.Type,
		Enabled: c.Enabled,
		Default: c.Default,
	}

	if c.MySQL != nil {
		newConfig.MySQL = c.MySQL.Clone().(*MySQL)
	}
	if c.PostgreSQL != nil {
		newConfig.PostgreSQL = c.PostgreSQL.Clone().(*PostgreSQL)
	}
	if c.SQLite != nil {
		newConfig.SQLite = c.SQLite.Clone().(*SQLite)
	}

	return newConfig
}

// Get 返回Database配置
func (c *Database) Get() interface{} {
	return c
}

// Set 设置Database配置
func (c *Database) Set(data interface{}) {
	if config, ok := data.(*Database); ok {
		*c = *config
	}
}

// Validate 验证Database配置的有效性
func (c *Database) Validate() error {
	if !c.Enabled {
		return nil
	}

	// 验证默认提供商设置
	if c.Default == "" {
		return fmt.Errorf("default database provider not specified")
	}

	defaultType := DBType(c.Default)
	if err := c.ValidateProvider(defaultType); err != nil {
		return fmt.Errorf("default provider validation failed: %w", err)
	}

	return nil
}

// EnsureDefaults 确保所有子配置都有默认值，避免空指针
func (c *Database) EnsureDefaults() {
	if c.MySQL == nil {
		c.MySQL = DefaultMySQL()
	}
	if c.PostgreSQL == nil {
		c.PostgreSQL = DefaultPostgreSQL()
	}
	if c.SQLite == nil {
		c.SQLite = DefaultSQLite()
	}

	// 如果没有设置默认类型，使用MySQL
	if c.Default == "" {
		c.Default = string(DBTypeMySQL)
	}
	if c.Type == "" {
		c.Type = DBTypeMySQL
	}
}

// GetSafe 安全获取配置，确保所有字段都有默认值
func (c *Database) GetSafe() interface{} {
	if c == nil {
		return DefaultDatabaseConfig()
	}
	c.EnsureDefaults()
	return c
}

// BeforeLoad 配置加载前的钩子
func (c *Database) BeforeLoad() error {
	c.EnsureDefaults()
	return nil
}

// AfterLoad 配置加载后的钩子
func (c *Database) AfterLoad() error {
	c.EnsureDefaults()

	// 验证默认提供商配置的有效性
	if c.Default != "" {
		if err := c.ValidateProvider(DBType(c.Default)); err != nil {
			// 如果默认提供商配置无效，回退到MySQL
			c.Default = string(DBTypeMySQL)
			c.Type = DBTypeMySQL
		}
	}

	return nil
}

// GetProviderName 获取提供商显示名称
func GetProviderName(dbType DBType) string {
	switch dbType {
	case DBTypeMySQL:
		return "MySQL"
	case DBTypePostgreSQL:
		return "PostgreSQL"
	case DBTypeSQLite:
		return "SQLite"
	default:
		return string(dbType)
	}
}

// ParseDBType 解析数据库类型字符串
func ParseDBType(typeStr string) (DBType, error) {
	switch strings.ToLower(typeStr) {
	case "mysql":
		return DBTypeMySQL, nil
	case "postgres", "postgresql":
		return DBTypePostgreSQL, nil
	case "sqlite", "sqlite3":
		return DBTypeSQLite, nil
	default:
		return "", fmt.Errorf("unsupported database type: %s", typeStr)
	}
}

// DefaultDatabaseConfig 返回默认的数据库配置
func DefaultDatabaseConfig() *Database {
	return NewDatabase()
}

// GetSupportedTypes 获取所有支持的数据库类型
func GetSupportedTypes() []DBType {
	return []DBType{
		DBTypeMySQL,
		DBTypePostgreSQL,
		DBTypeSQLite,
	}
}

// ========== Database 链式调用方法 ==========

// WithType 设置数据库类型
func (d *Database) WithType(dbType DBType) *Database {
	d.Type = dbType
	return d
}

// WithEnabled 设置是否启用数据库
func (d *Database) WithEnabled(enabled bool) *Database {
	d.Enabled = enabled
	return d
}

// EnableDatabase 启用数据库
func (d *Database) EnableDatabase() *Database {
	d.Enabled = true
	return d
}

// WithDefault 设置默认数据库
func (d *Database) WithDefault(defaultDB string) *Database {
	d.Default = defaultDB
	return d
}

// WithMySQL 设置MySQL配置
func (d *Database) WithMySQL(mysql *MySQL) *Database {
	d.MySQL = mysql
	return d
}

// EnableMySQL 启用MySQL并设置默认配置
func (d *Database) EnableMySQL() *Database {
	if d.MySQL == nil {
		d.MySQL = DefaultMySQL()
	}
	d.Type = DBTypeMySQL
	d.Default = "mysql"
	return d
}

// WithPostgreSQL 设置PostgreSQL配置
func (d *Database) WithPostgreSQL(postgresql *PostgreSQL) *Database {
	d.PostgreSQL = postgresql
	return d
}

// EnablePostgreSQL 启用PostgreSQL并设置默认配置
func (d *Database) EnablePostgreSQL() *Database {
	if d.PostgreSQL == nil {
		d.PostgreSQL = DefaultPostgreSQL()
	}
	d.Type = DBTypePostgreSQL
	d.Default = "postgresql"
	return d
}

// WithSQLite 设置SQLite配置
func (d *Database) WithSQLite(sqlite *SQLite) *Database {
	d.SQLite = sqlite
	return d
}

// EnableSQLite 启用SQLite并设置默认配置
func (d *Database) EnableSQLite() *Database {
	if d.SQLite == nil {
		d.SQLite = DefaultSQLite()
	}
	d.Type = DBTypeSQLite
	d.Default = "sqlite"
	return d
}
