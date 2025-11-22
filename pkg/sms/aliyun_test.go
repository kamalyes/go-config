/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\sms\aliyun_test.go
 * @Description: 阿里云短信配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package sms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunSms_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "aliyun-sms", config.ModuleName)
	assert.Equal(t, "sms_secret_id", config.SecretID)
	assert.Equal(t, "sms_secret_key", config.SecretKey)
	assert.Equal(t, "", config.Sign)
	assert.Equal(t, "", config.ResourceOwnerAccount)
	assert.Equal(t, int64(0), config.ResourceOwnerID)
	assert.Equal(t, "", config.TemplateCodeVerify)
	assert.Equal(t, "https://dysmsapi.aliyuncs.com", config.Endpoint)
}

func TestAliyunSms_WithModuleName(t *testing.T) {
	config := Default()
	result := config.WithModuleName("custom-sms")
	assert.Equal(t, "custom-sms", result.ModuleName)
	assert.Equal(t, config, result)
}

func TestAliyunSms_WithSecretID(t *testing.T) {
	config := Default()
	result := config.WithSecretID("test-secret-id")
	assert.Equal(t, "test-secret-id", result.SecretID)
	assert.Equal(t, config, result)
}

func TestAliyunSms_WithSecretKey(t *testing.T) {
	config := Default()
	result := config.WithSecretKey("test-secret-key")
	assert.Equal(t, "test-secret-key", result.SecretKey)
	assert.Equal(t, config, result)
}

func TestAliyunSms_WithSign(t *testing.T) {
	config := Default()
	result := config.WithSign("MyCompany")
	assert.Equal(t, "MyCompany", result.Sign)
	assert.Equal(t, config, result)
}

func TestAliyunSms_WithResourceOwnerAccount(t *testing.T) {
	config := Default()
	result := config.WithResourceOwnerAccount("owner@example.com")
	assert.Equal(t, "owner@example.com", result.ResourceOwnerAccount)
	assert.Equal(t, config, result)
}

func TestAliyunSms_WithResourceOwnerID(t *testing.T) {
	config := Default()
	result := config.WithResourceOwnerID(123456)
	assert.Equal(t, int64(123456), result.ResourceOwnerID)
	assert.Equal(t, config, result)
}

func TestAliyunSms_WithTemplateCodeVerify(t *testing.T) {
	config := Default()
	result := config.WithTemplateCodeVerify("SMS_123456")
	assert.Equal(t, "SMS_123456", result.TemplateCodeVerify)
	assert.Equal(t, config, result)
}

func TestAliyunSms_WithEndpoint(t *testing.T) {
	config := Default()
	result := config.WithEndpoint("https://custom-endpoint.com")
	assert.Equal(t, "https://custom-endpoint.com", result.Endpoint)
	assert.Equal(t, config, result)
}

func TestAliyunSms_Clone(t *testing.T) {
	config := Default()
	config.WithSecretID("id123").
		WithSecretKey("key123").
		WithSign("TestSign").
		WithResourceOwnerID(999)

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*AliyunSms)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.SecretID, clonedConfig.SecretID)
	assert.Equal(t, config.SecretKey, clonedConfig.SecretKey)
	assert.Equal(t, config.Sign, clonedConfig.Sign)
	assert.Equal(t, config.ResourceOwnerID, clonedConfig.ResourceOwnerID)

	// 验证深拷贝
	clonedConfig.SecretID = "modified"
	assert.NotEqual(t, config.SecretID, clonedConfig.SecretID)
}

func TestAliyunSms_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestAliyunSms_Set(t *testing.T) {
	config := Default()
	newConfig := &AliyunSms{
		ModuleName:           "new-sms",
		SecretID:             "new-id",
		SecretKey:            "new-key",
		Sign:                 "NewSign",
		ResourceOwnerAccount: "new-owner",
		ResourceOwnerID:      888,
		TemplateCodeVerify:   "SMS_888",
		Endpoint:             "https://new-endpoint.com",
	}

	config.Set(newConfig)
	assert.Equal(t, "new-sms", config.ModuleName)
	assert.Equal(t, "new-id", config.SecretID)
	assert.Equal(t, "new-key", config.SecretKey)
	assert.Equal(t, "NewSign", config.Sign)
	assert.Equal(t, "new-owner", config.ResourceOwnerAccount)
	assert.Equal(t, int64(888), config.ResourceOwnerID)
	assert.Equal(t, "SMS_888", config.TemplateCodeVerify)
	assert.Equal(t, "https://new-endpoint.com", config.Endpoint)
}

func TestAliyunSms_Validate(t *testing.T) {
	config := &AliyunSms{
		SecretID:             "test-id",
		SecretKey:            "test-key",
		Sign:                 "TestSign",
		ResourceOwnerAccount: "owner@example.com",
		ResourceOwnerID:      123456,
		TemplateCodeVerify:   "SMS_123456",
		Endpoint:             "https://dysmsapi.aliyuncs.com",
	}

	err := config.Validate()
	assert.NoError(t, err)
}

func TestAliyunSms_NewAliyunSms(t *testing.T) {
	opt := &AliyunSms{
		ModuleName: "test-sms",
		SecretID:   "test-id",
		Endpoint:   "https://test.com",
	}

	result := NewAliyunSms(opt)
	assert.NotNil(t, result)
	assert.Equal(t, "test-sms", result.ModuleName)
	assert.Equal(t, "test-id", result.SecretID)
	assert.Equal(t, "https://test.com", result.Endpoint)
}

func TestAliyunSms_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("my-sms").
		WithSecretID("my-secret-id").
		WithSecretKey("my-secret-key").
		WithSign("MyCompany").
		WithResourceOwnerAccount("admin@company.com").
		WithResourceOwnerID(123456789).
		WithTemplateCodeVerify("SMS_987654").
		WithEndpoint("https://custom.aliyuncs.com")

	assert.Equal(t, "my-sms", config.ModuleName)
	assert.Equal(t, "my-secret-id", config.SecretID)
	assert.Equal(t, "my-secret-key", config.SecretKey)
	assert.Equal(t, "MyCompany", config.Sign)
	assert.Equal(t, "admin@company.com", config.ResourceOwnerAccount)
	assert.Equal(t, int64(123456789), config.ResourceOwnerID)
	assert.Equal(t, "SMS_987654", config.TemplateCodeVerify)
	assert.Equal(t, "https://custom.aliyuncs.com", config.Endpoint)
}
