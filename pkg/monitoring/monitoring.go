/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 01:27:26
 * @FilePath: \go-config\pkg\monitoring\monitoring.go
 * @Description: 监控配置模块 - 统一管理所有监控相关功能
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package monitoring

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-config/pkg/grafana"
	"github.com/kamalyes/go-config/pkg/jaeger"
	"github.com/kamalyes/go-config/pkg/prometheus"
)

// Monitoring 监控配置 - 去掉Config后缀
type Monitoring struct {
	ModuleName string                 `mapstructure:"module_name" yaml:"module_name" json:"module_name"` // 模块名称
	Enabled    bool                   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`             // 是否启用监控
	Prometheus *prometheus.Prometheus `mapstructure:"prometheus" yaml:"prometheus" json:"prometheus"`    // Prometheus配置
	Grafana    *grafana.Grafana       `mapstructure:"grafana" yaml:"grafana" json:"grafana"`             // Grafana配置
	Jaeger     *jaeger.Jaeger         `mapstructure:"jaeger" yaml:"jaeger" json:"jaeger"`                // Jaeger配置
	Metrics    *Metrics               `mapstructure:"metrics" yaml:"metrics" json:"metrics"`             // 指标配置
	Alerting   *Alerting              `mapstructure:"alerting" yaml:"alerting" json:"alerting"`          // 告警配置
}

// Metrics 指标配置
type Metrics struct {
	Enabled           bool           `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                     // 是否启用指标
	RequestCount      bool           `mapstructure:"request_count" yaml:"request_count" json:"request_count"`                   // 是否记录请求数
	Duration          bool           `mapstructure:"duration" yaml:"duration" json:"duration"`                                  // 是否记录请求时长
	ResponseSize      bool           `mapstructure:"response_size" yaml:"response_size" json:"response_size"`                   // 是否记录响应大小
	RequestSize       bool           `mapstructure:"request_size" yaml:"request_size" json:"request_size"`                      // 是否记录请求大小
	Buckets           []float64      `mapstructure:"buckets" yaml:"buckets" json:"buckets"`                                     // 直方图桶配置
	EnableOpenMetrics bool           `mapstructure:"enable_open_metrics" yaml:"enable_open_metrics" json:"enable_open_metrics"` // 是否启用 OpenMetrics 格式
	CustomMetrics     []CustomMetric `mapstructure:"custom_metrics" yaml:"custom_metrics" json:"custom_metrics"`                // 自定义指标
	Endpoint          string         `mapstructure:"-" yaml:"_" json:"endpoint"`                                                // 指标端点（自动计算）
}

// CustomMetric 自定义指标配置
type CustomMetric struct {
	Name       string    `mapstructure:"name" yaml:"name" json:"name"`                   // 指标名称
	Type       string    `mapstructure:"type" yaml:"type" json:"type"`                   // 指标类型 (counter, gauge, histogram, summary)
	Help       string    `mapstructure:"help" yaml:"help" json:"help"`                   // 帮助信息
	Labels     []string  `mapstructure:"labels" yaml:"labels" json:"labels"`             // 标签列表
	Buckets    []float64 `mapstructure:"buckets" yaml:"buckets" json:"buckets"`          // 直方图桶 (仅histogram类型)
	Objectives string    `mapstructure:"objectives" yaml:"objectives" json:"objectives"` // 分位数目标 (仅summary类型)
}

// Alerting 告警配置
type Alerting struct {
	Enabled  bool         `mapstructure:"enabled" yaml:"enabled" json:"enabled"`    // 是否启用告警
	Rules    []AlertRule  `mapstructure:"rules" yaml:"rules" json:"rules"`          // 告警规则
	Webhooks []Webhook    `mapstructure:"webhooks" yaml:"webhooks" json:"webhooks"` // Webhook配置
	Email    *EmailConfig `mapstructure:"email" yaml:"email" json:"email"`          // 邮件配置
	Slack    *SlackConfig `mapstructure:"slack" yaml:"slack" json:"slack"`          // Slack配置
}

// AlertRule 告警规则
type AlertRule struct {
	Name        string            `mapstructure:"name" yaml:"name" json:"name"`                      // 规则名称
	Query       string            `mapstructure:"query" yaml:"query" json:"query"`                   // 查询表达式
	Condition   string            `mapstructure:"condition" yaml:"condition" json:"condition"`       // 条件 (>, <, ==, !=)
	Threshold   float64           `mapstructure:"threshold" yaml:"threshold" json:"threshold"`       // 阈值
	Duration    string            `mapstructure:"duration" yaml:"duration" json:"duration"`          // 持续时间
	Severity    string            `mapstructure:"severity" yaml:"severity" json:"severity"`          // 严重级别 (critical, warning, info)
	Labels      map[string]string `mapstructure:"labels" yaml:"labels" json:"labels"`                // 标签
	Annotations map[string]string `mapstructure:"annotations" yaml:"annotations" json:"annotations"` // 注释
}

// Webhook Webhook配置
type Webhook struct {
	Name    string            `mapstructure:"name" yaml:"name" json:"name"`          // Webhook名称
	URL     string            `mapstructure:"url" yaml:"url" json:"url"`             // Webhook URL
	Headers map[string]string `mapstructure:"headers" yaml:"headers" json:"headers"` // 请求头
	Timeout string            `mapstructure:"timeout" yaml:"timeout" json:"timeout"` // 超时时间
}

// EmailConfig 邮件配置
type EmailConfig struct {
	Enabled  bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`       // 是否启用邮件
	SMTPHost string   `mapstructure:"smtp_host" yaml:"smtp_host" json:"smtp_host"` // SMTP主机
	SMTPPort int      `mapstructure:"smtp_port" yaml:"smtp_port" json:"smtp_port"` // SMTP端口
	Username string   `mapstructure:"username" yaml:"username" json:"username"`    // 用户名
	Password string   `mapstructure:"password" yaml:"password" json:"password"`    // 密码
	From     string   `mapstructure:"from" yaml:"from" json:"from"`                // 发送者
	To       []string `mapstructure:"to" yaml:"to" json:"to"`                      // 接收者列表
	TLS      bool     `mapstructure:"tls" yaml:"tls" json:"tls"`                   // 是否使用TLS
}

// SlackConfig Slack配置
type SlackConfig struct {
	Enabled   bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`          // 是否启用Slack
	Token     string `mapstructure:"token" yaml:"token" json:"token"`                // Bot Token
	Channel   string `mapstructure:"channel" yaml:"channel" json:"channel"`          // 默认频道
	Username  string `mapstructure:"username" yaml:"username" json:"username"`       // Bot用户名
	IconEmoji string `mapstructure:"icon_emoji" yaml:"icon_emoji" json:"icon_emoji"` // 图标表情
}

// Default 创建默认监控配置
func Default() *Monitoring {
	m := &Monitoring{
		ModuleName: "monitoring",
		Enabled:    false,
		Prometheus: prometheus.Default(),
		Grafana:    grafana.Default(),
		Jaeger:     jaeger.Default(),
		Metrics: &Metrics{
			Enabled:           true,
			RequestCount:      true,
			Duration:          true,
			ResponseSize:      true,
			RequestSize:       true,
			Buckets:           []float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120},
			EnableOpenMetrics: false,
			CustomMetrics:     []CustomMetric{},
		},
		Alerting: &Alerting{
			Enabled:  false,
			Rules:    []AlertRule{},
			Webhooks: []Webhook{},
			Email: &EmailConfig{
				Enabled: false,
				To:      []string{},
			},
			Slack: &SlackConfig{
				Enabled: false,
			},
		},
	}
	internal.CallAfterLoad(m) // 自动计算 Endpoint
	return m
}

// BeforeLoad 配置加载前的钩子
func (m *Monitoring) BeforeLoad() error {
	return nil
}

// AfterLoad 配置加载后的钩子 - 计算 Metrics Endpoint
func (m *Monitoring) AfterLoad() error {
	// 如果 Prometheus 配置存在且有 Path，使用它来计算 Metrics Endpoint
	if m.Prometheus != nil && m.Prometheus.Path != "" {
		m.Metrics.Endpoint = m.Prometheus.Path
	} else if m.Metrics != nil {
		// 否则使用默认的 /metrics
		m.Metrics.Endpoint = "/metrics"
	}
	return nil
}

// Get 返回配置接口
func (m *Monitoring) Get() interface{} {
	return m
}

// Set 设置配置数据
func (m *Monitoring) Set(data interface{}) {
	if cfg, ok := data.(*Monitoring); ok {
		*m = *cfg
	}
}

// Clone 返回配置的副本
func (m *Monitoring) Clone() internal.Configurable {
	metrics := &Metrics{}
	alerting := &Alerting{}

	if m.Metrics != nil {
		*metrics = *m.Metrics
		metrics.Buckets = append([]float64(nil), m.Metrics.Buckets...)
		metrics.CustomMetrics = make([]CustomMetric, len(m.Metrics.CustomMetrics))
		for i, cm := range m.Metrics.CustomMetrics {
			metrics.CustomMetrics[i] = CustomMetric{
				Name:       cm.Name,
				Type:       cm.Type,
				Help:       cm.Help,
				Labels:     append([]string(nil), cm.Labels...),
				Buckets:    append([]float64(nil), cm.Buckets...),
				Objectives: cm.Objectives,
			}
		}
	}

	if m.Alerting != nil {
		*alerting = *m.Alerting
		alerting.Rules = make([]AlertRule, len(m.Alerting.Rules))
		copy(alerting.Rules, m.Alerting.Rules)
		alerting.Webhooks = make([]Webhook, len(m.Alerting.Webhooks))
		copy(alerting.Webhooks, m.Alerting.Webhooks)
		if m.Alerting.Email != nil {
			email := *m.Alerting.Email
			email.To = append([]string(nil), m.Alerting.Email.To...)
			alerting.Email = &email
		}
		if m.Alerting.Slack != nil {
			slack := *m.Alerting.Slack
			alerting.Slack = &slack
		}
	}

	cloned := &Monitoring{
		ModuleName: m.ModuleName,
		Enabled:    m.Enabled,
		Prometheus: m.Prometheus.Clone().(*prometheus.Prometheus),
		Grafana:    m.Grafana.Clone().(*grafana.Grafana),
		Jaeger:     m.Jaeger.Clone().(*jaeger.Jaeger),
		Metrics:    metrics,
		Alerting:   alerting,
	}
	internal.CallAfterLoad(cloned) // 重新计算 Endpoint
	return cloned
}

// Validate 验证配置
func (m *Monitoring) Validate() error {
	if err := internal.ValidateStruct(m); err != nil {
		return err
	}

	// 验证子模块
	if m.Prometheus != nil {
		if err := m.Prometheus.Validate(); err != nil {
			return err
		}
	}
	if m.Grafana != nil {
		if err := m.Grafana.Validate(); err != nil {
			return err
		}
	}
	if m.Jaeger != nil {
		if err := m.Jaeger.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// WithModuleName 设置模块名称
func (m *Monitoring) WithModuleName(moduleName string) *Monitoring {
	m.ModuleName = moduleName
	return m
}

// WithEnabled 设置是否启用监控
func (m *Monitoring) WithEnabled(enabled bool) *Monitoring {
	m.Enabled = enabled
	return m
}

// WithPrometheus 设置 Prometheus 配置
func (m *Monitoring) WithPrometheus(p *prometheus.Prometheus) *Monitoring {
	m.Prometheus = p
	internal.CallAfterLoad(m) // 重新计算 Endpoint
	return m
}

// WithGrafana 设置Grafana配置
func (m *Monitoring) WithGrafana(cfg *grafana.Grafana) *Monitoring {
	m.Grafana = cfg
	return m
}

// WithJaeger 设置Jaeger配置
func (m *Monitoring) WithJaeger(cfg *jaeger.Jaeger) *Monitoring {
	m.Jaeger = cfg
	return m
}

// WithMetrics 设置指标配置
func (m *Monitoring) WithMetrics(enabled, requestCount, duration, responseSize, requestSize bool) *Monitoring {
	if m.Metrics == nil {
		m.Metrics = &Metrics{}
	}
	m.Metrics.Enabled = enabled
	m.Metrics.RequestCount = requestCount
	m.Metrics.Duration = duration
	m.Metrics.ResponseSize = responseSize
	m.Metrics.RequestSize = requestSize
	return m
}

// AddCustomMetric 添加自定义指标
func (m *Monitoring) AddCustomMetric(name, metricType, help string, labels []string) *Monitoring {
	if m.Metrics == nil {
		m.Metrics = &Metrics{}
	}
	m.Metrics.CustomMetrics = append(m.Metrics.CustomMetrics, CustomMetric{
		Name:   name,
		Type:   metricType,
		Help:   help,
		Labels: labels,
	})
	return m
}

// AddAlertRule 添加告警规则
func (m *Monitoring) AddAlertRule(name, query, condition string, threshold float64, duration, severity string) *Monitoring {
	if m.Alerting == nil {
		m.Alerting = &Alerting{}
	}
	m.Alerting.Rules = append(m.Alerting.Rules, AlertRule{
		Name:        name,
		Query:       query,
		Condition:   condition,
		Threshold:   threshold,
		Duration:    duration,
		Severity:    severity,
		Labels:      make(map[string]string),
		Annotations: make(map[string]string),
	})
	return m
}

// AddWebhook 添加Webhook
func (m *Monitoring) AddWebhook(name, url string, headers map[string]string, timeout string) *Monitoring {
	if m.Alerting == nil {
		m.Alerting = &Alerting{}
	}
	m.Alerting.Webhooks = append(m.Alerting.Webhooks, Webhook{
		Name:    name,
		URL:     url,
		Headers: headers,
		Timeout: timeout,
	})
	return m
}

// EnablePrometheus 启用Prometheus
func (m *Monitoring) EnablePrometheus() *Monitoring {
	if m.Prometheus != nil {
		m.Prometheus.Enable()
	}
	return m
}

// EnableGrafana 启用Grafana
func (m *Monitoring) EnableGrafana() *Monitoring {
	if m.Grafana != nil {
		m.Grafana.Enable()
	}
	return m
}

// EnableJaeger 启用Jaeger
func (m *Monitoring) EnableJaeger() *Monitoring {
	if m.Jaeger != nil {
		m.Jaeger.Enable()
	}
	return m
}

// EnableMetrics 启用指标
func (m *Monitoring) EnableMetrics() *Monitoring {
	if m.Metrics == nil {
		m.Metrics = &Metrics{}
	}
	m.Metrics.Enabled = true
	return m
}

// EnableAlerting 启用告警
func (m *Monitoring) EnableAlerting() *Monitoring {
	if m.Alerting == nil {
		m.Alerting = &Alerting{}
	}
	m.Alerting.Enabled = true
	return m
}

// Enable 启用监控
func (m *Monitoring) Enable() *Monitoring {
	m.Enabled = true
	return m
}

// Disable 禁用监控
func (m *Monitoring) Disable() *Monitoring {
	m.Enabled = false
	return m
}

// IsEnabled 检查是否启用
func (m *Monitoring) IsEnabled() bool {
	return m.Enabled
}

// GetEndpoint 获取 Metrics 的 Endpoint（确保已计算）
func (m *Monitoring) GetEndpoint() string {
	if m.Metrics != nil && m.Metrics.Endpoint == "" {
		internal.CallAfterLoad(m)
	}
	if m.Metrics != nil {
		return m.Metrics.Endpoint
	}
	return "/metrics" // 默认值
}
