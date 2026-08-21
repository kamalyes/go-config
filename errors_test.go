/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 11:15:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-08-21 20:10:56
 * @FilePath: \go-config\errors_test.go
 * @Description: 错误测试
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package goconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorConstructors(t *testing.T) {
	inner := errors.New("inner err")
	tests := []struct {
		name   string
		err    error
		expect string
	}{
		{"CallbackIDExists", ErrCallbackIDExists("cb1"), "cb1"},
		{"CallbackIDNotFound", ErrCallbackIDNotFound("cb1"), "cb1"},
		{"EnvTypeRegistered", ErrEnvTypeRegistered(EnvProduction), "已注册"},
		{"EnvVarNotSet", ErrEnvVarNotSet("APP_ENV"), "APP_ENV"},
		{"ConfigTypeMismatch", ErrConfigTypeMismatch("a", 1), "配置类型不匹配"},
		{"ModuleNotFound", ErrModuleNotFound("redis"), "redis"},
		{"ModuleConfigEmpty", ErrModuleConfigEmpty("redis"), "redis"},
		{"ModuleDefaultFuncNil", ErrModuleDefaultFuncNil("redis"), "redis"},
		{"PartialModuleFailed", ErrPartialModuleFailed(2, 1), "失败 1"},
		{"ModuleValidationFailed", ErrModuleValidationFailed([]string{"a", "b"}), "a"},
		{"InvalidTimezone", ErrInvalidTimezone("UTC", inner), "UTC"},
		{"CronSpecEmpty", ErrCronSpecEmpty("job1"), "job1"},
		{"TimeoutInvalid", ErrTimeoutInvalid("job1"), "job1"},
		{"PriorityOutOfRange", ErrPriorityOutOfRange("job1"), "job1"},
		{"MaxConcurrentInvalid", ErrMaxConcurrentInvalid("job1"), "job1"},
		{"BreakerMaxFailuresInvalid", ErrBreakerMaxFailuresInvalid("job1"), "job1"},
		{"BreakerResetTimeoutInvalid", ErrBreakerResetTimeoutInvalid("job1"), "job1"},
		{"BreakerHalfOpenSuccessesInvalid", ErrBreakerHalfOpenSuccessesInvalid("job1"), "job1"},
		{"ReadConfigFile", ErrReadConfigFile(inner), "读取配置文件失败"},
		{"UnmarshalConfig", ErrUnmarshalConfig(inner), "解析配置失败"},
		{"CreateHotReloader", ErrCreateHotReloader(inner), "创建热更新器失败"},
		{"StartHotReloader", ErrStartHotReloader(inner), "启动热更新器失败"},
		{"StopHotReloader", ErrStopHotReloader(inner), "停止热更新器失败"},
		{"CreateWatcher", ErrCreateWatcher(inner), "创建文件监控器失败"},
		{"GetAbsPath", ErrGetAbsPath(inner), "绝对路径"},
		{"AddWatcher", ErrAddWatcher(inner), "添加配置目录监控失败"},
		{"CreateOutputDir", ErrCreateOutputDir(inner), "创建输出目录失败"},
		{"GenerateYAML", ErrGenerateYAML(inner), "生成YAML配置失败"},
		{"GenerateJSON", ErrGenerateJSON(inner), "生成JSON配置失败"},
		{"MarshalYAML", ErrMarshalYAML(inner), "序列化YAML失败"},
		{"WriteYAML", ErrWriteYAML(inner), "写入YAML文件失败"},
		{"MarshalJSON", ErrMarshalJSON(inner), "序列化JSON失败"},
		{"WriteJSON", ErrWriteJSON(inner), "写入JSON文件失败"},
		{"SearchPathNotExist", ErrSearchPathNotExist("/x"), "/x"},
		{"ConfigFileNotFound", ErrConfigFileNotFound("/x"), "/x"},
		{"CreateDir", ErrCreateDir(inner), "创建目录失败"},
		{"GenerateDefaultConfig", ErrGenerateDefaultConfig(inner), "生成默认配置失败"},
		{"WriteConfigFile", ErrWriteConfigFile(inner), "写入配置文件失败"},
		{"ScanDir", ErrScanDir(inner), "扫描目录失败"},
		{"UnsupportedFormat", ErrUnsupportedFormat(".xyz"), ".xyz"},
		{"ResolveConfigPath", ErrResolveConfigPath(inner), "解析配置路径失败"},
		{"CreateManager", ErrCreateManager(inner), "创建集成配置管理器失败"},
		{"StartManager", ErrStartManager(inner), "启动管理器失败"},
		{"FindConfigByPattern", ErrFindConfigByPattern(inner), "按模式查找配置文件失败"},
		{"NoMatchingConfig", ErrNoMatchingConfig("dev*"), "dev*"},
		{"ModuleConfigSerializeFailed", ErrModuleConfigSerializeFailed("redis", inner), "redis"},
		{"ModulesValidationFailed", ErrModulesValidationFailed("a, b"), "a, b"},
		{"DiscoverConfigFiles", ErrDiscoverConfigFiles(inner), "发现配置文件失败"},
		{"AutoDiscoverConfigFiles", ErrAutoDiscoverConfigFiles(inner), "自动发现配置文件失败"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Error(t, tt.err)
			assert.Contains(t, tt.err.Error(), tt.expect)
		})
	}
}

func TestErrorVars(t *testing.T) {
	vars := []error{
		ErrCallbackIDEmpty, ErrConfigEmpty, ErrNoConfigPath,
		ErrManagerRunning, ErrManagerNotRunning, ErrReloaderRunning, ErrReloaderNotRunning,
		ErrNoConfigCandidate, ErrCallbackFuncNil,
	}
	for _, v := range vars {
		assert.NotEmpty(t, v.Error())
	}
}

func TestErrorWrapping(t *testing.T) {
	inner := errors.New("root cause")
	wrapped := []error{
		ErrInvalidTimezone("UTC", inner),
		ErrReadConfigFile(inner), ErrUnmarshalConfig(inner),
		ErrCreateHotReloader(inner), ErrStartHotReloader(inner), ErrStopHotReloader(inner),
		ErrCreateWatcher(inner), ErrGetAbsPath(inner), ErrAddWatcher(inner),
		ErrCreateOutputDir(inner), ErrGenerateYAML(inner), ErrGenerateJSON(inner),
		ErrMarshalYAML(inner), ErrWriteYAML(inner), ErrMarshalJSON(inner), ErrWriteJSON(inner),
		ErrCreateDir(inner), ErrGenerateDefaultConfig(inner), ErrWriteConfigFile(inner),
		ErrScanDir(inner), ErrResolveConfigPath(inner), ErrCreateManager(inner),
		ErrStartManager(inner), ErrFindConfigByPattern(inner),
		ErrModuleConfigSerializeFailed("redis", inner),
		ErrDiscoverConfigFiles(inner), ErrAutoDiscoverConfigFiles(inner),
	}
	for _, err := range wrapped {
		assert.True(t, errors.Is(err, inner), "error should wrap inner: %v", err)
	}
}
