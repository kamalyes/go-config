/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 11:15:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 11:31:57
 * @FilePath: \go-config\env_test.go
 * @Description: 环境变量测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package goconfig

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvironmentJudgementFunctions(t *testing.T) {
	// 保存原始环境
	originalEnv := GetCurrentEnvironment()
	defer func() {
		SetCurrentEnvironment(originalEnv)
	}()

	tests := []struct {
		env       EnvironmentType
		isDev     bool
		isTest    bool
		isStaging bool
		isUAT     bool
		isProd    bool
		isLocal   bool
		isDebug   bool
		isDemo    bool
		isInteg   bool
		level     int
	}{
		{EnvDevelopment, true, false, false, false, false, false, false, false, false, 2},
		{EnvTest, false, true, false, false, false, false, false, false, false, 3},
		{EnvStaging, false, false, true, false, false, false, false, false, false, 5},
		{EnvUAT, false, false, false, true, false, false, false, false, false, 5},
		{EnvProduction, false, false, false, false, true, false, false, false, false, 10},
		{EnvLocal, false, false, false, false, false, true, false, false, false, 1},
		{EnvDebug, false, false, false, false, false, false, true, false, false, 1},
		{EnvDemo, false, false, false, false, false, false, false, true, false, 6},
		{EnvIntegration, false, false, false, false, false, false, false, false, true, 4},
	}

	for _, tt := range tests {
		t.Run(string(tt.env), func(t *testing.T) {
			// 设置测试环境
			SetCurrentEnvironment(tt.env)

			// 验证当前环境获取
			if current := GetCurrentEnvironment(); current != tt.env {
				t.Errorf("GetCurrentEnvironment() = %v, want %v", current, tt.env)
			}

			// 验证各个判断函数
			if got := IsDev(); got != tt.isDev {
				t.Errorf("IsDev() = %v, want %v", got, tt.isDev)
			}
			if got := IsTest(); got != tt.isTest {
				t.Errorf("IsTest() = %v, want %v", got, tt.isTest)
			}
			if got := IsStaging(); got != tt.isStaging {
				t.Errorf("IsStaging() = %v, want %v", got, tt.isStaging)
			}
			if got := IsUAT(); got != tt.isUAT {
				t.Errorf("IsUAT() = %v, want %v", got, tt.isUAT)
			}
			if got := IsProduction(); got != tt.isProd {
				t.Errorf("IsProduction() = %v, want %v", got, tt.isProd)
			}
			if got := IsLocal(); got != tt.isLocal {
				t.Errorf("IsLocal() = %v, want %v", got, tt.isLocal)
			}
			if got := IsDebug(); got != tt.isDebug {
				t.Errorf("IsDebug() = %v, want %v", got, tt.isDebug)
			}
			if got := IsDemo(); got != tt.isDemo {
				t.Errorf("IsDemo() = %v, want %v", got, tt.isDemo)
			}
			if got := IsIntegration(); got != tt.isInteg {
				t.Errorf("IsIntegration() = %v, want %v", got, tt.isInteg)
			}

			// 验证环境级别
			if got := GetEnvironmentLevel(tt.env); got != tt.level {
				t.Errorf("GetEnvironmentLevel(%v) = %v, want %v", tt.env, got, tt.level)
			}
			if got := GetCurrentEnvironmentLevel(); got != tt.level {
				t.Errorf("GetCurrentEnvironmentLevel() = %v, want %v", got, tt.level)
			}
		})
	}
}

func TestIsEnvironment(t *testing.T) {
	// 保存原始环境
	originalEnv := GetCurrentEnvironment()
	defer func() {
		SetCurrentEnvironment(originalEnv)
	}()

	SetCurrentEnvironment(EnvDevelopment)

	if !IsEnvironment(EnvDevelopment) {
		t.Error("IsEnvironment(EnvDevelopment) should return true when current env is development")
	}

	if IsEnvironment(EnvProduction) {
		t.Error("IsEnvironment(EnvProduction) should return false when current env is development")
	}
}

func TestIsAnyOf(t *testing.T) {
	// 保存原始环境
	originalEnv := GetCurrentEnvironment()
	defer func() {
		SetCurrentEnvironment(originalEnv)
	}()

	SetCurrentEnvironment(EnvTest)

	if !IsAnyOf(EnvTest, EnvDevelopment, EnvStaging) {
		t.Error("IsAnyOf should return true when current env matches one of the provided envs")
	}

	if IsAnyOf(EnvProduction, EnvLocal) {
		t.Error("IsAnyOf should return false when current env doesn't match any of the provided envs")
	}

	// 测试空参数
	if IsAnyOf() {
		t.Error("IsAnyOf with no parameters should return false")
	}
}

func TestEnvironmentLevelFunctions(t *testing.T) {
	// 保存原始环境
	originalEnv := GetCurrentEnvironment()
	defer func() {
		SetCurrentEnvironment(originalEnv)
	}()

	// 测试生产级别
	SetCurrentEnvironment(EnvProduction)
	if !IsProductionLevel() {
		t.Error("IsProductionLevel() should return true for production environment")
	}
	if IsTestingLevel() {
		t.Error("IsTestingLevel() should return false for production environment")
	}
	if IsDevelopmentLevel() {
		t.Error("IsDevelopmentLevel() should return false for production environment")
	}

	// 测试测试级别
	SetCurrentEnvironment(EnvStaging)
	if IsProductionLevel() {
		t.Error("IsProductionLevel() should return false for staging environment")
	}
	if !IsTestingLevel() {
		t.Error("IsTestingLevel() should return true for staging environment")
	}
	if IsDevelopmentLevel() {
		t.Error("IsDevelopmentLevel() should return false for staging environment")
	}

	// 测试开发级别
	SetCurrentEnvironment(EnvLocal)
	if IsProductionLevel() {
		t.Error("IsProductionLevel() should return false for local environment")
	}
	if IsTestingLevel() {
		t.Error("IsTestingLevel() should return false for local environment")
	}
	if !IsDevelopmentLevel() {
		t.Error("IsDevelopmentLevel() should return true for local environment")
	}
}

func TestGlobalEnvironmentAutoInit(t *testing.T) {
	// 测试全局环境实例是否已自动初始化
	env := GetGlobalEnvironment()
	if env == nil {
		t.Error("Global environment should be auto-initialized")
	}

	// 测试多次调用返回同一实例
	env2 := GetGlobalEnvironment()
	if env != env2 {
		t.Error("GetGlobalEnvironment() should return the same instance")
	}
}

func TestEnvironmentFromOSEnv(t *testing.T) {
	// 保存原始环境变量
	originalOSEnv := os.Getenv("APP_ENV")
	defer func() {
		if originalOSEnv != "" {
			os.Setenv("APP_ENV", originalOSEnv)
		} else {
			os.Unsetenv("APP_ENV")
		}
	}()

	// 测试从操作系统环境变量读取
	os.Setenv("APP_ENV", "production")

	// 创建新的环境实例
	env := NewEnvironment()
	defer env.StopWatch()              // 停止监控 goroutine，避免泄漏的 watchEnv 周期性覆盖全局 APP_ENV 干扰其他测试
	time.Sleep(100 * time.Millisecond) // 等待环境检查

	if env.Value != EnvProduction {
		t.Errorf("Environment should detect OS env variable, got %v, want %v", env.Value, EnvProduction)
	}
}

func TestConvenienceFunctions(t *testing.T) {
	// 保存原始环境
	originalEnv := GetCurrentEnvironment()
	defer func() {
		SetCurrentEnvironment(originalEnv)
	}()

	// 测试便捷设置函数
	SetCurrentEnvironment(EnvStaging)
	if current := GetCurrentEnvironment(); current != EnvStaging {
		t.Errorf("SetCurrentEnvironment/GetCurrentEnvironment failed, got %v, want %v", current, EnvStaging)
	}

	// 验证对应的判断函数
	if !IsStaging() {
		t.Error("IsStaging() should return true after SetCurrentEnvironment(EnvStaging)")
	}
}

// --- 包级别环境前缀管理 ---

func TestRegisterAndGetEnvPrefixes(t *testing.T) {
	customEnv := EnvironmentType("test-custom-env")
	t.Cleanup(func() {
		envPrefixesMutex.Lock()
		delete(DefaultEnvPrefixes, customEnv)
		envPrefixesMutex.Unlock()
	})

	RegisterEnvPrefixes(customEnv, "custom", "my-env")

	prefixes := GetEnvPrefixes(customEnv)
	assert.Contains(t, prefixes, "custom")
	assert.Contains(t, prefixes, "my-env")

	// 返回副本：外部修改不影响内部存储
	prefixes[0] = "modified"
	again := GetEnvPrefixes(customEnv)
	assert.NotContains(t, again, "modified")
}

func TestGetEnvPrefixes_NotRegistered(t *testing.T) {
	assert.Nil(t, GetEnvPrefixes(EnvironmentType("non-existent-env-xyz")))
}

func TestListAllEnvPrefixes(t *testing.T) {
	all := ListAllEnvPrefixes()
	assert.NotEmpty(t, all)
	assert.Contains(t, all, EnvDevelopment)

	// 返回副本
	all[EnvDevelopment][0] = "modified-copy"
	again := ListAllEnvPrefixes()
	assert.NotContains(t, again[EnvDevelopment], "modified-copy")
}

func TestEnvironmentType_IsValid(t *testing.T) {
	assert.True(t, EnvDevelopment.IsValid())
	assert.True(t, EnvProduction.IsValid())
	assert.False(t, EnvironmentType("invalid-env-xyz").IsValid())
}

func TestGetGlobalEnvManager(t *testing.T) {
	mgr := GetGlobalEnvManager()
	require.NotNil(t, mgr)
	assert.Same(t, mgr, GetGlobalEnvManager())
}

// --- EnvironmentManager ---

func TestEnvironmentManager_RegisterAndIsRegistered(t *testing.T) {
	mgr := NewEnvironmentManager()
	customEnv := EnvironmentType("test-mgr-env")
	assert.False(t, mgr.IsRegistered(customEnv))

	mgr.RegisterEnvironment(customEnv, "alias1", "alias2")
	assert.True(t, mgr.IsRegistered(customEnv))
	assert.True(t, mgr.IsEnvironment("alias1", customEnv))
	assert.True(t, mgr.IsEnvironment("ALIAS1", customEnv)) // 大小写不敏感
	assert.False(t, mgr.IsEnvironment("alias1", EnvDevelopment))
}

func TestEnvironmentManager_IsRegistered_Default(t *testing.T) {
	mgr := NewEnvironmentManager()
	assert.True(t, mgr.IsRegistered(EnvDevelopment))
	assert.False(t, mgr.IsRegistered(EnvironmentType("nope")))
}

// --- Environment 实例回调管理 ---

func TestEnvironment_RegisterCallbackAndList(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()

	cb := func(old, newEnv EnvironmentType) error { return nil }
	require.NoError(t, env.RegisterCallback("list-cb-1", cb, 1, false))
	require.NoError(t, env.RegisterCallback("list-cb-2", cb, 2, false))

	ids := env.ListCallbacks()
	assert.Contains(t, ids, "list-cb-1")
	assert.Contains(t, ids, "list-cb-2")
	assert.Len(t, ids, 2)
}

func TestEnvironment_RegisterCallback_Errors(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()

	// 空ID
	assert.Error(t, env.RegisterCallback("", func(o, n EnvironmentType) error { return nil }, 1, false))

	// 重复ID
	cb := func(o, n EnvironmentType) error { return nil }
	require.NoError(t, env.RegisterCallback("dup-cb", cb, 1, false))
	assert.Error(t, env.RegisterCallback("dup-cb", cb, 1, false))
}

func TestEnvironment_UnregisterCallback(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()

	require.NoError(t, env.RegisterCallback("unreg-cb", func(o, n EnvironmentType) error { return nil }, 1, false))
	require.NoError(t, env.UnregisterCallback("unreg-cb"))
	assert.NotContains(t, env.ListCallbacks(), "unreg-cb")

	// 注销不存在的回调返回错误
	assert.Error(t, env.UnregisterCallback("non-existent"))
}

func TestEnvironment_ClearCallbacks(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()

	require.NoError(t, env.RegisterCallback("clear-cb-1", func(o, n EnvironmentType) error { return nil }, 1, false))
	require.NoError(t, env.RegisterCallback("clear-cb-2", func(o, n EnvironmentType) error { return nil }, 2, false))

	env.ClearCallbacks()
	assert.Empty(t, env.ListCallbacks())
}

// --- Environment 注册环境类型 ---

func TestEnvironment_RegisterEnvironment(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()

	customEnv := EnvironmentType("test-reg-env")
	require.NoError(t, env.RegisterEnvironment(customEnv))
	// 重复注册返回错误
	assert.Error(t, env.RegisterEnvironment(customEnv))
}

// --- Environment 其他方法 ---

func TestEnvironment_SetCheckFrequency(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()

	result := env.SetCheckFrequency(500 * time.Millisecond)
	assert.Same(t, env, result)
	assert.Equal(t, 500*time.Millisecond, env.CheckFrequency)
}

func TestEnvironment_ClearEnv(t *testing.T) {
	origKey := GetContextKey()
	origVal := os.Getenv(string(origKey))
	t.Cleanup(func() {
		_ = os.Setenv(string(origKey), origVal)
	})

	env := NewEnvironment()
	defer env.StopWatch()

	env.ClearEnv()
	assert.Empty(t, os.Getenv(string(GetContextKey())))
}

func TestEnvironment_SetEnvironment_TriggersCallback(t *testing.T) {
	env := NewEnvironment()
	defer env.StopWatch()

	triggered := false
	require.NoError(t, env.RegisterCallback("trigger-cb", func(old, newEnv EnvironmentType) error {
		triggered = true
		assert.Equal(t, EnvDevelopment, old)
		assert.Equal(t, EnvProduction, newEnv)
		return nil
	}, 1, false))

	// 先设一个已知值，再变更以触发回调
	env.SetEnvironment(EnvDevelopment)
	env.SetEnvironment(EnvProduction)
	assert.True(t, triggered)
}

// --- 上下文键管理 ---

func TestGetAndSetContextKey(t *testing.T) {
	origKey := GetContextKey()
	origVal := os.Getenv(string(origKey))
	t.Cleanup(func() {
		// 必须同时恢复全局上下文键与其 OS 环境变量，否则改键会泄漏到后续测试
		// （如 TestEnvironmentFromOSEnv / TestEnvironmentManager_Callbacks 假定键为 APP_ENV）
		contextKeyMutex.Lock()
		envContextKey = origKey
		contextKeyMutex.Unlock()
		if origVal == "" {
			_ = os.Unsetenv(string(origKey))
		} else {
			_ = os.Setenv(string(origKey), origVal)
		}
	})

	SetContextKey(&ContextKeyOptions{Key: "TEST_CTX_KEY", Value: EnvTest})
	assert.Equal(t, ContextKey("TEST_CTX_KEY"), GetContextKey())
	assert.Equal(t, "test", os.Getenv("TEST_CTX_KEY"))
}

func TestSetContextKey_NilOptions(t *testing.T) {
	origKey := GetContextKey()
	origVal := os.Getenv(string(origKey))
	t.Cleanup(func() {
		_ = os.Setenv(string(origKey), origVal)
	})

	// nil options 使用默认值，不应 panic
	assert.NotPanics(t, func() {
		SetContextKey(nil)
	})
}

// --- 默认构造函数 ---

func TestDefaultEnvironmentFuncs(t *testing.T) {
	def := Default()
	require.NotNil(t, def)
	assert.Equal(t, DefaultEnv, def.Value)

	defVal := DefaultEnvironment()
	assert.Equal(t, DefaultEnv, defVal.Value)
	assert.NotZero(t, defVal.CheckFrequency)
}
