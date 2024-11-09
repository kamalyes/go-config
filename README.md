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
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/env"
)

// generateRandomConfigContent 生成随机配置内容
func generateRandomConfigContent(moduleName string) string {
	ip := fmt.Sprintf("192.168.1.%d", rand.Intn(256))   // 随机生成 IP 地址
	port := rand.Intn(10000) + 1000                     // 随机端口范围 1000-10999
	serviceName := fmt.Sprintf("%s-server", moduleName) // 服务名称
	contextPath := fmt.Sprintf("/%s", moduleName)       // 服务根路径

	return fmt.Sprintf(`# 服务实例配置
server:
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
    # 注册中心地址
    addr: 127.0.0.1:8500
    # 间隔 单位秒
    register-interval: 30

# 其他配置...
`, ip, port, serviceName, contextPath)
}

// createConfigFile 创建配置文件并写入内容
func createConfigFile(configOptions *goconfig.ConfigOptions) error {
	content := generateRandomConfigContent(configOptions.EnvValue.String()) // 生成随机配置内容
	// 确保目录存在
	filename := fmt.Sprintf("%s/%s%s.%s", configOptions.ConfigPath, configOptions.EnvValue, configOptions.ConfigSuffix, configOptions.ConfigType)
	dir := filepath.Dir(filename)                         // 获取文件目录
	if err := os.MkdirAll(dir, os.ModePerm); err != nil { // 创建目录
		return err
	}

	// 写入内容到文件
	return os.WriteFile(filename, []byte(content), 0644) // 写入配置文件
}

// printConfig 打印配置为 JSON 格式，并包含调用者信息
func printConfig(config interface{}, caller string) {
	log.Printf("Config in format from %s: %#v", caller, config) // 打印配置内容
}

// getCallerName 获取调用者的函数名称
func getCallerName() string {
	pc, _, _, ok := runtime.Caller(1) // 获取调用者信息
	if !ok {
		return "unknown" // 如果获取失败，返回 unknown
	}
	fn := runtime.FuncForPC(pc) // 获取函数信息
	if fn == nil {
		return "unknown" // 如果函数信息为空，返回 unknown
	}
	return fn.Name() // 返回函数名称
}

// simpleUse 简单例子
func simpleUse() *goconfig.SingleConfig {
	ctx := context.Background()                                     // 创建上下文
	customManager, err := goconfig.NewSingleConfigManager(ctx, nil) // 创建单个配置管理器
	if err != nil {
		log.Fatalf("simpleUse NewSingleConfigManager  err %s", err) // 错误处理
	}
	err = createConfigFile(&customManager.Options) // 创建配置文件
	if err != nil {
		log.Fatalf("simpleUse Error creating file %#v", err) // 错误处理
	}
	config := customManager.GetConfig()  // 获取配置
	printConfig(config, getCallerName()) // 打印配置
	return config
}

// customUse 自定义环境变量、读取配置
func customUse() *goconfig.SingleConfig {
	customEnv := env.EnvironmentType("custom_single")        // 定义自定义环境类型
	customContextKey := env.ContextKey("TEST_SINGLE_CONFIG") // 定义自定义上下文键
	env.NewEnvironment().
		SetCheckFrequency(1 * time.Second) // 初始化检测时间

	configOptions := &goconfig.ConfigOptions{
		ConfigSuffix:  "_config",            // 配置后缀
		ConfigPath:    "./custom_resources", // 配置路径
		ConfigType:    "yaml",               // 配置类型
		EnvValue:      customEnv,            // 环境变量值
		EnvContextKey: customContextKey,     // 环境变量Key
	}

	err := createConfigFile(configOptions) // 创建配置文件
	if err != nil {
		log.Fatalf("customUse Error creating file %#v", err) // 错误处理
	}
	ctx := context.Background()                                               // 创建上下文
	customManager, err := goconfig.NewSingleConfigManager(ctx, configOptions) // 创建单个配置管理器
	if err != nil {
		log.Fatalf("customUse NewSingleConfigManager err %s", err) // 错误处理
	}
	config := customManager.GetConfig()  // 获取配置
	printConfig(config, getCallerName()) // 打印配置
	return config
}

func main() {
	simpleUse() // 调用简单使用示例
	customUse() // 调用自定义使用示例
}

```

### 目录结构

```shell
├── internal/    # 仅供本项目使用的库代码
├── tests/       # 测试文件
├── go.mod       # Go Modules 文件
├── resources/   # 多配置文件相关配置(建议命名如下,目前提供了一个example_config.yaml可参考)
│   ├── dev_config.yaml    # 开发环境配置文件
│   ├── sit_config.yaml    # 功能测试环境配置文件
│   └── uat_config.yaml    # 产品验收测试环境配置文件
│   └── fat_config.yaml    # 预留环境配置文件
│   └── pro_config.yaml    # 生产环境配置文件
│   └── env_custom.yaml    # 当然还通过了个性化配置文件需要结合ConfigManager一起使用
└── pkg/          # 可供其他项目使用的库代码
    ├── captcha/  # 验证码图片尺寸配置
    ├── register/ # 注册中心、服务端口等相关配置
    ├── cors/     # 跨域配置
    ├── database/ # 数据库配置
    ├── elk/      # ELK配置
    ├── email/    # 邮件配置
    ├── env/      # 环境变量配置
    ├── ftp/      # 文件服务器配置
    ├── jwt/      # JWT token 生成和校验配置
    ├── mqtt/     # MQTT 物联网配置
    ├── oss/      # OSS 配置
    ├── queue/    # mqtt等队列相关配置
    ├── pay/      # 支付相关配置（支付宝和微信）
    ├── redis/    # redis(redis缓存数据库相关配置)
    ├── sms/      # 短信配置
    ├── sts/      # STS 配置
    ├── youzan/   # 有赞配置
    └── zap/      # 日志相关配置
    └── zero/     # zero相关配置
```
