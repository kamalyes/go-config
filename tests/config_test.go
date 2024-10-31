/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:19:51
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 21:26:37
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
	testConfig := &goconfig.MultiConfig{
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
		{"Empty slice", []captcha.Captcha{}, "captcha1", captcha.Captcha{}, true},
		{"Single module found", []captcha.Captcha{{ModuleName: "captcha1", KeyLen: 6}}, "captcha1", captcha.Captcha{ModuleName: "captcha1", KeyLen: 6}, false},
		{"Single module not found", []captcha.Captcha{{ModuleName: "captcha1", KeyLen: 6}}, "captcha2", captcha.Captcha{}, true},
		{"Multiple modules, one found", []captcha.Captcha{
			{ModuleName: "captcha1", KeyLen: 6},
			{ModuleName: "captcha2", KeyLen: 8},
		}, "captcha2", captcha.Captcha{ModuleName: "captcha2", KeyLen: 8}, false},
		{"Multiple modules, none found", []captcha.Captcha{
			{ModuleName: "captcha1", KeyLen: 6},
			{ModuleName: "captcha2", KeyLen: 8},
		}, "captcha3", captcha.Captcha{}, true},
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
