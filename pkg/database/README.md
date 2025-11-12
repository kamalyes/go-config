# Database Configuration Package

这个包提供了一个灵活且功能丰富的数据库配置管理解决方案，支持 MySQL、PostgreSQL 和 SQLite 数据库。

## 特性

- 支持多种数据库类型（MySQL、PostgreSQL、SQLite）
- 链式调用配置方法
- 配置验证
- 连接池配置
- 自动 DSN 生成
- 配置克隆和序列化
- 内置默认配置

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/kamalyes/go-config/pkg/database"
)

func main() {
    // 使用默认配置
    db := database.Default()
    fmt.Println(db.GetDSN()) // 输出: root:@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
    
    // 使用链式调用配置
    mysqlDB := database.Default().
        WithHost("192.168.1.100").
        WithPort("3306").
        WithUsername("admin").
        WithPassword("password123").
        WithDbname("myapp").
        WithModuleName("mysql-main")
    
    fmt.Println(mysqlDB.GetDSN())
}
```

### 不同数据库类型的配置

#### MySQL 配置

```go
mysql := database.DefaultMySQL().
    WithHost("localhost").
    WithPort("3306").
    WithUsername("root").
    WithPassword("password").
    WithDbname("myapp").
    WithMaxOpenConns(100).
    WithMaxIdleConns(20)

fmt.Println("MySQL DSN:", mysql.GetDSN())
```

#### PostgreSQL 配置

```go
postgres := database.DefaultPostgreSQL().
    WithHost("localhost").
    WithPort("5432").
    WithUsername("postgres").
    WithPassword("password").
    WithDbname("myapp").
    WithConfig("sslmode=disable&TimeZone=Asia/Shanghai")

fmt.Println("PostgreSQL DSN:", postgres.GetDSN())
```

#### SQLite 配置

```go
sqlite := database.DefaultSQLite().
    WithDbPath("./data/app.db").
    WithConfig("cache=shared&mode=rwc&_journal_mode=WAL").
    WithVacuum(true)

fmt.Println("SQLite DSN:", sqlite.GetDSN())
```

### 配置验证

```go
db := database.Default().
    WithMaxIdleConns(-1) // 无效配置

// 验证基本结构
if err := db.Validate(); err != nil {
    fmt.Printf("配置验证失败: %v\n", err)
}

// 验证连接参数
if err := db.ValidateConnections(); err != nil {
    fmt.Printf("连接参数验证失败: %v\n", err)
}
```

### 配置克隆

```go
original := database.Default().WithHost("original-host")
cloned := original.Clone().(*database.Database)

// 修改克隆的配置不会影响原始配置
cloned.WithHost("cloned-host")

fmt.Println("原始主机:", original.Host)     // 输出: original-host
fmt.Println("克隆主机:", cloned.Host)       // 输出: cloned-host
```

### 获取连接信息

```go
db := database.Default().WithModuleName("main-db")

info := db.GetConnectionInfo()
fmt.Printf("连接信息: %+v\n", info)
```

### 类型检查

```go
db := database.DefaultMySQL()

if db.IsMySQL() {
    fmt.Println("这是 MySQL 数据库")
}

if db.IsPostgreSQL() {
    fmt.Println("这是 PostgreSQL 数据库")
}

if db.IsSQLite() {
    fmt.Println("这是 SQLite 数据库")
}
```

## 配置字段说明

| 字段 | 类型 | 说明 | 默认值 |
|------|------|------|---------|
| `DBType` | string | 数据库类型 (mysql/postgres/sqlite) | "mysql" |
| `Host` | string | 数据库主机地址 | "localhost" |
| `Port` | string | 数据库端口 | "3306" |
| `Dbname` | string | 数据库名称 | "test" |
| `Username` | string | 数据库用户名 | "root" |
| `Password` | string | 数据库密码 | "" |
| `Config` | string | 连接参数 | "charset=utf8mb4&parseTime=True&loc=Local" |
| `LogLevel` | string | 日志级别 | "info" |
| `MaxIdleConns` | int | 最大空闲连接数 | 10 |
| `MaxOpenConns` | int | 最大连接数 | 100 |
| `ConnMaxIdleTime` | int | 连接最大空闲时间（秒） | 600 |
| `ConnMaxLifeTime` | int | 连接最大生命周期（秒） | 3600 |
| `DbPath` | string | SQLite 数据库文件路径 | "./data/app.db" |
| `Vacuum` | bool | 是否执行 SQLite 清理 | false |
| `ModuleName` | string | 模块名称 | "database" |

## YAML 配置文件示例

```yaml
database:
  db-type: "mysql"
  host: "localhost"
  port: "3306"
  db-name: "myapp"
  username: "root"
  password: "password"
  config: "charset=utf8mb4&parseTime=True&loc=Local"
  log-level: "info"
  max-idle-conns: 10
  max-open-conns: 100
  conn-max-idle-time: 600
  conn-max-life-time: 3600
  modulename: "database"
```

## 链式方法列表

- `WithModuleName(string)` - 设置模块名称
- `WithDBType(string)` - 设置数据库类型
- `WithHost(string)` - 设置主机地址
- `WithPort(string)` - 设置端口
- `WithConfig(string)` - 设置连接参数
- `WithLogLevel(string)` - 设置日志级别
- `WithDbname(string)` - 设置数据库名称
- `WithUsername(string)` - 设置用户名
- `WithPassword(string)` - 设置密码
- `WithMaxIdleConns(int)` - 设置最大空闲连接数
- `WithMaxOpenConns(int)` - 设置最大连接数
- `WithConnMaxIdleTime(int)` - 设置连接最大空闲时间
- `WithConnMaxLifeTime(int)` - 设置连接最大生命周期
- `WithDbPath(string)` - 设置 SQLite 文件路径
- `WithVacuum(bool)` - 设置是否执行清理

## 工具方法

- `GetDSN()` - 获取数据库连接字符串
- `GetConnMaxIdleTimeDuration()` - 获取空闲时间的 Duration
- `GetConnMaxLifeTimeDuration()` - 获取生命周期的 Duration
- `IsMySQL()` - 检查是否为 MySQL
- `IsPostgreSQL()` - 检查是否为 PostgreSQL
- `IsSQLite()` - 检查是否为 SQLite
- `GetPortAsInt()` - 获取端口的整数值
- `ValidateConnections()` - 验证连接参数
- `GetConnectionInfo()` - 获取连接信息摘要

## 许可证

Copyright (c) 2024 by kamalyes, All Rights Reserved.