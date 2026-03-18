/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 15:05:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 15:05:00
 * @FilePath: \go-config\pkg\common\attribute_source.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package common

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

// AttributeSourceType 属性提取来源类型
type AttributeSourceType string

const (
	// SourceTypeQuery 从 URL 查询参数提取
	SourceTypeQuery AttributeSourceType = "query"
	// SourceTypeHeader 从 HTTP Header 提取
	SourceTypeHeader AttributeSourceType = "header"
	// SourceTypeCookie 从 Cookie 提取
	SourceTypeCookie AttributeSourceType = "cookie"
)

// IsValid 验证来源类型是否有效
func (t AttributeSourceType) IsValid() bool {
	switch t {
	case SourceTypeQuery, SourceTypeHeader, SourceTypeCookie:
		return true
	default:
		return false
	}
}

// String 返回字符串表示
func (t AttributeSourceType) String() string {
	return string(t)
}

// AttributeSource 属性提取来源配置
type AttributeSource struct {
	Type AttributeSourceType `mapstructure:"type" yaml:"type" json:"type" validate:"required,oneof=query header cookie"` // 来源类型: query, header, cookie
	Key  string              `mapstructure:"key" yaml:"key" json:"key" validate:"required"`                              // 提取的字段名
}

// Validate 验证属性来源配置
func (a *AttributeSource) Validate() error {
	if !a.Type.IsValid() {
		return validator.New().Var(a.Type, "oneof=query header cookie")
	}
	if a.Key == "" {
		return validator.New().Var(a.Key, "required")
	}
	return nil
}

// ExtractAttribute 从请求中按优先级提取属性
func ExtractAttribute(r *http.Request, sources []AttributeSource) string {
	for _, source := range sources {
		if value := ExtractFromSource(r, source); value != "" {
			return value
		}
	}
	return ""
}

// ExtractFromSource 从指定来源提取值
func ExtractFromSource(r *http.Request, source AttributeSource) string {
	switch source.Type {
	case SourceTypeQuery:
		return r.URL.Query().Get(source.Key)
	case SourceTypeHeader:
		return r.Header.Get(source.Key)
	case SourceTypeCookie:
		if cookie, err := r.Cookie(source.Key); err == nil {
			return cookie.Value
		}
		return ""
	default:
		return ""
	}
}
