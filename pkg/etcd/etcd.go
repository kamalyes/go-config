/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-15 13:28:53
 * @FilePath: \go-config\pkg\etcd\etcd.go
 * @Description: Etcd配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package etcd

import (
	"github.com/kamalyes/go-config/internal"
)

// Etcd 结构体表示 Etcd 的配置
type Etcd struct {
	ModuleName         string   `mapstructure:"module_name" yaml:"module_name" json:"module_name"`                            // 模块名称
	Hosts              []string `mapstructure:"hosts" yaml:"hosts" json:"hosts" validate:"required,dive,required"`            // Etcd 主机列表
	Key                string   `mapstructure:"key" yaml:"key" json:"key" validate:"required"`                                // 注册的键
	ID                 int64    `mapstructure:"id" yaml:"id" json:"id" validate:"required"`                                   // ID
	User               string   `mapstructure:"user" yaml:"user" json:"user"`                                                 // 用户名
	Pass               string   `mapstructure:"pass" yaml:"pass" json:"pass"`                                                 // 密码
	CertFile           string   `mapstructure:"cert_file" yaml:"cert_file" json:"cert_file"`                                  // 证书文件
	CertKeyFile        string   `mapstructure:"cert_key_file" yaml:"cert_key_file" json:"cert_key_file"`                      // 证书密钥文件
	CACertFile         string   `mapstructure:"ca_cert_file" yaml:"ca_cert_file" json:"ca_cert_file"`                         // CA 证书文件
	InsecureSkipVerify bool     `mapstructure:"insecure_skip_verify" yaml:"insecure_skip_verify" json:"insecure_skip_verify"` // 是否跳过证书验证
	Namespace          string   `mapstructure:"namespace" yaml:"namespace" json:"namespace"`                                  // 命名空间
	DialTimeout        int      `mapstructure:"dial_timeout" yaml:"dial_timeout" json:"dial_timeout"`                         // 连接超时时间(秒)
	RequestTimeout     int      `mapstructure:"request_timeout" yaml:"request_timeout" json:"request_timeout"`                // 请求超时时间(秒)
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
		ModuleName:         e.ModuleName,
		Hosts:              append([]string(nil), e.Hosts...),
		Key:                e.Key,
		ID:                 e.ID,
		User:               e.User,
		Pass:               e.Pass,
		CertFile:           e.CertFile,
		CertKeyFile:        e.CertKeyFile,
		CACertFile:         e.CACertFile,
		InsecureSkipVerify: e.InsecureSkipVerify,
		Namespace:          e.Namespace,
		DialTimeout:        e.DialTimeout,
		RequestTimeout:     e.RequestTimeout,
	}
}

// Get 返回 Etcd 配置的所有字段
func (e *Etcd) Get() interface{} {
	return e
}

// Set 更新 Etcd 配置的字段
func (e *Etcd) Set(data interface{}) {
	if configData, ok := data.(*Etcd); ok {
		e.ModuleName = configData.ModuleName
		e.Hosts = configData.Hosts
		e.Key = configData.Key
		e.ID = configData.ID
		e.User = configData.User
		e.Pass = configData.Pass
		e.CertFile = configData.CertFile
		e.CertKeyFile = configData.CertKeyFile
		e.CACertFile = configData.CACertFile
		e.InsecureSkipVerify = configData.InsecureSkipVerify
		e.Namespace = configData.Namespace
		e.DialTimeout = configData.DialTimeout
		e.RequestTimeout = configData.RequestTimeout
	}
}

// Validate 验证 Etcd 配置的有效性
func (e *Etcd) Validate() error {
	// 设置默认值
	if len(e.Hosts) == 0 {
		e.Hosts = []string{"127.0.0.1:2379"}
	}
	if e.DialTimeout <= 0 {
		e.DialTimeout = 5
	}
	if e.RequestTimeout <= 0 {
		e.RequestTimeout = 10
	}
	if e.ModuleName == "" {
		e.ModuleName = "etcd"
	}
	return internal.ValidateStruct(e)
}

// DefaultEect 返回默认Etcd配置
func DefaultEect() Etcd {
	return Etcd{
		ModuleName:         "etcd",
		Hosts:              []string{"127.0.0.1:2379"},
		Key:                "app/config",
		ID:                 1,
		User:               "etcd_user",
		Pass:               "etcd_password",
		CertFile:           "/etc/etcd/client.crt",
		CertKeyFile:        "/etc/etcd/client.key",
		CACertFile:         "/etc/etcd/ca.crt",
		InsecureSkipVerify: true,
		Namespace:          "default",
		DialTimeout:        5,
		RequestTimeout:     10,
	}
}

// Default 返回默认Etcd配置的指针，支持链式调用
func Default() *Etcd {
	config := DefaultEect()
	return &config
}

// WithModuleName 设置模块名称
func (e *Etcd) WithModuleName(moduleName string) *Etcd {
	e.ModuleName = moduleName
	return e
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

// WithAuth 设置用户名和密码
func (e *Etcd) WithAuth(user, pass string) *Etcd {
	e.User = user
	e.Pass = pass
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

// WithCertFile 设置证书文件路径
func (e *Etcd) WithCertFile(certFile string) *Etcd {
	e.CertFile = certFile
	return e
}

// WithCertKeyFile 设置证书密钥文件路径
func (e *Etcd) WithCertKeyFile(certKeyFile string) *Etcd {
	e.CertKeyFile = certKeyFile
	return e
}

// WithCACertFile 设置CA证书文件路径
func (e *Etcd) WithCACertFile(caCertFile string) *Etcd {
	e.CACertFile = caCertFile
	return e
}

// WithInsecureSkipVerify 设置是否跳过证书验证
func (e *Etcd) WithInsecureSkipVerify(skip bool) *Etcd {
	e.InsecureSkipVerify = skip
	return e
}

// WithDialTimeout 设置连接超时时间
func (e *Etcd) WithDialTimeout(timeout int) *Etcd {
	e.DialTimeout = timeout
	return e
}

// WithRequestTimeout 设置请求超时时间
func (e *Etcd) WithRequestTimeout(timeout int) *Etcd {
	e.RequestTimeout = timeout
	return e
}

// WithCert 设置证书文件
func (e *Etcd) WithCert(certFile, certKeyFile, caCertFile string) *Etcd {
	e.CertFile = certFile
	e.CertKeyFile = certKeyFile
	e.CACertFile = caCertFile
	return e
}

// WithTLS 设置TLS配置
func (e *Etcd) WithTLS(certFile, certKeyFile, caCertFile string) *Etcd {
	e.CertFile = certFile
	e.CertKeyFile = certKeyFile
	e.CACertFile = caCertFile
	e.InsecureSkipVerify = false
	return e
}

// WithNamespace 设置命名空间
func (e *Etcd) WithNamespace(namespace string) *Etcd {
	e.Namespace = namespace
	return e
}

// WithTimeouts 设置超时配置
func (e *Etcd) WithTimeouts(dialTimeout, requestTimeout int) *Etcd {
	e.DialTimeout = dialTimeout
	e.RequestTimeout = requestTimeout
	return e
}

// EnableInsecure 启用不安全连接
func (e *Etcd) EnableInsecure() *Etcd {
	e.InsecureSkipVerify = true
	return e
}

// DisableInsecure 禁用不安全连接
func (e *Etcd) DisableInsecure() *Etcd {
	e.InsecureSkipVerify = false
	return e
}
