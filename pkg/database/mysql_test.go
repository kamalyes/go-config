/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-27 00:55:05
 * @FilePath: \go-config\pkg\database\mysql_test.go
 * @Description: MySQL数据库配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	// GORM 配置字段
	assert.Equal(t, 100, mysql.SlowThreshold)
	assert.Equal(t, false, mysql.IgnoreRecordNotFoundError)
	assert.Equal(t, false, mysql.SkipDefaultTransaction)
	assert.Equal(t, true, mysql.PrepareStmt)
	assert.Equal(t, true, mysql.DisableForeignKeyConstraintWhenMigrating)
	assert.Equal(t, false, mysql.DisableNestedTransaction)
	assert.Equal(t, false, mysql.AllowGlobalUpdate)
	assert.Equal(t, true, mysql.QueryFields)
	assert.Equal(t, 100, mysql.CreateBatchSize)
	assert.Equal(t, true, mysql.SingularTable)
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

func TestMySQL_GormConfigGetters(t *testing.T) {
	mysql := DefaultMySQL()

	// 测试 GORM 配置 Getter 方法
	assert.Equal(t, 100, mysql.GetSlowThreshold())
	assert.Equal(t, false, mysql.GetIgnoreRecordNotFoundError())
	assert.Equal(t, false, mysql.GetSkipDefaultTransaction())
	assert.Equal(t, true, mysql.GetPrepareStmt())
	assert.Equal(t, true, mysql.GetDisableForeignKeyConstraintWhenMigrating())
	assert.Equal(t, false, mysql.GetDisableNestedTransaction())
	assert.Equal(t, false, mysql.GetAllowGlobalUpdate())
	assert.Equal(t, true, mysql.GetQueryFields())
	assert.Equal(t, 100, mysql.GetCreateBatchSize())
	assert.Equal(t, true, mysql.GetSingularTable())
}

func TestMySQL_GormConfigCustom(t *testing.T) {
	mysql := &MySQL{
		ModuleName:                               "test_mysql",
		Host:                                     "localhost",
		Port:                                     "3306",
		Dbname:                                   "test",
		Username:                                 "root",
		Password:                                 "pass",
		SlowThreshold:                            200,
		IgnoreRecordNotFoundError:                true,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 true,
		AllowGlobalUpdate:                        true,
		QueryFields:                              false,
		CreateBatchSize:                          500,
		SingularTable:                            false,
	}

	assert.Equal(t, 200, mysql.GetSlowThreshold())
	assert.Equal(t, true, mysql.GetIgnoreRecordNotFoundError())
	assert.Equal(t, true, mysql.GetSkipDefaultTransaction())
	assert.Equal(t, false, mysql.GetPrepareStmt())
	assert.Equal(t, false, mysql.GetDisableForeignKeyConstraintWhenMigrating())
	assert.Equal(t, true, mysql.GetDisableNestedTransaction())
	assert.Equal(t, true, mysql.GetAllowGlobalUpdate())
	assert.Equal(t, false, mysql.GetQueryFields())
	assert.Equal(t, 500, mysql.GetCreateBatchSize())
	assert.Equal(t, false, mysql.GetSingularTable())
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
