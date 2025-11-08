/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:55:58
 * @FilePath: \go-config\pkg\cors\cors.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package cors

import (
	"github.com/kamalyes/go-config/internal"
)

// Cors 跨域资源共享配置
type Cors struct {
	AllowedOrigins      []string `mapstructure:"allowed-origins"       yaml:"allowed-origins"       json:"allowed_origins"       validate:"dive,required"` // 允许的来源
	AllowedMethods      []string `mapstructure:"allowed-methods"       yaml:"allowed-methods"       json:"allowed_methods"       validate:"dive,required"` // 允许的方法
	AllowedHeaders      []string `mapstructure:"allowed-headers"       yaml:"allowed-headers"       json:"allowed_headers"       validate:"dive,required"` // 允许的头部
	MaxAge              string   `mapstructure:"max-age"               yaml:"max-age"               json:"max_age"               validate:"required"`      // 最大缓存时间
	AllowedAllOrigins   bool     `mapstructure:"allowed-all-origins"   yaml:"allowed-all-origins"   json:"allowed_all_origins"`                            // 是否允许所有来源
	AllowedAllMethods   bool     `mapstructure:"allowed-all-methods"   yaml:"allowed-all-methods"   json:"allowed_all_methods"`                            // 是否允许所有方法
	ExposedHeaders      []string `mapstructure:"exposed-headers"       yaml:"exposed-headers"       json:"exposed_headers"`                                // 暴露的头部
	AllowCredentials    bool     `mapstructure:"allow-credentials"     yaml:"allow-credentials"     json:"allow_credentials"`                              // 允许凭证
	OptionsResponseCode int      `mapstructure:"options-response-code"  yaml:"options-response-code" json:"options_response_code"`                         // Options响应Code
	ModuleName          string   `mapstructure:"modulename"            yaml:"modulename"            json:"module_name"`                                    // 模块名称
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
	return internal.ValidateStruct(c)
}

// DefaultCors 返回默认Cors配置
func DefaultCors() Cors {
	return Cors{
		ModuleName:          "cors",
		AllowedOrigins:      []string{"*"},
		AllowedMethods:      []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:      []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:              "86400",
		AllowedAllOrigins:   false,
		AllowedAllMethods:   false,
		ExposedHeaders:      []string{},
		AllowCredentials:    true,
		OptionsResponseCode: 200,
	}
}

// Default 返回默认Cors配置的指针，支持链式调用
func Default() *Cors {
	config := DefaultCors()
	return &config
}

// WithModuleName 设置模块名称
func (c *Cors) WithModuleName(moduleName string) *Cors {
	c.ModuleName = moduleName
	return c
}

// WithAllowedOrigins 设置允许的来源
func (c *Cors) WithAllowedOrigins(allowedOrigins []string) *Cors {
	c.AllowedOrigins = allowedOrigins
	return c
}

// WithAllowedMethods 设置允许的方法
func (c *Cors) WithAllowedMethods(allowedMethods []string) *Cors {
	c.AllowedMethods = allowedMethods
	return c
}

// WithAllowedHeaders 设置允许的头部
func (c *Cors) WithAllowedHeaders(allowedHeaders []string) *Cors {
	c.AllowedHeaders = allowedHeaders
	return c
}

// WithMaxAge 设置最大缓存时间
func (c *Cors) WithMaxAge(maxAge string) *Cors {
	c.MaxAge = maxAge
	return c
}

// WithAllowedAllOrigins 设置是否允许所有来源
func (c *Cors) WithAllowedAllOrigins(allowedAllOrigins bool) *Cors {
	c.AllowedAllOrigins = allowedAllOrigins
	return c
}

// WithAllowedAllMethods 设置是否允许所有方法
func (c *Cors) WithAllowedAllMethods(allowedAllMethods bool) *Cors {
	c.AllowedAllMethods = allowedAllMethods
	return c
}

// WithExposedHeaders 设置暴露的头部
func (c *Cors) WithExposedHeaders(exposedHeaders []string) *Cors {
	c.ExposedHeaders = exposedHeaders
	return c
}

// WithAllowCredentials 设置是否允许凭证
func (c *Cors) WithAllowCredentials(allowCredentials bool) *Cors {
	c.AllowCredentials = allowCredentials
	return c
}

// WithOptionsResponseCode 设置Options响应码
func (c *Cors) WithOptionsResponseCode(optionsResponseCode int) *Cors {
	c.OptionsResponseCode = optionsResponseCode
	return c
}
