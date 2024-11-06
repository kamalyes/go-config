/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 13:55:55
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
	"github.com/kamalyes/go-config/pkg/env"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// assertEqualSingleConfig 用于比较两个 SingleConfig 对象的字段
func assertEqualSingleConfig(t *testing.T, expected, actual *goconfig.SingleConfig) {
	assert.Equal(t, expected.Server, actual.Server)
	assert.Equal(t, expected.Cors, actual.Cors)
	assert.Equal(t, expected.Consul, actual.Consul)
	assert.Equal(t, expected.Captcha, actual.Captcha)
	assert.Equal(t, expected.MySQL, actual.MySQL)
	assert.Equal(t, expected.PostgreSQL, actual.PostgreSQL)
	assert.Equal(t, expected.SQLite, actual.SQLite)
	assert.Equal(t, expected.Redis, actual.Redis)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.Ftp, actual.Ftp)
	assert.Equal(t, expected.JWT, actual.JWT)
	assert.Equal(t, expected.Minio, actual.Minio)
	assert.Equal(t, expected.AliyunOss, actual.AliyunOss)
	assert.Equal(t, expected.Mqtt, actual.Mqtt)
	assert.Equal(t, expected.Zap, actual.Zap)
	assert.Equal(t, expected.AliPay, actual.AliPay)
	assert.Equal(t, expected.WechatPay, actual.WechatPay)
	assert.Equal(t, expected.AliyunSms, actual.AliyunSms)
	assert.Equal(t, expected.AliyunSts, actual.AliyunSts)
	assert.Equal(t, expected.Youzan, actual.Youzan)
	assert.Equal(t, expected.ZeroServer, actual.ZeroServer)
	assert.Equal(t, expected.ZeroClient, actual.ZeroClient)
}

// createConfigFile 创建配置文件并写入内容
func createConfigFile(filename string, content string) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(filename, []byte(content), 0644)
}

// generateRandomConfigFile 生成随机模型并创建配置文件
func generateRandomConfigFile(model *goconfig.SingleConfig, filePath string) (*goconfig.SingleConfig, error) {
	resultModel, resultJson, err := random.GenerateRandomModel(model)
	if err != nil {
		return nil, err
	}
	resultJson = strings.ReplaceAll(resultJson, "module_name", "modulename")
	resultJson = strings.ReplaceAll(resultJson, "_", "-")
	if err := createConfigFile(filePath, resultJson); err != nil {
		return nil, err
	}
	return resultModel.(*goconfig.SingleConfig), nil
}

func assertEnvironment(t *testing.T, expected env.EnvironmentType) {
	currentOsEnv := env.GetEnvironment()
	assert.Equal(t, expected, currentOsEnv, "期望环境为 %s, 实际为 %s", expected, currentOsEnv)
}

// 测试全局配置加载
func TestDefaultSingleConfig(t *testing.T) {
	// 默认环境变量
	ctx := context.Background() // 创建一个新的背景上下文
	// 生成配置文件
	customModels, err := generateRandomConfigFile(&goconfig.SingleConfig{}, fmt.Sprintf("./resources/%s_config.yaml", env.GetEnvironment()))
	if err != nil {
		t.Fatalf("Error generating custom config: %v", err)
	}

	customManager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		t.Fatalf("Error creating config manager: %v", err)
	}

	assert.NotNil(t, customManager)
	config := customManager.GetConfig()
	assertEqualSingleConfig(t, customModels, config)
}
