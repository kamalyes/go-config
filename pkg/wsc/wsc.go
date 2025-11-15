/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-13 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-13 23:50:00
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
)

// WSC WebSocket 通信完整配置
type WSC struct {
	// === 基础配置 ===
	Enabled            bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                           // 是否启用
	NodeIP             string   `mapstructure:"node_ip" yaml:"node-ip" json:"node_ip"`                                           // 节点IP
	NodePort           int      `mapstructure:"node_port" yaml:"node-port" json:"node_port"`                                     // 节点端口
	HeartbeatInterval  int      `mapstructure:"heartbeat_interval" yaml:"heartbeat-interval" json:"heartbeat_interval"`          // 心跳间隔(秒)
	ClientTimeout      int      `mapstructure:"client_timeout" yaml:"client-timeout" json:"client_timeout"`                      // 客户端超时(秒)
	MessageBufferSize  int      `mapstructure:"message_buffer_size" yaml:"message-buffer-size" json:"message_buffer_size"`       // 消息缓冲区大小
	WebSocketOrigins   []string `mapstructure:"websocket_origins" yaml:"websocket-origins" json:"websocket_origins"`             // 允许的WebSocket Origin
	
	// === SSE 配置 ===
	SSEHeartbeat       int      `mapstructure:"sse_heartbeat" yaml:"sse-heartbeat" json:"sse_heartbeat"`                         // SSE心跳间隔(秒)
	SSETimeout         int      `mapstructure:"sse_timeout" yaml:"sse-timeout" json:"sse_timeout"`                               // SSE超时(秒)
	SSEMessageBuffer   int      `mapstructure:"sse_message_buffer" yaml:"sse-message-buffer" json:"sse_message_buffer"`          // SSE消息缓冲区大小
	
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
}

// Distributed 分布式节点配置
type Distributed struct {
	Enabled             bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                                       // 是否启用分布式
	NodeDiscovery       string `mapstructure:"node_discovery" yaml:"node-discovery" json:"node_discovery"`                                  // 节点发现方式: redis, etcd, consul
	NodeSyncInterval    int    `mapstructure:"node_sync_interval" yaml:"node-sync-interval" json:"node_sync_interval"`                      // 节点同步间隔(秒)
	MessageRouting      string `mapstructure:"message_routing" yaml:"message-routing" json:"message_routing"`                               // 消息路由策略: hash, random, round-robin
	EnableLoadBalance   bool   `mapstructure:"enable_load_balance" yaml:"enable-load-balance" json:"enable_load_balance"`                   // 是否启用负载均衡
	HealthCheckInterval int    `mapstructure:"health_check_interval" yaml:"health-check-interval" json:"health_check_interval"`             // 健康检查间隔(秒)
	NodeTimeout         int    `mapstructure:"node_timeout" yaml:"node-timeout" json:"node_timeout"`                                        // 节点超时(秒)
	ClusterName         string `mapstructure:"cluster_name" yaml:"cluster-name" json:"cluster_name"`                                        // 集群名称
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
	Enabled            bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                                          // 是否启用群组功能
	MaxGroupSize       int  `mapstructure:"max_group_size" yaml:"max-group-size" json:"max_group_size"`                                     // 最大群组人数
	MaxGroupsPerUser   int  `mapstructure:"max_groups_per_user" yaml:"max-groups-per-user" json:"max_groups_per_user"`                      // 每个用户最大群组数
	EnableBroadcast    bool `mapstructure:"enable_broadcast" yaml:"enable-broadcast" json:"enable_broadcast"`                               // 是否启用全局广播
	BroadcastRateLimit int  `mapstructure:"broadcast_rate_limit" yaml:"broadcast-rate-limit" json:"broadcast_rate_limit"`                   // 广播频率限制(次/分钟)
	GroupCacheExpire   int  `mapstructure:"group_cache_expire" yaml:"group-cache-expire" json:"group_cache_expire"`                         // 群组缓存过期时间(秒)
	AutoCreateGroup    bool `mapstructure:"auto_create_group" yaml:"auto-create-group" json:"auto_create_group"`                            // 是否自动创建群组
}

// Ticket 工单配置
type Ticket struct {
	Enabled              bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                                                    // 是否启用工单功能
	MaxTicketsPerAgent   int    `mapstructure:"max_tickets_per_agent" yaml:"max-tickets-per-agent" json:"max_tickets_per_agent"`                         // 每个客服最大工单数
	AutoAssign           bool   `mapstructure:"auto_assign" yaml:"auto-assign" json:"auto_assign"`                                                        // 是否自动分配工单
	AssignStrategy       string `mapstructure:"assign_strategy" yaml:"assign-strategy" json:"assign_strategy"`                                           // 分配策略: random, load-balance, skill-based
	TicketTimeout        int    `mapstructure:"ticket_timeout" yaml:"ticket-timeout" json:"ticket_timeout"`                                               // 工单超时(秒)
	EnableQueueing       bool   `mapstructure:"enable_queueing" yaml:"enable-queueing" json:"enable_queueing"`                                            // 是否启用排队
	QueueTimeout         int    `mapstructure:"queue_timeout" yaml:"queue-timeout" json:"queue_timeout"`                                                  // 排队超时(秒)
	NotifyTimeout        int    `mapstructure:"notify_timeout" yaml:"notify-timeout" json:"notify_timeout"`                                               // 通知超时(秒)
	EnableTransfer       bool   `mapstructure:"enable_transfer" yaml:"enable-transfer" json:"enable_transfer"`                                            // 是否启用工单转接
	TransferMaxTimes     int    `mapstructure:"transfer_max_times" yaml:"transfer-max-times" json:"transfer_max_times"`                                   // 最大转接次数
	EnableOfflineMessage bool   `mapstructure:"enable_offline_message" yaml:"enable-offline-message" json:"enable_offline_message"`                       // 是否启用离线消息
	OfflineMessageExpire int    `mapstructure:"offline_message_expire" yaml:"offline-message-expire" json:"offline_message_expire"`                       // 离线消息过期时间(秒)
}

// Performance 性能配置
type Performance struct {
	MaxConnectionsPerNode int  `mapstructure:"max_connections_per_node" yaml:"max-connections-per-node" json:"max_connections_per_node"` // 每个节点最大连接数
	ReadBufferSize        int  `mapstructure:"read_buffer_size" yaml:"read-buffer-size" json:"read_buffer_size"`                         // 读缓冲区大小(KB)
	WriteBufferSize       int  `mapstructure:"write_buffer_size" yaml:"write-buffer-size" json:"write_buffer_size"`                       // 写缓冲区大小(KB)
	EnableCompression     bool `mapstructure:"enable_compression" yaml:"enable-compression" json:"enable_compression"`                    // 是否启用压缩
	CompressionLevel      int  `mapstructure:"compression_level" yaml:"compression-level" json:"compression_level"`                       // 压缩级别(1-9)
	EnableMetrics         bool `mapstructure:"enable_metrics" yaml:"enable-metrics" json:"enable_metrics"`                                // 是否启用性能指标
	MetricsInterval       int  `mapstructure:"metrics_interval" yaml:"metrics-interval" json:"metrics_interval"`                          // 指标采集间隔(秒)
	EnableSlowLog         bool `mapstructure:"enable_slow_log" yaml:"enable-slow-log" json:"enable_slow_log"`                             // 是否启用慢日志
	SlowLogThreshold      int  `mapstructure:"slow_log_threshold" yaml:"slow-log-threshold" json:"slow_log_threshold"`                    // 慢日志阈值(毫秒)
}

// Security 安全配置
type Security struct {
	EnableAuth        bool     `mapstructure:"enable_auth" yaml:"enable-auth" json:"enable_auth"`                                  // 是否启用认证
	EnableEncryption  bool     `mapstructure:"enable_encryption" yaml:"enable-encryption" json:"enable_encryption"`                // 是否启用加密
	EnableRateLimit   bool     `mapstructure:"enable_rate_limit" yaml:"enable-rate-limit" json:"enable_rate_limit"`                // 是否启用限流
	MaxMessageSize    int      `mapstructure:"max_message_size" yaml:"max-message-size" json:"max_message_size"`                   // 最大消息大小(KB)
	AllowedUserTypes  []string `mapstructure:"allowed_user_types" yaml:"allowed-user-types" json:"allowed_user_types"`             // 允许的用户类型
	BlockedIPs        []string `mapstructure:"blocked_ips" yaml:"blocked-ips" json:"blocked_ips"`                                  // 黑名单IP
	WhitelistIPs      []string `mapstructure:"whitelist_ips" yaml:"whitelist-ips" json:"whitelist_ips"`                            // 白名单IP
	EnableIPWhitelist bool     `mapstructure:"enable_ip_whitelist" yaml:"enable-ip-whitelist" json:"enable_ip_whitelist"`          // 是否启用IP白名单
	TokenExpiration   int      `mapstructure:"token_expiration" yaml:"token-expiration" json:"token_expiration"`                   // Token过期时间(秒)
	MaxLoginAttempts  int      `mapstructure:"max_login_attempts" yaml:"max-login-attempts" json:"max_login_attempts"`             // 最大登录尝试次数
	LoginLockDuration int      `mapstructure:"login_lock_duration" yaml:"login-lock-duration" json:"login_lock_duration"`          // 登录锁定时长(秒)
}

// Default 创建默认 WSC 配置
func Default() *WSC {
	return &WSC{
		Enabled:            false,
		NodeIP:             "0.0.0.0",
		NodePort:           8080,
		HeartbeatInterval:  30,
		ClientTimeout:      90,
		MessageBufferSize:  256,
		WebSocketOrigins:   []string{"*"},
		SSEHeartbeat:       30,
		SSETimeout:         120,
		SSEMessageBuffer:   100,
		Distributed:        DefaultDistributed(),
		Redis:              DefaultRedis(),
		Group:              DefaultGroup(),
		Ticket:             DefaultTicket(),
		Performance:        DefaultPerformance(),
		Security:           DefaultSecurity(),
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

// DefaultGroup 默认群组配置
func DefaultGroup() *Group {
	return &Group{
		Enabled:            false,
		MaxGroupSize:       500,
		MaxGroupsPerUser:   100,
		EnableBroadcast:    true,
		BroadcastRateLimit: 10,
		GroupCacheExpire:   3600,
		AutoCreateGroup:    false,
	}
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
	}
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
		Enabled:            c.Enabled,
		NodeIP:             c.NodeIP,
		NodePort:           c.NodePort,
		HeartbeatInterval:  c.HeartbeatInterval,
		ClientTimeout:      c.ClientTimeout,
		MessageBufferSize:  c.MessageBufferSize,
		WebSocketOrigins:   origins,
		SSEHeartbeat:       c.SSEHeartbeat,
		SSETimeout:         c.SSETimeout,
		SSEMessageBuffer:   c.SSEMessageBuffer,
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

// WithWebSocketOrigins 设置允许的 WebSocket Origin
func (c *WSC) WithWebSocketOrigins(origins []string) *WSC {
	c.WebSocketOrigins = origins
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

