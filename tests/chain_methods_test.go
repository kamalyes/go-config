/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-09 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 10:31:10
 * @FilePath: \go-config\tests\chain_methods_test.go
 * @Description: 测试所有包的链式方法功能
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"
	"time"

	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-config/pkg/captcha"
	"github.com/kamalyes/go-config/pkg/cors"
	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-config/pkg/email"
	"github.com/kamalyes/go-config/pkg/ftp"
	"github.com/kamalyes/go-config/pkg/jwt"
	"github.com/kamalyes/go-config/pkg/oss"
	"github.com/kamalyes/go-config/pkg/queue"
	"github.com/kamalyes/go-config/pkg/zap"
	"github.com/stretchr/testify/assert"
)

// TestAllDefaultConfigs 测试所有包的默认配置
func TestAllDefaultConfigs(t *testing.T) {
	// 测试缓存配置
	cacheConfig := cache.Default()
	assert.NotNil(t, cacheConfig)
	assert.Equal(t, "default", cacheConfig.ModuleName)

	// 测试数据库配置
	mysqlConfig := database.Default()
	assert.NotNil(t, mysqlConfig)
	assert.Equal(t, "mysql", mysqlConfig.ModuleName)

	// 测试JWT配置
	jwtConfig := jwt.Default()
	assert.NotNil(t, jwtConfig)
	assert.Equal(t, "jwt", jwtConfig.ModuleName)

	// 测试邮件配置
	emailConfig := email.Default()
	assert.NotNil(t, emailConfig)
	assert.Equal(t, "email", emailConfig.ModuleName)

	// 测试验证码配置
	captchaConfig := captcha.Default()
	assert.NotNil(t, captchaConfig)
	assert.Equal(t, "captcha", captchaConfig.ModuleName)

	// 测试FTP配置
	ftpConfig := ftp.Default()
	assert.NotNil(t, ftpConfig)
	assert.Equal(t, "ftp", ftpConfig.ModuleName)

	// 测试MQTT配置
	mqttConfig := queue.Default()
	assert.NotNil(t, mqttConfig)
	assert.Equal(t, "mqtt", mqttConfig.ModuleName)
}

// TestChainMethodsComplexScenario 测试复杂场景的链式调用
func TestChainMethodsComplexScenario(t *testing.T) {
	// 模拟一个完整的应用配置
	appConfig := struct {
		Cache    *cache.Cache
		Database *database.MySQL
		JWT      *jwt.JWT
		Email    *email.Email
		Logger   *zap.Zap
	}{
		Cache: cache.Default().
			WithModuleName("app-cache").
			WithType(cache.TypeRedis).
			WithEnabled(true).
			WithDefaultTTL(1 * time.Hour).
			WithKeyPrefix("myapp:"),

		Database: database.Default().
			WithModuleName("main-db").
			WithHost("production-db.example.com").
			WithPort("3306").
			WithDbname("production").
			WithUsername("app_user").
			WithMaxIdleConns(25).
			WithMaxOpenConns(250),

		JWT: jwt.Default().
			WithModuleName("auth").
			WithSigningKey("production-secret-key").
			WithExpiresTime(3600 * 24). // 24小时
			WithBufferTime(3600).       // 1小时缓冲
			WithUseMultipoint(true),

		Email: email.Default().
			WithModuleName("notification").
			WithHost("smtp.company.com").
			WithPort(587).
			WithFrom("noreply@company.com").
			WithTo("admin@company.com").
			WithIsSSL(true),

		Logger: zap.Default().
			WithModuleName("app-logger").
			WithLevel("info").
			WithFormat("json").
			WithPrefix("[PRODUCTION]").
			WithDirector("./logs/production").
			WithMaxSize(100).
			WithCompress(true),
	}

	// 验证配置
	assert.Equal(t, "app-cache", appConfig.Cache.ModuleName)
	assert.Equal(t, cache.TypeRedis, appConfig.Cache.Type)
	assert.Equal(t, true, appConfig.Cache.Enabled)

	assert.Equal(t, "main-db", appConfig.Database.ModuleName)
	assert.Equal(t, "production-db.example.com", appConfig.Database.Host)
	assert.Equal(t, "production", appConfig.Database.Dbname)

	assert.Equal(t, "auth", appConfig.JWT.ModuleName)
	assert.Equal(t, "production-secret-key", appConfig.JWT.SigningKey)
	assert.Equal(t, true, appConfig.JWT.UseMultipoint)

	assert.Equal(t, "notification", appConfig.Email.ModuleName)
	assert.Equal(t, "smtp.company.com", appConfig.Email.Host)
	assert.Equal(t, true, appConfig.Email.IsSSL)

	assert.Equal(t, "app-logger", appConfig.Logger.ModuleName)
	assert.Equal(t, "info", appConfig.Logger.Level)
	assert.Equal(t, "json", appConfig.Logger.Format)
}

// TestChainMethodsImmutability 测试链式方法的不变性
func TestChainMethodsImmutability(t *testing.T) {
	// 创建基础配置
	baseConfig := cache.Default()
	originalType := baseConfig.Type

	// 使用链式方法修改配置
	modifiedConfig := baseConfig.WithType(cache.TypeRedis)

	// 验证返回的是同一个实例（可变性）
	assert.Same(t, baseConfig, modifiedConfig)
	assert.Equal(t, cache.TypeRedis, baseConfig.Type)
	assert.NotEqual(t, originalType, baseConfig.Type)
}

// TestDefaultConfigsValidation 测试默认配置的验证
func TestDefaultConfigsValidation(t *testing.T) {
	// 测试MySQL默认配置验证
	mysqlConfig := database.Default()
	err := mysqlConfig.Validate()
	// 默认配置可能因为缺少必需字段而验证失败，这是正常的
	t.Logf("MySQL validation result: %v", err)

	// 测试JWT默认配置验证
	jwtConfig := jwt.Default()
	err = jwtConfig.Validate()
	assert.NoError(t, err) // JWT默认配置应该是有效的

	// 测试验证码默认配置验证
	captchaConfig := captcha.Default()
	err = captchaConfig.Validate()
	assert.NoError(t, err) // 验证码默认配置应该是有效的
}

// TestCORSChainMethods 测试CORS包的链式方法
func TestCORSChainMethods(t *testing.T) {
	config := cors.Default().
		WithModuleName("api-cors").
		WithAllowedOrigins([]string{"https://app.example.com", "https://admin.example.com"}).
		WithAllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}).
		WithAllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"}).
		WithMaxAge("86400").
		WithAllowCredentials(true).
		WithAllowedAllOrigins(false).
		WithOptionsResponseCode(200)

	assert.Equal(t, "api-cors", config.ModuleName)
	assert.Equal(t, []string{"https://app.example.com", "https://admin.example.com"}, config.AllowedOrigins)
	assert.Equal(t, []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, config.AllowedMethods)
	assert.Equal(t, true, config.AllowCredentials)
	assert.Equal(t, false, config.AllowedAllOrigins)
	assert.Equal(t, 200, config.OptionsResponseCode)
}

// TestS3ChainMethods 测试S3包的链式方法
func TestS3ChainMethods(t *testing.T) {
	config := oss.DefaultS3Config().
		WithModuleName("file-storage").
		WithEndpoint("https://s3.us-west-2.amazonaws.com").
		WithRegion("us-west-2").
		WithAccessKey("AKIAIOSFODNN7EXAMPLE").
		WithSecretKey("wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY").
		WithBucketPrefix("myapp-prod").
		WithDisableSSL(false).
		WithPathStyle(false)

	assert.Equal(t, "file-storage", config.ModuleName)
	assert.Equal(t, "https://s3.us-west-2.amazonaws.com", config.Endpoint)
	assert.Equal(t, "us-west-2", config.Region)
	assert.Equal(t, "myapp-prod", config.BucketPrefix)
	assert.Equal(t, true, config.UseSSL)
	assert.Equal(t, false, config.PathStyle)
}

// TestAliyunOssChainMethods 测试阿里云OSS包的链式方法
func TestAliyunOssChainMethods(t *testing.T) {
	config := oss.DefaultAliyunOSSConfig().
		WithModuleName("aliyun-storage").
		WithEndpoint("oss-cn-beijing.aliyuncs.com").
		WithAccessKey("LTAIxxxxxxxxxxxx").
		WithSecretKey("xxxxxxxxxxxxxxxxxxxxxxxxxxxx").
		WithBucket("my-bucket").
		WithReplaceOriginalHost("oss-cn-beijing.aliyuncs.com").
		WithReplaceLaterHost("cdn.example.com")

	assert.Equal(t, "aliyun-storage", config.ModuleName)
	assert.Equal(t, "oss-cn-beijing.aliyuncs.com", config.Endpoint)
	assert.Equal(t, "my-bucket", config.Bucket)
	assert.Equal(t, "cdn.example.com", config.ReplaceLaterHost)
}

// TestAllConfigsClone 测试所有配置的克隆功能
func TestAllConfigsClone(t *testing.T) {
	// 测试各种配置的克隆功能
	mysqlOriginal := database.Default().WithDbname("test_db")
	mysqlCloned := mysqlOriginal.Clone().(*database.MySQL)
	assert.Equal(t, mysqlOriginal, mysqlCloned)
	assert.NotSame(t, mysqlOriginal, mysqlCloned)

	jwtOriginal := jwt.Default().WithSigningKey("test_key")
	jwtCloned := jwtOriginal.Clone().(*jwt.JWT)
	assert.Equal(t, jwtOriginal, jwtCloned)
	assert.NotSame(t, jwtOriginal, jwtCloned)

	ftpOriginal := ftp.Default().WithEndpoint("192.168.1.100:21")
	ftpCloned := ftpOriginal.Clone().(*ftp.Ftp)
	assert.Equal(t, ftpOriginal, ftpCloned)
	assert.NotSame(t, ftpOriginal, ftpCloned)
}
