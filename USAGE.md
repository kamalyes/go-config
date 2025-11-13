# 📖 go-config 详细使用文档

## 目录

- [🚀 快速开始](#-快速开始)
- [🎯 核心概念](#-核心概念)
- [🔗 链式调用 API](#-链式调用-api)
- [🌍 环境管理](#-环境管理)
- [🔄 热更新配置](#-热更新配置)
- [🔔 回调系统](#-回调系统)
- [🛡️ 类型安全](#️-类型安全)
- [📁 配置发现](#-配置发现)
- [💼 实际应用场景](#-实际应用场景)
- [🔧 高级用法](#-高级用法)
- [❓ 故障排除](#-故障排除)

---

## 🚀 快速开始

### 📦 安装

```bash
go get -u github.com/kamalyes/go-config
```

### ⚡ 5分钟快速上手

#### 1. 创建配置结构体

```go
package main

import (
    "context"
    "log"
    "time"
    
    goconfig "github.com/kamalyes/go-config"
    "github.com/kamalyes/go-config/pkg/gateway"
)

// 定义你的配置结构体
type AppConfig struct {
    Name     string `yaml:"name" json:"name"`
    Version  string `yaml:"version" json:"version"`
    Debug    bool   `yaml:"debug" json:"debug"`
    
    Server struct {
        Host         string        `yaml:"host" json:"host"`
        Port         int           `yaml:"port" json:"port"`
        ReadTimeout  time.Duration `yaml:"read-timeout" json:"read-timeout"`
        WriteTimeout time.Duration `yaml:"write-timeout" json:"write-timeout"`
    } `yaml:"server" json:"server"`
    
    Database struct {
        Host     string `yaml:"host" json:"host"`
        Port     int    `yaml:"port" json:"port"`
        Username string `yaml:"username" json:"username"`
        Password string `yaml:"password" json:"password"`
        Database string `yaml:"database" json:"database"`
    } `yaml:"database" json:"database"`
}
```

#### 2. 创建配置文件

创建 `config/app-dev.yaml`:

```yaml
name: "my-awesome-app"
version: "1.0.0"
debug: true

server:
  host: "0.0.0.0"
  port: 8080
  read-timeout: 30s
  write-timeout: 30s

database:
  host: "localhost"
  port: 3306
  username: "root"
  password: "password"
  database: "myapp"
```

#### 3. 使用链式 API 创建管理器

```go
func main() {
    // 🎯 推荐方式：链式调用API
    config := &AppConfig{}
    
    manager, err := goconfig.NewManager(config).
        WithSearchPath("./config").        // 设置搜索路径
        WithPrefix("app").                 // 配置文件前缀
        WithEnvironment(goconfig.GetEnvironment()). // 自动获取环境
        WithHotReload(nil).                // 启用热更新（使用默认配置）
        BuildAndStart()                    // 构建并启动
    
    if err != nil {
        log.Fatal("配置管理器启动失败:", err)
    }
    defer manager.Stop()
    
    // ✅ 类型安全的配置获取
    appConfig, err := goconfig.GetConfigAs[AppConfig](manager)
    if err != nil {
        log.Fatal("获取配置失败:", err)
    }
    
    // 🎉 使用配置
    log.Printf("🚀 应用启动: %s v%s", appConfig.Name, appConfig.Version)
    log.Printf("🌐 服务地址: %s:%d", appConfig.Server.Host, appConfig.Server.Port)
    log.Printf("🗄️ 数据库: %s@%s:%d/%s", 
        appConfig.Database.Username,
        appConfig.Database.Host,
        appConfig.Database.Port,
        appConfig.Database.Database)
    
    // 🔥 应用保持运行...
    select {}
}
```

---

## 🎯 核心概念

### 🏗️ 架构组件

| 组件 | 作用 | 特点 |
|------|------|------|
| **ManagerBuilder** | 配置管理器构建器 | 支持链式调用，类型安全 |
| **IntegratedConfigManager** | 集成配置管理器 | 统一管理配置、环境、热更新 |
| **HotReloader** | 热更新器 | 监听文件变化，自动重载 |
| **Environment** | 环境管理器 | 多环境切换和管理 |
| **ContextManager** | 上下文管理器 | 配置注入到 Context |

### 🎭 配置发现模式

1. **直接路径模式** - 明确指定配置文件路径
2. **前缀匹配模式** - 根据前缀和环境自动匹配
3. **模式匹配模式** - 使用glob模式灵活匹配
4. **自动发现模式** - 智能扫描目录中的配置文件

---

## 🔗 链式调用 API

### 🏗️ ManagerBuilder 构建器

链式调用API是推荐的使用方式，提供流畅的配置体验：

```go
// 🎯 完整的链式调用示例
manager, err := goconfig.NewManager(config).
    WithSearchPath("./config").           // 1. 设置搜索路径
    WithPrefix("myapp").                 // 2. 设置配置前缀
    WithEnvironment(goconfig.EnvProduction). // 3. 设置环境
    WithHotReload(&goconfig.HotReloadConfig{
        Enabled:  true,
        Debounce: time.Second * 3,       // 4. 配置热更新
    }).
    WithContext(&goconfig.ContextKeyOptions{
        Value: "production-service",      // 5. 设置上下文
    }).
    BuildAndStart()                      // 6. 构建并启动
```

### 📋 构建器方法详解

#### WithConfigPath - 直接路径

```go
// 直接指定配置文件路径
manager, err := goconfig.NewManager(config).
    WithConfigPath("./configs/production.yaml").
    BuildAndStart()
```

#### WithSearchPath - 搜索路径

```go
// 在指定目录中自动发现配置文件
manager, err := goconfig.NewManager(config).
    WithSearchPath("./config").    // 搜索 ./config 目录
    BuildAndStart()
```

#### WithPrefix - 前缀匹配

```go
// 使用前缀匹配配置文件
// 会查找: gateway-dev.yaml, gateway-prod.json 等
manager, err := goconfig.NewManager(config).
    WithSearchPath("./config").
    WithPrefix("gateway").
    WithEnvironment(goconfig.EnvDevelopment).
    BuildAndStart()
```

#### WithPattern - 模式匹配

```go
// 使用glob模式匹配
manager, err := goconfig.NewManager(config).
    WithSearchPath("./config").
    WithPattern("service-*.yaml").    // 匹配 service-*.yaml
    BuildAndStart()
```

#### WithEnvironment - 环境设置

```go
// 明确设置环境
manager, err := goconfig.NewManager(config).
    WithSearchPath("./config").
    WithEnvironment(goconfig.EnvProduction).  // 设置为生产环境
    BuildAndStart()
```

#### WithHotReload - 热更新配置

```go
// 使用默认热更新配置
manager, err := goconfig.NewManager(config).
    WithConfigPath("./app.yaml").
    WithHotReload(nil).              // 使用默认配置
    BuildAndStart()

// 自定义热更新配置
manager, err := goconfig.NewManager(config).
    WithConfigPath("./app.yaml").
    WithHotReload(&goconfig.HotReloadConfig{
        Enabled:    true,
        Debounce:   time.Second * 2,  // 防抖延迟
        MaxRetries: 5,                // 最大重试次数
    }).
    BuildAndStart()
```

### 🛠️ 构建方法

#### Build - 仅构建

```go
// 构建管理器但不启动
manager, err := goconfig.NewManager(config).
    WithConfigPath("./app.yaml").
    Build()

if err != nil {
    log.Fatal(err)
}

// 手动启动
ctx := context.Background()
if err := manager.Start(ctx); err != nil {
    log.Fatal(err)
}
```

#### BuildAndStart - 构建并启动

```go
// 一步完成构建和启动
manager, err := goconfig.NewManager(config).
    WithConfigPath("./app.yaml").
    BuildAndStart()
```

#### MustBuildAndStart - 必须成功

```go
// 失败时panic，适用于应用启动阶段
manager := goconfig.NewManager(config).
    WithConfigPath("./app.yaml").
    MustBuildAndStart()  // 失败时panic
```

---

## 🌍 环境管理

### 📋 支持的环境

| 环境 | 标识符 | 说明 | 配置文件模式 |
|------|--------|------|-------------|
| 开发环境 | `dev` | 开发调试 | `app-dev.yaml` |
| 测试环境 | `test` | 单元测试 | `app-test.yaml` |
| 集成环境 | `sit` | 系统集成测试 | `app-sit.yaml` |
| 验收环境 | `uat` | 用户验收测试 | `app-uat.yaml` |
| 预生产 | `staging` | 预生产验证 | `app-staging.yaml` |
| 生产环境 | `prod` | 正式生产 | `app-prod.yaml` |

### 🔧 环境设置方式

#### 1. 环境变量设置

```bash
# Linux/macOS
export APP_ENV=prod
export GO_ENV=production

# Windows
set APP_ENV=prod
set GO_ENV=production
```

#### 2. 代码中设置

```go
// 方式1: 直接设置环境
manager, err := goconfig.NewManager(config).
    WithEnvironment(goconfig.EnvProduction).
    BuildAndStart()

// 方式2: 运行时切换环境
err = manager.SetEnvironment(goconfig.EnvProduction)
```

#### 3. 自动环境检测

```go
// 自动从环境变量读取
currentEnv := goconfig.GetEnvironment()
fmt.Printf("当前环境: %s", currentEnv)

manager, err := goconfig.NewManager(config).
    WithEnvironment(currentEnv).  // 使用检测到的环境
    BuildAndStart()
```

### 🔄 环境切换回调

```go
// 注册环境变更监听
manager.RegisterEnvironmentCallback("env_logger", 
    func(oldEnv, newEnv goconfig.EnvironmentType) error {
        log.Printf("🌍 环境切换: %s → %s", oldEnv, newEnv)
        
        // 根据环境执行不同逻辑
        switch newEnv {
        case goconfig.EnvProduction:
            // 生产环境特殊处理
            enableProductionMode()
        case goconfig.EnvDevelopment:
            // 开发环境特殊处理
            enableDebugMode()
        }
        
        return nil
    }, 
    goconfig.CallbackPriorityHigh,  // 高优先级
    false)                          // 同步执行
```

---

## 🔄 热更新配置

### ⚙️ 热更新配置选项

```go
type HotReloadConfig struct {
    Enabled    bool          // 是否启用热更新
    Debounce   time.Duration // 防抖延迟，避免频繁更新
    MaxRetries int           // 配置加载失败时的最大重试次数
}

// 默认配置
defaultConfig := &goconfig.HotReloadConfig{
    Enabled:    true,
    Debounce:   time.Second * 1,
    MaxRetries: 3,
}
```

### 🎯 热更新使用示例

#### 基础热更新

```go
// 使用默认热更新配置
manager, err := goconfig.NewManager(config).
    WithConfigPath("./app.yaml").
    WithHotReload(nil).  // nil = 使用默认配置
    BuildAndStart()
```

#### 自定义热更新

```go
// 自定义热更新配置
hotReloadConfig := &goconfig.HotReloadConfig{
    Enabled:    true,
    Debounce:   time.Second * 5,  // 5秒防抖
    MaxRetries: 10,               // 最多重试10次
}

manager, err := goconfig.NewManager(config).
    WithConfigPath("./app.yaml").
    WithHotReload(hotReloadConfig).
    BuildAndStart()
```

#### 禁用热更新

```go
// 完全禁用热更新
disabledConfig := &goconfig.HotReloadConfig{
    Enabled: false,
}

manager, err := goconfig.NewManager(config).
    WithConfigPath("./app.yaml").
    WithHotReload(disabledConfig).
    BuildAndStart()
```

### 📋 手动重新加载

```go
// 手动触发配置重新加载
ctx := context.Background()
if err := manager.ReloadConfig(ctx); err != nil {
    log.Printf("重新加载失败: %v", err)
} else {
    log.Println("✅ 配置重新加载成功")
}
```

---

## 🔔 回调系统

### 📋 回调类型

```go
// 支持的回调类型
const (
    CallbackTypeConfigChanged  // 配置文件变更
    CallbackTypeReloaded      // 配置重新加载完成
    CallbackTypeError         // 错误事件
    CallbackTypeValidation    // 配置验证
)
```

### 🎯 回调优先级

```go
// 内置优先级常量
const (
    CallbackPriorityVeryHigh = -2000
    CallbackPriorityHigh     = -1000
    CallbackPriorityNormal   = 0
    CallbackPriorityLow      = 1000
    CallbackPriorityVeryLow  = 2000
)
```

### 🔧 回调注册示例

#### 配置变更回调

```go
// 配置变更回调
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    log.Printf("📝 配置已更新: %s", event.Source)
    
    // 获取新的配置
    newConfig, err := goconfig.GetConfigAs[AppConfig](manager)
    if err != nil {
        return fmt.Errorf("获取新配置失败: %w", err)
    }
    
    // 应用新配置
    if err := applyNewConfig(newConfig); err != nil {
        return fmt.Errorf("应用配置失败: %w", err)
    }
    
    return nil
}, goconfig.CallbackOptions{
    ID:       "config_applier",
    Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
    Priority: goconfig.CallbackPriorityHigh,
    Async:    false,           // 同步执行
    Timeout:  time.Second * 10, // 10秒超时
})
```

#### 错误处理回调

```go
// 错误处理回调
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    // 记录错误日志
    log.Printf("❌ 配置错误: %s, 来源: %s", event.Error, event.Source)
    
    // 发送告警通知
    if err := sendAlert(event.Error.Error()); err != nil {
        log.Printf("发送告警失败: %v", err)
    }
    
    // 记录到监控系统
    recordErrorMetric(event.Error)
    
    return nil
}, goconfig.CallbackOptions{
    ID:       "error_handler",
    Types:    []goconfig.CallbackType{goconfig.CallbackTypeError},
    Priority: goconfig.CallbackPriorityNormal,
    Async:    true,            // 异步执行，不阻塞主流程
    Timeout:  time.Second * 30, // 30秒超时
})
```

#### 多事件类型回调

```go
// 监听多种事件类型
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    switch event.Type {
    case goconfig.CallbackTypeConfigChanged:
        log.Println("🔄 配置已变更")
    case goconfig.CallbackTypeReloaded:
        log.Println("✅ 配置重载完成")
    case goconfig.CallbackTypeError:
        log.Printf("❌ 发生错误: %v", event.Error)
    }
    return nil
}, goconfig.CallbackOptions{
    ID: "multi_event_handler",
    Types: []goconfig.CallbackType{
        goconfig.CallbackTypeConfigChanged,
        goconfig.CallbackTypeReloaded,
        goconfig.CallbackTypeError,
    },
    Priority: goconfig.CallbackPriorityNormal,
})
```

### 🗑️ 回调管理

#### 注销回调

```go
// 注销指定ID的回调
if err := manager.UnregisterConfigCallback("config_applier"); err != nil {
    log.Printf("注销回调失败: %v", err)
}
```

#### 环境变更回调

```go
// 注册环境变更回调
manager.RegisterEnvironmentCallback("env_monitor",
    func(oldEnv, newEnv goconfig.EnvironmentType) error {
        log.Printf("🌍 环境变更: %s → %s", oldEnv, newEnv)
        
        // 根据新环境调整配置
        switch newEnv {
        case goconfig.EnvProduction:
            // 切换到生产模式
            enableProductionLogging()
        case goconfig.EnvDevelopment:
            // 切换到开发模式
            enableDebugLogging()
        }
        
        return nil
    },
    goconfig.CallbackPriorityHigh, // 高优先级
    false)                         // 同步执行
```

---

## 🛡️ 类型安全

### ✅ 泛型配置获取

新的API提供了完全类型安全的配置获取方式：

```go
// ✅ 推荐：类型安全的获取方式
config, err := goconfig.GetConfigAs[MyAppConfig](manager)
if err != nil {
    log.Fatal("配置类型转换失败:", err)
}

// ❌ 旧方式：需要手动类型断言
configInterface := manager.GetConfig()
config, ok := configInterface.(*MyAppConfig)
if !ok {
    log.Fatal("配置类型断言失败")
}
```

### 🎯 Must版本函数

对于确定类型正确的场景，可以使用Must版本：

```go
// 失败时panic，适用于应用启动阶段
config := goconfig.MustGetConfigAs[MyAppConfig](manager)

// 构建阶段的Must版本
manager := goconfig.NewManager(config).
    WithConfigPath("./app.yaml").
    MustBuildAndStart()  // 失败时panic
```

### 🔍 运行时类型检查

```go
// 验证配置类型
func validateConfigType[T any](manager *goconfig.IntegratedConfigManager) error {
    _, err := goconfig.GetConfigAs[T](manager)
    if err != nil {
        return fmt.Errorf("配置类型验证失败: %w", err)
    }
    return nil
}

// 使用
if err := validateConfigType[MyAppConfig](manager); err != nil {
    log.Fatal("配置类型不匹配:", err)
}
```

---

## 📁 配置发现

### 🎯 四种发现模式详解

#### 1. 直接路径模式

最简单直接的方式，适用于明确知道配置文件位置的场景：

```go
// 直接指定配置文件路径
manager, err := goconfig.NewManager(config).
    WithConfigPath("./configs/production.yaml").
    BuildAndStart()

// 支持相对路径和绝对路径
manager, err := goconfig.NewManager(config).
    WithConfigPath("/etc/myapp/config.json").
    BuildAndStart()
```

**适用场景：**
- 容器化部署，配置文件位置固定
- 简单应用，只有一个配置文件
- 生产环境，配置路径标准化

#### 2. 前缀匹配模式

根据配置前缀和环境自动匹配文件：

```go
// 前缀匹配示例
manager, err := goconfig.NewManager(config).
    WithSearchPath("./config").
    WithPrefix("gateway").
    WithEnvironment(goconfig.EnvDevelopment).
    BuildAndStart()
```

**匹配规则：**
```
搜索路径: ./config/
前缀: gateway
环境: dev

匹配顺序:
1. gateway-dev.yaml    ✅ 最优匹配
2. gateway-dev.yml
3. gateway-dev.json
4. gateway-dev.toml
5. gateway.yaml        📋 通用配置
6. gateway.yml
7. gateway.json
```

**目录结构示例：**
```
config/
├── gateway-dev.yaml      # 开发环境配置
├── gateway-test.yaml     # 测试环境配置
├── gateway-prod.yaml     # 生产环境配置
├── gateway.yaml          # 通用配置
└── database-dev.yaml     # 其他服务配置
```

#### 3. 模式匹配模式

使用glob模式进行灵活匹配：

```go
// glob模式匹配
manager, err := goconfig.NewManager(config).
    WithSearchPath("./config").
    WithPattern("service-*.yaml").
    WithEnvironment(goconfig.EnvProduction).
    BuildAndStart()
```

**支持的模式：**
```go
// 匹配所有yaml文件
WithPattern("*.yaml")

// 匹配以config开头的json文件
WithPattern("config-*.json")

// 匹配特定目录下的文件
WithPattern("*/app.yaml")

// 复杂模式匹配
WithPattern("**/config/**/*.{yaml,yml,json}")
```

#### 4. 自动发现模式

智能扫描目录，自动选择最佳配置文件：

```go
// 自动发现模式
manager, err := goconfig.NewManager(config).
    WithSearchPath("./config").
    WithEnvironment(goconfig.GetEnvironment()).
    BuildAndStart()
```

**发现优先级：**
1. 环境特定配置（如：`app-prod.yaml`）
2. 通用配置文件（如：`app.yaml`）
3. 默认配置文件（如：`config.yaml`）

### 🔍 配置文件格式支持

| 格式 | 扩展名 | 特点 | 推荐场景 |
|------|-------|------|----------|
| YAML | `.yaml`, `.yml` | 人类可读，支持注释 | 大多数场景 |
| JSON | `.json` | 标准格式，工具支持好 | API配置，CI/CD |
| TOML | `.toml` | 简洁清晰 | Go项目配置 |
| Properties | `.properties` | Java风格 | Spring项目迁移 |

### 📋 配置发现调试

使用内置工具查看配置发现过程：

```go
// 调试配置发现
configFiles, err := goconfig.ScanAndDisplayConfigs("./config", goconfig.EnvDevelopment)
if err != nil {
    log.Fatal(err)
}

// 输出示例：
// 📋 配置文件发现报告:
// 🔍 搜索路径: ./config
// 🌍 目标环境: dev
// 
// ✅ 发现的现有配置文件:
//    1. gateway-dev.yaml (环境: dev, 优先级: 100)
//    2. gateway.yaml (环境: default, 优先级: 50)
//    3. app-dev.json (环境: dev, 优先级: 100)
```

---

## 💼 实际应用场景

### 🌐 场景1：微服务网关

```go
package main

import (
    "context"
    "log"
    "net/http"
    "time"
    
    goconfig "github.com/kamalyes/go-config"
    "github.com/kamalyes/go-config/pkg/gateway"
)

func main() {
    // 网关配置结构
    config := &gateway.Gateway{}
    
    // 创建配置管理器
    manager, err := goconfig.NewManager(config).
        WithSearchPath("./config").
        WithPrefix("gateway").
        WithEnvironment(goconfig.GetEnvironment()).
        WithHotReload(&goconfig.HotReloadConfig{
            Enabled:  true,
            Debounce: time.Second * 2,
        }).
        BuildAndStart()
    
    if err != nil {
        log.Fatal("配置管理器启动失败:", err)
    }
    defer manager.Stop()
    
    // 注册HTTP服务器重启回调
    manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
        log.Println("🔄 配置变更，准备重启HTTP服务器...")
        return restartHTTPServer(ctx, manager)
    }, goconfig.CallbackOptions{
        ID:       "http_restart",
        Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
        Priority: goconfig.CallbackPriorityHigh,
    })
    
    // 启动HTTP服务器
    startHTTPServer(manager)
}

func startHTTPServer(manager *goconfig.IntegratedConfigManager) {
    gatewayConfig, err := goconfig.GetConfigAs[gateway.Gateway](manager)
    if err != nil {
        log.Fatal("获取配置失败:", err)
    }
    
    server := &http.Server{
        Addr:         fmt.Sprintf("%s:%d", gatewayConfig.HTTPServer.Host, gatewayConfig.HTTPServer.Port),
        ReadTimeout:  time.Duration(gatewayConfig.HTTPServer.ReadTimeout) * time.Second,
        WriteTimeout: time.Duration(gatewayConfig.HTTPServer.WriteTimeout) * time.Second,
    }
    
    log.Printf("🚀 HTTP服务器启动: %s", server.Addr)
    if err := server.ListenAndServe(); err != nil {
        log.Printf("HTTP服务器错误: %v", err)
    }
}

func restartHTTPServer(ctx context.Context, manager *goconfig.IntegratedConfigManager) error {
    // 优雅重启HTTP服务器的逻辑
    log.Println("✅ HTTP服务器重启完成")
    return nil
}
```

### 📊 场景2：数据处理服务

```go
package main

import (
    "database/sql"
    "log"
    "time"
    
    goconfig "github.com/kamalyes/go-config"
    _ "github.com/go-sql-driver/mysql"
)

type DataServiceConfig struct {
    Name string `yaml:"name"`
    
    Database struct {
        Host         string `yaml:"host"`
        Port         int    `yaml:"port"`
        Username     string `yaml:"username"`
        Password     string `yaml:"password"`
        Database     string `yaml:"database"`
        MaxOpenConns int    `yaml:"max-open-conns"`
        MaxIdleConns int    `yaml:"max-idle-conns"`
    } `yaml:"database"`
    
    Processing struct {
        BatchSize   int           `yaml:"batch-size"`
        Interval    time.Duration `yaml:"interval"`
        Concurrency int           `yaml:"concurrency"`
    } `yaml:"processing"`
}

func main() {
    config := &DataServiceConfig{}
    
    // 容器化部署配置
    manager, err := goconfig.NewManager(config).
        WithConfigPath("/etc/dataservice/config.yaml").  // 容器内固定路径
        WithHotReload(&goconfig.HotReloadConfig{
            Enabled:  true,
            Debounce: time.Second * 5,  // 数据处理服务需要更长的防抖时间
        }).
        BuildAndStart()
    
    if err != nil {
        log.Fatal("配置管理器启动失败:", err)
    }
    defer manager.Stop()
    
    // 数据库连接管理
    var db *sql.DB
    
    // 配置变更时重新连接数据库
    manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
        log.Println("📊 配置变更，重新建立数据库连接...")
        
        if db != nil {
            db.Close()
        }
        
        newConfig, err := goconfig.GetConfigAs[DataServiceConfig](manager)
        if err != nil {
            return err
        }
        
        db, err = connectDatabase(newConfig)
        return err
    }, goconfig.CallbackOptions{
        ID:       "db_reconnect",
        Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
        Priority: goconfig.CallbackPriorityHigh,
    })
    
    // 初始数据库连接
    serviceConfig, _ := goconfig.GetConfigAs[DataServiceConfig](manager)
    db, err = connectDatabase(serviceConfig)
    if err != nil {
        log.Fatal("数据库连接失败:", err)
    }
    
    // 启动数据处理服务
    startDataProcessing(db, serviceConfig)
}

func connectDatabase(config *DataServiceConfig) (*sql.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
        config.Database.Username,
        config.Database.Password,
        config.Database.Host,
        config.Database.Port,
        config.Database.Database)
    
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    db.SetMaxOpenConns(config.Database.MaxOpenConns)
    db.SetMaxIdleConns(config.Database.MaxIdleConns)
    
    return db, nil
}

func startDataProcessing(db *sql.DB, config *DataServiceConfig) {
    log.Printf("🔄 开始数据处理: 批量大小=%d, 间隔=%v, 并发数=%d",
        config.Processing.BatchSize,
        config.Processing.Interval,
        config.Processing.Concurrency)
    
    // 数据处理逻辑...
}
```

### 🔄 场景3：定时任务服务

```go
package main

import (
    "context"
    "log"
    "time"
    
    goconfig "github.com/kamalyes/go-config"
    "github.com/robfig/cron/v3"
)

type SchedulerConfig struct {
    Name    string `yaml:"name"`
    Enabled bool   `yaml:"enabled"`
    
    Jobs []JobConfig `yaml:"jobs"`
}

type JobConfig struct {
    Name     string `yaml:"name"`
    Cron     string `yaml:"cron"`
    Enabled  bool   `yaml:"enabled"`
    Timeout  string `yaml:"timeout"`
}

func main() {
    config := &SchedulerConfig{}
    
    // 使用模式匹配发现配置
    manager, err := goconfig.NewManager(config).
        WithSearchPath("./config").
        WithPattern("scheduler-*.yaml").
        WithEnvironment(goconfig.GetEnvironment()).
        WithHotReload(nil).
        BuildAndStart()
    
    if err != nil {
        log.Fatal("配置管理器启动失败:", err)
    }
    defer manager.Stop()
    
    // Cron调度器
    c := cron.New(cron.WithSeconds())
    
    // 配置变更时重新加载任务
    manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
        log.Println("⏰ 任务配置变更，重新加载定时任务...")
        return reloadJobs(c, manager)
    }, goconfig.CallbackOptions{
        ID:       "job_reload",
        Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
        Priority: goconfig.CallbackPriorityNormal,
    })
    
    // 初始加载任务
    if err := reloadJobs(c, manager); err != nil {
        log.Fatal("加载定时任务失败:", err)
    }
    
    // 启动调度器
    c.Start()
    log.Println("🚀 定时任务调度器已启动")
    
    // 保持运行
    select {}
}

func reloadJobs(c *cron.Cron, manager *goconfig.IntegratedConfigManager) error {
    // 停止所有现有任务
    c.Stop()
    
    config, err := goconfig.GetConfigAs[SchedulerConfig](manager)
    if err != nil {
        return err
    }
    
    if !config.Enabled {
        log.Println("⏸️ 调度器已禁用")
        return nil
    }
    
    // 重新创建调度器
    c = cron.New(cron.WithSeconds())
    
    // 添加任务
    for _, job := range config.Jobs {
        if !job.Enabled {
            continue
        }
        
        jobFunc := createJobFunc(job)
        if _, err := c.AddFunc(job.Cron, jobFunc); err != nil {
            log.Printf("❌ 添加任务失败 %s: %v", job.Name, err)
            continue
        }
        
        log.Printf("✅ 已添加任务: %s (%s)", job.Name, job.Cron)
    }
    
    // 重新启动
    c.Start()
    return nil
}

func createJobFunc(job JobConfig) func() {
    return func() {
        log.Printf("🔄 执行任务: %s", job.Name)
        // 任务执行逻辑...
    }
}
```

---

## 🔧 高级用法

### 🔄 配置校验

```go
// 自定义配置校验
type ValidatableConfig interface {
    Validate() error
}

// 实现校验接口
func (c *MyConfig) Validate() error {
    if c.Server.Port <= 0 || c.Server.Port > 65535 {
        return fmt.Errorf("端口号无效: %d", c.Server.Port)
    }
    
    if c.Database.Host == "" {
        return fmt.Errorf("数据库主机不能为空")
    }
    
    return nil
}

// 注册校验回调
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    if validator, ok := event.NewValue.(ValidatableConfig); ok {
        if err := validator.Validate(); err != nil {
            return fmt.Errorf("配置校验失败: %w", err)
        }
    }
    return nil
}, goconfig.CallbackOptions{
    ID:       "config_validator",
    Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
    Priority: goconfig.CallbackPriorityVeryHigh, // 最高优先级，优先校验
})
```

### 🌐 上下文集成

```go
// HTTP中间件集成
func ConfigMiddleware(manager *goconfig.IntegratedConfigManager) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // 将配置注入到请求上下文
            ctx := manager.WithContext(r.Context())
            r = r.WithContext(ctx)
            
            next.ServeHTTP(w, r)
        })
    }
}

// 在HTTP处理器中获取配置
func MyHandler(w http.ResponseWriter, r *http.Request) {
    // 从上下文获取配置
    config := goconfig.FromContext(r.Context())
    if config != nil {
        // 使用配置...
    }
}
```

### 📊 配置监控

```go
// 配置元数据监控
func monitorConfig(manager *goconfig.IntegratedConfigManager) {
    ticker := time.NewTicker(time.Minute * 5)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            metadata := manager.GetConfigMetadata()
            
            log.Printf("📊 配置监控:")
            log.Printf("  配置文件: %s", metadata["config_path"])
            log.Printf("  环境: %s", metadata["environment"])
            log.Printf("  运行状态: %v", metadata["running"])
            log.Printf("  更新时间: %v", metadata["updated_at"])
            log.Printf("  热更新: %v", metadata["hot_reload_enabled"])
        }
    }
}
```

### 🔧 自定义配置源

```go
// 自定义配置加载逻辑
func loadConfigFromMultipleSources(manager *goconfig.IntegratedConfigManager) error {
    // 1. 从配置文件加载基础配置
    baseConfig := manager.GetConfig()
    
    // 2. 从环境变量覆盖
    if err := loadFromEnvVars(baseConfig); err != nil {
        return err
    }
    
    // 3. 从远程配置中心加载
    if err := loadFromRemoteConfig(baseConfig); err != nil {
        return err
    }
    
    // 4. 从命令行参数覆盖
    if err := loadFromCmdArgs(baseConfig); err != nil {
        return err
    }
    
    return nil
}
```

---

## ❓ 故障排除

### 🔍 常见问题

#### 1. 配置文件找不到

```go
// 问题：配置文件找不到
// 错误：config file not found

// 解决方案1：检查路径
manager, err := goconfig.NewManager(config).
    WithConfigPath("./config/app.yaml"). // 确保路径正确
    BuildAndStart()

// 解决方案2：使用绝对路径
absPath, _ := filepath.Abs("./config/app.yaml")
manager, err := goconfig.NewManager(config).
    WithConfigPath(absPath).
    BuildAndStart()

// 解决方案3：调试配置发现
configs, err := goconfig.ScanAndDisplayConfigs("./config", goconfig.GetEnvironment())
// 查看输出，确认配置文件是否存在
```

#### 2. 类型转换失败

```go
// 问题：配置类型转换失败
// 错误：config type mismatch

// 解决方案1：确保结构体标签正确
type MyConfig struct {
    Port int `yaml:"port" json:"port"` // 确保标签匹配配置文件格式
}

// 解决方案2：验证配置文件格式
// YAML文件示例：
// port: 8080  # 确保是数字，不是字符串 "8080"

// 解决方案3：使用类型安全的获取方法
config, err := goconfig.GetConfigAs[MyConfig](manager)
if err != nil {
    log.Printf("类型转换失败: %v", err)
    // 检查配置结构体定义
}
```

#### 3. 热更新不工作

```go
// 问题：配置文件修改后热更新不生效

// 解决方案1：确保热更新已启用
manager, err := goconfig.NewManager(config).
    WithConfigPath("./app.yaml").
    WithHotReload(&goconfig.HotReloadConfig{
        Enabled:  true,  // 确保启用
        Debounce: time.Second * 1,
    }).
    BuildAndStart()

// 解决方案2：检查文件权限
// 确保应用有读取配置文件的权限

// 解决方案3：检查回调注册
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    log.Println("✅ 配置已更新") // 添加日志确认回调被触发
    return nil
}, goconfig.CallbackOptions{
    ID: "debug_callback",
    Types: []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
})
```

#### 4. 环境变量不生效

```go
// 问题：设置的环境变量不生效

// 解决方案1：检查环境变量名
// 支持的环境变量名：
// APP_ENV, GO_ENV, ENV, ENVIRONMENT

// 解决方案2：手动设置环境
manager, err := goconfig.NewManager(config).
    WithEnvironment(goconfig.EnvProduction). // 明确指定环境
    BuildAndStart()

// 解决方案3：调试环境检测
currentEnv := goconfig.GetEnvironment()
log.Printf("检测到的环境: %s", currentEnv)
```

### 🛠️ 调试技巧

#### 启用详细日志

```go
// 启用详细的内部日志
goconfig.SetLogLevel(goconfig.LogLevelDebug)

// 或者使用自定义日志器
customLogger := log.New(os.Stdout, "[CONFIG] ", log.LstdFlags)
goconfig.SetLogger(customLogger)
```

#### 配置验证

```go
// 添加配置验证回调
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    // 打印配置内容用于调试
    configJSON, _ := json.MarshalIndent(event.NewValue, "", "  ")
    log.Printf("当前配置:\n%s", configJSON)
    return nil
}, goconfig.CallbackOptions{
    ID: "debug_printer",
    Types: []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
})
```

#### 运行时诊断

```go
// 获取详细的运行时信息
func diagnoseConfigManager(manager *goconfig.IntegratedConfigManager) {
    metadata := manager.GetConfigMetadata()
    
    fmt.Println("📋 配置管理器诊断信息:")
    fmt.Printf("  运行状态: %v\n", manager.IsRunning())
    fmt.Printf("  配置文件: %s\n", metadata["config_path"])
    fmt.Printf("  当前环境: %s\n", manager.GetEnvironment())
    fmt.Printf("  热更新: %v\n", metadata["hot_reload_enabled"])
    fmt.Printf("  创建时间: %v\n", metadata["created_at"])
    fmt.Printf("  更新时间: %v\n", metadata["updated_at"])
    
    // 验证配置
    if err := manager.ValidateConfig(); err != nil {
        fmt.Printf("  ❌ 配置验证失败: %v\n", err)
    } else {
        fmt.Println("  ✅ 配置验证通过")
    }
}
```

---

## 📋 最佳实践

### ✅ 推荐做法

1. **使用链式API**：优先使用 `NewManager().WithXxx().BuildAndStart()`
2. **类型安全获取**：使用 `GetConfigAs[T]()` 替代类型断言
3. **合理设置防抖**：热更新防抖时间根据应用特性调整
4. **优雅错误处理**：注册错误回调进行统一处理
5. **配置验证**：实现 `Validate()` 方法确保配置正确性

### ❌ 避免的做法

1. **不要**忘记调用 `defer manager.Stop()`
2. **不要**在回调中执行耗时操作（使用异步回调）
3. **不要**忽略错误处理
4. **不要**在生产环境禁用热更新（除非有特殊需求）
5. **不要**在回调中panic

---

## 🎯 总结

go-config 提供了强大而灵活的配置管理能力：

- **🔗 流畅的链式API**：简洁易用的配置方式
- **🛡️ 类型安全**：编译期类型检查，运行期类型安全
- **🌍 多环境支持**：无缝的环境切换和管理
- **🔄 智能热更新**：实时响应配置变化
- **📁 灵活发现**：多种配置文件发现模式
- **🔔 事件驱动**：强大的回调机制

通过本文档的指导，你应该能够：
- 快速集成 go-config 到你的项目中
- 选择合适的配置发现模式
- 实现类型安全的配置管理
- 构建响应配置变化的健壮应用

如果你有任何问题或建议，欢迎在 [GitHub Issues](https://github.com/kamalyes/go-config/issues) 中反馈。