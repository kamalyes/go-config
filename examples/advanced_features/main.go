/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 12:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 12:00:00
 * @FilePath: \go-config\examples\advanced_features\main.go
 * @Description: 高级特性示例 - 演示配置热更新、外部Viper、动态配置等高级功能
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/env"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("🚀 go-config 高级特性示例")
	fmt.Println("==========================")

	// 创建配置目录
	if err := createConfigDir(); err != nil {
		log.Fatalf("创建配置目录失败: %v", err)
	}

	// 创建配置文件
	if err := createAdvancedConfig(); err != nil {
		log.Fatalf("创建配置文件失败: %v", err)
	}

	// 示例1: 配置热更新
	fmt.Println("\n🔄 示例1: 配置热更新")
	example1HotReload()

	// 示例2: 外部Viper集成
	fmt.Println("\n🔌 示例2: 外部Viper集成")
	example2ExternalViper()

	// 示例3: 动态配置管理
	fmt.Println("\n📊 示例3: 动态配置管理")
	example3DynamicConfig()

	// 示例4: 环境变量优先级
	fmt.Println("\n🌍 示例4: 环境变量优先级")
	example4EnvironmentPriority()

	// 示例5: 配置文件合并
	fmt.Println("\n🔀 示例5: 配置文件合并")
	example5ConfigMerging()

	fmt.Println("\n🎉 高级特性示例完成!")
}

// createConfigDir 创建配置目录
func createConfigDir() error {
	dirs := []string{"resources", "external_configs"}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

// createAdvancedConfig 创建高级配置文件
func createAdvancedConfig() error {
	// 基础配置
	baseConfig := `# 基础服务配置
server:
  addr: '0.0.0.0:8080'
  server-name: 'advanced-example'
  context-path: '/api'
  handle-method-not-allowed: true
  data-driver: 'mysql'

# 数据库配置
mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'root'
  password: 'password123'
  db-name: 'advanced_example'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'
  max-idle-conns: 10
  max-open-conns: 100

# Redis配置
redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 0
  pool-size: 100

# 日志配置
zap:
  level: 'info'
  format: 'console'
  prefix: '[ADVANCED]'
  director: 'logs'
  development: true
`

	// 外部配置
	externalConfig := `# 外部配置示例
external_service:
  api_url: "https://api.external.com"
  api_key: "external-api-key-123"
  timeout: 30
  retry_count: 3

feature_flags:
  enable_new_ui: true
  enable_cache: true
  enable_metrics: false
  maintenance_mode: false

monitoring:
  prometheus:
    enabled: true
    port: 9090
    path: "/metrics"
  jaeger:
    enabled: false
    endpoint: "http://localhost:14268/api/traces"

rate_limiting:
  requests_per_minute: 1000
  burst_size: 100
  whitelist:
    - "127.0.0.1"
    - "192.168.1.0/24"
`

	// 写入配置文件
	files := map[string]string{
		"resources/dev_config.yaml":       baseConfig,
		"resources/prod_config.yaml":      baseConfig,
		"external_configs/external.yaml": externalConfig,
	}

	for file, content := range files {
		if err := os.WriteFile(file, []byte(content), 0644); err != nil {
			return fmt.Errorf("创建配置文件 %s 失败: %v", file, err)
		}
	}

	return nil
}

// example1HotReload 示例1: 配置热更新
func example1HotReload() {
	ctx := context.Background()
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 创建配置管理器失败: %v", err)
		return
	}

	// 初始配置
	config := manager.GetConfig()
	fmt.Printf("  📄 初始配置 - 服务名称: %s, 地址: %s\n", 
		config.Server.ServerName, config.Server.Addr)

	// 模拟配置文件更新
	fmt.Println("  🔄 模拟配置文件更新...")
	
	updatedConfig := `# 更新后的服务配置
server:
  addr: '0.0.0.0:8090'
  server-name: 'advanced-example-updated'
  context-path: '/api/v2'
  handle-method-not-allowed: false
  data-driver: 'postgresql'

mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'root'
  password: 'password123'
  db-name: 'advanced_example'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'debug'

redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 1

zap:
  level: 'debug'
  format: 'json'
  prefix: '[UPDATED]'
  director: 'logs'
  development: false
`

	// 写入更新的配置
	if err := os.WriteFile("resources/dev_config.yaml", []byte(updatedConfig), 0644); err != nil {
		log.Printf("❌ 更新配置文件失败: %v", err)
		return
	}

	// 等待配置热更新生效
	time.Sleep(1 * time.Second)
	
	// 重新获取配置
	config = manager.GetConfig()
	fmt.Printf("  ✅ 更新后配置 - 服务名称: %s, 地址: %s\n", 
		config.Server.ServerName, config.Server.Addr)
	fmt.Printf("  ✅ 日志级别更新: %s -> %s\n", "info", config.Zap.Level)
	fmt.Printf("  ✅ Redis DB更新: %s -> %d\n", "0", config.Redis.DB)
}

// example2ExternalViper 示例2: 外部Viper集成
func example2ExternalViper() {
	ctx := context.Background()
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 创建配置管理器失败: %v", err)
		return
	}

	config := manager.GetConfig()

	// 创建外部Viper实例
	externalViper := viper.New()
	externalViper.SetConfigName("external")
	externalViper.SetConfigType("yaml")
	externalViper.AddConfigPath("./external_configs")
	
	if err := externalViper.ReadInConfig(); err != nil {
		log.Printf("❌ 读取外部配置失败: %v", err)
		return
	}

	// 添加到配置管理器
	config.AddExternalViper("external", externalViper)
	
	fmt.Printf("  🔌 外部Viper实例数量: %d\n", len(config.GetAllExternalViperKeys()))
	
	// 使用外部配置
	if extViper, exists := config.GetExternalViper("external"); exists {
		apiUrl := extViper.GetString("external_service.api_url")
		apiKey := extViper.GetString("external_service.api_key")
		timeout := extViper.GetInt("external_service.timeout")
		
		fmt.Printf("  ✅ 外部API: %s\n", apiUrl)
		fmt.Printf("  ✅ API密钥: %s\n", maskString(apiKey))
		fmt.Printf("  ✅ 超时时间: %d秒\n", timeout)

		// 获取功能开关
		enableNewUI := extViper.GetBool("feature_flags.enable_new_ui")
		enableCache := extViper.GetBool("feature_flags.enable_cache")
		fmt.Printf("  ✅ 新UI开关: %t\n", enableNewUI)
		fmt.Printf("  ✅ 缓存开关: %t\n", enableCache)
	}

	// 解析到结构体
	type ExternalServiceConfig struct {
		APIURL     string `mapstructure:"api_url"`
		APIKey     string `mapstructure:"api_key"`
		Timeout    int    `mapstructure:"timeout"`
		RetryCount int    `mapstructure:"retry_count"`
	}

	var serviceConfig ExternalServiceConfig
	if err := config.UnmarshalSubFromExternalViper("external", "external_service", &serviceConfig); err != nil {
		log.Printf("❌ 解析外部服务配置失败: %v", err)
	} else {
		fmt.Printf("  ✅ 解析到结构体: %+v\n", serviceConfig)
	}
}

// example3DynamicConfig 示例3: 动态配置管理
func example3DynamicConfig() {
	ctx := context.Background()
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 创建配置管理器失败: %v", err)
		return
	}

	config := manager.GetConfig()

	// 设置动态配置
	config.SetDynamicConfig("runtime_settings", map[string]interface{}{
		"max_connections":    1000,
		"enable_debug":       true,
		"cache_ttl_seconds":  3600,
		"worker_pool_size":   10,
	})

	config.SetDynamicConfig("business_rules", map[string]interface{}{
		"max_order_amount":   10000.00,
		"discount_enabled":   true,
		"free_shipping_min":  100.00,
		"vip_discount_rate":  0.15,
	})

	config.SetDynamicConfig("security", map[string]interface{}{
		"jwt_secret_rotation": true,
		"password_min_length": 8,
		"session_timeout":     1800,
		"max_login_attempts":  5,
	})

	// 获取动态配置列表
	keys := config.GetAllDynamicConfigKeys()
	fmt.Printf("  📊 动态配置项: %v\n", keys)

	// 使用动态配置
	if settings, exists := config.GetDynamicConfig("runtime_settings"); exists {
		fmt.Printf("  🔧 运行时设置:\n")
		if settingsMap, ok := settings.(map[string]interface{}); ok {
			for key, value := range settingsMap {
				fmt.Printf("    - %s: %v\n", key, value)
			}
		}
	}

	if business, exists := config.GetDynamicConfig("business_rules"); exists {
		fmt.Printf("  💼 业务规则:\n")
		if businessMap, ok := business.(map[string]interface{}); ok {
			for key, value := range businessMap {
				fmt.Printf("    - %s: %v\n", key, value)
			}
		}
	}

	// 更新动态配置
	config.SetDynamicConfig("runtime_settings", map[string]interface{}{
		"max_connections":   2000, // 更新值
		"enable_debug":      false,
		"cache_ttl_seconds": 7200,
		"worker_pool_size":  20,
		"new_feature":       true, // 新增字段
	})

	fmt.Printf("  🔄 更新后的运行时设置:\n")
	if settings, exists := config.GetDynamicConfig("runtime_settings"); exists {
		if settingsMap, ok := settings.(map[string]interface{}); ok {
			for key, value := range settingsMap {
				fmt.Printf("    - %s: %v\n", key, value)
			}
		}
	}

	// 删除动态配置
	if config.RemoveDynamicConfig("security") {
		fmt.Printf("  🗑️ 已删除安全配置\n")
	}

	// 再次查看配置列表
	keys = config.GetAllDynamicConfigKeys()
	fmt.Printf("  📊 更新后的动态配置项: %v\n", keys)
}

// example4EnvironmentPriority 示例4: 环境变量优先级
func example4EnvironmentPriority() {
	environments := []struct {
		name     string
		envType  env.EnvironmentType
		useLevel goconfig.EnvLevel
	}{
		{"操作系统环境变量优先", env.Prod, goconfig.EnvLevelOS},
		{"代码设置优先", env.Dev, goconfig.EnvLevelCtx},
	}

	for _, envConfig := range environments {
		fmt.Printf("  🌍 测试: %s\n", envConfig.name)

		// 设置操作系统环境变量
		os.Setenv("APP_ENV", "prod")

		// 设置代码中的环境变量
		env.SetContextKey(&env.ContextKeyOptions{
			Key:   env.ContextKey("APP_ENV"),
			Value: env.Dev,
		})

		ctx := context.Background()
		options := &goconfig.ConfigOptions{
			EnvValue:    envConfig.envType,
			UseEnvLevel: envConfig.useLevel,
		}

		manager, err := goconfig.NewSingleConfigManager(ctx, options)
		if err != nil {
			log.Printf("    ❌ 创建配置管理器失败: %v", err)
			continue
		}

		fmt.Printf("    ✅ 实际使用环境: %s\n", manager.Options.EnvValue)
		fmt.Printf("    ✅ 优先级策略: %s\n", envConfig.useLevel)
	}

	// 清理环境变量
	os.Unsetenv("APP_ENV")
}

// example5ConfigMerging 示例5: 配置文件合并
func example5ConfigMerging() {
	ctx := context.Background()
	
	// 创建主配置管理器
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 创建配置管理器失败: %v", err)
		return
	}

	config := manager.GetConfig()

	// 加载外部配置
	externalViper := viper.New()
	externalViper.SetConfigName("external")
	externalViper.SetConfigType("yaml")
	externalViper.AddConfigPath("./external_configs")
	
	if err := externalViper.ReadInConfig(); err != nil {
		log.Printf("❌ 读取外部配置失败: %v", err)
		return
	}

	// 合并配置
	config.AddExternalViper("monitoring", externalViper)
	
	// 设置运行时配置
	config.SetDynamicConfig("computed", map[string]interface{}{
		"total_pool_size": config.Redis.PoolSize + 50,
		"db_connection_string": fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			config.MySQL.Username, "***", config.MySQL.Host, config.MySQL.Port, config.MySQL.Dbname),
		"service_endpoints": []string{
			config.Server.Addr,
			"backup:8081",
		},
	})

	fmt.Printf("  🔀 配置合并完成:\n")
	fmt.Printf("    ✅ 主配置: 服务器=%s, 数据库=%s:%s\n", 
		config.Server.Addr, config.MySQL.Host, config.MySQL.Port)
	
	if extViper, exists := config.GetExternalViper("monitoring"); exists {
		prometheusEnabled := extViper.GetBool("monitoring.prometheus.enabled")
		prometheusPort := extViper.GetInt("monitoring.prometheus.port")
		fmt.Printf("    ✅ 外部配置: Prometheus=%t, 端口=%d\n", 
			prometheusEnabled, prometheusPort)
	}

	if computed, exists := config.GetDynamicConfig("computed"); exists {
		fmt.Printf("    ✅ 计算配置:\n")
		if computedMap, ok := computed.(map[string]interface{}); ok {
			for key, value := range computedMap {
				if key != "db_connection_string" { // 敏感信息不打印
					fmt.Printf("      - %s: %v\n", key, value)
				}
			}
		}
	}
}

// maskString 遮蔽敏感信息
func maskString(s string) string {
	if len(s) <= 8 {
		return "****"
	}
	return s[:4] + "****" + s[len(s)-4:]
}