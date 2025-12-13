/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-13 11:53:15
 * @FilePath: \go-config\pkg\gateway\grpc.go
 * @Description: GRPC配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package gateway

import (
	"fmt"

	"github.com/kamalyes/go-config/internal"
)

// GRPC GRPC服务配置（包含服务端和多个客户端）
type GRPC struct {
	Server  *GRPCServer            `mapstructure:"server" yaml:"server" json:"server"`    // GRPC服务端配置
	Clients map[string]*GRPCClient `mapstructure:"clients" yaml:"clients" json:"clients"` // GRPC客户端配置（key为服务名）
}

// GRPCServer GRPC服务端配置
type GRPCServer struct {
	Enable            bool   `mapstructure:"enable" yaml:"enable" json:"enable"`                                    // 是否启用GRPC服务
	Host              string `mapstructure:"host" yaml:"host" json:"host"`                                          // 主机地址
	Port              int    `mapstructure:"port" yaml:"port" json:"port"`                                          // 端口
	Network           string `mapstructure:"network" yaml:"network" json:"network"`                                 // 网络类型 (tcp, tcp4, tcp6, unix)
	MaxRecvMsgSize    int    `mapstructure:"max-recv-msg-size" yaml:"max-recv-msg-size" json:"maxRecvMsgSize"`      // 最大接收消息大小(字节)
	MaxSendMsgSize    int    `mapstructure:"max-send-msg-size" yaml:"max-send-msg-size" json:"maxSendMsgSize"`      // 最大发送消息大小(字节)
	KeepaliveTime     int    `mapstructure:"keepalive-time" yaml:"keepalive-time" json:"keepaliveTime"`             // Keepalive时间(秒)
	KeepaliveTimeout  int    `mapstructure:"keepalive-timeout" yaml:"keepalive-timeout" json:"keepaliveTimeout"`    // Keepalive超时(秒)
	ConnectionTimeout int    `mapstructure:"connection-timeout" yaml:"connection-timeout" json:"connectionTimeout"` // 连接超时(秒)
	EnableReflection  bool   `mapstructure:"enable-reflection" yaml:"enable-reflection" json:"enableReflection"`    // 是否启用反射
	Endpoint          string `mapstructure:"-" yaml:"-" json:""`                                                    // 完整的服务端点地址（自动计算）
}

// GRPCClient GRPC客户端配置
type GRPCClient struct {
	ServiceName       string   `mapstructure:"service-name" yaml:"service-name" json:"serviceName"`                     // 服务名称
	Endpoints         []string `mapstructure:"endpoints" yaml:"endpoints" json:"endpoints"`                             // 服务端点列表
	Network           string   `mapstructure:"network" yaml:"network" json:"network"`                                   // 网络类型 (tcp, tcp4, tcp6, unix)
	MaxRecvMsgSize    int      `mapstructure:"max-recv-msg-size" yaml:"max-recv-msg-size" json:"maxRecvMsgSize"`        // 最大接收消息大小(字节)
	MaxSendMsgSize    int      `mapstructure:"max-send-msg-size" yaml:"max-send-msg-size" json:"maxSendMsgSize"`        // 最大发送消息大小(字节)
	KeepaliveTime     int      `mapstructure:"keepalive-time" yaml:"keepalive-time" json:"keepaliveTime"`               // Keepalive时间(秒)
	KeepaliveTimeout  int      `mapstructure:"keepalive-timeout" yaml:"keepalive-timeout" json:"keepaliveTimeout"`      // Keepalive超时(秒)
	Timeout           int      `mapstructure:"timeout" yaml:"timeout" json:"timeout"`                                   // 请求超时(秒)
	RetryTimes        int      `mapstructure:"retry-times" yaml:"retry-times" json:"retryTimes"`                        // 重试次数
	EnableLoadBalance bool     `mapstructure:"enable-load-balance" yaml:"enable-load-balance" json:"enableLoadBalance"` // 是否启用负载均衡
	LoadBalancePolicy string   `mapstructure:"load-balance-policy" yaml:"load-balance-policy" json:"loadBalancePolicy"` // 负载均衡策略
	EnableTLS         bool     `mapstructure:"enable-tls" yaml:"enable-tls" json:"enableTls"`                           // 是否启用TLS
	TLSCertFile       string   `mapstructure:"tls-cert-file" yaml:"tls-cert-file" json:"tlsCertFile"`                   // TLS证书文件
	TLSKeyFile        string   `mapstructure:"tls-key-file" yaml:"tls-key-file" json:"tlsKeyFile"`                      // TLS密钥文件
	TLSCAFile         string   `mapstructure:"tls-ca-file" yaml:"tls-ca-file" json:"tlsCaFile"`                         // TLS CA文件
}

// DefaultGRPC 创建默认GRPC配置
func DefaultGRPC() *GRPC {
	return &GRPC{
		Server:  DefaultGRPCServer(),
		Clients: make(map[string]*GRPCClient),
	}
}

// DefaultGRPCServer 创建默认GRPC服务端配置
func DefaultGRPCServer() *GRPCServer {
	g := &GRPCServer{
		Enable:            false, // 默认不启用，需要显式配置
		Host:              "0.0.0.0",
		Port:              9090,
		Network:           "tcp4",          // 默认使用 tcp4 强制 IPv4
		MaxRecvMsgSize:    4 * 1024 * 1024, // 4MB
		MaxSendMsgSize:    4 * 1024 * 1024, // 4MB
		KeepaliveTime:     30,
		KeepaliveTimeout:  10,
		ConnectionTimeout: 5,
		EnableReflection:  true,
	}
	internal.CallAfterLoad(g) // 自动调用 AfterLoad 钩子
	return g
}

// BeforeLoad GRPC Server 配置加载前的钩子
func (g *GRPCServer) BeforeLoad() error {
	return nil
}

// AfterLoad GRPC Server 配置加载后的钩子 - 计算 Endpoint
func (g *GRPCServer) AfterLoad() error {
	g.Endpoint = fmt.Sprintf("%s:%d", g.Host, g.Port)
	return nil
}

// GetEndpoint 获取GRPC服务端点地址
func (g *GRPCServer) GetEndpoint() string {
	if g.Endpoint == "" {
		internal.CallAfterLoad(g)
	}
	return g.Endpoint
}

// DefaultGRPCClient 创建默认GRPC客户端配置
func DefaultGRPCClient(serviceName string, endpoints []string) *GRPCClient {
	return &GRPCClient{
		ServiceName:       serviceName,
		Endpoints:         endpoints,
		Network:           "tcp4",          // 默认使用 tcp4 强制 IPv4
		MaxRecvMsgSize:    4 * 1024 * 1024, // 4MB
		MaxSendMsgSize:    4 * 1024 * 1024, // 4MB
		KeepaliveTime:     30,
		KeepaliveTimeout:  10,
		Timeout:           30,
		RetryTimes:        3,
		EnableLoadBalance: true,
		LoadBalancePolicy: "round_robin",
		EnableTLS:         false,
	}
}

// Clone 返回配置的副本
func (g *GRPC) Clone() internal.Configurable {
	clients := make(map[string]*GRPCClient)
	for k, v := range g.Clients {
		clients[k] = v.Clone()
	}
	return &GRPC{
		Server:  g.Server.Clone(),
		Clients: clients,
	}
}

// Clone 返回GRPC服务端配置的副本
func (g *GRPCServer) Clone() *GRPCServer {
	cloned := &GRPCServer{
		Enable:            g.Enable,
		Host:              g.Host,
		Port:              g.Port,
		Network:           g.Network,
		MaxRecvMsgSize:    g.MaxRecvMsgSize,
		MaxSendMsgSize:    g.MaxSendMsgSize,
		KeepaliveTime:     g.KeepaliveTime,
		KeepaliveTimeout:  g.KeepaliveTimeout,
		ConnectionTimeout: g.ConnectionTimeout,
		EnableReflection:  g.EnableReflection,
	}
	internal.CallAfterLoad(cloned) // 重新计算 Endpoint
	return cloned
}

// Clone 返回GRPC客户端配置的副本
func (g *GRPCClient) Clone() *GRPCClient {
	endpoints := make([]string, len(g.Endpoints))
	copy(endpoints, g.Endpoints)
	return &GRPCClient{
		ServiceName:       g.ServiceName,
		Endpoints:         endpoints,
		Network:           g.Network,
		MaxRecvMsgSize:    g.MaxRecvMsgSize,
		MaxSendMsgSize:    g.MaxSendMsgSize,
		KeepaliveTime:     g.KeepaliveTime,
		KeepaliveTimeout:  g.KeepaliveTimeout,
		Timeout:           g.Timeout,
		RetryTimes:        g.RetryTimes,
		EnableLoadBalance: g.EnableLoadBalance,
		LoadBalancePolicy: g.LoadBalancePolicy,
		EnableTLS:         g.EnableTLS,
		TLSCertFile:       g.TLSCertFile,
		TLSKeyFile:        g.TLSKeyFile,
		TLSCAFile:         g.TLSCAFile,
	}
}

// Get 返回配置接口
func (g *GRPC) Get() interface{} {
	return g
}

// Set 设置配置数据
func (g *GRPC) Set(data interface{}) {
	if cfg, ok := data.(*GRPC); ok {
		*g = *cfg
	}
}

// Validate 验证配置
func (g *GRPC) Validate() error {
	if err := internal.ValidateStruct(g); err != nil {
		return err
	}
	if g.Server != nil {
		if err := g.Server.Validate(); err != nil {
			return err
		}
	}
	for _, client := range g.Clients {
		if err := client.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate 验证GRPC服务端配置
func (g *GRPCServer) Validate() error {
	return internal.ValidateStruct(g)
}

// Validate 验证GRPC客户端配置
func (g *GRPCClient) Validate() error {
	return internal.ValidateStruct(g)
}

// WithHost 设置主机地址
func (g *GRPCServer) WithHost(host string) *GRPCServer {
	g.Host = host
	internal.CallAfterLoad(g) // 重新计算 Endpoint
	return g
}

// WithPort 设置端口
func (g *GRPCServer) WithPort(port int) *GRPCServer {
	g.Port = port
	internal.CallAfterLoad(g) // 重新计算 Endpoint
	return g
}

// WithMaxMsgSize 设置最大消息大小
func (g *GRPCServer) WithMaxMsgSize(maxRecv, maxSend int) *GRPCServer {
	g.MaxRecvMsgSize = maxRecv
	g.MaxSendMsgSize = maxSend
	return g
}

// WithKeepalive 设置Keepalive配置
func (g *GRPCServer) WithKeepalive(time, timeout int) *GRPCServer {
	g.KeepaliveTime = time
	g.KeepaliveTimeout = timeout
	return g
}

// EnableReflectionService 启用反射服务
func (g *GRPCServer) EnableReflectionService() *GRPCServer {
	g.EnableReflection = true
	return g
}

// DisableReflectionService 禁用反射服务
func (g *GRPCServer) DisableReflectionService() *GRPCServer {
	g.EnableReflection = false
	return g
}

// AddClient 添加GRPC客户端配置
func (g *GRPC) AddClient(name string, client *GRPCClient) *GRPC {
	if g.Clients == nil {
		g.Clients = make(map[string]*GRPCClient)
	}
	g.Clients[name] = client
	return g
}

// GetClient 获取GRPC客户端配置
func (g *GRPC) GetClient(name string) *GRPCClient {
	if g.Clients == nil {
		return nil
	}
	return g.Clients[name]
}

// WithTimeout 设置超时
func (g *GRPCClient) WithTimeout(timeout int) *GRPCClient {
	g.Timeout = timeout
	return g
}

// WithRetry 设置重试次数
func (g *GRPCClient) WithRetry(retryTimes int) *GRPCClient {
	g.RetryTimes = retryTimes
	return g
}

// EnableLoadBalancing 启用负载均衡
func (g *GRPCClient) EnableLoadBalancing(policy string) *GRPCClient {
	g.EnableLoadBalance = true
	g.LoadBalancePolicy = policy
	return g
}

// WithTLS 设置TLS配置
func (g *GRPCClient) WithTLS(certFile, keyFile, caFile string) *GRPCClient {
	g.EnableTLS = true
	g.TLSCertFile = certFile
	g.TLSKeyFile = keyFile
	g.TLSCAFile = caFile
	return g
}

// WithEndpoints 设置服务端点
func (g *GRPCClient) WithEndpoints(endpoints []string) *GRPCClient {
	g.Endpoints = endpoints
	return g
}

// AddEndpoint 添加服务端点
func (g *GRPCClient) AddEndpoint(endpoint string) *GRPCClient {
	g.Endpoints = append(g.Endpoints, endpoint)
	return g
}

// WithNetwork 设置网络类型
func (g *GRPCClient) WithNetwork(network string) *GRPCClient {
	g.Network = network
	return g
}
