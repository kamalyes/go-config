/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:57:55
 * @FilePath: \go-config\tests\youzan_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/youzan"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 YouZan 配置参数
func generateYouZanTestParams() *youzan.YouZan {
	return &youzan.YouZan{
		ModuleName:    random.RandString(10, random.CAPITAL),
		Host:          fmt.Sprintf("https://%s.example.com", random.RandString(5, random.CAPITAL)), // 随机生成 Host
		ClientID:      random.RandString(16, random.CAPITAL),                                       // 随机生成客户端ID
		ClientSecret:  random.RandString(32, random.CAPITAL),                                       // 随机生成客户端密钥
		AuthorizeType: random.RandString(10, random.CAPITAL),                                       // 随机生成授权类型
		GrantID:       random.RandString(10, random.CAPITAL),                                       // 随机生成授权ID
		Refresh:       random.FRandBool(),                                                          // 随机生成是否刷新
	}
}

// 将 YouZan 的参数转换为 map
func youzanToMap(youzan *youzan.YouZan) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME":    youzan.ModuleName,
		"HOST":           youzan.Host,
		"CLIENT_ID":      youzan.ClientID,
		"CLIENT_SECRET":  youzan.ClientSecret,
		"AUTHORIZE_TYPE": youzan.AuthorizeType,
		"GRANT_ID":       youzan.GrantID,
		"REFRESH":        youzan.Refresh,
	}
}

// 验证 YouZan 的字段与期望的映射是否相等
func assertYouZanFields(t *testing.T, youzan *youzan.YouZan, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], youzan.ModuleName)
	assert.Equal(t, expected["HOST"], youzan.Host)
	assert.Equal(t, expected["CLIENT_ID"], youzan.ClientID)
	assert.Equal(t, expected["CLIENT_SECRET"], youzan.ClientSecret)
	assert.Equal(t, expected["AUTHORIZE_TYPE"], youzan.AuthorizeType)
	assert.Equal(t, expected["GRANT_ID"], youzan.GrantID)
	assert.Equal(t, expected["REFRESH"], youzan.Refresh)
}

func TestYouZanClone(t *testing.T) {
	params := generateYouZanTestParams()
	original := youzan.NewYouZan(params)
	cloned := original.Clone().(*youzan.YouZan)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestYouZanSet(t *testing.T) {
	oldParams := generateYouZanTestParams()
	newParams := generateYouZanTestParams()

	youzanInstance := youzan.NewYouZan(oldParams)
	newConfig := youzan.NewYouZan(newParams)

	youzanInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, youzanInstance.ModuleName)
	assert.Equal(t, newParams.Host, youzanInstance.Host)
	assert.Equal(t, newParams.ClientID, youzanInstance.ClientID)
	assert.Equal(t, newParams.ClientSecret, youzanInstance.ClientSecret)
	assert.Equal(t, newParams.AuthorizeType, youzanInstance.AuthorizeType)
	assert.Equal(t, newParams.GrantID, youzanInstance.GrantID)
	assert.Equal(t, newParams.Refresh, youzanInstance.Refresh)
}
