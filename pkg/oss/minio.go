/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-13 01:25:18
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
	Endpoint   string `mapstructure:"endpoint" yaml:"endpoint" json:"endpoint"`                            // 地区
	AccessKey  string `mapstructure:"access_key" yaml:"access-key" json:"access_key"  validate:"required"` // 签名用的 key
	SecretKey  string `mapstructure:"secret_key" yaml:"secret-key" json:"secret_key"  validate:"required"` // 签名用的钥匙
	Bucket     string `mapstructure:"bucket" yaml:"bucket" json:"bucket"`                                  // 桶
	ModuleName string `mapstructure:"module_name" yaml:"module_name" json:"module_name"`                   // 模块名称
	UseSSL     bool   `mapstructure:"use_ssl" yaml:"use-ssl" json:"use_ssl"`                               // 是否使用 HTTPS
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
		Bucket:     m.Bucket,
		UseSSL:     m.UseSSL,
	}
}

// OSSProvider 接口实现

// GetProviderType 获取提供商类型
func (m *Minio) GetProviderType() OSSType {
	return OSSTypeMinio
}

// GetEndpoint 获取端点地址
func (m *Minio) GetEndpoint() string {
	return m.Endpoint
}

// GetAccessKey 获取访问密钥
func (m *Minio) GetAccessKey() string {
	return m.AccessKey
}

// GetSecretKey 获取私密密钥
func (m *Minio) GetSecretKey() string {
	return m.SecretKey
}

// GetBucket 获取存储桶名称
func (m *Minio) GetBucket() string {
	return m.Bucket
}

// IsSSL 是否使用SSL
func (m *Minio) IsSSL() bool {
	return m.UseSSL
}

// GetModuleName 获取模块名称
func (m *Minio) GetModuleName() string {
	return m.ModuleName
}

// SetCredentials 设置凭证
func (m *Minio) SetCredentials(accessKey, secretKey string) {
	m.AccessKey = accessKey
	m.SecretKey = secretKey
}

// SetEndpoint 设置端点
func (m *Minio) SetEndpoint(endpoint string) {
	m.Endpoint = endpoint
}

// SetBucket 设置存储桶
func (m *Minio) SetBucket(bucket string) {
	m.Bucket = bucket
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
		Endpoint:   "localhost",
		AccessKey:  "minioadmin",
		SecretKey:  "minioadmin",
		Bucket:     "my-bucket",
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
