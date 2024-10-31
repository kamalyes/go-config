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
