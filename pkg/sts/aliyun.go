/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-02 23:49:51
 * @FilePath: \go-config\pkg\sts\aliyun.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package sts

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// AliyunSts 结构体表示 STS 配置
type AliyunSts struct {
	ModuleName      string `mapstructure:"MODULE_NAME"               yaml:"modulename"`        // 模块名称
	RegionID        string `mapstructure:"REGION_ID"                 yaml:"region-id"`         // 区域 ID
	AccessKeyID     string `mapstructure:"ACCESS_KEY_ID"             yaml:"access-key-id"`     // 访问密钥 ID
	AccessKeySecret string `mapstructure:"ACCESS_KEY_SECRET"         yaml:"access-key-secret"` // 访问密钥 Secret
	RoleArn         string `mapstructure:"ROLE_ARN"                  yaml:"role-arn"`          // 角色 ARN
	RoleSessionName string `mapstructure:"ROLE_SESSION_NAME"         yaml:"role-session-name"` // 角色会话名称
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
	if a.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if a.RegionID == "" {
		return errors.New("region ID cannot be empty")
	}
	if a.AccessKeyID == "" {
		return errors.New("access key ID cannot be empty")
	}
	if a.AccessKeySecret == "" {
		return errors.New("access key secret cannot be empty")
	}
	if a.RoleArn == "" {
		return errors.New("role ARN cannot be empty")
	}
	if a.RoleSessionName == "" {
		return errors.New("role session name cannot be empty")
	}
	return nil
}
