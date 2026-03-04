# go-config

[![Go Reference](https://pkg.go.dev/badge/github.com/kamalyes/go-config.svg)](https://pkg.go.dev/github.com/kamalyes/go-config)
[![Go Report Card](https://goreportcard.com/badge/github.com/kamalyes/go-config)](https://goreportcard.com/report/github.com/kamalyes/go-config)
[![Tests](https://github.com/kamalyes/go-config/actions/workflows/test.yml/badge.svg)](https://github.com/kamalyes/go-config/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

一个功能强大且易于使用的 Go 配置管理库，支持多种配置格式、智能发现、热更新和安全访问为企业级应用提供开箱即用的配置管理解决方案

## ✨ 核心特性

### 🎯 配置管理
- **多格式支持** - 支持 YAML、JSON、TOML、Properties 等多种配置格式
- **智能发现** - 自动发现和加载配置文件，支持多环境配置
- **配置验证** - 内置 validator 验证机制，确保配置正确性
- **安全访问** - 防止空指针异常的链式配置访问
- **配置导出** - 支持将配置导出为 YAML/JSON 格式

### 🔥 热更新机制
- **📁 文件监控** - 基于 fsnotify 实时监听配置文件变化
- **⏱️ 防抖处理** - 可配置的防抖延迟，避免频繁重载
- **🔔 回调系统** - 灵活的回调管理器，支持优先级、异步执行、超时控制
- **🔄 错误恢复** - 配置重载失败时自动重试，可配置重试次数
- **🌐 环境监控** - 支持监控环境变量变化并自动重载

### 🌍 环境管理
- **56 种环境** - 内置 9 种标准环境 + 47 个国家/地区环境
- **别名支持** - 每个环境支持多个别名（如 prod/production/prd）
- **自动初始化** - 包导入时自动初始化，无需手动调用
- **环境级别** - 按重要程度对环境分级（开发/测试/生产）
- **动态切换** - 运行时动态切换环境并触发回调
- **自定义注册** - 支持注册自定义环境类型

### 📦 丰富的配置模块（48+）

| 分类 | 配置模块 | 说明 |
|------|---------|------|
| **🌐 网关与服务** | Gateway | 网关统一配置（集成所有子模块） |
| | HTTP/GRPC | HTTP 和 gRPC 服务器配置 |
| | RPC Client/Server | RPC 客户端和服务端配置 |
| | RESTful | RESTful API 配置 |
| **💾 数据存储** | Database | 数据库统一配置（MySQL、PostgreSQL、SQLite） |
| | Redis | Redis 缓存配置 |
| | Cache | 多级缓存配置（Memory、Expiring、Ristretto、Sharded、TwoLevel） |
| | Elasticsearch | Elasticsearch 配置 |
| | Etcd | Etcd 配置 |
| **🔌 中间件** | CORS | 跨域资源共享配置 |
| | JWT | JWT 认证配置 |
| | RateLimit | 限流配置 |
| | Recovery | 恢复中间件配置 |
| | RequestID | 请求 ID 配置 |
| | Timeout | 超时配置 |
| | Middleware | 中间件统一配置 |
| **📊 监控运维** | Health | 健康检查配置 |
| | Metrics | 指标收集配置 |
| | Prometheus | Prometheus 监控配置 |
| | Jaeger | Jaeger 链路追踪配置 |
| | Tracing | 分布式追踪配置 |
| | Monitoring | 监控统一配置 |
| | Grafana | Grafana 配置 |
| | Pprof | 性能分析配置 |
| | Alerting | 告警配置 |
| **📨 消息队列** | Kafka | Kafka 配置 |
| | MQTT | MQTT 配置 |
| | Queue | 消息队列统一配置 |
| **🗄️ 对象存储** | OSS | 对象存储统一配置（阿里云、MinIO、S3、BoltDB） |
| **🔗 第三方服务** | 支付 | 支付宝、微信支付配置 |
| | 短信 | 阿里云短信配置 |
| | 邮件 | SMTP、Email 配置 |
| | 有赞 | 有赞平台配置 |
| | STS | 阿里云 STS 配置 |
| **⚙️ 其他功能** | Logging/Zap | 日志配置 |
| | I18n | 国际化配置 |
| | Security | 安全配置 |
| | Signature | 签名配置 |
| | Captcha | 验证码配置 |
| | Banner | 启动横幅配置 |
| | Swagger | API 文档配置 |
| | Jobs | 定时任务配置 |
| | WSC | WebSocket 通信配置 |
| | Breaker | 熔断器配置 |
| | Consul | Consul 配置 |
| | FTP | FTP 配置 |

### 🎨 开发体验
- **链式 API** - 优雅的构建器模式 API 设计
- **零配置启动** - 开箱即用的默认配置
- **类型安全** - 泛型支持，编译时类型检查
- **上下文集成** - 配置信息自动注入到 context.Context
- **错误处理** - 统一的错误处理机制，支持错误分类、严重程度、回调

## 🚀 快速开始

### 安装

```bash
go get github.com/kamalyes/go-config
```

### 📖 示例代码

所有功能的完整示例代码请查看 [examples](examples/) 目录

| 示例 | 说明 | 链接 |
|------|------|------|
| 🎯 基础使用 | 最基本的配置加载和使用方式 | [examples/basic](examples/basic/main.go) |
| 🔥 配置热更新 | 启用配置热更新，注册配置变更回调 | [examples/hot_reload](examples/hot_reload/main.go) |
| ⚠️ 错误处理 | 创建自定义错误处理器，根据错误严重程度采取不同措施 | [examples/error_handling](examples/error_handling/main.go) |
| 🌍 环境管理 | 使用环境判断函数，在不同环境下加载不同配置 | [examples/environment](examples/environment/main.go) |
| 📋 上下文集成 | 将配置信息注入到 context.Context，从上下文中获取配置 | [examples/context](examples/context/main.go) |
| 🔍 配置发现 | 配置文件的自动发现、扫描和创建功能 | [examples/discovery](examples/discovery/main.go) |
| 🏗️ 构建器模式 | 配置构建器的各种链式 API 用法 | [examples/builder](examples/builder/main.go) |

### 快速运行示例

```bash
# 进入示例目录
cd examples

# 运行基础示例
cd basic && go run main.go

# 运行热更新示例
cd hot_reload && go run main.go

# 查看所有示例的详细说明
cat README.md
```

### 🌍 环境管理

完整示例请查看 [examples/environment](examples/environment/main.go)

## 🔥 配置热更新

完整示例请查看 [examples/hot_reload](examples/hot_reload/main.go)

### 特性说明

- **📁 文件监控** - 基于 fsnotify 实时监听配置文件变化
- **⏱️ 防抖处理** - 可配置的防抖延迟，避免频繁重载
- **🔔 回调系统** - 灵活的回调管理器，支持优先级、异步执行、超时控制
- **🔄 错误恢复** - 配置重载失败时自动重试，可配置重试次数
- **🌐 环境监控** - 支持监控环境变量变化并自动重载

## 📋 配置构建器 API

完整示例请查看 [examples/builder](examples/builder/main.go)

### 配置发现

完整示例请查看 [examples/discovery](examples/discovery/main.go)

## 🎯 高级特性

### 上下文集成

完整示例请查看 [examples/context](examples/context/main.go)

### 错误处理

完整示例请查看 [examples/error_handling](examples/error_handling/main.go)

## 📚 详细文档

- [示例代码](examples/) - 各种使用场景的完整示例
- [环境管理器使用文档](ENV_USAGE.md) - 环境管理的详细说明
- [配置模块文档](pkg/) - 各配置模块的详细文档

## 🎨 最佳实践

### 1. 使用配置前缀进行环境隔离

```go
// 开发环境：gateway-xl-dev.yaml
// 测试环境：gateway-xl-test.yaml
// 生产环境：gateway-xl-prod.yaml
manager := goconfig.NewConfigBuilder(config).
    WithPrefix("gateway-xl").
    WithSearchPath("resources").
    MustBuildAndStart()
```

### 2. 利用热更新实现零停机配置变更

```go
manager.RegisterConfigCallback(
    func(ctx context.Context, event goconfig.CallbackEvent) error {
        // 重新初始化依赖配置的组件
        return reinitializeComponents(event.NewValue)
    },
    goconfig.CallbackOptions{
        ID:    "component-reinit",
        Types: []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
    },
)
```

### 3. 使用环境级别进行功能开关

```go
func setupFeatures() {
    if goconfig.IsDevelopmentLevel() {
        // 开发环境：启用调试功能
        pprof.EnableProfiling()
        gin.SetMode(gin.DebugMode)
    }
    
    if goconfig.IsProductionLevel() {
        // 生产环境：启用监控
        prometheus.EnableMetrics()
        sentry.InitErrorTracking()
    }
}
```

### 4. 全球化部署配置

```bash
# 中国区域
export APP_ENV=china
./app

# 美国区域
export APP_ENV=usa
./app

# 欧洲区域
export APP_ENV=germany
./app
```

## 🔧 配置示例

查看 [pkg/](pkg/) 目录下各模块的 `.yaml` 和 `.json` 示例文件

## 📜 许可证

本项目采用 [MIT 许可证](LICENSE)

---

**如果这个项目对你有帮助，请给我们一个 ⭐️**
