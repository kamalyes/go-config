/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 01:13:53
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
	Pid        string `mapstructure:"pid" yaml:"pid" json:"pid"             validate:"required"`                           // 商户 PID，即商户的账号 ID
	AppId      string `mapstructure:"app-id" yaml:"app-id" json:"appId"          validate:"required"`                      // 应用 ID
	PriKey     string `mapstructure:"pri-key" yaml:"pri-key" json:"priKey"         validate:"required"`                    // 私钥
	PubKey     string `mapstructure:"pub-key" yaml:"pub-key" json:"pubKey"         validate:"required"`                    // 公钥，主要是回调验签用
	SignType   string `mapstructure:"sign-type" yaml:"sign-type" json:"signType"       validate:"required,oneof=RSA2 RSA"` // 签名方式，支持 RSA2 和 RSA
	NotifyUrl  string `mapstructure:"notify-url" yaml:"notify-url" json:"notifyUrl"      validate:"required,url"`          // 支付宝回调的 URL
	Subject    string `mapstructure:"subject" yaml:"subject" json:"subject"`                                               // 默认订单标题
	ModuleName string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                                    // 模块名称
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

// DefaultAliPay 返回默认AliPay配置
func DefaultAliPay() AliPay {
	return AliPay{
		ModuleName: "alipay",
		Pid:        "",
		AppId:      "",
		PriKey:     "",
		PubKey:     "",
		SignType:   "RSA2",
		NotifyUrl:  "",
		Subject:    "默认订单",
	}
}

// DefaultAliPayConfig 返回默认AliPay配置的指针，支持链式调用
func DefaultAliPayConfig() *AliPay {
	config := DefaultAliPay()
	return &config
}

// WithModuleName 设置模块名称
func (a *AliPay) WithModuleName(moduleName string) *AliPay {
	a.ModuleName = moduleName
	return a
}

// WithPid 设置商户PID
func (a *AliPay) WithPid(pid string) *AliPay {
	a.Pid = pid
	return a
}

// WithAppID 设置应用ID
func (a *AliPay) WithAppID(appID string) *AliPay {
	a.AppId = appID
	return a
}

// WithPrivateKey 设置私钥
func (a *AliPay) WithPrivateKey(priKey string) *AliPay {
	a.PriKey = priKey
	return a
}

// WithPublicKey 设置公钥
func (a *AliPay) WithPublicKey(pubKey string) *AliPay {
	a.PubKey = pubKey
	return a
}

// WithSignType 设置签名方式
func (a *AliPay) WithSignType(signType string) *AliPay {
	a.SignType = signType
	return a
}

// WithNotifyURL 设置回调URL
func (a *AliPay) WithNotifyURL(notifyURL string) *AliPay {
	a.NotifyUrl = notifyURL
	return a
}

// WithSubject 设置默认订单标题
func (a *AliPay) WithSubject(subject string) *AliPay {
	a.Subject = subject
	return a
}
