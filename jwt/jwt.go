/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 09:24:50
 * @FilePath: \go-config\captcha\captcha.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package jwt

type JWT struct {

	/** jwt签名 */
	SigningKey string `mapstructure:"signing-key"         json:"signingKey"        yaml:"signing-key"`

	/** 过期时间 */
	ExpiresTime int64 `mapstructure:"expires-time"        json:"expiresTime"       yaml:"expires-time"`

	/** 缓冲时间 */
	BufferTime int64 `mapstructure:"buffer-time"         json:"bufferTime"        yaml:"buffer-time"`

	/** 多地登录拦截 true 拦截 fasle 不拦截 */
	UseMultipoint bool `mapstructure:"use-multipoint"      json:"useMultipoint"     yaml:"use-multipoint"`
}
