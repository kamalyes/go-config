/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\youzan\youzan_test.go
 * @Description: 有赞配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package youzan

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYouZan_Default(t *testing.T) {
	yz := Default()

	assert.NotNil(t, yz)
	assert.Equal(t, "youzan", yz.ModuleName)
	assert.Equal(t, "https://open.youzan.com", yz.Endpoint)
	assert.Equal(t, "demo_client_id", yz.ClientID)
	assert.Equal(t, "demo_client_secret", yz.ClientSecret)
	assert.Equal(t, "silent", yz.AuthorizeType)
	assert.Equal(t, "demo_grant_id", yz.GrantID)
	assert.True(t, yz.Refresh)
}

func TestYouZan_WithEndpoint(t *testing.T) {
	yz := Default().WithEndpoint("https://custom.youzan.com")
	assert.Equal(t, "https://custom.youzan.com", yz.Endpoint)
}

func TestYouZan_WithClientID(t *testing.T) {
	yz := Default().WithClientID("test-client-id")
	assert.Equal(t, "test-client-id", yz.ClientID)
}

func TestYouZan_WithClientSecret(t *testing.T) {
	yz := Default().WithClientSecret("test-secret")
	assert.Equal(t, "test-secret", yz.ClientSecret)
}

func TestYouZan_WithAuthorizeType(t *testing.T) {
	yz := Default().WithAuthorizeType("oauth")
	assert.Equal(t, "oauth", yz.AuthorizeType)
}

func TestYouZan_WithGrantID(t *testing.T) {
	yz := Default().WithGrantID("grant-123")
	assert.Equal(t, "grant-123", yz.GrantID)
}

func TestYouZan_WithRefresh(t *testing.T) {
	yz := Default().WithRefresh(false)
	assert.False(t, yz.Refresh)
}

func TestYouZan_Clone(t *testing.T) {
	original := Default()
	original.ClientID = "original-id"
	original.ClientSecret = "original-secret"
	original.Refresh = false

	cloned := original.Clone().(*YouZan)

	assert.Equal(t, original.ClientID, cloned.ClientID)
	assert.Equal(t, original.ClientSecret, cloned.ClientSecret)
	assert.Equal(t, original.Refresh, cloned.Refresh)

	cloned.ClientID = "new-id"
	assert.Equal(t, "original-id", original.ClientID)
	assert.Equal(t, "new-id", cloned.ClientID)
}

func TestYouZan_Get(t *testing.T) {
	yz := Default()
	result := yz.Get()

	assert.NotNil(t, result)
	resultYZ, ok := result.(*YouZan)
	assert.True(t, ok)
	assert.Equal(t, yz, resultYZ)
}

func TestYouZan_Set(t *testing.T) {
	yz := Default()
	newYZ := &YouZan{
		ModuleName:    "new_youzan",
		Endpoint:      "https://new.youzan.com",
		ClientID:      "new-client-id",
		ClientSecret:  "new-secret",
		AuthorizeType: "oauth",
		GrantID:       "new-grant",
		Refresh:       false,
	}

	yz.Set(newYZ)

	assert.Equal(t, "new_youzan", yz.ModuleName)
	assert.Equal(t, "https://new.youzan.com", yz.Endpoint)
	assert.Equal(t, "new-client-id", yz.ClientID)
	assert.Equal(t, "new-secret", yz.ClientSecret)
	assert.Equal(t, "oauth", yz.AuthorizeType)
	assert.Equal(t, "new-grant", yz.GrantID)
	assert.False(t, yz.Refresh)
}

func TestYouZan_Validate(t *testing.T) {
	yz := &YouZan{
		Endpoint:      "https://open.youzan.com",
		ClientID:      "test-id",
		ClientSecret:  "test-secret",
		AuthorizeType: "silent",
		GrantID:       "test-grant",
	}

	err := yz.Validate()
	assert.NoError(t, err)
}

func TestYouZan_ChainedCalls(t *testing.T) {
	yz := Default().
		WithEndpoint("https://chained.youzan.com").
		WithClientID("chain-id").
		WithClientSecret("chain-secret").
		WithAuthorizeType("oauth").
		WithGrantID("chain-grant").
		WithRefresh(false)

	assert.Equal(t, "https://chained.youzan.com", yz.Endpoint)
	assert.Equal(t, "chain-id", yz.ClientID)
	assert.Equal(t, "chain-secret", yz.ClientSecret)
	assert.Equal(t, "oauth", yz.AuthorizeType)
	assert.Equal(t, "chain-grant", yz.GrantID)
	assert.False(t, yz.Refresh)
}

func TestNewYouZan(t *testing.T) {
	opt := &YouZan{
		ModuleName:    "test",
		Endpoint:      "https://test.youzan.com",
		ClientID:      "test-id",
		ClientSecret:  "test-secret",
		AuthorizeType: "silent",
		GrantID:       "test-grant",
		Refresh:       true,
	}

	yz := NewYouZan(opt)
	assert.NotNil(t, yz)
	assert.Equal(t, opt, yz)
}

func TestDefaultYouZan(t *testing.T) {
	yz := DefaultYouZan()

	assert.Equal(t, "youzan", yz.ModuleName)
	assert.Equal(t, "https://open.youzan.com", yz.Endpoint)
	assert.Equal(t, "silent", yz.AuthorizeType)
	assert.True(t, yz.Refresh)
}

func TestDefaultYouZanConfig(t *testing.T) {
	yz := DefaultYouZanConfig()

	assert.NotNil(t, yz)
	assert.Equal(t, "youzan", yz.ModuleName)
	assert.Equal(t, "https://open.youzan.com", yz.Endpoint)
}
