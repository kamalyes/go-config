/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 22:29:12
 * @FilePath: \go-config\consul\consul.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package consul

// Consul 注册中心配置
type Consul struct {
	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"   json:"moduleName"    yaml:"modulename"`

	/** 注册中心地址 */
	Addr string `mapstructure:"addr"                 json:"addr"                 yaml:"addr"`

	/** 间隔 单位秒 */
	RegisterInterval int `mapstructure:"register-interval"    json:"registerInterval"     yaml:"register-interval"`
}

// 实现 Configurable 接口
func (c Consul) GetModuleName() string {
	return "consul"
}
