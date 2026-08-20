/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 11:30:00
 * @FilePath: \go-config\pkg\wsc\wsc_test.go
 * @Description: WebSocket配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package wsc

import (
	"strings"
	"testing"
	"time"

	"github.com/kamalyes/go-config/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestWSC_Clone(t *testing.T) {
	original := &WSC{
		Enabled:               true,
		Network:               "tcp",
		NodeIP:                "127.0.0.1",
		NodePort:              8080,
		Path:                  "/ws",
		HeartbeatInterval:     30 * time.Second,
		ClientTimeout:         300 * time.Second,
		MessageBufferSize:     1024,
		MaxPendingQueueSize:   10000,
		WebSocketOrigins:      []string{"http://localhost"},
		ReadTimeout:           60 * time.Second,
		WriteTimeout:          60 * time.Second,
		IdleTimeout:           300 * time.Second,
		MaxMessageSize:        1048576,
		MinRecTime:            2 * time.Second,
		MaxRecTime:            30 * time.Second,
		RecFactor:             1.5,
		AutoReconnect:         true,
		AckTimeout:            5 * time.Second,
		AckMaxRetries:         3,
		EnableAck:             true,
		MessageRecordTTL:      24 * time.Hour,
		RecordCleanupInterval: 30 * time.Minute,
		RetryPolicy: &RetryPolicy{
			MaxRetries:    3,
			BaseDelay:     1 * time.Second,
			MaxDelay:      30 * time.Second,
			BackoffFactor: 2.0,
			Jitter:        true,
		},
		SSEHeartbeat:     30,
		SSETimeout:       300,
		SSEMessageBuffer: 100,
	}

	cloned := original.Clone().(*WSC)

	// 验证克隆后的值相等
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.Network, cloned.Network)
	assert.Equal(t, original.NodeIP, cloned.NodeIP)
	assert.Equal(t, original.NodePort, cloned.NodePort)
	assert.Equal(t, original.Path, cloned.Path)
	assert.Equal(t, original.HeartbeatInterval, cloned.HeartbeatInterval)
	assert.Equal(t, original.ClientTimeout, cloned.ClientTimeout)
	assert.Equal(t, original.MessageBufferSize, cloned.MessageBufferSize)
	assert.Equal(t, original.AutoReconnect, cloned.AutoReconnect)
	assert.Equal(t, original.EnableAck, cloned.EnableAck)

	// 验证 slice 深拷贝
	assert.Equal(t, original.WebSocketOrigins, cloned.WebSocketOrigins)

	// 修改原始对象不应影响克隆对象
	original.NodePort = 9090
	original.Path = "/websocket"
	original.WebSocketOrigins[0] = "http://newhost"

	assert.NotEqual(t, original.NodePort, cloned.NodePort)
	assert.NotEqual(t, original.Path, cloned.Path)
	assert.NotEqual(t, original.WebSocketOrigins[0], cloned.WebSocketOrigins[0])
}

// TestDefaultClientAttributesHasNamespaceGroup 验证默认配置包含 NamespaceSources/GroupIDSources
func TestDefaultClientAttributesHasNamespaceGroup(t *testing.T) {
	ca := DefaultClientAttributes()
	assert.NotEmpty(t, ca.NamespaceSources, "默认配置应包含 NamespaceSources")
	assert.NotEmpty(t, ca.GroupIDSources, "默认配置应包含 GroupIDSources")

	// 验证默认来源包含 query 和 header
	hasNamespaceQuery, hasNamespaceHeader := false, false
	for _, src := range ca.NamespaceSources {
		if src.Type == common.SourceTypeQuery && src.Key == "namespace" {
			hasNamespaceQuery = true
		}
		if src.Type == common.SourceTypeHeader && src.Key == "X-Namespace" {
			hasNamespaceHeader = true
		}
	}
	assert.True(t, hasNamespaceQuery, "NamespaceSources 应包含 query:namespace")
	assert.True(t, hasNamespaceHeader, "NamespaceSources 应包含 header:X-Namespace")

	hasGroupQuery, hasGroupHeader := false, false
	for _, src := range ca.GroupIDSources {
		if src.Type == common.SourceTypeQuery && src.Key == "group_id" {
			hasGroupQuery = true
		}
		if src.Type == common.SourceTypeHeader && src.Key == "X-Group-ID" {
			hasGroupHeader = true
		}
	}
	assert.True(t, hasGroupQuery, "GroupIDSources 应包含 query:group_id")
	assert.True(t, hasGroupHeader, "GroupIDSources 应包含 header:X-Group-ID")
}

// TestClientAttributesValidateNamespaceGroup 验证 NamespaceSources/GroupIDSources 的 Validate
func TestClientAttributesValidateNamespaceGroup(t *testing.T) {
	// 合法配置
	ca := &ClientAttributes{
		NamespaceSources: []common.AttributeSource{
			{Type: common.SourceTypeQuery, Key: "namespace"},
		},
		GroupIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "X-Group-ID"},
		},
	}
	assert.NoError(t, ca.Validate())

	// 非法 Namespace 来源（空 Key）
	invalid := &ClientAttributes{
		NamespaceSources: []common.AttributeSource{
			{Type: common.SourceTypeQuery, Key: ""},
		},
	}
	err := invalid.Validate()
	assert.Error(t, err, "空 Key 的 NamespaceSources 应验证失败")

	// 非法 GroupID 来源（空 Key）
	invalid2 := &ClientAttributes{
		GroupIDSources: []common.AttributeSource{
			{Type: common.SourceTypeQuery, Key: ""},
		},
	}
	err2 := invalid2.Validate()
	assert.Error(t, err2, "空 Key 的 GroupIDSources 应验证失败")
}

// TestResolveTokens_LegacyConfig 旧单 ConnectionToken 自动包装为单 Default
func TestResolveTokens_LegacyConfig(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.SigningKey = "legacy-secret"
	cfg.Algorithm = "HS256"

	tokens, defaultID, err := cfg.ResolveTokens()
	if err != nil {
		t.Fatalf("legacy resolve failed: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token set, got %d", len(tokens))
	}
	set, ok := tokens[defaultTokenAppID]
	if !ok {
		t.Fatalf("default app id %q not in tokens", defaultTokenAppID)
	}
	if set.GetSigningKey() != "legacy-secret" {
		t.Errorf("signing key = %q, want legacy-secret", set.GetSigningKey())
	}
	if defaultID != defaultTokenAppID {
		t.Errorf("default id = %q, want %q", defaultID, defaultTokenAppID)
	}
}

// TestResolveTokens_NewMultiApp 新版多 appid 正常解析
func TestResolveTokens_NewMultiApp(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.DefaultAppID = "default"
	cfg.Tokens = map[string]*ConnectionTokenSet{
		"default": {SigningKey: "default-secret", Algorithm: "HS256"},
		"app-A":   {SigningKey: "appA-secret", Algorithm: "HS512"},
		"app-B":   {SigningKey: "appB-secret", Algorithm: "HS256"},
	}

	tokens, defaultID, err := cfg.ResolveTokens()
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 token sets, got %d", len(tokens))
	}
	if defaultID != "default" {
		t.Errorf("default id = %q, want default", defaultID)
	}
	// 验证 AppID 回填
	for appID, set := range tokens {
		if set.AppID != appID {
			t.Errorf("set.AppID = %q, want %q", set.AppID, appID)
		}
	}
}

// TestResolveTokens_BothLegacyAndTokens 同时配旧字段+tokens 时 tokens 优先
func TestResolveTokens_BothLegacyAndTokens(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.SigningKey = "legacy-secret" // 旧字段
	cfg.Tokens = map[string]*ConnectionTokenSet{
		"default": {SigningKey: "new-secret", Algorithm: "HS256"},
	}

	tokens, _, err := cfg.ResolveTokens()
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	// tokens 优先，旧字段被忽略
	set := tokens["default"]
	if set.GetSigningKey() != "new-secret" {
		t.Errorf("signing key = %q, want new-secret (tokens should take priority)", set.GetSigningKey())
	}
}

// TestValidateMultiAppID_EmptySigningKey 某 appid signing-key 为空报错
func TestValidateMultiAppID_EmptySigningKey(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.DefaultAppID = "default"
	cfg.Tokens = map[string]*ConnectionTokenSet{
		"default": {SigningKey: "valid", Algorithm: "HS256"},
		"app-A":   {SigningKey: "", Algorithm: "HS256"}, // 空 key
	}

	err := cfg.ValidateMultiAppID()
	if err == nil {
		t.Fatal("expected error for empty signing key, got nil")
	}
	if !strings.Contains(err.Error(), "signing-key is required") {
		t.Errorf("error = %q, want contains 'signing-key is required'", err.Error())
	}
	if !strings.Contains(err.Error(), "app-A") {
		t.Errorf("error = %q, want contains app-A", err.Error())
	}
}

// TestValidateMultiAppID_DefaultMissing DefaultAppID 不在 tokens map 报错
func TestValidateMultiAppID_DefaultMissing(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.DefaultAppID = "nonexistent"
	cfg.Tokens = map[string]*ConnectionTokenSet{
		"default": {SigningKey: "secret", Algorithm: "HS256"},
	}

	err := cfg.ValidateMultiAppID()
	if err == nil {
		t.Fatal("expected error for missing default app id, got nil")
	}
	if !strings.Contains(err.Error(), "not found in tokens") {
		t.Errorf("error = %q, want contains 'not found in tokens'", err.Error())
	}
}

// TestValidateMultiAppID_DuplicateIssuerKey 跨 appid (Issuer,SigningKey) 重复报错
func TestValidateMultiAppID_DuplicateIssuerKey(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.DefaultAppID = "default"
	cfg.Tokens = map[string]*ConnectionTokenSet{
		"default": {SigningKey: "same-secret", Algorithm: "HS256", Issuer: "same-issuer"},
		"app-A":   {SigningKey: "same-secret", Algorithm: "HS256", Issuer: "same-issuer"}, // 重复
	}

	err := cfg.ValidateMultiAppID()
	if err == nil {
		t.Fatal("expected error for duplicate issuer+key, got nil")
	}
	if !strings.Contains(err.Error(), "duplicate (issuer, signing-key)") {
		t.Errorf("error = %q, want contains 'duplicate (issuer, signing-key)'", err.Error())
	}
}

// TestValidateMultiAppID_InvalidAlgorithm 非法 algorithm 报错
func TestValidateMultiAppID_InvalidAlgorithm(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.DefaultAppID = "default"
	cfg.Tokens = map[string]*ConnectionTokenSet{
		"default": {SigningKey: "secret", Algorithm: "RS256"}, // 非法算法
	}

	err := cfg.ValidateMultiAppID()
	if err == nil {
		t.Fatal("expected error for invalid algorithm, got nil")
	}
	if !strings.Contains(err.Error(), "invalid algorithm") {
		t.Errorf("error = %q, want contains 'invalid algorithm'", err.Error())
	}
}

// TestValidateMultiAppID_DuplicateRedisPrefix 启用 Redis 时前缀重复报错
func TestValidateMultiAppID_DuplicateRedisPrefix(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.DefaultAppID = "default"
	cfg.Tokens = map[string]*ConnectionTokenSet{
		"default": {SigningKey: "s1", Algorithm: "HS256", UseRedis: true, RedisKeyPrefix: "wsc:dup:"},
		"app-A":   {SigningKey: "s2", Algorithm: "HS256", UseRedis: true, RedisKeyPrefix: "wsc:dup:"}, // 重复前缀
	}

	err := cfg.ValidateMultiAppID()
	if err == nil {
		t.Fatal("expected error for duplicate redis prefix, got nil")
	}
	if !strings.Contains(err.Error(), "duplicate redis-key-prefix") {
		t.Errorf("error = %q, want contains 'duplicate redis-key-prefix'", err.Error())
	}
}

// TestConnectionTokenSet_Getters 每套 set 的 Getter 方法返回正确值
func TestConnectionTokenSet_Getters(t *testing.T) {
	set := &ConnectionTokenSet{
		AppID:          "app-X",
		SigningKey:     "key-X",
		Issuer:         "issuer-X",
		Audience:       "aud-X",
		Algorithm:      "HS384",
		ExpiresTime:    30 * time.Minute,
		UseRedis:       true,
		RedisKeyPrefix: "wsc:appX:",
		TokenSource:    "header",
		TokenParamName: "x-token",
	}

	if set.GetAppID() != "app-X" {
		t.Errorf("GetAppID = %q", set.GetAppID())
	}
	if set.GetSigningKey() != "key-X" {
		t.Errorf("GetSigningKey = %q", set.GetSigningKey())
	}
	if set.GetIssuer() != "issuer-X" {
		t.Errorf("GetIssuer = %q", set.GetIssuer())
	}
	if set.GetAudience() != "aud-X" {
		t.Errorf("GetAudience = %q", set.GetAudience())
	}
	if set.GetAlgorithm() != "HS384" {
		t.Errorf("GetAlgorithm = %q", set.GetAlgorithm())
	}
	if set.GetExpiresTime() != 30*time.Minute {
		t.Errorf("GetExpiresTime = %v", set.GetExpiresTime())
	}
	if !set.IsRedisEnabled() {
		t.Error("IsRedisEnabled = false, want true")
	}
	if set.GetRedisKeyPrefix() != "wsc:appX:" {
		t.Errorf("GetRedisKeyPrefix = %q", set.GetRedisKeyPrefix())
	}
	if set.GetTokenSource() != "header" {
		t.Errorf("GetTokenSource = %q", set.GetTokenSource())
	}
	if set.GetTokenParamName() != "x-token" {
		t.Errorf("GetTokenParamName = %q", set.GetTokenParamName())
	}
}

// TestConnectionTokenSet_Defaults 空字段的默认值兜底
func TestConnectionTokenSet_Defaults(t *testing.T) {
	set := &ConnectionTokenSet{AppID: "app-Y", SigningKey: "key-Y"}

	if set.GetAlgorithm() != "HS256" {
		t.Errorf("default Algorithm = %q, want HS256", set.GetAlgorithm())
	}
	if set.GetExpiresTime() != 5*time.Minute {
		t.Errorf("default ExpiresTime = %v, want 5m", set.GetExpiresTime())
	}
	if set.GetTokenSource() != "query" {
		t.Errorf("default TokenSource = %q, want query", set.GetTokenSource())
	}
	if set.GetTokenParamName() != "token" {
		t.Errorf("default TokenParamName = %q, want token", set.GetTokenParamName())
	}
	// 空前缀时按 appid 自动生成
	if set.GetRedisKeyPrefix() != defaultConnTokenKeyPrefix+"app-Y:" {
		t.Errorf("default RedisKeyPrefix = %q, want %s", set.GetRedisKeyPrefix(), defaultConnTokenKeyPrefix+"app-Y:")
	}
}

// TestValidateMultiAppID_Disabled 未启用时不校验
func TestValidateMultiAppID_Disabled(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = false
	cfg.Tokens = map[string]*ConnectionTokenSet{
		"default": {SigningKey: "", Algorithm: "INVALID"}, // 非法配置
	}

	err := cfg.ValidateMultiAppID()
	if err != nil {
		t.Errorf("disabled config should not validate, got error: %v", err)
	}
}

// TestValidateMultiAppID_InvalidTokenSource 非法 token-source 报错
func TestValidateMultiAppID_InvalidTokenSource(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.DefaultAppID = "default"
	cfg.Tokens = map[string]*ConnectionTokenSet{
		"default": {SigningKey: "secret", Algorithm: "HS256", TokenSource: "cookie"},
	}

	err := cfg.ValidateMultiAppID()
	if err == nil {
		t.Fatal("expected error for invalid token source, got nil")
	}
	if !strings.Contains(err.Error(), "invalid token-source") {
		t.Errorf("error = %q, want contains 'invalid token-source'", err.Error())
	}
}

// TestValidateMultiAppID_LegacyConfigPass 旧单套配置通过校验
func TestValidateMultiAppID_LegacyConfigPass(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.SigningKey = "legacy-secret"
	cfg.Algorithm = "HS256"

	if err := cfg.ValidateMultiAppID(); err != nil {
		t.Errorf("legacy config should pass validation, got error: %v", err)
	}
}

// TestResolveTokens_NilConfig nil 配置返回错误
func TestResolveTokens_NilConfig(t *testing.T) {
	var cfg *ConnectionToken
	_, _, err := cfg.ResolveTokens()
	if err == nil {
		t.Fatal("expected error for nil config, got nil")
	}
}

// TestResolveTokens_NilSetInTokens tokens map 中某 set 为 nil 报错
func TestResolveTokens_NilSetInTokens(t *testing.T) {
	cfg := DefaultConnectionToken()
	cfg.Enabled = true
	cfg.Tokens = map[string]*ConnectionTokenSet{
		"default": nil,
	}

	_, _, err := cfg.ResolveTokens()
	if err == nil {
		t.Fatal("expected error for nil set, got nil")
	}
	if !strings.Contains(err.Error(), "is nil") {
		t.Errorf("error = %q, want contains 'is nil'", err.Error())
	}
}
