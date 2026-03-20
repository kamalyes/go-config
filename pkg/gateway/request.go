/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-03-20 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-03-21 00:16:15
 * @FilePath: \go-config\pkg\gateway\request.go
 * @Description: Gateway 统一请求字段提取配置
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package gateway

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-config/pkg/common"
	"github.com/kamalyes/go-toolbox/pkg/syncx"
)

// RequestContext Gateway 统一请求上下文提取配置
// 用于定义 RequestCommon 所需字段从哪些来源提取，供所有 middleware 共享
type RequestContext struct {
	// 标准请求头
	AuthorizationSources []common.AttributeSource `mapstructure:"authorization-sources" yaml:"authorization-sources" json:"authorizationSources"` // Authorization 提取来源
	UserAgentSources     []common.AttributeSource `mapstructure:"user-agent-sources" yaml:"user-agent-sources" json:"userAgentSources"`           // User-Agent 提取来源
	AcceptSources        []common.AttributeSource `mapstructure:"accept-sources" yaml:"accept-sources" json:"acceptSources"`                      // Accept 提取来源
	CacheControlSources  []common.AttributeSource `mapstructure:"cache-control-sources" yaml:"cache-control-sources" json:"cacheControlSources"`  // Cache-Control 提取来源
	ConnectionSources    []common.AttributeSource `mapstructure:"connection-sources" yaml:"connection-sources" json:"connectionSources"`          // Connection 提取来源

	// 自定义请求头
	RequestIDSources      []common.AttributeSource `mapstructure:"request-id-sources" yaml:"request-id-sources" json:"requestIdSources"`                // X-Request-Id 提取来源
	TraceIDSources        []common.AttributeSource `mapstructure:"trace-id-sources" yaml:"trace-id-sources" json:"traceIdSources"`                      // X-Trace-Id 提取来源
	RealIPSources         []common.AttributeSource `mapstructure:"real-ip-sources" yaml:"real-ip-sources" json:"realIpSources"`                         // X-Real-IP 提取来源
	ForwardedForSources   []common.AttributeSource `mapstructure:"forwarded-for-sources" yaml:"forwarded-for-sources" json:"forwardedForSources"`       // X-Forwarded-For 提取来源
	ForwardedProtoSources []common.AttributeSource `mapstructure:"forwarded-proto-sources" yaml:"forwarded-proto-sources" json:"forwardedProtoSources"` // X-Forwarded-Proto 提取来源

	// 用户上下文相关头部
	ClientIDSources  []common.AttributeSource `mapstructure:"client-id-sources" yaml:"client-id-sources" json:"clientIdSources"`    // X-Client-ID 提取来源
	UserIDSources    []common.AttributeSource `mapstructure:"user-id-sources" yaml:"user-id-sources" json:"userIdSources"`          // X-User-ID 提取来源
	UserTypeSources  []common.AttributeSource `mapstructure:"user-type-sources" yaml:"user-type-sources" json:"userTypeSources"`    // X-User-Type 提取来源
	TenantIDSources  []common.AttributeSource `mapstructure:"tenant-id-sources" yaml:"tenant-id-sources" json:"tenantIdSources"`    // X-Tenant-ID 提取来源
	SessionIDSources []common.AttributeSource `mapstructure:"session-id-sources" yaml:"session-id-sources" json:"sessionIdSources"` // X-Session-ID 提取来源
	TimezoneSources  []common.AttributeSource `mapstructure:"timezone-sources" yaml:"timezone-sources" json:"timezoneSources"`      // X-Timezone 提取来源

	// 设备和应用相关头部
	DeviceIDSources   []common.AttributeSource `mapstructure:"device-id-sources" yaml:"device-id-sources" json:"deviceIdSources"`       // X-Device-Id / X-Device-ID 提取来源
	AppIDSources      []common.AttributeSource `mapstructure:"app-id-sources" yaml:"app-id-sources" json:"appIdSources"`                // X-App-Id 提取来源
	AppVersionSources []common.AttributeSource `mapstructure:"app-version-sources" yaml:"app-version-sources" json:"appVersionSources"` // X-App-Version 提取来源
	PlatformSources   []common.AttributeSource `mapstructure:"platform-sources" yaml:"platform-sources" json:"platformSources"`         // X-Platform 提取来源
	TimestampSources  []common.AttributeSource `mapstructure:"timestamp-sources" yaml:"timestamp-sources" json:"timestampSources"`      // X-Timestamp 提取来源
	SignatureSources  []common.AttributeSource `mapstructure:"signature-sources" yaml:"signature-sources" json:"signatureSources"`      // X-Signature 提取来源
	NonceSources      []common.AttributeSource `mapstructure:"nonce-sources" yaml:"nonce-sources" json:"nonceSources"`                  // X-Nonce 提取来源
	AccessKeySources  []common.AttributeSource `mapstructure:"access-key-sources" yaml:"access-key-sources" json:"accessKeySources"`    // X-Access-Key 提取来源

	// 其他请求头
	OriginSources    []common.AttributeSource `mapstructure:"origin-sources" yaml:"origin-sources" json:"originSources"`            // Origin 提取来源
	CSRFTokenSources []common.AttributeSource `mapstructure:"csrf-token-sources" yaml:"csrf-token-sources" json:"csrfTokenSources"` // X-CSRF-Token 提取来源
}

// DefaultRequestContext 创建默认请求上下文提取配置
func DefaultRequestContext() *RequestContext {
	return &RequestContext{
		AuthorizationSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "Authorization"},
			{Type: common.SourceTypeQuery, Key: "authorization"},
			{Type: common.SourceTypeCookie, Key: "authorization"},
		},
		UserAgentSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "User-Agent"},
		},
		AcceptSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "Accept"},
		},
		CacheControlSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "Cache-Control"},
		},
		ConnectionSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "Connection"},
		},

		RequestIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Request-Id"},
			{Type: common.SourceTypeQuery, Key: "request_id"},
			{Type: common.SourceTypeQuery, Key: "requestId"},
			{Type: common.SourceTypeCookie, Key: "request_id"},
		},
		TraceIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Trace-Id"},
			{Type: common.SourceTypeQuery, Key: "trace_id"},
			{Type: common.SourceTypeQuery, Key: "traceId"},
			{Type: common.SourceTypeCookie, Key: "trace_id"},
		},
		RealIPSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Real-IP"},
		},
		ForwardedForSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Forwarded-For"},
		},
		ForwardedProtoSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Forwarded-Proto"},
		},

		ClientIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Client-ID"},
			{Type: common.SourceTypeQuery, Key: "client_id"},
			{Type: common.SourceTypeQuery, Key: "clientId"},
			{Type: common.SourceTypeCookie, Key: "client_id"},
		},
		UserIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-User-ID"},
			{Type: common.SourceTypeQuery, Key: "user_id"},
			{Type: common.SourceTypeQuery, Key: "userId"},
			{Type: common.SourceTypeCookie, Key: "user_id"},
		},
		UserTypeSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-User-Type"},
			{Type: common.SourceTypeQuery, Key: "user_type"},
			{Type: common.SourceTypeQuery, Key: "userType"},
			{Type: common.SourceTypeCookie, Key: "user_type"},
		},
		TenantIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Tenant-ID"},
			{Type: common.SourceTypeQuery, Key: "tenant_id"},
			{Type: common.SourceTypeQuery, Key: "tenantId"},
			{Type: common.SourceTypeCookie, Key: "tenant_id"},
		},
		SessionIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Session-ID"},
			{Type: common.SourceTypeQuery, Key: "session_id"},
			{Type: common.SourceTypeQuery, Key: "sessionId"},
			{Type: common.SourceTypeCookie, Key: "session_id"},
		},
		TimezoneSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Timezone"},
			{Type: common.SourceTypeQuery, Key: "timezone"},
			{Type: common.SourceTypeCookie, Key: "timezone"},
		},

		DeviceIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Device-Id"},
			{Type: common.SourceTypeHeader, Key: "X-Device-ID"},
			{Type: common.SourceTypeQuery, Key: "device_id"},
			{Type: common.SourceTypeQuery, Key: "deviceId"},
			{Type: common.SourceTypeCookie, Key: "device_id"},
		},
		AppIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-App-Id"},
			{Type: common.SourceTypeQuery, Key: "app_id"},
			{Type: common.SourceTypeQuery, Key: "appId"},
			{Type: common.SourceTypeCookie, Key: "app_id"},
		},
		AppVersionSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-App-Version"},
			{Type: common.SourceTypeQuery, Key: "app_version"},
			{Type: common.SourceTypeQuery, Key: "appVersion"},
			{Type: common.SourceTypeCookie, Key: "app_version"},
		},
		PlatformSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Platform"},
			{Type: common.SourceTypeQuery, Key: "platform"},
			{Type: common.SourceTypeCookie, Key: "platform"},
		},
		TimestampSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Timestamp"},
			{Type: common.SourceTypeQuery, Key: "timestamp"},
			{Type: common.SourceTypeCookie, Key: "timestamp"},
		},
		SignatureSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Signature"},
			{Type: common.SourceTypeQuery, Key: "signature"},
			{Type: common.SourceTypeCookie, Key: "signature"},
		},
		NonceSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Nonce"},
			{Type: common.SourceTypeQuery, Key: "nonce"},
			{Type: common.SourceTypeCookie, Key: "nonce"},
		},
		AccessKeySources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Access-Key"},
			{Type: common.SourceTypeQuery, Key: "access_key"},
			{Type: common.SourceTypeQuery, Key: "accessKey"},
			{Type: common.SourceTypeCookie, Key: "access_key"},
		},
		OriginSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "Origin"},
		},
		CSRFTokenSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-CSRF-Token"},
		},
	}
}

// Get 返回配置接口
func (c *RequestContext) Get() any {
	return c
}

// Set 设置配置
func (c *RequestContext) Set(data any) {
	if cfg, ok := data.(*RequestContext); ok {
		*c = *cfg
	}
}

// Clone 返回副本
func (c *RequestContext) Clone() internal.Configurable {
	var cloned RequestContext
	if err := syncx.DeepCopy(&cloned, c); err != nil {
		return &RequestContext{}
	}
	return &cloned
}

// Validate 验证配置
func (c *RequestContext) Validate() error {
	if err := internal.ValidateStruct(c); err != nil {
		return err
	}

	sourceGroups := [][]common.AttributeSource{
		c.AuthorizationSources,
		c.UserAgentSources,
		c.AcceptSources,
		c.CacheControlSources,
		c.ConnectionSources,
		c.RequestIDSources,
		c.TraceIDSources,
		c.RealIPSources,
		c.ForwardedForSources,
		c.ForwardedProtoSources,
		c.ClientIDSources,
		c.UserIDSources,
		c.UserTypeSources,
		c.TenantIDSources,
		c.SessionIDSources,
		c.TimezoneSources,
		c.DeviceIDSources,
		c.AppIDSources,
		c.AppVersionSources,
		c.PlatformSources,
		c.TimestampSources,
		c.SignatureSources,
		c.NonceSources,
		c.AccessKeySources,
		c.OriginSources,
		c.CSRFTokenSources,
	}

	for _, group := range sourceGroups {
		for i := range group {
			if err := group[i].Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
