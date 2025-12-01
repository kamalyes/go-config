# 环境管理器 - 使用说明

## 概述

环境管理器提供了便捷的环境类型判断和自动初始化功能，无需手动初始化即可使用。

## 自动初始化

包在导入时会自动初始化全局环境实例，无需手动调用任何初始化函数：

```go
import goconfig "path/to/your/go-config"

// 无需手动初始化，直接使用
func main() {
    if goconfig.IsDev() {
        log.Println("当前是开发环境")
    }
}
```

## 环境类型判断函数

### 基础判断函数

```go
// 判断是否为特定环境
if goconfig.IsDev() {
    // 开发环境逻辑
}

if goconfig.IsProduction() {
    // 生产环境逻辑
}

if goconfig.IsTest() {
    // 测试环境逻辑
}

if goconfig.IsStaging() {
    // 预发布环境逻辑
}

if goconfig.IsUAT() {
    // 用户验收测试环境逻辑
}

if goconfig.IsLocal() {
    // 本地环境逻辑
}

if goconfig.IsDebug() {
    // 调试环境逻辑
}

if goconfig.IsDemo() {
    // 演示环境逻辑
}

if goconfig.IsIntegration() {
    // 集成环境逻辑
}
```

### 通用判断函数

```go
// 判断是否为指定环境
if goconfig.IsEnvironment(goconfig.EnvDevelopment) {
    // 开发环境逻辑
}

// 判断是否为多个环境中的任意一个
if goconfig.IsAnyOf(goconfig.EnvDevelopment, goconfig.EnvTest, goconfig.EnvLocal) {
    // 开发相关环境逻辑
}
```

### 环境级别判断

环境级别用于按重要程度分类环境：

```go
// 获取环境级别（数字越小级别越高）
level := goconfig.GetCurrentEnvironmentLevel()

// 按级别判断
if goconfig.IsProductionLevel() {
    // 生产级别环境（级别 >= 10）
    log.SetLevel(log.WarnLevel)
} else if goconfig.IsTestingLevel() {
    // 测试级别环境（3 <= 级别 < 10）
    log.SetLevel(log.InfoLevel)
} else if goconfig.IsDevelopmentLevel() {
    // 开发级别环境（级别 <= 2）
    log.SetLevel(log.DebugLevel)
}
```

## 环境级别定义

| 环境类型 | 级别 | 说明 |
|---------|------|------|
| Local, Debug | 1 | 本地开发环境 |
| Development | 2 | 开发环境 |
| Test | 3 | 测试环境 |
| Integration | 4 | 集成环境 |
| Staging, UAT | 5 | 预发布/验收环境 |
| Demo | 6 | 演示环境 |
| Production | 10 | 生产环境 |

## 便捷函数

```go
// 获取当前环境
env := goconfig.GetCurrentEnvironment()
fmt.Printf("当前环境: %s\n", env)

// 设置当前环境
goconfig.SetCurrentEnvironment(goconfig.EnvProduction)

// 获取全局环境实例
envInstance := goconfig.GetGlobalEnvironment()
```

## 实际使用示例

### 数据库配置

```go
func setupDatabase() {
    if goconfig.IsProduction() {
        // 生产环境：使用高可用配置
        db.SetMaxOpenConns(100)
        db.SetConnMaxLifetime(time.Hour)
    } else if goconfig.IsTestingLevel() {
        // 测试环境：使用测试数据库
        db.SetMaxOpenConns(10)
        connectToTestDB()
    } else {
        // 开发环境：使用本地数据库
        db.SetMaxOpenConns(5)
        connectToLocalDB()
    }
}
```

### 日志配置

```go
func setupLogging() {
    if goconfig.IsProductionLevel() {
        log.SetLevel(log.WarnLevel)
        log.SetFormatter(&log.JSONFormatter{})
    } else if goconfig.IsDevelopmentLevel() {
        log.SetLevel(log.DebugLevel)
        log.SetFormatter(&log.TextFormatter{})
    } else {
        log.SetLevel(log.InfoLevel)
    }
}
```

### 功能开关

```go
func configureFeatures() {
    // 只在开发环境启用调试功能
    if goconfig.IsDevelopmentLevel() {
        pprof.EnableProfiling()
        gin.SetMode(gin.DebugMode)
    }

    // 只在生产环境启用性能监控
    if goconfig.IsProductionLevel() {
        prometheus.EnableMetrics()
        sentry.InitErrorTracking()
    }

    // 在测试相关环境启用测试工具
    if goconfig.IsAnyOf(goconfig.EnvTest, goconfig.EnvStaging, goconfig.EnvUAT) {
        testutils.EnableTestEndpoints()
    }
}
```

### HTTP服务器配置

```go
func configureServer() *http.Server {
    server := &http.Server{
        Addr: ":8080",
    }

    if goconfig.IsProductionLevel() {
        // 生产环境：启用TLS，设置超时
        server.TLSConfig = &tls.Config{}
        server.ReadTimeout = 30 * time.Second
        server.WriteTimeout = 30 * time.Second
    } else if goconfig.IsDevelopmentLevel() {
        // 开发环境：启用热重载
        server.Addr = ":8081" // 避免端口冲突
        gin.SetMode(gin.DebugMode)
    }

    return server
}
```

## 环境变量设置

设置 `APP_ENV` 环境变量来指定运行环境：

```bash
# 开发环境
export APP_ENV=dev

# 测试环境  
export APP_ENV=test

# 预发布环境
export APP_ENV=staging

# 生产环境
export APP_ENV=production
```

支持的环境别名：

- Development: `dev`, `develop`, `development`
- Test: `test`, `testing`, `qa`, `sit`
- Staging: `staging`, `stage`, `stg`, `pre`, `preprod`, `pre-prod`, `fat`, `gray`, `grey`, `canary`
- Production: `prod`, `production`, `prd`, `release`, `live`, `online`, `master`, `main`
- Local: `local`, `localhost`
- Debug: `debug`, `debugging`, `dbg`
- Demo: `demo`, `demonstration`, `showcase`, `preview`, `sandbox`
- UAT: `uat`, `acceptance`, `user-acceptance`, `beta`
- Integration: `integration`, `int`, `ci`, `integration-test`, `integ`

## 高级功能

### 环境变更回调

```go
func init() {
    env := goconfig.GetGlobalEnvironment()
    env.RegisterCallback("db-config", func(oldEnv, newEnv goconfig.EnvironmentType) error {
        log.Printf("环境从 %s 变更为 %s，重新配置数据库", oldEnv, newEnv)
        return setupDatabase()
    }, 1, false)
}
```

### 自定义环境类型

```go
func init() {
    // 注册自定义环境类型
    goconfig.RegisterEnvPrefixes("custom", "custom", "my-env", "myenv")
}
```

## 注意事项

1. 包会自动初始化，无需手动调用初始化函数
2. 环境检测基于 `APP_ENV` 环境变量
3. 如果未设置环境变量，默认使用 `development` 环境
4. 环境级别判断有助于按重要程度分类处理逻辑
5. 支持环境变更监控和回调机制
