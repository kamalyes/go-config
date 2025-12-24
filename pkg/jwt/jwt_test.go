/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-24 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-24 00:00:00
 * @FilePath: \go-config\pkg\jwt\jwt_test.go
 * @Description: JWT配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWT_Clone(t *testing.T) {
	jwt := &JWT{
		ModuleName:       "jwt",
		SigningKey:       "test-key",
		ExpiresTime:      3600,
		BufferTime:       300,
		UseMultipoint:    true,
		Issuer:           "test-issuer",
		Audience:         "test-audience",
		Algorithm:        "HS256",
		EnableRefresh:    true,
		RefreshTokenLife: 7200,
		Subject:          "test-subject",
		CustomClaims: map[string]interface{}{
			"role": "admin",
			"org":  "test-org",
		},
	}

	cloned := jwt.Clone()

	assert.NotNil(t, cloned)
	clonedJWT, ok := cloned.(*JWT)
	assert.True(t, ok)
	assert.Equal(t, jwt.SigningKey, clonedJWT.SigningKey)
	assert.Equal(t, jwt.ExpiresTime, clonedJWT.ExpiresTime)
	assert.Equal(t, jwt.UseMultipoint, clonedJWT.UseMultipoint)
	assert.Equal(t, jwt.Algorithm, clonedJWT.Algorithm)

	// 验证是独立副本
	clonedJWT.SigningKey = "modified-key"
	assert.NotEqual(t, jwt.SigningKey, clonedJWT.SigningKey)

	// 验证map的深拷贝
	clonedJWT.CustomClaims["role"] = "user"
	assert.NotEqual(t, jwt.CustomClaims["role"], clonedJWT.CustomClaims["role"])
}

func TestJWT_Get(t *testing.T) {
	jwt := &JWT{
		SigningKey:  "test-key",
		ExpiresTime: 3600,
	}

	result := jwt.Get()
	assert.Equal(t, jwt, result)
}

func TestJWT_Set(t *testing.T) {
	jwt := &JWT{}
	newJWT := &JWT{
		SigningKey:  "new-key",
		ExpiresTime: 7200,
	}

	jwt.Set(newJWT)
	assert.Equal(t, newJWT.SigningKey, jwt.SigningKey)
	assert.Equal(t, newJWT.ExpiresTime, jwt.ExpiresTime)
}
