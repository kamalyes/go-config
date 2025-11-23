/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 01:19:56
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
	AccessKey           string `mapstructure:"access-key" yaml:"access-key" json:"accessKey"       validate:"required"`       // 签名用的 key
	SecretKey           string `mapstructure:"secret-key" yaml:"secret-key" json:"secretKey"       validate:"required"`       // 签名用的钥匙
	Endpoint            string `mapstructure:"endpoint" yaml:"endpoint" json:"endpoint"`                                      // 地区
	Bucket              string `mapstructure:"bucket" yaml:"bucket" json:"bucket"`                                            // 桶
	Region              string `mapstructure:"region" yaml:"region" json:"region"         validate:"required"`                // AWS 区域，如：ap-southeast-1
	ReplaceOriginalHost string `mapstructure:"replace-original-host" yaml:"replace-original-host" json:"replaceOriginalHost"` // 替换的原始主机
	ReplaceLaterHost    string `mapstructure:"replace-later-host" yaml:"replace-later-host" json:"replaceLaterHost"`          // 替换后的主机
	ModuleName          string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                              // 模块名称
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
		Region:              m.Region,
		ReplaceOriginalHost: m.ReplaceOriginalHost,
		ReplaceLaterHost:    m.ReplaceLaterHost,
	}
}

// OSSProvider 接口实现

// GetProviderType 获取提供商类型
func (a *AliyunOss) GetProviderType() OSSType {
	return OSSTypeAliyunOSS
}

// GetEndpoint 获取端点地址
func (a *AliyunOss) GetEndpoint() string {
	return a.Endpoint
}

// GetAccessKey 获取访问密钥
func (a *AliyunOss) GetAccessKey() string {
	return a.AccessKey
}

// GetSecretKey 获取私密密钥
func (a *AliyunOss) GetSecretKey() string {
	return a.SecretKey
}

// GetBucket 获取存储桶名称
func (a *AliyunOss) GetBucket() string {
	return a.Bucket
}

// IsSSL 是否使用SSL (阿里云OSS默认使用HTTPS)
func (a *AliyunOss) IsSSL() bool {
	return true // 阿里云OSS默认使用HTTPS
}

// GetModuleName 获取模块名称
func (a *AliyunOss) GetModuleName() string {
	return a.ModuleName
}

// SetCredentials 设置凭证
func (a *AliyunOss) SetCredentials(accessKey, secretKey string) {
	a.AccessKey = accessKey
	a.SecretKey = secretKey
}

// SetEndpoint 设置端点
func (a *AliyunOss) SetEndpoint(endpoint string) {
	a.Endpoint = endpoint
}

// SetBucket 设置存储桶
func (a *AliyunOss) SetBucket(bucket string) {
	a.Bucket = bucket
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
		m.Region = configData.Region
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

// DefaultAliyunOss 返回默认AliyunOss配置
func DefaultAliyunOss() AliyunOss {
	return AliyunOss{
		ModuleName:          "aliyun-oss",
		AccessKey:           "demo_access_key",
		SecretKey:           "demo_secret_key",
		Endpoint:            "oss-cn-hangzhou.aliyuncs.com",
		Region:              "oss-cn-hangzhou",
		Bucket:              "demo-bucket",
		ReplaceOriginalHost: "oss-cn-hangzhou.aliyuncs.com",
		ReplaceLaterHost:    "cdn.example.com",
	}
}

// DefaultAliyunOSSConfig 返回默认AliyunOSS配置的指针，支持链式调用
func DefaultAliyunOSSConfig() *AliyunOss {
	config := DefaultAliyunOss()
	return &config
}

// WithModuleName 设置模块名称
func (m *AliyunOss) WithModuleName(moduleName string) *AliyunOss {
	m.ModuleName = moduleName
	return m
}

// WithAccessKey 设置访问密钥
func (m *AliyunOss) WithAccessKey(accessKey string) *AliyunOss {
	m.AccessKey = accessKey
	return m
}

// WithSecretKey 设置私有密钥
func (m *AliyunOss) WithSecretKey(secretKey string) *AliyunOss {
	m.SecretKey = secretKey
	return m
}

// WithEndpoint 设置OSS端点
func (m *AliyunOss) WithEndpoint(endpoint string) *AliyunOss {
	m.Endpoint = endpoint
	return m
}

// WithBucket 设置存储桶名称
func (m *AliyunOss) WithBucket(bucket string) *AliyunOss {
	m.Bucket = bucket
	return m
}

// WithRegion 设置区域
func (a *AliyunOss) WithRegion(region string) *AliyunOss {
	a.Region = region
	return a
}

// WithReplaceOriginalHost 设置原始主机替换
func (m *AliyunOss) WithReplaceOriginalHost(replaceOriginalHost string) *AliyunOss {
	m.ReplaceOriginalHost = replaceOriginalHost
	return m
}

// WithReplaceLaterHost 设置替换后的主机
func (m *AliyunOss) WithReplaceLaterHost(replaceLaterHost string) *AliyunOss {
	m.ReplaceLaterHost = replaceLaterHost
	return m
}
