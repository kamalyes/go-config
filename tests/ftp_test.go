/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:55:26
 * @FilePath: \go-config\tests\ftp_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/ftp"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 FTP 配置参数
func generateFtpTestParams() *ftp.Ftp {
	return &ftp.Ftp{
		ModuleName: random.RandString(10, random.CAPITAL),
		Endpoint:   fmt.Sprintf("%s:%d", random.RandString(5, random.CAPITAL), random.FRandInt(1024, 65535)), // 随机生成 IP 和端口
		Username:   random.RandString(8, random.CAPITAL),
		Password:   random.RandString(12, random.CAPITAL),
		Cwd:        random.RandString(10, random.CAPITAL),
	}
}

func TestFtpClone(t *testing.T) {
	params := generateFtpTestParams()
	original := ftp.NewFtp(params)
	cloned := original.Clone().(*ftp.Ftp)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestFtpSet(t *testing.T) {
	oldParams := generateFtpTestParams()
	newParams := generateFtpTestParams()

	ftpInstance := ftp.NewFtp(oldParams)
	newConfig := ftp.NewFtp(newParams)

	ftpInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, ftpInstance.ModuleName)
	assert.Equal(t, newParams.Endpoint, ftpInstance.Endpoint)
	assert.Equal(t, newParams.Username, ftpInstance.Username)
	assert.Equal(t, newParams.Password, ftpInstance.Password)
	assert.Equal(t, newParams.Cwd, ftpInstance.Cwd)
}

// TestFtpDefault 测试默认配置
func TestFtpDefault(t *testing.T) {
	defaultConfig := ftp.DefaultFtp()
	
	// 检查默认值
	assert.Equal(t, "ftp", defaultConfig.ModuleName)
	assert.Equal(t, "127.0.0.1:21", defaultConfig.Endpoint)
	assert.Equal(t, "", defaultConfig.Username)
	assert.Equal(t, "", defaultConfig.Password)
	assert.Equal(t, "/", defaultConfig.Cwd)
}

// TestFtpDefaultPointer 测试默认配置指针
func TestFtpDefaultPointer(t *testing.T) {
	config := ftp.Default()
	
	assert.NotNil(t, config)
	assert.Equal(t, "ftp", config.ModuleName)
	assert.Equal(t, "127.0.0.1:21", config.Endpoint)
}

// TestFtpChainMethods 测试链式方法
func TestFtpChainMethods(t *testing.T) {
	config := ftp.Default().
		WithModuleName("file-service").
		WithEndpoint("192.168.1.200:21").
		WithUsername("testuser").
		WithPassword("testpass").
		WithCwd("/uploads")
	
	assert.Equal(t, "file-service", config.ModuleName)
	assert.Equal(t, "192.168.1.200:21", config.Endpoint)
	assert.Equal(t, "testuser", config.Username)
	assert.Equal(t, "testpass", config.Password)
	assert.Equal(t, "/uploads", config.Cwd)
}

// TestFtpChainMethodsReturnPointer 测试链式方法返回指针
func TestFtpChainMethodsReturnPointer(t *testing.T) {
	config1 := ftp.Default()
	config2 := config1.WithEndpoint("localhost:21")
	
	// 应该返回同一个实例
	assert.Same(t, config1, config2)
	assert.Equal(t, "localhost:21", config1.Endpoint)
}
