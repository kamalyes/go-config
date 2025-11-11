# Go Config 配置文件详细说明

本文档详细说明了所有配置项的含义、用途及生产环境配置建议。

## 目录

- [服务器配置 (Server)](#服务器配置-server)
- [缓存配置 (Cache)](#缓存配置-cache) 
- [跨域配置 (CORS)](#跨域配置-cors)
- [服务注册 (Consul)](#服务注册-consul)
- [验证码 (Captcha)](#验证码-captcha)
- [数据库配置](#数据库配置)
- [Redis配置](#redis配置)
- [日志配置 (Zap)](#日志配置-zap)
- [JWT认证](#jwt认证)
- [对象存储](#对象存储)
- [消息队列 (MQTT)](#消息队列-mqtt)
- [支付配置](#支付配置)
- [短信服务](#短信服务)
- [监控配置](#监控配置)
- [安全配置](#安全配置)
- [中间件配置](#中间件配置)
- [网关配置](#网关配置)

## 服务器配置 (Server)

### 基础配置
```yaml
server:
  modulename: "production-server"        # 模块名称，用于标识服务
  endpoint: "https://api.example.com"    # 服务对外访问地址
  host: "0.0.0.0"                       # 监听地址，0.0.0.0表示所有网卡
  port: 8080                            # HTTP端口
  read-timeout: 30                      # 读取超时时间（秒）
  write-timeout: 30                     # 写入超时时间（秒）
  idle-timeout: 120                     # 空闲超时时间（秒）
  grpc-port: 9090                       # gRPC端口
  enable-http: true                     # 是否启用HTTP服务
  enable-grpc: true                     # 是否启用gRPC服务
  enable-tls: true                      # 是否启用TLS
  tls:
    cert-file: "/etc/ssl/certs/server.crt"   # TLS证书文件路径
    key-file: "/etc/ssl/private/server.key"   # TLS私钥文件路径
    ca-file: "/etc/ssl/certs/ca.crt"          # CA证书文件路径
  headers:                              # 自定义响应头
    X-Server-Version: "v1.0.0"
    X-Environment: "production"
```

**生产环境建议：**
- `host`: 使用 "0.0.0.0" 允许外部访问
- `port`: 标准端口 8080 或 80/443
- 超时时间根据业务需求调整，一般读写超时30秒，空闲超时120秒
- 生产环境必须启用TLS
- 设置适当的响应头用于调试和监控

## 缓存配置 (Cache)

### 多级缓存策略
```yaml
cache:
  modulename: "production-cache"
  type: "two-level"                     # 缓存类型: memory, redis, ristretto, sharded, two-level, expiring
  enabled: true                         # 是否启用缓存
  default-ttl: 3600                     # 默认TTL（秒），0表示永不过期
  key-prefix: "prod:cache:"             # 缓存键前缀
  serializer: "json"                    # 序列化方式: json, msgpack, gob
  
  # 内存缓存配置
  memory:
    modulename: "memory-cache"
    capacity: 10000                     # 缓存容量
    default-ttl: 1800                   # 默认TTL（秒）
    cleanup-size: 1000                  # 清理时删除的条目数
    max-size: 10000                     # 最大条目数
  
  # Ristretto缓存配置（高性能内存缓存）
  ristretto:
    modulename: "ristretto-cache"
    num-counters: 100000                # 计数器数量，通常是最大条目数的10倍
    max-cost: 1073741824               # 最大成本（字节），1GB
    buffer-items: 64                    # 缓冲区大小
    metrics: true                       # 是否启用指标收集
    ignore-internal-cost: false         # 是否忽略内部成本计算
    key-to-hash: true                   # 是否对键进行哈希
    cost: 1                             # 默认条目成本
  
  # Redis缓存配置
  redis:
    modulename: "redis-cache"
    addr: "redis-cluster.example.com:6379"  # Redis地址
    addrs:                              # Redis集群地址
      - "redis1.example.com:6379"
      - "redis2.example.com:6379"
      - "redis3.example.com:6379"
    username: "cache_user"               # Redis用户名
    password: "secure_password_123"      # Redis密码
    db: 1                               # Redis数据库编号
    max-retries: 3                      # 最大重试次数
    pool-size: 100                      # 连接池大小
    min-idle-conns: 10                  # 最小空闲连接数
    max-conn-age: 3600                  # 连接最大年龄（秒）
    pool-timeout: 5                     # 连接池获取连接超时（秒）
    idle-timeout: 300                   # 空闲连接超时（秒）
    read-timeout: 3                     # 读取超时（秒）
    write-timeout: 3                    # 写入超时（秒）
    cluster-mode: true                  # 是否启用集群模式
  
  # 分片缓存配置
  sharded:
    modulename: "sharded-cache"
    shard-count: 16                     # 分片数量，建议为2的幂
    base-type: "memory"                 # 基础缓存类型
    hash-func: "fnv32"                  # 哈希函数: fnv32, fnv64, crc32
    load-balancer: "consistent"         # 负载均衡策略: round-robin, consistent
  
  # 二级缓存配置
  two-level:
    modulename: "two-level-cache"
    l1-type: "memory"                   # L1缓存类型（通常是内存）
    l2-type: "redis"                    # L2缓存类型（通常是Redis）
    l1-ttl: 300                         # L1缓存TTL（秒）
    l2-ttl: 3600                        # L2缓存TTL（秒）
    sync-strategy: "write-through"       # 同步策略: write-through, write-behind, write-around
    l1-size: 1000                       # L1缓存大小
    l2-size: 10000                      # L2缓存大小
    promote-threshold: 2                # 提升到L1缓存的访问阈值
  
  # 过期缓存配置
  expiring:
    modulename: "expiring-cache"
    cleanup-interval: 300               # 清理间隔（秒）
    default-ttl: 1800                   # 默认TTL（秒）
    max-size: 10000                     # 最大条目数
    eviction-policy: "lru"              # 淘汰策略: lru, lfu, fifo, random
    enable-lazy-expiry: true            # 是否启用懒惰过期
    max-memory-usage: 134217728         # 最大内存使用（字节），128MB
```

**生产环境建议：**
- 使用二级缓存提高性能
- Redis集群模式提高可用性
- 合理设置TTL避免缓存雪崩
- 启用指标收集进行监控

## 跨域配置 (CORS)

```yaml
cors:
  modulename: "production-cors"
  allowed-origins:                      # 允许的来源域名
    - "https://web.example.com"
    - "https://admin.example.com"
    - "https://mobile.example.com"
  allowed-methods:                      # 允许的HTTP方法
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
    - "PATCH"
    - "OPTIONS"
  allowed-headers:                      # 允许的请求头
    - "Origin"
    - "Content-Type"
    - "Accept"
    - "Authorization"
    - "X-Request-ID"
    - "X-Requested-With"
  exposed-headers:                      # 暴露给客户端的响应头
    - "X-Total-Count"
    - "X-Page-Count"
    - "X-Rate-Limit-Remaining"
  max-age: "86400"                      # 预检请求缓存时间（秒）
  allowed-all-origins: false            # 是否允许所有来源（生产环境建议false）
  allowed-all-methods: false            # 是否允许所有方法（生产环境建议false）
  allow-credentials: true               # 是否允许凭证
  options-response-code: 204            # OPTIONS请求响应状态码
```

**生产环境建议：**
- 明确指定允许的域名，不要使用通配符
- 只开放必要的HTTP方法
- 合理设置max-age减少预检请求

## 服务注册 (Consul)

```yaml
consul:
  modulename: "production-consul"
  endpoint: "https://consul.example.com:8500"  # Consul服务地址
  register-interval: 30                        # 注册间隔（秒）
```

**生产环境建议：**
- 使用Consul集群提高可用性
- 合理设置注册间隔，避免过于频繁

## 验证码 (Captcha)

```yaml
captcha:
  modulename: "production-captcha"
  key-len: 4                           # 验证码长度
  img-width: 120                       # 图片宽度
  img-height: 40                       # 图片高度
  max-skew: 0.7                        # 最大倾斜度
  dot-count: 80                        # 干扰点数量
```

## 数据库配置

### MySQL配置
```yaml
mysql:
  modulename: "production-mysql"
  host: "mysql-cluster.example.com"    # MySQL主机地址
  port: "3306"                         # MySQL端口
  config: "charset=utf8mb4&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=30s"
  log-level: "warn"                    # 日志级别: silent, error, warn, info
  db-name: "production_db"             # 数据库名
  username: "app_user"                 # 数据库用户名
  password: "secure_db_password_123"   # 数据库密码
  max-idle-conns: 10                   # 最大空闲连接数
  max-open-conns: 100                  # 最大打开连接数
  conn-max-idle-time: 300              # 连接最大空闲时间（秒）
  conn-max-life-time: 3600             # 连接最大生命周期（秒）
```

### PostgreSQL配置
```yaml
postgre:
  modulename: "production-postgres"
  host: "postgres-cluster.example.com"
  port: "5432"
  config: "sslmode=require TimeZone=Asia/Shanghai application_name=production_app"
  log-level: "warn"
  db-name: "production_db"
  username: "app_user"
  password: "secure_db_password_123"
  max-idle-conns: 10
  max-open-conns: 100
  conn-max-idle-time: 300
  conn-max-life-time: 3600
```

### SQLite配置（适用于开发环境）
```yaml
sqlite:
  modulename: "development-sqlite"
  db-path: "/data/app/sqlite.db"       # 数据库文件路径
  max-idle-conns: 5
  max-open-conns: 10
  log-level: "warn"
  conn-max-idle-time: 300
  conn-max-life-time: 3600
  vacuum: true                         # 是否执行VACUUM操作
```

**生产环境建议：**
- 使用连接池避免连接泄露
- 设置合理的超时时间
- 启用SSL连接保证安全
- 定期备份数据库

## Redis配置

```yaml
redis:
  modulename: "production-redis"
  addr: "redis-cluster.example.com:6379"
  addrs:                               # 集群模式下的多个节点
    - "redis1.example.com:6379"
    - "redis2.example.com:6379" 
    - "redis3.example.com:6379"
  username: "redis_user"               # Redis 6.0+ 支持用户名
  password: "secure_redis_password"    # Redis密码
  db: 0                                # 数据库编号
  max-retries: 3                       # 最大重试次数
  pool-size: 100                       # 连接池大小
  min-idle-conns: 10                   # 最小空闲连接
  max-conn-age: 0                      # 连接最大年龄（0表示无限制）
  pool-timeout: 5                      # 获取连接超时
  idle-timeout: 300                    # 空闲超时
  read-timeout: 3                      # 读取超时
  write-timeout: 3                     # 写入超时
  cluster-mode: true                   # 集群模式
```

## 日志配置 (Zap)

```yaml
zap:
  modulename: "production-logging"
  level: "info"                        # 日志级别: debug, info, warn, error, panic, fatal
  format: "json"                       # 格式: json, console
  prefix: "[PROD-API]"                 # 日志前缀
  director: "/var/log/app"             # 日志目录
  max-size: 100                        # 单个日志文件最大大小(MB)
  max-age: 30                          # 日志文件保留天数
  max-backups: 10                      # 最大备份文件数
  compress: true                       # 是否压缩备份文件
  link-name: "/var/log/app/current.log" # 软链接文件名
  show-line: false                     # 是否显示代码行号（生产环境建议false）
  encode-level: "lowercase"            # 级别编码: lowercase, capital, color
  stacktrace-key: "stacktrace"         # 堆栈跟踪键名
  log-in-console: false                # 是否同时在控制台输出（生产环境建议false）
  development: false                   # 是否开发模式（生产环境必须false）
```

**生产环境建议：**
- 使用json格式便于日志分析
- 设置合理的日志轮转策略
- 关闭控制台输出提高性能
- 不要在生产环境显示代码行号

## JWT认证

```yaml
jwt:
  modulename: "production-jwt"
  signing-key: "your-256-bit-secret-key-here-must-be-very-secure" # 签名密钥
  expires-time: 86400                  # 访问令牌过期时间（秒），24小时
  buffer-time: 3600                    # 缓冲时间（秒）
  use-multipoint: true                 # 是否启用多点登录控制
  issuer: "api.example.com"           # 签发者
  audience: "app.example.com"         # 受众
  algorithm: "HS256"                   # 签名算法: HS256, HS384, HS512, RS256, RS384, RS512
  enable-refresh: true                 # 是否启用刷新令牌
  refresh-token-life: 604800          # 刷新令牌生命周期（秒），7天
  subject: "user-authentication"       # 主题
  custom-claims:                       # 自定义声明
    app-version: "1.0.0"
    environment: "production"
```

**生产环境建议：**
- 使用强签名密钥，建议256位
- 合理设置令牌过期时间
- 启用刷新令牌机制
- 考虑使用RS256算法提高安全性

## 对象存储

### MinIO配置
```yaml
minio:
  modulename: "production-minio"
  endpoint: "https://minio.example.com:9000" # MinIO服务地址
  access-key: "production-access-key"          # 访问密钥
  secret-key: "production-secret-key-secure"   # 密钥
  bucket: "production-files"                   # 存储桶名称
```

### 阿里云OSS配置
```yaml
aliyunoss:
  modulename: "production-aliyun-oss"
  access-key: "LTAI5tNhK8bE7X9Z2mA"         # 阿里云AccessKey
  secret-key: "8qK2dT1mN5xR9vH3wP6sF4jL"   # 阿里云SecretKey
  endpoint: "https://oss-cn-hangzhou.aliyuncs.com" # OSS端点
  bucket: "production-assets"                       # 存储桶
  region: "cn-hangzhou"                            # 区域
  replace-original-host: "oss-cn-hangzhou.aliyuncs.com"  # 原始主机
  replace-later-host: "cdn.example.com"                  # 替换后的CDN域名
```

### AWS S3配置
```yaml
s3:
  modulename: "production-s3"
  endpoint: "https://s3.us-west-2.amazonaws.com"  # S3端点
  region: "us-west-2"                             # AWS区域
  access-key: "AKIAIOSFODNN7EXAMPLE"              # AWS访问密钥
  secret-key: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" # AWS私钥
  bucket-prefix: "production"                      # 存储桶前缀
  session-token: ""                               # 会话令牌（临时凭证）
  use-ssl: true                                   # 使用SSL
  path-style: false                               # 使用虚拟主机样式
```

## 消息队列 (MQTT)

```yaml
mqtt:
  modulename: "production-mqtt"
  endpoint: "ssl://mqtt.example.com:8883"    # MQTT服务器地址
  protocol-version: 4                         # MQTT协议版本（4=v3.1.1, 5=v5.0）
  keep-alive: 60                             # 保活间隔（秒）
  max-reconnect-interval: 60                 # 最大重连间隔（秒）
  ping-timeout: 10                           # Ping超时（秒）
  write-timeout: 10                          # 写超时（秒）
  connect-timeout: 30                        # 连接超时（秒）
  username: "mqtt_user"                      # MQTT用户名
  password: "secure_mqtt_password"           # MQTT密码
  clean-session: true                        # 清理会话
  auto-reconnect: true                       # 自动重连
  will-topic: "devices/last-will"            # 遗嘱主题
```

## 支付配置

### 支付宝配置
```yaml
alipay:
  modulename: "production-alipay"
  pid: "2088123456789012"                    # 商户PID
  app-id: "2021123456789012"                 # 应用ID
  pri-key: "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC..."  # 私钥
  pub-key: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA..."        # 公钥
  sign-type: "RSA2"                          # 签名类型
  notify-url: "https://api.example.com/alipay/notify"  # 异步通知地址
  subject: "商品购买"                        # 默认订单主题
```

### 微信支付配置
```yaml
wechatpay:
  modulename: "production-wechatpay"
  app-id: "wx1234567890123456"               # 应用ID
  mch-id: "1234567890"                       # 商户号
  notify-url: "https://api.example.com/wechat/notify"  # 异步通知地址
  api-key: "abcdef1234567890abcdef1234567890" # API密钥
  cert-p12-path: "/etc/ssl/wechat/cert.p12"  # 证书路径
```

## 短信服务

### 阿里云短信
```yaml
aliyunsms:
  modulename: "production-aliyun-sms"
  secret-id: "LTAI5tNhK8bE7X9Z2mA"          # AccessKeyId
  secret-key: "8qK2dT1mN5xR9vH3wP6sF4jL"    # AccessKeySecret
  sign: "阿里云"                            # 短信签名
  resource-owner-account: ""                 # 资源拥有者账号
  resource-owner-id: 0                      # 资源拥有者ID
  template-code-verify: "SMS_123456789"     # 验证码模板ID
  endpoint: "dysmsapi.aliyuncs.com"         # 服务端点
```

### 阿里云STS
```yaml
aliyunsts:
  modulename: "production-aliyun-sts"
  region-id: "cn-hangzhou"                  # 区域ID
  access-key-id: "LTAI5tNhK8bE7X9Z2mA"     # AccessKeyId
  access-key-secret: "8qK2dT1mN5xR9vH3wP6sF4jL" # AccessKeySecret
  role-arn: "acs:ram::123456789:role/AliyunOSSRole" # 角色ARN
  role-session-name: "production-session"   # 角色会话名
```

## 监控配置

### Prometheus配置
```yaml
prometheus:
  modulename: "production-prometheus"
  enabled: true                            # 是否启用Prometheus
  path: "/metrics"                         # 指标路径
  port: 9090                              # Prometheus端口
  endpoint: "https://prometheus.example.com" # Prometheus服务地址
  push-gateway:                           # PushGateway配置
    enabled: true
    endpoint: "https://pushgateway.example.com"
    job-name: "production-api"
  scraping:                               # 抓取配置
    interval: "15s"                       # 抓取间隔
    timeout: "10s"                        # 抓取超时
    metrics-path: "/metrics"              # 指标路径
```

### Jaeger配置
```yaml
jaeger:
  modulename: "production-jaeger"
  enabled: true                           # 是否启用链路追踪
  endpoint: "https://jaeger.example.com"  # Jaeger收集器地址
  service-name: "production-api"          # 服务名称
  sample-rate: 1                          # 采样率（0-100）
  agent:                                  # Agent配置
    host: "jaeger-agent.example.com"
    port: 6831
  collector:                              # 收集器配置
    endpoint: "https://jaeger-collector.example.com:14268/api/traces"
    username: "jaeger_user"
    password: "jaeger_password"
  sampling:                               # 采样配置
    type: "probabilistic"                 # 采样类型
    param: 1.0                           # 采样参数
    max-traces-per-second: 100           # 每秒最大trace数
    operation-sampling:                   # 操作级采样
      - operation: "GET /api/users"
        max-traces-per-second: 10
        probabilistic-sampling: 0.1
  tags:                                   # 全局标签
    environment: "production"
    version: "1.0.0"
```

### PProf性能分析
```yaml
pprof:
  modulename: "production-pprof"
  enabled: true                           # 是否启用性能分析
  path-prefix: "/debug/pprof"            # 路径前缀
  port: 6060                             # pprof端口
  enable-profiles:                       # 启用的分析类型
    cpu: true                            # CPU分析
    memory: true                         # 内存分析
    goroutine: true                      # Goroutine分析
    block: true                          # 阻塞分析
    mutex: true                          # 互斥锁分析
    heap: true                           # 堆分析
    allocs: true                         # 内存分配分析
    threadcreate: true                   # 线程创建分析
    trace: true                          # 执行跟踪
  sampling:                              # 采样配置
    cpu-rate: 100                        # CPU采样率
    memory-rate: 524288                  # 内存采样率
    block-rate: 1                        # 阻塞采样率
    mutex-fraction: 1                    # 互斥锁采样分数
```

## 安全配置

```yaml
security:
  modulename: "production-security"
  enabled: true                          # 是否启用安全功能
  jwt:                                   # JWT安全配置
    enabled: true
    secret: "your-super-secret-jwt-key-256-bits-long"
    expiry: 86400                        # JWT过期时间（秒）
    issuer: "api.example.com"
    algorithm: "HS256"
  rate-limit:                            # 速率限制
    enabled: true
    requests-per-min: 100                # 每分钟最大请求数
    burst-size: 10                       # 突发大小
    cleanup-interval: 60                 # 清理间隔（秒）
  auth:                                  # 认证配置
    enabled: true
    type: "bearer"                       # 认证类型: basic, bearer, apikey, custom
    header-name: "Authorization"         # 认证头名称
    token-prefix: "Bearer "              # 令牌前缀
    basic:                               # 基础认证配置
      users:
        - username: "admin"
          password: "secure_admin_password"
          role: "admin"
          permissions: ["read", "write", "delete"]
    bearer:                              # Bearer令牌配置
      tokens:
        - "secure-bearer-token-1"
        - "secure-bearer-token-2"
    apikey:                              # API Key配置
      keys:
        - "api-key-1-secure"
        - "api-key-2-secure"
      header-name: "X-API-Key"
      query-param: "api_key"
    custom:                              # 自定义认证
      header-name: "X-Custom-Auth"
      expected-value: "custom-auth-value"
      headers:
        X-Client-Version: "1.0.0"
  protection:                            # 端点保护
    swagger:                             # Swagger文档保护
      enabled: true
      auth-required: true
      auth-type: "basic"
      ip-whitelist:
        - "192.168.1.0/24"
        - "10.0.0.0/8"
      require-https: true
      username: "swagger_user"
      password: "swagger_password"
    pprof:                               # PProf保护
      enabled: true
      auth-required: true
      auth-type: "basic"
      ip-whitelist:
        - "127.0.0.1"
        - "::1"
      require-https: true
      username: "pprof_user"
      password: "pprof_password"
    metrics:                             # 指标端点保护
      enabled: true
      auth-required: true
      auth-type: "bearer"
      ip-whitelist:
        - "192.168.1.0/24"
      require-https: true
      username: "metrics_user"
      password: "metrics_password"
    health:                              # 健康检查保护
      enabled: true
      auth-required: false               # 健康检查通常不需要认证
      auth-type: "none"
      ip-whitelist: []                   # 允许所有IP访问
      require-https: false
      username: ""
      password: ""
    api:                                 # API保护
      enabled: true
      auth-required: true
      auth-type: "jwt"
      ip-whitelist: []
      require-https: true
      username: ""
      password: ""
```

## 中间件配置

```yaml
middleware:
  modulename: "production-middleware"
  enabled: true
  logging:                               # 日志中间件
    modulename: "request-logging"
    enabled: true
    level: "info"
    format: "json"
    output: "file"
    file-path: "/var/log/app/requests.log"
    max-size: 100
    max-backups: 10
    max-age: 30
    compress: true
    skip-paths:                          # 跳过记录的路径
      - "/health"
      - "/metrics"
      - "/debug/pprof"
    enable-request: true                 # 记录请求
    enable-response: true                # 记录响应
  recovery:                              # 恢复中间件
    modulename: "panic-recovery"
    enabled: true
    print-stack: false                   # 是否打印堆栈（生产环境建议false）
    log-level: "error"
    enable-notify: true                  # 启用通知
  tracing:                               # 链路追踪中间件
    modulename: "request-tracing"
    enabled: true
    service-name: "production-api"
    service-version: "1.0.0"
    environment: "production"
    endpoint: "https://jaeger-collector.example.com"
    exporter-type: "jaeger"
    exporter-endpoint: "https://jaeger.example.com:14268"
    sample-rate: 1                       # 采样率（0-100）
    sampler-type: "probabilistic"
    sampler-probability: 1.0
    sampler-rate: 100
    headers:                             # 传播的头部
      - "X-Request-ID"
      - "X-User-ID"
    attributes:                          # 自定义属性
      service.environment: "production"
      service.version: "1.0.0"
  metrics:                               # 指标中间件
    modulename: "request-metrics"
    enabled: true
    path: "/metrics"
    subsystem: "http"
    skip-paths:
      - "/health"
    buckets:                             # 自定义延迟桶
      - 0.005
      - 0.01
      - 0.025
      - 0.05
      - 0.1
      - 0.25
      - 0.5
      - 1
      - 2.5
      - 5
      - 10
    request-count: true                  # 请求计数
    duration: true                       # 请求耗时
    request-size: true                   # 请求大小
    response-size: true                  # 响应大小
  request-id:                            # 请求ID中间件
    modulename: "request-id"
    enabled: true
    header-name: "X-Request-ID"
    generator: "uuid"                    # 生成器类型: uuid, nanoid, ulid
  i18n:                                  # 国际化中间件
    modulename: "internationalization"
    enabled: true
    default-language: "zh-CN"
    supported-languages:
      - "zh-CN"
      - "en-US"
      - "ja-JP"
      - "ko-KR"
    locale-dir: "/app/locales"
    header-name: "Accept-Language"
    query-param: "lang"
    cookie-name: "locale"
  pprof:                                 # PProf中间件
    modulename: "pprof-middleware"
    enabled: true
    path-prefix: "/debug/pprof"
    port: 6060
    enable-profiles:
      cpu: true
      memory: true
      goroutine: true
      block: true
      mutex: true
      heap: true
      allocs: true
      threadcreate: true
      trace: true
    sampling:
      cpu-rate: 100
      memory-rate: 524288
      block-rate: 1
      mutex-fraction: 1
```

## 网关配置

```yaml
gateway:
  modulename: "production-gateway"
  name: "API Gateway"
  enabled: true
  debug: false                           # 生产环境关闭调试
  version: "1.0.0"
  environment: "production"
  http:                                  # HTTP服务配置
    host: "0.0.0.0"
    port: 8080
    read-timeout: 30
    write-timeout: 30
    idle-timeout: 120
    max-header-bytes: 1048576           # 1MB
    enable-gzip-compress: true
    endpoint: "https://api.example.com"
  grpc:                                  # gRPC服务配置
    server:
      host: "0.0.0.0"
      port: 9090
      network: "tcp"
      max-recv-msg-size: 4194304        # 4MB
      max-send-msg-size: 4194304        # 4MB
      keepalive-time: 60
      keepalive-timeout: 20
      connection-timeout: 30
      enable-reflection: true
      endpoint: "grpc://api.example.com:9090"
    clients:                             # gRPC客户端配置
      user-service:
        endpoint: "grpc://user-service.example.com:9090"
        timeout: 5000
        max-recv-msg-size: 4194304
        max-send-msg-size: 4194304
      order-service:
        endpoint: "grpc://order-service.example.com:9090"
        timeout: 10000
        max-recv-msg-size: 4194304
        max-send-msg-size: 4194304
  # 其他配置项继承自对应的模块配置...
```

**生产环境部署建议：**

1. **资源配置**：根据业务量合理配置连接池、缓存大小等
2. **安全配置**：启用TLS、配置防火墙、使用强密码
3. **监控告警**：配置完整的监控指标和告警规则
4. **日志管理**：使用结构化日志，配置日志聚合
5. **备份策略**：定期备份数据库和配置文件
6. **容灾方案**：配置多区域部署和故障转移
7. **性能优化**：根据压测结果调优各项参数

## 配置文件管理

### 环境分离
- `dev_config.yaml` - 开发环境
- `test_config.yaml` - 测试环境  
- `staging_config.yaml` - 预生产环境
- `prod_config.yaml` - 生产环境

### 敏感信息管理
- 使用环境变量存储密码、密钥等敏感信息
- 配置文件中使用占位符，运行时替换
- 使用配置管理服务如 Consul、etcd
- 启用配置文件加密

### 配置验证
- 启动时验证必要配置项
- 使用配置Schema验证格式
- 提供配置检查命令
- 配置热重载支持

本文档涵盖了所有主要配置模块，每个配置项都包含详细说明和生产环境建议。请根据实际业务需求调整具体的配置值。