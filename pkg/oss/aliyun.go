/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:39:54
 * @FilePath: \go-config\pkg\oss\aliyun.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package oss

import (
	"github.com/kamalyes/go-config/internal"
)

// AliyunOss 结构体用于配置 AliyunOss 服务器的相关参数
type AliyunOss struct {
	ModuleName          string `mapstructure:"modulename"         yaml:"modulename"      json:"module_name"      validate:"required"` // 模块名称
	AccessKey           string `mapstructure:"access-key"         yaml:"access-key"      json:"access_key"       validate:"required"` // 签名用的 key
	SecretKey           string `mapstructure:"secret-key"         yaml:"secret-key"      json:"secret_key"       validate:"required"` // 签名用的钥匙
	Endpoint            string `mapstructure:"endpoint"           yaml:"endpoint"        json:"endpoint"`                             // 地区
	Bucket              string `mapstructure:"bucket"             yaml:"bucket"          json:"bucket"`                               // 桶
	ReplaceOriginalHost string `mapstructure:"replace-original-host" yaml:"replace-original-host" json:"replace_original_host"`       // 替换的原始主机
	ReplaceLaterHost    string `mapstructure:"replace-later-host"   yaml:"replace-later-host"   json:"replace_later_host"`            // 替换后的主机
}

// NewMinio 创建一个新的 AliyunOss 实例
func NewAliyunOss(opt *AliyunOss) *AliyunOss {
	var minioInstance *AliyunOss

	internal.LockFunc(func() {
		minioInstance = opt
	})
	return minioInstance
}

// Clone 返回 AliyunOss 配置的副本
func (m *AliyunOss) Clone() internal.Configurable {
	return &AliyunOss{
		ModuleName:          m.ModuleName,
		AccessKey:           m.AccessKey,
		SecretKey:           m.SecretKey,
		Endpoint:            m.Endpoint,
		Bucket:              m.Bucket,
		ReplaceOriginalHost: m.ReplaceOriginalHost,
		ReplaceLaterHost:    m.ReplaceLaterHost,
	}
}

// Get 返回 AliyunOss 配置的所有字段
func (m *AliyunOss) Get() interface{} {
	return m
}

// Set 更新 AliyunOss 配置的字段
func (m *AliyunOss) Set(data interface{}) {
	if configData, ok := data.(*AliyunOss); ok {
		m.ModuleName = configData.ModuleName
		m.Endpoint = configData.Endpoint
		m.Bucket = configData.Bucket
		m.AccessKey = configData.AccessKey
		m.SecretKey = configData.SecretKey
		m.ReplaceOriginalHost = configData.ReplaceOriginalHost
		m.ReplaceLaterHost = configData.ReplaceLaterHost
	}
}

// Validate 验证 AliyunOss 配置的有效性
func (m *AliyunOss) Validate() error {
	return internal.ValidateStruct(m)
}
