/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:50:03
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
	Host       string `mapstructure:"host"               yaml:"host"            json:"host"             validate:"required"`                 // 主机
	Port       int    `mapstructure:"port"               yaml:"port"            json:"port"             validate:"required,min=1,max=65535"` // 端口，范围 1-65535
	AccessKey  string `mapstructure:"access-key"         yaml:"access-key"      json:"access_key"       validate:"required"`                 // 签名用的 key
	SecretKey  string `mapstructure:"secret-key"         yaml:"secret-key"      json:"secret_key"       validate:"required"`                 // 签名用的钥匙
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
	return internal.ValidateStruct(m)
}
