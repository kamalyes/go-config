/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-04-23 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-04-23 00:00:00
 * @FilePath: \go-config\pkg\queue\nats.go
 * @Description: NATS 消息队列配置
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package queue

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-toolbox/pkg/syncx"
)

// Nats 结构体用于配置 NATS 相关参数
//
// NATS 是一款高性能、轻量级的云原生消息系统，常用于：
//   - 分布式服务间低延迟消息广播（如 Casbin 策略变更通知）
//   - 事件驱动架构中的发布/订阅
//   - 启用 JetStream 后支持消息持久化、重放
type Nats struct {
	URL            string `mapstructure:"url" yaml:"url" json:"url" validate:"required"`                                                    // NATS 服务器地址，如 nats://127.0.0.1:4222
	Name           string `mapstructure:"name" yaml:"name" json:"name"`                                                                     // 客户端名称，用于服务端识别
	Username       string `mapstructure:"username" yaml:"username" json:"username"`                                                         // 用户名（可选）
	Password       string `mapstructure:"password" yaml:"password" json:"password"`                                                         // 密码（可选）
	Token          string `mapstructure:"token" yaml:"token" json:"token"`                                                                  // Token 鉴权（可选）
	JetStream      bool   `mapstructure:"jet-stream" yaml:"jet-stream" json:"jetStream"`                                                    // 是否启用 JetStream 持久化
	StreamName     string `mapstructure:"stream-name" yaml:"stream-name" json:"streamName"`                                                 // JetStream Stream 名称（启用 JetStream 时生效）
	ConnectTimeout int    `mapstructure:"connect-timeout" yaml:"connect-timeout" json:"connectTimeout" validate:"min=1"`                    // 连接超时时间（秒）
	ReconnectWait  int    `mapstructure:"reconnect-wait" yaml:"reconnect-wait" json:"reconnectWait" validate:"min=1"`                       // 重连等待时间（秒）
	MaxReconnects  int    `mapstructure:"max-reconnects" yaml:"max-reconnects" json:"maxReconnects"`                                        // 最大重连次数，-1 表示无限
	ChannelPrefix  string `mapstructure:"channel-prefix" yaml:"channel-prefix" json:"channelPrefix"`                                        // Subject 前缀（上层业务可基于此构造租户级 Subject）
	Source         string `mapstructure:"source" yaml:"source" json:"source"`                                                               // 当前节点标识（为空时由业务方自动生成），用于消息去重与来源过滤
	ModuleName     string `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`                                                 // 模块名称
}

// NewNats 创建一个新的 Nats 实例
func NewNats(opt *Nats) *Nats {
	var natsInstance *Nats

	internal.LockFunc(func() {
		natsInstance = opt
	})
	return natsInstance
}

// Clone 返回 Nats 配置的副本
func (n *Nats) Clone() internal.Configurable {
	var cloned Nats
	if err := syncx.DeepCopy(&cloned, n); err != nil {
		return &Nats{}
	}
	return &cloned
}

// Get 返回 Nats 配置
func (n *Nats) Get() interface{} {
	return n
}

// Set 更新 Nats 配置的字段
func (n *Nats) Set(data interface{}) {
	if configData, ok := data.(*Nats); ok {
		n.ModuleName = configData.ModuleName
		n.URL = configData.URL
		n.Name = configData.Name
		n.Username = configData.Username
		n.Password = configData.Password
		n.Token = configData.Token
		n.JetStream = configData.JetStream
		n.StreamName = configData.StreamName
		n.ConnectTimeout = configData.ConnectTimeout
		n.ReconnectWait = configData.ReconnectWait
		n.MaxReconnects = configData.MaxReconnects
		n.ChannelPrefix = configData.ChannelPrefix
		n.Source = configData.Source
	}
}

// Validate 验证 Nats 配置的有效性
func (n *Nats) Validate() error {
	return internal.ValidateStruct(n)
}

// DefaultNats 返回默认 NATS 配置
func DefaultNats() Nats {
	return Nats{
		ModuleName:     "nats",
		URL:            "nats://127.0.0.1:4222",
		Name:           "go-config-nats-client",
		JetStream:      false,
		StreamName:     "",
		ConnectTimeout: 10, // 10 秒
		ReconnectWait:  2,  // 2 秒
		MaxReconnects:  10,
		ChannelPrefix:  "",
		Source:         "",
	}
}

// DefaultNatsPtr 返回默认 NATS 配置的指针，支持链式调用
func DefaultNatsPtr() *Nats {
	config := DefaultNats()
	return &config
}

// WithModuleName 设置模块名称
func (n *Nats) WithModuleName(moduleName string) *Nats {
	n.ModuleName = moduleName
	return n
}

// WithURL 设置 NATS 服务器地址
func (n *Nats) WithURL(url string) *Nats {
	n.URL = url
	return n
}

// WithName 设置客户端名称
func (n *Nats) WithName(name string) *Nats {
	n.Name = name
	return n
}

// WithUsername 设置用户名
func (n *Nats) WithUsername(username string) *Nats {
	n.Username = username
	return n
}

// WithPassword 设置密码
func (n *Nats) WithPassword(password string) *Nats {
	n.Password = password
	return n
}

// WithToken 设置 Token
func (n *Nats) WithToken(token string) *Nats {
	n.Token = token
	return n
}

// WithJetStream 设置是否启用 JetStream
func (n *Nats) WithJetStream(enabled bool) *Nats {
	n.JetStream = enabled
	return n
}

// WithStreamName 设置 JetStream Stream 名称
func (n *Nats) WithStreamName(streamName string) *Nats {
	n.StreamName = streamName
	return n
}

// WithConnectTimeout 设置连接超时时间（秒）
func (n *Nats) WithConnectTimeout(seconds int) *Nats {
	n.ConnectTimeout = seconds
	return n
}

// WithReconnectWait 设置重连等待时间（秒）
func (n *Nats) WithReconnectWait(seconds int) *Nats {
	n.ReconnectWait = seconds
	return n
}

// WithMaxReconnects 设置最大重连次数
func (n *Nats) WithMaxReconnects(max int) *Nats {
	n.MaxReconnects = max
	return n
}

// WithChannelPrefix 设置 Subject 前缀
func (n *Nats) WithChannelPrefix(prefix string) *Nats {
	n.ChannelPrefix = prefix
	return n
}

// WithSource 设置当前节点标识
func (n *Nats) WithSource(source string) *Nats {
	n.Source = source
	return n
}
