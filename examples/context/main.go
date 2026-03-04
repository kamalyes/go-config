package main

import (
	"context"
	"fmt"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/gateway"
)

// ContextIntegration 演示上下文集成
func ContextIntegration() {
	config := gateway.Default()

	manager := goconfig.NewConfigBuilder(config).
		WithPrefix("gateway-xl").
		WithSearchPath("../resources").
		WithEnvironment(goconfig.EnvDevelopment).
		WithContext(&goconfig.ContextKeyOptions{
			Key:   "APP_ENV",
			Value: goconfig.EnvDevelopment,
		}).
		MustBuildAndStart()

	defer manager.Stop()

	// 将配置注入到上下文
	ctx := manager.WithContext(context.Background())

	// 从上下文获取配置
	if cfg, ok := goconfig.GetConfigFromContext(ctx); ok {
		fmt.Printf("从上下文获取配置: %+v\n", cfg)
	}

	// 从上下文获取环境
	if env, ok := goconfig.GetEnvironmentFromContext(ctx); ok {
		fmt.Printf("当前环境: %s\n", env)
	}
}

func main() {
	ContextIntegration()
}
