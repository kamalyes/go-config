/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-26 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-08-21 20:09:07
 * @FilePath: \go-config\config_formatter_test.go
 * @Description: 配置格式化器测试
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package goconfig

import (
	"testing"
	"time"

	"github.com/kamalyes/go-logger"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigFormatter_Default(t *testing.T) {
	cf := NewConfigFormatter()
	assert.NotNil(t, cf)
}

func TestNewConfigFormatter_CustomLogger(t *testing.T) {
	cf := NewConfigFormatter(logger.NewLogger())
	assert.NotNil(t, cf)
}

func TestNewConfigFormatter_NilLogger(t *testing.T) {
	cf := NewConfigFormatter(nil)
	assert.NotNil(t, cf)
}

func TestLogConfigChanged_NoPanic(t *testing.T) {
	cf := NewConfigFormatter()
	event := CallbackEvent{
		Type:        CallbackTypeConfigChanged,
		Source:      "test",
		Timestamp:   time.Now(),
		Environment: EnvDevelopment,
	}
	assert.NotPanics(t, func() {
		cf.LogConfigChanged(event, struct{ Name string }{"cfg"})
	})
}
