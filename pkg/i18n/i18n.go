/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 18:00:00
 * @FilePath: \go-config\pkg\i18n\i18n.go
 * @Description: 国际化中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package i18n

import "github.com/kamalyes/go-config/internal"

// I18N 国际化中间件配置
type I18N struct {
	ModuleName         string   `mapstructure:"module_name" yaml:"module-name" json:"module_name"`                         // 模块名称
	Enabled            bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                     // 是否启用国际化
	DefaultLanguage    string   `mapstructure:"default_language" yaml:"default-language" json:"default_language"`          // 默认语言
	SupportedLanguages []string `mapstructure:"supported_languages" yaml:"supported-languages" json:"supported_languages"` // 支持的语言
	LocaleDir          string   `mapstructure:"locale_dir" yaml:"locale-dir" json:"locale_dir"`                            // 语言文件目录
	HeaderName         string   `mapstructure:"header_name" yaml:"header-name" json:"header_name"`                         // 语言头部名称
	QueryParam         string   `mapstructure:"query_param" yaml:"query-param" json:"query_param"`                         // 查询参数名称
	CookieName         string   `mapstructure:"cookie_name" yaml:"cookie-name" json:"cookie_name"`                         // Cookie名称
}

// Default 创建默认国际化配置
func Default() *I18N {
	return &I18N{
		ModuleName:         "i18n",
		Enabled:            false,
		DefaultLanguage:    "en",
		SupportedLanguages: []string{"en", "zh", "ja", "ko"},
		LocaleDir:          "./locales",
		HeaderName:         "Accept-Language",
		QueryParam:         "lang",
		CookieName:         "language",
	}
}

// Get 返回配置接口
func (i *I18N) Get() interface{} {
	return i
}

// Set 设置配置数据
func (i *I18N) Set(data interface{}) {
	if cfg, ok := data.(*I18N); ok {
		*i = *cfg
	}
}

// Clone 返回配置的副本
func (i *I18N) Clone() internal.Configurable {
	clone := &I18N{
		ModuleName:      i.ModuleName,
		Enabled:         i.Enabled,
		DefaultLanguage: i.DefaultLanguage,
		LocaleDir:       i.LocaleDir,
		HeaderName:      i.HeaderName,
		QueryParam:      i.QueryParam,
		CookieName:      i.CookieName,
	}
	clone.SupportedLanguages = append([]string(nil), i.SupportedLanguages...)
	return clone
}

// Validate 验证配置
func (i *I18N) Validate() error {
	return internal.ValidateStruct(i)
}

// WithDefaultLanguage 设置默认语言
func (i *I18N) WithDefaultLanguage(defaultLanguage string) *I18N {
	i.DefaultLanguage = defaultLanguage
	return i
}

// WithSupportedLanguages 设置支持的语言
func (i *I18N) WithSupportedLanguages(supportedLanguages []string) *I18N {
	i.SupportedLanguages = supportedLanguages
	return i
}

// WithLocaleDir 设置语言文件目录
func (i *I18N) WithLocaleDir(localeDir string) *I18N {
	i.LocaleDir = localeDir
	return i
}

// WithHeaderName 设置语言头部名称
func (i *I18N) WithHeaderName(headerName string) *I18N {
	i.HeaderName = headerName
	return i
}

// WithQueryParam 设置查询参数名称
func (i *I18N) WithQueryParam(queryParam string) *I18N {
	i.QueryParam = queryParam
	return i
}

// WithCookieName 设置Cookie名称
func (i *I18N) WithCookieName(cookieName string) *I18N {
	i.CookieName = cookieName
	return i
}

// Enable 启用国际化中间件
func (i *I18N) Enable() *I18N {
	i.Enabled = true
	return i
}

// Disable 禁用国际化中间件
func (i *I18N) Disable() *I18N {
	i.Enabled = false
	return i
}

// IsEnabled 检查是否启用
func (i *I18N) IsEnabled() bool {
	return i.Enabled
}
