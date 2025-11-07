/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 12:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 12:00:00
 * @FilePath: \go-config\examples\basic_single_config\main.go
 * @Description: 基础单配置示例 - 演示如何使用go-config进行基本的配置管理
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/env"
)

func main() {
	fmt.Println("🚀 go-config 基础单配置示例")
	fmt.Println("=============================")

	// 创建配置目录
	if err := createConfigDir(); err != nil {
		log.Fatalf("创建配置目录失败: %v", err)
	}

	// 创建示例配置文件
	if err := createExampleConfig(); err != nil {
		log.Fatalf("创建配置文件失败: %v", err)
	}

	// 示例1: 使用默认配置
	fmt.Println("\n📋 示例1: 使用默认配置")
	example1DefaultConfig()

	// 示例2: 使用自定义配置选项
	fmt.Println("\n⚙️ 示例2: 使用自定义配置选项")
	example2CustomOptions()

	// 示例3: 环境变量切换
	fmt.Println("\n🌍 示例3: 环境变量切换")
	example3EnvironmentSwitch()

	fmt.Println("\n✅ 所有示例运行完成!")
}

// createConfigDir 创建配置目录
func createConfigDir() error {
	return os.MkdirAll("resources", 0755)
}

// createExampleConfig 创建示例配置文件
func createExampleConfig() error {
	configContent := `# 服务配置
server:
  addr: '0.0.0.0:8080'
  service-name: 'example-api'
  context-path: '/api/v1'
  handle-method-not-allowed: true
  data-driver: 'mysql'

# 数据库配置
mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'root'
  password: 'password123'
  db-name: 'example_db'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'
  max-idle-conns: 10
  max-open-conns: 100
  conn-max-idle-time: 30
  conn-max-life-time: 300

# Redis 配置
redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 0
  pool-size: 100
  min-idle-conns: 5
  max-retries: 3

# 日志配置
zap:
  level: 'info'
  format: 'console'
  prefix: '[EXAMPLE]'
  director: 'logs'
  link-name: 'logs/example.log'
  show-line: true
  encode-level: 'LowercaseColorLevelEncoder'
  log-in-console: true
  development: true

# CORS 配置
cors:
  allowed-all-origins: false
  allowed-origins:
    - "http://localhost:3000"
    - "http://localhost:8080"
  allowed-methods:
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
  allowed-headers:
    - "Authorization"
    - "Content-Type"
  allow-credentials: true
  max-age: "86400"

# JWT 配置
jwt:
  signing-key: 'example-secret-key-123'
  expires-time: 604800
  buffer-time: 86400
  use-multipoint: true
`

	configFiles := []string{
		"resources/dev_config.yaml",
		"resources/prod_config.yaml",
		"resources/test_config.yaml",
	}

	for _, file := range configFiles {
		if err := os.WriteFile(file, []byte(configContent), 0644); err != nil {
			return fmt.Errorf("创建配置文件 %s 失败: %v", file, err)
		}
	}

	return nil
}

// example1DefaultConfig 示例1: 使用默认配置
func example1DefaultConfig() {
	ctx := context.Background()

	// 创建单配置管理器 (使用默认选项)
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 创建配置管理器失败: %v", err)
		return
	}

	// 获取配置
	config := manager.GetConfig()

	// 打印配置信息
	fmt.Printf("  ✅ 服务地址: %s\n", config.Server.Addr)
	fmt.Printf("  ✅ 服务名称: %s\n", config.Server.ServerName)
	fmt.Printf("  ✅ 上下文路径: %s\n", config.Server.ContextPath)
	fmt.Printf("  ✅ MySQL: %s:%s/%s (用户: %s)\n", 
		config.MySQL.Host, config.MySQL.Port, config.MySQL.Dbname, config.MySQL.Username)
	fmt.Printf("  ✅ Redis: %s (DB:%d, 连接池:%d)\n", 
		config.Redis.Addr, config.Redis.DB, config.Redis.PoolSize)
	fmt.Printf("  ✅ 日志级别: %s, 格式: %s\n", 
		config.Zap.Level, config.Zap.Format)
	fmt.Printf("  ✅ JWT密钥: %s (过期时间: %d秒)\n", 
		maskString(config.JWT.SigningKey), config.JWT.ExpiresTime)
}

// example2CustomOptions 示例2: 使用自定义配置选项
func example2CustomOptions() {
	ctx := context.Background()

	// 自定义配置选项
	options := &goconfig.ConfigOptions{
		ConfigType:    "yaml",
		ConfigPath:    "./resources",
		ConfigSuffix:  "_config",
		EnvValue:      env.Prod, // 使用生产环境配置
		EnvContextKey: env.ContextKey("CUSTOM_ENV"),
		UseEnvLevel:   goconfig.EnvLevelCtx,
	}

	// 创建配置管理器
	manager, err := goconfig.NewSingleConfigManager(ctx, options)
	if err != nil {
		log.Printf("❌ 创建自定义配置管理器失败: %v", err)
		return
	}

	config := manager.GetConfig()

	fmt.Printf("  ✅ 当前环境: %s\n", options.EnvValue)
	fmt.Printf("  ✅ 配置路径: %s\n", options.ConfigPath)
	fmt.Printf("  ✅ 配置类型: %s\n", options.ConfigType)
	fmt.Printf("  ✅ 服务配置: %s (%s)\n", config.Server.ServerName, config.Server.Addr)
}

// example3EnvironmentSwitch 示例3: 环境变量切换
func example3EnvironmentSwitch() {
	environments := []env.EnvironmentType{env.Dev, env.Prod}

	for _, envType := range environments {
		fmt.Printf("\n  🌍 切换到环境: %s\n", envType)

		// 设置环境变量
		env.SetContextKey(&env.ContextKeyOptions{
			Key:   env.ContextKey("DEMO_ENV"),
			Value: envType,
		})

		ctx := context.Background()
		options := &goconfig.ConfigOptions{
			EnvValue:      envType,
			EnvContextKey: env.ContextKey("DEMO_ENV"),
			UseEnvLevel:   goconfig.EnvLevelCtx,
		}

		manager, err := goconfig.NewSingleConfigManager(ctx, options)
		if err != nil {
			log.Printf("    ❌ 环境 %s 配置加载失败: %v", envType, err)
			continue
		}

		config := manager.GetConfig()
		fmt.Printf("    ✅ 当前环境: %s\n", envType)
		fmt.Printf("    ✅ 服务: %s:%s\n", config.Server.ServerName, config.Server.Addr)
		fmt.Printf("    ✅ 数据库: %s:%s\n", config.MySQL.Host, config.MySQL.Port)
	}
}

// maskString 遮蔽敏感信息
func maskString(s string) string {
	if len(s) <= 8 {
		return "****"
	}
	return s[:4] + "****" + s[len(s)-4:]
}