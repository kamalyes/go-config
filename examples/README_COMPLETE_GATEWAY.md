# 完整Gateway演示服务

这是一个完整的Gateway演示服务，展示了如何使用go-config框架构建一个功能完整的网关服务。

## 🚀 功能特性

### 核心功能
- ✅ **HTTP/GRPC网关服务** - 支持HTTP和GRPC协议
- ✅ **MySQL数据库支持** - 完整的数据库连接和配置管理
- ✅ **Redis缓存支持** - 支持单机和集群模式
- ✅ **配置热更新** - 实时配置变更，无需重启服务
- ✅ **Swagger API文档** - 自动生成和展示API文档

### 中间件支持
- ✅ **CORS跨域** - 完整的跨域资源共享支持
- ✅ **请求ID追踪** - 每个请求的唯一标识符
- ✅ **限流保护** - 防止服务过载
- ✅ **安全防护** - 多种安全头设置
- ✅ **错误恢复** - Panic恢复和错误处理
- ✅ **请求超时** - 防止长时间阻塞
- ✅ **日志记录** - 结构化日志输出

### 监控和运维
- ✅ **健康检查** - 服务和组件健康状态检查
- ✅ **监控指标** - Prometheus格式指标导出
- ✅ **Banner展示** - 服务启动横幅
- ✅ **优雅关闭** - 信号处理和优雅退出

## 📁 文件结构

```
examples/
├── complete_gateway_demo_v2.go          # 主要演示程序
├── complete-gateway-config.yaml        # 完整配置文件
├── start_complete_gateway.sh           # Linux/Mac启动脚本
├── start_complete_gateway.bat          # Windows启动脚本
└── README_COMPLETE_GATEWAY.md          # 本文档
```

## 🔧 环境要求

- **Go版本**: 1.19+
- **MySQL数据库**: 5.7+ (可选，演示用)
- **Redis服务**: 5.0+ (可选，演示用)

## 🚀 快速开始

### 方法一：使用启动脚本

**Linux/Mac系统:**
```bash
# 进入examples目录
cd examples

# 给脚本执行权限
chmod +x start_complete_gateway.sh

# 启动服务
./start_complete_gateway.sh
```

**Windows系统:**
```cmd
# 进入examples目录
cd examples

# 运行启动脚本
start_complete_gateway.bat
```

### 方法二：手动启动

```bash
# 设置环境变量
export APP_ENV=development
export CONFIG_PATH=./complete-gateway-config.yaml

# 启动服务
go run complete_gateway_demo_v2.go ./complete-gateway-config.yaml
```

## 📋 API接口说明

### 基础接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET  | `/` | 服务首页，显示基本信息和功能列表 |
| GET  | `/config` | 获取完整的服务配置信息 |
| GET  | `/status` | 获取服务运行状态和组件状态 |

### 健康检查和监控

| 方法 | 路径 | 说明 |
|------|------|------|
| GET  | `/health` | 健康检查，返回各组件健康状态 |
| GET  | `/metrics` | Prometheus格式的监控指标 |

### API文档

| 方法 | 路径 | 说明 |
|------|------|------|
| GET  | `/swagger/` | Swagger UI界面 |
| GET  | `/swagger/doc.json` | API文档JSON格式 |

### 业务接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET  | `/api/users` | 获取用户列表（演示接口） |
| GET  | `/api/users/{id}` | 获取用户详情（演示接口） |
| GET  | `/api/cache/test` | 缓存功能测试 |
| GET  | `/api/db/test` | 数据库连接测试 |

### 管理接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/admin/config/reload` | 手动重新加载配置 |
| GET  | `/admin/config/validate` | 验证当前配置 |

## 🔥 热更新测试

1. **修改配置文件**: 编辑 `complete-gateway-config.yaml`
2. **观察日志**: 查看控制台输出的配置变更日志
3. **验证变化**: 访问 `http://localhost:8080/config` 查看配置变化
4. **手动重载**: 使用 `curl -X POST http://localhost:8080/admin/config/reload`

### 示例热更新操作

```bash
# 查看当前配置
curl http://localhost:8080/config | jq

# 修改配置文件中的debug值
# debug: false

# 手动重载配置
curl -X POST http://localhost:8080/admin/config/reload

# 验证配置变化
curl http://localhost:8080/config | jq '.data.complete_config.gateway.debug'
```

## 📊 监控和指标

### 访问监控指标
```bash
# 获取Prometheus格式指标
curl http://localhost:8080/metrics

# 检查服务健康状态
curl http://localhost:8080/health
```

### 主要指标说明

- `gateway_uptime_seconds`: 服务运行时间
- `gateway_requests_total`: 请求总数
- `gateway_config_enabled`: 各组件启用状态

## 🛠️ 配置说明

### 数据库配置
```yaml
database:
  type: "mysql"
  enabled: true
  mysql:
    host: "127.0.0.1"
    port: "3306"
    username: "root"
    password: "password123"
    db-name: "gateway_demo"
```

### Redis配置
```yaml
redis:
  addr: "127.0.0.1:6379"
  password: ""
  db: 1
  pool-size: 20
```

### 中间件配置
```yaml
cors:
  allowed-all-origins: true
  allowed-methods: ["GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"]

ratelimit:
  enabled: true
  max_requests: 1000
  window: "1m"
```

## 🔧 开发和调试

### 开发模式启动
```bash
# 启用调试模式
export APP_ENV=development

# 启用详细日志
export LOG_LEVEL=debug

# 启动服务
go run complete_gateway_demo_v2.go
```

### 常用调试命令
```bash
# 检查服务状态
curl -s http://localhost:8080/status | jq

# 验证配置
curl -s http://localhost:8080/admin/config/validate | jq

# 测试缓存功能
curl -s http://localhost:8080/api/cache/test | jq

# 测试数据库连接
curl -s http://localhost:8080/api/db/test | jq
```

## 🚨 常见问题

### Q: 服务启动失败？
**A**: 检查以下事项：
- Go版本是否1.19+
- 配置文件是否存在且格式正确
- 端口8080是否被占用
- MySQL/Redis服务是否正常（可选）

### Q: 配置热更新不生效？
**A**: 确认：
- 配置文件格式正确
- 配置管理器正常运行
- 查看日志中的错误信息

### Q: 数据库连接失败？
**A**: 演示服务默认包含数据库配置测试，但不要求实际的数据库服务。如果需要真实连接，请：
- 启动MySQL服务
- 创建对应的数据库
- 修改配置文件中的连接参数

## 📚 扩展阅读

- [go-config框架文档](../README.md)
- [配置热更新详解](../HOT_RELOAD_README.md)
- [使用指南](../USAGE.md)

## 🤝 贡献

欢迎提交Issue和Pull Request来改进这个演示项目！

## 📄 许可证

MIT License