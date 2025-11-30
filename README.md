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
- 🎯 **智能发现** - 自动发现和加载配置文件（支持多环境）
- 🌍 **多环境支持** - 内置9种环境类型，支持自定义环境注册
- 📦 **丰富模块** - 内置40+配置模块，覆盖常见应用场景
- 🚀 **零配置启动** - 开箱即用的默认配置
- 🎨 **链式API** - 优雅的构建器模式API设计

## 🌍 环境与配置文件发现

### 内置环境类型

| 环境类型 | 常量 | 支持的配置文件后缀 |
|---------|------|-------------------|
| 开发环境 | `EnvDevelopment` | `dev`, `develop`, `development` |
| 本地环境 | `EnvLocal` | `local`, `localhost` |
| 测试环境 | `EnvTest` | `test`, `testing`, `qa`, `sit` |
| 预发布环境 | `EnvStaging` | `staging`, `stage`, `stg`, `pre`, `preprod`, `pre-prod`, `fat`, `gray`, `grey`, `canary` |
| 生产环境 | `EnvProduction` | `prod`, `production`, `prd`, `release`, `live`, `online`, `master`, `main` |
| 调试环境 | `EnvDebug` | `debug`, `debugging`, `dbg` |
| 演示环境 | `EnvDemo` | `demo`, `demonstration`, `showcase`, `preview`, `sandbox` |
| UAT环境 | `EnvUAT` | `uat`, `acceptance`, `user-acceptance`, `beta` |
| 集成环境 | `EnvIntegration` | `integration`, `int`, `ci`, `integration-test`, `integ` |

### 配置文件命名规则

配置文件命名格式：`{prefix}-{env-suffix}.{ext}`

例如，当 `APP_ENV=local` 且前缀为 `gateway-xl` 时，会按优先级查找：
- `gateway-xl-local.yaml`
- `gateway-xl-local.yml`
- `gateway-xl-localhost.yaml`
- ...

### 注册自定义环境

如果内置环境不满足需求，可以注册自定义环境：

```go
package main

import goconfig "github.com/kamalyes/go-config"

func init() {
    // 注册自定义环境 "custom"，支持后缀 "custom", "my-env", "myenv"
    // 配置文件可命名为: gateway-xl-custom.yaml, gateway-xl-my-env.yaml 等
    goconfig.RegisterEnvPrefixes("custom", "custom", "my-env", "myenv")
}
```

### 配置文件未找到时的错误提示

当配置文件未找到时，会输出详细的诊断信息：

```
❌ 未找到前缀为 'gateway-xl' 的配置文件
📍 搜索路径: resources
🌍 当前环境: custom-env
⚠️ 当前环境 'custom-env' 未在 DefaultEnvPrefixes 中注册
📋 已注册的环境及其后缀:
   - development: [dev develop development]
   - local: [local localhost]
   ...

💡 如需注册自定义环境，请在程序启动前注册:

   示例代码:
   func init() {
       goconfig.RegisterEnvPrefixes("custom-env", "custom-env", "custom-alias")
   }
```

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
