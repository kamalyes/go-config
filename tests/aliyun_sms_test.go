/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 13:10:06
 * @FilePath: \go-config\tests\aliyun_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/sms"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 AliyunSms 配置参数
func generateAliyunSmsTestParams() *sms.AliyunSms {
	return &sms.AliyunSms{
		ModuleName:           random.RandString(10, random.CAPITAL),
		SecretID:             random.RandString(10, random.CAPITAL),
		SecretKey:            random.RandString(10, random.CAPITAL),
		Sign:                 random.RandString(10, random.CAPITAL),
		ResourceOwnerAccount: random.RandString(10, random.CAPITAL),
		ResourceOwnerID:      int64(random.FRandInt(1, 10000)), // 随机生成大于0的整数
		TemplateCodeVerify:   random.RandString(10, random.CAPITAL),
		Endpoint:             "http://api.aliyun.com/" + random.RandString(5, random.CAPITAL),
	}
}

func TestAliyunSmsClone(t *testing.T) {
	params := generateAliyunSmsTestParams()
	original := sms.NewAliyunSms(params)
	cloned := original.Clone().(*sms.AliyunSms)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestAliyunSmsSet(t *testing.T) {
	oldParams := generateAliyunSmsTestParams()
	newParams := generateAliyunSmsTestParams()

	aliyunSms := sms.NewAliyunSms(oldParams)
	newConfig := sms.NewAliyunSms(newParams)

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
