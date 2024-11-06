/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 11:36:21
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
	ModuleName  string `mapstructure:"modulename"        yaml:"modulename"      json:"module_name"      validate:"required"`    // 模块名称
	AppId       string `mapstructure:"app-id"             yaml:"app-id"         json:"app_id"          validate:"required"`     // 应用 ID
	MchId       string `mapstructure:"mch-id"             yaml:"mch-id"         json:"mch_id"          validate:"required"`     // 微信商户号
	NotifyUrl   string `mapstructure:"notify-url"         yaml:"notify-url"     json:"notify_url"      validate:"required,url"` // 微信回调的 URL
	ApiKey      string `mapstructure:"api-key"            yaml:"api-key"        json:"api_key"         validate:"required"`     // 签名用的 key
	CertP12Path string `mapstructure:"cert-p12-path"      yaml:"cert-p12-path"  json:"cert_p12_path"   validate:"required"`     // 微信 P12 密钥文件存放位置
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
