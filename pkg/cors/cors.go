/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 10:05:55
 * @FilePath: \go-config\cors\cors.go
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

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	// 是否允许所有来源
	AllowedAllOrigins bool `mapstructure:"allowed-all-origins" json:"allowedAllOrigins" yaml:"allowed-all-origins"`

	// 是否允许所有方法
	AllowedAllMethods bool `mapstructure:"allowed-all-methods" json:"allowedAllMethods" yaml:"allowed-all-methods"`

	// 允许的来源
	AllowedOrigins []string `mapstructure:"allowed-origins" json:"allowedOrigins" yaml:"allowed-origins"`

	// 允许的方法
	AllowedMethods []string `mapstructure:"allowed-methods" json:"allowedMethods" yaml:"allowed-methods"`

	// 允许的头部
	AllowedHeaders []string `mapstructure:"allowed-headers" json:"allowedHeaders" yaml:"allowed-headers"`

	// 最大缓存时间
	MaxAge string `mapstructure:"max-age" json:"maxAge" yaml:"max-age"`

	// 暴露的头部
	ExposedHeaders []string `mapstructure:"exposed-headers" json:"exposedHeaders" yaml:"exposed-headers"`

	// 允许凭证
	AllowCredentials bool `mapstructure:"allow-credentials" json:"allowCredentials" yaml:"allow-credentials"`

	// Options响应Code
	OptionsResponseCode int `mapstructure:"options-response-code" json:"optionsResponseCode" yaml:"options-response-code"`
}

// NewCors 创建一个新的 Cors 实例
func NewCors(moduleName string, allowedAllOrigins, allowedAllMethods bool, allowedOrigins, allowedMethods, allowedHeaders, exposedHeaders []string, maxAge string, allowCredentials bool, optionsResponseCode int) *Cors {
	var corsInstance *Cors

	internal.LockFunc(func() {
		corsInstance = &Cors{
			ModuleName:          moduleName,
			AllowedAllOrigins:   allowedAllOrigins,
			AllowedAllMethods:   allowedAllMethods,
			AllowedOrigins:      allowedOrigins,
			AllowedMethods:      allowedMethods,
			AllowedHeaders:      allowedHeaders,
			MaxAge:              maxAge,
			ExposedHeaders:      exposedHeaders,
			AllowCredentials:    allowCredentials,
			OptionsResponseCode: optionsResponseCode,
		}
	})
	return corsInstance
}

// ToMap 将配置转换为映射
func (c *Cors) ToMap() map[string]interface{} {
	return internal.ToMap(c)
}

// FromMap 从映射中填充配置
func (c *Cors) FromMap(data map[string]interface{}) {
	internal.FromMap(c, data)
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
