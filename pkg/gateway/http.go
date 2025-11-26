/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-26 23:42:05
 * @FilePath: \go-config\pkg\gateway\http.go
 * @Description: 服务器配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package gateway

import (
	"fmt"
	"github.com/kamalyes/go-config/internal"
)

// HTTPServer HTTP服务器配置
type HTTPServer struct {
	ModuleName         string            `mapstructure:"module_name" yaml:"module_name" json:"module_name"`                            // 模块名称
	Host               string            `mapstructure:"host" yaml:"host" json:"host"`                                                 // 主机地址
	Port               int               `mapstructure:"port" yaml:"port" json:"port"`                                                 // 端口
	ReadTimeout        int               `mapstructure:"read_timeout" yaml:"read-timeout" json:"read_timeout"`                         // 读取超时(秒)
	WriteTimeout       int               `mapstructure:"write_timeout" yaml:"write-timeout" json:"write_timeout"`                      // 写入超时(秒)
	IdleTimeout        int               `mapstructure:"idle_timeout" yaml:"idle-timeout" json:"idle_timeout"`                         // 空闲超时(秒)
	MaxHeaderBytes     int               `mapstructure:"max_header_bytes" yaml:"max-header-bytes" json:"max_header_bytes"`             // 最大请求头字节数
	EnableTls          bool              `mapstructure:"enable_tls" yaml:"enable-tls" json:"enable_tls"`                               // 是否启用TLS
	TLS                *TLS              `mapstructure:"tls" yaml:"tls" json:"tls"`                                                    // TLS配置
	Headers            map[string]string `mapstructure:"headers" yaml:"headers" json:"headers"`                                        // 自定义头部
	Endpoint           string            `mapstructure:"_" yaml:"-" json:"_"`                                                          // 完整的服务端点地址（自动计算，不从配置文件读取）
	EnableGzipCompress bool              `mapstructure:"enable_gzip_compress" yaml:"enable-gzip-compress" json:"enable_gzip_compress"` // 是否启用Gzip压缩
}

// TLS TLS配置
type TLS struct {
	CertFile string `mapstructure:"cert_file" yaml:"cert-file" json:"cert_file"` // 证书文件路径
	KeyFile  string `mapstructure:"key_file" yaml:"key-file" json:"key_file"`    // 私钥文件路径
	CAFile   string `mapstructure:"ca_file" yaml:"ca-file" json:"ca_file"`       // CA文件路径
}

// DefaultHTTPServer 创建默认HTTP服务器配置
func DefaultHTTPServer() *HTTPServer {
	h := &HTTPServer{
		ModuleName:     "server",
		Host:           "0.0.0.0",
		Port:           8080,
		ReadTimeout:    30,
		WriteTimeout:   30,
		IdleTimeout:    60,
		MaxHeaderBytes: 1 << 20, // 1MB
		EnableTls:      false,
		TLS: &TLS{
			CertFile: "",
			KeyFile:  "",
			CAFile:   "",
		},
		Headers: map[string]string{
			"x-server": "go-config",
		},
		EnableGzipCompress: true,
	}
	internal.CallAfterLoad(h) // 自动调用 AfterLoad 钩子
	return h
}

// BeforeLoad 配置加载前的钩子函数
func (h *HTTPServer) BeforeLoad() error {
	// 这里可以添加加载前的预处理逻辑
	return nil
}

// AfterLoad 配置加载后的钩子函数 - 用于计算衍生字段
func (h *HTTPServer) AfterLoad() error {
	h.Endpoint = fmt.Sprintf("http://%s:%d", h.Host, h.Port)
	return nil
}

// GetEndpoint 获取服务端点地址（如果未设置则自动计算）
func (h *HTTPServer) GetEndpoint() string {
	if h.Endpoint == "" {
		internal.CallAfterLoad(h)
	}
	return h.Endpoint
}

// Clone 返回配置的副本
func (h *HTTPServer) Clone() *HTTPServer {
	cloned := &HTTPServer{
		ModuleName:         h.ModuleName,
		Host:               h.Host,
		Port:               h.Port,
		ReadTimeout:        h.ReadTimeout,
		WriteTimeout:       h.WriteTimeout,
		IdleTimeout:        h.IdleTimeout,
		MaxHeaderBytes:     h.MaxHeaderBytes,
		EnableTls:          h.EnableTls,
		EnableGzipCompress: h.EnableGzipCompress,
	}
	if h.TLS != nil {
		cloned.TLS = &TLS{
			CertFile: h.TLS.CertFile,
			KeyFile:  h.TLS.KeyFile,
			CAFile:   h.TLS.CAFile,
		}
	}
	cloned.Headers = make(map[string]string)
	for k, v := range h.Headers {
		cloned.Headers[k] = v
	}
	internal.CallAfterLoad(cloned) // 重新计算衍生字段
	return cloned
}

// Validate 验证配置
func (h *HTTPServer) Validate() error {
	return internal.ValidateStruct(h)
}

// WithHost 设置主机地址
func (h *HTTPServer) WithHost(host string) *HTTPServer {
	h.Host = host
	internal.CallAfterLoad(h) // 重新计算 Endpoint
	return h
}

// WithPort 设置端口
func (h *HTTPServer) WithPort(port int) *HTTPServer {
	h.Port = port
	internal.CallAfterLoad(h) // 重新计算 Endpoint
	return h
}

// WithTimeouts 设置超时配置
func (h *HTTPServer) WithTimeouts(read, write, idle int) *HTTPServer {
	h.ReadTimeout = read
	h.WriteTimeout = write
	h.IdleTimeout = idle
	return h
}

// WithTLS 设置TLS配置
func (h *HTTPServer) WithTLS(certFile, keyFile, caFile string) *HTTPServer {
	if h.TLS == nil {
		h.TLS = &TLS{}
	}
	h.TLS.CertFile = certFile
	h.TLS.KeyFile = keyFile
	h.TLS.CAFile = caFile
	h.EnableTls = true
	return h
}

// EnableTLSService 启用TLS
func (h *HTTPServer) EnableTLSService() *HTTPServer {
	h.EnableTls = true
	return h
}

// DisableTLS 禁用TLS
func (h *HTTPServer) DisableTLS() *HTTPServer {
	h.EnableTls = false
	return h
}

// EnableGzip 启用Gzip压缩
func (h *HTTPServer) EnableGzip() *HTTPServer {
	h.EnableGzipCompress = true
	return h
}

// DisableGzip 禁用Gzip压缩
func (h *HTTPServer) DisableGzip() *HTTPServer {
	h.EnableGzipCompress = false
	return h
}

// AddHeader 添加自定义头部
func (h *HTTPServer) AddHeader(key, value string) *HTTPServer {
	if h.Headers == nil {
		h.Headers = make(map[string]string)
	}
	h.Headers[key] = value
	return h
}
