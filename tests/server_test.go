/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 10:51:18
 * @FilePath: \go-config\tests\server_test.go
 * @Description: 服务器配置模块测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/server"
	"github.com/stretchr/testify/assert"
)

// 验证 Server 的字段与期望的映射是否相等
func assertServerFields(t *testing.T, actual *server.Server, expected *server.Server) {
	assert.Equal(t, expected.ModuleName, actual.ModuleName)
	assert.Equal(t, expected.Host, actual.Host)
	assert.Equal(t, expected.Port, actual.Port)
	assert.Equal(t, expected.ReadTimeout, actual.ReadTimeout)
	assert.Equal(t, expected.WriteTimeout, actual.WriteTimeout)
	assert.Equal(t, expected.IdleTimeout, actual.IdleTimeout)
	assert.Equal(t, expected.GrpcPort, actual.GrpcPort)
	assert.Equal(t, expected.EnableHttp, actual.EnableHttp)
	assert.Equal(t, expected.EnableGrpc, actual.EnableGrpc)
	assert.Equal(t, expected.EnableTls, actual.EnableTls)
	assert.Equal(t, expected.TLS, actual.TLS)
	assert.Equal(t, expected.Headers, actual.Headers)
}

func TestDefaultServer(t *testing.T) {
	defaultServer := server.Default()

	assert.NotNil(t, defaultServer)
	assert.Equal(t, "server", defaultServer.ModuleName)
	assert.Equal(t, "localhost", defaultServer.Host)
	assert.Equal(t, 8080, defaultServer.Port)
	assert.Equal(t, 30, defaultServer.ReadTimeout)
	assert.Equal(t, 30, defaultServer.WriteTimeout)
	assert.Equal(t, 60, defaultServer.IdleTimeout)
	assert.Equal(t, 9090, defaultServer.GrpcPort)
	assert.True(t, defaultServer.EnableHttp)
	assert.False(t, defaultServer.EnableGrpc)
	assert.False(t, defaultServer.EnableTls)
}

func TestSetServer(t *testing.T) {
	params := server.Default()
	params.Host = "127.0.0.1"
	params.Port = 9090

	newServer := server.Default()
	newServer.Set(params)

	assertServerFields(t, newServer, params)
}

func TestCloneServer(t *testing.T) {
	original := server.Default()
	original.AddHeader("X-Custom-Header", "value")

	cloned := original.Clone().(*server.Server)

	assertServerFields(t, cloned, original)
	assert.NotSame(t, original, cloned)               // 确保是不同的实例
	assert.Equal(t, original.Headers, cloned.Headers) // 确保头部也被克隆
}

func TestValidateServer(t *testing.T) {
	validParams := &server.Server{
		Endpoint:     "localhost:8080",
		Host:         "localhost",
		Port:         8080,
		ReadTimeout:  30,
		WriteTimeout: 30,
		IdleTimeout:  60,
		GrpcPort:     9090,
		EnableHttp:   true,
		EnableGrpc:   false,
		EnableTls:    false,
	}

	serverInstance := server.Default()
	serverInstance.Set(validParams)
	err := serverInstance.Validate()
	assert.NoError(t, err)
}

func TestAddHeader(t *testing.T) {
	srv := server.Default()
	srv.AddHeader("X-Test-Header", "TestValue")

	assert.Equal(t, "TestValue", srv.Headers["X-Test-Header"])
}

func TestWithTLS(t *testing.T) {
	srv := server.Default()
	srv.WithTLS("path/to/cert.pem", "path/to/key.pem", "path/to/ca.pem")

	assert.True(t, srv.EnableTls)
	assert.Equal(t, "path/to/cert.pem", srv.TLS.CertFile)
	assert.Equal(t, "path/to/key.pem", srv.TLS.KeyFile)
	assert.Equal(t, "path/to/ca.pem", srv.TLS.CAFile)
}
