/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-24 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-24 09:35:01
 * @FilePath: \go-config\pkg\jobs\jobs_test.go
 * @Description: Jobs配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package jobs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobs_Clone(t *testing.T) {
	jobs := Default()
	jobs.MaxRetries = 5
	jobs.TimeZone = "America/New_York"

	cloned := jobs.Clone()

	assert.NotNil(t, cloned)
	clonedJobs, ok := cloned.(*Jobs)
	assert.True(t, ok)
	assert.Equal(t, jobs.MaxRetries, clonedJobs.MaxRetries)
	assert.Equal(t, jobs.TimeZone, clonedJobs.TimeZone)
	assert.Equal(t, jobs.Enabled, clonedJobs.Enabled)

	// 验证是独立副本 - 修改基本字段
	clonedJobs.MaxRetries = 10
	clonedJobs.TimeZone = "UTC"
	assert.NotEqual(t, jobs.MaxRetries, clonedJobs.MaxRetries)
	assert.NotEqual(t, jobs.TimeZone, clonedJobs.TimeZone)

	// 验证 Tasks map 被深拷贝
	assert.Equal(t, len(jobs.Tasks), len(clonedJobs.Tasks))

	// 测试 1: 修改克隆的 map，不应影响原始 map
	if len(clonedJobs.Tasks) > 0 {
		// 添加新任务到克隆版本
		clonedJobs.Tasks["new_cloned_task"] = TaskCfg{
			Enabled:  true,
			CronSpec: "0 0 * * *",
			Timeout:  300,
		}

		// 原始配置不应该有这个新任务
		_, exists := jobs.Tasks["new_cloned_task"]
		assert.False(t, exists, "向克隆的 map 添加元素不应影响原始 map")

		// 从原始 map 添加任务
		jobs.Tasks["new_original_task"] = TaskCfg{
			Enabled:  false,
			CronSpec: "0 1 * * *",
			Timeout:  600,
		}

		// 克隆的配置不应该有这个新任务
		_, exists = clonedJobs.Tasks["new_original_task"]
		assert.False(t, exists, "向原始 map 添加元素不应影响克隆的 map")
	}

	// 测试 2: 修改克隆的 map 中已存在的值，不应影响原始 map 的值
	if len(jobs.Tasks) > 0 {
		for taskName := range jobs.Tasks {
			originalTask := jobs.Tasks[taskName]
			originalEnabled := originalTask.Enabled
			originalCronSpec := originalTask.CronSpec
			originalTimeout := originalTask.Timeout

			// 修改克隆的任务
			clonedTask := clonedJobs.Tasks[taskName]
			clonedTask.Enabled = !originalEnabled
			clonedTask.CronSpec = "modified_" + originalCronSpec
			clonedTask.Timeout = originalTimeout + 100
			clonedTask.Description = "Modified Description"
			clonedJobs.Tasks[taskName] = clonedTask

			// 验证原始任务未被修改
			assert.Equal(t, originalEnabled, jobs.Tasks[taskName].Enabled,
				"修改克隆 map 中的值不应影响原始 map 的 Enabled")
			assert.Equal(t, originalCronSpec, jobs.Tasks[taskName].CronSpec,
				"修改克隆 map 中的值不应影响原始 map 的 CronSpec")
			assert.Equal(t, originalTimeout, jobs.Tasks[taskName].Timeout,
				"修改克隆 map 中的值不应影响原始 map 的 Timeout")
			assert.NotEqual(t, "Modified Description", jobs.Tasks[taskName].Description,
				"修改克隆 map 中的值不应影响原始 map 的 Description")

			// 验证克隆的任务已被修改
			assert.Equal(t, !originalEnabled, clonedJobs.Tasks[taskName].Enabled)
			assert.Equal(t, "modified_"+originalCronSpec, clonedJobs.Tasks[taskName].CronSpec)
			assert.Equal(t, originalTimeout+100, clonedJobs.Tasks[taskName].Timeout)
			assert.Equal(t, "Modified Description", clonedJobs.Tasks[taskName].Description)

			break // 只测试第一个任务
		}
	}

	// 测试 3: 验证嵌套的 slice (Dependencies, Tags) 也被深拷贝
	if len(jobs.Tasks) > 0 {
		for taskName, task := range jobs.Tasks {
			if len(task.Dependencies) > 0 {
				originalDepCount := len(jobs.Tasks[taskName].Dependencies)

				// 修改克隆任务的 Dependencies
				clonedTask := clonedJobs.Tasks[taskName]
				clonedTask.Dependencies = append(clonedTask.Dependencies, DependencyTask{
					TaskName: "new_dependency",
				})
				clonedJobs.Tasks[taskName] = clonedTask

				// 验证原始任务的 Dependencies 未被修改
				assert.Equal(t, originalDepCount, len(jobs.Tasks[taskName].Dependencies),
					"修改克隆任务的 Dependencies 不应影响原始任务")
				assert.NotEqual(t, len(jobs.Tasks[taskName].Dependencies),
					len(clonedJobs.Tasks[taskName].Dependencies),
					"克隆的 Dependencies slice 应该是独立的")
				break
			}
		}
	}
}

func TestJobs_Validate(t *testing.T) {
	tests := []struct {
		name    string
		jobs    *Jobs
		wantErr bool
	}{
		{
			name:    "valid config",
			jobs:    Default(),
			wantErr: false,
		},
		{
			name: "empty timezone",
			jobs: &Jobs{
				Enabled:          true,
				TimeZone:         "",
				GracefulShutdown: 30,
			},
			wantErr: true,
		},
		{
			name: "invalid graceful shutdown",
			jobs: &Jobs{
				Enabled:          true,
				TimeZone:         "Asia/Shanghai",
				GracefulShutdown: 0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.jobs.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestJobs_Enable(t *testing.T) {
	jobs := &Jobs{Enabled: false}
	result := jobs.Enable()
	assert.True(t, result.Enabled)
	assert.Equal(t, jobs, result) // 验证链式调用
}

func TestJobs_Disable(t *testing.T) {
	jobs := Default()
	result := jobs.Disable()
	assert.False(t, result.Enabled)
	assert.Equal(t, jobs, result)
}

func TestJobs_AddTask(t *testing.T) {
	jobs := &Jobs{}
	task := TaskCfg{
		Enabled:  true,
		CronSpec: "0 0 * * *",
	}

	result := jobs.AddTask("test-task", task)
	assert.NotNil(t, result.Tasks)
	assert.Contains(t, result.Tasks, "test-task")
	assert.Equal(t, task.CronSpec, result.Tasks["test-task"].CronSpec)
}

func TestJobs_RemoveTask(t *testing.T) {
	jobs := Default()
	initialCount := len(jobs.Tasks)

	result := jobs.RemoveTask("cleanup")
	assert.Len(t, result.Tasks, initialCount-1)
	assert.NotContains(t, result.Tasks, "cleanup")
}

func TestJobs_IsTaskEnabled(t *testing.T) {
	jobs := Default()

	assert.True(t, jobs.IsTaskEnabled("cleanup"))
	assert.False(t, jobs.IsTaskEnabled("non-existent"))

	jobs.Enabled = false
	assert.False(t, jobs.IsTaskEnabled("cleanup"))
}

func TestTaskCfg_Validate(t *testing.T) {
	tests := []struct {
		name    string
		task    TaskCfg
		wantErr bool
	}{
		{
			name: "valid task",
			task: TaskCfg{
				Enabled:  true,
				CronSpec: "0 0 * * *",
				Timeout:  300,
				Priority: 1,
			},
			wantErr: false,
		},
		{
			name: "empty cron spec",
			task: TaskCfg{
				Enabled:  true,
				CronSpec: "",
			},
			wantErr: true,
		},
		{
			name: "negative timeout",
			task: TaskCfg{
				Enabled:  true,
				CronSpec: "0 0 * * *",
				Timeout:  -1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.task.Validate("test-task")
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestBreakerCfg_Validate(t *testing.T) {
	tests := []struct {
		name    string
		breaker BreakerCfg
		wantErr bool
	}{
		{
			name: "valid breaker",
			breaker: BreakerCfg{
				Enabled:           true,
				MaxFailures:       5,
				ResetTimeout:      30,
				HalfOpenSuccesses: 2,
			},
			wantErr: false,
		},
		{
			name: "disabled breaker",
			breaker: BreakerCfg{
				Enabled: false,
			},
			wantErr: false,
		},
		{
			name: "zero max failures",
			breaker: BreakerCfg{
				Enabled:     true,
				MaxFailures: 0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.breaker.Validate("test-task")
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestJobs_WithRetryJitter(t *testing.T) {
	jobs := &Jobs{}
	result := jobs.WithRetryJitter(0.5)
	assert.Equal(t, 0.5, result.RetryJitter)
	assert.Equal(t, jobs, result)
}

func TestJobs_WithMaxConcurrentJobs(t *testing.T) {
	jobs := &Jobs{}
	result := jobs.WithMaxConcurrentJobs(10)
	assert.Equal(t, 10, result.MaxConcurrentJobs)
	assert.Equal(t, jobs, result)
}

func TestJobs_EnableDistribute(t *testing.T) {
	jobs := &Jobs{Distribute: false}
	result := jobs.EnableDistribute()
	assert.True(t, result.Distribute)
	assert.Equal(t, jobs, result)
}

func TestJobs_DisableDistribute(t *testing.T) {
	jobs := &Jobs{Distribute: true}
	result := jobs.DisableDistribute()
	assert.False(t, result.Distribute)
	assert.Equal(t, jobs, result)
}

func TestJobs_GetAllTaskNames(t *testing.T) {
	jobs := Default()
	names := jobs.GetAllTaskNames()
	assert.Len(t, names, 2)
	assert.Contains(t, names, "cleanup")
	assert.Contains(t, names, "data-sync")
}

func TestJobs_GetEnabledTasks(t *testing.T) {
	jobs := Default()
	enabled := jobs.GetEnabledTasks()
	assert.Len(t, enabled, 2)
	assert.Contains(t, enabled, "cleanup")
}

func TestJobs_AddDependencyToTask(t *testing.T) {
	jobs := Default()
	dep := DependencyTask{TaskName: "new-dep"}

	result := jobs.AddDependencyToTask("cleanup", dep)
	assert.Equal(t, jobs, result)

	task, _ := jobs.GetTaskConfig("cleanup")
	assert.Len(t, task.Dependencies, 1)
	assert.Equal(t, "new-dep", task.Dependencies[0].TaskName)
}

func TestJobs_RemoveDependencyFromTask(t *testing.T) {
	jobs := Default()

	result := jobs.RemoveDependencyFromTask("data-sync", "cleanup")
	assert.Equal(t, jobs, result)

	task, _ := jobs.GetTaskConfig("data-sync")
	// 应该还有一个内联依赖
	assert.Len(t, task.Dependencies, 1)
	assert.NotNil(t, task.Dependencies[0].Inline)
}

func TestJobs_AddTagToTask(t *testing.T) {
	jobs := Default()

	result := jobs.AddTagToTask("cleanup", "critical")
	assert.Equal(t, jobs, result)

	task, _ := jobs.GetTaskConfig("cleanup")
	assert.Contains(t, task.Tags, "critical")

	// 测试重复添加
	jobs.AddTagToTask("cleanup", "critical")
	task, _ = jobs.GetTaskConfig("cleanup")
	count := 0
	for _, tag := range task.Tags {
		if tag == "critical" {
			count++
		}
	}
	assert.Equal(t, 1, count, "标签不应重复添加")
}

func TestJobs_RemoveTagFromTask(t *testing.T) {
	jobs := Default()
	jobs.AddTagToTask("cleanup", "test-tag")

	result := jobs.RemoveTagFromTask("cleanup", "test-tag")
	assert.Equal(t, jobs, result)

	task, _ := jobs.GetTaskConfig("cleanup")
	assert.NotContains(t, task.Tags, "test-tag")
}

func TestJobs_SetTaskPriority(t *testing.T) {
	jobs := Default()

	result := jobs.SetTaskPriority("cleanup", 5)
	assert.Equal(t, jobs, result)

	task, _ := jobs.GetTaskConfig("cleanup")
	assert.Equal(t, 5, task.Priority)
}

func TestJobs_SetTaskTimeout(t *testing.T) {
	jobs := Default()

	result := jobs.SetTaskTimeout("cleanup", 600)
	assert.Equal(t, jobs, result)

	task, _ := jobs.GetTaskConfig("cleanup")
	assert.Equal(t, 600, task.Timeout)
}

func TestJobs_SetTaskMaxRetries(t *testing.T) {
	jobs := Default()

	result := jobs.SetTaskMaxRetries("cleanup", 5)
	assert.Equal(t, jobs, result)

	task, _ := jobs.GetTaskConfig("cleanup")
	assert.Equal(t, 5, task.MaxRetries)
}

func TestJobs_SetTaskCronSpec(t *testing.T) {
	jobs := Default()

	result := jobs.SetTaskCronSpec("cleanup", "0 0 * * *")
	assert.Equal(t, jobs, result)

	task, _ := jobs.GetTaskConfig("cleanup")
	assert.Equal(t, "0 0 * * *", task.CronSpec)
}

func TestJobs_GetTasksByTag(t *testing.T) {
	jobs := Default()
	jobs.AddTagToTask("cleanup", "maintenance")
	jobs.AddTagToTask("data-sync", "sync")

	tasks := jobs.GetTasksByTag("maintenance")
	assert.Len(t, tasks, 1)
	assert.Contains(t, tasks, "cleanup")

	tasks = jobs.GetTasksByTag("non-existent")
	assert.Len(t, tasks, 0)
}

func TestJobs_GetTasksByPriority(t *testing.T) {
	jobs := Default()
	jobs.SetTaskPriority("cleanup", 1)
	jobs.SetTaskPriority("data-sync", 5)

	tasks := jobs.GetTasksByPriority(1, 3)
	assert.Len(t, tasks, 1)
	assert.Contains(t, tasks, "cleanup")

	tasks = jobs.GetTasksByPriority(0, 10)
	assert.Len(t, tasks, 2)
}

func TestJobs_GetTaskConfig(t *testing.T) {
	jobs := Default()

	task, exists := jobs.GetTaskConfig("cleanup")
	assert.True(t, exists)
	assert.Equal(t, "定期清理过期数据和无效缓存", task.Description)

	_, exists = jobs.GetTaskConfig("non-existent")
	assert.False(t, exists)
}

func TestJobs_GetTimeZoneLocation(t *testing.T) {
	jobs := Default()

	loc, err := jobs.GetTimeZoneLocation()
	assert.NoError(t, err)
	assert.NotNil(t, loc)
	assert.Equal(t, "Asia/Shanghai", loc.String())

	jobs.TimeZone = "Invalid/Zone"
	_, err = jobs.GetTimeZoneLocation()
	assert.Error(t, err)
}

func TestJobs_ChainedCalls(t *testing.T) {
	jobs := &Jobs{}
	result := jobs.
		Enable().
		WithTimeZone("UTC").
		WithGracefulShutdown(60).
		WithMaxRetries(5).
		WithRetryInterval(10).
		WithRetryJitter(0.3).
		WithMaxConcurrentJobs(20).
		EnableDistribute()

	assert.True(t, result.Enabled)
	assert.Equal(t, "UTC", result.TimeZone)
	assert.Equal(t, 60, result.GracefulShutdown)
	assert.Equal(t, 5, result.MaxRetries)
	assert.Equal(t, 10, result.RetryInterval)
	assert.Equal(t, 0.3, result.RetryJitter)
	assert.Equal(t, 20, result.MaxConcurrentJobs)
	assert.True(t, result.Distribute)
}

func TestJobs_GetSet(t *testing.T) {
	jobs := Default()

	// Test Get
	result := jobs.Get()
	assert.Equal(t, jobs, result)

	// Test Set
	newJobs := &Jobs{
		Enabled:  false,
		TimeZone: "UTC",
	}
	jobs.Set(newJobs)
	assert.Equal(t, newJobs.Enabled, jobs.Enabled)
	assert.Equal(t, newJobs.TimeZone, jobs.TimeZone)
}
