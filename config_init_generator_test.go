/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-26 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-08-21 20:09:28
 * @FilePath: \go-config\config_init_generator_test.go
 * @Description: 配置初始化生成器测试
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"os"
	"testing"

	gologger "github.com/kamalyes/go-logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- 链式选项 ---

func TestSmartConfigGenerator_ChainedOptions(t *testing.T) {
	generator := NewSmartConfigGenerator("./test_output_extra")

	customLogger := gologger.NewLogger()
	result := generator.
		WithLogger(customLogger).
		WithForceRegenerate(true).
		WithIncludeComments(false).
		WithBackupExisting(false)

	assert.Same(t, generator, result)
	assert.Same(t, customLogger, generator.Logger)
	assert.True(t, generator.ForceRegenerate)
	assert.False(t, generator.IncludeComments)
	assert.False(t, generator.BackupExisting)
}

// --- toKebabCase ---

func TestSmartConfigGenerator_ToKebabCase(t *testing.T) {
	generator := NewSmartConfigGenerator("./test_output_extra")

	tests := []struct {
		input  string
		expect string
	}{
		{"", ""},
		{"name", "name"},
		{"ModuleName", "module-name"},
		{"HTTPServer", "h-t-t-p-server"},
		{"simpleCase", "simple-case"},
		{"already-lower", "already-lower"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expect, generator.toKebabCase(tt.input))
		})
	}
}

// --- GenerateModulesByNames ---

func TestGenerateModulesByNames_Empty(t *testing.T) {
	// 空参数应等价于 GenerateAllConfigs
	tempDir := "./test_output_gen_empty"
	defer os.RemoveAll(tempDir)
	generator := NewSmartConfigGenerator(tempDir)
	require.NoError(t, generator.EnableOnlyModules("health"))
	assert.NoError(t, generator.GenerateModulesByNames())
}

func TestGenerateModulesByNames_Valid(t *testing.T) {
	tempDir := "./test_output_gen_valid"
	defer os.RemoveAll(tempDir)
	generator := NewSmartConfigGenerator(tempDir)
	// 生成单个有效模块
	err := generator.GenerateModulesByNames("health")
	assert.NoError(t, err)
	// 验证文件生成
	_, statErr := os.Stat(tempDir + "/pkg/health/health.yaml")
	assert.NoError(t, statErr)
}

func TestGenerateModulesByNames_InvalidAndValid(t *testing.T) {
	tempDir := "./test_output_gen_mix"
	defer os.RemoveAll(tempDir)
	generator := NewSmartConfigGenerator(tempDir)
	// 包含无效模块名应返回 ErrPartialModuleFailed
	err := generator.GenerateModulesByNames("health", "non-existent-module")
	assert.Error(t, err)
}

func TestGenerateModulesByNames_AllInvalid(t *testing.T) {
	tempDir := "./test_output_gen_invalid"
	defer os.RemoveAll(tempDir)
	generator := NewSmartConfigGenerator(tempDir)
	err := generator.GenerateModulesByNames("no-such-module-1", "no-such-module-2")
	assert.Error(t, err)
}

// --- ValidateModuleConfig 错误路径 ---

func TestValidateModuleConfig_NotFound(t *testing.T) {
	generator := NewSmartConfigGenerator("./test_output_validate")
	err := generator.ValidateModuleConfig("non-existent-module-xyz")
	assert.Error(t, err)
}
