/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\smtp\smtp_test.go
 * @Description: SMTP邮箱配置测试模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package smtp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmtp_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "smtp", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "127.0.0.1", config.SMTPHost)
	assert.Equal(t, 587, config.SMTPPort)
	assert.Equal(t, "smtp_user", config.Username)
	assert.Equal(t, "smtp_password", config.Password)
	assert.Equal(t, "noreply@example.com", config.FromAddress)
	assert.NotNil(t, config.ToAddresses)
	assert.False(t, config.EnableTLS)
	assert.NotNil(t, config.Headers)
	assert.Equal(t, 5, config.PoolSize)
}

func TestSmtp_WithModuleName(t *testing.T) {
	config := Default()
	result := config.WithModuleName("custom-smtp")
	assert.Equal(t, "custom-smtp", result.ModuleName)
	assert.Equal(t, config, result)
}

func TestSmtp_WithEnabled(t *testing.T) {
	config := Default()
	result := config.WithEnabled(false)
	assert.False(t, result.Enabled)
	assert.Equal(t, config, result)
}

func TestSmtp_WithSMTPHost(t *testing.T) {
	config := Default()
	result := config.WithSMTPHost("smtp.gmail.com")
	assert.Equal(t, "smtp.gmail.com", result.SMTPHost)
	assert.Equal(t, config, result)
}

func TestSmtp_WithSMTPPort(t *testing.T) {
	config := Default()
	result := config.WithSMTPPort(465)
	assert.Equal(t, 465, result.SMTPPort)
	assert.Equal(t, config, result)
}

func TestSmtp_WithUsername(t *testing.T) {
	config := Default()
	result := config.WithUsername("user@example.com")
	assert.Equal(t, "user@example.com", result.Username)
	assert.Equal(t, config, result)
}

func TestSmtp_WithPassword(t *testing.T) {
	config := Default()
	result := config.WithPassword("password123")
	assert.Equal(t, "password123", result.Password)
	assert.Equal(t, config, result)
}

func TestSmtp_WithFromAddress(t *testing.T) {
	config := Default()
	result := config.WithFromAddress("noreply@example.com")
	assert.Equal(t, "noreply@example.com", result.FromAddress)
	assert.Equal(t, config, result)
}

func TestSmtp_WithPoolSize(t *testing.T) {
	config := Default()
	result := config.WithPoolSize(10)
	assert.Equal(t, 10, result.PoolSize)
	assert.Equal(t, config, result)
}

func TestSmtp_AddToAddress(t *testing.T) {
	config := Default()
	result := config.AddToAddress("recipient@example.com")
	assert.Contains(t, result.ToAddresses, "recipient@example.com")
	assert.Equal(t, config, result)
}

func TestSmtp_EnableTLSService(t *testing.T) {
	config := Default()
	config.EnableTLS = false
	result := config.EnableTLSService()
	assert.True(t, result.EnableTLS)
	assert.Equal(t, config, result)
}

func TestSmtp_DisableTLS(t *testing.T) {
	config := Default()
	result := config.DisableTLS()
	assert.False(t, result.EnableTLS)
	assert.Equal(t, config, result)
}

func TestSmtp_AddHeader(t *testing.T) {
	config := Default()
	result := config.AddHeader("X-Custom", "value")
	assert.Equal(t, "value", result.Headers["X-Custom"])
	assert.Equal(t, config, result)
}

func TestSmtp_Clone(t *testing.T) {
	config := Default()
	config.WithSMTPHost("smtp.example.com").
		AddToAddress("user1@example.com").
		AddToAddress("user2@example.com").
		AddHeader("X-Priority", "high")

	clone := config.Clone()
	assert.NotNil(t, clone)

	clonedConfig, ok := clone.(*Smtp)
	assert.True(t, ok)
	assert.Equal(t, config.ModuleName, clonedConfig.ModuleName)
	assert.Equal(t, config.SMTPHost, clonedConfig.SMTPHost)
	assert.Equal(t, config.SMTPPort, clonedConfig.SMTPPort)

	// 验证深拷贝 - 切片
	clonedConfig.ToAddresses = append(clonedConfig.ToAddresses, "user3@example.com")
	assert.NotEqual(t, len(config.ToAddresses), len(clonedConfig.ToAddresses))

	// 验证深拷贝 - map
	clonedConfig.Headers["X-New"] = "new-value"
	_, exists := config.Headers["X-New"]
	assert.False(t, exists)
}

func TestSmtp_Get(t *testing.T) {
	config := Default()
	result := config.Get()
	assert.NotNil(t, result)
	assert.Equal(t, config, result)
}

func TestSmtp_Set(t *testing.T) {
	config := Default()
	newConfig := &Smtp{
		ModuleName:  "new-smtp",
		Enabled:     false,
		SMTPHost:    "smtp.new.com",
		SMTPPort:    465,
		Username:    "newuser",
		Password:    "newpass",
		FromAddress: "new@example.com",
		ToAddresses: []string{"recipient@example.com"},
		EnableTLS:   false,
		PoolSize:    10,
	}

	config.Set(newConfig)
	assert.Equal(t, "new-smtp", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "smtp.new.com", config.SMTPHost)
	assert.Equal(t, 465, config.SMTPPort)
	assert.Equal(t, "newuser", config.Username)
	assert.Equal(t, "newpass", config.Password)
	assert.Equal(t, "new@example.com", config.FromAddress)
	assert.False(t, config.EnableTLS)
	assert.Equal(t, 10, config.PoolSize)
}

func TestSmtp_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestSmtp_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("my-smtp").
		WithSMTPHost("smtp.gmail.com").
		WithSMTPPort(465).
		WithUsername("myuser@gmail.com").
		WithPassword("mypassword").
		WithFromAddress("noreply@company.com").
		WithPoolSize(15).
		AddToAddress("admin@company.com").
		AddToAddress("support@company.com").
		AddHeader("X-Priority", "high").
		AddHeader("X-Source", "notification-system").
		EnableTLSService()

	assert.Equal(t, "my-smtp", config.ModuleName)
	assert.Equal(t, "smtp.gmail.com", config.SMTPHost)
	assert.Equal(t, 465, config.SMTPPort)
	assert.Equal(t, "myuser@gmail.com", config.Username)
	assert.Equal(t, "mypassword", config.Password)
	assert.Equal(t, "noreply@company.com", config.FromAddress)
	assert.Equal(t, 15, config.PoolSize)
	assert.Contains(t, config.ToAddresses, "admin@company.com")
	assert.Contains(t, config.ToAddresses, "support@company.com")
	assert.Equal(t, "high", config.Headers["X-Priority"])
	assert.Equal(t, "notification-system", config.Headers["X-Source"])
	assert.True(t, config.EnableTLS)
}
