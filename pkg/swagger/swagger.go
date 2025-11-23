/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 17:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 23:16:53
 * @FilePath: \go-config\pkg\swagger\swagger.go
 * @Description: Swagger配置模块 - 基于go-config设计模式
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package swagger

import (
	"fmt"
	"github.com/kamalyes/go-config/internal"
)

// AuthType Swagger认证类型
type AuthType string

const (
	// AuthNone 无认证
	AuthNone AuthType = "none"
	// AuthBasic 基本认证
	AuthBasic AuthType = "basic"
	// AuthBearer Bearer Token认证
	AuthBearer AuthType = "bearer"
	// AuthCustom 自定义认证
	AuthCustom AuthType = "custom"
)

// Contact Swagger联系信息
type Contact struct {
	Name  string `mapstructure:"name" yaml:"name" json:"name"`    // 联系人姓名
	URL   string `mapstructure:"url" yaml:"url" json:"url"`       // 联系URL
	Email string `mapstructure:"email" yaml:"email" json:"email"` // 联系邮箱
}

// License Swagger许可证信息
type License struct {
	Name string `mapstructure:"name" yaml:"name" json:"name"` // 许可证名称
	URL  string `mapstructure:"url" yaml:"url" json:"url"`    // 许可证URL
}

// AuthConfig Swagger认证配置
type AuthConfig struct {
	Type        AuthType `mapstructure:"type" yaml:"type" json:"type"`                        // 认证类型
	Username    string   `mapstructure:"username" yaml:"username" json:"username"`            // 用户名（基本认证）
	Password    string   `mapstructure:"password" yaml:"password" json:"password"`            // 密码（基本认证）
	Token       string   `mapstructure:"token" yaml:"token" json:"token"`                     // Token（Bearer认证）
	HeaderName  string   `mapstructure:"header-name" yaml:"header-name" json:"headerName"`    // 自定义header名称
	HeaderValue string   `mapstructure:"header-value" yaml:"header-value" json:"headerValue"` // 自定义header值
}

// ServiceSpec 单个微服务Swagger规范配置
type ServiceSpec struct {
	Name        string   `mapstructure:"name" yaml:"name" json:"name"`                      // 服务名称
	Description string   `mapstructure:"description" yaml:"description" json:"description"` // 服务描述
	SpecPath    string   `mapstructure:"spec-path" yaml:"spec-path" json:"specPath"`        // Swagger规范文件路径
	URL         string   `mapstructure:"url" yaml:"url" json:"url"`                         // 远程Swagger文档URL
	Version     string   `mapstructure:"version" yaml:"version" json:"version"`             // 服务版本
	Enabled     bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`             // 是否启用
	BasePath    string   `mapstructure:"base-path" yaml:"base-path" json:"basePath"`        // API基础路径前缀
	Tags        []string `mapstructure:"tags" yaml:"tags" json:"tags"`                      // 服务标签
}

// AggregateConfig 聚合Swagger配置
type AggregateConfig struct {
	Enabled  bool           `mapstructure:"enabled" yaml:"enabled" json:"enabled"`      // 是否启用聚合
	Mode     string         `mapstructure:"mode" yaml:"mode" json:"mode"`               // 聚合模式: merge|selector
	Services []*ServiceSpec `mapstructure:"services" yaml:"services" json:"services"`   // 微服务列表
	UILayout string         `mapstructure:"ui-layout" yaml:"ui-layout" json:"uiLayout"` // UI布局: tabs|dropdown|list
}

// Swagger Swagger配置结构
type Swagger struct {
	ModuleName  string           `mapstructure:"module-name" yaml:"module-name" json:"moduleName"`  // 模块名称
	Enabled     bool             `mapstructure:"enabled" yaml:"enabled" json:"enabled"`             // 是否启用Swagger
	JSONPath    string           `mapstructure:"json-path" yaml:"json-path" json:"jsonPath"`        // Swagger JSON文件路径
	UIPath      string           `mapstructure:"ui-path" yaml:"ui-path" json:"uiPath"`              // Swagger UI路由路径
	YamlPath    string           `mapstructure:"yaml-path" yaml:"yaml-path" json:"yamlPath"`        // Swagger YAML文件路径
	SpecPath    string           `mapstructure:"spec-path" yaml:"spec-path" json:"specPath"`        // Swagger规范文件路径(自动检测格式)
	Title       string           `mapstructure:"title" yaml:"title" json:"title"`                   // 文档标题
	Description string           `mapstructure:"description" yaml:"description" json:"description"` // 文档描述
	Version     string           `mapstructure:"version" yaml:"version" json:"version"`             // Swagger版本
	Contact     *Contact         `mapstructure:"contact" yaml:"contact" json:"contact"`             // 联系信息
	License     *License         `mapstructure:"license" yaml:"license" json:"license"`             // 许可证信息
	Auth        *AuthConfig      `mapstructure:"auth" yaml:"auth" json:"auth"`                      // 认证配置
	Aggregate   *AggregateConfig `mapstructure:"aggregate" yaml:"aggregate" json:"aggregate"`       // 聚合配置
}

// Default 创建默认Swagger配置
func Default() *Swagger {
	return &Swagger{
		ModuleName:  "swagger",
		Enabled:     true, // 默认启用Swagger
		JSONPath:    "/swagger/doc.json",
		UIPath:      "/swagger",
		YamlPath:    "/swagger/doc.yaml",
		SpecPath:    "./docs/swagger.yaml", // 默认规范文件路径
		Title:       "API Documentation",
		Description: "API Documentation powered by Swagger UI",
		Version:     "1.0.0",
		Contact: &Contact{
			Name:  "API Support",
			URL:   "https://example.com/support",
			Email: "support@example.com",
		},
		License: &License{
			Name: "MIT",
			URL:  "https://opensource.org/licenses/MIT",
		},
		Auth: &AuthConfig{
			Type: AuthNone,
		},
		Aggregate: &AggregateConfig{
			Enabled:  false,
			Mode:     "merge",
			Services: []*ServiceSpec{},
			UILayout: "tabs",
		},
	}
}

// WithModuleName 设置模块名称
func (c *Swagger) WithModuleName(moduleName string) *Swagger {
	c.ModuleName = moduleName
	return c
}

// WithEnabled 设置是否启用Swagger
func (c *Swagger) WithEnabled(enabled bool) *Swagger {
	c.Enabled = enabled
	return c
}

// WithJSONPath 设置Swagger JSON文件路径
func (c *Swagger) WithJSONPath(path string) *Swagger {
	c.JSONPath = path
	return c
}

// WithUIPath 设置Swagger UI路由路径
func (c *Swagger) WithUIPath(path string) *Swagger {
	c.UIPath = path
	return c
}

// WithYamlPath 设置Swagger YAML文件路径
func (c *Swagger) WithYamlPath(path string) *Swagger {
	c.YamlPath = path
	return c
}

// WithSpecPath 设置Swagger规范文件路径
func (c *Swagger) WithSpecPath(path string) *Swagger {
	c.SpecPath = path
	return c
}

// WithTitle 设置文档标题
func (c *Swagger) WithTitle(title string) *Swagger {
	c.Title = title
	return c
}

// WithDescription 设置文档描述
func (c *Swagger) WithDescription(description string) *Swagger {
	c.Description = description
	return c
}

// WithAuth 设置认证配置
func (c *Swagger) WithAuth(auth *AuthConfig) *Swagger {
	c.Auth = auth
	return c
}

// WithVersion 设置Swagger版本
func (c *Swagger) WithVersion(version string) *Swagger {
	c.Version = version
	return c
}

// WithContact 设置联系信息
func (c *Swagger) WithContact(contact *Contact) *Swagger {
	c.Contact = contact
	return c
}

// WithLicense 设置许可证信息
func (c *Swagger) WithLicense(license *License) *Swagger {
	c.License = license
	return c
}

// WithBasicAuth 设置基本认证
func (c *Swagger) WithBasicAuth(username, password string) *Swagger {
	if c.Auth == nil {
		c.Auth = &AuthConfig{}
	}
	c.Auth.Type = AuthBasic
	c.Auth.Username = username
	c.Auth.Password = password
	return c
}

// WithBearerAuth 设置Bearer Token认证
func (c *Swagger) WithBearerAuth(token string) *Swagger {
	if c.Auth == nil {
		c.Auth = &AuthConfig{}
	}
	c.Auth.Type = AuthBearer
	c.Auth.Token = token
	return c
}

// WithAggregate 设置聚合配置
func (c *Swagger) WithAggregate(aggregate *AggregateConfig) *Swagger {
	c.Aggregate = aggregate
	return c
}

// WithAggregateServices 设置聚合服务列表
func (c *Swagger) WithAggregateServices(services []*ServiceSpec) *Swagger {
	if c.Aggregate == nil {
		c.Aggregate = &AggregateConfig{
			Enabled:  true,
			Mode:     "merge",
			UILayout: "tabs",
		}
	}
	c.Aggregate.Services = services
	c.Aggregate.Enabled = len(services) > 0
	return c
}

// AddAggregateService 添加单个聚合服务
func (c *Swagger) AddAggregateService(service *ServiceSpec) *Swagger {
	if c.Aggregate == nil {
		c.Aggregate = &AggregateConfig{
			Enabled:  true,
			Mode:     "merge",
			UILayout: "tabs",
			Services: []*ServiceSpec{},
		}
	}
	c.Aggregate.Services = append(c.Aggregate.Services, service)
	c.Aggregate.Enabled = true
	return c
}

// WithAggregateMode 设置聚合模式
func (c *Swagger) WithAggregateMode(mode string) *Swagger {
	if c.Aggregate == nil {
		c.Aggregate = &AggregateConfig{
			Enabled:  false,
			UILayout: "tabs",
			Services: []*ServiceSpec{},
		}
	}
	c.Aggregate.Mode = mode
	return c
}

// WithAggregateUILayout 设置UI布局
func (c *Swagger) WithAggregateUILayout(layout string) *Swagger {
	if c.Aggregate == nil {
		c.Aggregate = &AggregateConfig{
			Enabled:  false,
			Mode:     "merge",
			Services: []*ServiceSpec{},
		}
	}
	c.Aggregate.UILayout = layout
	return c
}

// EnableAggregate 启用聚合功能
func (c *Swagger) EnableAggregate() *Swagger {
	if c.Aggregate == nil {
		c.Aggregate = &AggregateConfig{
			Mode:     "merge",
			UILayout: "tabs",
			Services: []*ServiceSpec{},
		}
	}
	c.Aggregate.Enabled = true
	return c
}

// DisableAggregate 禁用聚合功能
func (c *Swagger) DisableAggregate() *Swagger {
	if c.Aggregate != nil {
		c.Aggregate.Enabled = false
	}
	return c
}

// IsAggregateEnabled 检查聚合功能是否启用
func (c *Swagger) IsAggregateEnabled() bool {
	return c.Aggregate != nil && c.Aggregate.Enabled
}

// WithDefaults 使用默认配置填充未设置的字段
func (c *Swagger) WithDefaults() *Swagger {
	defaults := Default()

	if c.JSONPath == "" {
		c.JSONPath = defaults.JSONPath
	}
	if c.UIPath == "" {
		c.UIPath = defaults.UIPath
	}
	if c.Title == "" {
		c.Title = defaults.Title
	}
	if c.Description == "" {
		c.Description = defaults.Description
	}
	if c.Auth == nil {
		c.Auth = defaults.Auth
	}
	if c.Version == "" {
		c.Version = defaults.Version
	}

	return c
}

// WithBasePath 设置基础路径（同时设置JSON和UI路径）
func (c *Swagger) WithBasePath(basePath string) *Swagger {
	c.JSONPath = basePath + "/doc.json"
	c.UIPath = basePath + "/*any"
	return c
}

// WithCustomPaths 自定义JSON和UI路径
func (c *Swagger) WithCustomPaths(jsonPath, uiPath string) *Swagger {
	c.JSONPath = jsonPath
	c.UIPath = uiPath
	return c
}

// WithDocumentInfo 设置文档信息（标题和描述）
func (c *Swagger) WithDocumentInfo(title, description string) *Swagger {
	c.Title = title
	c.Description = description
	return c
}

// Clone 克隆配置对象
func (c *Swagger) Clone() internal.Configurable {
	clone := &Swagger{
		ModuleName:  c.ModuleName,
		Enabled:     c.Enabled,
		JSONPath:    c.JSONPath,
		UIPath:      c.UIPath,
		Title:       c.Title,
		Description: c.Description,
		Version:     c.Version,
	}

	if c.Auth != nil {
		clone.Auth = &AuthConfig{
			Type:        c.Auth.Type,
			Username:    c.Auth.Username,
			Password:    c.Auth.Password,
			Token:       c.Auth.Token,
			HeaderName:  c.Auth.HeaderName,
			HeaderValue: c.Auth.HeaderValue,
		}
	}

	if c.Contact != nil {
		clone.Contact = &Contact{
			Name:  c.Contact.Name,
			URL:   c.Contact.URL,
			Email: c.Contact.Email,
		}
	}

	if c.License != nil {
		clone.License = &License{
			Name: c.License.Name,
			URL:  c.License.URL,
		}
	}

	if c.Aggregate != nil {
		clone.Aggregate = &AggregateConfig{
			Enabled:  c.Aggregate.Enabled,
			Mode:     c.Aggregate.Mode,
			UILayout: c.Aggregate.UILayout,
		}
		// 克隆服务列表
		if len(c.Aggregate.Services) > 0 {
			clone.Aggregate.Services = make([]*ServiceSpec, len(c.Aggregate.Services))
			for i, service := range c.Aggregate.Services {
				clone.Aggregate.Services[i] = &ServiceSpec{
					Name:        service.Name,
					Description: service.Description,
					SpecPath:    service.SpecPath,
					URL:         service.URL,
					Version:     service.Version,
					Enabled:     service.Enabled,
					BasePath:    service.BasePath,
				}
				if len(service.Tags) > 0 {
					clone.Aggregate.Services[i].Tags = make([]string, len(service.Tags))
					copy(clone.Aggregate.Services[i].Tags, service.Tags)
				}
			}
		}
	}

	return clone
}

// Get 返回 Swagger 配置的所有字段
func (c *Swagger) Get() interface{} {
	return c
}

// Set 更新 Swagger 配置的字段
func (c *Swagger) Set(data interface{}) {
	if configData, ok := data.(*Swagger); ok {
		c.ModuleName = configData.ModuleName
		c.Enabled = configData.Enabled
		c.JSONPath = configData.JSONPath
		c.UIPath = configData.UIPath
		c.Title = configData.Title
		c.Description = configData.Description
		c.Version = configData.Version
		c.Contact = configData.Contact
		c.License = configData.License
		c.Auth = configData.Auth
		c.Aggregate = configData.Aggregate
	}
}

// Validate 检查 Swagger 配置的有效性
func (c *Swagger) Validate() error {
	return internal.ValidateStruct(c)
}

// Reset 重置为默认配置
func (c *Swagger) Reset() *Swagger {
	defaults := Default()
	c.ModuleName = defaults.ModuleName
	c.Enabled = defaults.Enabled
	c.JSONPath = defaults.JSONPath
	c.UIPath = defaults.UIPath
	c.Title = defaults.Title
	c.Description = defaults.Description
	c.Version = defaults.Version
	c.Contact = defaults.Contact
	c.License = defaults.License
	c.Auth = defaults.Auth
	c.Aggregate = defaults.Aggregate
	return c
}

// Disable 禁用Swagger
func (c *Swagger) Disable() *Swagger {
	c.Enabled = false
	return c
}

// Enable 启用Swagger
func (c *Swagger) Enable() *Swagger {
	c.Enabled = true
	return c
}

// IsEnabled 检查Swagger是否启用
func (c *Swagger) IsEnabled() bool {
	return c.Enabled
}

// IsAuthEnabled 检查是否启用了认证
func (c *Swagger) IsAuthEnabled() bool {
	return c.Auth != nil && c.Auth.Type != AuthNone
}

// GetAuthType 获取认证类型
func (c *Swagger) GetAuthType() AuthType {
	if c.Auth == nil {
		return AuthNone
	}
	return c.Auth.Type
}

// GetFullUIPath 获取完整的UI路径
func (c *Swagger) GetFullUIPath() string {
	if c.UIPath == "" {
		return "/swagger/*any"
	}
	return c.UIPath
}

// GetFullJSONPath 获取完整的JSON路径
func (c *Swagger) GetFullJSONPath() string {
	if c.JSONPath == "" {
		return "/swagger/doc.json"
	}
	return c.JSONPath
}

// ValidateAuth 验证认证配置
func (c *Swagger) ValidateAuth() error {
	if c.Auth == nil || c.Auth.Type == AuthNone {
		return nil
	}

	switch c.Auth.Type {
	case AuthBasic:
		if c.Auth.Username == "" || c.Auth.Password == "" {
			return fmt.Errorf("basic auth requires username and password")
		}
	case AuthBearer:
		if c.Auth.Token == "" {
			return fmt.Errorf("bearer auth requires token")
		}
	case AuthCustom:
		if c.Auth.HeaderName == "" || c.Auth.HeaderValue == "" {
			return fmt.Errorf("custom auth requires header name and value")
		}
	default:
		return fmt.Errorf("unsupported auth type: %s", c.Auth.Type)
	}

	return nil
}

// NewServiceSpec 创建新的服务规范配置
func NewServiceSpec(name, description, specPath string) *ServiceSpec {
	return &ServiceSpec{
		Name:        name,
		Description: description,
		SpecPath:    specPath,
		Enabled:     true,
		Version:     "1.0.0",
		Tags:        []string{},
	}
}

// NewRemoteServiceSpec 创建远程服务规范配置
func NewRemoteServiceSpec(name, description, url string) *ServiceSpec {
	return &ServiceSpec{
		Name:        name,
		Description: description,
		URL:         url,
		Enabled:     true,
		Version:     "1.0.0",
		Tags:        []string{},
	}
}

// WithName 设置服务名称
func (s *ServiceSpec) WithName(name string) *ServiceSpec {
	s.Name = name
	return s
}

// WithDescription 设置服务描述
func (s *ServiceSpec) WithDescription(description string) *ServiceSpec {
	s.Description = description
	return s
}

// WithVersion 设置服务版本
func (s *ServiceSpec) WithVersion(version string) *ServiceSpec {
	s.Version = version
	return s
}

// WithBasePath 设置API基础路径
func (s *ServiceSpec) WithBasePath(basePath string) *ServiceSpec {
	s.BasePath = basePath
	return s
}

// WithTags 设置服务标签
func (s *ServiceSpec) WithTags(tags []string) *ServiceSpec {
	s.Tags = tags
	return s
}

// AddTag 添加单个标签
func (s *ServiceSpec) AddTag(tag string) *ServiceSpec {
	if s.Tags == nil {
		s.Tags = []string{}
	}
	s.Tags = append(s.Tags, tag)
	return s
}

// Enable 启用服务
func (s *ServiceSpec) Enable() *ServiceSpec {
	s.Enabled = true
	return s
}

// Disable 禁用服务
func (s *ServiceSpec) Disable() *ServiceSpec {
	s.Enabled = false
	return s
}
