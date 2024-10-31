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

// 实现 Configurable 接口
func (y YouZan) GetModuleName() string {
	return "youzan"
}
