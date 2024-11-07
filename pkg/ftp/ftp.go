/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:31:10
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
	Addr       string `mapstructure:"addr"        yaml:"addr"       json:"addr"        validate:"required"` // ftp 服务器 Ip 和端口
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
		Addr:       f.Addr,
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
		f.Addr = configData.Addr
		f.Username = configData.Username
		f.Password = configData.Password
		f.Cwd = configData.Cwd
	}
}

// Validate 检查 Email 配置的有效性
func (f *Ftp) Validate() error {
	return internal.ValidateStruct(f)
}
