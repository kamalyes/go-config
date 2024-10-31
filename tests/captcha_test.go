/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 09:55:15
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 10:55:28
 * @FilePath: \go-config\captcha\captcha_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"reflect"
	"testing"

	"github.com/kamalyes/go-config/pkg/captcha"
)

// 公共函数：生成测试数据
func generateTestCaptcha() *captcha.Captcha {
	return captcha.NewCaptcha("testModule", 6, 200, 100, 0.7, 100)
}

// 公共函数：生成测试映射数据
func generateCaptchaTestDataMap() map[string]interface{} {
	return map[string]interface{}{
		"keyLen":    6,
		"imgWidth":  200,
		"imgHeight": 100,
		"maxSkew":   0.7,
		"dotCount":  100,
	}
}

// 公共函数：校验 Captcha 结构体的字段
func assertCaptchaFields(t *testing.T, captcha *captcha.Captcha, testModuleName ...string) {
	// 设置默认的 moduleName
	moduleName := "testModule"
	if len(testModuleName) > 0 {
		moduleName = testModuleName[0] // 使用第一个参数作为 moduleName
	}

	// 校验 ModuleName 字段
	if captcha.ModuleName != moduleName {
		t.Errorf("Expected ModuleName to be '%s', got '%s'", moduleName, captcha.ModuleName)
	}
	if captcha.KeyLen != 6 {
		t.Errorf("Expected KeyLen to be 6, got %d", captcha.KeyLen)
	}
	if captcha.ImgWidth != 200 {
		t.Errorf("Expected ImgWidth to be 200, got %d", captcha.ImgWidth)
	}
	if captcha.ImgHeight != 100 {
		t.Errorf("Expected ImgHeight to be 100, got %d", captcha.ImgHeight)
	}
	if captcha.MaxSkew != 0.7 {
		t.Errorf("Expected MaxSkew to be 0.7, got %f", captcha.MaxSkew)
	}
	if captcha.DotCount != 100 {
		t.Errorf("Expected DotCount to be 100, got %d", captcha.DotCount)
	}
}

func TestNewCaptcha(t *testing.T) {
	captcha := generateTestCaptcha()
	assertCaptchaFields(t, captcha)
}

func TestCaptchaToMap(t *testing.T) {
	captcha := generateTestCaptcha()
	result := captcha.ToMap()

	expectedKeys := []string{"keyLen", "imgWidth", "imgHeight", "maxSkew", "dotCount"}
	for _, key := range expectedKeys {
		if _, ok := result[key]; !ok {
			t.Errorf("Expected key '%s' not found in result map", key)
		}
	}
}

func TestCaptchaFromMap(t *testing.T) {
	captcha := &captcha.Captcha{}
	data := generateCaptchaTestDataMap() // 使用公共函数生成测试映射数据
	data["moduleName"] = "newModule"
	captcha.FromMap(data)

	assertCaptchaFields(t, captcha, "newModule") // 使用公共函数校验字段
}

func TestCaptchaClone(t *testing.T) {
	data := generateTestCaptcha()
	clone := data.Clone().(*captcha.Captcha)

	if !reflect.DeepEqual(data, clone) {
		t.Error("Expected clone to be equal to original")
	}

	// Modify the clone and check that the original is unaffected
	clone.KeyLen = 8
	if data.KeyLen == clone.KeyLen {
		t.Error("Expected original KeyLen to be unaffected by clone modification")
	}
}
