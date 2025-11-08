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
	Endpoint    string `mapstructure:"endpoint"     yaml:"endpoint"     json:"endpoint" validate:"required"`           // Elasticsearch 端点地址
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
		Endpoint:    e.Endpoint,
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
		e.Endpoint = configData.Endpoint
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

// DefaultElasticsearch 返回默认Elasticsearch配置
func DefaultElasticsearch() Elasticsearch {
	return Elasticsearch{
		ModuleName:  "elasticsearch",
		Endpoint:    "http://127.0.0.1:9200",
		HealthCheck: 10,
		Sniff:       false,
		Gzip:        false,
		Timeout:     "10s",
	}
}

// Default 返回默认Elasticsearch配置的指针，支持链式调用
func Default() *Elasticsearch {
	config := DefaultElasticsearch()
	return &config
}

// WithModuleName 设置模块名称
func (e *Elasticsearch) WithModuleName(moduleName string) *Elasticsearch {
	e.ModuleName = moduleName
	return e
}

// WithEndpoint 设置Elasticsearch端点地址
func (e *Elasticsearch) WithEndpoint(endpoint string) *Elasticsearch {
	e.Endpoint = endpoint
	return e
}

// WithHealthCheck 设置健康检查间隔
func (e *Elasticsearch) WithHealthCheck(healthCheck int) *Elasticsearch {
	e.HealthCheck = healthCheck
	return e
}

// WithSniff 设置是否启用嗅探
func (e *Elasticsearch) WithSniff(sniff bool) *Elasticsearch {
	e.Sniff = sniff
	return e
}

// WithGzip 设置是否启用Gzip
func (e *Elasticsearch) WithGzip(gzip bool) *Elasticsearch {
	e.Gzip = gzip
	return e
}

// WithTimeout 设置超时时间
func (e *Elasticsearch) WithTimeout(timeout string) *Elasticsearch {
	e.Timeout = timeout
	return e
}
