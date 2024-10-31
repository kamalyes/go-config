/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:50:21
 * @FilePath: \go-config\pay\alipay.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package pay

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Alipay 支付宝 支付相关参数配置
type Alipay struct {

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** 商户pid 即商户的账号id，在某些业务场景下要用到*/
	Pid string `mapstructure:"pid"               json:"pid"             yaml:"pid"`

	/** 应用id */
	AppId string `mapstructure:"app-id"            json:"appId"           yaml:"app-id"`

	/** 私钥 */
	PriKey string `mapstructure:"pri-key"           json:"priKey"          yaml:"pri-key"`

	/** 公钥 主要是回调 验签用 */
	PubKey string `mapstructure:"pub-key"           json:"pubKey"          yaml:"pub-key"`

	/** 签名方式  商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，私钥 1024位RSA */
	SignType string `mapstructure:"sign-type"         json:"signType"        yaml:"sign-type"`

	/** 支付宝回调的url */
	NotifyUrl string `mapstructure:"notify-url"        json:"notifyUrl"       yaml:"notify-url"`

	/** 默认订单标题 */
	Subject string `mapstructure:"subject"           json:"subject"         yaml:"subject"`
}

// NewAlipay 创建一个新的 Alipay 实例
func NewAlipay(moduleName, pid, appId, priKey, pubKey, signType, notifyUrl, subject string) *Alipay {
	var alipayInstance *Alipay

	internal.LockFunc(func() {
		alipayInstance = &Alipay{
			ModuleName: moduleName,
			Pid:        pid,
			AppId:      appId,
			PriKey:     priKey,
			PubKey:     pubKey,
			SignType:   signType,
			NotifyUrl:  notifyUrl,
			Subject:    subject,
		}
	})
	return alipayInstance
}

// ToMap 将配置转换为映射
func (a *Alipay) ToMap() map[string]interface{} {
	return internal.ToMap(a)
}

// FromMap 从映射中填充配置
func (a *Alipay) FromMap(data map[string]interface{}) {
	internal.FromMap(a, data)
}

// Clone 返回 Alipay 配置的副本
func (a *Alipay) Clone() internal.Configurable {
	return &Alipay{
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

// Get 返回 Alipay 配置的所有字段
func (a *Alipay) Get() interface{} {
	return a
}

// Set 更新 Alipay 配置的字段
func (a *Alipay) Set(data interface{}) {
	if configData, ok := data.(*Alipay); ok {
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

// Validate 验证 Alipay 配置的有效性
func (a *Alipay) Validate() error {
	if a.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if a.Pid == "" {
		return errors.New("pid cannot be empty")
	}
	if a.AppId == "" {
		return errors.New("app id cannot be empty")
	}
	if a.PriKey == "" {
		return errors.New("private key cannot be empty")
	}
	if a.PubKey == "" {
		return errors.New("public key cannot be empty")
	}
	if a.SignType == "" {
		return errors.New("sign type cannot be empty")
	}
	if a.NotifyUrl == "" {
		return errors.New("notify url cannot be empty")
	}
	if a.Subject == "" {
		return errors.New("subject cannot be empty")
	}
	return nil
}
