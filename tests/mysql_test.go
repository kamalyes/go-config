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
