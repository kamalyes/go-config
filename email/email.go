/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 09:24:50
 * @FilePath: \go-config\email\email.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package email

// Email 邮件配置
type Email struct {

	/** 收件人:多个以英文逗号分隔 */
	To string `mapstructure:"to"         json:"to"           yaml:"to"`

	/** 设置发件人 */
	From string `mapstructure:"from"       json:"from"         yaml:"from"`

	/** 邮件服务器地址 */
	Host string `mapstructure:"host"       json:"host"         yaml:"host"`

	/** 端口 */
	Port int `mapstructure:"port"       json:"port"         yaml:"port"`

	/** 是否SSL */
	IsSSL bool `mapstructure:"is-ssl"     json:"isSSL"        yaml:"is-ssl"` // 是否SSL

	/** 密钥 */
	Secret string `mapstructure:"secret"     json:"secret"       yaml:"secret"`
}
