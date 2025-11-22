/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-21 23:58:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-21 23:58:00
 * @FilePath: \go-config\pkg\alerting\alerting_test.go
 * @Description: 告警配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package alerting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlerting_Default(t *testing.T) {
	alerting := Default()

	assert.NotNil(t, alerting)
	assert.Equal(t, "alerting", alerting.ModuleName)
	assert.False(t, alerting.Enabled)
	assert.Empty(t, alerting.Webhooks)
	assert.Empty(t, alerting.Channels)
}

func TestAlerting_WithWebhooks(t *testing.T) {
	alerting := Default()
	webhooks := []string{"https://example.com/webhook1", "https://example.com/webhook2"}
	result := alerting.WithWebhooks(webhooks)

	assert.Equal(t, webhooks, result.Webhooks)
	assert.Equal(t, alerting, result)
}

func TestAlerting_AddWebhook(t *testing.T) {
	alerting := Default()
	webhook := "https://example.com/webhook"
	result := alerting.AddWebhook(webhook)

	assert.Contains(t, result.Webhooks, webhook)
	assert.Len(t, result.Webhooks, 1)
	assert.Equal(t, alerting, result)
}

func TestAlerting_AddChannel(t *testing.T) {
	alerting := Default()
	channel := NotificationChannel{
		Name: "slack-channel",
		Type: "slack",
		Settings: map[string]string{
			"url": "https://hooks.slack.com/services/xxx",
		},
	}
	result := alerting.AddChannel(channel.Name, channel.Type, channel.Settings)

	assert.Len(t, result.Channels, 1)
	assert.Equal(t, "slack-channel", result.Channels[0].Name)
	assert.Equal(t, "slack", result.Channels[0].Type)
	assert.Equal(t, alerting, result)
}

func TestAlerting_WithChannels(t *testing.T) {
	alerting := Default()
	channels := []NotificationChannel{
		{
			Name: "email-channel",
			Type: "email",
			Settings: map[string]string{
				"to": "admin@example.com",
			},
		},
		{
			Name: "slack-channel",
			Type: "slack",
			Settings: map[string]string{
				"url": "https://hooks.slack.com/services/xxx",
			},
		},
	}
	result := alerting.WithChannels(channels)

	assert.Equal(t, channels, result.Channels)
	assert.Len(t, result.Channels, 2)
	assert.Equal(t, alerting, result)
}

func TestAlerting_Enable(t *testing.T) {
	alerting := Default()
	result := alerting.Enable()

	assert.True(t, result.Enabled)
	assert.True(t, result.IsEnabled())
	assert.Equal(t, alerting, result)
}

func TestAlerting_Disable(t *testing.T) {
	alerting := Default()
	alerting.Enabled = true
	result := alerting.Disable()

	assert.False(t, result.Enabled)
	assert.False(t, result.IsEnabled())
	assert.Equal(t, alerting, result)
}

func TestAlerting_IsEnabled(t *testing.T) {
	alerting := Default()
	assert.False(t, alerting.IsEnabled())

	alerting.Enabled = true
	assert.True(t, alerting.IsEnabled())
}

func TestAlerting_Get(t *testing.T) {
	alerting := Default()
	result := alerting.Get()

	assert.NotNil(t, result)
	assert.Equal(t, alerting, result)
}

func TestAlerting_Set(t *testing.T) {
	alerting := Default()
	newAlerting := &Alerting{
		ModuleName: "custom-alerting",
		Enabled:    true,
		Webhooks:   []string{"https://new.com/webhook"},
		Channels: []NotificationChannel{
			{
				Name: "new-channel",
				Type: "webhook",
				Settings: map[string]string{
					"url": "https://new.com",
				},
			},
		},
	}

	alerting.Set(newAlerting)

	assert.Equal(t, "custom-alerting", alerting.ModuleName)
	assert.True(t, alerting.Enabled)
	assert.Equal(t, []string{"https://new.com/webhook"}, alerting.Webhooks)
	assert.Len(t, alerting.Channels, 1)
	assert.Equal(t, "new-channel", alerting.Channels[0].Name)
}

func TestAlerting_Set_InvalidType(t *testing.T) {
	alerting := Default()
	originalModuleName := alerting.ModuleName

	alerting.Set("invalid type")

	assert.Equal(t, originalModuleName, alerting.ModuleName)
}

func TestAlerting_Clone(t *testing.T) {
	alerting := Default()
	alerting.Webhooks = []string{"https://example.com/webhook"}
	alerting.Channels = []NotificationChannel{
		{
			Name: "test-channel",
			Type: "slack",
			Settings: map[string]string{
				"url": "https://slack.com",
			},
		},
	}

	cloned := alerting.Clone()

	assert.NotNil(t, cloned)
	clonedAlerting, ok := cloned.(*Alerting)
	assert.True(t, ok)
	assert.Equal(t, alerting.ModuleName, clonedAlerting.ModuleName)
	assert.Equal(t, alerting.Webhooks, clonedAlerting.Webhooks)
	assert.Equal(t, len(alerting.Channels), len(clonedAlerting.Channels))

	// 验证是独立副本
	clonedAlerting.Webhooks[0] = "https://modified.com"
	assert.NotEqual(t, alerting.Webhooks[0], clonedAlerting.Webhooks[0])

	clonedAlerting.Channels[0].Name = "modified-channel"
	assert.NotEqual(t, alerting.Channels[0].Name, clonedAlerting.Channels[0].Name)

	clonedAlerting.Channels[0].Settings["url"] = "https://modified.com"
	assert.NotEqual(t, alerting.Channels[0].Settings["url"], clonedAlerting.Channels[0].Settings["url"])
}

func TestAlerting_Validate(t *testing.T) {
	alerting := Default()
	err := alerting.Validate()
	assert.NoError(t, err)
}

func TestNotificationChannel_Creation(t *testing.T) {
	channel := NotificationChannel{
		Name: "test-channel",
		Type: "email",
		Settings: map[string]string{
			"to":   "test@example.com",
			"from": "noreply@example.com",
		},
	}

	assert.Equal(t, "test-channel", channel.Name)
	assert.Equal(t, "email", channel.Type)
	assert.Equal(t, "test@example.com", channel.Settings["to"])
	assert.Equal(t, "noreply@example.com", channel.Settings["from"])
}

func TestAlerting_ChainedCalls(t *testing.T) {
	alerting := Default()
	result := alerting.
		AddWebhook("https://webhook1.com").
		AddWebhook("https://webhook2.com").
		AddChannel("slack", "slack", map[string]string{
			"url": "https://slack.com",
		}).
		Enable()

	assert.Len(t, result.Webhooks, 2)
	assert.Len(t, result.Channels, 1)
	assert.True(t, result.Enabled)
}
