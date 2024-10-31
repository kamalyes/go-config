/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 10:06:09
 * @FilePath: \go-config\consul\consul.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package consul

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Consul 注册中心配置
type Consul struct {
	/** 基础 **/
	ModuleName string `mapstructure:"modulename" json:"moduleName" yaml:"modulename"`

	/** 注册中心地址 */
	Addr string `mapstructure:"addr"                 json:"addr"                 yaml:"addr"`

	/** 间隔 单位秒 */
	RegisterInterval int `mapstructure:"register-interval"    json:"registerInterval"     yaml:"register-interval"`
}

// NewConsul 创建一个新的 Consul 实例
func NewConsul(moduleName, addr string, registerInterval int) *Consul {
	var consulInstance *Consul

	internal.LockFunc(func() {
		consulInstance = &Consul{
			ModuleName:       moduleName,
			Addr:             addr,
			RegisterInterval: registerInterval,
		}
	})
	return consulInstance
}

// ToMap 将配置转换为映射
func (c *Consul) ToMap() map[string]interface{} {
	return internal.ToMap(c)
}

// FromMap 从映射中填充配置
func (c *Consul) FromMap(data map[string]interface{}) {
	internal.FromMap(c, data)
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
	if c.Addr == "" {
		return errors.New("addr cannot be empty")
	}
	if c.RegisterInterval <= 0 {
		return errors.New("registerInterval must be greater than 0")
	}
	return nil
}
