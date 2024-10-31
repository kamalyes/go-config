/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 11:38:00
 * @FilePath: \go-config\pkg\register\consul.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package register

import (
	"github.com/kamalyes/go-config/internal"
)

// Consul 结构体用于配置 Consul 注册中心相关参数
type Consul struct {
	ModuleName       string `mapstructure:"modulename"              yaml:"modulename"        json:"module_name"       validate:"required"`      // 模块名称
	Addr             string `mapstructure:"addr"                     yaml:"addr"              json:"addr"              validate:"required,url"` // 注册中心地址
	RegisterInterval int    `mapstructure:"register-interval"        yaml:"register-interval" json:"register_interval" validate:"min=1"`        // 注册间隔，单位秒，最小值为 1 秒
}

// NewConsul 创建一个新的 Consul 实例
func NewConsul(opt *Consul) *Consul {
	var consulInstance *Consul

	internal.LockFunc(func() {
		consulInstance = opt
	})
	return consulInstance
}

// Clone 返回 Consul 配置的副本
func (c *Consul) Clone() internal.Configurable {
	return &Consul{
		ModuleName:       c.ModuleName,
		Addr:             c.Addr,
		RegisterInterval: c.RegisterInterval,
	}
}

// Get 返回 Consul 配置的所有字段
func (c *Consul) Get() interface{} {
	return c
}

// Set 更新 Consul 配置的字段
func (c *Consul) Set(data interface{}) {
	if configData, ok := data.(*Consul); ok {
		c.ModuleName = configData.ModuleName
		c.Addr = configData.Addr
		c.RegisterInterval = configData.RegisterInterval
	}
}

// Validate 检查 Consul 配置的有效性
func (c *Consul) Validate() error {
	return internal.ValidateStruct(c)
}
