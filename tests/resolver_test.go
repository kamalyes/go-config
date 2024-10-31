/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 13:06:18
 * @FilePath: \go-config\resolver.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"context"
	"fmt"
	"testing"

	goconfig "github.com/kamalyes/go-config"
	"github.com/stretchr/testify/assert"
)

// 测试全局配置加载
func TestGlobalConfig(t *testing.T) {
	// 测试获取全局配置
	ctx := context.Background()
	// 使用自定义值创建 ConfigManager
	customManager, _ := goconfig.NewConfigManager(ctx, nil)

	assert.NotNil(t, customManager)
	config := customManager.GetConfig()
	// 获取模块
	module, err := goconfig.GetModuleByName(config.Server, "")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found module: %+v\n", module)
	}
	assert.Equal(t, "0.0.0.0", module.Host)
	assert.Equal(t, "8880", module.Port)
}
