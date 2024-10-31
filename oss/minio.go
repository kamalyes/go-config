/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 22:35:17
 * @FilePath: \go-config\oss\minio.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package oss

// Minio 配置文件
type Minio struct {
	/** 模块名称 */
	ModuleName string `mapstructure:"modulename" json:"moduleName" yaml:"modulename"`

	/** Host */
	Host string `mapstructure:"host" json:"host" yaml:"host"`

	/** 端口 */
	Port int `mapstructure:"port" json:"port" yaml:"port"`

	/** 签名用的 key */
	AccessKey string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`

	/** 签名用的钥匙 */
	SecretKey string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
}

// 实现 Configurable 接口
func (m Minio) GetModuleName() string {
	return "minio"
}
