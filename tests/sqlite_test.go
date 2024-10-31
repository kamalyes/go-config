/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 10:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:05:55
 * @FilePath: \go-config\database\postgre_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/database"
)

// 公共测试数据

func getInvalidSqlLiteConfig() *database.SQLite {
	return &database.SQLite{
		ModuleName:      "testSQLiteModule",
		DbPath:          "/path/to/sqlite.db",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		LogLevel:        "info",
		Vacuum:          true,
		ConnMaxIdleTime: 30,
		ConnMaxLifeTime: 3600,
	}
}

// 公共断言函数
func assertEqual(t *testing.T, expected, actual interface{}, message string) {
	if expected != actual {
		t.Errorf("%s: expected %v, got %v", message, expected, actual)
	}
}

func TestSQLite_Validate(t *testing.T) {
	testSQLiteData := getInvalidSqlLiteConfig()
	// 测试有效配置
	if err := testSQLiteData.Validate(); err != nil {
		t.Errorf("SQLite Validate failed: %v", err)
	}

	// 测试无效配置
	invalidData := &database.SQLite{}
	if err := invalidData.Validate(); err == nil {
		t.Error("SQLite Validate should have failed for empty configuration")
	}
}

func TestSQLite_ToMap(t *testing.T) {
	testSQLiteData := getInvalidSqlLiteConfig()
	dataMap := testSQLiteData.ToMap()
	assertEqual(t, testSQLiteData.ModuleName, dataMap["moduleName"], "SQLite ToMap failed for ModuleName")
	assertEqual(t, testSQLiteData.DbPath, dataMap["dbpath"], "SQLite ToMap failed for DbPath")
	assertEqual(t, testSQLiteData.MaxIdleConns, dataMap["maxIdleConns"], "SQLite ToMap failed for MaxIdleConns")
	assertEqual(t, testSQLiteData.MaxOpenConns, dataMap["maxOpenConns"], "SQLite ToMap failed for MaxOpenConns")
	assertEqual(t, testSQLiteData.LogLevel, dataMap["logLevel"], "SQLite ToMap failed for LogLevel")
	assertEqual(t, testSQLiteData.Vacuum, dataMap["vacuum"], "SQLite ToMap failed for Vacuum")
	assertEqual(t, testSQLiteData.ConnMaxIdleTime, dataMap["connMaxIdleTime"], "SQLite ToMap failed for ConnMaxIdleTime")
	assertEqual(t, testSQLiteData.ConnMaxLifeTime, dataMap["connMaxLifeTime"], "SQLite ToMap failed for ConnMaxLifeTime")
}

func TestSQLite_FromMap(t *testing.T) {
	newData := map[string]interface{}{
		"moduleName":      "newSQLiteModule",
		"dbpath":          "/new/path/to/sqlite.db",
		"maxIdleConns":    20,
		"maxOpenConns":    200,
		"logLevel":        "debug",
		"vacuum":          false,
		"connMaxIdleTime": 60,
		"connMaxLifeTime": 7200,
	}
	testSQLiteData := getInvalidSqlLiteConfig()
	testSQLiteData.FromMap(newData)

	// 验证更新后的值
	assertEqual(t, "newSQLiteModule", testSQLiteData.ModuleName, "SQLite FromMap failed for ModuleName")
	assertEqual(t, "/new/path/to/sqlite.db", testSQLiteData.DbPath, "SQLite FromMap failed for DbPath")
	assertEqual(t, 20, testSQLiteData.MaxIdleConns, "SQLite FromMap failed for MaxIdleConns")
	assertEqual(t, 200, testSQLiteData.MaxOpenConns, "SQLite FromMap failed for MaxOpenConns")
	assertEqual(t, "debug", testSQLiteData.LogLevel, "SQLite FromMap failed for LogLevel")
	assertEqual(t, false, testSQLiteData.Vacuum, "SQLite FromMap failed for Vacuum")
	assertEqual(t, 60, testSQLiteData.ConnMaxIdleTime, "SQLite FromMap failed for ConnMaxIdleTime")
	assertEqual(t, 7200, testSQLiteData.ConnMaxLifeTime, "SQLite FromMap failed for ConnMaxLifeTime")
}

func TestSQLite_Clone(t *testing.T) {
	testSQLiteData := getInvalidSqlLiteConfig()
	clonedSQLite := testSQLiteData.Clone()
	assertEqual(t, testSQLiteData.ModuleName, clonedSQLite.(*database.SQLite).ModuleName, "SQLite Clone failed for ModuleName")
	assertEqual(t, testSQLiteData.DbPath, clonedSQLite.(*database.SQLite).DbPath, "SQLite Clone failed for DbPath")
	assertEqual(t, testSQLiteData.MaxIdleConns, clonedSQLite.(*database.SQLite).MaxIdleConns, "SQLite Clone failed for MaxIdleConns")
	assertEqual(t, testSQLiteData.MaxOpenConns, clonedSQLite.(*database.SQLite).MaxOpenConns, "SQLite Clone failed for MaxOpenConns")
	assertEqual(t, testSQLiteData.LogLevel, clonedSQLite.(*database.SQLite).LogLevel, "SQLite Clone failed for LogLevel")
	assertEqual(t, testSQLiteData.Vacuum, clonedSQLite.(*database.SQLite).Vacuum, "SQLite Clone failed for Vacuum")
	assertEqual(t, testSQLiteData.ConnMaxIdleTime, clonedSQLite.(*database.SQLite).ConnMaxIdleTime, "SQLite Clone failed for ConnMaxIdleTime")
	assertEqual(t, testSQLiteData.ConnMaxLifeTime, clonedSQLite.(*database.SQLite).ConnMaxLifeTime, "SQLite Clone failed for ConnMaxLifeTime")
}

func TestSQLite_Get(t *testing.T) {
	testSQLiteData := getInvalidSqlLiteConfig()
	data := testSQLiteData.Get().(*database.SQLite)
	assertEqual(t, testSQLiteData.ModuleName, data.ModuleName, "SQLite Get failed for ModuleName")
	assertEqual(t, testSQLiteData.DbPath, data.DbPath, "SQLite Get failed for DbPath")
}

func TestSQLite_Set(t *testing.T) {
	newConfig := &database.SQLite{
		ModuleName:      "updatedSQLiteModule",
		DbPath:          "/updated/path/to/sqlite.db",
		MaxIdleConns:    15,
		MaxOpenConns:    150,
		LogLevel:        "warn",
		Vacuum:          false,
		ConnMaxIdleTime: 45,
		ConnMaxLifeTime: 7200,
	}
	testSQLiteData := getInvalidSqlLiteConfig()
	testSQLiteData.Set(newConfig)

	// 验证更新后的值
	assertEqual(t, "updatedSQLiteModule", testSQLiteData.ModuleName, "SQLite Set failed for ModuleName")
	assertEqual(t, "/updated/path/to/sqlite.db", testSQLiteData.DbPath, "SQLite Set failed for DbPath")
	assertEqual(t, 15, testSQLiteData.MaxIdleConns, "SQLite Set failed for MaxIdleConns")
	assertEqual(t, 150, testSQLiteData.MaxOpenConns, "SQLite Set failed for MaxOpenConns")
	assertEqual(t, "warn", testSQLiteData.LogLevel, "SQLite Set failed for LogLevel")
	assertEqual(t, false, testSQLiteData.Vacuum, "SQLite Set failed for Vacuum")
	assertEqual(t, 45, testSQLiteData.ConnMaxIdleTime, "SQLite Set failed for ConnMaxIdleTime")
	assertEqual(t, 7200, testSQLiteData.ConnMaxLifeTime, "SQLite Set failed for ConnMaxLifeTime")
}
