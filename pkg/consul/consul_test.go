/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-21 23:59:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-21 23:59:00
 * @FilePath: \go-config\pkg\consul\consul_test.go
 * @Description: Consul配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package consul

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsul_DefaultConsulConfig(t *testing.T) {
	consul := DefaultConsulConfig()

	assert.NotNil(t, consul)
	assert.Equal(t, "consul", consul.ModuleName)
	assert.Equal(t, "http://127.0.0.1:8500", consul.Endpoint)
	assert.Equal(t, 10, consul.RegisterInterval)
}

func TestConsul_WithModuleName(t *testing.T) {
	consul := DefaultConsulConfig().WithModuleName("custom_consul")
	assert.Equal(t, "custom_consul", consul.ModuleName)
}

func TestConsul_WithEndpoint(t *testing.T) {
	consul := DefaultConsulConfig().WithEndpoint("http://192.168.1.100:8500")
	assert.Equal(t, "http://192.168.1.100:8500", consul.Endpoint)
}

func TestConsul_WithRegisterInterval(t *testing.T) {
	consul := DefaultConsulConfig().WithRegisterInterval(30)
	assert.Equal(t, 30, consul.RegisterInterval)
}

func TestConsul_Clone(t *testing.T) {
	original := DefaultConsulConfig()
	original.Endpoint = "http://custom:8500"
	original.RegisterInterval = 20

	cloned := original.Clone().(*Consul)

	assert.Equal(t, original.Endpoint, cloned.Endpoint)
	assert.Equal(t, original.RegisterInterval, cloned.RegisterInterval)

	cloned.RegisterInterval = 40
	assert.Equal(t, 20, original.RegisterInterval)
	assert.Equal(t, 40, cloned.RegisterInterval)
}

func TestConsul_Get(t *testing.T) {
	consul := DefaultConsulConfig()
	result := consul.Get()

	assert.NotNil(t, result)
	resultConsul, ok := result.(*Consul)
	assert.True(t, ok)
	assert.Equal(t, consul, resultConsul)
}

func TestConsul_Set(t *testing.T) {
	consul := DefaultConsulConfig()
	newConsul := &Consul{
		ModuleName:       "new_consul",
		Endpoint:         "http://new:8500",
		RegisterInterval: 60,
	}

	consul.Set(newConsul)

	assert.Equal(t, "new_consul", consul.ModuleName)
	assert.Equal(t, "http://new:8500", consul.Endpoint)
	assert.Equal(t, 60, consul.RegisterInterval)
}

func TestConsul_Validate(t *testing.T) {
	consul := DefaultConsulConfig()
	err := consul.Validate()
	assert.NoError(t, err)
}

func TestConsul_Validate_Invalid(t *testing.T) {
	consul := &Consul{
		ModuleName:       "",
		Endpoint:         "invalid-url",
		RegisterInterval: 0,
	}

	err := consul.Validate()
	assert.Error(t, err)
}

func TestConsul_ChainedCalls(t *testing.T) {
	consul := DefaultConsulConfig().
		WithModuleName("chained").
		WithEndpoint("http://chained:8500").
		WithRegisterInterval(15)

	assert.Equal(t, "chained", consul.ModuleName)
	assert.Equal(t, "http://chained:8500", consul.Endpoint)
	assert.Equal(t, 15, consul.RegisterInterval)

	err := consul.Validate()
	assert.NoError(t, err)
}

func TestNewConsul(t *testing.T) {
	opt := &Consul{
		ModuleName:       "test",
		Endpoint:         "http://test:8500",
		RegisterInterval: 25,
	}

	consul := NewConsul(opt)
	assert.NotNil(t, consul)
	assert.Equal(t, opt, consul)
}
