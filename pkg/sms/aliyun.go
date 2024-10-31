/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 13:51:19
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

// AliyunSms 表示 SMS 配置结构体
type AliyunSms struct {

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** 阿里云短信服务的 Secret ID */
	SecretID string `mapstructure:"secret-id" json:"secretId" yaml:"secret_id"`

	/** 阿里云短信服务的 Secret Key */
	SecretKey string `mapstructure:"secret-key" json:"secretKey" yaml:"secret_key"`

	/** 短信签名 */
	Sign string `mapstructure:"sign" json:"sign" yaml:"sign"`

	/** 资源所有者账户 */
	ResourceOwnerAccount string `mapstructure:"resource-owner-account" json:"resourceOwnerAccount" yaml:"resource_owner_account"`

	/** 资源所有者 ID */
	ResourceOwnerID int64 `mapstructure:"resource-owner-id" json:"resourceOwnerId" yaml:"resource_owner_id"`

	/** 短信模板代码 */
	TemplateCodeVerify string `mapstructure:"template-code-verify" json:"templateCodeVerify" yaml:"template_code_verify"`

	/** 短信服务的 API 端点 */
	Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
}

// NewAliyunSms 创建一个新的 AliyunSms 实例
func NewAliyunSms(moduleName, secretID, secretKey, sign, resourceOwnerAccount string, resourceOwnerID int64, templateCodeVerify, endpoint string) *AliyunSms {
	var aliyunSmsInstance *AliyunSms

	internal.LockFunc(func() {
		aliyunSmsInstance = &AliyunSms{
			ModuleName:           moduleName,
			SecretID:             secretID,
			SecretKey:            secretKey,
			Sign:                 sign,
			ResourceOwnerAccount: resourceOwnerAccount,
			ResourceOwnerID:      resourceOwnerID,
			TemplateCodeVerify:   templateCodeVerify,
			Endpoint:             endpoint,
		}
	})
	return aliyunSmsInstance
}

// ToMap 将配置转换为映射
func (a *AliyunSms) ToMap() map[string]interface{} {
	return internal.ToMap(a)
}

// FromMap 从映射中填充配置
func (a *AliyunSms) FromMap(data map[string]interface{}) {
	internal.FromMap(a, data)
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
