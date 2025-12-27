/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\security\security_test.go
 * @Description: 统一安全配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecurity_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "security", config.ModuleName)
	assert.NotNil(t, config.JWT)
	assert.NotNil(t, config.Auth)
	assert.NotNil(t, config.Protection)
	assert.NotNil(t, config.CSP)
}

func TestSecurity_JWT_Default(t *testing.T) {
	config := Default()
	assert.False(t, config.JWT.Enabled)
	assert.Equal(t, "jwt_secret_key_please_change_in_production", config.JWT.Secret)
	assert.Equal(t, 24, config.JWT.Expiry)
	assert.Equal(t, "go-rpc-gateway", config.JWT.Issuer)
	assert.Equal(t, "HS256", config.JWT.Algorithm)
}

func TestSecurity_Auth_Default(t *testing.T) {
	config := Default()
	assert.False(t, config.Auth.Enabled)
	assert.Equal(t, "bearer", config.Auth.Type)
	assert.Equal(t, "Authorization", config.Auth.HeaderName)
	assert.Equal(t, "Bearer ", config.Auth.TokenPrefix)
	assert.NotNil(t, config.Auth.Basic)
	assert.NotNil(t, config.Auth.Bearer)
	assert.NotNil(t, config.Auth.APIKey)
	assert.NotNil(t, config.Auth.Custom)
}

func TestSecurity_WithModuleName(t *testing.T) {
	config := Default()
	result := config.WithModuleName("custom-security")
	assert.Equal(t, "custom-security", result.ModuleName)
	assert.Equal(t, config, result)
}

func TestSecurity_WithJWT(t *testing.T) {
	config := Default()
	result := config.WithJWT(true, "secret-key", 48, "my-app", "HS512")
	assert.True(t, result.JWT.Enabled)
	assert.Equal(t, "secret-key", result.JWT.Secret)
	assert.Equal(t, 48, result.JWT.Expiry)
	assert.Equal(t, "my-app", result.JWT.Issuer)
	assert.Equal(t, "HS512", result.JWT.Algorithm)
	assert.Equal(t, config, result)
}

func TestSecurity_WithAuth(t *testing.T) {
	config := Default()
	result := config.WithAuth(true, "basic", "X-Auth", "Basic ")
	assert.True(t, result.Auth.Enabled)
	assert.Equal(t, "basic", result.Auth.Type)
	assert.Equal(t, "X-Auth", result.Auth.HeaderName)
	assert.Equal(t, "Basic ", result.Auth.TokenPrefix)
	assert.Equal(t, config, result)
}

func TestSecurity_EnableJWT(t *testing.T) {
	config := Default()
	result := config.EnableJWT()
	assert.True(t, result.JWT.Enabled)
	assert.Equal(t, config, result)
}

func TestSecurity_EnableAuth(t *testing.T) {
	config := Default()
	result := config.EnableAuth()
	assert.True(t, result.Auth.Enabled)
	assert.Equal(t, config, result)
}

func TestSecurity_Clone(t *testing.T) {
	config := Default()
	config.WithJWT(true, "secret", 24, "app", "HS256")

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*Security)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
}

func TestSecurity_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestSecurity_Set(t *testing.T) {
	config := Default()
	newConfig := &Security{
		ModuleName: "new-security",
	}

	config.Set(newConfig)
	assert.Equal(t, "new-security", config.ModuleName)
}

func TestSecurity_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestSecurity_Protection_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config.Protection.Swagger)
	assert.NotNil(t, config.Protection.PProf)
	assert.NotNil(t, config.Protection.Metrics)
	assert.NotNil(t, config.Protection.Health)
	assert.NotNil(t, config.Protection.API)

	// PProf 默认配置
	assert.False(t, config.Protection.PProf.Enabled)
	assert.True(t, config.Protection.PProf.AuthRequired)
	assert.Equal(t, "basic", config.Protection.PProf.AuthType)
	assert.True(t, config.Protection.PProf.RequireHTTPS)
	assert.Equal(t, "admin", config.Protection.PProf.Username)
}

func TestSecurity_CSP_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config.CSP)
	assert.False(t, config.CSP.Enabled)
	assert.Equal(t, "balanced", config.CSP.Mode)
	assert.Equal(t, "", config.CSP.Custom)
}

func TestSecurity_CSP_GetPolicy(t *testing.T) {
	tests := []struct {
		name     string
		enabled  bool
		mode     string
		custom   string
		expected string
	}{
		{
			name:     "Disabled CSP",
			enabled:  false,
			mode:     "strict",
			expected: "",
		},
		{
			name:     "Strict Mode",
			enabled:  true,
			mode:     "strict",
			expected: "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self' data:; font-src 'self'; connect-src 'self' ws: wss:; frame-src 'none'; object-src 'none'; base-uri 'self'; form-action 'self'; frame-ancestors 'none'",
		},
		{
			name:     "Development Mode",
			enabled:  true,
			mode:     "development",
			expected: "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval' https:; style-src 'self' 'unsafe-inline' https:; img-src 'self' data: blob: https:; font-src 'self' data: https:; connect-src 'self' ws: wss: http: https:; media-src 'self' https:; object-src 'none'; frame-src 'self'; base-uri 'self'; form-action 'self'",
		},
		{
			name:     "Relaxed Mode",
			enabled:  true,
			mode:     "relaxed",
			expected: "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline' https:; img-src 'self' data: blob: https:; font-src 'self' data: https:; connect-src 'self' ws: wss: http: https:; media-src 'self' https:; object-src 'none'; frame-src 'self' https:; base-uri 'self'; form-action 'self'",
		},
		{
			name:     "API Mode",
			enabled:  true,
			mode:     "api",
			expected: "default-src 'none'; frame-ancestors 'none'; base-uri 'none'; form-action 'none'",
		},
		{
			name:     "Balanced Mode (default)",
			enabled:  true,
			mode:     "balanced",
			expected: "default-src 'self'; script-src 'self' https:; style-src 'self' 'unsafe-inline' https:; img-src 'self' data: https:; font-src 'self' data: https:; connect-src 'self' ws: wss:; media-src 'self' https:; object-src 'none'; frame-src 'self'; base-uri 'self'; form-action 'self'; frame-ancestors 'self'",
		},
		{
			name:     "Custom Mode",
			enabled:  true,
			mode:     "custom",
			custom:   "default-src 'none'; script-src 'self'",
			expected: "default-src 'none'; script-src 'self'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csp := &CSP{
				Enabled: tt.enabled,
				Mode:    tt.mode,
				Custom:  tt.custom,
			}
			assert.Equal(t, tt.expected, csp.GetPolicy())
		})
	}
}

func TestSecurity_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("my-security").
		WithJWT(true, "my-secret", 72, "my-issuer", "HS512").
		WithAuth(true, "bearer", "Authorization", "Bearer ").
		EnableJWT().
		EnableAuth()

	assert.Equal(t, "my-security", config.ModuleName)
	assert.True(t, config.JWT.Enabled)
	assert.Equal(t, "my-secret", config.JWT.Secret)
	assert.Equal(t, 72, config.JWT.Expiry)
	assert.True(t, config.Auth.Enabled)
	assert.Equal(t, "bearer", config.Auth.Type)
}
