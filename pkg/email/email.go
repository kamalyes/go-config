/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:35:55
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
	ModuleName string `mapstructure:"modulename"             yaml:"modulename"          json:"module_name"      validate:"required"`                 // 模块名称
	To         string `mapstructure:"to"                      yaml:"to"                  json:"to"              validate:"required,email"`           // 收件人:多个以英文逗号分隔
	From       string `mapstructure:"from"                    yaml:"from"                json:"from"            validate:"required,email"`           // 发件人
	Host       string `mapstructure:"host"                    yaml:"host"                json:"host"            validate:"required"`                 // 邮件服务器地址
	Port       int    `mapstructure:"port"                    yaml:"port"                json:"port"            validate:"required,min=1,max=65535"` // 端口
	Secret     string `mapstructure:"secret"                  yaml:"secret"              json:"secret"          validate:"required"`                 // 密钥
	IsSSL      bool   `mapstructure:"is-ssl"                  yaml:"is-ssl"             json:"is_ssl"`                                               // 是否SSL
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
