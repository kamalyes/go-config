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

// DefaultEtcd 返回默认的 Etcd 值
func DefaultEtcd() Etcd {
	return Etcd{
		Hosts:              []string{"127.0.0.1:2379"},
		Key:                "default-key",
		ID:                 1,
		User:               "",
		Pass:               "",
		CertFile:           "",
		CertKeyFile:        "",
		CACertFile:         "",
		InsecureSkipVerify: false,
	}
}

// DefaultEtcdConfig 返回默认的 Etcd 指针，支持链式调用
func DefaultEtcdConfig() *Etcd {
	config := DefaultEtcd()
	return &config
}

// WithHosts 设置Etcd主机列表
func (e *Etcd) WithHosts(hosts []string) *Etcd {
	e.Hosts = hosts
	return e
}

// WithKey 设置注册的键
func (e *Etcd) WithKey(key string) *Etcd {
	e.Key = key
	return e
}

// WithID 设置ID
func (e *Etcd) WithID(id int64) *Etcd {
	e.ID = id
	return e
}

// WithUser 设置用户名
func (e *Etcd) WithUser(user string) *Etcd {
	e.User = user
	return e
}

// WithPass 设置密码
func (e *Etcd) WithPass(pass string) *Etcd {
	e.Pass = pass
	return e
}

// WithCertFile 设置证书文件
func (e *Etcd) WithCertFile(certFile string) *Etcd {
	e.CertFile = certFile
	return e
}

// WithCertKeyFile 设置证书密钥文件
func (e *Etcd) WithCertKeyFile(certKeyFile string) *Etcd {
	e.CertKeyFile = certKeyFile
	return e
}

// WithCACertFile 设置CA证书文件
func (e *Etcd) WithCACertFile(caCertFile string) *Etcd {
	e.CACertFile = caCertFile
	return e
}

// WithInsecureSkipVerify 设置是否跳过证书验证
func (e *Etcd) WithInsecureSkipVerify(insecureSkipVerify bool) *Etcd {
	e.InsecureSkipVerify = insecureSkipVerify
	return e
}
