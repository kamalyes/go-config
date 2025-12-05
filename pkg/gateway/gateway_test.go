/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-05 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-05 00:00:00
 * @FilePath: \go-config\pkg\gateway\gateway_test.go
 * @Description: Gateway配置合并测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package gateway

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/banner"
	"github.com/kamalyes/go-config/pkg/cors"
	"github.com/kamalyes/go-config/pkg/jwt"
	"github.com/kamalyes/go-toolbox/pkg/safe"
	"github.com/stretchr/testify/assert"
)

// TestMergeWithDefaultsNilGateway 测试 nil Gateway 使用默认配置
func TestMergeWithDefaultsNilGateway(t *testing.T) {
	var nilGateway *Gateway
	defaultGateway := Default()

	result := safe.MergeWithDefaults(nilGateway, defaultGateway)

	assert.NotNil(t, result, "Expected non-nil result")
	assert.Equal(t, "gateway", result.ModuleName)
	assert.Equal(t, "Go RPC Gateway", result.Name)
	assert.True(t, result.Enabled)
	assert.Equal(t, "v1.0.0", result.Version)
	assert.Equal(t, "dev", result.Environment)
}

// TestMergeWithDefaultsPartialGateway 测试部分配置合并
func TestMergeWithDefaultsPartialGateway(t *testing.T) {
	partialGateway := &Gateway{
		ModuleName:  "custom-gateway",
		Name:        "Custom Gateway",
		Environment: "prod",
	}
	defaultGateway := Default()

	result := safe.MergeWithDefaults(partialGateway, defaultGateway)

	assert.Equal(t, partialGateway.ModuleName, result.ModuleName, "Should keep custom module name")
	assert.Equal(t, partialGateway.Name, result.Name, "Should keep custom name")
	assert.Equal(t, partialGateway.Environment, result.Environment, "Should keep custom environment")
	assert.True(t, result.Enabled, "Should use default enabled value")
	assert.Equal(t, "v1.0.0", result.Version, "Should use default version")
	assert.NotNil(t, result.HTTPServer, "Should have HTTPServer from default")
	assert.NotNil(t, result.Cache, "Should have Cache from default")
}

// TestMergeWithDefaultsHTTPServer 测试 HTTPServer 配置合并
func TestMergeWithDefaultsHTTPServer(t *testing.T) {
	partialGateway := &Gateway{
		Name: "Test Gateway",
		HTTPServer: &HTTPServer{
			Host: "custom.example.com",
			Port: 9090,
		},
	}
	defaultGateway := Default()

	result := safe.MergeWithDefaults(partialGateway, defaultGateway)

	assert.NotNil(t, result.HTTPServer)
	assert.Equal(t, partialGateway.HTTPServer.Host, result.HTTPServer.Host, "Should keep custom host")
	assert.Equal(t, partialGateway.HTTPServer.Port, result.HTTPServer.Port, "Should keep custom port")
	assert.NotNil(t, result.HTTPServer.TLS, "Should have TLS from default")
}

// TestMergeWithDefaultsNilHTTPServer 测试 nil HTTPServer 使用默认配置
func TestMergeWithDefaultsNilHTTPServer(t *testing.T) {
	partialGateway := &Gateway{
		Name:       "Test Gateway",
		HTTPServer: nil,
	}
	defaultGateway := Default()

	result := safe.MergeWithDefaults(partialGateway, defaultGateway)

	assert.NotNil(t, result.HTTPServer, "Should use default HTTPServer")
	assert.Equal(t, defaultGateway.HTTPServer.Host, result.HTTPServer.Host)
	assert.Equal(t, defaultGateway.HTTPServer.Port, result.HTTPServer.Port)
}

// TestMergeWithDefaultsNestedConfig 测试嵌套配置合并
func TestMergeWithDefaultsNestedConfig(t *testing.T) {
	partialGateway := &Gateway{
		Name: "Test Gateway",
		Banner: &banner.Banner{
			Enabled: true,
			Title:   "Custom Title",
		},
	}
	defaultGateway := Default()

	result := safe.MergeWithDefaults(partialGateway, defaultGateway)

	assert.NotNil(t, result.Banner)
	assert.True(t, result.Banner.Enabled, "Should keep custom enabled")
	assert.Equal(t, partialGateway.Banner.Title, result.Banner.Title, "Should keep custom title")
	// 其他字段应该从默认配置获取
	assert.NotEmpty(t, result.Banner.Description, "Should have description from default")
}

// TestMergeWithDefaultsJWT 测试 JWT 配置合并
func TestMergeWithDefaultsJWT(t *testing.T) {
	partialGateway := &Gateway{
		Name: "Test Gateway",
		JWT: &jwt.JWT{
			SigningKey: "custom-secret-key",
			Issuer:     "custom-issuer",
		},
	}
	defaultGateway := Default()

	result := safe.MergeWithDefaults(partialGateway, defaultGateway)

	assert.NotNil(t, result.JWT)
	assert.Equal(t, partialGateway.JWT.SigningKey, result.JWT.SigningKey, "Should keep custom secret key")
	assert.Equal(t, partialGateway.JWT.Issuer, result.JWT.Issuer, "Should keep custom issuer")
	// 其他字段应该从默认配置获取
	assert.NotZero(t, result.JWT.ExpiresTime, "Should have expire time from default")
}

// TestMergeWithDefaultsCORS 测试 CORS 配置合并
func TestMergeWithDefaultsCORS(t *testing.T) {
	partialGateway := &Gateway{
		Name: "Test Gateway",
		CORS: &cors.Cors{
			AllowedOrigins: []string{"https://custom.example.com"},
		},
	}
	defaultGateway := Default()

	result := safe.MergeWithDefaults(partialGateway, defaultGateway)

	assert.NotNil(t, result.CORS)
	assert.Len(t, result.CORS.AllowedOrigins, len(partialGateway.CORS.AllowedOrigins), "Should keep custom origins")
	assert.Equal(t, partialGateway.CORS.AllowedOrigins[0], result.CORS.AllowedOrigins[0])
}

// TestMergeWithDefaultsMultipleDefaults 测试多个默认配置
func TestMergeWithDefaultsMultipleDefaults(t *testing.T) {
	partialGateway := &Gateway{
		Name: "Test Gateway",
	}
	default1 := &Gateway{
		ModuleName: "default1",
		Version:    "v1.0.0",
	}
	default2 := &Gateway{
		ModuleName:  "default2",
		Environment: "staging",
		Debug:       true,
	}

	result := safe.MergeWithDefaults(partialGateway, default1, default2)

	assert.Equal(t, partialGateway.Name, result.Name, "Should keep original name")
	assert.Equal(t, default1.ModuleName, result.ModuleName, "Should use first default's module name")
	assert.Equal(t, default1.Version, result.Version, "Should use first default's version")
	assert.Equal(t, default2.Environment, result.Environment, "Should use second default's environment")
	assert.True(t, result.Debug, "Should use second default's debug")
}

// TestMergeWithDefaultsZeroValues 测试零值字段合并
func TestMergeWithDefaultsZeroValues(t *testing.T) {
	partialGateway := &Gateway{
		Name:    "",
		Enabled: false,
		Debug:   false,
	}
	defaultGateway := Default()

	result := safe.MergeWithDefaults(partialGateway, defaultGateway)

	assert.Equal(t, defaultGateway.Name, result.Name, "Empty string should use default")
	assert.True(t, result.Enabled, "False bool should use default true")
	assert.True(t, result.Debug, "False debug should use default true")
}

// TestGatewayBuilderMethods 测试 Gateway 构建器方法
func TestGatewayBuilderMethods(t *testing.T) {
	gateway := Default().
		WithName("Test Gateway").
		WithVersion("v2.0.0").
		WithEnvironment("production").
		WithDebug(false).
		EnableTLS().
		EnableBanner().
		EnableSwagger()

	assert.Equal(t, "Test Gateway", gateway.Name)
	assert.Equal(t, "v2.0.0", gateway.Version)
	assert.Equal(t, "production", gateway.Environment)
	assert.False(t, gateway.Debug)
	assert.True(t, gateway.HTTPServer.EnableTls)
	assert.True(t, gateway.Banner.Enabled)
	assert.True(t, gateway.Swagger.Enabled)
}
