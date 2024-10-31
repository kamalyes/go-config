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
		Addr:       fmt.Sprintf("%s:%d", random.RandString(5, random.CAPITAL), random.FRandInt(1024, 65535)), // 随机生成 IP 和端口
		Username:   random.RandString(8, random.CAPITAL),
		Password:   random.RandString(12, random.CAPITAL),
		Cwd:        random.RandString(10, random.CAPITAL),
	}
}

// 将 Ftp 的参数转换为 map
func ftpToMap(ftp *ftp.Ftp) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME": ftp.ModuleName,
		"ADDR":        ftp.Addr,
		"USERNAME":    ftp.Username,
		"PASSWORD":    ftp.Password,
		"CWD":         ftp.Cwd,
	}
}

// 验证 Ftp 的字段与期望的映射是否相等
func assertFtpFields(t *testing.T, ftp *ftp.Ftp, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], ftp.ModuleName)
	assert.Equal(t, expected["ADDR"], ftp.Addr)
	assert.Equal(t, expected["USERNAME"], ftp.Username)
	assert.Equal(t, expected["PASSWORD"], ftp.Password)
	assert.Equal(t, expected["CWD"], ftp.Cwd)
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
	assert.Equal(t, newParams.Addr, ftpInstance.Addr)
	assert.Equal(t, newParams.Username, ftpInstance.Username)
	assert.Equal(t, newParams.Password, ftpInstance.Password)
	assert.Equal(t, newParams.Cwd, ftpInstance.Cwd)
}
