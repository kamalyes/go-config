/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 23:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 16:54:04
 * @FilePath: \go-config\pkg\database\mysql.go
 * @Description: MySQL数据库配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package database

import (
	"github.com/kamalyes/go-config/internal"
)

// MySQL MySQL数据库配置
type MySQL struct {
	ModuleName                               string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                                                                                                   // 模块名称
	Host                                     string `mapstructure:"host" yaml:"host" json:"host"            validate:"required"`                                                                                        // 数据库 IP 地址
	Port                                     string `mapstructure:"port" yaml:"port" json:"port"            validate:"required"`                                                                                        // 端口
	Config                                   string `mapstructure:"config" yaml:"config" json:"config"          validate:"required"`                                                                                    // 后缀配置 默认配置 charset=utf8mb4&parseTime=True&loc=Local
	LogLevel                                 string `mapstructure:"log-level" yaml:"log-level" json:"logLevel"       validate:"required"`                                                                               // SQL 日志等级
	SlowThreshold                            int    `mapstructure:"slow-threshold" yaml:"slow-threshold" json:"slowThreshold"`                                                                                          // 慢查询阈值（毫秒）
	IgnoreRecordNotFoundError                bool   `mapstructure:"ignore-record-not-found-error" yaml:"ignore-record-not-found-error" json:"ignoreRecordNotFoundError"`                                                // 是否忽略ErrRecordNotFound错误
	Dbname                                   string `mapstructure:"db-name" yaml:"db-name" json:"dbName"         validate:"required"`                                                                                   // 数据库名称
	Username                                 string `mapstructure:"username" yaml:"username" json:"username"        validate:"required"`                                                                                // 数据库用户名
	Password                                 string `mapstructure:"password" yaml:"password" json:"password"        validate:"required"`                                                                                // 数据库密码
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
}

// 为MySQL配置实现DatabaseProvider接口
func (m *MySQL) GetDBType() DBType                  { return DBTypeMySQL }
func (m *MySQL) GetHost() string                    { return m.Host }
func (m *MySQL) GetPort() string                    { return m.Port }
func (m *MySQL) GetDBName() string                  { return m.Dbname }
func (m *MySQL) GetUsername() string                { return m.Username }
func (m *MySQL) GetPassword() string                { return m.Password }
func (m *MySQL) GetConfig() string                  { return m.Config }
func (m *MySQL) GetModuleName() string              { return m.ModuleName }
func (m *MySQL) GetSlowThreshold() int              { return m.SlowThreshold }
func (m *MySQL) GetIgnoreRecordNotFoundError() bool { return m.IgnoreRecordNotFoundError }
func (m *MySQL) GetSkipDefaultTransaction() bool    { return m.SkipDefaultTransaction }
func (m *MySQL) GetPrepareStmt() bool               { return m.PrepareStmt }
func (m *MySQL) GetDisableForeignKeyConstraintWhenMigrating() bool {
	return m.DisableForeignKeyConstraintWhenMigrating
}
func (m *MySQL) GetDisableNestedTransaction() bool { return m.DisableNestedTransaction }
func (m *MySQL) GetAllowGlobalUpdate() bool        { return m.AllowGlobalUpdate }
func (m *MySQL) GetQueryFields() bool              { return m.QueryFields }
func (m *MySQL) GetCreateBatchSize() int           { return m.CreateBatchSize }
func (m *MySQL) GetSingularTable() bool            { return m.SingularTable }
func (m *MySQL) SetCredentials(username, password string) {
	m.Username, m.Password = username, password
}
func (m *MySQL) SetHost(host string)     { m.Host = host }
func (m *MySQL) SetPort(port string)     { m.Port = port }
func (m *MySQL) SetDBName(dbName string) { m.Dbname = dbName }

func (m *MySQL) Clone() internal.Configurable {
	return &MySQL{
		ModuleName:                               m.ModuleName,
		Host:                                     m.Host,
		Port:                                     m.Port,
		Config:                                   m.Config,
		LogLevel:                                 m.LogLevel,
		SlowThreshold:                            m.SlowThreshold,
		IgnoreRecordNotFoundError:                m.IgnoreRecordNotFoundError,
		Dbname:                                   m.Dbname,
		Username:                                 m.Username,
		Password:                                 m.Password,
		MaxIdleConns:                             m.MaxIdleConns,
		MaxOpenConns:                             m.MaxOpenConns,
		ConnMaxIdleTime:                          m.ConnMaxIdleTime,
		ConnMaxLifeTime:                          m.ConnMaxLifeTime,
		SkipDefaultTransaction:                   m.SkipDefaultTransaction,
		PrepareStmt:                              m.PrepareStmt,
		DisableForeignKeyConstraintWhenMigrating: m.DisableForeignKeyConstraintWhenMigrating,
		DisableNestedTransaction:                 m.DisableNestedTransaction,
		AllowGlobalUpdate:                        m.AllowGlobalUpdate,
		QueryFields:                              m.QueryFields,
		CreateBatchSize:                          m.CreateBatchSize,
		SingularTable:                            m.SingularTable,
	}
}

func (m *MySQL) Get() interface{} { return m }
func (m *MySQL) Set(data interface{}) {
	if cfg, ok := data.(*MySQL); ok {
		*m = *cfg
	}
}
func (m *MySQL) Validate() error { return internal.ValidateStruct(m) }

// DefaultMySQL 创建默认MySQL配置
func DefaultMySQL() *MySQL {
	return &MySQL{
		ModuleName:                               "mysql",
		Host:                                     "localhost",
		Port:                                     "3306",
		Config:                                   "charset=utf8mb4&parseTime=True&loc=Local",
		LogLevel:                                 "info",
		SlowThreshold:                            100,   // 100毫秒
		IgnoreRecordNotFoundError:                false, // 不忽略record not found错误
		Dbname:                                   "test",
		Username:                                 "root",
		Password:                                 "mysql_password",
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
	}
}

// NewMySQLConfig 创建一个新的MySQL配置实例
func NewMySQLConfig(opt *MySQL) *MySQL {
	var mysqlInstance *MySQL
	internal.LockFunc(func() {
		mysqlInstance = opt
	})
	return mysqlInstance
}

// WithModuleName 设置模块名称
func (m *MySQL) WithModuleName(moduleName string) *MySQL {
	m.ModuleName = moduleName
	return m
}

// WithHost 设置数据库主机地址
func (m *MySQL) WithHost(host string) *MySQL {
	m.Host = host
	return m
}

// WithPort 设置数据库端口
func (m *MySQL) WithPort(port string) *MySQL {
	m.Port = port
	return m
}

// WithConfig 设置数据库配置参数
func (m *MySQL) WithConfig(config string) *MySQL {
	m.Config = config
	return m
}

// WithLogLevel 设置日志级别
func (m *MySQL) WithLogLevel(logLevel string) *MySQL {
	m.LogLevel = logLevel
	return m
}

// WithDbname 设置数据库名称
func (m *MySQL) WithDbname(dbname string) *MySQL {
	m.Dbname = dbname
	return m
}

// WithUsername 设置数据库用户名
func (m *MySQL) WithUsername(username string) *MySQL {
	m.Username = username
	return m
}

// WithPassword 设置数据库密码
func (m *MySQL) WithPassword(password string) *MySQL {
	m.Password = password
	return m
}

// WithMaxIdleConns 设置最大空闲连接数
func (m *MySQL) WithMaxIdleConns(maxIdleConns int) *MySQL {
	m.MaxIdleConns = maxIdleConns
	return m
}

// WithMaxOpenConns 设置最大连接数
func (m *MySQL) WithMaxOpenConns(maxOpenConns int) *MySQL {
	m.MaxOpenConns = maxOpenConns
	return m
}

// WithConnMaxIdleTime 设置连接最大空闲时间
func (m *MySQL) WithConnMaxIdleTime(connMaxIdleTime int) *MySQL {
	m.ConnMaxIdleTime = connMaxIdleTime
	return m
}

// WithConnMaxLifeTime 设置连接最大生命周期
func (m *MySQL) WithConnMaxLifeTime(connMaxLifeTime int) *MySQL {
	m.ConnMaxLifeTime = connMaxLifeTime
	return m
}
