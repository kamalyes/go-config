/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-07 15:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 18:06:52
 * @FilePath: \go-config\pkg\zero\etcd.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

import (
	"github.com/kamalyes/go-config/internal"
)

// EtcdConfig 结构体表示 Etcd 的配置
type EtcdConfig struct {
	Hosts []string `mapstructure:"hosts" yaml:"hosts" json:"hosts" validate:"required,dive,required"` // Etcd 主机列表
	Key   string   `mapstructure:"key"   yaml:"key"   json:"key"   validate:"required"`               // 注册的键
}

// NewEtcdConfig 创建一个新的 EtcdConfig 实例
func NewEtcdConfig(opt *EtcdConfig) *EtcdConfig {
	var etcdInstance *EtcdConfig

	internal.LockFunc(func() {
		etcdInstance = opt
	})
	return etcdInstance
}

// Clone 返回 RpcServer 配置的副本
func (e *EtcdConfig) Clone() internal.Configurable {
	return &EtcdConfig{
		Hosts: e.Hosts,
		Key:   e.Key,
	}
}

// Get 返回 RpcServer 配置的所有字段
func (e *EtcdConfig) Get() interface{} {
	return e
}

// Set 更新 RpcServer 配置的字段
func (e *EtcdConfig) Set(data interface{}) {
	if configData, ok := data.(*EtcdConfig); ok {
		e.Hosts = configData.Hosts
		e.Key = configData.Key
	}
}

// Validate 验证 EtcdConfig 配置的有效性
func (e *EtcdConfig) Validate() error {
	return internal.ValidateStruct(e)
}
