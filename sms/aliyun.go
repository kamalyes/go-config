/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 13:51:19
 * @FilePath: \go-config\sms\aliyun.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package sms

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

// 实现 Configurable 接口
func (a AliyunSms) GetModuleName() string {
	return "aliyunsms"
}
