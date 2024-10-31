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
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Alipay 结构体用于配置支付宝支付相关参数
type Alipay struct {
	ModuleName string `mapstructure:"MODULE_NAME"              yaml:"modulename"` // 模块名称
	Pid        string `mapstructure:"PID"                      yaml:"pid"`        // 商户 PID，即商户的账号 ID，在某些业务场景下要用到
	AppId      string `mapstructure:"APP_ID"                   yaml:"app-id"`     // 应用 ID
	PriKey     string `mapstructure:"PRI_KEY"                  yaml:"pri-key"`    // 私钥
	PubKey     string `mapstructure:"PUB_KEY"                  yaml:"pub-key"`    // 公钥，主要是回调验签用
	SignType   string `mapstructure:"SIGN_TYPE"                yaml:"sign-type"`  // 签名方式，商户生成签名字符串所使用的签名算法类型，目前支持 RSA2 和 RSA，私钥 1024 位 RSA
	NotifyUrl  string `mapstructure:"NOTIFY_URL"               yaml:"notify-url"` // 支付宝回调的 URL
	Subject    string `mapstructure:"SUBJECT"                  yaml:"subject"`    // 默认订单标题
}

// NewAlipay 创建一个新的 Alipay 实例
func NewAlipay(opt *Alipay) *Alipay {
	var alipayInstance *Alipay

	internal.LockFunc(func() {
		alipayInstance = opt
	})
	return alipayInstance
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
