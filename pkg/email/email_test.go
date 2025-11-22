/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-21 23:59:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-21 23:59:00
 * @FilePath: \go-config\pkg\email\email_test.go
 * @Description: Email配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package email

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmail_Default(t *testing.T) {
	email := Default()

	assert.NotNil(t, email)
	assert.Equal(t, "email", email.ModuleName)
	assert.Equal(t, "recipient@example.com", email.To)
	assert.Equal(t, "sender@example.com", email.From)
	assert.Equal(t, "smtp.gmail.com", email.Host)
	assert.Equal(t, 587, email.Port)
	assert.Equal(t, "email_app_password", email.Secret)
	assert.True(t, email.IsSSL)
}

func TestEmail_WithModuleName(t *testing.T) {
	email := Default().WithModuleName("custom_email")
	assert.Equal(t, "custom_email", email.ModuleName)
}

func TestEmail_WithTo(t *testing.T) {
	email := Default().WithTo("user@example.com")
	assert.Equal(t, "user@example.com", email.To)
}

func TestEmail_WithFrom(t *testing.T) {
	email := Default().WithFrom("sender@example.com")
	assert.Equal(t, "sender@example.com", email.From)
}

func TestEmail_WithHost(t *testing.T) {
	email := Default().WithHost("smtp.outlook.com")
	assert.Equal(t, "smtp.outlook.com", email.Host)
}

func TestEmail_WithPort(t *testing.T) {
	email := Default().WithPort(587)
	assert.Equal(t, 587, email.Port)
}

func TestEmail_WithSecret(t *testing.T) {
	email := Default().WithSecret("secret-password")
	assert.Equal(t, "secret-password", email.Secret)
}

func TestEmail_WithIsSSL(t *testing.T) {
	email := Default().WithIsSSL(false)
	assert.False(t, email.IsSSL)
}

func TestEmail_Clone(t *testing.T) {
	original := Default()
	original.To = "test@example.com"
	original.From = "sender@example.com"
	original.Port = 587

	cloned := original.Clone().(*Email)

	assert.Equal(t, original.To, cloned.To)
	assert.Equal(t, original.From, cloned.From)
	assert.Equal(t, original.Port, cloned.Port)

	cloned.Port = 25
	assert.Equal(t, 587, original.Port)
	assert.Equal(t, 25, cloned.Port)
}

func TestEmail_Get(t *testing.T) {
	email := Default()
	result := email.Get()

	assert.NotNil(t, result)
	resultEmail, ok := result.(*Email)
	assert.True(t, ok)
	assert.Equal(t, email, resultEmail)
}

func TestEmail_Set(t *testing.T) {
	email := Default()
	newEmail := &Email{
		ModuleName: "new_email",
		To:         "newuser@example.com",
		From:       "newsender@example.com",
		Host:       "smtp.yahoo.com",
		Port:       465,
		Secret:     "new-secret",
		IsSSL:      true,
	}

	email.Set(newEmail)

	assert.Equal(t, "new_email", email.ModuleName)
	assert.Equal(t, "newuser@example.com", email.To)
	assert.Equal(t, "newsender@example.com", email.From)
	assert.Equal(t, "smtp.yahoo.com", email.Host)
	assert.Equal(t, 465, email.Port)
	assert.Equal(t, "new-secret", email.Secret)
	assert.True(t, email.IsSSL)
}

func TestEmail_Validate(t *testing.T) {
	email := &Email{
		To:     "user@example.com",
		From:   "sender@example.com",
		Host:   "smtp.gmail.com",
		Port:   465,
		Secret: "password",
	}

	err := email.Validate()
	assert.NoError(t, err)
}

func TestEmail_Validate_Invalid(t *testing.T) {
	email := &Email{
		To:     "invalid-email",
		From:   "sender@example.com",
		Host:   "smtp.gmail.com",
		Port:   465,
		Secret: "password",
	}

	err := email.Validate()
	assert.Error(t, err)
}

func TestEmail_ChainedCalls(t *testing.T) {
	email := Default().
		WithModuleName("chained").
		WithTo("user@example.com").
		WithFrom("sender@example.com").
		WithHost("smtp.gmail.com").
		WithPort(465).
		WithSecret("password").
		WithIsSSL(true)

	assert.Equal(t, "chained", email.ModuleName)
	assert.Equal(t, "user@example.com", email.To)
	assert.Equal(t, "sender@example.com", email.From)
	assert.Equal(t, "smtp.gmail.com", email.Host)
	assert.Equal(t, 465, email.Port)
	assert.Equal(t, "password", email.Secret)
	assert.True(t, email.IsSSL)

	err := email.Validate()
	assert.NoError(t, err)
}

func TestNewEmail(t *testing.T) {
	opt := &Email{
		ModuleName: "test",
		To:         "test@example.com",
		From:       "sender@example.com",
		Host:       "smtp.test.com",
		Port:       587,
		Secret:     "secret",
		IsSSL:      false,
	}

	email := NewEmail(opt)
	assert.NotNil(t, email)
	assert.Equal(t, opt, email)
}
