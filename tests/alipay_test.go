/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 11:35:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 12:05:57
 * @FilePath: \go-config\pay\alipay_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/pay"
	"github.com/stretchr/testify/assert"
)

// getValidAlipayPayConfig 返回有效的 AlipayPay 配置
func getValidAlipayPayConfig() *pay.Alipay {
	return pay.NewAlipay("test-module", "123456", "app-123", "private-key", "public-key", "RSA2", "http://notify.url", "Test Subject")
}

// getInvalidAlipayPayConfig 返回无效的 AlipayPay 配置
func getInvalidAlipayPayConfig() *pay.Alipay {
	return pay.NewAlipay("", "", "", "", "", "", "", "")
}

// TestAlipayCreation 测试 Alipay 配置的创建和验证
func TestAlipayCreation(t *testing.T) {
	validAlipay := getValidAlipayPayConfig()
	invalidAlipay := getInvalidAlipayPayConfig()

	// 测试有效的 Alipay 配置
	assert.NotNil(t, validAlipay) // 确保 alipay 对象不为 nil

	// 测试有效配置的验证
	assert.NoError(t, validAlipay.Validate()) // 确保没有错误

	// 测试无效配置
	assert.Error(t, invalidAlipay.Validate()) // 确保验证返回错误
}

// TestAlipayMapMethods 测试 Alipay 的映射方法
func TestAlipayMapMethods(t *testing.T) {
	validAlipay := getValidAlipayPayConfig()

	// 测试 ToMap 方法
	expectedMap := map[string]interface{}{
		"moduleName": "test-module",
		"pid":        "123456",
		"appId":      "app-123",
		"priKey":     "private-key",
		"pubKey":     "public-key",
		"signType":   "RSA2",
		"notifyUrl":  "http://notify.url",
		"subject":    "Test Subject",
	}
	assert.Equal(t, expectedMap, validAlipay.ToMap()) // 确保映射结果正确

	// 测试 FromMap 方法
	newAlipay := &pay.Alipay{}
	newAlipay.FromMap(validAlipay.ToMap())
	assert.Equal(t, validAlipay, newAlipay) // 确保从映射创建的对象与克隆对象相等
}

// TestAlipayPayCloneAndSetMethods 测试 Alipay 的克隆和设置方法
func TestAlipayPayCloneAndSetMethods(t *testing.T) {
	validAlipay := getValidAlipayPayConfig()

	// 测试 Set 方法
	newData := &pay.Alipay{
		ModuleName: "new-module",
		Pid:        "654321",
		AppId:      "app-456",
		PriKey:     "new-private-key",
		PubKey:     "new-public-key",
		SignType:   "RSA",
		NotifyUrl:  "http://new.notify.url",
		Subject:    "New Test Subject",
	}
	validAlipay.Set(newData)
	assert.Equal(t, newData, validAlipay) // 确保设置的新数据与克隆对象相等

	// 测试 Get 方法
	assert.Equal(t, validAlipay, validAlipay.Get()) // 确保获取的对象与克隆对象相等
}
