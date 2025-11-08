/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:55:15
 * @FilePath: \go-config\internal\common.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package internal

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

// Configurable 接口定义了配置的基本行为
type Configurable interface {
	Get() interface{}
	Set(data interface{})
	Clone() Configurable
	Validate() error
}

var mu sync.Mutex // 用于保护实例的创建

// 全局验证器实例
var validate *validator.Validate

// 初始化验证器
func init() {
	validate = validator.New()
}

// Validate 方法，通用验证
func ValidateStruct(c interface{}) error {
	return validate.Struct(c)
}

// 验证额外函数
func ValidateExtra(extraFunc func() error) error {
	if extraFunc != nil {
		return extraFunc()
	}
	return nil
}

// LockFunc 是一个公共方法，用于锁定和解锁
func LockFunc(fn func()) {
	mu.Lock()
	defer mu.Unlock()
	fn()
}
