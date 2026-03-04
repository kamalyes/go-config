/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 15:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 12:49:29
 * @FilePath: \go-config\config_formatter.go
 * @Description: 配置信息格式化输出工具 - 使用反射自动解析结构体
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"sync"
	"time"

	"github.com/kamalyes/go-logger"
)

var (
	// 全局自动日志开关，默认开启
	autoLogEnabled = true
	autoLogMutex   sync.RWMutex
)

// ConfigFormatter 配置格式化器
type ConfigFormatter struct {
	logger *logger.Logger
}

// NewConfigFormatter 创建配置格式化器
func NewConfigFormatter(lg ...*logger.Logger) *ConfigFormatter {
	var log *logger.Logger
	if len(lg) > 0 && lg[0] != nil {
		log = lg[0]
	} else {
		log = logger.GetGlobalLogger()
	}

	return &ConfigFormatter{
		logger: log,
	}
}

// LogConfigChanged 记录配置变更 - 主要入口函数
func (cf *ConfigFormatter) LogConfigChanged(event CallbackEvent, newConfig any) {
	cf.logger.Info("🔄 配置发生变更!")
	cf.logger.Info("   📂 来源: %s", event.Source)
	cf.logger.Info("   🕐 时间: %s", event.Timestamp.Format(time.DateTime))
	cf.logger.Info("   🌍 环境: %s", event.Environment)
	cf.logger.Info("   📋 事件类型: %s", event.Type)

	// 根据配置类型记录详细信息
	cf.logger.Info("🆕 配置已更新: %T", newConfig)
}
