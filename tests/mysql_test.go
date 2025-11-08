/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 17:02:02
 * @FilePath: \go-config\tests\mysql_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

func generateMySQLTestParams() *database.MySQL {
	return &database.MySQL{
		ModuleName:      random.RandString(10, random.CAPITAL),
		Host:            random.RandString(5, random.CAPITAL) + ".database.local", // 随机生成主机名
		Port:            random.RandString(5, random.NUMBER),                      // 随机生成端口
		Config:          "charset=utf8mb4&parseTime=True&loc=Local",               // 默认配置
		LogLevel:        random.RandString(5, random.CAPITAL),                     // 随机生成日志等级
		Dbname:          random.RandString(8, random.CAPITAL),                     // 随机生成数据库名称
		Username:        random.RandString(8, random.CAPITAL),                     // 随机生成用户名
		Password:        random.RandString(16, random.CAPITAL),                    // 随机生成密码
		MaxIdleConns:    random.FRandInt(1, 100),                                  // 随机生成最大空闲连接数
		MaxOpenConns:    random.FRandInt(1, 100),                                  // 随机生成最大连接数
		ConnMaxIdleTime: random.FRandInt(30, 300),                                 // 随机生成最大空闲时间（30到300秒）
		ConnMaxLifeTime: random.FRandInt(30, 300),                                 // 随机生成最大生命周期（30到300秒）
	}
}

func TestMySQLClone(t *testing.T) {
	params := generateMySQLTestParams()
	original := database.NewMySQL(params)
	cloned := original.Clone().(*database.MySQL)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestMySQLSet(t *testing.T) {
	oldParams := generateMySQLTestParams()
	newParams := generateMySQLTestParams()

	mysqlInstance := database.NewMySQL(oldParams)
	newConfig := database.NewMySQL(newParams)

	mysqlInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, mysqlInstance.ModuleName)
	assert.Equal(t, newParams.Host, mysqlInstance.Host)
	assert.Equal(t, newParams.Port, mysqlInstance.Port)
	assert.Equal(t, newParams.Config, mysqlInstance.Config)
	assert.Equal(t, newParams.LogLevel, mysqlInstance.LogLevel)
	assert.Equal(t, newParams.Dbname, mysqlInstance.Dbname)
	assert.Equal(t, newParams.Username, mysqlInstance.Username)
	assert.Equal(t, newParams.Password, mysqlInstance.Password)
	assert.Equal(t, newParams.MaxIdleConns, mysqlInstance.MaxIdleConns)
	assert.Equal(t, newParams.MaxOpenConns, mysqlInstance.MaxOpenConns)
	assert.Equal(t, newParams.ConnMaxIdleTime, mysqlInstance.ConnMaxIdleTime)
	assert.Equal(t, newParams.ConnMaxLifeTime, mysqlInstance.ConnMaxLifeTime)
}

// TestMySQLDefault 测试默认配置
func TestMySQLDefault(t *testing.T) {
	defaultConfig := database.DefaultMySQL()
	
	// 检查默认值
	assert.Equal(t, "mysql", defaultConfig.ModuleName)
	assert.Equal(t, "127.0.0.1", defaultConfig.Host)
	assert.Equal(t, "3306", defaultConfig.Port)
	assert.Equal(t, "charset=utf8mb4&parseTime=True&loc=Local", defaultConfig.Config)
	assert.Equal(t, "silent", defaultConfig.LogLevel)
	assert.Equal(t, 10, defaultConfig.MaxIdleConns)
	assert.Equal(t, 100, defaultConfig.MaxOpenConns)
	assert.Equal(t, 3600, defaultConfig.ConnMaxIdleTime)
	assert.Equal(t, 7200, defaultConfig.ConnMaxLifeTime)
}

// TestMySQLDefaultPointer 测试默认配置指针
func TestMySQLDefaultPointer(t *testing.T) {
	config := database.Default()
	
	assert.NotNil(t, config)
	assert.Equal(t, "mysql", config.ModuleName)
	assert.Equal(t, "127.0.0.1", config.Host)
	assert.Equal(t, "3306", config.Port)
}

// TestMySQLChainMethods 测试链式方法
func TestMySQLChainMethods(t *testing.T) {
	config := database.Default().
		WithModuleName("test-db").
		WithHost("192.168.1.100").
		WithPort("3307").
		WithDbname("testdb").
		WithUsername("testuser").
		WithPassword("testpass").
		WithMaxIdleConns(20).
		WithMaxOpenConns(200).
		WithConnMaxIdleTime(7200).
		WithConnMaxLifeTime(14400)
	
	assert.Equal(t, "test-db", config.ModuleName)
	assert.Equal(t, "192.168.1.100", config.Host)
	assert.Equal(t, "3307", config.Port)
	assert.Equal(t, "testdb", config.Dbname)
	assert.Equal(t, "testuser", config.Username)
	assert.Equal(t, "testpass", config.Password)
	assert.Equal(t, 20, config.MaxIdleConns)
	assert.Equal(t, 200, config.MaxOpenConns)
	assert.Equal(t, 7200, config.ConnMaxIdleTime)
	assert.Equal(t, 14400, config.ConnMaxLifeTime)
}

// TestMySQLChainMethodsReturnPointer 测试链式方法返回指针
func TestMySQLChainMethodsReturnPointer(t *testing.T) {
	config1 := database.Default()
	config2 := config1.WithHost("localhost")
	
	// 应该返回同一个实例
	assert.Same(t, config1, config2)
	assert.Equal(t, "localhost", config1.Host)
}
