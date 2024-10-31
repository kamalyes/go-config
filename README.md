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
go get -u github.com/kamalyes/go-config
```

### 例子

默认的程序根目录下必须包含 resources 文件夹

```go
package main

import (
 "fmt"
 "github.com/kamalyes/go-config"
)

func main {
    goconfig := GlobalConfig()
}


```

### 目录结构

```shell
├── internal/     # 仅供本项目使用的库代码
├── tests/         # 测试文件
├── go.mod        # Go Modules 文件
└── pkg/          # 可供其他项目使用的库代码
    ├── captcha/  # 验证码图片尺寸配置
    ├── consul/   # 注册中心配置
    ├── cors/     # 跨域配置
    ├── database/  # 数据库配置
    ├── email/    # 邮件配置
    ├── env/      # 环境变量配置
    ├── ftp/      # 文件服务器配置
    ├── global/   # 全局配置
    ├── jwt/      # JWT token 生成和校验配置
    ├── mqtt/     # MQTT 物联网配置
    ├── oss/      # OSS 配置
    ├── pay/      # 支付相关配置（支付宝和微信）
    ├── redis/      # redis(redis缓存数据库相关配置)
    ├── resources/   # 多配置文件相关配置
    │   ├── dev_config.yaml   # 开发环境配置文件
    │   ├── fat_config.yaml    # 功能验收测试环境配置文件
    │   ├── pro_config.yaml    # 生产环境配置文件
    │   └── uat_config.yaml    # 用户验收测试环境配置文件
    ├── server/   # 服务端口等相关配置
    ├── sms/      # 短信配置
    ├── sts/      # STS 配置
    ├── youzan/   # 有赞配置
    └── zap/      # 日志相关配置
```
