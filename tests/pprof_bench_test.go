/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 00:39:11
 * @FilePath: \go-config\tests\pprof_bench_test.go
 * @Description: pprof监控模块性能测试
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/register"
)

func BenchmarkNewPProf(b *testing.B) {
	params := generatePProfTestParams()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = register.NewPProf(params)
	}
}

func BenchmarkPProfConfigClone(b *testing.B) {
	params := generatePProfTestParams()
	config := register.NewPProf(params)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = config.Clone()
	}
}

func BenchmarkPProfConfigValidate(b *testing.B) {
	params := generatePProfTestParams()
	config := register.NewPProf(params)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = config.Validate()
	}
}
