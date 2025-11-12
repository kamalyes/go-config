/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 12:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 12:00:00
 * @FilePath: \go-config\examples\gateway_hot_reload_demo.go
 * @Description: Gateway配置热更新完整演示Demo
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/gateway"
	"github.com/kamalyes/go-logger"
)

// GatewayService 网关服务
type GatewayService struct {
	configManager *goconfig.IntegratedConfigManager
	server        *http.Server
	gatewayConfig *gateway.Gateway
}

// ConfigResponse 配置响应结构
type ConfigResponse struct {
	Environment   string                 `json:"environment"`
	Timestamp     string                 `json:"timestamp"`
	ConfigVersion string                 `json:"config_version"`
	Gateway       *gateway.Gateway       `json:"gateway"`
	Metadata      map[string]interface{} `json:"metadata"`
}

// StatusResponse 状态响应结构
type StatusResponse struct {
	Status      string                 `json:"status"`
	Uptime      string                 `json:"uptime"`
	Environment string                 `json:"environment"`
	Version     string                 `json:"version"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// NewGatewayService 创建新的网关服务
func NewGatewayService(configPath string) (*GatewayService, error) {
	// 创建Gateway配置实例
	gatewayConfig := &gateway.Gateway{}

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
			gatewayConfig,
			configPath,
			goconfig.GetEnvironment(),
			"gateway", // 指定配置类型为gateway
		)
	} else {
		// 如果是文件或路径不存在，使用传统方式
		logger.GetGlobalLogger().Info("📄 使用指定配置文件: %s", configPath)

		// 创建集成配置管理器
		manager, err = goconfig.CreateAndStartIntegratedManager(
			gatewayConfig,
			configPath,
			goconfig.GetEnvironment(),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("创建配置管理器失败: %w", err)
	}

	service := &GatewayService{
		configManager: manager,
		gatewayConfig: gatewayConfig,
	}

	// 注册配置变更回调
	service.registerCallbacks()

	return service, nil
}

// registerCallbacks 注册回调函数
func (gs *GatewayService) registerCallbacks() {
	// 注册配置变更回调
	gs.configManager.RegisterConfigCallback(gs.onConfigChanged, goconfig.CallbackOptions{
		ID:       "gateway_config_callback",
		Types:    []goconfig.CallbackType{goconfig.CallbackTypeConfigChanged, goconfig.CallbackTypeReloaded},
		Priority: goconfig.CallbackPriorityHigh,
		Async:    false,
		Timeout:  5 * time.Second,
		Metadata: map[string]interface{}{
			"component": "gateway",
			"callback":  "config_changed",
		},
	})

	// 注册环境变更回调
	gs.configManager.RegisterEnvironmentCallback("gateway_env_callback", gs.onEnvironmentChanged,
		goconfig.CallbackPriorityHigh, false)

	// 注册错误回调
	gs.configManager.RegisterConfigCallback(gs.onError, goconfig.CallbackOptions{
		ID:       "gateway_error_callback",
		Types:    []goconfig.CallbackType{goconfig.CallbackTypeError},
		Priority: goconfig.CallbackPriorityHighest,
		Async:    true,
		Timeout:  10 * time.Second,
	})

	logger.GetGlobalLogger().Info("✅ 已注册配置热更新回调")
}

// onConfigChanged 配置变更回调处理
func (gs *GatewayService) onConfigChanged(ctx context.Context, event goconfig.CallbackEvent) error {
	start := time.Now()
	defer func() {
		logger.GetGlobalLogger().Info("⏱️ 配置变更处理耗时: %v", time.Since(start))
	}()

	// 更新配置
	if newConfig, ok := event.NewValue.(*gateway.Gateway); ok {
		gs.gatewayConfig = newConfig

		// 🆕 现在框架会自动输出美化的配置变更日志
		// 无需手动调用 goconfig.LogConfigChange(event, newConfig)

		// 如果确实需要强制输出额外的日志信息，可以使用:
		// goconfig.ForceLogConfigChange(event, newConfig)

		logger.GetGlobalLogger().Info("📋 业务层配置变更处理:")
		logger.GetGlobalLogger().Info("├── 配置文件: %s", event.Source)
		logger.GetGlobalLogger().Info("├── 服务名称: %s", newConfig.Name)
		logger.GetGlobalLogger().Info("└── HTTP端点: %s", newConfig.HTTPServer.GetEndpoint())

		// 验证新配置
		if err := newConfig.Validate(); err != nil {
			logger.GetGlobalLogger().Warn("⚠️ 配置验证失败: %v", err)
			return fmt.Errorf("新配置验证失败: %w", err)
		}

		logger.GetGlobalLogger().Info("✅ 配置更新成功")
	}

	return nil
}

// onEnvironmentChanged 环境变更回调处理
func (gs *GatewayService) onEnvironmentChanged(oldEnv, newEnv goconfig.EnvironmentType) error {
	// 使用配置格式化工具记录环境变更
	goconfig.LogEnvChange(oldEnv, newEnv)
	return nil
}

// onError 错误回调处理
func (gs *GatewayService) onError(ctx context.Context, event goconfig.CallbackEvent) error {
	// 使用配置格式化工具记录错误
	goconfig.LogConfigError(event)
	// 这里可以实现错误报警逻辑
	// 例如：发送邮件、推送消息到钉钉/企业微信等
	return nil
}

// Start 启动网关服务
func (gs *GatewayService) Start() error {
	config := gs.configManager.GetConfig().(*gateway.Gateway)

	if !config.Enabled {
		return fmt.Errorf("网关服务已禁用")
	}

	// 创建HTTP路由
	mux := http.NewServeMux()
	gs.setupRoutes(mux)

	// 创建HTTP服务器
	addr := fmt.Sprintf("%s:%d", config.HTTPServer.Host, config.HTTPServer.Port)
	gs.server = &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Duration(config.HTTPServer.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.HTTPServer.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.HTTPServer.IdleTimeout) * time.Second,
	}

	logger.GetGlobalLogger().Info("🚀 网关服务启动中...")
	logger.GetGlobalLogger().Info("   📍 监听地址: %s", addr)
	logger.GetGlobalLogger().Info("   🔗 服务端点: %s", config.HTTPServer.GetEndpoint())
	logger.GetGlobalLogger().Info("   🌍 环境: %s", config.Environment)
	logger.GetGlobalLogger().Info("   📝 版本: %s", config.Version)

	// 启动HTTP服务器
	if err := gs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP服务器启动失败: %w", err)
	}

	return nil
}

// setupRoutes 设置路由
func (gs *GatewayService) setupRoutes(mux *http.ServeMux) {
	// 配置信息端点
	mux.HandleFunc("/config", gs.handleConfig)

	// 状态检查端点
	mux.HandleFunc("/status", gs.handleStatus)
	mux.HandleFunc("/health", gs.handleHealth)

	// 配置管理端点
	mux.HandleFunc("/config/reload", gs.handleReloadConfig)
	mux.HandleFunc("/config/validate", gs.handleValidateConfig)

	// 回调管理端点
	mux.HandleFunc("/callbacks", gs.handleListCallbacks)

	// 环境管理端点
	mux.HandleFunc("/environment", gs.handleEnvironment)

	logger.GetGlobalLogger().Info("📋 HTTP路由已设置:")
	logger.GetGlobalLogger().Info("   GET  /config          - 获取完整配置信息")
	logger.GetGlobalLogger().Info("   GET  /status          - 获取服务状态")
	logger.GetGlobalLogger().Info("   GET  /health          - 健康检查")
	logger.GetGlobalLogger().Info("   POST /config/reload   - 手动重新加载配置")
	logger.GetGlobalLogger().Info("   GET  /config/validate - 验证当前配置")
	logger.GetGlobalLogger().Info("   GET  /callbacks       - 列出所有回调")
	logger.GetGlobalLogger().Info("   GET  /environment     - 获取环境信息")
}

// handleConfig 处理配置信息请求
func (gs *GatewayService) handleConfig(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	config := gs.configManager.GetConfig().(*gateway.Gateway)
	metadata := gs.configManager.GetConfigMetadata()

	// 添加请求处理时间到元数据
	metadata["request_duration"] = time.Since(startTime)
	metadata["request_path"] = r.URL.Path
	metadata["request_method"] = r.Method
	metadata["client_ip"] = getClientIP(r)

	response := ConfigResponse{
		Environment:   string(gs.configManager.GetEnvironment()),
		Timestamp:     time.Now().Format(time.RFC3339),
		ConfigVersion: config.Version,
		Gateway:       config,
		Metadata:      metadata,
	}

	w.Header().Set("Content-Type", "application/json")

	// 如果启用了调试模式，返回格式化的JSON
	if config.Debug {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(response)
	} else {
		json.NewEncoder(w).Encode(response)
	}

	logger.GetGlobalLogger().Info("📋 /config - 客户端: %s, 耗时: %v", getClientIP(r), time.Since(startTime))
}

// handleStatus 处理状态检查请求
func (gs *GatewayService) handleStatus(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	config := gs.configManager.GetConfig().(*gateway.Gateway)

	uptime := time.Since(startTime) // 简化版本，实际应该记录真实启动时间

	response := StatusResponse{
		Status:      "healthy",
		Uptime:      uptime.String(),
		Environment: config.Environment,
		Version:     config.Version,
		Metadata: map[string]interface{}{
			"http_server": map[string]interface{}{
				"enabled":  config.HTTPServer.EnableHttp,
				"endpoint": config.HTTPServer.GetEndpoint(),
				"gzip":     config.HTTPServer.EnableGzipCompress,
			},
			"debug_mode": config.Debug,
			"timestamp":  time.Now().Format(time.RFC3339),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	logger.GetGlobalLogger().Info("📊 /status - 客户端: %s, 耗时: %v", getClientIP(r), time.Since(startTime))
}

// handleHealth 处理健康检查请求
func (gs *GatewayService) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
}

// handleReloadConfig 处理手动重新加载配置请求
func (gs *GatewayService) handleReloadConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	startTime := time.Now()
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	err := gs.configManager.ReloadConfig(ctx)
	duration := time.Since(startTime)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  false,
			"error":    err.Error(),
			"duration": duration.String(),
		})
		logger.GetGlobalLogger().Error("❌ 配置重新加载失败: %v, 耗时: %v", err, duration)
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  true,
			"message":  "配置重新加载成功",
			"duration": duration.String(),
		})
		logger.GetGlobalLogger().Info("✅ 配置重新加载成功, 耗时: %v", duration)
	}
}

// handleValidateConfig 处理配置验证请求
func (gs *GatewayService) handleValidateConfig(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	err := gs.configManager.ValidateConfig()
	duration := time.Since(startTime)

	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"valid":     err == nil,
		"duration":  duration.String(),
		"timestamp": time.Now().Format(time.RFC3339),
	}

	if err != nil {
		response["error"] = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(response)

	if err != nil {
		logger.GetGlobalLogger().Error("❌ 配置验证失败: %v", err)
	} else {
		logger.GetGlobalLogger().Info("✅ 配置验证成功")
	}
}

// handleListCallbacks 处理列出回调请求
func (gs *GatewayService) handleListCallbacks(w http.ResponseWriter, r *http.Request) {
	configCallbacks := gs.configManager.GetHotReloader().ListCallbacks()
	envCallbacks := gs.configManager.GetEnvironmentManager().ListCallbacks()

	response := map[string]interface{}{
		"config_callbacks": configCallbacks,
		"env_callbacks":    envCallbacks,
		"timestamp":        time.Now().Format(time.RFC3339),
		"total_count":      len(configCallbacks) + len(envCallbacks),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleEnvironment 处理环境信息请求
func (gs *GatewayService) handleEnvironment(w http.ResponseWriter, r *http.Request) {
	env := gs.configManager.GetEnvironment()

	response := map[string]interface{}{
		"current_environment": env,
		"context_key":         goconfig.GetContextKey(),
		"available_environments": []goconfig.EnvironmentType{
			goconfig.EnvDevelopment,
			goconfig.EnvTest,
			goconfig.EnvStaging,
			goconfig.EnvProduction,
		},
		"timestamp": time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Stop 停止网关服务
func (gs *GatewayService) Stop() error {
	if gs.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := gs.server.Shutdown(ctx); err != nil {
			return err
		}
	}

	if gs.configManager != nil {
		return gs.configManager.Stop()
	}

	return nil
}

// 辅助函数

// getClientIP 获取客户端IP
func getClientIP(r *http.Request) string {
	// 尝试从 X-Forwarded-For 头获取
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}
	// 尝试从 X-Real-IP 头获取
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	// 使用远程地址
	return r.RemoteAddr
}

// main 主函数
func main() {
	// 获取配置路径
	configPath := getConfigPath()

	// 设置环境变量
	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development")
		logger.GetGlobalLogger().Info("🌍 设置默认环境: development")
	}

	// 🆕 演示自动日志控制功能
	logger.GetGlobalLogger().Info("📚 演示自动日志控制功能:")
	logger.GetGlobalLogger().Info("├── 当前自动日志状态: %t", goconfig.IsAutoLogEnabled())

	// 演示如何控制自动日志输出
	if os.Getenv("DISABLE_AUTO_LOG") == "true" {
		logger.GetGlobalLogger().Info("├── 🔇 禁用自动日志输出")
		goconfig.DisableAutoLog()
	} else {
		logger.GetGlobalLogger().Info("├── 🎨 启用自动日志输出（推荐）")
		goconfig.EnableAutoLog()
	}

	logger.GetGlobalLogger().Info("└── 最终自动日志状态: %t", goconfig.IsAutoLogEnabled())

	logger.GetGlobalLogger().Info("🔧 当前环境: %s", goconfig.GetEnvironment())

	// 创建并启动网关服务
	service, err := NewGatewayService(configPath)
	if err != nil {
		logger.GetGlobalLogger().Fatal("创建网关服务失败: %v", err)
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
	config := service.configManager.GetConfig().(*gateway.Gateway)
	endpoint := config.HTTPServer.GetEndpoint()

	logger.GetGlobalLogger().Info("\n🎉 Gateway热更新演示服务已启动!")
	logger.GetGlobalLogger().Info("📍 服务地址: %s", endpoint)
	logger.GetGlobalLogger().Info("\n📋 可用的API端点:")
	logger.GetGlobalLogger().Info("   %s/config          - 获取完整配置信息", endpoint)
	logger.GetGlobalLogger().Info("   %s/status          - 获取服务状态", endpoint)
	logger.GetGlobalLogger().Info("   %s/health          - 健康检查", endpoint)
	logger.GetGlobalLogger().Info("   %s/config/reload   - 手动重新加载配置 (POST)", endpoint)
	logger.GetGlobalLogger().Info("   %s/config/validate - 验证当前配置", endpoint)
	logger.GetGlobalLogger().Info("   %s/callbacks       - 列出所有回调", endpoint)
	logger.GetGlobalLogger().Info("   %s/environment     - 获取环境信息", endpoint)

	logger.GetGlobalLogger().Info("\n🔥 热更新测试方法:")
	logger.GetGlobalLogger().Info("   1. 修改配置文件")
	logger.GetGlobalLogger().Info("   2. 观察控制台输出的回调日志")
	logger.GetGlobalLogger().Info("   3. 访问 %s/config 查看配置变化", endpoint)
	logger.GetGlobalLogger().Info("   4. 使用 curl -X POST %s/config/reload 手动重载", endpoint)

	logger.GetGlobalLogger().Info("\n🌍 环境变量测试:")
	logger.GetGlobalLogger().Info("   修改 APP_ENV 环境变量 (development, test, staging, production)")

	logger.GetGlobalLogger().Info("\n⚡ 按 Ctrl+C 优雅退出")

	// 等待退出信号
	<-sigChan
	logger.GetGlobalLogger().Info("\n🛑 接收到退出信号，正在优雅关闭...")

	// 停止服务
	if err := service.Stop(); err != nil {
		logger.GetGlobalLogger().Error("停止服务失败: %v", err)
	} else {
		logger.GetGlobalLogger().Info("✅ 服务已优雅关闭")
	}
}

// getConfigPath 获取配置路径
func getConfigPath() string {
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
