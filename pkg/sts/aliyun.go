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

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

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

// NewAliyunSts 创建一个新的 AliyunSts 实例
func NewAliyunSts(moduleName, regionID, accessKeyID, accessKeySecret, roleArn, roleSessionName string) *AliyunSts {
	var aliyunStsInstance *AliyunSts

	internal.LockFunc(func() {
		aliyunStsInstance = &AliyunSts{
			ModuleName:      moduleName,
			RegionID:        regionID,
			AccessKeyID:     accessKeyID,
			AccessKeySecret: accessKeySecret,
			RoleArn:         roleArn,
			RoleSessionName: roleSessionName,
		}
	})
	return aliyunStsInstance
}

// ToMap 将配置转换为映射
func (a *AliyunSts) ToMap() map[string]interface{} {
	return internal.ToMap(a)
}

// FromMap 从映射中填充配置
func (a *AliyunSts) FromMap(data map[string]interface{}) {
	internal.FromMap(a, data)
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
