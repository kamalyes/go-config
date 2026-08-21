/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-26 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-08-21 20:09:39
 * @FilePath: \go-config\config_unmarshaler_test.go
 * @Description: 配置反序列化器测试
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlexibleMatchName(t *testing.T) {
	tests := []struct {
		name      string
		mapKey    string
		fieldName string
		expect    bool
	}{
		{"direct match", "name", "name", true},
		{"case insensitive", "Name", "name", true},
		{"kebab to snake", "service-name", "service_name", true},
		{"snake to kebab", "service_name", "service-name", true},
		{"camelCase to snake", "serviceName", "service_name", true},
		{"PascalCase to snake", "ServiceName", "service_name", true},
		{"no match", "foo", "bar", false},
		{"empty both", "", "", true},
		{"empty mapKey nonempty field", "", "name", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, FlexibleMatchName(tt.mapKey, tt.fieldName))
		})
	}
}

type flexConfig struct {
	ModuleName  string `mapstructure:"module-name"`
	ReadTimeout int    `mapstructure:"read-timeout"`
}

func TestUnmarshalWithFlexibleNaming_JSONSnake(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	require.NoError(t, os.WriteFile(path, []byte(`{"module_name":"test","read_timeout":30}`), 0644))
	v := viper.New()
	v.SetConfigFile(path)
	require.NoError(t, v.ReadInConfig())

	var cfg flexConfig
	require.NoError(t, UnmarshalWithFlexibleNaming(v, &cfg))
	assert.Equal(t, "test", cfg.ModuleName)
	assert.Equal(t, 30, cfg.ReadTimeout)
}

func TestUnmarshalWithFlexibleNaming_YAMLKebab(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	require.NoError(t, os.WriteFile(path, []byte("module-name: test\nread-timeout: 30\n"), 0644))
	v := viper.New()
	v.SetConfigFile(path)
	require.NoError(t, v.ReadInConfig())

	var cfg flexConfig
	require.NoError(t, UnmarshalWithFlexibleNaming(v, &cfg))
	assert.Equal(t, "test", cfg.ModuleName)
	assert.Equal(t, 30, cfg.ReadTimeout)
}

type snakeConfig struct {
	ModuleName  string `mapstructure:"module_name"`
	ReadTimeout int    `mapstructure:"read_timeout"`
}

func TestUnmarshalWithKebabToSnake(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	require.NoError(t, os.WriteFile(path, []byte("module-name: test\nread-timeout: 30\n"), 0644))
	v := viper.New()
	v.SetConfigFile(path)
	require.NoError(t, v.ReadInConfig())

	var cfg snakeConfig
	require.NoError(t, UnmarshalWithKebabToSnake(v, &cfg))
	assert.Equal(t, "test", cfg.ModuleName)
	assert.Equal(t, 30, cfg.ReadTimeout)
}
