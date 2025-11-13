/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-13 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-13 15:10:42
 * @FilePath: \go-config\config_safe_access.go
 * @Description: 配置安全访问辅助工具，专门针对配置结构体优化
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"time"

	"github.com/kamalyes/go-toolbox/pkg/safe"
)

// ConfigSafe 专门针对配置的安全访问工具
type ConfigSafe struct {
	*safe.SafeAccess
}

// SafeConfig 创建配置安全访问器
func SafeConfig(config interface{}) *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: safe.Safe(config),
	}
}

// Health 安全访问Health配置
func (c *ConfigSafe) Health() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Health"),
	}
}

// Redis 安全访问Redis配置
func (c *ConfigSafe) Redis() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Redis"),
	}
}

// MySQL 安全访问MySQL配置
func (c *ConfigSafe) MySQL() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("MySQL"),
	}
}

// HTTP 安全访问HTTP配置
func (c *ConfigSafe) HTTP() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("HTTP"),
	}
}

// HTTPServer 安全访问HTTPServer配置
func (c *ConfigSafe) HTTPServer() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("HTTPServer"),
	}
}

// Server 安全访问Server配置
func (c *ConfigSafe) Server() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Server"),
	}
}

// Middleware 安全访问Middleware配置
func (c *ConfigSafe) Middleware() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Middleware"),
	}
}

// CORS 安全访问CORS配置
func (c *ConfigSafe) CORS() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("CORS"),
	}
}

// RateLimit 安全访问RateLimit配置
func (c *ConfigSafe) RateLimit() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("RateLimit"),
	}
}

// Enabled 获取Enabled字段值
func (c *ConfigSafe) Enabled(defaultValue ...bool) bool {
	return c.Field("Enabled").Bool(defaultValue...)
}

// Port 获取Port字段值
func (c *ConfigSafe) Port(defaultValue ...int) int {
	return c.Field("Port").Int(defaultValue...)
}

// Host 获取Host字段值
func (c *ConfigSafe) Host(defaultValue ...string) string {
	return c.Field("Host").String(defaultValue...)
}

// Timeout 获取Timeout字段值并转换为Duration
func (c *ConfigSafe) Timeout(defaultValue ...time.Duration) time.Duration {
	return c.Field("Timeout").Duration(defaultValue...)
}

// Name 获取Name字段值
func (c *ConfigSafe) Name(defaultValue ...string) string {
	return c.Field("Name").String(defaultValue...)
}

// Version 获取Version字段值
func (c *ConfigSafe) Version(defaultValue ...string) string {
	return c.Field("Version").String(defaultValue...)
}

// Debug 获取Debug字段值
func (c *ConfigSafe) Debug(defaultValue ...bool) bool {
	return c.Field("Debug").Bool(defaultValue...)
}

// 链式调用示例方法

// IsRedisHealthEnabled 检查Redis健康检查是否启用
func (c *ConfigSafe) IsRedisHealthEnabled() bool {
	return c.Health().Redis().Enabled()
}

// GetRedisHealthTimeout 获取Redis健康检查超时时间
func (c *ConfigSafe) GetRedisHealthTimeout(defaultTimeout time.Duration) time.Duration {
	return c.Health().Redis().Timeout(defaultTimeout)
}

// IsMySQLHealthEnabled 检查MySQL健康检查是否启用
func (c *ConfigSafe) IsMySQLHealthEnabled() bool {
	return c.Health().MySQL().Enabled()
}

// GetMySQLHealthTimeout 获取MySQL健康检查超时时间
func (c *ConfigSafe) GetMySQLHealthTimeout(defaultTimeout time.Duration) time.Duration {
	return c.Health().MySQL().Timeout(defaultTimeout)
}

// IsCORSEnabled 检查CORS是否启用
func (c *ConfigSafe) IsCORSEnabled() bool {
	return c.Middleware().CORS().Enabled()
}

// IsRateLimitEnabled 检查限流是否启用
func (c *ConfigSafe) IsRateLimitEnabled() bool {
	return c.Middleware().RateLimit().Enabled()
}

// GetHTTPPort 获取HTTP端口
func (c *ConfigSafe) GetHTTPPort(defaultPort int) int {
	return c.HTTP().Port(defaultPort)
}

// GetServerPort 获取服务器端口
func (c *ConfigSafe) GetServerPort(defaultPort int) int {
	return c.Server().Port(defaultPort)
}

// 健康检查相关方法

// IsHealthEnabled 检查健康检查是否启用
func (c *ConfigSafe) IsHealthEnabled() bool {
	return c.Health().Enabled()
}

// GetHealthPath 获取健康检查路径
func (c *ConfigSafe) GetHealthPath(defaultPath string) string {
	return c.Health().Field("Path").String(defaultPath)
}

// GetRedisHealthPath 获取Redis健康检查路径
func (c *ConfigSafe) GetRedisHealthPath(defaultPath string) string {
	return c.Health().Redis().Field("Path").String(defaultPath)
}

// GetMySQLHealthPath 获取MySQL健康检查路径
func (c *ConfigSafe) GetMySQLHealthPath(defaultPath string) string {
	return c.Health().MySQL().Field("Path").String(defaultPath)
}

// ======= PProf性能分析相关方法 =======

// IsPProfEnabled 检查PProf性能分析是否启用
func (c *ConfigSafe) IsPProfEnabled() bool {
	return c.Middleware().Field("PProf").Field("Enabled").Bool(false)
}

// GetPProfPathPrefix 获取PProf路径前缀
func (c *ConfigSafe) GetPProfPathPrefix(defaultPrefix string) string {
	return c.Middleware().Field("PProf").Field("PathPrefix").String(defaultPrefix)
}

// ======= Monitoring监控相关方法 =======

// IsMonitoringEnabled 检查监控是否启用
func (c *ConfigSafe) IsMonitoringEnabled() bool {
	return c.Field("Monitoring").Field("Enabled").Bool(false)
}

// IsMetricsEnabled 检查Metrics是否启用
func (c *ConfigSafe) IsMetricsEnabled() bool {
	return c.Field("Monitoring").Field("Metrics").Field("Enabled").Bool(false)
}

// GetMetricsEndpoint 获取Metrics端点路径
func (c *ConfigSafe) GetMetricsEndpoint(defaultEndpoint string) string {
	return c.Field("Monitoring").Field("Metrics").Field("Endpoint").String(defaultEndpoint)
}

// ======= Jaeger链路追踪相关方法 =======

// IsJaegerEnabled 检查Jaeger链路追踪是否启用
func (c *ConfigSafe) IsJaegerEnabled() bool {
	return c.Field("Monitoring").Field("Jaeger").Field("Enabled").Bool(false)
}

// GetJaegerServiceName 获取Jaeger服务名称
func (c *ConfigSafe) GetJaegerServiceName(defaultServiceName string) string {
	return c.Field("Monitoring").Field("Jaeger").Field("ServiceName").String(defaultServiceName)
}

// GetJaegerEndpoint 获取Jaeger端点
func (c *ConfigSafe) GetJaegerEndpoint(defaultEndpoint string) string {
	return c.Field("Monitoring").Field("Jaeger").Field("Endpoint").String(defaultEndpoint)
}

// GetJaegerSamplingType 获取Jaeger采样类型
func (c *ConfigSafe) GetJaegerSamplingType(defaultType string) string {
	return c.Field("Monitoring").Field("Jaeger").Field("Sampling").Field("Type").String(defaultType)
}
