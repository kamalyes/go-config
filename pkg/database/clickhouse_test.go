/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-04-23 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-04-23 00:00:00
 * @FilePath: \go-config\pkg\database\clickhouse_test.go
 * @Description: ClickHouse数据库配置测试
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClickHouse_DefaultClickHouse(t *testing.T) {
	ch := DefaultClickHouse()

	assert.NotNil(t, ch)
	assert.Equal(t, "clickhouse", ch.ModuleName)
	assert.Equal(t, "localhost", ch.Host)
	assert.Equal(t, "9000", ch.Port)
	assert.Equal(t, "info", ch.LogLevel)
	assert.Equal(t, "default", ch.Dbname)
	assert.Equal(t, "default", ch.Username)
	assert.Equal(t, "", ch.Password)
	assert.Equal(t, "native", ch.Protocol)
	assert.Equal(t, true, ch.Compress)
	assert.Equal(t, false, ch.Secure)
	assert.Equal(t, 10, ch.DialTimeout)
	assert.Equal(t, 20, ch.ReadTimeout)
	assert.Equal(t, 5, ch.MaxIdleConns)
	assert.Equal(t, 20, ch.MaxOpenConns)
	// GORM 字段
	assert.Equal(t, 200, ch.SlowThreshold)
	assert.Equal(t, true, ch.SkipDefaultTransaction)
	assert.Equal(t, false, ch.PrepareStmt)
	assert.Equal(t, true, ch.DisableForeignKeyConstraintWhenMigrating)
	assert.Equal(t, 1000, ch.CreateBatchSize)
	assert.Equal(t, true, ch.SingularTable)
}

func TestClickHouse_GetDBType(t *testing.T) {
	ch := DefaultClickHouse()
	assert.Equal(t, DBTypeClickHouse, ch.GetDBType())
}

func TestClickHouse_SettersAndGetters(t *testing.T) {
	ch := DefaultClickHouse()
	ch.SetHost("10.0.0.1")
	ch.SetPort("8123")
	ch.SetDBName("analytics")
	ch.SetCredentials("admin", "secret")

	assert.Equal(t, "10.0.0.1", ch.GetHost())
	assert.Equal(t, "8123", ch.GetPort())
	assert.Equal(t, "analytics", ch.GetDBName())
	assert.Equal(t, "admin", ch.GetUsername())
	assert.Equal(t, "secret", ch.GetPassword())
}

func TestClickHouse_WithMethods(t *testing.T) {
	ch := DefaultClickHouse().
		WithModuleName("ch").
		WithHost("ch.example.com").
		WithPort("8123").
		WithConfig("debug=true").
		WithLogLevel("debug").
		WithDbname("metrics").
		WithUsername("reader").
		WithPassword("p@ss").
		WithProtocol("http").
		WithCluster("mycluster").
		WithCompress(false).
		WithDebug(true).
		WithSecure(true).
		WithDialTimeout(5).
		WithReadTimeout(15).
		WithMaxIdleConns(8).
		WithMaxOpenConns(32).
		WithConnMaxIdleTime(600).
		WithConnMaxLifeTime(7200)

	assert.Equal(t, "ch", ch.ModuleName)
	assert.Equal(t, "ch.example.com", ch.Host)
	assert.Equal(t, "8123", ch.Port)
	assert.Equal(t, "http", ch.Protocol)
	assert.Equal(t, "mycluster", ch.Cluster)
	assert.Equal(t, true, ch.Debug)
	assert.Equal(t, true, ch.Secure)
	assert.Equal(t, 32, ch.MaxOpenConns)
}

func TestClickHouse_Clone(t *testing.T) {
	ch := DefaultClickHouse()
	ch.Host = "source-host"

	cloned := ch.Clone().(*ClickHouse)
	assert.Equal(t, "source-host", cloned.Host)

	cloned.Host = "changed"
	assert.Equal(t, "source-host", ch.Host, "原对象不应被修改")
}

func TestClickHouse_Validate(t *testing.T) {
	ch := DefaultClickHouse()
	assert.NoError(t, ch.Validate())

	// Host 为空应验证失败
	ch.Host = ""
	assert.Error(t, ch.Validate())
}

func TestClickHouse_GetSet(t *testing.T) {
	ch := DefaultClickHouse()
	got := ch.Get().(*ClickHouse)
	assert.Equal(t, ch, got)

	newCfg := DefaultClickHouse()
	newCfg.Host = "new-host"
	ch.Set(newCfg)
	assert.Equal(t, "new-host", ch.Host)
}
