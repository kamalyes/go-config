/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 01:14:39
 * @FilePath: \go-config\pkg\pay\wechat.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package pay

import (
	"github.com/kamalyes/go-config/internal"
)

// WechatPay 结构体用于配置微信支付相关参数
type WechatPay struct {
	AppId       string `mapstructure:"app-id" yaml:"app-id" json:"appId"          validate:"required"`                // 应用 ID
	MchId       string `mapstructure:"mch-id" yaml:"mch-id" json:"mchId"          validate:"required"`                // 微信商户号
	NotifyUrl   string `mapstructure:"notify-url" yaml:"notify-url" json:"notifyUrl"      validate:"required,url"`    // 微信回调的 URL
	ApiKey      string `mapstructure:"api-key" yaml:"api-key" json:"apiKey"         validate:"required"`              // 签名用的 key
	CertP12Path string `mapstructure:"cert_p12_path" yaml:"cert-p12-path" json:"cert_p12_path"   validate:"required"` // 微信 P12 密钥文件存放位置
	ModuleName  string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                              // 模块名称
}

// NewWechatPay 创建一个新的 Wechat 实例
func NewWechatPay(opt *WechatPay) *WechatPay {
	var wechatInstance *WechatPay

	internal.LockFunc(func() {
		wechatInstance = opt
	})
	return wechatInstance
}

// Clone 返回 WechatPay 配置的副本
func (w *WechatPay) Clone() internal.Configurable {
	return &WechatPay{
		ModuleName:  w.ModuleName,
		AppId:       w.AppId,
		MchId:       w.MchId,
		NotifyUrl:   w.NotifyUrl,
		ApiKey:      w.ApiKey,
		CertP12Path: w.CertP12Path,
	}
}

// Get 返回 WechatPay 配置的所有字段
func (w *WechatPay) Get() interface{} {
	return w
}

// Set 更新 WechatPay 配置的字段
func (w *WechatPay) Set(data interface{}) {
	if configData, ok := data.(*WechatPay); ok {
		w.ModuleName = configData.ModuleName
		w.AppId = configData.AppId
		w.MchId = configData.MchId
		w.NotifyUrl = configData.NotifyUrl
		w.ApiKey = configData.ApiKey
		w.CertP12Path = configData.CertP12Path
	}
}

// Validate 验证 WechatPay 配置的有效性
func (w *WechatPay) Validate() error {
	return internal.ValidateStruct(w)
}

// DefaultWechatPay 返回默认WechatPay配置
func DefaultWechatPay() WechatPay {
	return WechatPay{
		ModuleName:  "wechatpay",
		AppId:       "",
		MchId:       "",
		NotifyUrl:   "",
		ApiKey:      "",
		CertP12Path: "",
	}
}

// DefaultWechatPayConfig 返回默认WechatPay配置的指针，支持链式调用
func DefaultWechatPayConfig() *WechatPay {
	config := DefaultWechatPay()
	return &config
}

// WithModuleName 设置模块名称
func (w *WechatPay) WithModuleName(moduleName string) *WechatPay {
	w.ModuleName = moduleName
	return w
}

// WithAppID 设置应用ID
func (w *WechatPay) WithAppID(appId string) *WechatPay {
	w.AppId = appId
	return w
}

// WithMchID 设置微信商户号
func (w *WechatPay) WithMchID(mchId string) *WechatPay {
	w.MchId = mchId
	return w
}

// WithNotifyURL 设置回调URL
func (w *WechatPay) WithNotifyURL(notifyURL string) *WechatPay {
	w.NotifyUrl = notifyURL
	return w
}

// WithApiKey 设置签名密钥
func (w *WechatPay) WithApiKey(apiKey string) *WechatPay {
	w.ApiKey = apiKey
	return w
}

// WithCertP12Path 设置P12证书文件路径
func (w *WechatPay) WithCertP12Path(certP12Path string) *WechatPay {
	w.CertP12Path = certP12Path
	return w
}
