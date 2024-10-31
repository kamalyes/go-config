/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 12:50:58
 * @FilePath: \go-config\youzan\youzan.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"github.com/kamalyes/go-config/pkg/youzan"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getValidYouZanConfig 返回有效的 YouZan 配置
func getValidYouZanConfig() *youzan.YouZan {
	return youzan.NewYouZan("test-module", "https://api.youzan.com", "test-client-id", "test-client-secret", "authorization_code", "test-grant-id", true)
}

// getInvalidYouZanConfig 返回无效的 YouZan 配置
func getInvalidYouZanConfig() *youzan.YouZan {
	return youzan.NewYouZan("", "", "", "", "", "", false)
}

func TestNewYouZan(t *testing.T) {
	// 测试创建 YouZan 实例
	youzan := getValidYouZanConfig()

	assert.NotNil(t, youzan)
	assert.Equal(t, "test-module", youzan.ModuleName)
	assert.Equal(t, "https://api.youzan.com", youzan.Host)
	assert.Equal(t, "test-client-id", youzan.ClientID)
	assert.Equal(t, "test-client-secret", youzan.ClientSecret)
	assert.Equal(t, "authorization_code", youzan.AuthorizeType)
	assert.Equal(t, "test-grant-id", youzan.GrantID)
	assert.True(t, youzan.Refresh)
}

func TestYouZan_Validate(t *testing.T) {
	// 测试有效配置
	validYouZan := getValidYouZanConfig()
	assert.NoError(t, validYouZan.Validate())

	// 测试无效配置
	invalidYouZan := getInvalidYouZanConfig()
	assert.Error(t, invalidYouZan.Validate())
}
