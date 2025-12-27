/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-27 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-27 11:30:00
 * @FilePath: \go-config\pkg\monitoring\monitoring_test.go
 * @Description: 监控配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package monitoring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonitoring_Clone(t *testing.T) {
	original := Default()
	original.Enabled = true
	original.Metrics.Enabled = true
	original.Metrics.RequestCount = true
	original.Alerting.Enabled = true

	cloned := original.Clone().(*Monitoring)

	// 验证克隆后的值相等
	assert.Equal(t, original.ModuleName, cloned.ModuleName)
	assert.Equal(t, original.Enabled, cloned.Enabled)
	assert.NotSame(t, original.Metrics, cloned.Metrics)
	assert.Equal(t, original.Metrics.Enabled, cloned.Metrics.Enabled)
	assert.Equal(t, original.Metrics.RequestCount, cloned.Metrics.RequestCount)
	assert.NotSame(t, original.Alerting, cloned.Alerting)
	assert.Equal(t, original.Alerting.Enabled, cloned.Alerting.Enabled)

	// 修改原始对象不应影响克隆对象
	original.Enabled = false
	original.Metrics.Enabled = false
	assert.NotEqual(t, original.Enabled, cloned.Enabled)
	assert.NotEqual(t, original.Metrics.Enabled, cloned.Metrics.Enabled)
}
