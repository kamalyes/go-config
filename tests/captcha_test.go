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

// TestCaptchaDefault 测试默认配置
func TestCaptchaDefault(t *testing.T) {
	defaultConfig := captcha.DefaultCaptcha()
	
	// 检查默认值
	assert.Equal(t, "captcha", defaultConfig.ModuleName)
	assert.Equal(t, 4, defaultConfig.KeyLen)
	assert.Equal(t, 120, defaultConfig.ImgWidth)
	assert.Equal(t, 40, defaultConfig.ImgHeight)
	assert.Equal(t, 0.7, defaultConfig.MaxSkew)
	assert.Equal(t, 80, defaultConfig.DotCount)
}

// TestCaptchaDefaultPointer 测试默认配置指针
func TestCaptchaDefaultPointer(t *testing.T) {
	config := captcha.Default()
	
	assert.NotNil(t, config)
	assert.Equal(t, "captcha", config.ModuleName)
	assert.Equal(t, 4, config.KeyLen)
}

// TestCaptchaChainMethods 测试链式方法
func TestCaptchaChainMethods(t *testing.T) {
	config := captcha.Default().
		WithModuleName("security").
		WithKeyLen(6).
		WithImgWidth(150).
		WithImgHeight(50).
		WithMaxSkew(0.8).
		WithDotCount(100)
	
	assert.Equal(t, "security", config.ModuleName)
	assert.Equal(t, 6, config.KeyLen)
	assert.Equal(t, 150, config.ImgWidth)
	assert.Equal(t, 50, config.ImgHeight)
	assert.Equal(t, 0.8, config.MaxSkew)
	assert.Equal(t, 100, config.DotCount)
}

// TestCaptchaChainMethodsReturnPointer 测试链式方法返回指针
func TestCaptchaChainMethodsReturnPointer(t *testing.T) {
	config1 := captcha.Default()
	config2 := config1.WithKeyLen(8)
	
	// 应该返回同一个实例
	assert.Same(t, config1, config2)
	assert.Equal(t, 8, config1.KeyLen)
}
