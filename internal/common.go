/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 13:55:55
 * @FilePath: \go-config\internal\common.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package internal

import (
	"sync"
)

// Configurable 接口定义了配置的基本行为
type Configurable interface {
	Get() interface{}
	Set(data interface{})
	Validate() error
	Clone() Configurable
}

var mu sync.Mutex // 用于保护实例的创建

// LockFunc 是一个公共方法，用于锁定和解锁
func LockFunc(fn func()) {
	mu.Lock()
	defer mu.Unlock()
	fn()
}
