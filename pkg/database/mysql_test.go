/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-27 00:44:04
 * @FilePath: \go-config\pkg\database\mysql_test.go
 * @Description: MySQL数据库配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMySQL_DefaultMySQL(t *testing.T) {
	mysql := DefaultMySQL()

	assert.NotNil(t, mysql)
	assert.Equal(t, "mysql", mysql.ModuleName)
	assert.Equal(t, "localhost", mysql.Host)
	assert.Equal(t, "3306", mysql.Port)
	assert.Equal(t, "charset=utf8mb4&parseTime=True&loc=Local", mysql.Config)
	assert.Equal(t, "info", mysql.LogLevel)
	assert.Equal(t, "test", mysql.Dbname)
	assert.Equal(t, "root", mysql.Username)
	assert.Equal(t, "mysql_password", mysql.Password)
	assert.Equal(t, 10, mysql.MaxIdleConns)
	assert.Equal(t, 100, mysql.MaxOpenConns)
	assert.Equal(t, 300, mysql.ConnMaxIdleTime)
	assert.Equal(t, 3600, mysql.ConnMaxLifeTime)
}

func TestMySQL_GetDBType(t *testing.T) {
	mysql := DefaultMySQL()
	assert.Equal(t, DBTypeMySQL, mysql.GetDBType())
}

func TestMySQL_GetHost(t *testing.T) {
	mysql := DefaultMySQL()
	assert.Equal(t, "localhost", mysql.GetHost())
}

func TestMySQL_SetHost(t *testing.T) {
	mysql := DefaultMySQL()
	mysql.SetHost("192.168.1.100")
	assert.Equal(t, "192.168.1.100", mysql.Host)
}

func TestMySQL_GetPort(t *testing.T) {
	mysql := DefaultMySQL()
	assert.Equal(t, "3306", mysql.GetPort())
}

func TestMySQL_SetPort(t *testing.T) {
	mysql := DefaultMySQL()
	mysql.SetPort("3307")
	assert.Equal(t, "3307", mysql.Port)
}

func TestMySQL_GetDBName(t *testing.T) {
	mysql := DefaultMySQL()
	assert.Equal(t, "test", mysql.GetDBName())
}

func TestMySQL_SetDBName(t *testing.T) {
	mysql := DefaultMySQL()
	mysql.SetDBName("production")
	assert.Equal(t, "production", mysql.Dbname)
}

func TestMySQL_GetUsername(t *testing.T) {
	mysql := DefaultMySQL()
	assert.Equal(t, "root", mysql.GetUsername())
}

func TestMySQL_GetPassword(t *testing.T) {
	mysql := DefaultMySQL()
	assert.Equal(t, "mysql_password", mysql.GetPassword())
}

func TestMySQL_SetCredentials(t *testing.T) {
	mysql := DefaultMySQL()
	mysql.SetCredentials("admin", "secret123")
	assert.Equal(t, "admin", mysql.Username)
	assert.Equal(t, "secret123", mysql.Password)
}

func TestMySQL_GetConfig(t *testing.T) {
	mysql := DefaultMySQL()
	assert.Equal(t, "charset=utf8mb4&parseTime=True&loc=Local", mysql.GetConfig())
}

func TestMySQL_GetModuleName(t *testing.T) {
	mysql := DefaultMySQL()
	assert.Equal(t, "mysql", mysql.GetModuleName())
}

func TestMySQL_Clone(t *testing.T) {
	original := DefaultMySQL()
	original.Host = "custom-host"
	original.Port = "3307"
	original.Dbname = "custom_db"

	cloned := original.Clone().(*MySQL)

	assert.Equal(t, original.Host, cloned.Host)
	assert.Equal(t, original.Port, cloned.Port)
	assert.Equal(t, original.Dbname, cloned.Dbname)

	cloned.Host = "new-host"
	assert.Equal(t, "custom-host", original.Host)
	assert.Equal(t, "new-host", cloned.Host)
}

func TestMySQL_Get(t *testing.T) {
	mysql := DefaultMySQL()
	result := mysql.Get()

	assert.NotNil(t, result)
	resultMySQL, ok := result.(*MySQL)
	assert.True(t, ok)
	assert.Equal(t, mysql, resultMySQL)
}

func TestMySQL_Set(t *testing.T) {
	mysql := DefaultMySQL()
	newMySQL := &MySQL{
		ModuleName:      "new_mysql",
		Host:            "new-host",
		Port:            "3308",
		Config:          "new-config",
		LogLevel:        "debug",
		Dbname:          "new_db",
		Username:        "new_user",
		Password:        "new_pass",
		MaxIdleConns:    20,
		MaxOpenConns:    200,
		ConnMaxIdleTime: 7200,
		ConnMaxLifeTime: 14400,
	}

	mysql.Set(newMySQL)

	assert.Equal(t, "new_mysql", mysql.ModuleName)
	assert.Equal(t, "new-host", mysql.Host)
	assert.Equal(t, "3308", mysql.Port)
	assert.Equal(t, "new-config", mysql.Config)
	assert.Equal(t, "debug", mysql.LogLevel)
	assert.Equal(t, "new_db", mysql.Dbname)
	assert.Equal(t, "new_user", mysql.Username)
	assert.Equal(t, "new_pass", mysql.Password)
	assert.Equal(t, 20, mysql.MaxIdleConns)
	assert.Equal(t, 200, mysql.MaxOpenConns)
	assert.Equal(t, 7200, mysql.ConnMaxIdleTime)
	assert.Equal(t, 14400, mysql.ConnMaxLifeTime)
}

func TestMySQL_Validate(t *testing.T) {
	mysql := DefaultMySQL()
	mysql.Password = "test123" // 设置密码以满足验证要求
	err := mysql.Validate()
	assert.NoError(t, err)
}

func TestMySQL_ChainedCalls(t *testing.T) {
	mysql := DefaultMySQL()
	mysql.SetHost("chain-host")
	mysql.SetPort("3309")
	mysql.SetDBName("chain_db")
	mysql.SetCredentials("chain_user", "chain_pass")

	assert.Equal(t, "chain-host", mysql.Host)
	assert.Equal(t, "3309", mysql.Port)
	assert.Equal(t, "chain_db", mysql.Dbname)
	assert.Equal(t, "chain_user", mysql.Username)
	assert.Equal(t, "chain_pass", mysql.Password)

	err := mysql.Validate()
	assert.NoError(t, err)
}
