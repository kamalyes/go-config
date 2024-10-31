/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:52:22
 * @FilePath: \go-config\pay\wechat_test.go
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

// getValidWechatPayConfig 返回有效的 Wechat 配置
func getValidWechatPayConfig() *pay.Wechat {
	return pay.NewWechat("test-module", "app-123", "mch-123", "http://notify.url", "api-key", "/path/to/cert.p12")
}

// getInvalidWechatPayConfig 返回无效的 Wechat 配置
func getInvalidWechatPayConfig() *pay.Wechat {
	return pay.NewWechat("", "", "", "", "", "")
}

func TestNewWechatPay(t *testing.T) {
	// 测试创建 Wechat 实例
	validWechat := getValidWechatPayConfig()

	assert.NotNil(t, validWechat)                                 // 确保 validWechat 不为 nil
	assert.Equal(t, "test-module", validWechat.ModuleName)        // 检查模块名称
	assert.Equal(t, "app-123", validWechat.AppId)                 // 检查应用 ID
	assert.Equal(t, "mch-123", validWechat.MchId)                 // 检查商户 ID
	assert.Equal(t, "http://notify.url", validWechat.NotifyUrl)   // 检查通知 URL
	assert.Equal(t, "api-key", validWechat.ApiKey)                // 检查 API 密钥
	assert.Equal(t, "/path/to/cert.p12", validWechat.CertP12Path) // 检查证书路径
}

func TestWechatPay_Validate(t *testing.T) {
	// 测试有效配置
	validWechat := getValidWechatPayConfig()
	assert.NoError(t, validWechat.Validate()) // 确保没有错误

	// 测试无效配置
	invalidWechat := getInvalidWechatPayConfig()
	assert.Error(t, invalidWechat.Validate()) // 确保返回错误
}

func TestWechatPay_ToMap(t *testing.T) {
	// 测试将配置转换为映射
	wechatConfig := getValidWechatPayConfig()
	wechatMap := wechatConfig.ToMap()

	assert.Equal(t, "test-module", wechatMap["moduleName"]) // 检查模块名称
	assert.Equal(t, "app-123", wechatMap["appId"])          // 检查应用 ID
	assert.Equal(t, "mch-123", wechatMap["mchId"])          // 检查商户 ID
}

func TestWechatPayClone(t *testing.T) {
	originalWechat := getValidWechatPayConfig()          // 获取原始配置
	clonedWechat := originalWechat.Clone().(*pay.Wechat) // 克隆配置

	assert.Equal(t, originalWechat, clonedWechat)   // 确保原始和克隆的配置相等
	assert.NotSame(t, originalWechat, clonedWechat) // 确保它们是不同的实例
}
