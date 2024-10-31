/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:15:21
 * @FilePath: \go-config\email\email.go
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

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** 收件人:多个以英文逗号分隔 */
	To string `mapstructure:"to"         json:"to"           yaml:"to"`

	/** 设置发件人 */
	From string `mapstructure:"from"       json:"from"         yaml:"from"`

	/** 邮件服务器地址 */
	Host string `mapstructure:"host"       json:"host"         yaml:"host"`

	/** 端口 */
	Port int `mapstructure:"port"       json:"port"         yaml:"port"`

	/** 是否SSL */
	IsSSL bool `mapstructure:"is-ssl"     json:"isSSL"        yaml:"is-ssl"` // 是否SSL

	/** 密钥 */
	Secret string `mapstructure:"secret"     json:"secret"       yaml:"secret"`
}

// NewEmail 创建一个新的 Email 实例
func NewEmail(moduleName, to, from, host string, port int, isSSL bool, secret string) *Email {
	var emailInstance *Email

	internal.LockFunc(func() {
		emailInstance = &Email{
			ModuleName: moduleName,
			To:         to,
			From:       from,
			Host:       host,
			Port:       port,
			IsSSL:      isSSL,
			Secret:     secret,
		}
	})
	return emailInstance
}

// ToMap 将配置转换为映射
func (e *Email) ToMap() map[string]interface{} {
	return internal.ToMap(e)
}

// FromMap 从映射中填充配置
func (e *Email) FromMap(data map[string]interface{}) {
	internal.FromMap(e, data)
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
