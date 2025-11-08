# go-config 链式配置方法

本文档介绍了如何使用 go-config 包中新增的 Default 方法和链式 Withxxx 方法。

> **🎉 完成状态**: 已为 **34个包** 实现统一的Default方法和链式配置支持，覆盖率 **100%**，所有测试通过！

## 功能概述

每个 pkg 包现在都支持：
1. `DefaultXXX()` 函数 - 返回该类型的默认配置（值类型）
2. `Default()` 函数 - 返回默认配置的指针，支持链式调用（指针类型）
3. `WithXXX()` 方法 - 支持链式调用的配置方法

## 统一的Default方法格式

所有包都遵循统一的Default方法返回体格式：

### 返回体格式说明

1. **`Default()`** - 返回指针类型 `*struct`，支持链式调用
   - 用于最常见的使用场景：需要链式配置的情况
   - 示例：`cache.Default().WithType("redis").WithEnabled(true)`

2. **`DefaultXXX()`** - 返回值类型 `struct`，提供纯净的默认配置
   - 用于需要多个独立实例或不可变配置的场景
   - 示例：`baseConfig := cache.DefaultCache()`

### 设计优势

这种设计模式的优势：
- **最常用场景优先**：`Default()`直接支持链式调用
- **性能优化**：避免不必要的指针分配
- **类型安全**：编译时检查，避免空指针错误
- **API一致性**：所有包使用相同的命名约定

## 快速入门

```go
import (
    "github.com/kamalyes/go-config/pkg/cache"
    "github.com/kamalyes/go-config/pkg/database"
    "github.com/kamalyes/go-config/pkg/jwt"
)

// 简单使用默认配置
cacheConfig := cache.Default()

// 链式配置定制
customCache := cache.Default().
    WithType(cache.TypeRedis).
    WithEnabled(true).
    WithKeyPrefix("myapp:")

// 组合多个配置
appConfig := struct {
    Cache    *cache.Cache
    Database *database.MySQL  
    JWT      *jwt.JWT
}{
    Cache:    cache.Default().WithType(cache.TypeRedis),
    Database: database.Default().WithHost("localhost").WithDbname("myapp"),
    JWT:      jwt.Default().WithSigningKey("my-secret"),
}
```

## 使用方式

### 1. 基本用法

```go
// 使用默认配置指针（支持链式调用）
config := cache.Default()

// 获取默认配置值（用于复制或不可变场景）
defaultConfig := cache.DefaultCache()

// 两种方式的区别：
// 1. 指针类型 - 支持链式调用
chainConfig := cache.Default().
    WithType(cache.TypeRedis).
    WithEnabled(true)

// 2. 值类型 - 用于获取纯净默认值
baseConfig := cache.DefaultCache()
customConfig1 := baseConfig  // 独立副本1
customConfig2 := baseConfig  // 独立副本2
```

### 2. 链式配置

```go
// 缓存配置
cacheConfig := cache.Default().
    WithModuleName("my-cache").
    WithType(cache.TypeRedis).
    WithEnabled(true).
    WithDefaultTTL(30 * time.Minute).
    WithKeyPrefix("myapp:")

// MySQL 配置
mysqlConfig := database.Default().
    WithHost("127.0.0.1").
    WithPort("3306").
    WithDbname("myapp").
    WithUsername("root").
    WithPassword("password").
    WithMaxIdleConns(20)

// JWT 配置
jwtConfig := jwt.Default().
    WithSigningKey("my-secret-key").
    WithExpiresTime(3600 * 24).
    WithUseMultipoint(true)
```

## 支持的包

以下包已经支持统一的Default方法和链式Withxxx方法：

### 核心配置
- ✅ `cache` - 缓存配置
- ✅ `captcha` - 验证码配置
- ✅ `env` - 环境变量配置

### 数据库
- ✅ `database/mysql` - MySQL 数据库配置
- ✅ `database/postgre` - PostgreSQL 数据库配置
- ✅ `database/sqlite` - SQLite 数据库配置

### 认证与安全
- ✅ `jwt` - JWT 令牌配置
- ✅ `cors` - 跨域资源共享配置

### 通信与通知
- ✅ `email` - 邮件配置
- ✅ `ftp` - FTP 配置
- ✅ `sms/aliyun` - 阿里云短信配置
- ✅ `queue/mqtt` - MQTT 消息队列配置

### 服务注册与发现
- ✅ `register/server` - 服务器配置
- ✅ `register/consul` - Consul 注册中心配置
- ✅ `register/jaeger` - Jaeger 链路追踪配置
- ✅ `register/pprof` - PProf 性能分析配置

### 日志
- ✅ `zap` - Zap 日志配置

### 支付
- ✅ `pay/alipay` - 支付宝配置
- ✅ `pay/wechat` - 微信支付配置

### 对象存储
- ✅ `oss/s3` - AWS S3 对象存储配置
- ✅ `oss/aliyun` - 阿里云OSS配置
- ✅ `oss/minio` - MinIO 对象存储配置

### ELK 技术栈
- ✅ `elk/es` - Elasticsearch 配置
- ✅ `elk/kafka` - Kafka 配置

### 云服务
- ✅ `sts/aliyun` - 阿里云STS配置
- ✅ `youzan` - 有赞API配置

### Zero 微服务框架
- ✅ `zero/client` - RPC 客户端配置
- ✅ `zero/server` - RPC 服务端配置
- ✅ `zero/etcd` - Etcd 配置
- ✅ `zero/logx` - 日志配置
- ✅ `zero/prometheus` - Prometheus 配置
- ✅ `zero/trace` - 链路追踪配置
- ✅ `zero/restful` - RESTful 服务配置
- ✅ `zero/signature` - 签名配置

## 默认值说明

每个包的默认配置都提供了合理的默认值：

### Cache (缓存)
```go
ModuleName: "default"
Type: TypeMemory
Enabled: true
DefaultTTL: 30 * time.Minute
KeyPrefix: "cache:"
Serializer: "json"
```

### MySQL (数据库)
```go
ModuleName: "mysql"
Host: "127.0.0.1"
Port: "3306"
Config: "charset=utf8mb4&parseTime=True&loc=Local"
LogLevel: "silent"
MaxIdleConns: 10
MaxOpenConns: 100
ConnMaxIdleTime: 3600  // 1小时
ConnMaxLifeTime: 7200  // 2小时
```

### JWT (认证)
```go
ModuleName: "jwt"
SigningKey: "go-config-default-key"
ExpiresTime: 3600 * 24 * 7  // 7天
BufferTime: 3600             // 1小时
UseMultipoint: false
```

### Redis (缓存)
```go
ModuleName: "redis"
Addr: "127.0.0.1:6379"
DB: 0
MaxRetries: 3
PoolSize: 10
MaxConnAge: 30 * time.Minute
PoolTimeout: 4 * time.Second
IdleTimeout: 5 * time.Minute
ReadTimeout: 3 * time.Second
WriteTimeout: 3 * time.Second
```

## 优势

1. **简洁的API** - 链式调用使配置代码更简洁易读
2. **类型安全** - 编译时检查，减少配置错误
3. **合理默认值** - 每个包都有经过验证的默认配置
4. **一致性** - 所有包使用相同的模式
5. **可读性** - 配置意图清晰明确
6. **灵活性** - 可以只配置需要修改的字段

## 完整示例

参见 `examples/chain_usage/main.go` 文件，其中包含了所有支持包的使用示例。

运行示例：
```bash
go run examples/chain_usage/main.go
```

## 测试验证

所有的Default方法都通过了完整的测试验证：

### 运行测试
```bash
# 测试所有Default方法
go test ./tests -run "TestAllDefaultMethods" -v

# 测试所有链式方法
go test ./tests -run "TestAllDefaultChainMethods" -v

# 测试链式方法使用
go test ./tests -run "TestChainMethodUsage" -v

# 运行完整测试套件
go test ./tests -v
```

### 测试覆盖

**🎉 覆盖完成状态：34/34 包已完成，覆盖率 100%**

测试包括：
- ✅ 34个包的`Default()`方法测试
- ✅ 34个包的链式调用方法测试  
- ✅ 链式方法功能验证测试
- ✅ 配置验证测试
- ✅ 不可变性测试
- ✅ 复制功能测试

**所有测试均已通过验证！**

### 特殊说明

1. **elk包处理**：由于elk包含多个配置类型，使用以下方式：
   - `elk.Default()` - 返回Elasticsearch配置指针（主要类型，支持链式调用）
   - `elk.DefaultElasticsearch()` - 返回Elasticsearch配置值（纯净默认配置）
   - `elk.DefaultElasticsearchConfig()` - 返回Elasticsearch配置指针（支持链式调用）
   - `elk.DefaultKafka()` - 返回Kafka配置值（纯净默认配置）
   - `elk.DefaultKafkaConfig()` - 返回Kafka配置指针（支持链式调用）

2. **zero包处理**：由于zero包含多个配置类型，使用以下方式：
   - `zero.DefaultRpcServer()` - 返回RPC服务端配置值（纯净默认配置）
   - `zero.DefaultRpcServerConfig()` - 返回RPC服务端配置指针（支持链式调用）
   - `zero.DefaultRpcClient()` - 返回RPC客户端配置值（纯净默认配置）
   - `zero.DefaultRpcClientConfig()` - 返回RPC客户端配置指针（支持链式调用）
   - `zero.DefaultEtcd()` - 返回Etcd配置值（纯净默认配置）
   - `zero.DefaultEtcdConfig()` - 返回Etcd配置指针（支持链式调用）
   - `zero.DefaultLogConf()` - 返回日志配置值（纯净默认配置）
   - `zero.DefaultLogConfConfig()` - 返回日志配置指针（支持链式调用）
   - `zero.DefaultPrometheus()` - 返回Prometheus配置值（纯净默认配置）
   - `zero.DefaultPrometheusConfig()` - 返回Prometheus配置指针（支持链式调用）
   - `zero.DefaultTelemetry()` - 返回链路追踪配置值（纯净默认配置）
   - `zero.DefaultTelemetryConfig()` - 返回链路追踪配置指针（支持链式调用）
   - `zero.DefaultRestful()` - 返回RESTful配置值（纯净默认配置）
   - `zero.DefaultRestfulConfig()` - 返回RESTful配置指针（支持链式调用）
   - `zero.DefaultSignature()` - 返回签名配置值（纯净默认配置）
   - `zero.DefaultSignatureConfig()` - 返回签名配置指针（支持链式调用）

3. **pay包处理**：由于pay包含多个配置类型，使用以下方式：
   - `pay.Default()` - 返回支付宝配置指针（主要支付方式，支持链式调用）
   - `pay.DefaultAliPay()` - 返回支付宝配置值（纯净默认配置）
   - `pay.DefaultAliPayConfig()` - 返回支付宝配置指针（支持链式调用）
   - `pay.DefaultWechatPay()` - 返回微信支付配置值（纯净默认配置）
   - `pay.DefaultWechatPayConfig()` - 返回微信支付配置指针（支持链式调用）

4. **database包处理**：由于database包含多个配置类型，使用以下方式：
   - `database.Default()` - 返回MySQL配置指针（主要数据库，支持链式调用）
   - `database.DefaultMySQL()` - 返回MySQL配置值（纯净默认配置）
   - `database.DefaultMySQLConfig()` - 返回MySQL配置指针（支持链式调用）
   - `database.DefaultPostgreSQL()` - 返回PostgreSQL配置值（纯净默认配置）
   - `database.DefaultPostgreSQLConfig()` - 返回PostgreSQL配置指针（支持链式调用）
   - `database.DefaultSQLite()` - 返回SQLite配置值（纯净默认配置）
   - `database.DefaultSQLiteConfig()` - 返回SQLite配置指针（支持链式调用）

5. **oss包处理**：由于oss包含多个配置类型，使用以下方式：
   - `oss.DefaultS3()` - 返回S3配置值（纯净默认配置）
   - `oss.DefaultS3Config()` - 返回S3配置指针（支持链式调用）
   - `oss.DefaultAliyunOss()` - 返回阿里云OSS配置值（纯净默认配置）
   - `oss.DefaultAliyunOssConfig()` - 返回阿里云OSS配置指针（支持链式调用）
   - `oss.DefaultMinio()` - 返回MinIO配置值（纯净默认配置）
   - `oss.DefaultMinioConfig()` - 返回MinIO配置指针（支持链式调用）

6. **register包处理**：由于register包含多个配置类型，使用以下方式：
   - `register.Default()` - 返回服务器配置指针（主要服务，支持链式调用）
   - `register.DefaultServer()` - 返回服务器配置值（纯净默认配置）
   - `register.DefaultServerConfig()` - 返回服务器配置指针（支持链式调用）
   - `register.DefaultConsul()` - 返回Consul配置值（纯净默认配置）
   - `register.DefaultConsulConfig()` - 返回Consul配置指针（支持链式调用）
   - `register.DefaultJaeger()` - 返回Jaeger配置值（纯净默认配置）
   - `register.DefaultJaegerConfig()` - 返回Jaeger配置指针（支持链式调用）
   - `register.DefaultPProf()` - 返回PProf配置值（纯净默认配置）
   - `register.DefaultPProfConfig()` - 返回PProf配置指针（支持链式调用）

7. **cache子包**：通过cache主包统一提供Default方法，子包如memory、redis等通过主包访问

## 向后兼容

新的链式方法与现有的配置方式完全兼容：
- 现有的 `NewXXX()` 构造函数依然可用
- 现有的字段访问方式依然可用
- 现有的 `Validate()` 方法依然可用
- 现有的 `Clone()` 方法依然可用

## 扩展其他包

如果要为其他包添加相同的功能，请遵循以下统一模式：

1. 添加 `DefaultXXX()` 函数返回默认配置（值类型）
2. 添加 `Default()` 函数返回配置指针（支持链式调用）
3. 为每个字段添加 `WithXXX()` 方法，返回 `*XXX` 类型以支持链式调用

示例：
```go
// DefaultExample 返回默认Example配置值
func DefaultExample() Example {
    return Example{
        ModuleName: "example",
        Field1: "default_value",
        Field2: 42,
    }
}

// Default 返回默认Example配置的指针，支持链式调用
func Default() *Example {
    config := DefaultExample()
    return &config
}

// WithField1 设置字段1
func (e *Example) WithField1(field1 string) *Example {
    e.Field1 = field1
    return e
}

// WithField2 设置字段2
func (e *Example) WithField2(field2 int) *Example {
    e.Field2 = field2
    return e
}
```

### 统一格式要点

1. **命名约定**：
   - `Default()` 总是返回指针类型（支持链式调用）
   - `DefaultXXX()` 总是返回值类型（纯净默认配置）
   - `WithXXX()` 方法总是返回指针类型

2. **多类型包的特殊处理**：
   - 对于包含多个配置类型的包，使用 `DefaultXXXConfig()` 返回指针类型
   - 避免同包内 `Default()` 函数名冲突

3. **返回值类型**：
   - 指针类型用于链式调用
   - 值类型用于提供纯净默认配置

4. **使用场景**：
   - 需要链式配置时使用返回指针类型的函数
   - 需要多个独立实例时使用返回值类型的函数