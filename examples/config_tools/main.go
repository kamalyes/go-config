/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-08 12:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-08 12:00:00
 * @FilePath: \go-config\examples\config_tools\main.go
 * @Description: 配置工具示例 - 演示配置验证、比较、导出等实用工具功能
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/env"
)

func main() {
	fmt.Println("🚀 go-config 配置工具示例")
	fmt.Println("==========================")

	// 创建测试配置
	if err := setupTestConfigs(); err != nil {
		log.Fatalf("设置测试配置失败: %v", err)
	}

	// 工具1: 配置验证工具
	fmt.Println("\n✅ 工具1: 配置验证")
	tool1ConfigValidation()

	// 工具2: 配置比较工具
	fmt.Println("\n🔍 工具2: 配置比较")
	tool2ConfigComparison()

	// 工具3: 配置导出工具
	fmt.Println("\n📤 工具3: 配置导出")
	tool3ConfigExport()

	// 工具4: 配置诊断工具
	fmt.Println("\n🩺 工具4: 配置诊断")
	tool4ConfigDiagnosis()

	// 工具5: 配置模板生成器
	fmt.Println("\n📝 工具5: 配置模板生成")
	tool5ConfigTemplateGenerator()

	fmt.Println("\n🎉 配置工具示例完成!")
}

// setupTestConfigs 创建测试配置文件
func setupTestConfigs() error {
	if err := os.MkdirAll("resources", 0755); err != nil {
		return err
	}

	// 正确的配置
	validConfig := `# 有效配置
server:
  addr: '0.0.0.0:8080'
  server-name: 'valid-service'
  context-path: '/api'
  handle-method-not-allowed: true
  data-driver: 'mysql'

mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'root'
  password: 'password123'
  db-name: 'test_db'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'
  max-idle-conns: 10
  max-open-conns: 100
  conn-max-idle-time: 30
  conn-max-life-time: 300

redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 0
  pool-size: 100
  min-idle-conns: 10
  max-retries: 3

zap:
  level: 'info'
  format: 'console'
  prefix: '[VALID]'
  director: 'logs'
  development: true
`

	// 有问题的配置
	invalidConfig := `# 有问题的配置
server:
  addr: ''  # 空地址
  server-name: ''  # 空服务名
  context-path: '/api'
  data-driver: 'unknown'  # 未知驱动

mysql:
  host: ''  # 空主机
  port: 'invalid_port'  # 无效端口
  username: ''  # 空用户名
  password: ''
  db-name: ''  # 空数据库名
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'invalid_level'  # 无效日志级别
  max-idle-conns: -1  # 负数连接
  max-open-conns: 0   # 零连接

redis:
  addr: 'invalid:address'  # 无效地址
  db: 16   # 超出范围的DB
  pool-size: 0  # 零连接池
  min-idle-conns: -5  # 负数
  max-retries: -1  # 负重试次数

zap:
  level: 'unknown'  # 未知级别
  format: 'invalid'  # 无效格式
  director: ''  # 空目录
`

	configs := map[string]string{
		"resources/dev_config.yaml":     validConfig,
		"resources/invalid_config.yaml": invalidConfig,
	}

	for file, content := range configs {
		if err := os.WriteFile(file, []byte(content), 0644); err != nil {
			return fmt.Errorf("创建配置文件 %s 失败: %v", file, err)
		}
	}

	return nil
}

// tool1ConfigValidation 工具1: 配置验证工具
func tool1ConfigValidation() {
	configs := []struct {
		name        string
		env         env.EnvironmentType
		expectValid bool
	}{
		{"有效配置", env.Dev, true},
		{"无效配置", env.EnvironmentType("invalid"), false},
	}

	for _, testConfig := range configs {
		fmt.Printf("  🔍 验证 %s:\n", testConfig.name)

		ctx := context.Background()
		options := &goconfig.ConfigOptions{
			EnvValue:    testConfig.env,
			UseEnvLevel: goconfig.EnvLevelCtx,
		}

		manager, err := goconfig.NewSingleConfigManager(ctx, options)
		if err != nil {
			fmt.Printf("    ❌ 配置加载失败: %v\n", err)
			continue
		}

		config := manager.GetConfig()

		// 验证各个组件
		validationResults := validateAllConfigs(config)
		
		totalErrors := 0
		for component, errors := range validationResults {
			if len(errors) > 0 {
				fmt.Printf("    ❌ %s 验证失败:\n", component)
				for _, err := range errors {
					fmt.Printf("      - %s\n", err)
					totalErrors++
				}
			} else {
				fmt.Printf("    ✅ %s 验证通过\n", component)
			}
		}

		if totalErrors == 0 {
			fmt.Printf("    🎉 整体验证: 全部通过 (%d个组件)\n", len(validationResults))
		} else {
			fmt.Printf("    ⚠️ 整体验证: 发现 %d 个错误\n", totalErrors)
		}
		fmt.Println()
	}
}

// validateAllConfigs 验证所有配置组件
func validateAllConfigs(config *goconfig.SingleConfig) map[string][]string {
	results := make(map[string][]string)

	// 服务器配置验证
	var serverErrors []string
	if err := config.Server.Validate(); err != nil {
		serverErrors = append(serverErrors, err.Error())
	}
	if config.Server.Addr == "" {
		serverErrors = append(serverErrors, "服务地址不能为空")
	}
	if config.Server.ServerName == "" {
		serverErrors = append(serverErrors, "服务名称不能为空")
	}
	results["Server"] = serverErrors

	// MySQL配置验证
	var mysqlErrors []string
	if err := config.MySQL.Validate(); err != nil {
		mysqlErrors = append(mysqlErrors, err.Error())
	}
	if config.MySQL.Host == "" {
		mysqlErrors = append(mysqlErrors, "MySQL主机地址不能为空")
	}
	if config.MySQL.Dbname == "" {
		mysqlErrors = append(mysqlErrors, "数据库名不能为空")
	}
	if config.MySQL.MaxOpenConns <= 0 {
		mysqlErrors = append(mysqlErrors, "最大连接数必须大于0")
	}
	results["MySQL"] = mysqlErrors

	// Redis配置验证
	var redisErrors []string
	if err := config.Redis.Validate(); err != nil {
		redisErrors = append(redisErrors, err.Error())
	}
	if config.Redis.Addr == "" {
		redisErrors = append(redisErrors, "Redis地址不能为空")
	}
	if config.Redis.DB < 0 || config.Redis.DB > 15 {
		redisErrors = append(redisErrors, "Redis DB索引必须在0-15之间")
	}
	if config.Redis.PoolSize <= 0 {
		redisErrors = append(redisErrors, "连接池大小必须大于0")
	}
	results["Redis"] = redisErrors

	return results
}

// tool2ConfigComparison 工具2: 配置比较工具
func tool2ConfigComparison() {
	ctx := context.Background()

	// 加载开发环境配置
	devManager, err := goconfig.NewSingleConfigManager(ctx, &goconfig.ConfigOptions{
		EnvValue: env.Dev,
		UseEnvLevel: goconfig.EnvLevelCtx,
	})
	if err != nil {
		log.Printf("❌ 加载开发环境配置失败: %v", err)
		return
	}

	// 模拟生产环境配置
	prodConfig := `server:
  addr: '0.0.0.0:80'
  server-name: 'valid-service-prod'
  context-path: '/api'
  handle-method-not-allowed: false
  data-driver: 'mysql'

mysql:
  host: '10.0.1.100'
  port: '3306'
  username: 'prod_user'
  password: 'secure_password'
  db-name: 'production_db'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'error'
  max-idle-conns: 20
  max-open-conns: 200
  conn-max-idle-time: 60
  conn-max-life-time: 600

redis:
  addr: '10.0.1.101:6379'
  password: 'redis_password'
  db: 0
  pool-size: 200
  min-idle-conns: 20
  max-retries: 5

zap:
  level: 'error'
  format: 'json'
  prefix: '[PROD]'
  director: '/var/log/app'
  development: false
`

	// 创建生产环境配置文件
	if err := os.WriteFile("resources/prod_config.yaml", []byte(prodConfig), 0644); err != nil {
		log.Printf("❌ 创建生产配置失败: %v", err)
		return
	}

	prodManager, err := goconfig.NewSingleConfigManager(ctx, &goconfig.ConfigOptions{
		EnvValue: env.Prod,
		UseEnvLevel: goconfig.EnvLevelCtx,
	})
	if err != nil {
		log.Printf("❌ 加载生产环境配置失败: %v", err)
		return
	}

	// 比较配置
	devConfig := devManager.GetConfig()
	prodConfigObj := prodManager.GetConfig()

	fmt.Printf("  🔍 开发 vs 生产环境配置比较:\n")
	compareConfigs("Server", devConfig.Server, prodConfigObj.Server)
	compareConfigs("MySQL", devConfig.MySQL, prodConfigObj.MySQL)
	compareConfigs("Redis", devConfig.Redis, prodConfigObj.Redis)
	compareConfigs("Zap", devConfig.Zap, prodConfigObj.Zap)
}

// compareConfigs 比较两个配置对象
func compareConfigs(name string, dev, prod interface{}) {
	fmt.Printf("    📊 %s 配置差异:\n", name)
	
	devValue := reflect.ValueOf(dev)
	prodValue := reflect.ValueOf(prod)
	devType := reflect.TypeOf(dev)

	differences := 0
	for i := 0; i < devType.NumField(); i++ {
		field := devType.Field(i)
		devFieldValue := devValue.Field(i)
		prodFieldValue := prodValue.Field(i)

		if !reflect.DeepEqual(devFieldValue.Interface(), prodFieldValue.Interface()) {
			fmt.Printf("      - %s: dev=%v → prod=%v\n", 
				field.Name, devFieldValue.Interface(), prodFieldValue.Interface())
			differences++
		}
	}

	if differences == 0 {
		fmt.Printf("      ✅ 无差异\n")
	} else {
		fmt.Printf("      📈 发现 %d 处差异\n", differences)
	}
	fmt.Println()
}

// tool3ConfigExport 工具3: 配置导出工具
func tool3ConfigExport() {
	ctx := context.Background()
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 创建配置管理器失败: %v", err)
		return
	}

	config := manager.GetConfig()

	// 导出为JSON
	if err := exportToJSON(config, "config_export.json"); err != nil {
		log.Printf("❌ JSON导出失败: %v", err)
	} else {
		fmt.Printf("  ✅ 配置已导出为 JSON: config_export.json\n")
	}

	// 导出配置摘要
	if err := exportSummary(config, "config_summary.txt"); err != nil {
		log.Printf("❌ 摘要导出失败: %v", err)
	} else {
		fmt.Printf("  ✅ 配置摘要已导出: config_summary.txt\n")
	}

	// 导出环境变量格式
	if err := exportAsEnvVars(config, "config.env"); err != nil {
		log.Printf("❌ 环境变量导出失败: %v", err)
	} else {
		fmt.Printf("  ✅ 环境变量格式已导出: config.env\n")
	}
}

// exportToJSON 导出配置为JSON格式
func exportToJSON(config *goconfig.SingleConfig, filename string) error {
	// 创建导出结构，移除敏感信息
	exportConfig := struct {
		Server struct {
			Addr       string `json:"addr"`
			ServerName string `json:"server_name"`
			DataDriver string `json:"data_driver"`
		} `json:"server"`
		MySQL struct {
			Host           string `json:"host"`
			Port           string `json:"port"`
			Username       string `json:"username"`
			Dbname         string `json:"db_name"`
			MaxIdleConns   int    `json:"max_idle_conns"`
			MaxOpenConns   int    `json:"max_open_conns"`
		} `json:"mysql"`
		Redis struct {
			Addr     string `json:"addr"`
			DB       int    `json:"db"`
			PoolSize int    `json:"pool_size"`
		} `json:"redis"`
		Zap struct {
			Level       string `json:"level"`
			Format      string `json:"format"`
			Development bool   `json:"development"`
		} `json:"zap"`
	}{}

	// 复制非敏感数据
	exportConfig.Server.Addr = config.Server.Addr
	exportConfig.Server.ServerName = config.Server.ServerName
	exportConfig.Server.DataDriver = config.Server.DataDriver

	exportConfig.MySQL.Host = config.MySQL.Host
	exportConfig.MySQL.Port = config.MySQL.Port
	exportConfig.MySQL.Username = config.MySQL.Username
	exportConfig.MySQL.Dbname = config.MySQL.Dbname
	exportConfig.MySQL.MaxIdleConns = config.MySQL.MaxIdleConns
	exportConfig.MySQL.MaxOpenConns = config.MySQL.MaxOpenConns

	exportConfig.Redis.Addr = config.Redis.Addr
	exportConfig.Redis.DB = config.Redis.DB
	exportConfig.Redis.PoolSize = config.Redis.PoolSize

	exportConfig.Zap.Level = config.Zap.Level
	exportConfig.Zap.Format = config.Zap.Format
	exportConfig.Zap.Development = config.Zap.Development

	data, err := json.MarshalIndent(exportConfig, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// exportSummary 导出配置摘要
func exportSummary(config *goconfig.SingleConfig, filename string) error {
	summary := fmt.Sprintf(`配置摘要报告
=================

服务配置:
- 名称: %s
- 地址: %s
- 数据驱动: %s

数据库配置:
- 主机: %s:%s
- 数据库: %s
- 用户: %s
- 连接池: %d-%d

Redis配置:
- 地址: %s
- 数据库: %d
- 连接池大小: %d

日志配置:
- 级别: %s
- 格式: %s
- 开发模式: %t

生成时间: %s
`, 
		config.Server.ServerName, config.Server.Addr, config.Server.DataDriver,
		config.MySQL.Host, config.MySQL.Port, config.MySQL.Dbname, config.MySQL.Username,
		config.MySQL.MaxIdleConns, config.MySQL.MaxOpenConns,
		config.Redis.Addr, config.Redis.DB, config.Redis.PoolSize,
		config.Zap.Level, config.Zap.Format, config.Zap.Development,
		"2025-11-08 12:00:00")

	return os.WriteFile(filename, []byte(summary), 0644)
}

// exportAsEnvVars 导出为环境变量格式
func exportAsEnvVars(config *goconfig.SingleConfig, filename string) error {
	envVars := fmt.Sprintf(`# 服务配置环境变量
SERVER_ADDR=%s
SERVER_NAME=%s
SERVER_DATA_DRIVER=%s

# 数据库配置环境变量
DB_HOST=%s
DB_PORT=%s
DB_NAME=%s
DB_USER=%s
DB_MAX_IDLE_CONNS=%d
DB_MAX_OPEN_CONNS=%d

# Redis配置环境变量
REDIS_ADDR=%s
REDIS_DB=%d
REDIS_POOL_SIZE=%d

# 日志配置环境变量
LOG_LEVEL=%s
LOG_FORMAT=%s
LOG_DEVELOPMENT=%t
`,
		config.Server.Addr, config.Server.ServerName, config.Server.DataDriver,
		config.MySQL.Host, config.MySQL.Port, config.MySQL.Dbname, config.MySQL.Username,
		config.MySQL.MaxIdleConns, config.MySQL.MaxOpenConns,
		config.Redis.Addr, config.Redis.DB, config.Redis.PoolSize,
		config.Zap.Level, config.Zap.Format, config.Zap.Development)

	return os.WriteFile(filename, []byte(envVars), 0644)
}

// tool4ConfigDiagnosis 工具4: 配置诊断工具
func tool4ConfigDiagnosis() {
	ctx := context.Background()
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 创建配置管理器失败: %v", err)
		return
	}

	config := manager.GetConfig()

	fmt.Printf("  🩺 配置健康诊断:\n")

	// 诊断服务配置
	diagnoseServer(config.Server)

	// 诊断数据库配置
	diagnoseMySQL(config.MySQL)

	// 诊断Redis配置
	diagnoseRedis(config.Redis)

	// 诊断日志配置
	diagnoseZap(config.Zap)

	// 性能建议
	fmt.Printf("    💡 性能优化建议:\n")
	performanceAdvice(config)
}

// diagnoseServer 诊断服务器配置
func diagnoseServer(server interface{}) {
	// 使用反射访问字段
	v := reflect.ValueOf(server)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fmt.Printf("    🖥️ 服务器诊断:\n")
	
	// 检查地址格式
	addrField := v.FieldByName("Addr")
	if addrField.IsValid() && addrField.String() != "" {
		addr := addrField.String()
		if strings.Contains(addr, ":") {
			fmt.Printf("      ✅ 地址格式正确: %s\n", addr)
		} else {
			fmt.Printf("      ⚠️ 地址格式可能有问题: %s\n", addr)
		}
	}
}

// diagnoseMySQL 诊断MySQL配置
func diagnoseMySQL(mysql interface{}) {
	v := reflect.ValueOf(mysql)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fmt.Printf("    🗄️ MySQL诊断:\n")
	
	// 检查连接池配置
	maxIdleField := v.FieldByName("MaxIdleConns")
	maxOpenField := v.FieldByName("MaxOpenConns")
	
	if maxIdleField.IsValid() && maxOpenField.IsValid() {
		maxIdle := int(maxIdleField.Int())
		maxOpen := int(maxOpenField.Int())
		
		if maxIdle > maxOpen {
			fmt.Printf("      ⚠️ 空闲连接数(%d) > 最大连接数(%d)\n", maxIdle, maxOpen)
		} else {
			fmt.Printf("      ✅ 连接池配置合理: %d/%d\n", maxIdle, maxOpen)
		}
	}
}

// diagnoseRedis 诊断Redis配置
func diagnoseRedis(redis interface{}) {
	v := reflect.ValueOf(redis)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fmt.Printf("    📦 Redis诊断:\n")
	
	// 检查DB索引
	dbField := v.FieldByName("DB")
	if dbField.IsValid() {
		db := int(dbField.Int())
		if db >= 0 && db <= 15 {
			fmt.Printf("      ✅ DB索引有效: %d\n", db)
		} else {
			fmt.Printf("      ❌ DB索引无效: %d (应该在0-15之间)\n", db)
		}
	}
}

// diagnoseZap 诊断日志配置
func diagnoseZap(zapConfig interface{}) {
	v := reflect.ValueOf(zapConfig)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fmt.Printf("    📝 日志诊断:\n")
	
	// 检查日志级别
	levelField := v.FieldByName("Level")
	devField := v.FieldByName("Development")
	
	if levelField.IsValid() && devField.IsValid() {
		level := levelField.String()
		isDev := devField.Bool()
		
		if isDev && level == "error" {
			fmt.Printf("      ⚠️ 开发模式使用error级别可能不合适\n")
		} else if !isDev && level == "debug" {
			fmt.Printf("      ⚠️ 生产模式使用debug级别可能影响性能\n")
		} else {
			fmt.Printf("      ✅ 日志级别配置合理: %s (开发模式: %t)\n", level, isDev)
		}
	}
}

// performanceAdvice 性能建议
func performanceAdvice(config *goconfig.SingleConfig) {
	advice := []string{}

	// MySQL连接池建议
	if config.MySQL.MaxOpenConns > 0 && config.MySQL.MaxOpenConns < 50 {
		advice = append(advice, "考虑增加MySQL最大连接数以提高并发性能")
	}

	// Redis连接池建议
	if config.Redis.PoolSize > 0 && config.Redis.PoolSize < 100 {
		advice = append(advice, "考虑增加Redis连接池大小")
	}

	// 日志性能建议
	if config.Zap.Format == "json" && config.Zap.Development {
		advice = append(advice, "开发环境推荐使用console格式以提高可读性")
	}

	if len(advice) == 0 {
		fmt.Printf("      ✅ 当前配置性能表现良好\n")
	} else {
		for _, tip := range advice {
			fmt.Printf("      💡 %s\n", tip)
		}
	}
}

// tool5ConfigTemplateGenerator 工具5: 配置模板生成器
func tool5ConfigTemplateGenerator() {
	templates := map[string]string{
		"微服务模板": generateMicroserviceTemplate(),
		"Web应用模板": generateWebAppTemplate(),
		"批处理模板": generateBatchTemplate(),
	}

	fmt.Printf("  📝 生成配置模板:\n")

	for name, template := range templates {
		filename := fmt.Sprintf("template_%s.yaml", 
			strings.ToLower(strings.ReplaceAll(name, "模板", "")))
		
		if err := os.WriteFile(filename, []byte(template), 0644); err != nil {
			log.Printf("    ❌ 生成 %s 失败: %v", name, err)
		} else {
			fmt.Printf("    ✅ %s: %s\n", name, filename)
		}
	}

	// 生成README
	readme := generateTemplateReadme()
	if err := os.WriteFile("TEMPLATE_README.md", []byte(readme), 0644); err != nil {
		log.Printf("    ❌ 生成README失败: %v", err)
	} else {
		fmt.Printf("    📖 使用说明: TEMPLATE_README.md\n")
	}
}

// generateMicroserviceTemplate 生成微服务模板
func generateMicroserviceTemplate() string {
	return `# 微服务配置模板
server:
  addr: '0.0.0.0:8080'
  server-name: 'microservice-name'
  context-path: '/api/v1'
  handle-method-not-allowed: true
  data-driver: 'mysql'

mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'microservice_user'
  password: '${DB_PASSWORD}'
  db-name: 'microservice_db'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'
  max-idle-conns: 10
  max-open-conns: 100
  conn-max-idle-time: 30
  conn-max-life-time: 300

redis:
  addr: '127.0.0.1:6379'
  password: '${REDIS_PASSWORD}'
  db: 0
  pool-size: 100
  min-idle-conns: 10
  max-retries: 3

zap:
  level: 'info'
  format: 'json'
  prefix: '[MICROSERVICE]'
  director: 'logs'
  development: false
  log-in-console: false

consul:
  addr: '127.0.0.1:8500'
  register-interval: 30

jwt:
  signing-key: '${JWT_SECRET}'
  expires-time: 604800
  buffer-time: 86400
  use-multipoint: false
`
}

// generateWebAppTemplate 生成Web应用模板
func generateWebAppTemplate() string {
	return `# Web应用配置模板
server:
  addr: '0.0.0.0:3000'
  server-name: 'webapp-name'
  context-path: ''
  handle-method-not-allowed: true
  data-driver: 'mysql'

mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'webapp_user'
  password: '${DB_PASSWORD}'
  db-name: 'webapp_db'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'
  max-idle-conns: 5
  max-open-conns: 50

redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 0
  pool-size: 50
  min-idle-conns: 5

zap:
  level: 'info'
  format: 'console'
  prefix: '[WEBAPP]'
  director: 'logs'
  development: true
  log-in-console: true

cors:
  allowed-all-origins: true
  allowed-methods:
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
  allowed-headers:
    - "Authorization"
    - "Content-Type"
  allow-credentials: true

email:
  host: 'smtp.gmail.com'
  port: 587
  from: '${EMAIL_FROM}'
  secret: '${EMAIL_SECRET}'
  is-ssl: true
`
}

// generateBatchTemplate 生成批处理模板
func generateBatchTemplate() string {
	return `# 批处理任务配置模板
server:
  addr: '0.0.0.0:8080'
  server-name: 'batch-processor'
  context-path: '/batch'
  data-driver: 'mysql'

mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'batch_user'
  password: '${DB_PASSWORD}'
  db-name: 'batch_db'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'
  max-idle-conns: 20
  max-open-conns: 200
  conn-max-idle-time: 60
  conn-max-life-time: 600

redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 1
  pool-size: 200
  min-idle-conns: 20

zap:
  level: 'info'
  format: 'json'
  prefix: '[BATCH]'
  director: '/var/log/batch'
  development: false
  log-in-console: false

# 批处理特定配置可以放在这里
# batch_config:
#   batch_size: 1000
#   max_workers: 10
#   timeout: 300
`
}

// generateTemplateReadme 生成模板使用说明
func generateTemplateReadme() string {
	return `# 配置模板使用说明

## 模板文件

- **template_微服务.yaml**: 微服务应用配置模板
- **template_web应用.yaml**: Web应用配置模板  
- **template_批处理.yaml**: 批处理任务配置模板

## 使用步骤

1. 选择合适的模板文件
2. 复制模板到您的项目中
3. 根据环境修改配置文件名（如：dev_config.yaml）
4. 替换模板中的占位符变量
5. 根据实际需求调整配置参数

## 环境变量占位符

模板中使用了以下环境变量占位符：

- **${DB_PASSWORD}**: 数据库密码
- **${REDIS_PASSWORD}**: Redis密码（如果需要）
- **${JWT_SECRET}**: JWT签名密钥
- **${EMAIL_FROM}**: 邮件发送地址
- **${EMAIL_SECRET}**: 邮件服务密钥

## 安全建议

1. 不要在配置文件中硬编码敏感信息
2. 使用环境变量管理密钥和密码
3. 生产环境务必更改默认配置
4. 定期轮换密钥和密码

## 性能调优

- 根据并发需求调整数据库连接池大小
- Redis连接池大小应与应用负载匹配
- 生产环境建议使用JSON格式日志
- 合理设置日志级别避免性能影响

生成时间: 2025-11-08
`
}