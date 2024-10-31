/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:19:51
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 13:06:18
 * @FilePath: \go-config\tests\config_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/captcha"
	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-config/pkg/sms"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigByModuleName(t *testing.T) {
	// 创建一个示例配置
	testConfig := &goconfig.Config{
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
	resultConfig, err := goconfig.GetConfigByModuleName(testConfig, "mysql_module_1")
	assert.NoError(t, err)
	assert.Len(t, resultConfig.MySQL, 1)                                // 确保返回的 MySQL 切片长度为 1
	assert.Equal(t, "mysql_module_1", resultConfig.MySQL[0].ModuleName) // 验证模块名称

	// 测试获取不存在的模块名称
	_, err = goconfig.GetConfigByModuleName(testConfig, "non_existent_module")
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
			result, err := goconfig.GetModuleByName(tt.modules, tt.moduleName)

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
