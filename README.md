# go-config

[![Go Reference](https://pkg.go.dev/badge/github.com/kamalyes/go-config.svg)](https://pkg.go.dev/github.com/kamalyes/go-config)
[![Go Report Card](https://goreportcard.com/badge/github.com/kamalyes/go-config)](https://goreportcard.com/report/github.com/kamalyes/go-config)
[![Tests](https://github.com/kamalyes/go-config/actions/workflows/test.yml/badge.svg)](https://github.com/kamalyes/go-config/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

一个功能强大且易于使用的Go配置管理库，支持多种配置格式、智能发现、热更新和安全访问。为第三方开发者提供开箱即用的配置管理解决方案。

## ✨ 核心特性

- 🔧 **多格式支持** - 支持YAML、JSON、TOML等多种配置格式
- 🔥 **配置热更新** - 实时监听配置文件变化并自动重载
- 🛡️ **安全访问** - 防止空指针异常的链式配置访问
- 🎯 **智能发现** - 自动发现和加载配置文件
- 📦 **丰富模块** - 内置40+配置模块，覆盖常见应用场景
- 🚀 **零配置启动** - 开箱即用的默认配置
- 🎨 **链式API** - 优雅的构建器模式API设计

## 🚀 快速开始

### 安装

```bash
go get github.com/kamalyes/go-config
```

### 基础使用 - 配置热更新

```go
package main

import (
    "fmt"
    "time"
    
    goconfig "github.com/kamalyes/go-config"
    "github.com/kamalyes/go-config/pkg/gateway"
)

func main() {
    // 初始化HTTPServer配置
    config := gateway.DefaultHTTPServer()
    
    // 配置热更新回调
    hotReloadConfig := &goconfig.HotReloadConfig{
        Enabled: true,
        OnReloaded: func(oldConfig, newConfig interface{}) {
            fmt.Printf("配置已更新: %+v -> %+v\n", oldConfig, newConfig)
        },
        OnError: func(err error) {
            fmt.Printf("热更新错误: %v\n", err)
        },
    }
    
    // 创建并启动配置管理器
    manager := goconfig.NewConfigBuilder(config).
        WithConfigPath("config.yaml").
        WithEnvironment(goconfig.EnvDevelopment).
        WithHotReload(hotReloadConfig).
        MustBuildAndStart()
    
    defer manager.Stop()
    
    // 使用安全配置访问
    safeConfig := goconfig.SafeConfig(config)
    
    fmt.Printf("HTTP服务器启动在 %s:%d\n", 
        safeConfig.Host("localhost"), 
        safeConfig.Port(8080))
    
    fmt.Printf("启用HTTP: %v\n", 
        safeConfig.Field("EnableHttp").Bool(true))
    
    // 保持程序运行以观察热更新
    select {
    case <-time.After(time.Minute * 5):
        fmt.Println("程序退出")
    }
}
```

### 创建配置文件 `config.yaml`

```yaml
# HTTP服务器配置 - 注意字段名使用横线格式
module-name: "my-app-server"
host: "0.0.0.0" 
port: 8080
grpc-port: 9090
read-timeout: 30
write-timeout: 30
idle-timeout: 60
max-header-bytes: 1048576
enable-http: true
enable-grpc: false
enable-tls: false
enable-gzip-compress: true
tls:
  cert-file: ""
  key-file: ""
  ca-file: ""
headers:
  x-custom-header: "my-app"
  x-version: "1.0.0"
```

现在修改配置文件，程序会自动检测变化并重载配置！

## 🎯 支持的配置模块

| 类别 | 模块 | 描述 |
|------|------|------|
| **网关服务** | Gateway, HTTP, GRPC | 网关和服务配置 |
| **数据存储** | MySQL, PostgreSQL, SQLite, Redis | 数据库配置 |
| **中间件** | CORS, 限流, JWT, 恢复 | 常用中间件配置 |
| **监控运维** | Health, Metrics, Prometheus, Jaeger | 监控和链路追踪 |
| **消息队列** | Kafka, MQTT | 消息系统配置 |
| **第三方服务** | 支付宝, 微信支付, 阿里云短信 | 第三方集成 |

## 📜 许可证

本项目采用 [MIT 许可证](LICENSE)

---

**如果这个项目对你有帮助，请给我们一个 ⭐️**
