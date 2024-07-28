/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 09:33:47
 * @FilePath: \go-config\pay\wechat.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package pay

// Wechat 微信支付配置
type Wechat struct {

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
