package main

import (
	"fmt"

	goconfig "github.com/kamalyes/go-config"
)

// ConfigDiscovery 演示配置发现
func ConfigDiscovery() {
	fmt.Println("=== 配置发现示例 ===\n")

	// 1. 自动发现配置文件
	fmt.Println("1. 自动发现配置文件:")
	configFiles, err := goconfig.DiscoverConfig("../resources", goconfig.EnvDevelopment)
	if err != nil {
		fmt.Printf("   发现配置文件失败: %v\n", err)
	} else {
		fmt.Printf("   发现 %d 个配置文件:\n", len(configFiles))
		for i, file := range configFiles {
			if file != nil {
				fmt.Printf("   [%d] %s\n", i+1, file.Path)
			}
		}
	}

	// 2. 找到最佳配置文件
	fmt.Println("\n2. 查找最佳配置文件:")
	bestFile, err := goconfig.FindBestConfig("../resources", goconfig.EnvDevelopment)
	if err != nil {
		fmt.Printf("   查找失败: %v\n", err)
	} else {
		fmt.Printf("   最佳配置: %s\n", bestFile)
	}

	// 3. 扫描目录中的所有配置文件
	fmt.Println("\n3. 扫描目录中的所有配置文件:")
	allFiles, err := goconfig.GetGlobalConfigDiscovery().ScanDirectory("../resources")
	if err != nil {
		fmt.Printf("   扫描失败: %v\n", err)
	} else {
		fmt.Printf("   找到 %d 个配置文件:\n", len(allFiles))
		for i, file := range allFiles {
			if file != nil {
				fmt.Printf("   [%d] %s (环境: %s)\n", i+1, file.Path, file.Environment)
			}
		}
	}

	// 4. 使用模式匹配查找配置
	fmt.Println("\n4. 使用模式 'gateway-xl*' 查找配置:")
	discovery := goconfig.GetGlobalConfigDiscovery()
	files, err := discovery.FindConfigFileByPattern("../resources", "gateway-xl*", goconfig.EnvDevelopment)
	if err != nil {
		fmt.Printf("   查找失败: %v\n", err)
	} else if len(files) > 0 {
		fmt.Printf("   找到 %d 个匹配的配置文件:\n", len(files))
		for i, file := range files {
			if file != nil {
				fmt.Printf("   [%d] %s\n", i+1, file.Path)
			}
		}
	} else {
		fmt.Println("   未找到匹配的配置文件")
	}

	fmt.Println("\n=== 示例完成 ===")
}

func main() {
	ConfigDiscovery()
}
