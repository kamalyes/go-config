# go-config

[![Go Reference](https://pkg.go.dev/badge/github.com/kamalyes/go-config.svg)](https://pkg.go.dev/github.com/kamalyes/go-config)
[![Go Report Card](https://goreportcard.com/badge/github.com/kamalyes/go-config)](https://goreportcard.com/report/github.com/kamalyes/go-config)
[![Tests](https://github.com/kamalyes/go-config/actions/workflows/test.yml/badge.svg)](https://github.com/kamalyes/go-config/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

一个功能强大且易于使用的Go配置管理库，支持多种配置格式、智能发现、热更新和安全访问。为第三方开发者提供开箱即用的配置管理解决方案。

## ✨ 核心特性

- 🔧 **多格式支持** - 支持YAML、JSON、TOML等多种配置格式
- 🔥 **配置热更新** - 实时监听配置文件变化并自动重载
- 🛡️ **安全访问** - 防止空指针异常的链式配置访问
- 🎯 **智能发现** - 自动发现和加载配置文件（支持多环境）
- 🌍 **全球化环境支持** - 内置 56 种环境类型（9 种标准环境 + 47 个国家/地区），支持自定义环境注册
- 📦 **丰富模块** - 内置40+配置模块，覆盖常见应用场景
- 🚀 **零配置启动** - 开箱即用的默认配置
- 🎨 **链式API** - 优雅的构建器模式API设计

## 🌍 环境与配置文件发现

### 内置环境类型

#### 📋 标准环境（9 个）

| 环境类型 | 常量 | 支持的配置文件后缀 |
|---------|------|-------------------|
| 开发环境 | `EnvDevelopment` | `dev`, `develop`, `development` |
| 本地环境 | `EnvLocal` | `local`, `localhost` |
| 测试环境 | `EnvTest` | `test`, `testing`, `qa`, `sit` |
| 预发布环境 | `EnvStaging` | `staging`, `stage`, `stg`, `pre`, `preprod`, `pre-prod`, `fat`, `gray`, `grey`, `canary` |
| 生产环境 | `EnvProduction` | `prod`, `production`, `prd`, `release`, `live`, `online`, `master`, `main` |
| 调试环境 | `EnvDebug` | `debug`, `debugging`, `dbg` |
| 演示环境 | `EnvDemo` | `demo`, `demonstration`, `showcase`, `preview`, `sandbox` |
| UAT环境 | `EnvUAT` | `uat`, `acceptance`, `user-acceptance`, `beta` |
| 集成环境 | `EnvIntegration` | `integration`, `int`, `ci`, `integration-test`, `integ` |

#### 🌏 全球化环境支持（47 个国家/地区）

**🏯 亚洲（20 个）**

| 国家/地区 | 常量 | 别名 |
|----------|------|------|
| 中国 | `EnvChina` | `china`, `cn`, `chn` |
| 日本 | `EnvJapan` | `japan`, `jp`, `jpn` |
| 韩国 | `EnvKorea` | `korea`, `kr`, `kor`, `south-korea` |
| 印度 | `EnvIndia` | `india`, `in`, `ind` |
| 新加坡 | `EnvSingapore` | `singapore`, `sg`, `sgp` |
| 泰国 | `EnvThailand` | `thailand`, `th`, `tha`, `thai` |
| 越南 | `EnvVietnam` | `vietnam`, `vn`, `vnm`, `viet` |
| 马来西亚 | `EnvMalaysia` | `malaysia`, `my`, `mys` |
| 印度尼西亚 | `EnvIndonesia` | `indonesia`, `id`, `idn` |
| 菲律宾 | `EnvPhilippines` | `philippines`, `ph`, `phl` |
| 缅甸 | `EnvMyanmar` | `myanmar`, `mm`, `mmr`, `burma` |
| 老挝 | `EnvLaos` | `laos`, `la`, `lao` |
| 柬埔寨 | `EnvCambodia` | `cambodia`, `kh`, `khm` |
| 巴基斯坦 | `EnvPakistan` | `pakistan`, `pk`, `pak` |
| 孟加拉国 | `EnvBangladesh` | `bangladesh`, `bd`, `bgd` |
| 斯里兰卡 | `EnvSriLanka` | `srilanka`, `lk`, `lka`, `sri-lanka` |
| 尼泊尔 | `EnvNepal` | `nepal`, `np`, `npl` |
| 香港 | `EnvHongKong` | `hongkong`, `hk`, `hkg`, `hong-kong` |
| 台湾 | `EnvTaiwan` | `taiwan`, `tw`, `twn` |
| 澳门 | `EnvMacao` | `macao`, `mo`, `mac`, `macau` |

**🏰 欧洲（16 个）**

| 国家 | 常量 | 别名 |
|------|------|------|
| 英国 | `EnvUK` | `uk`, `gb`, `gbr`, `united-kingdom`, `britain`, `england` |
| 德国 | `EnvGermany` | `germany`, `de`, `deu`, `deutschland` |
| 法国 | `EnvFrance` | `france`, `fr`, `fra` |
| 意大利 | `EnvItaly` | `italy`, `it`, `ita`, `italia` |
| 西班牙 | `EnvSpain` | `spain`, `es`, `esp`, `espana` |
| 荷兰 | `EnvNetherlands` | `netherlands`, `nl`, `nld`, `holland` |
| 比利时 | `EnvBelgium` | `belgium`, `be`, `bel` |
| 瑞士 | `EnvSwitzerland` | `switzerland`, `ch`, `che` |
| 奥地利 | `EnvAustria` | `austria`, `at`, `aut` |
| 瑞典 | `EnvSweden` | `sweden`, `se`, `swe` |
| 挪威 | `EnvNorway` | `norway`, `no`, `nor` |
| 丹麦 | `EnvDenmark` | `denmark`, `dk`, `dnk` |
| 芬兰 | `EnvFinland` | `finland`, `fi`, `fin` |
| 波兰 | `EnvPoland` | `poland`, `pl`, `pol` |
| 俄罗斯 | `EnvRussia` | `russia`, `ru`, `rus` |
| 土耳其 | `EnvTurkey` | `turkey`, `tr`, `tur` |

**🗽 美洲（8 个）**

| 国家 | 常量 | 别名 |
|------|------|------|
| 美国 | `EnvUSA` | `usa`, `us`, `united-states`, `america` |
| 加拿大 | `EnvCanada` | `canada`, `ca`, `can` |
| 墨西哥 | `EnvMexico` | `mexico`, `mx`, `mex` |
| 巴西 | `EnvBrazil` | `brazil`, `br`, `bra`, `brasil` |
| 阿根廷 | `EnvArgentina` | `argentina`, `ar`, `arg` |
| 智利 | `EnvChile` | `chile`, `cl`, `chl` |
| 哥伦比亚 | `EnvColombia` | `colombia`, `co`, `col` |
| 秘鲁 | `EnvPeru` | `peru`, `pe`, `per` |

**🦘 其他地区（7 个）**

| 国家/地区 | 常量 | 别名 |
|----------|------|------|
| 澳大利亚 | `EnvAustralia` | `australia`, `au`, `aus` |
| 新西兰 | `EnvNewZealand` | `newzealand`, `nz`, `nzl`, `new-zealand` |
| 南非 | `EnvSouthAfrica` | `southafrica`, `za`, `zaf`, `south-africa` |
| 埃及 | `EnvEgypt` | `egypt`, `eg`, `egy` |
| 尼日利亚 | `EnvNigeria` | `nigeria`, `ng`, `nga` |
| 肯尼亚 | `EnvKenya` | `kenya`, `ke`, `ken` |
| 阿联酋 | `EnvUAE` | `uae`, `ae`, `are`, `emirates`, `dubai` |
| 沙特阿拉伯 | `EnvSaudiArabia` | `saudiarabia`, `sa`, `sau`, `saudi-arabia`, `saudi` |
| 以色列 | `EnvIsrael` | `israel`, `il`, `isr` |
| 卡塔尔 | `EnvQatar` | `qatar`, `qa`, `qat` |

**💡 使用示例：**

```go
// 设置中国环境
goconfig.SetCurrentEnvironment(goconfig.EnvChina)

// 判断是否为中国环境
if goconfig.IsEnvironment(goconfig.EnvChina) {
    // 使用中国特定配置
}

// 配置文件命名示例：
// gateway-xl-china.yaml
// gateway-xl-cn.yaml
// gateway-xl-chn.yaml
```

### 配置文件命名规则

配置文件命名格式：`{prefix}-{env-suffix}.{ext}`

**示例：**

```bash
# 标准环境
gateway-xl-dev.yaml          # 开发环境
gateway-xl-prod.yaml         # 生产环境
gateway-xl-staging.yaml      # 预发布环境

# 国家/地区环境
gateway-xl-china.yaml        # 中国环境
gateway-xl-cn.yaml           # 中国环境（别名）
gateway-xl-japan.yaml        # 日本环境
gateway-xl-usa.yaml          # 美国环境
gateway-xl-uk.yaml           # 英国环境
```

当 `APP_ENV=china` 时，会按优先级查找：
- `gateway-xl-china.yaml`
- `gateway-xl-china.yml`
- `gateway-xl-cn.yaml`
- `gateway-xl-chn.yaml`
- ...

### 注册自定义环境

如果内置的 56 种环境不满足需求，可以注册自定义环境：

```go
package main

import goconfig "github.com/kamalyes/go-config"

func init() {
    // 注册自定义环境 "custom"，支持后缀 "custom", "my-env", "myenv"
    // 配置文件可命名为: gateway-xl-custom.yaml, gateway-xl-my-env.yaml 等
    goconfig.RegisterEnvPrefixes("custom", "custom", "my-env", "myenv")
}
```

### 全球化部署示例

```go
package main

import (
    goconfig "github.com/kamalyes/go-config"
)

func main() {
    // 方式1：直接使用环境变量（推荐）
    // 设置环境变量：export APP_ENV=china
    // 或：export APP_ENV=cn
    // 或：export APP_ENV=usa
    // 配置管理器会自动识别并加载对应的配置文件
    
    manager := goconfig.NewConfigBuilder(config).
        WithConfigPrefix("gateway-xl").
        WithConfigPath("resources").
        MustBuildAndStart()
    
    defer manager.Stop()
    
    // 方式2：代码中动态设置（适用于特殊场景）
    // goconfig.SetCurrentEnvironment(goconfig.EnvChina)
    
    // 方式3：使用环境判断
    if goconfig.IsEnvironment(goconfig.EnvChina) {
        // 中国特定逻辑
        log.Info("使用中国区域配置")
    }
}
```

**部署配置示例：**

```bash
# 中国区域部署
export APP_ENV=china  # 或 cn, chn
./app

# 美国区域部署
export APP_ENV=usa    # 或 us
./app

# 日本区域部署
export APP_ENV=japan  # 或 jp, jpn
./app

# 欧洲区域部署
export APP_ENV=germany  # 或 de, deu
./app
```

### 配置文件未找到时的错误提示

当配置文件未找到时，会输出详细的诊断信息：

```
❌ 未找到前缀为 'gateway-xl' 的配置文件
📍 搜索路径: resources
🌍 当前环境: custom-env
⚠️ 当前环境 'custom-env' 未在 DefaultEnvPrefixes 中注册
📋 已注册的环境及其后缀:
   - development: [dev develop development]
   - local: [local localhost]
   ...

💡 如需注册自定义环境，请在程序启动前注册:

   示例代码:
   func init() {
       goconfig.RegisterEnvPrefixes("custom-env", "custom-env", "custom-alias")
   }
```

## 🚀 快速开始

### 安装

```bash
go get github.com/kamalyes/go-config
```

### 基础使用 - 配置热更新

```go
package main

import (
    "fmt"
    "time"
    
    goconfig "github.com/kamalyes/go-config"
    "github.com/kamalyes/go-config/pkg/gateway"
)

func main() {
    // 初始化HTTPServer配置
    config := gateway.DefaultHTTPServer()
    
    // 配置热更新回调
    hotReloadConfig := &goconfig.HotReloadConfig{
        Enabled: true,
        OnReloaded: func(oldConfig, newConfig interface{}) {
            fmt.Printf("配置已更新: %+v -> %+v\n", oldConfig, newConfig)
        },
        OnError: func(err error) {
            fmt.Printf("热更新错误: %v\n", err)
        },
    }
    
    // 创建并启动配置管理器
    manager := goconfig.NewConfigBuilder(config).
        WithConfigPath("config.yaml").
        WithEnvironment(goconfig.EnvDevelopment).
        WithHotReload(hotReloadConfig).
        MustBuildAndStart()
    
    defer manager.Stop()
    
    // 使用安全配置访问
    safeConfig := goconfig.SafeConfig(config)
    
    fmt.Printf("HTTP服务器启动在 %s:%d\n", 
        safeConfig.Host("localhost"), 
        safeConfig.Port(8080))
    
    fmt.Printf("启用HTTP: %v\n", 
        safeConfig.Field("EnableHttp").Bool(true))
    
    // 保持程序运行以观察热更新
    select {
    case <-time.After(time.Minute * 5):
        fmt.Println("程序退出")
    }
}
```

### 创建配置文件 `config.yaml`

```yaml
# HTTP服务器配置 - 注意字段名使用横线格式
module-name: "my-app-server"
host: "0.0.0.0" 
port: 8080
grpc-port: 9090
read-timeout: 30
write-timeout: 30
idle-timeout: 60
max-header-bytes: 1048576
enable-http: true
enable-grpc: false
enable-tls: false
enable-gzip-compress: true
tls:
  cert-file: ""
  key-file: ""
  ca-file: ""
headers:
  x-custom-header: "my-app"
  x-version: "1.0.0"
```

现在修改配置文件，程序会自动检测变化并重载配置！

## 🌟 环境管理功能

- ✅ **自动环境初始化** - 包导入时自动初始化，无需手动调用
- 🎯 **便捷判断函数** - 提供 `IsDev()`, `IsProduction()` 等直观的环境判断函数
- 📊 **环境级别管理** - 按重要程度对环境进行分级管理
- 🔄 **环境变更监听** - 支持环境变更回调机制
- 🛠️ **自定义环境注册** - 灵活注册自定义环境类型

### 快速使用环境判断

```go
import goconfig "github.com/kamalyes/go-config"

func main() {
    // 无需手动初始化，直接使用
    if goconfig.IsDev() {
        log.SetLevel(log.DebugLevel)
    } else if goconfig.IsProduction() {
        log.SetLevel(log.WarnLevel)
    }
    
    // 环境级别判断
    if goconfig.IsProductionLevel() {
        // 启用生产级别的监控和安全功能
        enableProductionFeatures()
    }
}
```

**📖 详细使用说明请参考：[环境管理器使用文档](ENV_USAGE.md)**

## 🎯 支持的配置模块

| 类别 | 模块 | 描述 |
|------|------|------|
| **网关服务** | Gateway, HTTP, GRPC | 网关和服务配置 |
| **数据存储** | MySQL, PostgreSQL, SQLite, Redis | 数据库配置 |
| **中间件** | CORS, 限流, JWT, 恢复 | 常用中间件配置 |
| **监控运维** | Health, Metrics, Prometheus, Jaeger | 监控和链路追踪 |
| **消息队列** | Kafka, MQTT | 消息系统配置 |
| **第三方服务** | 支付宝, 微信支付, 阿里云短信 | 第三方集成 |

## 📜 许可证

本项目采用 [MIT 许可证](LICENSE)

---

**如果这个项目对你有帮助，请给我们一个 ⭐️**
