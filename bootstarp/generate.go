/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-22 01:05:00
 * @FilePath: \go-config\cmd\generate-configs\main.go
 * @Description: 配置文件生成命令行工具
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package main

import (
	"flag"
	"fmt"
	goconfig "github.com/kamalyes/go-config"
	"os"
	"strings"
)

func main() {
	// 定义命令行参数
	// del /S /Q pkg\*.yaml pkg\*.json
	outputDir := flag.String("output", ".", "输出目录,配置文件将生成到该目录下的pkg文件夹")
	modules := flag.String("modules", "", "指定要生成的模块,多个模块用逗号分隔(为空则生成所有模块)")
	listModules := flag.Bool("list", false, "列出所有可用的模块")
	help := flag.Bool("help", false, "显示帮助信息")
	backup := flag.Bool("backup", false, "是否备份已有的配置文件")

	flag.Parse()

	// 显示帮助
	if *help {
		showHelp()
		return
	}

	// 创建生成器
	generator := goconfig.NewSmartConfigGenerator(*outputDir).WithBackupExisting(*backup)

	// 列出所有模块
	if *listModules {
		listAllModules(generator)
		return
	}

	// 生成配置文件
	var err error
	if *modules == "" {
		// 生成所有模块
		fmt.Println("正在生成所有模块的配置文件...")
		err = generator.GenerateAllConfigs()
	} else {
		// 生成指定模块
		moduleList := strings.Split(*modules, ",")
		for i, m := range moduleList {
			moduleList[i] = strings.TrimSpace(m)
		}
		fmt.Printf("正在生成指定模块的配置文件: %s\n", strings.Join(moduleList, ", "))
		err = generator.GenerateModulesByNames(moduleList...)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "生成配置文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ 配置文件生成成功!")
}

func showHelp() {
	fmt.Println("配置文件生成工具")
	fmt.Println()
	fmt.Println("用法:")
	fmt.Println("  generate-configs [选项]")
	fmt.Println()
	fmt.Println("选项:")
	fmt.Println("  -output string")
	fmt.Println("        输出目录,配置文件将生成到该目录下的pkg文件夹 (默认: 当前目录)")
	fmt.Println("  -modules string")
	fmt.Println("        指定要生成的模块,多个模块用逗号分隔 (为空则生成所有模块)")
	fmt.Println("  -list")
	fmt.Println("        列出所有可用的模块")
	fmt.Println("  -help")
	fmt.Println("        显示此帮助信息")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  # 生成所有模块配置")
	fmt.Println("  generate-configs")
	fmt.Println()
	fmt.Println("  # 只生成 banner 和 cache 模块")
	fmt.Println("  generate-configs -modules banner,cache")
	fmt.Println()
	fmt.Println("  # 生成到指定目录")
	fmt.Println("  generate-configs -output ./output")
	fmt.Println()
	fmt.Println("  # 列出所有可用模块")
	fmt.Println("  generate-configs -list")
}

func listAllModules(generator *goconfig.SmartConfigGenerator) {
	fmt.Println("可用的模块列表:")
	fmt.Println()

	modules := generator.GetModuleList()

	// 直接显示模块名列表
	fmt.Println("  可用模块:")
	for _, name := range modules {
		fmt.Printf("    - %s\n", name)
	}

	fmt.Println()
	fmt.Printf("总计: %d 个模块\n", len(modules))
}
