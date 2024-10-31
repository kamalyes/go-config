/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 13:57:26
 * @FilePath: \go-config\tests\server_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/register"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Server 配置参数
func generateServerTestParams() *register.Server {
	return &register.Server{
		ModuleName:             random.RandString(10, random.CAPITAL),
		Host:                   random.RandString(5, random.CAPITAL) + ".example.com", // 随机生成主机名
		Port:                   random.RandString(4, random.NUMBER),                   // 随机生成端口
		ServerName:             random.RandString(8, random.CAPITAL),                  // 随机生成服务名称
		ContextPath:            random.RandString(5, random.CAPITAL),                  // 随机生成请求根路径
		DataDriver:             random.RandString(5, random.CAPITAL),                  // 随机生成数据库类型
		HandleMethodNotAllowed: random.FRandBool(),                                    // 随机生成是否开启请求方式检测
		Language:               random.RandString(2, random.CAPITAL),                  // 随机生成语言
	}
}

func TestServerClone(t *testing.T) {
	params := generateServerTestParams()
	original := register.NewServer(params)
	cloned := original.Clone().(*register.Server)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestServerSet(t *testing.T) {
	oldParams := generateServerTestParams()
	newParams := generateServerTestParams()

	serverInstance := register.NewServer(oldParams)
	newConfig := register.NewServer(newParams)

	serverInstance.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, serverInstance.ModuleName)
	assert.Equal(t, newParams.Host, serverInstance.Host)
	assert.Equal(t, newParams.Port, serverInstance.Port)
	assert.Equal(t, newParams.ServerName, serverInstance.ServerName)
	assert.Equal(t, newParams.ContextPath, serverInstance.ContextPath)
	assert.Equal(t, newParams.DataDriver, serverInstance.DataDriver)
	assert.Equal(t, newParams.ContextPath, serverInstance.ContextPath)
	assert.Equal(t, newParams.HandleMethodNotAllowed, serverInstance.HandleMethodNotAllowed)
	assert.Equal(t, newParams.Language, serverInstance.Language)
}
