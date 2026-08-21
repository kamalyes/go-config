# go-config

[![Go Reference](https://pkg.go.dev/badge/github.com/kamalyes/go-config.svg)](https://pkg.go.dev/github.com/kamalyes/go-config)
[![Go Report Card](https://goreportcard.com/badge/github.com/kamalyes/go-config)](https://goreportcard.com/report/github.com/kamalyes/go-config)
[![Tests](https://github.com/kamalyes/go-config/actions/workflows/test.yml/badge.svg)](https://github.com/kamalyes/go-config/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

一个功能强大且易于使用的 Go 配置管理库，支持多种配置格式、智能发现、热更新和环境隔离，为企业级应用提供开箱即用的配置管理解决方案。

- **Go 版本**：`go 1.25.0`
- **模块路径**：`github.com/kamalyes/go-config`

## ✨ 核心特性

### 🎯 配置管理

- **多格式支持** - 基于 Viper，支持 YAML、JSON、TOML、Properties 等多种配置格式
- **智能发现** - 自动发现和加载配置文件，支持多环境配置（按文件名末尾段识别环境，如 `config-dev.yaml`、`config-prod.yaml`）
- **灵活反序列化** - `UnmarshalWithFlexibleNaming` / `UnmarshalWithKebabToSnake` 支持宽松的键名匹配与 kebab→snake 转换
- **配置校验** - `ValidateConfig` 校验配置可用性
- **泛型安全访问** - `GetConfigAs[T]` / `MustGetConfigAs[T]` 以编译期类型安全的方式获取配置
- **配置初始化生成** - `SmartConfigGenerator` 按模块生成带字段注释的 YAML/JSON 配置模板，支持备份已存在文件
- **配置格式化** - `ConfigFormatter` 输出配置变更日志

### 🔥 热更新机制

- **📁 文件监控** - 基于 fsnotify 实时监听配置文件变化
- **⏱️ 防抖处理** - 可配置的防抖延迟，避免频繁重载
- **🔔 回调系统** - `CommonCallbackManager` 灵活的回调管理器，支持优先级、异步执行、超时控制、按类型过滤
- **🔄 错误恢复** - 配置重载失败时自动重试，可配置重试次数
- **🌐 环境监控** - 支持监控环境变量变化并自动重载

### 🌍 环境管理

- **63 种环境** - 内置 9 种标准环境 + 54 个国家/地区环境
- **别名支持** - 每个环境支持多个别名（如 prod/production/prd）
- **自动初始化** - 包导入时自动初始化，无需手动调用
- **环境级别** - 按重要程度对环境分级（开发/测试/生产），提供 `IsDevelopmentLevel` / `IsTestingLevel` / `IsProductionLevel`
- **动态切换** - 运行时动态切换环境并触发回调
- **自定义注册** - `RegisterEnvPrefixes` 支持注册自定义环境类型与别名

### 📦 丰富的配置模块（49 个）

| 分类 | 配置模块 | 说明 |
| ------ | --------- | ------ |
| **🌐 网关与服务** | Gateway | 网关统一配置（含 HTTP/gRPC 服务） |
| | RESTful | RESTful API 配置 |
| | RPC Client / RPC Server | RPC 客户端和服务端配置 |
| **💾 数据存储** | Database | 数据库统一配置（MySQL、PostgreSQL、SQLite） |
| | Redis | Redis 缓存配置 |
| | Cache | 多级缓存配置（Memory、Expiring、Ristretto、Sharded、TwoLevel） |
| | Elasticsearch | Elasticsearch 配置 |
| | Etcd | Etcd 配置 |
| | ClickHouse | ClickHouse 列式数据库配置 |
| | TSDB | 时序数据库配置 |
| **🔌 中间件** | CORS | 跨域资源共享配置 |
| | JWT | JWT 认证配置 |
| | RateLimit | 限流配置 |
| | Recovery | 恢复中间件配置 |
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
| | Queue | 消息队列统一配置 |
| **🗄️ 对象存储** | OSS | 对象存储统一配置（阿里云、MinIO、S3、BoltDB） |
| **🔗 第三方服务** | Pay | 支付宝、微信支付配置 |
| | SMS | 阿里云短信配置 |
| | Email / SMTP | 邮件配置 |
| | Youzan | 有赞平台配置 |
| | STS | 阿里云 STS 配置 |
| **⚙️ 其他功能** | Logging / Zap | 日志配置 |
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
| | Common | 公共基础配置 |

### 🎨 开发体验

- **链式 API** - 泛型构建器 `NewConfigBuilder[T]` 优雅的链式调用
- **零配置启动** - 合理的默认值，开箱即用
- **类型安全** - 泛型支持，编译时类型检查
- **上下文集成** - 配置信息自动注入到 `context.Context`，通过 `ContextKeyHelper` 从上下文获取
- **统一错误处理** - 分类错误（`errors.go`）支持错误分类、严重程度、回调

## 🚀 快速开始

### 安装

```bash
go get github.com/kamalyes/go-config
```

### 最小示例

```go
package main

import (
 "fmt"
 "time"

 goconfig "github.com/kamalyes/go-config"
)

type AppConfig struct {
 Server struct {
  Host string `mapstructure:"host"`
  Port int    `mapstructure:"port"`
 } `mapstructure:"server"`
}

func main() {
 cfg := &AppConfig{}
 manager := goconfig.NewConfigBuilder(cfg).
  WithConfigPath("config.yaml").              // 指定配置文件路径
  WithEnvironment(goconfig.EnvDevelopment).    // 设置运行环境
  WithHotReload(&goconfig.HotReloadConfig{     // 启用热更新
   Enabled:       true,
   DebounceDelay: 1 * time.Second,
  }).
  MustBuildAndStart() // 构建并启动（失败时 panic）
 defer manager.Stop()

 fmt.Printf("服务启动在 %s:%d\n", cfg.Server.Host, cfg.Server.Port)
}
```

对应的 `config.yaml`：

```yaml
server:
  host: localhost
  port: 8080
```

### 📖 示例代码

所有功能的完整示例代码请查看 [examples](examples/) 目录

| 示例 | 说明 | 链接 |
| ------ | ------ | ------ |
| 🎯 基础使用 | 最基本的配置加载和使用方式 | [examples/basic](examples/basic/main.go) |
| 🔥 配置热更新 | 启用配置热更新，注册配置变更回调 | [examples/hot_reload](examples/hot_reload/main.go) |
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

## 📋 配置构建器 API

`ConfigBuilder[T]` 提供以下链式方法（详见 [config_builder.go](config_builder.go)）：

| 方法 | 说明 |
| ------ | ------ |
| `WithConfigPath(path)` | 直接指定配置文件路径 |
| `WithSearchPath(path)` | 设置配置搜索目录 |
| `WithPrefix(prefix)` | 设置配置文件名前缀（如 `gateway-xl`） |
| `WithPattern(pattern)` | 按模式匹配配置文件 |
| `WithEnvironment(env)` | 设置运行环境 |
| `WithHotReload(cfg)` | 配置热更新参数 |
| `WithContext(opts)` | 设置上下文键选项 |
| `Build()` / `BuildAndStart(ctx)` | 构建 / 构建并启动 |
| `MustBuildAndStart(ctx)` | 构建并启动，失败时 panic |

完整示例请查看 [examples/builder](examples/builder/main.go)

### 配置发现

`ConfigDiscovery` 支持目录扫描、按环境/模式发现配置文件，环境识别按文件名末尾段匹配。完整示例请查看 [examples/discovery](examples/discovery/main.go)

## 🎯 高级特性

### 上下文集成

`ContextManager` / `ContextKeyHelper` 将配置与环境信息注入 `context.Context`，便于跨层传递。完整示例请查看 [examples/context](examples/context/main.go)

### 集成配置管理器

`IntegratedConfigManager` 是统一入口，聚合配置、热重载、环境与上下文管理，主要方法：

| 方法 | 说明 |
| ------ | ------ |
| `Start(ctx)` / `Stop()` / `MustStart(ctx)` | 生命周期管理 |
| `IsRunning()` | 运行状态 |
| `GetConfig()` / `GetConfigAs[T]` / `MustGetConfigAs[T]` | 获取配置（泛型安全访问） |
| `GetViper()` / `GetContextManager()` / `GetEnvironmentManager()` / `GetHotReloader()` | 获取底层组件 |
| `GetConfigMetadata()` | 获取配置元数据 |
| `RegisterConfigCallback` / `UnregisterConfigCallback` | 注册/注销配置变更回调 |
| `RegisterEnvironmentCallback` / `UnregisterEnvironmentCallback` | 注册/注销环境变更回调 |
| `SetEnvironment(env)` | 动态切换环境 |
| `ValidateConfig()` | 校验配置 |
| `WithContext(ctx)` | 将配置注入 context |

## 📚 详细文档

- [示例代码](examples/) - 各种使用场景的完整示例
- [配置模块文档](pkg/) - 各配置模块的详细文档与示例配置文件
- [环境管理器](env.go) - 环境类型、别名、级别判断的源码定义

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
        // pprof.EnableProfiling()
        // gin.SetMode(gin.DebugMode)
    }

    if goconfig.IsProductionLevel() {
        // 生产环境：启用监控
        // prometheus.EnableMetrics()
        // sentry.InitErrorTracking()
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
