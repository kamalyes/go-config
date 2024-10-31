/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:59:28
 * @FilePath: \go-config\zero\client.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package zero

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

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

// NewRpcClient 创建一个新的 RpcClient 实例
func NewRpcClient(moduleName, target, app, token string, nonBlock bool, timeout int64) *RpcClient {
	var rpcClientInstance *RpcClient

	internal.LockFunc(func() {
		rpcClientInstance = &RpcClient{
			ModuleName: moduleName,
			Target:     target,
			App:        app,
			Token:      token,
			NonBlock:   nonBlock,
			Timeout:    timeout,
		}
	})
	return rpcClientInstance
}

// ToMap 将配置转换为映射
func (r *RpcClient) ToMap() map[string]interface{} {
	return internal.ToMap(r)
}

// FromMap 从映射中填充配置
func (r *RpcClient) FromMap(data map[string]interface{}) {
	internal.FromMap(r, data)
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
