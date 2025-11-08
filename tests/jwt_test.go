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

// TestJWTDefault 测试默认配置
func TestJWTDefault(t *testing.T) {
	defaultConfig := jwt.DefaultJWT()
	
	// 检查默认值
	assert.Equal(t, "jwt", defaultConfig.ModuleName)
	assert.Equal(t, "go-config-default-key", defaultConfig.SigningKey)
	assert.Equal(t, int64(3600*24*7), defaultConfig.ExpiresTime) // 7天
	assert.Equal(t, int64(3600), defaultConfig.BufferTime)        // 1小时
	assert.Equal(t, false, defaultConfig.UseMultipoint)
}

// TestJWTDefaultPointer 测试默认配置指针
func TestJWTDefaultPointer(t *testing.T) {
	config := jwt.Default()
	
	assert.NotNil(t, config)
	assert.Equal(t, "jwt", config.ModuleName)
	assert.Equal(t, "go-config-default-key", config.SigningKey)
}

// TestJWTChainMethods 测试链式方法
func TestJWTChainMethods(t *testing.T) {
	config := jwt.Default().
		WithModuleName("auth-service").
		WithSigningKey("my-secret-key").
		WithExpiresTime(7200).
		WithBufferTime(900).
		WithUseMultipoint(true)
	
	assert.Equal(t, "auth-service", config.ModuleName)
	assert.Equal(t, "my-secret-key", config.SigningKey)
	assert.Equal(t, int64(7200), config.ExpiresTime)
	assert.Equal(t, int64(900), config.BufferTime)
	assert.Equal(t, true, config.UseMultipoint)
}

// TestJWTChainMethodsReturnPointer 测试链式方法返回指针
func TestJWTChainMethodsReturnPointer(t *testing.T) {
	config1 := jwt.Default()
	config2 := config1.WithSigningKey("new-key")
	
	// 应该返回同一个实例
	assert.Same(t, config1, config2)
	assert.Equal(t, "new-key", config1.SigningKey)
}
