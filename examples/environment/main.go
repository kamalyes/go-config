package main

import (
	"fmt"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/pkg/gateway"
)

// EnvironmentManagement 演示环境管理
func EnvironmentManagement() {
	config := gateway.Default()

	// 基础判断
	if goconfig.IsDev() {
		fmt.Println("当前是开发环境")
	}

	if goconfig.IsProduction() {
		fmt.Println("当前是生产环境")
	}

	// 环境级别判断
	if goconfig.IsProductionLevel() {
		fmt.Println("启用生产级别的监控和安全功能")
	}

	// 通用判断
	if goconfig.IsEnvironment(goconfig.EnvChina) {
		fmt.Println("使用中国特定配置")
	}

	// 多环境判断
	if goconfig.IsAnyOf(goconfig.EnvDevelopment, goconfig.EnvTest, goconfig.EnvLocal) {
		fmt.Println("开发相关环境逻辑")
	}

	// 使用不同环境配置
	manager := goconfig.NewConfigBuilder(config).
		WithPrefix("gateway-xl").
		WithSearchPath("../resources").
		WithEnvironment(goconfig.EnvDevelopment).
		MustBuildAndStart()

	defer manager.Stop()

	fmt.Printf("当前环境: %s\n", goconfig.GetCurrentEnvironment())
}

func main() {
	EnvironmentManagement()
}
