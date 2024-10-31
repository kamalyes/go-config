/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-31 12:50:58
 * @FilePath: \go-config\zero\server.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

// RpcServer 是 RPC 服务器的配置
type RpcServer struct {
	/** 模块名称 */
	ModuleName string `mapstructure:"modulename" json:"moduleName" yaml:"modulename"`

	/** 监听地址 */
	ListenOn string `mapstructure:"listen-on" json:"listenOn" yaml:"listen-on"`

	/** 是否启用认证 */
	Auth bool `mapstructure:"auth" json:"auth" yaml:"auth"`

	/** 是否启用严格控制 */
	StrictControl bool `mapstructure:"strict-control" json:"strictControl" yaml:"strict-control"`

	/** 超时时间，单位毫秒 */
	Timeout int64 `mapstructure:"timeout" json:"timeout" yaml:"timeout"`

	/** CPU 使用率阈值 */
	CpuThreshold int64 `mapstructure:"cpu-threshold" json:"cpuThreshold" yaml:"cpu-threshold"`

	/** Etcd 配置 */
	Etcd EtcdConfig `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
}

// EtcdConfig 是 Etcd 的配置
type EtcdConfig struct {
	/** Etcd 主机列表 */
	Hosts []string `mapstructure:"hosts" json:"hosts" yaml:"hosts"`

	/** 注册的键 */
	Key string `mapstructure:"key" json:"key" yaml:"key"`
}

// 实现 Configurable 接口
func (r RpcServer) GetModuleName() string {
	return "rpc_server"
}
