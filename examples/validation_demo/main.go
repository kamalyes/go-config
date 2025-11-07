/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-08 17:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-08 17:00:00
 * @FilePath: \go-config\examples\validation_demo\main.go
 * @Description: go-config 配置验证和错误处理示例
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
	fmt.Println("🚀 go-config 配置验证和错误处理示例")
	fmt.Println("===================================")

	// 演示正确配置
	demonstrateValidConfig()

	// 演示配置验证
	demonstrateConfigValidation()

	// 演示错误处理
	demonstrateErrorHandling()

	// 演示配置缺失处理
	demonstrateMissingConfig()

	// 清理
	cleanup()
}

// demonstrateValidConfig 演示正确配置
func demonstrateValidConfig() {
	fmt.Println("\n✅ 正确配置示例")
	fmt.Println("----------------")

	validConfig := `# 正确的配置示例
server:
  addr: '0.0.0.0:8080'
  server-name: 'validation-demo'
  context-path: '/api/v1'
  data-driver: 'mysql'

mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'root'
  password: 'secure_password'
  db-name: 'valid_database'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'
  max-idle-conns: 10
  max-open-conns: 100

redis:
  addr: '127.0.0.1:6379'
  db: 0
  pool-size: 50
  min-idle-conns: 5

jwt:
  signing-key: 'valid-jwt-secret-key-with-sufficient-length'
  expires-time: 3600
  buffer-time: 300
  use-multipoint: true
`

	if err := createConfigFile("valid_config.yaml", validConfig); err != nil {
		log.Printf("❌ 创建配置文件失败: %v", err)
		return
	}

	ctx := context.Background()
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 创建配置管理器失败: %v", err)
		return
	}

	config := manager.GetConfig()
	
	// 验证各个模块的配置
	fmt.Println("🔍 配置验证结果:")
	
	// 验证服务器配置
	if err := config.Server.Validate(); err != nil {
		fmt.Printf("   ❌ 服务器配置无效: %v\n", err)
	} else {
		fmt.Printf("   ✅ 服务器配置有效: %s\n", config.Server.Addr)
	}

	// 验证MySQL配置
	if err := config.MySQL.Validate(); err != nil {
		fmt.Printf("   ❌ MySQL配置无效: %v\n", err)
	} else {
		fmt.Printf("   ✅ MySQL配置有效: %s:%s/%s\n", 
			config.MySQL.Host, config.MySQL.Port, config.MySQL.Dbname)
	}

	// 验证Redis配置
	if err := config.Redis.Validate(); err != nil {
		fmt.Printf("   ❌ Redis配置无效: %v\n", err)
	} else {
		fmt.Printf("   ✅ Redis配置有效: %s\n", config.Redis.Addr)
	}

	// 验证JWT配置
	if err := config.JWT.Validate(); err != nil {
		fmt.Printf("   ❌ JWT配置无效: %v\n", err)
	} else {
		fmt.Printf("   ✅ JWT配置有效: 过期时间%d秒\n", config.JWT.ExpiresTime)
	}
}

// demonstrateConfigValidation 演示配置验证
func demonstrateConfigValidation() {
	fmt.Println("\n⚠️ 配置验证错误示例")
	fmt.Println("-------------------")

	// 创建包含错误的配置
	invalidConfigs := []struct {
		name   string
		config string
		desc   string
	}{
		{
			name: "missing_required_fields.yaml",
			desc: "缺少必填字段",
			config: `# 缺少必填字段的配置
server:
  # 缺少 addr 字段
  server-name: 'test-service'
  
mysql:
  # 缺少 host, username, password 等字段
  port: '3306'
  db-name: 'test_db'
`,
		},
		{
			name: "invalid_values.yaml",
			desc: "无效的字段值",
			config: `# 包含无效值的配置
server:
  addr: ''  # 空地址
  server-name: ''  # 空服务名
  data-driver: ''  # 空数据驱动

mysql:
  host: '127.0.0.1'
  port: 'invalid_port'  # 无效端口
  username: 'root'
  password: 'pass'
  db-name: 'test'
  max-idle-conns: -1  # 负数连接数
  max-open-conns: -10  # 负数连接数

redis:
  addr: '127.0.0.1:6379'
  db: 16  # 超出有效范围 (0-15)
  pool-size: 0  # 无效池大小
`,
		},
		{
			name: "jwt_weak_key.yaml",
			desc: "JWT密钥过短",
			config: `# JWT密钥不安全的配置
server:
  addr: '0.0.0.0:8080'
  server-name: 'test-service'
  data-driver: 'mysql'

jwt:
  signing-key: 'weak'  # 过短的密钥
  expires-time: 0  # 无效的过期时间
  buffer-time: -100  # 负数缓冲时间
`,
		},
	}

	for _, testCase := range invalidConfigs {
		fmt.Printf("\n📋 测试案例: %s\n", testCase.desc)
		
		if err := createConfigFile(testCase.name, testCase.config); err != nil {
			log.Printf("❌ 创建配置文件失败: %v", err)
			continue
		}

		// 尝试加载配置
		ctx := context.Background()
		manager, err := goconfig.NewSingleConfigManager(ctx, nil)
		if err != nil {
			fmt.Printf("   ❌ 配置加载失败: %v\n", err)
			continue
		}

		config := manager.GetConfig()

		// 逐个验证组件
		validateComponent("Server", config.Server.Validate())
		validateComponent("MySQL", config.MySQL.Validate())
		validateComponent("Redis", config.Redis.Validate())
		validateComponent("JWT", config.JWT.Validate())
	}
}

// demonstrateErrorHandling 演示错误处理
func demonstrateErrorHandling() {
	fmt.Println("\n🔧 错误处理策略示例")
	fmt.Println("-------------------")

	// 演示优雅的错误处理
	ctx := context.Background()

	fmt.Println("📋 演示各种错误场景:")

	// 1. 配置文件不存在
	fmt.Printf("\n1️⃣ 配置文件不存在的处理:\n")
	if err := os.RemoveAll("./resources"); err != nil {
		log.Printf("清理失败: %v", err)
	}

	_, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		fmt.Printf("   ⚠️ 预期错误 - 配置文件不存在: %v\n", err)
		fmt.Printf("   💡 解决方案: 创建默认配置文件\n")
	}

	// 2. 无效的配置格式
	fmt.Printf("\n2️⃣ 无效配置格式的处理:\n")
	invalidFormatConfig := `
invalid yaml format:
  - item1
  item2  # 缩进错误
    - item3
`
	if err := createConfigFile("invalid_format.yaml", invalidFormatConfig); err == nil {
		_, err := goconfig.NewSingleConfigManager(ctx, nil)
		if err != nil {
			fmt.Printf("   ⚠️ 预期错误 - YAML格式无效: %v\n", err)
			fmt.Printf("   💡 解决方案: 检查YAML语法，确保正确缩进\n")
		}
	}

	// 3. 空配置文件
	fmt.Printf("\n3️⃣ 空配置文件的处理:\n")
	emptyConfig := `# 空配置文件`
	if err := createConfigFile("empty_config.yaml", emptyConfig); err == nil {
		manager, err := goconfig.NewSingleConfigManager(ctx, nil)
		if err != nil {
			fmt.Printf("   ⚠️ 配置管理器创建失败: %v\n", err)
		} else {
			config := manager.GetConfig()
			fmt.Printf("   ✅ 空配置加载成功，使用默认值\n")
			fmt.Printf("   📊 服务器地址: '%s' (默认值)\n", config.Server.Addr)
		}
	}

	// 4. 部分配置缺失
	fmt.Printf("\n4️⃣ 部分配置缺失的处理:\n")
	partialConfig := `
server:
  addr: '0.0.0.0:8080'
  server-name: 'partial-service'
# MySQL 和 Redis 配置完全缺失
`
	if err := createConfigFile("partial_config.yaml", partialConfig); err == nil {
		manager, err := goconfig.NewSingleConfigManager(ctx, nil)
		if err != nil {
			fmt.Printf("   ⚠️ 配置加载失败: %v\n", err)
		} else {
			config := manager.GetConfig()
			fmt.Printf("   ✅ 部分配置加载成功\n")
			fmt.Printf("   📊 服务器: %s (已配置)\n", config.Server.Addr)
			fmt.Printf("   📊 MySQL主机: '%s' (默认值/空值)\n", config.MySQL.Host)
			
			// 显示哪些配置有效/无效
			if err := config.Server.Validate(); err != nil {
				fmt.Printf("   ❌ 服务器配置验证失败: %v\n", err)
			} else {
				fmt.Printf("   ✅ 服务器配置验证通过\n")
			}
			
			if err := config.MySQL.Validate(); err != nil {
				fmt.Printf("   ❌ MySQL配置验证失败: %v\n", err)
				fmt.Printf("   💡 建议: 添加完整的MySQL配置\n")
			}
		}
	}
}

// demonstrateMissingConfig 演示配置缺失处理
func demonstrateMissingConfig() {
	fmt.Println("\n📋 配置缺失处理策略")
	fmt.Println("-------------------")

	// 创建一个只有基本配置的文件
	basicConfig := `
server:
  addr: '0.0.0.0:8080'
  server-name: 'basic-service'
  data-driver: 'mysql'

# 其他配置项完全缺失
`

	if err := createConfigFile("basic_only.yaml", basicConfig); err != nil {
		log.Printf("❌ 创建基本配置失败: %v", err)
		return
	}

	ctx := context.Background()
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		log.Printf("❌ 配置管理器创建失败: %v", err)
		return
	}

	config := manager.GetConfig()

	fmt.Println("🔍 检查各模块配置完整性:")

	// 检查必需的配置
	requiredConfigs := []struct {
		name      string
		validator func() error
		resolver  func() string
	}{
		{
			name:      "服务器",
			validator: config.Server.Validate,
			resolver: func() string {
				if config.Server.Addr != "" {
					return fmt.Sprintf("地址: %s", config.Server.Addr)
				}
				return "配置缺失"
			},
		},
		{
			name:      "MySQL",
			validator: config.MySQL.Validate,
			resolver: func() string {
				if config.MySQL.Host != "" {
					return fmt.Sprintf("主机: %s", config.MySQL.Host)
				}
				return "配置缺失，需要手动配置"
			},
		},
		{
			name:      "Redis",
			validator: config.Redis.Validate,
			resolver: func() string {
				if config.Redis.Addr != "" {
					return fmt.Sprintf("地址: %s", config.Redis.Addr)
				}
				return "配置缺失，将使用默认连接"
			},
		},
		{
			name:      "JWT",
			validator: config.JWT.Validate,
			resolver: func() string {
				if config.JWT.SigningKey != "" {
					return "密钥已配置"
				}
				return "配置缺失，JWT功能将被禁用"
			},
		},
	}

	for _, cfg := range requiredConfigs {
		if err := cfg.validator(); err != nil {
			fmt.Printf("   ❌ %s: %s (%v)\n", cfg.name, cfg.resolver(), err)
		} else {
			fmt.Printf("   ✅ %s: %s\n", cfg.name, cfg.resolver())
		}
	}

	fmt.Println("\n💡 配置缺失处理建议:")
	fmt.Println("   1. 为缺失的配置项提供默认值")
	fmt.Println("   2. 在应用启动时检查关键配置")
	fmt.Println("   3. 对于可选配置，提供优雅的降级方案")
	fmt.Println("   4. 记录配置警告信息，便于运维监控")
}

// validateComponent 验证组件配置
func validateComponent(name string, err error) {
	if err != nil {
		fmt.Printf("   ❌ %s配置验证失败: %v\n", name, err)
	} else {
		fmt.Printf("   ✅ %s配置验证通过\n", name)
	}
}

// createConfigFile 创建配置文件
func createConfigFile(filename, content string) error {
	resourcesDir := "./resources"
	if err := os.MkdirAll(resourcesDir, os.ModePerm); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	configFile := filepath.Join(resourcesDir, "dev_config.yaml")
	if err := os.WriteFile(configFile, []byte(content), 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %w", err)
	}

	return nil
}

// cleanup 清理测试文件
func cleanup() {
	if err := os.RemoveAll("./resources"); err != nil {
		log.Printf("⚠️ 清理文件失败: %v", err)
	}
	fmt.Println("\n🧹 验证示例清理完成")
}