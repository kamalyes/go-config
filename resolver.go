/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:18:37
 * @FilePath: \go-config\resolver.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"context"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/kamalyes/go-config/pkg/env"
	"github.com/spf13/viper"
)

const (
	defaultConfigSuffix = "_config"     // 默认配置文件后缀
	defaultConfigType   = "yaml"        // 默认配置文件类型
	defaultConfigPath   = "./resources" // 默认配置文件路径
	defaultUseEnvLevel  = EnvLevelOS
)

const (
	EnvLevelOS  EnvLevel = "os"
	EnvLevelCtx EnvLevel = "ctx"
)

type EnvLevel string

// ConfigOptions 定义配置选项结构体
type ConfigOptions struct {
	ConfigType    string              // 配置文件类型
	ConfigPath    string              // 配置文件路径
	ConfigSuffix  string              // 配置文件后缀
	EnvValue      env.EnvironmentType // 初始化环境
	EnvContextKey env.ContextKey      // 环境上下文Key
	UseEnvLevel   EnvLevel            // 使用环境级别，值为 "os" 或 "ctx"
}

// MultiConfigManager 负责加载和管理 MultiConfig
type MultiConfigManager struct {
	MultiConfig *MultiConfig
	Options     ConfigOptions
}

// SingleConfigManager 负责加载和管理 SingleConfig
type SingleConfigManager struct {
	SingleConfig *SingleConfig
	Options      ConfigOptions
}

// GetDefaultConfigOptions 返回 ConfigOptions 的默认值
func GetDefaultConfigOptions() *ConfigOptions {
	return &ConfigOptions{
		ConfigSuffix:  defaultConfigSuffix,
		ConfigType:    defaultConfigType,
		ConfigPath:    defaultConfigPath,
		EnvValue:      env.DefaultEnv,
		EnvContextKey: env.GetContextKey(),
		UseEnvLevel:   defaultUseEnvLevel,
	}
}

// NewMultiConfigManager 创建一个新的 MultiConfigManager
func NewMultiConfigManager(ctx context.Context, options *ConfigOptions) (*MultiConfigManager, error) {
	mcm := &MultiConfigManager{}
	if err := mcm.initialize(options, &MultiConfig{}); err != nil {
		return nil, err
	}
	return mcm, nil
}

// NewSingleConfigManager 创建一个新的 SingleConfigManager
func NewSingleConfigManager(ctx context.Context, options *ConfigOptions) (*SingleConfigManager, error) {
	scm := &SingleConfigManager{}
	if err := scm.initialize(options, &SingleConfig{}); err != nil {
		return nil, err
	}
	return scm, nil
}

// initialize 初始化配置选项并加载配置
func (m *MultiConfigManager) initialize(options *ConfigOptions, config interface{}) error {
	m.Options = initializeConfigOptions(options)
	multiConfig, newOptions, err := loadConfig(config, m.Options)
	if err != nil {
		return err
	}
	m.Options = newOptions
	m.MultiConfig = multiConfig.(*MultiConfig) // 类型断言
	return nil
}

func (m *SingleConfigManager) initialize(options *ConfigOptions, config interface{}) error {
	m.Options = initializeConfigOptions(options)
	singleConfig, newOptions, err := loadConfig(config, m.Options)
	if err != nil {
		return err
	}
	m.Options = newOptions
	m.SingleConfig = singleConfig.(*SingleConfig) // 类型断言
	return nil
}

// initializeConfigOptions 使用默认值替换空字段
func initializeConfigOptions(options *ConfigOptions) ConfigOptions {
	if options == nil {
		options = GetDefaultConfigOptions()
	}

	if options.ConfigType == "" {
		options.ConfigType = defaultConfigType
	}

	if options.ConfigPath == "" {
		options.ConfigPath = defaultConfigPath
	}

	if options.ConfigSuffix == "" {
		options.ConfigSuffix = defaultConfigSuffix
	}

	if options.EnvValue == "" {
		options.EnvValue = env.DefaultEnv
	}

	if options.EnvContextKey == "" {
		options.EnvContextKey = env.GetContextKey()
	}

	if options.UseEnvLevel == "" {
		options.UseEnvLevel = defaultUseEnvLevel
	}

	env.SetContextKey(&env.ContextKeyOptions{Key: options.EnvContextKey, Value: options.EnvValue})

	return *options
}

// loadConfig 加载配置文件并返回相应的配置对象
func loadConfig(config interface{}, options ConfigOptions) (interface{}, ConfigOptions, error) {
	// 获取OS环境变量中的值
	osEnv := env.GetEnvironment()
	// 使用变量将Os值存起来为当前
	contextEnv := osEnv

	// 如果俩环境变量不一致
	if osEnv != env.EnvironmentType(contextEnv) {
		// 优先级为Ctx、就使用Ctx进行替换
		if options.UseEnvLevel == EnvLevelCtx {
			contextEnv = options.EnvValue
		}
	}
	// 最后更新options EnvValue
	options.EnvValue = contextEnv

	// 确定使用的配置文件
	filename := contextEnv.String() + options.ConfigSuffix

	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(options.ConfigType)
	v.AddConfigPath(options.ConfigPath)
	log.Printf("读取配置文件: %s", filename)

	if err := v.ReadInConfig(); err != nil {
		return nil, options, fmt.Errorf("读取配置文件异常: %w", err)
	}

	if err := v.Unmarshal(config); err != nil {
		return nil, options, fmt.Errorf("读取配置文件异常: %w", err)
	}

	// 监控配置改变
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置文件内容发生改变: %s", e.Name)
		if err := v.Unmarshal(config); err != nil {
			log.Printf("读取配置文件异常: %s", err)
		}
	})

	return config, options, nil
}

// SubItem 从配置中获取指定的配置子项
func (m *MultiConfigManager) SubItem(ctx context.Context, subKey string, v interface{}) {
	if m.MultiConfig == nil {
		log.Println("MultiConfig is nil")
		return
	}
	value := m.MultiConfig.Viper.Sub(subKey)
	if value == nil {
		log.Printf("子配置项 %s 不存在", subKey)
		return
	}
	if err := value.Unmarshal(v); err != nil {
		log.Printf("读取子配置项异常: %s", err)
	}
}

// GetConfig 获取 MultiConfig 配置
func (m *MultiConfigManager) GetConfig() *MultiConfig {
	return m.MultiConfig
}

// GetConfig 获取 SingleConfig 配置
func (m *SingleConfigManager) GetConfig() *SingleConfig {
	return m.SingleConfig
}
