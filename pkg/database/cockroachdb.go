/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-03 20:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-04-28 11:16:58
 * @FilePath: \go-config\pkg\database\cockroachdb.go
 * @Description: CockroachDB统一配置管理
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package database

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-toolbox/pkg/syncx"
)

type CockroachDB struct {
	ModuleName                               string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`
	Host                                     string `mapstructure:"host" yaml:"host" json:"host" validate:"required"`
	Port                                     string `mapstructure:"port" yaml:"port" json:"port" validate:"required"`
	Config                                   string `mapstructure:"config" yaml:"config" json:"config" validate:"required"`
	LogLevel                                 string `mapstructure:"log-level" yaml:"log-level" json:"logLevel" validate:"required"`
	SlowThreshold                            int    `mapstructure:"slow-threshold" yaml:"slow-threshold" json:"slowThreshold"`
	IgnoreRecordNotFoundError                bool   `mapstructure:"ignore-record-not-found-error" yaml:"ignore-record-not-found-error" json:"ignoreRecordNotFoundError"`
	Dbname                                   string `mapstructure:"db-name" yaml:"db-name" json:"dbName" validate:"required"`
	Username                                 string `mapstructure:"username" yaml:"username" json:"username" validate:"required"`
	Password                                 string `mapstructure:"password" yaml:"password" json:"password" validate:"required"`
	MaxIdleConns                             int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns" json:"maxIdleConns" validate:"min=0"`
	MaxOpenConns                             int    `mapstructure:"max-open-conns" yaml:"max-open-conns" json:"maxOpenConns" validate:"min=0"`
	ConnMaxIdleTime                          int    `mapstructure:"conn-max-idle-time" yaml:"conn-max-idle-time" json:"connMaxIdleTime" validate:"min=0"`
	ConnMaxLifeTime                          int    `mapstructure:"conn-max-life-time" yaml:"conn-max-life-time" json:"connMaxLifeTime" validate:"min=0"`
	SkipDefaultTransaction                   bool   `mapstructure:"skip-default-transaction" yaml:"skip-default-transaction" json:"skipDefaultTransaction"`
	PrepareStmt                              bool   `mapstructure:"prepare-stmt" yaml:"prepare-stmt" json:"prepareStmt"`
	DisableForeignKeyConstraintWhenMigrating bool   `mapstructure:"disable-foreign-key-constraint-when-migrating" yaml:"disable-foreign-key-constraint-when-migrating" json:"disableForeignKeyConstraintWhenMigrating"`
	DisableNestedTransaction                 bool   `mapstructure:"disable-nested-transaction" yaml:"disable-nested-transaction" json:"disableNestedTransaction"`
	AllowGlobalUpdate                        bool   `mapstructure:"allow-global-update" yaml:"allow-global-update" json:"allowGlobalUpdate"`
	QueryFields                              bool   `mapstructure:"query-fields" yaml:"query-fields" json:"queryFields"`
	CreateBatchSize                          int    `mapstructure:"create-batch-size" yaml:"create-batch-size" json:"createBatchSize"`
	SingularTable                            bool   `mapstructure:"singular-table" yaml:"singular-table" json:"singularTable"`
}

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

func DefaultCockroachDB() *CockroachDB {
	return &CockroachDB{
		ModuleName:                               "cockroachdb",
		Host:                                     "localhost",
		Port:                                     "26257",
		Config:                                   "sslmode=disable",
		LogLevel:                                 "info",
		SlowThreshold:                            100,
		IgnoreRecordNotFoundError:                false,
		Dbname:                                   "defaultdb",
		Username:                                 "root",
		Password:                                 "",
		MaxIdleConns:                             10,
		MaxOpenConns:                             100,
		ConnMaxIdleTime:                          300,
		ConnMaxLifeTime:                          3600,
		SkipDefaultTransaction:                   false,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              true,
		CreateBatchSize:                          100,
		SingularTable:                            true,
	}
}

func NewCockroachDBConfig(opt *CockroachDB) *CockroachDB {
	var cockroachdbInstance *CockroachDB
	internal.LockFunc(func() {
		cockroachdbInstance = opt
	})
	return cockroachdbInstance
}

func (c *CockroachDB) WithModuleName(moduleName string) *CockroachDB {
	c.ModuleName = moduleName
	return c
}

func (c *CockroachDB) WithHost(host string) *CockroachDB {
	c.Host = host
	return c
}

func (c *CockroachDB) WithPort(port string) *CockroachDB {
	c.Port = port
	return c
}

func (c *CockroachDB) WithConfig(config string) *CockroachDB {
	c.Config = config
	return c
}

func (c *CockroachDB) WithLogLevel(logLevel string) *CockroachDB {
	c.LogLevel = logLevel
	return c
}

func (c *CockroachDB) WithDbname(dbname string) *CockroachDB {
	c.Dbname = dbname
	return c
}

func (c *CockroachDB) WithUsername(username string) *CockroachDB {
	c.Username = username
	return c
}

func (c *CockroachDB) WithPassword(password string) *CockroachDB {
	c.Password = password
	return c
}

func (c *CockroachDB) WithMaxIdleConns(maxIdleConns int) *CockroachDB {
	c.MaxIdleConns = maxIdleConns
	return c
}

func (c *CockroachDB) WithMaxOpenConns(maxOpenConns int) *CockroachDB {
	c.MaxOpenConns = maxOpenConns
	return c
}

func (c *CockroachDB) WithConnMaxIdleTime(connMaxIdleTime int) *CockroachDB {
	c.ConnMaxIdleTime = connMaxIdleTime
	return c
}

func (c *CockroachDB) WithConnMaxLifeTime(connMaxLifeTime int) *CockroachDB {
	c.ConnMaxLifeTime = connMaxLifeTime
	return c
}
