/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\database\postgresql_test.go
 * @Description: PostgreSQL数据库配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgreSQL_DefaultPostgreSQL(t *testing.T) {
	pg := DefaultPostgreSQL()

	assert.NotNil(t, pg)
	assert.Equal(t, "postgresql", pg.ModuleName)
	assert.Equal(t, "localhost", pg.Host)
	assert.Equal(t, "5432", pg.Port)
	assert.Equal(t, "sslmode=disable", pg.Config)
	assert.Equal(t, "info", pg.LogLevel)
	assert.Equal(t, "postgres", pg.Dbname)
	assert.Equal(t, "postgres", pg.Username)
	assert.Equal(t, "postgres_password", pg.Password)
	assert.Equal(t, 10, pg.MaxIdleConns)
	assert.Equal(t, 100, pg.MaxOpenConns)
}

func TestPostgreSQL_GetDBType(t *testing.T) {
	pg := DefaultPostgreSQL()
	assert.Equal(t, DBTypePostgreSQL, pg.GetDBType())
}

func TestPostgreSQL_GetHost(t *testing.T) {
	pg := DefaultPostgreSQL()
	assert.Equal(t, "localhost", pg.GetHost())
}

func TestPostgreSQL_SetHost(t *testing.T) {
	pg := DefaultPostgreSQL()
	pg.SetHost("192.168.1.200")
	assert.Equal(t, "192.168.1.200", pg.Host)
}

func TestPostgreSQL_GetPort(t *testing.T) {
	pg := DefaultPostgreSQL()
	assert.Equal(t, "5432", pg.GetPort())
}

func TestPostgreSQL_SetPort(t *testing.T) {
	pg := DefaultPostgreSQL()
	pg.SetPort("5433")
	assert.Equal(t, "5433", pg.Port)
}

func TestPostgreSQL_GetDBName(t *testing.T) {
	pg := DefaultPostgreSQL()
	assert.Equal(t, "postgres", pg.GetDBName())
}

func TestPostgreSQL_SetDBName(t *testing.T) {
	pg := DefaultPostgreSQL()
	pg.SetDBName("myapp")
	assert.Equal(t, "myapp", pg.Dbname)
}

func TestPostgreSQL_GetUsername(t *testing.T) {
	pg := DefaultPostgreSQL()
	assert.Equal(t, "postgres", pg.GetUsername())
}

func TestPostgreSQL_GetPassword(t *testing.T) {
	pg := DefaultPostgreSQL()
	assert.Equal(t, "postgres_password", pg.GetPassword())
}

func TestPostgreSQL_SetCredentials(t *testing.T) {
	pg := DefaultPostgreSQL()
	pg.SetCredentials("pgadmin", "pgpass123")
	assert.Equal(t, "pgadmin", pg.Username)
	assert.Equal(t, "pgpass123", pg.Password)
}

func TestPostgreSQL_GetConfig(t *testing.T) {
	pg := DefaultPostgreSQL()
	assert.Equal(t, "sslmode=disable", pg.GetConfig())
}

func TestPostgreSQL_GetModuleName(t *testing.T) {
	pg := DefaultPostgreSQL()
	assert.Equal(t, "postgresql", pg.GetModuleName())
}

func TestPostgreSQL_Clone(t *testing.T) {
	original := DefaultPostgreSQL()
	original.Host = "pg-host"
	original.Port = "5433"
	original.Dbname = "pg_db"

	cloned := original.Clone().(*PostgreSQL)

	assert.Equal(t, original.Host, cloned.Host)
	assert.Equal(t, original.Port, cloned.Port)
	assert.Equal(t, original.Dbname, cloned.Dbname)

	cloned.Host = "new-pg-host"
	assert.Equal(t, "pg-host", original.Host)
	assert.Equal(t, "new-pg-host", cloned.Host)
}

func TestPostgreSQL_Get(t *testing.T) {
	pg := DefaultPostgreSQL()
	result := pg.Get()

	assert.NotNil(t, result)
	resultPG, ok := result.(*PostgreSQL)
	assert.True(t, ok)
	assert.Equal(t, pg, resultPG)
}

func TestPostgreSQL_Set(t *testing.T) {
	pg := DefaultPostgreSQL()
	newPG := &PostgreSQL{
		ModuleName:      "new_pg",
		Host:            "new-pg-host",
		Port:            "5434",
		Config:          "sslmode=require",
		LogLevel:        "debug",
		Dbname:          "new_db",
		Username:        "new_user",
		Password:        "new_pass",
		MaxIdleConns:    15,
		MaxOpenConns:    150,
		ConnMaxIdleTime: 5000,
		ConnMaxLifeTime: 10000,
	}

	pg.Set(newPG)

	assert.Equal(t, "new_pg", pg.ModuleName)
	assert.Equal(t, "new-pg-host", pg.Host)
	assert.Equal(t, "5434", pg.Port)
	assert.Equal(t, "sslmode=require", pg.Config)
	assert.Equal(t, "debug", pg.LogLevel)
	assert.Equal(t, "new_db", pg.Dbname)
	assert.Equal(t, "new_user", pg.Username)
	assert.Equal(t, "new_pass", pg.Password)
	assert.Equal(t, 15, pg.MaxIdleConns)
	assert.Equal(t, 150, pg.MaxOpenConns)
}

func TestPostgreSQL_Validate(t *testing.T) {
	pg := DefaultPostgreSQL()
	pg.Password = "test123" // 设置密码以满足验证要求
	err := pg.Validate()
	assert.NoError(t, err)
}

func TestPostgreSQL_ChainedCalls(t *testing.T) {
	pg := DefaultPostgreSQL()
	pg.SetHost("chain-pg-host")
	pg.SetPort("5435")
	pg.SetDBName("chain_db")
	pg.SetCredentials("chain_user", "chain_pass")

	assert.Equal(t, "chain-pg-host", pg.Host)
	assert.Equal(t, "5435", pg.Port)
	assert.Equal(t, "chain_db", pg.Dbname)
	assert.Equal(t, "chain_user", pg.Username)
	assert.Equal(t, "chain_pass", pg.Password)

	err := pg.Validate()
	assert.NoError(t, err)
}
