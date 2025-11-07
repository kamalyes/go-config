# go-config 使用指南

![Go Version](https://img.shields.io/badge/Go-1.20+-blue.svg)
![License](https://img.shields.io/github/license/kamalyes/go-config)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)

## 项目概述

`go-config` 是一个功能强大的 Go 语言配置管理框架，专为现代微服务架构设计。

### 核心特性

- 🔧 **多环境支持**: dev、sit、fat、uat、prod
- 🔄 **配置热更新**: 基于 fsnotify 的实时配置监听  
- 📦 **模块化配置**: 支持 20+ 种常用服务配置
- ✅ **配置验证**: 内置数据验证机制
- 🎭 **双模式管理**: SingleConfig 和 MultiConfig
- 🛡️ **类型安全**: 强类型配置结构

### 支持的配置模块

| 分类 | 模块 | 说明 |
|------|------|------|
| **基础服务** | Server, CORS, JWT | HTTP服务、跨域、JWT认证 |
| **数据存储** | MySQL, PostgreSQL, SQLite, Redis | 关系型数据库和缓存 |
| **对象存储** | Minio, AliyunOSS, FTP | 文件存储服务 |
| **消息队列** | MQTT, Kafka | 消息中间件 |
| **监控日志** | Zap, Jaeger, Elasticsearch | 日志和链路追踪 |
| **第三方API** | 支付宝、微信支付、阿里云短信 | 第三方服务集成 |
| **服务治理** | Consul, Zero(go-zero) | 服务发现和微服务框架 |

## 安装指南

### 系统要求

- Go 1.20+
- 支持 Linux、macOS、Windows

### 安装步骤

```bash
# 1. 初始化项目
go mod init your-project-name

# 2. 安装 go-config
go get -u github.com/kamalyes/go-config

# 3. 整理依赖
go mod tidy
```

### 项目结构

```text
your-project/
├── resources/                 # 配置文件目录
│   ├── dev_config.yaml       # 开发环境配置
│   ├── sit_config.yaml       # 系统集成测试环境
│   ├── fat_config.yaml       # 功能验收测试环境
│   ├── uat_config.yaml       # 用户验收测试环境
│   └── prod_config.yaml      # 生产环境配置
├── main.go                   # 主程序
├── go.mod                    # Go modules文件
└── go.sum                    # 依赖校验文件
```

## 快速开始

### 1. 创建配置文件

创建 `resources/dev_config.yaml`:

```yaml
# 服务配置
server:
  addr: '0.0.0.0:8080'
  service-name: 'my-api'
  context-path: '/api/v1'

# MySQL 配置
mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'root'
  password: 'password'
  db-name: 'myapp_dev'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'

# Redis 配置
redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 0

# 日志配置
zap:
  level: 'info'
  format: 'console'
  director: 'logs'
  development: true
```

### 2. 基础使用示例

```go
package main

import (
    "context"
    "log"
    
    goconfig "github.com/kamalyes/go-config"
)

func main() {
    // 创建上下文
    ctx := context.Background()
    
    // 创建配置管理器
    manager, err := goconfig.NewSingleConfigManager(ctx, nil)
    if err != nil {
        log.Fatalf("创建配置管理器失败: %v", err)
    }
    
    // 获取配置
    config := manager.GetConfig()
    
    // 使用配置
    log.Printf("服务地址: %s", config.Server.Addr)
    log.Printf("数据库: %s:%s", config.MySQL.Host, config.MySQL.Port)
    log.Printf("Redis: %s", config.Redis.Addr)
    
    // 启动你的应用...
}
```

### 3. 运行程序

```bash
# 设置环境变量（可选，默认为 dev）
export APP_ENV=dev

# 运行程序
go run main.go
```

## 核心概念

### ConfigOptions

配置选项用于自定义配置加载行为：

```go
type ConfigOptions struct {
    ConfigType    string              // 配置文件类型，默认 "yaml"
    ConfigPath    string              // 配置文件路径，默认 "./resources"
    ConfigSuffix  string              // 配置文件后缀，默认 "_config"
    EnvValue      env.EnvironmentType // 环境类型
    EnvContextKey env.ContextKey      // 环境上下文Key
    UseEnvLevel   EnvLevel            // 环境级别，"os" 或 "ctx"
}
```

### SingleConfig vs MultiConfig

#### SingleConfig

适用于每种服务只有一个配置实例的场景：

```go
type SingleConfig struct {
    Server    register.Server   `yaml:"server"`
    MySQL     database.MySQL    `yaml:"mysql"`
    Redis     redis.Redis       `yaml:"redis"`
    // ... 其他单一配置
}
```

#### MultiConfig

适用于需要多个同类型服务实例的场景：

```go
type MultiConfig struct {
    Server    []register.Server   `yaml:"server"`
    MySQL     []database.MySQL    `yaml:"mysql"`
    Redis     []redis.Redis       `yaml:"redis"`
    // ... 其他数组配置
}
```

### 环境类型

```go
const (
    Dev        EnvironmentType = "dev"   // 开发环境
    Sit        EnvironmentType = "sit"   // 系统集成测试
    Fat        EnvironmentType = "fat"   // 功能验收测试
    Uat        EnvironmentType = "uat"   // 用户验收测试
    Prod       EnvironmentType = "prod"  // 生产环境
)
```

## 配置模块详解

### 1. 服务器配置 (Server)

```go
type Server struct {
    Addr                    string `yaml:"addr" validate:"required"`
    ServiceName            string `yaml:"service-name" validate:"required"`
    ContextPath            string `yaml:"context-path"`
    HandleMethodNotAllowed bool   `yaml:"handle-method-not-allowed"`
    DataDriver             string `yaml:"data-driver"`
    ModuleName             string `yaml:"modulename"`
}
```

**YAML 配置示例:**

```yaml
server:
  addr: '0.0.0.0:8080'
  service-name: 'user-service'
  context-path: '/api/v1'
  handle-method-not-allowed: true
  data-driver: 'mysql'
```

### 2. MySQL 数据库配置

```go
type MySQL struct {
    Host            string `yaml:"host" validate:"required"`
    Port            string `yaml:"port" validate:"required"`
    Username        string `yaml:"username" validate:"required"`
    Password        string `yaml:"password" validate:"required"`
    Dbname          string `yaml:"db-name" validate:"required"`
    Config          string `yaml:"config" validate:"required"`
    LogLevel        string `yaml:"log-level" validate:"required"`
    MaxIdleConns    int    `yaml:"max-idle-conns" validate:"min=0"`
    MaxOpenConns    int    `yaml:"max-open-conns" validate:"min=0"`
    ConnMaxIdleTime int    `yaml:"conn-max-idle-time" validate:"min=0"`
    ConnMaxLifeTime int    `yaml:"conn-max-life-time" validate:"min=0"`
    ModuleName      string `yaml:"modulename"`
}
```

**YAML 配置示例:**

```yaml
mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'root'
  password: 'password'
  db-name: 'myapp'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'
  max-idle-conns: 10
  max-open-conns: 100
  conn-max-idle-time: 30
  conn-max-life-time: 300
```

### 3. Redis 配置

```go
type Redis struct {
    Addr         string `yaml:"addr" validate:"required"`
    Password     string `yaml:"password"`
    DB           int    `yaml:"db" validate:"min=0,max=15"`
    PoolSize     int    `yaml:"pool-size" validate:"min=1"`
    MinIdleConns int    `yaml:"min-idle-conns" validate:"min=0"`
    ModuleName   string `yaml:"modulename"`
}
```

**YAML 配置示例:**

```yaml
redis:
  addr: '127.0.0.1:6379'
  password: 'redis_password'
  db: 0
  pool-size: 100
  min-idle-conns: 5
```

### 4. 日志配置 (Zap)

```go
type Zap struct {
    Level         string `yaml:"level" validate:"required"`
    Format        string `yaml:"format" validate:"required"`
    Prefix        string `yaml:"prefix"`
    Director      string `yaml:"director" validate:"required"`
    LinkName      string `yaml:"link-name"`
    ShowLine      bool   `yaml:"show-line"`
    EncodeLevel   string `yaml:"encode-level"`
    LogInConsole  bool   `yaml:"log-in-console"`
    Development   bool   `yaml:"development"`
    ModuleName    string `yaml:"modulename"`
}
```

**YAML 配置示例:**

```yaml
zap:
  level: 'info'                              # debug、info、warn、error
  format: 'console'                          # json、console
  prefix: '[MyApp]'
  director: 'logs'
  link-name: 'logs/app.log'
  show-line: true
  encode-level: 'LowercaseColorLevelEncoder'
  log-in-console: true
  development: true
```

### 5. CORS 配置

```yaml
cors:
  allowed-all-origins: false
  allowed-origins:
    - "http://localhost:3000"
    - "https://myapp.com"
  allowed-methods:
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
  allowed-headers:
    - "Authorization"
    - "Content-Type"
  allow-credentials: true
  max-age: "86400"
```

## 环境管理

### 环境变量设置

#### 1. 通过操作系统环境变量

```bash
# Linux/macOS
export APP_ENV=prod

# Windows
set APP_ENV=prod
```

#### 2. 通过代码设置

```go
import "github.com/kamalyes/go-config/pkg/env"

// 设置环境
env.SetContextKey(&env.ContextKeyOptions{
    Key:   env.ContextKey("MY_APP_ENV"),
    Value: env.Prod,
})
```

#### 3. 通过配置选项

```go
options := &goconfig.ConfigOptions{
    EnvValue: env.Prod,
    UseEnvLevel: goconfig.EnvLevelOS, // 或 EnvLevelCtx
}
```

### 环境优先级

1. **EnvLevelOS**: 优先使用操作系统环境变量
2. **EnvLevelCtx**: 优先使用代码中设置的环境变量

## 高级用法

### 1. 自定义配置路径和选项

```go
package main

import (
    "context"
    "log"
    
    goconfig "github.com/kamalyes/go-config"
    "github.com/kamalyes/go-config/pkg/env"
)

func advancedUsage() {
    ctx := context.Background()
    
    // 自定义配置选项
    options := &goconfig.ConfigOptions{
        ConfigType:    "yaml",                    // 配置文件类型
        ConfigPath:    "./custom_configs",       // 自定义配置路径
        ConfigSuffix:  "_settings",              // 自定义后缀
        EnvValue:      env.Prod,                 // 指定环境
        EnvContextKey: env.ContextKey("CUSTOM_ENV"),
        UseEnvLevel:   goconfig.EnvLevelCtx,     // 使用代码设置的环境
    }
    
    manager, err := goconfig.NewSingleConfigManager(ctx, options)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    
    config := manager.GetConfig()
    log.Printf("Custom config loaded: %+v", config.Server)
}
```

### 2. MultiConfig 使用示例

当您需要多个同类型的服务实例时：

```yaml
# 多 MySQL 实例配置
mysql:
  - modulename: "primary"
    host: '192.168.1.10'
    port: '3306'
    username: 'root'
    password: 'password1'
    db-name: 'primary_db'
    config: 'charset=utf8mb4&parseTime=True&loc=Local'
    log-level: 'info'
  
  - modulename: "secondary"
    host: '192.168.1.11'
    port: '3306'
    username: 'root'
    password: 'password2'
    db-name: 'secondary_db'
    config: 'charset=utf8mb4&parseTime=True&loc=Local'
    log-level: 'error'

# 多 Redis 实例配置
redis:
  - modulename: "cache"
    addr: '192.168.1.20:6379'
    password: 'cache_password'
    db: 0
  
  - modulename: "session"
    addr: '192.168.1.21:6379'
    password: 'session_password'
    db: 1
```

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    goconfig "github.com/kamalyes/go-config"
)

func multiConfigUsage() {
    ctx := context.Background()
    
    // 创建多配置管理器
    manager, err := goconfig.NewMultiConfigManager(ctx, nil)
    if err != nil {
        log.Fatalf("Error creating multi config manager: %v", err)
    }
    
    multiConfig := manager.GetConfig()
    
    // 获取指定模块的配置
    primaryMySQL, err := goconfig.GetModuleByName(multiConfig.MySQL, "primary")
    if err != nil {
        log.Printf("Error getting primary MySQL config: %v", err)
        return
    }
    
    cacheRedis, err := goconfig.GetModuleByName(multiConfig.Redis, "cache")
    if err != nil {
        log.Printf("Error getting cache Redis config: %v", err)
        return
    }
    
    fmt.Printf("Primary MySQL: %s:%s/%s\n", 
        primaryMySQL.Host, primaryMySQL.Port, primaryMySQL.Dbname)
    fmt.Printf("Cache Redis: %s (DB: %d)\n", 
        cacheRedis.Addr, cacheRedis.DB)
}
```

### 3. 配置验证

```go
func validateConfig() {
    ctx := context.Background()
    manager, _ := goconfig.NewSingleConfigManager(ctx, nil)
    config := manager.GetConfig()
    
    // 验证 MySQL 配置
    if err := config.MySQL.Validate(); err != nil {
        log.Fatalf("MySQL config validation failed: %v", err)
    }
    
    // 验证 Redis 配置
    if err := config.Redis.Validate(); err != nil {
        log.Fatalf("Redis config validation failed: %v", err)
    }
    
    log.Println("All configurations are valid!")
}
```

## 配置文件示例

### 开发环境配置 (dev_config.yaml)

```yaml
# 服务配置
server:
  addr: '0.0.0.0:8080'
  service-name: 'myapp-dev'
  context-path: '/api/v1'
  handle-method-not-allowed: true
  data-driver: 'mysql'

# MySQL 配置
mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'root'
  password: 'dev_password'
  db-name: 'myapp_dev'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'debug'
  max-idle-conns: 5
  max-open-conns: 50

# Redis 配置
redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 0
  pool-size: 50

# 日志配置
zap:
  level: 'debug'
  format: 'console'
  prefix: '[DEV]'
  director: 'logs'
  show-line: true
  development: true
  log-in-console: true

# CORS 配置 (开发环境允许所有来源)
cors:
  allowed-all-origins: true
  allow-credentials: true
```

### 生产环境配置 (prod_config.yaml)

```yaml
# 服务配置
server:
  addr: '0.0.0.0:80'
  service-name: 'myapp-prod'
  context-path: '/api/v1'
  handle-method-not-allowed: false
  data-driver: 'mysql'

# MySQL 配置
mysql:
  host: '${DB_HOST}'
  port: '${DB_PORT:3306}'
  username: '${DB_USER}'
  password: '${DB_PASSWORD}'
  db-name: '${DB_NAME}'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'error'
  max-idle-conns: 20
  max-open-conns: 200
  conn-max-idle-time: 300
  conn-max-life-time: 3600

# Redis 配置
redis:
  addr: '${REDIS_HOST}:${REDIS_PORT:6379}'
  password: '${REDIS_PASSWORD}'
  db: 0
  pool-size: 200
  min-idle-conns: 10

# 日志配置
zap:
  level: 'info'
  format: 'json'
  prefix: '[PROD]'
  director: '/var/log/myapp'
  show-line: false
  development: false
  log-in-console: false

# CORS 配置 (生产环境严格控制)
cors:
  allowed-all-origins: false
  allowed-origins:
    - "https://myapp.com"
    - "https://www.myapp.com"
  allowed-methods: ["GET", "POST", "PUT", "DELETE"]
  allowed-headers: ["Authorization", "Content-Type"]
  allow-credentials: true
  max-age: "86400"
```

## API 参考

### 核心函数

#### NewSingleConfigManager

```go
func NewSingleConfigManager(ctx context.Context, options *ConfigOptions) (*SingleConfigManager, error)
```

创建单一配置管理器。

**参数:**

- `ctx`: 上下文对象
- `options`: 配置选项，可为 nil 使用默认值

**返回:**

- `*SingleConfigManager`: 配置管理器实例
- `error`: 错误信息

#### NewMultiConfigManager

```go
func NewMultiConfigManager(ctx context.Context, options *ConfigOptions) (*MultiConfigManager, error)
```

创建多配置管理器。

#### GetModuleByName

```go
func GetModuleByName[T any](modules []T, moduleName string) (T, error)
```

从配置数组中根据模块名获取特定配置。

**参数:**

- `modules`: 配置模块数组
- `moduleName`: 模块名称

### 配置管理器方法

#### GetConfig

```go
func (m *SingleConfigManager) GetConfig() *SingleConfig
func (m *MultiConfigManager) GetConfig() *MultiConfig
```

获取配置对象。

#### SubItem

```go
func (m *MultiConfigManager) SubItem(ctx context.Context, subKey string, v interface{})
```

获取配置子项。

## 最佳实践

### 1. 配置文件组织

```text
resources/
├── common/                    # 公共配置
│   ├── database.yaml         # 数据库配置模板
│   └── logging.yaml          # 日志配置模板
├── dev_config.yaml           # 开发环境
├── sit_config.yaml           # 系统集成测试
├── uat_config.yaml           # 用户验收测试  
├── prod_config.yaml          # 生产环境
└── local_config.yaml         # 本地开发（git ignore）
```

### 2. 敏感信息处理

#### 方法1: 环境变量

```yaml
mysql:
  host: '${DB_HOST:127.0.0.1}'
  username: '${DB_USER:root}'
  password: '${DB_PASSWORD}'
```

#### 方法2: 外部配置文件

```go
// 加载外部敏感配置
func loadSecrets() {
    if secretFile := os.Getenv("SECRET_FILE"); secretFile != "" {
        // 加载外部密钥文件
    }
}
```

### 3. 配置验证策略

```go
func validateBusinessConfig(config *goconfig.SingleConfig) error {
    // 业务逻辑验证
    if config.MySQL.MaxOpenConns < config.MySQL.MaxIdleConns {
        return errors.New("max_open_conns should be >= max_idle_conns")
    }
    
    if config.Server.Addr == "" {
        return errors.New("server address is required")
    }
    
    return nil
}
```

### 4. 错误处理模式

```go
func initConfig() (*goconfig.SingleConfig, error) {
    ctx := context.Background()
    
    manager, err := goconfig.NewSingleConfigManager(ctx, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create config manager: %w", err)
    }
    
    config := manager.GetConfig()
    
    // 验证关键配置
    if err := validateCriticalConfig(config); err != nil {
        return nil, fmt.Errorf("config validation failed: %w", err)
    }
    
    return config, nil
}
```

## 故障排除

### 常见问题

#### 1. 配置文件未找到

**错误信息:**

```text
读取配置文件异常: Config File "dev_config" Not Found in "[./resources]"
```

**解决方法:**

- 检查配置文件路径是否正确
- 确认文件名格式: `{environment}_config.yaml`
- 检查文件权限

```go
// 自定义配置路径
options := &goconfig.ConfigOptions{
    ConfigPath: "/path/to/your/configs",
}
```

#### 2. 环境变量未设置

**错误信息:**

```text
环境变量 APP_ENV 未设置，使用默认环境: dev
```

**解决方法:**

```bash
# 设置环境变量
export APP_ENV=prod

# 或在代码中设置
env.SetContextKey(&env.ContextKeyOptions{
    Value: env.Prod,
})
```

#### 3. 配置验证失败

**错误信息:**

```text
MySQL config validation failed: Host is required
```

**解决方法:**

- 检查必填字段是否设置
- 验证数据格式是否正确
- 查看具体的验证规则

### 调试技巧

#### 1. 启用详细日志

```go
import "log"

// 在程序开始时启用详细日志
log.SetFlags(log.LstdFlags | log.Lshortfile)
```

#### 2. 配置内容检查

```go
func debugConfig(config *goconfig.SingleConfig) {
    log.Printf("Server Config: %+v", config.Server)
    log.Printf("MySQL Config: %+v", config.MySQL)
    log.Printf("Redis Config: %+v", config.Redis)
}
```

#### 3. 环境检查

```go
import "github.com/kamalyes/go-config/pkg/env"

func debugEnvironment() {
    log.Printf("Current Environment: %s", env.GetEnvironment())
    log.Printf("Context Key: %s", env.GetContextKey())
}
```

## 更多资源

- [项目主页](https://github.com/kamalyes/go-config)
- [API 文档](https://pkg.go.dev/github.com/kamalyes/go-config)
- [问题反馈](https://github.com/kamalyes/go-config/issues)

## 许可证

本项目采用 [MIT 许可证](LICENSE)。

---

**最后更新:** 2024年11月7日  
**版本:** v1.0.0  
**作者:** kamalyes