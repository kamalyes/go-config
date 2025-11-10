/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 10:50:24
 * @FilePath: \go-config\pkg\server\server.go
 * @Description: 服务器配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package server

import "github.com/kamalyes/go-config/internal"

// Server 服务器配置
type Server struct {
	ModuleName   string            `mapstructure:"module_name" yaml:"module-name" json:"module_name"`       // 模块名称
	Endpoint     string            `mapstructure:"endpoint"    yaml:"endpoint"   json:"endpoint" `          // 服务器端点地址
	Host         string            `mapstructure:"host" yaml:"host" json:"host"`                            // 服务器地址
	Port         int               `mapstructure:"port" yaml:"port" json:"port"`                            // 服务器端口
	ReadTimeout  int               `mapstructure:"read_timeout" yaml:"read-timeout" json:"read_timeout"`    // 读取超时(秒)
	WriteTimeout int               `mapstructure:"write_timeout" yaml:"write-timeout" json:"write_timeout"` // 写入超时(秒)
	IdleTimeout  int               `mapstructure:"idle_timeout" yaml:"idle-timeout" json:"idle_timeout"`    // 空闲超时(秒)
	GrpcPort     int               `mapstructure:"grpc_port" yaml:"grpc-port" json:"grpc_port"`             // GRPC端口
	EnableHttp   bool              `mapstructure:"enable_http" yaml:"enable-http" json:"enable_http"`       // 是否启用HTTP
	EnableGrpc   bool              `mapstructure:"enable_grpc" yaml:"enable-grpc" json:"enable_grpc"`       // 是否启用GRPC
	EnableTls    bool              `mapstructure:"enable_tls" yaml:"enable-tls" json:"enable_tls"`          // 是否启用TLS
	TLS          *TLS              `mapstructure:"tls" yaml:"tls" json:"tls"`                               // TLS配置
	Headers      map[string]string `mapstructure:"headers" yaml:"headers" json:"headers"`                   // 自定义头部
}

// TLS TLS配置
type TLS struct {
	CertFile string `mapstructure:"cert_file" yaml:"cert-file" json:"cert_file"` // 证书文件路径
	KeyFile  string `mapstructure:"key_file" yaml:"key-file" json:"key_file"`    // 私钥文件路径
	CAFile   string `mapstructure:"ca_file" yaml:"ca-file" json:"ca_file"`       // CA文件路径
}

// Default 创建默认服务器配置
func Default() *Server {
	return &Server{
		ModuleName:   "server",
		Endpoint:     "localhost:8080",
		Host:         "localhost",
		Port:         8080,
		ReadTimeout:  30,
		WriteTimeout: 30,
		IdleTimeout:  60,
		GrpcPort:     9090,
		EnableHttp:   true,
		EnableGrpc:   false,
		EnableTls:    false,
		TLS: &TLS{
			CertFile: "",
			KeyFile:  "",
			CAFile:   "",
		},
		Headers: make(map[string]string),
	}
}

// Get 返回配置接口
func (s *Server) Get() interface{} {
	return s
}

// Set 设置配置数据
func (s *Server) Set(data interface{}) {
	if cfg, ok := data.(*Server); ok {
		*s = *cfg
	}
}

// Clone 返回配置的副本
func (s *Server) Clone() internal.Configurable {
	clone := &Server{
		ModuleName:   s.ModuleName,
		Endpoint:     s.Endpoint,
		Host:         s.Host,
		Port:         s.Port,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		IdleTimeout:  s.IdleTimeout,
		GrpcPort:     s.GrpcPort,
		EnableHttp:   s.EnableHttp,
		EnableGrpc:   s.EnableGrpc,
		EnableTls:    s.EnableTls,
	}

	if s.TLS != nil {
		clone.TLS = &TLS{
			CertFile: s.TLS.CertFile,
			KeyFile:  s.TLS.KeyFile,
			CAFile:   s.TLS.CAFile,
		}
	}

	clone.Headers = make(map[string]string)
	for k, v := range s.Headers {
		clone.Headers[k] = v
	}

	return clone
}

// Validate 验证配置
func (s *Server) Validate() error {
	return internal.ValidateStruct(s)
}

// WithHost 设置主机地址
func (s *Server) WithHost(host string) *Server {
	s.Host = host
	return s
}

// WithPort 设置端口
func (s *Server) WithPort(port int) *Server {
	s.Port = port
	return s
}

// WithGrpcPort 设置GRPC端口
func (s *Server) WithGrpcPort(port int) *Server {
	s.GrpcPort = port
	return s
}

// WithTimeouts 设置超时配置
func (s *Server) WithTimeouts(read, write, idle int) *Server {
	s.ReadTimeout = read
	s.WriteTimeout = write
	s.IdleTimeout = idle
	return s
}

// WithTLS 设置TLS配置
func (s *Server) WithTLS(certFile, keyFile, caFile string) *Server {
	if s.TLS == nil {
		s.TLS = &TLS{}
	}
	s.TLS.CertFile = certFile
	s.TLS.KeyFile = keyFile
	s.TLS.CAFile = caFile
	s.EnableTls = true
	return s
}

// EnableHTTP 启用HTTP服务
func (s *Server) EnableHTTP() *Server {
	s.EnableHttp = true
	return s
}

// DisableHTTP 禁用HTTP服务
func (s *Server) DisableHTTP() *Server {
	s.EnableHttp = false
	return s
}

// EnableGRPC 启用GRPC服务
func (s *Server) EnableGRPC() *Server {
	s.EnableGrpc = true
	return s
}

// DisableGRPC 禁用GRPC服务
func (s *Server) DisableGRPC() *Server {
	s.EnableGrpc = false
	return s
}

// EnableTLSService 启用TLS
func (s *Server) EnableTLSService() *Server {
	s.EnableTls = true
	return s
}

// DisableTLS 禁用TLS
func (s *Server) DisableTLS() *Server {
	s.EnableTls = false
	return s
}

// AddHeader 添加自定义头部
func (s *Server) AddHeader(key, value string) *Server {
	if s.Headers == nil {
		s.Headers = make(map[string]string)
	}
	s.Headers[key] = value
	return s
}
