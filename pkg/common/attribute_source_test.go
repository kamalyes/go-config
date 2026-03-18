/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 15:05:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 15:05:00
 * @FilePath: \go-config\pkg\common\attribute_source_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package common

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeSource_Validate(t *testing.T) {
	tests := []struct {
		name        string
		source      AttributeSource
		expectError bool
	}{
		{"Valid Query Source", AttributeSource{Type: SourceTypeQuery, Key: "param"}, false},
		{"Valid Header Source", AttributeSource{Type: SourceTypeHeader, Key: "Authorization"}, false},
		{"Invalid Source Type", AttributeSource{Type: "invalid", Key: "param"}, true},
		{"Empty Key", AttributeSource{Type: SourceTypeQuery, Key: ""}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.source.Validate()
			assert.Equal(t, tt.expectError, err != nil)
		})
	}
}

func TestExtractAttribute(t *testing.T) {
	// 准备带 Header 的请求
	reqWithHeader := httptest.NewRequest(http.MethodGet, "/", nil)
	reqWithHeader.Header.Set("X-Test-Header", "header_value")

	// 准备带 Cookie 的请求
	reqWithCookie := httptest.NewRequest(http.MethodGet, "/", nil)
	reqWithCookie.AddCookie(&http.Cookie{Name: "session", Value: "cookie_value"})

	tests := []struct {
		name     string
		sources  []AttributeSource
		request  *http.Request
		expected string
	}{
		{"Extract from Query", []AttributeSource{{Type: SourceTypeQuery, Key: "test"}}, httptest.NewRequest(http.MethodGet, "/?test=value", nil), "value"},
		{"Extract from Header", []AttributeSource{{Type: SourceTypeHeader, Key: "X-Test-Header"}}, reqWithHeader, "header_value"},
		{"Extract from Cookie", []AttributeSource{{Type: SourceTypeCookie, Key: "session"}}, reqWithCookie, "cookie_value"},
		{"No Source Found", []AttributeSource{{Type: SourceTypeQuery, Key: "missing_param"}}, httptest.NewRequest(http.MethodGet, "/", nil), ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractAttribute(tt.request, tt.sources)
			assert.Equal(t, tt.expected, result)
		})
	}
}
