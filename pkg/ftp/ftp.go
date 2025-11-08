/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-09 01:32:31
 * @FilePath: \go-config\pkg\ftp\ftp.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package ftp

import (
	"github.com/kamalyes/go-config/internal"
)

// Ftp 结构体用于配置 FTP 服务器的相关参数
type Ftp struct {
	Endpoint   string `mapstructure:"endpoint"    yaml:"endpoint"   json:"endpoint"    validate:"required"` // FTP 服务器端点地址
	Username   string `mapstructure:"username"    yaml:"username"   json:"username"    validate:"required"` // 用户
	Password   string `mapstructure:"password"    yaml:"password"   json:"password"    validate:"required"` // 密码
	Cwd        string `mapstructure:"cwd"         yaml:"cwd"        json:"cwd"`                             // 指定目录
	ModuleName string `mapstructure:"modulename"  yaml:"modulename" json:"module_name"`                     // 模块名称
}

// NewFtp 创建一个新的 Ftp 实例
func NewFtp(opt *Ftp) *Ftp {
	var ftpInstance *Ftp

	internal.LockFunc(func() {
		ftpInstance = opt
	})

	return ftpInstance
}

// Clone 返回 Ftp 配置的副本
func (f *Ftp) Clone() internal.Configurable {
	return &Ftp{
		ModuleName: f.ModuleName,
		Endpoint:   f.Endpoint,
		Username:   f.Username,
		Password:   f.Password,
		Cwd:        f.Cwd,
	}
}

// Get 返回 Ftp 配置的所有字段
func (f *Ftp) Get() interface{} {
	return f
}

// Set 更新 Ftp 配置的字段
func (f *Ftp) Set(data interface{}) {
	if configData, ok := data.(*Ftp); ok {
		f.ModuleName = configData.ModuleName
		f.Endpoint = configData.Endpoint
		f.Username = configData.Username
		f.Password = configData.Password
		f.Cwd = configData.Cwd
	}
}

// Validate 检查 Ftp 配置的有效性
func (f *Ftp) Validate() error {
	return internal.ValidateStruct(f)
}

// DefaultFtp 返回默认Ftp配置
func DefaultFtp() Ftp {
	return Ftp{
		ModuleName: "ftp",
		Endpoint:   "127.0.0.1:21",
		Username:   "",
		Password:   "",
		Cwd:        "/",
	}
}

// Default 返回默认Ftp配置的指针，支持链式调用
func Default() *Ftp {
	config := DefaultFtp()
	return &config
}

// WithModuleName 设置模块名称
func (f *Ftp) WithModuleName(moduleName string) *Ftp {
	f.ModuleName = moduleName
	return f
}

// WithEndpoint 设置FTP服务器端点地址
func (f *Ftp) WithEndpoint(endpoint string) *Ftp {
	f.Endpoint = endpoint
	return f
}

// WithUsername 设置用户名
func (f *Ftp) WithUsername(username string) *Ftp {
	f.Username = username
	return f
}

// WithPassword 设置密码
func (f *Ftp) WithPassword(password string) *Ftp {
	f.Password = password
	return f
}

// WithCwd 设置工作目录
func (f *Ftp) WithCwd(cwd string) *Ftp {
	f.Cwd = cwd
	return f
}
