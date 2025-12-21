/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\database\sqlite_test.go
 * @Description: SQLite数据库配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite_DefaultSQLite(t *testing.T) {
	sqlite := DefaultSQLite()

	assert.NotNil(t, sqlite)
	assert.Equal(t, "sqlite", sqlite.ModuleName)
	assert.Equal(t, "./data.db", sqlite.DbPath)
	assert.Equal(t, "info", sqlite.LogLevel)
	assert.Equal(t, 10, sqlite.MaxIdleConns)
	assert.Equal(t, 1, sqlite.MaxOpenConns)
	// GORM 配置字段
	assert.Equal(t, 100, sqlite.SlowThreshold)
	assert.Equal(t, false, sqlite.IgnoreRecordNotFoundError)
	assert.Equal(t, false, sqlite.SkipDefaultTransaction)
	assert.Equal(t, true, sqlite.PrepareStmt)
	assert.Equal(t, true, sqlite.DisableForeignKeyConstraintWhenMigrating)
	assert.Equal(t, false, sqlite.DisableNestedTransaction)
	assert.Equal(t, false, sqlite.AllowGlobalUpdate)
	assert.Equal(t, true, sqlite.QueryFields)
	assert.Equal(t, 100, sqlite.CreateBatchSize)
	assert.Equal(t, true, sqlite.SingularTable)
}

func TestSQLite_GetDBType(t *testing.T) {
	sqlite := DefaultSQLite()
	assert.Equal(t, DBTypeSQLite, sqlite.GetDBType())
}

func TestSQLite_GetDBName(t *testing.T) {
	sqlite := DefaultSQLite()
	assert.Equal(t, "./data.db", sqlite.GetDBName())
}

func TestSQLite_SetDBName(t *testing.T) {
	sqlite := DefaultSQLite()
	sqlite.SetDBName("/custom/path/db.sqlite")
	assert.Equal(t, "/custom/path/db.sqlite", sqlite.DbPath)
}

func TestSQLite_GetModuleName(t *testing.T) {
	sqlite := DefaultSQLite()
	assert.Equal(t, "sqlite", sqlite.GetModuleName())
}

func TestSQLite_Clone(t *testing.T) {
	original := DefaultSQLite()
	original.DbPath = "/test/db.sqlite"
	original.MaxIdleConns = 20

	cloned := original.Clone().(*SQLite)

	assert.Equal(t, original.DbPath, cloned.DbPath)
	assert.Equal(t, original.MaxIdleConns, cloned.MaxIdleConns)

	cloned.DbPath = "/new/db.sqlite"
	assert.Equal(t, "/test/db.sqlite", original.DbPath)
	assert.Equal(t, "/new/db.sqlite", cloned.DbPath)
}

func TestSQLite_Get(t *testing.T) {
	sqlite := DefaultSQLite()
	result := sqlite.Get()

	assert.NotNil(t, result)
	resultSQLite, ok := result.(*SQLite)
	assert.True(t, ok)
	assert.Equal(t, sqlite, resultSQLite)
}

func TestSQLite_Set(t *testing.T) {
	sqlite := DefaultSQLite()
	newSQLite := &SQLite{
		ModuleName:      "new_sqlite",
		DbPath:          "/new/path/db.sqlite",
		LogLevel:        "debug",
		MaxIdleConns:    15,
		MaxOpenConns:    150,
		ConnMaxIdleTime: 5000,
		ConnMaxLifeTime: 10000,
	}

	sqlite.Set(newSQLite)

	assert.Equal(t, "new_sqlite", sqlite.ModuleName)
	assert.Equal(t, "/new/path/db.sqlite", sqlite.DbPath)
	assert.Equal(t, "debug", sqlite.LogLevel)
	assert.Equal(t, 15, sqlite.MaxIdleConns)
	assert.Equal(t, 150, sqlite.MaxOpenConns)
}

func TestSQLite_Validate(t *testing.T) {
	sqlite := DefaultSQLite()
	err := sqlite.Validate()
	assert.NoError(t, err)
}

func TestSQLite_ChainedCalls(t *testing.T) {
	sqlite := DefaultSQLite()
	sqlite.SetDBName("/chain/db.sqlite")

	assert.Equal(t, "/chain/db.sqlite", sqlite.DbPath)

	err := sqlite.Validate()
	assert.NoError(t, err)
}

func TestSQLite_GormConfigGetters(t *testing.T) {
	sqlite := DefaultSQLite()

	// 测试 GORM 配置 Getter 方法
	assert.Equal(t, 100, sqlite.GetSlowThreshold())
	assert.Equal(t, false, sqlite.GetIgnoreRecordNotFoundError())
	assert.Equal(t, false, sqlite.GetSkipDefaultTransaction())
	assert.Equal(t, true, sqlite.GetPrepareStmt())
	assert.Equal(t, true, sqlite.GetDisableForeignKeyConstraintWhenMigrating())
	assert.Equal(t, false, sqlite.GetDisableNestedTransaction())
	assert.Equal(t, false, sqlite.GetAllowGlobalUpdate())
	assert.Equal(t, true, sqlite.GetQueryFields())
	assert.Equal(t, 100, sqlite.GetCreateBatchSize())
	assert.Equal(t, true, sqlite.GetSingularTable())
}

func TestSQLite_GormConfigCustom(t *testing.T) {
	sqlite := &SQLite{
		ModuleName:                               "test_sqlite",
		DbPath:                                   "/test/db.sqlite",
		SlowThreshold:                            150,
		IgnoreRecordNotFoundError:                true,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 true,
		AllowGlobalUpdate:                        true,
		QueryFields:                              false,
		CreateBatchSize:                          200,
		SingularTable:                            false,
	}

	assert.Equal(t, 150, sqlite.GetSlowThreshold())
	assert.Equal(t, true, sqlite.GetIgnoreRecordNotFoundError())
	assert.Equal(t, true, sqlite.GetSkipDefaultTransaction())
	assert.Equal(t, false, sqlite.GetPrepareStmt())
	assert.Equal(t, false, sqlite.GetDisableForeignKeyConstraintWhenMigrating())
	assert.Equal(t, true, sqlite.GetDisableNestedTransaction())
	assert.Equal(t, true, sqlite.GetAllowGlobalUpdate())
	assert.Equal(t, false, sqlite.GetQueryFields())
	assert.Equal(t, 200, sqlite.GetCreateBatchSize())
	assert.Equal(t, false, sqlite.GetSingularTable())
}
