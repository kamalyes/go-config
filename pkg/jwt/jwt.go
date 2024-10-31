/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 11:22:58
 * @FilePath: \go-config\jwt\jwt.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package jwt

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

type JWT struct {
	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** jwt签名 */
	SigningKey string `mapstructure:"signing-key"         json:"signingKey"        yaml:"signing-key"`

	/** 过期时间 */
	ExpiresTime int64 `mapstructure:"expires-time"        json:"expiresTime"       yaml:"expires-time"`

	/** 缓冲时间 */
	BufferTime int64 `mapstructure:"buffer-time"         json:"bufferTime"        yaml:"buffer-time"`

	/** 多地登录拦截 true 拦截 fasle 不拦截 */
	UseMultipoint bool `mapstructure:"use-multipoint"      json:"useMultipoint"     yaml:"use-multipoint"`
}

// NewJWT 创建一个新的 JWT 实例
func NewJWT(moduleName, signingKey string, expiresTime, bufferTime int64, useMultipoint bool) *JWT {
	var jwtInstance *JWT

	internal.LockFunc(func() {
		jwtInstance = &JWT{
			ModuleName:    moduleName,
			SigningKey:    signingKey,
			ExpiresTime:   expiresTime,
			BufferTime:    bufferTime,
			UseMultipoint: useMultipoint,
		}
	})

	return jwtInstance
}

// ToMap 将配置转换为映射
func (j *JWT) ToMap() map[string]interface{} {
	return internal.ToMap(j)
}

// FromMap 从映射中填充配置
func (j *JWT) FromMap(data map[string]interface{}) {
	internal.FromMap(j, data)
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
