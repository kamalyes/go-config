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

import (
	"github.com/kamalyes/go-config/internal"
)

// MessageLoader 消息加载器接口
type MessageLoader interface {
	LoadMessages(language string) (map[string]string, error)
}

// I18N 国际化中间件配置
type I18N struct {
	ModuleName             string            `mapstructure:"module_name" yaml:"module-name" json:"module_name"`                                         // 模块名称
	Enabled                bool              `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                                     // 是否启用国际化
	DefaultLanguage        string            `mapstructure:"default_language" yaml:"default-language" json:"default_language"`                          // 默认语言
	SupportedLanguages     []string          `mapstructure:"supported_languages" yaml:"supported-languages" json:"supported_languages"`               // 支持的语言
	LanguageMapping        map[string]string `mapstructure:"language_mapping" yaml:"language-mapping" json:"language_mapping"`                         // 语言映射关系
	DetectionOrder         []string          `mapstructure:"detection_order" yaml:"detection-order" json:"detection_order"`                            // 语言检测顺序
	LanguageParam          string            `mapstructure:"language_param" yaml:"language-param" json:"language_param"`                               // 语言参数名称（用于query和cookie）
	LanguageHeader         string            `mapstructure:"language_header" yaml:"language-header" json:"language_header"`                            // 语言头名称
	MessagesPath           string            `mapstructure:"messages_path" yaml:"messages-path" json:"messages_path"`                                  // 消息文件路径
	CustomMessagePaths     map[string]string `mapstructure:"custom_message_paths" yaml:"custom-message-paths" json:"custom_message_paths"`             // 自定义消息文件路径映射
	EnableFallback         bool              `mapstructure:"enable_fallback" yaml:"enable-fallback" json:"enable_fallback"`                            // 是否启用回退到默认语言
	MessageLoader          MessageLoader     `mapstructure:"-" yaml:"-" json:"-"`                                                                      // 自定义消息加载器
	CookieName             string            `mapstructure:"cookie_name" yaml:"cookie-name" json:"cookie_name"`                                        // Cookie名称
}

// Default 创建默认国际化配置
func Default() *I18N {
	return &I18N{
		ModuleName:         "i18n",
		Enabled:            false,
		DefaultLanguage:    "en",
		SupportedLanguages: []string{"en", "zh", "ja", "ko"},
		LanguageMapping: map[string]string{
			"zh-cn": "zh",
			"zh-tw": "zh-tw",
			"en-us": "en",
			"fr-fr": "fr",
		},
		DetectionOrder:     []string{"header", "query", "cookie", "default"},
		LanguageParam:      "lang",
		LanguageHeader:     "Accept-Language",
		MessagesPath:       "./locales",
		CustomMessagePaths: make(map[string]string),
		EnableFallback:     true,
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
		ModuleName:         i.ModuleName,
		Enabled:            i.Enabled,
		DefaultLanguage:    i.DefaultLanguage,
		LanguageParam:      i.LanguageParam,
		LanguageHeader:     i.LanguageHeader,
		MessagesPath:       i.MessagesPath,
		EnableFallback:     i.EnableFallback,
		MessageLoader:      i.MessageLoader,
		CookieName:         i.CookieName,
	}
	clone.SupportedLanguages = append([]string(nil), i.SupportedLanguages...)
	clone.DetectionOrder = append([]string(nil), i.DetectionOrder...)
	
	// 深拷贝映射
	clone.LanguageMapping = make(map[string]string)
	for k, v := range i.LanguageMapping {
		clone.LanguageMapping[k] = v
	}
	
	clone.CustomMessagePaths = make(map[string]string)
	for k, v := range i.CustomMessagePaths {
		clone.CustomMessagePaths[k] = v
	}
	
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

// WithMessagesPath 设置消息文件路径
func (i *I18N) WithMessagesPath(messagesPath string) *I18N {
	i.MessagesPath = messagesPath
	return i
}

// WithLanguageHeader 设置语言头名称
func (i *I18N) WithLanguageHeader(languageHeader string) *I18N {
	i.LanguageHeader = languageHeader
	return i
}

// WithLanguageParam 设置语言参数名称
func (i *I18N) WithLanguageParam(languageParam string) *I18N {
	i.LanguageParam = languageParam
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
