/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 23:50:19
 * @FilePath: \go-config\tests\resolver_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// createConfigFile 创建配置文件并写入内容
func createConfigFile(filename string, content string) error {
	// 确保目录存在
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// 写入内容到文件
	return os.WriteFile(filename, []byte(content), 0644)
}

// 测试全局配置加载
func TestGlobalConfig(t *testing.T) {
	// 测试获取全局配置
	ctx := context.Background()
	// 使用自定义值创建 ConfigManager
	model := &goconfig.SingleConfig{}

	resultModel, resultJson, _ := random.GenerateRandomModel(model)

	createConfigFile("./resources/dev_config.yaml", resultJson)

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
	serverModel := resultModel.(*goconfig.SingleConfig)
	assert.Equal(t, serverModel.Server.Addr, module.Addr)
}
