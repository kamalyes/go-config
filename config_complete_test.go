/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-26 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-11 15:33:26
 * @FilePath: \go-config\config_complete_test.go
 * @Description: 配置生成器与热加载集成测试 - 完整验证所有字段包括嵌套结构
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"context"
	"testing"
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/kamalyes/go-config/pkg/banner"
	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-config/pkg/gateway"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConfigCompleteValidation 完整验证所有字段
func TestConfigCompleteValidation(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()

	// 创建配置生成器
	generator := NewSmartConfigGenerator(tmpDir)
	require.NotNil(t, generator)

	// 测试关键模块
	testModules := []string{"access", "banner", "cache", "gateway"}
	err := generator.EnableOnlyModules(testModules...)
	require.NoError(t, err)

	// 生成所有配置文件
	err = generator.GenerateAllConfigs()
	require.NoError(t, err)
	// 测试 Banner 模块 - 完整字段验证
	t.Run("Banner Complete", func(t *testing.T) {
		testCompleteModuleConfig(t, tmpDir, "banner", &banner.Banner{}, func(cfg interface{}) {
			config := cfg.(*banner.Banner)
			defaultCfg := banner.Default()

			// 验证所有字段
			assert.Equal(t, defaultCfg.ModuleName, config.ModuleName, "ModuleName should match")
			assert.Equal(t, defaultCfg.Enabled, config.Enabled, "Enabled should match")
			assert.Equal(t, defaultCfg.Template, config.Template, "Template should match")
			assert.Equal(t, defaultCfg.Title, config.Title, "Title should match")
			assert.Equal(t, defaultCfg.Description, config.Description, "Description should match")
			assert.Equal(t, defaultCfg.Author, config.Author, "Author should match")
			assert.Equal(t, defaultCfg.Email, config.Email, "Email should match")
			assert.Equal(t, defaultCfg.Website, config.Website, "Website should match")

			t.Logf("✅ Banner 模块所有字段验证通过")
		})
	})

	// 测试 Cache 模块 - 完整字段验证（包括嵌套结构）
	t.Run("Cache Complete", func(t *testing.T) {
		testCompleteModuleConfig(t, tmpDir, "cache", &cache.Cache{}, func(cfg interface{}) {
			config := cfg.(*cache.Cache)
			defaultCfg := cache.Default()

			// 验证顶层字段
			assert.Equal(t, defaultCfg.ModuleName, config.ModuleName, "ModuleName should match")
			assert.Equal(t, defaultCfg.Type, config.Type, "Type should match")
			assert.Equal(t, defaultCfg.Enabled, config.Enabled, "Enabled should match")
			assert.Equal(t, defaultCfg.DefaultTTL, config.DefaultTTL, "DefaultTTL should match")
			assert.Equal(t, defaultCfg.KeyPrefix, config.KeyPrefix, "KeyPrefix should match")
			assert.Equal(t, defaultCfg.Serializer, config.Serializer, "Serializer should match")

			// 验证 Memory 嵌套结构
			assert.Equal(t, defaultCfg.Memory.ModuleName, config.Memory.ModuleName, "Memory.ModuleName should match")
			assert.Equal(t, defaultCfg.Memory.Capacity, config.Memory.Capacity, "Memory.Capacity should match")
			assert.Equal(t, defaultCfg.Memory.DefaultTTL, config.Memory.DefaultTTL, "Memory.DefaultTTL should match")
			assert.Equal(t, defaultCfg.Memory.CleanupSize, config.Memory.CleanupSize, "Memory.CleanupSize should match")
			assert.Equal(t, defaultCfg.Memory.MaxSize, config.Memory.MaxSize, "Memory.MaxSize should match")

			// 验证 Ristretto 嵌套结构
			assert.Equal(t, defaultCfg.Ristretto.NumCounters, config.Ristretto.NumCounters, "Ristretto.NumCounters should match")
			assert.Equal(t, defaultCfg.Ristretto.MaxCost, config.Ristretto.MaxCost, "Ristretto.MaxCost should match")
			assert.Equal(t, defaultCfg.Ristretto.BufferItems, config.Ristretto.BufferItems, "Ristretto.BufferItems should match")

			// 验证 Redis 嵌套结构
			assert.Equal(t, defaultCfg.Redis.ModuleName, config.Redis.ModuleName, "Redis.ModuleName should match")
			assert.Equal(t, defaultCfg.Redis.Addr, config.Redis.Addr, "Redis.Addr should match")
			assert.Equal(t, defaultCfg.Redis.Addrs, config.Redis.Addrs, "Redis.Addrs should match")
			assert.Equal(t, defaultCfg.Redis.Username, config.Redis.Username, "Redis.Username should match")
			assert.Equal(t, defaultCfg.Redis.Password, config.Redis.Password, "Redis.Password should match")
			assert.Equal(t, defaultCfg.Redis.DB, config.Redis.DB, "Redis.DB should match")
			assert.Equal(t, defaultCfg.Redis.MaxRetries, config.Redis.MaxRetries, "Redis.MaxRetries should match")
			assert.Equal(t, defaultCfg.Redis.PoolSize, config.Redis.PoolSize, "Redis.PoolSize should match")
			assert.Equal(t, defaultCfg.Redis.MinIdleConns, config.Redis.MinIdleConns, "Redis.MinIdleConns should match")
			assert.Equal(t, defaultCfg.Redis.MaxConnAge, config.Redis.MaxConnAge, "Redis.MaxConnAge should match")
			assert.Equal(t, defaultCfg.Redis.PoolTimeout, config.Redis.PoolTimeout, "Redis.PoolTimeout should match")
			assert.Equal(t, defaultCfg.Redis.IdleTimeout, config.Redis.IdleTimeout, "Redis.IdleTimeout should match")
			assert.Equal(t, defaultCfg.Redis.ReadTimeout, config.Redis.ReadTimeout, "Redis.ReadTimeout should match")
			assert.Equal(t, defaultCfg.Redis.WriteTimeout, config.Redis.WriteTimeout, "Redis.WriteTimeout should match")
			assert.Equal(t, defaultCfg.Redis.ClusterMode, config.Redis.ClusterMode, "Redis.ClusterMode should match")

			// 验证 Sharded 嵌套结构
			assert.Equal(t, defaultCfg.Sharded.ModuleName, config.Sharded.ModuleName, "Sharded.ModuleName should match")
			assert.Equal(t, defaultCfg.Sharded.ShardCount, config.Sharded.ShardCount, "Sharded.ShardCount should match")
			assert.Equal(t, defaultCfg.Sharded.BaseType, config.Sharded.BaseType, "Sharded.BaseType should match")
			assert.Equal(t, defaultCfg.Sharded.HashFunc, config.Sharded.HashFunc, "Sharded.HashFunc should match")
			assert.Equal(t, defaultCfg.Sharded.LoadBalancer, config.Sharded.LoadBalancer, "Sharded.LoadBalancer should match")

			// 验证 TwoLevel 嵌套结构
			assert.Equal(t, defaultCfg.TwoLevel.L1Type, config.TwoLevel.L1Type, "TwoLevel.L1Type should match")
			assert.Equal(t, defaultCfg.TwoLevel.L2Type, config.TwoLevel.L2Type, "TwoLevel.L2Type should match")
			assert.Equal(t, defaultCfg.TwoLevel.L1TTL, config.TwoLevel.L1TTL, "TwoLevel.L1TTL should match")
			assert.Equal(t, defaultCfg.TwoLevel.L2TTL, config.TwoLevel.L2TTL, "TwoLevel.L2TTL should match")

			// 验证 Expiring 嵌套结构
			assert.Equal(t, defaultCfg.Expiring.ModuleName, config.Expiring.ModuleName, "Expiring.ModuleName should match")
			assert.Equal(t, defaultCfg.Expiring.CleanupInterval, config.Expiring.CleanupInterval, "Expiring.CleanupInterval should match")
			assert.Equal(t, defaultCfg.Expiring.DefaultTTL, config.Expiring.DefaultTTL, "Expiring.DefaultTTL should match")
			assert.Equal(t, defaultCfg.Expiring.MaxSize, config.Expiring.MaxSize, "Expiring.MaxSize should match")
			assert.Equal(t, defaultCfg.Expiring.EvictionPolicy, config.Expiring.EvictionPolicy, "Expiring.EvictionPolicy should match")
			assert.Equal(t, defaultCfg.Expiring.EnableLazyExpiry, config.Expiring.EnableLazyExpiry, "Expiring.EnableLazyExpiry should match")
			assert.Equal(t, defaultCfg.Expiring.MaxMemoryUsage, config.Expiring.MaxMemoryUsage, "Expiring.MaxMemoryUsage should match")

			t.Logf("✅ Cache 模块所有字段（包括嵌套结构）验证通过")
		})
	})

	// 测试 Gateway 模块 - 完整字段验证（包括多层嵌套结构）
	t.Run("Gateway Complete", func(t *testing.T) {
		testCompleteModuleConfig(t, tmpDir, "gateway", &gateway.Gateway{}, func(cfg interface{}) {
			config := cfg.(*gateway.Gateway)
			defaultCfg := gateway.Default()

			// 验证顶层字段
			assert.Equal(t, defaultCfg.ModuleName, config.ModuleName, "ModuleName should match")
			assert.Equal(t, defaultCfg.Name, config.Name, "Name should match")
			assert.Equal(t, defaultCfg.Enabled, config.Enabled, "Enabled should match")
			assert.Equal(t, defaultCfg.Debug, config.Debug, "Debug should match")
			assert.Equal(t, defaultCfg.Version, config.Version, "Version should match")
			assert.Equal(t, defaultCfg.Environment, config.Environment, "Environment should match")
			assert.Equal(t, defaultCfg.BuildTime, config.BuildTime, "BuildTime should match")
			assert.Equal(t, defaultCfg.BuildUser, config.BuildUser, "BuildUser should match")
			assert.Equal(t, defaultCfg.GoVersion, config.GoVersion, "GoVersion should match")
			assert.Equal(t, defaultCfg.GitCommit, config.GitCommit, "GitCommit should match")
			assert.Equal(t, defaultCfg.GitBranch, config.GitBranch, "GitBranch should match")
			assert.Equal(t, defaultCfg.GitTag, config.GitTag, "GitTag should match")

			// 验证 JSON 嵌套结构
			assert.Equal(t, defaultCfg.JSON.UseProtoNames, config.JSON.UseProtoNames, "JSON.UseProtoNames should match")
			assert.Equal(t, defaultCfg.JSON.EmitUnpopulated, config.JSON.EmitUnpopulated, "JSON.EmitUnpopulated should match")
			assert.Equal(t, defaultCfg.JSON.DiscardUnknown, config.JSON.DiscardUnknown, "JSON.DiscardUnknown should match")

			// 验证 HTTPServer 嵌套结构
			assert.Equal(t, defaultCfg.HTTPServer.ModuleName, config.HTTPServer.ModuleName, "HTTPServer.ModuleName should match")
			assert.Equal(t, defaultCfg.HTTPServer.Host, config.HTTPServer.Host, "HTTPServer.Host should match")
			assert.Equal(t, defaultCfg.HTTPServer.Port, config.HTTPServer.Port, "HTTPServer.Port should match")
			assert.Equal(t, defaultCfg.HTTPServer.ReadTimeout, config.HTTPServer.ReadTimeout, "HTTPServer.ReadTimeout should match")
			assert.Equal(t, defaultCfg.HTTPServer.WriteTimeout, config.HTTPServer.WriteTimeout, "HTTPServer.WriteTimeout should match")
			assert.Equal(t, defaultCfg.HTTPServer.IdleTimeout, config.HTTPServer.IdleTimeout, "HTTPServer.IdleTimeout should match")
			assert.Equal(t, defaultCfg.HTTPServer.MaxHeaderBytes, config.HTTPServer.MaxHeaderBytes, "HTTPServer.MaxHeaderBytes should match")
			assert.Equal(t, defaultCfg.HTTPServer.EnableTls, config.HTTPServer.EnableTls, "HTTPServer.EnableTls should match")
			assert.Equal(t, defaultCfg.HTTPServer.EnableGzipCompress, config.HTTPServer.EnableGzipCompress, "HTTPServer.EnableGzipCompress should match")
			assert.Equal(t, defaultCfg.HTTPServer.Headers, config.HTTPServer.Headers, "HTTPServer.Headers should match")

			// 验证 TLS 子嵌套结构
			assert.Equal(t, defaultCfg.HTTPServer.TLS.CertFile, config.HTTPServer.TLS.CertFile, "HTTPServer.TLS.CertFile should match")
			assert.Equal(t, defaultCfg.HTTPServer.TLS.KeyFile, config.HTTPServer.TLS.KeyFile, "HTTPServer.TLS.KeyFile should match")
			assert.Equal(t, defaultCfg.HTTPServer.TLS.CAFile, config.HTTPServer.TLS.CAFile, "HTTPServer.TLS.CAFile should match")

			// 验证 GRPC 嵌套结构
			assert.Equal(t, defaultCfg.GRPC.Server.Host, config.GRPC.Server.Host, "GRPC.Server.Host should match")
			assert.Equal(t, defaultCfg.GRPC.Server.Port, config.GRPC.Server.Port, "GRPC.Server.Port should match")
			assert.Equal(t, defaultCfg.GRPC.Server.Network, config.GRPC.Server.Network, "GRPC.Server.Network should match")
			assert.Equal(t, defaultCfg.GRPC.Server.MaxRecvMsgSize, config.GRPC.Server.MaxRecvMsgSize, "GRPC.Server.MaxRecvMsgSize should match")
			assert.Equal(t, defaultCfg.GRPC.Server.MaxSendMsgSize, config.GRPC.Server.MaxSendMsgSize, "GRPC.Server.MaxSendMsgSize should match")
			assert.Equal(t, defaultCfg.GRPC.Server.KeepaliveTime, config.GRPC.Server.KeepaliveTime, "GRPC.Server.KeepaliveTime should match")
			assert.Equal(t, defaultCfg.GRPC.Server.KeepaliveTimeout, config.GRPC.Server.KeepaliveTimeout, "GRPC.Server.KeepaliveTimeout should match")
			assert.Equal(t, defaultCfg.GRPC.Server.ConnectionTimeout, config.GRPC.Server.ConnectionTimeout, "GRPC.Server.ConnectionTimeout should match")
			assert.Equal(t, defaultCfg.GRPC.Server.EnableReflection, config.GRPC.Server.EnableReflection, "GRPC.Server.EnableReflection should match")

			// 验证 GRPC Clients
			assert.Equal(t, len(defaultCfg.GRPC.Clients), len(config.GRPC.Clients), "GRPC.Clients length should match")

			// 验证 Cache 嵌套结构（Gateway中的Cache）
			assert.Equal(t, defaultCfg.Cache.ModuleName, config.Cache.ModuleName, "Cache.ModuleName should match")
			assert.Equal(t, defaultCfg.Cache.Type, config.Cache.Type, "Cache.Type should match")
			assert.Equal(t, defaultCfg.Cache.Enabled, config.Cache.Enabled, "Cache.Enabled should match")
			assert.Equal(t, defaultCfg.Cache.DefaultTTL, config.Cache.DefaultTTL, "Cache.DefaultTTL should match")
			assert.Equal(t, defaultCfg.Cache.KeyPrefix, config.Cache.KeyPrefix, "Cache.KeyPrefix should match")
			assert.Equal(t, defaultCfg.Cache.Serializer, config.Cache.Serializer, "Cache.Serializer should match")

			// 验证 Cache.Memory 子结构
			assert.Equal(t, defaultCfg.Cache.Memory.ModuleName, config.Cache.Memory.ModuleName, "Cache.Memory.ModuleName should match")
			assert.Equal(t, defaultCfg.Cache.Memory.Capacity, config.Cache.Memory.Capacity, "Cache.Memory.Capacity should match")

			// 验证 Database 嵌套结构
			assert.Equal(t, defaultCfg.Database.Type, config.Database.Type, "Database.Type should match")
			assert.Equal(t, defaultCfg.Database.Enabled, config.Database.Enabled, "Database.Enabled should match")
			assert.Equal(t, defaultCfg.Database.Default, config.Database.Default, "Database.Default should match")

			// 验证 Database.MySQL 子结构
			assert.Equal(t, defaultCfg.Database.MySQL.ModuleName, config.Database.MySQL.ModuleName, "Database.MySQL.ModuleName should match")
			assert.Equal(t, defaultCfg.Database.MySQL.Host, config.Database.MySQL.Host, "Database.MySQL.Host should match")
			assert.Equal(t, defaultCfg.Database.MySQL.Port, config.Database.MySQL.Port, "Database.MySQL.Port should match")
			assert.Equal(t, defaultCfg.Database.MySQL.Dbname, config.Database.MySQL.Dbname, "Database.MySQL.Database should match")
			assert.Equal(t, defaultCfg.Database.MySQL.Username, config.Database.MySQL.Username, "Database.MySQL.Username should match")

			// 验证 Database.PostgreSQL 子结构
			assert.Equal(t, defaultCfg.Database.PostgreSQL.ModuleName, config.Database.PostgreSQL.ModuleName, "Database.PostgreSQL.ModuleName should match")
			assert.Equal(t, defaultCfg.Database.PostgreSQL.Host, config.Database.PostgreSQL.Host, "Database.PostgreSQL.Host should match")
			assert.Equal(t, defaultCfg.Database.PostgreSQL.Port, config.Database.PostgreSQL.Port, "Database.PostgreSQL.Port should match")

			// 验证 Database.SQLite 子结构
			assert.Equal(t, defaultCfg.Database.SQLite.ModuleName, config.Database.SQLite.ModuleName, "Database.SQLite.ModuleName should match")
			assert.Equal(t, defaultCfg.Database.SQLite.DbPath, config.Database.SQLite.DbPath, "Database.SQLite.DbPath should match")

			// 验证 Etcd 嵌套结构
			assert.Equal(t, defaultCfg.Etcd.ModuleName, config.Etcd.ModuleName, "Etcd.ModuleName should match")
			assert.Equal(t, defaultCfg.Etcd.Hosts, config.Etcd.Hosts, "Etcd.Hosts should match")
			assert.Equal(t, defaultCfg.Etcd.Key, config.Etcd.Key, "Etcd.Key should match")
			assert.Equal(t, defaultCfg.Etcd.ID, config.Etcd.ID, "Etcd.ID should match")
			assert.Equal(t, defaultCfg.Etcd.User, config.Etcd.User, "Etcd.User should match")
			assert.Equal(t, defaultCfg.Etcd.Pass, config.Etcd.Pass, "Etcd.Pass should match")
			assert.Equal(t, defaultCfg.Etcd.CertFile, config.Etcd.CertFile, "Etcd.CertFile should match")
			assert.Equal(t, defaultCfg.Etcd.CertKeyFile, config.Etcd.CertKeyFile, "Etcd.CertKeyFile should match")
			assert.Equal(t, defaultCfg.Etcd.CACertFile, config.Etcd.CACertFile, "Etcd.CACertFile should match")

			// 验证 Kafka 嵌套结构
			assert.Equal(t, defaultCfg.Kafka.ModuleName, config.Kafka.ModuleName, "Kafka.ModuleName should match")
			assert.Equal(t, defaultCfg.Kafka.Brokers, config.Kafka.Brokers, "Kafka.Brokers should match")
			assert.Equal(t, defaultCfg.Kafka.Topic, config.Kafka.Topic, "Kafka.Topic should match")
			assert.Equal(t, defaultCfg.Kafka.GroupID, config.Kafka.GroupID, "Kafka.GroupID should match")
			assert.Equal(t, defaultCfg.Kafka.Username, config.Kafka.Username, "Kafka.Username should match")
			assert.Equal(t, defaultCfg.Kafka.Partition, config.Kafka.Partition, "Kafka.Partition should match")
			assert.Equal(t, defaultCfg.Kafka.Offset, config.Kafka.Offset, "Kafka.Offset should match")
			assert.Equal(t, defaultCfg.Kafka.TryTimes, config.Kafka.TryTimes, "Kafka.TryTimes should match")

			// 验证 OSS 嵌套结构
			assert.Equal(t, defaultCfg.OSS.Type, config.OSS.Type, "OSS.Type should match")
			assert.Equal(t, defaultCfg.OSS.Enabled, config.OSS.Enabled, "OSS.Enabled should match")

			// 验证 OSS.Minio 子结构
			assert.Equal(t, defaultCfg.OSS.Minio.Endpoint, config.OSS.Minio.Endpoint, "OSS.Minio.Endpoint should match")
			assert.Equal(t, defaultCfg.OSS.Minio.AccessKey, config.OSS.Minio.AccessKey, "OSS.Minio.AccessKey should match")
			assert.Equal(t, defaultCfg.OSS.Minio.UseSSL, config.OSS.Minio.UseSSL, "OSS.Minio.UseSSL should match")
			assert.Equal(t, defaultCfg.OSS.Minio.Bucket, config.OSS.Minio.Bucket, "OSS.Minio.BucketName should match")

			// 验证 CORS 嵌套结构
			assert.Equal(t, defaultCfg.CORS.ModuleName, config.CORS.ModuleName, "CORS.ModuleName should match")
			assert.Equal(t, defaultCfg.CORS.Enabled, config.CORS.Enabled, "CORS.Enabled should match")
			assert.Equal(t, defaultCfg.CORS.AllowedOrigins, config.CORS.AllowedOrigins, "CORS.AllowedOrigins should match")
			assert.Equal(t, defaultCfg.CORS.AllowedMethods, config.CORS.AllowedMethods, "CORS.AllowedMethods should match")
			assert.Equal(t, defaultCfg.CORS.AllowedHeaders, config.CORS.AllowedHeaders, "CORS.AllowedHeaders should match")
			assert.Equal(t, defaultCfg.CORS.ExposedHeaders, config.CORS.ExposedHeaders, "CORS.ExposedHeaders should match")
			assert.Equal(t, defaultCfg.CORS.AllowCredentials, config.CORS.AllowCredentials, "CORS.AllowCredentials should match")
			assert.Equal(t, defaultCfg.CORS.MaxAge, config.CORS.MaxAge, "CORS.MaxAge should match")
			assert.Equal(t, defaultCfg.CORS.AllowedAllOrigins, config.CORS.AllowedAllOrigins, "CORS.AllowedAllOrigins should match")
			assert.Equal(t, defaultCfg.CORS.AllowedAllMethods, config.CORS.AllowedAllMethods, "CORS.AllowedAllMethods should match")
			assert.Equal(t, defaultCfg.CORS.OptionsResponseCode, config.CORS.OptionsResponseCode, "CORS.OptionsResponseCode should match")

			// 验证 JWT 嵌套结构
			assert.Equal(t, defaultCfg.JWT.ModuleName, config.JWT.ModuleName, "JWT.ModuleName should match")
			assert.Equal(t, defaultCfg.JWT.SigningKey, config.JWT.SigningKey, "JWT.SigningKey should match")
			assert.Equal(t, defaultCfg.JWT.ExpiresTime, config.JWT.ExpiresTime, "JWT.ExpiresTime should match")
			assert.Equal(t, defaultCfg.JWT.BufferTime, config.JWT.BufferTime, "JWT.BufferTime should match")
			assert.Equal(t, defaultCfg.JWT.Issuer, config.JWT.Issuer, "JWT.Issuer should match")
			assert.Equal(t, defaultCfg.JWT.Audience, config.JWT.Audience, "JWT.Audience should match")
			assert.Equal(t, defaultCfg.JWT.Algorithm, config.JWT.Algorithm, "JWT.Algorithm should match")
			assert.Equal(t, defaultCfg.JWT.UseMultipoint, config.JWT.UseMultipoint, "JWT.UseMultipoint should match")

			// 验证 Banner 嵌套结构
			assert.Equal(t, defaultCfg.Banner.ModuleName, config.Banner.ModuleName, "Banner.ModuleName should match")
			assert.Equal(t, defaultCfg.Banner.Enabled, config.Banner.Enabled, "Banner.Enabled should match")
			assert.Equal(t, defaultCfg.Banner.Template, config.Banner.Template, "Banner.Template should match")
			assert.Equal(t, defaultCfg.Banner.Title, config.Banner.Title, "Banner.Title should match")
			assert.Equal(t, defaultCfg.Banner.Description, config.Banner.Description, "Banner.Description should match")
			assert.Equal(t, defaultCfg.Banner.Author, config.Banner.Author, "Banner.Author should match")
			assert.Equal(t, defaultCfg.Banner.Email, config.Banner.Email, "Banner.Email should match")
			assert.Equal(t, defaultCfg.Banner.Website, config.Banner.Website, "Banner.Website should match")

			// 验证 WSC 嵌套结构（完整）
			assert.Equal(t, defaultCfg.WSC.Enabled, config.WSC.Enabled, "WSC.Enabled should match")
			assert.Equal(t, defaultCfg.WSC.NodeIP, config.WSC.NodeIP, "WSC.NodeIP should match")
			assert.Equal(t, defaultCfg.WSC.NodePort, config.WSC.NodePort, "WSC.NodePort should match")
			assert.Equal(t, defaultCfg.WSC.HeartbeatInterval, config.WSC.HeartbeatInterval, "WSC.HeartbeatInterval should match")
			assert.Equal(t, defaultCfg.WSC.ClientTimeout, config.WSC.ClientTimeout, "WSC.ClientTimeout should match")
			assert.Equal(t, defaultCfg.WSC.MessageBufferSize, config.WSC.MessageBufferSize, "WSC.MessageBufferSize should match")
			assert.Equal(t, defaultCfg.WSC.AutoReconnect, config.WSC.AutoReconnect, "WSC.AutoReconnect should match")
			assert.Equal(t, defaultCfg.WSC.MaxRetries, config.WSC.MaxRetries, "WSC.MaxRetries should match")

			// 验证 WSC.Performance 子结构
			assert.Equal(t, defaultCfg.WSC.Performance.MaxConnectionsPerNode, config.WSC.Performance.MaxConnectionsPerNode, "WSC.Performance.MaxConnectionsPerNode should match")
			assert.Equal(t, defaultCfg.WSC.Performance.EnableMetrics, config.WSC.Performance.EnableMetrics, "WSC.Performance.EnableMetrics should match")

			// 验证 WSC.Security 子结构
			assert.Equal(t, defaultCfg.WSC.Security.EnableAuth, config.WSC.Security.EnableAuth, "WSC.Security.EnableAuth should match")
			assert.Equal(t, defaultCfg.WSC.Security.EnableEncryption, config.WSC.Security.EnableEncryption, "WSC.Security.EnableEncryption should match")
			assert.Equal(t, defaultCfg.WSC.Security.EnableRateLimit, config.WSC.Security.EnableRateLimit, "WSC.Security.EnableRateLimit should match")

			// 验证 WSC.Database 子结构
			assert.Equal(t, defaultCfg.WSC.Database.Enabled, config.WSC.Database.Enabled, "WSC.Database.Enabled should match")
			assert.Equal(t, defaultCfg.WSC.Database.AutoMigrate, config.WSC.Database.AutoMigrate, "WSC.Database.AutoMigrate should match")

			// 验证 Mqtt 嵌套结构
			assert.Equal(t, defaultCfg.Mqtt.ModuleName, config.Mqtt.ModuleName, "Mqtt.ModuleName should match")
			assert.Equal(t, defaultCfg.Mqtt.Endpoint, config.Mqtt.Endpoint, "Mqtt.Endpoint should match")
			assert.Equal(t, defaultCfg.Mqtt.ClientID, config.Mqtt.ClientID, "Mqtt.ClientID should match")
			assert.Equal(t, defaultCfg.Mqtt.ProtocolVersion, config.Mqtt.ProtocolVersion, "Mqtt.ProtocolVersion should match")
			assert.Equal(t, defaultCfg.Mqtt.KeepAlive, config.Mqtt.KeepAlive, "Mqtt.KeepAlive should match")
			assert.Equal(t, defaultCfg.Mqtt.Username, config.Mqtt.Username, "Mqtt.Username should match")
			assert.Equal(t, defaultCfg.Mqtt.CleanSession, config.Mqtt.CleanSession, "Mqtt.CleanSession should match")
			assert.Equal(t, defaultCfg.Mqtt.AutoReconnect, config.Mqtt.AutoReconnect, "Mqtt.AutoReconnect should match")

			// 验证 Elasticsearch 嵌套结构
			assert.Equal(t, defaultCfg.Elasticsearch.ModuleName, config.Elasticsearch.ModuleName, "Elasticsearch.ModuleName should match")
			assert.Equal(t, defaultCfg.Elasticsearch.Endpoint, config.Elasticsearch.Endpoint, "Elasticsearch.Endpoint should match")
			assert.Equal(t, defaultCfg.Elasticsearch.HealthCheck, config.Elasticsearch.HealthCheck, "Elasticsearch.HealthCheck should match")
			assert.Equal(t, defaultCfg.Elasticsearch.Sniff, config.Elasticsearch.Sniff, "Elasticsearch.Sniff should match")
			assert.Equal(t, defaultCfg.Elasticsearch.Gzip, config.Elasticsearch.Gzip, "Elasticsearch.Gzip should match")
			assert.Equal(t, defaultCfg.Elasticsearch.Timeout, config.Elasticsearch.Timeout, "Elasticsearch.Timeout should match")

			// 验证 Smtp 嵌套结构
			assert.Equal(t, defaultCfg.Smtp.ModuleName, config.Smtp.ModuleName, "Smtp.ModuleName should match")
			assert.Equal(t, defaultCfg.Smtp.Enabled, config.Smtp.Enabled, "Smtp.Enabled should match")
			assert.Equal(t, defaultCfg.Smtp.SMTPHost, config.Smtp.SMTPHost, "Smtp.SMTPHost should match")
			assert.Equal(t, defaultCfg.Smtp.SMTPPort, config.Smtp.SMTPPort, "Smtp.SMTPPort should match")
			assert.Equal(t, defaultCfg.Smtp.Username, config.Smtp.Username, "Smtp.Username should match")
			assert.Equal(t, defaultCfg.Smtp.FromAddress, config.Smtp.FromAddress, "Smtp.FromAddress should match")
			assert.Equal(t, defaultCfg.Smtp.EnableTLS, config.Smtp.EnableTLS, "Smtp.EnableTLS should match")
			assert.Equal(t, defaultCfg.Smtp.PoolSize, config.Smtp.PoolSize, "Smtp.PoolSize should match")

			// 验证 Health 嵌套结构（完整）
			assert.Equal(t, defaultCfg.Health.ModuleName, config.Health.ModuleName, "Health.ModuleName should match")
			assert.Equal(t, defaultCfg.Health.Enabled, config.Health.Enabled, "Health.Enabled should match")
			assert.Equal(t, defaultCfg.Health.Path, config.Health.Path, "Health.Path should match")
			assert.Equal(t, defaultCfg.Health.Port, config.Health.Port, "Health.Port should match")
			assert.Equal(t, defaultCfg.Health.Timeout, config.Health.Timeout, "Health.Timeout should match")

			// 验证 Health.Redis 子结构
			assert.Equal(t, defaultCfg.Health.Redis.Enabled, config.Health.Redis.Enabled, "Health.Redis.Enabled should match")
			assert.Equal(t, defaultCfg.Health.Redis.Path, config.Health.Redis.Path, "Health.Redis.Path should match")
			assert.Equal(t, defaultCfg.Health.Redis.Timeout, config.Health.Redis.Timeout, "Health.Redis.Timeout should match")

			// 验证 Health.MySQL 子结构
			assert.Equal(t, defaultCfg.Health.MySQL.Enabled, config.Health.MySQL.Enabled, "Health.MySQL.Enabled should match")
			assert.Equal(t, defaultCfg.Health.MySQL.Path, config.Health.MySQL.Path, "Health.MySQL.Path should match")
			assert.Equal(t, defaultCfg.Health.MySQL.Timeout, config.Health.MySQL.Timeout, "Health.MySQL.Timeout should match")

			// 验证 Monitoring 嵌套结构（完整）
			assert.Equal(t, defaultCfg.Monitoring.ModuleName, config.Monitoring.ModuleName, "Monitoring.ModuleName should match")
			assert.Equal(t, defaultCfg.Monitoring.Enabled, config.Monitoring.Enabled, "Monitoring.Enabled should match")

			// 验证 Monitoring.Metrics 子结构
			assert.Equal(t, defaultCfg.Monitoring.Metrics.Enabled, config.Monitoring.Metrics.Enabled, "Monitoring.Metrics.Enabled should match")
			assert.Equal(t, defaultCfg.Monitoring.Metrics.RequestCount, config.Monitoring.Metrics.RequestCount, "Monitoring.Metrics.RequestCount should match")
			assert.Equal(t, defaultCfg.Monitoring.Metrics.Duration, config.Monitoring.Metrics.Duration, "Monitoring.Metrics.Duration should match")
			assert.Equal(t, defaultCfg.Monitoring.Metrics.ResponseSize, config.Monitoring.Metrics.ResponseSize, "Monitoring.Metrics.ResponseSize should match")
			assert.Equal(t, defaultCfg.Monitoring.Metrics.RequestSize, config.Monitoring.Metrics.RequestSize, "Monitoring.Metrics.RequestSize should match")

			// 验证 Monitoring.Prometheus 子结构
			assert.Equal(t, defaultCfg.Monitoring.Prometheus.Enabled, config.Monitoring.Prometheus.Enabled, "Monitoring.Prometheus.Enabled should match")

			// 验证 Monitoring.Grafana 子结构
			assert.Equal(t, defaultCfg.Monitoring.Grafana.Enabled, config.Monitoring.Grafana.Enabled, "Monitoring.Grafana.Enabled should match")

			// 验证 Monitoring.Jaeger 子结构
			assert.Equal(t, defaultCfg.Monitoring.Jaeger.Enabled, config.Monitoring.Jaeger.Enabled, "Monitoring.Jaeger.Enabled should match")

			// 验证 Monitoring.Alerting 子结构
			assert.Equal(t, defaultCfg.Monitoring.Alerting.Enabled, config.Monitoring.Alerting.Enabled, "Monitoring.Alerting.Enabled should match")

			// 验证 Security 嵌套结构（完整）
			assert.Equal(t, defaultCfg.Security.ModuleName, config.Security.ModuleName, "Security.ModuleName should match")

			// 验证 Security.CSP 子结构
			assert.NotNil(t, config.Security.CSP, "Security.CSP should not be nil")
			assert.Equal(t, defaultCfg.Security.CSP.Enabled, config.Security.CSP.Enabled, "Security.CSP.Enabled should match")
			assert.Equal(t, defaultCfg.Security.CSP.Mode, config.Security.CSP.Mode, "Security.CSP.Mode should match")
			assert.Equal(t, defaultCfg.Security.CSP.Custom, config.Security.CSP.Custom, "Security.CSP.Custom should match")

			// 验证 Security.JWT 子结构
			assert.Equal(t, defaultCfg.Security.JWT.Enabled, config.Security.JWT.Enabled, "Security.JWT.Enabled should match")
			assert.Equal(t, defaultCfg.Security.JWT.Secret, config.Security.JWT.Secret, "Security.JWT.Secret should match")
			assert.Equal(t, defaultCfg.Security.JWT.Expiry, config.Security.JWT.Expiry, "Security.JWT.Expiry should match")
			assert.Equal(t, defaultCfg.Security.JWT.Issuer, config.Security.JWT.Issuer, "Security.JWT.Issuer should match")
			assert.Equal(t, defaultCfg.Security.JWT.Algorithm, config.Security.JWT.Algorithm, "Security.JWT.Algorithm should match")

			// 验证 Security.Auth 子结构
			assert.Equal(t, defaultCfg.Security.Auth.Enabled, config.Security.Auth.Enabled, "Security.Auth.Enabled should match")
			assert.Equal(t, defaultCfg.Security.Auth.Type, config.Security.Auth.Type, "Security.Auth.Type should match")
			assert.Equal(t, defaultCfg.Security.Auth.HeaderName, config.Security.Auth.HeaderName, "Security.Auth.HeaderName should match")
			assert.Equal(t, defaultCfg.Security.Auth.TokenPrefix, config.Security.Auth.TokenPrefix, "Security.Auth.TokenPrefix should match")

			// 验证 Security.Auth.Basic 子结构
			assert.NotNil(t, config.Security.Auth.Basic, "Security.Auth.Basic should not be nil")
			assert.Equal(t, len(defaultCfg.Security.Auth.Basic.Users), len(config.Security.Auth.Basic.Users), "Security.Auth.Basic.Users length should match")

			// 验证 Security.Auth.Bearer 子结构
			assert.NotNil(t, config.Security.Auth.Bearer, "Security.Auth.Bearer should not be nil")
			assert.Equal(t, len(defaultCfg.Security.Auth.Bearer.Tokens), len(config.Security.Auth.Bearer.Tokens), "Security.Auth.Bearer.Tokens length should match")

			// 验证 Security.Auth.APIKey 子结构
			assert.NotNil(t, config.Security.Auth.APIKey, "Security.Auth.APIKey should not be nil")
			assert.Equal(t, len(defaultCfg.Security.Auth.APIKey.Keys), len(config.Security.Auth.APIKey.Keys), "Security.Auth.APIKey.Keys length should match")
			assert.Equal(t, defaultCfg.Security.Auth.APIKey.HeaderName, config.Security.Auth.APIKey.HeaderName, "Security.Auth.APIKey.HeaderName should match")
			assert.Equal(t, defaultCfg.Security.Auth.APIKey.QueryParam, config.Security.Auth.APIKey.QueryParam, "Security.Auth.APIKey.QueryParam should match")

			// 验证 Security.Auth.Custom 子结构
			assert.NotNil(t, config.Security.Auth.Custom, "Security.Auth.Custom should not be nil")
			assert.Equal(t, defaultCfg.Security.Auth.Custom.HeaderName, config.Security.Auth.Custom.HeaderName, "Security.Auth.Custom.HeaderName should match")
			assert.Equal(t, defaultCfg.Security.Auth.Custom.ExpectedValue, config.Security.Auth.Custom.ExpectedValue, "Security.Auth.Custom.ExpectedValue should match")
			assert.Equal(t, len(defaultCfg.Security.Auth.Custom.Headers), len(config.Security.Auth.Custom.Headers), "Security.Auth.Custom.Headers length should match")

			// 验证 Security.Protection 子结构
			assert.NotNil(t, config.Security.Protection, "Security.Protection should not be nil")

			// 验证 Security.Protection.Swagger 详细字段
			assert.NotNil(t, config.Security.Protection.Swagger, "Security.Protection.Swagger should not be nil")
			assert.Equal(t, defaultCfg.Security.Protection.Swagger.Enabled, config.Security.Protection.Swagger.Enabled, "Security.Protection.Swagger.Enabled should match")
			assert.Equal(t, defaultCfg.Security.Protection.Swagger.AuthRequired, config.Security.Protection.Swagger.AuthRequired, "Security.Protection.Swagger.AuthRequired should match")
			assert.Equal(t, defaultCfg.Security.Protection.Swagger.AuthType, config.Security.Protection.Swagger.AuthType, "Security.Protection.Swagger.AuthType should match")
			assert.Equal(t, defaultCfg.Security.Protection.Swagger.IPWhitelist, config.Security.Protection.Swagger.IPWhitelist, "Security.Protection.Swagger.IPWhitelist should match")
			assert.Equal(t, defaultCfg.Security.Protection.Swagger.RequireHTTPS, config.Security.Protection.Swagger.RequireHTTPS, "Security.Protection.Swagger.RequireHTTPS should match")
			assert.Equal(t, defaultCfg.Security.Protection.Swagger.Username, config.Security.Protection.Swagger.Username, "Security.Protection.Swagger.Username should match")
			assert.Equal(t, defaultCfg.Security.Protection.Swagger.Password, config.Security.Protection.Swagger.Password, "Security.Protection.Swagger.Password should match")

			// 验证 Security.Protection.PProf 详细字段
			assert.NotNil(t, config.Security.Protection.PProf, "Security.Protection.PProf should not be nil")
			assert.Equal(t, defaultCfg.Security.Protection.PProf.Enabled, config.Security.Protection.PProf.Enabled, "Security.Protection.PProf.Enabled should match")
			assert.Equal(t, defaultCfg.Security.Protection.PProf.AuthRequired, config.Security.Protection.PProf.AuthRequired, "Security.Protection.PProf.AuthRequired should match")
			assert.Equal(t, defaultCfg.Security.Protection.PProf.AuthType, config.Security.Protection.PProf.AuthType, "Security.Protection.PProf.AuthType should match")
			assert.Equal(t, defaultCfg.Security.Protection.PProf.IPWhitelist, config.Security.Protection.PProf.IPWhitelist, "Security.Protection.PProf.IPWhitelist should match")
			assert.Equal(t, defaultCfg.Security.Protection.PProf.RequireHTTPS, config.Security.Protection.PProf.RequireHTTPS, "Security.Protection.PProf.RequireHTTPS should match")
			assert.Equal(t, defaultCfg.Security.Protection.PProf.Username, config.Security.Protection.PProf.Username, "Security.Protection.PProf.Username should match")
			assert.Equal(t, defaultCfg.Security.Protection.PProf.Password, config.Security.Protection.PProf.Password, "Security.Protection.PProf.Password should match")

			// 验证 Security.Protection.Metrics 详细字段
			assert.NotNil(t, config.Security.Protection.Metrics, "Security.Protection.Metrics should not be nil")
			assert.Equal(t, defaultCfg.Security.Protection.Metrics.Enabled, config.Security.Protection.Metrics.Enabled, "Security.Protection.Metrics.Enabled should match")
			assert.Equal(t, defaultCfg.Security.Protection.Metrics.AuthRequired, config.Security.Protection.Metrics.AuthRequired, "Security.Protection.Metrics.AuthRequired should match")
			assert.Equal(t, defaultCfg.Security.Protection.Metrics.IPWhitelist, config.Security.Protection.Metrics.IPWhitelist, "Security.Protection.Metrics.IPWhitelist should match")

			// 验证 Security.Protection.Health 详细字段
			assert.NotNil(t, config.Security.Protection.Health, "Security.Protection.Health should not be nil")
			assert.Equal(t, defaultCfg.Security.Protection.Health.Enabled, config.Security.Protection.Health.Enabled, "Security.Protection.Health.Enabled should match")
			assert.Equal(t, defaultCfg.Security.Protection.Health.AuthRequired, config.Security.Protection.Health.AuthRequired, "Security.Protection.Health.AuthRequired should match")
			assert.Equal(t, defaultCfg.Security.Protection.Health.IPWhitelist, config.Security.Protection.Health.IPWhitelist, "Security.Protection.Health.IPWhitelist should match")

			// 验证 Security.Protection.API 详细字段
			assert.NotNil(t, config.Security.Protection.API, "Security.Protection.API should not be nil")
			assert.Equal(t, defaultCfg.Security.Protection.API.Enabled, config.Security.Protection.API.Enabled, "Security.Protection.API.Enabled should match")
			assert.Equal(t, defaultCfg.Security.Protection.API.AuthRequired, config.Security.Protection.API.AuthRequired, "Security.Protection.API.AuthRequired should match")
			assert.Equal(t, defaultCfg.Security.Protection.API.IPWhitelist, config.Security.Protection.API.IPWhitelist, "Security.Protection.API.IPWhitelist should match")

			// 验证 Middleware 嵌套结构（完整）
			assert.Equal(t, defaultCfg.Middleware.ModuleName, config.Middleware.ModuleName, "Middleware.ModuleName should match")
			assert.Equal(t, defaultCfg.Middleware.Enabled, config.Middleware.Enabled, "Middleware.Enabled should match")

			// 验证 Middleware 子模块
			assert.Equal(t, defaultCfg.Middleware.Logging.Enabled, config.Middleware.Logging.Enabled, "Middleware.Logging.Enabled should match")
			assert.Equal(t, defaultCfg.Middleware.Recovery.Enabled, config.Middleware.Recovery.Enabled, "Middleware.Recovery.Enabled should match")
			assert.Equal(t, defaultCfg.Middleware.Tracing.Enabled, config.Middleware.Tracing.Enabled, "Middleware.Tracing.Enabled should match")
			assert.Equal(t, defaultCfg.Middleware.Metrics.Enabled, config.Middleware.Metrics.Enabled, "Middleware.Metrics.Enabled should match")
			assert.Equal(t, defaultCfg.Middleware.RequestID.Enabled, config.Middleware.RequestID.Enabled, "Middleware.RequestID.Enabled should match")
			assert.Equal(t, defaultCfg.Middleware.PProf.Enabled, config.Middleware.PProf.Enabled, "Middleware.PProf.Enabled should match")

			// 验证 Middleware 额外子模块
			assert.NotNil(t, config.Middleware.I18N, "Middleware.I18N should not be nil")
			assert.Equal(t, defaultCfg.Middleware.I18N.Enabled, config.Middleware.I18N.Enabled, "Middleware.I18N.Enabled should match")
			assert.NotNil(t, config.Middleware.CircuitBreaker, "Middleware.CircuitBreaker should not be nil")
			assert.Equal(t, defaultCfg.Middleware.CircuitBreaker.Enabled, config.Middleware.CircuitBreaker.Enabled, "Middleware.CircuitBreaker.Enabled should match")
			assert.NotNil(t, config.Middleware.Alerting, "Middleware.Alerting should not be nil")
			assert.Equal(t, defaultCfg.Middleware.Alerting.Enabled, config.Middleware.Alerting.Enabled, "Middleware.Alerting.Enabled should match")
			assert.NotNil(t, config.Middleware.Signature, "Middleware.Signature should not be nil")
			assert.Equal(t, defaultCfg.Middleware.Signature.Enabled, config.Middleware.Signature.Enabled, "Middleware.Signature.Enabled should match")

			// 验证 Swagger 嵌套结构（完整）
			assert.Equal(t, defaultCfg.Swagger.ModuleName, config.Swagger.ModuleName, "Swagger.ModuleName should match")
			assert.Equal(t, defaultCfg.Swagger.Enabled, config.Swagger.Enabled, "Swagger.Enabled should match")
			assert.Equal(t, defaultCfg.Swagger.JSONPath, config.Swagger.JSONPath, "Swagger.JSONPath should match")
			assert.Equal(t, defaultCfg.Swagger.UIPath, config.Swagger.UIPath, "Swagger.UIPath should match")
			assert.Equal(t, defaultCfg.Swagger.Title, config.Swagger.Title, "Swagger.Title should match")
			assert.Equal(t, defaultCfg.Swagger.Description, config.Swagger.Description, "Swagger.Description should match")
			assert.Equal(t, defaultCfg.Swagger.Version, config.Swagger.Version, "Swagger.Version should match")

			// 验证 Swagger.Contact 子结构
			assert.Equal(t, defaultCfg.Swagger.Contact.Name, config.Swagger.Contact.Name, "Swagger.Contact.Name should match")
			assert.Equal(t, defaultCfg.Swagger.Contact.URL, config.Swagger.Contact.URL, "Swagger.Contact.URL should match")
			assert.Equal(t, defaultCfg.Swagger.Contact.Email, config.Swagger.Contact.Email, "Swagger.Contact.Email should match")

			// 验证 Swagger.License 子结构
			assert.Equal(t, defaultCfg.Swagger.License.Name, config.Swagger.License.Name, "Swagger.License.Name should match")
			assert.Equal(t, defaultCfg.Swagger.License.URL, config.Swagger.License.URL, "Swagger.License.URL should match")

			// 验证 Swagger.Auth 子结构
			assert.NotNil(t, config.Swagger.Auth, "Swagger.Auth should not be nil")
			assert.Equal(t, defaultCfg.Swagger.Auth.Enabled, config.Swagger.Auth.Enabled, "Swagger.Auth.Enabled should match")
			assert.Equal(t, defaultCfg.Swagger.Auth.Type, config.Swagger.Auth.Type, "Swagger.Auth.Type should match")

			// 验证 RateLimit 嵌套结构（完整）
			assert.Equal(t, defaultCfg.RateLimit.ModuleName, config.RateLimit.ModuleName, "RateLimit.ModuleName should match")
			assert.Equal(t, defaultCfg.RateLimit.Enabled, config.RateLimit.Enabled, "RateLimit.Enabled should match")
			assert.Equal(t, defaultCfg.RateLimit.Strategy, config.RateLimit.Strategy, "RateLimit.Strategy should match")
			assert.Equal(t, defaultCfg.RateLimit.DefaultScope, config.RateLimit.DefaultScope, "RateLimit.DefaultScope should match")

			// 验证 RateLimit.GlobalLimit 子结构
			assert.Equal(t, defaultCfg.RateLimit.GlobalLimit.RequestsPerSecond, config.RateLimit.GlobalLimit.RequestsPerSecond, "RateLimit.GlobalLimit.RequestsPerSecond should match")
			assert.Equal(t, defaultCfg.RateLimit.GlobalLimit.BurstSize, config.RateLimit.GlobalLimit.BurstSize, "RateLimit.GlobalLimit.BurstSize should match")

			// 验证 RateLimit.Storage 子结构
			assert.Equal(t, defaultCfg.RateLimit.Storage.Type, config.RateLimit.Storage.Type, "RateLimit.Storage.Type should match")
			assert.Equal(t, defaultCfg.RateLimit.Storage.KeyPrefix, config.RateLimit.Storage.KeyPrefix, "RateLimit.Storage.KeyPrefix should match")
			assert.Equal(t, defaultCfg.RateLimit.Storage.CleanInterval, config.RateLimit.Storage.CleanInterval, "RateLimit.Storage.CleanInterval should match")

			// 验证 RateLimit 详细字段
			assert.Equal(t, len(defaultCfg.RateLimit.Routes), len(config.RateLimit.Routes), "RateLimit.Routes length should match")
			assert.Equal(t, len(defaultCfg.RateLimit.IPRules), len(config.RateLimit.IPRules), "RateLimit.IPRules length should match")
			assert.Equal(t, len(defaultCfg.RateLimit.UserRules), len(config.RateLimit.UserRules), "RateLimit.UserRules length should match")
			assert.Equal(t, defaultCfg.RateLimit.CustomRuleLoader, config.RateLimit.CustomRuleLoader, "RateLimit.CustomRuleLoader should match")
			assert.Equal(t, defaultCfg.RateLimit.EnableDynamicRule, config.RateLimit.EnableDynamicRule, "RateLimit.EnableDynamicRule should match")

			// 验证 Jobs 嵌套结构（完整）
			assert.NotNil(t, config.Jobs, "Jobs should not be nil")
			assert.Equal(t, defaultCfg.Jobs.Enabled, config.Jobs.Enabled, "Jobs.Enabled should match")
			assert.Equal(t, defaultCfg.Jobs.TimeZone, config.Jobs.TimeZone, "Jobs.TimeZone should match")
			assert.Equal(t, defaultCfg.Jobs.GracefulShutdown, config.Jobs.GracefulShutdown, "Jobs.GracefulShutdown should match")
			assert.Equal(t, defaultCfg.Jobs.MaxRetries, config.Jobs.MaxRetries, "Jobs.MaxRetries should match")
			assert.Equal(t, defaultCfg.Jobs.RetryInterval, config.Jobs.RetryInterval, "Jobs.RetryInterval should match")
			assert.Equal(t, defaultCfg.Jobs.RetryJitter, config.Jobs.RetryJitter, "Jobs.RetryJitter should match")
			assert.Equal(t, defaultCfg.Jobs.MaxConcurrentJobs, config.Jobs.MaxConcurrentJobs, "Jobs.MaxConcurrentJobs should match")
			assert.Equal(t, defaultCfg.Jobs.Distribute, config.Jobs.Distribute, "Jobs.Distribute should match")
			assert.Equal(t, len(defaultCfg.Jobs.Tasks), len(config.Jobs.Tasks), "Jobs.Tasks length should match")

			t.Logf("✅ Gateway 模块所有字段（包括多层嵌套结构）验证通过")
		})
	})
}

// testCompleteModuleConfig 完整的模块配置测试函数（支持YAML和JSON）
func testCompleteModuleConfig(t *testing.T, baseDir, moduleName string, configType interface{}, validateFunc func(interface{})) {
	yamlPath := baseDir + "/pkg/" + moduleName + "/" + moduleName + ".yaml"

	// 创建 Viper 实例
	v := viper.New()
	v.SetConfigFile(yamlPath)
	err := v.ReadInConfig()
	require.NoError(t, err, "读取YAML配置文件失败")

	// 解析配置
	err = v.Unmarshal(configType, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "yaml"
		dc.WeaklyTypedInput = true
	})
	require.NoError(t, err, "解析YAML配置失败")

	// 创建热加载器
	hotReloader, err := NewHotReloader(configType, v, yamlPath, &HotReloadConfig{
		Enabled: false,
	})
	require.NoError(t, err, "创建热加载器失败")

	// 验证配置
	loadedConfig := hotReloader.GetConfig()
	require.NotNil(t, loadedConfig, "加载的配置不应为空")

	// 执行完整验证
	if validateFunc != nil {
		validateFunc(loadedConfig)
	}

	// 测试热重载
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = hotReloader.Reload(ctx)
	require.NoError(t, err, "配置重载失败")

	reloadedConfig := hotReloader.GetConfig()
	require.NotNil(t, reloadedConfig, "重载的配置不应为空")

	// 再次验证
	if validateFunc != nil {
		validateFunc(reloadedConfig)
	}
}
