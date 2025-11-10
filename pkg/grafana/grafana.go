/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 18:00:00
 * @FilePath: \go-config\pkg\grafana\grafana.go
 * @Description: Grafana配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package grafana

import "github.com/kamalyes/go-config/internal"

// Grafana Grafana配置
type Grafana struct {
	ModuleName string      `mapstructure:"module_name" yaml:"module_name" json:"module_name"` // 模块名称
	Enabled    bool        `mapstructure:"enabled" yaml:"enabled" json:"enabled"`             // 是否启用Grafana
	Endpoint   string      `mapstructure:"endpoint" yaml:"endpoint" json:"endpoint"`          // Grafana端点
	Username   string      `mapstructure:"username" yaml:"username" json:"username"`          // 用户名
	Password   string      `mapstructure:"password" yaml:"password" json:"password"`          // 密码
	APIKey     string      `mapstructure:"api_key" yaml:"api_key" json:"api_key"`             // API Key
	Datasource *Datasource `mapstructure:"datasource" yaml:"datasource" json:"datasource"`    // 数据源配置
	Dashboard  *Dashboard  `mapstructure:"dashboard" yaml:"dashboard" json:"dashboard"`       // 仪表盘配置
	Alerting   *Alerting   `mapstructure:"alerting" yaml:"alerting" json:"alerting"`          // 告警配置
}

// Datasource 数据源配置
type Datasource struct {
	Type     string `mapstructure:"type" yaml:"type" json:"type"`             // 数据源类型 (prometheus, influxdb, etc.)
	URL      string `mapstructure:"url" yaml:"url" json:"url"`                // 数据源URL
	Database string `mapstructure:"database" yaml:"database" json:"database"` // 数据库名称
	Username string `mapstructure:"username" yaml:"username" json:"username"` // 用户名
	Password string `mapstructure:"password" yaml:"password" json:"password"` // 密码
}

// Dashboard 仪表盘配置
type Dashboard struct {
	ImportPath      string   `mapstructure:"import_path" yaml:"import_path" json:"import_path"`                // 导入路径
	AutoImport      bool     `mapstructure:"auto_import" yaml:"auto_import" json:"auto_import"`                // 自动导入
	Templates       []string `mapstructure:"templates" yaml:"templates" json:"templates"`                      // 模板列表
	RefreshInterval string   `mapstructure:"refresh_interval" yaml:"refresh_interval" json:"refresh_interval"` // 刷新间隔
}

// Alerting 告警配置
type Alerting struct {
	Enabled  bool                  `mapstructure:"enabled" yaml:"enabled" json:"enabled"`    // 是否启用告警
	Webhooks []string              `mapstructure:"webhooks" yaml:"webhooks" json:"webhooks"` // Webhook列表
	Channels []NotificationChannel `mapstructure:"channels" yaml:"channels" json:"channels"` // 通知渠道
}

// NotificationChannel 通知渠道配置
type NotificationChannel struct {
	Name     string            `mapstructure:"name" yaml:"name" json:"name"`             // 渠道名称
	Type     string            `mapstructure:"type" yaml:"type" json:"type"`             // 渠道类型 (slack, email, webhook)
	Settings map[string]string `mapstructure:"settings" yaml:"settings" json:"settings"` // 渠道设置
}

// Default 创建默认Grafana配置
func Default() *Grafana {
	return &Grafana{
		ModuleName: "grafana",
		Enabled:    false,
		Endpoint:   "",
		Username:   "",
		Password:   "",
		APIKey:     "",
		Datasource: &Datasource{
			Type:     "prometheus",
			URL:      "http://localhost:9090",
			Database: "",
			Username: "",
			Password: "",
		},
		Dashboard: &Dashboard{
			ImportPath:      "",
			AutoImport:      false,
			Templates:       []string{},
			RefreshInterval: "30s",
		},
		Alerting: &Alerting{
			Enabled:  false,
			Webhooks: []string{},
			Channels: []NotificationChannel{},
		},
	}
}

// Get 返回配置接口
func (g *Grafana) Get() interface{} {
	return g
}

// Set 设置配置数据
func (g *Grafana) Set(data interface{}) {
	if cfg, ok := data.(*Grafana); ok {
		*g = *cfg
	}
}

// Clone 返回配置的副本
func (g *Grafana) Clone() internal.Configurable {
	datasource := &Datasource{}
	dashboard := &Dashboard{}
	alerting := &Alerting{}

	if g.Datasource != nil {
		*datasource = *g.Datasource
	}
	if g.Dashboard != nil {
		*dashboard = *g.Dashboard
		dashboard.Templates = append([]string(nil), g.Dashboard.Templates...)
	}
	if g.Alerting != nil {
		*alerting = *g.Alerting
		alerting.Webhooks = append([]string(nil), g.Alerting.Webhooks...)
		alerting.Channels = make([]NotificationChannel, len(g.Alerting.Channels))
		for i, ch := range g.Alerting.Channels {
			alerting.Channels[i] = NotificationChannel{
				Name:     ch.Name,
				Type:     ch.Type,
				Settings: make(map[string]string),
			}
			for k, v := range ch.Settings {
				alerting.Channels[i].Settings[k] = v
			}
		}
	}

	return &Grafana{
		ModuleName: g.ModuleName,
		Enabled:    g.Enabled,
		Endpoint:   g.Endpoint,
		Username:   g.Username,
		Password:   g.Password,
		APIKey:     g.APIKey,
		Datasource: datasource,
		Dashboard:  dashboard,
		Alerting:   alerting,
	}
}

// Validate 验证配置
func (g *Grafana) Validate() error {
	return internal.ValidateStruct(g)
}

// WithModuleName 设置模块名称
func (g *Grafana) WithModuleName(moduleName string) *Grafana {
	g.ModuleName = moduleName
	return g
}

// WithEnabled 设置是否启用
func (g *Grafana) WithEnabled(enabled bool) *Grafana {
	g.Enabled = enabled
	return g
}

// WithEndpoint 设置端点
func (g *Grafana) WithEndpoint(endpoint string) *Grafana {
	g.Endpoint = endpoint
	return g
}

// WithCredentials 设置凭证
func (g *Grafana) WithCredentials(username, password string) *Grafana {
	g.Username = username
	g.Password = password
	return g
}

// WithAPIKey 设置API Key
func (g *Grafana) WithAPIKey(apiKey string) *Grafana {
	g.APIKey = apiKey
	return g
}

// WithDatasource 设置数据源
func (g *Grafana) WithDatasource(dsType, url, database, username, password string) *Grafana {
	if g.Datasource == nil {
		g.Datasource = &Datasource{}
	}
	g.Datasource.Type = dsType
	g.Datasource.URL = url
	g.Datasource.Database = database
	g.Datasource.Username = username
	g.Datasource.Password = password
	return g
}

// WithDashboard 设置仪表盘
func (g *Grafana) WithDashboard(importPath string, autoImport bool, templates []string, refreshInterval string) *Grafana {
	if g.Dashboard == nil {
		g.Dashboard = &Dashboard{}
	}
	g.Dashboard.ImportPath = importPath
	g.Dashboard.AutoImport = autoImport
	g.Dashboard.Templates = templates
	g.Dashboard.RefreshInterval = refreshInterval
	return g
}

// WithAlerting 设置告警
func (g *Grafana) WithAlerting(enabled bool, webhooks []string) *Grafana {
	if g.Alerting == nil {
		g.Alerting = &Alerting{}
	}
	g.Alerting.Enabled = enabled
	g.Alerting.Webhooks = webhooks
	return g
}

// AddNotificationChannel 添加通知渠道
func (g *Grafana) AddNotificationChannel(name, channelType string, settings map[string]string) *Grafana {
	if g.Alerting == nil {
		g.Alerting = &Alerting{}
	}
	g.Alerting.Channels = append(g.Alerting.Channels, NotificationChannel{
		Name:     name,
		Type:     channelType,
		Settings: settings,
	})
	return g
}

// EnableAlerting 启用告警
func (g *Grafana) EnableAlerting() *Grafana {
	if g.Alerting == nil {
		g.Alerting = &Alerting{}
	}
	g.Alerting.Enabled = true
	return g
}

// Enable 启用Grafana
func (g *Grafana) Enable() *Grafana {
	g.Enabled = true
	return g
}

// Disable 禁用Grafana
func (g *Grafana) Disable() *Grafana {
	g.Enabled = false
	return g
}

// IsEnabled 检查是否启用
func (g *Grafana) IsEnabled() bool {
	return g.Enabled
}
