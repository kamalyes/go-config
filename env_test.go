package goconfig

import (
	"os"
	"testing"
	"time"
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
