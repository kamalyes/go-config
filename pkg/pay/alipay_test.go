/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 11:30:00
 * @FilePath: \go-config\pkg\pay\alipay_test.go
 * @Description: 支付宝支付配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package pay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAliPay_Clone(t *testing.T) {
	original := &AliPay{
		ModuleName: "test-alipay",
		Pid:        "test-pid",
		AppId:      "test-appid",
		PriKey:     "test-prikey",
		PubKey:     "test-pubkey",
		SignType:   "RSA2",
		NotifyUrl:  "https://example.com/notify",
		Subject:    "测试订单",
	}

	cloned := original.Clone().(*AliPay)

	// 验证克隆后的值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Pid, cloned.Pid)
	assert.Equal(t, original.AppId, cloned.AppId)
	assert.Equal(t, original.PriKey, cloned.PriKey)
	assert.Equal(t, original.PubKey, cloned.PubKey)
	assert.Equal(t, original.SignType, cloned.SignType)
	assert.Equal(t, original.NotifyUrl, cloned.NotifyUrl)
	assert.Equal(t, original.Subject, cloned.Subject)

	// 修改原始对象不应影响克隆对象
	original.AppId = "new-appid"
	original.Subject = "新订单"
	assert.NotEqual(t, original.AppId, cloned.AppId)
	assert.NotEqual(t, original.Subject, cloned.Subject)
}
