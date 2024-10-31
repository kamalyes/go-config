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
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Wechat 结构体用于配置微信支付相关参数
type Wechat struct {
	ModuleName  string `mapstructure:"MODULE_NAME"              yaml:"modulename"`    // 模块名称
	AppId       string `mapstructure:"APP_ID"                   yaml:"app-id"`        // 应用 ID
	MchId       string `mapstructure:"MCH_ID"                   yaml:"mch-id"`        // 微信商户号
	NotifyUrl   string `mapstructure:"NOTIFY_URL"               yaml:"notify-url"`    // 微信回调的 URL
	ApiKey      string `mapstructure:"API_KEY"                  yaml:"api-key"`       // 签名用的 key
	CertP12Path string `mapstructure:"CERT_P12_PATH"            yaml:"cert-p12-path"` // 微信 P12 密钥文件存放位置
}

// NewWechat 创建一个新的 Wechat 实例
func NewWechat(opt *Wechat) *Wechat {
	var wechatInstance *Wechat

	internal.LockFunc(func() {
		wechatInstance = opt
	})
	return wechatInstance
}

// Clone 返回 Wechat 配置的副本
func (w *Wechat) Clone() internal.Configurable {
	return &Wechat{
		ModuleName:  w.ModuleName,
		AppId:       w.AppId,
		MchId:       w.MchId,
		NotifyUrl:   w.NotifyUrl,
		ApiKey:      w.ApiKey,
		CertP12Path: w.CertP12Path,
	}
}

// Get 返回 Wechat 配置的所有字段
func (w *Wechat) Get() interface{} {
	return w
}

// Set 更新 Wechat 配置的字段
func (w *Wechat) Set(data interface{}) {
	if configData, ok := data.(*Wechat); ok {
		w.ModuleName = configData.ModuleName
		w.AppId = configData.AppId
		w.MchId = configData.MchId
		w.NotifyUrl = configData.NotifyUrl
		w.ApiKey = configData.ApiKey
		w.CertP12Path = configData.CertP12Path
	}
}

// Validate 验证 Wechat 配置的有效性
func (w *Wechat) Validate() error {
	if w.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if w.AppId == "" {
		return errors.New("app id cannot be empty")
	}
	if w.MchId == "" {
		return errors.New("merchant id cannot be empty")
	}
	if w.NotifyUrl == "" {
		return errors.New("notify url cannot be empty")
	}
	if w.ApiKey == "" {
		return errors.New("API key cannot be empty")
	}
	if w.CertP12Path == "" {
		return errors.New("certificate P12 path cannot be empty")
	}
	return nil
}
