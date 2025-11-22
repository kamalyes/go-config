/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\sts\aliyun_test.go
 * @Description: 阿里云STS配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package sts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunSts_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "aliyun-sts", config.ModuleName)
	assert.Equal(t, "cn-hangzhou", config.RegionID)
	assert.Equal(t, "demo_access_key_id", config.AccessKeyID)
	assert.Equal(t, "demo_access_key_secret", config.AccessKeySecret)
	assert.Equal(t, "acs:ram::account:role/demo-role", config.RoleArn)
	assert.Equal(t, "default-session", config.RoleSessionName)
}

func TestAliyunSts_DefaultAliyunSts(t *testing.T) {
	config := DefaultAliyunSts()
	assert.Equal(t, "aliyun-sts", config.ModuleName)
	assert.Equal(t, "cn-hangzhou", config.RegionID)
	assert.Equal(t, "demo_access_key_id", config.AccessKeyID)
	assert.Equal(t, "demo_access_key_secret", config.AccessKeySecret)
	assert.Equal(t, "acs:ram::account:role/demo-role", config.RoleArn)
	assert.Equal(t, "default-session", config.RoleSessionName)
}

func TestAliyunSts_DefaultAliyunStsConfig(t *testing.T) {
	config := DefaultAliyunStsConfig()
	assert.NotNil(t, config)
	assert.Equal(t, "aliyun-sts", config.ModuleName)
	assert.Equal(t, "cn-hangzhou", config.RegionID)
}

func TestAliyunSts_WithModuleName(t *testing.T) {
	config := Default()
	result := config.WithModuleName("custom-sts")
	assert.Equal(t, "custom-sts", result.ModuleName)
	assert.Equal(t, config, result)
}

func TestAliyunSts_WithRegionID(t *testing.T) {
	config := Default()
	result := config.WithRegionID("cn-beijing")
	assert.Equal(t, "cn-beijing", result.RegionID)
	assert.Equal(t, config, result)
}

func TestAliyunSts_WithAccessKeyID(t *testing.T) {
	config := Default()
	result := config.WithAccessKeyID("test-access-key-id")
	assert.Equal(t, "test-access-key-id", result.AccessKeyID)
	assert.Equal(t, config, result)
}

func TestAliyunSts_WithAccessKeySecret(t *testing.T) {
	config := Default()
	result := config.WithAccessKeySecret("test-access-key-secret")
	assert.Equal(t, "test-access-key-secret", result.AccessKeySecret)
	assert.Equal(t, config, result)
}

func TestAliyunSts_WithRoleArn(t *testing.T) {
	config := Default()
	result := config.WithRoleArn("acs:ram::123456:role/test-role")
	assert.Equal(t, "acs:ram::123456:role/test-role", result.RoleArn)
	assert.Equal(t, config, result)
}

func TestAliyunSts_WithRoleSessionName(t *testing.T) {
	config := Default()
	result := config.WithRoleSessionName("custom-session")
	assert.Equal(t, "custom-session", result.RoleSessionName)
	assert.Equal(t, config, result)
}

func TestAliyunSts_Clone(t *testing.T) {
	config := Default()
	config.WithRegionID("cn-shanghai").
		WithAccessKeyID("key123").
		WithAccessKeySecret("secret123").
		WithRoleArn("acs:ram::123456:role/test").
		WithRoleSessionName("test-session")

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*AliyunSts)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.RegionID, clonedConfig.RegionID)
	assert.Equal(t, config.AccessKeyID, clonedConfig.AccessKeyID)
	assert.Equal(t, config.AccessKeySecret, clonedConfig.AccessKeySecret)
	assert.Equal(t, config.RoleArn, clonedConfig.RoleArn)
	assert.Equal(t, config.RoleSessionName, clonedConfig.RoleSessionName)

	// 验证深拷贝
	clonedConfig.AccessKeyID = "modified"
	assert.NotEqual(t, config.AccessKeyID, clonedConfig.AccessKeyID)
}

func TestAliyunSts_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestAliyunSts_Set(t *testing.T) {
	config := Default()
	newConfig := &AliyunSts{
		ModuleName:      "new-sts",
		RegionID:        "cn-shenzhen",
		AccessKeyID:     "new-key-id",
		AccessKeySecret: "new-key-secret",
		RoleArn:         "acs:ram::999999:role/new-role",
		RoleSessionName: "new-session",
	}

	config.Set(newConfig)
	assert.Equal(t, "new-sts", config.ModuleName)
	assert.Equal(t, "cn-shenzhen", config.RegionID)
	assert.Equal(t, "new-key-id", config.AccessKeyID)
	assert.Equal(t, "new-key-secret", config.AccessKeySecret)
	assert.Equal(t, "acs:ram::999999:role/new-role", config.RoleArn)
	assert.Equal(t, "new-session", config.RoleSessionName)
}

func TestAliyunSts_Validate(t *testing.T) {
	config := &AliyunSts{
		ModuleName:      "test-sts",
		RegionID:        "cn-hangzhou",
		AccessKeyID:     "test-key-id",
		AccessKeySecret: "test-key-secret",
		RoleArn:         "acs:ram::123456:role/test-role",
		RoleSessionName: "test-session",
	}

	err := config.Validate()
	assert.NoError(t, err)
}

func TestAliyunSts_NewAliyunSts(t *testing.T) {
	opt := &AliyunSts{
		ModuleName:      "test-sts",
		RegionID:        "cn-beijing",
		AccessKeyID:     "test-id",
		AccessKeySecret: "test-secret",
	}

	result := NewAliyunSts(opt)
	assert.NotNil(t, result)
	assert.Equal(t, "test-sts", result.ModuleName)
	assert.Equal(t, "cn-beijing", result.RegionID)
	assert.Equal(t, "test-id", result.AccessKeyID)
	assert.Equal(t, "test-secret", result.AccessKeySecret)
}

func TestAliyunSts_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("my-sts").
		WithRegionID("cn-guangzhou").
		WithAccessKeyID("my-access-key-id").
		WithAccessKeySecret("my-access-key-secret").
		WithRoleArn("acs:ram::888888:role/my-role").
		WithRoleSessionName("my-session")

	assert.Equal(t, "my-sts", config.ModuleName)
	assert.Equal(t, "cn-guangzhou", config.RegionID)
	assert.Equal(t, "my-access-key-id", config.AccessKeyID)
	assert.Equal(t, "my-access-key-secret", config.AccessKeySecret)
	assert.Equal(t, "acs:ram::888888:role/my-role", config.RoleArn)
	assert.Equal(t, "my-session", config.RoleSessionName)
}
