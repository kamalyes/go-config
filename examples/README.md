# go-config 示例

本目录包含 go-config 的各种使用示例

## 📁 示例列表

### basic - 基础使用
演示最基本的配置加载和使用方式

```bash
cd basic && go run main.go
```

### hot_reload - 配置热更新
演示如何启用配置热更新，以及如何注册配置变更回调

```bash
cd hot_reload && go run main.go
```

**测试热更新：**

程序运行后，在另一个终端修改配置文件 `../resources/gateway-xl-dev.yaml`，观察配置自动重新加载。

```bash
# 使用测试脚本（Bash）
bash test_hot_reload.sh

# 或手动修改配置文件
# 修改 name, version, debug 等字段，保存后观察输出
```

**注意事项：**
- 热更新监控配置文件所在的目录（而不是文件本身）
- 支持编辑器的保存操作（包括临时文件和重命名）
- 使用防抖机制（1秒延迟）避免频繁重载
- 配置文件路径会自动转换为绝对路径

### error_handling - 错误处理
演示如何创建自定义错误处理器，以及如何根据错误严重程度采取不同措施

```bash
cd error_handling && go run main.go
```

### environment - 环境管理
演示如何使用环境判断函数，以及如何在不同环境下加载不同配置

```bash
cd environment && go run main.go
```

### context - 上下文集成
演示如何将配置信息注入到 context.Context，以及如何从上下文中获取配置

```bash
cd context && go run main.go
```

### discovery - 配置发现
演示配置文件的自动发现、扫描和创建功能

```bash
cd discovery && go run main.go
```

### builder - 构建器模式
演示配置构建器的各种链式 API 用法

```bash
cd builder && go run main.go
```

## 🚀 运行所有示例

在 `examples` 目录下运行：

**Linux/Mac:**
```bash
for dir in */; do
    if [ -f "$dir/main.go" ]; then
        echo "Running $dir..."
        (cd "$dir" && go run main.go)
        echo ""
    fi
done
```

**或者使用提供的脚本:**
```bash
# PowerShell
run_all.ps1

# CMD
run_all.bat

# Linux/Mac
chmod +x run_all.sh
run_all.sh
```

## 📝 配置文件

所有示例共享 `resources/` 目录下的配置文件：
- `gateway-xl-dev.yaml` - 开发环境配置

## 💡 提示

1. 确保在运行示例前已经安装了所有依赖：
   ```bash
   cd go-config
   go mod download
   ```

2. 某些示例需要配置文件存在，请确保 `resources/` 目录下有相应的配置文件

3. 热更新示例会持续运行 60 秒，你可以在此期间修改配置文件查看效果

4. 运行示例时请确保在 `examples` 目录下：
   ```bash
   # 从项目根目录
   cd go-config/examples
   
   # 运行单个示例
   cd basic && go run main.go
   ```
