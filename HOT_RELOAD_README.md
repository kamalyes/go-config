# 配置热更新和环境回调功能

本文档介绍 go-config 项目中新增的配置热更新和环境变更回调功能。

## 🚀 新功能概览

### 1. 配置热更新 (Hot Reload)

- 🔄 实时监控配置文件变化
- ⚡ 自动重新加载配置
- 🔔 配置变更回调通知
- 🛡️ 防抖机制避免频繁更新
- 🔧 可配置的监控参数

### 2. 环境变更监听

- 🌍 环境变量变化监控
- 📢 环境切换回调通知
- 🔗 与配置管理器集成
- ⚙️ 可自定义监控频率

### 3. 上下文集成

- 📦 将配置和环境信息注入上下文
- 🎯 便捷的上下文辅助工具
- 📊 配置元数据管理
- 🔄 自动同步配置更新

### 4. 集成管理器

- 🎛️ 统一的配置管理入口
- 🤝 环境与配置的协调管理
- 📋 简化的 API 接口
- 🎪 开箱即用的解决方案

## 📚 使用指南

### 基本使用

```go
package main

import (
    "context"
    "log"
    
    goconfig "github.com/kamalyes/go-config"
)

// 定义配置结构
type AppConfig struct {
    Server struct {
        Host string `yaml:"host"`
        Port int    `yaml:"port"`
    } `yaml:"server"`
    Database struct {
        URL string `yaml:"url"`
    } `yaml:"database"`
}

func main() {
    // 创建配置实例
    config := &AppConfig{}
    
    // 创建并启动集成配置管理器
    manager, err := goconfig.CreateAndStartIntegratedManager(
        config,
        "config.yaml",
        goconfig.EnvDevelopment,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer manager.Stop()
    
    // 使用配置
    currentConfig := manager.GetConfig().(*AppConfig)
    log.Printf("服务器地址: %s:%d", currentConfig.Server.Host, currentConfig.Server.Port)
}
```

### 注册回调监听

```go
// 注册配置变更回调
err := manager.RegisterConfigCallback(func(ctx context.Context, event goconfig.CallbackEvent) error {
    log.Printf("配置已更新: %s", event.Source)
    
    if newConfig, ok := event.NewValue.(*AppConfig); ok {
        log.Printf("新的服务器地址: %s:%d", newConfig.Server.Host, newConfig.Server.Port)
        
        // 在这里执行配置更新后的逻辑
        // 例如：重启服务器、重新连接数据库等
    }
    
    return nil
}, goconfig.CallbackOptions{
    ID:       "my_config_callback",
    Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged},
    Priority: 1,
    Async:    false,
    Timeout:  5 * time.Second,
})

// 注册环境变更回调
err = manager.RegisterEnvironmentCallback("my_env_callback", 
    func(oldEnv, newEnv goconfig.EnvironmentType) error {
        log.Printf("环境已切换: %s -> %s", oldEnv, newEnv)
        
        // 在这里执行环境切换后的逻辑
        // 例如：切换数据库连接、调整日志级别等
        
        return nil
    }, 1, false)
```

### 使用上下文功能

```go
// 创建带配置的上下文
ctx := manager.WithContext(context.Background())

// 从上下文获取环境信息
if env, ok := goconfig.GetEnvironmentFromContext(ctx); ok {
    log.Printf("当前环境: %s", env)
}

// 从上下文获取配置
if config, ok := goconfig.GetConfigFromContext(ctx); ok {
    if appConfig, ok := config.(*AppConfig); ok {
        log.Printf("数据库URL: %s", appConfig.Database.URL)
    }
}

// 使用上下文辅助工具
isDev := goconfig.ContextHelper.IsEnvironment(ctx, goconfig.EnvDevelopment)
if isDev {
    log.Println("当前运行在开发环境")
}

// 创建带超时的配置上下文
timeoutCtx, cancel := goconfig.ContextHelper.NewContextWithTimeout(30 * time.Second)
defer cancel()
```

### 高级配置

```go
// 自定义热更新配置
hotReloadConfig := &goconfig.HotReloadConfig{
    Enabled:         true,
    WatchInterval:   500 * time.Millisecond, // 监控间隔
    DebounceDelay:   1 * time.Second,        // 防抖延迟
    MaxRetries:      3,                      // 最大重试次数
    CallbackTimeout: 30 * time.Second,      // 回调超时
    EnableEnvWatch:  true,                  // 启用环境监控
}

// 创建自定义选项
options := &goconfig.IntegratedConfigOptions{
    ConfigPath:      "config/app.yaml",
    Environment:     goconfig.EnvProduction,
    HotReloadConfig: hotReloadConfig,
    ContextOptions: &goconfig.ContextKeyOptions{
        Key:   goconfig.ContextKey("CUSTOM_ENV"),
        Value: goconfig.EnvProduction,
    },
}

// 使用自定义选项创建管理器
manager, err := goconfig.NewIntegratedConfigManager(config, options)
```

## 🎯 回调类型

| 回调类型 | 说明 | 触发时机 |
|---------|-----|--------|
| `CallbackTypeConfigChanged` | 配置文件变更 | 配置文件被修改时 |
| `CallbackTypeEnvChanged` | 环境变量变更 | 环境变量被修改时 |
| `CallbackTypeReloaded` | 重新加载完成 | 手动或自动重载完成时 |
| `CallbackTypeError` | 错误回调 | 发生错误时 |

## ⚙️ 配置选项

### 热更新配置 (HotReloadConfig)

```go
type HotReloadConfig struct {
    Enabled         bool          // 是否启用热更新 (默认: true)
    WatchInterval   time.Duration // 监控间隔 (默认: 500ms)
    DebounceDelay   time.Duration // 防抖延迟 (默认: 1s)
    MaxRetries      int           // 最大重试次数 (默认: 3)
    CallbackTimeout time.Duration // 回调超时 (默认: 30s)
    EnableEnvWatch  bool          // 是否监控环境变量 (默认: true)
}
```

### 回调选项 (CallbackOptions)

```go
type CallbackOptions struct {
    ID          string            // 回调唯一标识 (必需)
    Types       []CallbackType    // 监听的事件类型 (必需)
    Priority    int               // 优先级，数字越小优先级越高 (默认: 0)
    Async       bool              // 是否异步执行 (默认: false)
    Timeout     time.Duration     // 超时时间 (默认: 30s)
    Retry       int               // 重试次数 (默认: 0)
    Metadata    map[string]interface{} // 附加元数据
}
```

## 🛠️ API 参考

### IntegratedConfigManager

| 方法 | 说明 |
|-----|-----|
| `Start(ctx)` | 启动配置管理器 |
| `Stop()` | 停止配置管理器 |
| `GetConfig()` | 获取当前配置 |
| `GetEnvironment()` | 获取当前环境 |
| `ReloadConfig(ctx)` | 手动重新加载配置 |
| `RegisterConfigCallback(callback, options)` | 注册配置回调 |
| `RegisterEnvironmentCallback(id, callback, priority, async)` | 注册环境回调 |
| `WithContext(ctx)` | 创建带配置的上下文 |

### ContextHelper

| 方法 | 说明 |
|-----|-----|
| `NewConfigContext()` | 创建新的配置上下文 |
| `NewContextWithTimeout(timeout)` | 创建带超时的配置上下文 |
| `IsEnvironment(ctx, env)` | 检查上下文中的环境 |
| `MustGetConfig(ctx)` | 从上下文获取配置（失败时panic） |
| `MustGetEnvironment(ctx)` | 从上下文获取环境（失败时panic） |

## 🎬 完整示例

查看 `examples/gateway_hot_reload_demo.go` 文件获取完整的使用示例。

运行示例：

```bash
# 进入示例目录
cd examples

# 运行示例
go run gateway_hot_reload_demo.go

# 修改生成的 example_config.yaml 文件来测试热更新
# 修改 APP_ENV 环境变量来测试环境变更
```

## 🧪 测试

运行测试：

```bash
# 运行所有测试
go test ./tests/...

# 运行热更新相关测试
go test ./tests/hot_reload_test.go -v

# 运行基准测试
go test ./tests/hot_reload_test.go -bench=.
```

## 📈 性能考虑

1. **防抖机制**: 避免频繁的配置文件变更导致过多的重载
2. **异步回调**: 支持异步执行回调，避免阻塞主流程
3. **优先级控制**: 回调按优先级执行，确保重要操作优先处理
4. **超时控制**: 防止回调执行时间过长
5. **错误重试**: 支持回调执行失败时的重试机制

## 🚨 注意事项

1. **配置文件权限**: 确保应用有读取配置文件的权限
2. **环境变量**: 环境变量的变更监控依赖于定期检查机制
3. **回调异常**: 回调函数中的异常会被捕获并记录，不会影响主流程
4. **内存占用**: 长时间运行的应用需要注意回调注册的数量
5. **并发安全**: 所有 API 都是并发安全的

## 🔗 相关文档

- [基础配置使用文档](README.md)

---

📝 **注意**: 这是 go-config 项目的扩展功能，确保你的项目已经集成了基础的 go-config 功能。
