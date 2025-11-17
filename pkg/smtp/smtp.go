/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-17 14:51:53
 * @FilePath: \im-access-control-service\go-config\pkg\smtp\smtp.go
 * @Description: 邮箱配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package smtp

import "github.com/kamalyes/go-config/internal"

// Smtp 邮箱配置
type Smtp struct {
	ModuleName  string            `mapstructure:"module_name" yaml:"module-name" json:"module_name"`    // 模块名称
	Enabled     bool              `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                // 是否启用
	SMTPHost    string            `mapstructure:"smtp_host" yaml:"smtp-host" json:"smtp_host"`          // SMTP主机
	SMTPPort    int               `mapstructure:"smtp_port" yaml:"smtp-port" json:"smtp_port"`          // SMTP端口
	Username    string            `mapstructure:"username" yaml:"username" json:"username"`             // 用户名
	Password    string            `mapstructure:"password" yaml:"password" json:"password"`             // 密码
	FromAddress string            `mapstructure:"from_address" yaml:"from-address" json:"from_address"` // 发件人地址
	ToAddresses []string          `mapstructure:"to_addresses" yaml:"to-addresses" json:"to_addresses"` // 收件人地址列表
	EnableTLS   bool              `mapstructure:"enable_tls" yaml:"enable-tls" json:"enable_tls"`       // 是否启用TLS
	Headers     map[string]string `mapstructure:"headers" yaml:"headers" json:"headers"`                // 自定义头部
	PoolSize    int               `mapstructure:"pool_size" yaml:"pool-size" json:"pool_size"`          // 连接池大小
}

// Default 创建默认邮箱配置
func Default() *Smtp {
	return &Smtp{
		ModuleName:  "smtp",
		Enabled:     true,
		SMTPHost:    "127.0.0.1",
		SMTPPort:    587,
		Username:    "",
		Password:    "",
		FromAddress: "",
		ToAddresses: []string{},
		EnableTLS:   true,
		Headers:     make(map[string]string),
		PoolSize:    5,
	}
}

// Get 返回配置接口
func (e *Smtp) Get() interface{} {
	return e
}

// Set 设置配置数据
func (e *Smtp) Set(data interface{}) {
	if cfg, ok := data.(*Smtp); ok {
		*e = *cfg
	}
}

// Clone 返回配置的副本
func (e *Smtp) Clone() internal.Configurable {
	clone := &Smtp{
		ModuleName:  e.ModuleName,
		SMTPHost:    e.SMTPHost,
		SMTPPort:    e.SMTPPort,
		Username:    e.Username,
		Password:    e.Password,
		FromAddress: e.FromAddress,
		ToAddresses: make([]string, len(e.ToAddresses)),
		EnableTLS:   e.EnableTLS,
		PoolSize:    e.PoolSize,
	}

	copy(clone.ToAddresses, e.ToAddresses)

	clone.Headers = make(map[string]string)
	for k, v := range e.Headers {
		clone.Headers[k] = v
	}

	return clone
}

// Validate 验证配置
func (e *Smtp) Validate() error {
	return internal.ValidateStruct(e)
}

// WithModuleName 设置模块名称
func (e *Smtp) WithModuleName(name string) *Smtp {
	e.ModuleName = name
	return e
}

// WithEnabled 设置是否启用
func (e *Smtp) WithEnabled(enabled bool) *Smtp {
	e.Enabled = enabled
	return e
}

// WithSMTPHost 设置SMTP主机
func (e *Smtp) WithSMTPHost(host string) *Smtp {
	e.SMTPHost = host
	return e
}

// WithSMTPPort 设置SMTP端口
func (e *Smtp) WithSMTPPort(port int) *Smtp {
	e.SMTPPort = port
	return e
}

// WithUsername 设置用户名
func (e *Smtp) WithUsername(username string) *Smtp {
	e.Username = username
	return e
}

// WithPassword 设置密码
func (e *Smtp) WithPassword(password string) *Smtp {
	e.Password = password
	return e
}

// WithFromAddress 设置发件人地址
func (e *Smtp) WithFromAddress(address string) *Smtp {
	e.FromAddress = address
	return e
}

// WithPoolSize 设置连接池大小
func (e *Smtp) WithPoolSize(size int) *Smtp {
	e.PoolSize = size
	return e
}

// AddToAddress 添加收件人地址
func (e *Smtp) AddToAddress(address string) *Smtp {
	e.ToAddresses = append(e.ToAddresses, address)
	return e
}

// EnableTLSService 启用TLS
func (e *Smtp) EnableTLSService() *Smtp {
	e.EnableTLS = true
	return e
}

// DisableTLS 禁用TLS
func (e *Smtp) DisableTLS() *Smtp {
	e.EnableTLS = false
	return e
}

// AddHeader 添加自定义头部
func (e *Smtp) AddHeader(key, value string) *Smtp {
	if e.Headers == nil {
		e.Headers = make(map[string]string)
	}
	e.Headers[key] = value
	return e
}
