/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-22 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-11 15:36:15
 * @FilePath: \go-config\config_init_generator.go
 * @Description: 配置文件自动生成器 - 智能生成和更新所有模块的配置文件
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/kamalyes/go-config/pkg/alerting"
	"github.com/kamalyes/go-config/pkg/banner"
	"github.com/kamalyes/go-config/pkg/breaker"
	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-config/pkg/captcha"
	"github.com/kamalyes/go-config/pkg/consul"
	"github.com/kamalyes/go-config/pkg/cors"
	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-config/pkg/elasticsearch"
	"github.com/kamalyes/go-config/pkg/email"
	"github.com/kamalyes/go-config/pkg/etcd"
	"github.com/kamalyes/go-config/pkg/ftp"
	"github.com/kamalyes/go-config/pkg/gateway"
	"github.com/kamalyes/go-config/pkg/grafana"
	"github.com/kamalyes/go-config/pkg/health"
	"github.com/kamalyes/go-config/pkg/i18n"
	"github.com/kamalyes/go-config/pkg/jaeger"
	"github.com/kamalyes/go-config/pkg/jobs"
	"github.com/kamalyes/go-config/pkg/jwt"
	"github.com/kamalyes/go-config/pkg/kafka"
	"github.com/kamalyes/go-config/pkg/logging"
	"github.com/kamalyes/go-config/pkg/metrics"
	"github.com/kamalyes/go-config/pkg/middleware"
	"github.com/kamalyes/go-config/pkg/monitoring"
	"github.com/kamalyes/go-config/pkg/oss"
	"github.com/kamalyes/go-config/pkg/pay"
	"github.com/kamalyes/go-config/pkg/pprof"
	"github.com/kamalyes/go-config/pkg/prometheus"
	"github.com/kamalyes/go-config/pkg/queue"
	"github.com/kamalyes/go-config/pkg/ratelimit"
	"github.com/kamalyes/go-config/pkg/recovery"
	"github.com/kamalyes/go-config/pkg/redis"
	"github.com/kamalyes/go-config/pkg/requestid"
	"github.com/kamalyes/go-config/pkg/restful"
	"github.com/kamalyes/go-config/pkg/rpcclient"
	"github.com/kamalyes/go-config/pkg/rpcserver"
	"github.com/kamalyes/go-config/pkg/security"
	"github.com/kamalyes/go-config/pkg/signature"
	"github.com/kamalyes/go-config/pkg/sms"
	"github.com/kamalyes/go-config/pkg/smtp"
	"github.com/kamalyes/go-config/pkg/sts"
	"github.com/kamalyes/go-config/pkg/swagger"
	"github.com/kamalyes/go-config/pkg/timeout"
	"github.com/kamalyes/go-config/pkg/tracing"
	"github.com/kamalyes/go-config/pkg/wsc"
	"github.com/kamalyes/go-config/pkg/youzan"
	"github.com/kamalyes/go-config/pkg/zap"
	gologger "github.com/kamalyes/go-logger"
	"gopkg.in/yaml.v3"
)

// ModuleConfig 模块配置信息
type ModuleConfig struct {
	Name          string             // 模块名称
	PackageName   string             // 包名
	DefaultFunc   func() interface{} // 默认配置函数
	OutputSubDir  string             // 输出子目录
	SourceFile    string             // 源文件路径
	StructName    string             // 结构体名称
	Description   string             // 模块描述
	Enabled       bool               // 是否启用此模块生成
	LastGenerated time.Time          // 最后生成时间
}

// SmartConfigGenerator 智能配置生成器
type SmartConfigGenerator struct {
	BaseOutputDir     string                       // 基础输出目录
	Logger            *gologger.Logger             // 日志记录器
	ModuleRegistry    map[string]ModuleConfig      // 模块注册表
	ForceRegenerate   bool                         // 是否强制重新生成
	IncludeComments   bool                         // 是否包含注释
	GenerateJSON      bool                         // 是否生成JSON
	GenerateYAML      bool                         // 是否生成YAML
	BackupExisting    bool                         // 是否备份现有文件
	OverwriteExisting bool                         // 是否覆盖现有文件
	commentCache      map[string]map[string]string // 注释缓存: packagePath -> fieldName -> comment
}

// NewSmartConfigGenerator 创建新的智能配置生成器
func NewSmartConfigGenerator(baseOutputDir string) *SmartConfigGenerator {
	config := gologger.DefaultConfig()
	logger := gologger.NewLogger(config)

	generator := &SmartConfigGenerator{
		BaseOutputDir:     baseOutputDir,
		Logger:            logger,
		ModuleRegistry:    make(map[string]ModuleConfig),
		ForceRegenerate:   false,
		IncludeComments:   true,
		GenerateJSON:      true,
		GenerateYAML:      true,
		BackupExisting:    true,
		OverwriteExisting: true,
		commentCache:      make(map[string]map[string]string),
	}

	// 自动注册所有模块
	generator.registerAllModules()
	return generator
}

// WithLogger 设置日志记录器
func (sg *SmartConfigGenerator) WithLogger(logger *gologger.Logger) *SmartConfigGenerator {
	sg.Logger = logger
	return sg
}

// WithForceRegenerate 设置是否强制重新生成
func (sg *SmartConfigGenerator) WithForceRegenerate(force bool) *SmartConfigGenerator {
	sg.ForceRegenerate = force
	return sg
}

// WithIncludeComments 设置是否包含注释
func (sg *SmartConfigGenerator) WithIncludeComments(include bool) *SmartConfigGenerator {
	sg.IncludeComments = include
	return sg
}

// WithBackupExisting 设置是否备份现有文件
func (sg *SmartConfigGenerator) WithBackupExisting(backup bool) *SmartConfigGenerator {
	sg.BackupExisting = backup
	return sg
}

// registerAllModules 自动注册所有模块
func (sg *SmartConfigGenerator) registerAllModules() {
	modules := []ModuleConfig{
		{Name: "alerting", PackageName: "alerting", DefaultFunc: func() interface{} { return alerting.Default() }, OutputSubDir: "alerting", Description: "告警模块", Enabled: true},
		{Name: "banner", PackageName: "banner", DefaultFunc: func() interface{} { return banner.Default() }, OutputSubDir: "banner", Description: "Banner显示模块", Enabled: true},
		{Name: "breaker", PackageName: "breaker", DefaultFunc: func() interface{} { return breaker.Default() }, OutputSubDir: "breaker", Description: "熔断器模块", Enabled: true},
		{Name: "cache", PackageName: "cache", DefaultFunc: func() interface{} { return cache.Default() }, OutputSubDir: "cache", Description: "缓存模块", Enabled: true},
		{Name: "captcha", PackageName: "captcha", DefaultFunc: func() interface{} { return captcha.Default() }, OutputSubDir: "captcha", Description: "验证码模块", Enabled: true},
		{Name: "consul", PackageName: "consul", DefaultFunc: func() interface{} { return consul.Default() }, OutputSubDir: "consul", Description: "Consul服务发现模块", Enabled: true},
		{Name: "cors", PackageName: "cors", DefaultFunc: func() interface{} { return cors.Default() }, OutputSubDir: "cors", Description: "CORS跨域模块", Enabled: true},
		{Name: "database", PackageName: "database", DefaultFunc: func() interface{} { return database.DefaultDatabaseConfig() }, OutputSubDir: "database", Description: "数据库模块", Enabled: true},
		{Name: "mysql", PackageName: "database", DefaultFunc: func() interface{} { return database.DefaultMySQL() }, OutputSubDir: "database", Description: "MySQL数据库", Enabled: true},
		{Name: "postgresql", PackageName: "database", DefaultFunc: func() interface{} { return database.DefaultPostgreSQL() }, OutputSubDir: "database", Description: "PostgreSQL数据库", Enabled: true},
		{Name: "sqlite", PackageName: "database", DefaultFunc: func() interface{} { return database.DefaultSQLite() }, OutputSubDir: "database", Description: "SQLite数据库", Enabled: true},
		{Name: "elasticsearch", PackageName: "elasticsearch", DefaultFunc: func() interface{} { return elasticsearch.Default() }, OutputSubDir: "elasticsearch", Description: "Elasticsearch搜索引擎", Enabled: true},
		{Name: "email", PackageName: "email", DefaultFunc: func() interface{} { return email.Default() }, OutputSubDir: "email", Description: "邮件发送模块", Enabled: true},
		{Name: "etcd", PackageName: "etcd", DefaultFunc: func() interface{} { return etcd.Default() }, OutputSubDir: "etcd", Description: "Etcd分布式键值存储", Enabled: true},
		{Name: "ftp", PackageName: "ftp", DefaultFunc: func() interface{} { return ftp.Default() }, OutputSubDir: "ftp", Description: "FTP文件传输模块", Enabled: true},
		{Name: "gateway", PackageName: "gateway", DefaultFunc: func() interface{} { return gateway.Default() }, OutputSubDir: "gateway", Description: "API网关模块", Enabled: true},
		{Name: "grafana", PackageName: "grafana", DefaultFunc: func() interface{} { return grafana.Default() }, OutputSubDir: "grafana", Description: "Grafana监控面板", Enabled: true},
		{Name: "health", PackageName: "health", DefaultFunc: func() interface{} { return health.Default() }, OutputSubDir: "health", Description: "健康检查模块", Enabled: true},
		{Name: "i18n", PackageName: "i18n", DefaultFunc: func() interface{} { return i18n.Default() }, OutputSubDir: "i18n", Description: "国际化模块", Enabled: true},
		{Name: "jaeger", PackageName: "jaeger", DefaultFunc: func() interface{} { return jaeger.Default() }, OutputSubDir: "jaeger", Description: "Jaeger链路追踪", Enabled: true},
		{Name: "jwt", PackageName: "jwt", DefaultFunc: func() interface{} { return jwt.Default() }, OutputSubDir: "jwt", Description: "JWT认证模块", Enabled: true},
		{Name: "kafka", PackageName: "kafka", DefaultFunc: func() interface{} { return kafka.Default() }, OutputSubDir: "kafka", Description: "Kafka消息队列", Enabled: true},
		{Name: "logging", PackageName: "logging", DefaultFunc: func() interface{} { return logging.Default() }, OutputSubDir: "logging", Description: "日志记录模块", Enabled: true},
		{Name: "metrics", PackageName: "metrics", DefaultFunc: func() interface{} { return metrics.Default() }, OutputSubDir: "metrics", Description: "指标收集模块", Enabled: true},
		{Name: "middleware", PackageName: "middleware", DefaultFunc: func() interface{} { return middleware.Default() }, OutputSubDir: "middleware", Description: "中间件模块", Enabled: true},
		{Name: "monitoring", PackageName: "monitoring", DefaultFunc: func() interface{} { return monitoring.Default() }, OutputSubDir: "monitoring", Description: "监控模块", Enabled: true},
		{Name: "oss", PackageName: "oss", DefaultFunc: func() interface{} { return oss.DefaultOSSConfig() }, OutputSubDir: "oss", Description: "对象存储模块", Enabled: true},
		{Name: "alipay", PackageName: "pay", DefaultFunc: func() interface{} { return pay.DefaultAliPay() }, OutputSubDir: "pay", Description: "支付宝支付", Enabled: true},
		{Name: "wechatpay", PackageName: "pay", DefaultFunc: func() interface{} { return pay.DefaultWechatPay() }, OutputSubDir: "pay", Description: "微信支付", Enabled: true},
		{Name: "pprof", PackageName: "pprof", DefaultFunc: func() interface{} { return pprof.Default() }, OutputSubDir: "pprof", Description: "性能分析模块", Enabled: true},
		{Name: "prometheus", PackageName: "prometheus", DefaultFunc: func() interface{} { return prometheus.Default() }, OutputSubDir: "prometheus", Description: "Prometheus指标", Enabled: true},
		{Name: "mqtt", PackageName: "queue", DefaultFunc: func() interface{} { return queue.Default() }, OutputSubDir: "queue", Description: "MQTT消息队列", Enabled: true},
		{Name: "ratelimit", PackageName: "ratelimit", DefaultFunc: func() interface{} { return ratelimit.Default() }, OutputSubDir: "ratelimit", Description: "限流模块", Enabled: true},
		{Name: "recovery", PackageName: "recovery", DefaultFunc: func() interface{} { return recovery.Default() }, OutputSubDir: "recovery", Description: "错误恢复模块", Enabled: true},
		{Name: "redis", PackageName: "redis", DefaultFunc: func() interface{} { return redis.NewRedis(&redis.Redis{}) }, OutputSubDir: "redis", Description: "Redis缓存", Enabled: true},
		{Name: "requestid", PackageName: "requestid", DefaultFunc: func() interface{} { return requestid.Default() }, OutputSubDir: "requestid", Description: "请求ID模块", Enabled: true},
		{Name: "restful", PackageName: "restful", DefaultFunc: func() interface{} { return restful.Default() }, OutputSubDir: "restful", Description: "RESTful API模块", Enabled: true},
		{Name: "rpcclient", PackageName: "rpcclient", DefaultFunc: func() interface{} { return rpcclient.Default() }, OutputSubDir: "rpcclient", Description: "RPC客户端", Enabled: true},
		{Name: "rpcserver", PackageName: "rpcserver", DefaultFunc: func() interface{} { return rpcserver.Default() }, OutputSubDir: "rpcserver", Description: "RPC服务端", Enabled: true},
		{Name: "security", PackageName: "security", DefaultFunc: func() interface{} { return security.Default() }, OutputSubDir: "security", Description: "安全模块", Enabled: true},
		{Name: "signature", PackageName: "signature", DefaultFunc: func() interface{} { return signature.Default() }, OutputSubDir: "signature", Description: "数字签名模块", Enabled: true},
		{Name: "sms", PackageName: "sms", DefaultFunc: func() interface{} { return sms.DefaultAliyunSms() }, OutputSubDir: "sms", Description: "短信发送模块", Enabled: true},
		{Name: "smtp", PackageName: "smtp", DefaultFunc: func() interface{} { return smtp.Default() }, OutputSubDir: "smtp", Description: "SMTP邮件模块", Enabled: true},
		{Name: "sts", PackageName: "sts", DefaultFunc: func() interface{} { return sts.DefaultAliyunSts() }, OutputSubDir: "sts", Description: "STS临时凭证", Enabled: true},
		{Name: "swagger", PackageName: "swagger", DefaultFunc: func() interface{} { return swagger.Default() }, OutputSubDir: "swagger", Description: "Swagger API文档", Enabled: true},
		{Name: "timeout", PackageName: "timeout", DefaultFunc: func() interface{} { return timeout.Default() }, OutputSubDir: "timeout", Description: "超时控制模块", Enabled: true},
		{Name: "tracing", PackageName: "tracing", DefaultFunc: func() interface{} { return tracing.Default() }, OutputSubDir: "tracing", Description: "链路追踪模块", Enabled: true},
		{Name: "wsc", PackageName: "wsc", DefaultFunc: func() interface{} { return wsc.Default() }, OutputSubDir: "wsc", Description: "WebSocket通信模块", Enabled: true},
		{Name: "youzan", PackageName: "youzan", DefaultFunc: func() interface{} { return youzan.Default() }, OutputSubDir: "youzan", Description: "有赞电商模块", Enabled: true},
		{Name: "zap", PackageName: "zap", DefaultFunc: func() interface{} { return zap.Default() }, OutputSubDir: "zap", Description: "Zap日志模块", Enabled: true},
		{Name: "jobs", PackageName: "jobs", DefaultFunc: func() interface{} { return jobs.Default() }, OutputSubDir: "jobs", Description: "任务调度模块", Enabled: true},
	}

	for _, module := range modules {
		sg.ModuleRegistry[module.Name] = module
	}

	sg.Logger.InfoKV("模块注册完成", "count", len(modules))
}

// GenerateAllConfigs 生成所有模块的配置文件
func (sg *SmartConfigGenerator) GenerateAllConfigs() error {
	sg.Logger.Info("开始生成所有模块配置文件")

	successCount := 0
	failCount := 0

	// 按模块名称排序
	modules := make([]ModuleConfig, 0, len(sg.ModuleRegistry))
	for _, module := range sg.ModuleRegistry {
		if module.Enabled {
			modules = append(modules, module)
		}
	}

	sort.Slice(modules, func(i, j int) bool {
		return modules[i].Name < modules[j].Name
	})

	for _, module := range modules {
		if err := sg.GenerateModuleConfig(module); err != nil {
			sg.Logger.ErrorKV("生成模块配置失败",
				"module", module.Name,
				"package", module.PackageName,
				"error", err.Error())
			failCount++
			continue
		}
		successCount++
	}

	sg.Logger.InfoKV("所有模块配置文件生成完成",
		"total", len(modules),
		"success", successCount,
		"failed", failCount)

	if failCount > 0 {
		return fmt.Errorf("部分模块配置生成失败: 成功 %d, 失败 %d", successCount, failCount)
	}

	return nil
}

// GenerateModuleConfig 生成单个模块的配置文件
func (sg *SmartConfigGenerator) GenerateModuleConfig(module ModuleConfig) error {
	// 构建输出目录: baseOutputDir/pkg/模块子目录
	outputDir := filepath.Join(sg.BaseOutputDir, "pkg", module.OutputSubDir)

	// 确保输出目录存在
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("创建输出目录失败: %w", err)
	}

	sg.Logger.DebugKV("生成模块配置",
		"module", module.Name,
		"package", module.PackageName,
		"output_dir", outputDir)

	// 获取默认配置
	defaultConfig := module.DefaultFunc()
	if defaultConfig == nil {
		return fmt.Errorf("模块 %s 的默认配置为空", module.Name)
	}

	// 生成YAML配置文件
	if sg.GenerateYAML {
		yamlPath := filepath.Join(outputDir, module.Name+".yaml")
		if err := sg.generateYAMLConfig(defaultConfig, yamlPath, module); err != nil {
			return fmt.Errorf("生成YAML配置失败: %w", err)
		}
	}

	// 生成JSON配置文件
	if sg.GenerateJSON {
		jsonPath := filepath.Join(outputDir, module.Name+".json")
		if err := sg.generateJSONConfig(defaultConfig, jsonPath, module); err != nil {
			return fmt.Errorf("生成JSON配置失败: %w", err)
		}
	}

	// 更新最后生成时间
	module.LastGenerated = time.Now()
	sg.ModuleRegistry[module.Name] = module

	sg.Logger.InfoKV("模块配置文件生成成功",
		"module", module.Name,
		"yaml", filepath.Join(outputDir, module.Name+".yaml"),
		"json", filepath.Join(outputDir, module.Name+".json"))

	return nil
}

// GenerateModulesByNames 根据模块名称生成指定模块的配置文件
func (sg *SmartConfigGenerator) GenerateModulesByNames(moduleNames ...string) error {
	if len(moduleNames) == 0 {
		return sg.GenerateAllConfigs()
	}

	sg.Logger.InfoKV("开始生成指定模块配置文件",
		"modules", strings.Join(moduleNames, ", "))

	successCount := 0
	failCount := 0

	for _, name := range moduleNames {
		module, exists := sg.ModuleRegistry[name]
		if !exists {
			sg.Logger.WarnKV("模块不存在", "module", name)
			failCount++
			continue
		}

		if err := sg.GenerateModuleConfig(module); err != nil {
			sg.Logger.ErrorKV("生成模块配置失败",
				"module", name,
				"error", err.Error())
			failCount++
			continue
		}
		successCount++
	}

	sg.Logger.InfoKV("指定模块配置文件生成完成",
		"total", len(moduleNames),
		"success", successCount,
		"failed", failCount)

	if failCount > 0 {
		return fmt.Errorf("部分模块配置生成失败: 成功 %d, 失败 %d", successCount, failCount)
	}

	return nil
}

// generateYAMLConfig 生成YAML配置文件
func (sg *SmartConfigGenerator) generateYAMLConfig(config interface{}, filePath string, module ModuleConfig) error {
	// 备份现有文件
	if sg.BackupExisting && sg.fileExists(filePath) {
		backupPath := filePath + ".backup." + time.Now().Format("20060102_150405")
		if err := sg.copyFile(filePath, backupPath); err != nil {
			sg.Logger.WarnKV("备份文件失败", "file", filePath, "backup", backupPath, "error", err.Error())
		}
	}

	// 如果文件已存在且不覆盖，则跳过
	if !sg.OverwriteExisting && sg.fileExists(filePath) {
		sg.Logger.DebugKV("跳过已存在的配置文件", "file", filePath)
		return nil
	}

	// 序列化为YAML（函数类型字段会被自动跳过，因为它们应该有 yaml:"-" 标签）
	yamlData, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("序列化YAML失败（可能包含未标记为 yaml:\"-\" 的函数字段）: %w", err)
	}

	// 添加文件头注释
	content := sg.generateFileHeader(module, "yaml") + string(yamlData)

	// 如果包含注释，添加字段注释
	if sg.IncludeComments {
		content = sg.addFieldComments(content, config)
	}

	// 写入文件
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("写入YAML文件失败: %w", err)
	}

	sg.Logger.DebugKV("YAML配置文件生成成功", "file", filePath)
	return nil
}

// generateJSONConfig 生成JSON配置文件
func (sg *SmartConfigGenerator) generateJSONConfig(config interface{}, filePath string, module ModuleConfig) error {
	// 备份现有文件
	if sg.BackupExisting && sg.fileExists(filePath) {
		backupPath := filePath + ".backup." + time.Now().Format("20060102_150405")
		if err := sg.copyFile(filePath, backupPath); err != nil {
			sg.Logger.WarnKV("备份文件失败", "file", filePath, "backup", backupPath, "error", err.Error())
		}
	}

	// 如果文件已存在且不覆盖，则跳过
	if !sg.OverwriteExisting && sg.fileExists(filePath) {
		sg.Logger.DebugKV("跳过已存在的配置文件", "file", filePath)
		return nil
	}

	// 序列化为JSON
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化JSON失败: %w", err)
	}

	// 添加文件头注释（JSON格式）
	content := sg.generateFileHeader(module, "json") + string(jsonData)

	// 写入文件
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("写入JSON文件失败: %w", err)
	}

	sg.Logger.DebugKV("JSON配置文件生成成功", "file", filePath)
	return nil
}

// generateFileHeader 生成文件头注释
func (sg *SmartConfigGenerator) generateFileHeader(module ModuleConfig, format string) string {
	commentPrefix := "#"
	if format == "json" {
		commentPrefix = "//"
	}

	header := fmt.Sprintf(`%s %s模块配置文件
%s 模块名称: %s
%s 包名: %s  
%s 描述: %s
%s 生成时间: %s
%s 注意: 此文件由配置生成器自动生成

`,
		commentPrefix, module.Name,
		commentPrefix, module.Name,
		commentPrefix, module.PackageName,
		commentPrefix, module.Description,
		commentPrefix, time.Now().Format("2006-01-02 15:04:05"),
		commentPrefix)

	return header
}

// addFieldComments 为YAML字段添加注释(改进版)
func (sg *SmartConfigGenerator) addFieldComments(content string, config interface{}) string {
	v := reflect.ValueOf(config)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return content
	}

	t := v.Type()
	lines := strings.Split(content, "\n")
	result := make([]string, 0, len(lines)*2) // 预留空间给注释行

	// 构建字段注释映射
	fieldComments := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		yamlTag := field.Tag.Get("yaml")

		if yamlTag == "" || yamlTag == "-" {
			continue
		}

		yamlName := strings.Split(yamlTag, ",")[0]
		if yamlName == "" {
			yamlName = sg.toKebabCase(field.Name)
		}

		comment := sg.extractDetailedComment(field, config)
		if comment != "" {
			fieldComments[yamlName] = comment
		}
	}

	// 处理每一行
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// 跳过空行和已有注释
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			result = append(result, line)
			continue
		}

		// 检查是否是配置行
		if colonIdx := strings.Index(trimmed, ":"); colonIdx > 0 {
			fieldName := strings.TrimSpace(trimmed[:colonIdx])

			// 查找对应的注释
			if comment, exists := fieldComments[fieldName]; exists {
				// 计算缩进
				indent := ""
				if len(line) > len(trimmed) {
					indent = line[:len(line)-len(trimmed)]
				}

				// 添加注释行
				result = append(result, indent+"# "+comment)
			}
		}

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

// toKebabCase 将驼峰转换为短横线命名
func (sg *SmartConfigGenerator) toKebabCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '-')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

// extractFieldComment 从结构体字段提取注释(改进版)
func (sg *SmartConfigGenerator) extractFieldComment(field reflect.StructField) string {
	// 尝试从tag中提取注释信息
	if comment := field.Tag.Get("comment"); comment != "" {
		return comment
	}

	if description := field.Tag.Get("description"); description != "" {
		return description
	}

	if doc := field.Tag.Get("doc"); doc != "" {
		return doc
	}

	// 根据字段名生成更详细的注释
	return sg.generateDetailedComment(field.Name, field.Type)
}

// extractDetailedComment 提取详细注释(优先从源代码,包含类型信息)
func (sg *SmartConfigGenerator) extractDetailedComment(field reflect.StructField, config interface{}) string {
	parts := []string{}

	// 1. 优先从源代码注释提取
	sourceComment := sg.getFieldCommentFromSource(config, field.Name)
	if sourceComment != "" {
		parts = append(parts, sourceComment)
		return sourceComment // 如果有源代码注释,直接返回,最准确
	}

	// 2. 从tag提取注释
	if comment := field.Tag.Get("comment"); comment != "" {
		parts = append(parts, comment)
	} else if desc := field.Tag.Get("description"); desc != "" {
		parts = append(parts, desc)
	} else if doc := field.Tag.Get("doc"); doc != "" {
		parts = append(parts, doc)
	}

	// 添加类型信息
	typeInfo := sg.getTypeInfo(field.Type)
	if typeInfo != "" {
		parts = append(parts, typeInfo)
	}

	// 如果没有注释,生成描述性注释
	if len(parts) == 0 {
		parts = append(parts, sg.generateDetailedComment(field.Name, field.Type))
	}

	return strings.Join(parts, " | ")
}

// generateDetailedComment 生成详细注释(包含类型信息)
func (sg *SmartConfigGenerator) generateDetailedComment(fieldName string, fieldType reflect.Type) string {
	// 将驼峰命名转换为可读文本
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	readable := re.ReplaceAllString(fieldName, "${1} ${2}")
	readable = strings.ToLower(readable)

	// 特殊字段名称映射
	fieldDescriptions := map[string]string{
		"modulename":    "模块标识名称",
		"module-name":   "模块标识名称",
		"name":          "服务名称",
		"enabled":       "是否启用该功能",
		"debug":         "是否启用调试模式",
		"version":       "版本号",
		"environment":   "运行环境 (dev/test/prod)",
		"host":          "主机地址",
		"port":          "端口号",
		"timeout":       "超时时间",
		"password":      "密码",
		"username":      "用户名",
		"endpoint":      "服务端点地址",
		"database":      "数据库名称",
		"maxretries":    "最大重试次数",
		"max-retries":   "最大重试次数",
		"idletimeout":   "空闲超时时间",
		"idle-timeout":  "空闲超时时间",
		"readtimeout":   "读取超时时间",
		"read-timeout":  "读取超时时间",
		"writetimeout":  "写入超时时间",
		"write-timeout": "写入超时时间",
		"poolsize":      "连接池大小",
		"pool-size":     "连接池大小",
	}

	// 检查是否有预定义的描述
	key := strings.ToLower(strings.ReplaceAll(fieldName, "-", ""))
	if desc, exists := fieldDescriptions[key]; exists {
		readable = desc
	}

	// 根据类型添加额外信息
	typeHint := ""
	switch fieldType.Kind() {
	case reflect.Bool:
		typeHint = "(true/false)"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if fieldType.String() == "time.Duration" {
			typeHint = "(时间间隔: 如 30s, 5m, 1h)"
		} else {
			typeHint = "(整数)"
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		typeHint = "(正整数)"
	case reflect.Float32, reflect.Float64:
		typeHint = "(浮点数)"
	case reflect.String:
		typeHint = "(字符串)"
	case reflect.Slice:
		elem := fieldType.Elem()
		elemType := elem.Kind().String()
		if elem.Kind() == reflect.String {
			typeHint = "(字符串数组)"
		} else {
			typeHint = fmt.Sprintf("(%s数组)", elemType)
		}
	case reflect.Map:
		typeHint = "(键值对映射)"
	case reflect.Struct:
		if fieldType.String() == "time.Duration" {
			typeHint = "(时间间隔: 如 30s, 5m, 1h)"
		} else if fieldType.String() == "time.Time" {
			typeHint = "(时间: 如 2006-01-02 15:04:05)"
		} else {
			typeHint = "(嵌套配置对象)"
		}
	case reflect.Ptr:
		elem := fieldType.Elem()
		if elem.Kind() == reflect.Struct {
			typeHint = "(嵌套配置对象)"
		}
	}

	if typeHint != "" {
		return readable + " " + typeHint
	}
	return readable
} // getTypeInfo 获取类型信息提示
func (sg *SmartConfigGenerator) getTypeInfo(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Slice:
		elem := t.Elem()
		return fmt.Sprintf("数组[%s]", elem.Kind().String())
	case reflect.Map:
		key := t.Key()
		val := t.Elem()
		return fmt.Sprintf("映射<%s, %s>", key.Kind().String(), val.Kind().String())
	case reflect.Struct:
		if t.String() == "time.Duration" {
			return "时间间隔(如: 30s, 5m, 1h)"
		}
		if t.String() == "time.Time" {
			return "时间(如: 2006-01-02 15:04:05)"
		}
		return "嵌套配置"
	}
	return ""
}

// generateGenericComment 根据字段名生成通用注释
func (sg *SmartConfigGenerator) generateGenericComment(fieldName string) string {
	// 将驼峰命名转换为有意义的注释
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	result := re.ReplaceAllString(fieldName, "${1} ${2}")
	return strings.ToLower(result) + "配置"
}

// fileExists 检查文件是否存在
func (sg *SmartConfigGenerator) fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// copyFile 复制文件
func (sg *SmartConfigGenerator) copyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, input, 0644)
}

// GetModuleList 获取所有注册的模块列表
func (sg *SmartConfigGenerator) GetModuleList() []string {
	modules := make([]string, 0, len(sg.ModuleRegistry))
	for name := range sg.ModuleRegistry {
		modules = append(modules, name)
	}
	sort.Strings(modules)
	return modules
}

// GetEnabledModules 获取已启用的模块列表
func (sg *SmartConfigGenerator) GetEnabledModules() []string {
	modules := make([]string, 0)
	for name, module := range sg.ModuleRegistry {
		if module.Enabled {
			modules = append(modules, name)
		}
	}
	sort.Strings(modules)
	return modules
}

// EnableModule 启用指定模块
func (sg *SmartConfigGenerator) EnableModule(moduleName string) error {
	module, exists := sg.ModuleRegistry[moduleName]
	if !exists {
		return fmt.Errorf("模块 %s 不存在", moduleName)
	}

	module.Enabled = true
	sg.ModuleRegistry[moduleName] = module
	sg.Logger.DebugKV("模块已启用", "module", moduleName)
	return nil
}

// DisableModule 禁用指定模块
func (sg *SmartConfigGenerator) DisableModule(moduleName string) error {
	module, exists := sg.ModuleRegistry[moduleName]
	if !exists {
		return fmt.Errorf("模块 %s 不存在", moduleName)
	}

	module.Enabled = false
	sg.ModuleRegistry[moduleName] = module
	sg.Logger.DebugKV("模块已禁用", "module", moduleName)
	return nil
}

// EnableOnlyModules 只启用指定的模块，其他模块全部禁用
func (sg *SmartConfigGenerator) EnableOnlyModules(moduleNames ...string) error {
	// 先禁用所有模块
	for name, module := range sg.ModuleRegistry {
		module.Enabled = false
		sg.ModuleRegistry[name] = module
	}

	// 再启用指定模块
	for _, name := range moduleNames {
		if err := sg.EnableModule(name); err != nil {
			return err
		}
	}

	sg.Logger.InfoKV("只启用指定模块", "modules", strings.Join(moduleNames, ", "))
	return nil
}

// PrintModuleStatus 打印模块状态
func (sg *SmartConfigGenerator) PrintModuleStatus() {
	sg.Logger.Info("=== 模块状态列表 ===")

	modules := make([]ModuleConfig, 0, len(sg.ModuleRegistry))
	for _, module := range sg.ModuleRegistry {
		modules = append(modules, module)
	}

	sort.Slice(modules, func(i, j int) bool {
		return modules[i].Name < modules[j].Name
	})

	enabledCount := 0
	for _, module := range modules {
		status := "❌ 禁用"
		if module.Enabled {
			status = "✅ 启用"
			enabledCount++
		}

		sg.Logger.InfoKV("模块状态",
			"name", module.Name,
			"package", module.PackageName,
			"status", status,
			"description", module.Description)
	}

	sg.Logger.InfoKV("模块统计",
		"total", len(modules),
		"enabled", enabledCount,
		"disabled", len(modules)-enabledCount)
}

// ValidateModuleConfig 验证模块配置
func (sg *SmartConfigGenerator) ValidateModuleConfig(moduleName string) error {
	module, exists := sg.ModuleRegistry[moduleName]
	if !exists {
		return fmt.Errorf("模块 %s 不存在", moduleName)
	}

	// 检查默认函数是否可用
	if module.DefaultFunc == nil {
		return fmt.Errorf("模块 %s 的默认函数为空", moduleName)
	}

	// 尝试调用默认函数
	config := module.DefaultFunc()
	if config == nil {
		return fmt.Errorf("模块 %s 的默认配置为空", moduleName)
	}

	// 尝试序列化
	if _, err := yaml.Marshal(config); err != nil {
		return fmt.Errorf("模块 %s 的配置序列化失败: %w", moduleName, err)
	}

	sg.Logger.DebugKV("模块配置验证通过", "module", moduleName)
	return nil
}

// ValidateAllModules 验证所有模块配置
func (sg *SmartConfigGenerator) ValidateAllModules() error {
	sg.Logger.Info("开始验证所有模块配置")

	failedModules := make([]string, 0)

	for name := range sg.ModuleRegistry {
		if err := sg.ValidateModuleConfig(name); err != nil {
			sg.Logger.ErrorKV("模块配置验证失败", "module", name, "error", err.Error())
			failedModules = append(failedModules, name)
		}
	}

	if len(failedModules) > 0 {
		return fmt.Errorf("以下模块配置验证失败: %s", strings.Join(failedModules, ", "))
	}

	sg.Logger.InfoKV("所有模块配置验证通过", "count", len(sg.ModuleRegistry))
	return nil
}

// CleanupBackupFiles 清理备份文件
func (sg *SmartConfigGenerator) CleanupBackupFiles(maxAge time.Duration) error {
	sg.Logger.InfoKV("开始清理备份文件", "max_age", maxAge.String())

	cleanupCount := 0

	for _, module := range sg.ModuleRegistry {
		outputDir := filepath.Join(sg.BaseOutputDir, "pkg", module.OutputSubDir)

		if err := filepath.Walk(outputDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil // 忽略错误，继续处理其他文件
			}

			// 只处理备份文件
			if !strings.Contains(info.Name(), ".backup.") {
				return nil
			}

			// 检查文件年龄
			if time.Since(info.ModTime()) > maxAge {
				if err := os.Remove(path); err != nil {
					sg.Logger.WarnKV("删除备份文件失败", "file", path, "error", err.Error())
				} else {
					cleanupCount++
					sg.Logger.DebugKV("已删除过期备份文件", "file", path)
				}
			}

			return nil
		}); err != nil {
			sg.Logger.WarnKV("遍历目录失败", "dir", outputDir, "error", err.Error())
		}
	}

	sg.Logger.InfoKV("备份文件清理完成", "cleaned", cleanupCount)
	return nil
}

// GetModuleInfo 获取模块详细信息
func (sg *SmartConfigGenerator) GetModuleInfo(moduleName string) (*ModuleConfig, error) {
	module, exists := sg.ModuleRegistry[moduleName]
	if !exists {
		return nil, fmt.Errorf("模块 %s 不存在", moduleName)
	}

	return &module, nil
}

// UpdateModuleConfig 更新模块配置
func (sg *SmartConfigGenerator) UpdateModuleConfig(moduleName string, updates map[string]interface{}) error {
	module, exists := sg.ModuleRegistry[moduleName]
	if !exists {
		return fmt.Errorf("模块 %s 不存在", moduleName)
	}

	// 根据updates更新模块配置
	if description, ok := updates["description"].(string); ok {
		module.Description = description
	}

	if enabled, ok := updates["enabled"].(bool); ok {
		module.Enabled = enabled
	}

	if outputSubDir, ok := updates["outputSubDir"].(string); ok {
		module.OutputSubDir = outputSubDir
	}

	sg.ModuleRegistry[moduleName] = module
	sg.Logger.DebugKV("模块配置已更新", "module", moduleName)
	return nil
}

// parseSourceComments 从Go源代码中解析字段注释
func (sg *SmartConfigGenerator) parseSourceComments(packagePath string, structName string) map[string]string {
	// 检查缓存
	cacheKey := packagePath + "." + structName
	if comments, exists := sg.commentCache[cacheKey]; exists {
		return comments
	}

	comments := make(map[string]string)

	// 查找包的源代码目录
	pkgDir := filepath.Join(sg.BaseOutputDir, "pkg", packagePath)

	// 解析目录中的所有Go文件
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, pkgDir, nil, parser.ParseComments)
	if err != nil {
		sg.Logger.WarnKV("解析源代码失败", "package", packagePath, "error", err.Error())
		return comments
	}

	// 遍历所有包和文件
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				// 查找结构体定义
				typeSpec, ok := n.(*ast.TypeSpec)
				if !ok || typeSpec.Name.Name != structName {
					return true
				}

				// 确保是结构体
				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					return true
				}

				// 遍历结构体字段
				for _, field := range structType.Fields.List {
					if field.Comment == nil || len(field.Names) == 0 {
						continue
					}

					// 提取注释文本
					comment := strings.TrimSpace(field.Comment.Text())
					comment = strings.TrimPrefix(comment, "//")
					comment = strings.TrimSpace(comment)

					// 保存到map
					for _, name := range field.Names {
						comments[name.Name] = comment
					}
				}

				return false
			})
		}
	}

	// 缓存结果
	sg.commentCache[cacheKey] = comments
	sg.Logger.DebugKV("解析源代码注释", "package", packagePath, "struct", structName, "count", len(comments))

	return comments
}

// getFieldCommentFromSource 从源代码获取字段注释
func (sg *SmartConfigGenerator) getFieldCommentFromSource(config interface{}, fieldName string) string {
	t := reflect.TypeOf(config)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return ""
	}

	// 获取包路径和结构体名称
	pkgPath := t.PkgPath()
	structName := t.Name()

	// 从包路径提取相对路径
	// 例如: github.com/kamalyes/go-config/pkg/gateway -> gateway
	parts := strings.Split(pkgPath, "/")
	if len(parts) > 0 {
		pkgPath = parts[len(parts)-1]
	}

	// 解析源代码注释
	comments := sg.parseSourceComments(pkgPath, structName)

	// 查找字段注释
	if comment, exists := comments[fieldName]; exists {
		return comment
	}

	return ""
}
