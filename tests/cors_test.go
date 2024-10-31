/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 13:15:26
 * @FilePath: \go-config\tests\cors_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/cors"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Cors 配置参数
func generateCorsTestParams() *cors.Cors {
	return &cors.Cors{
		ModuleName:          random.RandString(10, random.CAPITAL),
		AllowedAllOrigins:   random.FRandBool(),
		AllowedAllMethods:   random.FRandBool(),
		AllowedOrigins:      random.RandStringSlice(3, 15, random.CAPITAL), // 随机生成 3 个字符串
		AllowedMethods:      random.RandStringSlice(3, 5, random.CAPITAL),  // 随机生成 3 个字符串
		AllowedHeaders:      random.RandStringSlice(3, 30, random.CAPITAL), // 随机生成 3 个字符串
		MaxAge:              fmt.Sprintf("%d", random.FRandInt(1, 3600)),   // 随机生成 1 到 3600 的整数
		ExposedHeaders:      random.RandStringSlice(3, 30, random.CAPITAL), // 随机生成 3 个字符串
		AllowCredentials:    random.FRandBool(),
		OptionsResponseCode: random.RandInt(200, 600), // 随机生成 200 到 600 的整数
	}
}

func TestCorsClone(t *testing.T) {
	params := generateCorsTestParams()
	original := cors.NewCors(params)
	cloned := original.Clone().(*cors.Cors)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestCorsSet(t *testing.T) {
	oldParams := generateCorsTestParams()
	newParams := generateCorsTestParams()

	corsInstance := cors.NewCors(oldParams)
	newConfig := cors.NewCors(newParams)

	corsInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, corsInstance.ModuleName)
	assert.Equal(t, newParams.AllowedAllOrigins, corsInstance.AllowedAllOrigins)
	assert.Equal(t, newParams.AllowedAllMethods, corsInstance.AllowedAllMethods)
	assert.Equal(t, newParams.AllowedOrigins, corsInstance.AllowedOrigins)
	assert.Equal(t, newParams.AllowedMethods, corsInstance.AllowedMethods)
	assert.Equal(t, newParams.AllowedHeaders, corsInstance.AllowedHeaders)
	assert.Equal(t, newParams.MaxAge, corsInstance.MaxAge)
	assert.Equal(t, newParams.ExposedHeaders, corsInstance.ExposedHeaders)
	assert.Equal(t, newParams.AllowCredentials, corsInstance.AllowCredentials)
	assert.Equal(t, newParams.OptionsResponseCode, corsInstance.OptionsResponseCode)
}
