/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 14:55:23
 * @FilePath: \go-config\pkg\middleware\middleware.go
 * @Description: 中间件配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package middleware

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-config/pkg/access"
	"github.com/kamalyes/go-config/pkg/alerting"
	"github.com/kamalyes/go-config/pkg/breaker"
	"github.com/kamalyes/go-config/pkg/i18n"
	"github.com/kamalyes/go-config/pkg/logging"
	"github.com/kamalyes/go-config/pkg/metrics"
	"github.com/kamalyes/go-config/pkg/pprof"
	"github.com/kamalyes/go-config/pkg/recovery"
	"github.com/kamalyes/go-config/pkg/requestid"
	"github.com/kamalyes/go-config/pkg/signature"
	"github.com/kamalyes/go-config/pkg/tracing"
)

// Middleware 中间件配置
type Middleware struct {
	ModuleName     string                  `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`             // 模块名称
	Enabled        bool                    `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                        // 是否启用中间件
	Logging        *logging.Logging        `mapstructure:"logging" yaml:"logging" json:"logging"`                        // 日志中间件
	Recovery       *recovery.Recovery      `mapstructure:"recovery" yaml:"recovery" json:"recovery"`                     // 恢复中间件
	Tracing        *tracing.Tracing        `mapstructure:"tracing" yaml:"tracing" json:"tracing"`                        // 追踪中间件
	Metrics        *metrics.Metrics        `mapstructure:"metrics" yaml:"metrics" json:"metrics"`                        // 指标中间件
	RequestID      *requestid.RequestID    `mapstructure:"request-id" yaml:"request-id" json:"requestId"`                // 请求ID中间件
	I18N           *i18n.I18N              `mapstructure:"i18n" yaml:"i18n" json:"i18n"`                                 // 国际化中间件
	PProf          *pprof.PProf            `mapstructure:"pprof" yaml:"pprof" json:"pprof"`                              // PProf中间件
	CircuitBreaker *breaker.CircuitBreaker `mapstructure:"circuit-breaker" yaml:"circuit-breaker" json:"circuitBreaker"` // 断路器配置
	Alerting       *alerting.Alerting      `mapstructure:"alerting" yaml:"alerting" json:"alerting"`                     // 告警配置
	Access         *access.Access          `mapstructure:"access" yaml:"access" json:"access"`                           // 访问记录中间件
	Signature      *signature.Signature    `mapstructure:"signature" yaml:"signature" json:"signature"`                  // 签名验证中间件
}

// Default 创建默认中间件配置
func Default() *Middleware {
	return &Middleware{
		ModuleName:     "middleware",
		Enabled:        true,
		Logging:        logging.Default(),
		Recovery:       recovery.Default(),
		Tracing:        tracing.Default(),
		Metrics:        metrics.Default(),
		RequestID:      requestid.Default(),
		I18N:           i18n.Default(),
		PProf:          pprof.Default(),
		CircuitBreaker: breaker.Default(),
		Alerting:       alerting.Default(),
		Access:         access.Default(),
		Signature:      signature.Default(),
	}
}

// Get 返回配置接口
func (m *Middleware) Get() interface{} {
	return m
}

// Set 设置配置数据
func (m *Middleware) Set(data interface{}) {
	if cfg, ok := data.(*Middleware); ok {
		*m = *cfg
	}
}

// Clone 返回配置的副本
func (m *Middleware) Clone() internal.Configurable {
	clone := &Middleware{
		ModuleName: m.ModuleName,
		Enabled:    m.Enabled,
	}

	if m.Logging != nil {
		clone.Logging = m.Logging.Clone().(*logging.Logging)
	}
	if m.Recovery != nil {
		clone.Recovery = m.Recovery.Clone().(*recovery.Recovery)
	}
	if m.Tracing != nil {
		clone.Tracing = m.Tracing.Clone().(*tracing.Tracing)
	}
	if m.Metrics != nil {
		clone.Metrics = m.Metrics.Clone().(*metrics.Metrics)
	}
	if m.RequestID != nil {
		clone.RequestID = m.RequestID.Clone().(*requestid.RequestID)
	}
	if m.I18N != nil {
		clone.I18N = m.I18N.Clone().(*i18n.I18N)
	}
	if m.PProf != nil {
		clone.PProf = m.PProf.Clone().(*pprof.PProf)
	}
	if m.CircuitBreaker != nil {
		clone.CircuitBreaker = m.CircuitBreaker.Clone().(*breaker.CircuitBreaker)
	}
	if m.Alerting != nil {
		clone.Alerting = m.Alerting.Clone().(*alerting.Alerting)
	}
	if m.Access != nil {
		clone.Access = m.Access.Clone().(*access.Access)
	}
	if m.Signature != nil {
		clone.Signature = m.Signature.Clone().(*signature.Signature)
	}

	return clone
}

// Validate 验证配置
func (m *Middleware) Validate() error {
	return internal.ValidateStruct(m)
}

// WithModuleName 设置模块名称
func (m *Middleware) WithModuleName(moduleName string) *Middleware {
	m.ModuleName = moduleName
	return m
}

// WithEnabled 设置是否启用中间件
func (m *Middleware) WithEnabled(enabled bool) *Middleware {
	m.Enabled = enabled
	return m
}

// WithLogging 设置日志中间件配置
func (m *Middleware) WithLogging(logging *logging.Logging) *Middleware {
	m.Logging = logging
	return m
}

// WithRecovery 设置恢复中间件配置
func (m *Middleware) WithRecovery(recovery *recovery.Recovery) *Middleware {
	m.Recovery = recovery
	return m
}

// WithTracing 设置追踪中间件配置
func (m *Middleware) WithTracing(tracing *tracing.Tracing) *Middleware {
	m.Tracing = tracing
	return m
}

// WithMetrics 设置指标中间件配置
func (m *Middleware) WithMetrics(metrics *metrics.Metrics) *Middleware {
	m.Metrics = metrics
	return m
}

// WithRequestID 设置请求ID中间件配置
func (m *Middleware) WithRequestID(requestID *requestid.RequestID) *Middleware {
	m.RequestID = requestID
	return m
}

// WithI18N 设置国际化中间件配置
func (m *Middleware) WithI18N(i18n *i18n.I18N) *Middleware {
	m.I18N = i18n
	return m
}

// WithPProf 设置PProf中间件配置
func (m *Middleware) WithPProf(pprof *pprof.PProf) *Middleware {
	m.PProf = pprof
	return m
}

// EnableLogging 启用日志中间件
func (m *Middleware) EnableLogging() *Middleware {
	if m.Logging != nil {
		m.Logging.Enable()
	}
	return m
}

// EnableRecovery 启用恢复中间件
func (m *Middleware) EnableRecovery() *Middleware {
	if m.Recovery != nil {
		m.Recovery.Enable()
	}
	return m
}

// EnableTracing 启用追踪中间件
func (m *Middleware) EnableTracing() *Middleware {
	if m.Tracing != nil {
		m.Tracing.Enable()
	}
	return m
}

// EnableMetrics 启用指标中间件
func (m *Middleware) EnableMetrics() *Middleware {
	if m.Metrics != nil {
		m.Metrics.Enable()
	}
	return m
}

// EnableRequestID 启用请求ID中间件
func (m *Middleware) EnableRequestID() *Middleware {
	if m.RequestID != nil {
		m.RequestID.Enable()
	}
	return m
}

// EnableI18N 启用国际化中间件
func (m *Middleware) EnableI18N() *Middleware {
	if m.I18N != nil {
		m.I18N.Enable()
	}
	return m
}

// EnablePProf 启用PProf中间件
func (m *Middleware) EnablePProf() *Middleware {
	if m.PProf != nil {
		m.PProf.Enable()
	}
	return m
}

// Enable 启用中间件
func (m *Middleware) Enable() *Middleware {
	m.Enabled = true
	return m
}

// Disable 禁用中间件
func (m *Middleware) Disable() *Middleware {
	m.Enabled = false
	return m
}

// IsEnabled 检查是否启用
func (m *Middleware) IsEnabled() bool {
	return m.Enabled
}

// WithCircuitBreaker 设置断路器配置
func (m *Middleware) WithCircuitBreaker(circuitBreaker *breaker.CircuitBreaker) *Middleware {
	m.CircuitBreaker = circuitBreaker
	return m
}

// WithAlerting 设置告警配置
func (m *Middleware) WithAlerting(alerting *alerting.Alerting) *Middleware {
	m.Alerting = alerting
	return m
}

// WithAccess 设置访问记录中间件配置
func (m *Middleware) WithAccess(access *access.Access) *Middleware {
	m.Access = access
	return m
}

// WithSignature 设置签名验证中间件配置
func (m *Middleware) WithSignature(signature *signature.Signature) *Middleware {
	m.Signature = signature
	return m
}

// EnableCircuitBreaker 启用断路器
func (m *Middleware) EnableCircuitBreaker() *Middleware {
	if m.CircuitBreaker != nil {
		m.CircuitBreaker.Enable()
	}
	return m
}

// EnableAlerting 启用告警
func (m *Middleware) EnableAlerting() *Middleware {
	if m.Alerting != nil {
		m.Alerting.Enable()
	}
	return m
}

// EnableAccess 启用访问记录
func (m *Middleware) EnableAccess() *Middleware {
	if m.Access != nil {
		m.Access.Enable()
	}
	return m
}

// EnableSignature 启用签名验证
func (m *Middleware) EnableSignature() *Middleware {
	if m.Signature != nil {
		m.Signature.Enable()
	}
	return m
}
