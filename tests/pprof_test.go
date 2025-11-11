/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 03:11:28
 * @FilePath: \go-config\tests\pprof_test.go
 * @Description: pprof监控模块测试
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/pprof"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 PProf 配置参数
func generatePProfTestParams() *pprof.PProf {
	return &pprof.PProf{
		Enabled: true,
		PathPrefix:    "/debug/pprof",
		Port:    6060,
		EnableProfiles: &pprof.ProfilesConfig{
			CPU:          true,
			Memory:       true,
			Goroutine:    true,
			Block:        false,
			Mutex:        false,
			Heap:         true,
			Allocs:       true,
			ThreadCreate: false,
			Trace:        false,
		},
		Sampling: &pprof.SamplingConfig{
			CPURate:       100,
			MemoryRate:    512 * 1024, // 512KB
			BlockRate:     1,
			MutexFraction: 1,
		},
	}
}

// 验证 PProf 的字段与期望的映射是否相等
func assertPProfFields(t *testing.T, actual *pprof.PProf, expected *pprof.PProf) {
	assert.Equal(t, expected.Enabled, actual.Enabled)
	assert.Equal(t, expected.PathPrefix, actual.PathPrefix)
	assert.Equal(t, expected.Port, actual.Port)
	assert.Equal(t, expected.EnableProfiles, actual.EnableProfiles)
	assert.Equal(t, expected.Sampling, actual.Sampling)
}
