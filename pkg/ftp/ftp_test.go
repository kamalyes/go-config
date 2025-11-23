/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 16:51:22
 * @FilePath: \go-config\pkg\ftp\ftp_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package ftp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFtp_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "ftp", config.ModuleName)
	assert.Equal(t, "127.0.0.1:21", config.Endpoint)
	assert.Equal(t, "ftp_user", config.Username)
	assert.Equal(t, "ftp_password", config.Password)
	assert.Equal(t, "/", config.Cwd)
}

func TestFtp_DefaultFtp(t *testing.T) {
	config := DefaultFtp()
	assert.Equal(t, "ftp", config.ModuleName)
	assert.Equal(t, "127.0.0.1:21", config.Endpoint)
	assert.Equal(t, "ftp_user", config.Username)
	assert.Equal(t, "ftp_password", config.Password)
	assert.Equal(t, "/", config.Cwd)
}

func TestFtp_WithModuleName(t *testing.T) {
	config := Default().WithModuleName("custom-ftp")
	assert.Equal(t, "custom-ftp", config.ModuleName)
}

func TestFtp_WithEndpoint(t *testing.T) {
	config := Default().WithEndpoint("192.168.1.100:2121")
	assert.Equal(t, "192.168.1.100:2121", config.Endpoint)
}

func TestFtp_WithUsername(t *testing.T) {
	config := Default().WithUsername("testuser")
	assert.Equal(t, "testuser", config.Username)
}

func TestFtp_WithPassword(t *testing.T) {
	config := Default().WithPassword("testpass123")
	assert.Equal(t, "testpass123", config.Password)
}

func TestFtp_WithCwd(t *testing.T) {
	config := Default().WithCwd("/uploads")
	assert.Equal(t, "/uploads", config.Cwd)
}

func TestFtp_Clone(t *testing.T) {
	original := Default().
		WithModuleName("test-ftp").
		WithEndpoint("ftp.example.com:21").
		WithUsername("admin").
		WithPassword("secret").
		WithCwd("/data")

	cloned := original.Clone().(*Ftp)

	// 验证值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Endpoint, cloned.Endpoint)
	assert.Equal(t, original.Username, cloned.Username)
	assert.Equal(t, original.Password, cloned.Password)
	assert.Equal(t, original.Cwd, cloned.Cwd)

	// 验证独立性
	cloned.ModuleName = "modified-ftp"
	cloned.Endpoint = "new.ftp.com:21"
	assert.NotEqual(t, original.ModuleName, cloned.ModuleName)
	assert.NotEqual(t, original.Endpoint, cloned.Endpoint)
}

func TestFtp_Get(t *testing.T) {
	config := Default().WithUsername("testuser")
	got := config.Get()
	assert.NotNil(t, got)
	ftpConfig, ok := got.(*Ftp)
	assert.True(t, ok)
	assert.Equal(t, "testuser", ftpConfig.Username)
}

func TestFtp_Set(t *testing.T) {
	config := Default()
	newConfig := &Ftp{
		ModuleName: "new-ftp",
		Endpoint:   "new.server.com:21",
		Username:   "newuser",
		Password:   "newpass",
		Cwd:        "/new",
	}

	config.Set(newConfig)
	assert.Equal(t, "new-ftp", config.ModuleName)
	assert.Equal(t, "new.server.com:21", config.Endpoint)
	assert.Equal(t, "newuser", config.Username)
	assert.Equal(t, "newpass", config.Password)
	assert.Equal(t, "/new", config.Cwd)
}

func TestFtp_Validate(t *testing.T) {
	// 有效配置
	validConfig := Default().
		WithEndpoint("ftp.example.com:21").
		WithUsername("user").
		WithPassword("pass")
	err := validConfig.Validate()
	assert.NoError(t, err)

	// 无效配置 - 缺少必需字段
	invalidConfig := &Ftp{
		ModuleName: "test",
		Endpoint:   "", // required field
		Username:   "user",
		Password:   "pass",
	}
	err = invalidConfig.Validate()
	assert.Error(t, err)
}

func TestFtp_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("chain-ftp").
		WithEndpoint("chained.ftp.com:21").
		WithUsername("chainuser").
		WithPassword("chainpass").
		WithCwd("/chain/path")

	assert.Equal(t, "chain-ftp", config.ModuleName)
	assert.Equal(t, "chained.ftp.com:21", config.Endpoint)
	assert.Equal(t, "chainuser", config.Username)
	assert.Equal(t, "chainpass", config.Password)
	assert.Equal(t, "/chain/path", config.Cwd)
}

func TestNewFtp(t *testing.T) {
	opt := &Ftp{
		ModuleName: "new-instance",
		Endpoint:   "new.ftp.com:21",
		Username:   "newuser",
		Password:   "newpass",
		Cwd:        "/new/path",
	}

	instance := NewFtp(opt)
	assert.NotNil(t, instance)
	assert.Equal(t, "new-instance", instance.ModuleName)
	assert.Equal(t, "new.ftp.com:21", instance.Endpoint)
	assert.Equal(t, "newuser", instance.Username)
	assert.Equal(t, "newpass", instance.Password)
	assert.Equal(t, "/new/path", instance.Cwd)
}
