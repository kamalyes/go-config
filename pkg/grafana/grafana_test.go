/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 16:51:10
 * @FilePath: \go-config\pkg\grafana\grafana_test.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package grafana

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrafana_Default(t *testing.T) {
	config := Default()
	assert.NotNil(t, config)
	assert.Equal(t, "grafana", config.ModuleName)
	assert.False(t, config.Enabled)
	assert.Equal(t, "http://localhost:3000", config.Endpoint)
	assert.NotNil(t, config.Datasource)
	assert.NotNil(t, config.Dashboard)
	assert.NotNil(t, config.Alerting)
	assert.Equal(t, "prometheus", config.Datasource.Type)
	assert.Equal(t, "http://localhost:9090", config.Datasource.URL)
	assert.Equal(t, "30s", config.Dashboard.RefreshInterval)
	assert.False(t, config.Alerting.Enabled)
}

func TestGrafana_WithModuleName(t *testing.T) {
	config := Default().WithModuleName("custom-grafana")
	assert.Equal(t, "custom-grafana", config.ModuleName)
}

func TestGrafana_WithEnabled(t *testing.T) {
	config := Default().WithEnabled(true)
	assert.True(t, config.Enabled)
}

func TestGrafana_WithEndpoint(t *testing.T) {
	config := Default().WithEndpoint("http://grafana.example.com")
	assert.Equal(t, "http://grafana.example.com", config.Endpoint)
}

func TestGrafana_WithCredentials(t *testing.T) {
	config := Default().WithCredentials("admin", "password123")
	assert.Equal(t, "admin", config.Username)
	assert.Equal(t, "password123", config.Password)
}

func TestGrafana_WithAPIKey(t *testing.T) {
	config := Default().WithAPIKey("test-api-key-12345")
	assert.Equal(t, "test-api-key-12345", config.APIKey)
}

func TestGrafana_WithDatasource(t *testing.T) {
	config := Default().WithDatasource("influxdb", "http://influx:8086", "mydb", "user", "pass")
	assert.Equal(t, "influxdb", config.Datasource.Type)
	assert.Equal(t, "http://influx:8086", config.Datasource.URL)
	assert.Equal(t, "mydb", config.Datasource.Database)
	assert.Equal(t, "user", config.Datasource.Username)
	assert.Equal(t, "pass", config.Datasource.Password)
}

func TestGrafana_WithDashboard(t *testing.T) {
	templates := []string{"template1.json", "template2.json"}
	config := Default().WithDashboard("/dashboards", true, templates, "1m")
	assert.Equal(t, "/dashboards", config.Dashboard.ImportPath)
	assert.True(t, config.Dashboard.AutoImport)
	assert.Equal(t, templates, config.Dashboard.Templates)
	assert.Equal(t, "1m", config.Dashboard.RefreshInterval)
}

func TestGrafana_WithAlerting(t *testing.T) {
	webhooks := []string{"http://hook1", "http://hook2"}
	config := Default().WithAlerting(true, webhooks)
	assert.True(t, config.Alerting.Enabled)
	assert.Equal(t, webhooks, config.Alerting.Webhooks)
}

func TestGrafana_AddNotificationChannel(t *testing.T) {
	settings := map[string]string{
		"url": "http://slack.webhook",
	}
	config := Default().AddNotificationChannel("slack-channel", "slack", settings)
	assert.Len(t, config.Alerting.Channels, 1)
	assert.Equal(t, "slack-channel", config.Alerting.Channels[0].Name)
	assert.Equal(t, "slack", config.Alerting.Channels[0].Type)
	assert.Equal(t, settings, config.Alerting.Channels[0].Settings)
}

func TestGrafana_EnableAlerting(t *testing.T) {
	config := Default().EnableAlerting()
	assert.True(t, config.Alerting.Enabled)
}

func TestGrafana_Enable(t *testing.T) {
	config := Default().Enable()
	assert.True(t, config.Enabled)
}

func TestGrafana_Disable(t *testing.T) {
	config := Default().Enable().Disable()
	assert.False(t, config.Enabled)
}

func TestGrafana_IsEnabled(t *testing.T) {
	config := Default()
	assert.False(t, config.IsEnabled())
	config.Enable()
	assert.True(t, config.IsEnabled())
}

func TestGrafana_Clone(t *testing.T) {
	original := Default().
		WithModuleName("test-grafana").
		WithEnabled(true).
		WithEndpoint("http://grafana.test").
		WithCredentials("admin", "secret").
		WithAPIKey("api-key-123").
		AddNotificationChannel("email", "email", map[string]string{"to": "test@example.com"})

	cloned := original.Clone().(*Grafana)

	// 验证值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.Equal(t, original.Endpoint, cloned.Endpoint)
	assert.Equal(t, original.Username, cloned.Username)
	assert.Equal(t, original.Password, cloned.Password)
	assert.Equal(t, original.APIKey, cloned.APIKey)

	// 验证切片独立性
	cloned.Dashboard.Templates = append(cloned.Dashboard.Templates, "new-template")
	assert.NotEqual(t, len(original.Dashboard.Templates), len(cloned.Dashboard.Templates))

	// 验证map独立性
	cloned.Alerting.Channels[0].Settings["new-key"] = "new-value"
	assert.NotEqual(t, len(original.Alerting.Channels[0].Settings), len(cloned.Alerting.Channels[0].Settings))
}

func TestGrafana_Get(t *testing.T) {
	config := Default().WithEnabled(true)
	got := config.Get()
	assert.NotNil(t, got)
	grafanaConfig, ok := got.(*Grafana)
	assert.True(t, ok)
	assert.True(t, grafanaConfig.Enabled)
}

func TestGrafana_Set(t *testing.T) {
	config := Default()
	newConfig := &Grafana{
		ModuleName: "new-grafana",
		Enabled:    true,
		Endpoint:   "http://new.grafana.com",
		Username:   "newuser",
		Password:   "newpass",
	}

	config.Set(newConfig)
	assert.Equal(t, "new-grafana", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "http://new.grafana.com", config.Endpoint)
}

func TestGrafana_Validate(t *testing.T) {
	config := Default()
	err := config.Validate()
	assert.NoError(t, err)
}

func TestGrafana_ChainedCalls(t *testing.T) {
	config := Default().
		WithModuleName("chain-grafana").
		Enable().
		WithEndpoint("http://chain.grafana.com").
		WithCredentials("chainuser", "chainpass").
		WithAPIKey("chain-api-key").
		WithDatasource("prometheus", "http://prom:9090", "", "", "").
		EnableAlerting()

	assert.Equal(t, "chain-grafana", config.ModuleName)
	assert.True(t, config.Enabled)
	assert.Equal(t, "http://chain.grafana.com", config.Endpoint)
	assert.Equal(t, "chainuser", config.Username)
	assert.Equal(t, "chainpass", config.Password)
	assert.Equal(t, "chain-api-key", config.APIKey)
	assert.True(t, config.Alerting.Enabled)
}
