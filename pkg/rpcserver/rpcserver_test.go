/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:33:16
 * @FilePath: \go-config\pkg\rpcserver\rpcserver_test.go
 * @Description: RPC服务器配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package rpcserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRpcServer_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "rpc-server", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "rpc-service", config.Name)
	assert.Equal(t, "127.0.0.1:8080", config.ListenOn)
	assert.Equal(t, "dev", config.Mode)
	assert.Equal(t, 1000, config.MaxConns)
	assert.Equal(t, 4*1024*1024, config.MaxMsgSize)
	assert.Equal(t, 30, config.Timeout)
	assert.Equal(t, int64(900), config.CpuThreshold)
	assert.True(t, config.Health)
	assert.False(t, config.Auth)
	assert.False(t, config.StrictControl)
	assert.NotNil(t, config.Headers)
	assert.NotNil(t, config.Middlewares)
	assert.NotNil(t, config.TLS)
	assert.NotNil(t, config.RateLimit)
	assert.NotNil(t, config.Recovery)
	assert.NotNil(t, config.Tracing)
}

func TestRpcServer_WithModuleName(t *testing.T) {
	config := Default()
	result := config.WithModuleName("custom-server")
	assert.Equal(t, "custom-server", result.ModuleName)
	assert.Equal(t, config, result)
}

func TestRpcServer_WithName(t *testing.T) {
	config := Default()
	result := config.WithName("my-service")
	assert.Equal(t, "my-service", result.Name)
	assert.Equal(t, config, result)
}

func TestRpcServer_WithListenOn(t *testing.T) {
	config := Default()
	result := config.WithListenOn("0.0.0.0:9090")
	assert.Equal(t, "0.0.0.0:9090", result.ListenOn)
	assert.Equal(t, config, result)
}

func TestRpcServer_WithMode(t *testing.T) {
	config := Default()
	result := config.WithMode("prod")
	assert.Equal(t, "prod", result.Mode)
	assert.Equal(t, config, result)
}

func TestRpcServer_WithMaxConns(t *testing.T) {
	config := Default()
	result := config.WithMaxConns(2000)
	assert.Equal(t, 2000, result.MaxConns)
	assert.Equal(t, config, result)
}

func TestRpcServer_WithTimeout(t *testing.T) {
	config := Default()
	result := config.WithTimeout(60)
	assert.Equal(t, 60, result.Timeout)
	assert.Equal(t, config, result)
}

func TestRpcServer_WithTLS(t *testing.T) {
	config := Default()
	result := config.WithTLS(true, "/cert.pem", "/key.pem", "/ca.pem")
	assert.True(t, result.TLS.Enabled)
	assert.Equal(t, "/cert.pem", result.TLS.CertFile)
	assert.Equal(t, "/key.pem", result.TLS.KeyFile)
	assert.Equal(t, "/ca.pem", result.TLS.CACertFile)
	assert.Equal(t, config, result)
}

func TestRpcServer_WithRateLimit(t *testing.T) {
	config := Default()
	result := config.WithRateLimit(true, 10, 100)
	assert.True(t, result.RateLimit.Enabled)
	assert.Equal(t, 10, result.RateLimit.Seconds)
	assert.Equal(t, 100, result.RateLimit.Quota)
	assert.Equal(t, config, result)
}

func TestRpcServer_EnableAuth(t *testing.T) {
	config := Default()
	result := config.EnableAuth()
	assert.True(t, result.Auth)
	assert.Equal(t, config, result)
}

func TestRpcServer_EnableHealth(t *testing.T) {
	config := Default()
	config.Health = false
	result := config.EnableHealth()
	assert.True(t, result.Health)
	assert.Equal(t, config, result)
}

func TestRpcServer_EnableTracing(t *testing.T) {
	config := Default()
	result := config.EnableTracing("http://jaeger:14268", "my-service", 0.5)
	assert.True(t, result.Tracing.Enabled)
	assert.Equal(t, "http://jaeger:14268", result.Tracing.Endpoint)
	assert.Equal(t, "my-service", result.Tracing.ServiceName)
	assert.Equal(t, 0.5, result.Tracing.Sampler)
	assert.Equal(t, config, result)
}

func TestRpcServer_Enable(t *testing.T) {
	config := Default()
	result := config.Enable()
	assert.True(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestRpcServer_Disable(t *testing.T) {
	config := Default()
	config.Enabled = true
	result := config.Disable()
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestRpcServer_Clone(t *testing.T) {
	config := Default()
	config.WithName("test-service")
	config.Headers["X-Custom"] = "value"
	config.Middlewares = append(config.Middlewares, "logger")

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*RpcServer)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.Name, clonedConfig.Name)

	// 验证深拷贝 - 切片
	clonedConfig.Middlewares = append(clonedConfig.Middlewares, "extra")
	assert.NotEqual(t, len(config.Middlewares), len(clonedConfig.Middlewares))

	// 验证深拷贝 - map
	clonedConfig.Headers["X-New"] = "new-value"
	_, exists := config.Headers["X-New"]
	assert.False(t, exists)

	// 验证深拷贝 - 嵌套结构
	clonedConfig.TLS.CertFile = "modified"
	assert.NotEqual(t, config.TLS.CertFile, clonedConfig.TLS.CertFile)
}

func TestRpcServer_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestRpcServer_Set(t *testing.T) {
	config := Default()
	newConfig := &RpcServer{
		ModuleName: "new-server",
		Enabled:    true,
		ListenOn:   "127.0.0.1:8080",
	}

	config.Set(newConfig)
	assert.Equal(t, "new-server", config.ModuleName)
	assert.True(t, config.Enabled)
}

func TestRpcServer_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
	assert.Equal(t, "rpc-server", config.ModuleName)
	assert.Equal(t, "127.0.0.1:8080", config.ListenOn)
}

func TestRpcServer_NewRpcServer(t *testing.T) {
	opt := &RpcServer{
		ModuleName: "test-server",
		ListenOn:   "0.0.0.0:9090",
	}

	result := NewRpcServer(opt)
	assert.NotNil(t, result)
	assert.Equal(t, "test-server", result.ModuleName)
	assert.Equal(t, "0.0.0.0:9090", result.ListenOn)
}

func TestRpcServer_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("my-server").
		WithName("my-rpc-service").
		WithListenOn("0.0.0.0:9999").
		WithMode("prod").
		WithMaxConns(5000).
		WithTimeout(120).
		WithTLS(true, "/cert.pem", "/key.pem", "/ca.pem").
		WithRateLimit(true, 60, 1000).
		EnableAuth().
		EnableHealth().
		EnableTracing("http://jaeger:14268", "my-service", 0.1).
		Enable()

	assert.Equal(t, "my-server", config.ModuleName)
	assert.Equal(t, "my-rpc-service", config.Name)
	assert.Equal(t, "0.0.0.0:9999", config.ListenOn)
	assert.Equal(t, "prod", config.Mode)
	assert.Equal(t, 5000, config.MaxConns)
	assert.Equal(t, 120, config.Timeout)
	assert.True(t, config.TLS.Enabled)
	assert.True(t, config.RateLimit.Enabled)
	assert.True(t, config.Auth)
	assert.True(t, config.Health)
	assert.True(t, config.Tracing.Enabled)
	assert.True(t, config.Enabled)
}
