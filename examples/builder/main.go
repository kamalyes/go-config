package main

import (
	"fmt"
	"time"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/gateway"
)

// BuilderPatterns 演示构建器模式的各种用法
func BuilderPatterns() {
	config := gateway.Default()
	manager := goconfig.NewConfigBuilder(config).
		// 配置文件路径（直接指定）
		WithConfigPath("../resources/gateway-xl-dev.yaml").
		// 或使用自动发现
		// WithSearchPath("../resources").
		// WithPrefix("gateway-xl").
		// 或使用模式匹配
		// WithPattern("*-config").
		// 环境配置
		WithEnvironment(goconfig.EnvDevelopment).
		// 热更新配置
		WithHotReload(&goconfig.HotReloadConfig{
			Enabled: true,
		}).
		// 上下文配置
		WithContext(&goconfig.ContextKeyOptions{
			Key:   "APP_ENV",
			Value: goconfig.EnvDevelopment,
		}).
		// 构建并启动
		MustBuildAndStart()

	defer manager.Stop()

	fmt.Println("配置管理器已启动")
	time.Sleep(2 * time.Second)
}

func main() {
	BuilderPatterns()
}
