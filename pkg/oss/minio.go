/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 01:24:48
 * @FilePath: \go-config\pkg\oss\minio.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package oss

import (
	"github.com/kamalyes/go-config/internal"
)

// Minio 结构体用于配置 Minio 服务器的相关参数
type Minio struct {
	Endpoint            string `mapstructure:"endpoint"           yaml:"endpoint"        json:"endpoint"`                             // 地区
	AccessKey  string `mapstructure:"access-key"         yaml:"access-key"      json:"access_key"       validate:"required"`                 // 签名用的 key
	SecretKey  string `mapstructure:"secret-key"         yaml:"secret-key"      json:"secret_key"       validate:"required"`                 // 签名用的钥匙
	Bucket              string `mapstructure:"bucket"             yaml:"bucket"          json:"bucket"`                               // 桶
	ModuleName string `mapstructure:"modulename"         yaml:"modulename"      json:"module_name"`                                          // 模块名称
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
		Endpoint:   m.Endpoint,
		AccessKey:  m.AccessKey,
		SecretKey:  m.SecretKey,
		Bucket:    m.Bucket,
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
		m.Endpoint = configData.Endpoint
		m.AccessKey = configData.AccessKey
		m.SecretKey = configData.SecretKey
		m.Bucket = configData.Bucket
	}
}

// Validate 验证 Minio 配置的有效性
func (m *Minio) Validate() error {
	return internal.ValidateStruct(m)
}

// DefaultMinio 返回默认的 Minio 值
func DefaultMinio() Minio {
	return Minio{
		Endpoint:       "localhost",
		AccessKey:  "minioadmin",
		SecretKey:  "minioadmin",
		Bucket:    "my-bucket",
		ModuleName: "minio",
	}
}

// DefaultMinioConfig 返回默认的 Minio 指针，支持链式调用
func DefaultMinioConfig() *Minio {
	config := DefaultMinio()
	return &config
}


func (m *Minio) WithEndpoint(endpoint string) *Minio {
	m.Endpoint = endpoint
	return m
}

// WithAccessKey 设置AccessKey
func (m *Minio) WithAccessKey(accessKey string) *Minio {
	m.AccessKey = accessKey
	return m
}

// WithSecretKey 设置SecretKey
func (m *Minio) WithSecretKey(secretKey string) *Minio {
	m.SecretKey = secretKey
	return m
}

// WithModuleName 设置模块名称
func (m *Minio) WithModuleName(moduleName string) *Minio {
	m.ModuleName = moduleName
	return m
}

func (m *Minio) WithBucket(bucket string) *Minio {
	m.Bucket = bucket
	return m
}