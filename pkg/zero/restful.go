/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-07 15:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 18:18:50
 * @FilePath: \go-config\pkg\zero\restful.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

import "github.com/kamalyes/go-config/internal"

// Restful 是 REST 服务的配置
type Restful struct {
	RpcServer    *RpcServer `mapstructure:"rpc_server" yaml:"rpc_server" json:"rpc_server"`          // 嵌入 RpcServer 结构体指针
	Host         string     `mapstructure:"host"       yaml:"host"       json:"host"`                // 主机
	Port         int        `mapstructure:"port"       yaml:"port"       json:"port"`                // 端口
	CertFile     string     `mapstructure:"cert_file"  yaml:"cert_file"  json:"cert_file"`           // 证书文件
	KeyFile      string     `mapstructure:"key_file"   yaml:"key_file"   json:"key_file"`            // 密钥文件
	Verbose      bool       `mapstructure:"verbose"    yaml:"verbose"    json:"verbose"`             // 是否详细输出
	MaxConns     int        `mapstructure:"max_conns"  yaml:"max_conns"  json:"max_conns"`           // 最大连接数
	MaxBytes     int64      `mapstructure:"max_bytes"  yaml:"max_bytes"  json:"max_bytes"`           // 最大字节数
	Timeout      int64      `mapstructure:"timeout"    yaml:"timeout"    json:"timeout"`             // 超时时间
	CpuThreshold int64      `mapstructure:"cpu_threshold" yaml:"cpu_threshold" json:"cpu_threshold"` // CPU 阈值
	Signature    *Signature `mapstructure:"signature"  yaml:"signature"  json:"signature"`           // 签名配置
}

// NewRestful 创建一个新的 Restful 实例
func NewRestful(opt *Restful) *Restful {
	var zeroRestfulInstance *Restful

	internal.LockFunc(func() {
		zeroRestfulInstance = opt
	})
	return zeroRestfulInstance
}

// Clone 返回 Restful 配置的副本
func (z *Restful) Clone() internal.Configurable {
	var (
		rpcServer *RpcServer
	)

	if z.RpcServer != nil {
		rpcServer = z.RpcServer.Clone().(*RpcServer) // 确保克隆 RpcServer
	}

	return &Restful{
		RpcServer:    rpcServer,
		Host:         z.Host,
		Port:         z.Port,
		CertFile:     z.CertFile,
		KeyFile:      z.KeyFile,
		Verbose:      z.Verbose,
		MaxConns:     z.MaxConns,
		MaxBytes:     z.MaxBytes,
		Timeout:      z.Timeout,
		CpuThreshold: z.CpuThreshold,
		Signature:    z.Signature,
	}
}

// Get 返回 Restful 配置的所有字段
func (z *Restful) Get() interface{} {
	return z
}

// Set 更新 Restful 配置的字段
func (z *Restful) Set(data interface{}) {
	if configData, ok := data.(*Restful); ok {
		z.RpcServer = configData.RpcServer // 更新 RpcServer 配置
		z.Host = configData.Host
		z.Port = configData.Port
		z.CertFile = configData.CertFile
		z.KeyFile = configData.KeyFile
		z.Verbose = configData.Verbose
		z.MaxConns = configData.MaxConns
		z.MaxBytes = configData.MaxBytes
		z.Timeout = configData.Timeout
		z.CpuThreshold = configData.CpuThreshold
		z.Signature = configData.Signature
	}
}

// Validate 验证 Restful 配置的有效性
func (z *Restful) Validate() error {
	if err := internal.ValidateStruct(z); err != nil {
		return err
	}
	// 额外的验证逻辑可以在这里添加
	return nil
}
