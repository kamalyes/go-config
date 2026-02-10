/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-02-10 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-02-10 12:00:00
 * @FilePath: \go-config\pkg\wsc\client_capacity.go
 * @Description: WebSocket 客户端容量配置 - 根据客户端类型优化 SendChan 缓冲区大小
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package wsc

// ClientCapacity WebSocket 客户端容量配置
// 用于根据客户端类型优化 SendChan 缓冲区大小，节省内存并提高性能
// 第一阶段优化：根据客户端类型调整容量，节省 40-50% 内存
type ClientCapacity struct {
	// 代理：高优先级，消息频率高，不能丢失
	// 预期消息频率：1000+ msg/s
	// 处理延迟：100ms
	// 缓冲消息数：100+
	Agent int `mapstructure:"agent" yaml:"agent" json:"agent"`

	// 机器人：消息频率很高，需要较大缓冲
	// 预期消息频率：100-500 msg/s
	// 处理延迟：100ms
	// 缓冲消息数：10-50
	Bot int `mapstructure:"bot" yaml:"bot" json:"bot"`

	// 普通客户：消息频率中等，中等优先级
	// 预期消息频率：20-80 msg/s
	// 处理延迟：100ms
	// 缓冲消息数：2-8
	Customer int `mapstructure:"customer" yaml:"customer" json:"customer"`

	// 观察者：消息频率低，低优先级，可以丢失
	// 预期消息频率：1-5 msg/s
	// 处理延迟：100ms
	// 缓冲消息数：0-1
	Observer int `mapstructure:"observer" yaml:"observer" json:"observer"`

	// 管理员：高优先级，类似代理
	// 预期消息频率：500+ msg/s
	// 处理延迟：100ms
	// 缓冲消息数：50+
	Admin int `mapstructure:"admin" yaml:"admin" json:"admin"`

	// VIP 客户：高优先级
	// 预期消息频率：100-200 msg/s
	// 处理延迟：100ms
	// 缓冲消息数：10-20
	VIP int `mapstructure:"vip" yaml:"vip" json:"vip"`

	// 访客：低优先级，但可能有突发流量
	// 预期消息频率：10-40 msg/s
	// 处理延迟：100ms
	// 缓冲消息数：1-4
	Visitor int `mapstructure:"visitor" yaml:"visitor" json:"visitor"`

	// 系统消息：高优先级，不能丢失
	// 预期消息频率：100+ msg/s
	// 处理延迟：100ms
	// 缓冲消息数：10+
	System int `mapstructure:"system" yaml:"system" json:"system"`

	// 默认容量：用于未知类型的客户端
	Default int `mapstructure:"default" yaml:"default" json:"default"`
}

// DefaultClientCapacity 返回默认的客户端容量配置
func DefaultClientCapacity() *ClientCapacity {
	return &ClientCapacity{
		Agent:    256,
		Bot:      128,
		Customer: 96,
		Observer: 16,
		Admin:    256,
		VIP:      128,
		Visitor:  64,
		System:   256,
		Default:  64,
	}
}
