/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 12:08:56
 * @FilePath: \go-config\zero\common.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// EtcdConfig 是 Etcd 的配置
type EtcdConfig struct {
	/** Etcd 主机列表 */
	Hosts []string `mapstructure:"hosts" json:"hosts" yaml:"hosts"`

	/** 注册的键 */
	Key string `mapstructure:"key" json:"key" yaml:"key"`
}

// NewEtcdConfig 创建一个新的 EtcdConfig 实例
func NewEtcdConfig(hosts []string, key string) *EtcdConfig {
	var etcdInstance *EtcdConfig

	internal.LockFunc(func() {
		etcdInstance = &EtcdConfig{
			Hosts: hosts,
			Key:   key,
		}
	})
	return etcdInstance
}

// ToMap 将配置转换为映射
func (e *EtcdConfig) ToMap() map[string]interface{} {
	return internal.ToMap(e)
}

// FromMap 从映射中填充配置
func (e *EtcdConfig) FromMap(data map[string]interface{}) {
	internal.FromMap(e, data)
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
	if len(e.Hosts) == 0 {
		return errors.New("etcd hosts cannot be empty")
	}
	if e.Key == "" {
		return errors.New("etcd key cannot be empty")
	}
	return nil
}
