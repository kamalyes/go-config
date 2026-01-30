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
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestI18N_AddLanguageMapping(t *testing.T) {
	config := Default()
	config.AddLanguageMapping("zh-tw", "zh")
	assert.Equal(t, "zh", config.LanguageMapping["zh-tw"])

	config.AddLanguageMapping("zh-hk", "zh").AddLanguageMapping("en-gb", "en")
	assert.Equal(t, "zh", config.LanguageMapping["zh-hk"])
	assert.Equal(t, "en", config.LanguageMapping["en-gb"])
}

func TestI18N_AddLanguageMappings(t *testing.T) {
	config := Default()
	mappings := map[string]string{"zh-tw": "zh", "zh-hk": "zh", "en-gb": "en", "en-au": "en"}
	config.AddLanguageMappings(mappings)

	assert.Equal(t, "zh", config.LanguageMapping["zh-tw"])
	assert.Equal(t, "zh", config.LanguageMapping["zh-hk"])
	assert.Equal(t, "en", config.LanguageMapping["en-gb"])
	assert.Equal(t, "en", config.LanguageMapping["en-au"])
}

func TestI18N_GetMappedLanguage(t *testing.T) {
	config := Default()
	assert.Equal(t, "zh", config.GetMappedLanguage("zh-cn"))
	assert.Equal(t, "en", config.GetMappedLanguage("en-us"))
	assert.Equal(t, "fr", config.GetMappedLanguage("fr"))
	assert.Equal(t, "ja", config.GetMappedLanguage("ja"))
}

func TestI18N_AddLegacyLanguageMapping(t *testing.T) {
	config := Default()
	config.AddLegacyLanguageMapping("cn", "zh")
	assert.Equal(t, "zh", config.LegacyLanguageMapping["cn"])

	config.AddLegacyLanguageMapping("eng", "en").AddLegacyLanguageMapping("chn", "zh")
	assert.Equal(t, "en", config.LegacyLanguageMapping["eng"])
	assert.Equal(t, "zh", config.LegacyLanguageMapping["chn"])
}

func TestI18N_AddLegacyLanguageMappings(t *testing.T) {
	config := Default()
	legacyMappings := map[string]string{"cn": "zh", "chn": "zh", "eng": "en", "us": "en"}
	config.AddLegacyLanguageMappings(legacyMappings)

	assert.Equal(t, "zh", config.LegacyLanguageMapping["cn"])
	assert.Equal(t, "zh", config.LegacyLanguageMapping["chn"])
	assert.Equal(t, "en", config.LegacyLanguageMapping["eng"])
	assert.Equal(t, "en", config.LegacyLanguageMapping["us"])
}

func TestI18N_ResolveLanguage(t *testing.T) {
	config := Default()
	config.AddLegacyLanguageMapping("cn", "zh").AddLegacyLanguageMapping("eng", "en")
	config.AddLanguageMapping("zh-tw", "zh").AddLanguageMapping("en-gb", "en")

	tests := []struct {
		input, expected string
	}{
		{"cn", "zh"}, {"eng", "en"}, {"zh-cn", "zh"}, {"zh-tw", "zh"},
		{"en-us", "en"}, {"en-gb", "en"}, {"zh", "zh"}, {"en", "en"},
		{"fr", "en"}, {"ja", "en"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, config.ResolveLanguage(tt.input))
	}
}

func TestI18N_ResolveLanguage_ChainedMapping(t *testing.T) {
	config := Default()
	config.AddLegacyLanguageMapping("chn", "zh-cn")
	config.AddLanguageMapping("zh-cn", "zh")
	assert.Equal(t, "zh", config.ResolveLanguage("chn"))
}

func TestI18N_ResolveLanguage_UnsupportedAfterMapping(t *testing.T) {
	config := Default()
	config.AddLanguageMapping("fr-fr", "fr")
	assert.Equal(t, "en", config.ResolveLanguage("fr-fr"))
}

func TestI18N_IsSupportedLanguage(t *testing.T) {
	config := Default()
	assert.True(t, config.IsSupportedLanguage("en"))
	assert.True(t, config.IsSupportedLanguage("zh"))
	assert.False(t, config.IsSupportedLanguage("fr"))
	assert.False(t, config.IsSupportedLanguage("ja"))
	assert.False(t, config.IsSupportedLanguage("ko"))
}

func TestI18N_ResolveLanguage_WithCustomSupportedLanguages(t *testing.T) {
	config := Default().WithSupportedLanguages([]string{"en", "zh", "ja", "ko"})
	config.AddLanguageMapping("ja-jp", "ja").AddLanguageMapping("ko-kr", "ko")

	assert.Equal(t, "ja", config.ResolveLanguage("ja-jp"))
	assert.Equal(t, "ko", config.ResolveLanguage("ko-kr"))
	assert.Equal(t, "ja", config.ResolveLanguage("ja"))
	assert.Equal(t, "ko", config.ResolveLanguage("ko"))
	assert.Equal(t, "en", config.ResolveLanguage("fr"))
}

func TestI18N_CompleteWorkflow(t *testing.T) {
	config := Default().
		Enable().
		WithDefaultLanguage("en").
		WithSupportedLanguages([]string{"en", "zh", "ja"})

	config.AddLegacyLanguageMappings(map[string]string{"cn": "zh", "chn": "zh", "eng": "en"})
	config.AddLanguageMappings(map[string]string{
		"zh-cn": "zh", "zh-tw": "zh", "zh-hk": "zh",
		"en-us": "en", "en-gb": "en", "ja-jp": "ja",
	})

	tests := []struct {
		input, expected string
	}{
		{"cn", "zh"}, {"chn", "zh"}, {"eng", "en"},
		{"zh-cn", "zh"}, {"zh-tw", "zh"}, {"en-us", "en"}, {"ja-jp", "ja"},
		{"zh", "zh"}, {"en", "en"}, {"ja", "ja"},
		{"fr", "en"}, {"ko", "en"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, config.ResolveLanguage(tt.input))
	}
}

func TestI18N_ParseAcceptLanguage(t *testing.T) {
	config := Default()
	config.AddLanguageMapping("zh-cn", "zh").AddLanguageMapping("en-us", "en")

	tests := []struct {
		input, expected string
	}{
		{"zh-CN,zh;q=0.9,en;q=0.8", "zh"},
		{"en-US,en;q=0.9", "en"},
		{"zh-CN", "zh"},
		{"en", "en"},
		{"fr-FR,fr;q=0.9,en;q=0.8", "en"},
		{"ja,en;q=0.9", "en"},
		{"", "en"},
		{"zh-CN;q=0.9,en-US;q=0.8", "zh"},
		{"en-GB,en-US;q=0.9,en;q=0.8", "en"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, config.ParseAcceptLanguage(tt.input))
	}
}

func TestI18N_ParseAcceptLanguage_WithLegacyMapping(t *testing.T) {
	config := Default()
	config.AddLegacyLanguageMapping("cn", "zh")
	config.AddLanguageMapping("zh-cn", "zh")

	assert.Equal(t, "zh", config.ParseAcceptLanguage("cn,en;q=0.9"))
	assert.Equal(t, "zh", config.ParseAcceptLanguage("zh-CN,en;q=0.9"))
	assert.Equal(t, "en", config.ParseAcceptLanguage("fr,en;q=0.9"))
}

func TestI18N_ParseAcceptLanguage_ComplexScenarios(t *testing.T) {
	config := Default().WithSupportedLanguages([]string{"en", "zh", "ja", "ko"})
	config.AddLanguageMappings(map[string]string{
		"zh-cn": "zh", "zh-tw": "zh", "zh-hk": "zh",
		"en-us": "en", "en-gb": "en",
		"ja-jp": "ja", "ko-kr": "ko",
	})

	tests := []struct {
		input, expected string
	}{
		{"zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7", "zh"},
		{"fr-FR,de-DE;q=0.9,en;q=0.8", "en"},
		{"ja-JP,en-US;q=0.9,zh-CN;q=0.8", "ja"},
		{"ko-KR,ja-JP;q=0.9,zh-CN;q=0.8,en;q=0.7", "ko"},
		{"es-ES,pt-BR;q=0.9,fr-FR;q=0.8", "en"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, config.ParseAcceptLanguage(tt.input))
	}
}

func TestI18N_ResolutionOrder_LegacyFirst(t *testing.T) {
	config := Default().
		WithSupportedLanguages([]string{"en", "zh", "ja"}).
		WithResolutionOrder([]MappingType{"legacy", "standard"})

	config.AddLegacyLanguageMapping("cn", "zh")
	config.AddLanguageMapping("zh-cn", "zh")

	// 遗留映射优先
	assert.Equal(t, "zh", config.ResolveLanguage("cn"))
	assert.Equal(t, "zh", config.ResolveLanguage("zh-cn"))
}

func TestI18N_ResolutionOrder_StandardFirst(t *testing.T) {
	config := Default().
		WithSupportedLanguages([]string{"en", "zh", "ja"}).
		WithResolutionOrder([]MappingType{"standard", "legacy"})

	// 标准映射优先
	config.AddLanguageMapping("zh-cn", "zh")
	config.AddLegacyLanguageMapping("zh-cn", "ja") // 这个会被忽略

	assert.Equal(t, "zh", config.ResolveLanguage("zh-cn"))
}

func TestI18N_ResolutionOrder_LegacyOnly(t *testing.T) {
	config := Default().
		WithSupportedLanguages([]string{"en", "zh", "ja"}).
		WithResolutionOrder([]MappingType{"legacy"})

	config.AddLegacyLanguageMapping("cn", "zh")
	config.AddLanguageMapping("zh-cn", "zh")

	// 仅使用遗留映射
	assert.Equal(t, "zh", config.ResolveLanguage("cn"))
	// 标准映射不生效，zh-cn 不在支持列表中，返回默认语言
	assert.Equal(t, "en", config.ResolveLanguage("zh-cn"))
}

func TestI18N_ResolutionOrder_StandardOnly(t *testing.T) {
	config := Default().
		WithSupportedLanguages([]string{"en", "zh", "ja"}).
		WithResolutionOrder([]MappingType{"standard"})

	config.AddLegacyLanguageMapping("cn", "zh")
	config.AddLanguageMapping("zh-cn", "zh")

	// 仅使用标准映射
	assert.Equal(t, "zh", config.ResolveLanguage("zh-cn"))
	// 遗留映射不生效，cn 不在支持列表中，返回默认语言
	assert.Equal(t, "en", config.ResolveLanguage("cn"))
}

func TestI18N_ResolutionOrder_ChainedMapping(t *testing.T) {
	config := Default().
		WithSupportedLanguages([]string{"en", "zh", "ja"}).
		WithResolutionOrder([]MappingType{"legacy", "standard"})

	// 测试链式映射：chn -> zh-cn -> zh
	config.AddLegacyLanguageMapping("chn", "zh-cn")
	config.AddLanguageMapping("zh-cn", "zh")

	assert.Equal(t, "zh", config.ResolveLanguage("chn"))
}

func TestI18N_ResolutionOrder_CustomOrder(t *testing.T) {
	config := Default().
		WithSupportedLanguages([]string{"en", "zh", "ja"}).
		WithResolutionOrder([]MappingType{"standard", "legacy", "standard"})

	// 测试自定义顺序：可以多次应用同一类型的映射
	config.AddLegacyLanguageMapping("cn", "zh-cn")
	config.AddLanguageMapping("zh-cn", "zh")

	assert.Equal(t, "zh", config.ResolveLanguage("cn"))
}

func TestI18N_ResolutionOrder_EmptyOrder(t *testing.T) {
	config := Default().
		WithSupportedLanguages([]string{"en", "zh", "ja"}).
		WithResolutionOrder([]MappingType{})

	config.AddLegacyLanguageMapping("cn", "zh")
	config.AddLanguageMapping("zh-cn", "zh")

	// 空顺序，不应用任何映射
	assert.Equal(t, "zh", config.ResolveLanguage("zh"))
	assert.Equal(t, "en", config.ResolveLanguage("cn"))    // 不映射，返回默认
	assert.Equal(t, "en", config.ResolveLanguage("zh-cn")) // 不映射，返回默认
}

func TestI18N_WithResolutionOrder(t *testing.T) {
	config := Default().WithResolutionOrder([]MappingType{StandardMapping, LegacyMapping})
	assert.Equal(t, []MappingType{"standard", "legacy"}, config.ResolutionOrder)

	config.WithResolutionOrder([]MappingType{LegacyMapping})
	assert.Equal(t, []MappingType{"legacy"}, config.ResolutionOrder)
}
