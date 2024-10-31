/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 09:55:52
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:58:32
 * @FilePath: \go-config\tests\zero_server_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/zero"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 EtcdConfig 配置参数
func generateEtcdConfig() *zero.EtcdConfig {
	return &zero.EtcdConfig{
		Hosts: []string{fmt.Sprintf("http://%s:2379", random.RandString(5, random.CAPITAL))}, // 随机生成 Etcd 地址
		Key:   random.RandString(10, random.CAPITAL),                                         // 随机生成用户名
	}
}

// 生成随机的 RpcServer 配置参数
func generateRpcServerTestParams() *zero.RpcServer {
	return &zero.RpcServer{
		ModuleName:    random.RandString(10, random.CAPITAL),                  // 随机生成模块名称
		ListenOn:      fmt.Sprintf("0.0.0.0:%d", random.RandInt(1000, 65535)), // 随机生成监听地址
		Auth:          random.FRandBool(),                                     // 随机生成是否启用认证
		StrictControl: random.FRandBool(),                                     // 随机生成是否启用严格控制
		Timeout:       int64(random.RandInt(100, 5000)),                       // 随机生成超时时间（单位：毫秒）
		CpuThreshold:  int64(random.RandInt(1, 100)),                          // 随机生成 CPU 使用率阈值
		Etcd:          generateEtcdConfig(),                                   // 随机生成 Etcd 配置
	}
}

// 将 RpcServer 的参数转换为 map
func rpcServerToMap(server *zero.RpcServer) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME":    server.ModuleName,
		"LISTEN_ON":      server.ListenOn,
		"AUTH":           server.Auth,
		"STRICT_CONTROL": server.StrictControl,
		"TIMEOUT":        server.Timeout,
		"CPU_THRESHOLD":  server.CpuThreshold,
		"ETCD": &map[string]interface{}{
			"HOSTS": server.Etcd.Hosts,
			"KEY":   server.Etcd.Key,
		},
	}
}

// 验证 RpcServer 的字段与期望的映射是否相等
func assertRpcServerFields(t *testing.T, server *zero.RpcServer, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], server.ModuleName)
	assert.Equal(t, expected["LISTEN_ON"], server.ListenOn)
	assert.Equal(t, expected["AUTH"], server.Auth)
	assert.Equal(t, expected["STRICT_CONTROL"], server.StrictControl)
	assert.Equal(t, expected["TIMEOUT"], server.Timeout)
	assert.Equal(t, expected["CPU_THRESHOLD"], server.CpuThreshold)
	assert.Equal(t, expected["ETCD"], server.Etcd)
}

func TestRpcServerClone(t *testing.T) {
	params := generateRpcServerTestParams()
	original := zero.NewRpcServer(params)
	cloned := original.Clone().(*zero.RpcServer)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestRpcServerSet(t *testing.T) {
	oldParams := generateRpcServerTestParams()
	newParams := generateRpcServerTestParams()

	rpcServerInstance := zero.NewRpcServer(oldParams)
	newConfig := zero.NewRpcServer(newParams)

	rpcServerInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, rpcServerInstance.ModuleName)
	assert.Equal(t, newParams.ListenOn, rpcServerInstance.ListenOn)
	assert.Equal(t, newParams.Auth, rpcServerInstance.Auth)
	assert.Equal(t, newParams.StrictControl, rpcServerInstance.StrictControl)
	assert.Equal(t, newParams.Timeout, rpcServerInstance.Timeout)
	assert.Equal(t, newParams.CpuThreshold, rpcServerInstance.CpuThreshold)
	assert.Equal(t, newParams.Etcd, rpcServerInstance.Etcd)
}