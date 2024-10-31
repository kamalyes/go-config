/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:56:21
 * @FilePath: \go-config\tests\jwt_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/jwt"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 JWT 配置参数
func generateJwtTestParams() *jwt.JWT {
	return &jwt.JWT{
		ModuleName:    random.RandString(10, random.CAPITAL),
		SigningKey:    random.RandString(16, random.CAPITAL), // 随机生成签名密钥
		ExpiresTime:   int64(random.FRandInt(3600, 86400)),   // 随机生成过期时间（1小时到1天）
		BufferTime:    int64(random.FRandInt(60, 600)),       // 随机生成缓冲时间（1分钟到10分钟）
		UseMultipoint: random.FRandBool(),                    // 随机生成多地登录拦截
	}
}

// 将 JWT 的参数转换为 map
func jwtToMap(jwt *jwt.JWT) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME":    jwt.ModuleName,
		"SIGNING_KEY":    jwt.SigningKey,
		"EXPIRES_TIME":   jwt.ExpiresTime,
		"BUFFER_TIME":    jwt.BufferTime,
		"USE_MULTIPOINT": jwt.UseMultipoint,
	}
}

// 验证 JWT 的字段与期望的映射是否相等
func assertJwtFields(t *testing.T, jwt *jwt.JWT, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], jwt.ModuleName)
	assert.Equal(t, expected["SIGNING_KEY"], jwt.SigningKey)
	assert.Equal(t, expected["EXPIRES_TIME"], jwt.ExpiresTime)
	assert.Equal(t, expected["BUFFER_TIME"], jwt.BufferTime)
	assert.Equal(t, expected["USE_MULTIPOINT"], jwt.UseMultipoint)
}

func TestJwtClone(t *testing.T) {
	params := generateJwtTestParams()
	original := jwt.NewJWT(params)
	cloned := original.Clone().(*jwt.JWT)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestJwtSet(t *testing.T) {
	oldParams := generateJwtTestParams()
	newParams := generateJwtTestParams()

	jwtInstance := jwt.NewJWT(oldParams)
	newConfig := jwt.NewJWT(newParams)

	jwtInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, jwtInstance.ModuleName)
	assert.Equal(t, newParams.SigningKey, jwtInstance.SigningKey)
	assert.Equal(t, newParams.ExpiresTime, jwtInstance.ExpiresTime)
	assert.Equal(t, newParams.BufferTime, jwtInstance.BufferTime)
	assert.Equal(t, newParams.UseMultipoint, jwtInstance.UseMultipoint)
}
