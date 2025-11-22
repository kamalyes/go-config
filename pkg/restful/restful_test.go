/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\restful\restful_test.go
 * @Description: RESTful API配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package restful

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRestful_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "restful", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "restful-api", config.Name)
	assert.Equal(t, "127.0.0.1", config.Host)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, "dev", config.Mode)
	assert.Equal(t, 1000, config.MaxConns)
	assert.Equal(t, int64(32*1024*1024), config.MaxBytes)
	assert.Equal(t, 30, config.Timeout)
	assert.Equal(t, int64(900), config.CpuThreshold)
	assert.False(t, config.Auth)
	assert.False(t, config.PrintRoutes)
	assert.False(t, config.StrictSlash)
	assert.NotNil(t, config.Headers)
	assert.NotNil(t, config.Middlewares)
	assert.NotNil(t, config.Signature)
	assert.NotNil(t, config.CORS)
	assert.NotNil(t, config.TLS)
	assert.NotNil(t, config.RateLimit)
	assert.NotNil(t, config.Compression)
	assert.NotNil(t, config.Static)
}

func TestRestful_WithModuleName(t *testing.T) {
	config := Default()
	result := config.WithModuleName("custom-api")
	assert.Equal(t, "custom-api", result.ModuleName)
	assert.Equal(t, config, result)
}

func TestRestful_WithName(t *testing.T) {
	config := Default()
	result := config.WithName("my-api")
	assert.Equal(t, "my-api", result.Name)
	assert.Equal(t, config, result)
}

func TestRestful_WithHost(t *testing.T) {
	config := Default()
	result := config.WithHost("0.0.0.0")
	assert.Equal(t, "0.0.0.0", result.Host)
	assert.Equal(t, config, result)
}

func TestRestful_WithPort(t *testing.T) {
	config := Default()
	result := config.WithPort(9090)
	assert.Equal(t, 9090, result.Port)
	assert.Equal(t, config, result)
}

func TestRestful_WithMode(t *testing.T) {
	config := Default()
	result := config.WithMode("prod")
	assert.Equal(t, "prod", result.Mode)
	assert.Equal(t, config, result)
}

func TestRestful_WithTimeout(t *testing.T) {
	config := Default()
	result := config.WithTimeout(60)
	assert.Equal(t, 60, result.Timeout)
	assert.Equal(t, config, result)
}

func TestRestful_EnableCORS(t *testing.T) {
	config := Default()
	result := config.EnableCORS()
	assert.True(t, result.CORS.Enabled)
	assert.Equal(t, config, result)
}

func TestRestful_EnableTLS(t *testing.T) {
	config := Default()
	result := config.EnableTLS("/path/to/cert.pem", "/path/to/key.pem")
	assert.True(t, result.TLS.Enabled)
	assert.Equal(t, "/path/to/cert.pem", result.TLS.CertFile)
	assert.Equal(t, "/path/to/key.pem", result.TLS.KeyFile)
	assert.Equal(t, config, result)
}

func TestRestful_EnableStatic(t *testing.T) {
	config := Default()
	result := config.EnableStatic("/public", "/static")
	assert.True(t, result.Static.Enabled)
	assert.Equal(t, "/public", result.Static.Root)
	assert.Equal(t, "/static", result.Static.Prefix)
	assert.Equal(t, config, result)
}

func TestRestful_Enable(t *testing.T) {
	config := Default()
	result := config.Enable()
	assert.True(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestRestful_Disable(t *testing.T) {
	config := Default()
	config.Enabled = true
	result := config.Disable()
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestRestful_Clone(t *testing.T) {
	config := Default()
	config.WithName("test-api").WithPort(9090)
	config.Headers["X-Custom"] = "value"
	config.Middlewares = append(config.Middlewares, "logger")
	config.Signature.PrivateKeys = []string{"key1", "key2"}
	config.CORS.AllowOrigins = []string{"http://example.com"}

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*Restful)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.Name, clonedConfig.Name)
	assert.Equal(t, config.Port, clonedConfig.Port)

	// 验证深拷贝 - 切片
	assert.Equal(t, len(config.Middlewares), len(clonedConfig.Middlewares))
	clonedConfig.Middlewares = append(clonedConfig.Middlewares, "extra")
	assert.NotEqual(t, len(config.Middlewares), len(clonedConfig.Middlewares))

	// 验证深拷贝 - map
	clonedConfig.Headers["X-New"] = "new-value"
	_, exists := config.Headers["X-New"]
	assert.False(t, exists)

	// 验证深拷贝 - 嵌套切片
	clonedConfig.Signature.PrivateKeys = append(clonedConfig.Signature.PrivateKeys, "key3")
	assert.NotEqual(t, len(config.Signature.PrivateKeys), len(clonedConfig.Signature.PrivateKeys))
}

func TestRestful_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestRestful_Set(t *testing.T) {
	config := Default()
	newConfig := &Restful{
		ModuleName: "new-api",
		Enabled:    true,
		Port:       9090,
	}

	config.Set(newConfig)
	assert.Equal(t, "new-api", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, 9090, config.Port)
}

func TestRestful_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
	assert.Equal(t, "restful", config.ModuleName)
	assert.Equal(t, "127.0.0.1", config.Host)
	assert.Equal(t, 8080, config.Port)
}

func TestRestful_NewRestful(t *testing.T) {
	opt := &Restful{
		ModuleName: "test-api",
		Port:       9090,
	}

	result := NewRestful(opt)
	assert.NotNil(t, result)
	assert.Equal(t, "test-api", result.ModuleName)
	assert.Equal(t, 9090, result.Port)
}

func TestRestful_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("my-restful").
		WithName("my-service").
		WithHost("0.0.0.0").
		WithPort(8888).
		WithMode("prod").
		WithTimeout(120).
		EnableCORS().
		EnableTLS("/cert.pem", "/key.pem").
		EnableStatic("/public", "/static").
		Enable()

	assert.Equal(t, "my-restful", config.ModuleName)
	assert.Equal(t, "my-service", config.Name)
	assert.Equal(t, "0.0.0.0", config.Host)
	assert.Equal(t, 8888, config.Port)
	assert.Equal(t, "prod", config.Mode)
	assert.Equal(t, 120, config.Timeout)
	assert.True(t, config.CORS.Enabled)
	assert.True(t, config.TLS.Enabled)
	assert.True(t, config.Static.Enabled)
	assert.True(t, config.Enabled)
}
