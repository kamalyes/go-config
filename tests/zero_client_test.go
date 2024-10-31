/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 09:55:52
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 13:15:57
 * @FilePath: \go-config\zero\client_test.go
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

// generateZeroRpcClientTestParams 生成测试参数
func generateZeroRpcClientTestParams() zero.RpcClient {
	return zero.RpcClient{
		ModuleName: "test-module",
		Target:     "http://localhost:8080",
		App:        "test-app",
		Token:      "test-token",
		NonBlock:   true,
		Timeout:    int64(1000),
	}
}

func TestNewRpcClient(t *testing.T) {
	params := generateZeroRpcClientTestParams()
	rpcClient := zero.NewRpcClient(params.ModuleName, params.Target, params.App, params.Token, params.NonBlock, params.Timeout)

	assert.NotNil(t, rpcClient)
	assert.Equal(t, params.ModuleName, rpcClient.ModuleName)
	assert.Equal(t, params.Target, rpcClient.Target)
	assert.Equal(t, params.App, rpcClient.App)
	assert.Equal(t, params.Token, rpcClient.Token)
	assert.Equal(t, params.NonBlock, rpcClient.NonBlock)
	assert.Equal(t, params.Timeout, rpcClient.Timeout)
}

func TestRpcClientValidate(t *testing.T) {
	validParams := generateZeroRpcClientTestParams()
	validRpcClient := zero.NewRpcClient(validParams.ModuleName, validParams.Target, validParams.App, validParams.Token, validParams.NonBlock, validParams.Timeout)
	assert.NoError(t, validRpcClient.Validate())

	invalidRpcClient := zero.NewRpcClient("", "", "", "", false, -1)
	assert.Error(t, invalidRpcClient.Validate())
	assert.EqualError(t, invalidRpcClient.Validate(), "module name cannot be empty")
}

func TestRpcClientClone(t *testing.T) {
	params := generateZeroRpcClientTestParams()
	original := zero.NewRpcClient(params.ModuleName, params.Target, params.App, params.Token, params.NonBlock, params.Timeout)
	cloned := original.Clone().(*zero.RpcClient)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestRpcClientToMap(t *testing.T) {
	params := generateZeroRpcClientTestParams()
	rpcClient := zero.NewRpcClient(params.ModuleName, params.Target, params.App, params.Token, params.NonBlock, params.Timeout)
	expectedMap := map[string]interface{}{
		"moduleName": params.ModuleName,
		"target":     params.Target,
		"app":        params.App,
		"token":      params.Token,
		"nonBlock":   params.NonBlock,
		"timeout":    params.Timeout,
	}
	assert.Equal(t, expectedMap, rpcClient.ToMap())
}

func TestRpcClientFromMap(t *testing.T) {
	rpcClient := zero.NewRpcClient("", "", "", "", false, 0)
	data := map[string]interface{}{
		"moduleName": "new-module",
		"target":     "http://new-target:8080",
		"app":        "new-app",
		"token":      "new-token",
		"nonBlock":   false,
		"timeout":    int64(2000),
	}
	rpcClient.FromMap(data)

	// 验证填充后的数据是否正确
	assert.Equal(t, "new-module", rpcClient.ModuleName)
	assert.Equal(t, "http://new-target:8080", rpcClient.Target)
	assert.Equal(t, "new-app", rpcClient.App)
	assert.Equal(t, "new-token", rpcClient.Token)
	assert.Equal(t, false, rpcClient.NonBlock)
	assert.Equal(t, int64(2000), rpcClient.Timeout)
}

func TestRpcClientSet(t *testing.T) {
	oldParams := generateZeroRpcClientTestParams()
	newParams := zero.RpcClient{
		ModuleName: "new-module",
		Target:     "http://new-target:8080",
		App:        "new-app",
		Token:      "new-token",
		NonBlock:   false,
		Timeout:    2000,
	}

	rpcClient := zero.NewRpcClient(oldParams.ModuleName, oldParams.Target, oldParams.App, oldParams.Token, oldParams.NonBlock, oldParams.Timeout)
	newConfig := zero.NewRpcClient(newParams.ModuleName, newParams.Target, newParams.App, newParams.Token, newParams.NonBlock, newParams.Timeout)

	rpcClient.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, rpcClient.ModuleName)
	assert.Equal(t, newParams.Target, rpcClient.Target)
	assert.Equal(t, newParams.App, rpcClient.App)
	assert.Equal(t, newParams.Token, rpcClient.Token)
	assert.Equal(t, newParams.NonBlock, rpcClient.NonBlock)
	assert.Equal(t, newParams.Timeout, rpcClient.Timeout)
}
