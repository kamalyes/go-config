/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 15:22:31
 * @FilePath: \go-config\pkg\queue\mqtt.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package queue

import (
	"errors"

	"github.com/kamalyes/go-config/internal"
)

// Mqtt 配置文件
type Mqtt struct {
	/** 模块名称 */
	ModuleName string `mapstructure:"modulename"              json:"moduleName"             yaml:"modulename"`

	/** Mqtt代理服务器的Ip和端口 */
	Url string `mapstructure:"url"                      json:"url"                       yaml:"url"`

	/** 连接到代理服务器的用户名 */
	Username string `mapstructure:"username"                 json:"username"                  yaml:"username"`

	/** 密码 */
	Password string `mapstructure:"password"                 json:"password"                  yaml:"password"`

	/** Mqtt 协议版本号 4是3.1.1，3是3.1 */
	ProtocolVersion uint `mapstructure:"protocol-ver"             json:"protocolVer"               yaml:"protocol-ver"`

	/** 设置客户端掉线服务端是否清除session */
	CleanSession bool `mapstructure:"clean-session"            json:"cleanSession"              yaml:"clean-session"`

	/** 断开后是否重新连接 */
	AutoReconnect bool `mapstructure:"auto-reconnect"           json:"autoReconnect"             yaml:"auto-reconnect"`

	/** 保活时间间隔 */
	KeepAlive int `mapstructure:"keep-alive"               json:"keepAlive"                 yaml:"keep-alive"`

	/** 最大连接间隔时间 单位：秒 */
	MaxReconnectInterval int `mapstructure:"max-reconnect-interval"   json:"maxReconnectInterval"      yaml:"max-reconnect-interval"`

	/** ping 超时时间 单位：秒 */
	PingTimeout int `mapstructure:"pingTimeout"              json:"ping-timeout"              yaml:"ping-timeout"`

	/** 写超时时间 单位：秒 */
	WriteTimeout int `mapstructure:"write-timeout"            json:"writeTimeout"              yaml:"write-timeout"`

	/** 连接超时时间 单位：秒 */
	ConnectTimeout int `mapstructure:"connect-timeout"          json:"connectTimeout"            yaml:"connect-timeout"`

	/** 遗言发送的topic */
	WillTopic string `mapstructure:"will-topic"               json:"willTopic"                 yaml:"will-topic"`
}

// NewMqtt 创建一个新的 Mqtt 实例
func NewMqtt(moduleName, url, username, password string, protocolVersion uint, cleanSession, autoReconnect bool, keepAlive, maxReconnectInterval, pingTimeout, writeTimeout, connectTimeout int, willTopic string) *Mqtt {
	var mqttInstance *Mqtt

	internal.LockFunc(func() {
		mqttInstance = &Mqtt{
			ModuleName:           moduleName,
			Url:                  url,
			Username:             username,
			Password:             password,
			ProtocolVersion:      protocolVersion,
			CleanSession:         cleanSession,
			AutoReconnect:        autoReconnect,
			KeepAlive:            keepAlive,
			MaxReconnectInterval: maxReconnectInterval,
			PingTimeout:          pingTimeout,
			WriteTimeout:         writeTimeout,
			ConnectTimeout:       connectTimeout,
			WillTopic:            willTopic,
		}
	})
	return mqttInstance
}

// ToMap 将配置转换为映射
func (m *Mqtt) ToMap() map[string]interface{} {
	return internal.ToMap(m)
}

// FromMap 从映射中填充配置
func (m *Mqtt) FromMap(data map[string]interface{}) {
	internal.FromMap(m, data)
}

// Clone 返回 Mqtt 配置的副本
func (m *Mqtt) Clone() internal.Configurable {
	return &Mqtt{
		ModuleName:           m.ModuleName,
		Url:                  m.Url,
		Username:             m.Username,
		Password:             m.Password,
		ProtocolVersion:      m.ProtocolVersion,
		CleanSession:         m.CleanSession,
		AutoReconnect:        m.AutoReconnect,
		KeepAlive:            m.KeepAlive,
		MaxReconnectInterval: m.MaxReconnectInterval,
		PingTimeout:          m.PingTimeout,
		WriteTimeout:         m.WriteTimeout,
		ConnectTimeout:       m.ConnectTimeout,
		WillTopic:            m.WillTopic,
	}
}

// Get 返回 Mqtt 配置的所有字段
func (m *Mqtt) Get() interface{} {
	return m
}

// Set 更新 Mqtt 配置的字段
func (m *Mqtt) Set(data interface{}) {
	if configData, ok := data.(*Mqtt); ok {
		m.ModuleName = configData.ModuleName
		m.Url = configData.Url
		m.Username = configData.Username
		m.Password = configData.Password
		m.ProtocolVersion = configData.ProtocolVersion
		m.CleanSession = configData.CleanSession
		m.AutoReconnect = configData.AutoReconnect
		m.KeepAlive = configData.KeepAlive
		m.MaxReconnectInterval = configData.MaxReconnectInterval
		m.PingTimeout = configData.PingTimeout
		m.WriteTimeout = configData.WriteTimeout
		m.ConnectTimeout = configData.ConnectTimeout
		m.WillTopic = configData.WillTopic
	}
}

// Validate 验证 Mqtt 配置的有效性
func (m *Mqtt) Validate() error {
	if m.ModuleName == "" {
		return errors.New("module name cannot be empty")
	}
	if m.Url == "" {
		return errors.New("URL cannot be empty")
	}
	if m.Username == "" {
		return errors.New("username cannot be empty")
	}
	if m.Password == "" {
		return errors.New("password cannot be empty")
	}
	if m.ProtocolVersion != 3 && m.ProtocolVersion != 4 {
		return errors.New("protocol version must be 3 or 4")
	}
	if m.KeepAlive <= 0 {
		return errors.New("keep alive must be greater than 0")
	}
	if m.MaxReconnectInterval <= 0 {
		return errors.New("max reconnect interval must be greater than 0")
	}
	if m.PingTimeout <= 0 {
		return errors.New("ping timeout must be greater than 0")
	}
	if m.WriteTimeout <= 0 {
		return errors.New("write timeout must be greater than 0")
	}
	if m.ConnectTimeout <= 0 {
		return errors.New("connect timeout must be greater than 0")
	}
	return nil
}
