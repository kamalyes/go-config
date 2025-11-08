/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 11:35:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:43:25
 * @FilePath: \go-config\tests\aliyun_oss_test.go
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

// 生成随机的 AliyunOss 配置参数
func generateAliyunOssTestParams() *oss.AliyunOss {
	return &oss.AliyunOss{
		ModuleName:          random.RandString(10, random.CAPITAL),
		ReplaceOriginalHost: random.RandString(5, random.CAPITAL) + ".AliyunOss.local", // 随机待替换主机名
		ReplaceLaterHost:    random.RandString(5, random.CAPITAL) + ".AliyunOss.local", // 随机生成替换后的
		AccessKey:           random.RandString(16, random.CAPITAL),                     // 随机生成 Access Key
		SecretKey:           random.RandString(32, random.CAPITAL),                     // 随机生成 Secret Key
		Endpoint:            random.RandString(32, random.CAPITAL),
		Bucket:              random.RandString(32, random.CAPITAL),
	}
}

func TestAliyunOssClone(t *testing.T) {
	params := generateAliyunOssTestParams()
	original := oss.NewAliyunOss(params)
	cloned := original.Clone().(*oss.AliyunOss)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestAliyunOssSet(t *testing.T) {
	oldParams := generateAliyunOssTestParams()
	newParams := generateAliyunOssTestParams()

	minioInstance := oss.NewAliyunOss(oldParams)
	newConfig := oss.NewAliyunOss(newParams)

	minioInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, minioInstance.ModuleName)
	assert.Equal(t, newParams.ReplaceOriginalHost, minioInstance.ReplaceOriginalHost)
	assert.Equal(t, newParams.ReplaceLaterHost, minioInstance.ReplaceLaterHost)
	assert.Equal(t, newParams.AccessKey, minioInstance.AccessKey)
	assert.Equal(t, newParams.SecretKey, minioInstance.SecretKey)
	assert.Equal(t, newParams.Endpoint, minioInstance.Endpoint)
	assert.Equal(t, newParams.Bucket, minioInstance.Bucket)

}
