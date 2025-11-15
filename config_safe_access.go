/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-13 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-15 14:47:38
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

// SafeConfig 专门针对配置的安全访问工具
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

// GRPC 安全访问GRPC配置
func (c *ConfigSafe) GRPC() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("GRPC"),
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

// Cache 安全访问Cache配置
func (c *ConfigSafe) Cache() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Cache"),
	}
}

// WSC 安全访问WSC配置
func (c *ConfigSafe) WSC() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("WSC"),
	}
}

// Swagger 安全访问Swagger配置
func (c *ConfigSafe) Swagger() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Swagger"),
	}
}

// Monitoring 安全访问Monitoring配置
func (c *ConfigSafe) Monitoring() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Monitoring"),
	}
}

// Metrics 安全访问Metrics配置
func (c *ConfigSafe) Metrics() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Metrics"),
	}
}

// Prometheus 安全访问Prometheus配置
func (c *ConfigSafe) Prometheus() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Prometheus"),
	}
}

// Memory 安全访问Memory缓存配置
func (c *ConfigSafe) Memory() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Memory"),
	}
}

// Ristretto 安全访问Ristretto缓存配置
func (c *ConfigSafe) Ristretto() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Ristretto"),
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

// ======= JWT认证相关方法 =======

// JWT 安全访问JWT配置
func (c *ConfigSafe) JWT() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("JWT"),
	}
}

// IsJWTEnabled 检查JWT是否启用
func (c *ConfigSafe) IsJWTEnabled() bool {
	return c.JWT().Enabled(false)
}

// GetJWTSecret 获取JWT密钥
func (c *ConfigSafe) GetJWTSecret(defaultSecret string) string {
	return c.JWT().Field("Secret").String(defaultSecret)
}

// GetJWTExpiration 获取JWT过期时间
func (c *ConfigSafe) GetJWTExpiration(defaultExpiration time.Duration) time.Duration {
	return c.JWT().Field("Expiration").Duration(defaultExpiration)
}

// ======= 队列相关方法 =======

// Queue 安全访问Queue配置
func (c *ConfigSafe) Queue() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Queue"),
	}
}

// MQTT 安全访问MQTT配置
func (c *ConfigSafe) MQTT() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("MQTT"),
	}
}

// IsMQTTEnabled 检查MQTT是否启用
func (c *ConfigSafe) IsMQTTEnabled() bool {
	return c.MQTT().Enabled(false)
}

// GetMQTTBroker 获取MQTT代理地址
func (c *ConfigSafe) GetMQTTBroker(defaultBroker string) string {
	return c.MQTT().Field("Broker").String(defaultBroker)
}

// ======= 日志相关方法 =======

// Logging 安全访问Logging配置
func (c *ConfigSafe) Logging() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Logging"),
	}
}

// Zap 安全访问Zap日志配置
func (c *ConfigSafe) Zap() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Zap"),
	}
}

// IsLoggingEnabled 检查日志是否启用
func (c *ConfigSafe) IsLoggingEnabled() bool {
	return c.Logging().Enabled(true)
}

// GetLogLevel 获取日志级别
func (c *ConfigSafe) GetLogLevel(defaultLevel string) string {
	return c.Logging().Field("Level").String(defaultLevel)
}

// GetLogPath 获取日志文件路径
func (c *ConfigSafe) GetLogPath(defaultPath string) string {
	return c.Logging().Field("Path").String(defaultPath)
}

// ======= 对象存储相关方法 =======

// OSS 安全访问OSS配置
func (c *ConfigSafe) OSS() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("OSS"),
	}
}

// Aliyun 安全访问阿里云配置
func (c *ConfigSafe) Aliyun() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Aliyun"),
	}
}

// Minio 安全访问Minio配置
func (c *ConfigSafe) Minio() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Minio"),
	}
}

// S3 安全访问S3配置
func (c *ConfigSafe) S3() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("S3"),
	}
}

// IsOSSEnabled 检查对象存储是否启用
func (c *ConfigSafe) IsOSSEnabled() bool {
	return c.OSS().Enabled(false)
}

// GetOSSEndpoint 获取OSS端点
func (c *ConfigSafe) GetOSSEndpoint(defaultEndpoint string) string {
	return c.OSS().Field("Endpoint").String(defaultEndpoint)
}

// GetOSSBucket 获取OSS存储桶
func (c *ConfigSafe) GetOSSBucket(defaultBucket string) string {
	return c.OSS().Field("Bucket").String(defaultBucket)
}

// ======= 邮件相关方法 =======

// Email 安全访问Email配置
func (c *ConfigSafe) Email() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Email"),
	}
}

// SMTP 安全访问SMTP配置
func (c *ConfigSafe) SMTP() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("SMTP"),
	}
}

// IsEmailEnabled 检查邮件是否启用
func (c *ConfigSafe) IsEmailEnabled() bool {
	return c.Email().Enabled(false)
}

// GetSMTPHost 获取SMTP主机
func (c *ConfigSafe) GetSMTPHost(defaultHost string) string {
	return c.Email().SMTP().Host(defaultHost)
}

// GetSMTPPort 获取SMTP端口
func (c *ConfigSafe) GetSMTPPort(defaultPort int) int {
	return c.Email().SMTP().Port(defaultPort)
}

// ======= 支付相关方法 =======

// Pay 安全访问Pay配置
func (c *ConfigSafe) Pay() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Pay"),
	}
}

// Alipay 安全访问支付宝配置
func (c *ConfigSafe) Alipay() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Alipay"),
	}
}

// Wechat 安全访问微信支付配置
func (c *ConfigSafe) Wechat() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Wechat"),
	}
}

// IsPayEnabled 检查支付是否启用
func (c *ConfigSafe) IsPayEnabled() bool {
	return c.Pay().Enabled(false)
}

// IsAlipayEnabled 检查支付宝是否启用
func (c *ConfigSafe) IsAlipayEnabled() bool {
	return c.Pay().Alipay().Enabled(false)
}

// IsWechatPayEnabled 检查微信支付是否启用
func (c *ConfigSafe) IsWechatPayEnabled() bool {
	return c.Pay().Wechat().Enabled(false)
}

// ======= 短信相关方法 =======

// SMS 安全访问SMS配置
func (c *ConfigSafe) SMS() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("SMS"),
	}
}

// IsSMSEnabled 检查短信是否启用
func (c *ConfigSafe) IsSMSEnabled() bool {
	return c.SMS().Enabled(false)
}

// GetSMSAccessKeyID 获取短信AccessKeyID
func (c *ConfigSafe) GetSMSAccessKeyID(defaultKeyID string) string {
	return c.SMS().Field("AccessKeyID").String(defaultKeyID)
}

// ======= 验证码相关方法 =======

// Captcha 安全访问Captcha配置
func (c *ConfigSafe) Captcha() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Captcha"),
	}
}

// IsCaptchaEnabled 检查验证码是否启用
func (c *ConfigSafe) IsCaptchaEnabled() bool {
	return c.Captcha().Enabled(false)
}

// GetCaptchaType 获取验证码类型
func (c *ConfigSafe) GetCaptchaType(defaultType string) string {
	return c.Captcha().Field("Type").String(defaultType)
}

// ======= 国际化相关方法 =======

// I18n 安全访问国际化配置
func (c *ConfigSafe) I18n() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("I18n"),
	}
}

// IsI18nEnabled 检查国际化是否启用
func (c *ConfigSafe) IsI18nEnabled() bool {
	return c.I18n().Enabled(false)
}

// GetI18nDefaultLang 获取默认语言
func (c *ConfigSafe) GetI18nDefaultLang(defaultLang string) string {
	return c.I18n().Field("DefaultLang").String(defaultLang)
}

// ======= 服务发现相关方法 =======

// Consul 安全访问Consul配置
func (c *ConfigSafe) Consul() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Consul"),
	}
}

// Etcd 安全访问Etcd配置
func (c *ConfigSafe) Etcd() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Etcd"),
	}
}

// IsConsulEnabled 检查Consul是否启用
func (c *ConfigSafe) IsConsulEnabled() bool {
	return c.Consul().Enabled(false)
}

// IsEtcdEnabled 检查Etcd是否启用
func (c *ConfigSafe) IsEtcdEnabled() bool {
	return c.Etcd().Enabled(false)
}

// GetConsulAddress 获取Consul地址
func (c *ConfigSafe) GetConsulAddress(defaultAddress string) string {
	return c.Consul().Field("Address").String(defaultAddress)
}

// GetEtcdEndpoints 获取Etcd端点列表
func (c *ConfigSafe) GetEtcdEndpoints(defaultEndpoints []string) []string {
	endpoints := c.Etcd().Field("Endpoints").Value()
	if endpoints == nil {
		return defaultEndpoints
	}
	if slice, ok := endpoints.([]string); ok {
		return slice
	}
	if slice, ok := endpoints.([]interface{}); ok {
		result := make([]string, len(slice))
		for i, v := range slice {
			if str, ok := v.(string); ok {
				result[i] = str
			}
		}
		return result
	}
	return defaultEndpoints
}

// ======= 安全相关的nil检查方法 =======

// IsNil 检查当前配置是否为nil
func (c *ConfigSafe) IsNil() bool {
	return c.SafeAccess.Value() == nil
}

// IsValidConfig 检查配置是否有效（非nil且有值）
func (c *ConfigSafe) IsValidConfig() bool {
	return c.SafeAccess.IsValid()
}

// SafeField 安全获取字段，如果字段不存在或为nil则返回带有默认值的SafeAccess
func (c *ConfigSafe) SafeField(fieldName string) *ConfigSafe {
	if c.IsNil() {
		return &ConfigSafe{
			SafeAccess: safe.Safe(map[string]interface{}{}),
		}
	}

	// 检查字段是否存在
	fieldValue := c.Field(fieldName)
	if fieldValue.Value() == nil {
		// 字段不存在或为nil，返回空map以确保IsValidConfig返回true
		return &ConfigSafe{
			SafeAccess: safe.Safe(map[string]interface{}{}),
		}
	}

	return &ConfigSafe{
		SafeAccess: fieldValue,
	}
}

// WithDefault 为当前配置设置默认值（如果配置为nil或不存在）
func (c *ConfigSafe) WithDefault(defaultValue interface{}) *ConfigSafe {
	if !c.SafeAccess.IsValid() {
		return &ConfigSafe{
			SafeAccess: safe.Safe(defaultValue),
		}
	}
	return c
} // ======= 通用配置检查方法 =======

// HasField 检查是否包含指定字段
func (c *ConfigSafe) HasField(fieldName string) bool {
	if c.IsNil() {
		return false
	}
	return c.Field(fieldName).IsValid()
} // GetString 安全获取字符串值，支持多级路径
func (c *ConfigSafe) GetString(path string, defaultValue ...string) string {
	if c.IsNil() {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return c.Field(path).String(defaultValue...)
}

// GetInt 安全获取整数值
func (c *ConfigSafe) GetInt(path string, defaultValue ...int) int {
	if c.IsNil() {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return c.Field(path).Int(defaultValue...)
}

// SafeGetBool 安全获取布尔值
func (c *ConfigSafe) SafeGetBool(path string, defaultValue ...bool) bool {
	if c.IsNil() {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}
	return c.Field(path).Bool(defaultValue...)
}

// GetDuration 安全获取时间间隔值
func (c *ConfigSafe) GetDuration(path string, defaultValue ...time.Duration) time.Duration {
	if c.IsNil() {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return c.Field(path).Duration(defaultValue...)
}
