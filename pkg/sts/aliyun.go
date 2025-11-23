/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-02 23:55:51
 * @FilePath: \go-config\pkg\sts\aliyun.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package sts

import (
	"github.com/kamalyes/go-config/internal"
)

// AliyunSts 结构体表示 STS 配置
type AliyunSts struct {
	ModuleName      string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"           validate:"required"`            // 模块名称
	RegionID        string `mapstructure:"region-id" yaml:"region-id" json:"regionId"             validate:"required"`                // 区域 ID
	AccessKeyID     string `mapstructure:"access-key-id" yaml:"access-key-id" json:"accessKeyId"         validate:"required"`         // 访问密钥 ID
	AccessKeySecret string `mapstructure:"access-key-secret" yaml:"access-key-secret" json:"accessKeySecret"     validate:"required"` // 访问密钥 Secret
	RoleArn         string `mapstructure:"role-arn" yaml:"role-arn" json:"roleArn"              validate:"required"`                  // 角色 ARN
	RoleSessionName string `mapstructure:"role-session-name" yaml:"role-session-name" json:"roleSessionName"     validate:"required"` // 角色会话名称
}

// NewAliyunSts 创建一个新的 AliyunSts 实例
func NewAliyunSts(opt *AliyunSts) *AliyunSts {
	var aliyunStsInstance *AliyunSts

	internal.LockFunc(func() {
		aliyunStsInstance = opt
	})
	return aliyunStsInstance
}

// Clone 返回 AliyunSts 配置的副本
func (a *AliyunSts) Clone() internal.Configurable {
	return &AliyunSts{
		ModuleName:      a.ModuleName,
		RegionID:        a.RegionID,
		AccessKeyID:     a.AccessKeyID,
		AccessKeySecret: a.AccessKeySecret,
		RoleArn:         a.RoleArn,
		RoleSessionName: a.RoleSessionName,
	}
}

// Get 返回 AliyunSts 配置的所有字段
func (a *AliyunSts) Get() interface{} {
	return a
}

// Set 更新 AliyunSts 配置的字段
func (a *AliyunSts) Set(data interface{}) {
	if configData, ok := data.(*AliyunSts); ok {
		a.ModuleName = configData.ModuleName
		a.RegionID = configData.RegionID
		a.AccessKeyID = configData.AccessKeyID
		a.AccessKeySecret = configData.AccessKeySecret
		a.RoleArn = configData.RoleArn
		a.RoleSessionName = configData.RoleSessionName
	}
}

// Validate 验证 AliyunSts 配置的有效性
func (a *AliyunSts) Validate() error {
	return internal.ValidateStruct(a)
}

// Default 返回默认的 AliyunSts 指针，支持链式调用
func Default() *AliyunSts {
	config := DefaultAliyunSts()
	return &config
}

// DefaultAliyunSts 返回默认的 AliyunSts 值
func DefaultAliyunSts() AliyunSts {
	return AliyunSts{
		ModuleName:      "aliyun-sts",
		RegionID:        "cn-hangzhou",
		AccessKeyID:     "demo_access_key_id",
		AccessKeySecret: "demo_access_key_secret",
		RoleArn:         "acs:ram::account:role/demo-role",
		RoleSessionName: "default-session",
	}
}

// DefaultAliyunStsConfig 返回默认的 AliyunSts 指针，支持链式调用
func DefaultAliyunStsConfig() *AliyunSts {
	config := DefaultAliyunSts()
	return &config
}

// WithModuleName 设置模块名称
func (a *AliyunSts) WithModuleName(moduleName string) *AliyunSts {
	a.ModuleName = moduleName
	return a
}

// WithRegionID 设置区域 ID
func (a *AliyunSts) WithRegionID(regionID string) *AliyunSts {
	a.RegionID = regionID
	return a
}

// WithAccessKeyID 设置访问密钥 ID
func (a *AliyunSts) WithAccessKeyID(accessKeyID string) *AliyunSts {
	a.AccessKeyID = accessKeyID
	return a
}

// WithAccessKeySecret 设置访问密钥 Secret
func (a *AliyunSts) WithAccessKeySecret(accessKeySecret string) *AliyunSts {
	a.AccessKeySecret = accessKeySecret
	return a
}

// WithRoleArn 设置角色 ARN
func (a *AliyunSts) WithRoleArn(roleArn string) *AliyunSts {
	a.RoleArn = roleArn
	return a
}

// WithRoleSessionName 设置角色会话名称
func (a *AliyunSts) WithRoleSessionName(roleSessionName string) *AliyunSts {
	a.RoleSessionName = roleSessionName
	return a
}
