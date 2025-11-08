/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-08 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 02:02:15
 * @FilePath: \go-config\tests\default_methods_test.go
 * @Description: 测试所有包的Default方法
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-config/pkg/captcha"
	"github.com/kamalyes/go-config/pkg/cors"
	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-config/pkg/elk"
	"github.com/kamalyes/go-config/pkg/email"
	"github.com/kamalyes/go-config/pkg/env"
	"github.com/kamalyes/go-config/pkg/ftp"
	"github.com/kamalyes/go-config/pkg/jwt"
	"github.com/kamalyes/go-config/pkg/oss"
	"github.com/kamalyes/go-config/pkg/pay"
	"github.com/kamalyes/go-config/pkg/queue"
	"github.com/kamalyes/go-config/pkg/register"
	"github.com/kamalyes/go-config/pkg/sms"
	"github.com/kamalyes/go-config/pkg/sts"
	"github.com/kamalyes/go-config/pkg/youzan"
	"github.com/kamalyes/go-config/pkg/zap"
	"github.com/kamalyes/go-config/pkg/zero"
)

// TestAllDefaultMethods 测试所有包的Default方法
func TestAllDefaultMethods(t *testing.T) {
	testCases := []struct {
		name string
		test func() interface{}
	}{
		{"Cache", func() interface{} { return cache.Default() }},
		{"Captcha", func() interface{} { return captcha.Default() }},
		{"Cors", func() interface{} { return cors.Default() }},
		{"Email", func() interface{} { return email.Default() }},
		{"Environment", func() interface{} { return env.Default() }},
		{"Ftp", func() interface{} { return ftp.Default() }},
		{"JWT", func() interface{} { return jwt.Default() }},
		{"Mqtt", func() interface{} { return queue.Default() }},
		{"Zap", func() interface{} { return zap.Default() }},
		
		// Database
		{"MySQL", func() interface{} { return database.DefaultMySQL() }},
		{"PostgreSQL", func() interface{} { return database.DefaultPostgreSQL() }},
		{"SQLite", func() interface{} { return database.DefaultSQLite() }},
		
		// ELK
		{"Elasticsearch", func() interface{} { return elk.DefaultElasticsearch() }},
		{"Kafka", func() interface{} { return elk.DefaultKafka() }},
		
		// OSS
		{"AliyunOss", func() interface{} { return oss.DefaultAliyunOss() }},
		{"S3", func() interface{} { return oss.DefaultS3() }},
		{"Minio", func() interface{} { return oss.DefaultMinio() }},
		
		// Pay
		{"AliPay", func() interface{} { return pay.DefaultAliPay() }},
		{"WechatPay", func() interface{} { return pay.DefaultWechatPay() }},
		
		// Register
		{"Server", func() interface{} { return register.DefaultServer() }},
		{"Consul", func() interface{} { return register.DefaultConsul() }},
		{"Jaeger", func() interface{} { return register.DefaultJaeger() }},
		{"PProf", func() interface{} { return register.DefaultPProf() }},
		
		// SMS
		{"AliyunSms", func() interface{} { return sms.DefaultAliyunSms() }},
		
		// STS
		{"AliyunSts", func() interface{} { return sts.Default() }},
		
		// YouZan
		{"YouZan", func() interface{} { return youzan.Default() }},
		
		// Zero
		{"RpcClient", func() interface{} { return zero.DefaultRpcClient() }},
		{"RpcServer", func() interface{} { return zero.DefaultRpcServer() }},
		{"Etcd", func() interface{} { return zero.DefaultEtcd() }},
		{"LogConf", func() interface{} { return zero.DefaultLogConf() }},
		{"Prometheus", func() interface{} { return zero.DefaultPrometheus() }},
		{"Telemetry", func() interface{} { return zero.DefaultTelemetry() }},
		{"Restful", func() interface{} { return zero.DefaultRestful() }},
		{"Signature", func() interface{} { return zero.DefaultSignature() }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.test()
			if result == nil {
				t.Errorf("Default method for %s returned nil", tc.name)
			}
		})
	}
}

// TestAllDefaultChainMethods 测试所有包的链式Default方法
func TestAllDefaultChainMethods(t *testing.T) {
	testCases := []struct {
		name string
		test func() interface{}
	}{
		{"CacheChain", func() interface{} { return cache.DefaultCache() }},
		{"CaptchaChain", func() interface{} { return captcha.DefaultCaptcha() }},
		{"CorsChain", func() interface{} { return cors.DefaultCors() }},
		{"EmailChain", func() interface{} { return email.DefaultEmail() }},
		{"EnvironmentChain", func() interface{} { return env.DefaultEnvironment() }},
		{"FtpChain", func() interface{} { return ftp.DefaultFtp() }},
		{"JWTChain", func() interface{} { return jwt.DefaultJWT() }},
		{"MqttChain", func() interface{} { return queue.DefaultMqtt() }},
		{"ZapChain", func() interface{} { return zap.DefaultZap() }},
		
		// Database
		{"MySQLChain", func() interface{} { return database.Default() }},
		{"PostgreSQLChain", func() interface{} { return database.DefaultPostgreSQLConfig() }},
		{"SQLiteChain", func() interface{} { return database.DefaultSQLiteConfig() }},
		
		// ELK
		{"ElasticsearchChain", func() interface{} { return elk.DefaultElasticsearch() }},
		{"KafkaChain", func() interface{} { return elk.DefaultKafka() }},
		
		// OSS
		{"AliyunOssChain", func() interface{} { return oss.DefaultAliyunOSSConfig() }},
		{"S3Chain", func() interface{} { return oss.DefaultS3Config() }},
		{"MinioChain", func() interface{} { return oss.DefaultMinioConfig() }},
		
		// Pay
		{"AliPayChain", func() interface{} { return pay.DefaultAliPayConfig() }},
		{"WechatPayChain", func() interface{} { return pay.DefaultWechatPayConfig() }},
		
		// Register
		{"ServerChain", func() interface{} { return register.Default() }},
		{"ConsulChain", func() interface{} { return register.DefaultConsulConfig() }},
		{"JaegerChain", func() interface{} { return register.DefaultJaegerConfig() }},
		{"PProfChain", func() interface{} { return register.DefaultPProfConfig() }},
		
		// SMS
		{"AliyunSmsChain", func() interface{} { return sms.Default() }},
		
		// STS
		{"AliyunStsChain", func() interface{} { return sts.DefaultAliyunSts() }},
		
		// YouZan
		{"YouZanChain", func() interface{} { return youzan.DefaultYouZan() }},
		
		// Zero
		{"RpcClientChain", func() interface{} { return zero.DefaultRpcClientConfig() }},
		{"RpcServerChain", func() interface{} { return zero.DefaultRpcServerConfig() }},
		{"EtcdChain", func() interface{} { return zero.DefaultEtcdConfig() }},
		{"LogConfChain", func() interface{} { return zero.DefaultLogConfConfig() }},
		{"PrometheusChain", func() interface{} { return zero.DefaultPrometheusConfig() }},
		{"TelemetryChain", func() interface{} { return zero.DefaultTelemetryConfig() }},
		{"RestfulChain", func() interface{} { return zero.DefaultRestfulConfig() }},
		{"SignatureChain", func() interface{} { return zero.DefaultSignatureConfig() }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.test()
			if result == nil {
				t.Errorf("Default chain method for %s returned nil", tc.name)
			}
		})
	}
}

// TestChainMethodUsage 测试链式方法的使用
func TestChainMethodUsage(t *testing.T) {
	// 测试缓存链式配置
	cacheConfig := cache.Default().
		WithType("redis").
		WithEnabled(true)
	
	if cacheConfig.Type != "redis" || !cacheConfig.Enabled {
		t.Error("Cache chain methods not working correctly")
	}
	
	// 测试JWT链式配置
	jwtConfig := jwt.Default().
		WithSigningKey("test-key").
		WithExpiresTime(3600)
	
	if jwtConfig.SigningKey != "test-key" || jwtConfig.ExpiresTime != 3600 {
		t.Error("JWT chain methods not working correctly")
	}
	
	// 测试MySQL链式配置
	mysqlConfig := database.Default().
		WithHost("localhost").
		WithPort("3306").
		WithDbname("test")
	
	if mysqlConfig.Host != "localhost" || mysqlConfig.Port != "3306" || mysqlConfig.Dbname != "test" {
		t.Error("MySQL chain methods not working correctly")
	}
}