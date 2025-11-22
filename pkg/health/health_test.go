/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\health\health_test.go
 * @Description: 健康检查配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package health

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealth_Default(t *testing.T) {
	health := Default()

	assert.NotNil(t, health)
	assert.Equal(t, "health", health.ModuleName)
	assert.True(t, health.Enabled)
	assert.Equal(t, "/health", health.Path)
	assert.Equal(t, 8080, health.Port)
	assert.Equal(t, 30, health.Timeout)

	assert.NotNil(t, health.Redis)
	assert.False(t, health.Redis.Enabled)
	assert.Equal(t, "/health/redis", health.Redis.Path)
	assert.Equal(t, 5, health.Redis.Timeout)

	assert.NotNil(t, health.MySQL)
	assert.False(t, health.MySQL.Enabled)
	assert.Equal(t, "/health/mysql", health.MySQL.Path)
	assert.Equal(t, 5, health.MySQL.Timeout)
}

func TestHealth_WithModuleName(t *testing.T) {
	health := Default().WithModuleName("custom_health")
	assert.Equal(t, "custom_health", health.ModuleName)
}

func TestHealth_WithEnabled(t *testing.T) {
	health := Default().WithEnabled(false)
	assert.False(t, health.Enabled)
}

func TestHealth_WithPath(t *testing.T) {
	health := Default().WithPath("/custom/health")
	assert.Equal(t, "/custom/health", health.Path)
}

func TestHealth_WithPort(t *testing.T) {
	health := Default().WithPort(9090)
	assert.Equal(t, 9090, health.Port)
}

func TestHealth_WithTimeout(t *testing.T) {
	health := Default().WithTimeout(60)
	assert.Equal(t, 60, health.Timeout)
}

func TestHealth_WithRedisCheck(t *testing.T) {
	health := Default().WithRedisCheck(true, "/custom/redis", 10)
	assert.True(t, health.Redis.Enabled)
	assert.Equal(t, "/custom/redis", health.Redis.Path)
	assert.Equal(t, 10, health.Redis.Timeout)
}

func TestHealth_WithMySQLCheck(t *testing.T) {
	health := Default().WithMySQLCheck(true, "/custom/mysql", 15)
	assert.True(t, health.MySQL.Enabled)
	assert.Equal(t, "/custom/mysql", health.MySQL.Path)
	assert.Equal(t, 15, health.MySQL.Timeout)
}

func TestHealth_Enable(t *testing.T) {
	health := Default()
	health.Enabled = false
	health.Enable()
	assert.True(t, health.Enabled)
}

func TestHealth_Disable(t *testing.T) {
	health := Default().Disable()
	assert.False(t, health.Enabled)
}

func TestHealth_IsEnabled(t *testing.T) {
	health := Default()
	assert.True(t, health.IsEnabled())

	health.Enabled = false
	assert.False(t, health.IsEnabled())
}

func TestHealth_Clone(t *testing.T) {
	original := Default()
	original.Path = "/test/health"
	original.Port = 9999
	original.Redis.Enabled = true
	original.MySQL.Enabled = true

	cloned := original.Clone().(*Health)

	assert.Equal(t, original.Path, cloned.Path)
	assert.Equal(t, original.Port, cloned.Port)
	assert.Equal(t, original.Redis.Enabled, cloned.Redis.Enabled)
	assert.Equal(t, original.MySQL.Enabled, cloned.MySQL.Enabled)

	// 验证独立性
	cloned.Port = 7777
	cloned.Redis.Enabled = false
	assert.Equal(t, 9999, original.Port)
	assert.True(t, original.Redis.Enabled)
}

func TestHealth_Get(t *testing.T) {
	health := Default()
	result := health.Get()

	assert.NotNil(t, result)
	resultHealth, ok := result.(*Health)
	assert.True(t, ok)
	assert.Equal(t, health, resultHealth)
}

func TestHealth_Set(t *testing.T) {
	health := Default()
	newHealth := &Health{
		ModuleName: "new_health",
		Enabled:    false,
		Path:       "/new/health",
		Port:       3000,
		Timeout:    120,
		Redis:      &RedisConfig{Enabled: true, Path: "/new/redis", Timeout: 10},
		MySQL:      &MySQLConfig{Enabled: true, Path: "/new/mysql", Timeout: 10},
	}

	health.Set(newHealth)

	assert.Equal(t, "new_health", health.ModuleName)
	assert.False(t, health.Enabled)
	assert.Equal(t, "/new/health", health.Path)
	assert.Equal(t, 3000, health.Port)
	assert.Equal(t, 120, health.Timeout)
	assert.True(t, health.Redis.Enabled)
	assert.True(t, health.MySQL.Enabled)
}

func TestHealth_Validate(t *testing.T) {
	health := Default()
	err := health.Validate()
	assert.NoError(t, err)
}

func TestHealth_ChainedCalls(t *testing.T) {
	health := Default().
		WithModuleName("chained").
		WithEnabled(true).
		WithPath("/api/health").
		WithPort(8888).
		WithTimeout(45).
		WithRedisCheck(true, "/api/redis", 10).
		WithMySQLCheck(true, "/api/mysql", 15)

	assert.Equal(t, "chained", health.ModuleName)
	assert.True(t, health.Enabled)
	assert.Equal(t, "/api/health", health.Path)
	assert.Equal(t, 8888, health.Port)
	assert.Equal(t, 45, health.Timeout)
	assert.True(t, health.Redis.Enabled)
	assert.True(t, health.MySQL.Enabled)

	err := health.Validate()
	assert.NoError(t, err)
}
