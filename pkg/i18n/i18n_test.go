/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\i18n\i18n_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package i18n

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestI18N_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "i18n", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "en", config.DefaultLanguage)
	assert.Equal(t, []string{"en", "zh"}, config.SupportedLanguages)
	assert.Equal(t, "lang", config.LanguageParam)
	assert.Equal(t, "Accept-Language", config.LanguageHeader)
	assert.Equal(t, "./locales", config.MessagesPath)
	assert.True(t, config.EnableFallback)
	assert.Equal(t, "language", config.CookieName)
}

func TestI18N_WithDefaultLanguage(t *testing.T) {
	config := Default().WithDefaultLanguage("zh")
	assert.Equal(t, "zh", config.DefaultLanguage)
}

func TestI18N_WithSupportedLanguages(t *testing.T) {
	languages := []string{"en", "zh", "ja", "ko"}
	config := Default().WithSupportedLanguages(languages)
	assert.Equal(t, languages, config.SupportedLanguages)
}

func TestI18N_WithMessagesPath(t *testing.T) {
	config := Default().WithMessagesPath("/app/i18n")
	assert.Equal(t, "/app/i18n", config.MessagesPath)
}

func TestI18N_WithLanguageHeader(t *testing.T) {
	config := Default().WithLanguageHeader("X-Language")
	assert.Equal(t, "X-Language", config.LanguageHeader)
}

func TestI18N_WithLanguageParam(t *testing.T) {
	config := Default().WithLanguageParam("locale")
	assert.Equal(t, "locale", config.LanguageParam)
}

func TestI18N_WithCookieName(t *testing.T) {
	config := Default().WithCookieName("user_locale")
	assert.Equal(t, "user_locale", config.CookieName)
}

func TestI18N_Enable(t *testing.T) {
	config := Default().Enable()
	assert.True(t, config.Enabled)
}

func TestI18N_Disable(t *testing.T) {
	config := Default().Enable().Disable()
	assert.False(t, config.Enabled)
}

func TestI18N_IsEnabled(t *testing.T) {
	config := Default()
	assert.False(t, config.IsEnabled())
	config.Enable()
	assert.True(t, config.IsEnabled())
}

func TestI18N_Clone(t *testing.T) {
	original := Default().
		WithDefaultLanguage("zh").
		WithSupportedLanguages([]string{"en", "zh", "ja"}).
		WithMessagesPath("/custom/path").
		WithLanguageHeader("X-Custom-Lang").
		Enable()

	cloned := original.Clone().(*I18N)

	// 验证值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.DefaultLanguage, cloned.DefaultLanguage)
	assert.Equal(t, original.LanguageParam, cloned.LanguageParam)
	assert.Equal(t, original.LanguageHeader, cloned.LanguageHeader)
	assert.Equal(t, original.MessagesPath, cloned.MessagesPath)

	// 验证切片独立性
	cloned.SupportedLanguages = append(cloned.SupportedLanguages, "ko")
	assert.NotEqual(t, len(original.SupportedLanguages), len(cloned.SupportedLanguages))

	cloned.DetectionOrder = append(cloned.DetectionOrder, "custom")
	assert.NotEqual(t, len(original.DetectionOrder), len(cloned.DetectionOrder))

	// 验证map独立性
	cloned.LanguageMapping["test"] = "test"
	assert.NotEqual(t, len(original.LanguageMapping), len(cloned.LanguageMapping))

	cloned.CustomMessagePaths["test"] = "test"
	assert.NotEqual(t, len(original.CustomMessagePaths), len(cloned.CustomMessagePaths))
}

func TestI18N_Get(t *testing.T) {
	config := Default().WithDefaultLanguage("zh")
	got := config.Get()
	assert.NotNil(t, got)
	i18nConfig, ok := got.(*I18N)
	assert.True(t, ok)
	assert.Equal(t, "zh", i18nConfig.DefaultLanguage)
}

func TestI18N_Set(t *testing.T) {
	config := Default()
	newConfig := &I18N{
		ModuleName:      "new-i18n",
		Enabled:         true,
		DefaultLanguage: "ja",
		LanguageParam:   "locale",
	}

	config.Set(newConfig)
	assert.Equal(t, "new-i18n", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "ja", config.DefaultLanguage)
	assert.Equal(t, "locale", config.LanguageParam)
}

func TestI18N_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestI18N_ChainedCalls(t *testing.T) {
	config := Default().
		Enable().
		WithDefaultLanguage("zh").
		WithSupportedLanguages([]string{"en", "zh", "ja"}).
		WithMessagesPath("/i18n/messages").
		WithLanguageHeader("X-Locale").
		WithLanguageParam("locale").
		WithCookieName("app_locale")

	assert.True(t, config.Enabled)
	assert.Equal(t, "zh", config.DefaultLanguage)
	assert.Len(t, config.SupportedLanguages, 3)
	assert.Equal(t, "/i18n/messages", config.MessagesPath)
	assert.Equal(t, "X-Locale", config.LanguageHeader)
	assert.Equal(t, "locale", config.LanguageParam)
	assert.Equal(t, "app_locale", config.CookieName)
}

func TestI18N_LanguageMapping(t *testing.T) {
	config := Default()
	assert.NotNil(t, config.LanguageMapping)
	assert.Equal(t, "zh", config.LanguageMapping["zh-cn"])
	assert.Equal(t, "en", config.LanguageMapping["en-us"])
}

func TestI18N_DetectionOrder(t *testing.T) {
	config := Default()
	assert.NotNil(t, config.DetectionOrder)
	assert.Equal(t, []string{"header", "query", "cookie", "default"}, config.DetectionOrder)
}
