package goconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestSmartConfigGeneratorBasic 测试基础功能
func TestSmartConfigGeneratorBasic(t *testing.T) {
	generator := NewSmartConfigGenerator("./test_output")

	// 测试获取模块列表
	modules := generator.GetModuleList()
	if len(modules) == 0 {
		t.Error("模块列表为空")
	}

	fmt.Printf("注册的模块数量: %d\n", len(modules))
	fmt.Printf("前5个模块: %v\n", modules[:min(5, len(modules))])

	// 测试启用/禁用模块
	testModule := modules[0]

	err := generator.DisableModule(testModule)
	if err != nil {
		t.Errorf("禁用模块失败: %v", err)
	}

	err = generator.EnableModule(testModule)
	if err != nil {
		t.Errorf("启用模块失败: %v", err)
	}
}

// TestSmartConfigGeneratorValidation 测试验证功能
func TestSmartConfigGeneratorValidation(t *testing.T) {
	generator := NewSmartConfigGenerator("./test_output")

	// 测试验证所有模块
	err := generator.ValidateAllModules()
	if err != nil {
		t.Errorf("模块验证失败: %v", err)
	}

	fmt.Println("所有模块验证通过")
}

// TestSmartConfigGeneratorGeneration 测试生成功能
func TestSmartConfigGeneratorGeneration(t *testing.T) {
	// 创建临时输出目录
	tempDir := "./test_output_generation"
	defer os.RemoveAll(tempDir)

	generator := NewSmartConfigGenerator(tempDir)

	// 只启用几个模块进行测试
	err := generator.EnableOnlyModules("health", "redis", "monitoring")
	if err != nil {
		t.Errorf("启用模块失败: %v", err)
	}

	// 生成配置文件
	err = generator.GenerateAllConfigs()
	if err != nil {
		t.Errorf("生成配置文件失败: %v", err)
	}

	// 检查文件是否生成
	healthYamlPath := filepath.Join(tempDir, "pkg", "health", "health.yaml")
	if _, err := os.Stat(healthYamlPath); os.IsNotExist(err) {
		t.Errorf("health.yaml 文件未生成")
	}

	redisYamlPath := filepath.Join(tempDir, "pkg", "redis", "redis.yaml")
	if _, err := os.Stat(redisYamlPath); os.IsNotExist(err) {
		t.Errorf("redis.yaml 文件未生成")
	}

	fmt.Println("配置文件生成测试通过")
}

// TestSmartConfigGeneratorOverwrite 测试覆盖保护功能
func TestSmartConfigGeneratorOverwrite(t *testing.T) {
	// 创建临时输出目录
	tempDir := "./test_output_overwrite"
	defer os.RemoveAll(tempDir)

	generator := NewSmartConfigGenerator(tempDir)
	generator.OverwriteExisting = false // 禁用覆盖

	// 只启用一个模块
	err := generator.EnableOnlyModules("health")
	if err != nil {
		t.Errorf("启用模块失败: %v", err)
	}

	// 第一次生成
	err = generator.GenerateAllConfigs()
	if err != nil {
		t.Errorf("第一次生成失败: %v", err)
	}

	// 第二次生成（应该跳过已存在的文件）
	err = generator.GenerateAllConfigs()
	if err != nil {
		t.Errorf("第二次生成失败: %v", err)
	}

	fmt.Println("覆盖保护测试通过")
}

// TestSmartConfigGeneratorBackup 测试备份功能
func TestSmartConfigGeneratorBackup(t *testing.T) {
	// 创建临时输出目录
	tempDir := "./test_output_backup"
	defer os.RemoveAll(tempDir)

	generator := NewSmartConfigGenerator(tempDir)
	generator.BackupExisting = true
	generator.OverwriteExisting = true

	// 只启用一个模块
	err := generator.EnableOnlyModules("health")
	if err != nil {
		t.Errorf("启用模块失败: %v", err)
	}

	// 第一次生成
	err = generator.GenerateAllConfigs()
	if err != nil {
		t.Errorf("第一次生成失败: %v", err)
	}

	// 第二次生成（应该创建备份）
	err = generator.GenerateAllConfigs()
	if err != nil {
		t.Errorf("第二次生成失败: %v", err)
	}

	fmt.Println("备份功能测试通过")
}

// TestSmartConfigGeneratorCleanup 测试清理功能
func TestSmartConfigGeneratorCleanup(t *testing.T) {
	// 创建临时输出目录
	tempDir := "./test_output_cleanup"
	defer os.RemoveAll(tempDir)

	generator := NewSmartConfigGenerator(tempDir)
	generator.BackupExisting = true
	generator.OverwriteExisting = true

	// 只启用一个模块
	err := generator.EnableOnlyModules("health")
	if err != nil {
		t.Errorf("启用模块失败: %v", err)
	}

	// 生成几次来创建备份文件
	for i := 0; i < 3; i++ {
		err = generator.GenerateAllConfigs()
		if err != nil {
			t.Errorf("第%d次生成失败: %v", i+1, err)
		}
		time.Sleep(10 * time.Millisecond) // 确保文件时间戳不同
	}

	// 清理1纳秒以前的备份文件（应该清理所有备份）
	err = generator.CleanupBackupFiles(1 * time.Nanosecond)
	if err != nil {
		t.Errorf("清理备份文件失败: %v", err)
	}

	fmt.Println("清理功能测试通过")
}

// TestSmartConfigGeneratorModuleInfo 测试模块信息功能
func TestSmartConfigGeneratorModuleInfo(t *testing.T) {
	generator := NewSmartConfigGenerator("./test_output")

	modules := generator.GetModuleList()
	if len(modules) == 0 {
		t.Error("模块列表为空")
		return
	}

	// 测试获取模块信息
	testModule := modules[0]
	info, err := generator.GetModuleInfo(testModule)
	if err != nil {
		t.Errorf("获取模块信息失败: %v", err)
	}

	if info.Name != testModule {
		t.Errorf("模块名称不匹配: 期望 %s, 实际 %s", testModule, info.Name)
	}

	// 测试更新模块配置
	updates := map[string]interface{}{
		"description": "测试描述",
		"enabled":     false,
	}

	err = generator.UpdateModuleConfig(testModule, updates)
	if err != nil {
		t.Errorf("更新模块配置失败: %v", err)
	}

	// 验证更新
	updatedInfo, err := generator.GetModuleInfo(testModule)
	if err != nil {
		t.Errorf("获取更新后模块信息失败: %v", err)
	}

	if updatedInfo.Description != "测试描述" {
		t.Errorf("描述更新失败: 期望 '测试描述', 实际 '%s'", updatedInfo.Description)
	}

	if updatedInfo.Enabled != false {
		t.Error("启用状态更新失败")
	}

	fmt.Println("模块信息功能测试通过")
}

// TestSmartConfigGeneratorPrintStatus 测试状态打印功能
func TestSmartConfigGeneratorPrintStatus(t *testing.T) {
	generator := NewSmartConfigGenerator("./test_output")

	// 启用部分模块
	err := generator.EnableOnlyModules("health", "redis", "monitoring")
	if err != nil {
		t.Errorf("启用模块失败: %v", err)
	}

	// 打印状态（主要是视觉验证）
	fmt.Println("\n=== 模块状态打印测试 ===")
	generator.PrintModuleStatus()

	// 检查启用的模块数量
	enabledModules := generator.GetEnabledModules()
	if len(enabledModules) != 3 {
		t.Errorf("启用的模块数量不正确: 期望 3, 实际 %d", len(enabledModules))
	}

	fmt.Println("状态打印功能测试通过")
}

// 辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 运行所有测试的主函数
func main() {
	fmt.Println("开始运行智能配置生成器测试...")

	// 可以在这里运行特定的测试
	generator := NewSmartConfigGenerator("./example_output")

	fmt.Println("\n1. 打印模块状态:")
	generator.PrintModuleStatus()

	fmt.Println("\n2. 验证所有模块:")
	if err := generator.ValidateAllModules(); err != nil {
		fmt.Printf("验证失败: %v\n", err)
	} else {
		fmt.Println("所有模块验证通过")
	}

	fmt.Println("\n3. 启用部分模块并生成配置:")
	if err := generator.EnableOnlyModules("health", "redis", "monitoring", "jwt"); err != nil {
		fmt.Printf("启用模块失败: %v\n", err)
	} else {
		fmt.Println("已启用模块: health, redis, monitoring, jwt")

		if err := generator.GenerateAllConfigs(); err != nil {
			fmt.Printf("生成配置失败: %v\n", err)
		} else {
			fmt.Println("配置文件生成成功")
		}
	}

	fmt.Println("\n测试完成!")
}
