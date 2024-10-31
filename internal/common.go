/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 10:29:20
 * @FilePath: \go-config\internal\base.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package internal

import (
	"reflect"
	"sync"
)

// Configurable 接口定义了配置的基本行为
type Configurable interface {
	Get() interface{}
	Set(data interface{})
	Validate() error
	ToMap() map[string]interface{}
	FromMap(data map[string]interface{})
	Clone() Configurable
}

var mu sync.Mutex // 用于保护实例的创建

// LockFunc 是一个公共方法，用于锁定和解锁
func LockFunc(fn func()) {
	mu.Lock()
	defer mu.Unlock()
	fn()
}

// ToMap 将配置转换为映射
func ToMap[T Configurable](c T) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(c)

	// 确保传入的类型是指针，并获取其元素
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		key := fieldType.Tag.Get("json")
		if key != "" { // 确保标签不为空
			if field.Kind() == reflect.Ptr && !field.IsNil() {
				// 处理嵌套的 Configurable 类型
				if nestedConfig, ok := field.Interface().(Configurable); ok {
					result[key] = ToMap(nestedConfig)
				} else {
					result[key] = field.Interface()
				}
			} else if field.Kind() == reflect.Struct {
				// 处理嵌套的结构体
				if nestedConfig, ok := field.Interface().(Configurable); ok {
					result[key] = ToMap(nestedConfig)
				} else {
					result[key] = field.Interface()
				}
			} else {
				result[key] = field.Interface()
			}
		}
	}

	return result
}

// FromMap 从映射中填充配置
func FromMap[T Configurable](c T, data map[string]interface{}) {
	val := reflect.ValueOf(c)

	// 确保传入的类型是指针，并获取其元素
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldType := typ.Field(i)
		key := fieldType.Tag.Get("json")

		if value, ok := data[key]; ok {
			field := val.Field(i)

			if field.Kind() == reflect.Ptr {
				// 处理嵌套的 Configurable 类型
				if field.IsNil() {
					newConfig := reflect.New(fieldType.Type.Elem()).Interface().(Configurable)
					field.Set(reflect.ValueOf(newConfig))
				}
				nestedConfig := field.Interface().(Configurable)
				FromMap(nestedConfig, value.(map[string]interface{}))
			} else if field.Kind() == reflect.Struct {
				// 处理嵌套的结构体
				nestedConfig := field.Interface().(Configurable)
				FromMap(nestedConfig, value.(map[string]interface{}))
			} else {
				if field.CanSet() {
					field.Set(reflect.ValueOf(value))
				}
			}
		}
	}
}
