/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-04-23 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-04-23 00:00:00
 * @FilePath: \go-config\pkg\database\cockroachdb_test.go
 * @Description: CockroachDB数据库配置测试
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCockroachDB_DefaultCockroachDB(t *testing.T) {
	cr := DefaultCockroachDB()

	assert.NotNil(t, cr)
	assert.Equal(t, "cockroachdb", cr.ModuleName)
	assert.Equal(t, "localhost", cr.Host)
	assert.Equal(t, "26257", cr.Port)
	assert.Equal(t, "sslmode=disable", cr.Config)
	assert.Equal(t, "info", cr.LogLevel)
	assert.Equal(t, "defaultdb", cr.Dbname)
	assert.Equal(t, "root", cr.Username)
	assert.Equal(t, "disable", cr.SSLMode)
	assert.Equal(t, "go-config", cr.ApplicationName)
	assert.Equal(t, 10, cr.MaxIdleConns)
	assert.Equal(t, 100, cr.MaxOpenConns)
	assert.Equal(t, 3, cr.MaxRetries)

	// GORM 字段
	assert.Equal(t, 100, cr.SlowThreshold)
	assert.Equal(t, false, cr.SkipDefaultTransaction)
	assert.Equal(t, true, cr.PrepareStmt)
	assert.Equal(t, true, cr.DisableForeignKeyConstraintWhenMigrating)
	assert.Equal(t, 100, cr.CreateBatchSize)
	assert.Equal(t, true, cr.SingularTable)
}

func TestCockroachDB_GetDBType(t *testing.T) {
	cr := DefaultCockroachDB()
	assert.Equal(t, DBTypeCockroachDB, cr.GetDBType())
}

func TestCockroachDB_SettersAndGetters(t *testing.T) {
	cr := DefaultCockroachDB()
	cr.SetHost("10.0.0.2")
	cr.SetPort("26258")
	cr.SetDBName("mydb")
	cr.SetCredentials("admin", "secret")

	assert.Equal(t, "10.0.0.2", cr.GetHost())
	assert.Equal(t, "26258", cr.GetPort())
	assert.Equal(t, "mydb", cr.GetDBName())
	assert.Equal(t, "admin", cr.GetUsername())
	assert.Equal(t, "secret", cr.GetPassword())
}

func TestCockroachDB_WithMethods(t *testing.T) {
	cr := DefaultCockroachDB().
		WithModuleName("crdb").
		WithHost("crdb.example.com").
		WithPort("26257").
		WithConfig("sslmode=verify-full").
		WithLogLevel("warn").
		WithDbname("production").
		WithUsername("app_user").
		WithPassword("app_pass").
		WithSSLMode("verify-full").
		WithSSLRootCert("/certs/ca.crt").
		WithSSLCert("/certs/client.crt").
		WithSSLKey("/certs/client.key").
		WithApplicationName("my-app").
		WithMaxIdleConns(20).
		WithMaxOpenConns(200).
		WithConnMaxIdleTime(600).
		WithConnMaxLifeTime(7200).
		WithMaxRetries(5)

	assert.Equal(t, "crdb", cr.ModuleName)
	assert.Equal(t, "crdb.example.com", cr.Host)
	assert.Equal(t, "verify-full", cr.SSLMode)
	assert.Equal(t, "/certs/ca.crt", cr.SSLRootCert)
	assert.Equal(t, "my-app", cr.ApplicationName)
	assert.Equal(t, 200, cr.MaxOpenConns)
	assert.Equal(t, 5, cr.MaxRetries)
}

func TestCockroachDB_Clone(t *testing.T) {
	cr := DefaultCockroachDB()
	cr.Host = "source-host"

	cloned := cr.Clone().(*CockroachDB)
	assert.Equal(t, "source-host", cloned.Host)

	cloned.Host = "changed"
	assert.Equal(t, "source-host", cr.Host, "原对象不应被修改")
}

func TestCockroachDB_Validate(t *testing.T) {
	cr := DefaultCockroachDB()
	assert.NoError(t, cr.Validate())

	// Host 为空应验证失败
	cr.Host = ""
	assert.Error(t, cr.Validate())
}

func TestCockroachDB_GetSet(t *testing.T) {
	cr := DefaultCockroachDB()
	got := cr.Get().(*CockroachDB)
	assert.Equal(t, cr, got)

	newCfg := DefaultCockroachDB()
	newCfg.Host = "new-host"
	cr.Set(newCfg)
	assert.Equal(t, "new-host", cr.Host)
}

func TestDatabase_NewDatabaseIncludesClickHouseAndCockroachDB(t *testing.T) {
	db := NewDatabase()
	assert.NotNil(t, db.NonRelational)
	assert.NotNil(t, db.NonRelational.ClickHouse)
	assert.NotNil(t, db.Relational)
	assert.NotNil(t, db.Relational.CockroachDB)

	// GetProvider 测试
	chProvider, err := db.GetProvider(DBTypeClickHouse)
	assert.NoError(t, err)
	assert.Equal(t, DBTypeClickHouse, chProvider.GetDBType())

	crProvider, err := db.GetProvider(DBTypeCockroachDB)
	assert.NoError(t, err)
	assert.Equal(t, DBTypeCockroachDB, crProvider.GetDBType())

	// ParseDBType 测试
	ch, err := ParseDBType("clickhouse")
	assert.NoError(t, err)
	assert.Equal(t, DBTypeClickHouse, ch)

	cr, err := ParseDBType("cockroach")
	assert.NoError(t, err)
	assert.Equal(t, DBTypeCockroachDB, cr)

	cr2, err := ParseDBType("cockroachdb")
	assert.NoError(t, err)
	assert.Equal(t, DBTypeCockroachDB, cr2)

	// Supported types
	supported := GetSupportedTypes()
	assert.Contains(t, supported, DBTypeClickHouse)
	assert.Contains(t, supported, DBTypeCockroachDB)

	// ProviderName
	assert.Equal(t, "ClickHouse", GetProviderName(DBTypeClickHouse))
	assert.Equal(t, "CockroachDB", GetProviderName(DBTypeCockroachDB))
}

func TestDatabase_EnableClickHouseAndCockroachDB(t *testing.T) {
	db := NewDatabase()

	db.EnableClickHouse()
	assert.Equal(t, true, db.NonRelational.Enabled)
	assert.Equal(t, string(DBTypeClickHouse), db.NonRelational.Default)

	db.EnableCockroachDB()
	assert.Equal(t, true, db.Relational.Enabled)
	assert.Equal(t, string(DBTypeCockroachDB), db.Relational.Default)
}
