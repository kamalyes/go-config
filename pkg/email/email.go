/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:47:34
 * @FilePath: \go-config\pkg\email\email.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package email

import (
	"github.com/kamalyes/go-config/internal"
)

// Email 邮件配置结构体
type Email struct {
	To         string `mapstructure:"to" yaml:"to" json:"to"              validate:"required,email"`               // 收件人:多个以英文逗号分隔
	From       string `mapstructure:"from" yaml:"from" json:"from"            validate:"required,email"`           // 发件人
	Host       string `mapstructure:"host" yaml:"host" json:"host"            validate:"required"`                 // 邮件服务器地址
	Port       int    `mapstructure:"port" yaml:"port" json:"port"            validate:"required,min=1,max=65535"` // 端口
	Secret     string `mapstructure:"secret" yaml:"secret" json:"secret"          validate:"required"`             // 密钥
	IsSSL      bool   `mapstructure:"is-ssl" yaml:"is-ssl" json:"isSsl"`                                           // 是否SSL
	ModuleName string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                            // 模块名称
}

// NewEmail 创建一个新的 Email 实例
func NewEmail(opt *Email) *Email {
	var emailInstance *Email

	internal.LockFunc(func() {
		emailInstance = opt
	})
	return emailInstance
}

// Clone 返回 Email 配置的副本
func (e *Email) Clone() internal.Configurable {
	return &Email{
		ModuleName: e.ModuleName,
		To:         e.To,
		From:       e.From,
		Host:       e.Host,
		Port:       e.Port,
		IsSSL:      e.IsSSL,
		Secret:     e.Secret,
	}
}

// Get 返回 Email 配置的所有字段
func (e *Email) Get() interface{} {
	return e
}

// Set 更新 Email 配置的字段
func (e *Email) Set(data interface{}) {
	if configData, ok := data.(*Email); ok {
		e.ModuleName = configData.ModuleName
		e.To = configData.To
		e.From = configData.From
		e.Host = configData.Host
		e.Port = configData.Port
		e.IsSSL = configData.IsSSL
		e.Secret = configData.Secret
	}
}

// Validate 检查 Email 配置的有效性
func (e *Email) Validate() error {
	return internal.ValidateStruct(e)
}

// DefaultEmail 返回默认Email配置
func DefaultEmail() Email {
	return Email{
		ModuleName: "email",
		To:         "recipient@example.com",
		From:       "sender@example.com",
		Host:       "smtp.gmail.com",
		Port:       587,
		Secret:     "email_app_password",
		IsSSL:      true,
	}
}

// Default 返回默认Email配置的指针，支持链式调用
func Default() *Email {
	config := DefaultEmail()
	return &config
}

// WithModuleName 设置模块名称
func (e *Email) WithModuleName(moduleName string) *Email {
	e.ModuleName = moduleName
	return e
}

// WithTo 设置收件人
func (e *Email) WithTo(to string) *Email {
	e.To = to
	return e
}

// WithFrom 设置发件人
func (e *Email) WithFrom(from string) *Email {
	e.From = from
	return e
}

// WithHost 设置邮件服务器地址
func (e *Email) WithHost(host string) *Email {
	e.Host = host
	return e
}

// WithPort 设置端口
func (e *Email) WithPort(port int) *Email {
	e.Port = port
	return e
}

// WithSecret 设置密钥
func (e *Email) WithSecret(secret string) *Email {
	e.Secret = secret
	return e
}

// WithIsSSL 设置是否启用SSL
func (e *Email) WithIsSSL(isSSL bool) *Email {
	e.IsSSL = isSSL
	return e
}
