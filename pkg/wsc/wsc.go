/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-13 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-16 20:25:08
 * @FilePath: \go-config\pkg\wsc\wsc.go
 * @Description: WebSocket 通信完整配置模块（包含分布式、Redis、群组、广播等）
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package wsc

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-toolbox/pkg/safe"
	"time"
)

// WSC WebSocket 通信完整配置
type WSC struct {
	// === 基础配置 ===
	Enabled           bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                     // 是否启用
	NodeIP            string        `mapstructure:"node_ip" yaml:"node-ip" json:"node_ip"`                                     // 节点IP
	NodePort          int           `mapstructure:"node_port" yaml:"node-port" json:"node_port"`                               // 节点端口
	HeartbeatInterval int           `mapstructure:"heartbeat_interval" yaml:"heartbeat-interval" json:"heartbeat_interval"`    // 心跳间隔(秒)
	ClientTimeout     int           `mapstructure:"client_timeout" yaml:"client-timeout" json:"client_timeout"`                // 客户端超时(秒)
	MessageBufferSize int           `mapstructure:"message_buffer_size" yaml:"message-buffer-size" json:"message_buffer_size"` // 消息缓冲区大小
	WebSocketOrigins  []string      `mapstructure:"websocket_origins" yaml:"websocket-origins" json:"websocket_origins"`       // 允许的WebSocket Origin
	WriteWait         time.Duration `mapstructure:"write_wait" yaml:"write-wait" json:"write_wait"`                            // 写超时
	MaxMessageSize    int64         `mapstructure:"max_message_size" yaml:"max-message-size" json:"max_message_size"`          // 最大消息长度
	MinRecTime        time.Duration `mapstructure:"min_rec_time" yaml:"min-rec-time" json:"min_rec_time"`                      // 最小重连时间
	MaxRecTime        time.Duration `mapstructure:"max_rec_time" yaml:"max-rec-time" json:"max_rec_time"`                      // 最大重连时间
	RecFactor         float64       `mapstructure:"rec_factor" yaml:"rec-factor" json:"rec_factor"`                            // 重连因子
	AutoReconnect     bool          `mapstructure:"auto_reconnect" yaml:"auto-reconnect" json:"auto_reconnect"`                // 是否自动重连

	// === SSE 配置 ===
	SSEHeartbeat     int `mapstructure:"sse_heartbeat" yaml:"sse-heartbeat" json:"sse_heartbeat"`                // SSE心跳间隔(秒)
	SSETimeout       int `mapstructure:"sse_timeout" yaml:"sse-timeout" json:"sse_timeout"`                      // SSE超时(秒)
	SSEMessageBuffer int `mapstructure:"sse_message_buffer" yaml:"sse-message-buffer" json:"sse_message_buffer"` // SSE消息缓冲区大小

	// === 分布式节点配置 ===
	Distributed *Distributed `mapstructure:"distributed" yaml:"distributed,omitempty" json:"distributed,omitempty"` // 分布式配置

	// === Redis 配置（用于分布式消息）- 复用 cache.Redis ===
	Redis *cache.Redis `mapstructure:"redis" yaml:"redis,omitempty" json:"redis,omitempty"` // Redis配置

	// === 群组/广播配置 ===
	Group *Group `mapstructure:"group" yaml:"group,omitempty" json:"group,omitempty"` // 群组配置

	// === 工单配置 ===
	Ticket *Ticket `mapstructure:"ticket" yaml:"ticket,omitempty" json:"ticket,omitempty"` // 工单配置

	// === 性能配置 ===
	Performance *Performance `mapstructure:"performance" yaml:"performance,omitempty" json:"performance,omitempty"` // 性能配置

	// === 安全配置 ===
	Security *Security `mapstructure:"security" yaml:"security,omitempty" json:"security,omitempty"` // 安全配置

	// === VIP等级配置 ===
	VIP *VIP `mapstructure:"vip" yaml:"vip,omitempty" json:"vip,omitempty"` // VIP等级配置

	// === 增强功能配置 ===
	Enhancement *Enhancement `mapstructure:"enhancement" yaml:"enhancement,omitempty" json:"enhancement,omitempty"` // 增强功能配置
}

// Distributed 分布式节点配置
type Distributed struct {
	Enabled             bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                           // 是否启用分布式
	NodeDiscovery       string `mapstructure:"node_discovery" yaml:"node-discovery" json:"node_discovery"`                      // 节点发现方式: redis, etcd, consul
	NodeSyncInterval    int    `mapstructure:"node_sync_interval" yaml:"node-sync-interval" json:"node_sync_interval"`          // 节点同步间隔(秒)
	MessageRouting      string `mapstructure:"message_routing" yaml:"message-routing" json:"message_routing"`                   // 消息路由策略: hash, random, round-robin
	EnableLoadBalance   bool   `mapstructure:"enable_load_balance" yaml:"enable-load-balance" json:"enable_load_balance"`       // 是否启用负载均衡
	HealthCheckInterval int    `mapstructure:"health_check_interval" yaml:"health-check-interval" json:"health_check_interval"` // 健康检查间隔(秒)
	NodeTimeout         int    `mapstructure:"node_timeout" yaml:"node-timeout" json:"node_timeout"`                            // 节点超时(秒)
	ClusterName         string `mapstructure:"cluster_name" yaml:"cluster-name" json:"cluster_name"`                            // 集群名称
}

// DefaultRedis 默认Redis配置（复用cache包）
func DefaultRedis() *cache.Redis {
	return cache.DefaultRedis().
		WithModuleName("wsc").
		WithDB(0).
		WithPoolSize(10).
		WithMinIdleConns(2).
		WithMaxRetries(3)
}

// Group 群组/广播配置
type Group struct {
	Enabled             bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                           // 是否启用群组功能
	MaxGroupSize        int  `mapstructure:"max_group_size" yaml:"max-group-size" json:"max_group_size"`                      // 最大群组人数
	MaxGroupsPerUser    int  `mapstructure:"max_groups_per_user" yaml:"max-groups-per-user" json:"max_groups_per_user"`       // 每个用户最大群组数
	EnableBroadcast     bool `mapstructure:"enable_broadcast" yaml:"enable-broadcast" json:"enable_broadcast"`                // 是否启用全局广播
	BroadcastRateLimit  int  `mapstructure:"broadcast_rate_limit" yaml:"broadcast-rate-limit" json:"broadcast_rate_limit"`    // 广播频率限制(次/分钟)
	GroupCacheExpire    int  `mapstructure:"group_cache_expire" yaml:"group-cache-expire" json:"group_cache_expire"`          // 群组缓存过期时间(秒)
	AutoCreateGroup     bool `mapstructure:"auto_create_group" yaml:"auto-create-group" json:"auto_create_group"`             // 是否自动创建群组
	EnableMessageRecord bool `mapstructure:"enable_message_record" yaml:"enable-message-record" json:"enable_message_record"` // 是否启用消息记录
}

// Ticket 工单配置
type Ticket struct {
	Enabled              bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                              // 是否启用工单功能
	MaxTicketsPerAgent   int    `mapstructure:"max_tickets_per_agent" yaml:"max-tickets-per-agent" json:"max_tickets_per_agent"`    // 每个客服最大工单数
	AutoAssign           bool   `mapstructure:"auto_assign" yaml:"auto-assign" json:"auto_assign"`                                  // 是否自动分配工单
	AssignStrategy       string `mapstructure:"assign_strategy" yaml:"assign-strategy" json:"assign_strategy"`                      // 分配策略: random, load-balance, skill-based
	TicketTimeout        int    `mapstructure:"ticket_timeout" yaml:"ticket-timeout" json:"ticket_timeout"`                         // 工单超时(秒)
	EnableQueueing       bool   `mapstructure:"enable_queueing" yaml:"enable-queueing" json:"enable_queueing"`                      // 是否启用排队
	QueueTimeout         int    `mapstructure:"queue_timeout" yaml:"queue-timeout" json:"queue_timeout"`                            // 排队超时(秒)
	NotifyTimeout        int    `mapstructure:"notify_timeout" yaml:"notify-timeout" json:"notify_timeout"`                         // 通知超时(秒)
	EnableTransfer       bool   `mapstructure:"enable_transfer" yaml:"enable-transfer" json:"enable_transfer"`                      // 是否启用工单转接
	TransferMaxTimes     int    `mapstructure:"transfer_max_times" yaml:"transfer-max-times" json:"transfer_max_times"`             // 最大转接次数
	EnableOfflineMessage bool   `mapstructure:"enable_offline_message" yaml:"enable-offline-message" json:"enable_offline_message"` // 是否启用离线消息
	OfflineMessageExpire int    `mapstructure:"offline_message_expire" yaml:"offline-message-expire" json:"offline_message_expire"` // 离线消息过期时间(秒)
	EnableAck            bool   `mapstructure:"enable_ack" yaml:"enable-ack" json:"enable_ack"`                                     // 是否启用ACK确认
	AckTimeoutMs         int    `mapstructure:"ack_timeout_ms" yaml:"ack-timeout-ms" json:"ack_timeout_ms"`                         // ACK超时时间(毫秒)
	MaxRetry             int    `mapstructure:"max_retry" yaml:"max-retry" json:"max_retry"`                                        // 最大重试次数
}

// Performance 性能配置
type Performance struct {
	MaxConnectionsPerNode int  `mapstructure:"max_connections_per_node" yaml:"max-connections-per-node" json:"max_connections_per_node"` // 每个节点最大连接数
	ReadBufferSize        int  `mapstructure:"read_buffer_size" yaml:"read-buffer-size" json:"read_buffer_size"`                         // 读缓冲区大小(KB)
	WriteBufferSize       int  `mapstructure:"write_buffer_size" yaml:"write-buffer-size" json:"write_buffer_size"`                      // 写缓冲区大小(KB)
	EnableCompression     bool `mapstructure:"enable_compression" yaml:"enable-compression" json:"enable_compression"`                   // 是否启用压缩
	CompressionLevel      int  `mapstructure:"compression_level" yaml:"compression-level" json:"compression_level"`                      // 压缩级别(1-9)
	EnableMetrics         bool `mapstructure:"enable_metrics" yaml:"enable-metrics" json:"enable_metrics"`                               // 是否启用性能指标
	MetricsInterval       int  `mapstructure:"metrics_interval" yaml:"metrics-interval" json:"metrics_interval"`                         // 指标采集间隔(秒)
	EnableSlowLog         bool `mapstructure:"enable_slow_log" yaml:"enable-slow-log" json:"enable_slow_log"`                            // 是否启用慢日志
	SlowLogThreshold      int  `mapstructure:"slow_log_threshold" yaml:"slow-log-threshold" json:"slow_log_threshold"`                   // 慢日志阈值(毫秒)
}

// Security 安全配置
type Security struct {
	EnableAuth        bool     `mapstructure:"enable_auth" yaml:"enable-auth" json:"enable_auth"`                         // 是否启用认证
	EnableEncryption  bool     `mapstructure:"enable_encryption" yaml:"enable-encryption" json:"enable_encryption"`       // 是否启用加密
	EnableRateLimit   bool     `mapstructure:"enable_rate_limit" yaml:"enable-rate-limit" json:"enable_rate_limit"`       // 是否启用限流
	MaxMessageSize    int      `mapstructure:"max_message_size" yaml:"max-message-size" json:"max_message_size"`          // 最大消息大小(KB)
	AllowedUserTypes  []string `mapstructure:"allowed_user_types" yaml:"allowed-user-types" json:"allowed_user_types"`    // 允许的用户类型
	BlockedIPs        []string `mapstructure:"blocked_ips" yaml:"blocked-ips" json:"blocked_ips"`                         // 黑名单IP
	WhitelistIPs      []string `mapstructure:"whitelist_ips" yaml:"whitelist-ips" json:"whitelist_ips"`                   // 白名单IP
	EnableIPWhitelist bool     `mapstructure:"enable_ip_whitelist" yaml:"enable-ip-whitelist" json:"enable_ip_whitelist"` // 是否启用IP白名单
	TokenExpiration   int      `mapstructure:"token_expiration" yaml:"token-expiration" json:"token_expiration"`          // Token过期时间(秒)
	MaxLoginAttempts  int      `mapstructure:"max_login_attempts" yaml:"max-login-attempts" json:"max_login_attempts"`    // 最大登录尝试次数
	LoginLockDuration int      `mapstructure:"login_lock_duration" yaml:"login-lock-duration" json:"login_lock_duration"` // 登录锁定时长(秒)
}

// VIP VIP等级配置
type VIP struct {
	Enabled               bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                                 // 是否启用VIP等级
	MaxLevel              int      `mapstructure:"max_level" yaml:"max-level" json:"max_level"`                                           // 最大VIP等级 (0-8)
	DefaultLevel          int      `mapstructure:"default_level" yaml:"default-level" json:"default_level"`                               // 默认VIP等级
	PriorityMultiplier    float64  `mapstructure:"priority_multiplier" yaml:"priority-multiplier" json:"priority_multiplier"`             // VIP优先级乘数
	MessagePriorityBonus  int      `mapstructure:"message_priority_bonus" yaml:"message-priority-bonus" json:"message_priority_bonus"`    // VIP消息优先级加分
	QueuePriority         bool     `mapstructure:"queue_priority" yaml:"queue-priority" json:"queue_priority"`                            // VIP用户是否享受队列优先级
	CustomServicePriority bool     `mapstructure:"custom_service_priority" yaml:"custom-service-priority" json:"custom_service_priority"` // VIP用户是否享受专属客服
	SpecialFeatures       []string `mapstructure:"special_features" yaml:"special-features" json:"special_features"`                      // VIP特殊功能列表
	UpgradeRules          []string `mapstructure:"upgrade_rules" yaml:"upgrade-rules" json:"upgrade_rules"`                               // VIP升级规则
}

// Enhancement 增强功能配置
type Enhancement struct {
	Enabled              bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                              // 是否启用增强功能
	SmartRouting         bool   `mapstructure:"smart_routing" yaml:"smart-routing" json:"smart_routing"`                            // 启用智能路由
	LoadBalancing        bool   `mapstructure:"load_balancing" yaml:"load-balancing" json:"load_balancing"`                         // 启用负载均衡
	SmartQueue           bool   `mapstructure:"smart_queue" yaml:"smart-queue" json:"smart_queue"`                                  // 启用智能队列
	Monitoring           bool   `mapstructure:"monitoring" yaml:"monitoring" json:"monitoring"`                                     // 启用监控
	ClusterManagement    bool   `mapstructure:"cluster_management" yaml:"cluster-management" json:"cluster_management"`             // 启用集群管理
	RuleEngine           bool   `mapstructure:"rule_engine" yaml:"rule-engine" json:"rule_engine"`                                  // 启用规则引擎
	CircuitBreaker       bool   `mapstructure:"circuit_breaker" yaml:"circuit-breaker" json:"circuit_breaker"`                      // 启用熔断器
	MessageFiltering     bool   `mapstructure:"message_filtering" yaml:"message-filtering" json:"message_filtering"`                // 启用消息过滤
	PerformanceTracking  bool   `mapstructure:"performance_tracking" yaml:"performance-tracking" json:"performance_tracking"`       // 启用性能追踪
	AdvancedMetrics      bool   `mapstructure:"advanced_metrics" yaml:"advanced-metrics" json:"advanced_metrics"`                   // 启用高级指标
	AlertSystem          bool   `mapstructure:"alert_system" yaml:"alert-system" json:"alert_system"`                               // 启用警报系统
	FailureThreshold     int    `mapstructure:"failure_threshold" yaml:"failure-threshold" json:"failure_threshold"`                // 熔断器失败阈值
	SuccessThreshold     int    `mapstructure:"success_threshold" yaml:"success-threshold" json:"success_threshold"`                // 熔断器成功阈值
	CircuitTimeout       int    `mapstructure:"circuit_timeout" yaml:"circuit-timeout" json:"circuit_timeout"`                      // 熔断器超时(秒)
	MaxQueueSize         int    `mapstructure:"max_queue_size" yaml:"max-queue-size" json:"max_queue_size"`                         // 智能队列最大大小
	MetricsInterval      int    `mapstructure:"metrics_interval" yaml:"metrics-interval" json:"metrics_interval"`                   // 指标采集间隔(秒)
	MaxSamples           int    `mapstructure:"max_samples" yaml:"max-samples" json:"max_samples"`                                  // 性能样本最大数量
	LoadBalanceAlgorithm string `mapstructure:"load_balance_algorithm" yaml:"load-balance-algorithm" json:"load_balance_algorithm"` // 负载均衡算法: round-robin, least-connections, weighted-random, consistent-hash
}

// Default 创建默认 WSC 配置
func Default() *WSC {
	return &WSC{
		Enabled:           false,
		NodeIP:            "0.0.0.0",
		NodePort:          8080,
		HeartbeatInterval: 30,
		ClientTimeout:     90,
		MessageBufferSize: 256,
		WebSocketOrigins:  []string{"*"},
		WriteWait:         10 * time.Second,
		MaxMessageSize:    512,
		MinRecTime:        2 * time.Second,
		MaxRecTime:        60 * time.Second,
		RecFactor:         1.5,
		AutoReconnect:     true,
		SSEHeartbeat:      30,
		SSETimeout:        120,
		SSEMessageBuffer:  100,
		Distributed:       DefaultDistributed(),
		Redis:             DefaultRedis(),
		Group:             DefaultGroup(),
		Ticket:            DefaultTicket(),
		Performance:       DefaultPerformance(),
		Security:          DefaultSecurity(),
		VIP:               DefaultVIP(),
		Enhancement:       DefaultEnhancement(),
	}
}

// DefaultDistributed 默认分布式配置
func DefaultDistributed() *Distributed {
	return &Distributed{
		Enabled:             false,
		NodeDiscovery:       "redis",
		NodeSyncInterval:    30,
		MessageRouting:      "hash",
		EnableLoadBalance:   true,
		HealthCheckInterval: 10,
		NodeTimeout:         60,
		ClusterName:         "wsc-cluster",
	}
}

// ========== Distributed 自身链式调用方法 ==========

// Enable 启用分布式功能
func (d *Distributed) Enable() *Distributed {
	d.Enabled = true
	return d
}

// Disable 禁用分布式功能
func (d *Distributed) Disable() *Distributed {
	d.Enabled = false
	return d
}

// WithNodeDiscovery 设置节点发现方式
func (d *Distributed) WithNodeDiscovery(discovery string) *Distributed {
	d.NodeDiscovery = discovery
	return d
}

// WithSyncInterval 设置节点同步间隔
func (d *Distributed) WithSyncInterval(intervalSeconds int) *Distributed {
	d.NodeSyncInterval = intervalSeconds
	return d
}

// WithMessageRouting 设置消息路由策略
func (d *Distributed) WithMessageRouting(routing string) *Distributed {
	d.MessageRouting = routing
	return d
}

// WithLoadBalance 启用/禁用负载均衡
func (d *Distributed) WithLoadBalance(enabled bool) *Distributed {
	d.EnableLoadBalance = enabled
	return d
}

// WithHealthCheck 设置健康检查配置
func (d *Distributed) WithHealthCheck(intervalSeconds, timeoutSeconds int) *Distributed {
	d.HealthCheckInterval = intervalSeconds
	d.NodeTimeout = timeoutSeconds
	return d
}

// WithClusterName 设置集群名称
func (d *Distributed) WithClusterName(clusterName string) *Distributed {
	d.ClusterName = clusterName
	return d
}

// DefaultGroup 默认群组配置
func DefaultGroup() *Group {
	return &Group{
		Enabled:             false,
		MaxGroupSize:        500,
		MaxGroupsPerUser:    100,
		EnableBroadcast:     true,
		BroadcastRateLimit:  10,
		GroupCacheExpire:    3600,
		AutoCreateGroup:     false,
		EnableMessageRecord: true, // 默认启用消息记录
	}
}

// ========== Group 自身链式调用方法 ==========

// Enable 启用群组功能
func (g *Group) Enable() *Group {
	g.Enabled = true
	return g
}

// Disable 禁用群组功能
func (g *Group) Disable() *Group {
	g.Enabled = false
	return g
}

// WithMaxSize 设置群组最大人数
func (g *Group) WithMaxSize(maxSize int) *Group {
	g.MaxGroupSize = maxSize
	return g
}

// WithMaxGroupsPerUser 设置每个用户最大群组数
func (g *Group) WithMaxGroupsPerUser(maxGroups int) *Group {
	g.MaxGroupsPerUser = maxGroups
	return g
}

// WithBroadcast 启用/禁用全局广播
func (g *Group) WithBroadcast(enabled bool, rateLimit int) *Group {
	g.EnableBroadcast = enabled
	g.BroadcastRateLimit = rateLimit
	return g
}

// WithCacheExpire 设置群组缓存过期时间
func (g *Group) WithCacheExpire(expireSeconds int) *Group {
	g.GroupCacheExpire = expireSeconds
	return g
}

// WithAutoCreate 启用/禁用自动创建群组
func (g *Group) WithAutoCreate(enabled bool) *Group {
	g.AutoCreateGroup = enabled
	return g
}

// WithMessageRecord 启用/禁用消息记录
func (g *Group) WithMessageRecord(enabled bool) *Group {
	g.EnableMessageRecord = enabled
	return g
}

// DefaultTicket 默认工单配置
func DefaultTicket() *Ticket {
	return &Ticket{
		Enabled:              true,
		MaxTicketsPerAgent:   10,
		AutoAssign:           true,
		AssignStrategy:       "load-balance",
		TicketTimeout:        1800,
		EnableQueueing:       true,
		QueueTimeout:         300,
		NotifyTimeout:        30,
		EnableTransfer:       true,
		TransferMaxTimes:     3,
		EnableOfflineMessage: true,
		OfflineMessageExpire: 86400,
		EnableAck:            true,
		AckTimeoutMs:         5000, // 5秒
		MaxRetry:             3,
	}
}

// ========== Ticket 自身链式调用方法 ==========

// Enable 启用工单功能
func (t *Ticket) Enable() *Ticket {
	t.Enabled = true
	return t
}

// Disable 禁用工单功能
func (t *Ticket) Disable() *Ticket {
	t.Enabled = false
	return t
}

// WithMaxPerAgent 设置每个客服最大工单数
func (t *Ticket) WithMaxPerAgent(maxTickets int) *Ticket {
	t.MaxTicketsPerAgent = maxTickets
	return t
}

// WithAutoAssign 设置是否自动分配工单
func (t *Ticket) WithAutoAssign(enabled bool, strategy string) *Ticket {
	t.AutoAssign = enabled
	t.AssignStrategy = strategy
	return t
}

// WithTimeout 设置工单超时
func (t *Ticket) WithTimeout(ticketTimeoutSeconds int) *Ticket {
	t.TicketTimeout = ticketTimeoutSeconds
	return t
}

// WithQueueing 设置排队功能
func (t *Ticket) WithQueueing(enabled bool, queueTimeoutSeconds int) *Ticket {
	t.EnableQueueing = enabled
	t.QueueTimeout = queueTimeoutSeconds
	return t
}

// WithNotifyTimeout 设置通知超时
func (t *Ticket) WithNotifyTimeout(timeoutSeconds int) *Ticket {
	t.NotifyTimeout = timeoutSeconds
	return t
}

// WithTransfer 设置工单转接功能
func (t *Ticket) WithTransfer(enabled bool, maxTimes int) *Ticket {
	t.EnableTransfer = enabled
	t.TransferMaxTimes = maxTimes
	return t
}

// WithOfflineMessage 设置离线消息功能
func (t *Ticket) WithOfflineMessage(enabled bool, expireSeconds int) *Ticket {
	t.EnableOfflineMessage = enabled
	t.OfflineMessageExpire = expireSeconds
	return t
}

// WithAck 设置ACK相关配置
func (t *Ticket) WithAck(enabled bool, timeoutMs, maxRetry int) *Ticket {
	t.EnableAck = enabled
	t.AckTimeoutMs = timeoutMs
	t.MaxRetry = maxRetry
	return t
}

// DefaultPerformance 默认性能配置
func DefaultPerformance() *Performance {
	return &Performance{
		MaxConnectionsPerNode: 10000,
		ReadBufferSize:        4,
		WriteBufferSize:       4,
		EnableCompression:     false,
		CompressionLevel:      6,
		EnableMetrics:         true,
		MetricsInterval:       60,
		EnableSlowLog:         true,
		SlowLogThreshold:      1000,
	}
}

// ========== Performance 自身链式调用方法 ==========

// WithMaxConnections 设置最大连接数
func (p *Performance) WithMaxConnections(maxConnections int) *Performance {
	p.MaxConnectionsPerNode = maxConnections
	return p
}

// WithBufferSize 设置缓冲区大小
func (p *Performance) WithBufferSize(readSize, writeSize int) *Performance {
	p.ReadBufferSize = readSize
	p.WriteBufferSize = writeSize
	return p
}

// WithCompression 设置压缩配置
func (p *Performance) WithCompression(enabled bool, level int) *Performance {
	p.EnableCompression = enabled
	p.CompressionLevel = level
	return p
}

// WithMetrics 设置性能指标配置
func (p *Performance) WithMetrics(enabled bool, intervalSeconds int) *Performance {
	p.EnableMetrics = enabled
	p.MetricsInterval = intervalSeconds
	return p
}

// WithSlowLog 设置慢日志配置
func (p *Performance) WithSlowLog(enabled bool, thresholdMs int) *Performance {
	p.EnableSlowLog = enabled
	p.SlowLogThreshold = thresholdMs
	return p
}

// DefaultSecurity 默认安全配置
func DefaultSecurity() *Security {
	return &Security{
		EnableAuth:        true,
		EnableEncryption:  false,
		EnableRateLimit:   true,
		MaxMessageSize:    1024,
		AllowedUserTypes:  []string{"customer", "agent", "admin"},
		BlockedIPs:        []string{},
		WhitelistIPs:      []string{},
		EnableIPWhitelist: false,
		TokenExpiration:   3600,
		MaxLoginAttempts:  5,
		LoginLockDuration: 300,
	}
}

// ========== Security 自身链式调用方法 ==========

// WithAuth 设置认证配置
func (s *Security) WithAuth(enabled bool) *Security {
	s.EnableAuth = enabled
	return s
}

// WithEncryption 设置加密配置
func (s *Security) WithEncryption(enabled bool) *Security {
	s.EnableEncryption = enabled
	return s
}

// WithRateLimit 设置限流配置
func (s *Security) WithRateLimit(enabled bool) *Security {
	s.EnableRateLimit = enabled
	return s
}

// WithMaxMessageSize 设置最大消息大小
func (s *Security) WithMaxMessageSize(maxSizeKB int) *Security {
	s.MaxMessageSize = maxSizeKB
	return s
}

// WithAllowedUserTypes 设置允许的用户类型
func (s *Security) WithAllowedUserTypes(userTypes []string) *Security {
	s.AllowedUserTypes = userTypes
	return s
}

// WithBlockedIPs 设置黑名单IP
func (s *Security) WithBlockedIPs(ips []string) *Security {
	s.BlockedIPs = ips
	return s
}

// WithWhitelist 设置IP白名单
func (s *Security) WithWhitelist(enabled bool, ips []string) *Security {
	s.EnableIPWhitelist = enabled
	s.WhitelistIPs = ips
	return s
}

// WithTokenExpiration 设置Token过期时间
func (s *Security) WithTokenExpiration(expireSeconds int) *Security {
	s.TokenExpiration = expireSeconds
	return s
}

// WithLoginSecurity 设置登录安全配置
func (s *Security) WithLoginSecurity(maxAttempts int, lockDurationSeconds int) *Security {
	s.MaxLoginAttempts = maxAttempts
	s.LoginLockDuration = lockDurationSeconds
	return s
}

// Get 返回配置接口
func (c *WSC) Get() interface{} {
	return c
}

// Set 设置配置数据
func (c *WSC) Set(data interface{}) {
	if cfg, ok := data.(*WSC); ok {
		*c = *cfg
	}
}

// Clone 返回配置的副本
func (c *WSC) Clone() internal.Configurable {
	origins := make([]string, len(c.WebSocketOrigins))
	copy(origins, c.WebSocketOrigins)

	cloned := &WSC{
		Enabled:           c.Enabled,
		NodeIP:            c.NodeIP,
		NodePort:          c.NodePort,
		HeartbeatInterval: c.HeartbeatInterval,
		ClientTimeout:     c.ClientTimeout,
		MessageBufferSize: c.MessageBufferSize,
		WebSocketOrigins:  origins,
		SSEHeartbeat:      c.SSEHeartbeat,
		SSETimeout:        c.SSETimeout,
		SSEMessageBuffer:  c.SSEMessageBuffer,
	}

	// 克隆嵌套配置
	if c.Distributed != nil {
		dist := *c.Distributed
		cloned.Distributed = &dist
	}

	if c.Redis != nil {
		redis := c.Redis.Clone().(*cache.Redis)
		cloned.Redis = redis
	}

	if c.Group != nil {
		group := *c.Group
		cloned.Group = &group
	}

	if c.Ticket != nil {
		ticket := *c.Ticket
		cloned.Ticket = &ticket
	}

	if c.Performance != nil {
		perf := *c.Performance
		cloned.Performance = &perf
	}

	if c.Security != nil {
		sec := *c.Security
		if len(c.Security.AllowedUserTypes) > 0 {
			sec.AllowedUserTypes = make([]string, len(c.Security.AllowedUserTypes))
			copy(sec.AllowedUserTypes, c.Security.AllowedUserTypes)
		}
		if len(c.Security.BlockedIPs) > 0 {
			sec.BlockedIPs = make([]string, len(c.Security.BlockedIPs))
			copy(sec.BlockedIPs, c.Security.BlockedIPs)
		}
		if len(c.Security.WhitelistIPs) > 0 {
			sec.WhitelistIPs = make([]string, len(c.Security.WhitelistIPs))
			copy(sec.WhitelistIPs, c.Security.WhitelistIPs)
		}
		cloned.Security = &sec
	}

	return cloned
}

// Validate 验证配置
func (c *WSC) Validate() error {
	return internal.ValidateStruct(c)
}

// Enable 启用 WebSocket 通信
func (c *WSC) Enable() *WSC {
	c.Enabled = true
	return c
}

// Disable 禁用 WebSocket 通信
func (c *WSC) Disable() *WSC {
	c.Enabled = false
	return c
}

// IsEnabled 检查是否启用
func (c *WSC) IsEnabled() bool {
	return c.Enabled
}

// WithNodeIP 设置节点IP
func (c *WSC) WithNodeIP(ip string) *WSC {
	c.NodeIP = ip
	return c
}

// WithNodePort 设置节点端口
func (c *WSC) WithNodePort(port int) *WSC {
	c.NodePort = port
	return c
}

// WithHeartbeatInterval 设置心跳间隔
func (c *WSC) WithHeartbeatInterval(interval int) *WSC {
	c.HeartbeatInterval = interval
	return c
}

// WithClientTimeout 设置客户端超时
func (c *WSC) WithClientTimeout(timeout int) *WSC {
	c.ClientTimeout = timeout
	return c
}

// WithMessageBufferSize 设置消息缓冲区大小
func (c *WSC) WithMessageBufferSize(size int) *WSC {
	c.MessageBufferSize = size
	return c
}

// WithWriteWait 设置写超时并返回当前配置对象
func (c *WSC) WithWriteWait(d time.Duration) *WSC {
	c.WriteWait = d
	return c
}

// WithMaxMessageSize 设置最大消息长度并返回当前配置对象
func (c *WSC) WithMaxMessageSize(size int64) *WSC {
	c.MaxMessageSize = size
	return c
}

// WithWebSocketOrigins 设置允许的 WebSocket Origin
func (c *WSC) WithWebSocketOrigins(origins []string) *WSC {
	c.WebSocketOrigins = origins
	return c
}

// WithMinRecTime 设置最小重连时间并返回当前配置对象
func (c *WSC) WithMinRecTime(d time.Duration) *WSC {
	c.MinRecTime = d
	return c
}

// WithMaxRecTime 设置最大重连时间并返回当前配置对象
func (c *WSC) WithMaxRecTime(d time.Duration) *WSC {
	c.MaxRecTime = d
	return c
}

// WithRecFactor 设置重连因子并返回当前配置对象
func (c *WSC) WithRecFactor(factor float64) *WSC {
	c.RecFactor = factor
	return c
}

// WithAutoReconnect 设置自动重连开关
func (c *WSC) WithAutoReconnect(enabled bool) *WSC {
	c.AutoReconnect = enabled
	return c
}

// WithSSEHeartbeat 设置 SSE 心跳间隔
func (c *WSC) WithSSEHeartbeat(interval int) *WSC {
	c.SSEHeartbeat = interval
	return c
}

// WithSSETimeout 设置 SSE 超时
func (c *WSC) WithSSETimeout(timeout int) *WSC {
	c.SSETimeout = timeout
	return c
}

// WithSSEMessageBuffer 设置 SSE 消息缓冲区大小
func (c *WSC) WithSSEMessageBuffer(size int) *WSC {
	c.SSEMessageBuffer = size
	return c
}

// WithDistributed 设置分布式配置
func (c *WSC) WithDistributed(distributed *Distributed) *WSC {
	c.Distributed = distributed
	return c
}

// EnableDistributed 启用分布式
func (c *WSC) EnableDistributed() *WSC {
	if c.Distributed == nil {
		c.Distributed = DefaultDistributed()
	}
	c.Distributed.Enabled = true
	return c
}

// WithRedis 设置Redis配置
func (c *WSC) WithRedis(redis *cache.Redis) *WSC {
	c.Redis = redis
	return c
}

// EnableRedis 启用Redis
func (c *WSC) EnableRedis() *WSC {
	if c.Redis == nil {
		c.Redis = DefaultRedis()
	}
	return c
}

// WithGroup 设置群组配置
func (c *WSC) WithGroup(group *Group) *WSC {
	c.Group = group
	return c
}

// EnableGroup 启用群组功能
func (c *WSC) EnableGroup() *WSC {
	if c.Group == nil {
		c.Group = DefaultGroup()
	}
	c.Group.Enabled = true
	return c
}

// WithTicket 设置工单配置
func (c *WSC) WithTicket(ticket *Ticket) *WSC {
	c.Ticket = ticket
	return c
}

// EnableTicket 启用工单功能
func (c *WSC) EnableTicket() *WSC {
	if c.Ticket == nil {
		c.Ticket = DefaultTicket()
	}
	c.Ticket.Enabled = true
	return c
}

// WithPerformance 设置性能配置
func (c *WSC) WithPerformance(performance *Performance) *WSC {
	c.Performance = performance
	return c
}

// WithSecurity 设置安全配置
func (c *WSC) WithSecurity(security *Security) *WSC {
	c.Security = security
	return c
}

// WithVIP 设置VIP配置
func (c *WSC) WithVIP(vip *VIP) *WSC {
	c.VIP = vip
	return c
}

// EnableVIP 启用VIP功能（使用默认配置）
func (c *WSC) EnableVIP() *WSC {
	if c.VIP == nil {
		c.VIP = DefaultVIP()
	}
	c.VIP.Enabled = true
	return c
}

// WithEnhancement 设置增强功能配置
func (c *WSC) WithEnhancement(enhancement *Enhancement) *WSC {
	c.Enhancement = enhancement
	return c
}

// EnableEnhancement 启用增强功能（使用默认配置）
func (c *WSC) EnableEnhancement() *WSC {
	if c.Enhancement == nil {
		c.Enhancement = DefaultEnhancement()
	}
	c.Enhancement.Enabled = true
	return c
}

// ========== Safe 安全访问方法 ==========

// WSCSafe WSC 配置安全访问器
type WSCSafe struct {
	*safe.SafeAccess
}

// Safe 创建 WSC 安全访问器
func (c *WSC) Safe() *WSCSafe {
	return &WSCSafe{
		SafeAccess: safe.Safe(c),
	}
}

// SafeWSC 创建 WSC 安全访问器（全局函数）
func SafeWSC(config interface{}) *WSCSafe {
	return &WSCSafe{
		SafeAccess: safe.Safe(config),
	}
}

// Distributed 安全访问 Distributed 配置
func (s *WSCSafe) Distributed() *WSCSafe {
	return &WSCSafe{SafeAccess: s.Field("Distributed")}
}

// Redis 安全访问 Redis 配置
func (s *WSCSafe) Redis() *WSCSafe {
	return &WSCSafe{SafeAccess: s.Field("Redis")}
}

// Group 安全访问 Group 配置
func (s *WSCSafe) Group() *WSCSafe {
	return &WSCSafe{SafeAccess: s.Field("Group")}
}

// Ticket 安全访问 Ticket 配置
func (s *WSCSafe) Ticket() *WSCSafe {
	return &WSCSafe{SafeAccess: s.Field("Ticket")}
}

// Performance 安全访问 Performance 配置
func (s *WSCSafe) Performance() *WSCSafe {
	return &WSCSafe{SafeAccess: s.Field("Performance")}
}

// Security 安全访问 Security 配置
func (s *WSCSafe) Security() *WSCSafe {
	return &WSCSafe{SafeAccess: s.Field("Security")}
}

// VIP 安全访问 VIP 配置
func (s *WSCSafe) VIP() *WSCSafe {
	return &WSCSafe{SafeAccess: s.Field("VIP")}
}

// Enhancement 安全访问 Enhancement 配置
func (s *WSCSafe) Enhancement() *WSCSafe {
	return &WSCSafe{SafeAccess: s.Field("Enhancement")}
}

// Enabled 安全获取 Enabled 字段
func (s *WSCSafe) Enabled(defaultValue ...bool) bool {
	return s.Field("Enabled").Bool(defaultValue...)
}

// NodeIP 安全获取 NodeIP 字段
func (s *WSCSafe) NodeIP(defaultValue ...string) string {
	return s.Field("NodeIP").String(defaultValue...)
}

// NodePort 安全获取 NodePort 字段
func (s *WSCSafe) NodePort(defaultValue ...int) int {
	return s.Field("NodePort").Int(defaultValue...)
}

// HeartbeatInterval 安全获取 HeartbeatInterval 字段
func (s *WSCSafe) HeartbeatInterval(defaultValue ...int) int {
	return s.Field("HeartbeatInterval").Int(defaultValue...)
}

// ClientTimeout 安全获取 ClientTimeout 字段
func (s *WSCSafe) ClientTimeout(defaultValue ...int) int {
	return s.Field("ClientTimeout").Int(defaultValue...)
}

// MessageBufferSize 安全获取 MessageBufferSize 字段
func (s *WSCSafe) MessageBufferSize(defaultValue ...int) int {
	return s.Field("MessageBufferSize").Int(defaultValue...)
}

// WebSocketOrigins 安全获取 WebSocketOrigins 字段
func (s *WSCSafe) WebSocketOrigins() []string {
	val := s.Field("WebSocketOrigins").Value()
	if slice, ok := val.([]string); ok {
		return slice
	}
	return []string{}
}

// SSEHeartbeat 安全获取 SSEHeartbeat 字段
func (s *WSCSafe) SSEHeartbeat(defaultValue ...int) int {
	return s.Field("SSEHeartbeat").Int(defaultValue...)
}

// SSETimeout 安全获取 SSETimeout 字段
func (s *WSCSafe) SSETimeout(defaultValue ...int) int {
	return s.Field("SSETimeout").Int(defaultValue...)
}

// SSEMessageBuffer 安全获取 SSEMessageBuffer 字段
func (s *WSCSafe) SSEMessageBuffer(defaultValue ...int) int {
	return s.Field("SSEMessageBuffer").Int(defaultValue...)
}

// MaxGroupSize 安全获取 MaxGroupSize 字段
func (s *WSCSafe) MaxGroupSize(defaultValue ...int) int {
	return s.Field("MaxGroupSize").Int(defaultValue...)
}

// MaxTicketsPerAgent 安全获取 MaxTicketsPerAgent 字段
func (s *WSCSafe) MaxTicketsPerAgent(defaultValue ...int) int {
	return s.Field("MaxTicketsPerAgent").Int(defaultValue...)
}

// AssignStrategy 安全获取 AssignStrategy 字段
func (s *WSCSafe) AssignStrategy(defaultValue ...string) string {
	return s.Field("AssignStrategy").String(defaultValue...)
}

// PubSubChannel 安全获取 PubSubChannel 字段
func (s *WSCSafe) PubSubChannel(defaultValue ...string) string {
	return s.Field("PubSubChannel").String(defaultValue...)
}

// ClusterName 安全获取 ClusterName 字段
func (s *WSCSafe) ClusterName(defaultValue ...string) string {
	return s.Field("ClusterName").String(defaultValue...)
}

// DefaultVIP 创建默认VIP配置
func DefaultVIP() *VIP {
	return &VIP{
		Enabled:               true,
		MaxLevel:              8, // V0-V8
		DefaultLevel:          0, // V0
		PriorityMultiplier:    2.0,
		MessagePriorityBonus:  10,
		QueuePriority:         true,
		CustomServicePriority: true,
		SpecialFeatures:       []string{"priority_queue", "dedicated_support"},
		UpgradeRules:          []string{"spending_based", "loyalty_based"},
	}
}

// ========== VIP 自身链式调用方法 ==========

// WithMaxLevel 设置VIP最大等级
func (v *VIP) WithMaxLevel(maxLevel int) *VIP {
	v.MaxLevel = maxLevel
	return v
}

// WithDefaultLevel 设置默认VIP等级
func (v *VIP) WithDefaultLevel(defaultLevel int) *VIP {
	v.DefaultLevel = defaultLevel
	return v
}

// WithPriorityMultiplier 设置VIP优先级倍数
func (v *VIP) WithPriorityMultiplier(multiplier float64) *VIP {
	v.PriorityMultiplier = multiplier
	return v
}

// WithMessagePriorityBonus 设置VIP消息优先级加分
func (v *VIP) WithMessagePriorityBonus(bonus int) *VIP {
	v.MessagePriorityBonus = bonus
	return v
}

// WithQueuePriority 设置VIP队列优先级
func (v *VIP) WithQueuePriority(enabled bool) *VIP {
	v.QueuePriority = enabled
	return v
}

// WithCustomServicePriority 设置VIP专属客服
func (v *VIP) WithCustomServicePriority(enabled bool) *VIP {
	v.CustomServicePriority = enabled
	return v
}

// WithSpecialFeatures 设置VIP特殊功能列表
func (v *VIP) WithSpecialFeatures(features []string) *VIP {
	v.SpecialFeatures = features
	return v
}

// WithUpgradeRules 设置VIP升级规则
func (v *VIP) WithUpgradeRules(rules []string) *VIP {
	v.UpgradeRules = rules
	return v
}

// Enable 启用VIP功能
func (v *VIP) Enable() *VIP {
	v.Enabled = true
	return v
}

// Disable 禁用VIP功能
func (v *VIP) Disable() *VIP {
	v.Enabled = false
	return v
}

// DefaultEnhancement 创建默认增强功能配置
func DefaultEnhancement() *Enhancement {
	return &Enhancement{
		Enabled:              true,
		SmartRouting:         true,
		LoadBalancing:        true,
		SmartQueue:           true,
		Monitoring:           true,
		ClusterManagement:    false, // 默认关闭，需要集群时开启
		RuleEngine:           true,
		CircuitBreaker:       true,
		MessageFiltering:     true,
		PerformanceTracking:  true,
		AdvancedMetrics:      false, // 默认关闭，性能敏感时开启
		AlertSystem:          true,
		FailureThreshold:     5,
		SuccessThreshold:     3,
		CircuitTimeout:       30,
		MaxQueueSize:         1000,
		MetricsInterval:      60,
		MaxSamples:           1000,
		LoadBalanceAlgorithm: "round-robin",
	}
}

// ========== Enhancement 自身链式调用方法 ==========

// Enable 启用增强功能
func (e *Enhancement) Enable() *Enhancement {
	e.Enabled = true
	return e
}

// Disable 禁用增强功能
func (e *Enhancement) Disable() *Enhancement {
	e.Enabled = false
	return e
}

// WithSmartRouting 启用/禁用智能路由
func (e *Enhancement) WithSmartRouting(enabled bool) *Enhancement {
	e.SmartRouting = enabled
	return e
}

// WithLoadBalancing 启用/禁用负载均衡
func (e *Enhancement) WithLoadBalancing(enabled bool, algorithm string) *Enhancement {
	e.LoadBalancing = enabled
	e.LoadBalanceAlgorithm = algorithm
	return e
}

// WithSmartQueue 启用/禁用智能队列
func (e *Enhancement) WithSmartQueue(enabled bool, maxSize int) *Enhancement {
	e.SmartQueue = enabled
	e.MaxQueueSize = maxSize
	return e
}

// WithMonitoring 启用/禁用监控
func (e *Enhancement) WithMonitoring(enabled bool) *Enhancement {
	e.Monitoring = enabled
	return e
}

// WithClusterManagement 启用/禁用集群管理
func (e *Enhancement) WithClusterManagement(enabled bool) *Enhancement {
	e.ClusterManagement = enabled
	return e
}

// WithRuleEngine 启用/禁用规则引擎
func (e *Enhancement) WithRuleEngine(enabled bool) *Enhancement {
	e.RuleEngine = enabled
	return e
}

// WithCircuitBreaker 启用/禁用熔断器
func (e *Enhancement) WithCircuitBreaker(enabled bool, failureThreshold, successThreshold, timeoutSeconds int) *Enhancement {
	e.CircuitBreaker = enabled
	e.FailureThreshold = failureThreshold
	e.SuccessThreshold = successThreshold
	e.CircuitTimeout = timeoutSeconds
	return e
}

// WithMessageFiltering 启用/禁用消息过滤
func (e *Enhancement) WithMessageFiltering(enabled bool) *Enhancement {
	e.MessageFiltering = enabled
	return e
}

// WithPerformanceTracking 启用/禁用性能追踪
func (e *Enhancement) WithPerformanceTracking(enabled bool) *Enhancement {
	e.PerformanceTracking = enabled
	return e
}

// WithAdvancedMetrics 启用/禁用高级指标
func (e *Enhancement) WithAdvancedMetrics(enabled bool) *Enhancement {
	e.AdvancedMetrics = enabled
	return e
}

// WithAlertSystem 启用/禁用警报系统
func (e *Enhancement) WithAlertSystem(enabled bool) *Enhancement {
	e.AlertSystem = enabled
	return e
}

// WithMetrics 设置指标采集间隔和最大样本数
func (e *Enhancement) WithMetrics(metricsInterval, maxSamples int) *Enhancement {
	e.MetricsInterval = metricsInterval
	e.MaxSamples = maxSamples
	return e
}
