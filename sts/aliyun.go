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

package sts

// AliyunSts 配置文件
type AliyunSts struct {
	/** 模块名称 */
	ModuleName string `mapstructure:"modulename" json:"moduleName" yaml:"modulename"`

	/** 区域 ID */
	RegionID string `mapstructure:"region-id" json:"regionId" yaml:"region-id"`

	/** 访问密钥 ID */
	AccessKeyID string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`

	/** 访问密钥 Secret */
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`

	/** 角色 ARN */
	RoleArn string `mapstructure:"role-arn" json:"roleArn" yaml:"role-arn"`

	/** 角色会话名称 */
	RoleSessionName string `mapstructure:"role-session-name" json:"roleSessionName" yaml:"role-session-name"`
}

// 实现 Configurable 接口
func (a AliyunSts) GetModuleName() string {
	return "aliyunsts"
}
