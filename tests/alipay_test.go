/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:39:54
 * @FilePath: \go-config\tests\alipay_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/pay"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Alipay 配置参数
func generateAlipayTestParams() *pay.Alipay {
	return &pay.Alipay{
		ModuleName: random.RandString(10, random.CAPITAL),
		Pid:        random.RandString(10, random.CAPITAL),
		AppId:      random.RandString(10, random.CAPITAL),
		PriKey:     random.RandString(10, random.CAPITAL),
		PubKey:     random.RandString(10, random.CAPITAL),
		SignType:   random.RandString(5, random.CAPITAL),
		NotifyUrl:  "http://notify-url/" + random.RandString(5, random.CAPITAL),
		Subject:    random.RandString(10, random.CAPITAL),
	}
}

// 将 Alipay 的参数转换为 map
func alipayToMap(alipay *pay.Alipay) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME": alipay.ModuleName,
		"PID":         alipay.Pid,
		"APP_ID":      alipay.AppId,
		"PRI_KEY":     alipay.PriKey,
		"PUB_KEY":     alipay.PubKey,
		"SIGN_TYPE":   alipay.SignType,
		"NOTIFY_URL":  alipay.NotifyUrl,
		"SUBJECT":     alipay.Subject,
	}
}

// 验证 Alipay 的字段与期望的映射是否相等
func assertAlipayFields(t *testing.T, alipay *pay.Alipay, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], alipay.ModuleName)
	assert.Equal(t, expected["PID"], alipay.Pid)
	assert.Equal(t, expected["APP_ID"], alipay.AppId)
	assert.Equal(t, expected["PRI_KEY"], alipay.PriKey)
	assert.Equal(t, expected["PUB_KEY"], alipay.PubKey)
	assert.Equal(t, expected["SIGN_TYPE"], alipay.SignType)
	assert.Equal(t, expected["NOTIFY_URL"], alipay.NotifyUrl)
	assert.Equal(t, expected["SUBJECT"], alipay.Subject)
}

func TestAlipayClone(t *testing.T) {
	params := generateAlipayTestParams()
	original := pay.NewAlipay(params)
	cloned := original.Clone().(*pay.Alipay)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestAlipaySet(t *testing.T) {
	oldParams := generateAlipayTestParams()
	newParams := generateAlipayTestParams()

	alipay := pay.NewAlipay(oldParams)
	newConfig := pay.NewAlipay(newParams)

	alipay.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, alipay.ModuleName)
	assert.Equal(t, newParams.Pid, alipay.Pid)
	assert.Equal(t, newParams.AppId, alipay.AppId)
	assert.Equal(t, newParams.PriKey, alipay.PriKey)
	assert.Equal(t, newParams.PubKey, alipay.PubKey)
	assert.Equal(t, newParams.SignType, alipay.SignType)
	assert.Equal(t, newParams.NotifyUrl, alipay.NotifyUrl)
	assert.Equal(t, newParams.Subject, alipay.Subject)
}
