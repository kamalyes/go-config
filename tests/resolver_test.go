/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-05 09:49:36
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
	"strings"
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
	model := &goconfig.MultiConfig{}

	resultModel, resultJson, _ := random.GenerateRandomModel(model)
	resultJson = strings.ReplaceAll(resultJson, "module_name", "modulename")
	resultJson = strings.ReplaceAll(resultJson, "_", "-")

	createConfigFile("./resources/dev_config.yaml", resultJson)

	customManager, _ := goconfig.NewConfigManager(ctx, nil)

	assert.NotNil(t, customManager)
	config := customManager.GetConfig()
	// 获取模块
	modelx := resultModel.(*goconfig.MultiConfig)
	searchKey := modelx.Server[0].ModuleName
	searchValue := modelx.Server[0].Addr

	if modelx.Server != nil {
		module, err := goconfig.GetSingleConfigByModuleName(*config, searchKey)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			assert.Equal(t, searchValue, module.Server.Addr)
			fmt.Printf("Found module: %+v\n", module)
		}
	}
}
