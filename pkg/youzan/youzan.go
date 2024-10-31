/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:37:33
 * @FilePath: \go-config\pkg\youzan\youzan.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package youzan

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// YouZan 结构体表示 YouZan 配置
type YouZan struct {
	ModuleName    string `mapstructure:"MODULE_NAME"              yaml:"modulename"`     // 模块名称
	Host          string `mapstructure:"HOST"                     yaml:"host"`           // Host
	ClientID      string `mapstructure:"CLIENT_ID"                yaml:"client-id"`      // 客户端ID
	ClientSecret  string `mapstructure:"CLIENT_SECRET"            yaml:"client-secret"`  // 客户端密钥
	AuthorizeType string `mapstructure:"AUTHORIZE_TYPE"           yaml:"authorize-type"` // 授权类型
	GrantID       string `mapstructure:"GRANT_ID"                 yaml:"grant-id"`       // 授权ID
	Refresh       bool   `mapstructure:"REFRESH"                  yaml:"refresh"`        // 是否刷新
}

// NewYouZan 创建一个新的 YouZan 实例
func NewYouZan(opt *YouZan) *YouZan {
	var youzanInstance *YouZan

	internal.LockFunc(func() {
		youzanInstance = opt
	})
	return youzanInstance
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
