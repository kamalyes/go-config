package goconfig

import (
	"fmt"
	"github.com/kamalyes/go-config/pkg/health"
	"github.com/kamalyes/go-config/pkg/jwt"
	"testing"
)

func TestDebugHealthStructAccess(t *testing.T) {
	// 创建Health配置
	healthConfig := health.Default()
	healthConfig.Enabled = true
	healthConfig.Redis.Enabled = true
	healthConfig.Redis.Timeout = 30 // 秒
	healthConfig.MySQL.Enabled = false
	healthConfig.MySQL.Timeout = 10 // 秒

	// 创建SafeConfig
	safeConfig := SafeConfig(healthConfig)

	// 调试输出
	fmt.Printf("原始配置: %+v\n", healthConfig)
	fmt.Printf("Redis配置: %+v\n", healthConfig.Redis)
	fmt.Printf("SafeAccess值: %+v\n", safeConfig.SafeAccess.Value())
	fmt.Printf("Health: %+v\n", safeConfig.Health().Value())
	fmt.Printf("Redis: %+v\n", safeConfig.Health().Redis().Value())
	fmt.Printf("Redis.Timeout: %+v\n", safeConfig.Health().Redis().Field("Timeout").Value())

	// 测试Int方法
	timeout := safeConfig.Health().Redis().Field("Timeout").Int(0)
	fmt.Printf("Redis.Timeout.Int(): %d\n", timeout)

	// 测试完整方法
	result := safeConfig.GetRedisHealthTimeout(5)
	fmt.Printf("GetRedisHealthTimeout(5s): %v\n", result)
}

func TestDebugStructAccess(t *testing.T) {
	// 创建JWT配置
	jwtConfig := jwt.Default()
	jwtConfig.SigningKey = "test-secret-key"
	jwtConfig.ExpiresTime = 86400

	// 创建SafeConfig
	safeConfig := SafeConfig(jwtConfig)

	// 调试输出
	fmt.Printf("原始配置: %+v\n", jwtConfig)
	fmt.Printf("SafeAccess值: %+v\n", safeConfig.SafeAccess.Value())
	fmt.Printf("SigningKey直接访问: %+v\n", safeConfig.Field("SigningKey").Value())
	fmt.Printf("ExpiresTime直接访问: %+v\n", safeConfig.Field("ExpiresTime").Value())

	// 测试String方法
	signingKey := safeConfig.Field("SigningKey").String("")
	fmt.Printf("SigningKey.String(): %s\n", signingKey)

	// 测试Int方法
	expiresTime := safeConfig.Field("ExpiresTime").Int(0)
	fmt.Printf("ExpiresTime.Int(): %d\n", expiresTime)

	// 测试Int64方法
	expiresTime64 := safeConfig.Field("ExpiresTime").Int64(0)
	fmt.Printf("ExpiresTime.Int64(): %d\n", expiresTime64)

	// 直接访问GetJWTSecret
	secret := safeConfig.GetJWTSecret("")
	fmt.Printf("GetJWTSecret: %s\n", secret)

	expiration := safeConfig.GetJWTExpiration(0)
	fmt.Printf("GetJWTExpiration: %v\n", expiration)
}
