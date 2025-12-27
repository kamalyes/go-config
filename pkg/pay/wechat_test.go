/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 11:30:00
 * @FilePath: \go-config\pkg\pay\wechat_test.go
 * @Description: 微信支付配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package pay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWechatPay_Clone(t *testing.T) {
	original := &WechatPay{
		ModuleName:  "test-wechat",
		AppId:       "test-appid",
		MchId:       "test-mchid",
		NotifyUrl:   "https://example.com/notify",
		ApiKey:      "test-apikey",
		CertP12Path: "/path/to/cert.p12",
	}

	cloned := original.Clone().(*WechatPay)

	// 验证克隆后的值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.AppId, cloned.AppId)
	assert.Equal(t, original.MchId, cloned.MchId)
	assert.Equal(t, original.NotifyUrl, cloned.NotifyUrl)
	assert.Equal(t, original.ApiKey, cloned.ApiKey)
	assert.Equal(t, original.CertP12Path, cloned.CertP12Path)

	// 修改原始对象不应影响克隆对象
	original.AppId = "new-appid"
	original.MchId = "new-mchid"
	assert.NotEqual(t, original.AppId, cloned.AppId)
	assert.NotEqual(t, original.MchId, cloned.MchId)
}
