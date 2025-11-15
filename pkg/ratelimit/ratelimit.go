/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-14 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-14 00:00:00
 * @FilePath: \go-config\pkg\ratelimit\ratelimit.go
 * @Description: 增强的限流配置（支持路由级别、用户级别、IP级别等）
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package ratelimit

import (
	"time"

	"github.com/kamalyes/go-config/internal"
)

// Strategy 限流策略类型
type Strategy string

const (
	StrategyTokenBucket   Strategy = "token-bucket"   // 令牌桶算法
	StrategyLeakyBucket   Strategy = "leaky-bucket"   // 漏桶算法
	StrategySlidingWindow Strategy = "sliding-window" // 滑动窗口
	StrategyFixedWindow   Strategy = "fixed-window"   // 固定窗口
)

// Scope 限流作用域
type Scope string

const (
	ScopeGlobal   Scope = "global"    // 全局限流
	ScopePerIP    Scope = "per-ip"    // 按IP限流
	ScopePerUser  Scope = "per-user"  // 按用户限流
	ScopePerRoute Scope = "per-route" // 按路由限流
)

// RateLimit 增强限流配置
type RateLimit struct {
	ModuleName        string        `mapstructure:"module_name" yaml:"module-name" json:"module_name"`                      // 模块名称
	Enabled           bool          `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                  // 是否启用
	Strategy          Strategy      `mapstructure:"strategy" yaml:"strategy" json:"strategy"`                               // 限流策略
	DefaultScope      Scope         `mapstructure:"default_scope" yaml:"default-scope" json:"default_scope"`                // 默认作用域
	GlobalLimit       *LimitRule    `mapstructure:"global_limit" yaml:"global-limit,omitempty" json:"global_limit,omitempty"` // 全局限流规则
	Routes            []RouteLimit  `mapstructure:"routes" yaml:"routes,omitempty" json:"routes,omitempty"`                 // 路由级别限流
	IPRules           []IPRule      `mapstructure:"ip_rules" yaml:"ip-rules,omitempty" json:"ip_rules,omitempty"`           // IP规则
	UserRules         []UserRule    `mapstructure:"user_rules" yaml:"user-rules,omitempty" json:"user_rules,omitempty"`     // 用户规则
	Storage           StorageConfig `mapstructure:"storage" yaml:"storage" json:"storage"`                                  // 存储配置
	CustomRuleLoader  string        `mapstructure:"custom_rule_loader" yaml:"custom-rule-loader" json:"custom_rule_loader"` // 自定义规则加载器名称
	EnableDynamicRule bool          `mapstructure:"enable_dynamic_rule" yaml:"enable-dynamic-rule" json:"enable_dynamic_rule"` // 是否启用动态规则
}

// LimitRule 限流规则
type LimitRule struct {
	RequestsPerSecond int           `mapstructure:"requests_per_second" yaml:"requests-per-second" json:"requests_per_second"` // 每秒请求数
	BurstSize         int           `mapstructure:"burst_size" yaml:"burst-size" json:"burst_size"`                            // 突发大小
	WindowSize        time.Duration `mapstructure:"window_size" yaml:"window-size" json:"window_size"`                         // 时间窗口
	BlockDuration     time.Duration `mapstructure:"block_duration" yaml:"block-duration" json:"block_duration"`                // 阻塞时长
}

// RouteLimit 路由级别限流
type RouteLimit struct {
	Path      string     `mapstructure:"path" yaml:"path" json:"path"`                                        // 路由路径（支持通配符）
	Methods   []string   `mapstructure:"methods" yaml:"methods,omitempty" json:"methods,omitempty"`           // HTTP方法
	Limit     *LimitRule `mapstructure:"limit" yaml:"limit" json:"limit"`                                     // 限流规则
	PerUser   bool       `mapstructure:"per_user" yaml:"per-user" json:"per_user"`                            // 是否按用户限流
	PerIP     bool       `mapstructure:"per_ip" yaml:"per-ip" json:"per_ip"`                                  // 是否按IP限流
	Whitelist []string   `mapstructure:"whitelist" yaml:"whitelist,omitempty" json:"whitelist,omitempty"`     // 白名单
	Blacklist []string   `mapstructure:"blacklist" yaml:"blacklist,omitempty" json:"blacklist,omitempty"`     // 黑名单
}

// IPRule IP限流规则
type IPRule struct {
	IP       string     `mapstructure:"ip" yaml:"ip" json:"ip"`                               // IP地址（支持CIDR）
	Limit    *LimitRule `mapstructure:"limit" yaml:"limit" json:"limit"`                      // 限流规则
	Type     string     `mapstructure:"type" yaml:"type" json:"type"`                         // 类型: whitelist, blacklist, custom
	Priority int        `mapstructure:"priority" yaml:"priority" json:"priority"`             // 优先级
}

// UserRule 用户限流规则
type UserRule struct {
	UserID   string     `mapstructure:"user_id" yaml:"user-id" json:"user_id"`       // 用户ID（支持通配符）
	UserType string     `mapstructure:"user_type" yaml:"user-type" json:"user_type"` // 用户类型
	Role     string     `mapstructure:"role" yaml:"role" json:"role"`                // 用户角色
	Limit    *LimitRule `mapstructure:"limit" yaml:"limit" json:"limit"`             // 限流规则
	Priority int        `mapstructure:"priority" yaml:"priority" json:"priority"`    // 优先级
}

// StorageConfig 存储配置
type StorageConfig struct {
	Type          string        `mapstructure:"type" yaml:"type" json:"type"`                                       // 存储类型: memory, redis
	KeyPrefix     string        `mapstructure:"key_prefix" yaml:"key-prefix" json:"key_prefix"`                     // Key前缀
	CleanInterval time.Duration `mapstructure:"clean_interval" yaml:"clean-interval" json:"clean_interval"`         // 清理间隔
	RedisConfig   *RedisStorage `mapstructure:"redis" yaml:"redis,omitempty" json:"redis,omitempty"`                // Redis配置（可选）
}

// RedisStorage Redis存储配置（复用cache.Redis字段）
type RedisStorage struct {
	Addresses    []string      `mapstructure:"addresses" yaml:"addresses" json:"addresses"`                         // Redis地址
	Password     string        `mapstructure:"password" yaml:"password" json:"password"`                            // 密码
	DB           int           `mapstructure:"db" yaml:"db" json:"db"`                                              // 数据库
	PoolSize     int           `mapstructure:"pool_size" yaml:"pool-size" json:"pool_size"`                         // 连接池
	MinIdleConns int           `mapstructure:"min_idle_conns" yaml:"min-idle-conns" json:"min_idle_conns"`          // 最小空闲连接
	MaxRetries   int           `mapstructure:"max_retries" yaml:"max-retries" json:"max_retries"`                   // 最大重试
	ReadTimeout  time.Duration `mapstructure:"read_timeout" yaml:"read-timeout" json:"read_timeout"`                // 读超时
	WriteTimeout time.Duration `mapstructure:"write_timeout" yaml:"write-timeout" json:"write_timeout"`             // 写超时
	ClusterMode  bool          `mapstructure:"cluster_mode" yaml:"cluster-mode" json:"cluster_mode"`                // 集群模式
}

// Default 默认限流配置
func Default() *RateLimit {
	return &RateLimit{
		ModuleName:   "ratelimit",
		Enabled:      true,
		Strategy:     StrategyTokenBucket,
		DefaultScope: ScopeGlobal,
		GlobalLimit: &LimitRule{
			RequestsPerSecond: 100,
			BurstSize:         200,
			WindowSize:        time.Minute,
			BlockDuration:     time.Minute,
		},
		Storage: StorageConfig{
			Type:          "memory",
			KeyPrefix:     "rate_limit:",
			CleanInterval: 5 * time.Minute,
		},
		Routes:            []RouteLimit{},
		IPRules:           []IPRule{},
		UserRules:         []UserRule{},
		EnableDynamicRule: false,
	}
}

// Clone 返回配置副本
func (r *RateLimit) Clone() internal.Configurable {
	cloned := &RateLimit{
		ModuleName:        r.ModuleName,
		Enabled:           r.Enabled,
		Strategy:          r.Strategy,
		DefaultScope:      r.DefaultScope,
		Storage:           r.Storage,
		CustomRuleLoader:  r.CustomRuleLoader,
		EnableDynamicRule: r.EnableDynamicRule,
	}

	if r.GlobalLimit != nil {
		limit := *r.GlobalLimit
		cloned.GlobalLimit = &limit
	}

	if len(r.Routes) > 0 {
		cloned.Routes = make([]RouteLimit, len(r.Routes))
		copy(cloned.Routes, r.Routes)
	}

	if len(r.IPRules) > 0 {
		cloned.IPRules = make([]IPRule, len(r.IPRules))
		copy(cloned.IPRules, r.IPRules)
	}

	if len(r.UserRules) > 0 {
		cloned.UserRules = make([]UserRule, len(r.UserRules))
		copy(cloned.UserRules, r.UserRules)
	}

	return cloned
}

// Get 返回配置
func (r *RateLimit) Get() interface{} {
	return r
}

// Set 设置配置
func (r *RateLimit) Set(data interface{}) {
	if cfg, ok := data.(*RateLimit); ok {
		*r = *cfg
	}
}

// Validate 验证配置
func (r *RateLimit) Validate() error {
	return internal.ValidateStruct(r)
}

// WithGlobalLimit 设置全局限流
func (r *RateLimit) WithGlobalLimit(limit *LimitRule) *RateLimit {
	r.GlobalLimit = limit
	return r
}

// AddRouteLimit 添加路由限流
func (r *RateLimit) AddRouteLimit(route RouteLimit) *RateLimit {
	r.Routes = append(r.Routes, route)
	return r
}

// AddIPRule 添加IP规则
func (r *RateLimit) AddIPRule(rule IPRule) *RateLimit {
	r.IPRules = append(r.IPRules, rule)
	return r
}

// AddUserRule 添加用户规则
func (r *RateLimit) AddUserRule(rule UserRule) *RateLimit {
	r.UserRules = append(r.UserRules, rule)
	return r
}

// WithStrategy 设置限流策略
func (r *RateLimit) WithStrategy(strategy Strategy) *RateLimit {
	r.Strategy = strategy
	return r
}

// WithStorage 设置存储配置
func (r *RateLimit) WithStorage(storage StorageConfig) *RateLimit {
	r.Storage = storage
	return r
}

// EnableDynamic 启用动态规则
func (r *RateLimit) EnableDynamic() *RateLimit {
	r.EnableDynamicRule = true
	return r
}

// WithCustomLoader 设置自定义加载器
func (r *RateLimit) WithCustomLoader(loaderName string) *RateLimit {
	r.CustomRuleLoader = loaderName
	return r
}

// Enable 启用限流中间件
func (r *RateLimit) Enable() *RateLimit {
	r.Enabled = true
	return r
}

// Disable 禁用限流中间件
func (r *RateLimit) Disable() *RateLimit {
	r.Enabled = false
	return r
}

// IsEnabled 检查是否启用
func (r *RateLimit) IsEnabled() bool {
	return r.Enabled
}