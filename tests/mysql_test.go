/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 13:21:27
 * @FilePath: \go-config\tests\mysql_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/database"
	"github.com/stretchr/testify/assert"
)

// 公共函数：生成测试数据
func generateTestMySQL() *database.MySQL {
	return database.NewMySQL("testModule", "localhost", "3306", "testdb", "user", "password", 10, 100, 30, 300)
}

// 公共函数：生成无效的 MySQL 数据
func generateInvalidMySQL() *database.MySQL {
	return database.NewMySQL("", "", "", "", "", "", -1, -1, -1, -1) // 无效值
}

// 公共函数：生成测试映射数据
func generateTestDataMap() map[string]interface{} {
	return map[string]interface{}{
		"moduleName":      "testModule",
		"host":            "localhost",
		"port":            "3306",
		"dbname":          "testdb",
		"username":        "user",
		"password":        "password",
		"maxIdleConns":    10,
		"maxOpenConns":    100,
		"connMaxIdleTime": 30,
		"connMaxLifeTime": 300,
	}
}

// 公共函数：校验 MySQL 结构体的字段
func assertMySQLFields(t *testing.T, mysql *database.MySQL) {
	assert.Equal(t, "testModule", mysql.ModuleName)
	assert.Equal(t, "localhost", mysql.Host)
	assert.Equal(t, "3306", mysql.Port)
	assert.Equal(t, "testdb", mysql.Dbname)
	assert.Equal(t, "user", mysql.Username)
	assert.Equal(t, "password", mysql.Password)
	assert.Equal(t, 10, mysql.MaxIdleConns)
	assert.Equal(t, 100, mysql.MaxOpenConns)
	assert.Equal(t, 30, mysql.ConnMaxIdleTime)
	assert.Equal(t, 300, mysql.ConnMaxLifeTime)
}

// 公共函数：校验映射数据
func assertMapFields(t *testing.T, m map[string]interface{}) {
	assert.Equal(t, "testModule", m["moduleName"])
	assert.Equal(t, "localhost", m["host"])
	assert.Equal(t, "3306", m["port"])
	assert.Equal(t, "testdb", m["dbname"])
	assert.Equal(t, "user", m["username"])
	assert.Equal(t, "password", m["password"])
	assert.Equal(t, 10, m["maxIdleConns"])
	assert.Equal(t, 100, m["maxOpenConns"])
	assert.Equal(t, 30, m["connMaxIdleTime"])
	assert.Equal(t, 300, m["connMaxLifeTime"])
}

func TestNewMySQL(t *testing.T) {
	mysql := generateTestMySQL()

	assert.NotNil(t, mysql)
	assertMySQLFields(t, mysql)
}

func TestMySQLToMap(t *testing.T) {
	mysql := generateTestMySQL()
	m := mysql.ToMap()

	assertMapFields(t, m)
}

func TestMySQLFromMap(t *testing.T) {
	mysql := database.NewMySQL("", "", "", "", "", "", 0, 0, 0, 0)
	data := generateTestDataMap() // 使用公共函数生成测试映射数据

	mysql.FromMap(data)

	assertMySQLFields(t, mysql)
}

func TestMySQLClone(t *testing.T) {
	mysql := generateTestMySQL()
	clone := mysql.Clone().(*database.MySQL)

	assert.Equal(t, mysql, clone)
	assert.NotSame(t, mysql, clone) // 确保是不同的实例
}

func TestMySQLValidate(t *testing.T) {
	mysql := generateTestMySQL()
	err := mysql.Validate()
	assert.NoError(t, err)

	// 测试无效情况
	invalidMySQL := generateInvalidMySQL()
	err = invalidMySQL.Validate()
	assert.Error(t, err)
}
