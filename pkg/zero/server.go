/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:56:26
 * @FilePath: \go-config\pkg\zero\server.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

import (
	"github.com/kamalyes/go-config/internal"
)

// RpcServer 结构体表示 RPC 服务器的配置
type RpcServer struct {
	ListenOn      string      `mapstructure:"listen-on"                yaml:"listen-on"        json:"listen_on"        validate:"required"` // 监听地址
	Timeout       int64       `mapstructure:"timeout"                  yaml:"timeout"          json:"timeout"          validate:"gte=0"`    // 超时时间，单位毫秒
	CpuThreshold  int64       `mapstructure:"cpu-threshold"            yaml:"cpu-threshold"    json:"cpu_threshold"    validate:"gte=0"`    // CPU 使用率阈值
	Etcd          *Etcd       `mapstructure:"etcd"                     yaml:"etcd"             json:"etcd"             validate:"required"` // Etcd 配置
	Auth          bool        `mapstructure:"auth"                     yaml:"auth"             json:"auth"`                                 // 是否启用认证
	StrictControl bool        `mapstructure:"strict-control"           yaml:"strict-control"   json:"strict_control"`                       // 是否启用严格控制
	LogConf       *LogConf    `mapstructure:"log-conf"                 yaml:"log-conf"         json:"log_conf"`                             // Log 配置
	Name          string      `mapstructure:"name"                     yaml:"name"             json:"name"`                                 // 服务器名称
	Mode          string      `mapstructure:"mode"                     yaml:"mode"             json:"mode"`                                 // 运行模式
	MetricsUrl    string      `mapstructure:"metrics-url"              yaml:"metrics-url"      json:"metrics_url"`                          // 指标 URL
	Prometheus    *Prometheus `mapstructure:"prometheus"               yaml:"prometheus"       json:"prometheus"`                           // Prometheus 配置
	Telemetry     *Telemetry  `mapstructure:"telemetry"                yaml:"telemetry"        json:"telemetry"`                            // 追踪配置
	ModuleName    string      `mapstructure:"modulename"               yaml:"modulename"       json:"module_name"`                          // 模块名称
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
	var (
		logConfClone    *LogConf
		etcdConfigClone *Etcd
		prometheusClone *Prometheus
		telemetry       *Telemetry
	)
	if r.LogConf != nil {
		logConfClone = r.LogConf.Clone().(*LogConf) // 确保克隆 LogConf
	}

	if r.Etcd != nil {
		etcdConfigClone = r.Etcd.Clone().(*Etcd) // 确保克隆 EtcdConfig
	}

	if r.Prometheus != nil {
		prometheusClone = r.Prometheus.Clone().(*Prometheus) // 确保克隆 Prometheus
	}

	if r.Telemetry != nil {
		telemetry = r.Telemetry.Clone().(*Telemetry) // 确保克隆 Telemetry
	}

	return &RpcServer{
		ModuleName:    r.ModuleName,
		ListenOn:      r.ListenOn,
		Auth:          r.Auth,
		StrictControl: r.StrictControl,
		Timeout:       r.Timeout,
		CpuThreshold:  r.CpuThreshold,
		Etcd:          etcdConfigClone,
		LogConf:       logConfClone,
		Name:          r.Name,
		Mode:          r.Mode,
		MetricsUrl:    r.MetricsUrl,
		Prometheus:    prometheusClone,
		Telemetry:     telemetry,
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
		r.LogConf = configData.LogConf
		r.Name = configData.Name
		r.Mode = configData.Mode
		r.MetricsUrl = configData.MetricsUrl
		r.Prometheus = configData.Prometheus
		r.Telemetry = configData.Telemetry
	}
}

// Validate 验证 RpcServer 配置的有效性
func (r *RpcServer) Validate() error {
	return internal.ValidateStruct(r)
}
