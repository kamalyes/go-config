/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 21:52:32
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
)

// ConfigOptions 配置选项结构体
type ConfigOptions struct {
	ConfigSuffix string
	ConfigType   string
	ConfigPath   string
}

// ConfigManager 负责加载和管理配置
type ConfigManager struct {
	config  *MultiConfig
	options ConfigOptions
}

// NewConfigManager 创建一个新的 ConfigManager
func NewConfigManager(ctx context.Context, options *ConfigOptions) (*ConfigManager, error) {
	// 如果没有传入 options，则使用默认值
	if options == nil {
		options = &ConfigOptions{
			ConfigSuffix: defaultConfigSuffix,
			ConfigType:   defaultConfigType,
			ConfigPath:   defaultConfigPath,
		}
	}

	// 使用默认值替换空字段
	if options.ConfigType == "" {
		options.ConfigType = defaultConfigType
	}
	if options.ConfigPath == "" {
		options.ConfigPath = defaultConfigPath
	}
	if options.ConfigSuffix == "" {
		options.ConfigSuffix = defaultConfigSuffix
	}

	cm := &ConfigManager{
		options: *options,
	}

	// 直接使用 cm 调用 loadConfig
	config, err := cm.loadConfig(ctx)
	if err != nil {
		return nil, err
	}
	cm.config = config
	return cm, nil
}

// loadConfig 加载配置文件并返回 Config 对象
func (cm *ConfigManager) loadConfig(ctx context.Context) (*MultiConfig, error) {
	// 从上下文获取当前环境
	contextEnv := env.FromContext(ctx)

	// 确定使用的环境
	filename := contextEnv.String() + cm.options.ConfigSuffix

	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(cm.options.ConfigType)
	v.AddConfigPath(cm.options.ConfigPath)
	log.Printf("读取配置文件: %s", filename)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件异常: %w", err)
	}

	globalConfig := &MultiConfig{Viper: v}

	if err := v.Unmarshal(globalConfig); err != nil {
		return nil, fmt.Errorf("读取配置文件异常: %w", err)
	}

	// 监控配置改变
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置文件内容发生改变: %s", e.Name)
		if err := v.Unmarshal(globalConfig); err != nil {
			log.Fatalf("读取配置文件异常: %s", err)
		}
	})

	return globalConfig, nil
}

// SubItem 从配置中获取指定的配置子项
func (cm *ConfigManager) SubItem(ctx context.Context, subKey string, v interface{}) {
	value := cm.config.Viper.Sub(subKey)
	if value == nil {
		return
	}
	if err := value.Unmarshal(v); err != nil {
		log.Printf("读取子配置项异常: %s", err)
	}
}

// GetConfig 获取全局配置
func (cm *ConfigManager) GetConfig() *MultiConfig {
	return cm.config
}
