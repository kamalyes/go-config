/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\signature\signature_test.go
 * @Description: 签名验证中间件配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package signature

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSignature_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "signature", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "default-secret-key-change-in-production", config.SecretKey)
	assert.Equal(t, "X-Signature", config.SignatureHeader)
	assert.Equal(t, "X-Timestamp", config.TimestampHeader)
	assert.Equal(t, "X-Nonce", config.NonceHeader)
	assert.Equal(t, "sha256", config.Algorithm)
	assert.Equal(t, time.Minute*5, config.TimeoutWindow)
	assert.NotNil(t, config.IgnorePaths)
	assert.NotNil(t, config.RequiredHeaders)
	assert.False(t, config.SkipQuery)
	assert.False(t, config.SkipBody)
}

func TestSignature_WithSecretKey(t *testing.T) {
	config := Default()
	result := config.WithSecretKey("my-secret-key")
	assert.Equal(t, "my-secret-key", result.SecretKey)
	assert.Equal(t, config, result)
}

func TestSignature_WithAlgorithm(t *testing.T) {
	config := Default()
	result := config.WithAlgorithm("sha512")
	assert.Equal(t, "sha512", result.Algorithm)
	assert.Equal(t, config, result)
}

func TestSignature_WithTimeoutWindow(t *testing.T) {
	config := Default()
	result := config.WithTimeoutWindow(time.Minute * 10)
	assert.Equal(t, time.Minute*10, result.TimeoutWindow)
	assert.Equal(t, config, result)
}

func TestSignature_WithSignatureHeader(t *testing.T) {
	config := Default()
	result := config.WithSignatureHeader("X-Custom-Signature")
	assert.Equal(t, "X-Custom-Signature", result.SignatureHeader)
	assert.Equal(t, config, result)
}

func TestSignature_WithTimestampHeader(t *testing.T) {
	config := Default()
	result := config.WithTimestampHeader("X-Custom-Timestamp")
	assert.Equal(t, "X-Custom-Timestamp", result.TimestampHeader)
	assert.Equal(t, config, result)
}

func TestSignature_WithNonceHeader(t *testing.T) {
	config := Default()
	result := config.WithNonceHeader("X-Custom-Nonce")
	assert.Equal(t, "X-Custom-Nonce", result.NonceHeader)
	assert.Equal(t, config, result)
}

func TestSignature_AddIgnorePath(t *testing.T) {
	config := Default()
	initialLen := len(config.IgnorePaths)
	result := config.AddIgnorePath("/api/public")
	assert.Equal(t, initialLen+1, len(result.IgnorePaths))
	assert.Contains(t, result.IgnorePaths, "/api/public")
	assert.Equal(t, config, result)
}

func TestSignature_AddRequiredHeader(t *testing.T) {
	config := Default()
	initialLen := len(config.RequiredHeaders)
	result := config.AddRequiredHeader("X-Custom-Header")
	assert.Equal(t, initialLen+1, len(result.RequiredHeaders))
	assert.Contains(t, result.RequiredHeaders, "X-Custom-Header")
	assert.Equal(t, config, result)
}

func TestSignature_WithSkipQuery(t *testing.T) {
	config := Default()
	result := config.WithSkipQuery(true)
	assert.True(t, result.SkipQuery)
	assert.Equal(t, config, result)
}

func TestSignature_WithSkipBody(t *testing.T) {
	config := Default()
	result := config.WithSkipBody(true)
	assert.True(t, result.SkipBody)
	assert.Equal(t, config, result)
}

func TestSignature_Enable(t *testing.T) {
	config := Default()
	result := config.Enable()
	assert.True(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestSignature_Disable(t *testing.T) {
	config := Default()
	config.Enabled = true
	result := config.Disable()
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestSignature_IsEnabled(t *testing.T) {
	config := Default()
	assert.False(t, config.IsEnabled())

	config.Enabled = true
	assert.True(t, config.IsEnabled())
}

func TestSignature_Clone(t *testing.T) {
	config := Default()
	config.WithSecretKey("test-key").AddIgnorePath("/custom").AddRequiredHeader("X-Custom")

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*Signature)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.SecretKey, clonedConfig.SecretKey)
	assert.Equal(t, config.TimeoutWindow, clonedConfig.TimeoutWindow)

	// 验证深拷贝 - 切片
	clonedConfig.IgnorePaths = append(clonedConfig.IgnorePaths, "/extra")
	assert.NotEqual(t, len(config.IgnorePaths), len(clonedConfig.IgnorePaths))

	clonedConfig.RequiredHeaders = append(clonedConfig.RequiredHeaders, "X-Extra")
	assert.NotEqual(t, len(config.RequiredHeaders), len(clonedConfig.RequiredHeaders))
}

func TestSignature_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestSignature_Set(t *testing.T) {
	config := Default()
	newConfig := &Signature{
		ModuleName:    "custom-signature",
		Enabled:       true,
		SecretKey:     "new-key",
		Algorithm:     "md5",
		TimeoutWindow: time.Minute * 15,
	}

	config.Set(newConfig)
	assert.Equal(t, "custom-signature", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "new-key", config.SecretKey)
	assert.Equal(t, "md5", config.Algorithm)
	assert.Equal(t, time.Minute*15, config.TimeoutWindow)
}

func TestSignature_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestSignature_ChainedCalls(t *testing.T) {
	config := Default().
		WithSecretKey("production-secret").
		WithAlgorithm("sha512").
		WithTimeoutWindow(time.Minute * 3).
		WithSignatureHeader("X-API-Signature").
		WithTimestampHeader("X-API-Timestamp").
		WithNonceHeader("X-API-Nonce").
		AddIgnorePath("/api/v1/public").
		AddIgnorePath("/api/v1/status").
		AddRequiredHeader("X-API-Key").
		WithSkipQuery(true).
		WithSkipBody(false).
		Enable()

	assert.Equal(t, "production-secret", config.SecretKey)
	assert.Equal(t, "sha512", config.Algorithm)
	assert.Equal(t, time.Minute*3, config.TimeoutWindow)
	assert.Equal(t, "X-API-Signature", config.SignatureHeader)
	assert.Equal(t, "X-API-Timestamp", config.TimestampHeader)
	assert.Equal(t, "X-API-Nonce", config.NonceHeader)
	assert.Contains(t, config.IgnorePaths, "/api/v1/public")
	assert.Contains(t, config.IgnorePaths, "/api/v1/status")
	assert.Contains(t, config.RequiredHeaders, "X-API-Key")
	assert.True(t, config.SkipQuery)
	assert.False(t, config.SkipBody)
	assert.True(t, config.Enabled)
}
