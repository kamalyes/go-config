/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 11:08:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 15:29:58
 * @FilePath: \go-config\tests\email_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"encoding/json"
	"testing"

	"github.com/kamalyes/go-config/pkg/email"
)

// 公共的测试数据
var testEmail = email.NewEmail(
	"testModule",
	"recipient@example.com",
	"sender@example.com",
	"smtp.example.com",
	587,
	true,
	"secret-key",
)

// 公共的期望值映射
var expectedMap = map[string]interface{}{
	"moduleName": "testModule",
	"to":         "recipient@example.com",
	"from":       "sender@example.com",
	"host":       "smtp.example.com",
	"port":       587,
	"isSSL":      true,
	"secret":     "secret-key",
}

// 验证 Email 实例的字段
func validateEmail(t *testing.T, email *email.Email) {
	if email.ModuleName != testEmail.ModuleName {
		t.Errorf("expected ModuleName %s, got %s", testEmail.ModuleName, email.ModuleName)
	}
	if email.To != testEmail.To {
		t.Errorf("expected To %s, got %s", testEmail.To, email.To)
	}
	if email.From != testEmail.From {
		t.Errorf("expected From %s, got %s", testEmail.From, email.From)
	}
	if email.Host != testEmail.Host {
		t.Errorf("expected Host %s, got %s", testEmail.Host, email.Host)
	}
	if email.Port != testEmail.Port {
		t.Errorf("expected Port %d, got %d", testEmail.Port, email.Port)
	}
	if email.IsSSL != testEmail.IsSSL {
		t.Errorf("expected IsSSL %v, got %v", testEmail.IsSSL, email.IsSSL)
	}
	if email.Secret != testEmail.Secret {
		t.Errorf("expected Secret %s, got %s", testEmail.Secret, email.Secret)
	}
}

func TestEmailSerialization(t *testing.T) {
	// 将 Email 实例序列化为 JSON
	jsonData, err := json.Marshal(testEmail)
	if err != nil {
		t.Fatalf("failed to marshal email to JSON: %v", err)
	}

	// 反序列化 JSON 到新的 Email 实例
	var deserializedEmail email.Email
	err = json.Unmarshal(jsonData, &deserializedEmail)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON to email: %v", err)
	}

	// 验证反序列化后的 Email 实例与原始实例相同
	validateEmail(t, &deserializedEmail)
}

func TestEmailValidation(t *testing.T) {
	tests := []struct {
		email     email.Email
		expectErr bool
	}{
		{email.Email{To: "", From: "sender@example.com", Host: "smtp.example.com", Port: 587}, true},                       // To is empty
		{email.Email{To: "recipient@example.com", From: "", Host: "smtp.example.com", Port: 587}, true},                    // From is empty
		{email.Email{To: "recipient@example.com", From: "sender@example.com", Host: "", Port: 587}, true},                  // Host is empty
		{email.Email{To: "recipient@example.com", From: "sender@example.com", Host: "smtp.example.com", Port: 0}, true},    // Port is 0
		{email.Email{To: "recipient@example.com", From: "sender@example.com", Host: "smtp.example.com", Port: 587}, false}, // Valid email
	}

	for _, tt := range tests {
		t.Run(tt.email.To, func(t *testing.T) {
			err := tt.email.Validate()
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err != nil)
			}
		})
	}
}

func TestEmailToMap(t *testing.T) {
	resultMap := testEmail.ToMap()
	for key, expectedValue := range expectedMap {
		if resultValue, exists := resultMap[key]; !exists || resultValue != expectedValue {
			t.Errorf("expected %s: %v, got: %v", key, expectedValue, resultValue)
		}
	}
}

func TestEmailFromMap(t *testing.T) {
	email := &email.Email{}
	email.FromMap(expectedMap)

	validateEmail(t, email)
}

func TestEmailClone(t *testing.T) {
	clonedEmail := testEmail.Clone().(*email.Email)
	validateEmail(t, clonedEmail)
}

func TestEmailGet(t *testing.T) {
	result := testEmail.Clone().Get().(*email.Email)
	validateEmail(t, result)
}

func TestEmailSet(t *testing.T) {
	clonedEmail := testEmail.Clone().(*email.Email)
	newEmail := email.NewEmail("newModule", "newRecipient@example.com", "newSender@example.com", "new.smtp.example.com", 465, false, "new-secret-key")
	clonedEmail.Set(newEmail)

	// 验证设置后的 Email 实例
	validateEmail(t, testEmail)
}
