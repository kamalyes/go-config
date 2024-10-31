/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 15:21:39
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

// Wechat 微信支付配置
type Wechat struct {

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** 应用id */
	AppId string `mapstructure:"app-id"            json:"appId"           yaml:"app-id"`

	/** 微信商户号 */
	MchId string `mapstructure:"mch-id"            json:"mchId"           yaml:"mch-id"`

	/** 微信回调的url */
	NotifyUrl string `mapstructure:"notify-url"        json:"notifyUrl"       yaml:"notify-url"`

	/** 签名用的 key */
	ApiKey string `mapstructure:"api-key"           json:"apiKey"          yaml:"api-key"`

	/** 微信p12密钥文件存放位置 */
	CertP12Path string `mapstructure:"cert-p12-path"     json:"certP12Path"     yaml:"cert-p12-path"`
}

// NewWechat 创建一个新的 Wechat 实例
func NewWechat(moduleName, appId, mchId, notifyUrl, apiKey, certP12Path string) *Wechat {
	var wechatInstance *Wechat

	internal.LockFunc(func() {
		wechatInstance = &Wechat{
			ModuleName:  moduleName,
			AppId:       appId,
			MchId:       mchId,
			NotifyUrl:   notifyUrl,
			ApiKey:      apiKey,
			CertP12Path: certP12Path,
		}
	})
	return wechatInstance
}

// ToMap 将配置转换为映射
func (w *Wechat) ToMap() map[string]interface{} {
	return internal.ToMap(w)
}

// FromMap 从映射中填充配置
func (w *Wechat) FromMap(data map[string]interface{}) {
	internal.FromMap(w, data)
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
