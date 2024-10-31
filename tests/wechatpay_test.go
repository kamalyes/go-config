/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:57:58
 * @FilePath: \go-config\tests\wechatpay_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/pay"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Wechat 配置参数
func generateWechatTestParams() *pay.Wechat {
	return &pay.Wechat{
		ModuleName:  random.RandString(10, random.CAPITAL),
		AppId:       random.RandString(18, random.CAPITAL),                                                // 随机生成应用 ID
		MchId:       random.RandString(10, random.NUMBER),                                                 // 随机生成商户号
		NotifyUrl:   fmt.Sprintf("https://%s.example.com/callback", random.RandString(5, random.CAPITAL)), // 随机生成回调 URL
		ApiKey:      random.RandString(32, random.CAPITAL),                                                // 随机生成签名用的 key
		CertP12Path: fmt.Sprintf("/path/to/cert/%s.p12", random.RandString(5, random.CAPITAL)),            // 随机生成 P12 密钥文件路径
	}
}

// 将 Wechat 的参数转换为 map
func wechatToMap(wechat *pay.Wechat) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME":   wechat.ModuleName,
		"APP_ID":        wechat.AppId,
		"MCH_ID":        wechat.MchId,
		"NOTIFY_URL":    wechat.NotifyUrl,
		"API_KEY":       wechat.ApiKey,
		"CERT_P12_PATH": wechat.CertP12Path,
	}
}

// 验证 Wechat 的字段与期望的映射是否相等
func assertWechatFields(t *testing.T, wechat *pay.Wechat, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], wechat.ModuleName)
	assert.Equal(t, expected["APP_ID"], wechat.AppId)
	assert.Equal(t, expected["MCH_ID"], wechat.MchId)
	assert.Equal(t, expected["NOTIFY_URL"], wechat.NotifyUrl)
	assert.Equal(t, expected["API_KEY"], wechat.ApiKey)
	assert.Equal(t, expected["CERT_P12_PATH"], wechat.CertP12Path)
}

func TestWechatClone(t *testing.T) {
	params := generateWechatTestParams()
	original := pay.NewWechat(params)
	cloned := original.Clone().(*pay.Wechat)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestWechatSet(t *testing.T) {
	oldParams := generateWechatTestParams()
	newParams := generateWechatTestParams()

	wechatInstance := pay.NewWechat(oldParams)
	newConfig := pay.NewWechat(newParams)

	wechatInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, wechatInstance.ModuleName)
	assert.Equal(t, newParams.AppId, wechatInstance.AppId)
	assert.Equal(t, newParams.MchId, wechatInstance.MchId)
	assert.Equal(t, newParams.NotifyUrl, wechatInstance.NotifyUrl)
	assert.Equal(t, newParams.ApiKey, wechatInstance.ApiKey)
	assert.Equal(t, newParams.CertP12Path, wechatInstance.CertP12Path)
}
