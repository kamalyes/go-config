/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-08 09:00:36
 * @FilePath: \go-config\pkg\zero\client.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package zero

import (
	"time"

	"github.com/kamalyes/go-config/internal"
)

// RpcClient 结构体表示 RPC 客户端的配置
type RpcClient struct {
	Etcd          *Etcd         `mapstructure:"etcd"                    yaml:"etcd"                  json:"etcd"`      // Etcd 配置
	Endpoints     []string      `mapstructure:"endpoints"               yaml:"endpoints"             json:"endpoints"` // 端点列表
	Target        string        `mapstructure:"target"                  yaml:"target"                json:"target"`    // 目标服务器地址
	App           string        `mapstructure:"app"                     yaml:"app"                   json:"app"`       // 应用名称
	Token         string        `mapstructure:"token"                   yaml:"token"                 json:"token"`     // 认证 token
	NonBlock      bool          `mapstructure:"non-block"               yaml:"non-block"             json:"non_block"` // 是否非阻塞
	Timeout       int64         `mapstructure:"timeout"                 yaml:"timeout"               json:"timeout"`   // 超时时间，单位毫秒，必须大于0
	KeepaliveTime time.Duration `mapstructure:"keepalive-time"          yaml:"keepalive-time"        json:"keepalive_time"`
	ModuleName    string        `mapstructure:"modulename"              yaml:"modulename"            json:"module_name"` // 模块名称
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
	var etcdClone *Etcd
	if r.Etcd != nil {
		etcdClone = r.Etcd.Clone().(*Etcd) // 确保克隆 Etcd
	}
	return &RpcClient{
		ModuleName:    r.ModuleName,
		Target:        r.Target,
		App:           r.App,
		Token:         r.Token,
		NonBlock:      r.NonBlock,
		Timeout:       r.Timeout,
		Etcd:          etcdClone,
		Endpoints:     r.Endpoints,
		KeepaliveTime: r.KeepaliveTime,
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
		r.Target = configData.Target
		r.App = configData.App
		r.Token = configData.Token
		r.NonBlock = configData.NonBlock
		r.Timeout = configData.Timeout
		r.Etcd = configData.Etcd
		r.Endpoints = configData.Endpoints
		r.KeepaliveTime = configData.KeepaliveTime
	}
}

// Validate 验证 RpcClient 配置的有效性
func (r *RpcClient) Validate() error {
	return internal.ValidateStruct(r)
}
