/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\rpcclient\rpcclient_test.go
 * @Description: RPC客户端配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package rpcclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRpcClient_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "rpc-client", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.NotNil(t, config.Endpoints)
	assert.Equal(t, []string{"localhost:8080"}, config.Endpoints)
	assert.Equal(t, 5000, config.Timeout)
	assert.Equal(t, 3, config.DialTimeout)
	assert.Equal(t, 30, config.KeepaliveTime)
	assert.Equal(t, 3, config.MaxRetries)
	assert.Equal(t, 1000, config.RetryInterval)
	assert.Equal(t, "round_robin", config.LoadBalance)
	assert.NotNil(t, config.Headers)
	assert.NotNil(t, config.TLS)
	assert.NotNil(t, config.CircuitBreaker)
}

func TestRpcClient_WithModuleName(t *testing.T) {
	config := Default()
	result := config.WithModuleName("custom-client")
	assert.Equal(t, "custom-client", result.ModuleName)
	assert.Equal(t, config, result)
}

func TestRpcClient_WithEndpoints(t *testing.T) {
	config := Default()
	endpoints := []string{"localhost:8080", "localhost:8081"}
	result := config.WithEndpoints(endpoints)
	assert.Equal(t, endpoints, result.Endpoints)
	assert.Equal(t, config, result)
}

func TestRpcClient_WithTarget(t *testing.T) {
	config := Default()
	result := config.WithTarget("service-name")
	assert.Equal(t, "service-name", result.Target)
	assert.Equal(t, config, result)
}

func TestRpcClient_WithApp(t *testing.T) {
	config := Default()
	result := config.WithApp("my-app")
	assert.Equal(t, "my-app", result.App)
	assert.Equal(t, config, result)
}

func TestRpcClient_WithAuth(t *testing.T) {
	config := Default()
	result := config.WithAuth("token123")
	assert.Equal(t, "token123", result.Token)
	assert.Equal(t, config, result)
}

func TestRpcClient_WithTimeout(t *testing.T) {
	config := Default()
	result := config.WithTimeout(10000)
	assert.Equal(t, 10000, result.Timeout)
	assert.Equal(t, config, result)
}

func TestRpcClient_WithRetry(t *testing.T) {
	config := Default()
	result := config.WithRetry(5, 2000)
	assert.Equal(t, 5, result.MaxRetries)
	assert.Equal(t, 2000, result.RetryInterval)
	assert.Equal(t, config, result)
}

func TestRpcClient_WithLoadBalance(t *testing.T) {
	config := Default()
	result := config.WithLoadBalance("random")
	assert.Equal(t, "random", result.LoadBalance)
	assert.Equal(t, config, result)
}

func TestRpcClient_WithTLS(t *testing.T) {
	config := Default()
	result := config.WithTLS(true, "/cert.pem", "/key.pem", "/ca.pem")
	assert.True(t, result.TLS.Enabled)
	assert.Equal(t, "/cert.pem", result.TLS.CertFile)
	assert.Equal(t, "/key.pem", result.TLS.KeyFile)
	assert.Equal(t, "/ca.pem", result.TLS.CACertFile)
	assert.Equal(t, config, result)
}

func TestRpcClient_EnableCircuitBreaker(t *testing.T) {
	config := Default()
	config.CircuitBreaker.Enabled = false
	result := config.EnableCircuitBreaker()
	assert.True(t, result.CircuitBreaker.Enabled)
	assert.Equal(t, config, result)
}

func TestRpcClient_Enable(t *testing.T) {
	config := Default()
	result := config.Enable()
	assert.True(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestRpcClient_Disable(t *testing.T) {
	config := Default()
	config.Enabled = true
	result := config.Disable()
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestRpcClient_Clone(t *testing.T) {
	config := Default()
	config.WithEndpoints([]string{"localhost:8080"})
	config.Headers["X-Custom"] = "value"

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*RpcClient)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.Timeout, clonedConfig.Timeout)

	// 验证深拷贝 - 切片
	clonedConfig.Endpoints = append(clonedConfig.Endpoints, "localhost:8082")
	assert.NotEqual(t, len(config.Endpoints), len(clonedConfig.Endpoints))

	// 验证深拷贝 - map
	clonedConfig.Headers["X-New"] = "new-value"
	_, exists := config.Headers["X-New"]
	assert.False(t, exists)

	// 验证深拷贝 - 嵌套结构
	clonedConfig.TLS.CertFile = "modified"
	assert.NotEqual(t, config.TLS.CertFile, clonedConfig.TLS.CertFile)
}

func TestRpcClient_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestRpcClient_Set(t *testing.T) {
	config := Default()
	newConfig := &RpcClient{
		ModuleName: "new-client",
		Enabled:    true,
		Timeout:    8000,
	}

	config.Set(newConfig)
	assert.Equal(t, "new-client", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, 8000, config.Timeout)
}

func TestRpcClient_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
	assert.Equal(t, "rpc-client", config.ModuleName)
	assert.Equal(t, 5000, config.Timeout)
}

func TestRpcClient_NewRpcClient(t *testing.T) {
	opt := &RpcClient{
		ModuleName: "test-client",
		Timeout:    6000,
	}

	result := NewRpcClient(opt)
	assert.NotNil(t, result)
	assert.Equal(t, "test-client", result.ModuleName)
	assert.Equal(t, 6000, result.Timeout)
}

func TestRpcClient_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("my-client").
		WithEndpoints([]string{"localhost:8080"}).
		WithTarget("my-service").
		WithApp("my-app").
		WithAuth("token123").
		WithTimeout(10000).
		WithRetry(5, 2000).
		WithLoadBalance("weighted").
		WithTLS(true, "/cert.pem", "/key.pem", "/ca.pem").
		EnableCircuitBreaker().
		Enable()

	assert.Equal(t, "my-client", config.ModuleName)
	assert.Equal(t, []string{"localhost:8080"}, config.Endpoints)
	assert.Equal(t, "my-service", config.Target)
	assert.Equal(t, "my-app", config.App)
	assert.Equal(t, "token123", config.Token)
	assert.Equal(t, 10000, config.Timeout)
	assert.Equal(t, 5, config.MaxRetries)
	assert.Equal(t, 2000, config.RetryInterval)
	assert.Equal(t, "weighted", config.LoadBalance)
	assert.True(t, config.TLS.Enabled)
	assert.True(t, config.CircuitBreaker.Enabled)
	assert.True(t, config.Enabled)
}
