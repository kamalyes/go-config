/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:53:07
 * @FilePath: \go-config\pkg\cors\cors.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package cors

import (
	"github.com/kamalyes/go-config/internal"
)

// Cors 跨域配置结构体
type Cors struct {
	ModuleName          string   `mapstructure:"MODULE_NAME"   yaml:"modulename"`                    // 模块名称
	AllowedAllOrigins   bool     `mapstructure:"ALLOWED_ALL_ORIGINS" yaml:"allowed-all-origins"`     // 是否允许所有来源
	AllowedAllMethods   bool     `mapstructure:"ALLOWED_ALL_METHODS" yaml:"allowed-all-methods"`     // 是否允许所有方法
	AllowedOrigins      []string `mapstructure:"ALLOWED_ORIGINS" yaml:"allowed-origins"`             // 允许的来源
	AllowedMethods      []string `mapstructure:"ALLOWED_METHODS" yaml:"allowed-methods"`             // 允许的方法
	AllowedHeaders      []string `mapstructure:"ALLOWED_HEADERS" yaml:"allowed-headers"`             // 允许的头部
	MaxAge              string   `mapstructure:"MAX_AGE" yaml:"max-age"`                             // 最大缓存时间
	ExposedHeaders      []string `mapstructure:"EXPOSED_HEADERS" yaml:"exposed-headers"`             // 暴露的头部
	AllowCredentials    bool     `mapstructure:"ALLOW_CREDENTIALS" yaml:"allow-credentials"`         // 允许凭证
	OptionsResponseCode int      `mapstructure:"OPTIONS_RESPONSE_CODE" yaml:"options-response-code"` // Options响应Code
}

// NewCors 创建一个新的 Cors 实例
func NewCors(opt *Cors) *Cors {
	var corsInstance *Cors

	internal.LockFunc(func() {
		corsInstance = opt
	})
	return corsInstance
}

// Clone 返回 Cors 配置的副本
func (c *Cors) Clone() internal.Configurable {
	return &Cors{
		ModuleName:          c.ModuleName,
		AllowedAllOrigins:   c.AllowedAllOrigins,
		AllowedAllMethods:   c.AllowedAllMethods,
		AllowedOrigins:      c.AllowedOrigins,
		AllowedMethods:      c.AllowedMethods,
		AllowedHeaders:      c.AllowedHeaders,
		MaxAge:              c.MaxAge,
		ExposedHeaders:      c.ExposedHeaders,
		AllowCredentials:    c.AllowCredentials,
		OptionsResponseCode: c.OptionsResponseCode,
	}
}

// Get 返回 Cors 配置的所有字段
func (c *Cors) Get() interface{} {
	return c
}

// Set 更新 Cors 配置的字段
func (c *Cors) Set(data interface{}) {
	if configData, ok := data.(*Cors); ok {
		c.ModuleName = configData.ModuleName
		c.AllowedAllOrigins = configData.AllowedAllOrigins
		c.AllowedAllMethods = configData.AllowedAllMethods
		c.AllowedOrigins = configData.AllowedOrigins
		c.AllowedMethods = configData.AllowedMethods
		c.AllowedHeaders = configData.AllowedHeaders
		c.MaxAge = configData.MaxAge
		c.ExposedHeaders = configData.ExposedHeaders
		c.AllowCredentials = configData.AllowCredentials
		c.OptionsResponseCode = configData.OptionsResponseCode
	}
}

// Validate 检查 Cors 配置的有效性
func (c *Cors) Validate() error {
	return nil
}
