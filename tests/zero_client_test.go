/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 09:55:52
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 17:08:58
 * @FilePath: \go-config\tests\zero_client_test.go
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

// 生成随机的 RpcClient 配置参数
func generateRpcClientTestParams() *zero.RpcClient {
	return &zero.RpcClient{
		ModuleName: random.RandString(10, random.CAPITAL),                                      // 随机生成模块名称
		Target:     fmt.Sprintf("http://%s.example.com", random.RandString(5, random.CAPITAL)), // 随机生成目标地址
		App:        random.RandString(10, random.CAPITAL),                                      // 随机生成应用名称
		Token:      random.RandString(32, random.CAPITAL),                                      // 随机生成认证令牌
		NonBlock:   random.FRandBool(),                                                         // 随机生成是否非阻塞
		Timeout:    int64(random.RandInt(100, 5000)),                                           // 随机生成超时时间（单位：毫秒）
	}
}

// 将 RpcClient 的参数转换为 map
func rpcClientToMap(client *zero.RpcClient) map[string]interface{} {
	return map[string]interface{}{
		"MODULE_NAME": client.ModuleName,
		"TARGET":      client.Target,
		"APP":         client.App,
		"TOKEN":       client.Token,
		"NON_BLOCK":   client.NonBlock,
		"TIMEOUT":     client.Timeout,
	}
}

// 验证 RpcClient 的字段与期望的映射是否相等
func assertRpcClientFields(t *testing.T, client *zero.RpcClient, expected map[string]interface{}) {
	assert.Equal(t, expected["MODULE_NAME"], client.ModuleName)
	assert.Equal(t, expected["TARGET"], client.Target)
	assert.Equal(t, expected["APP"], client.App)
	assert.Equal(t, expected["TOKEN"], client.Token)
	assert.Equal(t, expected["NON_BLOCK"], client.NonBlock)
	assert.Equal(t, expected["TIMEOUT"], client.Timeout)
}

func TestRpcClientClone(t *testing.T) {
	params := generateRpcClientTestParams()
	original := zero.NewRpcClient(params)
	cloned := original.Clone().(*zero.RpcClient)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestRpcClientSet(t *testing.T) {
	oldParams := generateRpcClientTestParams()
	newParams := generateRpcClientTestParams()

	rpcClientInstance := zero.NewRpcClient(oldParams)
	newConfig := zero.NewRpcClient(newParams)

	rpcClientInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, rpcClientInstance.ModuleName)
	assert.Equal(t, newParams.Target, rpcClientInstance.Target)
	assert.Equal(t, newParams.App, rpcClientInstance.App)
	assert.Equal(t, newParams.Token, rpcClientInstance.Token)
	assert.Equal(t, newParams.NonBlock, rpcClientInstance.NonBlock)
	assert.Equal(t, newParams.Timeout, rpcClientInstance.Timeout)
}
