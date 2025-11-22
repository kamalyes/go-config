/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 00:00:00
 * @FilePath: \go-config\pkg\elasticsearch\elasticsearch_test.go
 * @Description: Elasticsearch配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package elasticsearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElasticsearch_Default(t *testing.T) {
	es := Default()

	assert.NotNil(t, es)
	assert.Equal(t, "elasticsearch", es.ModuleName)
	assert.Equal(t, "http://127.0.0.1:9200", es.Endpoint)
	assert.Equal(t, 10, es.HealthCheck)
	assert.False(t, es.Sniff)
	assert.False(t, es.Gzip)
	assert.Equal(t, "10s", es.Timeout)
}

func TestElasticsearch_WithModuleName(t *testing.T) {
	es := Default().WithModuleName("custom_es")
	assert.Equal(t, "custom_es", es.ModuleName)
}

func TestElasticsearch_WithEndpoint(t *testing.T) {
	es := Default().WithEndpoint("http://es-cluster:9200")
	assert.Equal(t, "http://es-cluster:9200", es.Endpoint)
}

func TestElasticsearch_WithHealthCheck(t *testing.T) {
	es := Default().WithHealthCheck(30)
	assert.Equal(t, 30, es.HealthCheck)
}

func TestElasticsearch_WithSniff(t *testing.T) {
	es := Default().WithSniff(true)
	assert.True(t, es.Sniff)
}

func TestElasticsearch_WithGzip(t *testing.T) {
	es := Default().WithGzip(true)
	assert.True(t, es.Gzip)
}

func TestElasticsearch_WithTimeout(t *testing.T) {
	es := Default().WithTimeout("30s")
	assert.Equal(t, "30s", es.Timeout)
}

func TestElasticsearch_Clone(t *testing.T) {
	original := Default()
	original.Endpoint = "http://custom:9200"
	original.HealthCheck = 20
	original.Sniff = true

	cloned := original.Clone().(*Elasticsearch)

	assert.Equal(t, original.Endpoint, cloned.Endpoint)
	assert.Equal(t, original.HealthCheck, cloned.HealthCheck)
	assert.Equal(t, original.Sniff, cloned.Sniff)

	cloned.HealthCheck = 40
	assert.Equal(t, 20, original.HealthCheck)
	assert.Equal(t, 40, cloned.HealthCheck)
}

func TestElasticsearch_Get(t *testing.T) {
	es := Default()
	result := es.Get()

	assert.NotNil(t, result)
	resultES, ok := result.(*Elasticsearch)
	assert.True(t, ok)
	assert.Equal(t, es, resultES)
}

func TestElasticsearch_Set(t *testing.T) {
	es := Default()
	newES := &Elasticsearch{
		ModuleName:  "new_es",
		Endpoint:    "http://new:9200",
		HealthCheck: 25,
		Sniff:       true,
		Gzip:        true,
		Timeout:     "20s",
	}

	es.Set(newES)

	assert.Equal(t, "new_es", es.ModuleName)
	assert.Equal(t, "http://new:9200", es.Endpoint)
	assert.Equal(t, 25, es.HealthCheck)
	assert.True(t, es.Sniff)
	assert.True(t, es.Gzip)
	assert.Equal(t, "20s", es.Timeout)
}

func TestElasticsearch_Validate(t *testing.T) {
	es := Default()
	err := es.Validate()
	assert.NoError(t, err)
}

func TestElasticsearch_ChainedCalls(t *testing.T) {
	es := Default().
		WithModuleName("chained").
		WithEndpoint("http://chained:9200").
		WithHealthCheck(15).
		WithSniff(true).
		WithGzip(true).
		WithTimeout("15s")

	assert.Equal(t, "chained", es.ModuleName)
	assert.Equal(t, "http://chained:9200", es.Endpoint)
	assert.Equal(t, 15, es.HealthCheck)
	assert.True(t, es.Sniff)
	assert.True(t, es.Gzip)
	assert.Equal(t, "15s", es.Timeout)

	err := es.Validate()
	assert.NoError(t, err)
}

func TestNewElasticsearch(t *testing.T) {
	opt := &Elasticsearch{
		ModuleName:  "test",
		Endpoint:    "http://test:9200",
		HealthCheck: 5,
		Sniff:       false,
		Gzip:        false,
		Timeout:     "5s",
	}

	es := NewElasticsearch(opt)
	assert.NotNil(t, es)
	assert.Equal(t, opt, es)
}

func TestDefaultElasticsearch(t *testing.T) {
	es := DefaultElasticsearch()

	assert.Equal(t, "elasticsearch", es.ModuleName)
	assert.Equal(t, "http://127.0.0.1:9200", es.Endpoint)
	assert.Equal(t, 10, es.HealthCheck)
	assert.False(t, es.Sniff)
	assert.False(t, es.Gzip)
	assert.Equal(t, "10s", es.Timeout)
}
