/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:59:11
 * @FilePath: \go-config\pkg\elk\es.go
 * @Description:
 *
 * Copyright (j) 2024 by kamalyes, All Rights Reserved.
 */
package elk

import "github.com/kamalyes/go-config/internal"

// Elasticsearch 结构体
type Elasticsearch struct {
	URL         string `mapstructure:"url"          yaml:"url"          json:"url" validate:"required"`                // Elasticsearch URL
	HealthCheck int    `mapstructure:"health-check" yaml:"health-check" json:"health_check" validate:"required,min=1"` // 健康检查间隔
	Sniff       bool   `mapstructure:"sniff"        yaml:"sniff"        json:"sniff"`                                  // 是否启用嗅探
	Gzip        bool   `mapstructure:"gzip"         yaml:"gzip"         json:"gzip"`                                   // 是否启用 Gzip
	Timeout     string `mapstructure:"timeout"      yaml:"timeout"      json:"timeout" validate:"required"`            // 超时时间
	ModuleName  string `mapstructure:"modulename"   yaml:"modulename"   json:"module_name"`                            // 模块名称
}

// NewElasticsearch 创建一个新的 Elasticsearch 实例
func NewElasticsearch(opt *Elasticsearch) *Elasticsearch {
	var esInstance *Elasticsearch

	internal.LockFunc(func() {
		esInstance = opt
	})
	return esInstance
}

// Clone 返回 Elasticsearch 配置的副本
func (e *Elasticsearch) Clone() internal.Configurable {
	return &Elasticsearch{
		URL:         e.URL,
		HealthCheck: e.HealthCheck,
		Sniff:       e.Sniff,
		Gzip:        e.Gzip,
		Timeout:     e.Timeout,
		ModuleName:  e.ModuleName,
	}
}

// Get 返回 Elasticsearch 配置的所有字段
func (e *Elasticsearch) Get() interface{} {
	return e
}

// Set 更新 Elasticsearch 配置的字段
func (e *Elasticsearch) Set(data interface{}) {
	if configData, ok := data.(*Elasticsearch); ok {
		e.URL = configData.URL
		e.HealthCheck = configData.HealthCheck
		e.Sniff = configData.Sniff
		e.Gzip = configData.Gzip
		e.Timeout = configData.Timeout
		e.ModuleName = configData.ModuleName
	}
}

// Validate 检查 Elasticsearch 配置的有效性
func (e *Elasticsearch) Validate() error {
	return internal.ValidateStruct(e)
}
