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
func generateWechatTestParams() *pay.WechatPay {
	return &pay.WechatPay{
		ModuleName:  random.RandString(10, random.CAPITAL),
		AppId:       random.RandString(18, random.CAPITAL),                                                // 随机生成应用 ID
		MchId:       random.RandString(10, random.NUMBER),                                                 // 随机生成商户号
		NotifyUrl:   fmt.Sprintf("https://%s.example.com/callback", random.RandString(5, random.CAPITAL)), // 随机生成回调 URL
		ApiKey:      random.RandString(32, random.CAPITAL),                                                // 随机生成签名用的 key
		CertP12Path: fmt.Sprintf("/path/to/cert/%s.p12", random.RandString(5, random.CAPITAL)),            // 随机生成 P12 密钥文件路径
	}
}

func TestWechatClone(t *testing.T) {
	params := generateWechatTestParams()
	original := pay.NewWechatPay(params)
	cloned := original.Clone().(*pay.WechatPay)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestWechatSet(t *testing.T) {
	oldParams := generateWechatTestParams()
	newParams := generateWechatTestParams()

	wechatInstance := pay.NewWechatPay(oldParams)
	newConfig := pay.NewWechatPay(newParams)

	wechatInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, wechatInstance.ModuleName)
	assert.Equal(t, newParams.AppId, wechatInstance.AppId)
	assert.Equal(t, newParams.MchId, wechatInstance.MchId)
	assert.Equal(t, newParams.NotifyUrl, wechatInstance.NotifyUrl)
	assert.Equal(t, newParams.ApiKey, wechatInstance.ApiKey)
	assert.Equal(t, newParams.CertP12Path, wechatInstance.CertP12Path)
}

// TestWechatPayDefault 测试默认配置
func TestWechatPayDefault(t *testing.T) {
	defaultConfig := pay.DefaultWechatPay()
	
	// 检查默认值
	assert.Equal(t, "wechatpay", defaultConfig.ModuleName)
	assert.Equal(t, "", defaultConfig.AppId)
	assert.Equal(t, "", defaultConfig.MchId)
	assert.Equal(t, "", defaultConfig.NotifyUrl)
	assert.Equal(t, "", defaultConfig.ApiKey)
	assert.Equal(t, "", defaultConfig.CertP12Path)
}

// TestWechatPayDefaultPointer 测试默认配置指针
func TestWechatPayDefaultPointer(t *testing.T) {
	config := pay.DefaultWechatPayConfig()
	
	assert.NotNil(t, config)
	assert.Equal(t, "wechatpay", config.ModuleName)
}

// TestWechatPayChainMethods 测试链式方法
func TestWechatPayChainMethods(t *testing.T) {
	config := pay.DefaultWechatPayConfig().
		WithModuleName("payment-service").
		WithAppID("wx1234567890123456").
		WithMchID("1234567890").
		WithNotifyURL("https://example.com/notify").
		WithApiKey("my-api-key").
		WithCertP12Path("./certs/apiclient_cert.p12")
	
	assert.Equal(t, "payment-service", config.ModuleName)
	assert.Equal(t, "wx1234567890123456", config.AppId)
	assert.Equal(t, "1234567890", config.MchId)
	assert.Equal(t, "https://example.com/notify", config.NotifyUrl)
	assert.Equal(t, "my-api-key", config.ApiKey)
	assert.Equal(t, "./certs/apiclient_cert.p12", config.CertP12Path)
}

// TestWechatPayChainMethodsReturnPointer 测试链式方法返回指针
func TestWechatPayChainMethodsReturnPointer(t *testing.T) {
	config1 := pay.DefaultWechatPayConfig()
	config2 := config1.WithAppID("new-app-id")
	
	// 应该返回同一个实例
	assert.Same(t, config1, config2)
	assert.Equal(t, "new-app-id", config1.AppId)
}
