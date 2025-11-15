/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-15 13:53:37
 * @FilePath: \go-config\config_safe_access_test.go
 * @Description: 配置安全访问辅助工具测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
