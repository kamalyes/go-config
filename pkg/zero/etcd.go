/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-07 15:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 19:55:50
 * @FilePath: \go-config\pkg\zero\etcd.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

import (
	"github.com/kamalyes/go-config/internal"
)

// Etcd 结构体表示 Etcd 的配置
type Etcd struct {
	Hosts              []string `mapstructure:"hosts"               yaml:"hosts"               json:"hosts"               validate:"required,dive,required"` // Etcd 主机列表
	Key                string   `mapstructure:"key"                 yaml:"key"                 json:"key"                 validate:"required"`               // 注册的键
	ID                 int64    `mapstructure:"id"                  yaml:"id"                  json:"id"                  validate:"required"`               // ID
	User               string   `mapstructure:"user"                yaml:"user"                json:"user"`                                                  // 用户名
	Pass               string   `mapstructure:"pass"                yaml:"pass"                json:"pass"`                                                  // 密码
	CertFile           string   `mapstructure:"cert-file"           yaml:"cert-file"           json:"cert_file"`                                             // 证书文件
	CertKeyFile        string   `mapstructure:"cert-key-file"       yaml:"cert-key-file"       json:"cert_key_file"`                                         // 证书密钥文件
	CACertFile         string   `mapstructure:"ca-cert-file"        yaml:"ca-cert-file"        json:"ca_cert_file"`                                          // CA 证书文件
	InsecureSkipVerify bool     `mapstructure:"insecure-skip-verify" yaml:"insecure-skip-verify" json:"insecure_skip_verify"`                                // 是否跳过证书验证
}

// NewEtcd 创建一个新的 Etcd 实例
func NewEtcd(opt *Etcd) *Etcd {
	var etcdInstance *Etcd

	internal.LockFunc(func() {
		etcdInstance = opt
	})
	return etcdInstance
}

// Clone 返回 Etcd 配置的副本
func (e *Etcd) Clone() internal.Configurable {
	return &Etcd{
		Hosts:              e.Hosts,
		Key:                e.Key,
		ID:                 e.ID,
		User:               e.User,
		Pass:               e.Pass,
		CertFile:           e.CertFile,
		CertKeyFile:        e.CertKeyFile,
		CACertFile:         e.CACertFile,
		InsecureSkipVerify: e.InsecureSkipVerify,
	}
}

// Get 返回 Etcd 配置的所有字段
func (e *Etcd) Get() interface{} {
	return e
}

// Set 更新 Etcd 配置的字段
func (e *Etcd) Set(data interface{}) {
	if configData, ok := data.(*Etcd); ok {
		e.Hosts = configData.Hosts
		e.Key = configData.Key
		e.ID = configData.ID
		e.User = configData.User
		e.Pass = configData.Pass
		e.CertFile = configData.CertFile
		e.CertKeyFile = configData.CertKeyFile
		e.CACertFile = configData.CACertFile
		e.InsecureSkipVerify = configData.InsecureSkipVerify
	}
}

// Validate 验证 Etcd 配置的有效性
func (e *Etcd) Validate() error {
	return internal.ValidateStruct(e)
}
