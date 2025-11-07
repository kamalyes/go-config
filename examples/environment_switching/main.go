/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 12:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 12:00:00
 * @FilePath: \go-config\examples\environment_switching\main.go
 * @Description: 环境切换示例 - 演示如何在不同环境间切换配置
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
	fmt.Println("🚀 go-config 环境切换示例")
	fmt.Println("==========================")

	// 创建配置目录和文件
	if err := setupEnvironmentConfigs(); err != nil {
		log.Fatalf("设置环境配置失败: %v", err)
	}

	// 示例1: 通过环境变量切换
	fmt.Println("\n🌍 示例1: 通过环境变量切换")
	example1EnvironmentVariableSwitch()

	// 示例2: 通过代码配置切换
	fmt.Println("\n⚙️ 示例2: 通过代码配置切换")
	example2CodeConfigSwitch()

	// 示例3: 动态环境切换
	fmt.Println("\n🔄 示例3: 动态环境切换")
	example3DynamicSwitch()

	// 示例4: 自定义环境配置
	fmt.Println("\n🎨 示例4: 自定义环境配置")
	example4CustomEnvironment()

	fmt.Println("\n✅ 环境切换示例完成!")
}

// setupEnvironmentConfigs 设置不同环境的配置文件
func setupEnvironmentConfigs() error {
	if err := os.MkdirAll("resources", 0755); err != nil {
		return err
	}

	// 开发环境配置
	devConfig := `# 开发环境配置
server:
  addr: '0.0.0.0:8080'
  server-name: 'myapp-dev'
  context-path: '/api/dev'
  handle-method-not-allowed: true
  data-driver: 'sqlite'

mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'dev_user'
  password: 'dev_password'
  db-name: 'myapp_dev'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'debug'
  max-idle-conns: 5
  max-open-conns: 50

redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 0
  pool-size: 50

zap:
  level: 'debug'
  format: 'console'
  prefix: '[DEV]'
  director: 'logs/dev'
  development: true
  log-in-console: true
`

	// 测试环境配置
	testConfig := `# 测试环境配置
server:
  addr: '0.0.0.0:8081'
  server-name: 'myapp-test'
  context-path: '/api/test'
  handle-method-not-allowed: true
  data-driver: 'mysql'

mysql:
  host: '192.168.1.100'
  port: '3306'
  username: 'test_user'
  password: 'test_password'
  db-name: 'myapp_test'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'
  max-idle-conns: 10
  max-open-conns: 100

redis:
  addr: '192.168.1.101:6379'
  password: 'test_redis_pass'
  db: 1
  pool-size: 100

zap:
  level: 'info'
  format: 'json'
  prefix: '[TEST]'
  director: 'logs/test'
  development: false
  log-in-console: false
`

	// 生产环境配置
	prodConfig := `# 生产环境配置
server:
  addr: '0.0.0.0:80'
  server-name: 'myapp-prod'
  context-path: '/api'
  handle-method-not-allowed: false
  data-driver: 'mysql'

mysql:
  host: '10.0.1.100'
  port: '3306'
  username: 'prod_user'
  password: 'super_secure_password'
  db-name: 'myapp_production'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'error'
  max-idle-conns: 20
  max-open-conns: 200

redis:
  addr: '10.0.1.101:6379'
  password: 'prod_redis_secure_pass'
  db: 0
  pool-size: 200

zap:
  level: 'error'
  format: 'json'
  prefix: '[PROD]'
  director: '/var/log/myapp'
  development: false
  log-in-console: false
`

	// 写入配置文件
	configs := map[string]string{
		"resources/dev_config.yaml":  devConfig,
		"resources/sit_config.yaml": testConfig,
		"resources/prod_config.yaml": prodConfig,
	}

	for file, content := range configs {
		if err := os.WriteFile(file, []byte(content), 0644); err != nil {
			return fmt.Errorf("写入配置文件 %s 失败: %v", file, err)
		}
	}

	return nil
}

// example1EnvironmentVariableSwitch 示例1: 通过环境变量切换
func example1EnvironmentVariableSwitch() {
	environments := []string{"dev", "sit", "prod"}

	for _, envName := range environments {
		fmt.Printf("  🌍 设置环境变量: APP_ENV=%s\n", envName)

		// 设置操作系统环境变量
		os.Setenv("APP_ENV", envName)

		ctx := context.Background()
		manager, err := goconfig.NewSingleConfigManager(ctx, nil)
		if err != nil {
			log.Printf("    ❌ 创建配置管理器失败: %v", err)
			continue
		}

		config := manager.GetConfig()
		
		fmt.Printf("    ✅ 当前环境: %s\n", manager.Options.EnvValue)
		fmt.Printf("    ✅ 服务配置: %s (%s)\n", 
			config.Server.ServerName, config.Server.Addr)
		fmt.Printf("    ✅ 数据库: %s@%s:%s/%s\n", 
			config.MySQL.Username, config.MySQL.Host, config.MySQL.Port, config.MySQL.Dbname)
		fmt.Printf("    ✅ Redis: %s (DB:%d)\n", 
			config.Redis.Addr, config.Redis.DB)
		fmt.Printf("    ✅ 日志: %s级别, %s格式\n\n", 
			config.Zap.Level, config.Zap.Format)
	}

	// 清理环境变量
	os.Unsetenv("APP_ENV")
}

// example2CodeConfigSwitch 示例2: 通过代码配置切换
func example2CodeConfigSwitch() {
	environments := []env.EnvironmentType{env.Dev, env.Sit, env.Prod}

	for _, envType := range environments {
		fmt.Printf("  ⚙️ 代码设置环境: %s\n", envType)

		ctx := context.Background()
		options := &goconfig.ConfigOptions{
			ConfigType:    "yaml",
			ConfigPath:    "./resources",
			ConfigSuffix:  "_config",
			EnvValue:      envType,
			EnvContextKey: env.ContextKey("CUSTOM_ENV"),
			UseEnvLevel:   goconfig.EnvLevelCtx, // 使用代码设置优先
		}

		manager, err := goconfig.NewSingleConfigManager(ctx, options)
		if err != nil {
			log.Printf("    ❌ 创建配置管理器失败: %v", err)
			continue
		}

		config := manager.GetConfig()
		
		fmt.Printf("    ✅ 环境: %s\n", envType)
		fmt.Printf("    ✅ 服务: %s\n", config.Server.ServerName)
		fmt.Printf("    ✅ 数据库类型: %s\n", config.Server.DataDriver)
		fmt.Printf("    ✅ 日志目录: %s\n", config.Zap.Director)
		fmt.Printf("    ✅ 是否开发模式: %t\n\n", config.Zap.Development)
	}
}

// example3DynamicSwitch 示例3: 动态环境切换
func example3DynamicSwitch() {
	var currentManager *goconfig.SingleConfigManager
	var currentEnv env.EnvironmentType

	switchToEnvironment := func(envType env.EnvironmentType) error {
		ctx := context.Background()
		options := &goconfig.ConfigOptions{
			EnvValue:      envType,
			UseEnvLevel:   goconfig.EnvLevelCtx,
		}

		manager, err := goconfig.NewSingleConfigManager(ctx, options)
		if err != nil {
			return err
		}

		currentManager = manager
		currentEnv = envType
		return nil
	}

	// 模拟应用启动时的环境切换
	startupSequence := []env.EnvironmentType{env.Dev, env.Sit, env.Prod}

	for i, envType := range startupSequence {
		fmt.Printf("  🔄 步骤 %d: 切换到 %s 环境\n", i+1, envType)

		if err := switchToEnvironment(envType); err != nil {
			log.Printf("    ❌ 切换失败: %v", err)
			continue
		}

		config := currentManager.GetConfig()
		
		fmt.Printf("    ✅ 当前环境: %s\n", currentEnv)
		fmt.Printf("    ✅ 服务端口: %s\n", config.Server.Addr)
		fmt.Printf("    ✅ 数据库连接池: %d-%d\n", 
			config.MySQL.MaxIdleConns, config.MySQL.MaxOpenConns)
		
		// 根据环境进行不同的配置验证
		switch currentEnv {
		case env.Dev:
			if config.Zap.Development {
				fmt.Printf("    ✅ 开发环境验证通过: 开启开发模式\n")
			}
		case env.Prod:
			if !config.Zap.Development && config.Zap.Level == "error" {
				fmt.Printf("    ✅ 生产环境验证通过: 错误级别日志\n")
			}
		}
		fmt.Println()
	}
}

// example4CustomEnvironment 示例4: 自定义环境配置
func example4CustomEnvironment() {
	// 创建自定义环境配置
	customConfig := `# 自定义环境配置
server:
  addr: '0.0.0.0:9999'
  server-name: 'myapp-custom'
  context-path: '/custom'
  handle-method-not-allowed: true
  data-driver: 'postgresql'

mysql:
  host: 'custom.db.example.com'
  port: '5432'
  username: 'custom_user'
  password: 'custom_password'
  db-name: 'custom_database'
  config: 'sslmode=disable TimeZone=Asia/Shanghai'
  log-level: 'warn'
  max-idle-conns: 15
  max-open-conns: 150

redis:
  addr: 'custom.redis.example.com:6379'
  password: 'custom_redis_password'
  db: 5
  pool-size: 75

zap:
  level: 'warn'
  format: 'json'
  prefix: '[CUSTOM]'
  director: 'logs/custom'
  development: false
  log-in-console: true
`

	// 写入自定义配置
	customEnvFile := "resources/custom_config.yaml"
	if err := os.WriteFile(customEnvFile, []byte(customConfig), 0644); err != nil {
		log.Printf("❌ 创建自定义配置失败: %v", err)
		return
	}

	fmt.Printf("  🎨 创建自定义环境配置\n")

	// 使用自定义环境
	customEnv := env.EnvironmentType("custom")
	
	ctx := context.Background()
	options := &goconfig.ConfigOptions{
		ConfigType:    "yaml",
		ConfigPath:    "./resources",
		ConfigSuffix:  "_config",
		EnvValue:      customEnv,
		EnvContextKey: env.ContextKey("CUSTOM_DEMO_ENV"),
		UseEnvLevel:   goconfig.EnvLevelCtx,
	}

	manager, err := goconfig.NewSingleConfigManager(ctx, options)
	if err != nil {
		log.Printf("❌ 创建自定义配置管理器失败: %v", err)
		return
	}

	config := manager.GetConfig()
	
	fmt.Printf("    ✅ 自定义环境: %s\n", customEnv)
	fmt.Printf("    ✅ 服务配置: %s (%s)\n", 
		config.Server.ServerName, config.Server.Addr)
	fmt.Printf("    ✅ 数据库: %s@%s:%s\n", 
		config.MySQL.Username, config.MySQL.Host, config.MySQL.Port)
	fmt.Printf("    ✅ 数据库类型: %s\n", config.Server.DataDriver)
	fmt.Printf("    ✅ Redis: %s (DB:%d)\n", 
		config.Redis.Addr, config.Redis.DB)

	// 配置比较
	fmt.Printf("\n  📊 与标准环境的差异:\n")
	
	// 与开发环境比较
	devOptions := &goconfig.ConfigOptions{
		EnvValue: env.Dev,
		UseEnvLevel: goconfig.EnvLevelCtx,
	}
	devManager, _ := goconfig.NewSingleConfigManager(ctx, devOptions)
	devConfig := devManager.GetConfig()
	
	fmt.Printf("    - 端口差异: dev=%s, custom=%s\n", 
		devConfig.Server.Addr, config.Server.Addr)
	fmt.Printf("    - 数据库类型: dev=%s, custom=%s\n", 
		devConfig.Server.DataDriver, config.Server.DataDriver)
	fmt.Printf("    - Redis DB: dev=%d, custom=%d\n", 
		devConfig.Redis.DB, config.Redis.DB)

	// 清理自定义配置文件
	os.Remove(customEnvFile)
	fmt.Printf("    🗑️ 清理自定义配置文件\n")
}