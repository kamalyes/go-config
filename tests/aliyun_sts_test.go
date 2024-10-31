/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 12:55:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 12:55:56
 * @FilePath: \go-config\sms\aliyun_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/sts"
	"github.com/stretchr/testify/assert"
)

// getValidAliyunStsConfig 返回有效的 AliyunSts 配置
func getValidAliyunStsConfig() *sts.AliyunSts {
	return sts.NewAliyunSts("test-module", "cn-hangzhou", "test-access-key-id", "test-access-key-secret", "test-role-arn", "test-role-session-name")
}

// getInvalidAliyunStsConfig 返回无效的 AliyunSts 配置
func getInvalidAliyunStsConfig() *sts.AliyunSts {
	return sts.NewAliyunSts("", "", "", "", "", "")
}

func TestNewAliyunSts(t *testing.T) {
	// 测试创建 AliyunSts 实例
	sts := getValidAliyunStsConfig()

	assert.NotNil(t, sts)
	assert.Equal(t, "test-module", sts.ModuleName)
	assert.Equal(t, "cn-hangzhou", sts.RegionID)
	assert.Equal(t, "test-access-key-id", sts.AccessKeyID)
	assert.Equal(t, "test-access-key-secret", sts.AccessKeySecret)
	assert.Equal(t, "test-role-arn", sts.RoleArn)
	assert.Equal(t, "test-role-session-name", sts.RoleSessionName)
}

func TestAliyunSts_Validate(t *testing.T) {
	// 测试有效配置
	validSts := getValidAliyunStsConfig()
	assert.NoError(t, validSts.Validate())

	// 测试无效配置
	invalidSts := getInvalidAliyunStsConfig()
	assert.Error(t, invalidSts.Validate())
}

func TestAliyunSts_ToMap(t *testing.T) {
	// 测试将配置转换为映射
	sts := getValidAliyunStsConfig()
	stsMap := sts.ToMap()

	assert.Equal(t, "test-module", stsMap["moduleName"])
	assert.Equal(t, "cn-hangzhou", stsMap["regionId"])
	assert.Equal(t, "test-access-key-id", stsMap["accessKeyId"])
	assert.Equal(t, "test-access-key-secret", stsMap["accessKeySecret"])
	assert.Equal(t, "test-role-arn", stsMap["roleArn"])
	assert.Equal(t, "test-role-session-name", stsMap["roleSessionName"])
}

func TestAliyunSts_FromMap(t *testing.T) {
	// 测试从映射填充配置
	data := map[string]interface{}{
		"moduleName":      "test-module",
		"regionId":        "cn-hangzhou",
		"accessKeyId":     "test-access-key-id",
		"accessKeySecret": "test-access-key-secret",
		"roleArn":         "test-role-arn",
		"roleSessionName": "test-role-session-name",
	}

	sts := &sts.AliyunSts{}
	sts.FromMap(data)

	assert.Equal(t, "test-module", sts.ModuleName)
	assert.Equal(t, "cn-hangzhou", sts.RegionID)
	assert.Equal(t, "test-access-key-id", sts.AccessKeyID)
	assert.Equal(t, "test-access-key-secret", sts.AccessKeySecret)
	assert.Equal(t, "test-role-arn", sts.RoleArn)
	assert.Equal(t, "test-role-session-name", sts.RoleSessionName)
}
