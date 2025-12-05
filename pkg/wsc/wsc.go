/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-13 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-05 16:11:55
 * @FilePath: \go-config\pkg\wsc\wsc.go
 * @Description: WebSocket 通信核心配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package wsc

import (
	"time"

	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-config/pkg/logging"
)

// WSC WebSocket 通信核心配置
type WSC struct {
	// === 基础配置 ===
	Enabled            bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                      // 是否启用
	Network            string        `mapstructure:"network" yaml:"network" json:"network"`                                      // 网络类型: tcp, tcp4, tcp6
	NodeIP             string        `mapstructure:"node-ip" yaml:"node-ip" json:"nodeIp"`                                       // 节点IP
	NodePort           int           `mapstructure:"node-port" yaml:"node-port" json:"nodePort"`                                 // 节点端口
	HeartbeatInterval  int           `mapstructure:"heartbeat-interval" yaml:"heartbeat-interval" json:"heartbeatInterval"`      // 心跳间隔(秒)
	ClientTimeout      int           `mapstructure:"client-timeout" yaml:"client-timeout" json:"clientTimeout"`                  // 客户端超时(秒)
	MessageBufferSize  int           `mapstructure:"message-buffer-size" yaml:"message-buffer-size" json:"messageBufferSize"`    // 消息缓冲区大小
	WebSocketOrigins   []string      `mapstructure:"websocket-origins" yaml:"websocket-origins" json:"websocketOrigins"`         // 允许的WebSocket Origin
	ReadTimeout        time.Duration `mapstructure:"read-timeout" yaml:"read-timeout" json:"readTimeout"`                        // 读取超时
	WriteTimeout       time.Duration `mapstructure:"write-timeout" yaml:"write-timeout" json:"writeTimeout"`                     // 写入超时
	IdleTimeout        time.Duration `mapstructure:"idle-timeout" yaml:"idle-timeout" json:"idleTimeout"`                        // 空闲超时
	MaxMessageSize     int64         `mapstructure:"max-message-size" yaml:"max-message-size" json:"maxMessageSize"`             // 最大消息长度
	MinRecTime         time.Duration `mapstructure:"min-rec-time" yaml:"min-rec-time" json:"minRecTime"`                         // 最小重连时间
	MaxRecTime         time.Duration `mapstructure:"max-rec-time" yaml:"max-rec-time" json:"maxRecTime"`                         // 最大重连时间
	RecFactor          float64       `mapstructure:"rec-factor" yaml:"rec-factor" json:"recFactor"`                              // 重连因子
	AutoReconnect      bool          `mapstructure:"auto-reconnect" yaml:"auto-reconnect" json:"autoReconnect"`                  // 是否自动重连
	MaxRetries         int           `mapstructure:"max-retries" yaml:"max-retries" json:"maxRetries"`                           // 最大重试次数
	BaseDelay          time.Duration `mapstructure:"base-delay" yaml:"base-delay" json:"baseDelay"`                              // 基本重试延迟
	MaxDelay           time.Duration `mapstructure:"max-delay" yaml:"max-delay" json:"maxDelay"`                                 // 最大重试延迟
	AckTimeout         time.Duration `mapstructure:"ack-timeout" yaml:"ack-timeout" json:"ackTimeout"`                           // 消息确认超时
	AckMaxRetries      int           `mapstructure:"ack-max-retries" yaml:"ack-max-retries" json:"ackMaxRetries"`                // 消息确认最大重试次数
	EnableAck          bool          `mapstructure:"enable-ack" yaml:"enable-ack" json:"enableAck"`                              // 是否启用消息确认
	BackoffFactor      float64       `mapstructure:"backoff-factor" yaml:"backoff-factor" json:"backoffFactor"`                  // 重试延迟倍数
	Jitter             bool          `mapstructure:"jitter" yaml:"jitter" json:"jitter"`                                         // 是否添加随机抖动
	RetryableErrors    []string      `mapstructure:"retryable-errors" yaml:"retryable-errors" json:"retryableErrors"`            // 可重试的错误类型
	NonRetryableErrors []string      `mapstructure:"non-retryable-errors" yaml:"non-retryable-errors" json:"nonRetryableErrors"` // 不可重试的错误类型

	// === SSE 配置 ===
	SSEHeartbeat     int `mapstructure:"sse-heartbeat" yaml:"sse-heartbeat" json:"sseHeartbeat"`               // SSE心跳间隔(秒)
	SSETimeout       int `mapstructure:"sse-timeout" yaml:"sse-timeout" json:"sseTimeout"`                     // SSE超时(秒)
	SSEMessageBuffer int `mapstructure:"sse-message-buffer" yaml:"sse-message-buffer" json:"sseMessageBuffer"` // SSE消息缓冲区大小

	// === 性能配置 ===
	Performance *Performance `mapstructure:"performance" yaml:"performance" json:"performance"` // 性能配置

	// === 安全配置 ===
	Security *Security `mapstructure:"security" yaml:"security" json:"security"` // 安全配置

	// === 数据库配置 ===
	Database *Database `mapstructure:"database" yaml:"database" json:"database"`

	// === 日志配置 ===
	Logging *logging.Logging `mapstructure:"logging" yaml:"logging" json:"logging"` // 日志配置
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

	// 消息加密配置
	MessageEncryption *MessageEncryption `mapstructure:"message-encryption" yaml:"message-encryption" json:"messageEncryption"` // 消息加密配置

	// 消息风控配置
	MessageRateLimit *MessageRateLimit `mapstructure:"message-rate-limit" yaml:"message-rate-limit" json:"messageRateLimit"` // 消息风控配置
}

// MessageEncryption 消息加密配置
type MessageEncryption struct {
	Enabled         bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                             // 是否启用消息数据加密
	Algorithm       string `mapstructure:"algorithm" yaml:"algorithm" json:"algorithm"`                       // 加密算法: AES-256-GCM, AES-128-GCM
	Key             string `mapstructure:"key" yaml:"key" json:"key"`                                         // 加密密钥 (32字节用于AES-256, 16字节用于AES-128)
	EnableKeyRotate bool   `mapstructure:"enable-key-rotate" yaml:"enable-key-rotate" json:"enableKeyRotate"` // 是否启用密钥轮换
	KeyRotateHours  int    `mapstructure:"key-rotate-hours" yaml:"key-rotate-hours" json:"keyRotateHours"`    // 密钥轮换间隔(小时)
	Compress        bool   `mapstructure:"compress" yaml:"compress" json:"compress"`                          // 加密前是否压缩数据
	EncryptPrefix   string `mapstructure:"encrypt-prefix" yaml:"encrypt-prefix" json:"encryptPrefix"`         // 加密数据前缀标识
	BackupKeys      int    `mapstructure:"backup-keys" yaml:"backup-keys" json:"backupKeys"`                  // 保留的备份密钥数量
}

// MessageRateLimit 消息风控配置
type MessageRateLimit struct {
	Enabled          bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                // 是否启用消息风控
	Window           time.Duration `mapstructure:"window" yaml:"window" json:"window"`                                   // 时间窗口
	MaxMessages      int           `mapstructure:"max-messages" yaml:"max-messages" json:"maxMessages"`                  // 窗口内最大消息数
	AlertThreshold   int           `mapstructure:"alert-threshold" yaml:"alert-threshold" json:"alertThreshold"`         // 预警阈值(百分比)
	BlockDuration    time.Duration `mapstructure:"block-duration" yaml:"block-duration" json:"blockDuration"`            // 封禁时长
	UseRedis         bool          `mapstructure:"use-redis" yaml:"use-redis" json:"useRedis"`                           // 是否使用Redis存储
	RedisKeyPrefix   string        `mapstructure:"redis-key-prefix" yaml:"redis-key-prefix" json:"redisKeyPrefix"`       // Redis键前缀
	EnableEmailAlert bool          `mapstructure:"enable-email-alert" yaml:"enable-email-alert" json:"enableEmailAlert"` // 是否启用邮件预警
	EmailAlertConfig *EmailAlert   `mapstructure:"email-alert" yaml:"email-alert" json:"emailAlert"`                     // 邮件预警配置
}

// EmailAlert 邮件预警配置
type EmailAlert struct {
	SMTPHost      string   `mapstructure:"smtp-host" yaml:"smtp-host" json:"smtpHost"`                // SMTP服务器地址
	SMTPPort      int      `mapstructure:"smtp-port" yaml:"smtp-port" json:"smtpPort"`                // SMTP端口
	Username      string   `mapstructure:"username" yaml:"username" json:"username"`                  // SMTP用户名
	Password      string   `mapstructure:"password" yaml:"password" json:"password"`                  // SMTP密码
	From          string   `mapstructure:"from" yaml:"from" json:"from"`                              // 发件人地址
	To            []string `mapstructure:"to" yaml:"to" json:"to"`                                    // 收件人列表
	EnableTLS     bool     `mapstructure:"enable-tls" yaml:"enable-tls" json:"enableTls"`             // 是否启用TLS
	SubjectAlert  string   `mapstructure:"subject-alert" yaml:"subject-alert" json:"subjectAlert"`    // 预警邮件主题
	SubjectBlock  string   `mapstructure:"subject-block" yaml:"subject-block" json:"subjectBlock"`    // 封禁邮件主题
	TemplateAlert string   `mapstructure:"template-alert" yaml:"template-alert" json:"templateAlert"` // 预警邮件HTML模板
	TemplateBlock string   `mapstructure:"template-block" yaml:"template-block" json:"templateBlock"` // 封禁邮件HTML模板
	AppName       string   `mapstructure:"app-name" yaml:"app-name" json:"appName"`                   // 应用名称
}

// Database 数据库持久化配置
type Database struct {
	Enabled       bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                     // 是否启用数据库持久化
	AutoMigrate   bool          `mapstructure:"auto-migrate" yaml:"auto-migrate" json:"autoMigrate"`       // 是否自动迁移表结构
	TablePrefix   string        `mapstructure:"table-prefix" yaml:"table-prefix" json:"tablePrefix"`       // 表前缀
	LogLevel      string        `mapstructure:"log-level" yaml:"log-level" json:"logLevel"`                // 日志级别
	SlowThreshold time.Duration `mapstructure:"slow-threshold" yaml:"slow-threshold" json:"slowThreshold"` // 慢查询阈值
}

// Default 创建默认 WSC 配置
func Default() *WSC {
	return &WSC{
		Enabled:            false,
		Network:            "tcp4",
		NodeIP:             "0.0.0.0",
		NodePort:           8080,
		HeartbeatInterval:  30,
		ClientTimeout:      90,
		MessageBufferSize:  256,
		WebSocketOrigins:   []string{"*"},
		WriteTimeout:       10 * time.Second,
		ReadTimeout:        10 * time.Second,
		IdleTimeout:        120 * time.Second,
		MaxMessageSize:     512,
		MinRecTime:         2 * time.Second,
		MaxRecTime:         60 * time.Second,
		RecFactor:          1.5,
		AutoReconnect:      true,
		MaxRetries:         3,
		BaseDelay:          100 * time.Millisecond,
		MaxDelay:           5 * time.Second,
		AckTimeout:         500 * time.Millisecond,
		EnableAck:          false,
		AckMaxRetries:      3,
		BackoffFactor:      2.0,
		Jitter:             true,
		RetryableErrors:    []string{"queue_full", "timeout", "conn_error", "channel_closed"},
		NonRetryableErrors: []string{"user_offline", "permission", "validation"},
		SSETimeout:         120,
		SSEMessageBuffer:   100,
		Database:           DefaultDatabase(),
		Performance:        DefaultPerformance(),
		Security:           DefaultSecurity(),
		Logging:            DefaultLogging(),
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
		MessageEncryption: DefaultMessageEncryption(),
		MessageRateLimit:  DefaultMessageRateLimit(),
	}
}

// DefaultMessageEncryption 默认消息加密配置
func DefaultMessageEncryption() *MessageEncryption {
	return &MessageEncryption{
		Enabled:         false,
		Algorithm:       "AES-256-GCM",
		Key:             "your-32-byte-secret-key-for-aes256", // 32字节密钥,实际使用时应该通过环境变量或安全配置提供
		EnableKeyRotate: false,
		KeyRotateHours:  24,
		Compress:        true,
		EncryptPrefix:   "ENC:",
		BackupKeys:      3,
	}
}

// DefaultMessageRateLimit 默认消息风控配置
func DefaultMessageRateLimit() *MessageRateLimit {
	return &MessageRateLimit{
		Enabled:          true,
		Window:           time.Minute,
		MaxMessages:      100,
		AlertThreshold:   80,
		BlockDuration:    5 * time.Minute,
		UseRedis:         true,
		RedisKeyPrefix:   "wsc:rate_limit:",
		EnableEmailAlert: false,
		EmailAlertConfig: DefaultEmailAlert(),
	}
}

// DefaultEmailAlert 默认邮件预警配置
func DefaultEmailAlert() *EmailAlert {
	return &EmailAlert{
		SMTPHost:      "",
		SMTPPort:      587,
		Username:      "",
		Password:      "",
		From:          "",
		To:            []string{},
		EnableTLS:     true,
		SubjectAlert:  "[WebSocket风控预警] 用户消息频率异常",
		SubjectBlock:  "[WebSocket风控封禁] 用户已被封禁",
		TemplateAlert: defaultAlertEmailTemplate(),
		TemplateBlock: defaultBlockEmailTemplate(),
		AppName:       "WebSocket消息系统",
	}
}

// defaultAlertEmailTemplate 默认预警邮件模板
func defaultAlertEmailTemplate() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background: #ff9800; color: white; padding: 20px; border-radius: 5px 5px 0 0; }
        .content { background: #f9f9f9; padding: 20px; border: 1px solid #ddd; border-top: none; }
        .info-table { width: 100%; border-collapse: collapse; margin: 15px 0; }
        .info-table td { padding: 10px; border-bottom: 1px solid #ddd; }
        .info-table td:first-child { font-weight: bold; width: 150px; }
        .footer { background: #f1f1f1; padding: 15px; text-align: center; font-size: 12px; color: #666; }
        .warning { background: #fff3cd; border-left: 4px solid #ff9800; padding: 12px; margin: 15px 0; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h2>⚠️ 消息频率预警 - {{.AppName}}</h2>
        </div>
        <div class="content">
            <div class="warning">
                <strong>用户消息发送频率达到预警阈值，请关注该用户行为</strong>
            </div>
            <table class="info-table">
                <tr><td>用户ID</td><td>{{.UserID}}</td></tr>
                <tr><td>用户类型</td><td>{{.UserType}}</td></tr>
                <tr><td>当前分钟消息数</td><td><strong style="color: #ff9800; font-size: 18px;">{{.MinuteCount}} 条</strong></td></tr>
                <tr><td>当前小时消息数</td><td><strong style="color: #ff9800; font-size: 18px;">{{.HourCount}} 条</strong></td></tr>
                <tr><td>触发时间</td><td>{{.TriggerTime}}</td></tr>
            </table>
            <p><strong>建议操作：</strong></p>
            <ul>
                <li>立即检查该用户的消息内容</li>
                <li>确认是否为恶意刷屏行为</li>
                <li>必要时联系用户或执行封禁操作</li>
            </ul>
        </div>
        <div class="footer">
            <p>此邮件由 {{.AppName}} 自动发送，请勿直接回复</p>
            <p>Generated at {{.GenerateTime}}</p>
        </div>
    </div>
</body>
</html>`
}

// defaultBlockEmailTemplate 默认封禁邮件模板
func defaultBlockEmailTemplate() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background: #f44336; color: white; padding: 20px; border-radius: 5px 5px 0 0; }
        .content { background: #f9f9f9; padding: 20px; border: 1px solid #ddd; border-top: none; }
        .info-table { width: 100%; border-collapse: collapse; margin: 15px 0; }
        .info-table td { padding: 10px; border-bottom: 1px solid #ddd; }
        .info-table td:first-child { font-weight: bold; width: 150px; }
        .footer { background: #f1f1f1; padding: 15px; text-align: center; font-size: 12px; color: #666; }
        .warning { background: #ffebee; border-left: 4px solid #f44336; padding: 12px; margin: 15px 0; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h2>🚫 用户已被封禁 - {{.AppName}}</h2>
        </div>
        <div class="content">
            <div class="warning">
                <strong>用户消息发送频率超过限制，已被临时封禁</strong>
            </div>
            <table class="info-table">
                <tr><td>用户ID</td><td>{{.UserID}}</td></tr>
                <tr><td>用户类型</td><td>{{.UserType}}</td></tr>
                <tr><td>当前分钟消息数</td><td><strong style="color: #f44336; font-size: 18px;">{{.MinuteCount}} 条</strong></td></tr>
                <tr><td>当前小时消息数</td><td><strong style="color: #f44336; font-size: 18px;">{{.HourCount}} 条</strong></td></tr>
                <tr><td>封禁时间</td><td>{{.TriggerTime}}</td></tr>
            </table>
            <p><strong>紧急处理：</strong></p>
            <ul>
                <li>立即审查该用户的所有消息内容</li>
                <li>评估是否需要永久封禁</li>
                <li>通知相关运维人员</li>
                <li>记录异常行为日志</li>
            </ul>
        </div>
        <div class="footer">
            <p>此邮件由 {{.AppName}} 自动发送，请勿直接回复</p>
            <p>Generated at {{.GenerateTime}}</p>
        </div>
    </div>
</body>
</html>`
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

// ========== 配置接口方法 ==========

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
		ReadTimeout:        c.ReadTimeout,
		WriteTimeout:       c.WriteTimeout,
		IdleTimeout:        c.IdleTimeout,
		MaxMessageSize:     c.MaxMessageSize,
		MinRecTime:         c.MinRecTime,
		MaxRecTime:         c.MaxRecTime,
		RecFactor:          c.RecFactor,
		AutoReconnect:      c.AutoReconnect,
		MaxRetries:         c.MaxRetries,
		BaseDelay:          c.BaseDelay,
		MaxDelay:           c.MaxDelay,
		AckTimeout:         c.AckTimeout,
		AckMaxRetries:      c.AckMaxRetries,
		EnableAck:          c.EnableAck,
		BackoffFactor:      c.BackoffFactor,
		Jitter:             c.Jitter,
		RetryableErrors:    append([]string{}, c.RetryableErrors...),
		NonRetryableErrors: append([]string{}, c.NonRetryableErrors...),
		SSEHeartbeat:       c.SSEHeartbeat,
		SSETimeout:         c.SSETimeout,
		SSEMessageBuffer:   c.SSEMessageBuffer,
	}

	// 克隆嵌套配置
	if c.Performance != nil {
		perf := *c.Performance
		cloned.Performance = &perf
	}

	if c.Security != nil {
		sec := *c.Security
		cloned.Security = &sec
	}

	if c.Database != nil {
		db := *c.Database
		cloned.Database = &db
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

// ========== WSC 链式调用方法 ==========

// Enable 启用 WebSocket 通信
func (c *WSC) Enable() *WSC {
	c.Enabled = true
	return c
}

// WithNetwork 设置网络类型
func (g *WSC) WithNetwork(network string) *WSC {
	g.Network = network
	return g
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

// WithWriteTimeout 设置写超时
func (c *WSC) WithWriteTimeout(timeout time.Duration) *WSC {
	c.WriteTimeout = timeout
	return c
}

// WithReadTimeout 设置读超时
func (c *WSC) WithReadTimeout(timeout time.Duration) *WSC {
	c.ReadTimeout = timeout
	return c
}

// WithIdleTimeout 设置空闲超时
func (c *WSC) WithIdleTimeout(timeout time.Duration) *WSC {
	c.IdleTimeout = timeout
	return c
}

// WithMaxMessageSize 设置最大消息长度
func (c *WSC) WithMaxMessageSize(size int64) *WSC {
	c.MaxMessageSize = size
	return c
}

// WithWebSocketOrigins 设置允许的 WebSocket Origin
func (c *WSC) WithWebSocketOrigins(origins []string) *WSC {
	c.WebSocketOrigins = origins
	return c
}

// WithMinRecTime 设置最小重连时间
func (c *WSC) WithMinRecTime(d time.Duration) *WSC {
	c.MinRecTime = d
	return c
}

// WithMaxRecTime 设置最大重连时间
func (c *WSC) WithMaxRecTime(d time.Duration) *WSC {
	c.MaxRecTime = d
	return c
}

// WithAck 设置消息确认相关配置
func (c *WSC) WithAck(d time.Duration) *WSC {
	c.EnableAck = true
	c.AckTimeout = d
	return c
}

// WithAckRetries 设置ACK最大重试次数
func (c *WSC) WithAckRetries(maxRetries int) *WSC {
	c.AckMaxRetries = maxRetries
	return c
}

// WithRecFactor 设置重连因子
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

// WithLogging 设置日志配置
func (c *WSC) WithLogging(logging *logging.Logging) *WSC {
	c.Logging = logging
	return c
}

// ========== Performance 链式调用方法 ==========

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

// ========== Security 链式调用方法 ==========

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

// ========== Database 链式调用方法 ==========

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
