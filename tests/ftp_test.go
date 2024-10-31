/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-01 11:16:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 15:55:16
 * @FilePath: \go-config\tests\ftp_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/ftp"
)

var testFtp = ftp.NewFtp("testModule", "127.0.0.1:21", "user", "pass", "/path")

func TestFtp_Validate(t *testing.T) {
	tests := []struct {
		name     string
		ftp      *ftp.Ftp
		expected bool
	}{
		{"Valid Ftp", testFtp.Clone().(*ftp.Ftp), true},
		{"Empty Addr", ftp.NewFtp("testModule", "", "user", "pass", "/path"), false},
		{"Empty Username", ftp.NewFtp("testModule", "127.0.0.1:21", "", "pass", "/path"), false},
		{"Empty Password", ftp.NewFtp("testModule", "127.0.0.1:21", "user", "", "/path"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ftp.Validate()
			if (err == nil) != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, err == nil)
			}
		})
	}
}

func TestFtp_Clone(t *testing.T) {
	clonedFtp := testFtp.Clone().(*ftp.Ftp)

	if testFtp.ModuleName != clonedFtp.ModuleName || testFtp.Addr != clonedFtp.Addr ||
		testFtp.Username != clonedFtp.Username || testFtp.Password != clonedFtp.Password ||
		testFtp.Cwd != clonedFtp.Cwd {
		t.Errorf("Clone method did not copy fields correctly")
	}
}

func TestFtp_ToMap(t *testing.T) {
	expectedMap := map[string]interface{}{
		"moduleName": "testModule",
		"addr":       "127.0.0.1:21",
		"username":   "user",
		"password":   "pass",
		"cwd":        "/path",
	}

	clonedFtp := testFtp.Clone().(*ftp.Ftp).ToMap()
	for key, expectedValue := range expectedMap {
		if clonedFtp[key] != expectedValue {
			t.Errorf("ToMap() for key %s: expected %v, got %v", key, expectedValue, clonedFtp[key])
		}
	}
}

func TestFtp_FromMap(t *testing.T) {
	data := map[string]interface{}{
		"moduleName": "newModule",
		"addr":       "192.168.0.1:21",
		"username":   "newUser",
		"password":   "newPass",
		"cwd":        "/newpath",
	}

	clonedFtp := testFtp.Clone().(*ftp.Ftp)
	clonedFtp.FromMap(data)

	if clonedFtp.ModuleName != "newModule" || clonedFtp.Addr != "192.168.0.1:21" ||
		clonedFtp.Username != "newUser" || clonedFtp.Password != "newPass" ||
		clonedFtp.Cwd != "/newpath" {
		t.Errorf("FromMap() did not set fields correctly")
	}
}

func TestFtp_Set(t *testing.T) {
	clonedFtp := testFtp.Clone().(*ftp.Ftp) // 克隆原始 FTP 实例
	newFtp := ftp.NewFtp("anotherModule", "10.0.0.1:21", "anotherUser", "anotherPass", "/anotherpath")
	clonedFtp.Set(newFtp) // 设置新 FTP 实例

	// 验证克隆的 FTP 实例是否被正确更新
	if clonedFtp.ModuleName != "anotherModule" || clonedFtp.Addr != "10.0.0.1:21" ||
		clonedFtp.Username != "anotherUser" || clonedFtp.Password != "anotherPass" ||
		clonedFtp.Cwd != "/anotherpath" {
		t.Errorf("Set() did not update fields correctly")
	}

	// 验证原始 FTP 实例是否未受到影响
	if testFtp.ModuleName != "testModule" || testFtp.Addr != "127.0.0.1:21" ||
		testFtp.Username != "user" || testFtp.Password != "pass" ||
		testFtp.Cwd != "/path" {
		t.Errorf("Original FTP instance should not be modified")
	}
}
