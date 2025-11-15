/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-15 13:47:26
 * @FilePath: \go-config\config_generator.go
 * @Description: 通用配置生成器，支持所有go-config结构体
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/kamalyes/go-config/internal"
	gologger "github.com/kamalyes/go-logger"
)

// ConfigGenerator 通用配置生成器
type ConfigGenerator struct {
	OutputDir     string           // 输出目录
	FilePrefix    string           // 文件名前缀
	EnableConsole bool             // 是否在控制台输出
	EnableYAML    bool             // 是否生成YAML文件
	EnableJSON    bool             // 是否生成JSON文件
	Logger        *gologger.Logger // 日志记录器
}

// FieldInfo 字段信息
type FieldInfo struct {
	Name     string            // 字段名
	Type     string            // 字段类型
	Tags     map[string]string // 标签信息
	Comment  string            // 注释信息
	JSONName string            // JSON标签名
	YAMLName string            // YAML标签名
	MapName  string            // Mapstructure标签名
}

// StructInfo 结构体信息
type StructInfo struct {
	Name     string      // 结构体名称
	Package  string      // 包名
	Comment  string      // 结构体注释
	Fields   []FieldInfo // 字段列表
	FilePath string      // 文件路径
}

// NewConfigGenerator 创建新的配置生成器
func NewConfigGenerator() *ConfigGenerator {
	config := gologger.DefaultConfig()
	logger := gologger.NewLogger(config)

	return &ConfigGenerator{
		OutputDir:     "./config",
		FilePrefix:    "config",
		EnableConsole: true,
		EnableYAML:    true,
		EnableJSON:    true,
		Logger:        logger,
	}
}

// WithOutputDir 设置输出目录
func (cg *ConfigGenerator) WithOutputDir(dir string) *ConfigGenerator {
	cg.OutputDir = dir
	return cg
}

// WithFilePrefix 设置文件名前缀
func (cg *ConfigGenerator) WithFilePrefix(prefix string) *ConfigGenerator {
	cg.FilePrefix = prefix
	return cg
}

// WithConsoleOutput 设置是否在控制台输出
func (cg *ConfigGenerator) WithConsoleOutput(enable bool) *ConfigGenerator {
	cg.EnableConsole = enable
	return cg
}

// WithLogger 设置自定义日志记录器
func (cg *ConfigGenerator) WithLogger(logger *gologger.Logger) *ConfigGenerator {
	cg.Logger = logger
	return cg
}

// GenerateConfig 生成配置文件
func (cg *ConfigGenerator) GenerateConfig(config interface{}, configName string) error {
	cg.Logger.InfoKV("开始生成配置文件",
		"config_name", configName,
		"output_dir", cg.OutputDir)

	// 确保输出目录存在
	if err := cg.ensureDir(cg.OutputDir); err != nil {
		cg.Logger.ErrorKV("创建输出目录失败",
			"error", err.Error(),
			"dir", cg.OutputDir)
		return err
	}

	// 生成YAML文件
	if cg.EnableYAML {
		if err := cg.generateYAMLFile(config, configName); err != nil {
			cg.Logger.ErrorKV("生成YAML文件失败",
				"error", err.Error(),
				"config_name", configName)
			return err
		}
	}

	// 生成JSON文件
	if cg.EnableJSON {
		if err := cg.generateJSONFile(config, configName); err != nil {
			cg.Logger.ErrorKV("生成JSON文件失败",
				"error", err.Error(),
				"config_name", configName)
			return err
		}
	}

	cg.Logger.InfoKV("配置文件生成完成",
		"config_name", configName)

	return nil
}

// GenerateFromDefaultFunc 从默认函数生成配置
func (cg *ConfigGenerator) GenerateFromDefaultFunc(defaultFunc func() interface{}, configName string) error {
	config := defaultFunc()
	return cg.GenerateConfig(config, configName)
}

// ExtractStructInfo 提取结构体信息包括注释
func (cg *ConfigGenerator) ExtractStructInfo(structType reflect.Type, filePath string) (*StructInfo, error) {
	cg.Logger.DebugKV("提取结构体信息",
		"struct_name", structType.Name(),
		"file_path", filePath)

	structInfo := &StructInfo{
		Name:     structType.Name(),
		Package:  structType.PkgPath(),
		Fields:   make([]FieldInfo, 0),
		FilePath: filePath,
	}

	// 解析源文件获取注释
	comments, err := cg.parseFileComments(filePath)
	if err != nil {
		cg.Logger.WarnKV("解析文件注释失败",
			"error", err.Error(),
			"file_path", filePath)
	}

	// 提取字段信息
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldInfo := cg.extractFieldInfo(field, comments)
		structInfo.Fields = append(structInfo.Fields, fieldInfo)
	}

	return structInfo, nil
}

// extractFieldInfo 提取字段信息
func (cg *ConfigGenerator) extractFieldInfo(field reflect.StructField, comments map[string]string) FieldInfo {
	fieldInfo := FieldInfo{
		Name: field.Name,
		Type: field.Type.String(),
		Tags: make(map[string]string),
	}

	// 解析标签
	if jsonTag := field.Tag.Get("json"); jsonTag != "" {
		fieldInfo.JSONName = strings.Split(jsonTag, ",")[0]
		fieldInfo.Tags["json"] = jsonTag
	}

	if yamlTag := field.Tag.Get("yaml"); yamlTag != "" {
		fieldInfo.YAMLName = strings.Split(yamlTag, ",")[0]
		fieldInfo.Tags["yaml"] = yamlTag
	}

	if mapTag := field.Tag.Get("mapstructure"); mapTag != "" {
		fieldInfo.MapName = strings.Split(mapTag, ",")[0]
		fieldInfo.Tags["mapstructure"] = mapTag
	}

	// 从注释映射中获取注释
	if comment, exists := comments[field.Name]; exists {
		fieldInfo.Comment = comment
	}

	return fieldInfo
}

// parseFileComments 解析文件注释
func (cg *ConfigGenerator) parseFileComments(filePath string) (map[string]string, error) {
	comments := make(map[string]string)

	if filePath == "" {
		return comments, nil
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return comments, err
	}

	// 遍历AST查找结构体和字段注释
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.StructType:
			if x.Fields != nil {
				for _, field := range x.Fields.List {
					if field.Comment != nil && len(field.Names) > 0 {
						fieldName := field.Names[0].Name
						comment := strings.TrimSpace(field.Comment.Text())
						// 提取 // 后面的注释内容
						if strings.Contains(comment, "//") {
							parts := strings.Split(comment, "//")
							if len(parts) > 1 {
								comment = strings.TrimSpace(parts[len(parts)-1])
							}
						}
						comments[fieldName] = comment
					}
				}
			}
		}
		return true
	})

	return comments, nil
}

// generateYAMLFile 生成YAML文件
func (cg *ConfigGenerator) generateYAMLFile(config interface{}, configName string) error {
	yamlContent, err := internal.ExportToYAML(config)
	if err != nil {
		return fmt.Errorf("生成YAML内容失败: %w", err)
	}

	filename := filepath.Join(cg.OutputDir, fmt.Sprintf("%s-%s.yaml", cg.FilePrefix, strings.ToLower(configName)))
	if err := os.WriteFile(filename, []byte(yamlContent), 0644); err != nil {
		return fmt.Errorf("保存YAML文件失败: %w", err)
	}

	if cg.EnableConsole {
		cg.Logger.InfoKV("YAML文件生成成功",
			"file_path", filename,
			"config_name", configName)
	}

	return nil
}

// generateJSONFile 生成JSON文件
func (cg *ConfigGenerator) generateJSONFile(config interface{}, configName string) error {
	jsonContent, err := internal.ExportToJSON(config)
	if err != nil {
		return fmt.Errorf("生成JSON内容失败: %w", err)
	}

	filename := filepath.Join(cg.OutputDir, fmt.Sprintf("%s-%s.json", cg.FilePrefix, strings.ToLower(configName)))
	if err := os.WriteFile(filename, []byte(jsonContent), 0644); err != nil {
		return fmt.Errorf("保存JSON文件失败: %w", err)
	}

	if cg.EnableConsole {
		cg.Logger.InfoKV("JSON文件生成成功",
			"file_path", filename,
			"config_name", configName)
	}

	return nil
}

// ensureDir 确保目录存在
func (cg *ConfigGenerator) ensureDir(dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}
	return nil
}

// GenerateConfigDocumentation 生成配置文档
func (cg *ConfigGenerator) GenerateConfigDocumentation(structInfo *StructInfo) error {
	cg.Logger.InfoKV("生成配置文档",
		"struct_name", structInfo.Name)

	docContent := cg.buildDocumentationContent(structInfo)
	filename := filepath.Join(cg.OutputDir, fmt.Sprintf("%s-%s-doc.md", cg.FilePrefix, strings.ToLower(structInfo.Name)))

	if err := os.WriteFile(filename, []byte(docContent), 0644); err != nil {
		return fmt.Errorf("保存文档文件失败: %w", err)
	}

	cg.Logger.InfoKV("配置文档生成成功",
		"file_path", filename)

	return nil
}

// buildDocumentationContent 构建文档内容
func (cg *ConfigGenerator) buildDocumentationContent(structInfo *StructInfo) string {
	var doc strings.Builder

	doc.WriteString(fmt.Sprintf("# %s 配置文档\n\n", structInfo.Name))
	doc.WriteString(fmt.Sprintf("**包路径**: %s\n\n", structInfo.Package))

	if structInfo.Comment != "" {
		doc.WriteString(fmt.Sprintf("**描述**: %s\n\n", structInfo.Comment))
	}

	doc.WriteString("## 配置字段\n\n")
	doc.WriteString("| 字段名 | 类型 | JSON标签 | YAML标签 | 描述 |\n")
	doc.WriteString("|--------|------|----------|----------|------|\n")

	for _, field := range structInfo.Fields {
		jsonName := field.JSONName
		if jsonName == "" {
			jsonName = "-"
		}

		yamlName := field.YAMLName
		if yamlName == "" {
			yamlName = "-"
		}

		comment := field.Comment
		if comment == "" {
			comment = "-"
		}

		doc.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
			field.Name, field.Type, jsonName, yamlName, comment))
	}

	return doc.String()
}
