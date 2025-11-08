/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-09 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 00:00:00
 * @FilePath: \go-config\tests\benchmark_test.go
 * @Description: 链式方法性能基准测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"testing"
	"time"

	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-config/pkg/jwt"
)

// BenchmarkCacheChainMethods 缓存链式方法基准测试
func BenchmarkCacheChainMethods(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config := cache.Default().
			WithModuleName("benchmark").
			WithType(cache.TypeRedis).
			WithEnabled(true).
			WithDefaultTTL(30 * time.Minute).
			WithKeyPrefix("bench:")
		_ = config
	}
}

// BenchmarkCacheTraditionalMethods 缓存传统方法基准测试
func BenchmarkCacheTraditionalMethods(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config := &cache.Cache{
			ModuleName: "benchmark",
			Type:       cache.TypeRedis,
			Enabled:    true,
			DefaultTTL: 30 * time.Minute,
			KeyPrefix:  "bench:",
		}
		_ = config
	}
}

// BenchmarkMySQLChainMethods MySQL链式方法基准测试
func BenchmarkMySQLChainMethods(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config := database.Default().
			WithModuleName("benchmark").
			WithHost("localhost").
			WithPort("3306").
			WithDbname("test").
			WithUsername("user").
			WithPassword("pass").
			WithMaxIdleConns(10).
			WithMaxOpenConns(100)
		_ = config
	}
}

// BenchmarkMySQLTraditionalMethods MySQL传统方法基准测试
func BenchmarkMySQLTraditionalMethods(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config := &database.MySQL{
			ModuleName:   "benchmark",
			Host:         "localhost",
			Port:         "3306",
			Dbname:       "test",
			Username:     "user",
			Password:     "pass",
			MaxIdleConns: 10,
			MaxOpenConns: 100,
		}
		_ = config
	}
}

// BenchmarkJWTChainMethods JWT链式方法基准测试
func BenchmarkJWTChainMethods(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config := jwt.Default().
			WithModuleName("benchmark").
			WithSigningKey("secret").
			WithExpiresTime(3600).
			WithBufferTime(300).
			WithUseMultipoint(false)
		_ = config
	}
}

// BenchmarkJWTTraditionalMethods JWT传统方法基准测试
func BenchmarkJWTTraditionalMethods(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config := &jwt.JWT{
			ModuleName:    "benchmark",
			SigningKey:    "secret",
			ExpiresTime:   3600,
			BufferTime:    300,
			UseMultipoint: false,
		}
		_ = config
	}
}

// BenchmarkConfigCreationMemoryUsage 内存使用基准测试
func BenchmarkConfigCreationMemoryUsage(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// 创建多个配置实例
		cacheConfig := cache.Default().WithModuleName("mem-test")
		dbConfig := database.Default().WithHost("localhost")
		jwtConfig := jwt.Default().WithSigningKey("test-key")
		
		// 防止编译器优化
		_ = cacheConfig
		_ = dbConfig
		_ = jwtConfig
	}
}