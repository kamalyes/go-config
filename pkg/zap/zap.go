/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-26 13:52:55
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
	Level         string `mapstructure:"level" yaml:"level" json:"level"                 validate:"required,oneof=debug info warn error fatal"` // 日志级别
	Format        string `mapstructure:"format" yaml:"format" json:"format"                validate:"required,oneof=json console"`              // 日志格式
	Prefix        string `mapstructure:"prefix" yaml:"prefix" json:"prefix"`                                                                    // 日志前缀
	Director      string `mapstructure:"director" yaml:"director" json:"director"              validate:"required"`                             // 日志目录
	MaxSize       int    `mapstructure:"max-size" yaml:"max-size" json:"maxSize"              validate:"required,min=1"`                        // 日志文件的最大大小（以MB为单位）
	MaxAge        int    `mapstructure:"max-age" yaml:"max-age" json:"maxAge"               validate:"required,min=1"`                          // 日志最大保留时间 单位：天
	MaxBackups    int    `mapstructure:"max-backups" yaml:"max-backups" json:"maxBackups"           validate:"required,min=0"`                  // 保留旧文件的最大个数
	Compress      bool   `mapstructure:"compress" yaml:"compress" json:"compress"`                                                              // 是否压缩
	LinkName      string `mapstructure:"link-name" yaml:"link-name" json:"linkName"`                                                            // 日志软连接文件
	ShowLine      bool   `mapstructure:"show-line" yaml:"show-line" json:"showLine"`                                                            // 是否在日志中输出源码所在的行
	EncodeLevel   string `mapstructure:"encode-level" yaml:"encode-level" json:"encodeLevel"`                                                   // 日志编码等级，指定不通过等级可以有不同颜色
	StacktraceKey string `mapstructure:"stacktrace-key" yaml:"stacktrace-key" json:"stacktraceKey"`                                             // 堆栈捕捉标识
	LogInConsole  bool   `mapstructure:"log-in-console" yaml:"log-in-console" json:"logInConsole"`                                              // 是否在控制台打印日志
	Development   bool   `mapstructure:"development" yaml:"development" json:"development"`                                                     // 是否为开发者模式
	ModuleName    string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                                                      // 模块名称
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
		Development:   z.Development,
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
		z.Development = configData.Development
		z.LogInConsole = configData.LogInConsole
	}
}

// Validate 验证 Zap 配置的有效性
func (z *Zap) Validate() error {
	return internal.ValidateStruct(z)
}

// DefaultZap 返回默认Zap配置
func DefaultZap() Zap {
	return Zap{
		ModuleName:    "zap",
		Level:         "info",
		Format:        "console",
		Prefix:        "[GO-CONFIG]",
		Director:      "./logs",
		MaxSize:       10,
		MaxAge:        30,
		MaxBackups:    5,
		Compress:      false,
		LinkName:      "latest_log",
		ShowLine:      true,
		EncodeLevel:   "capitalColor",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
		Development:   false,
	}
}

// Default 返回默认Zap配置的指针，支持链式调用
func Default() *Zap {
	config := DefaultZap()
	return &config
}

// WithModuleName 设置模块名称
func (z *Zap) WithModuleName(moduleName string) *Zap {
	z.ModuleName = moduleName
	return z
}

// WithLevel 设置日志级别
func (z *Zap) WithLevel(level string) *Zap {
	z.Level = level
	return z
}

// WithFormat 设置日志格式
func (z *Zap) WithFormat(format string) *Zap {
	z.Format = format
	return z
}

// WithPrefix 设置日志前缀
func (z *Zap) WithPrefix(prefix string) *Zap {
	z.Prefix = prefix
	return z
}

// WithDirector 设置日志目录
func (z *Zap) WithDirector(director string) *Zap {
	z.Director = director
	return z
}

// WithMaxSize 设置日志文件最大大小
func (z *Zap) WithMaxSize(maxSize int) *Zap {
	z.MaxSize = maxSize
	return z
}

// WithMaxAge 设置日志最大保留时间
func (z *Zap) WithMaxAge(maxAge int) *Zap {
	z.MaxAge = maxAge
	return z
}

// WithMaxBackups 设置保留旧文件最大个数
func (z *Zap) WithMaxBackups(maxBackups int) *Zap {
	z.MaxBackups = maxBackups
	return z
}

// WithCompress 设置是否压缩
func (z *Zap) WithCompress(compress bool) *Zap {
	z.Compress = compress
	return z
}

// WithLinkName 设置日志软连接文件名
func (z *Zap) WithLinkName(linkName string) *Zap {
	z.LinkName = linkName
	return z
}

// WithShowLine 设置是否在日志中输出源码行
func (z *Zap) WithShowLine(showLine bool) *Zap {
	z.ShowLine = showLine
	return z
}

// WithEncodeLevel 设置日志编码等级
func (z *Zap) WithEncodeLevel(encodeLevel string) *Zap {
	z.EncodeLevel = encodeLevel
	return z
}

// WithStacktraceKey 设置堆栈捕捉标识
func (z *Zap) WithStacktraceKey(stacktraceKey string) *Zap {
	z.StacktraceKey = stacktraceKey
	return z
}

// WithLogInConsole 设置是否在控制台打印日志
func (z *Zap) WithLogInConsole(logInConsole bool) *Zap {
	z.LogInConsole = logInConsole
	return z
}

// WithDevelopment 设置是否为开发者模式
func (z *Zap) WithDevelopment(development bool) *Zap {
	z.Development = development
	return z
}
