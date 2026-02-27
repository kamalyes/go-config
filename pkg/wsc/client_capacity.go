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

// CapacityEstimation Hub 容量估算配置
// 用于优化 Hub 内部 map 的初始容量，减少扩容次数，提升性能
type CapacityEstimation struct {
	// Nodes 预估的节点数量
	// 用于优化 nodes map 的初始容量
	// 默认值：10
	// 建议：根据实际部署规模设置
	//   - 单机部署：1
	//   - 小型集群：3-5
	//   - 中型集群：10-20
	//   - 大型集群：50+
	Nodes int `mapstructure:"nodes" yaml:"nodes" json:"nodes"`

	// Clients 预估的客户端连接数
	// 用于优化内存分配，减少 map 扩容次数
	// 默认值：1000
	// 建议：根据实际业务场景设置
	//   - 小型应用：100-500
	//   - 中型应用：1000-5000
	//   - 大型应用：10000+
	Clients int `mapstructure:"clients" yaml:"clients" json:"clients"`

	// AgentRatio 预估客服连接占比
	// 用于计算 agentClients map 的初始容量
	// 默认值：0.1（10%）
	// 说明：客服通常占总连接数的 5-15%
	AgentRatio float64 `mapstructure:"agent-ratio" yaml:"agent-ratio" json:"agentRatio"`

	// ObserverRatio 预估观察者（管理员）连接占比
	// 用于计算 observerClients map 的初始容量
	// 默认值：0.05（5%）
	// 说明：观察者通常占总连接数的 2-10%
	ObserverRatio float64 `mapstructure:"observer-ratio" yaml:"observer-ratio" json:"observerRatio"`

	// SSERatio 预估 SSE 连接占比
	// 用于计算 sseClients map 的初始容量
	// 默认值：0.1（10%）
	// 说明：SSE 连接通常占总连接数的 5-20%
	SSERatio float64 `mapstructure:"sse-ratio" yaml:"sse-ratio" json:"sseRatio"`
}

// DefaultCapacityEstimation 返回默认的容量估算配置
func DefaultCapacityEstimation() *CapacityEstimation {
	return &CapacityEstimation{
		Clients:       1000, // 默认预估 1000 个并发连接
		Nodes:         10,
		AgentRatio:    0.1,  // 客服占 10%
		ObserverRatio: 0.05, // 观察者占 5%
		SSERatio:      0.1,  // SSE 连接占 10%
	}
}

// GetClients 获取预估客户端连接数
func (c *CapacityEstimation) GetClients() int {
	return c.Clients
}

// GetAgentRatio 获取预估客服连接占比
func (c *CapacityEstimation) GetAgentRatio() float64 {
	return c.AgentRatio
}

// GetObserverRatio 获取预估观察者连接占比
func (c *CapacityEstimation) GetObserverRatio() float64 {
	return c.ObserverRatio
}

// GetSSERatio 获取预估 SSE 连接占比
func (c *CapacityEstimation) GetSSERatio() float64 {
	return c.SSERatio
}

// CalculateCapacities 计算各个 map 的初始容量
// 返回值：clients, userToClients, agentClients, observerClients, sseClients 的容量
func (c *CapacityEstimation) CalculateCapacities() (int, int, int, int, int) {
	estimatedClients := c.GetClients()
	agentRatio := c.GetAgentRatio()
	observerRatio := c.GetObserverRatio()
	sseRatio := c.GetSSERatio()

	// 计算各个 map 的容量
	clientsCapacity := estimatedClients
	userToClientsCapacity := estimatedClients // 用户数通常接近连接数
	agentClientsCapacity := int(float64(estimatedClients) * agentRatio)
	observerClientsCapacity := int(float64(estimatedClients) * observerRatio)
	sseClientsCapacity := int(float64(estimatedClients) * sseRatio)

	return clientsCapacity, userToClientsCapacity, agentClientsCapacity, observerClientsCapacity, sseClientsCapacity
}
