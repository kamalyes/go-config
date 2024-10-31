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

// 将 MySQL 的参数转换为 map
func mysqlToMap(mysql *database.MySQL) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME":        mysql.ModuleName,
		"HOST":               mysql.Host,
		"PORT":               mysql.Port,
		"CONFIG":             mysql.Config,
		"LOG_LEVEL":          mysql.LogLevel,
		"DB_NAME":            mysql.Dbname,
		"USERNAME":           mysql.Username,
		"PASSWORD":           mysql.Password,
		"MAX_IDLE_CONNS":     mysql.MaxIdleConns,
		"MAX_OPEN_CONNS":     mysql.MaxOpenConns,
		"CONN_MAX_IDLE_TIME": mysql.ConnMaxIdleTime,
		"CONN_MAX_LIFE_TIME": mysql.ConnMaxLifeTime,
	}
}

// 验证 MySQL 的字段与期望的映射是否相等
func assertMySQLFields(t *testing.T, mysql *database.MySQL, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], mysql.ModuleName)
	assert.Equal(t, expected["HOST"], mysql.Host)
	assert.Equal(t, expected["PORT"], mysql.Port)
	assert.Equal(t, expected["CONFIG"], mysql.Config)
	assert.Equal(t, expected["LOG_LEVEL"], mysql.LogLevel)
	assert.Equal(t, expected["DB_NAME"], mysql.Dbname)
	assert.Equal(t, expected["USERNAME"], mysql.Username)
	assert.Equal(t, expected["PASSWORD"], mysql.Password)
	assert.Equal(t, expected["MAX_IDLE_CONNS"], mysql.MaxIdleConns)
	assert.Equal(t, expected["MAX_OPEN_CONNS"], mysql.MaxOpenConns)
	assert.Equal(t, expected["CONN_MAX_IDLE_TIME"], mysql.ConnMaxIdleTime)
	assert.Equal(t, expected["CONN_MAX_LIFE_TIME"], mysql.ConnMaxLifeTime)
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
