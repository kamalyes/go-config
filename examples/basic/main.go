package main

import (
	"fmt"
	"time"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/gateway"
)

// Basic 演示基础配置使用
func Basic() {
	// 1. 初始化配置结构
	config := gateway.Default()

	// 2. 创建配置管理器（自动发现配置文件）
	manager := goconfig.NewConfigBuilder(config).
		WithPrefix("gateway-xl").                 // 配置文件前缀
		WithSearchPath("../resources").           // 搜索路径
		WithEnvironment(goconfig.EnvDevelopment). // 环境类型
		WithHotReload(&goconfig.HotReloadConfig{  // 热更新配置
			Enabled:       true,
			DebounceDelay: 1 * time.Second,
		}).
		MustBuildAndStart()

	defer manager.Stop()

	// 3. 使用配置
	fmt.Printf("服务启动在 %s:%d\n",
		config.HTTPServer.Host,
		config.HTTPServer.Port)
}

func main() {
	Basic()
}
