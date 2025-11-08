/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 21:19:10
 * @FilePath: \go-config\tests\consul_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/register"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Consul 配置参数
func generateConsulTestParams() *register.Consul {
	return &register.Consul{
		ModuleName:       "TestModule",
		Endpoint:         "127.0.0.1:8500",
		RegisterInterval: 10, // 随机间隔设置为 10 秒
	}
}

// 验证 Consul 的字段与期望的映射是否相等
func assertConsulFields(t *testing.T, consul *register.Consul, expected *register.Consul) {
	assert.Equal(t, expected.ModuleName, consul.ModuleName)
	assert.Equal(t, expected.Endpoint, consul.Endpoint)
	assert.Equal(t, expected.RegisterInterval, consul.RegisterInterval)
}

func TestNewConsul(t *testing.T) {
	params := generateConsulTestParams()
	consulInstance := register.NewConsul(params)

	assertConsulFields(t, consulInstance, params)
}

func TestConsulClone(t *testing.T) {
	params := generateConsulTestParams()
	original := register.NewConsul(params)
	cloned := original.Clone().(*register.Consul)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestConsulSet(t *testing.T) {
	oldParams := generateConsulTestParams()
	newParams := &register.Consul{
		ModuleName:       "UpdatedModule",
		Endpoint:         "192.168.1.1:8500",
		RegisterInterval: 15,
	}

	consulInstance := register.NewConsul(oldParams)
	consulInstance.Set(newParams)

	assert.Equal(t, newParams.ModuleName, consulInstance.ModuleName)
	assert.Equal(t, newParams.Endpoint, consulInstance.Endpoint)
	assert.Equal(t, newParams.RegisterInterval, consulInstance.RegisterInterval)
}
