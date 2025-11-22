/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-13 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-21 23:49:01
 * @FilePath: \go-config\config_safe_access.go
 * @Description: 配置安全访问辅助工具，专门针对配置结构体优化
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"github.com/kamalyes/go-toolbox/pkg/safe"
	"time"
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
	// 检测当前配置是否已经是Health类型
	if val := c.SafeAccess.Value(); val != nil {
		// 通过ModuleName字段判断
		if moduleName := c.Field("ModuleName").String(""); moduleName == "health" {
			return c
		}
	}
	return &ConfigSafe{
		SafeAccess: c.SafeAccess.Field("Health"),
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
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "wsc" {
			return c
		}
	}
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
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "monitoring" {
			return c
		}
	}
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
	// Health.Redis.Timeout 是 int 类型,单位为秒
	timeout := c.Health().Redis().Field("Timeout").Int(0)
	if timeout > 0 {
		return time.Duration(timeout) * time.Second
	}
	// 尝试Duration类型
	return c.Health().Redis().Timeout(defaultTimeout)
}

// IsMySQLHealthEnabled 检查MySQL健康检查是否启用
func (c *ConfigSafe) IsMySQLHealthEnabled() bool {
	return c.Health().MySQL().Enabled()
}

// GetMySQLHealthTimeout 获取MySQL健康检查超时时间
func (c *ConfigSafe) GetMySQLHealthTimeout(defaultTimeout time.Duration) time.Duration {
	// Health.MySQL.Timeout 是 int 类型,单位为秒
	timeout := c.Health().MySQL().Field("Timeout").Int(0)
	if timeout > 0 {
		return time.Duration(timeout) * time.Second
	}
	// 尝试Duration类型
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
	return c.Monitoring().Field("Metrics").Field("Enabled").Bool(false)
}

// GetMetricsEndpoint 获取Metrics端点路径
func (c *ConfigSafe) GetMetricsEndpoint(defaultEndpoint string) string {
	return c.Monitoring().Field("Metrics").Field("Endpoint").String(defaultEndpoint)
}

// ======= Jaeger链路追踪相关方法 =======

// IsJaegerEnabled 检查Jaeger链路追踪是否启用
func (c *ConfigSafe) IsJaegerEnabled() bool {
	return c.Monitoring().Field("Jaeger").Field("Enabled").Bool(false)
}

// GetJaegerServiceName 获取Jaeger服务名称
func (c *ConfigSafe) GetJaegerServiceName(defaultServiceName string) string {
	return c.Monitoring().Field("Jaeger").Field("ServiceName").String(defaultServiceName)
}

// GetJaegerEndpoint 获取Jaeger端点
func (c *ConfigSafe) GetJaegerEndpoint(defaultEndpoint string) string {
	return c.Monitoring().Field("Jaeger").Field("Endpoint").String(defaultEndpoint)
}

// GetJaegerSamplingType 获取Jaeger采样类型
func (c *ConfigSafe) GetJaegerSamplingType(defaultType string) string {
	return c.Monitoring().Field("Jaeger").Field("Sampling").Field("Type").String(defaultType)
}

// ======= JWT认证相关方法 =======

// JWT 安全访问JWT配置
func (c *ConfigSafe) JWT() *ConfigSafe {
	// 检测当前配置是否已经是JWT类型
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "jwt" {
			return c
		}
	}
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
	// 直接从根配置获取SigningKey字段
	secret := c.Field("SigningKey").String("")
	if secret != "" {
		return secret
	}
	// 尝试从JWT子配置获取
	secret = c.JWT().Field("SigningKey").String("")
	if secret == "" {
		secret = c.JWT().Field("Secret").String(defaultSecret)
	}
	return secret
}

// GetJWTExpiration 获取JWT过期时间
func (c *ConfigSafe) GetJWTExpiration(defaultExpiration time.Duration) time.Duration {
	// 直接从根配置获取ExpiresTime字段(int64秒) - SafeAccess的Int64方法会自动处理类型转换
	expiresTime := c.Field("ExpiresTime").Int64(0)
	if expiresTime > 0 {
		return time.Duration(expiresTime) * time.Second
	}

	// 尝试从JWT子配置获取
	expiresTime = c.JWT().Field("ExpiresTime").Int64(0)
	if expiresTime > 0 {
		return time.Duration(expiresTime) * time.Second
	}
	// 尝试Duration类型
	return c.JWT().Field("Expiration").Duration(defaultExpiration)
} // ======= 队列相关方法 =======

// Queue 安全访问Queue配置
func (c *ConfigSafe) Queue() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Queue"),
	}
}

// MQTT 安全访问MQTT配置
func (c *ConfigSafe) MQTT() *ConfigSafe {
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "mqtt" {
			return c
		}
	}
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
	// Queue/MQTT使用Endpoint字段
	broker := c.MQTT().Field("Endpoint").String("")
	if broker == "" {
		broker = c.MQTT().Field("Broker").String(defaultBroker)
	}
	return broker
}

// ======= 日志相关方法 =======

// Logging 安全访问Logging配置
func (c *ConfigSafe) Logging() *ConfigSafe {
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "logging" {
			return c
		}
	}
	return &ConfigSafe{
		SafeAccess: c.Field("Logging"),
	}
}

// Zap 安全访问Zap日志配置
func (c *ConfigSafe) Zap() *ConfigSafe {
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "zap" {
			return c
		}
	}
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
	// Logging结构体使用FilePath字段
	path := c.Logging().Field("FilePath").String("")
	if path == "" {
		path = c.Logging().Field("Path").String(defaultPath)
	}
	return path
}

// ======= 对象存储相关方法 =======

// OSS 安全访问OSS配置
func (c *ConfigSafe) OSS() *ConfigSafe {
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "oss" {
			return c
		}
	}
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
	// 先尝试从根OSS配置获取
	endpoint := c.OSS().Field("Endpoint").String("")
	if endpoint != "" {
		return endpoint
	}
	// 尝试从AliyunOSS获取
	endpoint = c.OSS().Field("AliyunOSS").Field("Endpoint").String("")
	if endpoint != "" {
		return endpoint
	}
	// 尝试从Aliyun获取
	endpoint = c.OSS().Aliyun().Field("Endpoint").String(defaultEndpoint)
	return endpoint
}

// GetOSSBucket 获取OSS存储桶
func (c *ConfigSafe) GetOSSBucket(defaultBucket string) string {
	// 先尝试从根OSS配置获取
	bucket := c.OSS().Field("Bucket").String("")
	if bucket != "" {
		return bucket
	}
	// 尝试从AliyunOSS获取
	bucket = c.OSS().Field("AliyunOSS").Field("Bucket").String("")
	if bucket != "" {
		return bucket
	}
	// 尝试从Aliyun获取
	bucket = c.OSS().Aliyun().Field("Bucket").String(defaultBucket)
	return bucket
}

// ======= 邮件相关方法 =======

// Email 安全访问Email配置
func (c *ConfigSafe) Email() *ConfigSafe {
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "email" {
			return c
		}
	}
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
	// Email结构体本身就有Host字段,不是SMTP子配置
	host := c.Email().Host(defaultHost)
	if host == defaultHost || host == "" {
		// 尝试SMTP子配置
		host = c.Email().SMTP().Host(defaultHost)
	}
	return host
}

// GetSMTPPort 获取SMTP端口
func (c *ConfigSafe) GetSMTPPort(defaultPort int) int {
	// Email结构体本身就有Port字段,不是SMTP子配置
	port := c.Email().Port(0)
	if port == 0 {
		// 尝试SMTP子配置
		port = c.Email().SMTP().Port(defaultPort)
	}
	return port
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
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "i18n" {
			return c
		}
	}
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
	// I18n结构体使用DefaultLanguage字段
	lang := c.I18n().Field("DefaultLanguage").String("")
	if lang == "" {
		lang = c.I18n().Field("DefaultLang").String(defaultLang)
	}
	return lang
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
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "etcd" {
			return c
		}
	}
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
		// 尝试从Hosts字段获取(兼容不同的配置结构)
		endpoints = c.Etcd().Field("Hosts").Value()
	}
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
}

// ======= 通用配置检查方法 =======

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

// ======= 更多配置访问器方法 =======

// Database 安全访问Database配置
func (c *ConfigSafe) Database() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Database"),
	}
}

// Banner 安全访问Banner配置
func (c *ConfigSafe) Banner() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Banner"),
	}
}

// PostgreSQL 安全访问PostgreSQL配置
func (c *ConfigSafe) PostgreSQL() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("PostgreSQL"),
	}
}

// SQLite 安全访问SQLite配置
func (c *ConfigSafe) SQLite() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("SQLite"),
	}
}

// Signature 安全访问Signature配置
func (c *ConfigSafe) Signature() *ConfigSafe {
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "signature" {
			return c
		}
	}
	return &ConfigSafe{
		SafeAccess: c.Field("Signature"),
	}
}

// Tracing 安全访问Tracing配置
func (c *ConfigSafe) Tracing() *ConfigSafe {
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "tracing" {
			return c
		}
	}
	return &ConfigSafe{
		SafeAccess: c.Field("Tracing"),
	}
}

// TimeoutConfig 安全访问Timeout配置
func (c *ConfigSafe) TimeoutConfig() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Timeout"),
	}
}

// Kafka 安全访问Kafka配置
func (c *ConfigSafe) Kafka() *ConfigSafe {
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "kafka" {
			return c
		}
	}
	return &ConfigSafe{
		SafeAccess: c.Field("Kafka"),
	}
}

// Access 安全访问Access配置
func (c *ConfigSafe) Access() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Access"),
	}
}

// Alerting 安全访问Alerting配置
func (c *ConfigSafe) Alerting() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Alerting"),
	}
}

// Breaker 安全访问Breaker配置
func (c *ConfigSafe) Breaker() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Breaker"),
	}
}

// CircuitBreaker 安全访问CircuitBreaker配置(Breaker的别名)
func (c *ConfigSafe) CircuitBreaker() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("CircuitBreaker"),
	}
}

// WebSocketBreaker 安全访问WebSocketBreaker配置
func (c *ConfigSafe) WebSocketBreaker() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("WebSocketBreaker"),
	}
}

// Elasticsearch 安全访问Elasticsearch配置
func (c *ConfigSafe) Elasticsearch() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Elasticsearch"),
	}
}

// FTP 安全访问FTP配置
func (c *ConfigSafe) FTP() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("FTP"),
	}
}

// Ftp 安全访问Ftp配置(FTP的别名)
func (c *ConfigSafe) Ftp() *ConfigSafe {
	return c.FTP()
}

// Gateway 安全访问Gateway配置
func (c *ConfigSafe) Gateway() *ConfigSafe {
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "gateway" {
			return c
		}
	}
	return &ConfigSafe{
		SafeAccess: c.Field("Gateway"),
	}
}

// JSON 安全访问JSON配置
func (c *ConfigSafe) JSON() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("JSON"),
	}
}

// TLS 安全访问TLS配置
func (c *ConfigSafe) TLS() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("TLS"),
	}
}

// Grafana 安全访问Grafana配置
func (c *ConfigSafe) Grafana() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Grafana"),
	}
}

// Request 安全访问Request配置
func (c *ConfigSafe) Request() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Request"),
	}
}

// RequestID 安全访问RequestID配置
func (c *ConfigSafe) RequestID() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("RequestID"),
	}
}

// Recovery 安全访问Recovery配置
func (c *ConfigSafe) Recovery() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Recovery"),
	}
}

// Restful 安全访问Restful配置
func (c *ConfigSafe) Restful() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Restful"),
	}
}

// RESTful 安全访问RESTful配置(Restful的别名)
func (c *ConfigSafe) RESTful() *ConfigSafe {
	return c.Restful()
}

// RPCClient 安全访问RPCClient配置
func (c *ConfigSafe) RPCClient() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("RPCClient"),
	}
}

// RpcClient 安全访问RpcClient配置(RPCClient的别名)
func (c *ConfigSafe) RpcClient() *ConfigSafe {
	return c.RPCClient()
}

// RPCServer 安全访问RPCServer配置
func (c *ConfigSafe) RPCServer() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("RPCServer"),
	}
}

// RpcServer 安全访问RpcServer配置(RPCServer的别名)
func (c *ConfigSafe) RpcServer() *ConfigSafe {
	return c.RPCServer()
}

// Security 安全访问Security配置
func (c *ConfigSafe) Security() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Security"),
	}
}

// Auth 安全访问Auth配置(Security的别名)
func (c *ConfigSafe) Auth() *ConfigSafe {
	return c.Security()
}

// Protection 安全访问Protection配置(Security的别名)
func (c *ConfigSafe) Protection() *ConfigSafe {
	return c.Security()
}

// STS 安全访问STS配置
func (c *ConfigSafe) STS() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("STS"),
	}
}

// Youzan 安全访问Youzan配置
func (c *ConfigSafe) Youzan() *ConfigSafe {
	if val := c.SafeAccess.Value(); val != nil {
		if moduleName := c.Field("ModuleName").String(""); moduleName == "youzan" {
			return c
		}
	}
	return &ConfigSafe{
		SafeAccess: c.Field("Youzan"),
	}
}

// Distributed 安全访问Distributed配置
func (c *ConfigSafe) Distributed() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Distributed"),
	}
}

// Group 安全访问Group配置
func (c *ConfigSafe) Group() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Group"),
	}
}

// Ticket 安全访问Ticket配置
func (c *ConfigSafe) Ticket() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Ticket"),
	}
}

// Performance 安全访问Performance配置
func (c *ConfigSafe) Performance() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Performance"),
	}
}

// VIP 安全访问VIP配置
func (c *ConfigSafe) VIP() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("VIP"),
	}
}

// Enhancement 安全访问Enhancement配置
func (c *ConfigSafe) Enhancement() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("Enhancement"),
	}
}

// BoltDB 安全访问BoltDB配置
func (c *ConfigSafe) BoltDB() *ConfigSafe {
	return &ConfigSafe{
		SafeAccess: c.Field("BoltDB"),
	}
}

// ======= 具体配置项的Getter方法 =======

// GetKafkaTopic 获取Kafka主题
func (c *ConfigSafe) GetKafkaTopic(defaultTopic string) string {
	return c.Kafka().Field("Topic").String(defaultTopic)
}

// GetKafkaBrokers 获取Kafka代理列表
func (c *ConfigSafe) GetKafkaBrokers(defaultBrokers []string) []string {
	brokers := c.Kafka().Field("Brokers").Value()
	if brokers == nil {
		return defaultBrokers
	}
	if slice, ok := brokers.([]string); ok {
		return slice
	}
	if slice, ok := brokers.([]interface{}); ok {
		result := make([]string, len(slice))
		for i, v := range slice {
			if str, ok := v.(string); ok {
				result[i] = str
			}
		}
		return result
	}
	return defaultBrokers
}

// GetSignatureSecretKey 获取签名密钥
func (c *ConfigSafe) GetSignatureSecretKey(defaultKey string) string {
	return c.Signature().Field("SecretKey").String(defaultKey)
}

// GetSignatureAlgorithm 获取签名算法
func (c *ConfigSafe) GetSignatureAlgorithm(defaultAlgorithm string) string {
	return c.Signature().Field("Algorithm").String(defaultAlgorithm)
}

// GetTimeoutDuration 获取超时时间
func (c *ConfigSafe) GetTimeoutDuration(defaultDuration time.Duration) time.Duration {
	return c.TimeoutConfig().Field("Duration").Duration(defaultDuration)
}

// GetTracingServiceName 获取链路追踪服务名称
func (c *ConfigSafe) GetTracingServiceName(defaultName string) string {
	return c.Tracing().Field("ServiceName").String(defaultName)
}

// GetTracingEndpoint 获取链路追踪端点
func (c *ConfigSafe) GetTracingEndpoint(defaultEndpoint string) string {
	return c.Tracing().Field("Endpoint").String(defaultEndpoint)
}

// GetAccessServiceName 获取访问日志服务名称
func (c *ConfigSafe) GetAccessServiceName(defaultName string) string {
	return c.Access().Field("ServiceName").String(defaultName)
}

// GetAccessRetentionDays 获取访问日志保留天数
func (c *ConfigSafe) GetAccessRetentionDays(defaultDays int) int {
	return c.Access().Field("RetentionDays").Int(defaultDays)
}

// GetBannerTitle 获取Banner标题
func (c *ConfigSafe) GetBannerTitle(defaultTitle string) string {
	return c.Banner().Field("Title").String(defaultTitle)
}

// GetBreakerFailureThreshold 获取熔断器失败阈值
func (c *ConfigSafe) GetBreakerFailureThreshold(defaultThreshold int) int {
	breaker := c.CircuitBreaker()
	if !breaker.IsValid() {
		breaker = c.Breaker()
	}
	return breaker.Field("FailureThreshold").Int(defaultThreshold)
}

// GetPostgreSQLHost 获取PostgreSQL主机
func (c *ConfigSafe) GetPostgreSQLHost(defaultHost string) string {
	return c.Database().PostgreSQL().Host(defaultHost)
}

// GetSQLitePath 获取SQLite数据库路径
func (c *ConfigSafe) GetSQLitePath(defaultPath string) string {
	return c.Database().SQLite().Field("Path").String(defaultPath)
}

// GetElasticsearchEndpoint 获取Elasticsearch端点
func (c *ConfigSafe) GetElasticsearchEndpoint(defaultEndpoint string) string {
	return c.Elasticsearch().Field("Endpoint").String(defaultEndpoint)
}

// GetFTPEndpoint 获取FTP端点
func (c *ConfigSafe) GetFTPEndpoint(defaultEndpoint string) string {
	return c.FTP().Field("Endpoint").String(defaultEndpoint)
}

// GetGrafanaEndpoint 获取Grafana端点
func (c *ConfigSafe) GetGrafanaEndpoint(defaultEndpoint string) string {
	return c.Grafana().Field("Endpoint").String(defaultEndpoint)
}

// GetRequestTraceID 获取请求跟踪ID
func (c *ConfigSafe) GetRequestTraceID(defaultID string) string {
	return c.Request().Field("TraceID").String(defaultID)
}

// GetRequestIDHeaderName 获取请求ID头名称
func (c *ConfigSafe) GetRequestIDHeaderName(defaultName string) string {
	return c.RequestID().Field("HeaderName").String(defaultName)
}

// GetSTSRegionID 获取STS区域ID
func (c *ConfigSafe) GetSTSRegionID(defaultRegion string) string {
	return c.STS().Field("RegionID").String(defaultRegion)
}

// GetSTSAccessKeyID 获取STS访问密钥ID
func (c *ConfigSafe) GetSTSAccessKeyID(defaultKeyID string) string {
	return c.STS().Field("AccessKeyID").String(defaultKeyID)
}

// GetYouzanEndpoint 获取有赞端点
func (c *ConfigSafe) GetYouzanEndpoint(defaultEndpoint string) string {
	return c.Youzan().Field("Endpoint").String(defaultEndpoint)
}

// GetYouzanClientID 获取有赞客户端ID
func (c *ConfigSafe) GetYouzanClientID(defaultClientID string) string {
	return c.Youzan().Field("ClientID").String(defaultClientID)
}

// GetWSCNodeIP 获取WSC节点IP
func (c *ConfigSafe) GetWSCNodeIP(defaultIP string) string {
	return c.WSC().Field("NodeIP").String(defaultIP)
}

// GetWSCNodePort 获取WSC节点端口
func (c *ConfigSafe) GetWSCNodePort(defaultPort int) int {
	return c.WSC().Field("NodePort").Int(defaultPort)
}

// GetWSCHeartbeatInterval 获取WSC心跳间隔
func (c *ConfigSafe) GetWSCHeartbeatInterval(defaultInterval time.Duration) time.Duration {
	// WSC.HeartbeatInterval是int类型,单位为秒
	interval := c.WSC().Field("HeartbeatInterval").Int(0)
	if interval > 0 {
		return time.Duration(interval) * time.Second
	}
	// 尝试Duration类型
	return c.WSC().Field("HeartbeatInterval").Duration(defaultInterval)
}

// GetZapLevel 获取Zap日志级别
func (c *ConfigSafe) GetZapLevel(defaultLevel string) string {
	return c.Zap().Field("Level").String(defaultLevel)
}

// GetZapFormat 获取Zap日志格式
func (c *ConfigSafe) GetZapFormat(defaultFormat string) string {
	return c.Zap().Field("Format").String(defaultFormat)
}

// GetZapDirector 获取Zap日志目录
func (c *ConfigSafe) GetZapDirector(defaultDir string) string {
	return c.Zap().Field("Director").String(defaultDir)
}

// GetBoltDBPath 获取BoltDB路径
func (c *ConfigSafe) GetBoltDBPath(defaultPath string) string {
	return c.BoltDB().Field("Path").String(defaultPath)
}

// GetBoltDBBucket 获取BoltDB存储桶
func (c *ConfigSafe) GetBoltDBBucket(defaultBucket string) string {
	return c.BoltDB().Field("Bucket").String(defaultBucket)
}

// ======= IsXxxEnabled便捷方法(使用通用Enabled方法) =======

// IsAccessEnabled 检查访问日志是否启用
func (c *ConfigSafe) IsAccessEnabled() bool {
	return c.Access().Enabled(false)
}

// IsAlertingEnabled 检查告警是否启用
func (c *ConfigSafe) IsAlertingEnabled() bool {
	return c.Alerting().Enabled(false)
}

// IsBannerEnabled 检查Banner是否启用
func (c *ConfigSafe) IsBannerEnabled() bool {
	return c.Banner().Enabled(false)
}

// IsBreakerEnabled 检查熔断器是否启用
func (c *ConfigSafe) IsBreakerEnabled() bool {
	breaker := c.CircuitBreaker()
	if !breaker.IsValid() {
		breaker = c.Breaker()
	}
	return breaker.Enabled(false)
}

// IsWebSocketBreakerEnabled 检查WebSocket熔断器是否启用
func (c *ConfigSafe) IsWebSocketBreakerEnabled() bool {
	return c.WebSocketBreaker().Enabled(false)
}

// IsDatabaseEnabled 检查数据库是否启用
func (c *ConfigSafe) IsDatabaseEnabled() bool {
	return c.Database().Enabled(false)
}

// IsGatewayEnabled 检查网关是否启用
func (c *ConfigSafe) IsGatewayEnabled() bool {
	return c.Gateway().Enabled(false)
}

// IsGrafanaEnabled 检查Grafana是否启用
func (c *ConfigSafe) IsGrafanaEnabled() bool {
	return c.Grafana().Enabled(false)
}

// IsRequestIDEnabled 检查请求ID是否启用
func (c *ConfigSafe) IsRequestIDEnabled() bool {
	return c.RequestID().Enabled(false)
}

// IsRecoveryEnabled 检查恢复中间件是否启用
func (c *ConfigSafe) IsRecoveryEnabled() bool {
	return c.Recovery().Enabled(false)
}

// IsSecurityEnabled 检查安全配置是否启用
func (c *ConfigSafe) IsSecurityEnabled() bool {
	return c.Security().Enabled(false)
}

// IsTracingEnabled 检查链路追踪是否启用
func (c *ConfigSafe) IsTracingEnabled() bool {
	return c.Tracing().Enabled(false)
}

// IsTimeoutEnabled 检查超时配置是否启用
func (c *ConfigSafe) IsTimeoutEnabled() bool {
	return c.TimeoutConfig().Enabled(false)
}

// IsSignatureEnabled 检查签名是否启用
func (c *ConfigSafe) IsSignatureEnabled() bool {
	return c.Signature().Enabled(false)
}

// IsWSCEnabled 检查WSC是否启用
func (c *ConfigSafe) IsWSCEnabled() bool {
	return c.WSC().Enabled(false)
}

// IsDistributedEnabled 检查分布式配置是否启用
func (c *ConfigSafe) IsDistributedEnabled() bool {
	return c.WSC().Distributed().Enabled(false)
}

// IsGroupEnabled 检查群组配置是否启用
func (c *ConfigSafe) IsGroupEnabled() bool {
	return c.WSC().Group().Enabled(false)
}

// IsTicketEnabled 检查票据配置是否启用
func (c *ConfigSafe) IsTicketEnabled() bool {
	return c.WSC().Ticket().Enabled(false)
}

// IsVIPEnabled 检查VIP配置是否启用
func (c *ConfigSafe) IsVIPEnabled() bool {
	return c.WSC().VIP().Enabled(false)
}

// IsEnhancementEnabled 检查增强配置是否启用
func (c *ConfigSafe) IsEnhancementEnabled() bool {
	return c.WSC().Enhancement().Enabled(false)
}

// IsZapEnabled 检查Zap日志是否启用
func (c *ConfigSafe) IsZapEnabled() bool {
	return c.Zap().Enabled(true)
}
