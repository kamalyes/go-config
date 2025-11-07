/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-08 16:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-08 16:30:00
 * @FilePath: \go-config\examples\web_service\main.go
 * @Description: go-config Web 服务应用示例 - 演示在实际Web服务中的使用
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	goconfig "github.com/kamalyes/go-config"
)

// AppConfig 应用配置
type AppConfig struct {
	*goconfig.SingleConfig
	Server *http.Server
}

func main() {
	fmt.Println("🚀 go-config Web服务应用示例")
	fmt.Println("===========================")

	// 创建配置文件
	if err := createWebServiceConfig(); err != nil {
		log.Fatalf("❌ 创建配置文件失败: %v", err)
	}

	// 初始化应用配置
	app, err := initializeApp()
	if err != nil {
		log.Fatalf("❌ 初始化应用失败: %v", err)
	}

	// 启动Web服务
	startWebService(app)

	// 清理
	cleanup()
}

// createWebServiceConfig 创建Web服务配置
func createWebServiceConfig() error {
	configContent := `# Web服务配置
server:
  addr: '0.0.0.0:8080'
  server-name: 'demo-web-service'
  context-path: '/api/v1'
  handle-method-not-allowed: true
  data-driver: 'mysql'

# 数据库配置
mysql:
  host: '127.0.0.1'
  port: '3306'
  username: 'root'
  password: 'password'
  db-name: 'demo_web'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  log-level: 'info'
  max-idle-conns: 10
  max-open-conns: 100
  conn-max-idle-time: 300
  conn-max-life-time: 3600

# Redis配置
redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 0
  pool-size: 100
  min-idle-conns: 10
  max-retries: 3

# JWT配置
jwt:
  signing-key: 'demo-web-service-jwt-secret-key'
  expires-time: 86400    # 24小时
  buffer-time: 3600      # 1小时
  use-multipoint: true

# 跨域配置
cors:
  allowed-all-origins: true
  allowed-methods:
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
    - "OPTIONS"
  allowed-headers:
    - "Authorization"
    - "Content-Type"
    - "Accept"
    - "X-Requested-With"
  allow-credentials: true
  max-age: "86400"

# 日志配置
zap:
  level: 'info'
  format: 'console'
  prefix: '[WEB-DEMO]'
  director: 'logs'
  link-name: 'logs/web-demo.log'
  show-line: true
  encode-level: 'LowercaseColorLevelEncoder'
  log-in-console: true
  development: true

# 邮件配置
email:
  to: 'admin@example.com'
  from: 'noreply@demo-service.com'
  host: 'smtp.example.com'
  port: 587
  is-ssl: true
  secret: 'email_password'
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

	fmt.Printf("✅ Web服务配置已创建: %s\n", configFile)
	return nil
}

// initializeApp 初始化应用
func initializeApp() (*AppConfig, error) {
	ctx := context.Background()

	// 创建配置管理器
	manager, err := goconfig.NewSingleConfigManager(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("创建配置管理器失败: %w", err)
	}

	config := manager.GetConfig()

	// 打印加载的配置
	printConfiguration(config)

	// 模拟初始化数据库连接
	if err := initializeDatabase(config); err != nil {
		return nil, fmt.Errorf("初始化数据库失败: %w", err)
	}

	// 模拟初始化Redis连接
	if err := initializeRedis(config); err != nil {
		return nil, fmt.Errorf("初始化Redis失败: %w", err)
	}

	app := &AppConfig{
		SingleConfig: config,
	}

	return app, nil
}

// printConfiguration 打印配置信息
func printConfiguration(config *goconfig.SingleConfig) {
	fmt.Println("\n📋 应用配置信息:")
	fmt.Println("================")

	fmt.Printf("🌐 服务器配置:\n")
	fmt.Printf("   - 地址: %s\n", config.Server.Addr)
	fmt.Printf("   - 服务名: %s\n", config.Server.ServerName)
	fmt.Printf("   - API路径: %s\n", config.Server.ContextPath)

	fmt.Printf("\n💾 数据库配置:\n")
	fmt.Printf("   - 主机: %s:%s\n", config.MySQL.Host, config.MySQL.Port)
	fmt.Printf("   - 数据库: %s\n", config.MySQL.Dbname)
	fmt.Printf("   - 连接池: 最大%d, 空闲%d\n", config.MySQL.MaxOpenConns, config.MySQL.MaxIdleConns)

	fmt.Printf("\n⚡ Redis配置:\n")
	fmt.Printf("   - 地址: %s\n", config.Redis.Addr)
	fmt.Printf("   - 数据库: %d\n", config.Redis.DB)
	fmt.Printf("   - 连接池: %d\n", config.Redis.PoolSize)

	fmt.Printf("\n🔐 JWT配置:\n")
	fmt.Printf("   - 过期时间: %d秒\n", config.JWT.ExpiresTime)
	fmt.Printf("   - 多点登录控制: %t\n", config.JWT.UseMultipoint)

	fmt.Printf("\n🌍 CORS配置:\n")
	fmt.Printf("   - 允许所有来源: %t\n", config.Cors.AllowedAllOrigins)
	fmt.Printf("   - 允许凭证: %t\n", config.Cors.AllowCredentials)

	fmt.Printf("\n📋 日志配置:\n")
	fmt.Printf("   - 级别: %s\n", config.Zap.Level)
	fmt.Printf("   - 格式: %s\n", config.Zap.Format)
	fmt.Printf("   - 开发模式: %t\n", config.Zap.Development)

	fmt.Printf("\n📧 邮件配置:\n")
	fmt.Printf("   - SMTP主机: %s:%d\n", config.Email.Host, config.Email.Port)
	fmt.Printf("   - 发件人: %s\n", config.Email.From)
	fmt.Printf("   - 收件人: %s\n", config.Email.To)
}

// initializeDatabase 模拟初始化数据库
func initializeDatabase(config *goconfig.SingleConfig) error {
	fmt.Printf("\n🔧 初始化数据库连接...\n")
	
	// 这里通常会创建真实的数据库连接
	// db, err := sql.Open("mysql", dsn)
	
	fmt.Printf("   ✅ 连接到MySQL: %s:%s/%s\n", 
		config.MySQL.Host, config.MySQL.Port, config.MySQL.Dbname)
	fmt.Printf("   ✅ 连接池配置: 最大%d个连接\n", config.MySQL.MaxOpenConns)
	
	// 模拟连接测试
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("   ✅ 数据库连接测试成功\n")
	
	return nil
}

// initializeRedis 模拟初始化Redis
func initializeRedis(config *goconfig.SingleConfig) error {
	fmt.Printf("\n🔧 初始化Redis连接...\n")
	
	// 这里通常会创建真实的Redis连接
	// rdb := redis.NewClient(&redis.Options{...})
	
	fmt.Printf("   ✅ 连接到Redis: %s\n", config.Redis.Addr)
	fmt.Printf("   ✅ 数据库索引: %d\n", config.Redis.DB)
	fmt.Printf("   ✅ 连接池大小: %d\n", config.Redis.PoolSize)
	
	// 模拟连接测试
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("   ✅ Redis连接测试成功\n")
	
	return nil
}

// startWebService 启动Web服务
func startWebService(app *AppConfig) {
	fmt.Printf("\n🚀 启动Web服务...\n")

	// 创建HTTP多路复用器
	mux := http.NewServeMux()

	// 注册路由处理器
	registerHandlers(mux, app)

	// 创建HTTP服务器
	app.Server = &http.Server{
		Addr:         app.SingleConfig.Server.Addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 在goroutine中启动服务器
	go func() {
		fmt.Printf("🌐 服务器正在监听: %s\n", app.Server.Addr)
		fmt.Printf("📡 API端点:\n")
		fmt.Printf("   - GET  %s/health     - 健康检查\n", app.SingleConfig.Server.ContextPath)
		fmt.Printf("   - GET  %s/config     - 配置信息\n", app.SingleConfig.Server.ContextPath)
		fmt.Printf("   - POST %s/login      - 用户登录\n", app.SingleConfig.Server.ContextPath)
		fmt.Printf("   - GET  %s/profile    - 用户资料\n", app.SingleConfig.Server.ContextPath)
		fmt.Printf("\n💡 在浏览器中访问: http://localhost%s%s/health\n", 
			app.Server.Addr[strings.Index(app.Server.Addr, ":"):], app.SingleConfig.Server.ContextPath)
		
		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("❌ HTTP服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号
	waitForShutdown(app)
}

// registerHandlers 注册路由处理器
func registerHandlers(mux *http.ServeMux, app *AppConfig) {
	contextPath := app.SingleConfig.Server.ContextPath

	// 健康检查端点
	mux.HandleFunc(contextPath+"/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		health := map[string]interface{}{
			"status":    "ok",
			"timestamp": time.Now().Unix(),
			"service":   app.SingleConfig.Server.ServerName,
			"version":   "1.0.0",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(health)
	})

	// 配置信息端点
	mux.HandleFunc(contextPath+"/config", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		config := map[string]interface{}{
			"server": map[string]interface{}{
				"name":    app.SingleConfig.Server.ServerName,
				"address": app.SingleConfig.Server.Addr,
				"context": app.SingleConfig.Server.ContextPath,
			},
			"database": map[string]interface{}{
				"host":     app.SingleConfig.MySQL.Host,
				"database": app.SingleConfig.MySQL.Dbname,
				"max_conn": app.SingleConfig.MySQL.MaxOpenConns,
			},
			"redis": map[string]interface{}{
				"address":   app.SingleConfig.Redis.Addr,
				"database":  app.SingleConfig.Redis.DB,
				"pool_size": app.SingleConfig.Redis.PoolSize,
			},
			"jwt": map[string]interface{}{
				"expires_in": app.SingleConfig.JWT.ExpiresTime,
				"multipoint": app.SingleConfig.JWT.UseMultipoint,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(config)
	})

	// 模拟登录端点
	mux.HandleFunc(contextPath+"/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 模拟JWT生成（实际应用中会使用配置的JWT密钥）
		response := map[string]interface{}{
			"success": true,
			"message": "登录成功",
			"token":   "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.demo.token",
			"expires_in": app.SingleConfig.JWT.ExpiresTime,
			"user": map[string]interface{}{
				"id":   1,
				"name": "demo_user",
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// 模拟用户资料端点
	mux.HandleFunc(contextPath+"/profile", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 模拟从Redis获取用户缓存信息
		profile := map[string]interface{}{
			"user_id": 1,
			"username": "demo_user",
			"email": app.SingleConfig.Email.To,
			"cache_info": map[string]interface{}{
				"redis_addr": app.SingleConfig.Redis.Addr,
				"cache_db":   app.SingleConfig.Redis.DB,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(profile)
	})
}

// waitForShutdown 等待关闭信号
func waitForShutdown(app *AppConfig) {
	// 创建信号通道
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 等待信号
	<-quit
	fmt.Printf("\n🛑 接收到关闭信号，正在停止服务器...\n")

	// 创建超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := app.Server.Shutdown(ctx); err != nil {
		log.Printf("❌ 服务器关闭失败: %v", err)
	} else {
		fmt.Printf("✅ 服务器已优雅关闭\n")
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
	fmt.Println("🧹 Web服务示例清理完成")
}