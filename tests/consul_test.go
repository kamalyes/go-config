/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 10:06:09
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 10:25:51
 * @FilePath: \go-config\consul\consul_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/consul"
	"github.com/stretchr/testify/assert"
)

var testData = consul.Consul{
	ModuleName:       "test-module",
	Addr:             "127.0.0.1:8500",
	RegisterInterval: 10,
}

var updatedData = consul.Consul{
	ModuleName:       "new-module",
	Addr:             "192.168.1.1:8500",
	RegisterInterval: 20,
}

func TestNewConsul(t *testing.T) {
	consul := consul.NewConsul(testData.ModuleName, testData.Addr, testData.RegisterInterval)

	assert.NotNil(t, consul)
	assert.Equal(t, testData.ModuleName, consul.ModuleName)
	assert.Equal(t, testData.Addr, consul.Addr)
	assert.Equal(t, testData.RegisterInterval, consul.RegisterInterval)
}

func TestConsulToMap(t *testing.T) {
	consul := consul.NewConsul(testData.ModuleName, testData.Addr, testData.RegisterInterval)
	expectedMap := map[string]interface{}{
		"moduleName":       testData.ModuleName,
		"addr":             testData.Addr,
		"registerInterval": testData.RegisterInterval,
	}

	result := consul.ToMap()
	assert.Equal(t, expectedMap, result)
}

func TestConsulFromMap(t *testing.T) {
	consul := consul.NewConsul(testData.ModuleName, testData.Addr, testData.RegisterInterval)
	data := map[string]interface{}{
		"moduleName":       updatedData.ModuleName,
		"addr":             updatedData.Addr,
		"registerInterval": updatedData.RegisterInterval,
	}

	consul.FromMap(data)

	assert.Equal(t, updatedData.ModuleName, consul.ModuleName)
	assert.Equal(t, updatedData.Addr, consul.Addr)
	assert.Equal(t, updatedData.RegisterInterval, consul.RegisterInterval)
}

func TestConsulClone(t *testing.T) {
	consulInit := consul.NewConsul(testData.ModuleName, testData.Addr, testData.RegisterInterval)
	clone := consulInit.Clone().(*consul.Consul)

	assert.Equal(t, consulInit.ModuleName, clone.ModuleName)
	assert.Equal(t, consulInit.Addr, clone.Addr)
	assert.Equal(t, consulInit.RegisterInterval, clone.RegisterInterval)
}

func TestConsulGet(t *testing.T) {
	consulInit := consul.NewConsul(testData.ModuleName, testData.Addr, testData.RegisterInterval)
	result := consulInit.Get().(*consul.Consul)

	assert.Equal(t, consulInit.ModuleName, result.ModuleName)
	assert.Equal(t, consulInit.Addr, result.Addr)
	assert.Equal(t, consulInit.RegisterInterval, result.RegisterInterval)
}

func TestConsulSet(t *testing.T) {
	consulInit := consul.NewConsul(testData.ModuleName, testData.Addr, testData.RegisterInterval)
	newConsul := consul.NewConsul(updatedData.ModuleName, updatedData.Addr, updatedData.RegisterInterval)
	consulInit.Set(newConsul)

	assert.Equal(t, updatedData.ModuleName, consulInit.ModuleName)
	assert.Equal(t, updatedData.Addr, consulInit.Addr)
	assert.Equal(t, updatedData.RegisterInterval, consulInit.RegisterInterval)
}

func TestConsulValidate(t *testing.T) {
	validConsul := consul.NewConsul(testData.ModuleName, testData.Addr, testData.RegisterInterval)
	invalidConsul1 := consul.NewConsul("", "127.0.0.1:8500", 10)            // ModuleName 为空
	invalidConsul2 := consul.NewConsul("test-module", "", 10)               // Addr 为空
	invalidConsul3 := consul.NewConsul("test-module", "127.0.0.1:8500", -1) // RegisterInterval <= 0

	assert.NoError(t, validConsul.Validate())
	assert.NoError(t, invalidConsul1.Validate())
	assert.Error(t, invalidConsul2.Validate())
	assert.Error(t, invalidConsul3.Validate())
}
