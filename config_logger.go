/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-28 00:23:44
 * @FilePath: \go-config\config_logger.go
 * @Description: 配置日志输出工具 - 封装配置信息的格式化输出
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"fmt"
	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-config/pkg/gateway"
	"github.com/kamalyes/go-logger"
	"reflect"
	"time"
)

// ConfigLogger 配置日志输出器
type ConfigLogger struct {
	logger *logger.Logger
}

// NewConfigLogger 创建新的配置日志输出器
func NewConfigLogger(loggerInstance ...*logger.Logger) *ConfigLogger {
	var log *logger.Logger
	if len(loggerInstance) > 0 && loggerInstance[0] != nil {
		log = loggerInstance[0]
	} else {
		// 使用默认全局日志器
		log = logger.GetGlobalLogger()
	}

	return &ConfigLogger{
		logger: log,
	}
}

// LogConfigChangeEvent 记录配置变更事件
func (cl *ConfigLogger) LogConfigChangeEvent(event CallbackEvent, newConfig interface{}) {
	cl.logger.Info("🔄 配置发生变更!")
	cl.logger.Info("   📂 来源: %s", event.Source)
	cl.logger.Info("   🕐 时间: %s", event.Timestamp.Format(time.DateTime))
	cl.logger.Info("   🌍 环境: %s", event.Environment)
	cl.logger.Info("   📋 事件类型: %s", event.Type)

	if event.Duration > 0 {
		cl.logger.Info("   ⏱️ 处理耗时: %v", event.Duration)
	}

	// 根据配置类型记录详细信息
	switch config := newConfig.(type) {
	case *gateway.Gateway:
		cl.LogGatewayConfig(config)
	default:
		cl.LogGenericConfig(config)
	}

	cl.logger.Info("✅ 配置更新成功")
}

// LogEnvironmentChangeEvent 记录环境变更事件
func (cl *ConfigLogger) LogEnvironmentChangeEvent(oldEnv, newEnv EnvironmentType) {
	cl.logger.Info("🌍 环境发生变更!")
	cl.logger.Info("   📤 旧环境: %s", oldEnv)
	cl.logger.Info("   📥 新环境: %s", newEnv)
	cl.logger.Info("   🕐 变更时间: %s", time.Now().Format(time.DateTime))

	// 根据环境类型显示不同的提示信息
	switch newEnv {
	case EnvDevelopment:
		cl.logger.Info("🔧 切换到开发环境模式")
	case EnvProduction:
		cl.logger.Info("🚀 切换到生产环境模式")
	case EnvTest:
		cl.logger.Info("🧪 切换到测试环境模式")
	case EnvStaging:
		cl.logger.Info("🏗️ 切换到预发布环境模式")
	}
}

// LogErrorEvent 记录错误事件
func (cl *ConfigLogger) LogErrorEvent(event CallbackEvent) {
	cl.logger.Error("❌ 发生错误: %s", event.Error)
	cl.logger.Error("   📂 来源: %s", event.Source)
	cl.logger.Error("   🕐 时间: %s", event.Timestamp.Format(time.DateTime))

	if event.Metadata != nil {
		if errorType, ok := event.Metadata["error_type"]; ok {
			cl.logger.Error("   🏷️ 错误类型: %v", errorType)
		}
	}
}

// LogGatewayConfig 记录Gateway配置详情
func (cl *ConfigLogger) LogGatewayConfig(config *gateway.Gateway) {
	cl.logger.Info("🆕 新配置信息:")
	cl.logger.Info("   📌 网关名称: %s", config.Name)
	cl.logger.Info("   🔢 版本: %s", config.Version)
	cl.logger.Info("   🌐 环境: %s", config.Environment)
	cl.logger.Info("   🚦 状态: %s", cl.formatEnabledStatus(config.Enabled))
	cl.logger.Info("   🔍 调试模式: %s", cl.formatEnabledStatus(config.Debug))

	// HTTP服务器配置
	if config.HTTPServer != nil {
		cl.LogHTTPServerConfig(config.HTTPServer)
	}

	// 数据库配置
	if config.Database != nil {
		cl.LogDatabaseConfig(config.Database)
	}

	// 缓存配置
	if config.Cache != nil {
		cl.LogCacheConfig(config.Cache)
	}

	// GRPC配置
	if config.GRPC != nil {
		cl.LogGRPCConfig(config.GRPC)
	}
}

// LogHTTPServerConfig 记录HTTP服务器配置
func (cl *ConfigLogger) LogHTTPServerConfig(httpConfig *gateway.HTTPServer) {
	cl.logger.Info("   🌐 HTTP服务器:")
	cl.logger.Info("      📍 地址: %s:%d", httpConfig.Host, httpConfig.Port)
	cl.logger.Info("      🔗 端点: %s", httpConfig.GetEndpoint())
	cl.logger.Info("      ⏱️ 读取超时: %ds", httpConfig.ReadTimeout)
	cl.logger.Info("      ⏱️ 写入超时: %ds", httpConfig.WriteTimeout)
	cl.logger.Info("      ⏰ 空闲超时: %ds", httpConfig.IdleTimeout)
	cl.logger.Info("      🗜️ Gzip压缩: %s", cl.formatEnabledStatus(httpConfig.EnableGzipCompress))
	cl.logger.Info("      🔒 TLS启用: %s", cl.formatEnabledStatus(httpConfig.EnableTls))

	if len(httpConfig.Headers) > 0 {
		cl.logger.Info("      📋 自定义头部:")
		for key, value := range httpConfig.Headers {
			cl.logger.Info("         %s: %s", key, value)
		}
	}
}

// LogGRPCConfig 记录GRPC配置
func (cl *ConfigLogger) LogGRPCConfig(grpcConfig *gateway.GRPC) {
	cl.logger.Info("   🔌 GRPC配置:")

	// GRPC服务端配置
	if grpcConfig.Server != nil {
		cl.logger.Info("      📍 服务端地址: %s:%d", grpcConfig.Server.Host, grpcConfig.Server.Port)
		cl.logger.Info("      🌐 网络类型: %s", grpcConfig.Server.Network)
		cl.logger.Info("      ⏱️ 连接超时: %ds", grpcConfig.Server.ConnectionTimeout)
		cl.logger.Info("      📦 最大接收消息: %d bytes", grpcConfig.Server.MaxRecvMsgSize)
		cl.logger.Info("      📦 最大发送消息: %d bytes", grpcConfig.Server.MaxSendMsgSize)
		cl.logger.Info("      � 启用反射: %s", cl.formatEnabledStatus(grpcConfig.Server.EnableReflection))
	}

	// GRPC客户端配置
	if len(grpcConfig.Clients) > 0 {
		cl.logger.Info("      👥 客户端配置: %d个", len(grpcConfig.Clients))
		for name, client := range grpcConfig.Clients {
			cl.logger.Info("         � 服务: %s", name)
			cl.logger.Info("         🎯 端点: %v", client.Endpoints)
		}
	}
}

// LogDatabaseConfig 记录数据库配置
func (cl *ConfigLogger) LogDatabaseConfig(dbConfig *database.Database) {
	cl.logger.Info("   🗄️ 数据库配置:")
	cl.logger.Info("      🚦 启用: %s", cl.formatEnabledStatus(dbConfig.Enabled))
	cl.logger.Info("      🏷️ 默认类型: %s", dbConfig.Default)

	// 获取默认提供商配置
	if provider, err := dbConfig.GetDefaultProvider(); err == nil {
		cl.logger.Info("      � 当前类型: %s", provider.GetDBType())
		if provider.GetHost() != "" {
			cl.logger.Info("      � 地址: %s:%s", provider.GetHost(), provider.GetPort())
		}
		cl.logger.Info("      � 数据库: %s", provider.GetDBName())
		if provider.GetUsername() != "" {
			cl.logger.Info("      👤 用户: %s", provider.GetUsername())
		}
		cl.logger.Info("      📊 模块名: %s", provider.GetModuleName())
	} else {
		cl.logger.Info("      ⚠️ 默认数据库提供商配置无效: %v", err)
	}
}

// LogCacheConfig 记录缓存配置
func (cl *ConfigLogger) LogCacheConfig(cacheConfig *cache.Cache) {
	cl.logger.Info("   💾 缓存配置:")
	cl.logger.Info("      🏷️ 类型: %s", cacheConfig.Type)
	cl.logger.Info("      🚦 启用: %s", cl.formatEnabledStatus(cacheConfig.Enabled))
	cl.logger.Info("      ⏰ 默认TTL: %v", cacheConfig.DefaultTTL)
	cl.logger.Info("      🔑 键前缀: %s", cacheConfig.KeyPrefix)
	cl.logger.Info("      📦 序列化: %s", cacheConfig.Serializer)

	// 根据缓存类型显示特定配置
	switch cacheConfig.Type {
	case "memory":
		if cacheConfig.Memory.Capacity > 0 {
			cl.logger.Info("      📊 内存缓存容量: %d", cacheConfig.Memory.Capacity)
			cl.logger.Info("      🧹 清理大小: %d", cacheConfig.Memory.CleanupSize)
		}
	case "ristretto":
		if cacheConfig.Ristretto.NumCounters > 0 {
			cl.logger.Info("      📊 计数器数量: %d", cacheConfig.Ristretto.NumCounters)
			cl.logger.Info("      📦 缓冲项目: %d", cacheConfig.Ristretto.BufferItems)
		}
	case "redis":
		if cacheConfig.Redis.Addr != "" {
			cl.logger.Info("      📍 Redis地址: %s", cacheConfig.Redis.Addr)
			cl.logger.Info("      💾 数据库: %d", cacheConfig.Redis.DB)
		}
	}
}

// LogConfigValidation 记录配置验证结果
func (cl *ConfigLogger) LogConfigValidation(isValid bool, err error, duration time.Duration) {
	if isValid {
		cl.logger.Info("✅ 配置验证成功 (耗时: %v)", duration)
	} else {
		cl.logger.Error("❌ 配置验证失败 (耗时: %v): %v", duration, err)
	}
}

// LogGenericConfig 记录通用配置
func (cl *ConfigLogger) LogGenericConfig(config interface{}) {
	cl.logger.Info("🆕 配置更新:")

	// 使用反射获取配置类型和基本信息
	configType := reflect.TypeOf(config)
	configValue := reflect.ValueOf(config)

	if configType.Kind() == reflect.Ptr {
		configType = configType.Elem()
		if configValue.IsValid() && !configValue.IsNil() {
			configValue = configValue.Elem()
		}
	}

	cl.logger.Info("   🏷️ 配置类型: %s", configType.Name())

	// 尝试获取一些通用字段
	if configValue.IsValid() {
		if nameField := configValue.FieldByName("Name"); nameField.IsValid() && nameField.Kind() == reflect.String {
			cl.logger.Info("   📌 名称: %s", nameField.String())
		}
		if versionField := configValue.FieldByName("Version"); versionField.IsValid() && versionField.Kind() == reflect.String {
			cl.logger.Info("   🔢 版本: %s", versionField.String())
		}
		if enabledField := configValue.FieldByName("Enabled"); enabledField.IsValid() && enabledField.Kind() == reflect.Bool {
			cl.logger.Info("   🚦 启用: %s", cl.formatEnabledStatus(enabledField.Bool()))
		}
	}
}

// LogServiceStartup 记录服务启动信息
func (cl *ConfigLogger) LogServiceStartup(serviceName, endpoint, environment, version string) {
	cl.logger.Info("🚀 服务启动信息:")
	cl.logger.Info("   📌 服务名称: %s", serviceName)
	cl.logger.Info("   📍 监听地址: %s", endpoint)
	cl.logger.Info("   🌍 运行环境: %s", environment)
	cl.logger.Info("   📝 服务版本: %s", version)
}

// LogServiceShutdown 记录服务关闭信息
func (cl *ConfigLogger) LogServiceShutdown(serviceName string, duration time.Duration) {
	cl.logger.Info("🛑 服务 %s 正在关闭... (运行时长: %v)", serviceName, duration)
}

// LogAPIEndpoints 记录API端点信息
func (cl *ConfigLogger) LogAPIEndpoints(endpoints map[string]string) {
	if len(endpoints) == 0 {
		return
	}

	cl.logger.Info("📋 可用的API端点:")
	for method, endpoint := range endpoints {
		cl.logger.Info("   %s", cl.formatEndpoint(method, endpoint))
	}
}

// LogCallbackRegistration 记录回调注册信息
func (cl *ConfigLogger) LogCallbackRegistration(callbackID string, callbackType []CallbackType, priority int) {
	cl.logger.Info("✅ 注册回调: %s", callbackID)
	cl.logger.Info("   📋 类型: %v", callbackType)
	cl.logger.Info("   📊 优先级: %d", priority)
}

// formatEnabledStatus 格式化启用状态
func (cl *ConfigLogger) formatEnabledStatus(enabled bool) string {
	if enabled {
		return "✅ 启用"
	}
	return "❌ 禁用"
}

// formatEndpoint 格式化端点信息
func (cl *ConfigLogger) formatEndpoint(method, endpoint string) string {
	methodPart := fmt.Sprintf("%-4s", method)
	return fmt.Sprintf("%s %s", methodPart, endpoint)
}

// SetLogger 设置自定义日志器
func (cl *ConfigLogger) SetLogger(logger *logger.Logger) {
	cl.logger = logger
}

// GetLogger 获取日志器实例
func (cl *ConfigLogger) GetLogger() *logger.Logger {
	return cl.logger
}

// 全局配置日志器实例
var globalConfigLogger *ConfigLogger

// GetGlobalConfigLogger 获取全局配置日志器
func GetGlobalConfigLogger() *ConfigLogger {
	if globalConfigLogger == nil {
		globalConfigLogger = NewConfigLogger()
	}
	return globalConfigLogger
}

// SetGlobalConfigLogger 设置全局配置日志器
func SetGlobalConfigLogger(logger *ConfigLogger) {
	globalConfigLogger = logger
}
