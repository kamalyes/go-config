/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:19:51
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 13:33:39
 * @FilePath: \go-config\config_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"os"
	"testing"

	"github.com/kamalyes/go-config/captcha"
	"github.com/kamalyes/go-config/database"
	"github.com/kamalyes/go-config/sms"
	"github.com/stretchr/testify/assert"
)

// 测试加载配置
func TestLoadConfig(t *testing.T) {
	// 创建一个临时的 YAML 配置文件
	tempFile, err := os.CreateTemp("", "config.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // 清理临时文件

	// 写入示例配置内容
	configContent := `
server:
  host: "localhost"
  port: 8080
cors:
  allow_origins: ["*"]
mysql:
  user: "root"
  password: "password"
`
	if _, err := tempFile.Write([]byte(configContent)); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	// 关闭文件以便 Viper 可以读取
	if err := tempFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	// 加载配置
	config, err := LoadConfig(tempFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestGetConfigByModuleName(t *testing.T) {
	// 创建一个示例配置
	testConfig := &Config{
		AliyunSms: []sms.AliyunSms{
			{ModuleName: "sms_module_1"},
			{ModuleName: "sms_module_2"},
		},
		MySQL: []database.MySQL{
			{ModuleName: "mysql_module_1"},
			{ModuleName: "mysql_module_2"},
		},
	}

	// 测试获取指定模块名称的配置
	resultConfig, err := GetConfigByModuleName(testConfig, "mysql_module_1")
	assert.NoError(t, err)
	assert.Len(t, resultConfig.MySQL, 1) // 确保返回的 MySQL 切片长度为 1
	assert.Equal(t, "mysql_module_1", resultConfig.MySQL[0].ModuleName) // 验证模块名称

	// 测试获取不存在的模块名称
	_, err = GetConfigByModuleName(testConfig, "non_existent_module")
	assert.Error(t, err) // 应该返回错误
}

// 测试 GetModuleByName 函数
func TestGetModuleByName(t *testing.T) {
	tests := []struct {
		name       string
		modules    []captcha.Captcha
		moduleName string
		expected   captcha.Captcha
		expectErr  bool
	}{
		{
			name:       "Empty slice",
			modules:    []captcha.Captcha{},
			moduleName: "captcha1",
			expectErr:  true,
		},
		{
			name:       "Single module found",
			modules:    []captcha.Captcha{{ModuleName: "captcha1", KeyLen: 6}},
			moduleName: "captcha1",
			expected:   captcha.Captcha{ModuleName: "captcha1", KeyLen: 6},
			expectErr:  false,
		},
		{
			name:       "Single module not found",
			modules:    []captcha.Captcha{{ModuleName: "captcha1", KeyLen: 6}},
			moduleName: "captcha2",
			expectErr:  true,
		},
		{
			name: "Multiple modules, one found",
			modules: []captcha.Captcha{
				{ModuleName: "captcha1", KeyLen: 6},
				{ModuleName: "captcha2", KeyLen: 8},
			},
			moduleName: "captcha2",
			expected:   captcha.Captcha{ModuleName: "captcha2", KeyLen: 8},
			expectErr:  false,
		},
		{
			name: "Multiple modules, none found",
			modules: []captcha.Captcha{
				{ModuleName: "captcha1", KeyLen: 6},
				{ModuleName: "captcha2", KeyLen: 8},
			},
			moduleName: "captcha3",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetModuleByName(tt.modules, tt.moduleName)

			if tt.expectErr {
				if err == nil {
					t.Errorf("%s expected an error but got none", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("%s expected no error but got: %v", tt.name, err)
				}
				if result != tt.expected {
					t.Errorf("%s expected %v, but got %v", tt.name, tt.expected, result)
				}
			}
		})
	}
}
