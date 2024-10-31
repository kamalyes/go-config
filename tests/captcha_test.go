/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 16:56:26
 * @FilePath: \go-config\tests\captcha_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/captcha"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Captcha 配置参数
func generateCaptchaTestParams() *captcha.Captcha {
	return &captcha.Captcha{
		ModuleName: random.RandString(10, random.CAPITAL),
		KeyLen:     random.RandInt(4, 10),      // 随机生成 4 到 10 的整数
		ImgWidth:   random.RandInt(100, 300),   // 随机生成 100 到 300 的整数
		ImgHeight:  random.RandInt(30, 100),    // 随机生成 30 到 100 的整数
		MaxSkew:    random.RandFloat(0.5, 1.0), // 随机生成 0.5 到 1.0 的浮点数
		DotCount:   random.RandInt(80, 120),    // 随机生成 80 到 120 的整数
	}
}

func TestCaptchaClone(t *testing.T) {
	params := generateCaptchaTestParams()
	original := captcha.NewCaptcha(params)
	cloned := original.Clone().(*captcha.Captcha)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestCaptchaSet(t *testing.T) {
	oldParams := generateCaptchaTestParams()
	newParams := generateCaptchaTestParams()

	captchaInstance := captcha.NewCaptcha(oldParams)
	newConfig := captcha.NewCaptcha(newParams)

	captchaInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, captchaInstance.ModuleName)
	assert.Equal(t, newParams.KeyLen, captchaInstance.KeyLen)
	assert.Equal(t, newParams.ImgWidth, captchaInstance.ImgWidth)
	assert.Equal(t, newParams.ImgHeight, captchaInstance.ImgHeight)
	assert.Equal(t, newParams.MaxSkew, captchaInstance.MaxSkew)
	assert.Equal(t, newParams.DotCount, captchaInstance.DotCount)
}
