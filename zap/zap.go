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

// Zap 日志配置
type Zap struct {

	/** 日志级别 */
	Level                string       `mapstructure:"level"               json:"level"             yaml:"level"`

	/** 日志格式 */
	Format               string       `mapstructure:"format"              json:"format"            yaml:"format"`

	/** 日志前缀 */
	Prefix               string       `mapstructure:"prefix"              json:"prefix"            yaml:"prefix"`

	/** 日志目录 */
	Director             string       `mapstructure:"director"            json:"director"          yaml:"director"`

	/** 在进行切割之前，日志文件的最大大小（以MB为单位） */
	MaxSize              int          `mapstructure:"max-size"            json:"maxSize"           yaml:"max-size"`

	/** 日志最大保留时间 单位：天 */
	MaxAge               int          `mapstructure:"max-age"             json:"maxAge"            yaml:"max-age"`

	/** 保留旧文件的最大个数 */
	MaxBackups           int          `mapstructure:"max-backups"         json:"maxBackups"        yaml:"max-backups"`

	/** 是否压缩 */
	Compress             bool         `mapstructure:"compress"            json:"compress"          yaml:"compress"`

	/** 日志软连接文件 */
	LinkName             string       `mapstructure:"link-name"           json:"linkName"          yaml:"link-name"`

	/** 是否在日志中输出源码所在的行 */
	ShowLine             bool         `mapstructure:"show-line"           json:"showLine"          yaml:"showLine"`

	/** 日志编码等级，指定不通过等级可以有不同颜色 */
	EncodeLevel          string       `mapstructure:"encode-level"        json:"encodeLevel"       yaml:"encode-level"`

	/** 堆栈捕捉标识 */
	StacktraceKey        string       `mapstructure:"stacktrace-key"      json:"stacktraceKey"     yaml:"stacktrace-key"`

	/** 是否在控制台打印日志 */
	LogInConsole         bool         `mapstructure:"log-in-console"      json:"logInConsole"      yaml:"log-in-console"`
}
