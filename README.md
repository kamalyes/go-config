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

### 📦 配置模块（35 个）

| 分类 | 配置模块 | 说明 |
| ------ | --------- | ------ |
| **🌐 网关与服务** | Gateway | 网关统一配置（含 HTTP/gRPC 服务） |
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
| | SMTP | 邮件发送配置 |
| **⚙️ 其他功能** | Logging | 日志配置 |
| | I18n | 国际化配置 |
| | Security | 安全配置 |
| | Signature | 签名配置 |
| | Banner | 启动横幅配置 |
| | Swagger | API 文档配置 |
| | Jobs | 定时任务配置 |
| | WSC | WebSocket 通信配置 |
| | Breaker | 熔断器配置 |
| | Common | 公共基础配置 |
### 🎨 开发体验

- **链式 API** - 泛型构建器 `NewConfigBuilder[T]` 优雅的链式调用
- **零配置启动** - 合理的默认值，开箱即用
- **类型安全** - 泛型支持，编译时类型检查
- **上下文集成** - 配置信息自动注入到 `context.Context`，通过 `ContextKeyHelper` 从上下文获取
- **统一错误处理** - 分类错误（`errors.go`）支持错误分类、严重程度、回调

## ⚠️ 用之前先知道这几条

这几条都是实际踩过或实测过的，不是风格建议。

### 配置键一律用连字符

`mapstructure` 的键名**必须**写成连字符形式（`module-name`、`l1-type`），不要用下划线。
仓库里 1200 多个键有 764 处连字符，这是约定，而初次加载的反序列化**不做** kebab↔snake 转换。

违反的代价是静默的。本库曾有 6 个字段把 `mapstructure` 写成 `l1_type` 而 `yaml` tag 写
`l1-type`，于是仓库自带的示例配置里那几行**从来没生效过**：

```yaml
# 配置里写 l1-type，而 tag 是 l1_type
l1-type: memory      # -> 字段保持零值，没有任何报错
l1_type: memory      # -> 才是当时唯一生效的写法
```

新增字段时让 `mapstructure` / `yaml` 两个 tag 完全一致，`json` 用驼峰。

### 初次加载与热重载的键名匹配规则不同

这是上一条之所以危险的放大器：

| 路径 | 用的反序列化 | 是否做 kebab→snake |
|------|-------------|-------------------|
| 初次加载 | 裸 `viper.Unmarshal` | **不做** |
| 热重载 | `UnmarshalWithFlexibleNaming` | **做** |

所以一个键名风格不匹配的字段，会在**启动时是空的、热重载之后突然有值**。实测：

```text
配置里写 l1-type / l2-type：
  初次加载   tag=l1_type 收到 ""        tag=l2-type 收到 "redis"
  热重载     tag=l1_type 收到 "memory"  tag=l2-type 收到 "redis"
```

只要遵守上一条（两个 tag 一致、一律连字符），两条路径的行为就一致。

### 环境变量注入：目前不支持

配置值**不能**从环境变量注入。host、password 这类都得写在配置文件里。实测三种常见预期：

```text
① YAML 里写 ${DB_HOST} 占位符   -> host="${DB_HOST}"   原样字面量，不展开
② 只设环境变量 HOST / PASSWORD  -> host="file-host"    环境变量完全无效
③ 手动开 viper.AutomaticEnv()   -> host="env-host"     生效
                                   password=""         但仍取不到
```

第 ③ 条是关键：`AutomaticEnv()` 只解析 viper **已经知道**的键。配置文件里没出现过的
`password`，即使设了环境变量也取不到——而那恰好是密码最需要的场景（不写进文件）。
所以"打开 AutomaticEnv 就行"是不够的，还要配合 `SetDefault` 或 `BindEnv` 逐键注册。

环境变量目前只用来**选环境**（`GO_ENV=production` 之类），不用来注入配置值。

### 多实例配置：只有两处支持

想配 `master` / `slave1` / `slave2`，或多个 Kafka 集群，目前**配不出来**。

| 配置 | 多实例 | 形式 |
|------|--------|------|
| `gateway.GRPC.Clients` | ✅ | `map[string]*GRPCClient`，key 为服务名 |
| `wsc.Tokens` | ✅ | `map[string]*ConnectionTokenSet`，key 为 appid |
| 数据库 / Kafka / NATS / Redis / ClickHouse / Elasticsearch / Etcd / MQTT / SMTP | ❌ | 单实例 |

数据库的容器按**驱动类型**分键（`mysql` / `postgresql` / `sqlite` / `cockroachdb`），
每个驱动一个实例，不是按实例名分键。

支持多实例的那两处用的是「map key 回填」模式：实例名不写在字段里，而是加载后从 map key
回填到一个 `mapstructure:"-"` 的字段上。新增多实例配置时照这个模式写。

### 衍生字段用 AfterLoad / Validate 计算

不从配置文件读、而是算出来的字段（如拼好的 `Endpoint`），tag 写 `mapstructure:"-" yaml:"-" json:"-"`，
在 `AfterLoad()`（`internal.CallAfterLoad` 调用）或 `Validate()` 里计算。

注意是 `"-"` 不是 `"_"`：写成 `"_"` 的话这个字段仍然可以被配置里一个叫 `_` 的键覆盖掉。

---

## ⚡ 请求属性提取与性能

`common.ExtractAttribute(r, sources)` 按来源清单顺序取第一个非空值，来源可以是 header、
query 或 cookie。

### 取多个属性时用 AttributeExtractor

标准库的 `r.URL.Query()` 与 `r.Cookie(name)` **每次调用都重新解析**。一个属性一次
`ExtractAttribute` 的用法下，同一个查询串会被解析几十遍——网关默认配置有 43 个 query 来源、
25 个 cookie 来源，分布在 28 个属性的清单里。

所以取多个属性时，建一个请求级的提取器共用：

```go
// 推荐：一次请求建一个，查询串与 Cookie 头各只解析一次
e := common.NewAttributeExtractor(r)
traceID := e.Extract(cfg.TraceIDSources)
userID  := e.Extract(cfg.UserIDSources)
tenantID := e.Extract(cfg.TenantIDSources)

// 只取一个属性时，这样就够了
traceID := common.ExtractAttribute(r, cfg.TraceIDSources)
```

两者语义完全一致。提取器是请求级的，不加锁，不要跨 goroutine 共享。

实测（按真实默认配置取齐 28 个属性，Apple M4，基准见 `pkg/gateway/request_bench_test.go`）：

| 请求形态 | 一个属性一次调用 | 复用提取器 | 提速 |
|---------|----------------|-----------|------|
| 属性都不存在 | 644 ns / 0 allocs | 649 ns / 0 allocs | — |
| 从 header 命中 | 885 ns / 0 allocs | 778 ns / 0 allocs | 1.14x |
| **带查询串** | 9617 ns / 200 allocs | **1210 ns / 8 allocs** | **7.95x** |
| 带 Cookie 头 | 3024 ns / 28 allocs | 1278 ns / 7 allocs | 2.37x |
| 混合 | 6160 ns / 116 allocs | 1594 ns / 9 allocs | 3.86x |

`query` 形态剩下的 8 次分配全部来自那唯一一次 `url.ParseQuery`——构建 `url.Values` 必须付的，
是当前语义下的地板。

### header 来源的 Key 写规范形式

`AttributeSource` 里 header 来源的 Key 请写 HTTP 规范形式：`X-User-Id`，不要写 `X-User-ID`。

两种拼法的查找结果完全一样，但 `Header.Get` 每次都会对 key 调 `CanonicalMIMEHeaderKey`——
key 本来就规范时原样返回、不分配；不规范则每次查找都要现算并分配一个新字符串。

默认配置里的 Key 已经是规范形式。用户自己在 YAML 里写 `x-user-id` 也没关系：
`RequestContext.Validate()` 会在加载后归一化一次，之后的查找同样零分配
（实测 111 ns / 3 allocs → 49 ns / 0 allocs）。

query 与 cookie 的键名**大小写敏感**，不会被归一化。

---

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
