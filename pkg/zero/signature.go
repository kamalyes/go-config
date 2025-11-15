/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-07 15:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 18:25:35
 * @FilePath: \go-config\pkg\zero\signature.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package zero

import (
	"time"

	"github.com/kamalyes/go-config/internal"
)

// PrivateKeyConf 是一个私钥配置
type PrivateKeyConf struct {
	Fingerprint string `json:"fingerprint" mapstructure:"fingerprint" yaml:"fingerprint"` // 私钥指纹
	KeyFile     string `json:"key-file" mapstructure:"key-file" yaml:"key-file"`          // 私钥文件路径
}

// Signature 是一个签名配置
type Signature struct {
	Strict      bool             `json:"strict" mapstructure:"strict" yaml:"strict" default:"false"`   // 是否严格验证
	Expiry      time.Duration    `json:"expiry" mapstructure:"expiry" yaml:"expiry" default:"1h"`      // 签名过期时间
	PrivateKeys []PrivateKeyConf `json:"private-keys" mapstructure:"private-keys" yaml:"private-keys"` // 私钥列表
}

// NewSignature 创建一个新的 Signature 实例
func NewSignature(opt *Signature) *Signature {
	var etcdInstance *Signature

	internal.LockFunc(func() {
		etcdInstance = opt
	})
	return etcdInstance
}

// Clone 返回 Signature 配置的副本
func (s *Signature) Clone() internal.Configurable {
	return &Signature{
		Strict:      s.Strict,
		Expiry:      s.Expiry,
		PrivateKeys: s.PrivateKeys,
	}
}

// Get 返回 Signature 配置的所有字段
func (s *Signature) Get() interface{} {
	return s
}

// Set 更新 Signature 配置的字段
func (s *Signature) Set(data interface{}) {
	if configData, ok := data.(*Signature); ok {
		s.Strict = configData.Strict
		s.Expiry = configData.Expiry
		s.PrivateKeys = configData.PrivateKeys
	}
}

// Validate 验证 Signature 配置的有效性
func (s *Signature) Validate() error {
	return internal.ValidateStruct(s)
}

// DefaultSignature 返回默认的 Signature 值
func DefaultSignature() Signature {
	return Signature{
		Strict:      false,
		Expiry:      time.Hour,
		PrivateKeys: []PrivateKeyConf{},
	}
}

// DefaultSignatureConfig 返回默认的 Signature 指针，支持链式调用
func DefaultSignatureConfig() *Signature {
	config := DefaultSignature()
	return &config
}

// WithStrict 设置是否严格验证
func (s *Signature) WithStrict(strict bool) *Signature {
	s.Strict = strict
	return s
}

// WithExpiry 设置签名过期时间
func (s *Signature) WithExpiry(expiry time.Duration) *Signature {
	s.Expiry = expiry
	return s
}

// WithPrivateKeys 设置私钥列表
func (s *Signature) WithPrivateKeys(privateKeys []PrivateKeyConf) *Signature {
	s.PrivateKeys = privateKeys
	return s
}

// AddPrivateKey 添加私钥
func (s *Signature) AddPrivateKey(fingerprint, keyFile string) *Signature {
	s.PrivateKeys = append(s.PrivateKeys, PrivateKeyConf{
		Fingerprint: fingerprint,
		KeyFile:     keyFile,
	})
	return s
}
