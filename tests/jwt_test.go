/*
 * @Author: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:29:08
 * @LastEditors: kamalyes 501893067@qq.com
 * @FilePath: \go-config\jwt\jwt_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/jwt"
)

// 辅助函数用于验证 JWT 结构体的字段
func validateJWTFields(t *testing.T, jwtInstance *jwt.JWT, expectedModuleName, expectedSigningKey string, expectedExpiresTime, expectedBufferTime int64, expectedUseMultipoint bool) {
	if jwtInstance.ModuleName != expectedModuleName {
		t.Errorf("expected ModuleName to be '%s', got '%s'", expectedModuleName, jwtInstance.ModuleName)
	}
	if jwtInstance.SigningKey != expectedSigningKey {
		t.Errorf("expected SigningKey to be '%s', got '%s'", expectedSigningKey, jwtInstance.SigningKey)
	}
	if jwtInstance.ExpiresTime != expectedExpiresTime {
		t.Errorf("expected ExpiresTime to be %d, got %d", expectedExpiresTime, jwtInstance.ExpiresTime)
	}
	if jwtInstance.BufferTime != expectedBufferTime {
		t.Errorf("expected BufferTime to be %d, got %d", expectedBufferTime, jwtInstance.BufferTime)
	}
	if jwtInstance.UseMultipoint != expectedUseMultipoint {
		t.Errorf("expected UseMultipoint to be %v, got %v", expectedUseMultipoint, jwtInstance.UseMultipoint)
	}
}

func TestNewJWTInstance(t *testing.T) { // 重命名测试函数
	jwtInstance := jwt.NewJWT("testModule", "testKey", 3600, 300, true)
	validateJWTFields(t, jwtInstance, "testModule", "testKey", 3600, 300, true)
}

func TestJWTConvertToMap(t *testing.T) { // 重命名测试函数
	jwtInstance := jwt.NewJWT("testModule", "testKey", 3600, 300, true)
	expectedMap := map[string]interface{}{
		"moduleName":    "testModule",
		"signingKey":    "testKey",
		"expiresTime":   int64(3600),
		"bufferTime":    int64(300),
		"useMultipoint": true,
	}

	result := jwtInstance.ToMap()
	// 验证结果映射是否与预期一致
	for key, expectedValue := range expectedMap {
		if result[key] != expectedValue {
			t.Errorf("expected %s to be %v, got %v", key, expectedValue, result[key])
		}
	}
}

func TestJWTPopulateFromMap(t *testing.T) { // 重命名测试函数
	jwtInstance := jwt.NewJWT("", "", 0, 0, false)
	data := map[string]interface{}{
		"moduleName":    "testModule",
		"signingKey":    "testKey",
		"expiresTime":   int64(3600),
		"bufferTime":    int64(300),
		"useMultipoint": true,
	}

	jwtInstance.FromMap(data)
	validateJWTFields(t, jwtInstance, "testModule", "testKey", 3600, 300, true)
}

func TestJWTCloneInstance(t *testing.T) { // 重命名测试函数
	jwtInstance := jwt.NewJWT("testModule", "testKey", 3600, 300, true)
	clone := jwtInstance.Clone()

	if clone == jwtInstance {
		t.Error("Clone should return a new instance, not the same instance")
	}
	// 尝试将 clone 转换为 *jwt.JWT 类型
	clonedJWT, ok := clone.(*jwt.JWT)
	if !ok {
		t.Fatalf("Clone should return a *jwt.JWT, got %T", clone)
	}
	validateJWTFields(t, clonedJWT, "testModule", "testKey", 3600, 300, true)
}

func TestJWTRetrieveInstance(t *testing.T) { // 重命名测试函数
	jwtInstance := jwt.NewJWT("testModule", "testKey", 3600, 300, true)
	result := jwtInstance.Get().(*jwt.JWT)

	if result != jwtInstance {
		t.Error("Get should return the same instance")
	}
}

func TestJWTUpdateInstance(t *testing.T) { // 重命名测试函数
	jwtInstance := jwt.NewJWT("testModule", "testKey", 3600, 300, true)
	newData := jwt.NewJWT("newModule", "newKey", 7200, 600, false)

	jwtInstance.Set(newData)

	validateJWTFields(t, jwtInstance, "newModule", "newKey", 7200, 600, false)
}

func TestJWTValidation(t *testing.T) { // 重命名测试函数
	jwtInstance := jwt.NewJWT("", "", 0, 0, false)
	err := jwtInstance.Validate()

	if err == nil {
		t.Error("expected validation to fail due to empty fields")
	}

	jwtInstance = jwt.NewJWT("testModule", "testKey", 3600, 300, true)
	err = jwtInstance.Validate()

	if err != nil {
		t.Errorf("expected validation to succeed, got error: %v", err)
	}
}
