/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:19:51
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-05 10:09:55
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
	"github.com/kamalyes/go-config/pkg/cors"
	"github.com/kamalyes/go-config/pkg/register"
)

func TestGetSingleConfigByModuleName(t *testing.T) {
	// 创建一个示例 MultiConfig
	multiConfig := goconfig.MultiConfig{
		Server: []register.Server{
			{ModuleName: "server1", Addr: "localhost:8080"},
			{ModuleName: "server2", Addr: "localhost:9000"},
		},
		Cors: []cors.Cors{
			{ModuleName: "cors1", AllowedOrigins: []string{"*"}},
			{ModuleName: "cors2", AllowedOrigins: []string{"http://example.com"}},
		},
		Consul: []register.Consul{
			{ModuleName: "consul1", Addr: "127.0.0.1:8080"},
			{ModuleName: "consul2", Addr: "127.0.0.1:7777"},
		},
		Captcha: []captcha.Captcha{
			{ModuleName: "captcha1", KeyLen: 55},
			{ModuleName: "captcha2", KeyLen: 66},
		},
	}

	tests := []struct {
		moduleName  string
		expected    goconfig.SingleConfig
		expectError bool
	}{
		// 正向测试用例
		{
			moduleName: "cors1",
			expected: goconfig.SingleConfig{
				Cors: cors.Cors{ModuleName: "cors1", AllowedOrigins: []string{"*"}},
			},
			expectError: false,
		},
		{
			moduleName: "server2",
			expected: goconfig.SingleConfig{
				Server: register.Server{ModuleName: "server2", Addr: "localhost:9000"},
			},
			expectError: false,
		},
		{
			moduleName: "consul1",
			expected: goconfig.SingleConfig{
				Consul: register.Consul{ModuleName: "consul1", Addr: "127.0.0.1:8080"},
			},
			expectError: false,
		},
		{
			moduleName: "captcha2",
			expected: goconfig.SingleConfig{
				Captcha: captcha.Captcha{ModuleName: "captcha2", KeyLen: 66},
			},
			expectError: false,
		},
		// 逆向测试用例
		{
			moduleName:  "nonexistent",
			expected:    goconfig.SingleConfig{},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.moduleName, func(t *testing.T) {
			singleConfig, err := goconfig.GetSingleConfigByModuleName(multiConfig, test.moduleName)

			if (err != nil) != test.expectError {
				t.Errorf("expected error: %v, got: %v", test.expectError, err != nil)
			}

			if !test.expectError {
				// 手动比较每个字段
				if !compareSingleConfig(*singleConfig, test.expected) {
					t.Errorf("expected: %+v, got: %+v", test.expected, singleConfig)
				}
			}
		})
	}
}

// compareSingleConfig 手动比较 SingleConfig 结构体
func compareSingleConfig(a, b goconfig.SingleConfig) bool {
	if a.Server.ModuleName != b.Server.ModuleName || a.Server.Addr != b.Server.Addr {
		return false
	}
	if a.Cors.ModuleName != b.Cors.ModuleName || !equalStringSlices(a.Cors.AllowedOrigins, b.Cors.AllowedOrigins) {
		return false
	}
	if a.Consul.ModuleName != b.Consul.ModuleName || a.Consul.Addr != b.Consul.Addr {
		return false
	}
	if a.Captcha.ModuleName != b.Captcha.ModuleName || a.Captcha.KeyLen != b.Captcha.KeyLen {
		return false
	}
	return true
}

// equalStringSlices 比较两个字符串切片
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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
