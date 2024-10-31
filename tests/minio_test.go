/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 11:35:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:56:36
 * @FilePath: \go-config\tests\minio_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/oss"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Minio 配置参数
func generateMinioTestParams() *oss.Minio {
	return &oss.Minio{
		ModuleName: random.RandString(10, random.CAPITAL),
		Host:       random.RandString(5, random.CAPITAL) + ".minio.local", // 随机生成主机名
		Port:       random.FRandInt(9000, 9999),                           // 随机生成端口（9000到9999）
		AccessKey:  random.RandString(16, random.CAPITAL),                 // 随机生成 Access Key
		SecretKey:  random.RandString(32, random.CAPITAL),                 // 随机生成 Secret Key
	}
}

func TestMinioClone(t *testing.T) {
	params := generateMinioTestParams()
	original := oss.NewMinio(params)
	cloned := original.Clone().(*oss.Minio)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestMinioSet(t *testing.T) {
	oldParams := generateMinioTestParams()
	newParams := generateMinioTestParams()

	minioInstance := oss.NewMinio(oldParams)
	newConfig := oss.NewMinio(newParams)

	minioInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, minioInstance.ModuleName)
	assert.Equal(t, newParams.Host, minioInstance.Host)
	assert.Equal(t, newParams.Port, minioInstance.Port)
	assert.Equal(t, newParams.AccessKey, minioInstance.AccessKey)
	assert.Equal(t, newParams.SecretKey, minioInstance.SecretKey)
}
