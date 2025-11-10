/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 17:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-11 10:25:11
 * @FilePath: \go-rpc-gateway\go-config\pkg\swagger\swagger.go
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

// AuthConfig Swagger认证配置
type AuthConfig struct {
	Type        AuthType `mapstructure:"type" yaml:"type" json:"type"`                         // 认证类型
	Username    string   `mapstructure:"username" yaml:"username" json:"username"`             // 用户名（基本认证）
	Password    string   `mapstructure:"password" yaml:"password" json:"password"`             // 密码（基本认证）
	Token       string   `mapstructure:"token" yaml:"token" json:"token"`                      // Token（Bearer认证）
	HeaderName  string   `mapstructure:"header_name" yaml:"header_name" json:"header_name"`    // 自定义header名称
	HeaderValue string   `mapstructure:"header_value" yaml:"header_value" json:"header_value"` // 自定义header值
}

// Swagger Swagger配置结构
type Swagger struct {
	ModuleName  string      `mapstructure:"module_name" yaml:"module_name" json:"module_name"` // 模块名称
	Enabled     bool        `mapstructure:"enabled" yaml:"enabled" json:"enabled"`             // 是否启用Swagger
	JSONPath    string      `mapstructure:"json_path" yaml:"json_path" json:"json_path"`       // Swagger JSON文件路径
	UIPath      string      `mapstructure:"ui_path" yaml:"ui_path" json:"ui_path"`             // Swagger UI路由路径
	Title       string      `mapstructure:"title" yaml:"title" json:"title"`                   // 文档标题
	Description string      `mapstructure:"description" yaml:"description" json:"description"` // 文档描述
	Auth        *AuthConfig `mapstructure:"auth" yaml:"auth" json:"auth"`                      // 认证配置
}

// Default 创建默认Swagger配置
func Default() *Swagger {
	return &Swagger{
		ModuleName:  "swagger",
		Enabled:     false,
		JSONPath:    "/swagger/doc.json",
		UIPath:      "/swagger/*any",
		Title:       "API Documentation",
		Description: "API Documentation powered by Swagger UI",
		Auth: &AuthConfig{
			Type: AuthNone,
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

// WithCustomAuth 设置自定义认证
func (c *Swagger) WithCustomAuth(headerName, headerValue string) *Swagger {
	if c.Auth == nil {
		c.Auth = &AuthConfig{}
	}
	c.Auth.Type = AuthCustom
	c.Auth.HeaderName = headerName
	c.Auth.HeaderValue = headerValue
	return c
}

// WithoutAuth 禁用认证
func (c *Swagger) WithoutAuth() *Swagger {
	if c.Auth == nil {
		c.Auth = &AuthConfig{}
	}
	c.Auth.Type = AuthNone
	return c
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
		c.Auth = configData.Auth
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
	c.Auth = defaults.Auth
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
