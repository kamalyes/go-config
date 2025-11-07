/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 00:41:08
 * @FilePath: \go-config\tests\pprof_test.go
 * @Description: pprof监控模块测试
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/kamalyes/go-config/pkg/register"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 PProf 配置参数
func generatePProfTestParams() *register.PProf {
	return &register.PProf{
		Enabled:        true,
		PathPrefix:     "/debug/pprof",
		AllowedIPs:     []string{"127.0.0.1", "::1"},
		RequireAuth:    random.FRandInt(0, 1) == 1, // 随机认证设置
		AuthToken:      random.RandString(32, random.LOWERCASE),
		EnableLogging:  true,
		Timeout:        random.FRandInt(10, 60),
		CustomHandlers: make(map[string]http.HandlerFunc),
	}
}

// 验证 PProf 的字段与期望的映射是否相等
func assertPProfFields(t *testing.T, actual *register.PProf, expected *register.PProf) {
	assert.Equal(t, expected.Enabled, actual.Enabled)
	assert.Equal(t, expected.PathPrefix, actual.PathPrefix)
	assert.Equal(t, expected.AllowedIPs, actual.AllowedIPs)
	assert.Equal(t, expected.RequireAuth, actual.RequireAuth)
	assert.Equal(t, expected.AuthToken, actual.AuthToken)
	assert.Equal(t, expected.EnableLogging, actual.EnableLogging)
	assert.Equal(t, expected.Timeout, actual.Timeout)
}

func TestNewPProf(t *testing.T) {
	params := generatePProfTestParams()
	pprofInstance := register.NewPProf(params)

	assertPProfFields(t, pprofInstance, params)
}

func TestNewPProfWithDefaults(t *testing.T) {
	pprofInstance := register.NewPProf(nil)

	// 验证默认值
	assert.True(t, pprofInstance.PathPrefix != "")
	assert.True(t, pprofInstance.Timeout > 0)
}

func TestPProfClone(t *testing.T) {
	params := generatePProfTestParams()
	// 添加一个自定义处理器用于测试
	params.CustomHandlers["test"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test")
	}

	original := register.NewPProf(params)
	cloned := original.Clone().(*register.PProf)

	assertPProfFields(t, cloned, original)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestPProfSet(t *testing.T) {
	oldParams := generatePProfTestParams()
	newParams := generatePProfTestParams()

	pprofInstance := register.NewPProf(oldParams)
	pprofInstance.Set(newParams)

	assertPProfFields(t, pprofInstance, newParams)
}

func TestPProfValidate(t *testing.T) {
	// 测试有效配置
	validParams := &register.PProf{
		PathPrefix:     "/debug/pprof",
		Timeout:        30,
	}
	
	pprofInstance := register.NewPProf(validParams)
	err := pprofInstance.Validate()
	assert.NoError(t, err)
}