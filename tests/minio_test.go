/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 11:35:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:50:07
 * @FilePath: \go-config\oss\minio_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/oss"
	"github.com/stretchr/testify/assert"
)

// getValidMinioConfig 返回有效的 Minio 配置
func getValidMinioConfig() *oss.Minio {
	return oss.NewMinio("test-module", "localhost", 9000, "minioadmin", "minioadmin")
}

// getInvalidMinioConfig 返回无效的 Minio 配置
func getInvalidMinioConfig() *oss.Minio {
	return oss.NewMinio("", "", -1, "", "")
}

func TestNewMinio(t *testing.T) {
	// 测试创建 Minio 实例
	minio := getValidMinioConfig()

	assert.NotNil(t, minio)
	assert.Equal(t, "test-module", minio.ModuleName)
	assert.Equal(t, "localhost", minio.Host)
	assert.Equal(t, 9000, minio.Port)
	assert.Equal(t, "minioadmin", minio.AccessKey)
	assert.Equal(t, "minioadmin", minio.SecretKey)
}

func TestMinio_Validate(t *testing.T) {
	// 测试有效配置
	validMinio := getValidMinioConfig()
	assert.NoError(t, validMinio.Validate())

	// 测试无效配置
	invalidMinio := getInvalidMinioConfig()
	assert.Error(t, invalidMinio.Validate())
}

func TestMinio_ToMap(t *testing.T) {
	// 测试将配置转换为映射
	m := getValidMinioConfig()
	mMap := m.ToMap()

	assert.Equal(t, "test-module", mMap["moduleName"])
	assert.Equal(t, "localhost", mMap["host"])
	assert.Equal(t, 9000, mMap["port"])
	assert.Equal(t, "minioadmin", mMap["accessKey"])
	assert.Equal(t, "minioadmin", mMap["secretKey"])
}

func TestMinio_FromMap(t *testing.T) {
	// 测试从映射中填充配置
	m := &oss.Minio{}
	data := map[string]interface{}{
		"moduleName": "test-module",
		"host":       "localhost",
		"port":       9000,
		"accessKey":  "minioadmin",
		"secretKey":  "minioadmin",
	}
	m.FromMap(data)

	assert.Equal(t, "test-module", m.ModuleName)
	assert.Equal(t, "localhost", m.Host)
	assert.Equal(t, 9000, m.Port)
	assert.Equal(t, "minioadmin", m.AccessKey)
	assert.Equal(t, "minioadmin", m.SecretKey)
}
