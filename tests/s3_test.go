/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-08 12:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-08 12:00:00
 * @FilePath: \go-config\tests\s3_test.go
 * @Description: S3 OSS 配置测试
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"context"
	"testing"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/oss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestS3Config(t *testing.T) {
	t.Run("TestS3BasicConfig", testS3BasicConfig)
	t.Run("TestS3MultipleInstances", testS3MultipleInstances)
	t.Run("TestS3Validation", testS3Validation)
	t.Run("TestS3Methods", testS3Methods)
}

func testS3BasicConfig(t *testing.T) {
	// 创建S3配置实例
	s3Config := &oss.S3{
		Endpoint:     "https://s3.ap-southeast-1.amazonaws.com",
		Region:       "ap-southeast-1",
		AccessKey:    "your-access-key-id",
		SecretKey:    "your-secret-access-key",
		BucketPrefix: "aicsqa",
		UseSSL:       true,
		PathStyle:    false,
		ModuleName:   "default",
	}

	// 测试新建实例
	newS3 := oss.NewS3(s3Config)
	assert.NotNil(t, newS3)
	assert.Equal(t, s3Config.Endpoint, newS3.Endpoint)
	assert.Equal(t, s3Config.Region, newS3.Region)
	assert.Equal(t, s3Config.AccessKey, newS3.AccessKey)
	assert.Equal(t, s3Config.BucketPrefix, newS3.BucketPrefix)

	// 测试验证
	err := newS3.Validate()
	assert.NoError(t, err)
}

func testS3MultipleInstances(t *testing.T) {
	ctx := context.Background()

	// 创建多配置管理器
	manager, err := goconfig.NewMultiConfigManager(ctx, &goconfig.ConfigOptions{
		ConfigType:   "yaml",
		ConfigPath:   "./resources",
		ConfigSuffix: "_config",
	})
	require.NoError(t, err)

	multiConfig := manager.GetConfig()

	// 模拟多个S3实例配置
	s3Configs := []oss.S3{
		{
			ModuleName:   "primary",
			Endpoint:     "https://s3.us-west-2.amazonaws.com",
			Region:       "us-west-2",
			AccessKey:    "primary-access-key",
			SecretKey:    "primary-secret-key",
			BucketPrefix: "primary",
			UseSSL:       true,
			PathStyle:    false,
		},
		{
			ModuleName:   "backup",
			Endpoint:     "https://s3.eu-west-1.amazonaws.com",
			Region:       "eu-west-1",
			AccessKey:    "backup-access-key",
			SecretKey:    "backup-secret-key",
			BucketPrefix: "backup",
			UseSSL:       true,
			PathStyle:    false,
		},
	}

	// 将配置添加到多配置中
	multiConfig.S3 = s3Configs

	// 测试根据模块名获取配置
	primaryS3, err := goconfig.GetModuleByName(multiConfig.S3, "primary")
	assert.NoError(t, err)
	assert.Equal(t, "primary", primaryS3.ModuleName)
	assert.Equal(t, "us-west-2", primaryS3.Region)

	backupS3, err := goconfig.GetModuleByName(multiConfig.S3, "backup")
	assert.NoError(t, err)
	assert.Equal(t, "backup", backupS3.ModuleName)
	assert.Equal(t, "eu-west-1", backupS3.Region)

	// 测试获取不存在的模块
	_, err = goconfig.GetModuleByName(multiConfig.S3, "nonexistent")
	assert.Error(t, err)
}

func testS3Validation(t *testing.T) {
	tests := []struct {
		name      string
		config    oss.S3
		expectErr bool
	}{
		{
			name: "Valid Config",
			config: oss.S3{
				Endpoint:  "https://s3.amazonaws.com",
				Region:    "us-east-1",
				AccessKey: "valid-access-key",
				SecretKey: "valid-secret-key",
			},
			expectErr: false,
		},
		{
			name: "Missing Endpoint",
			config: oss.S3{
				Region:    "us-east-1",
				AccessKey: "valid-access-key",
				SecretKey: "valid-secret-key",
			},
			expectErr: true,
		},
		{
			name: "Missing Region",
			config: oss.S3{
				Endpoint:  "https://s3.amazonaws.com",
				AccessKey: "valid-access-key",
				SecretKey: "valid-secret-key",
			},
			expectErr: true,
		},
		{
			name: "Missing AccessKey",
			config: oss.S3{
				Endpoint:  "https://s3.amazonaws.com",
				Region:    "us-east-1",
				SecretKey: "valid-secret-key",
			},
			expectErr: true,
		},
		{
			name: "Missing SecretKey",
			config: oss.S3{
				Endpoint:  "https://s3.amazonaws.com",
				Region:    "us-east-1",
				AccessKey: "valid-access-key",
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.expectErr {
				assert.Error(t, err, "Expected validation error for %s", tt.name)
			} else {
				assert.NoError(t, err, "Expected no validation error for %s", tt.name)
			}
		})
	}
}

func testS3Methods(t *testing.T) {
	original := &oss.S3{
		Endpoint:     "https://s3.amazonaws.com",
		Region:       "us-east-1",
		AccessKey:    "test-access-key",
		SecretKey:    "test-secret-key",
		BucketPrefix: "test",
		SessionToken: "test-session-token",
		UseSSL:       true,
		PathStyle:    false,
		ModuleName:   "test",
	}

	// 测试 Clone 方法
	t.Run("Clone", func(t *testing.T) {
		cloned := original.Clone()
		clonedS3, ok := cloned.(*oss.S3)
		require.True(t, ok)

		assert.Equal(t, original.Endpoint, clonedS3.Endpoint)
		assert.Equal(t, original.Region, clonedS3.Region)
		assert.Equal(t, original.AccessKey, clonedS3.AccessKey)
		assert.Equal(t, original.SecretKey, clonedS3.SecretKey)
		assert.Equal(t, original.BucketPrefix, clonedS3.BucketPrefix)
		assert.Equal(t, original.SessionToken, clonedS3.SessionToken)
		assert.Equal(t, original.UseSSL, clonedS3.UseSSL)
		assert.Equal(t, original.PathStyle, clonedS3.PathStyle)
		assert.Equal(t, original.ModuleName, clonedS3.ModuleName)

		// 修改克隆对象不应影响原对象
		clonedS3.Endpoint = "https://modified.endpoint.com"
		assert.NotEqual(t, original.Endpoint, clonedS3.Endpoint)
	})

	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		result := original.Get()
		resultS3, ok := result.(*oss.S3)
		require.True(t, ok)
		assert.Equal(t, original, resultS3)
	})

	// 测试 Set 方法
	t.Run("Set", func(t *testing.T) {
		target := &oss.S3{}
		target.Set(original)

		assert.Equal(t, original.Endpoint, target.Endpoint)
		assert.Equal(t, original.Region, target.Region)
		assert.Equal(t, original.AccessKey, target.AccessKey)
		assert.Equal(t, original.SecretKey, target.SecretKey)
		assert.Equal(t, original.BucketPrefix, target.BucketPrefix)
		assert.Equal(t, original.SessionToken, target.SessionToken)
		assert.Equal(t, original.UseSSL, target.UseSSL)
		assert.Equal(t, original.PathStyle, target.PathStyle)
		assert.Equal(t, original.ModuleName, target.ModuleName)
	})

	// 测试 GetBucketName 方法
	t.Run("GetBucketName", func(t *testing.T) {
		tests := []struct {
			prefix   string
			suffix   string
			expected string
		}{
			{"test", "bucket", "test-bucket"},
			{"", "bucket", "bucket"},
			{"test", "", "test"},
			{"", "", ""},
		}

		for _, tt := range tests {
			s3Config := &oss.S3{BucketPrefix: tt.prefix}
			result := s3Config.GetBucketName(tt.suffix)
			assert.Equal(t, tt.expected, result)
		}
	})

	// 测试 IsTemporaryCredentials 方法
	t.Run("IsTemporaryCredentials", func(t *testing.T) {
		// 有会话令牌
		s3WithToken := &oss.S3{SessionToken: "test-token"}
		assert.True(t, s3WithToken.IsTemporaryCredentials())

		// 无会话令牌
		s3WithoutToken := &oss.S3{}
		assert.False(t, s3WithoutToken.IsTemporaryCredentials())
	})
}