/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:54:09
 * @FilePath: \go-config\pkg\register\server.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package register

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Server 结构体用于配置服务相关参数
type Server struct {
	ModuleName             string `mapstructure:"MODULE_NAME"                     yaml:"modulename"`                // 模块名称
	Host                   string `mapstructure:"HOST"                            yaml:"host"`                      // 地址
	Port                   string `mapstructure:"PORT"                            yaml:"port"`                      // 端口
	ServerName             string `mapstructure:"SERVER_NAME"                     yaml:"server-name"`               // 服务名称
	ContextPath            string `mapstructure:"CONTEXT_PATH"                    yaml:"context-path"`              // 请求根路径
	DataDriver             string `mapstructure:"DATA_DRIVER"                     yaml:"data-driver"`               // 数据库类型
	HandleMethodNotAllowed bool   `mapstructure:"HANDLE_METHOD_NOT_ALLOWED"       yaml:"handle-method-not-allowed"` // 是否开启请求方式检测
	Language               string `mapstructure:"LANGUAGE"                        yaml:"language"`                  // 语言
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
		Host:                   s.Host,
		Port:                   s.Port,
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
		s.Host = configData.Host
		s.Port = configData.Port
		s.ServerName = configData.ServerName
		s.ContextPath = configData.ContextPath
		s.DataDriver = configData.DataDriver
		s.HandleMethodNotAllowed = configData.HandleMethodNotAllowed
		s.Language = configData.Language
	}
}

// Validate 验证 Server 配置的有效性
func (s *Server) Validate() error {
	if s.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if s.Host == "" {
		return errors.New("host cannot be empty")
	}
	if s.Port == "" {
		return errors.New("port cannot be empty")
	}
	if s.ServerName == "" {
		return errors.New("server name cannot be empty")
	}
	if s.ContextPath == "" {
		return errors.New("context path cannot be empty")
	}
	if s.DataDriver == "" {
		return errors.New("data driver cannot be empty")
	}
	if s.Language == "" {
		return errors.New("language cannot be empty")
	}
	return nil
}
