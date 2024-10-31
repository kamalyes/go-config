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
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Ftp 结构体用于配置 FTP 服务器的相关参数
type Ftp struct {
	ModuleName string `mapstructure:"MODULE_NAME"              yaml:"modulename"` // 模块名称
	Addr       string `mapstructure:"ADDR"                     yaml:"addr"`       // ftp 服务器Ip和端口
	Username   string `mapstructure:"USERNAME"                 yaml:"username"`   // 用户
	Password   string `mapstructure:"PASSWORD"                 yaml:"password"`   // 密码
	Cwd        string `mapstructure:"CWD"                      yaml:"cwd"`        // 指定目录
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

// Validate 检查 Ftp 配置的有效性
func (f *Ftp) Validate() error {
	if f.Addr == "" {
		return errors.New("addr cannot be empty")
	}
	if f.Username == "" {
		return errors.New("username cannot be empty")
	}
	if f.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}
