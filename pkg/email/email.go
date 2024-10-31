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
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Email 邮件配置
type Email struct {
	ModuleName string `mapstructure:"MODULE_NAME"             yaml:"modulename"` // 模块名称
	To         string `mapstructure:"TO"                      yaml:"to"`         // 收件人:多个以英文逗号分隔
	From       string `mapstructure:"FROM"                    yaml:"from"`       // 发件人
	Host       string `mapstructure:"HOST"                    yaml:"host"`       // 邮件服务器地址
	Port       int    `mapstructure:"PORT"                    yaml:"port"`       // 端口
	IsSSL      bool   `mapstructure:"IS_SSL"                  yaml:"is-ssl"`     // 是否SSL
	Secret     string `mapstructure:"SECRET"                  yaml:"secret"`     // 密钥
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
	if e.To == "" {
		return errors.New("to cannot be empty")
	}
	if e.From == "" {
		return errors.New("from cannot be empty")
	}
	if e.Host == "" {
		return errors.New("host cannot be empty")
	}
	if e.Port <= 0 {
		return errors.New("port must be greater than 0")
	}
	return nil
}
