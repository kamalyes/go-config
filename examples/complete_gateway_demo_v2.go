/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-13 12:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-13 12:00:00
 * @FilePath: \go-config\examples\complete_gateway_demo_v2.go
 * @Description: 完整的Gateway演示Demo，包含数据库、Redis、RPC、所有中间件和Swagger
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
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
	"strconv"
	"syscall"
	"time"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-config/pkg/gateway"
	"github.com/kamalyes/go-logger"

	// 导入中间件包
	"github.com/kamalyes/go-config/pkg/banner"
	"github.com/kamalyes/go-config/pkg/cors"
	"github.com/kamalyes/go-config/pkg/health"
	"github.com/kamalyes/go-config/pkg/logging"
	"github.com/kamalyes/go-config/pkg/middleware"
	"github.com/kamalyes/go-config/pkg/monitoring"
	"github.com/kamalyes/go-config/pkg/ratelimit"
	"github.com/kamalyes/go-config/pkg/recovery"
	"github.com/kamalyes/go-config/pkg/requestid"
	"github.com/kamalyes/go-config/pkg/security"
	"github.com/kamalyes/go-config/pkg/swagger"
	"github.com/kamalyes/go-config/pkg/timeout"
)

// 常量定义
const (
	HeaderRequestID     = "X-Request-ID"
	HeaderContentType   = "Content-Type"
	ContentTypeJSON     = "application/json"
	ContentTypeHTML     = "text/html"
	ContentTypePlain    = "text/plain"
)

// CompleteGatewayConfig 完整的网关配置结构
type CompleteGatewayConfig struct {
	Gateway    *gateway.Gateway    `mapstructure:"gateway"    yaml:"gateway"    json:"gateway"`
	Database   *database.Database  `mapstructure:"database"   yaml:"database"   json:"database"`
	Cache      *cache.Cache        `mapstructure:"cache"      yaml:"cache"      json:"cache"`
	Redis      *cache.Redis        `mapstructure:"redis"      yaml:"redis"      json:"redis"`
	CORS       *cors.Cors          `mapstructure:"cors"       yaml:"cors"       json:"cors"`
	Swagger    *swagger.Swagger    `mapstructure:"swagger"    yaml:"swagger"    json:"swagger"`
	Health     *health.Health      `mapstructure:"health"     yaml:"health"     json:"health"`
	Monitoring *monitoring.Monitoring `mapstructure:"monitoring" yaml:"monitoring" json:"monitoring"`
	Banner     *banner.Banner      `mapstructure:"banner"     yaml:"banner"     json:"banner"`
	Logging    *logging.Logging    `mapstructure:"logging"    yaml:"logging"    json:"logging"`
	Security   *security.Security  `mapstructure:"security"   yaml:"security"   json:"security"`
	RateLimit  *ratelimit.RateLimit `mapstructure:"ratelimit" yaml:"ratelimit" json:"ratelimit"`
	Recovery   *recovery.Recovery  `mapstructure:"recovery"   yaml:"recovery"   json:"recovery"`
	RequestID  *requestid.RequestID `mapstructure:"requestid" yaml:"requestid" json:"requestid"`
	Timeout    *timeout.Timeout    `mapstructure:"timeout"    yaml:"timeout"    json:"timeout"`
	Middleware *middleware.Middleware `mapstructure:"middleware" yaml:"middleware" json:"middleware"`
}

// CompleteGatewayService 完整网关服务
type CompleteGatewayService struct {
	configManager *goconfig.IntegratedConfigManager
	server        *http.Server
	config        *CompleteGatewayConfig
	startTime     time.Time
}

// User 用户模型（示例）
type User struct {
	ID       int    `json:"id" example:"1"`
	Username string `json:"username" example:"admin"`
	Email    string `json:"email" example:"admin@example.com"`
	Status   string `json:"status" example:"active"`
}

// ApiResponse 统一API响应结构
type ApiResponse struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	TraceID string      `json:"trace_id,omitempty"`
}

// NewCompleteGatewayService 创建新的完整网关服务
func NewCompleteGatewayService(configPath string) (*CompleteGatewayService, error) {
	// 创建默认配置实例
	defaultCache := cache.Default()
	
	// 创建完整配置实例
	config := &CompleteGatewayConfig{
		Gateway:    &gateway.Gateway{},
		Database:   database.NewDatabase(),
		Cache:      defaultCache,
		Redis:      cache.DefaultRedis(),
		CORS:       cors.Default(),
		Swagger:    &swagger.Swagger{},
		Health:     &health.Health{},
		Monitoring: monitoring.Default(),
		Banner:     banner.Default(),
		Logging:    &logging.Logging{},
		Security:   &security.Security{},
		RateLimit:  &ratelimit.RateLimit{},
		Recovery:   &recovery.Recovery{},
		RequestID:  &requestid.RequestID{},
		Timeout:    &timeout.Timeout{},
		Middleware: &middleware.Middleware{},
	}

	var manager *goconfig.IntegratedConfigManager
	var err error

	// 检查configPath是文件还是目录
	if stat, statErr := os.Stat(configPath); statErr == nil && stat.IsDir() {
		// 如果是目录，使用自动发现
		logger.GetGlobalLogger().Info("🔍 使用自动发现模式，搜索路径: %s", configPath)

		// 显示可用配置文件
		_, scanErr := goconfig.ScanAndDisplayConfigs(configPath, goconfig.GetEnvironment())
		if scanErr != nil {
			logger.GetGlobalLogger().Warn("⚠️ 扫描配置文件时出错: %v", scanErr)
		}

		// 使用自动发现创建管理器
		manager, err = goconfig.CreateAndStartIntegratedManagerWithAutoDiscovery(
			config,
			configPath,
			goconfig.GetEnvironment(),
			"complete-gateway", // 指定配置类型为完整网关
		)
	} else {
		// 如果是文件或路径不存在，使用传统方式
		logger.GetGlobalLogger().Info("📄 使用指定配置文件: %s", configPath)

		// 创建集成配置管理器
		manager, err = goconfig.CreateAndStartIntegratedManager(
			config,
			configPath,
			goconfig.GetEnvironment(),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("创建配置管理器失败: %w", err)
	}

	service := &CompleteGatewayService{
		configManager: manager,
		config:        config,
		startTime:     time.Now(),
	}

	// 注册配置变更回调
	service.registerCallbacks()

	return service, nil
}

// registerCallbacks 注册回调函数
func (cgs *CompleteGatewayService) registerCallbacks() {
	// 注册配置变更回调
	cgs.configManager.RegisterConfigCallback(cgs.onConfigChanged, goconfig.CallbackOptions{
		ID:       "complete_gateway_config_callback",
		Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged, goconfig.CallbackTypeReloaded},
		Priority: goconfig.CallbackPriorityHigh,
		Async:    false,
		Timeout:  10 * time.Second,
		Metadata: map[string]interface{}{
			"component": "complete_gateway",
			"callback":  "config_changed",
		},
	})

	// 注册环境变更回调
	cgs.configManager.RegisterEnvironmentCallback("complete_gateway_env_callback", cgs.onEnvironmentChanged,
		goconfig.CallbackPriorityHigh, false)

	// 注册错误回调
	cgs.configManager.RegisterConfigCallback(cgs.onError, goconfig.CallbackOptions{
		ID:       "complete_gateway_error_callback",
		Types:    []goconfig.CallbackType{goconfig.CallbackTypeError},
		Priority: goconfig.CallbackPriorityHighest,
		Async:    true,
		Timeout:  10 * time.Second,
	})

	logger.GetGlobalLogger().Info("✅ 已注册完整网关配置热更新回调")
}

// onConfigChanged 配置变更回调处理
func (cgs *CompleteGatewayService) onConfigChanged(ctx context.Context, event goconfig.CallbackEvent) error {
	start := time.Now()
	defer func() {
		logger.GetGlobalLogger().Info("⏱️ 配置变更处理耗时: %v", time.Since(start))
	}()

	// 更新配置
	if newConfig, ok := event.NewValue.(*CompleteGatewayConfig); ok {
		cgs.config = newConfig

		logger.GetGlobalLogger().Info("📋 完整网关配置变更处理:")
		logger.GetGlobalLogger().Info("├── 配置文件: %s", event.Source)
		logger.GetGlobalLogger().Info("├── 服务名称: %s", newConfig.Gateway.Name)
		logger.GetGlobalLogger().Info("├── HTTP端点: %s", newConfig.Gateway.HTTPServer.GetEndpoint())
		logger.GetGlobalLogger().Info("├── 数据库类型: %s", newConfig.Database.Type)
		logger.GetGlobalLogger().Info("├── 缓存类型: %s", newConfig.Cache.Type)
		logger.GetGlobalLogger().Info("├── Swagger启用: %t", newConfig.Swagger.Enabled)
		logger.GetGlobalLogger().Info("└── 监控启用: %t", newConfig.Monitoring.Enabled)

		// 验证新配置
		if err := cgs.validateConfig(newConfig); err != nil {
			logger.GetGlobalLogger().Warn("⚠️ 配置验证失败: %v", err)
			return fmt.Errorf("新配置验证失败: %w", err)
		}

		logger.GetGlobalLogger().Info("✅ 完整网关配置更新成功")
	}

	return nil
}

// validateConfig 验证配置
func (cgs *CompleteGatewayService) validateConfig(config *CompleteGatewayConfig) error {
	if config.Gateway != nil {
		if err := config.Gateway.Validate(); err != nil {
			return fmt.Errorf("Gateway配置验证失败: %w", err)
		}
	}

	if config.Database != nil {
		if err := config.Database.Validate(); err != nil {
			return fmt.Errorf("Database配置验证失败: %w", err)
		}
	}

	if config.Cache != nil {
		if err := config.Cache.Validate(); err != nil {
			return fmt.Errorf("Cache配置验证失败: %w", err)
		}
	}

	if config.Redis != nil {
		if err := config.Redis.Validate(); err != nil {
			return fmt.Errorf("Redis配置验证失败: %w", err)
		}
	}

	return nil
}

// onEnvironmentChanged 环境变更回调处理
func (cgs *CompleteGatewayService) onEnvironmentChanged(oldEnv, newEnv goconfig.EnvironmentType) error {
	goconfig.LogEnvChange(oldEnv, newEnv)
	return nil
}

// onError 错误回调处理
func (cgs *CompleteGatewayService) onError(ctx context.Context, event goconfig.CallbackEvent) error {
	goconfig.LogConfigError(event)
	// 这里可以实现错误报警逻辑
	return nil
}

// Start 启动完整网关服务
func (cgs *CompleteGatewayService) Start() error {
	if !cgs.config.Gateway.Enabled {
		return fmt.Errorf("网关服务已禁用")
	}

	// 显示Banner
	if cgs.config.Banner.Enabled {
		cgs.displayBanner()
	}

	// 创建HTTP路由
	mux := http.NewServeMux()
	cgs.setupRoutes(mux)

	// 创建HTTP服务器
	addr := fmt.Sprintf("%s:%d", cgs.config.Gateway.HTTPServer.Host, cgs.config.Gateway.HTTPServer.Port)
	cgs.server = &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Duration(cgs.config.Gateway.HTTPServer.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cgs.config.Gateway.HTTPServer.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cgs.config.Gateway.HTTPServer.IdleTimeout) * time.Second,
	}

	logger.GetGlobalLogger().Info("🚀 完整网关服务启动中...")
	logger.GetGlobalLogger().Info("   📍 监听地址: %s", addr)
	logger.GetGlobalLogger().Info("   🔗 服务端点: %s", cgs.config.Gateway.HTTPServer.GetEndpoint())
	logger.GetGlobalLogger().Info("   🌍 环境: %s", cgs.config.Gateway.Environment)
	logger.GetGlobalLogger().Info("   📝 版本: %s", cgs.config.Gateway.Version)

	// 启动HTTP服务器
	if err := cgs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP服务器启动失败: %w", err)
	}

	return nil
}

// displayBanner 显示Banner
func (cgs *CompleteGatewayService) displayBanner() {
	// 如果配置中有自定义的banner模板，使用它
	if cgs.config.Banner.Template != "" {
		fmt.Println(cgs.config.Banner.Template)
		
		// 显示附加信息
		if cgs.config.Banner.Title != "" {
			fmt.Printf("标题: %s\n", cgs.config.Banner.Title)
		}
		if cgs.config.Banner.Description != "" {
			fmt.Printf("描述: %s\n", cgs.config.Banner.Description)
		}
		if cgs.config.Banner.Author != "" {
			fmt.Printf("作者: %s\n", cgs.config.Banner.Author)
		}
	} else {
		// 使用默认的banner
		banner := `
╔══════════════════════════════════════════════════════════════════╗
║                        COMPLETE GATEWAY                         ║
║                      完整网关服务演示                            ║
╠══════════════════════════════════════════════════════════════════╣
║ 🚀 功能特性:                                                     ║
║   • HTTP/GRPC 网关服务                                          ║
║   • MySQL 数据库支持                                            ║
║   • Redis 缓存支持                                              ║
║   • Swagger API 文档                                            ║
║   • 所有中间件支持                                              ║
║   • 配置热更新                                                  ║
║   • 健康检查                                                    ║
║   • 监控指标                                                    ║
║   • 安全防护                                                    ║
╚══════════════════════════════════════════════════════════════════╝
`
		fmt.Println(banner)
	}
}

// setupRoutes 设置路由
func (cgs *CompleteGatewayService) setupRoutes(mux *http.ServeMux) {
	// 应用中间件的包装函数
	middlewareChain := cgs.buildMiddlewareChain()

	// 基础信息端点
	mux.HandleFunc("/", middlewareChain(cgs.handleIndex))
	mux.HandleFunc("/config", middlewareChain(cgs.handleConfig))
	mux.HandleFunc("/status", middlewareChain(cgs.handleStatus))

	// 健康检查端点
	if cgs.config.Health.Enabled && cgs.config.Health.Path != "" {
		mux.HandleFunc(cgs.config.Health.Path, middlewareChain(cgs.handleHealth))
	}

	// 监控端点
	if cgs.config.Monitoring.Enabled {
		metricsPath := cgs.config.Monitoring.GetEndpoint()
		if metricsPath != "" {
			mux.HandleFunc(metricsPath, middlewareChain(cgs.handleMetrics))
		}
	}

	// Swagger文档端点
	if cgs.config.Swagger.Enabled {
		if cgs.config.Swagger.JSONPath != "" {
			mux.HandleFunc(cgs.config.Swagger.JSONPath, middlewareChain(cgs.handleSwaggerJSON))
		}
		mux.HandleFunc("/swagger/", middlewareChain(cgs.handleSwaggerUI))
	}

	// API端点
	mux.HandleFunc("/api/users", middlewareChain(cgs.handleUsers))
	mux.HandleFunc("/api/users/", middlewareChain(cgs.handleUserDetail))
	mux.HandleFunc("/api/cache/test", middlewareChain(cgs.handleCacheTest))
	mux.HandleFunc("/api/db/test", middlewareChain(cgs.handleDatabaseTest))

	// 配置管理端点
	mux.HandleFunc("/admin/config/reload", middlewareChain(cgs.handleReloadConfig))
	mux.HandleFunc("/admin/config/validate", middlewareChain(cgs.handleValidateConfig))

	logger.GetGlobalLogger().Info("📋 完整网关HTTP路由已设置")
}

// buildMiddlewareChain 构建中间件链
func (cgs *CompleteGatewayService) buildMiddlewareChain() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 应用请求ID中间件
			if cgs.config.RequestID.Enabled {
				cgs.applyRequestIDMiddleware(w, r)
			}

			// 应用CORS中间件
			if cgs.config.CORS.AllowedAllOrigins {
				cgs.applyCORSMiddleware(w, r)
			}

			// 应用恢复中间件
			if cgs.config.Recovery.Enabled {
				defer cgs.applyRecoveryMiddleware(w, r)
			}

			// 应用超时中间件
			if cgs.config.Timeout.Enabled {
				ctx, cancel := context.WithTimeout(r.Context(), cgs.config.Timeout.Duration)
				defer cancel()
				r = r.WithContext(ctx)
			}

			// 执行下一个处理器
			next(w, r)
		}
	}
}

// 中间件实现

// applyRequestIDMiddleware 应用请求ID中间件
func (cgs *CompleteGatewayService) applyRequestIDMiddleware(w http.ResponseWriter, r *http.Request) {
	requestID := fmt.Sprintf("%d", time.Now().UnixNano())
	w.Header().Set(HeaderRequestID, requestID)
	r.Header.Set(HeaderRequestID, requestID)
}

// applyCORSMiddleware 应用CORS中间件
func (cgs *CompleteGatewayService) applyCORSMiddleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")
	
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

// applyRecoveryMiddleware 应用恢复中间件
func (cgs *CompleteGatewayService) applyRecoveryMiddleware(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		logger.GetGlobalLogger().Error("❌ 请求处理发生panic: %v", err)
		
		response := ApiResponse{
			Code:    500,
			Message: "内部服务器错误",
			Error:   "服务器发生了未预期的错误",
			TraceID: r.Header.Get(HeaderRequestID),
		}
		
		w.Header().Set(HeaderContentType, ContentTypeJSON)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
	}
}

// 路由处理器

// handleIndex 处理首页请求
func (cgs *CompleteGatewayService) handleIndex(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(cgs.startTime)
	
	data := map[string]interface{}{
		"service":     "Complete Gateway Service",
		"version":     cgs.config.Gateway.Version,
		"environment": cgs.config.Gateway.Environment,
		"uptime":      uptime.String(),
		"features": map[string]bool{
			"database":   cgs.config.Database.Enabled,
			"cache":      cgs.config.Cache.Enabled,
			"swagger":    cgs.config.Swagger.Enabled,
			"monitoring": cgs.config.Monitoring.Enabled,
			"health":     cgs.config.Health.Enabled,
		},
		"endpoints": []string{
			"/",
			"/config",
			"/status",
			"/health",
			"/metrics",
			"/swagger/",
			"/api/users",
			"/api/cache/test",
			"/api/db/test",
		},
	}

	response := ApiResponse{
		Code:    200,
		Message: "欢迎使用完整网关服务",
		Data:    data,
		TraceID: r.Header.Get(HeaderRequestID),
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

// handleConfig 处理配置信息请求
func (cgs *CompleteGatewayService) handleConfig(w http.ResponseWriter, r *http.Request) {
	metadata := cgs.configManager.GetConfigMetadata()
	metadata["request_path"] = r.URL.Path
	metadata["request_method"] = r.Method
	metadata["client_ip"] = getClientIPAddress(r)

	response := ApiResponse{
		Code:    200,
		Message: "配置信息获取成功",
		Data: map[string]interface{}{
			"environment":     string(cgs.configManager.GetEnvironment()),
			"timestamp":       time.Now().Format(time.RFC3339),
			"config_version":  cgs.config.Gateway.Version,
			"complete_config": cgs.config,
			"metadata":        metadata,
		},
		TraceID: r.Header.Get(HeaderRequestID),
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	
	// 如果启用了调试模式，返回格式化的JSON
	if cgs.config.Gateway.Debug {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(response)
	} else {
		json.NewEncoder(w).Encode(response)
	}
}

// handleStatus 处理状态检查请求
func (cgs *CompleteGatewayService) handleStatus(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(cgs.startTime)

	data := map[string]interface{}{
		"status":      "healthy",
		"uptime":      uptime.String(),
		"environment": cgs.config.Gateway.Environment,
		"version":     cgs.config.Gateway.Version,
		"components": map[string]interface{}{
			"http_server": map[string]interface{}{
				"enabled":  cgs.config.Gateway.HTTPServer.EnableHttp,
				"endpoint": cgs.config.Gateway.HTTPServer.GetEndpoint(),
				"gzip":     cgs.config.Gateway.HTTPServer.EnableGzipCompress,
			},
			"database": map[string]interface{}{
				"enabled": cgs.config.Database.Enabled,
				"type":    cgs.config.Database.Type,
			},
			"cache": map[string]interface{}{
				"enabled": cgs.config.Cache.Enabled,
				"type":    cgs.config.Cache.Type,
			},
			"redis": map[string]interface{}{
				"enabled":      len(cgs.config.Redis.Addrs) > 0,
				"cluster_mode": cgs.config.Redis.ClusterMode,
			},
		},
		"debug_mode": cgs.config.Gateway.Debug,
		"timestamp":  time.Now().Format(time.RFC3339),
	}

	response := ApiResponse{
		Code:    200,
		Message: "服务状态正常",
		Data:    data,
		TraceID: r.Header.Get(HeaderRequestID),
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

// handleHealth 处理健康检查请求
func (cgs *CompleteGatewayService) handleHealth(w http.ResponseWriter, r *http.Request) {
	// 检查各组件健康状态
	health := map[string]string{
		"service":  "ok",
		"database": "ok", // 这里应该实际检查数据库连接
		"cache":    "ok", // 这里应该实际检查缓存连接
		"redis":    "ok", // 这里应该实际检查Redis连接
	}

	response := ApiResponse{
		Code:    200,
		Message: "健康检查通过",
		Data: map[string]interface{}{
			"status":    "ok",
			"timestamp": time.Now().Format(time.RFC3339),
			"checks":    health,
		},
		TraceID: r.Header.Get(HeaderRequestID),
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// handleMetrics 处理监控指标请求
func (cgs *CompleteGatewayService) handleMetrics(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(cgs.startTime)
	
	metrics := fmt.Sprintf(`# HELP gateway_uptime_seconds 服务运行时间
# TYPE gateway_uptime_seconds counter
gateway_uptime_seconds %.0f

# HELP gateway_requests_total 总请求数
# TYPE gateway_requests_total counter
gateway_requests_total{method="%s",endpoint="%s"} 1

# HELP gateway_config_enabled 配置启用状态
# TYPE gateway_config_enabled gauge
gateway_config_enabled{component="database"} %v
gateway_config_enabled{component="cache"} %v
gateway_config_enabled{component="swagger"} %v
gateway_config_enabled{component="monitoring"} %v
`,
		uptime.Seconds(),
		r.Method,
		r.URL.Path,
		boolToFloat(cgs.config.Database.Enabled),
		boolToFloat(cgs.config.Cache.Enabled),
		boolToFloat(cgs.config.Swagger.Enabled),
		boolToFloat(cgs.config.Monitoring.Enabled),
	)

	w.Header().Set(HeaderContentType, ContentTypePlain)
	w.Write([]byte(metrics))
}

// handleSwaggerJSON 处理Swagger JSON请求
func (cgs *CompleteGatewayService) handleSwaggerJSON(w http.ResponseWriter, r *http.Request) {
	swaggerDoc := map[string]interface{}{
		"swagger": "2.0",
		"info": map[string]interface{}{
			"title":       cgs.config.Swagger.Title,
			"description": cgs.config.Swagger.Description,
			"version":     cgs.config.Gateway.Version,
		},
		"host":     r.Host,
		"basePath": "/api",
		"schemes":  []string{"http", "https"},
		"paths": map[string]interface{}{
			"/users": map[string]interface{}{
				"get": map[string]interface{}{
					"tags":        []string{"用户"},
					"summary":     "获取用户列表",
					"description": "获取所有用户的列表",
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "成功",
						},
					},
				},
			},
		},
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(swaggerDoc)
}

// handleSwaggerUI 处理Swagger UI请求
func (cgs *CompleteGatewayService) handleSwaggerUI(w http.ResponseWriter, r *http.Request) {
	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>%s</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@3.25.0/swagger-ui.css" />
    <style>
        html { box-sizing: border-box; overflow: -moz-scrollbars-vertical; overflow-y: scroll; }
        *, *:before, *:after { box-sizing: inherit; }
        body { margin:0; background: #fafafa; }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@3.25.0/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@3.25.0/swagger-ui-standalone-preset.js"></script>
    <script>
        window.onload = function() {
            const ui = SwaggerUIBundle({
                url: '%s',
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIStandalonePreset
                ],
                plugins: [
                    SwaggerUIBundle.plugins.DownloadUrl
                ],
                layout: "StandaloneLayout"
            });
        };
    </script>
</body>
</html>`,
		cgs.config.Swagger.Title,
		cgs.config.Swagger.JSONPath,
	)

	w.Header().Set(HeaderContentType, ContentTypeHTML)
	w.Write([]byte(html))
}

// handleUsers 处理用户列表请求
func (cgs *CompleteGatewayService) handleUsers(w http.ResponseWriter, r *http.Request) {
	// 模拟从数据库获取用户列表
	users := []User{
		{ID: 1, Username: "admin", Email: "admin@example.com", Status: "active"},
		{ID: 2, Username: "user1", Email: "user1@example.com", Status: "active"},
		{ID: 3, Username: "user2", Email: "user2@example.com", Status: "inactive"},
	}

	response := ApiResponse{
		Code:    200,
		Message: "用户列表获取成功",
		Data:    users,
		TraceID: r.Header.Get(HeaderRequestID),
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

// handleUserDetail 处理用户详情请求
func (cgs *CompleteGatewayService) handleUserDetail(w http.ResponseWriter, r *http.Request) {
	// 简单的路径解析，实际项目中应该使用路由框架
	path := r.URL.Path
	userIDStr := path[len("/api/users/"):]
	
	// 解析用户ID
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response := ApiResponse{
			Code:    400,
			Message: "无效的用户ID",
			Error:   err.Error(),
			TraceID: r.Header.Get(HeaderRequestID),
		}
		w.Header().Set(HeaderContentType, ContentTypeJSON)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// 模拟从数据库获取用户详情
	user := User{
		ID:       userID,
		Username: fmt.Sprintf("user%d", userID),
		Email:    fmt.Sprintf("user%d@example.com", userID),
		Status:   "active",
	}

	response := ApiResponse{
		Code:    200,
		Message: "用户详情获取成功",
		Data:    user,
		TraceID: r.Header.Get(HeaderRequestID),
	}

	logger.GetGlobalLogger().Info("📋 获取用户详情: ID=%d", userID)

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

// handleCacheTest 处理缓存测试请求
func (cgs *CompleteGatewayService) handleCacheTest(w http.ResponseWriter, r *http.Request) {
	// 模拟缓存操作
	cacheKey := "test:cache:key"
	cacheValue := fmt.Sprintf("cached_at_%d", time.Now().Unix())

	data := map[string]interface{}{
		"cache_type":    cgs.config.Cache.Type,
		"cache_enabled": cgs.config.Cache.Enabled,
		"redis_config": map[string]interface{}{
			"addrs":        cgs.config.Redis.Addrs,
			"cluster_mode": cgs.config.Redis.ClusterMode,
		},
		"test_operation": map[string]interface{}{
			"key":   cacheKey,
			"value": cacheValue,
			"ttl":   cgs.config.Cache.DefaultTTL.String(),
		},
		"timestamp": time.Now().Format(time.RFC3339),
	}

	response := ApiResponse{
		Code:    200,
		Message: "缓存测试执行成功",
		Data:    data,
		TraceID: r.Header.Get(HeaderRequestID),
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

// handleDatabaseTest 处理数据库测试请求
func (cgs *CompleteGatewayService) handleDatabaseTest(w http.ResponseWriter, r *http.Request) {
	// 模拟数据库操作
	provider, err := cgs.config.Database.GetDefaultProvider()
	if err != nil {
		response := ApiResponse{
			Code:    500,
			Message: "数据库配置错误",
			Error:   err.Error(),
			TraceID: r.Header.Get(HeaderRequestID),
		}

		w.Header().Set(HeaderContentType, ContentTypeJSON)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	data := map[string]interface{}{
		"database_type":    cgs.config.Database.Type,
		"database_enabled": cgs.config.Database.Enabled,
		"connection_info": map[string]interface{}{
			"host":    provider.GetHost(),
			"port":    provider.GetPort(),
			"dbname":  provider.GetDBName(),
			"type":    provider.GetDBType(),
		},
		"test_query": "SELECT 1 as test_result",
		"timestamp":  time.Now().Format(time.RFC3339),
	}

	response := ApiResponse{
		Code:    200,
		Message: "数据库测试执行成功",
		Data:    data,
		TraceID: r.Header.Get(HeaderRequestID),
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

// handleReloadConfig 处理手动重新加载配置请求
func (cgs *CompleteGatewayService) handleReloadConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	startTime := time.Now()
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	err := cgs.configManager.ReloadConfig(ctx)
	duration := time.Since(startTime)

	response := ApiResponse{
		Code:    200,
		Message: "配置重新加载成功",
		Data: map[string]interface{}{
			"success":  err == nil,
			"duration": duration.String(),
		},
		TraceID: r.Header.Get(HeaderRequestID),
	}

	if err != nil {
		response.Code = 500
		response.Message = "配置重新加载失败"
		response.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		logger.GetGlobalLogger().Error("❌ 配置重新加载失败: %v", err)
	} else {
		logger.GetGlobalLogger().Info("✅ 配置重新加载成功")
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

// handleValidateConfig 处理配置验证请求
func (cgs *CompleteGatewayService) handleValidateConfig(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	err := cgs.validateConfig(cgs.config)
	duration := time.Since(startTime)

	response := ApiResponse{
		Code:    200,
		Message: "配置验证成功",
		Data: map[string]interface{}{
			"valid":     err == nil,
			"duration":  duration.String(),
			"timestamp": time.Now().Format(time.RFC3339),
		},
		TraceID: r.Header.Get(HeaderRequestID),
	}

	if err != nil {
		response.Code = 400
		response.Message = "配置验证失败"
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		logger.GetGlobalLogger().Error("❌ 配置验证失败: %v", err)
	} else {
		logger.GetGlobalLogger().Info("✅ 配置验证成功")
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

// Stop 停止完整网关服务
func (cgs *CompleteGatewayService) Stop() error {
	if cgs.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := cgs.server.Shutdown(ctx); err != nil {
			return err
		}
	}

	if cgs.configManager != nil {
		return cgs.configManager.Stop()
	}

	return nil
}

// 辅助函数

// getClientIPAddress 获取客户端IP
func getClientIPAddress(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	return r.RemoteAddr
}

// boolToFloat 布尔值转浮点数（用于监控指标）
func boolToFloat(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

// main 主函数
func main() {
	// 获取配置路径
	configPath := getConfigurationPath()

	// 设置环境变量
	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development")
		logger.GetGlobalLogger().Info("🌍 设置默认环境: development")
	}

	// 启用自动日志输出
	goconfig.EnableAutoLog()
	logger.GetGlobalLogger().Info("🔧 当前环境: %s", goconfig.GetEnvironment())

	// 创建并启动完整网关服务
	service, err := NewCompleteGatewayService(configPath)
	if err != nil {
		log.Fatalf("创建完整网关服务失败: %v", err)
	}

	// 设置优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 启动服务
	go func() {
		if err := service.Start(); err != nil {
			logger.GetGlobalLogger().Error("服务启动失败: %v", err)
		}
	}()

	// 显示使用说明
	endpoint := service.config.Gateway.HTTPServer.GetEndpoint()

	logger.GetGlobalLogger().Info("\n🎉 完整Gateway演示服务已启动!")
	logger.GetGlobalLogger().Info("📍 服务地址: %s", endpoint)
	logger.GetGlobalLogger().Info("\n📋 可用的API端点:")
	logger.GetGlobalLogger().Info("   %s/              - 服务首页", endpoint)
	logger.GetGlobalLogger().Info("   %s/config        - 完整配置信息", endpoint)
	logger.GetGlobalLogger().Info("   %s/status        - 服务状态", endpoint)
	logger.GetGlobalLogger().Info("   %s/health        - 健康检查", endpoint)
	logger.GetGlobalLogger().Info("   %s/metrics       - 监控指标", endpoint)
	logger.GetGlobalLogger().Info("   %s/swagger/      - API文档", endpoint)
	logger.GetGlobalLogger().Info("   %s/api/users     - 用户列表", endpoint)
	logger.GetGlobalLogger().Info("   %s/api/cache/test - 缓存测试", endpoint)
	logger.GetGlobalLogger().Info("   %s/api/db/test   - 数据库测试", endpoint)

	logger.GetGlobalLogger().Info("\n🔥 热更新测试方法:")
	logger.GetGlobalLogger().Info("   1. 修改配置文件")
	logger.GetGlobalLogger().Info("   2. 观察控制台输出的回调日志") 
	logger.GetGlobalLogger().Info("   3. 访问 %s/config 查看配置变化", endpoint)
	logger.GetGlobalLogger().Info("   4. 使用 curl -X POST %s/admin/config/reload 手动重载", endpoint)

	logger.GetGlobalLogger().Info("\n⚡ 按 Ctrl+C 优雅退出")

	// 等待退出信号
	<-sigChan
	logger.GetGlobalLogger().Info("\n🛑 接收到退出信号，正在优雅关闭...")

	// 停止服务
	if err := service.Stop(); err != nil {
		logger.GetGlobalLogger().Error("停止服务失败: %v", err)
	} else {
		logger.GetGlobalLogger().Info("✅ 完整网关服务已优雅关闭")
	}
}

// getConfigurationPath 获取配置路径
func getConfigurationPath() string {
	// 1. 检查命令行参数
	if len(os.Args) > 1 {
		configPath := os.Args[1]
		logger.GetGlobalLogger().Info("📄 使用命令行指定的配置路径: %s", configPath)
		return configPath
	}

	// 2. 检查环境变量
	if envConfigPath := os.Getenv("CONFIG_PATH"); envConfigPath != "" {
		logger.GetGlobalLogger().Info("📄 使用环境变量指定的配置路径: %s", envConfigPath)
		return envConfigPath
	}

	// 3. 检查当前目录是否有配置文件
	currentDir, err := os.Getwd()
	if err != nil {
		logger.GetGlobalLogger().Warn("⚠️ 获取当前目录失败: %v", err)
		currentDir = "."
	}

	// 使用配置发现器检查
	discovery := goconfig.GetGlobalConfigDiscovery()
	configInfo, err := discovery.FindBestConfigFile(currentDir, goconfig.GetEnvironment())
	if err == nil && configInfo.Exists {
		logger.GetGlobalLogger().Info("🔍 在当前目录发现配置文件: %s", configInfo.Name)
		return configInfo.Path
	}

	// 4. 如果都没有找到，返回当前目录让自动发现处理
	logger.GetGlobalLogger().Info("🔍 未找到现有配置文件，将使用自动发现模式")
	return currentDir
}