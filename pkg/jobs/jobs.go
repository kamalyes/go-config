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
	Enabled           bool               `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                   // 是否启用Job管理器
	TimeZone          string             `mapstructure:"timezone" yaml:"timezone" json:"timezone"`                                // 时区配置(例如: Asia/Shanghai)
	GracefulShutdown  int                `mapstructure:"graceful-shutdown" yaml:"graceful-shutdown" json:"gracefulShutdown"`      // 优雅关闭超时时间(秒)
	MaxRetries        int                `mapstructure:"max-retries" yaml:"max-retries" json:"maxRetries"`                        // 失败最大重试次数
	RetryInterval     int                `mapstructure:"retry-interval" yaml:"retry-interval" json:"retryInterval"`               // 重试间隔(秒)
	RetryJitter       float64            `mapstructure:"retry-jitter" yaml:"retry-jitter" json:"retryJitter"`                     // 重试间隔抖动百分比(0-1)
	MaxConcurrentJobs int                `mapstructure:"max-concurrent-jobs" yaml:"max-concurrent-jobs" json:"maxConcurrentJobs"` // 最大并发任务数，0表示不限制
	EnableDistribute  bool               `mapstructure:"enable-distribute" yaml:"enable-distribute" json:"enableDistribute"`      // 是否启用分布式调度
	Tasks             map[string]TaskCfg `mapstructure:"tasks" yaml:"tasks" json:"tasks"`                                         // 任务配置
}

// TaskCfg 任务配置
type TaskCfg struct {
	Enabled        bool             `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                        // 是否启用
	CronSpec       string           `mapstructure:"cron-spec" yaml:"cron-spec" json:"cronSpec"`                   // Cron表达式
	ImmediateStart bool             `mapstructure:"immediate-start" yaml:"immediate-start" json:"immediateStart"` // 启动时立即执行一次
	Timeout        int              `mapstructure:"timeout" yaml:"timeout" json:"timeout"`                        // 任务超时时间(秒)，0表示无限制
	OverlapPrevent bool             `mapstructure:"overlap-prevent" yaml:"overlap-prevent" json:"overlapPrevent"` // 是否阻止任务重叠执行
	MaxRetries     int              `mapstructure:"max-retries" yaml:"max-retries" json:"maxRetries"`             // 任务级别的最大重试次数，0使用全局配置
	RetryInterval  int              `mapstructure:"retry-interval" yaml:"retry-interval" json:"retryInterval"`    // 任务级别的重试间隔(秒)，0使用全局配置
	RetryJitter    float64          `mapstructure:"retry-jitter" yaml:"retry-jitter" json:"retryJitter"`          // 任务级别的重试抖动，0使用全局配置
	Priority       int              `mapstructure:"priority" yaml:"priority" json:"priority"`                     // 任务优先级(0-5, 0=最低, 5=最高)
	Dependencies   []DependencyTask `mapstructure:"dependencies" yaml:"dependencies" json:"dependencies"`         // 任务依赖列表(工作流模式)，支持引用或内联配置
	MaxConcurrent  int              `mapstructure:"max-concurrent" yaml:"max-concurrent" json:"maxConcurrent"`    // 最大并发执行数，0表示不限制
	Tags           []string         `mapstructure:"tags" yaml:"tags" json:"tags"`                                 // 任务标签(用于分组和筛选)
	Description    string           `mapstructure:"description" yaml:"description" json:"description"`            // 任务描述
	Breaker        BreakerCfg       `mapstructure:"breaker" yaml:"breaker" json:"breaker"`                        // 熔断器配置
}

// DependencyTask 任务依赖配置
type DependencyTask struct {
	TaskName string   `mapstructure:"task-name" yaml:"task-name" json:"taskName"` // 依赖的任务名称(引用已定义的任务)
	Inline   *TaskCfg `mapstructure:"inline" yaml:"inline" json:"inline"`         // 内联任务配置(可选，用于定义临时依赖任务)
}

// BreakerCfg 熔断器配置（任务保护）
type BreakerCfg struct {
	Enabled           bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                                   // 是否启用熔断器
	MaxFailures       int  `mapstructure:"max-failures" yaml:"max-failures" json:"maxFailures"`                     // 最大连续失败次数触发熔断
	ResetTimeout      int  `mapstructure:"reset-timeout" yaml:"reset-timeout" json:"resetTimeout"`                  // 熔断恢复超时时间(秒)
	HalfOpenSuccesses int  `mapstructure:"half-open-successes" yaml:"half-open-successes" json:"halfOpenSuccesses"` // 半开状态需要的成功次数才能完全恢复
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
				MaxRetries:     3,
				RetryInterval:  10,
				Description:    "定期清理过期数据和无效缓存",
				Breaker: BreakerCfg{
					Enabled:           true,
					MaxFailures:       5,
					ResetTimeout:      30,
					HalfOpenSuccesses: 2,
				},
			},
			// 数据同步Job - 每小时执行，依赖清理任务
			"data-sync": {
				Enabled:        true,
				CronSpec:       "0 0 * * * *",
				ImmediateStart: false,
				Timeout:        600,
				OverlapPrevent: true,
				MaxRetries:     3,
				RetryInterval:  15,
				Description:    "同步数据到远程服务器",
				Dependencies: []DependencyTask{
					{
						TaskName: "cleanup", // 引用已定义的清理任务
					},
					{
						Inline: &TaskCfg{ // 内联定义数据验证任务
							Enabled:        true,
							CronSpec:       "0 */30 * * * *",
							Timeout:        60,
							OverlapPrevent: true,
							Description:    "数据验证前置任务",
						},
					},
				},
				Breaker: BreakerCfg{
					Enabled:           true,
					MaxFailures:       3,
					ResetTimeout:      60,
					HalfOpenSuccesses: 2,
				},
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

	if t.Priority < 0 || t.Priority > 99999 {
		return fmt.Errorf("任务[%s]的priority必须在0-99999之间", name)
	}

	if t.MaxConcurrent < 0 {
		return fmt.Errorf("任务[%s]的max_concurrent不能小于0", name)
	}

	// 验证熔断器配置
	if err := t.Breaker.Validate(name); err != nil {
		return err
	}

	return nil
}

// Validate 验证熔断器配置
func (b *BreakerCfg) Validate(taskName string) error {
	if !b.Enabled {
		return nil
	}

	if b.MaxFailures <= 0 {
		return fmt.Errorf("任务[%s]的熔断器max_failures必须大于0", taskName)
	}

	if b.ResetTimeout <= 0 {
		return fmt.Errorf("任务[%s]的熔断器reset_timeout必须大于0", taskName)
	}

	if b.HalfOpenSuccesses <= 0 {
		return fmt.Errorf("任务[%s]的熔断器half_open_successes必须大于0", taskName)
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
		// 深拷贝 Tags
		tags := make([]string, len(v.Tags))
		copy(tags, v.Tags)

		// 深拷贝 Dependencies
		deps := make([]DependencyTask, len(v.Dependencies))
		for i, dep := range v.Dependencies {
			deps[i] = DependencyTask{
				TaskName: dep.TaskName,
			}
			// 深拷贝内联任务配置
			if dep.Inline != nil {
				inlineCopy := *dep.Inline
				// 递归拷贝内联任务的 Tags
				if len(inlineCopy.Tags) > 0 {
					inlineCopy.Tags = make([]string, len(inlineCopy.Tags))
					copy(inlineCopy.Tags, dep.Inline.Tags)
				}
				// 递归拷贝内联任务的 Dependencies
				if len(inlineCopy.Dependencies) > 0 {
					inlineCopy.Dependencies = make([]DependencyTask, len(dep.Inline.Dependencies))
					copy(inlineCopy.Dependencies, dep.Inline.Dependencies)
				}
				deps[i].Inline = &inlineCopy
			}
		}

		// 深拷贝每个任务配置
		tasks[k] = TaskCfg{
			Enabled:        v.Enabled,
			CronSpec:       v.CronSpec,
			ImmediateStart: v.ImmediateStart,
			Timeout:        v.Timeout,
			OverlapPrevent: v.OverlapPrevent,
			MaxRetries:     v.MaxRetries,
			RetryInterval:  v.RetryInterval,
			RetryJitter:    v.RetryJitter,
			Priority:       v.Priority,
			Dependencies:   deps,
			MaxConcurrent:  v.MaxConcurrent,
			Tags:           tags,
			Description:    v.Description,
			Breaker:        v.Breaker, // BreakerCfg 是值类型，自动深拷贝
		}
	}

	return &Jobs{
		Enabled:           c.Enabled,
		TimeZone:          c.TimeZone,
		GracefulShutdown:  c.GracefulShutdown,
		MaxRetries:        c.MaxRetries,
		RetryInterval:     c.RetryInterval,
		RetryJitter:       c.RetryJitter,
		MaxConcurrentJobs: c.MaxConcurrentJobs,
		EnableDistribute:  c.EnableDistribute,
		Tasks:             tasks,
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
