# go-config 示例程序

这个目录包含了 go-config 的各种使用示例，每个示例都是可以独立运行的完整程序。

## 📋 示例列表

### 1. 基础使用示例 (`basic_usage/`)
**文件**: `basic_usage/main.go`

演示 go-config 的基本使用方法，包括：
- 创建配置文件
- 加载单一配置
- 访问各种配置项
- 基础配置验证

**运行方式**:
```bash
cd examples/basic_usage
go mod init basic-usage-demo
go get github.com/kamalyes/go-config
go run main.go
```

**预期输出**: 显示各种配置项的值和状态

---

### 2. 多配置实例示例 (`multi_config/`)
**文件**: `multi_config/main.go`

演示多实例配置管理，包括：
- 多数据库配置
- 多Redis实例
- 多服务器配置
- 按模块名称获取特定配置

**运行方式**:
```bash
cd examples/multi_config
go mod init multi-config-demo
go get github.com/kamalyes/go-config
go run main.go
```

**特点**:
- 展示如何管理多个同类型服务实例
- 演示根据模块名获取特定配置
- 适用于微服务或分布式架构

---

### 3. 环境变量和自定义配置 (`environment_config/`)
**文件**: `environment_config/main.go`

演示环境管理和自定义配置，包括：
- 多环境配置文件 (dev/sit/prod)
- 环境变量优先级
- 自定义配置路径和选项
- 动态环境切换

**运行方式**:
```bash
cd examples/environment_config
go mod init environment-config-demo
go get github.com/kamalyes/go-config
go run main.go
```

**特点**:
- 展示不同环境的配置差异
- 演示环境变量优先级控制
- 适用于多环境部署场景

---

### 4. Web服务应用示例 (`web_service/`)
**文件**: `web_service/main.go`

演示在实际Web服务中的使用，包括：
- HTTP服务器配置
- 数据库连接初始化
- Redis缓存配置
- RESTful API端点
- 优雅关闭

**运行方式**:
```bash
cd examples/web_service
go mod init web-service-demo
go get github.com/kamalyes/go-config
go run main.go
```

**API端点**:
- `GET /api/v1/health` - 健康检查
- `GET /api/v1/config` - 配置信息
- `POST /api/v1/login` - 模拟登录
- `GET /api/v1/profile` - 用户资料

**测试方式**:
```bash
# 健康检查
curl http://localhost:8080/api/v1/health

# 查看配置
curl http://localhost:8080/api/v1/config

# 模拟登录
curl -X POST http://localhost:8080/api/v1/login

# 获取资料
curl http://localhost:8080/api/v1/profile
```

---

### 5. 配置验证和错误处理 (`validation_demo/`)
**文件**: `validation_demo/main.go`

演示配置验证和错误处理，包括：
- 正确配置验证
- 配置错误检测
- 错误处理策略
- 配置缺失处理

**运行方式**:
```bash
cd examples/validation_demo
go mod init validation-demo
go get github.com/kamalyes/go-config
go run main.go
```

**特点**:
- 展示各种配置错误场景
- 演示错误处理最佳实践
- 适用于生产环境配置验证

---

## 🚀 快速开始

### 前置条件
- Go 1.20+
- 网络连接（用于下载依赖）

### 通用运行步骤

1. **进入示例目录**:
   ```bash
   cd examples/[示例名称]
   ```

2. **初始化Go模块**:
   ```bash
   go mod init [示例名称]-demo
   ```

3. **安装依赖**:
   ```bash
   go get github.com/kamalyes/go-config
   ```

4. **运行示例**:
   ```bash
   go run main.go
   ```

### 批量运行所有示例

创建一个脚本来运行所有示例：

```bash
#!/bin/bash
# run_all_examples.sh

examples=("basic_usage" "multi_config" "environment_config" "web_service" "validation_demo")

for example in "${examples[@]}"; do
    echo "🚀 运行示例: $example"
    echo "=========================="
    
    cd examples/$example
    go mod init ${example}-demo
    go get github.com/kamalyes/go-config
    go run main.go
    
    echo ""
    echo "✅ $example 示例完成"
    echo ""
    
    cd ../..
done
```

## 📖 配置文件说明

每个示例都会自动创建所需的配置文件，通常包括：

- `resources/dev_config.yaml` - 开发环境配置
- `resources/sit_config.yaml` - 测试环境配置  
- `resources/prod_config.yaml` - 生产环境配置

配置文件会在程序运行结束时自动清理。

## 🔧 自定义运行

### 指定环境变量
```bash
export APP_ENV=prod
go run main.go
```

### 修改配置
你可以修改示例代码中的配置内容来测试不同的场景。

### 持久化配置
如果想保留生成的配置文件，可以注释掉 `cleanup()` 函数调用。

## 📚 学习路径

建议按以下顺序学习示例：

1. **basic_usage** - 了解基本用法
2. **environment_config** - 理解环境管理
3. **multi_config** - 学习多实例配置
4. **validation_demo** - 掌握错误处理
5. **web_service** - 应用到实际项目

## ❓ 常见问题

### Q: 为什么示例运行后没有看到配置文件？
A: 示例程序会在结束时自动清理生成的文件。如需保留，请注释掉 `cleanup()` 函数。

### Q: 如何修改示例中的配置？
A: 修改 `createXXXConfig()` 函数中的配置内容字符串即可。

### Q: 示例可以用于生产环境吗？
A: 这些示例主要用于学习和测试。生产环境请根据实际需求进行配置优化。

### Q: 如何集成到现有项目？
A: 参考 `web_service` 示例，它展示了完整的项目集成方法。

## 🤝 贡献

欢迎提交新的示例或改进现有示例！请确保：
- 代码有充分的注释
- 示例能够独立运行
- 包含清理逻辑
- 添加到此README中

## 📄 许可证

这些示例代码遵循与主项目相同的 MIT 许可证。