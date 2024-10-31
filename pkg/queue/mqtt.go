/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 10:11:00
 * @FilePath: \go-config\pkg\queue\mqtt.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package queue

import (
	"github.com/kamalyes/go-config/internal"
)

// Mqtt 结构体用于配置 MQTT 相关参数
type Mqtt struct {
	ModuleName           string `mapstructure:"modulename"              yaml:"modulename"      json:"module_name"      validate:"required"`            // 模块名称
	Url                  string `mapstructure:"url"                      yaml:"url"             json:"url"             validate:"required,url"`        // Mqtt 代理服务器的 IP 和端口
	ProtocolVersion      uint   `mapstructure:"protocol-version"         yaml:"protocol-ver"    json:"protocol_version" validate:"required"`           // Mqtt 协议版本号，4 是 3.1.1，3 是 3.1
	KeepAlive            int    `mapstructure:"keep-alive"               yaml:"keep-alive"      json:"keep_alive"      validate:"required,min=1"`      // 保活时间间隔，最小值为 1 秒
	MaxReconnectInterval int    `mapstructure:"max-reconnect-interval"   yaml:"max-reconnect-interval" json:"max_reconnect_interval" validate:"min=1"` // 最大连接间隔时间，单位：秒，最小值为 1 秒
	PingTimeout          int    `mapstructure:"ping-timeout"             yaml:"ping-timeout"    json:"ping_timeout"      validate:"min=1"`             // ping 超时时间，单位：秒，最小值为 1 秒
	WriteTimeout         int    `mapstructure:"write-timeout"            yaml:"write-timeout"   json:"write_timeout"     validate:"min=1"`             // 写超时时间，单位：秒，最小值为 1 秒
	ConnectTimeout       int    `mapstructure:"connect-timeout"          yaml:"connect-timeout" json:"connect_timeout"   validate:"min=1"`             // 连接超时时间，单位：秒，最小值为 1 秒
	Username             string `mapstructure:"username"                 yaml:"username"        json:"username"`                                       // 连接到代理服务器的用户名
	Password             string `mapstructure:"password"                 yaml:"password"        json:"password"`                                       // 密码
	CleanSession         bool   `mapstructure:"clean-session"            yaml:"clean-session"   json:"clean_session"`                                  // 设置客户端掉线服务端是否清除 session
	AutoReconnect        bool   `mapstructure:"auto-reconnect"           yaml:"auto-reconnect"  json:"auto_reconnect"`                                 // 断开后是否重新连接
	WillTopic            string `mapstructure:"will-topic"               yaml:"will-topic"      json:"will_topic"`                                     // 遗言发送的 topic
}

// NewMqtt 创建一个新的 Mqtt 实例
func NewMqtt(opt *Mqtt) *Mqtt {
	var mqttInstance *Mqtt

	internal.LockFunc(func() {
		mqttInstance = opt
	})
	return mqttInstance
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
	return internal.ValidateStruct(m)
}
