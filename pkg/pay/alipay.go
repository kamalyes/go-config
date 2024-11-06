/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:39:39
 * @FilePath: \go-config\pkg\pay\alipay.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package pay

import (
	"github.com/kamalyes/go-config/internal"
)

// AliPay 结构体用于配置支付宝支付相关参数
type AliPay struct {
	ModuleName string `mapstructure:"modulename"        yaml:"modulename"     json:"module_name"     validate:"required"`                 // 模块名称
	Pid        string `mapstructure:"pid"                yaml:"pid"            json:"pid"             validate:"required"`                // 商户 PID，即商户的账号 ID
	AppId      string `mapstructure:"app-id"             yaml:"app-id"         json:"app_id"          validate:"required"`                // 应用 ID
	PriKey     string `mapstructure:"pri-key"            yaml:"pri-key"        json:"pri_key"         validate:"required"`                // 私钥
	PubKey     string `mapstructure:"pub-key"            yaml:"pub-key"        json:"pub_key"         validate:"required"`                // 公钥，主要是回调验签用
	SignType   string `mapstructure:"sign-type"          yaml:"sign-type"      json:"sign_type"       validate:"required,oneof=RSA2 RSA"` // 签名方式，支持 RSA2 和 RSA
	NotifyUrl  string `mapstructure:"notify-url"         yaml:"notify-url"     json:"notify_url"      validate:"required,url"`            // 支付宝回调的 URL
	Subject    string `mapstructure:"subject"            yaml:"subject"        json:"subject"`                                            // 默认订单标题
}

// NewAliPay 创建一个新的 AliPay 实例
func NewAliPay(opt *AliPay) *AliPay {
	var alipayInstance *AliPay

	internal.LockFunc(func() {
		alipayInstance = opt
	})
	return alipayInstance
}

// Clone 返回 AliPay 配置的副本
func (a *AliPay) Clone() internal.Configurable {
	return &AliPay{
		ModuleName: a.ModuleName,
		Pid:        a.Pid,
		AppId:      a.AppId,
		PriKey:     a.PriKey,
		PubKey:     a.PubKey,
		SignType:   a.SignType,
		NotifyUrl:  a.NotifyUrl,
		Subject:    a.Subject,
	}
}

// Get 返回 AliPay 配置的所有字段
func (a *AliPay) Get() interface{} {
	return a
}

// Set 更新 AliPay 配置的字段
func (a *AliPay) Set(data interface{}) {
	if configData, ok := data.(*AliPay); ok {
		a.ModuleName = configData.ModuleName
		a.Pid = configData.Pid
		a.AppId = configData.AppId
		a.PriKey = configData.PriKey
		a.PubKey = configData.PubKey
		a.SignType = configData.SignType
		a.NotifyUrl = configData.NotifyUrl
		a.Subject = configData.Subject
	}
}

// Validate 验证 AliPay 配置的有效性
func (a *AliPay) Validate() error {
	return internal.ValidateStruct(a)
}
