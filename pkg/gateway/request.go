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
	AuthPayloadSources   []common.AttributeSource `mapstructure:"auth-payload-sources" yaml:"auth-payload-sources" json:"authPayloadSources"`     // Auth-Payload 提取来源
	JtiSources           []common.AttributeSource `mapstructure:"jti-sources" yaml:"jti-sources" json:"jtiSources"`                               // Jti 提取来源
	FamilyIdSources      []common.AttributeSource `mapstructure:"family-id-sources" yaml:"family-id-sources" json:"familyIdSources"`              // FamilyId 提取来源
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
	ForwardedHostSources  []common.AttributeSource `mapstructure:"forwarded-host-sources" yaml:"forwarded-host-sources" json:"forwardedHostSources"`    // X-Forwarded-Host 提取来源

	// 用户上下文相关头部
	ClientIDSources   []common.AttributeSource `mapstructure:"client-id-sources" yaml:"client-id-sources" json:"clientIdSources"`       // X-Client-ID 提取来源
	UserIDSources     []common.AttributeSource `mapstructure:"user-id-sources" yaml:"user-id-sources" json:"userIdSources"`             // X-User-ID 提取来源
	UserTypeSources   []common.AttributeSource `mapstructure:"user-type-sources" yaml:"user-type-sources" json:"userTypeSources"`       // X-User-Type 提取来源
	TenantIDSources   []common.AttributeSource `mapstructure:"tenant-id-sources" yaml:"tenant-id-sources" json:"tenantIdSources"`       // X-Tenant-ID 提取来源
	TenantCodeSources []common.AttributeSource `mapstructure:"tenant-code-sources" yaml:"tenant-code-sources" json:"tenantCodeSources"` // X-Tenant-Code 提取来源
	SessionIDSources  []common.AttributeSource `mapstructure:"session-id-sources" yaml:"session-id-sources" json:"sessionIdSources"`    // X-Session-ID 提取来源
	TimezoneSources   []common.AttributeSource `mapstructure:"timezone-sources" yaml:"timezone-sources" json:"timezoneSources"`         // X-Timezone 提取来源
	IDSources         []common.AttributeSource `mapstructure:"id-sources" yaml:"id-sources" json:"idSources"`                           // X-ID 提取来源
	DomainSources     []common.AttributeSource `mapstructure:"domain-sources" yaml:"domain-sources" json:"domainSources"`               // X-Domain 提取来源
	RoleCodeSources   []common.AttributeSource `mapstructure:"role-code-sources" yaml:"role-code-sources" json:"roleCodeSources"`       // X-Role-Code 提取来源
	PushTokenSources  []common.AttributeSource `mapstructure:"push-token-sources" yaml:"push-token-sources" json:"pushTokenSources"`    // X-Push-Token 提取来源
	TokenSources      []common.AttributeSource `mapstructure:"token-sources" yaml:"token-sources" json:"tokenSources"`                  // X-Token 提取来源

	// 设备和应用相关头部
	DeviceIDSources     []common.AttributeSource `mapstructure:"device-id-sources" yaml:"device-id-sources" json:"deviceIdSources"`             // X-Device-Id / X-Device-ID 提取来源
	AppIDSources        []common.AttributeSource `mapstructure:"app-id-sources" yaml:"app-id-sources" json:"appIdSources"`                      // X-App-Id 提取来源
	AppVersionSources   []common.AttributeSource `mapstructure:"app-version-sources" yaml:"app-version-sources" json:"appVersionSources"`       // X-App-Version 提取来源
	PlatformSources     []common.AttributeSource `mapstructure:"platform-sources" yaml:"platform-sources" json:"platformSources"`               // X-Platform 提取来源
	PlatformIDSources   []common.AttributeSource `mapstructure:"platform-id-sources" yaml:"platform-id-sources" json:"platformIDSources"`       // X-Platform-ID 提取来源
	PlatformCodeSources []common.AttributeSource `mapstructure:"platform-code-sources" yaml:"platform-code-sources" json:"platformCodeSources"` // X-Platform-Code 提取来源
	RegionIDSources     []common.AttributeSource `mapstructure:"region-id-sources" yaml:"region-id-sources" json:"regionIDSources"`             // X-Region-ID 提取来源
	RegionCodeSources   []common.AttributeSource `mapstructure:"region-code-sources" yaml:"region-code-sources" json:"regionCodeSources"`       // X-Region-Code 提取来源
	AgentLineIDSources  []common.AttributeSource `mapstructure:"agent-line-id-sources" yaml:"agent-line-id-sources" json:"agentLineIDSources"`  // X-Agent-Line 提取来源
	TimestampSources    []common.AttributeSource `mapstructure:"timestamp-sources" yaml:"timestamp-sources" json:"timestampSources"`            // X-Timestamp 提取来源
	SignatureSources    []common.AttributeSource `mapstructure:"signature-sources" yaml:"signature-sources" json:"signatureSources"`            // X-Signature 提取来源
	NonceSources        []common.AttributeSource `mapstructure:"nonce-sources" yaml:"nonce-sources" json:"nonceSources"`                        // X-Nonce 提取来源
	AccessKeySources    []common.AttributeSource `mapstructure:"access-key-sources" yaml:"access-key-sources" json:"accessKeySources"`          // X-Access-Key 提取来源

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
		AuthPayloadSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Auth-Payload"},
		},
		JtiSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Jti"},
		},
		FamilyIdSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Family-Id"},
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
		ForwardedHostSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Forwarded-Host"},
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
		TenantCodeSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Tenant-Code"},
			{Type: common.SourceTypeQuery, Key: "tenant_code"},
			{Type: common.SourceTypeQuery, Key: "tenantCode"},
			{Type: common.SourceTypeCookie, Key: "tenant_code"},
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
		IDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-ID"},
			{Type: common.SourceTypeQuery, Key: "id"},
			{Type: common.SourceTypeCookie, Key: "id"},
		},
		DomainSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Domain"},
			{Type: common.SourceTypeQuery, Key: "domain"},
			{Type: common.SourceTypeCookie, Key: "domain"},
		},
		RoleCodeSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Role-Code"},
			{Type: common.SourceTypeQuery, Key: "role_code"},
			{Type: common.SourceTypeQuery, Key: "roleCode"},
			{Type: common.SourceTypeCookie, Key: "role_code"},
		},
		AgentLineIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Agent-Line-ID"},
			{Type: common.SourceTypeQuery, Key: "agent_line_id"},
			{Type: common.SourceTypeQuery, Key: "agentLineId"},
			{Type: common.SourceTypeCookie, Key: "agent_line_id"},
		},
		PushTokenSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Push-Token"},
			{Type: common.SourceTypeQuery, Key: "push_token"},
			{Type: common.SourceTypeQuery, Key: "pushToken"},
			{Type: common.SourceTypeCookie, Key: "push_token"},
		},
		TokenSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Token"},
			{Type: common.SourceTypeQuery, Key: "token"},
			{Type: common.SourceTypeQuery, Key: "token"},
			{Type: common.SourceTypeCookie, Key: "token"},
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
		PlatformIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Platform-Id"},
			{Type: common.SourceTypeQuery, Key: "platform_id"},
			{Type: common.SourceTypeQuery, Key: "platformId"},
			{Type: common.SourceTypeCookie, Key: "platform_id"},
		},
		PlatformCodeSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Platform-Code"},
			{Type: common.SourceTypeQuery, Key: "platform_code"},
			{Type: common.SourceTypeQuery, Key: "platformCode"},
			{Type: common.SourceTypeCookie, Key: "platform_code"},
		},
		RegionIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Region-ID"},
			{Type: common.SourceTypeQuery, Key: "region_id"},
			{Type: common.SourceTypeQuery, Key: "regionId"},
			{Type: common.SourceTypeCookie, Key: "region_id"},
		},
		RegionCodeSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Region-Code"},
			{Type: common.SourceTypeQuery, Key: "region_code"},
			{Type: common.SourceTypeQuery, Key: "regionCode"},
			{Type: common.SourceTypeCookie, Key: "region_code"},
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

// FieldSourceKeys 存储单个字段的所有来源 key（header / query / cookie）
type FieldSourceKeys struct {
	Header string
	Query  string
	Cookie string
}

// SourceKeys 存储每个字段对应的各来源类型的首选 Key
// 用于外部包动态获取 header / query / cookie key，而非依赖硬编码常量
type SourceKeys struct {
	ID             FieldSourceKeys
	TraceID        FieldSourceKeys
	RequestID      FieldSourceKeys
	Authorization  FieldSourceKeys
	AuthPayload    FieldSourceKeys
	Jti            FieldSourceKeys
	FamilyId       FieldSourceKeys
	UserAgent      FieldSourceKeys
	ClientID       FieldSourceKeys
	UserID         FieldSourceKeys
	UserType       FieldSourceKeys
	Domain         FieldSourceKeys
	RoleCode       FieldSourceKeys
	TenantID       FieldSourceKeys
	TenantCode     FieldSourceKeys
	SessionID      FieldSourceKeys
	Timezone       FieldSourceKeys
	RealIP         FieldSourceKeys
	ForwardedFor   FieldSourceKeys
	ForwardedProto FieldSourceKeys
	ForwardedHost  FieldSourceKeys
	DeviceID       FieldSourceKeys
	AppID          FieldSourceKeys
	AppVersion     FieldSourceKeys
	PlatformID     FieldSourceKeys
	PlatformCode   FieldSourceKeys
	RegionID       FieldSourceKeys
	RegionCode     FieldSourceKeys
	AgentLineID    FieldSourceKeys
	Timestamp      FieldSourceKeys
	Signature      FieldSourceKeys
	Nonce          FieldSourceKeys
	AccessKey      FieldSourceKeys
	PushToken      FieldSourceKeys
	Token          FieldSourceKeys
	Origin         FieldSourceKeys
	CSRFToken      FieldSourceKeys
}

// firstSourceKey 返回 sources 中第一个指定类型来源的 Key
func firstSourceKey(sources []common.AttributeSource, sourceType common.AttributeSourceType) string {
	for _, s := range sources {
		if s.Type == sourceType {
			return s.Key
		}
	}
	return ""
}

// extractFieldSourceKeys 从 sources 中提取所有来源类型的首选 Key
func extractFieldSourceKeys(sources []common.AttributeSource) FieldSourceKeys {
	return FieldSourceKeys{
		Header: firstSourceKey(sources, common.SourceTypeHeader),
		Query:  firstSourceKey(sources, common.SourceTypeQuery),
		Cookie: firstSourceKey(sources, common.SourceTypeCookie),
	}
}

// GetSourceKeys 从当前 RequestContext 配置中提取所有字段的各来源首选 key
// 当配置变更时，调用此方法可获取最新的 key
func (c *RequestContext) GetSourceKeys() *SourceKeys {
	return &SourceKeys{
		ID:             extractFieldSourceKeys(c.IDSources),
		TraceID:        extractFieldSourceKeys(c.TraceIDSources),
		RequestID:      extractFieldSourceKeys(c.RequestIDSources),
		Authorization:  extractFieldSourceKeys(c.AuthorizationSources),
		AuthPayload:    extractFieldSourceKeys(c.AuthPayloadSources),
		Jti:            extractFieldSourceKeys(c.JtiSources),
		FamilyId:       extractFieldSourceKeys(c.FamilyIdSources),
		UserAgent:      extractFieldSourceKeys(c.UserAgentSources),
		ClientID:       extractFieldSourceKeys(c.ClientIDSources),
		UserID:         extractFieldSourceKeys(c.UserIDSources),
		UserType:       extractFieldSourceKeys(c.UserTypeSources),
		Domain:         extractFieldSourceKeys(c.DomainSources),
		RoleCode:       extractFieldSourceKeys(c.RoleCodeSources),
		TenantID:       extractFieldSourceKeys(c.TenantIDSources),
		TenantCode:     extractFieldSourceKeys(c.TenantCodeSources),
		SessionID:      extractFieldSourceKeys(c.SessionIDSources),
		Timezone:       extractFieldSourceKeys(c.TimezoneSources),
		RealIP:         extractFieldSourceKeys(c.RealIPSources),
		ForwardedFor:   extractFieldSourceKeys(c.ForwardedForSources),
		ForwardedProto: extractFieldSourceKeys(c.ForwardedProtoSources),
		ForwardedHost:  extractFieldSourceKeys(c.ForwardedHostSources),
		DeviceID:       extractFieldSourceKeys(c.DeviceIDSources),
		AppID:          extractFieldSourceKeys(c.AppIDSources),
		AppVersion:     extractFieldSourceKeys(c.AppVersionSources),
		PlatformID:     extractFieldSourceKeys(c.PlatformIDSources),
		PlatformCode:   extractFieldSourceKeys(c.PlatformCodeSources),
		RegionID:       extractFieldSourceKeys(c.RegionIDSources),
		RegionCode:     extractFieldSourceKeys(c.RegionCodeSources),
		AgentLineID:    extractFieldSourceKeys(c.AgentLineIDSources),
		Timestamp:      extractFieldSourceKeys(c.TimestampSources),
		Signature:      extractFieldSourceKeys(c.SignatureSources),
		Nonce:          extractFieldSourceKeys(c.NonceSources),
		AccessKey:      extractFieldSourceKeys(c.AccessKeySources),
		PushToken:      extractFieldSourceKeys(c.PushTokenSources),
		Token:          extractFieldSourceKeys(c.TokenSources),
		Origin:         extractFieldSourceKeys(c.OriginSources),
		CSRFToken:      extractFieldSourceKeys(c.CSRFTokenSources),
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
		c.AuthPayloadSources,
		c.JtiSources,
		c.FamilyIdSources,
		c.UserAgentSources,
		c.AcceptSources,
		c.CacheControlSources,
		c.ConnectionSources,
		c.RequestIDSources,
		c.TraceIDSources,
		c.RealIPSources,
		c.ForwardedForSources,
		c.ForwardedProtoSources,
		c.ForwardedHostSources,
		c.ClientIDSources,
		c.UserIDSources,
		c.UserTypeSources,
		c.TenantIDSources,
		c.TenantCodeSources,
		c.SessionIDSources,
		c.TimezoneSources,
		c.IDSources,
		c.DomainSources,
		c.RoleCodeSources,
		c.DeviceIDSources,
		c.AppIDSources,
		c.AppVersionSources,
		c.PlatformIDSources,
		c.PlatformCodeSources,
		c.RegionIDSources,
		c.RegionCodeSources,
		c.AgentLineIDSources,
		c.TimestampSources,
		c.SignatureSources,
		c.NonceSources,
		c.AccessKeySources,
		c.PushTokenSources,
		c.TokenSources,
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
