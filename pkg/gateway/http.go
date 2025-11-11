/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 00:00:00
 * @FilePath: \go-config\pkg\gateway\http.go
 * @Description: HTTP服务器配置
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
	Host                string `mapstructure:"host" yaml:"host" json:"host"`                                           // 主机地址
	Port                int    `mapstructure:"port" yaml:"port" json:"port"`                                           // 端口
	ReadTimeout         int    `mapstructure:"read_timeout" yaml:"read_timeout" json:"read_timeout"`                   // 读取超时(秒)
	WriteTimeout        int    `mapstructure:"write_timeout" yaml:"write_timeout" json:"write_timeout"`                // 写入超时(秒)
	IdleTimeout         int    `mapstructure:"idle_timeout" yaml:"idle_timeout" json:"idle_timeout"`                   // 空闲超时(秒)
	MaxHeaderBytes      int    `mapstructure:"max_header_bytes" yaml:"max_header_bytes" json:"max_header_bytes"`       // 最大请求头字节数
	EnableGzipCompress  bool   `mapstructure:"enable_gzip_compress" yaml:"enable_gzip_compress" json:"enable_gzip_compress"` // 是否启用Gzip压缩
	Endpoint            string `mapstructure:"-" yaml:"-" json:"endpoint"`                                             // 完整的服务端点地址（自动计算，不从配置文件读取）
}

// DefaultHTTPServer 创建默认HTTP服务器配置
func DefaultHTTPServer() *HTTPServer {
	h := &HTTPServer{
		Host:               "0.0.0.0",
		Port:               8080,
		ReadTimeout:        30,
		WriteTimeout:       30,
		IdleTimeout:        60,
		MaxHeaderBytes:     1 << 20, // 1MB
		EnableGzipCompress: true,
	}
	internal.CallAfterLoad(h) // 自动调用 AfterLoad 钩子
	return h
}

// BeforeLoad 配置加载前的钩子函数
// 在配置从文件加载之前调用，可用于设置前置条件
func (h *HTTPServer) BeforeLoad() error {
	// 这里可以添加加载前的预处理逻辑
	return nil
}

// AfterLoad 配置加载后的钩子函数 - 用于计算衍生字段
// 这个方法会在配置加载后自动调用
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
		Host:               h.Host,
		Port:               h.Port,
		ReadTimeout:        h.ReadTimeout,
		WriteTimeout:       h.WriteTimeout,
		IdleTimeout:        h.IdleTimeout,
		MaxHeaderBytes:     h.MaxHeaderBytes,
		EnableGzipCompress: h.EnableGzipCompress,
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
