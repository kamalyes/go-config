/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 12:50:58
 * @FilePath: \go-config\youzan\youzan.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package youzan

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// YouZan 配置文件
type YouZan struct {
	/** 模块名称 */
	ModuleName string `mapstructure:"modulename" json:"moduleName" yaml:"modulename"`

	/** Host */
	Host string `mapstructure:"host" json:"host" yaml:"host"`

	/** 客户端ID */
	ClientID string `mapstructure:"client-id" json:"clientId" yaml:"client-id"`

	/** 客户端密钥 */
	ClientSecret string `mapstructure:"client-secret" json:"clientSecret" yaml:"client-secret"`

	/** 授权类型 */
	AuthorizeType string `mapstructure:"authorize-type" json:"authorizeType" yaml:"authorize-type"`

	/** 授权ID */
	GrantID string `mapstructure:"grant-id" json:"grantId" yaml:"grant-id"`

	/** 是否刷新 */
	Refresh bool `mapstructure:"refresh" json:"refresh" yaml:"refresh"`
}

// NewYouZan 创建一个新的 YouZan 实例
func NewYouZan(moduleName, host, clientID, clientSecret, authorizeType, grantID string, refresh bool) *YouZan {
	var youzanInstance *YouZan

	internal.LockFunc(func() {
		youzanInstance = &YouZan{
			ModuleName:    moduleName,
			Host:          host,
			ClientID:      clientID,
			ClientSecret:  clientSecret,
			AuthorizeType: authorizeType,
			GrantID:       grantID,
			Refresh:       refresh,
		}
	})
	return youzanInstance
}

// ToMap 将配置转换为映射
func (y *YouZan) ToMap() map[string]interface{} {
	return internal.ToMap(y)
}

// FromMap 从映射中填充配置
func (y *YouZan) FromMap(data map[string]interface{}) {
	internal.FromMap(y, data)
}

// Clone 返回 YouZan 配置的副本
func (y *YouZan) Clone() internal.Configurable {
	return &YouZan{
		ModuleName:    y.ModuleName,
		Host:          y.Host,
		ClientID:      y.ClientID,
		ClientSecret:  y.ClientSecret,
		AuthorizeType: y.AuthorizeType,
		GrantID:       y.GrantID,
		Refresh:       y.Refresh,
	}
}

// Get 返回 YouZan 配置的所有字段
func (y *YouZan) Get() interface{} {
	return y
}

// Set 更新 YouZan 配置的字段
func (y *YouZan) Set(data interface{}) {
	if configData, ok := data.(*YouZan); ok {
		y.ModuleName = configData.ModuleName
		y.Host = configData.Host
		y.ClientID = configData.ClientID
		y.ClientSecret = configData.ClientSecret
		y.AuthorizeType = configData.AuthorizeType
		y.GrantID = configData.GrantID
		y.Refresh = configData.Refresh
	}
}

// Validate 验证 YouZan 配置的有效性
func (y *YouZan) Validate() error {
	if y.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if y.Host == "" {
		return errors.New("host cannot be empty")
	}
	if y.ClientID == "" {
		return errors.New("client ID cannot be empty")
	}
	if y.ClientSecret == "" {
		return errors.New("client secret cannot be empty")
	}
	if y.AuthorizeType == "" {
		return errors.New("authorize type cannot be empty")
	}
	if y.GrantID == "" {
		return errors.New("grant ID cannot be empty")
	}
	return nil
}
