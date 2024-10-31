/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 11:35:54
 * @FilePath: \go-config\resolver.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试全局配置加载
func TestGlobalConfig(t *testing.T) {
	// 测试获取全局配置
	config := GlobalConfig()
	assert.NotNil(t, config)
	// 获取模块
	module, err := GetModuleByName(config.Server, "common")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found module: %+v\n", module)
	}
	assert.Equal(t, "0.0.0.0", module.Host)
	assert.Equal(t, "8880", module.Port)
}
