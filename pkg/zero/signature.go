/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-07 15:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 18:07:55
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
	KeyFile     string `json:"key-file" mapstructure:"key-file" yaml:"key_file"`          // 私钥文件路径
}

// SignatureConf 是一个签名配置
type SignatureConf struct {
	Strict      bool             `json:"strict" mapstructure:"strict" yaml:"strict" default:"false"`   // 是否严格验证
	Expiry      time.Duration    `json:"expiry" mapstructure:"expiry" yaml:"expiry" default:"1h"`      // 签名过期时间
	PrivateKeys []PrivateKeyConf `json:"private-keys" mapstructure:"private-keys" yaml:"private_keys"` // 私钥列表
}

// Signature 表示签名配置
type Signature struct {
	Algorithm SignatureConf `mapstructure:"algorithm" yaml:"algorithm" json:"algorithm"` // 签名算法
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
func (e *Signature) Clone() internal.Configurable {
	return &Signature{
		Algorithm: e.Algorithm,
	}
}

// Get 返回 Signature 配置的所有字段
func (e *Signature) Get() interface{} {
	return e
}

// Set 更新 Signature 配置的字段
func (e *Signature) Set(data interface{}) {
	if configData, ok := data.(*Signature); ok {
		e.Algorithm = configData.Algorithm
	}
}

// Validate 验证 Signature 配置的有效性
func (e *Signature) Validate() error {
	return internal.ValidateStruct(e)
}
