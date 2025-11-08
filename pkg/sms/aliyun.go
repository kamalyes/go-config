/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:53:49
 * @FilePath: \go-config\pkg\sms\aliyun.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package sms

import (
	"github.com/kamalyes/go-config/internal"
)

// AliyunSms 结构体表示 SMS 配置
type AliyunSms struct {
	SecretID             string `mapstructure:"secret-id"                       yaml:"secret-id"               json:"secret_id"             validate:"required"`     // 阿里云短信服务的 Secret ID
	SecretKey            string `mapstructure:"secret-key"                      yaml:"secret-key"              json:"secret_key"            validate:"required"`     // 阿里云短信服务的 Secret Key
	Sign                 string `mapstructure:"sign"                            yaml:"sign"                    json:"sign"                  validate:"required"`     // 短信签名
	ResourceOwnerAccount string `mapstructure:"resource-owner-account"          yaml:"resource-owner-account"  json:"resource_owner_account" validate:"required"`    // 资源所有者账户
	ResourceOwnerID      int64  `mapstructure:"resource-owner-id"               yaml:"resource-owner-id"       json:"resource_owner_id"     validate:"required"`     // 资源所有者 ID
	TemplateCodeVerify   string `mapstructure:"template-code-verify"            yaml:"template-code-verify"    json:"template_code_verify"   validate:"required"`    // 短信模板代码
	Endpoint             string `mapstructure:"endpoint"                        yaml:"endpoint"                json:"endpoint"              validate:"required,url"` // 短信服务的 API 端点，必须是有效的 URL
	ModuleName           string `mapstructure:"modulename"                      yaml:"modulename"              json:"module_name"`                                   // 模块名称
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
	return internal.ValidateStruct(a)
}

// DefaultAliyunSms 返回默认AliyunSms配置
func DefaultAliyunSms() AliyunSms {
	return AliyunSms{
		ModuleName:           "aliyun-sms",
		SecretID:             "",
		SecretKey:            "",
		Sign:                 "",
		ResourceOwnerAccount: "",
		ResourceOwnerID:      0,
		TemplateCodeVerify:   "",
		Endpoint:             "https://dysmsapi.aliyuncs.com",
	}
}

// Default 返回默认AliyunSms配置的指针，支持链式调用
func Default() *AliyunSms {
	config := DefaultAliyunSms()
	return &config
}

// WithModuleName 设置模块名称
func (a *AliyunSms) WithModuleName(moduleName string) *AliyunSms {
	a.ModuleName = moduleName
	return a
}

// WithSecretID 设置阿里云短信服务的Secret ID
func (a *AliyunSms) WithSecretID(secretID string) *AliyunSms {
	a.SecretID = secretID
	return a
}

// WithSecretKey 设置阿里云短信服务的Secret Key
func (a *AliyunSms) WithSecretKey(secretKey string) *AliyunSms {
	a.SecretKey = secretKey
	return a
}

// WithSign 设置短信签名
func (a *AliyunSms) WithSign(sign string) *AliyunSms {
	a.Sign = sign
	return a
}

// WithResourceOwnerAccount 设置资源所有者账户
func (a *AliyunSms) WithResourceOwnerAccount(resourceOwnerAccount string) *AliyunSms {
	a.ResourceOwnerAccount = resourceOwnerAccount
	return a
}

// WithResourceOwnerID 设置资源所有者ID
func (a *AliyunSms) WithResourceOwnerID(resourceOwnerID int64) *AliyunSms {
	a.ResourceOwnerID = resourceOwnerID
	return a
}

// WithTemplateCodeVerify 设置短信模板代码
func (a *AliyunSms) WithTemplateCodeVerify(templateCodeVerify string) *AliyunSms {
	a.TemplateCodeVerify = templateCodeVerify
	return a
}

// WithEndpoint 设置短信服务的API端点
func (a *AliyunSms) WithEndpoint(endpoint string) *AliyunSms {
	a.Endpoint = endpoint
	return a
}
