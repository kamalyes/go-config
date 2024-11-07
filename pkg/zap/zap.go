/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 10:56:22
 * @FilePath: \go-config\pkg\zap\zap.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package zap

import (
	"github.com/kamalyes/go-config/internal"
)

// Zap 结构体表示 Zap 日志配置
type Zap struct {
	Level         string `mapstructure:"level"                    yaml:"level"               json:"level"                 validate:"required,oneof=debug info warn error fatal"` // 日志级别
	Format        string `mapstructure:"format"                   yaml:"format"              json:"format"                validate:"required,oneof=json console"`                // 日志格式
	Prefix        string `mapstructure:"prefix"                   yaml:"prefix"              json:"prefix"`                                                                      // 日志前缀
	Director      string `mapstructure:"director"                 yaml:"director"            json:"director"              validate:"required"`                                   // 日志目录
	MaxSize       int    `mapstructure:"max-size"                 yaml:"max-size"            json:"max_size"              validate:"required,min=1"`                             // 日志文件的最大大小（以MB为单位）
	MaxAge        int    `mapstructure:"max-age"                  yaml:"max-age"             json:"max_age"               validate:"required,min=1"`                             // 日志最大保留时间 单位：天
	MaxBackups    int    `mapstructure:"max-backups"              yaml:"max-backups"         json:"max_backups"           validate:"required,min=0"`                             // 保留旧文件的最大个数
	Compress      bool   `mapstructure:"compress"                 yaml:"compress"            json:"compress"`                                                                    // 是否压缩
	LinkName      string `mapstructure:"link-name"                yaml:"link-name"           json:"link_name"`                                                                   // 日志软连接文件
	ShowLine      bool   `mapstructure:"show-line"                yaml:"show-line"           json:"show_line"`                                                                   // 是否在日志中输出源码所在的行
	EncodeLevel   string `mapstructure:"encode-level"             yaml:"encode-level"        json:"encode_level"`                                                                // 日志编码等级，指定不通过等级可以有不同颜色
	StacktraceKey string `mapstructure:"stacktrace-key"           yaml:"stacktrace-key"      json:"stacktrace_key"`                                                              // 堆栈捕捉标识
	LogInConsole  bool   `mapstructure:"log-in-console"           yaml:"log-in-console"      json:"log_in_console"`                                                              // 是否在控制台打印日志
	ModuleName    string `mapstructure:"modulename"               yaml:"modulename"          json:"module_name"`                                                                 // 模块名称
}

// NewZap 创建一个新的 Zap 实例
func NewZap(opt *Zap) *Zap {
	var zapInstance *Zap

	internal.LockFunc(func() {
		zapInstance = opt
	})
	return zapInstance
}

// Clone 返回 Zap 配置的副本
func (z *Zap) Clone() internal.Configurable {
	return &Zap{
		ModuleName:    z.ModuleName,
		Level:         z.Level,
		Format:        z.Format,
		Prefix:        z.Prefix,
		Director:      z.Director,
		MaxSize:       z.MaxSize,
		MaxAge:        z.MaxAge,
		MaxBackups:    z.MaxBackups,
		Compress:      z.Compress,
		LinkName:      z.LinkName,
		ShowLine:      z.ShowLine,
		EncodeLevel:   z.EncodeLevel,
		StacktraceKey: z.StacktraceKey,
		LogInConsole:  z.LogInConsole,
	}
}

// Get 返回 Zap 配置的所有字段
func (z *Zap) Get() interface{} {
	return z
}

// Set 更新 Zap 配置的字段
func (z *Zap) Set(data interface{}) {
	if configData, ok := data.(*Zap); ok {
		z.ModuleName = configData.ModuleName
		z.Level = configData.Level
		z.Format = configData.Format
		z.Prefix = configData.Prefix
		z.Director = configData.Director
		z.MaxSize = configData.MaxSize
		z.MaxAge = configData.MaxAge
		z.MaxBackups = configData.MaxBackups
		z.Compress = configData.Compress
		z.LinkName = configData.LinkName
		z.ShowLine = configData.ShowLine
		z.EncodeLevel = configData.EncodeLevel
		z.StacktraceKey = configData.StacktraceKey
		z.LogInConsole = configData.LogInConsole
	}
}

// Validate 验证 Zap 配置的有效性
func (z *Zap) Validate() error {
	return internal.ValidateStruct(z)
}
