/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-13 11:32:55
 * @FilePath: \go-config\pkg\oss\oss.go
 * @Description: OSS对象存储统一配置管理
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package oss

import (
	"fmt"
	"strings"

	"github.com/kamalyes/go-config/internal"
)

// OSSType 定义OSS存储类型
type OSSType string

const (
	OSSTypeMinio     OSSType = "minio"
	OSSTypeS3        OSSType = "s3"
	OSSTypeAliyunOSS OSSType = "aliyun"
	OSSTypeBoltDB    OSSType = "boltdb"
)

// OSSProvider OSS提供商接口，统一不同OSS服务的操作
type OSSProvider interface {
	internal.Configurable

	// GetProviderType 获取提供商类型
	GetProviderType() OSSType

	// GetEndpoint 获取端点地址
	GetEndpoint() string

	// GetAccessKey 获取访问密钥
	GetAccessKey() string

	// GetSecretKey 获取私密密钥
	GetSecretKey() string

	// GetBucket 获取存储桶名称
	GetBucket() string

	// IsSSL 是否使用SSL
	IsSSL() bool

	// GetModuleName 获取模块名称
	GetModuleName() string

	// SetCredentials 设置凭证
	SetCredentials(accessKey, secretKey string)

	// SetEndpoint 设置端点
	SetEndpoint(endpoint string)

	// SetBucket 设置存储桶
	SetBucket(bucket string)
}

// OSSConfig OSS统一配置结构
type OSSConfig struct {
	Type      OSSType    `mapstructure:"type" yaml:"type" json:"type"`          // OSS类型
	Enabled   bool       `mapstructure:"enabled" yaml:"enabled" json:"enabled"` // 是否启用
	Minio     *Minio     `mapstructure:"minio" yaml:"minio" json:"minio"`       // Minio配置
	S3        *S3        `mapstructure:"s3" yaml:"s3" json:"s3"`                // AWS S3配置
	AliyunOSS *AliyunOss `mapstructure:"aliyun" yaml:"aliyun" json:"aliyun"`    // 阿里云OSS配置
	BoltDB    *BoltDB    `mapstructure:"boltdb" yaml:"boltdb" json:"boltdb"`    // BoltDB本地存储配置
}

// NewOSSConfig 创建新的OSS配置
func NewOSSConfig() *OSSConfig {
	return &OSSConfig{
		Type:      OSSTypeMinio,
		Enabled:   false,
		Minio:     DefaultMinioConfig(),
		S3:        DefaultS3Config(),
		AliyunOSS: DefaultAliyunOSSConfig(),
		BoltDB:    DefaultBoltDB(),
	}
}

// GetProvider 获取指定类型的OSS提供商
func (c *OSSConfig) GetProvider(ossType OSSType) (OSSProvider, error) {
	switch ossType {
	case OSSTypeMinio:
		if c.Minio == nil {
			return nil, fmt.Errorf("minio config not found")
		}
		return c.Minio, nil
	case OSSTypeS3:
		if c.S3 == nil {
			return nil, fmt.Errorf("s3 config not found")
		}
		return c.S3, nil
	case OSSTypeAliyunOSS:
		if c.AliyunOSS == nil {
			return nil, fmt.Errorf("aliyun oss config not found")
		}
		return c.AliyunOSS, nil
	case OSSTypeBoltDB:
		if c.BoltDB == nil {
			return nil, fmt.Errorf("boltdb config not found")
		}
		return c.BoltDB, nil
	default:
		return nil, fmt.Errorf("unsupported oss type: %s", ossType)
	}
}

// ListAvailableProviders 列出所有可用的OSS提供商
func (c *OSSConfig) ListAvailableProviders() []OSSType {
	var providers []OSSType

	if c.Minio != nil {
		providers = append(providers, OSSTypeMinio)
	}
	if c.S3 != nil {
		providers = append(providers, OSSTypeS3)
	}
	if c.AliyunOSS != nil {
		providers = append(providers, OSSTypeAliyunOSS)
	}
	if c.BoltDB != nil {
		providers = append(providers, OSSTypeBoltDB)
	}

	return providers
}

// ValidateProvider 验证指定提供商的配置
func (c *OSSConfig) ValidateProvider(ossType OSSType) error {
	provider, err := c.GetProvider(ossType)
	if err != nil {
		return err
	}

	return provider.Validate()
}

// ValidateAll 验证所有配置的提供商
func (c *OSSConfig) ValidateAll() error {
	providers := c.ListAvailableProviders()

	for _, providerType := range providers {
		if err := c.ValidateProvider(providerType); err != nil {
			return fmt.Errorf("validation failed for %s: %w", providerType, err)
		}
	}

	return nil
}

// Clone 返回OSS配置的副本
func (c *OSSConfig) Clone() internal.Configurable {
	newConfig := &OSSConfig{
		Type:    c.Type,
		Enabled: c.Enabled,
	}

	if c.Minio != nil {
		newConfig.Minio = c.Minio.Clone().(*Minio)
	}
	if c.S3 != nil {
		newConfig.S3 = c.S3.Clone().(*S3)
	}
	if c.AliyunOSS != nil {
		newConfig.AliyunOSS = c.AliyunOSS.Clone().(*AliyunOss)
	}

	return newConfig
}

// Get 返回OSS配置
func (c *OSSConfig) Get() interface{} {
	return c
}

// Set 设置OSS配置
func (c *OSSConfig) Set(data interface{}) {
	if config, ok := data.(*OSSConfig); ok {
		*c = *config
	}
}

// Validate 验证OSS配置的有效性
func (c *OSSConfig) Validate() error {
	if !c.Enabled {
		return nil
	}

	return nil
}

// EnsureDefaults 确保OSS配置的默认值
func (c *OSSConfig) EnsureDefaults() {
	if c.Minio == nil {
		c.Minio = DefaultMinioConfig()
	}
	if c.S3 == nil {
		c.S3 = DefaultS3Config()
	}
	if c.AliyunOSS == nil {
		c.AliyunOSS = DefaultAliyunOSSConfig()
	}
	// 如果没有设置默认类型，使用MinIO
	if c.Type == "" {
		c.Type = OSSTypeMinio
	}
}

// GetSafe 安全获取配置，确保所有字段都有默认值
func (c *OSSConfig) GetSafe() interface{} {
	if c == nil {
		return DefaultOSSConfig()
	}
	c.EnsureDefaults()
	return c
}

// BeforeLoad 配置加载前的钩子
func (c *OSSConfig) BeforeLoad() error {
	c.EnsureDefaults()
	return nil
}

// AfterLoad 配置加载后的钩子
func (c *OSSConfig) AfterLoad() error {
	c.EnsureDefaults()
	return nil
}

// GetProviderName 获取提供商显示名称
func GetProviderName(ossType OSSType) string {
	switch ossType {
	case OSSTypeMinio:
		return "MinIO"
	case OSSTypeS3:
		return "Amazon S3"
	case OSSTypeAliyunOSS:
		return "Alibaba Cloud OSS"
	default:
		return string(ossType)
	}
}

// ParseOSSType 解析OSS类型字符串
func ParseOSSType(typeStr string) (OSSType, error) {
	switch strings.ToLower(typeStr) {
	case "minio":
		return OSSTypeMinio, nil
	case "s3", "aws", "amazon":
		return OSSTypeS3, nil
	case "aliyun", "alicloud", "oss":
		return OSSTypeAliyunOSS, nil
	default:
		return "", fmt.Errorf("unsupported oss type: %s", typeStr)
	}
}

// DefaultOSSConfig 返回默认的OSS配置
func DefaultOSSConfig() *OSSConfig {
	return NewOSSConfig()
}

// GetSupportedTypes 获取所有支持的OSS类型
func GetSupportedTypes() []OSSType {
	return []OSSType{
		OSSTypeMinio,
		OSSTypeS3,
		OSSTypeAliyunOSS,
	}
}
