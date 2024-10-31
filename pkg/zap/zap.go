/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 09:26:18
 * @FilePath: \go-config\zap\zap.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package zap

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Zap 日志配置
type Zap struct {

	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** 日志级别 */
	Level string `mapstructure:"level"               json:"level"             yaml:"level"`

	/** 日志格式 */
	Format string `mapstructure:"format"              json:"format"            yaml:"format"`

	/** 日志前缀 */
	Prefix string `mapstructure:"prefix"              json:"prefix"            yaml:"prefix"`

	/** 日志目录 */
	Director string `mapstructure:"director"            json:"director"          yaml:"director"`

	/** 在进行切割之前，日志文件的最大大小（以MB为单位） */
	MaxSize int `mapstructure:"max-size"            json:"maxSize"           yaml:"max-size"`

	/** 日志最大保留时间 单位：天 */
	MaxAge int `mapstructure:"max-age"             json:"maxAge"            yaml:"max-age"`

	/** 保留旧文件的最大个数 */
	MaxBackups int `mapstructure:"max-backups"         json:"maxBackups"        yaml:"max-backups"`

	/** 是否压缩 */
	Compress bool `mapstructure:"compress"            json:"compress"          yaml:"compress"`

	/** 日志软连接文件 */
	LinkName string `mapstructure:"link-name"           json:"linkName"          yaml:"link-name"`

	/** 是否在日志中输出源码所在的行 */
	ShowLine bool `mapstructure:"show-line"           json:"showLine"          yaml:"showLine"`

	/** 日志编码等级，指定不通过等级可以有不同颜色 */
	EncodeLevel string `mapstructure:"encode-level"        json:"encodeLevel"       yaml:"encode-level"`

	/** 堆栈捕捉标识 */
	StacktraceKey string `mapstructure:"stacktrace-key"      json:"stacktraceKey"     yaml:"stacktrace-key"`

	/** 是否在控制台打印日志 */
	LogInConsole bool `mapstructure:"log-in-console"      json:"logInConsole"      yaml:"log-in-console"`
}

// NewZap 创建一个新的 Zap 实例
func NewZap(moduleName, level, format, prefix, director string, maxSize, maxAge, maxBackups int, compress bool, linkName string, showLine bool, encodeLevel, stacktraceKey string, logInConsole bool) *Zap {
	var zapInstance *Zap

	internal.LockFunc(func() {
		zapInstance = &Zap{
			ModuleName:    moduleName,
			Level:         level,
			Format:        format,
			Prefix:        prefix,
			Director:      director,
			MaxSize:       maxSize,
			MaxAge:        maxAge,
			MaxBackups:    maxBackups,
			Compress:      compress,
			LinkName:      linkName,
			ShowLine:      showLine,
			EncodeLevel:   encodeLevel,
			StacktraceKey: stacktraceKey,
			LogInConsole:  logInConsole,
		}
	})
	return zapInstance
}

// ToMap 将配置转换为映射
func (z *Zap) ToMap() map[string]interface{} {
	return internal.ToMap(z)
}

// FromMap 从映射中填充配置
func (z *Zap) FromMap(data map[string]interface{}) {
	internal.FromMap(z, data)
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
