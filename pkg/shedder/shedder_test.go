/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-11-01 17:11:06
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-11-01 21:30:15
 * @FilePath: \go-config\pkg\shedder\shedder_test.go
 * @Description: 自适应降载配置测试
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package shedder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShedder_Default(t *testing.T) {
	shedder := Default()

	assert.NotNil(t, shedder)
	assert.Equal(t, "shedder", shedder.ModuleName)
	// 降载默认关闭：主动丢流量的策略必须由使用方显式开启
	assert.False(t, shedder.Enabled)
	assert.Equal(t, int64(900), shedder.CPUThreshold)
	assert.Equal(t, int64(0), shedder.MaxConcurrency)
	assert.Equal(t, int64(5000), shedder.WindowMS)
	assert.Equal(t, 50, shedder.WindowBuckets)
	assert.Equal(t, int64(1000), shedder.CooldownMS)
	assert.Equal(t, int64(1000), shedder.MinRTMS)
	assert.Equal(t, 0.9, shedder.FlyingBeta)
	assert.Equal(t, []string{"/health", "/healthz", "/readyz", "/livez", "/metrics", "/debug"}, shedder.ExcludePaths)
}

func TestShedder_GetLabel(t *testing.T) {
	shedder := Default()
	assert.Equal(t, "shedder", shedder.GetLabel())

	shedder.ModuleName = "custom_shedder"
	assert.Equal(t, "custom_shedder", shedder.GetLabel())
}

func TestShedder_Init(t *testing.T) {
	shedder := Default()
	assert.NoError(t, shedder.Init())
}

func TestShedder_Validate(t *testing.T) {
	shedder := Default()
	assert.NoError(t, shedder.Validate())
}

func TestShedder_EnableDisable(t *testing.T) {
	shedder := Default()
	assert.False(t, shedder.IsEnabled())

	shedder.Enable()
	assert.True(t, shedder.IsEnabled())

	shedder.Disable()
	assert.False(t, shedder.IsEnabled())
}

func TestShedder_Clone(t *testing.T) {
	shedder := Default().Enable().WithCPUThreshold(850)

	cloned, ok := shedder.Clone().(*Shedder)
	assert.True(t, ok)
	assert.Equal(t, shedder.CPUThreshold, cloned.CPUThreshold)
	assert.Equal(t, shedder.Enabled, cloned.Enabled)
	assert.Equal(t, shedder.ExcludePaths, cloned.ExcludePaths)

	// 副本必须独立：改副本不影响原配置
	cloned.CPUThreshold = 500
	cloned.ExcludePaths = append(cloned.ExcludePaths, "/extra")
	assert.Equal(t, int64(850), shedder.CPUThreshold)
	assert.NotContains(t, shedder.ExcludePaths, "/extra")
}

func TestShedder_GetSet(t *testing.T) {
	shedder := Default()
	assert.Equal(t, shedder, shedder.Get())

	other := Default().WithModuleName("other").WithMaxConcurrency(256)
	shedder.Set(other)
	assert.Equal(t, "other", shedder.ModuleName)
	assert.Equal(t, int64(256), shedder.MaxConcurrency)

	// 类型不符时不应改动原配置
	shedder.Set("not-a-shedder")
	assert.Equal(t, "other", shedder.ModuleName)
}

func TestShedder_Builders(t *testing.T) {
	shedder := Default().
		WithModuleName("custom").
		WithCPUThreshold(800).
		WithMaxConcurrency(1024).
		WithWindow(3000, 30).
		WithCooldownMS(1500).
		WithMinRTMS(600).
		WithFlyingBeta(0.8).
		WithExcludePaths([]string{"/status"})

	assert.Equal(t, "custom", shedder.ModuleName)
	assert.Equal(t, int64(800), shedder.CPUThreshold)
	assert.Equal(t, int64(1024), shedder.MaxConcurrency)
	assert.Equal(t, int64(3000), shedder.WindowMS)
	assert.Equal(t, 30, shedder.WindowBuckets)
	assert.Equal(t, int64(1500), shedder.CooldownMS)
	assert.Equal(t, int64(600), shedder.MinRTMS)
	assert.Equal(t, 0.8, shedder.FlyingBeta)
	assert.Equal(t, []string{"/status"}, shedder.ExcludePaths)
}

func TestShedder_ExcludePathsCoverProbeEndpoints(t *testing.T) {
	shedder := Default()

	// 这几个端点被降载会直接导致编排系统误判实例已死
	for _, path := range []string{"/health", "/healthz", "/readyz", "/livez", "/metrics"} {
		assert.Contains(t, shedder.ExcludePaths, path)
	}
}
