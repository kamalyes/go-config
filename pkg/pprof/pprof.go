/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 01:12:30
 * @FilePath: \go-rpc-gateway\go-config\pkg\pprof\pprof.go
 * @Description: PProf性能分析配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package pprof

import "github.com/kamalyes/go-config/internal"

// PProf 性能分析配置
type PProf struct {
	ModuleName     string          `mapstructure:"module_name" yaml:"module-name" json:"module_name"`             // 模块名称
	Enabled        bool            `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                         // 是否启用PProf
	PathPrefix     string          `mapstructure:"path_prefix" yaml:"path_prefix" json:"path_prefix"`             // PProf路径前缀
	Port           int             `mapstructure:"port" yaml:"port" json:"port"`                                  // PProf服务端口
	EnableProfiles *ProfilesConfig `mapstructure:"enable_profiles" yaml:"enable-profiles" json:"enable_profiles"` // 启用的性能分析
	Sampling       *SamplingConfig `mapstructure:"sampling" yaml:"sampling" json:"sampling"`                      // 采样配置
}

// ProfilesConfig 性能分析配置
type ProfilesConfig struct {
	CPU          bool `mapstructure:"cpu" yaml:"cpu" json:"cpu"`                            // CPU性能分析
	Memory       bool `mapstructure:"memory" yaml:"memory" json:"memory"`                   // 内存性能分析
	Goroutine    bool `mapstructure:"goroutine" yaml:"goroutine" json:"goroutine"`          // 协程性能分析
	Block        bool `mapstructure:"block" yaml:"block" json:"block"`                      // 阻塞性能分析
	Mutex        bool `mapstructure:"mutex" yaml:"mutex" json:"mutex"`                      // 互斥锁性能分析
	Heap         bool `mapstructure:"heap" yaml:"heap" json:"heap"`                         // 堆内存分析
	Allocs       bool `mapstructure:"allocs" yaml:"allocs" json:"allocs"`                   // 内存分配分析
	ThreadCreate bool `mapstructure:"threadcreate" yaml:"threadcreate" json:"threadcreate"` // 线程创建分析
	Trace        bool `mapstructure:"trace" yaml:"trace" json:"trace"`                      // 追踪分析
}

// SamplingConfig 采样配置
type SamplingConfig struct {
	CPURate       int `mapstructure:"cpu_rate" yaml:"cpu-rate" json:"cpu_rate"`                   // CPU采样率(Hz)
	MemoryRate    int `mapstructure:"memory_rate" yaml:"memory-rate" json:"memory_rate"`          // 内存采样率
	BlockRate     int `mapstructure:"block_rate" yaml:"block-rate" json:"block_rate"`             // 阻塞采样率
	MutexFraction int `mapstructure:"mutex_fraction" yaml:"mutex-fraction" json:"mutex_fraction"` // 互斥锁采样比例
}

// Default 创建默认PProf配置
func Default() *PProf {
	return &PProf{
		ModuleName: "PProf",
		Enabled:    false,
		PathPrefix:       "/debug/PProf",
		Port:       6060,
		EnableProfiles: &ProfilesConfig{
			CPU:          true,
			Memory:       true,
			Goroutine:    true,
			Block:        false,
			Mutex:        false,
			Heap:         true,
			Allocs:       true,
			ThreadCreate: false,
			Trace:        false,
		},
		Sampling: &SamplingConfig{
			CPURate:       100,
			MemoryRate:    512 * 1024, // 512KB
			BlockRate:     1,
			MutexFraction: 1,
		},
	}
}

// Get 返回配置接口
func (c *PProf) Get() interface{} {
	return c
}

// Set 设置配置数据
func (c *PProf) Set(data interface{}) {
	if cfg, ok := data.(*PProf); ok {
		*c = *cfg
	}
}

// Clone 返回配置的副本
func (c *PProf) Clone() internal.Configurable {
	profiles := &ProfilesConfig{}
	sampling := &SamplingConfig{}

	if c.EnableProfiles != nil {
		*profiles = *c.EnableProfiles
	}
	if c.Sampling != nil {
		*sampling = *c.Sampling
	}

	return &PProf{
		ModuleName:     c.ModuleName,
		Enabled:        c.Enabled,
		PathPrefix:     c.PathPrefix,
		Port:           c.Port,
		EnableProfiles: profiles,
		Sampling:       sampling,
	}
}

// Validate 验证配置
func (c *PProf) Validate() error {
	return internal.ValidateStruct(c)
}

// WithModuleName 设置模块名称
func (c *PProf) WithModuleName(moduleName string) *PProf {
	c.ModuleName = moduleName
	return c
}

// WithEnabled 设置是否启用PProf
func (c *PProf) WithEnabled(enabled bool) *PProf {
	c.Enabled = enabled
	return c
}

// WithPathPrefix 设置PProf路径前缀
func (c *PProf) WithPathPrefix(pathPrefix string) *PProf {
	c.PathPrefix = pathPrefix
	return c
}

// WithPort 设置PProf服务端口
func (c *PProf) WithPort(port int) *PProf {
	c.Port = port
	return c
}

// WithProfiles 设置启用的性能分析
func (c *PProf) WithProfiles(cpu, memory, goroutine, block, mutex, heap, allocs, threadCreate, trace bool) *PProf {
	if c.EnableProfiles == nil {
		c.EnableProfiles = &ProfilesConfig{}
	}
	c.EnableProfiles.CPU = cpu
	c.EnableProfiles.Memory = memory
	c.EnableProfiles.Goroutine = goroutine
	c.EnableProfiles.Block = block
	c.EnableProfiles.Mutex = mutex
	c.EnableProfiles.Heap = heap
	c.EnableProfiles.Allocs = allocs
	c.EnableProfiles.ThreadCreate = threadCreate
	c.EnableProfiles.Trace = trace
	return c
}

// WithSampling 设置采样配置
func (c *PProf) WithSampling(cpuRate, memoryRate, blockRate, mutexFraction int) *PProf {
	if c.Sampling == nil {
		c.Sampling = &SamplingConfig{}
	}
	c.Sampling.CPURate = cpuRate
	c.Sampling.MemoryRate = memoryRate
	c.Sampling.BlockRate = blockRate
	c.Sampling.MutexFraction = mutexFraction
	return c
}

// EnableCPUProfile 启用CPU性能分析
func (c *PProf) EnableCPUProfile() *PProf {
	if c.EnableProfiles == nil {
		c.EnableProfiles = &ProfilesConfig{}
	}
	c.EnableProfiles.CPU = true
	return c
}

// EnableMemoryProfile 启用内存性能分析
func (c *PProf) EnableMemoryProfile() *PProf {
	if c.EnableProfiles == nil {
		c.EnableProfiles = &ProfilesConfig{}
	}
	c.EnableProfiles.Memory = true
	return c
}

// EnableGoroutineProfile 启用协程性能分析
func (c *PProf) EnableGoroutineProfile() *PProf {
	if c.EnableProfiles == nil {
		c.EnableProfiles = &ProfilesConfig{}
	}
	c.EnableProfiles.Goroutine = true
	return c
}

// EnableBlockProfile 启用阻塞性能分析
func (c *PProf) EnableBlockProfile() *PProf {
	if c.EnableProfiles == nil {
		c.EnableProfiles = &ProfilesConfig{}
	}
	c.EnableProfiles.Block = true
	return c
}

// EnableMutexProfile 启用互斥锁性能分析
func (c *PProf) EnableMutexProfile() *PProf {
	if c.EnableProfiles == nil {
		c.EnableProfiles = &ProfilesConfig{}
	}
	c.EnableProfiles.Mutex = true
	return c
}

// Enable 启用PProf
func (c *PProf) Enable() *PProf {
	c.Enabled = true
	return c
}

// Disable 禁用PProf
func (c *PProf) Disable() *PProf {
	c.Enabled = false
	return c
}

// IsEnabled 检查是否启用
func (c *PProf) IsEnabled() bool {
	return c.Enabled
}
