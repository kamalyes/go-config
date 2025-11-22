/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-21 23:59:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 16:51:46
 * @FilePath: \go-config\pkg\etcd\etcd_test.go
 * @Description: Etcd配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package etcd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEtcd_Default(t *testing.T) {
	etcd := Default()

	assert.NotNil(t, etcd)
	assert.Equal(t, "etcd", etcd.ModuleName)
	assert.Equal(t, []string{"127.0.0.1:2379"}, etcd.Hosts)
	assert.Equal(t, "app/config", etcd.Key)
	assert.Equal(t, int64(1), etcd.ID)
	assert.Equal(t, "etcd_user", etcd.User)
	assert.Equal(t, "etcd_password", etcd.Pass)
	assert.True(t, etcd.InsecureSkipVerify)
	assert.Equal(t, 5, etcd.DialTimeout)
	assert.Equal(t, 10, etcd.RequestTimeout)
}

func TestEtcd_WithModuleName(t *testing.T) {
	etcd := Default().WithModuleName("custom_etcd")
	assert.Equal(t, "custom_etcd", etcd.ModuleName)
}

func TestEtcd_WithHosts(t *testing.T) {
	hosts := []string{"host1:2379", "host2:2379"}
	etcd := Default().WithHosts(hosts)
	assert.Equal(t, hosts, etcd.Hosts)
}

func TestEtcd_WithKey(t *testing.T) {
	etcd := Default().WithKey("/app/config")
	assert.Equal(t, "/app/config", etcd.Key)
}

func TestEtcd_WithID(t *testing.T) {
	etcd := Default().WithID(123)
	assert.Equal(t, int64(123), etcd.ID)
}

func TestEtcd_WithUser(t *testing.T) {
	etcd := Default().WithUser("admin")
	assert.Equal(t, "admin", etcd.User)
}

func TestEtcd_WithPass(t *testing.T) {
	etcd := Default().WithPass("password")
	assert.Equal(t, "password", etcd.Pass)
}

func TestEtcd_WithCertFile(t *testing.T) {
	etcd := Default().WithCertFile("/path/to/cert.pem")
	assert.Equal(t, "/path/to/cert.pem", etcd.CertFile)
}

func TestEtcd_WithDialTimeout(t *testing.T) {
	etcd := Default().WithDialTimeout(10)
	assert.Equal(t, 10, etcd.DialTimeout)
}

func TestEtcd_WithRequestTimeout(t *testing.T) {
	etcd := Default().WithRequestTimeout(5)
	assert.Equal(t, 5, etcd.RequestTimeout)
}

func TestEtcd_Clone(t *testing.T) {
	original := Default()
	original.Hosts = []string{"host1:2379", "host2:2379"}
	original.User = "admin"
	original.Pass = "password"

	cloned := original.Clone().(*Etcd)

	assert.Equal(t, original.Hosts, cloned.Hosts)
	assert.Equal(t, original.User, cloned.User)
	assert.Equal(t, original.Pass, cloned.Pass)

	// 验证切片独立性
	cloned.Hosts[0] = "host3:2379"
	assert.Equal(t, "host1:2379", original.Hosts[0])
	assert.Equal(t, "host3:2379", cloned.Hosts[0])
}

func TestEtcd_Get(t *testing.T) {
	etcd := Default()
	result := etcd.Get()

	assert.NotNil(t, result)
	resultEtcd, ok := result.(*Etcd)
	assert.True(t, ok)
	assert.Equal(t, etcd, resultEtcd)
}

func TestEtcd_Set(t *testing.T) {
	etcd := Default()
	newEtcd := &Etcd{
		ModuleName:     "new_etcd",
		Hosts:          []string{"new:2379"},
		Key:            "/new/key",
		ID:             456,
		User:           "newuser",
		Pass:           "newpass",
		DialTimeout:    10,
		RequestTimeout: 5,
	}

	etcd.Set(newEtcd)

	assert.Equal(t, "new_etcd", etcd.ModuleName)
	assert.Equal(t, []string{"new:2379"}, etcd.Hosts)
	assert.Equal(t, "/new/key", etcd.Key)
	assert.Equal(t, int64(456), etcd.ID)
	assert.Equal(t, "newuser", etcd.User)
	assert.Equal(t, "newpass", etcd.Pass)
	assert.Equal(t, 10, etcd.DialTimeout)
	assert.Equal(t, 5, etcd.RequestTimeout)
}

func TestEtcd_Validate(t *testing.T) {
	etcd := Default()
	err := etcd.Validate()
	assert.NoError(t, err)
}

func TestEtcd_ChainedCalls(t *testing.T) {
	etcd := Default().
		WithModuleName("chained").
		WithHosts([]string{"host1:2379", "host2:2379"}).
		WithKey("/chained/config").
		WithID(789).
		WithUser("chainuser").
		WithPass("chainpass").
		WithDialTimeout(15).
		WithRequestTimeout(10)

	assert.Equal(t, "chained", etcd.ModuleName)
	assert.Equal(t, []string{"host1:2379", "host2:2379"}, etcd.Hosts)
	assert.Equal(t, "/chained/config", etcd.Key)
	assert.Equal(t, int64(789), etcd.ID)
	assert.Equal(t, "chainuser", etcd.User)
	assert.Equal(t, "chainpass", etcd.Pass)
	assert.Equal(t, 15, etcd.DialTimeout)
	assert.Equal(t, 10, etcd.RequestTimeout)

	err := etcd.Validate()
	assert.NoError(t, err)
}

func TestNewEtcd(t *testing.T) {
	opt := &Etcd{
		ModuleName: "test",
		Hosts:      []string{"test:2379"},
		Key:        "/test/key",
		ID:         999,
	}

	etcd := NewEtcd(opt)
	assert.NotNil(t, etcd)
	assert.Equal(t, opt, etcd)
}
