/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-31 12:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 18:28:31
 * @FilePath: \go-config\pkg\zero\logx.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package zero

import (
	"github.com/kamalyes/go-config/internal"
)

// LogConf 结构体表示日志配置
type LogConf struct {
	ServiceName         string `mapstructure:"service-name"       yaml:"service-name"       json:"service_name"`                                             // 服务名称
	Mode                string `mapstructure:"mode"               yaml:"mode"               json:"mode" default:"console" options:"[console,file,volume]"`   // 日志模式
	Encoding            string `mapstructure:"encoding"           yaml:"encoding"           json:"encoding" default:"json" options:"[json,plain]"`           // 编码类型
	TimeFormat          string `mapstructure:"time-format"        yaml:"time-format"        json:"time_format"`                                              // 时间格式
	Path                string `mapstructure:"path"               yaml:"path"               json:"path" default:"logs"`                                      // 日志文件路径
	Level               string `mapstructure:"level"              yaml:"level"              json:"level" default:"info" options:"[debug,info,error,severe]"` // 日志级别
	MaxContentLength    uint32 `mapstructure:"max-content-length"  yaml:"max-content-length" json:"max_content_length"`                                      // 最大内容长度
	Compress            bool   `mapstructure:"compress"           yaml:"compress"           json:"compress"`                                                 // 是否压缩
	Stat                bool   `mapstructure:"stat"               yaml:"stat"               json:"stat" default:"true"`                                      // 记录统计信息
	KeepDays            int    `mapstructure:"keep-days"          yaml:"keep-days"          json:"keep_days"`                                                // 保留天数
	StackCooldownMillis int    `mapstructure:"stack-cooldown-millis" yaml:"stack-cooldown-millis" json:"stack_cooldown_millis" default:"100"`                // 堆栈冷却时间
	MaxBackups          int    `mapstructure:"max-backups"       yaml:"max-backups"       json:"max_backups" default:"0"`                                    // 最大备份数
	MaxSize             int    `mapstructure:"max-size"          yaml:"max-size"          json:"max-size" default:"0"`                                       // 最大文件大小
	Rotation            string `mapstructure:"rotation"          yaml:"rotation"          json:"rotation" default:"daily" options:"[daily,size]"`            // 轮换规则
	FileTimeFormat      string `mapstructure:"file-time-format"   yaml:"file-time-format"   json:"file_time_format"`                                         // 文件时间格式
}

// NewLogConf 创建一个新的 LogConf 实例
func NewLogConf(opt *LogConf) *LogConf {
	var logInstance *LogConf

	internal.LockFunc(func() {
		logInstance = opt
	})
	return logInstance
}

// Clone 返回 LogConf 配置的副本
func (l *LogConf) Clone() internal.Configurable {
	return &LogConf{
		ServiceName:         l.ServiceName,
		Mode:                l.Mode,
		Encoding:            l.Encoding,
		TimeFormat:          l.TimeFormat,
		Path:                l.Path,
		Level:               l.Level,
		MaxContentLength:    l.MaxContentLength,
		Compress:            l.Compress,
		Stat:                l.Stat,
		KeepDays:            l.KeepDays,
		StackCooldownMillis: l.StackCooldownMillis,
		MaxBackups:          l.MaxBackups,
		MaxSize:             l.MaxSize,
		Rotation:            l.Rotation,
		FileTimeFormat:      l.FileTimeFormat,
	}
}

// Get 返回 LogConf 配置的所有字段
func (l *LogConf) Get() interface{} {
	return l
}

// Set 更新 LogConf 配置的字段
func (l *LogConf) Set(data interface{}) {
	if configData, ok := data.(*LogConf); ok {
		l.ServiceName = configData.ServiceName
		l.Mode = configData.Mode
		l.Encoding = configData.Encoding
		l.TimeFormat = configData.TimeFormat
		l.Path = configData.Path
		l.Level = configData.Level
		l.MaxContentLength = configData.MaxContentLength
		l.Compress = configData.Compress
		l.Stat = configData.Stat
		l.KeepDays = configData.KeepDays
		l.StackCooldownMillis = configData.StackCooldownMillis
		l.MaxBackups = configData.MaxBackups
		l.MaxSize = configData.MaxSize
		l.Rotation = configData.Rotation
		l.FileTimeFormat = configData.FileTimeFormat
	}
}

// Validate 验证 LogConf 配置的有效性
func (l *LogConf) Validate() error {
	return internal.ValidateStruct(l)
}
