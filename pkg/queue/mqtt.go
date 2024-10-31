/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:20:22
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

// Mqtt 结构体用于配置 MQTT 相关参数
type Mqtt struct {
	ModuleName           string `mapstructure:"MODULE_NAME"              yaml:"modulename"`             // 模块名称
	Url                  string `mapstructure:"URL"                      yaml:"url"`                    // Mqtt 代理服务器的 IP 和端口
	Username             string `mapstructure:"USERNAME"                 yaml:"username"`               // 连接到代理服务器的用户名
	Password             string `mapstructure:"PASSWORD"                 yaml:"password"`               // 密码
	ProtocolVersion      uint   `mapstructure:"PROTOCOL_VERSION"         yaml:"protocol-ver"`           // Mqtt 协议版本号，4 是 3.1.1，3 是 3.1
	CleanSession         bool   `mapstructure:"CLEAN_SESSION"            yaml:"clean-session"`          // 设置客户端掉线服务端是否清除 session
	AutoReconnect        bool   `mapstructure:"AUTO_RECONNECT"           yaml:"auto-reconnect"`         // 断开后是否重新连接
	KeepAlive            int    `mapstructure:"KEEP_ALIVE"               yaml:"keep-alive"`             // 保活时间间隔
	MaxReconnectInterval int    `mapstructure:"MAX_RECONNECT_INTERVAL"   yaml:"max-reconnect-interval"` // 最大连接间隔时间，单位：秒
	PingTimeout          int    `mapstructure:"PING_TIMEOUT"             yaml:"ping-timeout"`           // ping 超时时间，单位：秒
	WriteTimeout         int    `mapstructure:"WRITE_TIMEOUT"            yaml:"write-timeout"`          // 写超时时间，单位：秒
	ConnectTimeout       int    `mapstructure:"CONNECT_TIMEOUT"          yaml:"connect-timeout"`        // 连接超时时间，单位：秒
	WillTopic            string `mapstructure:"WILL_TOPIC"               yaml:"will-topic"`             // 遗言发送的 topic
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
