/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 15:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 12:49:29
 * @FilePath: \go-config\config_formatter.go
 * @Description: 配置信息格式化输出工具 - 专门用于格式化配置变更信息
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-config/pkg/gateway"
	"github.com/kamalyes/go-logger"
	"sync"
	"time"
)

var (
	// 全局自动日志开关，默认开启
	autoLogEnabled = true
	autoLogMutex   sync.RWMutex
)

// ConfigFormatter 配置格式化器
type ConfigFormatter struct {
	logger *logger.Logger
}

// NewConfigFormatter 创建配置格式化器
func NewConfigFormatter(lg ...*logger.Logger) *ConfigFormatter {
	var log *logger.Logger
	if len(lg) > 0 && lg[0] != nil {
		log = lg[0]
	} else {
		log = logger.GetGlobalLogger()
	}

	return &ConfigFormatter{
		logger: log,
	}
}

// LogConfigChanged 记录配置变更 - 主要入口函数
func (cf *ConfigFormatter) LogConfigChanged(event CallbackEvent, newConfig interface{}) {
	cf.logger.Info("🔄 配置发生变更!")
	cf.logger.Info("   📂 来源: %s", event.Source)
	cf.logger.Info("   🕐 时间: %s", event.Timestamp.Format(time.DateTime))
	cf.logger.Info("   🌍 环境: %s", event.Environment)
	cf.logger.Info("   📋 事件类型: %s", event.Type)

	// 根据配置类型记录详细信息
	switch config := newConfig.(type) {
	case *gateway.Gateway:
		cf.LogGatewayDetails(config)
	default:
		cf.logger.Info("🆕 配置已更新: %T", config)
	}

	cf.logger.Info("✅ 配置更新成功")
}

// LogGatewayDetails 记录Gateway详细配置信息
func (cf *ConfigFormatter) LogGatewayDetails(config *gateway.Gateway) {
	cf.logger.Info("🆕 新配置信息:")
	cf.logger.Info("   📌 网关名称: %s", config.Name)
	cf.logger.Info("   🔢 版本: %s", config.Version)
	cf.logger.Info("   🌐 环境: %s", config.Environment)
	cf.logger.Info("   🚦 状态: %s", formatStatus(config.Enabled))
	cf.logger.Info("   🔍 调试模式: %s", formatStatus(config.Debug))

	// HTTP服务器配置
	if config.HTTPServer != nil {
		cf.LogHTTPServer(config.HTTPServer)
	}

	// 数据库配置
	if config.Database != nil {
		cf.LogDatabase(config.Database)
	}

	// 缓存配置
	if config.Cache != nil {
		cf.LogCache(config.Cache)
	}
}

// LogHTTPServer 记录HTTP服务器配置 - 您关注的重点
func (cf *ConfigFormatter) LogHTTPServer(httpConfig *gateway.HTTPServer) {
	cf.logger.Info("   🌐 HTTP服务器:")
	cf.logger.Info("      📍 地址: %s:%d", httpConfig.Host, httpConfig.Port)
	cf.logger.Info("      🔗 端点: %s", httpConfig.GetEndpoint())
	cf.logger.Info("      ⏱️ 读取超时: %ds", httpConfig.ReadTimeout)
	cf.logger.Info("      ⏱️ 写入超时: %ds", httpConfig.WriteTimeout)
	cf.logger.Info("      🗜️ Gzip压缩: %s", formatStatus(httpConfig.EnableGzipCompress))

	if httpConfig.EnableTls {
		cf.logger.Info("      🔒 TLS启用: %s", formatStatus(httpConfig.EnableTls))
	}

	if len(httpConfig.Headers) > 0 {
		cf.logger.Info("      📋 自定义头部:")
		for key, value := range httpConfig.Headers {
			cf.logger.Info("         %s: %s", key, value)
		}
	}
}

// LogDatabase 记录数据库配置
func (cf *ConfigFormatter) LogDatabase(dbConfig *database.Database) {
	cf.logger.Info("   🗄️ 数据库配置:")

	// 获取默认提供商配置
	if provider, err := dbConfig.GetDefaultProvider(); err == nil {
		cf.logger.Info("      📍 类型: %s", provider.GetDBType())
		if provider.GetHost() != "" {
			cf.logger.Info("      📍 地址: %s:%s", provider.GetHost(), provider.GetPort())
		}
		cf.logger.Info("      💾 数据库: %s", provider.GetDBName())
		if provider.GetUsername() != "" {
			cf.logger.Info("      👤 用户: %s", provider.GetUsername())
		}
	} else {
		cf.logger.Info("      ⚠️ 默认数据库提供商配置无效: %v", err)
	}
}

// LogCache 记录缓存配置
func (cf *ConfigFormatter) LogCache(cacheConfig *cache.Cache) {
	cf.logger.Info("   💾 缓存配置:")
	cf.logger.Info("      🏷️ 类型: %s", cacheConfig.Type)
	cf.logger.Info("      🚦 启用: %s", formatStatus(cacheConfig.Enabled))
	cf.logger.Info("      ⏰ 默认TTL: %v", cacheConfig.DefaultTTL)
}

// LogEnvironmentChanged 记录环境变更
func (cf *ConfigFormatter) LogEnvironmentChanged(oldEnv, newEnv EnvironmentType) {
	cf.logger.Info("🌍 环境发生变更!")
	cf.logger.Info("   📤 旧环境: %s", oldEnv)
	cf.logger.Info("   📥 新环境: %s", newEnv)
	cf.logger.Info("   🕐 变更时间: %s", time.Now().Format(time.DateTime))

	switch newEnv {
	case EnvDevelopment:
		cf.logger.Info("🔧 切换到开发环境模式")
	case EnvProduction:
		cf.logger.Info("🚀 切换到生产环境模式")
	case EnvTest:
		cf.logger.Info("🧪 切换到测试环境模式")
	case EnvStaging:
		cf.logger.Info("🏗️ 切换到预发布环境模式")
	default:
		cf.logger.Info("❓ 切换到未知环境模式: %s", newEnv)
	}
}

// LogError 记录错误信息
func (cf *ConfigFormatter) LogError(event CallbackEvent) {
	cf.logger.Error("❌ 发生错误: %s", event.Error)
	cf.logger.Error("   📂 来源: %s", event.Source)
	cf.logger.Error("   🕐 时间: %s", event.Timestamp.Format(time.DateTime))
}

// LogValidation 记录配置验证结果
func (cf *ConfigFormatter) LogValidation(isValid bool, err error, duration time.Duration) {
	if isValid {
		cf.logger.Info("✅ 配置验证成功 (耗时: %v)", duration)
	} else {
		cf.logger.Error("❌ 配置验证失败 (耗时: %v): %v", duration, err)
	}
}

// LogServiceStartup 记录服务启动信息
func (cf *ConfigFormatter) LogServiceStartup(serviceName, endpoint, environment, version string) {
	cf.logger.Info("🚀 服务启动信息:")
	cf.logger.Info("   📌 服务名称: %s", serviceName)
	cf.logger.Info("   📍 监听地址: %s", endpoint)
	cf.logger.Info("   🌍 运行环境: %s", environment)
	cf.logger.Info("   📝 服务版本: %s", version)
}

// formatStatus 格式化状态显示
func formatStatus(enabled bool) string {
	if enabled {
		return "✅ 启用"
	}
	return "❌ 禁用"
}

// 全局实例
var globalConfigFormatter *ConfigFormatter

// GetGlobalConfigFormatter 获取全局配置格式化器
func GetGlobalConfigFormatter() *ConfigFormatter {
	if globalConfigFormatter == nil {
		globalConfigFormatter = NewConfigFormatter()
	}
	return globalConfigFormatter
}

// SetGlobalConfigFormatter 设置全局配置格式化器
func SetGlobalConfigFormatter(cf *ConfigFormatter) {
	globalConfigFormatter = cf
}

// 便利函数 - 直接使用全局实例

// LogConfigChange 记录配置变更 - 全局函数
func LogConfigChange(event CallbackEvent, newConfig interface{}) {
	if isAutoLogEnabled() {
		GetGlobalConfigFormatter().LogConfigChanged(event, newConfig)
	}
}

// LogEnvChange 记录环境变更 - 全局函数
func LogEnvChange(oldEnv, newEnv EnvironmentType) {
	if isAutoLogEnabled() {
		GetGlobalConfigFormatter().LogEnvironmentChanged(oldEnv, newEnv)
	}
}

// LogConfigError 记录配置错误 - 全局函数
func LogConfigError(event CallbackEvent) {
	if isAutoLogEnabled() {
		GetGlobalConfigFormatter().LogError(event)
	}
}

// ======================== 自动日志控制功能 ========================

// SetAutoLogEnabled 设置是否自动开启美化日志输出
func SetAutoLogEnabled(enabled bool) {
	autoLogMutex.Lock()
	defer autoLogMutex.Unlock()
	autoLogEnabled = enabled

	status := "禁用"
	if enabled {
		status = "启用"
	}
	logger.GetGlobalLogger().Info("🎨 自动美化日志输出已%s", status)
}

// IsAutoLogEnabled 检查是否开启了自动日志输出
func IsAutoLogEnabled() bool {
	autoLogMutex.RLock()
	defer autoLogMutex.RUnlock()
	return autoLogEnabled
}

// isAutoLogEnabled 内部使用的检查函数（不导出）
func isAutoLogEnabled() bool {
	return IsAutoLogEnabled()
}

// EnableAutoLog 启用自动美化日志输出（便捷函数）
func EnableAutoLog() {
	SetAutoLogEnabled(true)
}

// DisableAutoLog 禁用自动美化日志输出（便捷函数）
func DisableAutoLog() {
	SetAutoLogEnabled(false)
}

// ======================== 强制日志输出功能 ========================

// ForceLogConfigChange 强制记录配置变更（忽略自动日志开关）
func ForceLogConfigChange(event CallbackEvent, newConfig interface{}) {
	GetGlobalConfigFormatter().LogConfigChanged(event, newConfig)
}

// ForceLogEnvChange 强制记录环境变更（忽略自动日志开关）
func ForceLogEnvChange(oldEnv, newEnv EnvironmentType) {
	GetGlobalConfigFormatter().LogEnvironmentChanged(oldEnv, newEnv)
}

// ForceLogConfigError 强制记录配置错误（忽略自动日志开关）
func ForceLogConfigError(event CallbackEvent) {
	GetGlobalConfigFormatter().LogError(event)
}
