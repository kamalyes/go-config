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

// TestHTTPServer_Clone 测试 HTTPServer 克隆
func TestHTTPServer_Clone(t *testing.T) {
	original := &HTTPServer{
		ModuleName:         "test-server",
		Host:               "localhost",
		Port:               8080,
		Network:            "tcp",
		ReadTimeout:        30,
		WriteTimeout:       30,
		IdleTimeout:        60,
		MaxHeaderBytes:     1048576,
		EnableTls:          true,
		TLS:                &TLS{CertFile: "cert.pem", KeyFile: "key.pem", CAFile: "ca.pem"},
		Headers:            map[string]string{"X-Custom": "value"},
		Endpoint:           "http://localhost:8080",
		EnableGzipCompress: true,
	}

	cloned := original.Clone().(*HTTPServer)

	// 验证克隆后的值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Host, cloned.Host)
	assert.Equal(t, original.Port, cloned.Port)
	assert.Equal(t, original.Network, cloned.Network)
	assert.Equal(t, original.ReadTimeout, cloned.ReadTimeout)
	assert.Equal(t, original.WriteTimeout, cloned.WriteTimeout)
	assert.Equal(t, original.IdleTimeout, cloned.IdleTimeout)
	assert.Equal(t, original.MaxHeaderBytes, cloned.MaxHeaderBytes)
	assert.Equal(t, original.EnableTls, cloned.EnableTls)
	assert.Equal(t, original.EnableGzipCompress, cloned.EnableGzipCompress)
	assert.Equal(t, original.Endpoint, cloned.Endpoint)

	// 验证 TLS 深拷贝
	assert.NotSame(t, original.TLS, cloned.TLS)
	assert.Equal(t, original.TLS.CertFile, cloned.TLS.CertFile)
	assert.Equal(t, original.TLS.KeyFile, cloned.TLS.KeyFile)
	assert.Equal(t, original.TLS.CAFile, cloned.TLS.CAFile)

	// 验证 Headers 深拷贝
	assert.Equal(t, original.Headers, cloned.Headers)

	// 修改原始对象不应影响克隆对象
	original.Port = 9090
	original.Headers["X-Custom"] = "new-value"
	original.TLS.CertFile = "new-cert.pem"

	assert.NotEqual(t, original.Port, cloned.Port)
	assert.NotEqual(t, original.Headers["X-Custom"], cloned.Headers["X-Custom"])
	assert.NotEqual(t, original.TLS.CertFile, cloned.TLS.CertFile)
}

// TestGRPC_Clone 测试 GRPC 克隆
func TestGRPC_Clone(t *testing.T) {
	original := &GRPC{
		Server: &GRPCServer{
			Enable:            true,
			Host:              "localhost",
			Port:              50051,
			Network:           "tcp",
			MaxRecvMsgSize:    1024,
			MaxSendMsgSize:    1024,
			KeepaliveTime:     30,
			KeepaliveTimeout:  10,
			ConnectionTimeout: 5,
			EnableReflection:  true,
			Endpoint:          "localhost:50051",
		},
		Clients: map[string]*GRPCClient{
			"service1": {
				ServiceName:       "service1",
				Endpoints:         []string{"localhost:50052", "localhost:50053"},
				Network:           "tcp",
				MaxRecvMsgSize:    2048,
				MaxSendMsgSize:    2048,
				KeepaliveTime:     60,
				KeepaliveTimeout:  20,
				ConnectionTimeout: 10,
			},
		},
	}

	cloned := original.Clone().(*GRPC)

	// 验证 Server 深拷贝
	assert.NotSame(t, original.Server, cloned.Server)
	assert.Equal(t, original.Server.Enable, cloned.Server.Enable)
	assert.Equal(t, original.Server.Host, cloned.Server.Host)
	assert.Equal(t, original.Server.Port, cloned.Server.Port)
	assert.Equal(t, original.Server.Endpoint, cloned.Server.Endpoint)

	// 验证 Clients 深拷贝
	assert.Equal(t, len(original.Clients), len(cloned.Clients))
	assert.NotSame(t, original.Clients["service1"], cloned.Clients["service1"])
	assert.Equal(t, original.Clients["service1"].ServiceName, cloned.Clients["service1"].ServiceName)
	assert.Equal(t, original.Clients["service1"].Endpoints, cloned.Clients["service1"].Endpoints)

	// 修改原始对象不应影响克隆对象
	original.Server.Port = 50052
	original.Clients["service1"].ServiceName = "new-service"

	assert.NotEqual(t, original.Server.Port, cloned.Server.Port)
	assert.NotEqual(t, original.Clients["service1"].ServiceName, cloned.Clients["service1"].ServiceName)
}

// TestGRPCServer_Clone 测试 GRPCServer 克隆
func TestGRPCServer_Clone(t *testing.T) {
	original := &GRPCServer{
		Enable:            true,
		Host:              "0.0.0.0",
		Port:              50051,
		Network:           "tcp",
		MaxRecvMsgSize:    4194304,
		MaxSendMsgSize:    4194304,
		KeepaliveTime:     30,
		KeepaliveTimeout:  10,
		ConnectionTimeout: 5,
		EnableReflection:  true,
		Endpoint:          "0.0.0.0:50051",
	}

	cloned := original.Clone()

	// 验证所有字段相等
	assert.Equal(t, original.Enable, cloned.Enable)
	assert.Equal(t, original.Host, cloned.Host)
	assert.Equal(t, original.Port, cloned.Port)
	assert.Equal(t, original.Network, cloned.Network)
	assert.Equal(t, original.MaxRecvMsgSize, cloned.MaxRecvMsgSize)
	assert.Equal(t, original.MaxSendMsgSize, cloned.MaxSendMsgSize)
	assert.Equal(t, original.KeepaliveTime, cloned.KeepaliveTime)
	assert.Equal(t, original.KeepaliveTimeout, cloned.KeepaliveTimeout)
	assert.Equal(t, original.ConnectionTimeout, cloned.ConnectionTimeout)
	assert.Equal(t, original.EnableReflection, cloned.EnableReflection)
	assert.Equal(t, original.Endpoint, cloned.Endpoint)

	// 修改原始对象不应影响克隆对象
	original.Port = 50052
	original.Endpoint = "0.0.0.0:50052"
	assert.NotEqual(t, original.Port, cloned.Port)
	assert.NotEqual(t, original.Endpoint, cloned.Endpoint)
}

// TestGRPCClient_Clone 测试 GRPCClient 克隆
func TestGRPCClient_Clone(t *testing.T) {
	original := &GRPCClient{
		ServiceName:       "test-service",
		Endpoints:         []string{"localhost:50051", "localhost:50052"},
		Network:           "tcp",
		MaxRecvMsgSize:    4194304,
		MaxSendMsgSize:    4194304,
		KeepaliveTime:     60,
		KeepaliveTimeout:  20,
		ConnectionTimeout: 10,
	}

	cloned := original.Clone()

	// 验证所有字段相等
	assert.Equal(t, original.ServiceName, cloned.ServiceName)
	assert.Equal(t, original.Endpoints, cloned.Endpoints)
	assert.Equal(t, original.Network, cloned.Network)
	assert.Equal(t, original.MaxRecvMsgSize, cloned.MaxRecvMsgSize)
	assert.Equal(t, original.MaxSendMsgSize, cloned.MaxSendMsgSize)
	assert.Equal(t, original.KeepaliveTime, cloned.KeepaliveTime)
	assert.Equal(t, original.KeepaliveTimeout, cloned.KeepaliveTimeout)
	assert.Equal(t, original.ConnectionTimeout, cloned.ConnectionTimeout)

	// 修改原始对象不应影响克隆对象
	original.ServiceName = "new-service"
	original.Endpoints[0] = "localhost:50053"
	assert.NotEqual(t, original.ServiceName, cloned.ServiceName)
	assert.NotEqual(t, original.Endpoints[0], cloned.Endpoints[0])
}
