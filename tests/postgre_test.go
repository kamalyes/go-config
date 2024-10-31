/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 10:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:03:55
 * @FilePath: \go-config\database\postgre_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"github.com/kamalyes/go-config/pkg/database"
	"testing"
)

// 公共测试数据

var testPostgreSQLData = &database.PostgreSQL{
	ModuleName:      "testModule",
	Host:            "127.0.0.1",
	Port:            "5432",
	Dbname:          "testdb",
	Username:        "testuser",
	Password:        "testpass",
	MaxIdleConns:    10,
	MaxOpenConns:    100,
	ConnMaxIdleTime: 30,
	ConnMaxLifeTime: 3600,
}

// 公共断言函数
func assertPostgreSQLEqual(t *testing.T, expected, actual interface{}, message string) {
	if expected != actual {
		t.Errorf("%s: expected %v, got %v", message, expected, actual)
	}
}

func TestPostgreSQL_Validate(t *testing.T) {
	// 测试有效配置
	if err := testPostgreSQLData.Validate(); err != nil {
		t.Errorf("Validate failed: %v", err)
	}

	// 测试无效配置
	invalidData := &database.PostgreSQL{}
	if err := invalidData.Validate(); err == nil {
		t.Error("Validate should have failed for empty configuration")
	}
}

func TestPostgreSQL_ToMap(t *testing.T) {
	dataMap := testPostgreSQLData.ToMap()
	assertPostgreSQLEqual(t, testPostgreSQLData.ModuleName, dataMap["moduleName"], "ToMap failed for ModuleName")
	assertPostgreSQLEqual(t, testPostgreSQLData.Host, dataMap["host"], "ToMap failed for Host")
}

func TestPostgreSQL_FromMap(t *testing.T) {
	newData := map[string]interface{}{
		"moduleName":         "newModule",
		"host":               "192.168.1.1",
		"port":               "5433",
		"dbname":             "newdb",
		"username":           "newuser",
		"password":           "newpass",
		"max-idle-conns":     20,
		"max-open-conns":     200,
		"conn-max-idle-time": 60,
		"conn-max-life-time": 7200,
	}
	testPostgreSQLData.FromMap(newData)

	// 验证更新后的值
	assertPostgreSQLEqual(t, "newModule", testPostgreSQLData.ModuleName, "FromMap failed for ModuleName")
	assertPostgreSQLEqual(t, "192.168.1.1", testPostgreSQLData.Host, "FromMap failed for Host")
}

func TestPostgreSQLClone(t *testing.T) {
	clonedPg := testPostgreSQLData.Clone().(*database.PostgreSQL)
	assertPostgreSQLEqual(t, testPostgreSQLData.ModuleName, clonedPg.ModuleName, "Clone failed for ModuleName")
	assertPostgreSQLEqual(t, testPostgreSQLData.Host, clonedPg.Host, "Clone failed for Host")
}

func TestPostgreSQLGet(t *testing.T) {
	data := testPostgreSQLData.Get().(*database.PostgreSQL)
	assertPostgreSQLEqual(t, testPostgreSQLData.ModuleName, data.ModuleName, "Get failed for ModuleName")
}

func TestPostgreSQL_Set(t *testing.T) {
	newConfig := &database.PostgreSQL{
		ModuleName:      "updatedModule",
		Host:            "10.0.0.1",
		Port:            "5434",
		Dbname:          "updateddb",
		Username:        "updateduser",
		Password:        "updatedpass",
		MaxIdleConns:    15,
		MaxOpenConns:    150,
		ConnMaxIdleTime: 45,
		ConnMaxLifeTime: 7200,
	}
	testPostgreSQLData.Set(newConfig)

	// 验证更新后的值
	assertPostgreSQLEqual(t, "updatedModule", testPostgreSQLData.ModuleName, "Set failed for ModuleName")
	assertPostgreSQLEqual(t, "10.0.0.1", testPostgreSQLData.Host, "Set failed for Host")
}
