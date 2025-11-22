/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-21 23:59:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-21 23:59:00
 * @FilePath: \go-config\pkg\captcha\captcha_test.go
 * @Description: 验证码配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package captcha

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaptcha_Default(t *testing.T) {
	captcha := Default()

	assert.NotNil(t, captcha)
	assert.Equal(t, "captcha", captcha.ModuleName)
	assert.Equal(t, 4, captcha.KeyLen)
	assert.Equal(t, 120, captcha.ImgWidth)
	assert.Equal(t, 40, captcha.ImgHeight)
	assert.Equal(t, 0.7, captcha.MaxSkew)
	assert.Equal(t, 80, captcha.DotCount)
}

func TestCaptcha_WithModuleName(t *testing.T) {
	captcha := Default().WithModuleName("custom_captcha")

	assert.Equal(t, "custom_captcha", captcha.ModuleName)
}

func TestCaptcha_WithKeyLen(t *testing.T) {
	captcha := Default().WithKeyLen(6)

	assert.Equal(t, 6, captcha.KeyLen)
}

func TestCaptcha_WithImgWidth(t *testing.T) {
	captcha := Default().WithImgWidth(150)

	assert.Equal(t, 150, captcha.ImgWidth)
}

func TestCaptcha_WithImgHeight(t *testing.T) {
	captcha := Default().WithImgHeight(50)

	assert.Equal(t, 50, captcha.ImgHeight)
}

func TestCaptcha_WithMaxSkew(t *testing.T) {
	captcha := Default().WithMaxSkew(0.8)

	assert.Equal(t, 0.8, captcha.MaxSkew)
}

func TestCaptcha_WithDotCount(t *testing.T) {
	captcha := Default().WithDotCount(100)

	assert.Equal(t, 100, captcha.DotCount)
}

func TestCaptcha_Get(t *testing.T) {
	captcha := Default()
	result := captcha.Get()

	assert.NotNil(t, result)
	resultCaptcha, ok := result.(*Captcha)
	assert.True(t, ok)
	assert.Equal(t, captcha, resultCaptcha)
}

func TestCaptcha_Set(t *testing.T) {
	captcha := Default()
	newCaptcha := &Captcha{
		ModuleName: "new_captcha",
		KeyLen:     8,
		ImgWidth:   200,
		ImgHeight:  80,
		MaxSkew:    0.9,
		DotCount:   150,
	}

	captcha.Set(newCaptcha)

	assert.Equal(t, "new_captcha", captcha.ModuleName)
	assert.Equal(t, 8, captcha.KeyLen)
	assert.Equal(t, 200, captcha.ImgWidth)
	assert.Equal(t, 80, captcha.ImgHeight)
	assert.Equal(t, 0.9, captcha.MaxSkew)
	assert.Equal(t, 150, captcha.DotCount)
}

func TestCaptcha_Clone(t *testing.T) {
	original := Default()
	original.ModuleName = "original"
	original.KeyLen = 5
	original.ImgWidth = 130
	original.ImgHeight = 45
	original.MaxSkew = 0.75
	original.DotCount = 90

	cloned := original.Clone().(*Captcha)

	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.KeyLen, cloned.KeyLen)
	assert.Equal(t, original.ImgWidth, cloned.ImgWidth)
	assert.Equal(t, original.ImgHeight, cloned.ImgHeight)
	assert.Equal(t, original.MaxSkew, cloned.MaxSkew)
	assert.Equal(t, original.DotCount, cloned.DotCount)

	// 验证独立性
	cloned.KeyLen = 10
	assert.Equal(t, 5, original.KeyLen)
	assert.Equal(t, 10, cloned.KeyLen)
}

func TestCaptcha_Validate(t *testing.T) {
	captcha := Default()
	err := captcha.Validate()
	assert.NoError(t, err)
}

func TestCaptcha_Validate_InvalidKeyLen(t *testing.T) {
	captcha := Default()
	captcha.KeyLen = 0

	err := captcha.Validate()
	assert.Error(t, err)
}

func TestCaptcha_Validate_InvalidImgWidth(t *testing.T) {
	captcha := Default()
	captcha.ImgWidth = 0

	err := captcha.Validate()
	assert.Error(t, err)
}

func TestCaptcha_Validate_InvalidImgHeight(t *testing.T) {
	captcha := Default()
	captcha.ImgHeight = 0

	err := captcha.Validate()
	assert.Error(t, err)
}

func TestCaptcha_Validate_InvalidMaxSkew(t *testing.T) {
	captcha := Default()
	captcha.MaxSkew = 0.3 // 低于 0.5

	err := captcha.Validate()
	assert.Error(t, err)

	captcha.MaxSkew = 1.5 // 高于 1.0
	err = captcha.Validate()
	assert.Error(t, err)
}

func TestCaptcha_Validate_InvalidDotCount(t *testing.T) {
	captcha := Default()
	captcha.DotCount = 0

	err := captcha.Validate()
	assert.Error(t, err)
}

func TestCaptcha_ChainedCalls(t *testing.T) {
	captcha := Default().
		WithModuleName("chained").
		WithKeyLen(7).
		WithImgWidth(180).
		WithImgHeight(60).
		WithMaxSkew(0.85).
		WithDotCount(120)

	assert.Equal(t, "chained", captcha.ModuleName)
	assert.Equal(t, 7, captcha.KeyLen)
	assert.Equal(t, 180, captcha.ImgWidth)
	assert.Equal(t, 60, captcha.ImgHeight)
	assert.Equal(t, 0.85, captcha.MaxSkew)
	assert.Equal(t, 120, captcha.DotCount)

	err := captcha.Validate()
	assert.NoError(t, err)
}

func TestNewCaptcha(t *testing.T) {
	opt := &Captcha{
		ModuleName: "test",
		KeyLen:     5,
		ImgWidth:   140,
		ImgHeight:  45,
		MaxSkew:    0.8,
		DotCount:   90,
	}

	captcha := NewCaptcha(opt)

	assert.NotNil(t, captcha)
	assert.Equal(t, opt, captcha)
}

func TestDefaultCaptcha(t *testing.T) {
	captcha := DefaultCaptcha()

	assert.Equal(t, "captcha", captcha.ModuleName)
	assert.Equal(t, 4, captcha.KeyLen)
	assert.Equal(t, 120, captcha.ImgWidth)
	assert.Equal(t, 40, captcha.ImgHeight)
	assert.Equal(t, 0.7, captcha.MaxSkew)
	assert.Equal(t, 80, captcha.DotCount)
}
