/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-18 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-18 13:07:56
 * @FilePath: \engine-im-service\go-config\pkg\gateway\json.go
 * @Description: JSON序列化配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package gateway

import (
	"github.com/kamalyes/go-config/internal"
)

// JSON JSON序列化配置
type JSON struct {
	UseProtoNames   bool `mapstructure:"use_proto_names" yaml:"use-proto-names" json:"use_proto_names"`    // 使用proto字段名(snake_case)，如 snake_case 而不是 snakeCase
	EmitUnpopulated bool `mapstructure:"emit_unpopulated" yaml:"emit-unpopulated" json:"emit_unpopulated"` // 输出所有字段，包括零值字段
	DiscardUnknown  bool `mapstructure:"discard_unknown" yaml:"discard-unknown" json:"discard_unknown"`    // 反序列化时忽略未知字段
}

// DefaultJSON 创建默认JSON配置
func DefaultJSON() *JSON {
	return &JSON{
		UseProtoNames:   true,
		EmitUnpopulated: true,
		DiscardUnknown:  true,
	}
}

// BeforeLoad 配置加载前的钩子函数
func (j *JSON) BeforeLoad() error {
	return nil
}

// AfterLoad 配置加载后的钩子函数
func (j *JSON) AfterLoad() error {
	return nil
}

// Clone 返回配置的副本
func (j *JSON) Clone() *JSON {
	return &JSON{
		UseProtoNames:   j.UseProtoNames,
		EmitUnpopulated: j.EmitUnpopulated,
		DiscardUnknown:  j.DiscardUnknown,
	}
}

// Validate 验证配置
func (j *JSON) Validate() error {
	return internal.ValidateStruct(j)
}

// WithUseProtoNames 设置是否使用proto字段名
func (j *JSON) WithUseProtoNames(use bool) *JSON {
	j.UseProtoNames = use
	return j
}

// WithEmitUnpopulated 设置是否输出零值字段
func (j *JSON) WithEmitUnpopulated(emit bool) *JSON {
	j.EmitUnpopulated = emit
	return j
}

// WithDiscardUnknown 设置是否忽略未知字段
func (j *JSON) WithDiscardUnknown(discard bool) *JSON {
	j.DiscardUnknown = discard
	return j
}
