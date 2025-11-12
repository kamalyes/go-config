# Go-Config 配置热更新完整使用指南

## 概述

`go-config` 是一个强大的Go语言配置管理库，支持智能配置发现、配置热更新、环境管理、回调机制和上下文集成。本文档详细介绍如何使用该库，包括配置自动发现、数据提取、回调配置、热更新等核心功能。

## 目录

1. [快速开始](#快速开始)
2. [核心概念](#核心概念)
3. [智能配置发现](#智能配置发现)
4. [配置数据提取](#配置数据提取)
5. [回调机制详解](#回调机制详解)
6. [热更新配置](#热更新配置)
7. [环境管理](#环境管理)
8. [上下文集成](#上下文集成)
9. [美化日志输出](#美化日志输出)
10. [完整示例](#完整示例)
11. [最佳实践](#最佳实践)

## 快速开始

### 安装

```bash
go mod init your-project
go get github.com/kamalyes/go-config
```

### 🚀 推荐方式：智能配置发现

```go
package main

import (
    "log"
    
    goconfig "github.com/kamalyes/go-config"
    "github.com/kamalyes/go-config/pkg/gateway"
)

func main() {
    // 1. 创建配置实例
    config := &gateway.Gateway{}
    
    // 2. 使用智能配置发现创建管理器（推荐）
    manager, err := goconfig.CreateAndStartIntegratedManagerWithAutoDiscovery(
        config,
        "./config",  // 配置目录，自动发现配置文件
        goconfig.GetEnvironment(),
        "gateway",   // 配置类型，用于匹配文件名
    )
    if err != nil {
        log.Fatal("创建配置管理器失败:", err)
    }
    defer manager.Stop()
    
    // 3. 获取配置数据
    gatewayConfig := manager.GetConfig().(*gateway.Gateway)
    log.Printf("服务名称: %s", gatewayConfig.Name)
    
    // 4. 保持程序运行
    select {}
}
```

## 核心概念

### 1. 配置管理器类型

- **IntegratedConfigManager**: 集成配置管理器，统一管理所有功能
- **HotReloader**: 热更新管理器，监控文件变化
- **EnvironmentManager**: 环境管理器，管理环境变量
- **ContextManager**: 上下文管理器，处理上下文集成

### 2. 配置类型

- **Gateway**: 网关配置，包含HTTP、GRPC、数据库、缓存等
- **SingleConfig**: 单一配置类型（自定义配置）
- **MultiConfig**: 多配置类型（自定义配置）

### 3. 回调类型

```go
type CallbackType int

const (
    CallbackTypeConfigChanged CallbackType = iota
    CallbackTypeReloaded
    CallbackTypeError
    CallbackTypeValidation
)
```

## 智能配置发现

### 配置发现概述

智能配置发现是框架的核心特性之一，能够根据环境和配置类型自动查找和加载合适的配置文件，无需手动指定文件路径。

### 配置文件命名规则

框架按照以下优先级顺序查找配置文件：

1. **环境优先模式**: `{type}-{env}.{ext}`
   - 示例：`gateway-dev.yaml`, `gateway-prod.json`

2. **类型优先模式**: `{type}.{ext}`
   - 示例：`gateway.yaml`, `database.json`

3. **环境目录模式**: `{env}/{type}.{ext}`
   - 示例：`dev/gateway.yaml`, `prod/database.json`

4. **通用配置模式**: `config.{ext}`, `application.{ext}`
   - 示例：`config.yaml`, `application.json`

### 支持的配置格式

- ✅ **YAML**: `.yaml`, `.yml`
- ✅ **JSON**: `.json`
- ✅ **TOML**: `.toml`
- ✅ **Properties**: `.properties`

### 使用方法

#### 1. 目录自动发现

```go
package main

import (
    goconfig "github.com/kamalyes/go-config"
    "github.com/kamalyes/go-config/pkg/gateway"
)

func main() {
    config := &gateway.Gateway{}
    
    // 在 ./config 目录中自动查找 gateway 相关配置
    manager, err := goconfig.CreateAndStartIntegratedManagerWithAutoDiscovery(
        config,
        "./config",        // 配置目录
        goconfig.EnvironmentDevelopment, // 环境类型
        "gateway",         // 配置类型
    )
    if err != nil {
        log.Fatalf("配置管理器创建失败: %v", err)
    }
    defer manager.Stop()
    
    // 获取配置
    gatewayConfig := manager.GetConfig().(*gateway.Gateway)
    fmt.Printf("服务名称: %s\n", gatewayConfig.Name)
}
```

#### 2. 文件模式匹配

```go
// 使用通配符模式查找配置文件
manager, err := goconfig.CreateIntegratedManagerWithPattern(
    config,
    "gateway-*.yaml",  // 文件模式
    goconfig.EnvironmentDevelopment,
)

// 匹配文件例子：
// ✅ gateway-dev.yaml
// ✅ gateway-prod.yaml  
// ✅ gateway-staging.yaml
// ❌ database.yaml (不匹配模式)
```

#### 3. 配置扫描和展示

```go
// 扫描指定目录并展示所有发现的配置文件
discoveredFiles := goconfig.ScanAndDisplayConfigs("./config", "gateway")

// 输出示例：
// 🔍 [配置发现] 在目录 ./config 中扫描 gateway 配置
// ├── ✅ 找到配置文件: ./config/gateway-dev.yaml
// ├── ✅ 找到配置文件: ./config/gateway.json  
// └── 📁 共发现 2 个配置文件

fmt.Printf("发现 %d 个配置文件\n", len(discoveredFiles))
```

### 最佳配置文件组织

#### 推荐的目录结构

```text
project/
├── config/
│   ├── gateway-dev.yaml      # 开发环境网关配置
│   ├── gateway-prod.yaml     # 生产环境网关配置
│   ├── database-dev.yaml     # 开发环境数据库配置
│   ├── database-prod.yaml    # 生产环境数据库配置
│   ├── redis.yaml            # 通用Redis配置
│   └── common/
│       ├── logging.yaml      # 日志配置
│       └── monitoring.yaml   # 监控配置
├── environments/
│   ├── dev/
│   │   ├── gateway.yaml      # 开发环境特定配置
│   │   └── secrets.yaml      # 开发环境密钥
│   └── prod/
│       ├── gateway.yaml      # 生产环境特定配置
│       └── secrets.yaml      # 生产环境密钥
└── main.go
```

#### 环境变量集成

```go
// 通过环境变量指定配置目录
configDir := os.Getenv("CONFIG_DIR")
if configDir == "" {
    configDir = "./config"  // 默认配置目录
}

// 通过环境变量指定环境类型
env := goconfig.GetEnvironmentFromEnv("APP_ENV", goconfig.EnvironmentDevelopment)

manager, err := goconfig.CreateAndStartIntegratedManagerWithAutoDiscovery(
    config,
    configDir,
    env,
    "gateway",
)
```

### 配置发现的错误处理

```go
manager, err := goconfig.CreateAndStartIntegratedManagerWithAutoDiscovery(
    config,
    "./config",
    goconfig.EnvironmentDevelopment,
    "gateway",
)

switch {
case errors.Is(err, goconfig.ErrConfigFileNotFound):
    log.Println("⚠️ 未找到配置文件，使用默认配置")
    // 创建默认配置文件
    defaultConfigPath := "./config/gateway-dev.yaml"
    if createErr := goconfig.CreateDefaultConfigFile(defaultConfigPath, config); createErr != nil {
        log.Fatalf("创建默认配置失败: %v", createErr)
    }
    
case errors.Is(err, goconfig.ErrInvalidConfigFormat):
    log.Fatalf("配置文件格式无效: %v", err)
    
case err != nil:
    log.Fatalf("配置管理器初始化失败: %v", err)
}
```

## 配置数据提取

### 1. 基本数据获取

```go
// 获取完整配置对象
config := manager.GetConfig()

// 类型断言获取具体配置
gatewayConfig := config.(*gateway.Gateway)

// 获取配置元数据
metadata := manager.GetConfigMetadata()
```

### 2. 特定配置提取

```go
// HTTP服务器配置
httpConfig := gatewayConfig.HTTPServer
if httpConfig != nil {
    fmt.Printf("服务地址: %s:%d\n", httpConfig.Host, httpConfig.Port)
    fmt.Printf("端点: %s\n", httpConfig.GetEndpoint())
}

// 数据库配置
dbConfig := gatewayConfig.Database
if dbConfig != nil {
    fmt.Printf("数据库: %s@%s:%s/%s\n", 
        dbConfig.Username, dbConfig.Host, dbConfig.Port, dbConfig.Dbname)
}

// 缓存配置
cacheConfig := gatewayConfig.Cache
if cacheConfig != nil {
    fmt.Printf("缓存类型: %s, TTL: %v\n", 
        cacheConfig.Type, cacheConfig.DefaultTTL)
}
```

### 3. 环境相关数据

```go
// 获取当前环境
env := manager.GetEnvironment()
fmt.Printf("当前环境: %s\n", env)

// 从上下文获取环境
ctx := context.Background()
ctxWithEnv := manager.GetContextManager().WithEnvironment(ctx, env)
envFromCtx := goconfig.GetEnvironmentFromContext(ctxWithEnv)
```

## 回调机制详解

### 1. 回调函数类型

```go
// 配置回调函数
type ConfigCallbackFunc func(ctx context.Context, event CallbackEvent) error

// 环境回调函数  
type EnvironmentCallbackFunc func(oldEnv, newEnv EnvironmentType) error
```

### 2. 注册配置回调

```go
// 基础回调注册
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    log.Printf("配置变更: %s -> %s", event.Source, event.Type)
    
    // 使用配置格式化工具输出详细信息
    goconfig.LogConfigChange(event, event.NewValue)
    
    return nil
}, goconfig.CallbackOptions{
    ID:       "my_config_callback",
    Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
    Priority: goconfig.CallbackPriorityHigh,
    Async:    false,
    Timeout:  5 * time.Second,
})

// 高级回调配置
manager.RegisterConfigCallback(configChangeHandler, goconfig.CallbackOptions{
    ID:       "advanced_callback",
    Types:    []goconfig.CallbackType{
        goconfig.CallbackTypeConfigChanged,
        goconfig.CallbackTypeReloaded,
        goconfig.CallbackTypeError,
    },
    Priority: goconfig.CallbackPriorityHighest,
    Async:    true,
    Timeout:  10 * time.Second,
    Metadata: map[string]interface{}{
        "component": "gateway",
        "version":   "v1.0.0",
    },
})
```

### 3. 注册环境回调

```go
manager.RegisterEnvironmentCallback(
    "env_change_handler",
    func(oldEnv, newEnv goconfig.EnvironmentType) error {
        // 使用配置格式化工具
        goconfig.LogEnvChange(oldEnv, newEnv)
        
        // 自定义逻辑
        switch newEnv {
        case goconfig.EnvProduction:
            // 生产环境特殊处理
            return setupProductionMode()
        case goconfig.EnvDevelopment:
            // 开发环境特殊处理
            return setupDevelopmentMode()
        }
        return nil
    },
    goconfig.CallbackPriorityHigh,
    false, // 同步执行
)
```

### 4. 回调优先级

```go
const (
    CallbackPriorityLowest   = 0
    CallbackPriorityLow      = 100
    CallbackPriorityMedium   = 200
    CallbackPriorityHigh     = 300
    CallbackPriorityHighest  = 400
)
```

### 5. 错误处理回调

```go
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    if event.Type == goconfig.CallbackTypeError {
        // 记录错误
        goconfig.LogConfigError(event)
        
        // 发送告警
        return sendAlert(event.Error)
    }
    return nil
}, goconfig.CallbackOptions{
    ID:       "error_handler",
    Types:    []goconfig.CallbackType{goconfig.CallbackTypeError},
    Priority: goconfig.CallbackPriorityHighest,
    Async:    true,
})
```

## 热更新配置

### 1. 自动热更新

```go
// 创建配置管理器时自动启用热更新
manager, err := goconfig.CreateAndStartIntegratedManager(
    config,
    "config.yaml",  // 监控此文件
    goconfig.GetEnvironment(),
)
// 文件变更时自动触发热更新
```

### 2. 手动触发热更新

```go
// 手动重新加载配置
ctx := context.WithTimeout(context.Background(), 10*time.Second)
err := manager.ReloadConfig(ctx)
if err != nil {
    log.Printf("手动重载失败: %v", err)
}
```

### 3. 验证配置

```go
// 验证当前配置
err := manager.ValidateConfig()
if err != nil {
    log.Printf("配置验证失败: %v", err)
}
```

### 4. 热更新回调

```go
// 监听热更新事件
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    switch event.Type {
    case goconfig.CallbackTypeConfigChanged:
        log.Println("配置文件已更改")
        // 重新初始化相关组件
        return reinitializeComponents(event.NewValue)
        
    case goconfig.CallbackTypeReloaded:
        log.Println("配置已重新加载")
        // 通知其他服务配置更新
        return notifyOtherServices()
        
    case goconfig.CallbackTypeError:
        log.Printf("热更新错误: %s", event.Error)
        return nil
    }
    return nil
}, goconfig.CallbackOptions{
    ID:    "hot_reload_handler",
    Types: []goconfig.CallbackType{
        goconfig.CallbackTypeConfigChanged,
        goconfig.CallbackTypeReloaded,
        goconfig.CallbackTypeError,
    },
})
```

## 环境管理

### 1. 环境类型

```go
const (
    EnvDevelopment EnvironmentType = "development"
    EnvTest        EnvironmentType = "test"
    EnvStaging     EnvironmentType = "staging"
    EnvProduction  EnvironmentType = "production"
)
```

### 2. 设置和获取环境

```go
// 设置环境变量
os.Setenv("APP_ENV", "production")

// 获取当前环境
env := goconfig.GetEnvironment()

// 获取环境管理器
envManager := manager.GetEnvironmentManager()
```

### 3. 环境变更监听

```go
// 监听环境变更
manager.RegisterEnvironmentCallback("env_monitor", func(oldEnv, newEnv goconfig.EnvironmentType) error {
    log.Printf("环境从 %s 切换到 %s", oldEnv, newEnv)
    
    // 环境相关配置调整
    switch newEnv {
    case goconfig.EnvProduction:
        // 关闭调试日志
        logger.SetGlobalLevel(logger.INFO)
        // 启用性能监控
        return enablePerformanceMonitoring()
    case goconfig.EnvDevelopment:
        // 启用调试日志
        logger.SetGlobalLevel(logger.DEBUG)
        // 启用开发工具
        return enableDevelopmentTools()
    }
    return nil
}, goconfig.CallbackPriorityHigh, false)
```

## 上下文集成

### 1. 基本上下文操作

```go
// 创建带配置的上下文
ctx := context.Background()
ctxManager := manager.GetContextManager()

// 将配置注入上下文
ctxWithConfig := ctxManager.WithConfig(ctx, config)

// 将环境注入上下文
ctxWithEnv := ctxManager.WithEnvironment(ctx, goconfig.EnvProduction)
```

### 2. 从上下文获取数据

```go
// 从上下文获取配置
config := goconfig.GetConfigFromContext(ctx)
if gatewayConfig, ok := config.(*gateway.Gateway); ok {
    log.Printf("从上下文获取到网关配置: %s", gatewayConfig.Name)
}

// 从上下文获取环境
env := goconfig.GetEnvironmentFromContext(ctx)
log.Printf("从上下文获取到环境: %s", env)
```

### 3. 中间件中的使用

```go
func configMiddleware(manager *goconfig.IntegratedConfigManager) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // 将配置注入请求上下文
            config := manager.GetConfig()
            ctx := manager.GetContextManager().WithConfig(r.Context(), config)
            
            // 将环境注入上下文
            env := manager.GetEnvironment()
            ctx = manager.GetContextManager().WithEnvironment(ctx, env)
            
            // 继续处理请求
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

## 美化日志输出

### 日志输出概述

框架内置了强大的配置格式化器，使用 `github.com/kamalyes/go-logger` 提供美观的Emoji丰富的日志输出，让配置变更一目了然。

### 1. 自动日志输出

框架会自动为配置变更生成美化的日志输出：

```go
// 在回调中自动触发美化日志
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    // 框架自动输出美化的配置变更日志
    goconfig.LogConfigChange(event, event.NewValue)
    
    // 你的业务逻辑...
    return handleConfigUpdate(event.NewValue)
}, options)

// 手动记录环境变更
goconfig.LogEnvChange(oldEnv, newEnv)
```

### 2. 配置变更日志示例

当Gateway配置发生变更时，会自动输出如下美化日志：

```text
📝 [配置变更] Gateway 配置已更新
├── 📁 配置文件: ./config/gateway-dev.yaml
├── 🌍 运行环境: development  
├── 🕐 变更时间: 2024-01-15 14:30:25
├── 🏷️  服务名称: user-gateway-service
├── 📦 服务版本: v1.2.0
├── 🌐 HTTP服务器: http://0.0.0.0:8080
│   ├── 📍 监听地址: 0.0.0.0:8080
│   ├── ⏱️  读取超时: 30s
│   ├── ⏱️  写入超时: 30s  
│   ├── 🔄 保持连接: true
│   └── 🗜️ Gzip压缩: 启用
├── 🔌 GRPC服务器: 0.0.0.0:9090
├── 💾 数据库配置:
│   ├── 📊 类型: MySQL
│   ├── 📍 地址: 127.0.0.1:3306
│   ├── 🗄️  数据库: userdb
│   ├── 👤 用户: root
│   ├── 🔗 最大连接: 100  
│   └── ⏱️  连接超时: 30s
├── ⚡ Redis配置:
│   ├── 📍 地址: localhost:6379
│   ├── 🗂️  数据库: 0
│   ├── 🏊 连接池: 100
│   └── ⏱️  超时: 5s
├── 📊 监控配置:
│   ├── 🔍 性能分析: 启用 (:6060)
│   ├── 📈 Prometheus: 启用 (:9090/metrics)
│   └── 🔎 链路追踪: Jaeger
├── 🔒 安全配置:
│   ├── 🌐 CORS: 启用
│   ├── 🎫 JWT: 启用
│   └── 🛡️  限流: 100 req/min
└── ✅ 配置加载成功
```

### 3. 环境变更日志示例

```text  
🌍 [环境变更] development -> production
├── 🕐 变更时间: 2024-01-15 14:32:18
├── 📋 触发原因: 环境变量更新
├── ⚠️  注意事项: 生产环境配置已生效
└── ✅ 环境切换完成
```

### 4. 错误日志示例

```text
❌ [配置错误] 配置文件加载失败
├── 📁 文件路径: ./config/gateway-prod.yaml
├── 🕐 错误时间: 2024-01-15 14:35:42  
├── 🐛 错误类型: 文件格式错误
├── 📋 错误详情: yaml: line 15: found character that cannot start any token
├── 💡 解决建议: 检查YAML文件第15行语法
└── 🔧 自动回滚: 使用上一次正确配置
```

### 5. 自定义日志格式化器

```go
// 创建自定义格式化器
customLogger := logger.GetGlobalLogger()
formatter := goconfig.NewConfigFormatter(customLogger)

// 可以自定义日志级别和格式
formatter.SetLogLevel(logger.InfoLevel)
formatter.SetShowTimestamp(true)
formatter.SetShowCaller(false)

// 设置为全局格式化器
goconfig.SetGlobalConfigFormatter(formatter)

// 手动输出格式化的配置信息
formatter.LogGatewayDetails(gatewayConfig)
```

### 6. 禁用自动日志输出

如果你不需要自动美化日志输出，可以禁用：

```go
// 禁用自动日志输出
goconfig.SetAutoLogEnabled(false)

// 或在创建管理器时指定
manager, err := goconfig.CreateAndStartIntegratedManagerWithOptions(
    config,
    configPath,
    env,
    &goconfig.ManagerOptions{
        AutoLogEnabled: false,  // 禁用自动日志
        LogLevel: logger.WarnLevel, // 只输出警告和错误
    },
)
      💾 数据库: gateway_db
      👤 用户: root
      🔗 最大连接: 100
   💾 缓存配置:
      🏷️ 类型: memory
      🚦 启用: ✅ 启用
      ⏰ 默认TTL: 30m0s
✅ 配置更新成功
```

## 完整示例

### Gateway服务示例

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    
    goconfig "github.com/kamalyes/go-config"
    "github.com/kamalyes/go-config/pkg/gateway"
)

func main() {
    // 1. 设置环境
    os.Setenv("APP_ENV", "development")
    
    // 2. 创建配置实例
    config := &gateway.Gateway{}
    
    // 3. 创建集成配置管理器
    manager, err := goconfig.CreateAndStartIntegratedManager(
        config,
        "gateway_config.yaml",
        goconfig.GetEnvironment(),
    )
    if err != nil {
        log.Fatalf("创建配置管理器失败: %v", err)
    }
    defer manager.Stop()
    
    // 4. 注册配置变更回调
    manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
        // 使用格式化工具输出详细配置信息
        goconfig.LogConfigChange(event, event.NewValue)
        
        // 自定义处理逻辑
        if gatewayConfig, ok := event.NewValue.(*gateway.Gateway); ok {
            // 验证配置
            if err := gatewayConfig.Validate(); err != nil {
                return fmt.Errorf("配置验证失败: %w", err)
            }
            
            // 动态更新服务
            return updateHTTPServer(gatewayConfig.HTTPServer)
        }
        return nil
    }, goconfig.CallbackOptions{
        ID:       "gateway_config_handler",
        Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
        Priority: goconfig.CallbackPriorityHigh,
        Async:    false,
        Timeout:  5 * time.Second,
    })
    
    // 5. 注册环境变更回调
    manager.RegisterEnvironmentCallback("gateway_env_handler", func(oldEnv, newEnv goconfig.EnvironmentType) error {
        goconfig.LogEnvChange(oldEnv, newEnv)
        return adjustEnvironmentSettings(newEnv)
    }, goconfig.CallbackPriorityHigh, false)
    
    // 6. 启动HTTP服务器
    startHTTPServer(manager)
}

func updateHTTPServer(httpConfig *gateway.HTTPServer) error {
    log.Printf("🔄 更新HTTP服务器配置: %s", httpConfig.GetEndpoint())
    // 实现服务器配置更新逻辑
    return nil
}

func adjustEnvironmentSettings(env goconfig.EnvironmentType) error {
    switch env {
    case goconfig.EnvProduction:
        log.Println("🚀 切换到生产环境配置")
        // 生产环境特殊设置
    case goconfig.EnvDevelopment:
        log.Println("🔧 切换到开发环境配置")
        // 开发环境特殊设置
    }
    return nil
}

func startHTTPServer(manager *goconfig.IntegratedConfigManager) {
    mux := http.NewServeMux()
    
    // 配置信息端点
    mux.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
        config := manager.GetConfig().(*gateway.Gateway)
        env := manager.GetEnvironment()
        
        response := map[string]interface{}{
            "name":        config.Name,
            "version":     config.Version,
            "environment": env,
            "enabled":     config.Enabled,
            "debug":       config.Debug,
            "timestamp":   time.Now().Format(time.RFC3339),
        }
        
        if config.HTTPServer != nil {
            response["http_server"] = map[string]interface{}{
                "host":     config.HTTPServer.Host,
                "port":     config.HTTPServer.Port,
                "endpoint": config.HTTPServer.GetEndpoint(),
            }
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    })
    
    // 手动重载端点
    mux.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        
        ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
        defer cancel()
        
        if err := manager.ReloadConfig(ctx); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "配置重新加载成功")
    })
    
    // 启动服务器
    config := manager.GetConfig().(*gateway.Gateway)
    addr := fmt.Sprintf("%s:%d", config.HTTPServer.Host, config.HTTPServer.Port)
    
    log.Printf("🚀 启动HTTP服务器: %s", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}
```

## 最佳实践

### 1. 配置文件组织

```yaml
# gateway_config.yaml
module-name: "gateway"
name: "网关服务"
enabled: true
debug: true
version: "v1.0.0"
environment: "development"

http:
  host: "0.0.0.0"
  port: 8080
  read-timeout: 30
  write-timeout: 30
  enable_gzip_compress: true

database:
  mysql:
    host: "127.0.0.1"
    port: "3306"
    username: "root"
    password: "password"
    db-name: "gateway_db"

cache:
  type: "memory"
  enabled: true
  default_ttl: "30m"
```

### 2. 错误处理

```go
// 配置加载错误处理
manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    if event.Type == goconfig.CallbackTypeError {
        // 记录错误
        goconfig.LogConfigError(event)
        
        // 发送告警
        if err := sendAlert(event.Error); err != nil {
            log.Printf("发送告警失败: %v", err)
        }
        
        // 尝试回滚到上一个有效配置
        return rollbackToLastValidConfig()
    }
    return nil
}, goconfig.CallbackOptions{
    ID:    "error_handler",
    Types: []goconfig.CallbackType{goconfig.CallbackTypeError},
})
```

### 3. 性能优化

```go
// 异步处理非关键回调
manager.RegisterConfigCallback(analyticsHandler, goconfig.CallbackOptions{
    ID:       "analytics",
    Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
    Priority: goconfig.CallbackPriorityLow,
    Async:    true,  // 异步执行
    Timeout:  30 * time.Second,
})

// 同步处理关键回调
manager.RegisterConfigCallback(coreConfigHandler, goconfig.CallbackOptions{
    ID:       "core_config",
    Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
    Priority: goconfig.CallbackPriorityHighest,
    Async:    false, // 同步执行
    Timeout:  5 * time.Second,
})
```

### 4. 测试配置

```go
func TestConfigHotReload(t *testing.T) {
    // 创建临时配置文件
    configFile := createTempConfigFile(t)
    defer os.Remove(configFile)
    
    // 创建配置管理器
    config := &gateway.Gateway{}
    manager, err := goconfig.CreateAndStartIntegratedManager(
        config, configFile, goconfig.EnvTest,
    )
    assert.NoError(t, err)
    defer manager.Stop()
    
    // 设置回调验证
    callbackCalled := false
    manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
        callbackCalled = true
        return nil
    }, goconfig.CallbackOptions{
        ID:    "test_callback",
        Types: []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
    })
    
    // 修改配置文件
    updateConfigFile(t, configFile)
    
    // 等待回调触发
    time.Sleep(100 * time.Millisecond)
    assert.True(t, callbackCalled)
}
```

### 5. 生产环境部署

```go
func setupProduction(manager *goconfig.IntegratedConfigManager) {
    // 设置生产环境特定的回调
    manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
        // 生产环境配置变更需要更严格的验证
        if err := validateProductionConfig(event.NewValue); err != nil {
            return fmt.Errorf("生产环境配置验证失败: %w", err)
        }
        
        // 记录配置变更到审计日志
        auditLog.LogConfigChange(event)
        
        // 通知运维团队
        return notifyOpsTeam(event)
    }, goconfig.CallbackOptions{
        ID:       "production_handler",
        Priority: goconfig.CallbackPriorityHighest,
        Timeout:  10 * time.Second,
    })
}
```

## 总结

`go-config` 提供了完整的配置管理解决方案，包括：

1. **简单易用的API**: 快速集成和使用
2. **强大的热更新**: 自动监控文件变化，支持手动触发
3. **灵活的回调机制**: 支持优先级、异步执行、超时控制
4. **完整的环境管理**: 环境变量集成和变更监听
5. **便捷的上下文集成**: 配置和环境信息注入上下文
6. **美观的日志格式化**: 自动格式化配置变更信息
7. **生产就绪**: 错误处理、性能优化、测试支持

通过本指南，您可以充分利用 `go-config` 的所有功能，构建健壮、可维护的配置管理系统。 
 