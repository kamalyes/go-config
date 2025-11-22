/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\pprof\pprof_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package pprof

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPProf_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "PProf", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "/debug/pprof", config.PathPrefix)
	assert.Equal(t, 6060, config.Port)
	assert.NotNil(t, config.EnableProfiles)
	assert.NotNil(t, config.Sampling)
	assert.NotNil(t, config.Authentication)
	assert.NotNil(t, config.Gateway)
	assert.NotNil(t, config.WebInterface)
}

func TestPProf_WithModuleName(t *testing.T) {
	config := Default().WithModuleName("custom-pprof")
	assert.Equal(t, "custom-pprof", config.ModuleName)
}

func TestPProf_WithEnabled(t *testing.T) {
	config := Default().WithEnabled(true)
	assert.True(t, config.Enabled)
}

func TestPProf_WithPathPrefix(t *testing.T) {
	config := Default().WithPathPrefix("/profile")
	assert.Equal(t, "/profile", config.PathPrefix)
}

func TestPProf_WithPort(t *testing.T) {
	config := Default().WithPort(8080)
	assert.Equal(t, 8080, config.Port)
}

func TestPProf_WithProfiles(t *testing.T) {
	config := Default().WithProfiles(true, true, true, true, true, true, true, true, true)
	assert.True(t, config.EnableProfiles.CPU)
	assert.True(t, config.EnableProfiles.Memory)
	assert.True(t, config.EnableProfiles.Goroutine)
	assert.True(t, config.EnableProfiles.Block)
	assert.True(t, config.EnableProfiles.Mutex)
	assert.True(t, config.EnableProfiles.Heap)
	assert.True(t, config.EnableProfiles.Allocs)
	assert.True(t, config.EnableProfiles.ThreadCreate)
	assert.True(t, config.EnableProfiles.Trace)
}

func TestPProf_WithSampling(t *testing.T) {
	config := Default().WithSampling(200, 1024*1024, 2, 2)
	assert.Equal(t, 200, config.Sampling.CPURate)
	assert.Equal(t, 1024*1024, config.Sampling.MemoryRate)
	assert.Equal(t, 2, config.Sampling.BlockRate)
	assert.Equal(t, 2, config.Sampling.MutexFraction)
}

func TestPProf_EnableCPUProfile(t *testing.T) {
	config := Default()
	config.EnableProfiles.CPU = false
	config.EnableCPUProfile()
	assert.True(t, config.EnableProfiles.CPU)
}

func TestPProf_EnableMemoryProfile(t *testing.T) {
	config := Default()
	config.EnableProfiles.Memory = false
	config.EnableMemoryProfile()
	assert.True(t, config.EnableProfiles.Memory)
}

func TestPProf_EnableGoroutineProfile(t *testing.T) {
	config := Default()
	config.EnableProfiles.Goroutine = false
	config.EnableGoroutineProfile()
	assert.True(t, config.EnableProfiles.Goroutine)
}

func TestPProf_EnableBlockProfile(t *testing.T) {
	config := Default()
	config.EnableBlockProfile()
	assert.True(t, config.EnableProfiles.Block)
}

func TestPProf_EnableMutexProfile(t *testing.T) {
	config := Default()
	config.EnableMutexProfile()
	assert.True(t, config.EnableProfiles.Mutex)
}

func TestPProf_WithAuthToken(t *testing.T) {
	config := Default().WithAuthToken("test-token-123")
	assert.Equal(t, "test-token-123", config.Authentication.AuthToken)
	assert.True(t, config.Authentication.RequireAuth)
	assert.True(t, config.Authentication.Enabled)
}

func TestPProf_WithAllowedIPs(t *testing.T) {
	ips := []string{"127.0.0.1", "192.168.1.100"}
	config := Default().WithAllowedIPs(ips)
	assert.Equal(t, ips, config.Authentication.AllowedIPs)
}

func TestPProf_EnableForDevelopment(t *testing.T) {
	config := Default().EnableForDevelopment()
	assert.True(t, config.Enabled)
	assert.Equal(t, "dev-debug-token", config.Authentication.AuthToken)
	assert.True(t, config.Gateway.Enabled)
	assert.True(t, config.Gateway.DevModeOnly)
}

func TestPProf_EnableGateway(t *testing.T) {
	config := Default().EnableGateway(true, true)
	assert.True(t, config.Gateway.Enabled)
	assert.True(t, config.Gateway.DevModeOnly)
	assert.True(t, config.Gateway.EnableLogging)
	assert.True(t, config.Gateway.RegisterWebInterface)
}

func TestPProf_WithWebInterface(t *testing.T) {
	config := Default().WithWebInterface(true, "Custom Title", "Custom Description")
	assert.True(t, config.WebInterface.Enabled)
	assert.Equal(t, "Custom Title", config.WebInterface.Title)
	assert.Equal(t, "Custom Description", config.WebInterface.Description)
	assert.True(t, config.WebInterface.ShowScenarios)
}

func TestPProf_Enable(t *testing.T) {
	config := Default().Enable()
	assert.True(t, config.Enabled)
}

func TestPProf_Disable(t *testing.T) {
	config := Default().Enable().Disable()
	assert.False(t, config.Enabled)
}

func TestPProf_IsEnabled(t *testing.T) {
	config := Default()
	assert.False(t, config.IsEnabled())
	config.Enable()
	assert.True(t, config.IsEnabled())
}

func TestPProf_Clone(t *testing.T) {
	original := Default().
		WithModuleName("test-pprof").
		Enable().
		WithPathPrefix("/custom/pprof").
		WithPort(9090).
		WithAuthToken("test-token").
		WithAllowedIPs([]string{"127.0.0.1"}).
		EnableCPUProfile().
		EnableMemoryProfile()

	cloned := original.Clone().(*PProf)

	// 验证值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.PathPrefix, cloned.PathPrefix)
	assert.Equal(t, original.Port, cloned.Port)

	// 验证嵌套结构
	assert.Equal(t, original.EnableProfiles.CPU, cloned.EnableProfiles.CPU)
	assert.Equal(t, original.Sampling.CPURate, cloned.Sampling.CPURate)
	assert.Equal(t, original.Authentication.AuthToken, cloned.Authentication.AuthToken)

	// 验证切片独立性
	cloned.Authentication.AllowedIPs = append(cloned.Authentication.AllowedIPs, "192.168.1.1")
	assert.NotEqual(t, len(original.Authentication.AllowedIPs), len(cloned.Authentication.AllowedIPs))
}

func TestPProf_Get(t *testing.T) {
	config := Default().WithEnabled(true)
	got := config.Get()
	assert.NotNil(t, got)
	pprofConfig, ok := got.(*PProf)
	assert.True(t, ok)
	assert.True(t, pprofConfig.Enabled)
}

func TestPProf_Set(t *testing.T) {
	config := Default()
	newConfig := &PProf{
		ModuleName: "new-pprof",
		Enabled:    true,
		PathPrefix: "/new/pprof",
		Port:       7070,
	}

	config.Set(newConfig)
	assert.Equal(t, "new-pprof", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "/new/pprof", config.PathPrefix)
	assert.Equal(t, 7070, config.Port)
}

func TestPProf_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestPProf_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("chain-pprof").
		Enable().
		WithPathPrefix("/chain/pprof").
		WithPort(8888).
		EnableCPUProfile().
		EnableMemoryProfile().
		EnableGoroutineProfile().
		WithAuthToken("chain-token").
		EnableGateway(true, false)

	assert.Equal(t, "chain-pprof", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "/chain/pprof", config.PathPrefix)
	assert.Equal(t, 8888, config.Port)
	assert.True(t, config.EnableProfiles.CPU)
	assert.True(t, config.EnableProfiles.Memory)
	assert.True(t, config.EnableProfiles.Goroutine)
	assert.Equal(t, "chain-token", config.Authentication.AuthToken)
	assert.True(t, config.Gateway.Enabled)
}

func TestPProf_DefaultProfiles(t *testing.T) {
	config := Default()
	assert.True(t, config.EnableProfiles.CPU)
	assert.True(t, config.EnableProfiles.Memory)
	assert.True(t, config.EnableProfiles.Goroutine)
	assert.False(t, config.EnableProfiles.Block)
	assert.False(t, config.EnableProfiles.Mutex)
	assert.True(t, config.EnableProfiles.Heap)
	assert.True(t, config.EnableProfiles.Allocs)
	assert.False(t, config.EnableProfiles.ThreadCreate)
	assert.False(t, config.EnableProfiles.Trace)
}

func TestPProf_DefaultSampling(t *testing.T) {
	config := Default()
	assert.Equal(t, 100, config.Sampling.CPURate)
	assert.Equal(t, 512*1024, config.Sampling.MemoryRate)
	assert.Equal(t, 1, config.Sampling.BlockRate)
	assert.Equal(t, 1, config.Sampling.MutexFraction)
}
