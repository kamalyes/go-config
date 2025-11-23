/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-13 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-24 14:23:00
 * @FilePath: \go-config\pkg\wsc\wsc.go
 * @Description: WebSocket 通信完整配置模块（包含分布式、Redis、群组、广播等）
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package wsc

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-config/pkg/logging"
	"time"
)

// WSC WebSocket 通信完整配置
type WSC struct {
	// === 基础配置 ===
	Enabled              bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                            // 是否启用
	NodeIP               string        `mapstructure:"node-ip" yaml:"node-ip" json:"nodeIp"`                                             // 节点IP
	NodePort             int           `mapstructure:"node-port" yaml:"node-port" json:"nodePort"`                                       // 节点端口
	HeartbeatInterval    int           `mapstructure:"heartbeat-interval" yaml:"heartbeat-interval" json:"heartbeatInterval"`            // 心跳间隔(秒)
	ClientTimeout        int           `mapstructure:"client-timeout" yaml:"client-timeout" json:"clientTimeout"`                        // 客户端超时(秒)
	MessageBufferSize    int           `mapstructure:"message-buffer-size" yaml:"message-buffer-size" json:"messageBufferSize"`          // 消息缓冲区大小
	WebSocketOrigins     []string      `mapstructure:"websocket-origins" yaml:"websocket-origins" json:"websocketOrigins"`               // 允许的WebSocket Origin
	ReadTimeout          time.Duration `mapstructure:"read-timeout" yaml:"read-timeout" json:"readTimeout"`                              // 读取超时(秒)
	WriteTimeout         time.Duration `mapstructure:"write-timeout" yaml:"write-timeout" json:"writeTimeout"`                           // 写入超时(秒)
	IdleTimeout          time.Duration `mapstructure:"idle-timeout" yaml:"idle-timeout" json:"idleTimeout"`                              // 空闲超时(秒)
	MaxMessageSize       int64         `mapstructure:"max-message-size" yaml:"max-message-size" json:"maxMessageSize"`                   // 最大消息长度
	MinRecTime           time.Duration `mapstructure:"min-rec-time" yaml:"min-rec-time" json:"minRecTime"`                               // 最小重连时间
	MaxRecTime           time.Duration `mapstructure:"max-rec-time" yaml:"max-rec-time" json:"maxRecTime"`                               // 最大重连时间
	RecFactor            float64       `mapstructure:"rec-factor" yaml:"rec-factor" json:"recFactor"`                                    // 重连因子
	AutoReconnect        bool          `mapstructure:"auto-reconnect" yaml:"auto-reconnect" json:"autoReconnect"`                        // 是否自动重连
	MaxRetries           int           `mapstructure:"max-retries" yaml:"max-retries" json:"maxRetries"`                                 // 最大重试次数
	BaseDelay            time.Duration `mapstructure:"base-delay" yaml:"base-delay" json:"baseDelay"`                                    // 基本重试延迟
	MaxDelay             time.Duration `mapstructure:"max-delay" yaml:"max-delay" json:"maxDelay"`                                       // 最大重试延迟
	AckTimeoutMs         time.Duration `mapstructure:"ack-timeout-ms" yaml:"ack-timeout-ms" json:"ackTimeoutMs"`                         // 消息确认超时(毫秒)
	AckMaxRetries        int           `mapstructure:"ack-max-retries" yaml:"ack-max-retries" json:"ackMaxRetries"`                      // 消息确认最大重试次数
	EnableAck            bool          `mapstructure:"enable-ack" yaml:"enable-ack" json:"enableAck"`                                    // 是否启用消息确认
	BackoffFactor        float64       `mapstructure:"backoff-factor" yaml:"backoff-factor" json:"backoffFactor"`                        // 重试延迟倍数
	Jitter               bool          `mapstructure:"jitter" yaml:"jitter" json:"jitter"`                                               // 是否添加随机抖动
	RetryableErrors      []string      `mapstructure:"retryable-errors" yaml:"retryable-errors" json:"retryableErrors"`                  // 可重试的错误类型
	NonRetryableErrors   []string      `mapstructure:"non-retryable-errors" yaml:"non-retryable-errors" json:"nonRetryableErrors"`       // 不可重试的错误类型
	EnableCircuitBreaker bool          `mapstructure:"enable-circuit-breaker" yaml:"enable-circuit-breaker" json:"enableCircuitBreaker"` // 是否启用熔断器

	// === SSE 配置 ===
	SSEHeartbeat     int `mapstructure:"sse-heartbeat" yaml:"sse-heartbeat" json:"sseHeartbeat"`               // SSE心跳间隔(秒)
	SSETimeout       int `mapstructure:"sse-timeout" yaml:"sse-timeout" json:"sseTimeout"`                     // SSE超时(秒)
	SSEMessageBuffer int `mapstructure:"sse-message-buffer" yaml:"sse-message-buffer" json:"sseMessageBuffer"` // SSE消息缓冲区大小

	// === 分布式节点配置 ===
	Distributed *Distributed `mapstructure:"distributed" yaml:"distributed" json:"distributed"` // 分布式配置

	// === 群组/广播配置 ===
	Group *Group `mapstructure:"group" yaml:"group" json:"group"` // 群组配置

	// === 性能配置 ===
	Performance *Performance `mapstructure:"performance" yaml:"performance" json:"performance"` // 性能配置

	// === 安全配置 ===
	Security *Security `mapstructure:"security" yaml:"security" json:"security"` // 安全配置

	// === 增强功能配置 ===
	Enhancement *Enhancement `mapstructure:"enhancement" yaml:"enhancement" json:"enhancement"` // 增强功能配置

	// === 数据库配置 ===
	Database *Database `mapstructure:"database" yaml:"database" json:"database"`

	// === 消息队列配置 ===
	Queue *Queue `mapstructure:"queue" yaml:"queue" json:"queue"` // 队列配置

	// === 定时任务配置 ===
	Jobs *Jobs `mapstructure:"jobs" yaml:"jobs" json:"jobs"` // 定时任务配置

	// === 日志配置 ===
	Logging *logging.Logging `mapstructure:"logging" yaml:"logging" json:"logging"` // 日志配置
}

// Distributed 分布式节点配置
type Distributed struct {
	Enabled             bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                         // 是否启用分布式
	NodeDiscovery       string `mapstructure:"node-discovery" yaml:"node-discovery" json:"nodeDiscovery"`                     // 节点发现方式: redis, etcd, consul
	NodeSyncInterval    int    `mapstructure:"node-sync-interval" yaml:"node-sync-interval" json:"nodeSyncInterval"`          // 节点同步间隔(秒)
	MessageRouting      string `mapstructure:"message-routing" yaml:"message-routing" json:"messageRouting"`                  // 消息路由策略: hash, random, round-robin
	EnableLoadBalance   bool   `mapstructure:"enable-load-balance" yaml:"enable-load-balance" json:"enableLoadBalance"`       // 是否启用负载均衡
	HealthCheckInterval int    `mapstructure:"health-check-interval" yaml:"health-check-interval" json:"healthCheckInterval"` // 健康检查间隔(秒)
	NodeTimeout         int    `mapstructure:"node-timeout" yaml:"node-timeout" json:"nodeTimeout"`                           // 节点超时(秒)
	ClusterName         string `mapstructure:"cluster-name" yaml:"cluster-name" json:"clusterName"`                           // 集群名称
}

// Group 群组/广播配置
type Group struct {
	Enabled             bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                         // 是否启用群组功能
	MaxGroupSize        int  `mapstructure:"max-group-size" yaml:"max-group-size" json:"maxGroupSize"`                      // 最大群组人数
	MaxGroupsPerUser    int  `mapstructure:"max-groups-per-user" yaml:"max-groups-per-user" json:"maxGroupsPerUser"`        // 每个用户最大群组数
	EnableBroadcast     bool `mapstructure:"enable-broadcast" yaml:"enable-broadcast" json:"enableBroadcast"`               // 是否启用全局广播
	BroadcastRateLimit  int  `mapstructure:"broadcast-rate-limit" yaml:"broadcast-rate-limit" json:"broadcastRateLimit"`    // 广播频率限制(次/分钟)
	GroupCacheExpire    int  `mapstructure:"group-cache-expire" yaml:"group-cache-expire" json:"groupCacheExpire"`          // 群组缓存过期时间(秒)
	AutoCreateGroup     bool `mapstructure:"auto-create-group" yaml:"auto-create-group" json:"autoCreateGroup"`             // 是否自动创建群组
	EnableMessageRecord bool `mapstructure:"enable-message-record" yaml:"enable-message-record" json:"enableMessageRecord"` // 是否启用消息记录
}

// Performance 性能配置
type Performance struct {
	MaxConnectionsPerNode int  `mapstructure:"max-connections-per-node" yaml:"max-connections-per-node" json:"maxConnectionsPerNode"` // 每个节点最大连接数
	ReadBufferSize        int  `mapstructure:"read-buffer-size" yaml:"read-buffer-size" json:"readBufferSize"`                        // 读缓冲区大小(KB)
	WriteBufferSize       int  `mapstructure:"write-buffer-size" yaml:"write-buffer-size" json:"writeBufferSize"`                     // 写缓冲区大小(KB)
	EnableCompression     bool `mapstructure:"enable-compression" yaml:"enable-compression" json:"enableCompression"`                 // 是否启用压缩
	CompressionLevel      int  `mapstructure:"compression-level" yaml:"compression-level" json:"compressionLevel"`                    // 压缩级别(1-9)
	EnableMetrics         bool `mapstructure:"enable-metrics" yaml:"enable-metrics" json:"enableMetrics"`                             // 是否启用性能指标
	MetricsInterval       int  `mapstructure:"metrics-interval" yaml:"metrics-interval" json:"metricsInterval"`                       // 指标采集间隔(秒)
	EnableSlowLog         bool `mapstructure:"enable-slow-log" yaml:"enable-slow-log" json:"enableSlowLog"`                           // 是否启用慢日志
	SlowLogThreshold      int  `mapstructure:"slow-log-threshold" yaml:"slow-log-threshold" json:"slowLogThreshold"`                  // 慢日志阈值(毫秒)
}

// Security 安全配置
type Security struct {
	EnableAuth        bool     `mapstructure:"enable-auth" yaml:"enable-auth" json:"enableAuth"`                        // 是否启用认证
	EnableEncryption  bool     `mapstructure:"enable-encryption" yaml:"enable-encryption" json:"enableEncryption"`      // 是否启用加密
	EnableRateLimit   bool     `mapstructure:"enable-rate-limit" yaml:"enable-rate-limit" json:"enableRateLimit"`       // 是否启用限流
	MaxMessageSize    int      `mapstructure:"max-message-size" yaml:"max-message-size" json:"maxMessageSize"`          // 最大消息大小(KB)
	AllowedUserTypes  []string `mapstructure:"allowed-user-types" yaml:"allowed-user-types" json:"allowedUserTypes"`    // 允许的用户类型
	BlockedIPs        []string `mapstructure:"blocked-ips" yaml:"blocked-ips" json:"blockedIps"`                        // 黑名单IP
	WhitelistIPs      []string `mapstructure:"whitelist-ips" yaml:"whitelist-ips" json:"whitelistIps"`                  // 白名单IP
	EnableIPWhitelist bool     `mapstructure:"enable-ip-whitelist" yaml:"enable-ip-whitelist" json:"enableIpWhitelist"` // 是否启用IP白名单
	TokenExpiration   int      `mapstructure:"token-expiration" yaml:"token-expiration" json:"tokenExpiration"`         // Token过期时间(秒)
	MaxLoginAttempts  int      `mapstructure:"max-login-attempts" yaml:"max-login-attempts" json:"maxLoginAttempts"`    // 最大登录尝试次数
	LoginLockDuration int      `mapstructure:"login-lock-duration" yaml:"login-lock-duration" json:"loginLockDuration"` // 登录锁定时长(秒)
}

// Database 数据库持久化配置
type Database struct {
	Enabled       bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                     // 是否启用数据库持久化
	AutoMigrate   bool          `mapstructure:"auto-migrate" yaml:"auto-migrate" json:"autoMigrate"`       // 是否自动迁移表结构
	TablePrefix   string        `mapstructure:"table-prefix" yaml:"table-prefix" json:"tablePrefix"`       // 表前缀
	LogLevel      string        `mapstructure:"log-level" yaml:"log-level" json:"logLevel"`                // 日志级别
	SlowThreshold time.Duration `mapstructure:"slow-threshold" yaml:"slow-threshold" json:"slowThreshold"` // 慢查询阈值
}

// Queue 消息队列配置
type Queue struct {
	Enabled        bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                        // 是否启用队列
	Type           string        `mapstructure:"type" yaml:"type" json:"type"`                                 // 队列类型: redis, rabbitmq, kafka
	BatchSize      int           `mapstructure:"batch-size" yaml:"batch-size" json:"batchSize"`                // 批处理大小
	MaxRetries     int           `mapstructure:"max-retries" yaml:"max-retries" json:"maxRetries"`             // 最大重试次数
	RetryInterval  time.Duration `mapstructure:"retry-interval" yaml:"retry-interval" json:"retryInterval"`    // 重试间隔
	ProcessTimeout time.Duration `mapstructure:"process-timeout" yaml:"process-timeout" json:"processTimeout"` // 处理超时时间
	PrefetchCount  int           `mapstructure:"prefetch-count" yaml:"prefetch-count" json:"prefetchCount"`    // 预取数量
	DeadLetterTTL  time.Duration `mapstructure:"dead-letter-ttl" yaml:"dead-letter-ttl" json:"deadLetterTtl"`  // 死信队列TTL
	Priority       bool          `mapstructure:"priority" yaml:"priority" json:"priority"`                     // 是否支持优先级
	Persistent     bool          `mapstructure:"persistent" yaml:"persistent" json:"persistent"`               // 是否持久化
}

// Jobs 定时任务配置
type Jobs struct {
	Enabled bool                `mapstructure:"enabled" yaml:"enabled" json:"enabled"` // 是否启用定时任务
	Tasks   map[string]*JobTask `mapstructure:"tasks" yaml:"tasks" json:"tasks"`       // 任务列表
}

// JobTask 单个定时任务配置
type JobTask struct {
	Enabled       bool                   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                       // 是否启用
	Cron          string                 `mapstructure:"cron" yaml:"cron" json:"cron"`                                // Cron表达式
	Interval      time.Duration          `mapstructure:"interval" yaml:"interval" json:"interval"`                    // 执行间隔(与cron二选一)
	Timeout       time.Duration          `mapstructure:"timeout" yaml:"timeout" json:"timeout"`                       // 执行超时时间
	MaxRetries    int                    `mapstructure:"max-retries" yaml:"max-retries" json:"maxRetries"`            // 最大重试次数
	Description   string                 `mapstructure:"description" yaml:"description" json:"description"`           // 任务描述
	Concurrency   int                    `mapstructure:"concurrency" yaml:"concurrency" json:"concurrency"`           // 并发数
	SkipIfRunning bool                   `mapstructure:"skip-if-running" yaml:"skip-if-running" json:"skipIfRunning"` // 如果正在运行则跳过
	Params        map[string]interface{} `mapstructure:"params" yaml:"params" json:"params"`                          // 任务自定义参数
}

// Enhancement 增强功能配置
type Enhancement struct {
	Enabled              bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                            // 是否启用增强功能
	SmartRouting         bool   `mapstructure:"smart-routing" yaml:"smart-routing" json:"smartRouting"`                           // 启用智能路由
	LoadBalancing        bool   `mapstructure:"load-balancing" yaml:"load-balancing" json:"loadBalancing"`                        // 启用负载均衡
	SmartQueue           bool   `mapstructure:"smart-queue" yaml:"smart-queue" json:"smartQueue"`                                 // 启用智能队列
	Monitoring           bool   `mapstructure:"monitoring" yaml:"monitoring" json:"monitoring"`                                   // 启用监控
	ClusterManagement    bool   `mapstructure:"cluster-management" yaml:"cluster-management" json:"clusterManagement"`            // 启用集群管理
	RuleEngine           bool   `mapstructure:"rule-engine" yaml:"rule-engine" json:"ruleEngine"`                                 // 启用规则引擎
	CircuitBreaker       bool   `mapstructure:"circuit-breaker" yaml:"circuit-breaker" json:"circuitBreaker"`                     // 启用熔断器
	MessageFiltering     bool   `mapstructure:"message-filtering" yaml:"message-filtering" json:"messageFiltering"`               // 启用消息过滤
	PerformanceTracking  bool   `mapstructure:"performance-tracking" yaml:"performance-tracking" json:"performanceTracking"`      // 启用性能追踪
	AdvancedMetrics      bool   `mapstructure:"advanced-metrics" yaml:"advanced-metrics" json:"advancedMetrics"`                  // 启用高级指标
	AlertSystem          bool   `mapstructure:"alert-system" yaml:"alert-system" json:"alertSystem"`                              // 启用警报系统
	FailureThreshold     int    `mapstructure:"failure-threshold" yaml:"failure-threshold" json:"failureThreshold"`               // 熔断器失败阈值
	SuccessThreshold     int    `mapstructure:"success-threshold" yaml:"success-threshold" json:"successThreshold"`               // 熔断器成功阈值
	CircuitTimeout       int    `mapstructure:"circuit-timeout" yaml:"circuit-timeout" json:"circuitTimeout"`                     // 熔断器超时(秒)
	MaxQueueSize         int    `mapstructure:"max-queue-size" yaml:"max-queue-size" json:"maxQueueSize"`                         // 智能队列最大大小
	MetricsInterval      int    `mapstructure:"metrics-interval" yaml:"metrics-interval" json:"metricsInterval"`                  // 指标采集间隔(秒)
	MaxSamples           int    `mapstructure:"max-samples" yaml:"max-samples" json:"maxSamples"`                                 // 性能样本最大数量
	LoadBalanceAlgorithm string `mapstructure:"load-balance-algorithm" yaml:"load-balance-algorithm" json:"loadBalanceAlgorithm"` // 负载均衡算法: round-robin, least-connections, weighted-random, consistent-hash
}

// Default 创建默认 WSC 配置
func Default() *WSC {
	return &WSC{
		Enabled:              false,
		NodeIP:               "0.0.0.0",
		NodePort:             8080,
		HeartbeatInterval:    30,
		ClientTimeout:        90,
		MessageBufferSize:    256,
		WebSocketOrigins:     []string{"*"},
		WriteTimeout:         10 * time.Second,
		ReadTimeout:          10 * time.Second,
		IdleTimeout:          120 * time.Second,
		MaxMessageSize:       512,
		MinRecTime:           2 * time.Second,
		MaxRecTime:           60 * time.Second,
		RecFactor:            1.5,
		AutoReconnect:        true,
		MaxRetries:           3,
		BaseDelay:            100 * time.Millisecond,
		MaxDelay:             5 * time.Second,
		AckTimeoutMs:         500 * time.Millisecond,
		EnableAck:            false,
		AckMaxRetries:        3,
		BackoffFactor:        2.0,
		Jitter:               true,
		RetryableErrors:      []string{"queue_full", "timeout", "conn_error", "channel_closed"},
		NonRetryableErrors:   []string{"user_offline", "permission", "validation"},
		EnableCircuitBreaker: false,
		SSETimeout:           120,
		SSEMessageBuffer:     100,
		Database:             DefaultDatabase(),
		Distributed:          DefaultDistributed(),
		Group:                DefaultGroup(),
		Performance:          DefaultPerformance(),
		Security:             DefaultSecurity(),
		Enhancement:          DefaultEnhancement(),
		Queue:                DefaultQueue(),
		Jobs:                 DefaultJobs(),
		Logging:              DefaultLogging(),
	}
}

// DefaultDatabase 创建默认数据库配置
func DefaultDatabase() *Database {
	return &Database{
		Enabled:       true,
		AutoMigrate:   true,
		TablePrefix:   "wsc_",
		LogLevel:      "warn",
		SlowThreshold: 200 * time.Millisecond,
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
	d.Enabled = false
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
	g.Enabled = false
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

// WithPerformance 设置性能配置
func (c *WSC) WithPerformance(performance *Performance) *WSC {
	c.Performance = performance
	return c
}

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

// WithSecurity 设置安全配置
func (c *WSC) WithSecurity(security *Security) *WSC {
	c.Security = security
	return c
}

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

	if c.Group != nil {
		group := *c.Group
		cloned.Group = &group
	}

	if c.Enhancement != nil {
		enhancement := *c.Enhancement
		cloned.Enhancement = &enhancement
	}

	if c.Queue != nil {
		queue := *c.Queue
		cloned.Queue = &queue
	}

	if c.Jobs != nil {
		jobs := &Jobs{
			Enabled: c.Jobs.Enabled,
			Tasks:   make(map[string]*JobTask),
		}
		for k, v := range c.Jobs.Tasks {
			task := *v
			jobs.Tasks[k] = &task
		}
		cloned.Jobs = jobs
	}

	if c.Logging != nil {
		logging := c.Logging.Clone().(*logging.Logging)
		cloned.Logging = logging
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

// WithWriteTimeout 设置写超时并返回当前配置对象
func (c *WSC) WithWriteTimeout(timeout time.Duration) *WSC {
	c.WriteTimeout = timeout
	return c
}

// WithReadTimeout 设置读超时并返回当前配置对象
func (c *WSC) WithReadTimeout(timeout time.Duration) *WSC {
	c.ReadTimeout = timeout
	return c
}

// WithIdleTimeout 设置空闲超时并返回当前配置对象
func (c *WSC) WithIdleTimeout(timeout time.Duration) *WSC {
	c.IdleTimeout = timeout
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

// WithAck 设置消息确认相关配置并返回当前配置对象
func (c *WSC) WithAck(d time.Duration) *WSC {
	c.EnableAck = true
	c.AckTimeoutMs = d
	return c
}

// WithAckRetries 设置ACK最大重试次数并返回当前配置对象
func (c *WSC) WithAckRetries(maxRetries int) *WSC {
	c.AckMaxRetries = maxRetries
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

// WithMaxRetries 设置最大重试次数
func (c *WSC) WithMaxRetries(retries int) *WSC {
	c.MaxRetries = retries
	return c
}

// WithBaseDelay 设置基础重连延迟
func (c *WSC) WithBaseDelay(d time.Duration) *WSC {
	c.BaseDelay = d
	return c
}

// WithMaxDelay 设置最大重连延迟
func (c *WSC) WithMaxDelay(d time.Duration) *WSC {
	c.MaxDelay = d
	return c
}

// WithBackoffFactor 设置重连回退因子
func (c *WSC) WithBackoffFactor(factor float64) *WSC {
	c.BackoffFactor = factor
	return c
}

// WithJitter 设置是否启用抖动
func (c *WSC) WithJitter(enabled bool) *WSC {
	c.Jitter = enabled
	return c
}

// WithRetryableErrors 设置可重试错误列表
func (c *WSC) WithRetryableErrors(errors []string) *WSC {
	c.RetryableErrors = errors
	return c
}

// WithNonRetryableErrors 设置不可重试错误列表
func (c *WSC) WithNonRetryableErrors(errors []string) *WSC {
	c.NonRetryableErrors = errors
	return c
}

// WithEnableCircuitBreaker 设置是否启用熔断器
func (c *WSC) WithEnableCircuitBreaker(enabled bool) *WSC {
	c.EnableCircuitBreaker = enabled
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

// DefaultEnhancement 创建默认增强功能配置
func DefaultEnhancement() *Enhancement {
	return &Enhancement{
		Enabled:              false,
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

// DefaultQueue 默认队列配置
func DefaultQueue() *Queue {
	return &Queue{
		Enabled:        false, // 默认关闭，需要时启用
		Type:           "redis",
		BatchSize:      10,
		MaxRetries:     3,
		RetryInterval:  5 * time.Second,
		ProcessTimeout: 30 * time.Second,
		PrefetchCount:  1,
		DeadLetterTTL:  24 * time.Hour,
		Priority:       true,
		Persistent:     true,
	}
}

// DefaultJobs 默认定时任务配置
func DefaultJobs() *Jobs {
	return &Jobs{
		Enabled: false, // 默认关闭，需要时启用
		Tasks: map[string]*JobTask{
			"cleanup": {
				Enabled:       true,
				Interval:      5 * time.Minute,
				Timeout:       60 * time.Second,
				MaxRetries:    3,
				Description:   "清理过期连接和数据",
				Concurrency:   1,
				SkipIfRunning: true,
			},
			"offline-message": {
				Enabled:       true,
				Interval:      1 * time.Minute,
				Timeout:       30 * time.Second,
				MaxRetries:    3,
				Description:   "推送离线消息",
				Concurrency:   1,
				SkipIfRunning: true,
			},
			"queue-process": {
				Enabled:       true,
				Interval:      10 * time.Second,
				Timeout:       30 * time.Second,
				MaxRetries:    5,
				Description:   "处理队列任务",
				Concurrency:   2,
				SkipIfRunning: true,
			},
			"status-sync": {
				Enabled:       true,
				Interval:      30 * time.Second,
				Timeout:       15 * time.Second,
				MaxRetries:    3,
				Description:   "同步用户在线状态",
				Concurrency:   1,
				SkipIfRunning: true,
			},
			"stats-collect": {
				Enabled:       true,
				Cron:          "0 */5 * * * *", // 每5分钟执行
				Timeout:       60 * time.Second,
				MaxRetries:    3,
				Description:   "收集统计数据",
				Concurrency:   1,
				SkipIfRunning: true,
			},
			"message-archive": {
				Enabled:       true,
				Cron:          "0 0 2 * * *", // 每天凌晨2点执行
				Timeout:       30 * time.Minute,
				MaxRetries:    3,
				Description:   "归档旧消息（3天前归档，60天后删除）",
				Concurrency:   1,
				SkipIfRunning: true,
			},
		},
	}
}

// DefaultLogging 创建默认日志配置
func DefaultLogging() *logging.Logging {
	return logging.Default().
		WithModuleName("wsc").
		WithLevel("info").
		WithFormat("json").
		WithOutput("stdout").
		WithFilePath("/var/log/wsc.log").
		WithEnabled(true)
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

// WithQueue 设置队列配置
func (c *WSC) WithQueue(queue *Queue) *WSC {
	c.Queue = queue
	return c
}

// EnableQueue 启用消息队列
func (c *WSC) EnableQueue() *WSC {
	if c.Queue == nil {
		c.Queue = DefaultQueue()
	}
	c.Queue.Enabled = true
	return c
}

// WithJobs 设置定时任务配置
func (c *WSC) WithJobs(jobs *Jobs) *WSC {
	c.Jobs = jobs
	return c
}

// EnableJobs 启用定时任务
func (c *WSC) EnableJobs() *WSC {
	if c.Jobs == nil {
		c.Jobs = DefaultJobs()
	}
	c.Jobs.Enabled = true
	return c
}

// WithLogging 设置日志配置
func (c *WSC) WithLogging(logging *logging.Logging) *WSC {
	c.Logging = logging
	return c
}

// ========== Database 自身链式调用方法 ==========

// Enable 启用数据库
func (d *Database) Enable() *Database {
	d.Enabled = true
	return d
}

// WithMigration 设置自动迁移和表前缀
func (d *Database) WithMigration(autoMigrate bool, tablePrefix string) *Database {
	d.AutoMigrate = autoMigrate
	d.TablePrefix = tablePrefix
	return d
}

// WithLogging 设置数据库日志
func (d *Database) WithLogging(logLevel string, slowThreshold time.Duration) *Database {
	d.LogLevel = logLevel
	d.SlowThreshold = slowThreshold
	return d
}

// ========== Queue 自身链式调用方法 ==========

// Enable 启用队列
func (q *Queue) Enable() *Queue {
	q.Enabled = true
	return q
}

// Disable 禁用队列
func (q *Queue) Disable() *Queue {
	q.Enabled = false
	return q
}

// WithType 设置队列类型
func (q *Queue) WithType(queueType string) *Queue {
	q.Type = queueType
	return q
}

// WithBatch 设置批处理参数
func (q *Queue) WithBatch(batchSize int) *Queue {
	q.BatchSize = batchSize
	return q
}

// WithRetry 设置重试参数
func (q *Queue) WithRetry(maxRetries int, interval time.Duration) *Queue {
	q.MaxRetries = maxRetries
	q.RetryInterval = interval
	return q
}

// WithTimeout 设置超时时间
func (q *Queue) WithTimeout(timeout time.Duration) *Queue {
	q.ProcessTimeout = timeout
	return q
}

// WithPrefetch 设置预取数量
func (q *Queue) WithPrefetch(count int) *Queue {
	q.PrefetchCount = count
	return q
}

// WithDeadLetter 设置死信队列TTL
func (q *Queue) WithDeadLetter(ttl time.Duration) *Queue {
	q.DeadLetterTTL = ttl
	return q
}

// WithPriority 启用/禁用优先级支持
func (q *Queue) WithPriority(enabled bool) *Queue {
	q.Priority = enabled
	return q
}

// WithPersistent 设置是否持久化
func (q *Queue) WithPersistent(persistent bool) *Queue {
	q.Persistent = persistent
	return q
}

// ========== Jobs 自身链式调用方法 ==========

// Enable 启用定时任务
func (j *Jobs) Enable() *Jobs {
	j.Enabled = true
	return j
}

// Disable 禁用定时任务
func (j *Jobs) Disable() *Jobs {
	j.Enabled = false
	return j
}

// WithTask 添加任务配置
func (j *Jobs) WithTask(name string, task *JobTask) *Jobs {
	if j.Tasks == nil {
		j.Tasks = make(map[string]*JobTask)
	}
	j.Tasks[name] = task
	return j
}

// EnableTask 启用指定任务
func (j *Jobs) EnableTask(name string) *Jobs {
	if j.Tasks != nil && j.Tasks[name] != nil {
		j.Tasks[name].Enabled = true
	}
	return j
}

// DisableTask 禁用指定任务
func (j *Jobs) DisableTask(name string) *Jobs {
	if j.Tasks != nil && j.Tasks[name] != nil {
		j.Tasks[name].Enabled = false
	}
	return j
}

// ========== JobTask 自身链式调用方法 ==========

// Enable 启用任务
func (t *JobTask) Enable() *JobTask {
	t.Enabled = true
	return t
}

// Disable 禁用任务
func (t *JobTask) Disable() *JobTask {
	t.Enabled = false
	return t
}

// WithCron 设置Cron表达式
func (t *JobTask) WithCron(cron string) *JobTask {
	t.Cron = cron
	t.Interval = 0 // 清空interval，使用cron
	return t
}

// WithInterval 设置执行间隔
func (t *JobTask) WithInterval(interval time.Duration) *JobTask {
	t.Interval = interval
	t.Cron = "" // 清空cron，使用interval
	return t
}

// WithTimeout 设置超时时间
func (t *JobTask) WithTimeout(timeout time.Duration) *JobTask {
	t.Timeout = timeout
	return t
}

// WithRetry 设置重试参数
func (t *JobTask) WithRetry(maxRetries int) *JobTask {
	t.MaxRetries = maxRetries
	return t
}

// WithDescription 设置任务描述
func (t *JobTask) WithDescription(description string) *JobTask {
	t.Description = description
	return t
}

// WithConcurrency 设置并发数
func (t *JobTask) WithConcurrency(concurrency int) *JobTask {
	t.Concurrency = concurrency
	return t
}

// WithSkipIfRunning 设置是否跳过运行中的任务
func (t *JobTask) WithSkipIfRunning(skip bool) *JobTask {
	t.SkipIfRunning = skip
	return t
}
