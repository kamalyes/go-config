/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-07 15:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-08 00:03:22
 * @FilePath: \go-config\pkg\zero\restful.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

import "github.com/kamalyes/go-config/internal"

// Restful 是 REST 服务的配置
type Restful struct {
	RpcServer    *RpcServer `mapstructure:"rpc_server" yaml:"rpc-server" json:"rpc_server"`          // 嵌入 RpcServer 结构体指针
	Host         string     `mapstructure:"host"       yaml:"host"       json:"host"`                // 主机
	Port         int        `mapstructure:"port"       yaml:"port"       json:"port"`                // 端口
	CertFile     string     `mapstructure:"cert_file"  yaml:"cert-file"  json:"cert_file"`           // 证书文件
	KeyFile      string     `mapstructure:"key_file"   yaml:"key-file"   json:"key_file"`            // 密钥文件
	Verbose      bool       `mapstructure:"verbose"    yaml:"verbose"    json:"verbose"`             // 是否详细输出
	MaxConns     int        `mapstructure:"max_conns"  yaml:"max-conns"  json:"max_conns"`           // 最大连接数
	MaxBytes     int64      `mapstructure:"max_bytes"  yaml:"max-bytes"  json:"max_bytes"`           // 最大字节数
	Timeout      int64      `mapstructure:"timeout"    yaml:"timeout"    json:"timeout"`             // 超时时间
	CpuThreshold int64      `mapstructure:"cpu_threshold" yaml:"cpu-threshold" json:"cpu_threshold"` // CPU 阈值
	Signature    *Signature `mapstructure:"signature"  yaml:"signature"  json:"signature"`           // 签名配置
	ModuleName   string     `mapstructure:"modulename" yaml:"modulename" json:"module_name"`         // 模块名称
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

// DefaultRestful 返回默认的 Restful 值
func DefaultRestful() Restful {
	return Restful{
		RpcServer:    nil,
		Host:         "0.0.0.0",
		Port:         8080,
		CertFile:     "",
		KeyFile:      "",
		Verbose:      false,
		MaxConns:     10000,
		MaxBytes:     1048576, // 1MB
		Timeout:      3000,
		CpuThreshold: 900,
		Signature:    nil,
		ModuleName:   "restful",
	}
}

// DefaultRestfulConfig 返回默认的 Restful 指针，支持链式调用
func DefaultRestfulConfig() *Restful {
	config := DefaultRestful()
	return &config
}

// WithRpcServer 设置RpcServer配置
func (z *Restful) WithRpcServer(rpcServer *RpcServer) *Restful {
	z.RpcServer = rpcServer
	return z
}

// WithHost 设置主机
func (z *Restful) WithHost(host string) *Restful {
	z.Host = host
	return z
}

// WithPort 设置端口
func (z *Restful) WithPort(port int) *Restful {
	z.Port = port
	return z
}

// WithCertFile 设置证书文件
func (z *Restful) WithCertFile(certFile string) *Restful {
	z.CertFile = certFile
	return z
}

// WithKeyFile 设置密钥文件
func (z *Restful) WithKeyFile(keyFile string) *Restful {
	z.KeyFile = keyFile
	return z
}

// WithVerbose 设置是否详细输出
func (z *Restful) WithVerbose(verbose bool) *Restful {
	z.Verbose = verbose
	return z
}

// WithMaxConns 设置最大连接数
func (z *Restful) WithMaxConns(maxConns int) *Restful {
	z.MaxConns = maxConns
	return z
}

// WithMaxBytes 设置最大字节数
func (z *Restful) WithMaxBytes(maxBytes int64) *Restful {
	z.MaxBytes = maxBytes
	return z
}

// WithTimeout 设置超时时间
func (z *Restful) WithTimeout(timeout int64) *Restful {
	z.Timeout = timeout
	return z
}

// WithCpuThreshold 设置CPU阈值
func (z *Restful) WithCpuThreshold(cpuThreshold int64) *Restful {
	z.CpuThreshold = cpuThreshold
	return z
}

// WithSignature 设置签名配置
func (z *Restful) WithSignature(signature *Signature) *Restful {
	z.Signature = signature
	return z
}

// WithModuleName 设置模块名称
func (z *Restful) WithModuleName(moduleName string) *Restful {
	z.ModuleName = moduleName
	return z
}
