/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-08 15:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-08 15:30:00
 * @FilePath: \go-config\examples\multi_config\main.go
 * @Description: go-config 多配置实例示例 - 演示如何管理多个数据库和Redis实例
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
	fmt.Println("🚀 go-config 多配置实例示例")
	fmt.Println("============================")

	// 创建多实例配置文件
	if err := createMultiConfig(); err != nil {
		log.Fatalf("❌ 创建配置文件失败: %v", err)
	}

	// 多配置示例
	multiConfigExample()

	// 配置管理示例
	configManagementExample()

	// 清理
	cleanup()
}

// createMultiConfig 创建多实例配置文件
func createMultiConfig() error {
	configContent := `# 多服务器配置
server:
  - modulename: "api"
    addr: '0.0.0.0:8080'
    server-name: 'api-server'
    context-path: '/api/v1'
    data-driver: 'mysql'
  
  - modulename: "admin"
    addr: '0.0.0.0:8081'
    server-name: 'admin-server'
    context-path: '/admin/v1'
    data-driver: 'mysql'

# 多 MySQL 数据库配置
mysql:
  - modulename: "primary"
    host: '192.168.1.100'
    port: '3306'
    username: 'root'
    password: 'primary_pass'
    db-name: 'primary_db'
    config: 'charset=utf8mb4&parseTime=True&loc=Local'
    log-level: 'info'
    max-idle-conns: 20
    max-open-conns: 200
    conn-max-idle-time: 60
    conn-max-life-time: 600
  
  - modulename: "replica"
    host: '192.168.1.101'
    port: '3306'
    username: 'readonly'
    password: 'replica_pass'
    db-name: 'primary_db'
    config: 'charset=utf8mb4&parseTime=True&loc=Local'
    log-level: 'warn'
    max-idle-conns: 10
    max-open-conns: 100
    conn-max-idle-time: 60
    conn-max-life-time: 600
  
  - modulename: "analytics"
    host: '192.168.1.102'
    port: '3306'
    username: 'analytics'
    password: 'analytics_pass'
    db-name: 'analytics_db'
    config: 'charset=utf8mb4&parseTime=True&loc=Local'
    log-level: 'error'
    max-idle-conns: 5
    max-open-conns: 50
    conn-max-idle-time: 120
    conn-max-life-time: 1200

# 多 Redis 实例配置
redis:
  - modulename: "cache"
    addr: '192.168.1.200:6379'
    password: 'cache_redis_pass'
    db: 0
    pool-size: 100
    min-idle-conns: 10
  
  - modulename: "session"
    addr: '192.168.1.201:6379'
    password: 'session_redis_pass'
    db: 1
    pool-size: 50
    min-idle-conns: 5
  
  - modulename: "queue"
    addr: '192.168.1.202:6379'
    password: 'queue_redis_pass'
    db: 0
    pool-size: 30
    min-idle-conns: 3

# 多日志配置
zap:
  - modulename: "api"
    level: 'info'
    format: 'json'
    prefix: '[API]'
    director: 'logs/api'
    link-name: 'logs/api/api.log'
    development: false
  
  - modulename: "admin"
    level: 'debug'
    format: 'console'
    prefix: '[ADMIN]'
    director: 'logs/admin'
    link-name: 'logs/admin/admin.log'
    development: true
    show-line: true
    log-in-console: true

# JWT 配置
jwt:
  - modulename: "api"
    signing-key: 'api-jwt-secret-key-123456'
    expires-time: 604800
    buffer-time: 86400
    use-multipoint: true
  
  - modulename: "admin"
    signing-key: 'admin-jwt-secret-key-789012'
    expires-time: 28800
    buffer-time: 3600
    use-multipoint: false
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

	fmt.Printf("✅ 多配置文件已创建: %s\n", configFile)
	return nil
}

// multiConfigExample 多配置示例
func multiConfigExample() {
	fmt.Println("\n📊 多配置实例管理示例")
	fmt.Println("---------------------")

	ctx := context.Background()

	// 创建多配置管理器
	manager, err := goconfig.NewMultiConfigManager(ctx, nil)
	if err != nil {
		log.Fatalf("❌ 创建多配置管理器失败: %v", err)
	}

	// 获取多配置
	multiConfig := manager.GetConfig()

	// 显示服务器配置
	fmt.Printf("🌐 服务器实例 (%d个):\n", len(multiConfig.Server))
	for i, server := range multiConfig.Server {
		fmt.Printf("   [%d] %s - %s (%s)\n", 
			i+1, server.ModuleName, server.Addr, server.ServerName)
	}

	// 显示数据库配置
	fmt.Printf("\n💾 MySQL 实例 (%d个):\n", len(multiConfig.MySQL))
	for i, mysql := range multiConfig.MySQL {
		fmt.Printf("   [%d] %s - %s:%s/%s (连接池: %d)\n", 
			i+1, mysql.ModuleName, mysql.Host, mysql.Port, 
			mysql.Dbname, mysql.MaxOpenConns)
	}

	// 显示Redis配置
	fmt.Printf("\n⚡ Redis 实例 (%d个):\n", len(multiConfig.Redis))
	for i, redis := range multiConfig.Redis {
		fmt.Printf("   [%d] %s - %s DB:%d (池大小: %d)\n", 
			i+1, redis.ModuleName, redis.Addr, redis.DB, redis.PoolSize)
	}

	// 显示日志配置
	fmt.Printf("\n📋 日志实例 (%d个):\n", len(multiConfig.Zap))
	for i, zap := range multiConfig.Zap {
		fmt.Printf("   [%d] %s - %s级别, %s格式\n", 
			i+1, zap.ModuleName, zap.Level, zap.Format)
	}

	// 显示JWT配置
	fmt.Printf("\n🔐 JWT 实例 (%d个):\n", len(multiConfig.JWT))
	for i, jwt := range multiConfig.JWT {
		fmt.Printf("   [%d] %s - 过期时间: %d秒, 多点登录: %t\n", 
			i+1, jwt.ModuleName, jwt.ExpiresTime, jwt.UseMultipoint)
	}
}

// configManagementExample 配置管理示例
func configManagementExample() {
	fmt.Println("\n🎯 特定配置获取示例")
	fmt.Println("-------------------")

	ctx := context.Background()
	manager, _ := goconfig.NewMultiConfigManager(ctx, nil)
	multiConfig := manager.GetConfig()

	// 获取主数据库配置
	fmt.Println("📊 获取特定模块配置:")
	
	// 获取主数据库
	if primaryDB, err := goconfig.GetModuleByName(multiConfig.MySQL, "primary"); err == nil {
		fmt.Printf("✅ 主数据库: %s:%s/%s\n", 
			primaryDB.Host, primaryDB.Port, primaryDB.Dbname)
		fmt.Printf("   - 用户名: %s\n", primaryDB.Username)
		fmt.Printf("   - 最大连接: %d\n", primaryDB.MaxOpenConns)
		fmt.Printf("   - 日志级别: %s\n", primaryDB.LogLevel)
	} else {
		fmt.Printf("❌ 获取主数据库配置失败: %v\n", err)
	}

	// 获取只读副本
	if replicaDB, err := goconfig.GetModuleByName(multiConfig.MySQL, "replica"); err == nil {
		fmt.Printf("✅ 只读副本: %s:%s/%s\n", 
			replicaDB.Host, replicaDB.Port, replicaDB.Dbname)
		fmt.Printf("   - 用户名: %s (只读)\n", replicaDB.Username)
	} else {
		fmt.Printf("❌ 获取只读副本配置失败: %v\n", err)
	}

	// 获取缓存Redis
	if cacheRedis, err := goconfig.GetModuleByName(multiConfig.Redis, "cache"); err == nil {
		fmt.Printf("✅ 缓存Redis: %s DB:%d\n", 
			cacheRedis.Addr, cacheRedis.DB)
		fmt.Printf("   - 连接池: %d, 最小空闲: %d\n", 
			cacheRedis.PoolSize, cacheRedis.MinIdleConns)
	} else {
		fmt.Printf("❌ 获取缓存Redis配置失败: %v\n", err)
	}

	// 获取会话Redis
	if sessionRedis, err := goconfig.GetModuleByName(multiConfig.Redis, "session"); err == nil {
		fmt.Printf("✅ 会话Redis: %s DB:%d\n", 
			sessionRedis.Addr, sessionRedis.DB)
	} else {
		fmt.Printf("❌ 获取会话Redis配置失败: %v\n", err)
	}

	// 获取API JWT配置
	if apiJWT, err := goconfig.GetModuleByName(multiConfig.JWT, "api"); err == nil {
		fmt.Printf("✅ API JWT: 过期时间 %d秒\n", apiJWT.ExpiresTime)
		fmt.Printf("   - 多点登录拦截: %t\n", apiJWT.UseMultipoint)
	} else {
		fmt.Printf("❌ 获取API JWT配置失败: %v\n", err)
	}

	// 尝试获取不存在的配置
	if _, err := goconfig.GetModuleByName(multiConfig.MySQL, "nonexistent"); err != nil {
		fmt.Printf("⚠️ 预期错误 - 获取不存在的配置: %v\n", err)
	}
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