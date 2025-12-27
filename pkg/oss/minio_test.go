/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 11:30:00
 * @FilePath: \go-config\pkg\oss\minio_test.go
 * @Description: Minio OSS配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package oss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinio_Clone(t *testing.T) {
	original := &Minio{
		ModuleName: "test-minio",
		Endpoint:   "localhost:9000",
		AccessKey:  "test-access",
		SecretKey:  "test-secret",
		Bucket:     "test-bucket",
		UseSSL:     true,
	}

	cloned := original.Clone().(*Minio)

	// 验证克隆后的值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Endpoint, cloned.Endpoint)
	assert.Equal(t, original.AccessKey, cloned.AccessKey)
	assert.Equal(t, original.SecretKey, cloned.SecretKey)
	assert.Equal(t, original.Bucket, cloned.Bucket)
	assert.Equal(t, original.UseSSL, cloned.UseSSL)

	// 修改原始对象不应影响克隆对象
	original.Endpoint = "newhost:9000"
	original.Bucket = "new-bucket"
	assert.NotEqual(t, original.Endpoint, cloned.Endpoint)
	assert.NotEqual(t, original.Bucket, cloned.Bucket)
}
