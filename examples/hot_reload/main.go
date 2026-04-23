package main

import (
	"context"
	"fmt"
	"time"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/gateway"
)

// printCurrentConfig 打印当前配置的关键信息
func printCurrentConfig(config *gateway.Gateway) {
	fmt.Println("\n--- 当前配置 ---")
	fmt.Printf("服务名称: %s\n", config.Name)
	fmt.Printf("版本: %s\n", config.Version)
	fmt.Printf("环境: %s\n", config.Environment)
	fmt.Printf("调试模式: %v\n", config.Debug)
	if config.HTTPServer != nil {
		fmt.Printf("HTTP服务: %s:%d\n", config.HTTPServer.Host, config.HTTPServer.Port)
	}
	if config.Database != nil && config.Database.Relational != nil && config.Database.Relational.MySQL != nil {
		fmt.Printf("数据库: %s:%s/%s\n",
			config.Database.Relational.MySQL.Host,
			config.Database.Relational.MySQL.Port,
			config.Database.Relational.MySQL.Dbname)
	}
	fmt.Println("----------------")
}

// HotReload 演示配置热更新
func HotReload() {
	config := gateway.Default()

	// 配置热更新选项
	hotReloadConfig := &goconfig.HotReloadConfig{
		Enabled:         true,                   // 启用热更新
		WatchInterval:   500 * time.Millisecond, // 监控间隔
		DebounceDelay:   1 * time.Second,        // 防抖延迟
		MaxRetries:      3,                      // 最大重试次数
		CallbackTimeout: 30 * time.Second,       // 回调超时
		EnableEnvWatch:  true,                   // 监控环境变量
	}

	manager := goconfig.NewConfigBuilder(config).
		WithConfigPath("../resources/gateway-xl-dev.yaml").
		WithHotReload(hotReloadConfig).
		MustBuildAndStart()

	defer manager.Stop()

	// 打印初始配置
	fmt.Println("=== 配置热更新示例 ===")
	fmt.Printf("📂 监控文件: %s\n", "../resources/gateway-xl-dev.yaml")
	printCurrentConfig(config)

	// 注册配置变更回调
	manager.RegisterConfigCallback(
		func(ctx context.Context, event goconfig.CallbackEvent) error {
			fmt.Printf("\n🔄 检测到配置变更！\n")
			fmt.Printf("变更时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("事件类型: %v\n", event.Type)
			fmt.Printf("事件源: %s\n", event.Source)

			// 调试：打印事件值的类型
			fmt.Printf("NewValue 类型: %T\n", event.NewValue)
			fmt.Printf("OldValue 类型: %T\n", event.OldValue)

			// 从事件中获取新配置并打印
			if newConfig, ok := event.NewValue.(*gateway.Gateway); ok {
				fmt.Println("✅ 成功获取新配置")
				printCurrentConfig(newConfig)
			} else {
				fmt.Printf("⚠️  无法转换为 *gateway.Gateway，实际类型: %T\n", event.NewValue)
				// 尝试打印原始值
				fmt.Printf("NewValue 内容: %+v\n", event.NewValue)
			}

			return nil
		},
		goconfig.CallbackOptions{
			ID: "config-change-logger",
			Types: []goconfig.CallbackType{
				goconfig.CallbackTypeConfigChanged,
				goconfig.CallbackTypeReloaded,
				goconfig.CallbackTypeFileChanged,
			},
			Priority: 100,
			Async:    false,
			Timeout:  5 * time.Second,
		},
	)

	// 注册环境变更回调
	manager.RegisterEnvironmentCallback(
		"env-callback",
		func(oldEnv, newEnv goconfig.EnvironmentType) error {
			fmt.Printf("\n🌍 环境变更: %s -> %s\n", oldEnv, newEnv)
			return nil
		},
		100,   // 优先级
		false, // 是否异步
	)

	fmt.Println("\n✅ 热更新已启用！")
	fmt.Println("📝 请修改配置文件 '../resources/gateway-xl-dev.yaml' 查看效果")
	fmt.Println("💡 提示：修改 name, version, debug 等字段可以看到变化")
	fmt.Println("⏱️  程序将运行 60 秒...")

	// 每 5 秒打印一次提示
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	timeout := time.After(60 * time.Second)
	counter := 0

	for {
		select {
		case <-ticker.C:
			counter += 5
			fmt.Printf("⏰ 已运行 %d 秒，还剩 %d 秒...\n", counter, 60-counter)
		case <-timeout:
			fmt.Println("\n⏹️  示例结束")
			return
		}
	}
}

func main() {
	HotReload()
}
