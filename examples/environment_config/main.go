/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-08 16:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-08 16:00:00
 * @FilePath: \go-config\examples\environment_config\main.go
 * @Description: go-config 环境变量和自定义配置示例
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
	"time"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/env"
)

func main() {
	fmt.Println("🚀 go-config 环境变量和自定义配置示例")
	fmt.Println("=====================================")

	// 创建多环境配置文件
	if err := createEnvironmentConfigs(); err != nil {
		log.Fatalf("❌ 创建配置文件失败: %v", err)
	}

	// 演示不同环境的配置
	demonstrateEnvironments()

	// 演示自定义配置选项
	demonstrateCustomOptions()

	// 演示环境切换
	demonstrateEnvironmentSwitching()

	// 清理
	cleanup()
}

// createEnvironmentConfigs 创建多环境配置文件
func createEnvironmentConfigs() error {
	environments := map[string]string{
		"dev": `# 开发环境配置
server:
  addr: '127.0.0.1:8080'
  server-name: 'dev-service'
  context-path: '/dev/api'
  data-driver: 'sqlite'

mysql:
  host: 'localhost'
  port: '3306'
  username: 'dev_user'
  password: 'dev_pass'
  db-name: 'dev_database'
  log-level: 'debug'
  max-open-conns: 50

redis:
  addr: 'localhost:6379'
  db: 0
  pool-size: 50

zap:
  level: 'debug'
  format: 'console'
  prefix: '[DEV]'
  development: true
  log-in-console: true
  show-line: true
`,
		"sit": `# 系统集成测试环境配置
server:
  addr: '0.0.0.0:8080'
  server-name: 'sit-service'
  context-path: '/sit/api'
  data-driver: 'mysql'

mysql:
  host: 'sit-mysql.internal'
  port: '3306'
  username: 'sit_user'
  password: 'sit_password'
  db-name: 'sit_database'
  log-level: 'info'
  max-open-conns: 100

redis:
  addr: 'sit-redis.internal:6379'
  db: 1
  pool-size: 100

zap:
  level: 'info'
  format: 'json'
  prefix: '[SIT]'
  development: false
  log-in-console: false
`,
		"prod": `# 生产环境配置
server:
  addr: '0.0.0.0:80'
  server-name: 'prod-service'
  context-path: '/api'
  data-driver: 'mysql'

mysql:
  host: 'prod-mysql-cluster.internal'
  port: '3306'
  username: 'prod_user'
  password: 'super_secure_prod_password'
  db-name: 'production_database'
  log-level: 'warn'
  max-open-conns: 200
  max-idle-conns: 50

redis:
  addr: 'prod-redis-cluster.internal:6379'
  password: 'redis_prod_password'
  db: 0
  pool-size: 200
  min-idle-conns: 20

zap:
  level: 'warn'
  format: 'json'
  prefix: '[PROD]'
  director: '/var/log/app'
  development: false
  log-in-console: false

cors:
  allowed-all-origins: false
  allowed-origins:
    - "https://myapp.com"
    - "https://www.myapp.com"
  allow-credentials: true
  max-age: "86400"
`,
	}

	resourcesDir := "./resources"
	if err := os.MkdirAll(resourcesDir, os.ModePerm); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	for envName, content := range environments {
		configFile := filepath.Join(resourcesDir, fmt.Sprintf("%s_config.yaml", envName))
		if err := os.WriteFile(configFile, []byte(content), 0644); err != nil {
			return fmt.Errorf("写入 %s 配置文件失败: %w", envName, err)
		}
		fmt.Printf("✅ %s 环境配置已创建: %s\n", envName, configFile)
	}

	return nil
}

// demonstrateEnvironments 演示不同环境的配置
func demonstrateEnvironments() {
	fmt.Println("\n🌍 不同环境配置演示")
	fmt.Println("------------------")

	environments := []env.EnvironmentType{env.Dev, env.Sit, env.Prod}
	ctx := context.Background()

	for _, environment := range environments {
		fmt.Printf("\n📋 加载 %s 环境配置:\n", environment)
		
		// 设置环境变量
		os.Setenv("APP_ENV", string(environment))
		
		// 创建配置管理器
		manager, err := goconfig.NewSingleConfigManager(ctx, nil)
		if err != nil {
			log.Printf("❌ 创建 %s 环境配置管理器失败: %v", environment, err)
			continue
		}

		config := manager.GetConfig()

		// 显示配置信息
		fmt.Printf("   🌐 服务器: %s (%s)\n", config.Server.Addr, config.Server.ServerName)
		fmt.Printf("   💾 数据库: %s:%s/%s\n", config.MySQL.Host, config.MySQL.Port, config.MySQL.Dbname)
		fmt.Printf("   ⚡ Redis: %s DB:%d\n", config.Redis.Addr, config.Redis.DB)
		fmt.Printf("   📋 日志: %s级别, %s格式\n", config.Zap.Level, config.Zap.Format)
		
		// 特殊配置显示
		if environment == env.Prod && len(config.Cors.AllowedOrigins) > 0 {
			fmt.Printf("   🌍 CORS 允许来源: %v\n", config.Cors.AllowedOrigins)
		}
	}

	// 恢复默认环境
	os.Setenv("APP_ENV", "dev")
}

// demonstrateCustomOptions 演示自定义配置选项
func demonstrateCustomOptions() {
	fmt.Println("\n⚙️ 自定义配置选项演示")
	fmt.Println("--------------------")

	// 创建自定义配置文件
	customConfigContent := `# 自定义配置示例
server:
  addr: '0.0.0.0:9999'
  server-name: 'custom-service'
  context-path: '/custom/api'

mysql:
  host: 'custom-db.local'
  port: '3306'
  username: 'custom_user'
  password: 'custom_pass'
  db-name: 'custom_db'

zap:
  level: 'info'
  format: 'json'
  prefix: '[CUSTOM]'
`

	customDir := "./custom_configs"
	if err := os.MkdirAll(customDir, os.ModePerm); err != nil {
		log.Printf("❌ 创建自定义目录失败: %v", err)
		return
	}

	customFile := filepath.Join(customDir, "special_settings.yaml")
	if err := os.WriteFile(customFile, []byte(customConfigContent), 0644); err != nil {
		log.Printf("❌ 创建自定义配置文件失败: %v", err)
		return
	}

	ctx := context.Background()

	// 方式1: 自定义路径和文件名
	fmt.Println("📁 方式1: 自定义配置路径和文件名")
	customOptions1 := &goconfig.ConfigOptions{
		ConfigType:   "yaml",
		ConfigPath:   "./custom_configs",
		ConfigSuffix: "_settings",
		EnvValue:     env.EnvironmentType("special"),
	}

	manager1, err := goconfig.NewSingleConfigManager(ctx, customOptions1)
	if err != nil {
		log.Printf("❌ 自定义配置1失败: %v", err)
	} else {
		config1 := manager1.GetConfig()
		fmt.Printf("   ✅ 服务器: %s (%s)\n", config1.Server.Addr, config1.Server.ServerName)
		fmt.Printf("   ✅ 数据库: %s/%s\n", config1.MySQL.Host, config1.MySQL.Dbname)
	}

	// 方式2: 自定义环境变量Key
	fmt.Println("\n🔑 方式2: 自定义环境变量Key")
	customEnv := env.EnvironmentType("custom")
	customContextKey := env.ContextKey("MY_CUSTOM_ENV")
	
	// 设置自定义环境变量
	env.SetContextKey(&env.ContextKeyOptions{
		Key:   customContextKey,
		Value: customEnv,
	})

	customOptions2 := &goconfig.ConfigOptions{
		ConfigType:    "yaml",
		ConfigPath:    "./resources",
		ConfigSuffix:  "_config",
		EnvValue:      env.Dev, // 使用dev环境的配置
		EnvContextKey: customContextKey,
		UseEnvLevel:   goconfig.EnvLevelCtx, // 优先使用代码设置的环境
	}

	manager2, err := goconfig.NewSingleConfigManager(ctx, customOptions2)
	if err != nil {
		log.Printf("❌ 自定义配置2失败: %v", err)
	} else {
		config2 := manager2.GetConfig()
		fmt.Printf("   ✅ 环境Key: %s\n", customContextKey)
		fmt.Printf("   ✅ 服务器: %s\n", config2.Server.Addr)
	}

	// 清理自定义文件
	os.RemoveAll(customDir)
}

// demonstrateEnvironmentSwitching 演示环境切换
func demonstrateEnvironmentSwitching() {
	fmt.Println("\n🔄 环境切换演示")
	fmt.Println("---------------")

	// 创建环境管理器
	envManager := env.NewEnvironment()
	defer envManager.StopWatch() // 停止监控

	// 设置检查频率为1秒（用于演示）
	envManager.SetCheckFrequency(1 * time.Second)

	ctx := context.Background()

	// 初始环境为dev
	fmt.Printf("📍 当前环境: %s\n", env.GetEnvironment())

	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 创建配置管理器失败: %v", err)
		return
	}

	config := manager.GetConfig()
	fmt.Printf("   ✅ 初始服务器地址: %s\n", config.Server.Addr)

	// 模拟环境切换
	fmt.Println("\n🔄 切换到 sit 环境...")
	envManager.SetEnvironment(env.Sit)
	time.Sleep(2 * time.Second) // 等待监控器检测到变化

	fmt.Printf("📍 当前环境: %s\n", env.GetEnvironment())

	// 重新创建管理器以加载新环境配置
	newManager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 重新创建配置管理器失败: %v", err)
		return
	}

	newConfig := newManager.GetConfig()
	fmt.Printf("   ✅ 新服务器地址: %s\n", newConfig.Server.Addr)
	fmt.Printf("   ✅ 新数据库主机: %s\n", newConfig.MySQL.Host)

	// 演示不同的UseEnvLevel
	fmt.Println("\n📊 环境优先级演示:")
	
	// OS环境变量优先
	os.Setenv("APP_ENV", "prod")
	osOptions := &goconfig.ConfigOptions{
		EnvValue:    env.Dev, // 代码中指定dev
		UseEnvLevel: goconfig.EnvLevelOS, // 但优先使用OS环境变量
	}
	
	osManager, _ := goconfig.NewSingleConfigManager(ctx, osOptions)
	osConfig := osManager.GetConfig()
	fmt.Printf("   🖥️ OS优先 (APP_ENV=prod): %s\n", osConfig.Server.ServerName)

	// 代码环境变量优先
	ctxOptions := &goconfig.ConfigOptions{
		EnvValue:    env.Dev, // 代码中指定dev
		UseEnvLevel: goconfig.EnvLevelCtx, // 优先使用代码设置
	}
	
	ctxManager, _ := goconfig.NewSingleConfigManager(ctx, ctxOptions)
	ctxConfig := ctxManager.GetConfig()
	fmt.Printf("   💻 代码优先 (强制dev): %s\n", ctxConfig.Server.ServerName)

	// 恢复环境
	envManager.SetEnvironment(env.Dev)
	os.Setenv("APP_ENV", "dev")
}

// cleanup 清理测试文件
func cleanup() {
	if err := os.RemoveAll("./resources"); err != nil {
		log.Printf("⚠️ 清理resources失败: %v", err)
	}
	if err := os.RemoveAll("./custom_configs"); err != nil {
		log.Printf("⚠️ 清理custom_configs失败: %v", err)
	}
	if err := os.RemoveAll("./logs"); err != nil {
		log.Printf("⚠️ 清理日志文件失败: %v", err)
	}
	fmt.Println("\n🧹 环境切换示例清理完成")
}