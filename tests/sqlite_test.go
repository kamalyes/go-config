/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 10:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 17:05:16
 * @FilePath: \go-config\tests\sqlite_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 SQLite 配置参数
func generateSQLiteTestParams() *database.SQLite {
	return &database.SQLite{
		ModuleName:      random.RandString(10, random.CAPITAL),
		DbPath:          fmt.Sprintf("/path/to/db/%s.db", random.RandString(5, random.CAPITAL)), // 随机生成数据库文件路径
		MaxIdleConns:    random.FRandInt(1, 100),                                                // 随机最大空闲连接数
		MaxOpenConns:    random.FRandInt(1, 100),                                                // 随机最大连接数
		LogLevel:        random.RandString(5, random.CAPITAL),                                   // 随机生成日志等级
		Vacuum:          random.FRandBool(),                                                     // 随机生成是否执行清除命令
		ConnMaxIdleTime: random.FRandInt(1, 3600),                                               // 随机生成连接最大空闲时间
		ConnMaxLifeTime: random.FRandInt(1, 3600),                                               // 随机生成连接最大生命周期
	}
}

// 将 SQLite 的参数转换为 map
func sqliteToMap(sqlite *database.SQLite) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME":        sqlite.ModuleName,
		"DB_PATH":            sqlite.DbPath,
		"MAX_IDLE_CONNS":     sqlite.MaxIdleConns,
		"MAX_OPEN_CONNS":     sqlite.MaxOpenConns,
		"LOG_LEVEL":          sqlite.LogLevel,
		"VACUUM":             sqlite.Vacuum,
		"CONN_MAX_IDLE_TIME": sqlite.ConnMaxIdleTime,
		"CONN_MAX_LIFE_TIME": sqlite.ConnMaxLifeTime,
	}
}

// 验证 SQLite 的字段与期望的映射是否相等
func assertSQLiteFields(t *testing.T, sqlite *database.SQLite, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], sqlite.ModuleName)
	assert.Equal(t, expected["DB_PATH"], sqlite.DbPath)
	assert.Equal(t, expected["MAX_IDLE_CONNS"], sqlite.MaxIdleConns)
	assert.Equal(t, expected["MAX_OPEN_CONNS"], sqlite.MaxOpenConns)
	assert.Equal(t, expected["LOG_LEVEL"], sqlite.LogLevel)
	assert.Equal(t, expected["VACUUM"], sqlite.Vacuum)
	assert.Equal(t, expected["CONN_MAX_IDLE_TIME"], sqlite.ConnMaxIdleTime)
	assert.Equal(t, expected["CONN_MAX_LIFE_TIME"], sqlite.ConnMaxLifeTime)
}

func TestSQLiteClone(t *testing.T) {
	params := generateSQLiteTestParams()
	original := database.NewSQLite(params)
	cloned := original.Clone().(*database.SQLite)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestSQLiteSet(t *testing.T) {
	oldParams := generateSQLiteTestParams()
	newParams := generateSQLiteTestParams()

	sqliteInstance := database.NewSQLite(oldParams)
	newConfig := database.NewSQLite(newParams)

	sqliteInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, sqliteInstance.ModuleName)
	assert.Equal(t, newParams.DbPath, sqliteInstance.DbPath)
	assert.Equal(t, newParams.MaxIdleConns, sqliteInstance.MaxIdleConns)
	assert.Equal(t, newParams.MaxOpenConns, sqliteInstance.MaxOpenConns)
	assert.Equal(t, newParams.LogLevel, sqliteInstance.LogLevel)
	assert.Equal(t, newParams.Vacuum, sqliteInstance.Vacuum)
	assert.Equal(t, newParams.ConnMaxIdleTime, sqliteInstance.ConnMaxIdleTime)
	assert.Equal(t, newParams.ConnMaxLifeTime, sqliteInstance.ConnMaxLifeTime)
}
