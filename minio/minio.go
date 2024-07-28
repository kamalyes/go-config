/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 22:35:17
 * @FilePath: \go-config\minio\minio.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package minio

// Minio 配置文件
type Minio struct {

	/** Mqtt代理服务器的Ip和端口 */
	Url string `mapstructure:"url"     json:"url"           yaml:"url"`

	/** 端口 */
	Port int `mapstructure:"port"       json:"port"         yaml:"port"`

	/** 签名用的 key */
	AccessKey string `mapstructure:"access-key"   json:"access-key"   yaml:"access-key"`

	/** 签名用的 钥匙 */
	SecretKey string `mapstructure:"secret-key"   json:"secret-key"    yaml:"secret-key"`
}
