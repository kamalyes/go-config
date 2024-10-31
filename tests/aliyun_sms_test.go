/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 20:34:17
 * @FilePath: \go-config\tests\aliyun_sms_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/sms"
	"github.com/stretchr/testify/assert"
)

// generateAliyunSmsTestParams 生成测试参数
func generateAliyunSmsTestParams() sms.AliyunSms {
	return sms.AliyunSms{
		ModuleName:           "test-module",
		SecretID:             "secret-id",
		SecretKey:            "secret-key",
		Sign:                 "test-sign",
		ResourceOwnerAccount: "owner-account",
		ResourceOwnerID:      123456,
		TemplateCodeVerify:   "template-code",
		Endpoint:             "http://endpoint",
	}
}

func TestNewAliyunSms(t *testing.T) {
	params := generateAliyunSmsTestParams()
	aliyunSms := sms.NewAliyunSms(params.ModuleName, params.SecretID, params.SecretKey, params.Sign, params.ResourceOwnerAccount, params.ResourceOwnerID, params.TemplateCodeVerify, params.Endpoint)
	assert.NotNil(t, aliyunSms)
	assert.Equal(t, params.ModuleName, aliyunSms.ModuleName)
	assert.Equal(t, params.SecretID, aliyunSms.SecretID)
	assert.Equal(t, params.SecretKey, aliyunSms.SecretKey)
	assert.Equal(t, params.Sign, aliyunSms.Sign)
	assert.Equal(t, params.ResourceOwnerAccount, aliyunSms.ResourceOwnerAccount)
	assert.Equal(t, params.ResourceOwnerID, aliyunSms.ResourceOwnerID)
	assert.Equal(t, params.TemplateCodeVerify, aliyunSms.TemplateCodeVerify)
	assert.Equal(t, params.Endpoint, aliyunSms.Endpoint)
}

func TestAliyunSmsValidation(t *testing.T) {
	validParams := generateAliyunSmsTestParams()
	validSms := sms.NewAliyunSms(validParams.ModuleName, validParams.SecretID, validParams.SecretKey, validParams.Sign, validParams.ResourceOwnerAccount, validParams.ResourceOwnerID, validParams.TemplateCodeVerify, validParams.Endpoint)
	assert.NoError(t, validSms.Validate())

	invalidSms := sms.NewAliyunSms("", "", "", "", "", -1, "", "")
	assert.Error(t, invalidSms.Validate())
	assert.EqualError(t, invalidSms.Validate(), "module name cannot be empty")
}

func TestAliyunSmsClone(t *testing.T) {
	params := generateAliyunSmsTestParams()
	original := sms.NewAliyunSms(params.ModuleName, params.SecretID, params.SecretKey, params.Sign, params.ResourceOwnerAccount, params.ResourceOwnerID, params.TemplateCodeVerify, params.Endpoint)
	cloned := original.Clone().(*sms.AliyunSms)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestAliyunSmsToMap(t *testing.T) {
	params := generateAliyunSmsTestParams()
	aliyunSms := sms.NewAliyunSms(params.ModuleName, params.SecretID, params.SecretKey, params.Sign, params.ResourceOwnerAccount, params.ResourceOwnerID, params.TemplateCodeVerify, params.Endpoint)
	expectedMap := map[string]interface{}{
		"moduleName":           params.ModuleName,
		"secretId":             params.SecretID,
		"secretKey":            params.SecretKey,
		"sign":                 params.Sign,
		"resourceOwnerAccount": params.ResourceOwnerAccount,
		"resourceOwnerId":      params.ResourceOwnerID,
		"templateCodeVerify":   params.TemplateCodeVerify,
		"endpoint":             params.Endpoint,
	}
	assert.Equal(t, expectedMap, aliyunSms.ToMap())
}

func TestAliyunSmsFromMap(t *testing.T) {
	aliyunSms := sms.NewAliyunSms("", "", "", "", "", 0, "", "")
	data := map[string]interface{}{
		"moduleName":           "new-module",
		"secretId":             "new-secret-id",
		"secretKey":            "new-secret-key",
		"sign":                 "new-sign",
		"resourceOwnerAccount": "new-owner-account",
		"resourceOwnerId":      int64(654321),
		"templateCodeVerify":   "new-template-code",
		"endpoint":             "http://new-endpoint",
	}
	aliyunSms.FromMap(data)

	// 验证填充后的数据是否正确
	assert.Equal(t, "new-module", aliyunSms.ModuleName)
	assert.Equal(t, "new-secret-id", aliyunSms.SecretID)
	assert.Equal(t, "new-secret-key", aliyunSms.SecretKey)
	assert.Equal(t, "new-sign", aliyunSms.Sign)
	assert.Equal(t, "new-owner-account", aliyunSms.ResourceOwnerAccount)
	assert.Equal(t, int64(654321), aliyunSms.ResourceOwnerID)
	assert.Equal(t, "new-template-code", aliyunSms.TemplateCodeVerify)
	assert.Equal(t, "http://new-endpoint", aliyunSms.Endpoint)
}

func TestAliyunSmsSet(t *testing.T) {
	oldParams := sms.AliyunSms{
		ModuleName:           "old-module",
		SecretID:             "old-secret-id",
		SecretKey:            "old-secret-key",
		Sign:                 "old-sign",
		ResourceOwnerAccount: "old-owner-account",
		ResourceOwnerID:      1,
		TemplateCodeVerify:   "old-template-code",
		Endpoint:             "http://old-endpoint",
	}
	newParams := sms.AliyunSms{
		ModuleName:           "new-module",
		SecretID:             "new-secret-id",
		SecretKey:            "new-secret-key",
		Sign:                 "new-sign",
		ResourceOwnerAccount: "new-owner-account",
		ResourceOwnerID:      2,
		TemplateCodeVerify:   "new-template-code",
		Endpoint:             "http://new-endpoint",
	}

	aliyunSms := sms.NewAliyunSms(oldParams.ModuleName, oldParams.SecretID, oldParams.SecretKey, oldParams.Sign, oldParams.ResourceOwnerAccount, oldParams.ResourceOwnerID, oldParams.TemplateCodeVerify, oldParams.Endpoint)
	newConfig := sms.NewAliyunSms(newParams.ModuleName, newParams.SecretID, newParams.SecretKey, newParams.Sign, newParams.ResourceOwnerAccount, newParams.ResourceOwnerID, newParams.TemplateCodeVerify, newParams.Endpoint)

	aliyunSms.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, aliyunSms.ModuleName)
	assert.Equal(t, newParams.SecretID, aliyunSms.SecretID)
	assert.Equal(t, newParams.SecretKey, aliyunSms.SecretKey)
	assert.Equal(t, newParams.Sign, aliyunSms.Sign)
	assert.Equal(t, newParams.ResourceOwnerAccount, aliyunSms.ResourceOwnerAccount)
	assert.Equal(t, newParams.ResourceOwnerID, aliyunSms.ResourceOwnerID)
	assert.Equal(t, newParams.TemplateCodeVerify, aliyunSms.TemplateCodeVerify)
	assert.Equal(t, newParams.Endpoint, aliyunSms.Endpoint)
}
