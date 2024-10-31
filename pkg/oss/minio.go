/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 22:35:17
 * @FilePath: \go-config\oss\minio.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package oss

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Minio 配置文件
type Minio struct {
	/** 模块名称 */
	ModuleName string `mapstructure:"modulename" json:"moduleName" yaml:"modulename"`

	/** Host */
	Host string `mapstructure:"host" json:"host" yaml:"host"`

	/** 端口 */
	Port int `mapstructure:"port" json:"port" yaml:"port"`

	/** 签名用的 key */
	AccessKey string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`

	/** 签名用的钥匙 */
	SecretKey string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
}

// NewMinio 创建一个新的 Minio 实例
func NewMinio(moduleName, host string, port int, accessKey, secretKey string) *Minio {
	var minioInstance *Minio

	internal.LockFunc(func() {
		minioInstance = &Minio{
			ModuleName: moduleName,
			Host:       host,
			Port:       port,
			AccessKey:  accessKey,
			SecretKey:  secretKey,
		}
	})
	return minioInstance
}

// ToMap 将配置转换为映射
func (m *Minio) ToMap() map[string]interface{} {
	return internal.ToMap(m)
}

// FromMap 从映射中填充配置
func (m *Minio) FromMap(data map[string]interface{}) {
	internal.FromMap(m, data)
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
