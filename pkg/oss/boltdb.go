/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-08 12:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-17 16:30:00
 * @FilePath: \go-config\pkg\oss\boltdb.go
 * @Description: BoltDB 本地对象存储配置
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package oss

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-toolbox/pkg/syncx"
)

// BoltDB 结构体用于配置 BoltDB 本地存储
type BoltDB struct {
	Path       string `mapstructure:"path" yaml:"path" json:"path"        validate:"required"` // 数据库文件路径
	ModuleName string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`        // 模块名称
}

// NewBoltDB 创建一个新的 BoltDB 实例
func NewBoltDB(opt *BoltDB) *BoltDB {
	var boltDBInstance *BoltDB

	internal.LockFunc(func() {
		boltDBInstance = opt
	})
	return boltDBInstance
}

// Clone 返回 BoltDB 配置的副本
func (b *BoltDB) Clone() internal.Configurable {
	var cloned BoltDB
	if err := syncx.DeepCopy(&cloned, b); err != nil {
		// 如果深拷贝失败，返回空配置
		return &BoltDB{}
	}
	return &cloned
}

// OSSProvider 接口实现

// GetProviderType 获取提供商类型
func (b *BoltDB) GetProviderType() OSSType {
	return OSSTypeBoltDB
}

// GetEndpoint 获取端点地址 (BoltDB本地存储,返回文件路径)
func (b *BoltDB) GetEndpoint() string {
	return b.Path
}

// GetAccessKey BoltDB不需要访问密钥
func (b *BoltDB) GetAccessKey() string {
	return ""
}

// GetSecretKey BoltDB不需要私密密钥
func (b *BoltDB) GetSecretKey() string {
	return ""
}

// GetBucket BoltDB使用文件路径
func (b *BoltDB) GetBucket() string {
	return b.Path
}

// IsSSL BoltDB是本地存储
func (b *BoltDB) IsSSL() bool {
	return false
}

// GetModuleName 获取模块名称
func (b *BoltDB) GetModuleName() string {
	return b.ModuleName
}

// SetCredentials BoltDB不需要凭证
func (b *BoltDB) SetCredentials(accessKey, secretKey string) {
	// BoltDB不需要凭证
}

// SetEndpoint 设置端点(文件路径)
func (b *BoltDB) SetEndpoint(endpoint string) {
	b.Path = endpoint
}

// SetBucket 设置存储桶(文件路径)
func (b *BoltDB) SetBucket(bucket string) {
	b.Path = bucket
}

// Get 返回 BoltDB 配置的所有字段
func (b *BoltDB) Get() interface{} {
	return b
}

// Set 更新 BoltDB 配置的字段
func (b *BoltDB) Set(data interface{}) {
	if configData, ok := data.(*BoltDB); ok {
		b.ModuleName = configData.ModuleName
		b.Path = configData.Path
	}
}

// Validate 验证 BoltDB 配置的有效性
func (b *BoltDB) Validate() error {
	return internal.ValidateStruct(b)
}

// DefaultBoltDB 返回默认BoltDB配置
func DefaultBoltDB() *BoltDB {
	return &BoltDB{
		Path:       "./data/storage.db",
		ModuleName: "boltdb",
	}
}
