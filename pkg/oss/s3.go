/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-08 12:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-08 12:00:00
 * @FilePath: \go-config\pkg\oss\s3.go
 * @Description: AWS S3 对象存储服务配置
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package oss

import (
	"github.com/kamalyes/go-config/internal"
)

// S3 结构体用于配置 AWS S3 服务器的相关参数
type S3 struct {
	Endpoint     string `mapstructure:"endpoint"      yaml:"endpoint"      json:"endpoint"       validate:"required"` // S3 端点地址，如：https://s3.ap-southeast-1.amazonaws.com
	Region       string `mapstructure:"region"        yaml:"region"        json:"region"         validate:"required"` // AWS 区域，如：ap-southeast-1
	AccessKey    string `mapstructure:"access-key"    yaml:"access-key"    json:"access_key"     validate:"required"` // AWS Access Key ID
	SecretKey    string `mapstructure:"secret-key"    yaml:"secret-key"    json:"secret_key"     validate:"required"` // AWS Secret Access Key
	BucketPrefix string `mapstructure:"bucket-prefix" yaml:"bucket-prefix" json:"bucket_prefix"`                      // 存储桶前缀，如：aicsqa
	SessionToken string `mapstructure:"session-token" yaml:"session-token" json:"session_token"`                      // 会话令牌（用于临时凭证）
	UseSSL       bool   `mapstructure:"use-ssl"       yaml:"use-ssl"       json:"use_ssl"`                            // 是否使用 HTTPS
	PathStyle    bool   `mapstructure:"path-style"    yaml:"path-style"    json:"path_style"`                         // 是否使用路径样式访问
	ModuleName   string `mapstructure:"modulename"    yaml:"modulename"    json:"module_name"`                        // 模块名称
}

// NewS3 创建一个新的 S3 实例
func NewS3(opt *S3) *S3 {
	var s3Instance *S3

	internal.LockFunc(func() {
		s3Instance = opt
	})
	return s3Instance
}

// Clone 返回 S3 配置的副本
func (s *S3) Clone() internal.Configurable {
	return &S3{
		ModuleName:   s.ModuleName,
		Endpoint:     s.Endpoint,
		Region:       s.Region,
		AccessKey:    s.AccessKey,
		SecretKey:    s.SecretKey,
		BucketPrefix: s.BucketPrefix,
		SessionToken: s.SessionToken,
		UseSSL:       s.UseSSL,
		PathStyle:    s.PathStyle,
	}
}

// OSSProvider 接口实现

// GetProviderType 获取提供商类型
func (s *S3) GetProviderType() OSSType {
	return OSSTypeS3
}

// GetEndpoint 获取端点地址
func (s *S3) GetEndpoint() string {
	return s.Endpoint
}

// GetAccessKey 获取访问密钥
func (s *S3) GetAccessKey() string {
	return s.AccessKey
}

// GetSecretKey 获取私密密钥
func (s *S3) GetSecretKey() string {
	return s.SecretKey
}

// GetBucket 获取存储桶名称 (S3使用前缀)
func (s *S3) GetBucket() string {
	return s.BucketPrefix
}

// IsSSL 是否使用SSL
func (s *S3) IsSSL() bool {
	return s.UseSSL
}

// GetModuleName 获取模块名称
func (s *S3) GetModuleName() string {
	return s.ModuleName
}

// SetCredentials 设置凭证
func (s *S3) SetCredentials(accessKey, secretKey string) {
	s.AccessKey = accessKey
	s.SecretKey = secretKey
}

// SetEndpoint 设置端点
func (s *S3) SetEndpoint(endpoint string) {
	s.Endpoint = endpoint
}

// SetBucket 设置存储桶
func (s *S3) SetBucket(bucket string) {
	s.BucketPrefix = bucket
}

// Get 返回 S3 配置的所有字段
func (s *S3) Get() interface{} {
	return s
}

// Set 更新 S3 配置的字段
func (s *S3) Set(data interface{}) {
	if configData, ok := data.(*S3); ok {
		s.ModuleName = configData.ModuleName
		s.Endpoint = configData.Endpoint
		s.Region = configData.Region
		s.AccessKey = configData.AccessKey
		s.SecretKey = configData.SecretKey
		s.BucketPrefix = configData.BucketPrefix
		s.SessionToken = configData.SessionToken
		s.UseSSL = configData.UseSSL
		s.PathStyle = configData.PathStyle
	}
}

// Validate 验证 S3 配置的有效性
func (s *S3) Validate() error {
	return internal.ValidateStruct(s)
}

// GetBucketName 根据前缀和后缀生成完整的存储桶名称
func (s *S3) GetBucketName(suffix string) string {
	if s.BucketPrefix == "" {
		return suffix
	}
	if suffix == "" {
		return s.BucketPrefix
	}
	return s.BucketPrefix + "-" + suffix
}

// IsTemporaryCredentials 判断是否使用临时凭证
func (s *S3) IsTemporaryCredentials() bool {
	return s.SessionToken != ""
}

// DefaultS3 返回默认S3配置
func DefaultS3() S3 {
	return S3{
		ModuleName:   "s3",
		Endpoint:     "https://s3.amazonaws.com",
		Region:       "us-east-1",
		AccessKey:    "demo_access_key",
		SecretKey:    "s3_secret_key",
		BucketPrefix: "demo-bucket",
		SessionToken: "",
		UseSSL:       true,
		PathStyle:    false,
	}
}

// Default 返回默认S3配置的指针，支持链式调用
func DefaultS3Config() *S3 {
	config := DefaultS3()
	return &config
}

// WithModuleName 设置模块名称
func (s *S3) WithModuleName(moduleName string) *S3 {
	s.ModuleName = moduleName
	return s
}

// WithEndpoint 设置S3端点地址
func (s *S3) WithEndpoint(endpoint string) *S3 {
	s.Endpoint = endpoint
	return s
}

// WithRegion 设置AWS区域
func (s *S3) WithRegion(region string) *S3 {
	s.Region = region
	return s
}

// WithAccessKey 设置AWS Access Key ID
func (s *S3) WithAccessKey(accessKey string) *S3 {
	s.AccessKey = accessKey
	return s
}

// WithSecretKey 设置AWS Secret Access Key
func (s *S3) WithSecretKey(secretKey string) *S3 {
	s.SecretKey = secretKey
	return s
}

// WithBucketPrefix 设置存储桶前缀
func (s *S3) WithBucketPrefix(bucketPrefix string) *S3 {
	s.BucketPrefix = bucketPrefix
	return s
}

// WithSessionToken 设置会话令牌
func (s *S3) WithSessionToken(sessionToken string) *S3 {
	s.SessionToken = sessionToken
	return s
}

// WithDisableSSL 设置是否使用HTTPS
func (s *S3) WithDisableSSL(disableSSL bool) *S3 {
	s.UseSSL = !disableSSL
	return s
}

// WithPathStyle 设置是否使用路径样式访问
func (s *S3) WithPathStyle(pathStyle bool) *S3 {
	s.PathStyle = pathStyle
	return s
}
