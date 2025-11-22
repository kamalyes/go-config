/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-21 23:58:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-21 23:58:00
 * @FilePath: \go-config\pkg\banner\banner_test.go
 * @Description: Banner配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package banner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBanner_Default(t *testing.T) {
	banner := Default()

	assert.NotNil(t, banner)
	assert.Equal(t, "banner", banner.ModuleName)
	assert.True(t, banner.Enabled)
	assert.NotEmpty(t, banner.Template)
	assert.Equal(t, "Go RPC Gateway", banner.Title)
	assert.Equal(t, "A high-performance RPC gateway service", banner.Description)
	assert.Equal(t, "kamalyes", banner.Author)
	assert.Equal(t, "501893067@qq.com", banner.Email)
	assert.Equal(t, "https://github.com/kamalyes/go-rpc-gateway", banner.Website)
}

func TestBanner_WithTitle(t *testing.T) {
	banner := Default()
	result := banner.WithTitle("Custom Title")

	assert.Equal(t, "Custom Title", result.Title)
	assert.Equal(t, banner, result) // 验证链式调用返回同一实例
}

func TestBanner_WithTemplate(t *testing.T) {
	banner := Default()
	customTemplate := "Custom Banner Template"
	result := banner.WithTemplate(customTemplate)

	assert.Equal(t, customTemplate, result.Template)
	assert.Equal(t, banner, result)
}

func TestBanner_WithDescription(t *testing.T) {
	banner := Default()
	result := banner.WithDescription("New description")

	assert.Equal(t, "New description", result.Description)
	assert.Equal(t, banner, result)
}

func TestBanner_WithAuthor(t *testing.T) {
	banner := Default()
	result := banner.WithAuthor("John Doe")

	assert.Equal(t, "John Doe", result.Author)
	assert.Equal(t, banner, result)
}

func TestBanner_WithEmail(t *testing.T) {
	banner := Default()
	result := banner.WithEmail("john@example.com")

	assert.Equal(t, "john@example.com", result.Email)
	assert.Equal(t, banner, result)
}

func TestBanner_WithWebsite(t *testing.T) {
	banner := Default()
	result := banner.WithWebsite("https://example.com")

	assert.Equal(t, "https://example.com", result.Website)
	assert.Equal(t, banner, result)
}

func TestBanner_Enable(t *testing.T) {
	banner := Default()
	banner.Enabled = false
	result := banner.Enable()

	assert.True(t, result.Enabled)
	assert.True(t, result.IsEnabled())
	assert.Equal(t, banner, result)
}

func TestBanner_Disable(t *testing.T) {
	banner := Default()
	result := banner.Disable()

	assert.False(t, result.Enabled)
	assert.False(t, result.IsEnabled())
	assert.Equal(t, banner, result)
}

func TestBanner_IsEnabled(t *testing.T) {
	banner := Default()
	assert.True(t, banner.IsEnabled())

	banner.Enabled = false
	assert.False(t, banner.IsEnabled())
}

func TestBanner_Get(t *testing.T) {
	banner := Default()
	result := banner.Get()

	assert.NotNil(t, result)
	assert.Equal(t, banner, result)
}

func TestBanner_Set(t *testing.T) {
	banner := Default()
	newBanner := &Banner{
		ModuleName:  "custom",
		Enabled:     false,
		Title:       "New Title",
		Description: "New Description",
		Author:      "New Author",
		Email:       "new@example.com",
		Website:     "https://new.com",
	}

	banner.Set(newBanner)

	assert.Equal(t, "custom", banner.ModuleName)
	assert.False(t, banner.Enabled)
	assert.Equal(t, "New Title", banner.Title)
	assert.Equal(t, "New Description", banner.Description)
	assert.Equal(t, "New Author", banner.Author)
	assert.Equal(t, "new@example.com", banner.Email)
	assert.Equal(t, "https://new.com", banner.Website)
}

func TestBanner_Set_InvalidType(t *testing.T) {
	banner := Default()
	originalTitle := banner.Title

	// 传入错误类型不应改变配置
	banner.Set("invalid type")

	assert.Equal(t, originalTitle, banner.Title)
}

func TestBanner_Clone(t *testing.T) {
	banner := Default()
	banner.Title = "Original Title"

	cloned := banner.Clone()

	assert.NotNil(t, cloned)
	clonedBanner, ok := cloned.(*Banner)
	assert.True(t, ok)
	assert.Equal(t, banner.Title, clonedBanner.Title)
	assert.Equal(t, banner.ModuleName, clonedBanner.ModuleName)

	// 验证是独立副本
	clonedBanner.Title = "Modified Title"
	assert.NotEqual(t, banner.Title, clonedBanner.Title)
}

func TestBanner_Validate(t *testing.T) {
	banner := Default()
	err := banner.Validate()
	assert.NoError(t, err)
}

func TestBanner_ChainedCalls(t *testing.T) {
	banner := Default()
	result := banner.
		WithTitle("Chained Title").
		WithDescription("Chained Description").
		WithAuthor("Chained Author").
		WithEmail("chained@example.com").
		WithWebsite("https://chained.com").
		Enable()

	assert.Equal(t, "Chained Title", result.Title)
	assert.Equal(t, "Chained Description", result.Description)
	assert.Equal(t, "Chained Author", result.Author)
	assert.Equal(t, "chained@example.com", result.Email)
	assert.Equal(t, "https://chained.com", result.Website)
	assert.True(t, result.Enabled)
}
