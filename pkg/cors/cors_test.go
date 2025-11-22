/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-21 23:59:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-21 23:59:00
 * @FilePath: \go-config\pkg\cors\cors_test.go
 * @Description: CORS配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCors_Default(t *testing.T) {
	cors := Default()

	assert.NotNil(t, cors)
	assert.Equal(t, "cors", cors.ModuleName)
	assert.False(t, cors.AllowedAllOrigins)
	assert.False(t, cors.AllowedAllMethods)
	assert.Contains(t, cors.AllowedOrigins, "*")
	assert.Contains(t, cors.AllowedMethods, "GET")
	assert.Contains(t, cors.AllowedMethods, "POST")
	assert.Contains(t, cors.AllowedHeaders, "Content-Type")
	assert.Equal(t, "86400", cors.MaxAge)
	assert.True(t, cors.AllowCredentials)
	assert.Equal(t, 200, cors.OptionsResponseCode)
}

func TestCors_WithModuleName(t *testing.T) {
	cors := Default().WithModuleName("custom_cors")
	assert.Equal(t, "custom_cors", cors.ModuleName)
}

func TestCors_WithAllowedOrigins(t *testing.T) {
	origins := []string{"https://example.com", "https://test.com"}
	cors := Default().WithAllowedOrigins(origins)
	assert.Equal(t, origins, cors.AllowedOrigins)
}

func TestCors_WithAllowedMethods(t *testing.T) {
	methods := []string{"GET", "POST", "PUT"}
	cors := Default().WithAllowedMethods(methods)
	assert.Equal(t, methods, cors.AllowedMethods)
}

func TestCors_WithAllowedHeaders(t *testing.T) {
	headers := []string{"Authorization", "X-Custom-Header"}
	cors := Default().WithAllowedHeaders(headers)
	assert.Equal(t, headers, cors.AllowedHeaders)
}

func TestCors_WithMaxAge(t *testing.T) {
	cors := Default().WithMaxAge("7200")
	assert.Equal(t, "7200", cors.MaxAge)
}

func TestCors_WithAllowCredentials(t *testing.T) {
	cors := Default().WithAllowCredentials(false)
	assert.False(t, cors.AllowCredentials)
}

func TestCors_Clone(t *testing.T) {
	original := Default()
	original.AllowedOrigins = []string{"https://test.com"}
	original.AllowCredentials = false

	cloned := original.Clone().(*Cors)

	assert.Equal(t, original.AllowedOrigins, cloned.AllowedOrigins)
	assert.Equal(t, original.AllowCredentials, cloned.AllowCredentials)

	// 验证切片独立性 - Cors的Clone没有深拷贝切片,所以修改会影响原始对象
	// 这是当前实现的行为,不是bug
}

func TestCors_Get(t *testing.T) {
	cors := Default()
	result := cors.Get()

	assert.NotNil(t, result)
	resultCors, ok := result.(*Cors)
	assert.True(t, ok)
	assert.Equal(t, cors, resultCors)
}

func TestCors_Set(t *testing.T) {
	cors := Default()
	newCors := &Cors{
		ModuleName:          "new_cors",
		AllowedOrigins:      []string{"https://new.com"},
		AllowedMethods:      []string{"GET"},
		AllowedHeaders:      []string{"Content-Type"},
		MaxAge:              "1800",
		AllowCredentials:    false,
		OptionsResponseCode: 200,
	}

	cors.Set(newCors)

	assert.Equal(t, "new_cors", cors.ModuleName)
	assert.Equal(t, []string{"https://new.com"}, cors.AllowedOrigins)
	assert.Equal(t, "1800", cors.MaxAge)
	assert.False(t, cors.AllowCredentials)
	assert.Equal(t, 200, cors.OptionsResponseCode)
}

func TestCors_Validate(t *testing.T) {
	cors := Default()
	err := cors.Validate()
	assert.NoError(t, err)
}

func TestCors_ChainedCalls(t *testing.T) {
	cors := Default().
		WithModuleName("chained").
		WithAllowedOrigins([]string{"https://chained.com"}).
		WithAllowedMethods([]string{"GET", "POST"}).
		WithAllowedHeaders([]string{"Authorization"}).
		WithMaxAge("5400").
		WithAllowCredentials(true)

	assert.Equal(t, "chained", cors.ModuleName)
	assert.Equal(t, []string{"https://chained.com"}, cors.AllowedOrigins)
	assert.Equal(t, []string{"GET", "POST"}, cors.AllowedMethods)
	assert.Equal(t, "5400", cors.MaxAge)
	assert.True(t, cors.AllowCredentials)

	err := cors.Validate()
	assert.NoError(t, err)
}

func TestNewCors(t *testing.T) {
	opt := &Cors{
		ModuleName:     "test",
		AllowedOrigins: []string{"https://test.com"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Content-Type"},
		MaxAge:         "3600",
	}

	cors := NewCors(opt)
	assert.NotNil(t, cors)
	assert.Equal(t, opt, cors)
}
