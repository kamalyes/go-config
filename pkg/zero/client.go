/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:55:18
 * @FilePath: \go-config\pkg\zero\client.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package zero

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// RpcClient 结构体表示 RPC 客户端的配置
type RpcClient struct {
	ModuleName string `mapstructure:"MODULE_NAME"              yaml:"modulename"` // 模块名称
	Target     string `mapstructure:"TARGET"                   yaml:"target"`     // 目标地址
	App        string `mapstructure:"APP"                      yaml:"app"`        // 应用名称
	Token      string `mapstructure:"TOKEN"                    yaml:"token"`      // 认证令牌
	NonBlock   bool   `mapstructure:"NON_BLOCK"                yaml:"non-block"`  // 是否非阻塞
	Timeout    int64  `mapstructure:"TIMEOUT"                  yaml:"timeout"`    // 超时时间，单位毫秒
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
	return &RpcClient{
		ModuleName: r.ModuleName,
		Target:     r.Target,
		App:        r.App,
		Token:      r.Token,
		NonBlock:   r.NonBlock,
		Timeout:    r.Timeout,
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
	}
}

// Validate 验证 RpcClient 配置的有效性
func (r *RpcClient) Validate() error {
	if r.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if r.Target == "" {
		return errors.New("target cannot be empty")
	}
	if r.App == "" {
		return errors.New("app cannot be empty")
	}
	if r.Token == "" {
		return errors.New("token cannot be empty")
	}
	if r.Timeout <= 0 {
		return errors.New("timeout must be greater than 0")
	}
	return nil
}
