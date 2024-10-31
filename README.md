# go-config

> go-config is characterized by daily work requirements and extension development, encapsulated generic tool classes

[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/kamalyes/go-config)
[![license](https://img.shields.io/github/license/kamalyes/go-config)]()
[![download](https://img.shields.io/github/downloads/kamalyes/go-config/total)]()
[![release](https://img.shields.io/github/v/release/kamalyes/go-config)]()
[![commit](https://img.shields.io/github/last-commit/kamalyes/go-config)]()
[![issues](https://img.shields.io/github/issues/kamalyes/go-config)]()
[![pull](https://img.shields.io/github/issues-pr/kamalyes/go-config)]()
[![fork](https://img.shields.io/github/forks/kamalyes/go-config)]()
[![star](https://img.shields.io/github/stars/kamalyes/go-config)]()
[![go](https://img.shields.io/github/go-mod/go-version/kamalyes/go-config)]()
[![size](https://img.shields.io/github/repo-size/kamalyes/go-config)]()
[![contributors](https://img.shields.io/github/contributors/kamalyes/go-config)]()
[![codecov](https://codecov.io/gh/kamalyes/go-config/branch/master/graph/badge.svg)](https://codecov.io/gh/kamalyes/go-config)
[![Go Report Card](https://goreportcard.com/badge/github.com/kamalyes/go-config)](https://goreportcard.com/report/github.com/kamalyes/go-config)
[![Go Reference](https://pkg.go.dev/badge/github.com/kamalyes/go-config?status.svg)](https://pkg.go.dev/github.com/kamalyes/go-config?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/kamalyes/go-config/-/badge.svg)](https://sourcegraph.com/github.com/kamalyes/go-config?badge)

### 介绍

go 开发中常用的一些配置

### 安装

```bash
go mod init "go-config-examples"
go get -u github.com/kamalyes/go-config
go mod tidy
```

### 例子

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"

	goconfig "github.com/kamalyes/go-config"

	"github.com/kamalyes/go-config/pkg/env"
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

// generateRandomConfigContent 生成随机配置内容
func generateRandomConfigContent(moduleName string) string {
	// 随机生成 IP 地址
	ip := fmt.Sprintf("192.168.1.%d", rand.Intn(256))
	// 随机生成端口号
	port := rand.Intn(10000) + 1000 // 随机端口范围 1000-10999
	// 随机生成服务名称
	serviceName := fmt.Sprintf("%s-server", moduleName)
	// 随机生成服务根路径
	contextPath := fmt.Sprintf("/%s", moduleName)

	return fmt.Sprintf(`# 服务实例配置
server:
  # 模块名称
  - modulename: "%s"
    # 服务绑定的IP
    host: "%s"
    # 端口
    port: '%d'
    # 服务名称
    service-name: "%s"
    # 服务根路径
    context-path: "%s"
    # 是否开启请求方式检测
    handle-method-not-allowed: true
    # 数据库类型
    data-driver: 'mysql'

# consul 注册中心
consul:
  # 模块名称
  - modulename: "%s01"
    # 注册中心地址
    addr: 127.0.0.1:8500
    # 间隔 单位秒
    register-interval: 30
  - modulename: "%s02"
    # 注册中心地址
    addr: 127.0.0.1:8501
    # 间隔 单位秒
    register-interval: 30

# 其他配置...
`, moduleName, ip, port, serviceName, contextPath, moduleName, moduleName)
}

// printConfigAsJSON 打印配置为 JSON 格式，并包含调用者信息
func printConfigAsJSON(config interface{}, caller string) {
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Printf("Error marshaling config to JSON: %v", err)
		return
	}
	log.Printf("Config in JSON format from %s: %s", caller, jsonData)
}

// getCallerName 获取调用者的函数名称
func getCallerName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

// envUse 使用环境变量读取配置
func envUse() *goconfig.Config {
	currentAppEnv := env.Active().String()
	content := generateRandomConfigContent(currentAppEnv)
	envFile := fmt.Sprintf("./resources/%s_config.yaml", currentAppEnv)
	err := createConfigFile(envFile, content)
	if err != nil {
		log.Fatalf("envUse Error creating %s: %#v", envFile, err)
	}
	ctx := context.Background()
	customManager, _ := goconfig.NewConfigManager(ctx, nil)
	config := customManager.GetConfig()
	printConfigAsJSON(config, getCallerName())
	return config
}

// envUse 使用上下文读取配置
func ctxUse(ctx context.Context) *goconfig.Config {
	currentAppEnv := env.Active().String()
	content := generateRandomConfigContent(currentAppEnv)
	envFile := fmt.Sprintf("./resources/%s_config.yaml", currentAppEnv)
	err := createConfigFile(envFile, content)
	if err != nil {
		log.Fatalf("ctxUse Error creating %s: %#v", envFile, err)
	}
	customManager, _ := goconfig.NewConfigManager(ctx, nil)
	config := customManager.GetConfig()
	printConfigAsJSON(config, getCallerName())
	return config
}

// envCtxUse 使用环境变量+上下文读取配置
func envCtxUse(ctx context.Context) *goconfig.Config {
	currentAppEnv := env.FromContext(ctx).String()
	content := generateRandomConfigContent(currentAppEnv)
	envFile := fmt.Sprintf("./resources/%s_config.yaml", currentAppEnv)
	err := createConfigFile(envFile, content)
	if err != nil {
		log.Fatalf("envCtxUse Error creating %s: %#v", envFile, err)
	}
	customManager, _ := goconfig.NewConfigManager(ctx, nil)
	config := customManager.GetConfig()
	printConfigAsJSON(config, getCallerName())
	return config
}

// customManagerUse 使用自定义参数+上下文读取配置
func customManagerUse(ctx context.Context, options *goconfig.ConfigOptions) *goconfig.Config {
	currentAppEnv := env.FromContext(ctx).String()
	content := generateRandomConfigContent(currentAppEnv)
	envFile := fmt.Sprintf("./%s/%s%s.%s", options.ConfigPath, currentAppEnv, options.ConfigSuffix, options.ConfigType)
	err := createConfigFile(envFile, content)
	if err != nil {
		log.Fatalf("envCtxUse Error creating %s: %#v", envFile, err)
	}
	customManager, _ := goconfig.NewConfigManager(ctx, options)
	config := customManager.GetConfig()
	printConfigAsJSON(config, getCallerName())
	return config
}

// getSingleAssignModel 当定义为数组时 可以获取指定Model
func getSingleAssignModel() {
	config := envUse()
	currentAppEnv := env.Active().String()
	modelName := fmt.Sprintf("%s01", currentAppEnv)
	consulConfig, err := goconfig.GetModuleByName(config.Consul, modelName)
	if err != nil {
		log.Fatalf("getSingleAssignModel %s: %#v", modelName, err)
	}
	printConfigAsJSON(consulConfig, fmt.Sprintf(getCallerName(), modelName))
	modelName = fmt.Sprintf("%s02", currentAppEnv)
	consulConfig, err = goconfig.GetModuleByName(config.Consul, modelName)
	if err != nil {
		log.Fatalf("getSingleAssignModel %s: %#v", modelName, err)
	}
	printConfigAsJSON(consulConfig, getCallerName())
}

func main() {
	envUse()

	singCtx := context.Background()
	ctxUse(singCtx)

	// 创建一个上下文并设置环境为 UA
	envCtx := env.NewEnv(context.Background(), env.Uat)
	envCtxUse(envCtx)

	// 创建一个上下文并设置环境变量为自定义
	customManagerCtx := env.NewEnv(envCtx, env.EnvironmentType("custom"))
	newConfigOptions := &goconfig.ConfigOptions{
		ConfigSuffix: "_env",
		ConfigPath:   "custom_resources",
		ConfigType:   "yaml",
	}
	customManagerUse(customManagerCtx, newConfigOptions)
	getSingleAssignModel()
}

```

### 目录结构

```shell
├── internal/    # 仅供本项目使用的库代码
├── tests/       # 测试文件
├── go.mod       # Go Modules 文件
├── resources/   # 多配置文件相关配置(建议命名如下,目前提供了一个example_config.yaml可参考)
│   ├── dev_config.yaml   # 开发环境配置文件
│   ├── fat_config.yaml    # 功能验收测试环境配置文件
│   ├── pro_config.yaml    # 生产环境配置文件
│   └── uat_config.yaml    # 用户验收测试环境配置文件
└── pkg/          # 可供其他项目使用的库代码
    ├── captcha/  # 验证码图片尺寸配置
    ├── register/ # 注册中心、服务端口等相关配置
    ├── cors/     # 跨域配置
    ├── database/ # 数据库配置
    ├── email/    # 邮件配置
    ├── env/      # 环境变量配置
    ├── ftp/      # 文件服务器配置
    ├── jwt/      # JWT token 生成和校验配置
    ├── mqtt/     # MQTT 物联网配置
    ├── oss/      # OSS 配置
    ├── pay/      # 支付相关配置（支付宝和微信）
    ├── redis/    # redis(redis缓存数据库相关配置)
    ├── sms/      # 短信配置
    ├── sts/      # STS 配置
    ├── youzan/   # 有赞配置
    └── zap/      # 日志相关配置
```
