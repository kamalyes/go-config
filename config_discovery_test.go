/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-01-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-01-30 00:00:00
 * @FilePath: \go-config\config_discovery_test.go
 * @Description: 配置文件发现器测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewConfigDiscovery 测试创建配置发现器
func TestNewConfigDiscovery(t *testing.T) {
	discovery := NewConfigDiscovery()

	assert.NotNil(t, discovery)
	assert.NotEmpty(t, discovery.SupportedExtensions)
	assert.NotEmpty(t, discovery.DefaultNames)
	assert.NotEmpty(t, discovery.EnvPrefixes)
}

// TestConfigDiscovery_DiscoverConfigFiles 测试发现配置文件
func TestConfigDiscovery_DiscoverConfigFiles(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()

	// 创建测试配置文件
	testFiles := []string{
		"config.yaml",
		"config.json",
		"config-dev.yaml",
		"config-prod.yaml",
	}

	for _, file := range testFiles {
		path := filepath.Join(tmpDir, file)
		err := os.WriteFile(path, []byte("test: value"), 0644)
		require.NoError(t, err)
	}

	discovery := NewConfigDiscovery()
	files, err := discovery.DiscoverConfigFiles(tmpDir, EnvDevelopment)

	require.NoError(t, err)
	assert.NotEmpty(t, files)

	// 验证发现的文件
	existingFiles := 0
	for _, file := range files {
		if file.Exists {
			existingFiles++
		}
	}
	// 至少应该找到我们创建的文件（可能不是全部，取决于发现器的逻辑）
	assert.GreaterOrEqual(t, existingFiles, 3, "应该至少发现3个配置文件")
	assert.LessOrEqual(t, existingFiles, len(testFiles), "发现的文件数不应超过创建的文件数")
}

// TestConfigDiscovery_DiscoverConfigFiles_NonExistentPath 测试不存在的路径
func TestConfigDiscovery_DiscoverConfigFiles_NonExistentPath(t *testing.T) {
	discovery := NewConfigDiscovery()
	files, err := discovery.DiscoverConfigFiles("/non/existent/path", EnvDevelopment)

	assert.Error(t, err)
	assert.Nil(t, files)
}

// TestConfigDiscovery_DiscoverConfigFiles_EmptyDirectory 测试空目录
func TestConfigDiscovery_DiscoverConfigFiles_EmptyDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	discovery := NewConfigDiscovery()
	files, err := discovery.DiscoverConfigFiles(tmpDir, EnvDevelopment)

	require.NoError(t, err)
	assert.NotEmpty(t, files)

	// 验证所有文件都不存在
	for _, file := range files {
		assert.False(t, file.Exists)
	}
}

// TestConfigDiscovery_DiscoverConfigFiles_Priority 测试文件优先级
func TestConfigDiscovery_DiscoverConfigFiles_Priority(t *testing.T) {
	tmpDir := t.TempDir()

	// 创建不同优先级的配置文件
	testFiles := []string{
		"config.yaml",
		"config-dev.yaml",
		"config-prod.yaml",
	}

	for _, file := range testFiles {
		path := filepath.Join(tmpDir, file)
		err := os.WriteFile(path, []byte("test: value"), 0644)
		require.NoError(t, err)
	}

	discovery := NewConfigDiscovery()
	files, err := discovery.DiscoverConfigFiles(tmpDir, EnvDevelopment)

	require.NoError(t, err)
	assert.NotEmpty(t, files)

	// 验证文件按优先级排序
	var existingFiles []*ConfigFileInfo
	for _, file := range files {
		if file.Exists {
			existingFiles = append(existingFiles, file)
		}
	}

	// 第一个文件应该是优先级最高的
	assert.NotEmpty(t, existingFiles)
	assert.LessOrEqual(t, existingFiles[0].Priority, existingFiles[len(existingFiles)-1].Priority)
}

// TestConfigDiscovery_DiscoverConfigFiles_MultipleEnvironments 测试多环境发现
func TestConfigDiscovery_DiscoverConfigFiles_MultipleEnvironments(t *testing.T) {
	tmpDir := t.TempDir()

	// 创建多环境配置文件
	environments := []EnvironmentType{
		EnvDevelopment,
		EnvTest,
		EnvProduction,
	}

	for _, env := range environments {
		fileName := "config-" + string(env) + ".yaml"
		path := filepath.Join(tmpDir, fileName)
		err := os.WriteFile(path, []byte("test: value"), 0644)
		require.NoError(t, err)
	}

	discovery := NewConfigDiscovery()

	for _, env := range environments {
		files, err := discovery.DiscoverConfigFiles(tmpDir, env)
		require.NoError(t, err)
		assert.NotEmpty(t, files)

		// 验证环境特定的文件被发现
		found := false
		for _, file := range files {
			if file.Exists && file.Environment == env {
				found = true
				break
			}
		}
		assert.True(t, found, "环境 %s 的配置文件未被发现", env)
	}
}

// TestConfigDiscovery_DiscoverConfigFiles_DifferentExtensions 测试不同扩展名
func TestConfigDiscovery_DiscoverConfigFiles_DifferentExtensions(t *testing.T) {
	tmpDir := t.TempDir()

	// 创建不同扩展名的配置文件
	extensions := []string{".yaml", ".yml", ".json", ".toml"}
	for _, ext := range extensions {
		fileName := "config" + ext
		path := filepath.Join(tmpDir, fileName)
		err := os.WriteFile(path, []byte("test: value"), 0644)
		require.NoError(t, err)
	}

	discovery := NewConfigDiscovery()
	files, err := discovery.DiscoverConfigFiles(tmpDir, EnvDevelopment)

	require.NoError(t, err)
	assert.NotEmpty(t, files)

	// 验证所有扩展名的文件都被发现
	foundExtensions := make(map[string]bool)
	for _, file := range files {
		if file.Exists {
			foundExtensions[file.Extension] = true
		}
	}

	for _, ext := range extensions {
		assert.True(t, foundExtensions[ext], "扩展名 %s 的文件未被发现", ext)
	}
}

// TestConfigFileInfo_Fields 测试配置文件信息字段
func TestConfigFileInfo_Fields(t *testing.T) {
	info := &ConfigFileInfo{
		Path:        "/path/to/config.yaml",
		Name:        "config.yaml",
		BaseName:    "config",
		Extension:   ".yaml",
		Environment: EnvDevelopment,
		Priority:    1,
		Exists:      true,
	}

	assert.Equal(t, "/path/to/config.yaml", info.Path)
	assert.Equal(t, "config.yaml", info.Name)
	assert.Equal(t, "config", info.BaseName)
	assert.Equal(t, ".yaml", info.Extension)
	assert.Equal(t, EnvDevelopment, info.Environment)
	assert.Equal(t, 1, info.Priority)
	assert.True(t, info.Exists)
}

// TestConfigDiscovery_CustomExtensions 测试自定义扩展名
func TestConfigDiscovery_CustomExtensions(t *testing.T) {
	tmpDir := t.TempDir()

	// 创建自定义扩展名的配置文件
	customExt := ".custom"
	fileName := "config" + customExt
	path := filepath.Join(tmpDir, fileName)
	err := os.WriteFile(path, []byte("test: value"), 0644)
	require.NoError(t, err)

	discovery := NewConfigDiscovery()
	discovery.SupportedExtensions = append(discovery.SupportedExtensions, customExt)

	files, err := discovery.DiscoverConfigFiles(tmpDir, EnvDevelopment)
	require.NoError(t, err)

	// 验证自定义扩展名的文件被发现
	found := false
	for _, file := range files {
		if file.Exists && file.Extension == customExt {
			found = true
			break
		}
	}
	assert.True(t, found, "自定义扩展名 %s 的文件未被发现", customExt)
}

// TestConfigDiscovery_CustomDefaultNames 测试自定义默认名称
func TestConfigDiscovery_CustomDefaultNames(t *testing.T) {
	tmpDir := t.TempDir()

	// 创建自定义名称的配置文件
	customName := "myapp"
	fileName := customName + ".yaml"
	path := filepath.Join(tmpDir, fileName)
	err := os.WriteFile(path, []byte("test: value"), 0644)
	require.NoError(t, err)

	discovery := NewConfigDiscovery()
	discovery.DefaultNames = []string{customName}

	files, err := discovery.DiscoverConfigFiles(tmpDir, EnvDevelopment)
	require.NoError(t, err)

	// 验证自定义名称的文件被发现
	found := false
	for _, file := range files {
		if file.Exists && file.BaseName == customName {
			found = true
			break
		}
	}
	assert.True(t, found, "自定义名称 %s 的文件未被发现", customName)
}

// TestConfigDiscovery_FindBestMatch 测试查找最佳匹配
func TestConfigDiscovery_FindBestMatch(t *testing.T) {
	tmpDir := t.TempDir()

	// 创建多个配置文件
	testFiles := []string{
		"config.yaml",
		"config-dev.yaml",
		"config-prod.yaml",
	}

	for _, file := range testFiles {
		path := filepath.Join(tmpDir, file)
		err := os.WriteFile(path, []byte("test: value"), 0644)
		require.NoError(t, err)
	}

	discovery := NewConfigDiscovery()
	files, err := discovery.DiscoverConfigFiles(tmpDir, EnvDevelopment)
	require.NoError(t, err)

	// 查找第一个存在的文件（优先级最高）
	var bestMatch *ConfigFileInfo
	for _, file := range files {
		if file.Exists {
			bestMatch = file
			break
		}
	}

	assert.NotNil(t, bestMatch)
	assert.True(t, bestMatch.Exists)
}
