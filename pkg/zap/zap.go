/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:37:51
 * @FilePath: \go-config\pkg\zap\zap.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package zap

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Zap 结构体表示 Zap 日志配置
type Zap struct {
	ModuleName    string `mapstructure:"MODULE_NAME"              yaml:"modulename"`     // 模块名称
	Level         string `mapstructure:"LEVEL"                    yaml:"level"`          // 日志级别
	Format        string `mapstructure:"FORMAT"                   yaml:"format"`         // 日志格式
	Prefix        string `mapstructure:"PREFIX"                   yaml:"prefix"`         // 日志前缀
	Director      string `mapstructure:"DIRECTOR"                 yaml:"director"`       // 日志目录
	MaxSize       int    `mapstructure:"MAX_SIZE"                 yaml:"max-size"`       // 日志文件的最大大小（以MB为单位）
	MaxAge        int    `mapstructure:"MAX_AGE"                  yaml:"max-age"`        // 日志最大保留时间 单位：天
	MaxBackups    int    `mapstructure:"MAX_BACKUPS"              yaml:"max-backups"`    // 保留旧文件的最大个数
	Compress      bool   `mapstructure:"COMPRESS"                 yaml:"compress"`       // 是否压缩
	LinkName      string `mapstructure:"LINK_NAME"                yaml:"link-name"`      // 日志软连接文件
	ShowLine      bool   `mapstructure:"SHOW_LINE"                yaml:"show-line"`      // 是否在日志中输出源码所在的行
	EncodeLevel   string `mapstructure:"ENCODE_LEVEL"             yaml:"encode-level"`   // 日志编码等级，指定不通过等级可以有不同颜色
	StacktraceKey string `mapstructure:"STACKTRACE_KEY"           yaml:"stacktrace-key"` // 堆栈捕捉标识
	LogInConsole  bool   `mapstructure:"LOG_IN_CONSOLE"           yaml:"log-in-console"` // 是否在控制台打印日志
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
	if z.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if z.Level == "" {
		return errors.New("log level cannot be empty")
	}
	if z.Format == "" {
		return errors.New("log format cannot be empty")
	}
	if z.Director == "" {
		return errors.New("log directory cannot be empty")
	}
	if z.MaxSize <= 0 {
		return errors.New("max size must be greater than 0")
	}
	if z.MaxAge < 0 {
		return errors.New("max age cannot be negative")
	}
	if z.MaxBackups < 0 {
		return errors.New("max backups cannot be negative")
	}
	return nil
}
