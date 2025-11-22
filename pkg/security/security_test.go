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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecurity_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "security", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.NotNil(t, config.JWT)
	assert.NotNil(t, config.Auth)
	assert.NotNil(t, config.Protection)
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

func TestSecurity_WithEnabled(t *testing.T) {
	config := Default()
	result := config.WithEnabled(false)
	assert.False(t, result.Enabled)
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

func TestSecurity_Enable(t *testing.T) {
	config := Default()
	config.Enabled = false
	result := config.Enable()
	assert.True(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestSecurity_Disable(t *testing.T) {
	config := Default()
	result := config.Disable()
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestSecurity_IsEnabled(t *testing.T) {
	config := Default()
	assert.True(t, config.IsEnabled())

	config.Enabled = false
	assert.False(t, config.IsEnabled())
}

func TestSecurity_Clone(t *testing.T) {
	config := Default()
	config.WithJWT(true, "secret", 24, "app", "HS256")

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*Security)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.Enabled, clonedConfig.Enabled)
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
		Enabled:    false,
	}

	config.Set(newConfig)
	assert.Equal(t, "new-security", config.ModuleName)
	assert.False(t, config.Enabled)
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

func TestSecurity_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("my-security").
		WithJWT(true, "my-secret", 72, "my-issuer", "HS512").
		WithAuth(true, "bearer", "Authorization", "Bearer ").
		EnableJWT().
		EnableAuth().
		Enable()

	assert.Equal(t, "my-security", config.ModuleName)
	assert.True(t, config.JWT.Enabled)
	assert.Equal(t, "my-secret", config.JWT.Secret)
	assert.Equal(t, 72, config.JWT.Expiry)
	assert.True(t, config.Auth.Enabled)
	assert.Equal(t, "bearer", config.Auth.Type)
	assert.True(t, config.Enabled)
}
