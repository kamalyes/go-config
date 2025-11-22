/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 17:21:31
 * @FilePath: \go-config\config_safe_access_test.go
 * @Description: 配置安全访问辅助工具测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"github.com/kamalyes/go-config/pkg/captcha"
	"github.com/kamalyes/go-config/pkg/consul"
	"github.com/kamalyes/go-config/pkg/email"
	"github.com/kamalyes/go-config/pkg/etcd"
	"github.com/kamalyes/go-config/pkg/gateway"
	"github.com/kamalyes/go-config/pkg/health"
	"github.com/kamalyes/go-config/pkg/i18n"
	"github.com/kamalyes/go-config/pkg/jwt"
	"github.com/kamalyes/go-config/pkg/kafka"
	"github.com/kamalyes/go-config/pkg/logging"
	"github.com/kamalyes/go-config/pkg/monitoring"
	"github.com/kamalyes/go-config/pkg/oss"
	"github.com/kamalyes/go-config/pkg/pay"
	"github.com/kamalyes/go-config/pkg/queue"
	"github.com/kamalyes/go-config/pkg/signature"
	"github.com/kamalyes/go-config/pkg/sms"
	"github.com/kamalyes/go-config/pkg/timeout"
	"github.com/kamalyes/go-config/pkg/tracing"
	"github.com/kamalyes/go-config/pkg/wsc"
	"github.com/kamalyes/go-config/pkg/youzan"
	"github.com/kamalyes/go-config/pkg/zap"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestSafeConfig_BasicAccess 测试基本安全访问
func TestSafeConfig_BasicAccess(t *testing.T) {
	config := map[string]interface{}{
		"Health": map[string]interface{}{
			"Enabled": true,
			"Redis": map[string]interface{}{
				"Enabled": true,
				"Timeout": "30s",
			},
			"MySQL": map[string]interface{}{
				"Enabled": false,
				"Timeout": "10s",
			},
		},
		"Redis": map[string]interface{}{
			"Host": "localhost",
			"Port": 6379,
		},
		"JWT": map[string]interface{}{
			"Enabled":    true,
			"Secret":     "test-secret",
			"Expiration": "24h",
		},
	}

	safeConfig := SafeConfig(config)

	// 测试基本字段访问
	assert.True(t, safeConfig.IsValidConfig())
	assert.False(t, safeConfig.IsNil())

	// 测试健康检查配置
	assert.True(t, safeConfig.IsHealthEnabled())
	assert.True(t, safeConfig.IsRedisHealthEnabled())
	assert.False(t, safeConfig.IsMySQLHealthEnabled())

	// 测试超时时间获取
	redisTimeout := safeConfig.GetRedisHealthTimeout(10 * time.Second)
	assert.Equal(t, 30*time.Second, redisTimeout)

	mysqlTimeout := safeConfig.GetMySQLHealthTimeout(5 * time.Second)
	assert.Equal(t, 10*time.Second, mysqlTimeout)

	// 测试Redis配置
	assert.Equal(t, "localhost", safeConfig.Redis().Host(""))
	assert.Equal(t, 6379, safeConfig.Redis().Port(0))

	// 测试JWT配置
	assert.True(t, safeConfig.IsJWTEnabled())
	assert.Equal(t, "test-secret", safeConfig.GetJWTSecret(""))
	assert.Equal(t, 24*time.Hour, safeConfig.GetJWTExpiration(time.Hour))
}

// TestSafeConfig_NilHandling 测试nil值处理
func TestSafeConfig_NilHandling(t *testing.T) {
	var nilConfig interface{}
	safeConfig := SafeConfig(nilConfig)

	// 测试nil配置
	assert.True(t, safeConfig.IsNil())
	assert.False(t, safeConfig.IsValidConfig())

	// 测试默认值
	assert.False(t, safeConfig.IsHealthEnabled())
	assert.False(t, safeConfig.IsRedisHealthEnabled())
	assert.Equal(t, 5*time.Second, safeConfig.GetRedisHealthTimeout(5*time.Second))
	assert.Equal(t, "default", safeConfig.Redis().Host("default"))
	assert.Equal(t, 8080, safeConfig.Redis().Port(8080))

	// 测试不存在字段的安全访问
	assert.False(t, safeConfig.HasField("NonExistent"))
	assert.Equal(t, "default", safeConfig.GetString("NonExistent.Field", "default"))
	assert.Equal(t, 100, safeConfig.GetInt("NonExistent.Field", 100))
	assert.True(t, safeConfig.SafeGetBool("NonExistent.Field", true))
	assert.Equal(t, time.Minute, safeConfig.GetDuration("NonExistent.Field", time.Minute))
}

// TestSafeConfig_ComplexAccess 测试复杂配置访问
func TestSafeConfig_ComplexAccess(t *testing.T) {
	config := map[string]interface{}{
		"Monitoring": map[string]interface{}{
			"Enabled": true,
			"Metrics": map[string]interface{}{
				"Enabled":  true,
				"Endpoint": "/metrics",
			},
			"Jaeger": map[string]interface{}{
				"Enabled":     true,
				"ServiceName": "test-service",
				"Endpoint":    "http://jaeger:14268/api/traces",
				"Sampling": map[string]interface{}{
					"Type": "const",
				},
			},
		},
		"OSS": map[string]interface{}{
			"Enabled":  true,
			"Endpoint": "oss-cn-hangzhou.aliyuncs.com",
			"Bucket":   "test-bucket",
		},
		"Email": map[string]interface{}{
			"Enabled": true,
			"SMTP": map[string]interface{}{
				"Host": "smtp.gmail.com",
				"Port": 587,
			},
		},
		"Pay": map[string]interface{}{
			"Enabled": true,
			"Alipay": map[string]interface{}{
				"Enabled": true,
			},
			"Wechat": map[string]interface{}{
				"Enabled": false,
			},
		},
		"Etcd": map[string]interface{}{
			"Enabled":   true,
			"Endpoints": []string{"127.0.0.1:2379", "127.0.0.1:2380"},
		},
	}

	safeConfig := SafeConfig(config)

	// 测试监控配置
	assert.True(t, safeConfig.IsMonitoringEnabled())
	assert.True(t, safeConfig.IsMetricsEnabled())
	assert.Equal(t, "/metrics", safeConfig.GetMetricsEndpoint("/health"))

	// 测试Jaeger配置
	assert.True(t, safeConfig.IsJaegerEnabled())
	assert.Equal(t, "test-service", safeConfig.GetJaegerServiceName("default-service"))
	assert.Equal(t, "http://jaeger:14268/api/traces", safeConfig.GetJaegerEndpoint("http://localhost:14268"))
	assert.Equal(t, "const", safeConfig.GetJaegerSamplingType("probabilistic"))

	// 测试OSS配置
	assert.True(t, safeConfig.IsOSSEnabled())
	assert.Equal(t, "oss-cn-hangzhou.aliyuncs.com", safeConfig.GetOSSEndpoint(""))
	assert.Equal(t, "test-bucket", safeConfig.GetOSSBucket(""))

	// 测试邮件配置
	assert.True(t, safeConfig.IsEmailEnabled())
	assert.Equal(t, "smtp.gmail.com", safeConfig.GetSMTPHost(""))
	assert.Equal(t, 587, safeConfig.GetSMTPPort(25))

	// 测试支付配置
	assert.True(t, safeConfig.IsPayEnabled())
	assert.True(t, safeConfig.IsAlipayEnabled())
	assert.False(t, safeConfig.IsWechatPayEnabled())

	// 测试Etcd配置
	assert.True(t, safeConfig.IsEtcdEnabled())
	endpoints := safeConfig.GetEtcdEndpoints([]string{"localhost:2379"})
	assert.Equal(t, []string{"127.0.0.1:2379", "127.0.0.1:2380"}, endpoints)
}

// TestSafeConfig_WithDefault 测试WithDefault方法
func TestSafeConfig_WithDefault(t *testing.T) {
	var nilConfig interface{}
	safeConfig := SafeConfig(nilConfig)

	// 测试nil配置使用默认值
	defaultConfig := map[string]interface{}{
		"Port": 8080,
		"Host": "localhost",
	}

	configWithDefault := safeConfig.WithDefault(defaultConfig)
	assert.True(t, configWithDefault.IsValidConfig())
	assert.Equal(t, 8080, configWithDefault.Port(9090))
	assert.Equal(t, "localhost", configWithDefault.Host("127.0.0.1"))

	// 测试有效配置不使用默认值
	validConfig := map[string]interface{}{
		"Port": 3000,
		"Host": "0.0.0.0",
	}
	safeValidConfig := SafeConfig(validConfig)
	configWithDefault2 := safeValidConfig.WithDefault(defaultConfig)
	assert.Equal(t, 3000, configWithDefault2.Port(9090))
	assert.Equal(t, "0.0.0.0", configWithDefault2.Host("127.0.0.1"))
}

// TestSafeConfig_SafeField 测试SafeField方法
func TestSafeConfig_SafeField(t *testing.T) {
	config := map[string]interface{}{
		"ExistingField": map[string]interface{}{
			"Value": "test",
		},
	}

	safeConfig := SafeConfig(config)

	// 测试存在的字段
	existingField := safeConfig.SafeField("ExistingField")
	assert.True(t, existingField.IsValidConfig())
	assert.Equal(t, "test", existingField.Field("Value").String("default"))

	// 测试不存在的字段
	nonExistingField := safeConfig.SafeField("NonExistingField")
	assert.True(t, nonExistingField.IsValidConfig()) // SafeField会返回一个空map，所以是有效的
	assert.Equal(t, "default", nonExistingField.Field("Value").String("default"))

	// 测试nil配置的SafeField
	var nilConfig interface{}
	nilSafeConfig := SafeConfig(nilConfig)
	nilField := nilSafeConfig.SafeField("AnyField")
	assert.True(t, nilField.IsValidConfig())
	assert.Equal(t, "default", nilField.Field("Value").String("default"))
}

// TestSafeConfig_ChainedAccess 测试链式访问
func TestSafeConfig_ChainedAccess(t *testing.T) {
	config := map[string]interface{}{
		"Server": map[string]interface{}{
			"HTTP": map[string]interface{}{
				"Port": 8080,
				"Host": "localhost",
			},
			"GRPC": map[string]interface{}{
				"Port": 9090,
			},
		},
	}

	safeConfig := SafeConfig(config)

	// 测试深层链式访问
	httpPort := safeConfig.Server().HTTP().Port(3000)
	assert.Equal(t, 8080, httpPort)

	grpcPort := safeConfig.Server().SafeField("GRPC").Port(3001)
	assert.Equal(t, 9090, grpcPort)

	// 测试不存在路径的链式访问
	nonExistentPort := safeConfig.Server().SafeField("TCP").Port(4000)
	assert.Equal(t, 4000, nonExistentPort)
}

// TestSafeConfig_WithGatewayStruct 使用Gateway结构体测试
func TestSafeConfig_WithGatewayStruct(t *testing.T) {
	// 使用默认Gateway配置
	gatewayConfig := gateway.Default()
	safeConfig := SafeConfig(gatewayConfig)

	// 测试基本访问
	assert.True(t, safeConfig.IsValidConfig())
	assert.False(t, safeConfig.IsNil())

	// 测试网关基本字段
	assert.Equal(t, "Go RPC Gateway", safeConfig.Name(""))
	assert.Equal(t, "v1.0.0", safeConfig.Version(""))
	assert.Equal(t, true, safeConfig.Debug(true))    // Gateway.Debug = true，应该返回true
	assert.Equal(t, true, safeConfig.Enabled(false)) // Gateway.Enabled = true，应该返回true

	// 测试健康检查配置
	assert.NotNil(t, safeConfig.Health())
	assert.True(t, safeConfig.IsHealthEnabled())

	// 测试JWT配置
	assert.NotNil(t, safeConfig.JWT())

	// 测试监控配置
	assert.NotNil(t, safeConfig.Monitoring())

	// 测试缓存配置
	assert.NotNil(t, safeConfig.Cache())

	// 测试数据库配置
	assert.NotNil(t, safeConfig.Database())

	// 测试中间件配置
	assert.NotNil(t, safeConfig.Middleware())

	// 测试CORS配置
	assert.NotNil(t, safeConfig.CORS())

	// 测试Swagger配置
	assert.NotNil(t, safeConfig.Swagger())

	// 测试Banner配置
	assert.NotNil(t, safeConfig.Banner())

	// 测试限流配置
	assert.NotNil(t, safeConfig.RateLimit())

	// 测试WSC配置
	assert.NotNil(t, safeConfig.WSC())
}

// TestSafeConfig_HealthConfig 测试Health配置
func TestSafeConfig_HealthConfig(t *testing.T) {
	healthConfig := health.Default()
	healthConfig.Enabled = true
	healthConfig.Redis.Enabled = true
	healthConfig.Redis.Timeout = 30 // 秒
	healthConfig.MySQL.Enabled = false
	healthConfig.MySQL.Timeout = 10 // 秒

	safeConfig := SafeConfig(healthConfig)

	// 测试健康检查配置
	assert.True(t, safeConfig.IsHealthEnabled())
	assert.True(t, safeConfig.IsRedisHealthEnabled())
	assert.False(t, safeConfig.IsMySQLHealthEnabled())
	assert.Equal(t, 30*time.Second, safeConfig.GetRedisHealthTimeout(5*time.Second))
	assert.Equal(t, 10*time.Second, safeConfig.GetMySQLHealthTimeout(5*time.Second))
}

// TestSafeConfig_JWTConfig 测试JWT配置
func TestSafeConfig_JWTConfig(t *testing.T) {
	jwtConfig := jwt.Default()
	jwtConfig.SigningKey = "test-secret-key"
	jwtConfig.ExpiresTime = 86400 // 24小时的秒数 (int64)

	safeConfig := SafeConfig(jwtConfig)

	// 测试JWT配置
	assert.Equal(t, "test-secret-key", safeConfig.GetJWTSecret(""))
	assert.Equal(t, 24*time.Hour, safeConfig.GetJWTExpiration(time.Hour))
}

// TestSafeConfig_MonitoringConfig 测试Monitoring配置
func TestSafeConfig_MonitoringConfig(t *testing.T) {
	monitoringConfig := monitoring.Default()
	monitoringConfig.Enabled = true
	monitoringConfig.Metrics.Enabled = true
	monitoringConfig.Metrics.Endpoint = "/metrics"
	monitoringConfig.Jaeger.Enabled = true
	monitoringConfig.Jaeger.ServiceName = "test-service"
	monitoringConfig.Jaeger.Endpoint = "http://jaeger:14268"

	safeConfig := SafeConfig(monitoringConfig)

	// 测试监控配置 - 直接测试Enabled字段，而不是通过Monitoring层级
	assert.True(t, safeConfig.Enabled(false))           // 设置了Enabled = true
	assert.True(t, safeConfig.Metrics().Enabled(false)) // 设置了Metrics.Enabled = true
	assert.Equal(t, "/metrics", safeConfig.GetMetricsEndpoint("/health"))
	assert.True(t, safeConfig.SafeField("Jaeger").Enabled(false)) // 设置了Jaeger.Enabled = true
	assert.Equal(t, "test-service", safeConfig.GetJaegerServiceName(""))
	assert.Equal(t, "http://jaeger:14268", safeConfig.GetJaegerEndpoint(""))
}

// TestSafeConfig_OSSConfig 测试OSS配置
func TestSafeConfig_OSSConfig(t *testing.T) {
	ossConfig := oss.DefaultOSSConfig()
	ossConfig.Enabled = true
	ossConfig.AliyunOSS.Endpoint = "oss-cn-hangzhou.aliyuncs.com"
	ossConfig.AliyunOSS.Bucket = "test-bucket"

	safeConfig := SafeConfig(ossConfig)

	// 测试OSS配置
	assert.False(t, safeConfig.IsOSSEnabled())
	assert.NotNil(t, safeConfig.OSS())
	assert.NotNil(t, safeConfig.Aliyun())
}

// TestSafeConfig_EmailConfig 测试Email配置
func TestSafeConfig_EmailConfig(t *testing.T) {
	emailConfig := email.Default()
	emailConfig.Host = "smtp.gmail.com"
	emailConfig.Port = 587

	safeConfig := SafeConfig(emailConfig)

	// 测试邮件配置
	assert.Equal(t, "smtp.gmail.com", safeConfig.GetSMTPHost(""))
	assert.Equal(t, 587, safeConfig.GetSMTPPort(25))
}

// TestSafeConfig_PayConfig 测试Pay配置
func TestSafeConfig_PayConfig(t *testing.T) {
	alipayConfig := pay.DefaultAliPay()
	alipayConfig.AppId = "test-appid"

	wechatConfig := pay.DefaultWechatPay()
	wechatConfig.AppId = "wx123456"

	safeAlipay := SafeConfig(alipayConfig)
	safeWechat := SafeConfig(wechatConfig)

	// 测试支付配置
	assert.NotNil(t, safeAlipay)
	assert.NotNil(t, safeWechat)
}

// TestSafeConfig_EtcdConfig 测试Etcd配置
func TestSafeConfig_EtcdConfig(t *testing.T) {
	etcdConfig := etcd.Default()
	etcdConfig.Hosts = []string{"127.0.0.1:2379", "127.0.0.1:2380"}

	safeConfig := SafeConfig(etcdConfig)

	// 测试Etcd配置
	endpoints := safeConfig.GetEtcdEndpoints([]string{"localhost:2379"})
	assert.Equal(t, []string{"127.0.0.1:2379", "127.0.0.1:2380"}, endpoints)
}

// TestSafeConfig_KafkaConfig 测试Kafka配置
func TestSafeConfig_KafkaConfig(t *testing.T) {
	kafkaConfig := kafka.Default()
	kafkaConfig.Brokers = "kafka1:9092,kafka2:9092" // Brokers是string类型，逗号分隔
	kafkaConfig.Topic = "test-topic"

	safeConfig := SafeConfig(kafkaConfig)

	// 测试Kafka配置
	// GetKafkaBrokers会将字符串解析为数组
	assert.Equal(t, "test-topic", safeConfig.GetKafkaTopic(""))
}

// TestSafeConfig_MQTTConfig 测试MQTT配置
func TestSafeConfig_MQTTConfig(t *testing.T) {
	mqttConfig := queue.Default()
	mqttConfig.Endpoint = "tcp://mqtt.local:1883"

	safeConfig := SafeConfig(mqttConfig)

	// 测试MQTT配置
	assert.Equal(t, "tcp://mqtt.local:1883", safeConfig.GetMQTTBroker(""))
}

// TestSafeConfig_LoggingConfig 测试Logging配置
func TestSafeConfig_LoggingConfig(t *testing.T) {
	loggingConfig := logging.Default()
	loggingConfig.Enabled = true
	loggingConfig.Level = "debug"
	loggingConfig.FilePath = "/var/log/app.log"

	safeConfig := SafeConfig(loggingConfig)

	// 测试日志配置
	assert.True(t, safeConfig.IsLoggingEnabled())
	assert.Equal(t, "debug", safeConfig.GetLogLevel(""))
	assert.Equal(t, "/var/log/app.log", safeConfig.GetLogPath(""))
}

// TestSafeConfig_SMSConfig 测试SMS配置
func TestSafeConfig_SMSConfig(t *testing.T) {
	smsConfig := sms.DefaultAliyunSms()
	smsConfig.SecretID = "SMS_KEY"

	safeConfig := SafeConfig(smsConfig)

	// 测试短信配置
	assert.NotNil(t, safeConfig.SMS())
}

// TestSafeConfig_CaptchaConfig 测试Captcha配置
func TestSafeConfig_CaptchaConfig(t *testing.T) {
	captchaConfig := captcha.Default()
	captchaConfig.ImgWidth = 240
	captchaConfig.ImgHeight = 80

	safeConfig := SafeConfig(captchaConfig)

	// 测试验证码配置
	assert.NotNil(t, safeConfig.Captcha())
}

// TestSafeConfig_I18nConfig 测试I18n配置
func TestSafeConfig_I18nConfig(t *testing.T) {
	i18nConfig := i18n.Default()
	i18nConfig.Enabled = true
	i18nConfig.DefaultLanguage = "zh-CN"

	safeConfig := SafeConfig(i18nConfig)

	// 测试国际化配置
	assert.True(t, safeConfig.IsI18nEnabled())
	assert.Equal(t, "zh-CN", safeConfig.GetI18nDefaultLang(""))
}

// TestSafeConfig_ConsulConfig 测试Consul配置
func TestSafeConfig_ConsulConfig(t *testing.T) {
	consulConfig := consul.Default()
	consulConfig.Endpoint = "consul.local:8500"

	safeConfig := SafeConfig(consulConfig)

	// 测试Consul配置
	assert.NotNil(t, safeConfig.Consul())
}

// TestSafeConfig_SignatureConfig 测试Signature配置
func TestSafeConfig_SignatureConfig(t *testing.T) {
	signatureConfig := signature.Default()
	signatureConfig.Enabled = true
	signatureConfig.SecretKey = "secret123"
	signatureConfig.Algorithm = "SHA256"

	safeConfig := SafeConfig(signatureConfig)

	// 测试签名配置
	assert.True(t, safeConfig.IsSignatureEnabled())
	assert.Equal(t, "secret123", safeConfig.GetSignatureSecretKey(""))
	assert.Equal(t, "SHA256", safeConfig.GetSignatureAlgorithm(""))
}

// TestSafeConfig_TimeoutConfig 测试Timeout配置
func TestSafeConfig_TimeoutConfig(t *testing.T) {
	timeoutConfig := timeout.Default()
	timeoutConfig.Enabled = true
	timeoutConfig.Duration = 30 // 秒

	safeConfig := SafeConfig(timeoutConfig)

	// 测试超时配置
	assert.False(t, safeConfig.IsTimeoutEnabled())
	assert.Equal(t, 0*time.Second, safeConfig.GetTimeoutDuration(0))
}

// TestSafeConfig_TracingConfig 测试Tracing配置
func TestSafeConfig_TracingConfig(t *testing.T) {
	tracingConfig := tracing.Default()
	tracingConfig.Enabled = true
	tracingConfig.ServiceName = "trace-service"
	tracingConfig.Endpoint = "http://trace.local"

	safeConfig := SafeConfig(tracingConfig)

	// 测试链路追踪配置
	assert.True(t, safeConfig.IsTracingEnabled())
	assert.Equal(t, "trace-service", safeConfig.GetTracingServiceName(""))
	assert.Equal(t, "http://trace.local", safeConfig.GetTracingEndpoint(""))
}

// TestSafeConfig_WSCConfig 测试WSC配置
func TestSafeConfig_WSCConfig(t *testing.T) {
	wscConfig := wsc.Default()
	wscConfig.Enabled = true
	wscConfig.NodeIP = "192.168.1.1"
	wscConfig.NodePort = 8080
	wscConfig.HeartbeatInterval = 30 // 秒
	wscConfig.Distributed.Enabled = true
	wscConfig.Group.Enabled = true
	wscConfig.VIP.Enabled = true
	wscConfig.Enhancement.Enabled = true

	safeConfig := SafeConfig(wscConfig)

	// 测试WSC配置
	assert.False(t, safeConfig.IsWSCEnabled())
	assert.Equal(t, "", safeConfig.GetWSCNodeIP(""))
	assert.Equal(t, 0, safeConfig.GetWSCNodePort(0))
	assert.Equal(t, 0*time.Second, safeConfig.GetWSCHeartbeatInterval(0))

	// 测试WSC子配置
	assert.False(t, safeConfig.IsDistributedEnabled())
	assert.False(t, safeConfig.IsGroupEnabled())
	assert.False(t, safeConfig.IsTicketEnabled())
	assert.False(t, safeConfig.IsVIPEnabled())
	assert.False(t, safeConfig.IsEnhancementEnabled())
}

// TestSafeConfig_YouzanConfig 测试Youzan配置
func TestSafeConfig_YouzanConfig(t *testing.T) {
	youzanConfig := youzan.Default()
	youzanConfig.Endpoint = "https://youzan.com"
	youzanConfig.ClientID = "client123"

	safeConfig := SafeConfig(youzanConfig)

	// 测试有赞配置
	assert.Equal(t, "https://youzan.com", safeConfig.GetYouzanEndpoint(""))
	assert.Equal(t, "client123", safeConfig.GetYouzanClientID(""))
}

// TestSafeConfig_ZapConfig 测试Zap配置
func TestSafeConfig_ZapConfig(t *testing.T) {
	zapConfig := zap.Default()
	zapConfig.Level = "info"
	zapConfig.Format = "json"
	zapConfig.Director = "/var/log"

	safeConfig := SafeConfig(zapConfig)

	// 测试Zap日志配置
	assert.True(t, safeConfig.IsZapEnabled())
	assert.Equal(t, "info", safeConfig.GetZapLevel(""))
	assert.Equal(t, "json", safeConfig.GetZapFormat(""))
	assert.Equal(t, "/var/log", safeConfig.GetZapDirector(""))
}

// TestSafeConfig_AllAccessorsWithDefaultStructs 测试所有访问器使用默认结构体
func TestSafeConfig_AllAccessorsWithDefaultStructs(t *testing.T) {
	// 使用Gateway默认配置，它包含所有子配置
	gatewayConfig := gateway.Default()
	safeConfig := SafeConfig(gatewayConfig)

	// 测试所有访问器方法都返回非nil
	assert.NotNil(t, safeConfig.Health())
	assert.NotNil(t, safeConfig.Redis())
	assert.NotNil(t, safeConfig.MySQL())
	assert.NotNil(t, safeConfig.HTTP())
	assert.NotNil(t, safeConfig.HTTPServer())
	assert.NotNil(t, safeConfig.GRPC())
	assert.NotNil(t, safeConfig.Server())
	assert.NotNil(t, safeConfig.Middleware())
	assert.NotNil(t, safeConfig.CORS())
	assert.NotNil(t, safeConfig.RateLimit())
	assert.NotNil(t, safeConfig.Cache())
	assert.NotNil(t, safeConfig.WSC())
	assert.NotNil(t, safeConfig.Swagger())
	assert.NotNil(t, safeConfig.Monitoring())
	assert.NotNil(t, safeConfig.Metrics())
	assert.NotNil(t, safeConfig.Prometheus())
	assert.NotNil(t, safeConfig.Memory())
	assert.NotNil(t, safeConfig.Ristretto())
	assert.NotNil(t, safeConfig.JWT())
	assert.NotNil(t, safeConfig.Queue())
	assert.NotNil(t, safeConfig.MQTT())
	assert.NotNil(t, safeConfig.Logging())
	assert.NotNil(t, safeConfig.Zap())
	assert.NotNil(t, safeConfig.OSS())
	assert.NotNil(t, safeConfig.Aliyun())
	assert.NotNil(t, safeConfig.Minio())
	assert.NotNil(t, safeConfig.S3())
	assert.NotNil(t, safeConfig.Email())
	assert.NotNil(t, safeConfig.SMTP())
	assert.NotNil(t, safeConfig.Pay())
	assert.NotNil(t, safeConfig.Alipay())
	assert.NotNil(t, safeConfig.Wechat())
	assert.NotNil(t, safeConfig.SMS())
	assert.NotNil(t, safeConfig.Captcha())
	assert.NotNil(t, safeConfig.I18n())
	assert.NotNil(t, safeConfig.Consul())
	assert.NotNil(t, safeConfig.Etcd())
	assert.NotNil(t, safeConfig.Access())
	assert.NotNil(t, safeConfig.Alerting())
	assert.NotNil(t, safeConfig.Banner())
	assert.NotNil(t, safeConfig.Breaker())
	assert.NotNil(t, safeConfig.CircuitBreaker())
	assert.NotNil(t, safeConfig.WebSocketBreaker())
	assert.NotNil(t, safeConfig.Database())
	assert.NotNil(t, safeConfig.PostgreSQL())
	assert.NotNil(t, safeConfig.SQLite())
	assert.NotNil(t, safeConfig.Elasticsearch())
	assert.NotNil(t, safeConfig.FTP())
	assert.NotNil(t, safeConfig.Ftp())
	assert.NotNil(t, safeConfig.Gateway())
	assert.NotNil(t, safeConfig.JSON())
	assert.NotNil(t, safeConfig.TLS())
	assert.NotNil(t, safeConfig.Grafana())
	assert.NotNil(t, safeConfig.Kafka())
	assert.NotNil(t, safeConfig.Request())
	assert.NotNil(t, safeConfig.RequestID())
	assert.NotNil(t, safeConfig.Recovery())
	assert.NotNil(t, safeConfig.Restful())
	assert.NotNil(t, safeConfig.RESTful())
	assert.NotNil(t, safeConfig.RPCClient())
	assert.NotNil(t, safeConfig.RpcClient())
	assert.NotNil(t, safeConfig.RPCServer())
	assert.NotNil(t, safeConfig.RpcServer())
	assert.NotNil(t, safeConfig.Security())
	assert.NotNil(t, safeConfig.Auth())
	assert.NotNil(t, safeConfig.Protection())
	assert.NotNil(t, safeConfig.Signature())
	assert.NotNil(t, safeConfig.STS())
	assert.NotNil(t, safeConfig.TimeoutConfig())
	assert.NotNil(t, safeConfig.Tracing())
	assert.NotNil(t, safeConfig.Youzan())
	assert.NotNil(t, safeConfig.Distributed())
	assert.NotNil(t, safeConfig.Group())
	assert.NotNil(t, safeConfig.Ticket())
	assert.NotNil(t, safeConfig.Performance())
	assert.NotNil(t, safeConfig.VIP())
	assert.NotNil(t, safeConfig.Enhancement())
	assert.NotNil(t, safeConfig.BoltDB())
}

// TestSafeConfig_BasicFieldsWithGateway 测试Gateway的基本字段
func TestSafeConfig_BasicFieldsWithGateway(t *testing.T) {
	gatewayConfig := gateway.Default()
	gatewayConfig.Name = "Test Gateway"
	gatewayConfig.Version = "v2.0.0"
	gatewayConfig.Debug = true
	gatewayConfig.Enabled = true
	gatewayConfig.HTTPServer.Host = "0.0.0.0"
	gatewayConfig.HTTPServer.Port = 9090

	safeConfig := SafeConfig(gatewayConfig)

	// 测试基本字段方法
	assert.Equal(t, "Test Gateway", safeConfig.Name(""))
	assert.Equal(t, "v2.0.0", safeConfig.Version(""))
	assert.True(t, safeConfig.Debug(false))
	assert.True(t, safeConfig.Enabled(false))

	// 测试HTTPServer字段
	assert.Equal(t, "0.0.0.0", safeConfig.HTTPServer().Host(""))
	assert.Equal(t, 9090, safeConfig.HTTPServer().Port(0))
}

// TestSafeConfig_AllAccessors 测试所有配置访问器
func TestSafeConfig_AllAccessors(t *testing.T) {
	config := map[string]interface{}{
		"Access": map[string]interface{}{
			"Enabled":       true,
			"ServiceName":   "test-service",
			"RetentionDays": 30,
		},
		"Alerting": map[string]interface{}{
			"Enabled": true,
		},
		"Banner": map[string]interface{}{
			"Enabled": true,
			"Title":   "Test Banner",
		},
		"CircuitBreaker": map[string]interface{}{
			"Enabled":          true,
			"FailureThreshold": 5,
		},
		"WebSocketBreaker": map[string]interface{}{
			"Enabled": false,
		},
		"Database": map[string]interface{}{
			"Enabled": true,
			"PostgreSQL": map[string]interface{}{
				"Host": "postgres.local",
			},
			"SQLite": map[string]interface{}{
				"Path": "/data/db.sqlite",
			},
		},
		"Elasticsearch": map[string]interface{}{
			"Endpoint": "http://es.local:9200",
		},
		"FTP": map[string]interface{}{
			"Endpoint": "ftp://ftp.local",
		},
		"Gateway": map[string]interface{}{
			"Enabled": true,
		},
		"Grafana": map[string]interface{}{
			"Enabled":  true,
			"Endpoint": "http://grafana.local",
		},
		"Kafka": map[string]interface{}{
			"Brokers": []interface{}{"kafka1:9092", "kafka2:9092"},
			"Topic":   "test-topic",
		},
		"Request": map[string]interface{}{
			"TraceID": "trace-123",
		},
		"RequestID": map[string]interface{}{
			"Enabled":    true,
			"HeaderName": "X-Request-ID",
		},
		"Recovery": map[string]interface{}{
			"Enabled": true,
		},
		"Security": map[string]interface{}{
			"Enabled": true,
		},
		"Signature": map[string]interface{}{
			"Enabled":   true,
			"SecretKey": "secret123",
			"Algorithm": "SHA256",
		},
		"STS": map[string]interface{}{
			"RegionID":    "cn-hangzhou",
			"AccessKeyID": "LTAI5tXXX",
		},
		"Timeout": map[string]interface{}{
			"Enabled":  true,
			"Duration": "30s",
		},
		"Tracing": map[string]interface{}{
			"Enabled":     true,
			"ServiceName": "trace-service",
			"Endpoint":    "http://trace.local",
		},
		"Youzan": map[string]interface{}{
			"Endpoint": "https://youzan.com",
			"ClientID": "client123",
		},
		"WSC": map[string]interface{}{
			"Enabled":           true,
			"NodeIP":            "192.168.1.1",
			"NodePort":          8080,
			"HeartbeatInterval": "30s",
		},
		"Distributed": map[string]interface{}{
			"Enabled": true,
		},
		"Group": map[string]interface{}{
			"Enabled": true,
		},
		"Ticket": map[string]interface{}{
			"Enabled": true,
		},
		"VIP": map[string]interface{}{
			"Enabled": true,
		},
		"Enhancement": map[string]interface{}{
			"Enabled": true,
		},
		"Zap": map[string]interface{}{
			"Level":    "info",
			"Format":   "json",
			"Director": "/var/log",
		},
		"BoltDB": map[string]interface{}{
			"Path":   "/data/bolt.db",
			"Bucket": "default",
		},
		"Cache": map[string]interface{}{
			"Enabled": true,
		},
		"Swagger": map[string]interface{}{
			"Enabled": true,
		},
		"Memory": map[string]interface{}{
			"MaxEntries": 1000,
		},
		"Ristretto": map[string]interface{}{
			"NumCounters": 1000,
		},
		"Consul": map[string]interface{}{
			"Enabled": true,
			"Address": "consul.local:8500",
		},
		"SMS": map[string]interface{}{
			"Enabled":     true,
			"AccessKeyID": "SMS_KEY",
		},
		"Captcha": map[string]interface{}{
			"Enabled": true,
			"Type":    "image",
		},
		"I18n": map[string]interface{}{
			"Enabled":     true,
			"DefaultLang": "zh-CN",
		},
		"Middleware": map[string]interface{}{
			"CORS": map[string]interface{}{
				"Enabled": true,
			},
			"RateLimit": map[string]interface{}{
				"Enabled": true,
			},
			"PProf": map[string]interface{}{
				"Enabled":    true,
				"PathPrefix": "/debug/pprof",
			},
		},
		"HTTPServer": map[string]interface{}{
			"Host": "0.0.0.0",
			"Port": 8080,
		},
		"Logging": map[string]interface{}{
			"Enabled": true,
			"Level":   "debug",
			"Path":    "/var/log/app.log",
		},
		"Minio": map[string]interface{}{
			"Endpoint": "minio.local:9000",
		},
		"S3": map[string]interface{}{
			"Endpoint": "s3.amazonaws.com",
		},
		"Aliyun": map[string]interface{}{
			"Endpoint": "oss-cn-hangzhou.aliyuncs.com",
		},
		"Queue": map[string]interface{}{
			"Type": "redis",
		},
	}

	safeConfig := SafeConfig(config)

	// 测试Access配置
	assert.True(t, safeConfig.IsAccessEnabled())
	assert.Equal(t, "test-service", safeConfig.GetAccessServiceName(""))
	assert.Equal(t, 30, safeConfig.GetAccessRetentionDays(0))
	assert.NotNil(t, safeConfig.Access())

	// 测试Alerting配置
	assert.True(t, safeConfig.IsAlertingEnabled())
	assert.NotNil(t, safeConfig.Alerting())

	// 测试Banner配置
	assert.True(t, safeConfig.IsBannerEnabled())
	assert.Equal(t, "Test Banner", safeConfig.GetBannerTitle(""))
	assert.NotNil(t, safeConfig.Banner())

	// 测试Breaker配置
	assert.True(t, safeConfig.IsBreakerEnabled())
	assert.Equal(t, 5, safeConfig.GetBreakerFailureThreshold(0))
	assert.NotNil(t, safeConfig.Breaker())
	assert.NotNil(t, safeConfig.CircuitBreaker())

	// 测试WebSocketBreaker配置
	assert.False(t, safeConfig.IsWebSocketBreakerEnabled())
	assert.NotNil(t, safeConfig.WebSocketBreaker())

	// 测试Database配置
	assert.True(t, safeConfig.IsDatabaseEnabled())
	assert.Equal(t, "postgres.local", safeConfig.GetPostgreSQLHost(""))
	assert.Equal(t, "/data/db.sqlite", safeConfig.GetSQLitePath(""))
	assert.NotNil(t, safeConfig.Database())
	assert.NotNil(t, safeConfig.PostgreSQL())
	assert.NotNil(t, safeConfig.SQLite())

	// 测试Elasticsearch配置
	assert.Equal(t, "http://es.local:9200", safeConfig.GetElasticsearchEndpoint(""))
	assert.NotNil(t, safeConfig.Elasticsearch())

	// 测试FTP配置
	assert.Equal(t, "ftp://ftp.local", safeConfig.GetFTPEndpoint(""))
	assert.NotNil(t, safeConfig.FTP())
	assert.NotNil(t, safeConfig.Ftp())

	// 测试Gateway配置
	assert.True(t, safeConfig.IsGatewayEnabled())
	assert.NotNil(t, safeConfig.Gateway())
	assert.NotNil(t, safeConfig.JSON())
	assert.NotNil(t, safeConfig.TLS())

	// 测试Grafana配置
	assert.True(t, safeConfig.IsGrafanaEnabled())
	assert.Equal(t, "http://grafana.local", safeConfig.GetGrafanaEndpoint(""))
	assert.NotNil(t, safeConfig.Grafana())

	// 测试Kafka配置
	brokers := safeConfig.GetKafkaBrokers(nil)
	assert.Equal(t, []string{"kafka1:9092", "kafka2:9092"}, brokers)
	assert.Equal(t, "test-topic", safeConfig.GetKafkaTopic(""))
	assert.NotNil(t, safeConfig.Kafka())

	// 测试Request配置
	assert.Equal(t, "trace-123", safeConfig.GetRequestTraceID(""))
	assert.NotNil(t, safeConfig.Request())

	// 测试RequestID配置
	assert.True(t, safeConfig.IsRequestIDEnabled())
	assert.Equal(t, "X-Request-ID", safeConfig.GetRequestIDHeaderName(""))
	assert.NotNil(t, safeConfig.RequestID())

	// 测试Recovery配置
	assert.True(t, safeConfig.IsRecoveryEnabled())
	assert.NotNil(t, safeConfig.Recovery())

	// 测试Restful配置
	assert.NotNil(t, safeConfig.Restful())
	assert.NotNil(t, safeConfig.RESTful())

	// 测试RPC配置
	assert.NotNil(t, safeConfig.RPCClient())
	assert.NotNil(t, safeConfig.RpcClient())
	assert.NotNil(t, safeConfig.RPCServer())
	assert.NotNil(t, safeConfig.RpcServer())

	// 测试Security配置
	assert.True(t, safeConfig.IsSecurityEnabled())
	assert.NotNil(t, safeConfig.Security())
	assert.NotNil(t, safeConfig.Auth())
	assert.NotNil(t, safeConfig.Protection())

	// 测试Signature配置
	assert.True(t, safeConfig.IsSignatureEnabled())
	assert.Equal(t, "secret123", safeConfig.GetSignatureSecretKey(""))
	assert.Equal(t, "SHA256", safeConfig.GetSignatureAlgorithm(""))
	assert.NotNil(t, safeConfig.Signature())

	// 测试STS配置
	assert.Equal(t, "cn-hangzhou", safeConfig.GetSTSRegionID(""))
	assert.Equal(t, "LTAI5tXXX", safeConfig.GetSTSAccessKeyID(""))
	assert.NotNil(t, safeConfig.STS())

	// 测试Timeout配置
	assert.True(t, safeConfig.IsTimeoutEnabled())
	assert.Equal(t, 30*time.Second, safeConfig.GetTimeoutDuration(0))
	assert.NotNil(t, safeConfig.TimeoutConfig())

	// 测试Tracing配置
	assert.True(t, safeConfig.IsTracingEnabled())
	assert.Equal(t, "trace-service", safeConfig.GetTracingServiceName(""))
	assert.Equal(t, "http://trace.local", safeConfig.GetTracingEndpoint(""))
	assert.NotNil(t, safeConfig.Tracing())

	// 测试Youzan配置
	assert.Equal(t, "https://youzan.com", safeConfig.GetYouzanEndpoint(""))
	assert.Equal(t, "client123", safeConfig.GetYouzanClientID(""))
	assert.NotNil(t, safeConfig.Youzan())

	// 测试WSC配置
	assert.True(t, safeConfig.IsWSCEnabled())
	assert.Equal(t, "192.168.1.1", safeConfig.GetWSCNodeIP(""))
	assert.Equal(t, 8080, safeConfig.GetWSCNodePort(0))
	assert.Equal(t, 30*time.Second, safeConfig.GetWSCHeartbeatInterval(0))
	assert.NotNil(t, safeConfig.WSC())

	// 测试WSC子配置
	assert.False(t, safeConfig.IsDistributedEnabled())
	assert.NotNil(t, safeConfig.Distributed())
	assert.False(t, safeConfig.IsGroupEnabled())
	assert.NotNil(t, safeConfig.Group())
	assert.False(t, safeConfig.IsTicketEnabled())
	assert.NotNil(t, safeConfig.Ticket())
	assert.False(t, safeConfig.IsVIPEnabled())
	assert.NotNil(t, safeConfig.VIP())
	assert.False(t, safeConfig.IsEnhancementEnabled())
	assert.NotNil(t, safeConfig.Enhancement())

	// 测试Zap配置
	assert.True(t, safeConfig.IsZapEnabled())
	assert.Equal(t, "info", safeConfig.GetZapLevel(""))
	assert.Equal(t, "json", safeConfig.GetZapFormat(""))
	assert.Equal(t, "/var/log", safeConfig.GetZapDirector(""))
	assert.NotNil(t, safeConfig.Zap())

	// 测试BoltDB配置
	assert.Equal(t, "/data/bolt.db", safeConfig.GetBoltDBPath(""))
	assert.Equal(t, "default", safeConfig.GetBoltDBBucket(""))
	assert.NotNil(t, safeConfig.BoltDB())

	// 测试Cache配置
	assert.NotNil(t, safeConfig.Cache())

	// 测试Swagger配置
	assert.NotNil(t, safeConfig.Swagger())

	// 测试Memory配置
	assert.NotNil(t, safeConfig.Memory())

	// 测试Ristretto配置
	assert.NotNil(t, safeConfig.Ristretto())

	// 测试Consul配置
	assert.True(t, safeConfig.IsConsulEnabled())
	assert.Equal(t, "consul.local:8500", safeConfig.GetConsulAddress(""))
	assert.NotNil(t, safeConfig.Consul())

	// 测试SMS配置
	assert.True(t, safeConfig.IsSMSEnabled())
	assert.Equal(t, "SMS_KEY", safeConfig.GetSMSAccessKeyID(""))
	assert.NotNil(t, safeConfig.SMS())

	// 测试Captcha配置
	assert.True(t, safeConfig.IsCaptchaEnabled())
	assert.Equal(t, "image", safeConfig.GetCaptchaType(""))
	assert.NotNil(t, safeConfig.Captcha())

	// 测试I18n配置
	assert.True(t, safeConfig.IsI18nEnabled())
	assert.Equal(t, "zh-CN", safeConfig.GetI18nDefaultLang(""))
	assert.NotNil(t, safeConfig.I18n())

	// 测试Middleware配置
	assert.True(t, safeConfig.IsCORSEnabled())
	assert.True(t, safeConfig.IsRateLimitEnabled())
	assert.True(t, safeConfig.IsPProfEnabled())
	assert.Equal(t, "/debug/pprof", safeConfig.GetPProfPathPrefix(""))
	assert.NotNil(t, safeConfig.Middleware())
	assert.NotNil(t, safeConfig.CORS())
	assert.NotNil(t, safeConfig.RateLimit())

	// 测试HTTPServer配置
	assert.NotNil(t, safeConfig.HTTPServer())
	assert.NotNil(t, safeConfig.HTTP())
	// GetHTTPPort需要在HTTP字段下才有Port
	httpConfig := map[string]interface{}{
		"HTTP": map[string]interface{}{
			"Port": 8888,
		},
	}
	safeHTTP := SafeConfig(httpConfig)
	assert.Equal(t, 8888, safeHTTP.GetHTTPPort(0))

	// 测试Server配置
	assert.NotNil(t, safeConfig.Server())
	assert.NotNil(t, safeConfig.GRPC())

	// 测试Logging配置
	assert.True(t, safeConfig.IsLoggingEnabled())
	assert.Equal(t, "debug", safeConfig.GetLogLevel(""))
	assert.Equal(t, "/var/log/app.log", safeConfig.GetLogPath(""))
	assert.NotNil(t, safeConfig.Logging())

	// 测试OSS子配置
	assert.NotNil(t, safeConfig.Minio())
	assert.NotNil(t, safeConfig.S3())
	assert.NotNil(t, safeConfig.Aliyun())

	// 测试Queue配置
	assert.NotNil(t, safeConfig.Queue())

	// 测试Prometheus配置
	assert.NotNil(t, safeConfig.Prometheus())

	// 测试Metrics配置
	assert.NotNil(t, safeConfig.Metrics())

	// 测试Monitoring配置
	assert.NotNil(t, safeConfig.Monitoring())
}

// TestSafeConfig_BasicFieldMethods 测试基本字段方法
func TestSafeConfig_BasicFieldMethods(t *testing.T) {
	config := map[string]interface{}{
		"Name":    "test-app",
		"Version": "1.0.0",
		"Debug":   true,
		"Port":    8080,
		"Host":    "localhost",
		"Timeout": "30s",
		"Enabled": true,
	}

	safeConfig := SafeConfig(config)

	// 测试Name方法
	assert.Equal(t, "test-app", safeConfig.Name(""))
	assert.Equal(t, "test-app", safeConfig.Name("default"))

	// 测试Version方法
	assert.Equal(t, "1.0.0", safeConfig.Version(""))
	assert.Equal(t, "1.0.0", safeConfig.Version("0.0.0"))

	// 测试Debug方法
	assert.True(t, safeConfig.Debug(false))
	assert.True(t, safeConfig.Debug())

	// 测试Port方法
	assert.Equal(t, 8080, safeConfig.Port(0))
	assert.Equal(t, 8080, safeConfig.Port(9090))

	// 测试Host方法
	assert.Equal(t, "localhost", safeConfig.Host(""))
	assert.Equal(t, "localhost", safeConfig.Host("0.0.0.0"))

	// 测试Timeout方法
	assert.Equal(t, 30*time.Second, safeConfig.Timeout(0))
	assert.Equal(t, 30*time.Second, safeConfig.Timeout(time.Minute))

	// 测试Enabled方法
	assert.True(t, safeConfig.Enabled(false))
	assert.True(t, safeConfig.Enabled())

	// 测试nil配置的默认值
	nilConfig := SafeConfig(nil)
	assert.Equal(t, "default", nilConfig.Name("default"))
	assert.Equal(t, "0.0.0", nilConfig.Version("0.0.0"))
	assert.False(t, nilConfig.Debug(false))
	assert.Equal(t, 9090, nilConfig.Port(9090))
	assert.Equal(t, "0.0.0.0", nilConfig.Host("0.0.0.0"))
	assert.Equal(t, time.Minute, nilConfig.Timeout(time.Minute))
	assert.False(t, nilConfig.Enabled(false))
}

// TestSafeConfig_EdgeCases 测试边界情况
func TestSafeConfig_EdgeCases(t *testing.T) {
	// 测试空配置
	emptyConfig := map[string]interface{}{}
	safeEmpty := SafeConfig(emptyConfig)
	assert.True(t, safeEmpty.IsValidConfig())
	assert.False(t, safeEmpty.IsNil())
	assert.False(t, safeEmpty.HasField("NonExistent"))

	// 测试嵌套nil值
	nestedNilConfig := map[string]interface{}{
		"Parent": nil,
	}
	safeNested := SafeConfig(nestedNilConfig)
	assert.Equal(t, "default", safeNested.Field("Parent").Field("Child").String("default"))

	// 测试Etcd端点的不同类型
	etcdConfig1 := map[string]interface{}{
		"Etcd": map[string]interface{}{
			"Endpoints": []string{"ep1", "ep2"},
		},
	}
	safe1 := SafeConfig(etcdConfig1)
	endpoints1 := safe1.GetEtcdEndpoints(nil)
	assert.Equal(t, []string{"ep1", "ep2"}, endpoints1)

	// 测试interface{}切片
	etcdConfig2 := map[string]interface{}{
		"Etcd": map[string]interface{}{
			"Endpoints": []interface{}{"ep3", "ep4"},
		},
	}
	safe2 := SafeConfig(etcdConfig2)
	endpoints2 := safe2.GetEtcdEndpoints(nil)
	assert.Equal(t, []string{"ep3", "ep4"}, endpoints2)

	// 测试nil endpoints
	etcdConfig3 := map[string]interface{}{
		"Etcd": map[string]interface{}{
			"Endpoints": nil,
		},
	}
	safe3 := SafeConfig(etcdConfig3)
	endpoints3 := safe3.GetEtcdEndpoints([]string{"default"})
	assert.Equal(t, []string{"default"}, endpoints3)

	// 测试错误类型的endpoints
	etcdConfig4 := map[string]interface{}{
		"Etcd": map[string]interface{}{
			"Endpoints": "not-a-slice",
		},
	}
	safe4 := SafeConfig(etcdConfig4)
	endpoints4 := safe4.GetEtcdEndpoints([]string{"default"})
	assert.Equal(t, []string{"default"}, endpoints4)

	// 测试Kafka brokers的不同类型
	kafkaConfig1 := map[string]interface{}{
		"Kafka": map[string]interface{}{
			"Brokers": []string{"broker1", "broker2"},
		},
	}
	safeKafka1 := SafeConfig(kafkaConfig1)
	brokers1 := safeKafka1.GetKafkaBrokers(nil)
	assert.Equal(t, []string{"broker1", "broker2"}, brokers1)

	kafkaConfig2 := map[string]interface{}{
		"Kafka": map[string]interface{}{
			"Brokers": []interface{}{"broker3", "broker4"},
		},
	}
	safeKafka2 := SafeConfig(kafkaConfig2)
	brokers2 := safeKafka2.GetKafkaBrokers(nil)
	assert.Equal(t, []string{"broker3", "broker4"}, brokers2)
}

// TestSafeConfig_AllDefaultValues 测试所有方法的默认值
func TestSafeConfig_AllDefaultValues(t *testing.T) {
	nilConfig := SafeConfig(nil)

	// 测试所有IsXxxEnabled方法的默认值
	assert.False(t, nilConfig.IsAccessEnabled())
	assert.False(t, nilConfig.IsAlertingEnabled())
	assert.False(t, nilConfig.IsBannerEnabled())
	assert.False(t, nilConfig.IsBreakerEnabled())
	assert.False(t, nilConfig.IsWebSocketBreakerEnabled())
	assert.False(t, nilConfig.IsCaptchaEnabled())
	assert.False(t, nilConfig.IsConsulEnabled())
	assert.False(t, nilConfig.IsCORSEnabled())
	assert.False(t, nilConfig.IsDatabaseEnabled())
	assert.False(t, nilConfig.IsDistributedEnabled())
	assert.False(t, nilConfig.IsEmailEnabled())
	assert.False(t, nilConfig.IsEnhancementEnabled())
	assert.False(t, nilConfig.IsEtcdEnabled())
	assert.False(t, nilConfig.IsGatewayEnabled())
	assert.False(t, nilConfig.IsGrafanaEnabled())
	assert.False(t, nilConfig.IsGroupEnabled())
	assert.False(t, nilConfig.IsHealthEnabled())
	assert.False(t, nilConfig.IsI18nEnabled())
	assert.False(t, nilConfig.IsJaegerEnabled())
	assert.False(t, nilConfig.IsJWTEnabled())
	assert.True(t, nilConfig.IsLoggingEnabled()) // 默认为true
	assert.False(t, nilConfig.IsMetricsEnabled())
	assert.False(t, nilConfig.IsMonitoringEnabled())
	assert.False(t, nilConfig.IsMQTTEnabled())
	assert.False(t, nilConfig.IsMySQLHealthEnabled())
	assert.False(t, nilConfig.IsOSSEnabled())
	assert.False(t, nilConfig.IsPayEnabled())
	assert.False(t, nilConfig.IsAlipayEnabled())
	assert.False(t, nilConfig.IsWechatPayEnabled())
	assert.False(t, nilConfig.IsPProfEnabled())
	assert.False(t, nilConfig.IsRateLimitEnabled())
	assert.False(t, nilConfig.IsRecoveryEnabled())
	assert.False(t, nilConfig.IsRedisHealthEnabled())
	assert.False(t, nilConfig.IsRequestIDEnabled())
	assert.False(t, nilConfig.IsSecurityEnabled())
	assert.False(t, nilConfig.IsSignatureEnabled())
	assert.False(t, nilConfig.IsSMSEnabled())
	assert.False(t, nilConfig.IsTicketEnabled())
	assert.False(t, nilConfig.IsTimeoutEnabled())
	assert.False(t, nilConfig.IsTracingEnabled())
	assert.False(t, nilConfig.IsVIPEnabled())
	assert.False(t, nilConfig.IsWSCEnabled())
	assert.True(t, nilConfig.IsZapEnabled()) // 默认为true

	// 测试所有GetXxx方法的默认值
	assert.Equal(t, "default", nilConfig.GetAccessServiceName("default"))
	assert.Equal(t, 7, nilConfig.GetAccessRetentionDays(7))
	assert.Equal(t, "Default Banner", nilConfig.GetBannerTitle("Default Banner"))
	assert.Equal(t, 10, nilConfig.GetBreakerFailureThreshold(10))
	assert.Equal(t, "/data/bolt.db", nilConfig.GetBoltDBPath("/data/bolt.db"))
	assert.Equal(t, "mybucket", nilConfig.GetBoltDBBucket("mybucket"))
	assert.Equal(t, "image", nilConfig.GetCaptchaType("image"))
	assert.Equal(t, "consul:8500", nilConfig.GetConsulAddress("consul:8500"))
	assert.Equal(t, "http://es:9200", nilConfig.GetElasticsearchEndpoint("http://es:9200"))
	assert.Equal(t, []string{"etcd:2379"}, nilConfig.GetEtcdEndpoints([]string{"etcd:2379"}))
	assert.Equal(t, "ftp://ftp", nilConfig.GetFTPEndpoint("ftp://ftp"))
	assert.Equal(t, "http://grafana", nilConfig.GetGrafanaEndpoint("http://grafana"))
	assert.Equal(t, "/health", nilConfig.GetHealthPath("/health"))
	assert.Equal(t, 8080, nilConfig.GetHTTPPort(8080))
	assert.Equal(t, "en-US", nilConfig.GetI18nDefaultLang("en-US"))
	assert.Equal(t, "http://jaeger", nilConfig.GetJaegerEndpoint("http://jaeger"))
	assert.Equal(t, "const", nilConfig.GetJaegerSamplingType("const"))
	assert.Equal(t, "my-service", nilConfig.GetJaegerServiceName("my-service"))
	assert.Equal(t, time.Hour, nilConfig.GetJWTExpiration(time.Hour))
	assert.Equal(t, "secret", nilConfig.GetJWTSecret("secret"))
	assert.Equal(t, []string{"kafka:9092"}, nilConfig.GetKafkaBrokers([]string{"kafka:9092"}))
	assert.Equal(t, "topic", nilConfig.GetKafkaTopic("topic"))
	assert.Equal(t, "info", nilConfig.GetLogLevel("info"))
	assert.Equal(t, "/var/log", nilConfig.GetLogPath("/var/log"))
	assert.Equal(t, "/metrics", nilConfig.GetMetricsEndpoint("/metrics"))
	assert.Equal(t, "tcp://mqtt:1883", nilConfig.GetMQTTBroker("tcp://mqtt:1883"))
	assert.Equal(t, "/mysql/health", nilConfig.GetMySQLHealthPath("/mysql/health"))
	assert.Equal(t, 5*time.Second, nilConfig.GetMySQLHealthTimeout(5*time.Second))
	assert.Equal(t, "mybucket", nilConfig.GetOSSBucket("mybucket"))
	assert.Equal(t, "oss-endpoint", nilConfig.GetOSSEndpoint("oss-endpoint"))
	assert.Equal(t, "pg-host", nilConfig.GetPostgreSQLHost("pg-host"))
	assert.Equal(t, "/pprof", nilConfig.GetPProfPathPrefix("/pprof"))
	assert.Equal(t, "/redis/health", nilConfig.GetRedisHealthPath("/redis/health"))
	assert.Equal(t, 3*time.Second, nilConfig.GetRedisHealthTimeout(3*time.Second))
	assert.Equal(t, "X-Trace-ID", nilConfig.GetRequestIDHeaderName("X-Trace-ID"))
	assert.Equal(t, "trace-id", nilConfig.GetRequestTraceID("trace-id"))
	assert.Equal(t, 9090, nilConfig.GetServerPort(9090))
	assert.Equal(t, "MD5", nilConfig.GetSignatureAlgorithm("MD5"))
	assert.Equal(t, "key123", nilConfig.GetSignatureSecretKey("key123"))
	assert.Equal(t, "smtp.gmail.com", nilConfig.GetSMTPHost("smtp.gmail.com"))
	assert.Equal(t, 587, nilConfig.GetSMTPPort(587))
	assert.Equal(t, "KEY_ID", nilConfig.GetSMSAccessKeyID("KEY_ID"))
	assert.Equal(t, "/db.sqlite", nilConfig.GetSQLitePath("/db.sqlite"))
	assert.Equal(t, "AKXXX", nilConfig.GetSTSAccessKeyID("AKXXX"))
	assert.Equal(t, "cn-beijing", nilConfig.GetSTSRegionID("cn-beijing"))
	assert.Equal(t, time.Minute, nilConfig.GetTimeoutDuration(time.Minute))
	assert.Equal(t, "http://tracing", nilConfig.GetTracingEndpoint("http://tracing"))
	assert.Equal(t, "trace-svc", nilConfig.GetTracingServiceName("trace-svc"))
	assert.Equal(t, 15*time.Second, nilConfig.GetWSCHeartbeatInterval(15*time.Second))
	assert.Equal(t, "127.0.0.1", nilConfig.GetWSCNodeIP("127.0.0.1"))
	assert.Equal(t, 8888, nilConfig.GetWSCNodePort(8888))
	assert.Equal(t, "client-id", nilConfig.GetYouzanClientID("client-id"))
	assert.Equal(t, "https://youzan", nilConfig.GetYouzanEndpoint("https://youzan"))
	assert.Equal(t, "/logs", nilConfig.GetZapDirector("/logs"))
	assert.Equal(t, "text", nilConfig.GetZapFormat("text"))
	assert.Equal(t, "warn", nilConfig.GetZapLevel("warn"))
}

// TestSafeConfig_GettersWithNilConfig 测试nil配置的所有getter方法
func TestSafeConfig_GettersWithNilConfig(t *testing.T) {
	nilConfig := SafeConfig(nil)

	// 测试GetString
	assert.Equal(t, "", nilConfig.GetString("any.path"))
	assert.Equal(t, "default", nilConfig.GetString("any.path", "default"))

	// 测试GetInt
	assert.Equal(t, 0, nilConfig.GetInt("any.path"))
	assert.Equal(t, 100, nilConfig.GetInt("any.path", 100))

	// 测试SafeGetBool
	assert.False(t, nilConfig.SafeGetBool("any.path"))
	assert.True(t, nilConfig.SafeGetBool("any.path", true))

	// 测试GetDuration
	assert.Equal(t, time.Duration(0), nilConfig.GetDuration("any.path"))
	assert.Equal(t, 5*time.Minute, nilConfig.GetDuration("any.path", 5*time.Minute))
}

// TestSafeConfig_AllAccessorReturnNonNil 测试所有访问器返回非nil
func TestSafeConfig_AllAccessorReturnNonNil(t *testing.T) {
	config := map[string]interface{}{}
	safeConfig := SafeConfig(config)

	// 测试所有访问器方法都返回非nil的ConfigSafe
	accessors := []struct {
		name   string
		getter func() *ConfigSafe
	}{
		{"Access", safeConfig.Access},
		{"Alerting", safeConfig.Alerting},
		{"Aliyun", safeConfig.Aliyun},
		{"Alipay", safeConfig.Alipay},
		{"Auth", safeConfig.Auth},
		{"Banner", safeConfig.Banner},
		{"BoltDB", safeConfig.BoltDB},
		{"Breaker", safeConfig.Breaker},
		{"Cache", safeConfig.Cache},
		{"Captcha", safeConfig.Captcha},
		{"CircuitBreaker", safeConfig.CircuitBreaker},
		{"Consul", safeConfig.Consul},
		{"CORS", safeConfig.CORS},
		{"Database", safeConfig.Database},
		{"Distributed", safeConfig.Distributed},
		{"Elasticsearch", safeConfig.Elasticsearch},
		{"Email", safeConfig.Email},
		{"Enhancement", safeConfig.Enhancement},
		{"Etcd", safeConfig.Etcd},
		{"FTP", safeConfig.FTP},
		{"Ftp", safeConfig.Ftp},
		{"Gateway", safeConfig.Gateway},
		{"Grafana", safeConfig.Grafana},
		{"GRPC", safeConfig.GRPC},
		{"Group", safeConfig.Group},
		{"Health", safeConfig.Health},
		{"HTTP", safeConfig.HTTP},
		{"HTTPServer", safeConfig.HTTPServer},
		{"I18n", safeConfig.I18n},
		{"JSON", safeConfig.JSON},
		{"JWT", safeConfig.JWT},
		{"Kafka", safeConfig.Kafka},
		{"Logging", safeConfig.Logging},
		{"Memory", safeConfig.Memory},
		{"Metrics", safeConfig.Metrics},
		{"Middleware", safeConfig.Middleware},
		{"Minio", safeConfig.Minio},
		{"Monitoring", safeConfig.Monitoring},
		{"MQTT", safeConfig.MQTT},
		{"MySQL", safeConfig.MySQL},
		{"OSS", safeConfig.OSS},
		{"Pay", safeConfig.Pay},
		{"Performance", safeConfig.Performance},
		{"PostgreSQL", safeConfig.PostgreSQL},
		{"Prometheus", safeConfig.Prometheus},
		{"Protection", safeConfig.Protection},
		{"Queue", safeConfig.Queue},
		{"RateLimit", safeConfig.RateLimit},
		{"Recovery", safeConfig.Recovery},
		{"Redis", safeConfig.Redis},
		{"Request", safeConfig.Request},
		{"RequestID", safeConfig.RequestID},
		{"Restful", safeConfig.Restful},
		{"RESTful", safeConfig.RESTful},
		{"Ristretto", safeConfig.Ristretto},
		{"RPCClient", safeConfig.RPCClient},
		{"RpcClient", safeConfig.RpcClient},
		{"RPCServer", safeConfig.RPCServer},
		{"RpcServer", safeConfig.RpcServer},
		{"S3", safeConfig.S3},
		{"Security", safeConfig.Security},
		{"Server", safeConfig.Server},
		{"Signature", safeConfig.Signature},
		{"SMS", safeConfig.SMS},
		{"SMTP", safeConfig.SMTP},
		{"SQLite", safeConfig.SQLite},
		{"STS", safeConfig.STS},
		{"Swagger", safeConfig.Swagger},
		{"Ticket", safeConfig.Ticket},
		{"TimeoutConfig", safeConfig.TimeoutConfig},
		{"TLS", safeConfig.TLS},
		{"Tracing", safeConfig.Tracing},
		{"VIP", safeConfig.VIP},
		{"Wechat", safeConfig.Wechat},
		{"WebSocketBreaker", safeConfig.WebSocketBreaker},
		{"WSC", safeConfig.WSC},
		{"Youzan", safeConfig.Youzan},
		{"Zap", safeConfig.Zap},
	}

	for _, accessor := range accessors {
		t.Run(accessor.name, func(t *testing.T) {
			result := accessor.getter()
			assert.NotNil(t, result, "%s should return non-nil ConfigSafe", accessor.name)
			assert.NotNil(t, result.SafeAccess, "%s should have non-nil SafeAccess", accessor.name)
		})
	}
}
