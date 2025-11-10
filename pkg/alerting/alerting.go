/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 18:00:00
 * @FilePath: \go-config\pkg\alerting\alerting.go
 * @Description: 告警配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package alerting

import "github.com/kamalyes/go-config/internal"

// Alerting 告警配置
type Alerting struct {
	ModuleName string                `mapstructure:"module_name" yaml:"module-name" json:"module_name"` // 模块名称
	Enabled    bool                  `mapstructure:"enabled" yaml:"enabled" json:"enabled"`             // 是否启用告警
	Webhooks   []string              `mapstructure:"webhooks" yaml:"webhooks" json:"webhooks"`          // Webhook列表
	Channels   []NotificationChannel `mapstructure:"channels" yaml:"channels" json:"channels"`          // 通知渠道
}

// NotificationChannel 通知渠道配置
type NotificationChannel struct {
	Name     string            `mapstructure:"name" yaml:"name" json:"name"`             // 渠道名称
	Type     string            `mapstructure:"type" yaml:"type" json:"type"`             // 渠道类型 (slack, email, webhook)
	Settings map[string]string `mapstructure:"settings" yaml:"settings" json:"settings"` // 渠道设置
}

// Default 创建默认告警配置
func Default() *Alerting {
	return &Alerting{
		ModuleName: "alerting",
		Enabled:    false,
		Webhooks:   []string{},
		Channels:   []NotificationChannel{},
	}
}

// Get 返回配置接口
func (a *Alerting) Get() interface{} {
	return a
}

// Set 设置配置数据
func (a *Alerting) Set(data interface{}) {
	if cfg, ok := data.(*Alerting); ok {
		*a = *cfg
	}
}

// Clone 返回配置的副本
func (a *Alerting) Clone() internal.Configurable {
	clone := &Alerting{
		ModuleName: a.ModuleName,
		Enabled:    a.Enabled,
	}

	// 复制Webhooks
	clone.Webhooks = append([]string(nil), a.Webhooks...)

	// 复制Channels
	clone.Channels = make([]NotificationChannel, len(a.Channels))
	for i, channel := range a.Channels {
		clone.Channels[i] = NotificationChannel{
			Name: channel.Name,
			Type: channel.Type,
		}
		// 复制Settings
		clone.Channels[i].Settings = make(map[string]string)
		for k, v := range channel.Settings {
			clone.Channels[i].Settings[k] = v
		}
	}

	return clone
}

// Validate 验证配置
func (a *Alerting) Validate() error {
	return internal.ValidateStruct(a)
}

// WithWebhooks 设置Webhook列表
func (a *Alerting) WithWebhooks(webhooks []string) *Alerting {
	a.Webhooks = webhooks
	return a
}

// AddWebhook 添加Webhook
func (a *Alerting) AddWebhook(webhook string) *Alerting {
	a.Webhooks = append(a.Webhooks, webhook)
	return a
}

// WithChannels 设置通知渠道列表
func (a *Alerting) WithChannels(channels []NotificationChannel) *Alerting {
	a.Channels = channels
	return a
}

// AddChannel 添加通知渠道
func (a *Alerting) AddChannel(name, channelType string, settings map[string]string) *Alerting {
	channel := NotificationChannel{
		Name:     name,
		Type:     channelType,
		Settings: settings,
	}
	a.Channels = append(a.Channels, channel)
	return a
}

// AddSlackChannel 添加Slack通知渠道
func (a *Alerting) AddSlackChannel(name, webhook, channel string) *Alerting {
	settings := map[string]string{
		"webhook": webhook,
		"channel": channel,
	}
	return a.AddChannel(name, "slack", settings)
}

// AddEmailChannel 添加邮件通知渠道
func (a *Alerting) AddEmailChannel(name, smtpHost, smtpPort, username, password string, recipients []string) *Alerting {
	settings := map[string]string{
		"smtp_host": smtpHost,
		"smtp_port": smtpPort,
		"username":  username,
		"password":  password,
	}
	// 将recipients转换为逗号分隔的字符串
	recipientStr := ""
	for i, recipient := range recipients {
		if i > 0 {
			recipientStr += ","
		}
		recipientStr += recipient
	}
	settings["recipients"] = recipientStr

	return a.AddChannel(name, "email", settings)
}

// Enable 启用告警
func (a *Alerting) Enable() *Alerting {
	a.Enabled = true
	return a
}

// Disable 禁用告警
func (a *Alerting) Disable() *Alerting {
	a.Enabled = false
	return a
}

// IsEnabled 检查是否启用
func (a *Alerting) IsEnabled() bool {
	return a.Enabled
}
