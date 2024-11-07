/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:54:29
 * @FilePath: \go-config\pkg\youzan\youzan.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package youzan

import (
	"github.com/kamalyes/go-config/internal"
)

// YouZan 结构体表示 YouZan 配置
type YouZan struct {
	Host          string `mapstructure:"host"                     yaml:"host"                json:"host"                  validate:"required,url"` // Host，必须是有效的 URL
	ClientID      string `mapstructure:"client-id"                yaml:"client-id"           json:"client_id"             validate:"required"`     // 客户端 ID
	ClientSecret  string `mapstructure:"client-secret"            yaml:"client-secret"       json:"client_secret"         validate:"required"`     // 客户端密钥
	AuthorizeType string `mapstructure:"authorize-type"           yaml:"authorize-type"      json:"authorize_type"        validate:"required"`     // 授权类型
	GrantID       string `mapstructure:"grant-id"                 yaml:"grant-id"            json:"grant_id"              validate:"required"`     // 授权 ID
	Refresh       bool   `mapstructure:"refresh"                  yaml:"refresh"             json:"refresh"`                                       // 是否刷新
	ModuleName    string `mapstructure:"modulename"               yaml:"modulename"          json:"module_name"`                                   // 模块名称
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
	return internal.ValidateStruct(y)
}
