/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:31:55
 * @FilePath: \go-config\pkg\jwt\jwt.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package jwt

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// JWT 结构体用于配置 JSON Web Token 的相关参数
type JWT struct {
	ModuleName    string `mapstructure:"MODULE_NAME"              yaml:"modulename"`     // 模块名称
	SigningKey    string `mapstructure:"SIGNING_KEY"              yaml:"signing-key"`    // jwt签名
	ExpiresTime   int64  `mapstructure:"EXPIRES_TIME"             yaml:"expires-time"`   // 过期时间（单位：秒）
	BufferTime    int64  `mapstructure:"BUFFER_TIME"              yaml:"buffer-time"`    // 缓冲时间（单位：秒）
	UseMultipoint bool   `mapstructure:"USE_MULTIPOINT"           yaml:"use-multipoint"` // 多地登录拦截，true 表示拦截，false 表示不拦截
}

// NewJWT 创建一个新的 JWT 实例
func NewJWT(opt *JWT) *JWT {
	var jwtInstance *JWT

	internal.LockFunc(func() {
		jwtInstance = opt
	})

	return jwtInstance
}

// Clone 返回 JWT 配置的副本
func (j *JWT) Clone() internal.Configurable {
	return &JWT{
		ModuleName:    j.ModuleName,
		SigningKey:    j.SigningKey,
		ExpiresTime:   j.ExpiresTime,
		BufferTime:    j.BufferTime,
		UseMultipoint: j.UseMultipoint,
	}
}

// Get 返回 JWT 配置的所有字段
func (j *JWT) Get() interface{} {
	return j
}

// Set 更新 JWT 配置的字段
func (j *JWT) Set(data interface{}) {
	if configData, ok := data.(*JWT); ok {
		j.ModuleName = configData.ModuleName
		j.SigningKey = configData.SigningKey
		j.ExpiresTime = configData.ExpiresTime
		j.BufferTime = configData.BufferTime
		j.UseMultipoint = configData.UseMultipoint
	}
}

// Validate 检查 JWT 配置的有效性
func (j *JWT) Validate() error {
	if j.SigningKey == "" {
		return errors.New("signing key cannot be empty")
	}
	if j.ExpiresTime <= 0 {
		return errors.New("expires time must be greater than 0")
	}
	return nil
}
