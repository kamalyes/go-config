/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 10:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 00:21:52
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
	assert.Equal(t, newParams.ConnMaxIdleTime, sqliteInstance.ConnMaxIdleTime)
	assert.Equal(t, newParams.ConnMaxLifeTime, sqliteInstance.ConnMaxLifeTime)
	assert.Equal(t, newParams.Vacuum, sqliteInstance.Vacuum)
}

// TestSQLiteDefault 测试默认配置
func TestSQLiteDefault(t *testing.T) {
	defaultConfig := database.DefaultSQLite()
	
	// 检查默认值
	assert.Equal(t, "sqlite", defaultConfig.ModuleName)
	assert.Equal(t, "./data.db", defaultConfig.DbPath)
	assert.Equal(t, 10, defaultConfig.MaxIdleConns)
	assert.Equal(t, 100, defaultConfig.MaxOpenConns)
	assert.Equal(t, "silent", defaultConfig.LogLevel)
	assert.Equal(t, 3600, defaultConfig.ConnMaxIdleTime)
	assert.Equal(t, 7200, defaultConfig.ConnMaxLifeTime)
	assert.Equal(t, false, defaultConfig.Vacuum)
}

// TestSQLiteDefaultPointer 测试默认配置指针
func TestSQLiteDefaultPointer(t *testing.T) {
	config := database.DefaultSQLiteConfig()
	
	assert.NotNil(t, config)
	assert.Equal(t, "sqlite", config.ModuleName)
	assert.Equal(t, "./data.db", config.DbPath)
}

// TestSQLiteChainMethods 测试链式方法
func TestSQLiteChainMethods(t *testing.T) {
	config := database.DefaultSQLiteConfig().
		WithModuleName("app-db").
		WithDbPath("./app.db").
		WithMaxIdleConns(20).
		WithMaxOpenConns(200).
		WithLogLevel("info").
		WithConnMaxIdleTime(7200).
		WithConnMaxLifeTime(14400).
		WithVacuum(true)
	
	assert.Equal(t, "app-db", config.ModuleName)
	assert.Equal(t, "./app.db", config.DbPath)
	assert.Equal(t, 20, config.MaxIdleConns)
	assert.Equal(t, 200, config.MaxOpenConns)
	assert.Equal(t, "info", config.LogLevel)
	assert.Equal(t, 7200, config.ConnMaxIdleTime)
	assert.Equal(t, 14400, config.ConnMaxLifeTime)
	assert.Equal(t, true, config.Vacuum)
}

// TestSQLiteChainMethodsReturnPointer 测试链式方法返回指针
func TestSQLiteChainMethodsReturnPointer(t *testing.T) {
	config1 := database.DefaultSQLiteConfig()
	config2 := config1.WithDbPath("./test.db")
	
	// 应该返回同一个实例
	assert.Same(t, config1, config2)
	assert.Equal(t, "./test.db", config1.DbPath)
}
