/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 09:55:52
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 12:16:29
 * @FilePath: \go-config\zero\server_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/zero"
	"github.com/stretchr/testify/assert"
)

// GetValidRpcServerConfig 返回有效的 RpcServer 配置
func GetValidRpcServerConfig() *zero.RpcServer {
	etcdConfig := zero.NewEtcdConfig([]string{"localhost:2379"}, "my_key")
	return zero.NewRpcServer("test-module", "0.0.0.0:8080", true, false, 5000, 80, etcdConfig)
}

// GetInvalidRpcServerConfig 返回无效的 RpcServer 配置
func GetInvalidRpcServerConfig() *zero.RpcServer {
	etcdConfig := zero.NewEtcdConfig([]string{}, "")                   // Invalid EtcdConfig
	return zero.NewRpcServer("", "", false, false, -1, -1, etcdConfig) // Invalid RpcServer
}

func TestNewRpcServer(t *testing.T) {
	// 测试创建有效的 RpcServer 配置
	validConfig := GetValidRpcServerConfig()
	assert.NotNil(t, validConfig)

	// 测试创建无效的 RpcServer 配置
	invalidConfig := GetInvalidRpcServerConfig()
	assert.Error(t, invalidConfig.Validate())
}

func TestRpcServer_Validate(t *testing.T) {
	validConfig := GetValidRpcServerConfig()
	assert.NoError(t, validConfig.Validate())

	invalidConfig := GetInvalidRpcServerConfig()
	assert.Error(t, invalidConfig.Validate())
}

func TestRpcServer_ToMap(t *testing.T) {
	config := GetValidRpcServerConfig()
	mapped := config.ToMap()
	assert.Equal(t, config.ModuleName, mapped["moduleName"])
	assert.Equal(t, config.ListenOn, mapped["listenOn"])
	assert.Equal(t, config.Auth, mapped["auth"])
	assert.Equal(t, config.StrictControl, mapped["strictControl"])
	assert.Equal(t, config.Timeout, mapped["timeout"])
	assert.Equal(t, config.CpuThreshold, mapped["cpuThreshold"])
	etcdMap := mapped["etcd"].(map[string]interface{})
	assert.Equal(t, config.Etcd.Hosts, etcdMap["hosts"])
	assert.Equal(t, config.Etcd.Key, etcdMap["key"])
}

func TestRpcServer_FromMap(t *testing.T) {
	config := GetValidRpcServerConfig()
	mapped := config.ToMap()

	newConfig := &zero.RpcServer{}
	newConfig.FromMap(mapped)

	assert.Equal(t, config.ModuleName, newConfig.ModuleName)
	assert.Equal(t, config.ListenOn, newConfig.ListenOn)
	assert.Equal(t, config.Auth, newConfig.Auth)
	assert.Equal(t, config.StrictControl, newConfig.StrictControl)
	assert.Equal(t, config.Timeout, newConfig.Timeout)
	assert.Equal(t, config.CpuThreshold, newConfig.CpuThreshold)
	assert.Equal(t, config.Etcd.Hosts, newConfig.Etcd.Hosts)
	assert.Equal(t, config.Etcd.Key, newConfig.Etcd.Key)
}
