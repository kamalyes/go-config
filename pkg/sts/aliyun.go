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
	ModuleName      string `mapstructure:"modulename"                yaml:"modulename"           json:"module_name"           validate:"required"` // 模块名称
	RegionID        string `mapstructure:"region-id"                 yaml:"region-id"            json:"region_id"             validate:"required"` // 区域 ID
	AccessKeyID     string `mapstructure:"access-key-id"             yaml:"access-key-id"        json:"access_key_id"         validate:"required"` // 访问密钥 ID
	AccessKeySecret string `mapstructure:"access-key-secret"         yaml:"access-key-secret"    json:"access_key_secret"     validate:"required"` // 访问密钥 Secret
	RoleArn         string `mapstructure:"role-arn"                  yaml:"role-arn"             json:"role_arn"              validate:"required"` // 角色 ARN
	RoleSessionName string `mapstructure:"role-session-name"         yaml:"role-session-name"    json:"role_session_name"     validate:"required"` // 角色会话名称
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
