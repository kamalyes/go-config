/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 10:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 17:26:26
 * @FilePath: \go-config\tests\postgre_test.go
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

// 生成随机的 PostgreSQL 配置参数
func generatePostgreSQLTestParams() *database.PostgreSQL {
	return &database.PostgreSQL{
		ModuleName:      random.RandString(10, random.CAPITAL),
		Host:            random.RandString(5, random.CAPITAL) + ".postgresql.local", // 随机生成主机名
		Port:            random.RandString(5, random.NUMBER),                        // 随机生成端口
		Config:          "sslmode=disable",                                          // 默认配置
		LogLevel:        random.RandString(5, random.CAPITAL),                       // 随机生成日志等级
		Dbname:          random.RandString(8, random.CAPITAL),                       // 随机生成数据库名称
		Username:        random.RandString(8, random.CAPITAL),                       // 随机生成用户名
		Password:        random.RandString(16, random.CAPITAL),                      // 随机生成密码
		MaxIdleConns:    random.FRandInt(1, 100),                                    // 随机生成最大空闲连接数
		MaxOpenConns:    random.FRandInt(1, 100),                                    // 随机生成最大连接数
		ConnMaxIdleTime: random.FRandInt(30, 300),                                   // 随机生成最大空闲时间（30到300秒）
		ConnMaxLifeTime: random.FRandInt(30, 300),                                   // 随机生成最大生命周期（30到300秒）
	}
}

// 将 PostgreSQL 的参数转换为 map
func postgresqlToMap(pg *database.PostgreSQL) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME":        pg.ModuleName,
		"HOST":               pg.Host,
		"PORT":               pg.Port,
		"CONFIG":             pg.Config,
		"LOG_LEVEL":          pg.LogLevel,
		"DB_NAME":            pg.Dbname,
		"USERNAME":           pg.Username,
		"PASSWORD":           pg.Password,
		"MAX_IDLE_CONNS":     pg.MaxIdleConns,
		"MAX_OPEN_CONNS":     pg.MaxOpenConns,
		"CONN_MAX_IDLE_TIME": pg.ConnMaxIdleTime,
		"CONN_MAX_LIFE_TIME": pg.ConnMaxLifeTime,
	}
}

// 验证 PostgreSQL 的字段与期望的映射是否相等
func assertPostgreSQLFields(t *testing.T, pg *database.PostgreSQL, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], pg.ModuleName)
	assert.Equal(t, expected["HOST"], pg.Host)
	assert.Equal(t, expected["PORT"], pg.Port)
	assert.Equal(t, expected["CONFIG"], pg.Config)
	assert.Equal(t, expected["LOG_LEVEL"], pg.LogLevel)
	assert.Equal(t, expected["DB_NAME"], pg.Dbname)
	assert.Equal(t, expected["USERNAME"], pg.Username)
	assert.Equal(t, expected["PASSWORD"], pg.Password)
	assert.Equal(t, expected["MAX_IDLE_CONNS"], pg.MaxIdleConns)
	assert.Equal(t, expected["MAX_OPEN_CONNS"], pg.MaxOpenConns)
	assert.Equal(t, expected["CONN_MAX_IDLE_TIME"], pg.ConnMaxIdleTime)
	assert.Equal(t, expected["CONN_MAX_LIFE_TIME"], pg.ConnMaxLifeTime)
}

func TestPostgreSQLClone(t *testing.T) {
	params := generatePostgreSQLTestParams()
	original := database.NewPostgreSQL(params)
	cloned := original.Clone().(*database.PostgreSQL)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不相同的实例
}

func TestPostgreSQLSet(t *testing.T) {
	oldParams := generatePostgreSQLTestParams()
	newParams := generatePostgreSQLTestParams()

	pgInstance := database.NewPostgreSQL(oldParams)
	newConfig := database.NewPostgreSQL(newParams)

	pgInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, pgInstance.ModuleName)
	assert.Equal(t, newParams.Host, pgInstance.Host)
	assert.Equal(t, newParams.Port, pgInstance.Port)
	assert.Equal(t, newParams.Config, pgInstance.Config)
	assert.Equal(t, newParams.LogLevel, pgInstance.LogLevel)
	assert.Equal(t, newParams.Dbname, pgInstance.Dbname)
	assert.Equal(t, newParams.Username, pgInstance.Username)
	assert.Equal(t, newParams.Password, pgInstance.Password)
	assert.Equal(t, newParams.MaxIdleConns, pgInstance.MaxIdleConns)
	assert.Equal(t, newParams.MaxOpenConns, pgInstance.MaxOpenConns)
	assert.Equal(t, newParams.ConnMaxIdleTime, pgInstance.ConnMaxIdleTime)
	assert.Equal(t, newParams.ConnMaxLifeTime, pgInstance.ConnMaxLifeTime)
}
