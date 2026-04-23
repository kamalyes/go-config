/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-04-23 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-04-23 00:00:00
 * @FilePath: \go-config\pkg\database\cockroachdb.go
 * @Description: CockroachDB数据库配置（兼容PostgreSQL协议的分布式数据库）
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package database

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-toolbox/pkg/syncx"
)

// CockroachDB CockroachDB数据库配置
// CockroachDB 使用 PostgreSQL 协议，因此大部分配置与 PostgreSQL 类似
type CockroachDB struct {
	ModuleName                               string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                                                                                                   // 模块名称
	Host                                     string `mapstructure:"host" yaml:"host" json:"host"            validate:"required"`                                                                                        // 数据库 IP 地址
	Port                                     string `mapstructure:"port" yaml:"port" json:"port"            validate:"required"`                                                                                        // 端口（默认26257）
	Config                                   string `mapstructure:"config" yaml:"config" json:"config"          validate:"required"`                                                                                    // 后缀配置
	LogLevel                                 string `mapstructure:"log-level" yaml:"log-level" json:"logLevel"       validate:"required"`                                                                               // SQL 日志等级
	SlowThreshold                            int    `mapstructure:"slow-threshold" yaml:"slow-threshold" json:"slowThreshold"`                                                                                          // 慢查询阈值（毫秒）
	IgnoreRecordNotFoundError                bool   `mapstructure:"ignore-record-not-found-error" yaml:"ignore-record-not-found-error" json:"ignoreRecordNotFoundError"`                                                // 是否忽略ErrRecordNotFound错误
	Dbname                                   string `mapstructure:"db-name" yaml:"db-name" json:"dbName"         validate:"required"`                                                                                   // 数据库名称
	Username                                 string `mapstructure:"username" yaml:"username" json:"username"        validate:"required"`                                                                                // 数据库用户名
	Password                                 string `mapstructure:"password" yaml:"password" json:"password"`                                                                                                           // 数据库密码（CockroachDB insecure 模式下可为空）
	SSLMode                                  string `mapstructure:"ssl-mode" yaml:"ssl-mode" json:"sslMode"`                                                                                                            // SSL模式: disable, require, verify-ca, verify-full
	SSLRootCert                              string `mapstructure:"ssl-root-cert" yaml:"ssl-root-cert" json:"sslRootCert"`                                                                                              // SSL根证书路径
	SSLCert                                  string `mapstructure:"ssl-cert" yaml:"ssl-cert" json:"sslCert"`                                                                                                            // SSL客户端证书路径
	SSLKey                                   string `mapstructure:"ssl-key" yaml:"ssl-key" json:"sslKey"`                                                                                                               // SSL客户端密钥路径
	ApplicationName                          string `mapstructure:"application-name" yaml:"application-name" json:"applicationName"`                                                                                    // 应用名称（便于CockroachDB追踪）
	MaxIdleConns                             int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns" json:"maxIdleConns"  validate:"min=0"`                                                                         // 最大空闲连接数
	MaxOpenConns                             int    `mapstructure:"max-open-conns" yaml:"max-open-conns" json:"maxOpenConns"  validate:"min=0"`                                                                         // 最大连接数
	ConnMaxIdleTime                          int    `mapstructure:"conn-max-idle-time" yaml:"conn-max-idle-time" json:"connMaxIdleTime" validate:"min=0"`                                                               // 连接最大空闲时间 单位：秒
	ConnMaxLifeTime                          int    `mapstructure:"conn-max-life-time" yaml:"conn-max-life-time" json:"connMaxLifeTime" validate:"min=0"`                                                               // 连接最大生命周期 单位：秒
	SkipDefaultTransaction                   bool   `mapstructure:"skip-default-transaction" yaml:"skip-default-transaction" json:"skipDefaultTransaction"`                                                             // 跳过默认事务
	PrepareStmt                              bool   `mapstructure:"prepare-stmt" yaml:"prepare-stmt" json:"prepareStmt"`                                                                                                // 预编译语句
	DisableForeignKeyConstraintWhenMigrating bool   `mapstructure:"disable-foreign-key-constraint-when-migrating" yaml:"disable-foreign-key-constraint-when-migrating" json:"disableForeignKeyConstraintWhenMigrating"` // 禁用自动创建外键约束
	DisableNestedTransaction                 bool   `mapstructure:"disable-nested-transaction" yaml:"disable-nested-transaction" json:"disableNestedTransaction"`                                                       // 禁用嵌套事务
	AllowGlobalUpdate                        bool   `mapstructure:"allow-global-update" yaml:"allow-global-update" json:"allowGlobalUpdate"`                                                                            // 允许全局更新
	QueryFields                              bool   `mapstructure:"query-fields" yaml:"query-fields" json:"queryFields"`                                                                                                // 执行查询时选择所有字段
	CreateBatchSize                          int    `mapstructure:"create-batch-size" yaml:"create-batch-size" json:"createBatchSize"`                                                                                  // 批量创建大小
	SingularTable                            bool   `mapstructure:"singular-table" yaml:"singular-table" json:"singularTable"`                                                                                          // 使用单数表名
	MaxRetries                               int    `mapstructure:"max-retries" yaml:"max-retries" json:"maxRetries" validate:"min=0"`                                                                                  // 事务重试次数（CockroachDB乐观事务遇到冲突时重试）
}

// 为CockroachDB配置实现DatabaseProvider接口
func (c *CockroachDB) GetDBType() DBType                  { return DBTypeCockroachDB }
func (c *CockroachDB) GetHost() string                    { return c.Host }
func (c *CockroachDB) GetPort() string                    { return c.Port }
func (c *CockroachDB) GetDBName() string                  { return c.Dbname }
func (c *CockroachDB) GetUsername() string                { return c.Username }
func (c *CockroachDB) GetPassword() string                { return c.Password }
func (c *CockroachDB) GetConfig() string                  { return c.Config }
func (c *CockroachDB) GetModuleName() string              { return c.ModuleName }
func (c *CockroachDB) GetSlowThreshold() int              { return c.SlowThreshold }
func (c *CockroachDB) GetIgnoreRecordNotFoundError() bool { return c.IgnoreRecordNotFoundError }
func (c *CockroachDB) GetSkipDefaultTransaction() bool    { return c.SkipDefaultTransaction }
func (c *CockroachDB) GetPrepareStmt() bool               { return c.PrepareStmt }
func (c *CockroachDB) GetDisableForeignKeyConstraintWhenMigrating() bool {
	return c.DisableForeignKeyConstraintWhenMigrating
}
func (c *CockroachDB) GetDisableNestedTransaction() bool { return c.DisableNestedTransaction }
func (c *CockroachDB) GetAllowGlobalUpdate() bool        { return c.AllowGlobalUpdate }
func (c *CockroachDB) GetQueryFields() bool              { return c.QueryFields }
func (c *CockroachDB) GetCreateBatchSize() int           { return c.CreateBatchSize }
func (c *CockroachDB) GetSingularTable() bool            { return c.SingularTable }
func (c *CockroachDB) SetCredentials(username, password string) {
	c.Username, c.Password = username, password
}
func (c *CockroachDB) SetHost(host string)     { c.Host = host }
func (c *CockroachDB) SetPort(port string)     { c.Port = port }
func (c *CockroachDB) SetDBName(dbName string) { c.Dbname = dbName }

func (c *CockroachDB) Clone() internal.Configurable {
	var cloned CockroachDB
	if err := syncx.DeepCopy(&cloned, c); err != nil {
		// 如果深拷贝失败，返回空配置
		return &CockroachDB{}
	}
	return &cloned
}
func (c *CockroachDB) Get() interface{} { return c }
func (c *CockroachDB) Set(data interface{}) {
	if cfg, ok := data.(*CockroachDB); ok {
		*c = *cfg
	}
}
func (c *CockroachDB) Validate() error { return internal.ValidateStruct(c) }

// DefaultCockroachDB 创建默认CockroachDB配置
func DefaultCockroachDB() *CockroachDB {
	return &CockroachDB{
		ModuleName:                               "cockroachdb",
		Host:                                     "localhost",
		Port:                                     "26257",
		Config:                                   "sslmode=disable",
		LogLevel:                                 "info",
		SlowThreshold:                            100, // 100毫秒
		IgnoreRecordNotFoundError:                false,
		Dbname:                                   "defaultdb",
		Username:                                 "root",
		Password:                                 "",
		SSLMode:                                  "disable",
		SSLRootCert:                              "",
		SSLCert:                                  "",
		SSLKey:                                   "",
		ApplicationName:                          "go-config",
		MaxIdleConns:                             10,
		MaxOpenConns:                             100,
		ConnMaxIdleTime:                          300,   // 5分钟
		ConnMaxLifeTime:                          3600,  // 1小时
		SkipDefaultTransaction:                   false, // 使用默认事务
		PrepareStmt:                              true,  // 启用预编译语句缓存
		DisableForeignKeyConstraintWhenMigrating: true,  // 禁用外键约束
		DisableNestedTransaction:                 false, // 允许嵌套事务
		AllowGlobalUpdate:                        false, // 禁止全局更新
		QueryFields:                              true,  // 查询时选择所有字段
		CreateBatchSize:                          100,   // 批量创建大小
		SingularTable:                            true,  // 使用单数表名
		MaxRetries:                               3,     // 事务冲突重试次数
	}
}

// NewCockroachDBConfig 创建一个新的CockroachDB配置实例
func NewCockroachDBConfig(opt *CockroachDB) *CockroachDB {
	var cockroachdbInstance *CockroachDB
	internal.LockFunc(func() {
		cockroachdbInstance = opt
	})
	return cockroachdbInstance
}

// WithModuleName 设置模块名称
func (c *CockroachDB) WithModuleName(moduleName string) *CockroachDB {
	c.ModuleName = moduleName
	return c
}

// WithHost 设置数据库主机地址
func (c *CockroachDB) WithHost(host string) *CockroachDB {
	c.Host = host
	return c
}

// WithPort 设置数据库端口
func (c *CockroachDB) WithPort(port string) *CockroachDB {
	c.Port = port
	return c
}

// WithConfig 设置数据库配置参数
func (c *CockroachDB) WithConfig(config string) *CockroachDB {
	c.Config = config
	return c
}

// WithLogLevel 设置日志级别
func (c *CockroachDB) WithLogLevel(logLevel string) *CockroachDB {
	c.LogLevel = logLevel
	return c
}

// WithDbname 设置数据库名称
func (c *CockroachDB) WithDbname(dbname string) *CockroachDB {
	c.Dbname = dbname
	return c
}

// WithUsername 设置数据库用户名
func (c *CockroachDB) WithUsername(username string) *CockroachDB {
	c.Username = username
	return c
}

// WithPassword 设置数据库密码
func (c *CockroachDB) WithPassword(password string) *CockroachDB {
	c.Password = password
	return c
}

// WithSSLMode 设置SSL模式
func (c *CockroachDB) WithSSLMode(sslMode string) *CockroachDB {
	c.SSLMode = sslMode
	return c
}

// WithSSLRootCert 设置SSL根证书路径
func (c *CockroachDB) WithSSLRootCert(sslRootCert string) *CockroachDB {
	c.SSLRootCert = sslRootCert
	return c
}

// WithSSLCert 设置SSL客户端证书路径
func (c *CockroachDB) WithSSLCert(sslCert string) *CockroachDB {
	c.SSLCert = sslCert
	return c
}

// WithSSLKey 设置SSL客户端密钥路径
func (c *CockroachDB) WithSSLKey(sslKey string) *CockroachDB {
	c.SSLKey = sslKey
	return c
}

// WithApplicationName 设置应用名称
func (c *CockroachDB) WithApplicationName(applicationName string) *CockroachDB {
	c.ApplicationName = applicationName
	return c
}

// WithMaxIdleConns 设置最大空闲连接数
func (c *CockroachDB) WithMaxIdleConns(maxIdleConns int) *CockroachDB {
	c.MaxIdleConns = maxIdleConns
	return c
}

// WithMaxOpenConns 设置最大连接数
func (c *CockroachDB) WithMaxOpenConns(maxOpenConns int) *CockroachDB {
	c.MaxOpenConns = maxOpenConns
	return c
}

// WithConnMaxIdleTime 设置连接最大空闲时间
func (c *CockroachDB) WithConnMaxIdleTime(connMaxIdleTime int) *CockroachDB {
	c.ConnMaxIdleTime = connMaxIdleTime
	return c
}

// WithConnMaxLifeTime 设置连接最大生命周期
func (c *CockroachDB) WithConnMaxLifeTime(connMaxLifeTime int) *CockroachDB {
	c.ConnMaxLifeTime = connMaxLifeTime
	return c
}

// WithMaxRetries 设置事务重试次数
func (c *CockroachDB) WithMaxRetries(maxRetries int) *CockroachDB {
	c.MaxRetries = maxRetries
	return c
}
