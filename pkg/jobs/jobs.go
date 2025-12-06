/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-11 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-11 13:30:55
 * @FilePath: \go-config\pkg\jobs\jobs.go
 * @Description: Job调度配置
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package jobs

import (
	"fmt"
	"time"

	"github.com/kamalyes/go-config/internal"
)

// Jobs Job调度配置
type Jobs struct {
	// 全局配置
	Enabled          bool               `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                              // 是否启用Job管理器
	TimeZone         string             `mapstructure:"timezone" yaml:"timezone" json:"timezone"`                           // 时区配置(例如: Asia/Shanghai)
	GracefulShutdown int                `mapstructure:"graceful-shutdown" yaml:"graceful-shutdown" json:"gracefulShutdown"` // 优雅关闭超时时间(秒)
	MaxRetries       int                `mapstructure:"max-retries" yaml:"max-retries" json:"maxRetries"`                   // 失败最大重试次数
	RetryInterval    int                `mapstructure:"retry-interval" yaml:"retry-interval" json:"retryInterval"`          // 重试间隔(秒)
	Tasks            map[string]TaskCfg `mapstructure:"tasks" yaml:"tasks" json:"tasks"`                                    // 任务配置
}

// TaskCfg 任务配置
type TaskCfg struct {
	Enabled        bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                        // 是否启用
	CronSpec       string `mapstructure:"cron-spec" yaml:"cron-spec" json:"cronSpec"`                   // Cron表达式
	ImmediateStart bool   `mapstructure:"immediate-start" yaml:"immediate-start" json:"immediateStart"` // 启动时立即执行一次
	Timeout        int    `mapstructure:"timeout" yaml:"timeout" json:"timeout"`                        // 任务超时时间(秒)，0表示无限制
	OverlapPrevent bool   `mapstructure:"overlap-prevent" yaml:"overlap-prevent" json:"overlapPrevent"` // 是否阻止任务重叠执行
	Description    string `mapstructure:"description" yaml:"description" json:"description"`            // 任务描述
}

// Default 默认Job配置
func Default() *Jobs {
	return &Jobs{
		Enabled:          true,
		TimeZone:         "Asia/Shanghai",
		GracefulShutdown: 30,
		MaxRetries:       3,
		RetryInterval:    5,
		Tasks: map[string]TaskCfg{
			// 清理Job - 每5分钟执行
			"cleanup": {
				Enabled:        true,
				CronSpec:       "0 */5 * * * *",
				ImmediateStart: false,
				Timeout:        300,
				OverlapPrevent: true,
				Description:    "定期清理过期数据和无效缓存",
			},
		},
	}
}

// Validate 验证配置
func (c *Jobs) Validate() error {
	if !c.Enabled {
		return nil // 如果未启用，跳过验证
	}

	// 验证时区
	if c.TimeZone == "" {
		return fmt.Errorf("timezone不能为空")
	}
	if _, err := time.LoadLocation(c.TimeZone); err != nil {
		return fmt.Errorf("无效的时区: %s, 错误: %w", c.TimeZone, err)
	}

	// 验证优雅关闭超时
	if c.GracefulShutdown <= 0 {
		return fmt.Errorf("graceful_shutdown必须大于0")
	}

	// 验证任务配置
	for name, task := range c.Tasks {
		if err := task.Validate(name); err != nil {
			return err
		}
	}

	return nil
}

// Validate 验证任务配置
func (t *TaskCfg) Validate(name string) error {
	if !t.Enabled {
		return nil // 如果未启用，跳过验证
	}

	if t.CronSpec == "" {
		return fmt.Errorf("任务[%s]的cron_spec不能为空", name)
	}

	if t.Timeout < 0 {
		return fmt.Errorf("任务[%s]的timeout不能小于0", name)
	}

	return nil
}

// GetTimeZoneLocation 获取时区Location
func (c *Jobs) GetTimeZoneLocation() (*time.Location, error) {
	return time.LoadLocation(c.TimeZone)
}

// GetTaskConfig 获取指定任务的配置
func (c *Jobs) GetTaskConfig(name string) (TaskCfg, bool) {
	task, exists := c.Tasks[name]
	return task, exists
}

// IsTaskEnabled 检查任务是否启用
func (c *Jobs) IsTaskEnabled(name string) bool {
	if !c.Enabled {
		return false
	}
	task, exists := c.Tasks[name]
	return exists && task.Enabled
}

// Clone 返回配置的副本
func (c *Jobs) Clone() internal.Configurable {
	// 深拷贝 Tasks map
	tasks := make(map[string]TaskCfg, len(c.Tasks))
	for k, v := range c.Tasks {
		tasks[k] = v
	}

	return &Jobs{
		Enabled:          c.Enabled,
		TimeZone:         c.TimeZone,
		GracefulShutdown: c.GracefulShutdown,
		MaxRetries:       c.MaxRetries,
		RetryInterval:    c.RetryInterval,
		Tasks:            tasks,
	}
}

// Get 返回配置接口
func (c *Jobs) Get() interface{} {
	return c
}

// Set 设置配置数据
func (c *Jobs) Set(data interface{}) {
	if cfg, ok := data.(*Jobs); ok {
		*c = *cfg
	}
}

// Enable 启用Job管理器
func (c *Jobs) Enable() *Jobs {
	c.Enabled = true
	return c
}

// Disable 禁用Job管理器
func (c *Jobs) Disable() *Jobs {
	c.Enabled = false
	return c
}

// WithTimeZone 设置时区
func (c *Jobs) WithTimeZone(tz string) *Jobs {
	c.TimeZone = tz
	return c
}

// WithGracefulShutdown 设置优雅关闭超时
func (c *Jobs) WithGracefulShutdown(seconds int) *Jobs {
	c.GracefulShutdown = seconds
	return c
}

// WithMaxRetries 设置最大重试次数
func (c *Jobs) WithMaxRetries(retries int) *Jobs {
	c.MaxRetries = retries
	return c
}

// WithRetryInterval 设置重试间隔
func (c *Jobs) WithRetryInterval(seconds int) *Jobs {
	c.RetryInterval = seconds
	return c
}

// AddTask 添加任务配置
func (c *Jobs) AddTask(name string, task TaskCfg) *Jobs {
	if c.Tasks == nil {
		c.Tasks = make(map[string]TaskCfg)
	}
	c.Tasks[name] = task
	return c
}

// RemoveTask 移除任务配置
func (c *Jobs) RemoveTask(name string) *Jobs {
	if c.Tasks != nil {
		delete(c.Tasks, name)
	}
	return c
}

// EnableTask 启用指定任务
func (c *Jobs) EnableTask(name string) *Jobs {
	if task, exists := c.Tasks[name]; exists {
		task.Enabled = true
		c.Tasks[name] = task
	}
	return c
}

// DisableTask 禁用指定任务
func (c *Jobs) DisableTask(name string) *Jobs {
	if task, exists := c.Tasks[name]; exists {
		task.Enabled = false
		c.Tasks[name] = task
	}
	return c
}
