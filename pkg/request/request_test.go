/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\request\request_test.go
 * @Description: 通用请求结构测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultHeaders(t *testing.T) {
	headers := DefaultHeaders()
	assert.NotNil(t, headers)
	assert.Equal(t, "X-Timestamp", headers.Timestamp)
	assert.Equal(t, "X-Signature", headers.Signature)
	assert.Equal(t, "X-Trace-Id", headers.TraceID)
	assert.Equal(t, "X-Request-Id", headers.RequestID)
	assert.Equal(t, "Authorization", headers.Authorization)
	assert.Equal(t, "X-Device-Id", headers.DeviceID)
	assert.Equal(t, "X-App-Version", headers.AppVersion)
	assert.Equal(t, "X-Platform", headers.Platform)
	assert.Equal(t, "User-Agent", headers.UserAgent)
	assert.Equal(t, "Content-Type", headers.ContentType)
	assert.Equal(t, "Accept", headers.Accept)
	assert.Equal(t, "Accept-Language", headers.AcceptLanguage)
}

func TestHeaders_GetHeaderName(t *testing.T) {
	headers := DefaultHeaders()
	
	tests := []struct {
		fieldType string
		expected  string
	}{
		{"timestamp", "X-Timestamp"},
		{"signature", "X-Signature"},
		{"traceId", "X-Trace-Id"},
		{"requestId", "X-Request-Id"},
		{"authorization", "Authorization"},
		{"deviceId", "X-Device-Id"},
		{"appVersion", "X-App-Version"},
		{"platform", "X-Platform"},
		{"userAgent", "User-Agent"},
		{"contentType", "Content-Type"},
		{"accept", "Accept"},
		{"acceptLanguage", "Accept-Language"},
		{"unknown", ""},
	}
	
	for _, tt := range tests {
		t.Run(tt.fieldType, func(t *testing.T) {
			result := headers.GetHeaderName(tt.fieldType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBaseRequest_WithTimestamp(t *testing.T) {
	req := &BaseRequest{}
	result := req.WithTimestamp("2023-01-01T00:00:00Z")
	assert.Equal(t, "2023-01-01T00:00:00Z", result.Timestamp)
	assert.Equal(t, req, result)
}

func TestBaseRequest_WithSignature(t *testing.T) {
	req := &BaseRequest{}
	result := req.WithSignature("abc123")
	assert.Equal(t, "abc123", result.Signature)
	assert.Equal(t, req, result)
}

func TestBaseRequest_WithTraceID(t *testing.T) {
	req := &BaseRequest{}
	result := req.WithTraceID("trace-123")
	assert.Equal(t, "trace-123", result.TraceID)
	assert.Equal(t, req, result)
}

func TestBaseRequest_WithRequestID(t *testing.T) {
	req := &BaseRequest{}
	result := req.WithRequestID("req-123")
	assert.Equal(t, "req-123", result.RequestID)
	assert.Equal(t, req, result)
}

func TestBaseRequest_WithAuthorization(t *testing.T) {
	req := &BaseRequest{}
	result := req.WithAuthorization("Bearer token123")
	assert.Equal(t, "Bearer token123", result.Authorization)
	assert.Equal(t, req, result)
}

func TestBaseRequest_WithDeviceID(t *testing.T) {
	req := &BaseRequest{}
	result := req.WithDeviceID("device-123")
	assert.Equal(t, "device-123", result.DeviceID)
	assert.Equal(t, req, result)
}

func TestBaseRequest_WithAppVersion(t *testing.T) {
	req := &BaseRequest{}
	result := req.WithAppVersion("1.0.0")
	assert.Equal(t, "1.0.0", result.AppVersion)
	assert.Equal(t, req, result)
}

func TestBaseRequest_WithPlatform(t *testing.T) {
	req := &BaseRequest{}
	result := req.WithPlatform("iOS")
	assert.Equal(t, "iOS", result.Platform)
	assert.Equal(t, req, result)
}

func TestBaseRequest_Validate(t *testing.T) {
	req := &BaseRequest{
		Timestamp: "2023-01-01T00:00:00Z",
		Signature: "abc123",
	}
	err := req.Validate()
	assert.NoError(t, err)
}

func TestBaseRequest_IsEmpty(t *testing.T) {
	req := &BaseRequest{}
	assert.True(t, req.IsEmpty())
	
	req.WithTimestamp("2023-01-01T00:00:00Z")
	assert.False(t, req.IsEmpty())
}

func TestBaseRequest_Clone(t *testing.T) {
	req := &BaseRequest{
		Timestamp:     "2023-01-01T00:00:00Z",
		Signature:     "abc123",
		TraceID:       "trace-123",
		RequestID:     "req-123",
		Authorization: "Bearer token123",
		DeviceID:      "device-123",
		AppVersion:    "1.0.0",
		Platform:      "iOS",
	}
	
	clone := req.Clone()
	assert.NotNil(t, clone)
	assert.Equal(t, req.Timestamp, clone.Timestamp)
	assert.Equal(t, req.Signature, clone.Signature)
	assert.Equal(t, req.TraceID, clone.TraceID)
	assert.Equal(t, req.RequestID, clone.RequestID)
	assert.Equal(t, req.Authorization, clone.Authorization)
	assert.Equal(t, req.DeviceID, clone.DeviceID)
	assert.Equal(t, req.AppVersion, clone.AppVersion)
	assert.Equal(t, req.Platform, clone.Platform)
	
	// 验证深拷贝
	clone.Timestamp = "modified"
	assert.NotEqual(t, req.Timestamp, clone.Timestamp)
}

func TestBaseRequest_ChainedCalls(t *testing.T) {
	req := &BaseRequest{}
	result := req.
		WithTimestamp("2023-01-01T00:00:00Z").
		WithSignature("abc123").
		WithTraceID("trace-123").
		WithRequestID("req-123").
		WithAuthorization("Bearer token123").
		WithDeviceID("device-123").
		WithAppVersion("1.0.0").
		WithPlatform("iOS")
	
	assert.Equal(t, "2023-01-01T00:00:00Z", result.Timestamp)
	assert.Equal(t, "abc123", result.Signature)
	assert.Equal(t, "trace-123", result.TraceID)
	assert.Equal(t, "req-123", result.RequestID)
	assert.Equal(t, "Bearer token123", result.Authorization)
	assert.Equal(t, "device-123", result.DeviceID)
	assert.Equal(t, "1.0.0", result.AppVersion)
	assert.Equal(t, "iOS", result.Platform)
}
