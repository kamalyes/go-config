/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:55:59
 * @FilePath: \go-config\pkg\server\server.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package server

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

type Server struct {

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** 地址 */
	Host string `mapstructure:"host"                     json:"host"                     yaml:"host"`

	/** 端口 */
	Port string `mapstructure:"port"                     json:"port"                     yaml:"port"`

	/** 服务名称 */
	ServerName string `mapstructure:"server-name"                 json:"serverName"               yaml:"server-name"`

	/** 请求根路径 */
	ContextPath string `mapstructure:"context-path"                json:"contextPath"              yaml:"context-path"`

	/** 数据库类型 */
	DataDriver string `mapstructure:"data-driver"                 json:"dataDriver"               yaml:"data-driver"`

	/** 是否开启请求方式检测 */
	HandleMethodNotAllowed bool `mapstructure:"handle-method-not-allowed"   json:"handleMethodNotAllowed"   yaml:"handle-method-not-allowed"`

	/** 语言    */
	Language string `mapstructure:"language"                    json:"language"                 yaml:"language"`
}

// NewServer 创建一个新的 Server 实例
func NewServer(moduleName, host, port, serverName, contextPath, dataDriver string, handleMethodNotAllowed bool, language string) *Server {
	var serverInstance *Server

	internal.LockFunc(func() {
		serverInstance = &Server{
			ModuleName:             moduleName,
			Host:                   host,
			Port:                   port,
			ServerName:             serverName,
			ContextPath:            contextPath,
			DataDriver:             dataDriver,
			HandleMethodNotAllowed: handleMethodNotAllowed,
			Language:               language,
		}
	})
	return serverInstance
}

// ToMap 将配置转换为映射
func (s *Server) ToMap() map[string]interface{} {
	return internal.ToMap(s)
}

// FromMap 从映射中填充配置
func (s *Server) FromMap(data map[string]interface{}) {
	internal.FromMap(s, data)
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
