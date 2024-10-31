/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 11:39:04
 * @FilePath: \go-config\pkg\sms\aliyun.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package sms

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// AliyunSms 结构体表示 SMS 配置
type AliyunSms struct {
	ModuleName           string `mapstructure:"MODULE_NAME"                     yaml:"modulename"`             // 模块名称
	SecretID             string `mapstructure:"SECRET_ID"                       yaml:"secret_id"`              // 阿里云短信服务的 Secret ID
	SecretKey            string `mapstructure:"SECRET_KEY"                      yaml:"secret_key"`             // 阿里云短信服务的 Secret Key
	Sign                 string `mapstructure:"SIGN"                            yaml:"sign"`                   // 短信签名
	ResourceOwnerAccount string `mapstructure:"RESOURCE_OWNER_ACCOUNT"          yaml:"resource_owner_account"` // 资源所有者账户
	ResourceOwnerID      int64  `mapstructure:"RESOURCE_OWNER_ID"               yaml:"resource_owner_id"`      // 资源所有者 ID
	TemplateCodeVerify   string `mapstructure:"TEMPLATE_CODE_VERIFY"            yaml:"template_code_verify"`   // 短信模板代码
	Endpoint             string `mapstructure:"ENDPOINT"                        yaml:"endpoint"`               // 短信服务的 API 端点
}

// NewAliyunSms 创建一个新的 AliyunSms 实例
func NewAliyunSms(opt *AliyunSms) *AliyunSms {
	var aliyunSmsInstance *AliyunSms

	internal.LockFunc(func() {
		aliyunSmsInstance = opt
	})
	return aliyunSmsInstance
}

// Clone 返回 AliyunSms 配置的副本
func (a *AliyunSms) Clone() internal.Configurable {
	return &AliyunSms{
		ModuleName:           a.ModuleName,
		SecretID:             a.SecretID,
		SecretKey:            a.SecretKey,
		Sign:                 a.Sign,
		ResourceOwnerAccount: a.ResourceOwnerAccount,
		ResourceOwnerID:      a.ResourceOwnerID,
		TemplateCodeVerify:   a.TemplateCodeVerify,
		Endpoint:             a.Endpoint,
	}
}

// Get 返回 AliyunSms 配置的所有字段
func (a *AliyunSms) Get() interface{} {
	return a
}

// Set 更新 AliyunSms 配置的字段
func (a *AliyunSms) Set(data interface{}) {
	if configData, ok := data.(*AliyunSms); ok {
		a.ModuleName = configData.ModuleName
		a.SecretID = configData.SecretID
		a.SecretKey = configData.SecretKey
		a.Sign = configData.Sign
		a.ResourceOwnerAccount = configData.ResourceOwnerAccount
		a.ResourceOwnerID = configData.ResourceOwnerID
		a.TemplateCodeVerify = configData.TemplateCodeVerify
		a.Endpoint = configData.Endpoint
	}
}

// Validate 验证 AliyunSms 配置的有效性
func (a *AliyunSms) Validate() error {
	if a.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if a.SecretID == "" {
		return errors.New("secret ID cannot be empty")
	}
	if a.SecretKey == "" {
		return errors.New("secret key cannot be empty")
	}
	if a.Sign == "" {
		return errors.New("sign cannot be empty")
	}
	if a.ResourceOwnerAccount == "" {
		return errors.New("resource owner account cannot be empty")
	}
	if a.ResourceOwnerID <= 0 {
		return errors.New("resource owner ID must be greater than zero")
	}
	if a.TemplateCodeVerify == "" {
		return errors.New("template code verify cannot be empty")
	}
	if a.Endpoint == "" {
		return errors.New("endpoint cannot be empty")
	}
	return nil
}
