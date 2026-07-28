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
	"fmt"

	"github.com/kamalyes/go-logger"
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
	lines := []string{
		"🔄 配置发生变更!",
		fmt.Sprintf("   📂 来源: %s", event.Source),
		fmt.Sprintf("   🕐 时间: %s", event.Timestamp),
		fmt.Sprintf("   🌍 环境: %s", event.Environment),
		fmt.Sprintf("   📋 事件类型: %s", event.Type),
		fmt.Sprintf("   🔄 新配置: %T", newConfig),
	}

	cf.logger.DebugLines(lines...)
}
