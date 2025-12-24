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
	"github.com/kamalyes/go-toolbox/pkg/syncx"
)

// MessageLoader 消息加载器接口
type MessageLoader interface {
	LoadMessages(language string) (map[string]string, error)
}

// I18N 国际化中间件配置
type I18N struct {
	ModuleName         string            `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                           // 模块名称
	Enabled            bool              `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                      // 是否启用国际化
	DefaultLanguage    string            `mapstructure:"default-language" yaml:"default-language" json:"defaultLanguage"`            // 默认语言
	SupportedLanguages []string          `mapstructure:"supported-languages" yaml:"supported-languages" json:"supportedLanguages"`   // 支持的语言
	LanguageMapping    map[string]string `mapstructure:"language-mapping" yaml:"language-mapping" json:"languageMapping"`            // 语言映射关系
	DetectionOrder     []string          `mapstructure:"detection-order" yaml:"detection-order" json:"detectionOrder"`               // 语言检测顺序
	LanguageParam      string            `mapstructure:"language-param" yaml:"language-param" json:"languageParam"`                  // 语言参数名称（用于query和cookie）
	LanguageHeader     string            `mapstructure:"language-header" yaml:"language-header" json:"languageHeader"`               // 语言头名称
	MessagesPath       string            `mapstructure:"messages-path" yaml:"messages-path" json:"messagesPath"`                     // 消息文件路径
	CustomMessagePaths map[string]string `mapstructure:"custom-message-paths" yaml:"custom-message-paths" json:"customMessagePaths"` // 自定义消息文件路径映射
	EnableFallback     bool              `mapstructure:"enable-fallback" yaml:"enable-fallback" json:"enableFallback"`               // 是否启用回退到默认语言
	MessageLoader      MessageLoader     `mapstructure:"-" yaml:"-" json:""`                                                         // 自定义消息加载器
	CookieName         string            `mapstructure:"cookie-name" yaml:"cookie-name" json:"cookieName"`                           // Cookie名称
}

// Default 创建默认国际化配置
func Default() *I18N {
	return &I18N{
		ModuleName:         "i18n",
		Enabled:            false,
		DefaultLanguage:    "en",
		SupportedLanguages: []string{"en", "zh"},
		LanguageMapping: map[string]string{
			"zh-cn": "zh",
			"en-us": "en",
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
	var cloned I18N
	if err := syncx.DeepCopy(&cloned, i); err != nil {
		// 如果深拷贝失败，返回空配置
		return &I18N{}
	}
	return &cloned
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
