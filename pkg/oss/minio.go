/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:31:58
 * @FilePath: \go-config\pkg\oss\minio.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package oss

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Minio 结构体用于配置 Minio 服务器的相关参数
type Minio struct {
	ModuleName string `mapstructure:"MODULE_NAME"              yaml:"modulename"` // 模块名称
	Host       string `mapstructure:"HOST"                     yaml:"host"`       // 主机
	Port       int    `mapstructure:"PORT"                     yaml:"port"`       // 端口
	AccessKey  string `mapstructure:"ACCESS_KEY"               yaml:"access-key"` // 签名用的 key
	SecretKey  string `mapstructure:"SECRET_KEY"               yaml:"secret-key"` // 签名用的钥匙
}

// NewMinio 创建一个新的 Minio 实例
func NewMinio(opt *Minio) *Minio {
	var minioInstance *Minio

	internal.LockFunc(func() {
		minioInstance = opt
	})
	return minioInstance
}

// Clone 返回 Minio 配置的副本
func (m *Minio) Clone() internal.Configurable {
	return &Minio{
		ModuleName: m.ModuleName,
		Host:       m.Host,
		Port:       m.Port,
		AccessKey:  m.AccessKey,
		SecretKey:  m.SecretKey,
	}
}

// Get 返回 Minio 配置的所有字段
func (m *Minio) Get() interface{} {
	return m
}

// Set 更新 Minio 配置的字段
func (m *Minio) Set(data interface{}) {
	if configData, ok := data.(*Minio); ok {
		m.ModuleName = configData.ModuleName
		m.Host = configData.Host
		m.Port = configData.Port
		m.AccessKey = configData.AccessKey
		m.SecretKey = configData.SecretKey
	}
}

// Validate 验证 Minio 配置的有效性
func (m *Minio) Validate() error {
	if m.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if m.Host == "" {
		return errors.New("host cannot be empty")
	}
	if m.Port <= 0 {
		return errors.New("port must be greater than 0")
	}
	if m.AccessKey == "" {
		return errors.New("access key cannot be empty")
	}
	if m.SecretKey == "" {
		return errors.New("secret key cannot be empty")
	}
	return nil
}
