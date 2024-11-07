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
		Host:       fmt.Sprintf("http://%s.example.com", random.RandString(5, random.CAPITAL)), // 随机生成目标地址
		Port:       int64(random.RandInt(100, 5000)),                                           // 随机生成目标端口
		App:        random.RandString(10, random.CAPITAL),                                      // 随机生成应用名称
		Token:      random.RandString(32, random.CAPITAL),                                      // 随机生成认证令牌
		NonBlock:   random.FRandBool(),                                                         // 随机生成是否非阻塞
		Timeout:    int64(random.RandInt(100, 5000)),                                           // 随机生成超时时间（单位：毫秒）
	}
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
	assert.Equal(t, newParams.Host, rpcClientInstance.Host)
	assert.Equal(t, newParams.Port, rpcClientInstance.Port)
	assert.Equal(t, newParams.App, rpcClientInstance.App)
	assert.Equal(t, newParams.Token, rpcClientInstance.Token)
	assert.Equal(t, newParams.NonBlock, rpcClientInstance.NonBlock)
	assert.Equal(t, newParams.Timeout, rpcClientInstance.Timeout)
}
