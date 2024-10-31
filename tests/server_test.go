/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 12:08:55
 * @FilePath: \go-config\server\server_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/server"
	"github.com/stretchr/testify/assert"
)

// getValidServerConfig 返回有效的 Server 配置
func getValidServerConfig() *server.Server {
	return server.NewServer("test-module", "localhost", "8080", "TestServer", "/api", "mysql", true, "en")
}

// getInvalidServerConfig 返回无效的 Server 配置
func getInvalidServerConfig() *server.Server {
	return server.NewServer("", "", "", "", "", "", false, "")
}

func TestNewServer(t *testing.T) {
	// 测试创建 Server 实例
	server := getValidServerConfig()

	assert.NotNil(t, server)
	assert.Equal(t, "test-module", server.ModuleName)
	assert.Equal(t, "localhost", server.Host)
	assert.Equal(t, "8080", server.Port)
	assert.Equal(t, "TestServer", server.ServerName)
	assert.Equal(t, "/api", server.ContextPath)
	assert.Equal(t, "mysql", server.DataDriver)
	assert.True(t, server.HandleMethodNotAllowed)
	assert.Equal(t, "en", server.Language)
}

func TestServer_Validate(t *testing.T) {
	// 测试有效配置
	validServer := getValidServerConfig()
	assert.NoError(t, validServer.Validate())

	// 测试无效配置
	invalidServer := getInvalidServerConfig()
	assert.Error(t, invalidServer.Validate())
}

func TestServer_ToMap(t *testing.T) {
	// 测试将配置转换为映射
	s := getValidServerConfig()
	sMap := s.ToMap()

	assert.Equal(t, "test-module", sMap["moduleName"])
	assert.Equal(t, "localhost", sMap["host"])
	assert.Equal(t, "8080", sMap["port"])
	assert.Equal(t, "TestServer", sMap["serverName"])
	assert.Equal(t, "/api", sMap["contextPath"])
	assert.Equal(t, "mysql", sMap["dataDriver"])
	assert.Equal(t, true, sMap["handleMethodNotAllowed"])
	assert.Equal(t, "en", sMap["language"])
}

func TestServer_FromMap(t *testing.T) {
	// 测试从映射填充配置
	data := map[string]interface{}{
		"moduleName":             "test-module",
		"host":                   "localhost",
		"port":                   "8080",
		"serverName":             "TestServer",
		"contextPath":            "/api",
		"dataDriver":             "mysql",
		"handleMethodNotAllowed": true,
		"language":               "en",
	}

	s := &server.Server{}
	s.FromMap(data)

	assert.Equal(t, "test-module", s.ModuleName)
	assert.Equal(t, "localhost", s.Host)
	assert.Equal(t, "8080", s.Port)
	assert.Equal(t, "TestServer", s.ServerName)
	assert.Equal(t, "/api", s.ContextPath)
	assert.Equal(t, "mysql", s.DataDriver)
	assert.True(t, s.HandleMethodNotAllowed)
	assert.Equal(t, "en", s.Language)
}
