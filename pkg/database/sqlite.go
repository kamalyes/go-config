/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 23:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-16 19:03:40
 * @FilePath: \go-config\pkg\database\sqlite.go
 * @Description: SQLite数据库配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package database

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-toolbox/pkg/syncx"
)

// SQLite SQLite数据库配置
type SQLite struct {
	ModuleName                               string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                                                                                                   // 模块名称
	Config                                   string `mapstructure:"config" yaml:"config" json:"config"          validate:"required"`                                                                                    // 后缀配置
	LogLevel                                 string `mapstructure:"log-level" yaml:"log-level" json:"logLevel"       validate:"required"`                                                                               // SQL 日志等级
	SlowThreshold                            int    `mapstructure:"slow-threshold" yaml:"slow-threshold" json:"slowThreshold"`                                                                                          // 慢查询阈值（毫秒）
	IgnoreRecordNotFoundError                bool   `mapstructure:"ignore-record-not-found-error" yaml:"ignore-record-not-found-error" json:"ignoreRecordNotFoundError"`                                                // 是否忽略ErrRecordNotFound错误
	DbPath                                   string `mapstructure:"db-path" yaml:"db-path" json:"dbPath"`                                                                                                               // SQLite 文件存放位置
	Vacuum                                   bool   `mapstructure:"vacuum" yaml:"vacuum" json:"vacuum"`                                                                                                                 // 是否执行清除命令
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

// 为SQLite配置实现DatabaseProvider接口
func (s *SQLite) GetDBType() DBType                  { return DBTypeSQLite }
func (s *SQLite) GetHost() string                    { return "" }       // SQLite无需host
func (s *SQLite) GetPort() string                    { return "" }       // SQLite无需port
func (s *SQLite) GetDBName() string                  { return s.DbPath } // SQLite使用文件路径
func (s *SQLite) GetUsername() string                { return "" }       // SQLite无需用户名
func (s *SQLite) GetPassword() string                { return "" }       // SQLite无需密码
func (s *SQLite) GetConfig() string                  { return s.Config }
func (s *SQLite) GetModuleName() string              { return s.ModuleName }
func (s *SQLite) GetSlowThreshold() int              { return s.SlowThreshold }
func (s *SQLite) GetIgnoreRecordNotFoundError() bool { return s.IgnoreRecordNotFoundError }
func (s *SQLite) GetSkipDefaultTransaction() bool    { return s.SkipDefaultTransaction }
func (s *SQLite) GetPrepareStmt() bool               { return s.PrepareStmt }
func (s *SQLite) GetDisableForeignKeyConstraintWhenMigrating() bool {
	return s.DisableForeignKeyConstraintWhenMigrating
}
func (s *SQLite) GetDisableNestedTransaction() bool        { return s.DisableNestedTransaction }
func (s *SQLite) GetAllowGlobalUpdate() bool               { return s.AllowGlobalUpdate }
func (s *SQLite) GetQueryFields() bool                     { return s.QueryFields }
func (s *SQLite) GetCreateBatchSize() int                  { return s.CreateBatchSize }
func (s *SQLite) GetSingularTable() bool                   { return s.SingularTable }
func (s *SQLite) SetCredentials(username, password string) {} // SQLite不支持凭证
func (s *SQLite) SetHost(host string)                      {} // SQLite不支持host
func (s *SQLite) SetPort(port string)                      {} // SQLite不支持port
func (s *SQLite) SetDBName(dbName string)                  { s.DbPath = dbName }

func (s *SQLite) Clone() internal.Configurable {
	var cloned SQLite
	if err := syncx.DeepCopy(&cloned, s); err != nil {
		// 如果深拷贝失败，返回空配置
		return &SQLite{}
	}
	return &cloned
}

func (s *SQLite) Get() interface{} { return s }
func (s *SQLite) Set(data interface{}) {
	if cfg, ok := data.(*SQLite); ok {
		*s = *cfg
	}
}
func (s *SQLite) Validate() error { return internal.ValidateStruct(s) }

// DefaultSQLite 创建默认SQLite配置
func DefaultSQLite() *SQLite {
	return &SQLite{
		ModuleName:                               "sqlite",
		DbPath:                                   "./data.db",
		Config:                                   "_foreign_keys=on",
		LogLevel:                                 "info",
		SlowThreshold:                            100,   // 100毫秒
		IgnoreRecordNotFoundError:                false, // 不忽略record not found错误
		Vacuum:                                   false,
		MaxIdleConns:                             10,
		MaxOpenConns:                             1, // SQLite 通常使用单连接
		ConnMaxIdleTime:                          300,
		ConnMaxLifeTime:                          3600,
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

// NewSQLiteConfig 创建一个新的SQLite配置实例
func NewSQLiteConfig(opt *SQLite) *SQLite {
	var sqliteInstance *SQLite
	internal.LockFunc(func() {
		sqliteInstance = opt
	})
	return sqliteInstance
}

// WithModuleName 设置模块名称
func (s *SQLite) WithModuleName(moduleName string) *SQLite {
	s.ModuleName = moduleName
	return s
}

// WithConfig 设置数据库配置参数
func (s *SQLite) WithConfig(config string) *SQLite {
	s.Config = config
	return s
}

// WithLogLevel 设置日志级别
func (s *SQLite) WithLogLevel(logLevel string) *SQLite {
	s.LogLevel = logLevel
	return s
}

// WithDbPath 设置SQLite数据库文件路径
func (s *SQLite) WithDbPath(dbPath string) *SQLite {
	s.DbPath = dbPath
	return s
}

// WithVacuum 设置是否执行清除命令
func (s *SQLite) WithVacuum(vacuum bool) *SQLite {
	s.Vacuum = vacuum
	return s
}

// WithMaxIdleConns 设置最大空闲连接数
func (s *SQLite) WithMaxIdleConns(maxIdleConns int) *SQLite {
	s.MaxIdleConns = maxIdleConns
	return s
}

// WithMaxOpenConns 设置最大连接数
func (s *SQLite) WithMaxOpenConns(maxOpenConns int) *SQLite {
	s.MaxOpenConns = maxOpenConns
	return s
}

// WithConnMaxIdleTime 设置连接最大空闲时间
func (s *SQLite) WithConnMaxIdleTime(connMaxIdleTime int) *SQLite {
	s.ConnMaxIdleTime = connMaxIdleTime
	return s
}

// WithConnMaxLifeTime 设置连接最大生命周期
func (s *SQLite) WithConnMaxLifeTime(connMaxLifeTime int) *SQLite {
	s.ConnMaxLifeTime = connMaxLifeTime
	return s
}
