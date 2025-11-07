/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-08 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-08 15:00:00
 * @FilePath: \go-config\examples\basic_usage\main.go
 * @Description: go-config 基础使用示例
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	goconfig "github.com/kamalyes/go-config"
)

func main() {
	fmt.Println("🚀 go-config 基础使用示例")
	fmt.Println("========================")

	// 创建配置文件
	if err := createExampleConfig(); err != nil {
		log.Fatalf("❌ 创建配置文件失败: %v", err)
	}

	// 基础使用示例
	basicExample()

	// 清理测试文件
	cleanup()
}

// createExampleConfig 创建示例配置文件
func createExampleConfig() error {
	configContent := `# 服务配置
server:
  addr: '0.0.0.0:8080'
  server-name: 'example-service'
  context-path: '/api/v1'
  handle-method-not-allowed: true
  data-driver: 'mysql'

# MySQL 配置
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
  conn-max-idle-time: 60
  conn-max-life-time: 600

# Redis 配置
redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 0
  pool-size: 100
  min-idle-conns: 5

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

# JWT 配置
jwt:
  signing-key: 'example-jwt-secret-key-123456'
  expires-time: 604800
  buffer-time: 86400
  use-multipoint: false

# 跨域配置
cors:
  allowed-all-origins: true
  allow-credentials: true
  max-age: "86400"
`

	// 确保目录存在
	resourcesDir := "./resources"
	if err := os.MkdirAll(resourcesDir, os.ModePerm); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	// 写入配置文件
	configFile := filepath.Join(resourcesDir, "dev_config.yaml")
	if err := os.WriteFile(configFile, []byte(configContent), 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %w", err)
	}

	fmt.Printf("✅ 配置文件已创建: %s\n", configFile)
	return nil
}

// basicExample 基础使用示例
func basicExample() {
	fmt.Println("\n📋 基础配置加载示例")
	fmt.Println("-------------------")

	ctx := context.Background()

	// 创建单配置管理器
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Fatalf("❌ 创建配置管理器失败: %v", err)
	}

	// 获取配置
	config := manager.GetConfig()

	// 打印服务器配置
	fmt.Printf("🌐 服务器配置:\n")
	fmt.Printf("   - 地址: %s\n", config.Server.Addr)
	fmt.Printf("   - 服务名: %s\n", config.Server.ServerName)
	fmt.Printf("   - 上下文路径: %s\n", config.Server.ContextPath)
	fmt.Printf("   - 数据驱动: %s\n", config.Server.DataDriver)

	// 打印数据库配置
	fmt.Printf("\n💾 MySQL 配置:\n")
	fmt.Printf("   - 地址: %s:%s\n", config.MySQL.Host, config.MySQL.Port)
	fmt.Printf("   - 数据库: %s\n", config.MySQL.Dbname)
	fmt.Printf("   - 用户名: %s\n", config.MySQL.Username)
	fmt.Printf("   - 最大连接数: %d\n", config.MySQL.MaxOpenConns)
	fmt.Printf("   - 最大空闲连接: %d\n", config.MySQL.MaxIdleConns)

	// 打印Redis配置
	fmt.Printf("\n⚡ Redis 配置:\n")
	fmt.Printf("   - 地址: %s\n", config.Redis.Addr)
	fmt.Printf("   - 数据库索引: %d\n", config.Redis.DB)
	fmt.Printf("   - 连接池大小: %d\n", config.Redis.PoolSize)

	// 打印日志配置
	fmt.Printf("\n📋 日志配置:\n")
	fmt.Printf("   - 日志级别: %s\n", config.Zap.Level)
	fmt.Printf("   - 输出格式: %s\n", config.Zap.Format)
	fmt.Printf("   - 前缀: %s\n", config.Zap.Prefix)
	fmt.Printf("   - 开发模式: %t\n", config.Zap.Development)

	// 打印JWT配置
	fmt.Printf("\n🔐 JWT 配置:\n")
	fmt.Printf("   - 过期时间: %d 秒\n", config.JWT.ExpiresTime)
	fmt.Printf("   - 缓冲时间: %d 秒\n", config.JWT.BufferTime)
	fmt.Printf("   - 多点登录拦截: %t\n", config.JWT.UseMultipoint)

	// 打印CORS配置
	fmt.Printf("\n🌍 CORS 配置:\n")
	fmt.Printf("   - 允许所有来源: %t\n", config.Cors.AllowedAllOrigins)
	fmt.Printf("   - 允许凭证: %t\n", config.Cors.AllowCredentials)
	fmt.Printf("   - 最大缓存时间: %s\n", config.Cors.MaxAge)

	fmt.Println("\n✅ 基础配置加载完成!")
}

// cleanup 清理测试文件
func cleanup() {
	if err := os.RemoveAll("./resources"); err != nil {
		log.Printf("⚠️ 清理文件失败: %v", err)
	}
	if err := os.RemoveAll("./logs"); err != nil {
		log.Printf("⚠️ 清理日志文件失败: %v", err)
	}
	fmt.Println("\n🧹 清理完成")
}