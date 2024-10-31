/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 09:26:11
 * @FilePath: \go-config\server\server.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package server

type Server struct {

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** 地址 */
	Host string `mapstructure:"host"                     json:"host"                     yaml:"host"`

	/** 端口 */
	Port string `mapstructure:"port"                     json:"port"                     yaml:"port"`

	/** 服务名称 */
	ServerName string `mapstructure:"server-name"                 json:"serverName"               yaml:"server-name"`

	/** 请求根路径 */
	ContextPath string `mapstructure:"context-path"                json:"contextPath"              yaml:"context-path"`

	/** 数据库类型 */
	DataDriver string `mapstructure:"data-driver"                 json:"dataDriver"               yaml:"data-driver"`

	/** 是否开启请求方式检测 */
	HandleMethodNotAllowed bool `mapstructure:"handle-method-not-allowed"   json:"handleMethodNotAllowed"   yaml:"handle-method-not-allowed"`

	/** 语言    */
	Language string `mapstructure:"language"                    json:"language"                 yaml:"language"`
}

// 实现 Configurable 接口
func (s Server) GetModuleName() string {
	return "server"
}
