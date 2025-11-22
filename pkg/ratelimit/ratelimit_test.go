/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\ratelimit\ratelimit_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package ratelimit

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRateLimit_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "ratelimit", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, StrategyTokenBucket, config.Strategy)
	assert.Equal(t, ScopeGlobal, config.DefaultScope)
	assert.NotNil(t, config.GlobalLimit)
	assert.Equal(t, 100, config.GlobalLimit.RequestsPerSecond)
	assert.Equal(t, 200, config.GlobalLimit.BurstSize)
	assert.Equal(t, time.Minute, config.GlobalLimit.WindowSize)
	assert.Equal(t, time.Minute, config.GlobalLimit.BlockDuration)
	assert.Equal(t, "memory", config.Storage.Type)
	assert.Equal(t, "rate_limit:", config.Storage.KeyPrefix)
	assert.Equal(t, 5*time.Minute, config.Storage.CleanInterval)
	assert.Empty(t, config.Routes)
	assert.Empty(t, config.IPRules)
	assert.Empty(t, config.UserRules)
	assert.False(t, config.EnableDynamicRule)
}

func TestRateLimit_WithGlobalLimit(t *testing.T) {
	limit := &LimitRule{
		RequestsPerSecond: 200,
		BurstSize:         400,
		WindowSize:        2 * time.Minute,
		BlockDuration:     2 * time.Minute,
	}
	config := Default().WithGlobalLimit(limit)
	assert.Equal(t, 200, config.GlobalLimit.RequestsPerSecond)
	assert.Equal(t, 400, config.GlobalLimit.BurstSize)
}

func TestRateLimit_AddRouteLimit(t *testing.T) {
	route := RouteLimit{
		Path:    "/api/users",
		Methods: []string{"GET", "POST"},
		Limit: &LimitRule{
			RequestsPerSecond: 50,
			BurstSize:         100,
		},
		PerUser: true,
		PerIP:   false,
	}
	config := Default().AddRouteLimit(route)
	assert.Len(t, config.Routes, 1)
	assert.Equal(t, "/api/users", config.Routes[0].Path)
	assert.Equal(t, []string{"GET", "POST"}, config.Routes[0].Methods)
	assert.True(t, config.Routes[0].PerUser)
}

func TestRateLimit_AddIPRule(t *testing.T) {
	rule := IPRule{
		IP: "192.168.1.0/24",
		Limit: &LimitRule{
			RequestsPerSecond: 1000,
			BurstSize:         2000,
		},
		Type:     "whitelist",
		Priority: 10,
	}
	config := Default().AddIPRule(rule)
	assert.Len(t, config.IPRules, 1)
	assert.Equal(t, "192.168.1.0/24", config.IPRules[0].IP)
	assert.Equal(t, "whitelist", config.IPRules[0].Type)
	assert.Equal(t, 10, config.IPRules[0].Priority)
}

func TestRateLimit_AddUserRule(t *testing.T) {
	rule := UserRule{
		UserID:   "premium-user-*",
		UserType: "premium",
		Role:     "admin",
		Limit: &LimitRule{
			RequestsPerSecond: 500,
			BurstSize:         1000,
		},
		Priority: 5,
	}
	config := Default().AddUserRule(rule)
	assert.Len(t, config.UserRules, 1)
	assert.Equal(t, "premium-user-*", config.UserRules[0].UserID)
	assert.Equal(t, "premium", config.UserRules[0].UserType)
	assert.Equal(t, "admin", config.UserRules[0].Role)
	assert.Equal(t, 5, config.UserRules[0].Priority)
}

func TestRateLimit_WithStrategy(t *testing.T) {
	config := Default().WithStrategy(StrategySlidingWindow)
	assert.Equal(t, StrategySlidingWindow, config.Strategy)
}

func TestRateLimit_WithStorage(t *testing.T) {
	storage := StorageConfig{
		Type:          "redis",
		KeyPrefix:     "custom_prefix:",
		CleanInterval: 10 * time.Minute,
	}
	config := Default().WithStorage(storage)
	assert.Equal(t, "redis", config.Storage.Type)
	assert.Equal(t, "custom_prefix:", config.Storage.KeyPrefix)
	assert.Equal(t, 10*time.Minute, config.Storage.CleanInterval)
}

func TestRateLimit_EnableDynamic(t *testing.T) {
	config := Default().EnableDynamic()
	assert.True(t, config.EnableDynamicRule)
}

func TestRateLimit_WithCustomLoader(t *testing.T) {
	config := Default().WithCustomLoader("custom-loader")
	assert.Equal(t, "custom-loader", config.CustomRuleLoader)
}

func TestRateLimit_Enable(t *testing.T) {
	config := Default().Disable().Enable()
	assert.True(t, config.Enabled)
}

func TestRateLimit_Disable(t *testing.T) {
	config := Default().Disable()
	assert.False(t, config.Enabled)
}

func TestRateLimit_IsEnabled(t *testing.T) {
	config := Default()
	assert.True(t, config.IsEnabled())
	config.Disable()
	assert.False(t, config.IsEnabled())
}

func TestRateLimit_Clone(t *testing.T) {
	original := Default().
		WithStrategy(StrategySlidingWindow).
		AddRouteLimit(RouteLimit{Path: "/api/test"}).
		AddIPRule(IPRule{IP: "127.0.0.1", Type: "whitelist"}).
		AddUserRule(UserRule{UserID: "user1", Priority: 1}).
		EnableDynamic().
		WithCustomLoader("test-loader")

	cloned := original.Clone().(*RateLimit)

	// 验证值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.Strategy, cloned.Strategy)
	assert.Equal(t, original.DefaultScope, cloned.DefaultScope)
	assert.Equal(t, original.CustomRuleLoader, cloned.CustomRuleLoader)
	assert.Equal(t, original.EnableDynamicRule, cloned.EnableDynamicRule)

	// 验证GlobalLimit独立性
	if original.GlobalLimit != nil && cloned.GlobalLimit != nil {
		assert.Equal(t, original.GlobalLimit.RequestsPerSecond, cloned.GlobalLimit.RequestsPerSecond)
		cloned.GlobalLimit.RequestsPerSecond = 999
		assert.NotEqual(t, original.GlobalLimit.RequestsPerSecond, cloned.GlobalLimit.RequestsPerSecond)
	}

	// 验证切片独立性
	cloned.Routes = append(cloned.Routes, RouteLimit{Path: "/new"})
	assert.NotEqual(t, len(original.Routes), len(cloned.Routes))

	cloned.IPRules = append(cloned.IPRules, IPRule{IP: "10.0.0.1"})
	assert.NotEqual(t, len(original.IPRules), len(cloned.IPRules))

	cloned.UserRules = append(cloned.UserRules, UserRule{UserID: "user2"})
	assert.NotEqual(t, len(original.UserRules), len(cloned.UserRules))
}

func TestRateLimit_Get(t *testing.T) {
	config := Default().WithStrategy(StrategyLeakyBucket)
	got := config.Get()
	assert.NotNil(t, got)
	rateLimitConfig, ok := got.(*RateLimit)
	assert.True(t, ok)
	assert.Equal(t, StrategyLeakyBucket, rateLimitConfig.Strategy)
}

func TestRateLimit_Set(t *testing.T) {
	config := Default()
	newConfig := &RateLimit{
		ModuleName:   "new-ratelimit",
		Enabled:      false,
		Strategy:     StrategyFixedWindow,
		DefaultScope: ScopePerIP,
	}

	config.Set(newConfig)
	assert.Equal(t, "new-ratelimit", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, StrategyFixedWindow, config.Strategy)
	assert.Equal(t, ScopePerIP, config.DefaultScope)
}

func TestRateLimit_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestRateLimit_ChainedCalls(t *testing.T) {
	config := Default().
		Enable().
		WithStrategy(StrategySlidingWindow).
		WithGlobalLimit(&LimitRule{RequestsPerSecond: 300}).
		AddRouteLimit(RouteLimit{Path: "/api/chain"}).
		AddIPRule(IPRule{IP: "192.168.0.1", Type: "blacklist"}).
		AddUserRule(UserRule{UserID: "chain-user", Priority: 10}).
		EnableDynamic().
		WithCustomLoader("chain-loader")

	assert.True(t, config.Enabled)
	assert.Equal(t, StrategySlidingWindow, config.Strategy)
	assert.Equal(t, 300, config.GlobalLimit.RequestsPerSecond)
	assert.Len(t, config.Routes, 1)
	assert.Len(t, config.IPRules, 1)
	assert.Len(t, config.UserRules, 1)
	assert.True(t, config.EnableDynamicRule)
	assert.Equal(t, "chain-loader", config.CustomRuleLoader)
}

func TestRateLimit_Strategies(t *testing.T) {
	assert.Equal(t, Strategy("token-bucket"), StrategyTokenBucket)
	assert.Equal(t, Strategy("leaky-bucket"), StrategyLeakyBucket)
	assert.Equal(t, Strategy("sliding-window"), StrategySlidingWindow)
	assert.Equal(t, Strategy("fixed-window"), StrategyFixedWindow)
}

func TestRateLimit_Scopes(t *testing.T) {
	assert.Equal(t, Scope("global"), ScopeGlobal)
	assert.Equal(t, Scope("per-ip"), ScopePerIP)
	assert.Equal(t, Scope("per-user"), ScopePerUser)
	assert.Equal(t, Scope("per-route"), ScopePerRoute)
}

func TestRateLimit_RouteLimit(t *testing.T) {
	route := RouteLimit{
		Path:      "/api/test",
		Methods:   []string{"GET", "POST"},
		Limit:     &LimitRule{RequestsPerSecond: 100},
		PerUser:   true,
		PerIP:     true,
		Whitelist: []string{"admin"},
		Blacklist: []string{"banned"},
	}

	assert.Equal(t, "/api/test", route.Path)
	assert.Len(t, route.Methods, 2)
	assert.True(t, route.PerUser)
	assert.True(t, route.PerIP)
	assert.Len(t, route.Whitelist, 1)
	assert.Len(t, route.Blacklist, 1)
}
