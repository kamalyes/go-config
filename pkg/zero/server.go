/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:55:27
 * @FilePath: \go-config\pkg\zero\server.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// RpcServer 结构体表示 RPC 服务器的配置
type RpcServer struct {
	ModuleName    string      `mapstructure:"MODULE_NAME"              yaml:"modulename"`     // 模块名称
	ListenOn      string      `mapstructure:"LISTEN_ON"                yaml:"listen-on"`      // 监听地址
	Auth          bool        `mapstructure:"AUTH"                     yaml:"auth"`           // 是否启用认证
	StrictControl bool        `mapstructure:"STRICT_CONTROL"           yaml:"strict-control"` // 是否启用严格控制
	Timeout       int64       `mapstructure:"TIMEOUT"                  yaml:"timeout"`        // 超时时间，单位毫秒
	CpuThreshold  int64       `mapstructure:"CPU_THRESHOLD"            yaml:"cpu-threshold"`  // CPU 使用率阈值
	Etcd          *EtcdConfig `mapstructure:"ETCD"                     yaml:"etcd"`           // Etcd 配置
}

// NewRpcServer 创建一个新的 RpcServer 实例
func NewRpcServer(opt *RpcServer) *RpcServer {
	var rpcServerInstance *RpcServer

	internal.LockFunc(func() {
		rpcServerInstance = opt
	})
	return rpcServerInstance
}

// Clone 返回 RpcServer 配置的副本
func (r *RpcServer) Clone() internal.Configurable {
	return &RpcServer{
		ModuleName:    r.ModuleName,
		ListenOn:      r.ListenOn,
		Auth:          r.Auth,
		StrictControl: r.StrictControl,
		Timeout:       r.Timeout,
		CpuThreshold:  r.CpuThreshold,
		Etcd:          r.Etcd,
	}
}

// Get 返回 RpcServer 配置的所有字段
func (r *RpcServer) Get() interface{} {
	return r
}

// Set 更新 RpcServer 配置的字段
func (r *RpcServer) Set(data interface{}) {
	if configData, ok := data.(*RpcServer); ok {
		r.ModuleName = configData.ModuleName
		r.ListenOn = configData.ListenOn
		r.Auth = configData.Auth
		r.StrictControl = configData.StrictControl
		r.Timeout = configData.Timeout
		r.CpuThreshold = configData.CpuThreshold
		r.Etcd = configData.Etcd
	}
}

// Validate 验证 RpcServer 配置的有效性
func (r *RpcServer) Validate() error {
	if r.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if r.ListenOn == "" {
		return errors.New("listen address cannot be empty")
	}
	if r.Timeout <= 0 {
		return errors.New("timeout must be greater than 0")
	}
	if r.CpuThreshold < 0 {
		return errors.New("CPU threshold cannot be negative")
	}
	if len(r.Etcd.Hosts) == 0 {
		return errors.New("etcd hosts cannot be empty")
	}
	return nil
}
