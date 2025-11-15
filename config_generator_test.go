/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-15 13:53:47
 * @FilePath: \go-config\config_generator_test.go
 * @Description: 通用配置生成器测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-config/pkg/gateway"
	"github.com/kamalyes/go-config/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConfigGenerator_NewConfigGenerator 测试创建配置生成器
func TestConfigGenerator_NewConfigGenerator(t *testing.T) {
	generator := NewConfigGenerator()

	assert.NotNil(t, generator, "配置生成器不应为nil")
	assert.NotNil(t, generator.Logger, "日志记录器不应为nil")
	assert.Equal(t, "./config", generator.OutputDir, "默认输出目录应为./config")
	assert.Equal(t, "config", generator.FilePrefix, "默认文件前缀应为config")
	assert.True(t, generator.EnableConsole, "默认应启用控制台输出")
	assert.True(t, generator.EnableYAML, "默认应启用YAML生成")
	assert.True(t, generator.EnableJSON, "默认应启用JSON生成")
}

// TestConfigGenerator_WithMethods 测试链式配置方法
func TestConfigGenerator_WithMethods(t *testing.T) {
	generator := NewConfigGenerator().
		WithOutputDir("./test-output").
		WithFilePrefix("test-config").
		WithConsoleOutput(false)

	assert.Equal(t, "./test-output", generator.OutputDir, "输出目录应被正确设置")
	assert.Equal(t, "test-config", generator.FilePrefix, "文件前缀应被正确设置")
	assert.False(t, generator.EnableConsole, "控制台输出应被禁用")
}

// TestConfigGenerator_GenerateConfig 测试生成配置文件
func TestConfigGenerator_GenerateConfig(t *testing.T) {
	tempDir := t.TempDir()

	generator := NewConfigGenerator().
		WithOutputDir(tempDir).
		WithFilePrefix("test").
		WithConsoleOutput(false)

	// 测试生成JWT配置
	jwtConfig := jwt.Default()
	err := generator.GenerateConfig(jwtConfig, "JWT")

	require.NoError(t, err, "生成JWT配置不应出错")

	// 验证YAML文件是否生成
	yamlFile := filepath.Join(tempDir, "test-jwt.yaml")
	assert.FileExists(t, yamlFile, "YAML文件应该存在")

	// 验证JSON文件是否生成
	jsonFile := filepath.Join(tempDir, "test-jwt.json")
	assert.FileExists(t, jsonFile, "JSON文件应该存在")

	// 验证文件内容不为空
	yamlContent, err := os.ReadFile(yamlFile)
	require.NoError(t, err, "读取YAML文件不应出错")
	assert.NotEmpty(t, yamlContent, "YAML文件内容不应为空")

	jsonContent, err := os.ReadFile(jsonFile)
	require.NoError(t, err, "读取JSON文件不应出错")
	assert.NotEmpty(t, jsonContent, "JSON文件内容不应为空")
}

// TestConfigGenerator_GenerateFromDefaultFunc 测试从默认函数生成配置
func TestConfigGenerator_GenerateFromDefaultFunc(t *testing.T) {
	tempDir := t.TempDir()

	generator := NewConfigGenerator().
		WithOutputDir(tempDir).
		WithFilePrefix("default-test").
		WithConsoleOutput(false)

	// 测试生成Cache默认配置
	err := generator.GenerateFromDefaultFunc(func() interface{} {
		return cache.Default()
	}, "Cache")

	require.NoError(t, err, "从默认函数生成Cache配置不应出错")

	// 验证文件是否生成
	yamlFile := filepath.Join(tempDir, "default-test-cache.yaml")
	jsonFile := filepath.Join(tempDir, "default-test-cache.json")

	assert.FileExists(t, yamlFile, "Cache YAML文件应该存在")
	assert.FileExists(t, jsonFile, "Cache JSON文件应该存在")
}

// TestConfigGenerator_ExtractStructInfo 测试提取结构体信息
func TestConfigGenerator_ExtractStructInfo(t *testing.T) {
	generator := NewConfigGenerator()

	// 测试提取JWT结构体信息
	jwtType := reflect.TypeOf(jwt.JWT{})
	structInfo, err := generator.ExtractStructInfo(jwtType, "")

	require.NoError(t, err, "提取JWT结构体信息不应出错")
	assert.NotNil(t, structInfo, "结构体信息不应为nil")
	assert.Equal(t, "JWT", structInfo.Name, "结构体名称应为JWT")
	assert.NotEmpty(t, structInfo.Fields, "字段列表不应为空")

	// 验证字段信息
	var modulenameField *FieldInfo
	for _, field := range structInfo.Fields {
		if field.Name == "ModuleName" {
			modulenameField = &field
			break
		}
	}

	assert.NotNil(t, modulenameField, "应该找到ModuleName字段")
	if modulenameField != nil {
		assert.Equal(t, "ModuleName", modulenameField.Name, "字段名应为ModuleName")
		assert.Equal(t, "string", modulenameField.Type, "ModuleName字段类型应为string")
	}
}

// TestConfigGenerator_ExtractFieldInfo 测试提取字段信息
func TestConfigGenerator_ExtractFieldInfo(t *testing.T) {
	generator := NewConfigGenerator()

	// 使用JWT结构体的字段进行测试
	jwtType := reflect.TypeOf(jwt.JWT{})

	// 查找SigningKey字段
	var signingKeyField reflect.StructField
	found := false
	for i := 0; i < jwtType.NumField(); i++ {
		field := jwtType.Field(i)
		if field.Name == "SigningKey" {
			signingKeyField = field
			found = true
			break
		}
	}

	require.True(t, found, "应该找到SigningKey字段")

	// 测试提取字段信息
	comments := make(map[string]string)
	fieldInfo := generator.extractFieldInfo(signingKeyField, comments)

	assert.Equal(t, "SigningKey", fieldInfo.Name, "字段名应为SigningKey")
	assert.Equal(t, "string", fieldInfo.Type, "字段类型应为string")
	assert.NotEmpty(t, fieldInfo.Tags, "标签信息不应为空")

	// 验证JSON标签
	if jsonTag, exists := fieldInfo.Tags["json"]; exists {
		assert.NotEmpty(t, jsonTag, "JSON标签不应为空")
	}

	// 验证YAML标签
	if yamlTag, exists := fieldInfo.Tags["yaml"]; exists {
		assert.NotEmpty(t, yamlTag, "YAML标签不应为空")
	}
}

// TestConfigGenerator_ParseFileComments 测试解析文件注释
func TestConfigGenerator_ParseFileComments(t *testing.T) {
	generator := NewConfigGenerator()

	// 创建临时Go文件进行测试
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test_struct.go")

	testContent := `package test

// TestStruct 测试结构体
type TestStruct struct {
	Name    string ` + "`" + `json:"name" yaml:"name"` + "`" + ` // 名称字段
	Age     int    ` + "`" + `json:"age" yaml:"age"` + "`" + `   // 年龄字段
	Enabled bool   ` + "`" + `json:"enabled"` + "`" + `         // 启用状态
}
`

	err := os.WriteFile(testFile, []byte(testContent), 0644)
	require.NoError(t, err, "写入测试文件不应出错")

	// 解析注释
	comments, err := generator.parseFileComments(testFile)
	require.NoError(t, err, "解析文件注释不应出错")

	// 验证注释提取结果
	assert.Contains(t, comments, "Name", "应该包含Name字段的注释")
	assert.Contains(t, comments, "Age", "应该包含Age字段的注释")
	assert.Contains(t, comments, "Enabled", "应该包含Enabled字段的注释")

	assert.Equal(t, "名称字段", comments["Name"], "Name字段注释应为'名称字段'")
	assert.Equal(t, "年龄字段", comments["Age"], "Age字段注释应为'年龄字段'")
	assert.Equal(t, "启用状态", comments["Enabled"], "Enabled字段注释应为'启用状态'")
}

// TestConfigGenerator_GenerateConfigDocumentation 测试生成配置文档
func TestConfigGenerator_GenerateConfigDocumentation(t *testing.T) {
	tempDir := t.TempDir()

	generator := NewConfigGenerator().
		WithOutputDir(tempDir).
		WithFilePrefix("doc-test").
		WithConsoleOutput(false)

	// 创建测试结构体信息
	structInfo := &StructInfo{
		Name:    "TestConfig",
		Package: "test",
		Fields: []FieldInfo{
			{
				Name:     "ModuleName",
				Type:     "string",
				Comment:  "模块名称",
				JSONName: "module_name",
				YAMLName: "module-name",
			},
			{
				Name:     "Enabled",
				Type:     "bool",
				Comment:  "是否启用",
				JSONName: "enabled",
				YAMLName: "enabled",
			},
		},
	}

	err := generator.GenerateConfigDocumentation(structInfo)
	require.NoError(t, err, "生成配置文档不应出错")

	// 验证文档文件是否生成
	docFile := filepath.Join(tempDir, "doc-test-testconfig-doc.md")
	assert.FileExists(t, docFile, "文档文件应该存在")

	// 验证文档内容
	content, err := os.ReadFile(docFile)
	require.NoError(t, err, "读取文档文件不应出错")

	docContent := string(content)
	assert.Contains(t, docContent, "# TestConfig 配置文档", "文档应包含标题")
	assert.Contains(t, docContent, "ModuleName", "文档应包含ModuleName字段")
	assert.Contains(t, docContent, "模块名称", "文档应包含字段注释")
	assert.Contains(t, docContent, "| 字段名 | 类型 |", "文档应包含表格头")
}

// TestConfigGenerator_BuildDocumentationContent 测试构建文档内容
func TestConfigGenerator_BuildDocumentationContent(t *testing.T) {
	generator := NewConfigGenerator()

	structInfo := &StructInfo{
		Name:    "Gateway",
		Package: "github.com/kamalyes/go-config/pkg/gateway",
		Comment: "网关统一配置",
		Fields: []FieldInfo{
			{
				Name:     "ModuleName",
				Type:     "string",
				Comment:  "模块名称",
				JSONName: "module_name",
				YAMLName: "module-name",
			},
			{
				Name:     "Name",
				Type:     "string",
				Comment:  "网关名称",
				JSONName: "name",
				YAMLName: "name",
			},
		},
	}

	docContent := generator.buildDocumentationContent(structInfo)

	assert.NotEmpty(t, docContent, "文档内容不应为空")
	assert.Contains(t, docContent, "# Gateway 配置文档", "应包含结构体名称标题")
	assert.Contains(t, docContent, "github.com/kamalyes/go-config/pkg/gateway", "应包含包路径")
	assert.Contains(t, docContent, "ModuleName", "应包含字段名称")
	assert.Contains(t, docContent, "模块名称", "应包含字段注释")
	assert.Contains(t, docContent, "module_name", "应包含JSON标签")
	assert.Contains(t, docContent, "module-name", "应包含YAML标签")
}

// BenchmarkConfigGenerator_GenerateConfig 性能测试
func BenchmarkConfigGenerator_GenerateConfig(b *testing.B) {
	tempDir := b.TempDir()

	generator := NewConfigGenerator().
		WithOutputDir(tempDir).
		WithFilePrefix("benchmark").
		WithConsoleOutput(false)

	gatewayConfig := gateway.Default()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := generator.GenerateConfig(gatewayConfig, "Gateway")
		if err != nil {
			b.Fatalf("生成配置失败: %v", err)
		}
	}
}

// TestConfigGenerator_EnsureDir 测试确保目录存在
func TestConfigGenerator_EnsureDir(t *testing.T) {
	generator := NewConfigGenerator()
	tempDir := t.TempDir()

	testDir := filepath.Join(tempDir, "nested", "test", "dir")

	err := generator.ensureDir(testDir)
	assert.NoError(t, err, "创建嵌套目录不应出错")
	assert.DirExists(t, testDir, "嵌套目录应该存在")
}

// TestConfigGenerator_MultipleCalls 测试多次调用生成器
func TestConfigGenerator_MultipleCalls(t *testing.T) {
	tempDir := t.TempDir()

	generator := NewConfigGenerator().
		WithOutputDir(tempDir).
		WithFilePrefix("multi").
		WithConsoleOutput(false)

	// 生成多个不同的配置
	configs := map[string]func() interface{}{
		"JWT":     func() interface{} { return jwt.Default() },
		"Cache":   func() interface{} { return cache.Default() },
		"Gateway": func() interface{} { return gateway.Default() },
	}

	for name, defaultFunc := range configs {
		err := generator.GenerateFromDefaultFunc(defaultFunc, name)
		assert.NoError(t, err, "生成%s配置不应出错", name)

		// 验证文件存在
		yamlFile := filepath.Join(tempDir, "multi-"+strings.ToLower(name)+".yaml")
		jsonFile := filepath.Join(tempDir, "multi-"+strings.ToLower(name)+".json")

		assert.FileExists(t, yamlFile, "%s YAML文件应该存在", name)
		assert.FileExists(t, jsonFile, "%s JSON文件应该存在", name)
	}
}
