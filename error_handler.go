/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-12 15:40:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 15:40:00
 * @FilePath: \go-config\error_handler.go
 * @Description: 统一错误处理组件，提供分类错误处理和日志记录
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package goconfig

import (
	"context"
	"errors"
	"fmt"
	"github.com/kamalyes/go-logger"
	"runtime"
	"strings"
	"sync"
	"time"
)

// ErrorType 错误类型枚举
type ErrorType string

const (
	ErrorTypeConfig        ErrorType = "config"        // 配置相关错误
	ErrorTypeFileSystem    ErrorType = "filesystem"    // 文件系统错误
	ErrorTypeNetwork       ErrorType = "network"       // 网络错误
	ErrorTypeValidation    ErrorType = "validation"    // 验证错误
	ErrorTypeSerialization ErrorType = "serialization" // 序列化错误
	ErrorTypeCallback      ErrorType = "callback"      // 回调执行错误
	ErrorTypePermission    ErrorType = "permission"    // 权限错误
	ErrorTypeTimeout       ErrorType = "timeout"       // 超时错误
	ErrorTypeInternal      ErrorType = "internal"      // 内部错误
	ErrorTypeExternal      ErrorType = "external"      // 外部错误
)

// ErrorSeverity 错误严重程度
type ErrorSeverity int

const (
	SeverityDebug    ErrorSeverity = iota // 调试信息
	SeverityInfo                          // 信息
	SeverityWarn                          // 警告
	SeverityError                         // 错误
	SeverityCritical                      // 严重错误
	SeverityFatal                         // 致命错误
)

// String 实现 Stringer 接口
func (s ErrorSeverity) String() string {
	switch s {
	case SeverityDebug:
		return "DEBUG"
	case SeverityInfo:
		return "INFO"
	case SeverityWarn:
		return "WARN"
	case SeverityError:
		return "ERROR"
	case SeverityCritical:
		return "CRITICAL"
	case SeverityFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// ConfigError 配置错误结构
type ConfigError struct {
	Type       ErrorType              `json:"type"`        // 错误类型
	Severity   ErrorSeverity          `json:"severity"`    // 严重程度
	Message    string                 `json:"message"`     // 错误消息
	Source     string                 `json:"source"`      // 错误来源
	Timestamp  time.Time              `json:"timestamp"`   // 错误时间
	StackTrace string                 `json:"stack_trace"` // 堆栈跟踪
	Context    map[string]interface{} `json:"context"`     // 错误上下文
	Cause      error                  `json:"cause"`       // 原始错误
	Retryable  bool                   `json:"retryable"`   // 是否可重试
	Code       string                 `json:"code"`        // 错误代码
}

// Error 实现 error 接口
func (e *ConfigError) Error() string {
	return fmt.Sprintf("[%s:%s] %s", e.Type, e.Severity, e.Message)
}

// Unwrap 实现 errors.Unwrap 接口
func (e *ConfigError) Unwrap() error {
	return e.Cause
}

// Is 实现 errors.Is 接口
func (e *ConfigError) Is(target error) bool {
	if target == nil {
		return false
	}

	if configErr, ok := target.(*ConfigError); ok {
		return e.Type == configErr.Type && e.Code == configErr.Code
	}

	return errors.Is(e.Cause, target)
}

// WithContext 添加上下文信息
func (e *ConfigError) WithContext(key string, value interface{}) *ConfigError {
	if e.Context == nil {
		e.Context = make(map[string]interface{})
	}
	e.Context[key] = value
	return e
}

// ErrorHandler 错误处理器接口
type ErrorHandler interface {
	// HandleError 处理错误
	HandleError(ctx context.Context, err error) *ConfigError

	// ClassifyError 分类错误
	ClassifyError(err error) (*ConfigError, ErrorType, ErrorSeverity)

	// RegisterErrorCallback 注册错误回调
	RegisterErrorCallback(callback ErrorCallback, filter ErrorFilter) error

	// UnregisterErrorCallback 注销错误回调
	UnregisterErrorCallback(id string) error

	// GetErrorStats 获取错误统计
	GetErrorStats() *ErrorStats

	// ClearErrorStats 清除错误统计
	ClearErrorStats()
}

// ErrorCallback 错误回调函数类型
type ErrorCallback func(ctx context.Context, configErr *ConfigError) error

// ErrorFilter 错误过滤器
type ErrorFilter struct {
	ID         string          `json:"id"`         // 回调ID
	Types      []ErrorType     `json:"types"`      // 监听的错误类型
	Severities []ErrorSeverity `json:"severities"` // 监听的严重程度
	Sources    []string        `json:"sources"`    // 监听的错误来源
}

// ErrorStats 错误统计
type ErrorStats struct {
	TotalErrors      int                   `json:"total_errors"`       // 总错误数
	ErrorsByType     map[ErrorType]int     `json:"errors_by_type"`     // 按类型分类
	ErrorsBySeverity map[ErrorSeverity]int `json:"errors_by_severity"` // 按严重程度分类
	ErrorsBySource   map[string]int        `json:"errors_by_source"`   // 按来源分类
	LastError        *ConfigError          `json:"last_error"`         // 最后一个错误
	FirstSeen        time.Time             `json:"first_seen"`         // 第一次看到错误的时间
	LastSeen         time.Time             `json:"last_seen"`          // 最后一次看到错误的时间
	RecentErrors     []*ConfigError        `json:"recent_errors"`      // 最近的错误（最多保留100个）
}

// errorCallback 内部错误回调信息
type errorCallback struct {
	Callback ErrorCallback
	Filter   ErrorFilter
}

// ConfigErrorHandler 配置错误处理器实现
type ConfigErrorHandler struct {
	mu        sync.RWMutex
	callbacks map[string]*errorCallback
	stats     *ErrorStats
}

// NewErrorHandler 创建新的错误处理器
func NewErrorHandler() ErrorHandler {
	return &ConfigErrorHandler{
		callbacks: make(map[string]*errorCallback),
		stats: &ErrorStats{
			ErrorsByType:     make(map[ErrorType]int),
			ErrorsBySeverity: make(map[ErrorSeverity]int),
			ErrorsBySource:   make(map[string]int),
			RecentErrors:     make([]*ConfigError, 0, 100),
		},
	}
}

// HandleError 处理错误
func (h *ConfigErrorHandler) HandleError(ctx context.Context, err error) *ConfigError {
	if err == nil {
		return nil
	}

	// 如果已经是 ConfigError，直接使用
	var configErr *ConfigError
	if errors.As(err, &configErr) {
		configErr.Timestamp = time.Now()
	} else {
		// 分类和包装错误
		configErr, _, _ = h.ClassifyError(err)
	}

	// 更新统计信息
	h.updateStats(configErr)

	// 记录日志
	h.logError(configErr)

	// 触发回调
	h.triggerCallbacks(ctx, configErr)

	return configErr
}

// ClassifyError 分类错误
func (h *ConfigErrorHandler) ClassifyError(err error) (*ConfigError, ErrorType, ErrorSeverity) {
	if err == nil {
		return nil, "", SeverityInfo
	}

	message := err.Error()
	errorType, severity := h.detectErrorType(message, err)

	configErr := &ConfigError{
		Type:       errorType,
		Severity:   severity,
		Message:    message,
		Source:     h.detectSource(),
		Timestamp:  time.Now(),
		StackTrace: h.captureStackTrace(),
		Context:    make(map[string]interface{}),
		Cause:      err,
		Retryable:  h.isRetryable(errorType),
		Code:       h.generateErrorCode(errorType, severity),
	}

	return configErr, errorType, severity
}

// detectErrorType 检测错误类型和严重程度
func (h *ConfigErrorHandler) detectErrorType(message string, err error) (ErrorType, ErrorSeverity) {
	lowerMessage := strings.ToLower(message)

	// 检查特定错误类型
	switch {
	case strings.Contains(lowerMessage, "config") || strings.Contains(lowerMessage, "配置"):
		return ErrorTypeConfig, SeverityError
	case strings.Contains(lowerMessage, "file") || strings.Contains(lowerMessage, "directory") || strings.Contains(lowerMessage, "文件"):
		return ErrorTypeFileSystem, SeverityError
	case strings.Contains(lowerMessage, "network") || strings.Contains(lowerMessage, "connection") || strings.Contains(lowerMessage, "网络"):
		return ErrorTypeNetwork, SeverityWarn
	case strings.Contains(lowerMessage, "validation") || strings.Contains(lowerMessage, "invalid") || strings.Contains(lowerMessage, "验证"):
		return ErrorTypeValidation, SeverityWarn
	case strings.Contains(lowerMessage, "json") || strings.Contains(lowerMessage, "yaml") || strings.Contains(lowerMessage, "unmarshal"):
		return ErrorTypeSerialization, SeverityError
	case strings.Contains(lowerMessage, "callback") || strings.Contains(lowerMessage, "回调"):
		return ErrorTypeCallback, SeverityWarn
	case strings.Contains(lowerMessage, "permission") || strings.Contains(lowerMessage, "access denied") || strings.Contains(lowerMessage, "权限"):
		return ErrorTypePermission, SeverityError
	case strings.Contains(lowerMessage, "timeout") || strings.Contains(lowerMessage, "超时"):
		return ErrorTypeTimeout, SeverityWarn
	case strings.Contains(lowerMessage, "panic") || strings.Contains(lowerMessage, "fatal"):
		return ErrorTypeInternal, SeverityFatal
	default:
		return ErrorTypeInternal, SeverityError
	}
}

// detectSource 检测错误来源
func (h *ConfigErrorHandler) detectSource() string {
	// 获取调用栈信息
	if pc, file, line, ok := runtime.Caller(3); ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			return fmt.Sprintf("%s:%d (%s)", file, line, fn.Name())
		}
		return fmt.Sprintf("%s:%d", file, line)
	}
	return "unknown"
}

// captureStackTrace 捕获堆栈跟踪
func (h *ConfigErrorHandler) captureStackTrace() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

// isRetryable 判断错误是否可重试
func (h *ConfigErrorHandler) isRetryable(errorType ErrorType) bool {
	switch errorType {
	case ErrorTypeNetwork, ErrorTypeTimeout, ErrorTypeExternal:
		return true
	case ErrorTypePermission, ErrorTypeValidation, ErrorTypeSerialization:
		return false
	default:
		return false
	}
}

// generateErrorCode 生成错误代码
func (h *ConfigErrorHandler) generateErrorCode(errorType ErrorType, severity ErrorSeverity) string {
	return fmt.Sprintf("%s_%s_%d", strings.ToUpper(string(errorType)), severity.String(), time.Now().Unix()%10000)
}

// updateStats 更新统计信息
func (h *ConfigErrorHandler) updateStats(configErr *ConfigError) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.stats.TotalErrors++
	h.stats.ErrorsByType[configErr.Type]++
	h.stats.ErrorsBySeverity[configErr.Severity]++
	h.stats.ErrorsBySource[configErr.Source]++
	h.stats.LastError = configErr
	h.stats.LastSeen = configErr.Timestamp

	if h.stats.FirstSeen.IsZero() {
		h.stats.FirstSeen = configErr.Timestamp
	}

	// 保持最近100个错误
	h.stats.RecentErrors = append(h.stats.RecentErrors, configErr)
	if len(h.stats.RecentErrors) > 100 {
		h.stats.RecentErrors = h.stats.RecentErrors[1:]
	}
}

// logError 记录错误日志
func (h *ConfigErrorHandler) logError(configErr *ConfigError) {
	switch configErr.Severity {
	case SeverityFatal:
		logger.GetGlobalLogger().Fatal("💀 %s [%s]: %s", configErr.Code, configErr.Source, configErr.Message)
	case SeverityCritical:
		logger.GetGlobalLogger().Error("🔥 %s [%s]: %s", configErr.Code, configErr.Source, configErr.Message)
	case SeverityError:
		logger.GetGlobalLogger().Error("❌ %s [%s]: %s", configErr.Code, configErr.Source, configErr.Message)
	case SeverityWarn:
		logger.GetGlobalLogger().Warn("⚠️ %s [%s]: %s", configErr.Code, configErr.Source, configErr.Message)
	case SeverityInfo:
		logger.GetGlobalLogger().Info("ℹ️ %s [%s]: %s", configErr.Code, configErr.Source, configErr.Message)
	case SeverityDebug:
		logger.GetGlobalLogger().Debug("🐛 %s [%s]: %s", configErr.Code, configErr.Source, configErr.Message)
	}
}

// triggerCallbacks 触发错误回调
func (h *ConfigErrorHandler) triggerCallbacks(ctx context.Context, configErr *ConfigError) {
	h.mu.RLock()
	callbacks := make([]*errorCallback, 0)
	for _, cb := range h.callbacks {
		if h.shouldTriggerCallback(cb, configErr) {
			callbacks = append(callbacks, cb)
		}
	}
	h.mu.RUnlock()

	// 异步执行回调
	for _, cb := range callbacks {
		go func(callback *errorCallback) {
			defer func() {
				if r := recover(); r != nil {
					logger.GetGlobalLogger().Error("错误回调执行panic: %v", r)
				}
			}()

			if err := callback.Callback(ctx, configErr); err != nil {
				logger.GetGlobalLogger().Error("错误回调执行失败: %v", err)
			}
		}(cb)
	}
}

// shouldTriggerCallback 判断是否应该触发回调
func (h *ConfigErrorHandler) shouldTriggerCallback(cb *errorCallback, configErr *ConfigError) bool {
	// 检查错误类型
	if len(cb.Filter.Types) > 0 {
		found := false
		for _, t := range cb.Filter.Types {
			if t == configErr.Type {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// 检查严重程度
	if len(cb.Filter.Severities) > 0 {
		found := false
		for _, s := range cb.Filter.Severities {
			if s == configErr.Severity {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// 检查来源
	if len(cb.Filter.Sources) > 0 {
		found := false
		for _, source := range cb.Filter.Sources {
			if strings.Contains(configErr.Source, source) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// RegisterErrorCallback 注册错误回调
func (h *ConfigErrorHandler) RegisterErrorCallback(callback ErrorCallback, filter ErrorFilter) error {
	if callback == nil {
		return fmt.Errorf("回调函数不能为空")
	}

	if filter.ID == "" {
		return fmt.Errorf("回调ID不能为空")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.callbacks[filter.ID]; exists {
		return fmt.Errorf("回调ID %s 已存在", filter.ID)
	}

	h.callbacks[filter.ID] = &errorCallback{
		Callback: callback,
		Filter:   filter,
	}

	logger.GetGlobalLogger().Debug("✅ 错误回调已注册: %s", filter.ID)
	return nil
}

// UnregisterErrorCallback 注销错误回调
func (h *ConfigErrorHandler) UnregisterErrorCallback(id string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.callbacks[id]; !exists {
		return fmt.Errorf("回调ID %s 不存在", id)
	}

	delete(h.callbacks, id)
	logger.GetGlobalLogger().Debug("🗑️ 错误回调已注销: %s", id)
	return nil
}

// GetErrorStats 获取错误统计
func (h *ConfigErrorHandler) GetErrorStats() *ErrorStats {
	h.mu.RLock()
	defer h.mu.RUnlock()

	// 深度复制统计信息
	statsCopy := &ErrorStats{
		TotalErrors:      h.stats.TotalErrors,
		ErrorsByType:     make(map[ErrorType]int),
		ErrorsBySeverity: make(map[ErrorSeverity]int),
		ErrorsBySource:   make(map[string]int),
		LastError:        h.stats.LastError,
		FirstSeen:        h.stats.FirstSeen,
		LastSeen:         h.stats.LastSeen,
		RecentErrors:     make([]*ConfigError, len(h.stats.RecentErrors)),
	}

	for k, v := range h.stats.ErrorsByType {
		statsCopy.ErrorsByType[k] = v
	}
	for k, v := range h.stats.ErrorsBySeverity {
		statsCopy.ErrorsBySeverity[k] = v
	}
	for k, v := range h.stats.ErrorsBySource {
		statsCopy.ErrorsBySource[k] = v
	}
	copy(statsCopy.RecentErrors, h.stats.RecentErrors)

	return statsCopy
}

// ClearErrorStats 清除错误统计
func (h *ConfigErrorHandler) ClearErrorStats() {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.stats = &ErrorStats{
		ErrorsByType:     make(map[ErrorType]int),
		ErrorsBySeverity: make(map[ErrorSeverity]int),
		ErrorsBySource:   make(map[string]int),
		RecentErrors:     make([]*ConfigError, 0, 100),
	}

	logger.GetGlobalLogger().Info("🧹 错误统计已清除")
}

// 全局错误处理器实例
var globalErrorHandler ErrorHandler

// GetGlobalErrorHandler 获取全局错误处理器
func GetGlobalErrorHandler() ErrorHandler {
	if globalErrorHandler == nil {
		globalErrorHandler = NewErrorHandler()
	}
	return globalErrorHandler
}

// SetGlobalErrorHandler 设置全局错误处理器
func SetGlobalErrorHandler(handler ErrorHandler) {
	globalErrorHandler = handler
}

// HandleError 全局错误处理函数
func HandleError(ctx context.Context, err error) *ConfigError {
	return GetGlobalErrorHandler().HandleError(ctx, err)
}

// NewConfigError 创建新的配置错误
func NewConfigError(errorType ErrorType, severity ErrorSeverity, message string) *ConfigError {
	return &ConfigError{
		Type:      errorType,
		Severity:  severity,
		Message:   message,
		Timestamp: time.Now(),
		Context:   make(map[string]interface{}),
		Retryable: false,
		Code:      fmt.Sprintf("%s_%s_%d", strings.ToUpper(string(errorType)), severity.String(), time.Now().Unix()%10000),
	}
}
