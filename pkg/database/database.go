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
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-toolbox/pkg/syncx"
	"strings"
)

// DBType 定义数据库类型
type DBType string

const (
	DBTypeMySQL       DBType = "mysql"       // MySQL数据库
	DBTypePostgreSQL  DBType = "postgres"    // PostgreSQL数据库
	DBTypeSQLite      DBType = "sqlite"      // SQLite数据库
	DBTypeClickHouse  DBType = "clickhouse"  // ClickHouse数据库
	DBTypeCockroachDB DBType = "cockroachdb" // CockroachDB数据库
)

// DataCategory 定义数据库分类（关系型/非关系型）
type DataCategory string

const (
	CategoryRelational    DataCategory = "relational"     // 关系型数据库
	CategoryNonRelational DataCategory = "non-relational" // 非关系型数据库
)

// Category 根据数据库类型返回其所属分类
func (t DBType) Category() DataCategory {
	switch t {
	case DBTypeMySQL, DBTypePostgreSQL, DBTypeSQLite, DBTypeCockroachDB:
		return CategoryRelational
	case DBTypeClickHouse:
		return CategoryNonRelational
	default:
		return ""
	}
}

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

	// GetSlowThreshold 获取慢查询阈值（毫秒）
	GetSlowThreshold() int

	// GetIgnoreRecordNotFoundError 获取是否忽略RecordNotFound错误
	GetIgnoreRecordNotFoundError() bool

	// GORM 性能配置 Getter 方法
	// GetSkipDefaultTransaction 获取是否跳过默认事务
	GetSkipDefaultTransaction() bool

	// GetPrepareStmt 获取是否启用预编译语句
	GetPrepareStmt() bool

	// GetDisableForeignKeyConstraintWhenMigrating 获取迁移时是否禁用外键约束
	GetDisableForeignKeyConstraintWhenMigrating() bool

	// GetDisableNestedTransaction 获取是否禁用嵌套事务
	GetDisableNestedTransaction() bool

	// GetAllowGlobalUpdate 获取是否允许全局更新
	GetAllowGlobalUpdate() bool

	// GetQueryFields 获取是否在查询时选择所有字段
	GetQueryFields() bool

	// GetCreateBatchSize 获取批量创建的大小
	GetCreateBatchSize() int

	// GetSingularTable 获取是否使用单数表名
	GetSingularTable() bool

	// SetCredentials 设置凭证
	SetCredentials(username, password string)

	// SetHost 设置主机地址
	SetHost(host string)

	// SetPort 设置端口
	SetPort(port string)

	// SetDBName 设置数据库名称
	SetDBName(dbName string)
}

// RelationalDatabase 关系型数据库配置
type RelationalDatabase struct {
	Enabled     bool         `mapstructure:"enabled" yaml:"enabled" json:"enabled"`             // 是否启用关系型数据库
	Default     string       `mapstructure:"default" yaml:"default" json:"default"`             // 默认关系型数据库类型
	MySQL       *MySQL       `mapstructure:"mysql" yaml:"mysql" json:"mysql"`                   // MySQL配置
	PostgreSQL  *PostgreSQL  `mapstructure:"postgresql" yaml:"postgresql" json:"postgresql"`    // PostgreSQL配置
	SQLite      *SQLite      `mapstructure:"sqlite" yaml:"sqlite" json:"sqlite"`                // SQLite配置
	CockroachDB *CockroachDB `mapstructure:"cockroachdb" yaml:"cockroachdb" json:"cockroachdb"` // CockroachDB配置
}

// NewRelationalDatabase 创建默认关系型数据库配置
func NewRelationalDatabase() *RelationalDatabase {
	return &RelationalDatabase{
		Enabled:     false,
		Default:     string(DBTypeMySQL),
		MySQL:       DefaultMySQL(),
		PostgreSQL:  DefaultPostgreSQL(),
		SQLite:      DefaultSQLite(),
		CockroachDB: DefaultCockroachDB(),
	}
}

// GetProvider 根据数据库类型获取对应的关系型数据库提供商
func (r *RelationalDatabase) GetProvider(dbType DBType) (DatabaseProvider, error) {
	switch dbType {
	case DBTypeMySQL:
		if r.MySQL == nil {
			return nil, fmt.Errorf("mysql config not found")
		}
		return r.MySQL, nil
	case DBTypePostgreSQL:
		if r.PostgreSQL == nil {
			return nil, fmt.Errorf("postgresql config not found")
		}
		return r.PostgreSQL, nil
	case DBTypeSQLite:
		if r.SQLite == nil {
			return nil, fmt.Errorf("sqlite config not found")
		}
		return r.SQLite, nil
	case DBTypeCockroachDB:
		if r.CockroachDB == nil {
			return nil, fmt.Errorf("cockroachdb config not found")
		}
		return r.CockroachDB, nil
	default:
		return nil, fmt.Errorf("unsupported relational database type: %s", dbType)
	}
}

// GetDefaultProvider 获取默认的关系型数据库提供商
func (r *RelationalDatabase) GetDefaultProvider() (DatabaseProvider, error) {
	if r.Default == "" {
		return nil, fmt.Errorf("relational default database not specified")
	}
	return r.GetProvider(DBType(r.Default))
}

// SetDefault 设置默认的关系型数据库类型
func (r *RelationalDatabase) SetDefault(dbType DBType) {
	r.Default = string(dbType)
}

// ListAvailableProviders 列出所有可用的关系型数据库提供商
func (r *RelationalDatabase) ListAvailableProviders() []DBType {
	var providers []DBType
	if r.MySQL != nil {
		providers = append(providers, DBTypeMySQL)
	}
	if r.PostgreSQL != nil {
		providers = append(providers, DBTypePostgreSQL)
	}
	if r.SQLite != nil {
		providers = append(providers, DBTypeSQLite)
	}
	if r.CockroachDB != nil {
		providers = append(providers, DBTypeCockroachDB)
	}
	return providers
}

// ValidateProvider 验证指定的关系型数据库提供商配置
func (r *RelationalDatabase) ValidateProvider(dbType DBType) error {
	provider, err := r.GetProvider(dbType)
	if err != nil {
		return err
	}
	return provider.Validate()
}

// ValidateAll 验证所有关系型数据库提供商配置
func (r *RelationalDatabase) ValidateAll() error {
	providers := r.ListAvailableProviders()
	for _, providerType := range providers {
		if err := r.ValidateProvider(providerType); err != nil {
			return fmt.Errorf("validation failed for %s: %w", providerType, err)
		}
	}
	return nil
}

// Validate 验证关系型数据库配置
func (r *RelationalDatabase) Validate() error {
	if !r.Enabled {
		return nil
	}
	if r.Default == "" {
		return fmt.Errorf("relational default database not specified")
	}
	if err := r.ValidateProvider(DBType(r.Default)); err != nil {
		return fmt.Errorf("relational default provider validation failed: %w", err)
	}
	return nil
}

// EnsureDefaults 确保关系型数据库配置的默认值
func (r *RelationalDatabase) EnsureDefaults() {
	if r.MySQL == nil {
		r.MySQL = DefaultMySQL()
	}
	if r.PostgreSQL == nil {
		r.PostgreSQL = DefaultPostgreSQL()
	}
	if r.SQLite == nil {
		r.SQLite = DefaultSQLite()
	}
	if r.CockroachDB == nil {
		r.CockroachDB = DefaultCockroachDB()
	}
	if r.Default == "" {
		r.Default = string(DBTypeMySQL)
	}
}

// NonRelationalDatabase 非关系型数据库配置
type NonRelationalDatabase struct {
	Enabled    bool        `mapstructure:"enabled" yaml:"enabled" json:"enabled"`          // 是否启用非关系型数据库
	Default    string      `mapstructure:"default" yaml:"default" json:"default"`          // 默认非关系型数据库类型
	ClickHouse *ClickHouse `mapstructure:"clickhouse" yaml:"clickhouse" json:"clickhouse"` // ClickHouse配置
}

// NewNonRelationalDatabase 创建默认非关系型数据库配置
func NewNonRelationalDatabase() *NonRelationalDatabase {
	return &NonRelationalDatabase{
		Enabled:    false,
		Default:    string(DBTypeClickHouse),
		ClickHouse: DefaultClickHouse(),
	}
}

// GetProvider 根据数据库类型获取对应的非关系型数据库提供商
func (n *NonRelationalDatabase) GetProvider(dbType DBType) (DatabaseProvider, error) {
	switch dbType {
	case DBTypeClickHouse:
		if n.ClickHouse == nil {
			return nil, fmt.Errorf("clickhouse config not found")
		}
		return n.ClickHouse, nil
	default:
		return nil, fmt.Errorf("unsupported non-relational database type: %s", dbType)
	}
}

// GetDefaultProvider 获取默认的非关系型数据库提供商
func (n *NonRelationalDatabase) GetDefaultProvider() (DatabaseProvider, error) {
	if n.Default == "" {
		return nil, fmt.Errorf("non-relational default database not specified")
	}
	return n.GetProvider(DBType(n.Default))
}

// SetDefault 设置默认的非关系型数据库类型
func (n *NonRelationalDatabase) SetDefault(dbType DBType) {
	n.Default = string(dbType)
}

// ListAvailableProviders 列出所有可用的非关系型数据库提供商
func (n *NonRelationalDatabase) ListAvailableProviders() []DBType {
	var providers []DBType
	if n.ClickHouse != nil {
		providers = append(providers, DBTypeClickHouse)
	}
	return providers
}

// ValidateProvider 验证指定的非关系型数据库提供商配置
func (n *NonRelationalDatabase) ValidateProvider(dbType DBType) error {
	provider, err := n.GetProvider(dbType)
	if err != nil {
		return err
	}
	return provider.Validate()
}

// ValidateAll 验证所有非关系型数据库提供商配置
func (n *NonRelationalDatabase) ValidateAll() error {
	providers := n.ListAvailableProviders()
	for _, providerType := range providers {
		if err := n.ValidateProvider(providerType); err != nil {
			return fmt.Errorf("validation failed for %s: %w", providerType, err)
		}
	}
	return nil
}

// Validate 验证非关系型数据库配置
func (n *NonRelationalDatabase) Validate() error {
	if !n.Enabled {
		return nil
	}
	if n.Default == "" {
		return fmt.Errorf("non-relational default database not specified")
	}
	if err := n.ValidateProvider(DBType(n.Default)); err != nil {
		return fmt.Errorf("non-relational default provider validation failed: %w", err)
	}
	return nil
}

// EnsureDefaults 确保非关系型数据库配置的默认值
func (n *NonRelationalDatabase) EnsureDefaults() {
	if n.ClickHouse == nil {
		n.ClickHouse = DefaultClickHouse()
	}
	if n.Default == "" {
		n.Default = string(DBTypeClickHouse)
	}
}

// Database 数据库统一配置，支持关系型和非关系型数据库同时使用
type Database struct {
	Enabled       bool                   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                     // 是否启用数据库
	Relational    *RelationalDatabase    `mapstructure:"relational" yaml:"relational" json:"relational"`            // 关系型数据库配置
	NonRelational *NonRelationalDatabase `mapstructure:"non-relational" yaml:"non-relational" json:"nonRelational"` // 非关系型数据库配置
}

// NewDatabase 创建默认数据库配置（关系型和非关系型均初始化）
func NewDatabase() *Database {
	return &Database{
		Enabled:       false,
		Relational:    NewRelationalDatabase(),
		NonRelational: NewNonRelationalDatabase(),
	}
}

// GetProvider 根据数据库类型获取对应的数据库提供商（自动判断分类）
func (c *Database) GetProvider(dbType DBType) (DatabaseProvider, error) {
	category := dbType.Category()
	switch category {
	case CategoryRelational:
		if c.Relational == nil {
			return nil, fmt.Errorf("relational database config not found")
		}
		return c.Relational.GetProvider(dbType)
	case CategoryNonRelational:
		if c.NonRelational == nil {
			return nil, fmt.Errorf("non-relational database config not found")
		}
		return c.NonRelational.GetProvider(dbType)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}

// GetDefaultProviderByCategory 根据分类获取默认数据库提供商
func (c *Database) GetDefaultProviderByCategory(category DataCategory) (DatabaseProvider, error) {
	switch category {
	case CategoryRelational:
		if c.Relational == nil {
			return nil, fmt.Errorf("relational database config not found")
		}
		return c.Relational.GetDefaultProvider()
	case CategoryNonRelational:
		if c.NonRelational == nil {
			return nil, fmt.Errorf("non-relational database config not found")
		}
		return c.NonRelational.GetDefaultProvider()
	default:
		return nil, fmt.Errorf("unsupported category: %s", category)
	}
}

// GetDefaultRelationalProvider 获取默认的关系型数据库提供商
func (c *Database) GetDefaultRelationalProvider() (DatabaseProvider, error) {
	return c.GetDefaultProviderByCategory(CategoryRelational)
}

// GetDefaultNonRelationalProvider 获取默认的非关系型数据库提供商
func (c *Database) GetDefaultNonRelationalProvider() (DatabaseProvider, error) {
	return c.GetDefaultProviderByCategory(CategoryNonRelational)
}

// ListAvailableProviders 列出所有可用的数据库提供商
func (c *Database) ListAvailableProviders() []DBType {
	var providers []DBType
	if c.Relational != nil {
		providers = append(providers, c.Relational.ListAvailableProviders()...)
	}
	if c.NonRelational != nil {
		providers = append(providers, c.NonRelational.ListAvailableProviders()...)
	}
	return providers
}

// ListAvailableProvidersByCategory 根据分类列出可用的数据库提供商
func (c *Database) ListAvailableProvidersByCategory(category DataCategory) []DBType {
	switch category {
	case CategoryRelational:
		if c.Relational == nil {
			return nil
		}
		return c.Relational.ListAvailableProviders()
	case CategoryNonRelational:
		if c.NonRelational == nil {
			return nil
		}
		return c.NonRelational.ListAvailableProviders()
	default:
		return nil
	}
}

// ValidateProvider 验证指定数据库提供商的配置
func (c *Database) ValidateProvider(dbType DBType) error {
	provider, err := c.GetProvider(dbType)
	if err != nil {
		return err
	}

	return provider.Validate()
}

// ValidateAll 验证所有数据库提供商配置
func (c *Database) ValidateAll() error {
	if c.Relational != nil {
		if err := c.Relational.ValidateAll(); err != nil {
			return err
		}
	}
	if c.NonRelational != nil {
		if err := c.NonRelational.ValidateAll(); err != nil {
			return err
		}
	}
	return nil
}

// Clone 深拷贝数据库配置
func (c *Database) Clone() internal.Configurable {
	var cloned Database
	if err := syncx.DeepCopy(&cloned, c); err != nil {
		// 如果深拷贝失败，返回空配置
		return &Database{}
	}
	return &cloned
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
	if c.Relational != nil {
		if err := c.Relational.Validate(); err != nil {
			return err
		}
	}
	if c.NonRelational != nil {
		if err := c.NonRelational.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// EnsureDefaults 确保数据库配置的默认值
func (c *Database) EnsureDefaults() {
	if c.Relational == nil {
		c.Relational = NewRelationalDatabase()
	} else {
		c.Relational.EnsureDefaults()
	}
	if c.NonRelational == nil {
		c.NonRelational = NewNonRelationalDatabase()
	} else {
		c.NonRelational.EnsureDefaults()
	}
}

// GetSafe 安全获取数据库配置，如果为空则返回默认配置
func (c *Database) GetSafe() interface{} {
	if c == nil {
		return DefaultDatabaseConfig()
	}
	c.EnsureDefaults()
	return c
}

// BeforeLoad 加载前确保默认值
func (c *Database) BeforeLoad() error {
	c.EnsureDefaults()
	return nil
}

// AfterLoad 加载后校验并修正默认数据库类型
func (c *Database) AfterLoad() error {
	c.EnsureDefaults()

	// 校验关系型默认数据库类型，无效则回退为MySQL
	if c.Relational != nil && c.Relational.Default != "" {
		if err := c.Relational.ValidateProvider(DBType(c.Relational.Default)); err != nil {
			c.Relational.Default = string(DBTypeMySQL)
		}
	}
	// 校验非关系型默认数据库类型，无效则回退为ClickHouse
	if c.NonRelational != nil && c.NonRelational.Default != "" {
		if err := c.NonRelational.ValidateProvider(DBType(c.NonRelational.Default)); err != nil {
			c.NonRelational.Default = string(DBTypeClickHouse)
		}
	}

	return nil
}

// GetProviderName 获取数据库类型对应的人类可读名称
func GetProviderName(dbType DBType) string {
	switch dbType {
	case DBTypeMySQL:
		return "MySQL"
	case DBTypePostgreSQL:
		return "PostgreSQL"
	case DBTypeSQLite:
		return "SQLite"
	case DBTypeClickHouse:
		return "ClickHouse"
	case DBTypeCockroachDB:
		return "CockroachDB"
	default:
		return string(dbType)
	}
}

// ParseDBType 将字符串解析为数据库类型
func ParseDBType(typeStr string) (DBType, error) {
	switch strings.ToLower(typeStr) {
	case "mysql":
		return DBTypeMySQL, nil
	case "postgres", "postgresql":
		return DBTypePostgreSQL, nil
	case "sqlite", "sqlite3":
		return DBTypeSQLite, nil
	case "clickhouse":
		return DBTypeClickHouse, nil
	case "cockroachdb", "cockroach", "crdb":
		return DBTypeCockroachDB, nil
	default:
		return "", fmt.Errorf("unsupported database type: %s", typeStr)
	}
}

// GetCategoryByDBType 根据数据库类型获取其分类
func GetCategoryByDBType(dbType DBType) DataCategory {
	return dbType.Category()
}

// DefaultDatabaseConfig 创建默认数据库配置
func DefaultDatabaseConfig() *Database {
	return NewDatabase()
}

// GetSupportedTypes 获取所有支持的数据库类型
func GetSupportedTypes() []DBType {
	return []DBType{
		DBTypeMySQL,
		DBTypePostgreSQL,
		DBTypeSQLite,
		DBTypeClickHouse,
		DBTypeCockroachDB,
	}
}

// GetSupportedTypesByCategory 根据分类获取支持的数据库类型
func GetSupportedTypesByCategory(category DataCategory) []DBType {
	var result []DBType
	for _, t := range GetSupportedTypes() {
		if t.Category() == category {
			result = append(result, t)
		}
	}
	return result
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

// WithRelational 设置关系型数据库配置
func (d *Database) WithRelational(relational *RelationalDatabase) *Database {
	d.Relational = relational
	return d
}

// EnableRelational 启用关系型数据库
func (d *Database) EnableRelational() *Database {
	if d.Relational == nil {
		d.Relational = NewRelationalDatabase()
	}
	d.Relational.Enabled = true
	return d
}

// WithNonRelational 设置非关系型数据库配置
func (d *Database) WithNonRelational(nonRelational *NonRelationalDatabase) *Database {
	d.NonRelational = nonRelational
	return d
}

// EnableNonRelational 启用非关系型数据库
func (d *Database) EnableNonRelational() *Database {
	if d.NonRelational == nil {
		d.NonRelational = NewNonRelationalDatabase()
	}
	d.NonRelational.Enabled = true
	return d
}

// WithMySQL 设置MySQL配置
func (d *Database) WithMySQL(mysql *MySQL) *Database {
	if d.Relational == nil {
		d.Relational = NewRelationalDatabase()
	}
	d.Relational.MySQL = mysql
	return d
}

// EnableMySQL 启用MySQL并设为默认关系型数据库
func (d *Database) EnableMySQL() *Database {
	if d.Relational == nil {
		d.Relational = NewRelationalDatabase()
	}
	if d.Relational.MySQL == nil {
		d.Relational.MySQL = DefaultMySQL()
	}
	d.Relational.Enabled = true
	d.Relational.Default = "mysql"
	return d
}

// WithPostgreSQL 设置PostgreSQL配置
func (d *Database) WithPostgreSQL(postgresql *PostgreSQL) *Database {
	if d.Relational == nil {
		d.Relational = NewRelationalDatabase()
	}
	d.Relational.PostgreSQL = postgresql
	return d
}

// EnablePostgreSQL 启用PostgreSQL并设为默认关系型数据库
func (d *Database) EnablePostgreSQL() *Database {
	if d.Relational == nil {
		d.Relational = NewRelationalDatabase()
	}
	if d.Relational.PostgreSQL == nil {
		d.Relational.PostgreSQL = DefaultPostgreSQL()
	}
	d.Relational.Enabled = true
	d.Relational.Default = "postgresql"
	return d
}

// WithSQLite 设置SQLite配置
func (d *Database) WithSQLite(sqlite *SQLite) *Database {
	if d.Relational == nil {
		d.Relational = NewRelationalDatabase()
	}
	d.Relational.SQLite = sqlite
	return d
}

// EnableSQLite 启用SQLite并设为默认关系型数据库
func (d *Database) EnableSQLite() *Database {
	if d.Relational == nil {
		d.Relational = NewRelationalDatabase()
	}
	if d.Relational.SQLite == nil {
		d.Relational.SQLite = DefaultSQLite()
	}
	d.Relational.Enabled = true
	d.Relational.Default = "sqlite"
	return d
}

// WithClickHouse 设置ClickHouse配置
func (d *Database) WithClickHouse(clickhouse *ClickHouse) *Database {
	if d.NonRelational == nil {
		d.NonRelational = NewNonRelationalDatabase()
	}
	d.NonRelational.ClickHouse = clickhouse
	return d
}

// EnableClickHouse 启用ClickHouse并设为默认非关系型数据库
func (d *Database) EnableClickHouse() *Database {
	if d.NonRelational == nil {
		d.NonRelational = NewNonRelationalDatabase()
	}
	if d.NonRelational.ClickHouse == nil {
		d.NonRelational.ClickHouse = DefaultClickHouse()
	}
	d.NonRelational.Enabled = true
	d.NonRelational.Default = string(DBTypeClickHouse)
	return d
}

// WithCockroachDB 设置CockroachDB配置
func (d *Database) WithCockroachDB(cockroachdb *CockroachDB) *Database {
	if d.Relational == nil {
		d.Relational = NewRelationalDatabase()
	}
	d.Relational.CockroachDB = cockroachdb
	return d
}

// EnableCockroachDB 启用CockroachDB并设为默认关系型数据库
func (d *Database) EnableCockroachDB() *Database {
	if d.Relational == nil {
		d.Relational = NewRelationalDatabase()
	}
	if d.Relational.CockroachDB == nil {
		d.Relational.CockroachDB = DefaultCockroachDB()
	}
	d.Relational.Enabled = true
	d.Relational.Default = string(DBTypeCockroachDB)
	return d
}
