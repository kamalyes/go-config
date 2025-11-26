/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-13 10:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-13 10:00:00
 * @FilePath: \go-config\pkg\request\request.go
 * @Description: 通用请求结构模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package request

// BaseRequest 通用请求基础结构
type BaseRequest struct {
	Timestamp     string `json:"timestamp" header:"X-Timestamp"`       // 时间戳
	Signature     string `json:"signature" header:"X-Signature"`       // 签名
	TraceID       string `json:"traceId" header:"X-Trace-Id"`          // 链路追踪ID
	RequestID     string `json:"requestId" header:"X-Request-Id"`      // 请求ID
	Authorization string `json:"authorization" header:"Authorization"` // 授权信息
	DeviceID      string `json:"deviceId" header:"X-Device-Id"`        // 设备ID
	AppVersion    string `json:"appVersion" header:"X-App-Version"`    // 应用版本
	Platform      string `json:"platform" header:"X-Platform"`         // 平台
}

// Headers 请求头常量
type Headers struct {
	Timestamp      string // X-Timestamp
	Signature      string // X-Signature
	TraceID        string // X-Trace-Id
	RequestID      string // X-Request-Id
	Authorization  string // Authorization
	DeviceID       string // X-Device-Id
	AppVersion     string // X-App-Version
	Platform       string // X-Platform
	UserAgent      string // User-Agent
	ContentType    string // Content-Type
	Accept         string // Accept
	AcceptLanguage string // Accept-Language
}

// DefaultHeaders 返回默认请求头常量
func DefaultHeaders() *Headers {
	return &Headers{
		Timestamp:      "X-Timestamp",
		Signature:      "X-Signature",
		TraceID:        "X-Trace-Id",
		RequestID:      "X-Request-Id",
		Authorization:  "Authorization",
		DeviceID:       "X-Device-Id",
		AppVersion:     "X-App-Version",
		Platform:       "X-Platform",
		UserAgent:      "User-Agent",
		ContentType:    "Content-Type",
		Accept:         "Accept",
		AcceptLanguage: "Accept-Language",
	}
}

// GetHeaderName 根据字段类型获取头部名称
func (h *Headers) GetHeaderName(fieldType string) string {
	switch fieldType {
	case "timestamp":
		return h.Timestamp
	case "signature":
		return h.Signature
	case "traceId":
		return h.TraceID
	case "requestId":
		return h.RequestID
	case "authorization":
		return h.Authorization
	case "deviceId":
		return h.DeviceID
	case "appVersion":
		return h.AppVersion
	case "platform":
		return h.Platform
	case "userAgent":
		return h.UserAgent
	case "contentType":
		return h.ContentType
	case "accept":
		return h.Accept
	case "acceptLanguage":
		return h.AcceptLanguage
	default:
		return ""
	}
}

// ExtendedRequest 扩展请求结构（可被继承）
type ExtendedRequest struct {
	BaseRequest
	Nonce     string            `json:"nonce" header:"X-Nonce"`        // 随机数
	ClientIP  string            `json:"clientIp" header:"X-Client-IP"` // 客户端IP
	UserAgent string            `json:"userAgent" header:"User-Agent"` // 用户代理
	Referer   string            `json:"referer" header:"Referer"`      // 引用页面
	Headers   map[string]string `json:"headers"`                       // 自定义头部
}

// WithTimestamp 设置时间戳
func (r *BaseRequest) WithTimestamp(timestamp string) *BaseRequest {
	r.Timestamp = timestamp
	return r
}

// WithSignature 设置签名
func (r *BaseRequest) WithSignature(signature string) *BaseRequest {
	r.Signature = signature
	return r
}

// WithTraceID 设置链路追踪ID
func (r *BaseRequest) WithTraceID(traceID string) *BaseRequest {
	r.TraceID = traceID
	return r
}

// WithRequestID 设置请求ID
func (r *BaseRequest) WithRequestID(requestID string) *BaseRequest {
	r.RequestID = requestID
	return r
}

// WithAuthorization 设置授权信息
func (r *BaseRequest) WithAuthorization(authorization string) *BaseRequest {
	r.Authorization = authorization
	return r
}

// WithDeviceID 设置设备ID
func (r *BaseRequest) WithDeviceID(deviceID string) *BaseRequest {
	r.DeviceID = deviceID
	return r
}

// WithAppVersion 设置应用版本
func (r *BaseRequest) WithAppVersion(appVersion string) *BaseRequest {
	r.AppVersion = appVersion
	return r
}

// WithPlatform 设置平台
func (r *BaseRequest) WithPlatform(platform string) *BaseRequest {
	r.Platform = platform
	return r
}

// Validate 基础验证
func (r *BaseRequest) Validate() error {
	// 可以在这里添加基础验证逻辑
	return nil
}

// IsEmpty 检查是否为空请求
func (r *BaseRequest) IsEmpty() bool {
	return r.Timestamp == "" && r.Signature == "" && r.RequestID == "" && r.TraceID == ""
}

// Clone 克隆请求
func (r *BaseRequest) Clone() *BaseRequest {
	return &BaseRequest{
		Timestamp:     r.Timestamp,
		Signature:     r.Signature,
		TraceID:       r.TraceID,
		RequestID:     r.RequestID,
		Authorization: r.Authorization,
		DeviceID:      r.DeviceID,
		AppVersion:    r.AppVersion,
		Platform:      r.Platform,
	}
}
