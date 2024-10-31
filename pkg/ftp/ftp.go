/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 13:50:55
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

type Ftp struct {

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** ftp 服务器Ip和端口 */
	Addr string `mapstructure:"addr"                json:"addr"              yaml:"addr"`

	/** 用户 */
	Username string `mapstructure:"username"            json:"username"          yaml:"username"`

	/** 密码 */
	Password string `mapstructure:"password"            json:"password"          yaml:"password"`

	/** 指定目录 */
	Cwd string `mapstructure:"cwd"                 json:"cwd"               yaml:"cwd"`
}

// NewFtp 创建一个新的 Ftp 实例
func NewFtp(moduleName, addr, username, password, cwd string) *Ftp {
	var ftpInstance *Ftp

	internal.LockFunc(func() {
		ftpInstance = &Ftp{
			ModuleName: moduleName,
			Addr:       addr,
			Username:   username,
			Password:   password,
			Cwd:        cwd,
		}
	})

	return ftpInstance
}

// ToMap 将配置转换为映射
func (f *Ftp) ToMap() map[string]interface{} {
	return internal.ToMap(f)
}

// FromMap 从映射中填充配置
func (f *Ftp) FromMap(data map[string]interface{}) {
	internal.FromMap(f, data)
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
