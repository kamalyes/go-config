/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 17:56:56
 * @FilePath: \go-config\pkg\zero\client.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package zero

import (
	"github.com/kamalyes/go-config/internal"
)

// RpcClient 结构体表示 RPC 客户端的配置
type RpcClient struct {
	ModuleName string   `mapstructure:"modulename"                yaml:"modulename"           json:"module_name"           validate:"required"`       // 模块名称
	Host       string   `mapstructure:"host"                      yaml:"host"                 json:"host"                  validate:"required"`       // 地址
	Port       int64    `mapstructure:"port"                      yaml:"port"                 json:"port"                  validate:"required"`       // 端口
	App        string   `mapstructure:"app"                       yaml:"app"                  json:"app"                   validate:"required"`       // 应用名称
	Token      string   `mapstructure:"token"                     yaml:"token"                json:"token"                 validate:"required"`       // 认证令牌
	Timeout    int64    `mapstructure:"timeout"                   yaml:"timeout"              json:"timeout"               validate:"required,min=1"` // 超时时间，单位毫秒，必须大于0
	NonBlock   bool     `mapstructure:"non-block"                 yaml:"non-block"            json:"non_block"`                                       // 是否非阻塞
	LogConf    *LogConf `mapstructure:"log-conf"                  yaml:"log-conf"             json:"log_conf"`                                        // Log 配置
}

// NewRpcClient 创建一个新的 RpcClient 实例
func NewRpcClient(opt *RpcClient) *RpcClient {
	var rpcClientInstance *RpcClient

	internal.LockFunc(func() {
		rpcClientInstance = opt
	})
	return rpcClientInstance
}

// Clone 返回 RpcClient 配置的副本
func (r *RpcClient) Clone() internal.Configurable {
	var logConfClone *LogConf
	if r.LogConf != nil {
		logConfClone = r.LogConf.Clone().(*LogConf) // 确保克隆 LogConf
	}
	return &RpcClient{
		ModuleName: r.ModuleName,
		Host:       r.Host,
		Port:       r.Port,
		App:        r.App,
		Token:      r.Token,
		NonBlock:   r.NonBlock,
		Timeout:    r.Timeout,
		LogConf:    logConfClone,
	}
}

// Get 返回 RpcClient 配置的所有字段
func (r *RpcClient) Get() interface{} {
	return r
}

// Set 更新 RpcClient 配置的字段
func (r *RpcClient) Set(data interface{}) {
	if configData, ok := data.(*RpcClient); ok {
		r.ModuleName = configData.ModuleName
		r.Host = configData.Host
		r.Port = configData.Port
		r.App = configData.App
		r.Token = configData.Token
		r.NonBlock = configData.NonBlock
		r.Timeout = configData.Timeout
		r.LogConf = configData.LogConf
	}
}

// Validate 验证 RpcClient 配置的有效性
func (r *RpcClient) Validate() error {
	return internal.ValidateStruct(r)
}
