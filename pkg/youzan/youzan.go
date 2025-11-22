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
	Endpoint      string `mapstructure:"endpoint"                 yaml:"endpoint"            json:"endpoint"              validate:"required,url"` // API端点地址，必须是有效的 URL
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
		Endpoint:      y.Endpoint,
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
		y.Endpoint = configData.Endpoint
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

// Default 返回默认的 YouZan 指针，支持链式调用
func Default() *YouZan {
	config := DefaultYouZan()
	return &config
}

// DefaultYouZan 返回默认的 YouZan 值
func DefaultYouZan() YouZan {
	return YouZan{
		Endpoint:      "https://open.youzan.com",
		ClientID:      "demo_client_id",
		ClientSecret:  "demo_client_secret",
		AuthorizeType: "silent",
		GrantID:       "demo_grant_id",
		Refresh:       true,
		ModuleName:    "youzan",
	}
}

// DefaultYouZanConfig 返回默认的 YouZan 指针，支持链式调用
func DefaultYouZanConfig() *YouZan {
	config := DefaultYouZan()
	return &config
}

// WithEndpoint 设置API端点地址
func (y *YouZan) WithEndpoint(endpoint string) *YouZan {
	y.Endpoint = endpoint
	return y
}

// WithClientID 设置客户端 ID
func (y *YouZan) WithClientID(clientID string) *YouZan {
	y.ClientID = clientID
	return y
}

// WithClientSecret 设置客户端密钥
func (y *YouZan) WithClientSecret(clientSecret string) *YouZan {
	y.ClientSecret = clientSecret
	return y
}

// WithAuthorizeType 设置授权类型
func (y *YouZan) WithAuthorizeType(authorizeType string) *YouZan {
	y.AuthorizeType = authorizeType
	return y
}

// WithGrantID 设置授权 ID
func (y *YouZan) WithGrantID(grantID string) *YouZan {
	y.GrantID = grantID
	return y
}

// WithRefresh 设置是否刷新
func (y *YouZan) WithRefresh(refresh bool) *YouZan {
	y.Refresh = refresh
	return y
}

// WithModuleName 设置模块名称
func (y *YouZan) WithModuleName(moduleName string) *YouZan {
	y.ModuleName = moduleName
	return y
}
