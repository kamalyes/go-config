/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 16:34:49
 * @FilePath: \go-config\zero\client.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package zero

// RpcClient 是 RPC 客户端的配置
type RpcClient struct {
	/** 模块名称 */
	ModuleName string `mapstructure:"modulename" json:"moduleName" yaml:"modulename"`

	/** 目标地址 */
	Target string `mapstructure:"target" json:"target" yaml:"target"`

	/** 应用名称 */
	App string `mapstructure:"app" json:"app" yaml:"app"`

	/** 认证令牌 */
	Token string `mapstructure:"token" json:"token" yaml:"token"`

	/** 是否非阻塞 */
	NonBlock bool `mapstructure:"non-block" json:"nonBlock" yaml:"non-block"`

	/** 超时时间，单位毫秒 */
	Timeout int64 `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}

// 实现 Configurable 接口
func (c RpcClient) GetModuleName() string {
	return "rpc_client"
}
