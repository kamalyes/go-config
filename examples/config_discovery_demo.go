/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 16:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 16:30:00
 * @FilePath: \go-config\examples\config_discovery_demo.go
 * @Description: 配置文件自动发现演示
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package main

import (
	"fmt"
	"log"
	"os"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/gateway"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 设置环境
	env := goconfig.EnvDevelopment
	if envVar := os.Getenv("APP_ENV"); envVar != "" {
		env = goconfig.EnvironmentType(envVar)
	}

	// 演示不同的配置发现方式
	demonstrateConfigDiscovery(env)
}

func demonstrateConfigDiscovery(env goconfig.EnvironmentType) {
	fmt.Println("🔍 配置文件自动发现演示")
	fmt.Println("=================================")

	// 1. 扫描当前目录
	fmt.Println("\n1️⃣ 扫描当前目录中的配置文件:")
	currentDir, _ := os.Getwd()
	_, err := goconfig.ScanAndDisplayConfigs(currentDir, env)
	if err != nil {
		log.Printf("扫描失败: %v", err)
	}

	// 2. 自动发现最佳配置文件
	fmt.Println("\n2️⃣ 自动发现最佳配置文件:")
	bestConfig, err := goconfig.FindBestConfig(currentDir, env)
	if err != nil {
		log.Printf("发现失败: %v", err)

		// 3. 自动创建配置文件
		fmt.Println("\n3️⃣ 自动创建默认配置文件:")
		createdConfig, createErr := goconfig.AutoCreateConfig(currentDir, env, "gateway")
		if createErr != nil {
			log.Printf("创建失败: %v", createErr)
		} else {
			fmt.Printf("✅ 已创建: %s\n", createdConfig.Path)
			bestConfig = createdConfig
		}
	} else {
		fmt.Printf("✅ 找到最佳配置: %s (优先级: %d)\n", bestConfig.Path, bestConfig.Priority)
	}

	if bestConfig == nil {
		log.Println("❌ 无法获取配置文件")
		return
	}

	// 4. 使用发现的配置文件创建管理器
	fmt.Println("\n4️⃣ 使用自动发现创建配置管理器:")

	config := &gateway.Gateway{}
	manager, err := goconfig.CreateIntegratedManagerWithAutoDiscovery(
		config,
		currentDir,
		env,
		"gateway",
	)
	if err != nil {
		log.Printf("创建管理器失败: %v", err)
		return
	}
	defer manager.Stop()

	fmt.Printf("✅ 配置管理器创建成功\n")

	// 5. 显示配置信息
	fmt.Println("\n5️⃣ 配置信息:")
	gatewayConfig := manager.GetConfig().(*gateway.Gateway)
	fmt.Printf("   📌 服务名称: %s\n", gatewayConfig.Name)
	fmt.Printf("   🔢 版本: %s\n", gatewayConfig.Version)
	fmt.Printf("   🌍 环境: %s\n", gatewayConfig.Environment)

	if gatewayConfig.HTTPServer != nil {
		fmt.Printf("   🌐 HTTP服务器: %s\n", gatewayConfig.HTTPServer.GetEndpoint())
	}

	// 6. 演示模式匹配查找
	fmt.Println("\n6️⃣ 模式匹配查找示例:")
	patterns := []string{"gateway", "config", "app"}
	for _, pattern := range patterns {
		matchedConfigs, matchErr := goconfig.GetGlobalConfigDiscovery().FindConfigFileByPattern(currentDir, pattern, env)
		if matchErr == nil && len(matchedConfigs) > 0 {
			fmt.Printf("   🔍 模式 '%s': 找到 %d 个文件\n", pattern, len(matchedConfigs))
			for i, matchedConfig := range matchedConfigs {
				if i < 2 { // 只显示前2个
					fmt.Printf("      - %s\n", matchedConfig.Name)
				}
			}
		} else {
			fmt.Printf("   🔍 模式 '%s': 未找到匹配文件\n", pattern)
		}
	}

	// 7. 显示支持的文件类型
	fmt.Println("\n7️⃣ 支持的配置文件类型:")
	discovery := goconfig.GetGlobalConfigDiscovery()
	fmt.Printf("   📄 扩展名: %v\n", discovery.SupportedExtensions)
	fmt.Printf("   📝 默认名称: %v\n", discovery.DefaultNames)
	fmt.Printf("   🌍 环境前缀: %v\n", discovery.EnvPrefixes[env])

	fmt.Println("\n✅ 演示完成!")
}
