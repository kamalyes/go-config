/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 01:34:55
 * @FilePath: \go-config\pkg\register\server.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package register

import (
	"github.com/kamalyes/go-config/internal"
)

// Server 结构体用于配置服务相关参数
type Server struct {
	Endpoint               string `mapstructure:"endpoint"                        yaml:"endpoint"          json:"endpoint"         validate:"required"` // 服务端点地址
	ServerName             string `mapstructure:"server-name"                     yaml:"server-name"       json:"server_name"      validate:"required"` // 服务名称
	DataDriver             string `mapstructure:"data-driver"                     yaml:"data-driver"       json:"data_driver"      validate:"required"` // 数据库类型
	HandleMethodNotAllowed bool   `mapstructure:"handle-method-not-allowed"       yaml:"handle-method-not-allowed" json:"handle_method_not_allowed"`    // 是否开启请求方式检测
	ContextPath            string `mapstructure:"context-path"                    yaml:"context-path"      json:"context_path"`                         // 请求根路径
	Language               string `mapstructure:"language"                        yaml:"language"          json:"language"`                             // 语言
	ModuleName             string `mapstructure:"modulename"                      yaml:"modulename"        json:"module_name"`                          // 模块名称
}

// NewServer 创建一个新的 Server 实例
func NewServer(opt *Server) *Server {
	var serverInstance *Server

	internal.LockFunc(func() {
		serverInstance = opt
	})
	return serverInstance
}

// Clone 返回 Server 配置的副本
func (s *Server) Clone() internal.Configurable {
	return &Server{
		ModuleName:             s.ModuleName,
		Endpoint:               s.Endpoint,
		ServerName:             s.ServerName,
		ContextPath:            s.ContextPath,
		DataDriver:             s.DataDriver,
		HandleMethodNotAllowed: s.HandleMethodNotAllowed,
		Language:               s.Language,
	}
}

// Get 返回 Server 配置的所有字段
func (s *Server) Get() interface{} {
	return s
}

// Set 更新 Server 配置的字段
func (s *Server) Set(data interface{}) {
	if configData, ok := data.(*Server); ok {
		s.ModuleName = configData.ModuleName
		s.Endpoint = configData.Endpoint
		s.ServerName = configData.ServerName
		s.ContextPath = configData.ContextPath
		s.DataDriver = configData.DataDriver
		s.HandleMethodNotAllowed = configData.HandleMethodNotAllowed
		s.Language = configData.Language
	}
}

// Validate 验证 Server 配置的有效性
func (s *Server) Validate() error {
	return internal.ValidateStruct(s)
}

// DefaultServer 返回默认Server配置
func DefaultServer() Server {
	return Server{
		ModuleName:             "server",
		Endpoint:               ":8080",
		ServerName:             "go-config-server",
		DataDriver:             "mysql",
		HandleMethodNotAllowed: false,
		ContextPath:            "/",
		Language:               "zh-cn",
	}
}

// Default 返回默认Server配置的指针，支持链式调用
func Default() *Server {
	config := DefaultServer()
	return &config
}

// WithModuleName 设置模块名称
func (s *Server) WithModuleName(moduleName string) *Server {
	s.ModuleName = moduleName
	return s
}

// WithEndpoint 设置服务端点地址
func (s *Server) WithEndpoint(endpoint string) *Server {
	s.Endpoint = endpoint
	return s
}

// WithServerName 设置服务名称
func (s *Server) WithServerName(serverName string) *Server {
	s.ServerName = serverName
	return s
}

// WithDataDriver 设置数据库类型
func (s *Server) WithDataDriver(dataDriver string) *Server {
	s.DataDriver = dataDriver
	return s
}

// WithHandleMethodNotAllowed 设置是否开启请求方式检测
func (s *Server) WithHandleMethodNotAllowed(handleMethodNotAllowed bool) *Server {
	s.HandleMethodNotAllowed = handleMethodNotAllowed
	return s
}

// WithContextPath 设置请求根路径
func (s *Server) WithContextPath(contextPath string) *Server {
	s.ContextPath = contextPath
	return s
}

// WithLanguage 设置语言
func (s *Server) WithLanguage(language string) *Server {
	s.Language = language
	return s
}
