/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-11 18:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-12 23:58:45
 * @FilePath: \go-config\pkg\gateway\gateway.go
 * @Description: Gateway网关统一配置模块
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package gateway

import (
	"github.com/kamalyes/go-config/internal"
	"github.com/kamalyes/go-config/pkg/banner"
	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/kamalyes/go-config/pkg/cors"
	"github.com/kamalyes/go-config/pkg/database"
	"github.com/kamalyes/go-config/pkg/elk"
	"github.com/kamalyes/go-config/pkg/health"
	"github.com/kamalyes/go-config/pkg/jwt"
	"github.com/kamalyes/go-config/pkg/middleware"
	"github.com/kamalyes/go-config/pkg/monitoring"
	"github.com/kamalyes/go-config/pkg/oss"
	"github.com/kamalyes/go-config/pkg/queue"
	"github.com/kamalyes/go-config/pkg/ratelimit"
	"github.com/kamalyes/go-config/pkg/security"
	"github.com/kamalyes/go-config/pkg/smtp"
	"github.com/kamalyes/go-config/pkg/swagger"
)

// Gateway网关统一配置
type Gateway struct {
	ModuleName    string                 `mapstructure:"module_name" yaml:"module-name" json:"module_name"`       // 模块名称
	Name          string                 `mapstructure:"name" yaml:"name" json:"name"`                            // 网关名称
	Enabled       bool                   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`                   // 是否启用网关
	Debug         bool                   `mapstructure:"debug" yaml:"debug" json:"debug"`                         // 是否启用调试模式
	Version       string                 `mapstructure:"version" yaml:"version" json:"version"`                   // 版本号
	Environment   string                 `mapstructure:"environment" yaml:"environment" json:"environment"`       // 环境 (dev, test, prod)
	HTTPServer    *HTTPServer            `mapstructure:"http" yaml:"http" json:"http"`                            // HTTP服务器配置
	GRPC          *GRPC                  `mapstructure:"grpc" yaml:"grpc" json:"grpc"`                            // GRPC配置
	Cache         *cache.Cache           `mapstructure:"cache" yaml:"cache" json:"cache"`                         // 缓存配置(包含Redis)
	Database      *database.Database     `mapstructure:"database" yaml:"database" json:"database"`              // 数据库统一配置
	OSS           *oss.OSSConfig         `mapstructure:"oss" yaml:"oss" json:"oss"`                               // 对象存储统一配置
	Mqtt          *queue.Mqtt            `mapstructure:"mqtt" yaml:"mqtt" json:"mqtt"`                            // MQTT配置
	Elasticsearch *elk.Elasticsearch     `mapstructure:"elasticsearch" yaml:"elasticsearch" json:"elasticsearch"` // Elasticsearch配置
	Smtp          *smtp.Smtp             `mapstructure:"smtp" yaml:"smtp" json:"smtp"`                            // SMTP邮件服务配置
	Health        *health.Health         `mapstructure:"health" yaml:"health" json:"health"`                      // 健康检查配置
	Monitoring    *monitoring.Monitoring `mapstructure:"monitoring" yaml:"monitoring" json:"monitoring"`          // 监控配置
	Security      *security.Security     `mapstructure:"security" yaml:"security" json:"security"`                // 安全配置
	Middleware    *middleware.Middleware `mapstructure:"middleware" yaml:"middleware" json:"middleware"`          // 中间件配置
	CORS          *cors.Cors             `mapstructure:"cors" yaml:"cors" json:"cors"`                            // CORS配置
	JWT           *jwt.JWT               `mapstructure:"jwt" yaml:"jwt" json:"jwt"`                               // JWT配置
	Swagger       *swagger.Swagger       `mapstructure:"swagger" yaml:"swagger" json:"swagger"`                   // Swagger配置
	Banner        *banner.Banner         `mapstructure:"banner" yaml:"banner" json:"banner"`                      // Banner配置
	RateLimit     *ratelimit.RateLimit   `mapstructure:"rate_limit" yaml:"rate_limit" json:"rate_limit"`          // 限流配置
}

// Default 创建默认Gateway配置
func Default() *Gateway {
	return &Gateway{
		ModuleName:    "gateway",
		Name:          "Go RPC Gateway",
		Enabled:       true,
		Debug:         false,
		Version:       "v1.0.0",
		Environment:   "dev",
		HTTPServer:    DefaultHTTPServer(),
		GRPC:          DefaultGRPC(),
		Cache:         cache.Default(),
		Database:      database.DefaultDatabaseConfig(),
		OSS:           oss.DefaultOSSConfig(),
		Mqtt:          queue.Default(),
		Elasticsearch: elk.Default(),
		Smtp:          smtp.Default(),
		Health:        health.Default(),
		Monitoring:    monitoring.Default(),
		Security:      security.Default(),
		Middleware:    middleware.Default(),
		CORS:          cors.Default(),
		JWT:           jwt.Default(),
		RateLimit: ratelimit.Default(),
		Swagger:       swagger.Default(),
		Banner:        banner.Default(),
	}
}

// Get 返回配置接口
func (c *Gateway) Get() interface{} {
	return c
}

// Set 设置配置数据
func (c *Gateway) Set(data interface{}) {
	if cfg, ok := data.(*Gateway); ok {
		*c = *cfg
	}
}

// Clone 返回配置的副本
func (c *Gateway) Clone() internal.Configurable {
	return &Gateway{
		ModuleName:    c.ModuleName,
		Name:          c.Name,
		Enabled:       c.Enabled,
		Debug:         c.Debug,
		Version:       c.Version,
		Environment:   c.Environment,
		HTTPServer:    c.HTTPServer.Clone(),
		GRPC:          c.GRPC.Clone(),
		Cache:         c.Cache.Clone().(*cache.Cache),
		Database:      c.Database.Clone().(*database.Database),
		OSS:           c.OSS.Clone().(*oss.OSSConfig),
		Mqtt:          c.Mqtt.Clone().(*queue.Mqtt),
		Elasticsearch: c.Elasticsearch.Clone().(*elk.Elasticsearch),
		Health:        c.Health.Clone().(*health.Health),
		Monitoring:    c.Monitoring.Clone().(*monitoring.Monitoring),
		Security:      c.Security.Clone().(*security.Security),
		Middleware:    c.Middleware.Clone().(*middleware.Middleware),
		CORS:          c.CORS.Clone().(*cors.Cors),
		JWT:           c.JWT.Clone().(*jwt.JWT),
		RateLimit:     c.RateLimit.Clone().(*ratelimit.RateLimit),
		Swagger:       c.Swagger.Clone().(*swagger.Swagger),
		Banner:        c.Banner.Clone().(*banner.Banner),
	}
}

// Validate 验证配置
func (c *Gateway) Validate() error {
	if err := internal.ValidateStruct(c); err != nil {
		return err
	}

	// 验证子配置
	if c.Cache != nil {
		if err := c.Cache.Validate(); err != nil {
			return err
		}
	}
	if c.Database != nil {
		if err := c.Database.Validate(); err != nil {
			return err
		}
	}
	if c.Mqtt != nil {
		if err := c.Mqtt.Validate(); err != nil {
			return err
		}
	}
	if c.OSS != nil {
		if err := c.OSS.Validate(); err != nil {
			return err
		}
	}
	if c.Elasticsearch != nil {
		if err := c.Elasticsearch.Validate(); err != nil {
			return err
		}
	}
	if c.Smtp != nil {
		if err := c.Smtp.Validate(); err != nil {
			return err
		}
	}
	if c.Health != nil {
		if err := c.Health.Validate(); err != nil {
			return err
		}
	}
	if c.Monitoring != nil {
		if err := c.Monitoring.Validate(); err != nil {
			return err
		}
	}
	if c.Security != nil {
		if err := c.Security.Validate(); err != nil {
			return err
		}
	}
	if c.Middleware != nil {
		if err := c.Middleware.Validate(); err != nil {
			return err
		}
	}
	if c.CORS != nil {
		if err := c.CORS.Validate(); err != nil {
			return err
		}
	}
	if c.CORS != nil {
		if err := c.CORS.Validate(); err != nil {
			return err
		}
	}
	if c.JWT != nil {
		if err := c.JWT.Validate(); err != nil {
			return err
		}
	}
	if c.RateLimit != nil {
		if err := c.RateLimit.Validate(); err != nil {
			return err
		}
	}
	if c.Swagger != nil {
		if err := c.Swagger.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// WithModuleName 设置模块名称
func (c *Gateway) WithModuleName(moduleName string) *Gateway {
	c.ModuleName = moduleName
	return c
}

// WithName 设置网关名称
func (c *Gateway) WithName(name string) *Gateway {
	c.Name = name
	return c
}

// WithEnabled 设置是否启用网关
func (c *Gateway) WithEnabled(enabled bool) *Gateway {
	c.Enabled = enabled
	return c
}

// WithDebug 设置调试模式
func (c *Gateway) WithDebug(debug bool) *Gateway {
	c.Debug = debug
	return c
}

// WithVersion 设置版本号
func (c *Gateway) WithVersion(version string) *Gateway {
	c.Version = version
	return c
}

// WithEnvironment 设置环境
func (c *Gateway) WithEnvironment(environment string) *Gateway {
	c.Environment = environment
	return c
}

// WithServer 设置服务器配置
func (c *Gateway) WithServer(host string, port int, grpcPort int) *Gateway {
	if c.HTTPServer == nil {
		c.HTTPServer = &HTTPServer{}
	}
	c.HTTPServer.Host = host
	c.HTTPServer.Port = port
	c.HTTPServer.GrpcPort = grpcPort
	return c
}

// WithTLS 设置TLS配置
func (c *Gateway) WithTLS(enabled bool, certFile, keyFile, caFile string) *Gateway {
	if c.HTTPServer == nil {
		c.HTTPServer = &HTTPServer{}
	}
	if c.HTTPServer.TLS == nil {
		c.HTTPServer.TLS = &TLS{}
	}
	c.HTTPServer.EnableTls = enabled
	c.HTTPServer.TLS.CertFile = certFile
	c.HTTPServer.TLS.KeyFile = keyFile
	c.HTTPServer.TLS.CAFile = caFile
	return c
}

// WithBanner 设置Banner配置
func (c *Gateway) WithBanner(enabled bool, title, description, author, email, website string) *Gateway {
	if c.Banner == nil {
		c.Banner = &banner.Banner{}
	}
	c.Banner.Enabled = enabled
	c.Banner.Title = title
	c.Banner.Description = description
	c.Banner.Author = author
	c.Banner.Email = email
	c.Banner.Website = website
	return c
}

// WithHealth 设置健康检查配置
func (c *Gateway) WithHealth(cfg *health.Health) *Gateway {
	c.Health = cfg
	return c
}

// WithMonitoring 设置监控配置
func (c *Gateway) WithMonitoring(cfg *monitoring.Monitoring) *Gateway {
	c.Monitoring = cfg
	return c
}

// WithSecurity 设置安全配置
func (c *Gateway) WithSecurity(cfg *security.Security) *Gateway {
	c.Security = cfg
	return c
}

// WithMiddleware 设置中间件配置
func (c *Gateway) WithMiddleware(cfg *middleware.Middleware) *Gateway {
	c.Middleware = cfg
	return c
}

// WithCORS 设置CORS配置
func (c *Gateway) WithCORS(cfg *cors.Cors) *Gateway {
	c.CORS = cfg
	return c
}

// WithJWT 设置JWT配置
func (c *Gateway) WithJWT(cfg *jwt.JWT) *Gateway {
	c.JWT = cfg
	return c
}

// WithSwagger 设置Swagger配置
func (c *Gateway) WithSwagger(cfg *swagger.Swagger) *Gateway {
	c.Swagger = cfg
	return c
}

// EnableHTTP 启用HTTP服务
func (c *Gateway) EnableHTTP() *Gateway {
	if c.HTTPServer == nil {
		c.HTTPServer = &HTTPServer{}
	}
	c.HTTPServer.EnableHttp = true
	return c
}

// EnableGRPC 启用GRPC服务
func (c *Gateway) EnableGRPC() *Gateway {
	if c.HTTPServer == nil {
		c.HTTPServer = &HTTPServer{}
	}
	c.HTTPServer.EnableGrpc = true
	return c
}

// EnableTLS 启用TLS
func (c *Gateway) EnableTLS() *Gateway {
	if c.HTTPServer == nil {
		c.HTTPServer = &HTTPServer{}
	}
	c.HTTPServer.EnableTls = true
	return c
}

// EnableBanner 启用Banner
func (c *Gateway) EnableBanner() *Gateway {
	if c.Banner == nil {
		c.Banner = &banner.Banner{}
	}
	c.Banner.Enabled = true
	return c
}

// EnableHealth 启用健康检查
func (c *Gateway) EnableHealth() *Gateway {
	if c.Health != nil {
		c.Health.Enable()
	}
	return c
}

// EnableMonitoring 启用监控
func (c *Gateway) EnableMonitoring() *Gateway {
	if c.Monitoring != nil {
		c.Monitoring.Enable()
	}
	return c
}

// EnableSecurity 启用安全功能
func (c *Gateway) EnableSecurity() *Gateway {
	if c.Security != nil {
		c.Security.Enable()
	}
	return c
}

// EnableMiddleware 启用中间件
func (c *Gateway) EnableMiddleware() *Gateway {
	if c.Middleware != nil {
		c.Middleware.Enable()
	}
	return c
}

// EnableSwagger 启用Swagger
func (c *Gateway) EnableSwagger() *Gateway {
	if c.Swagger != nil {
		c.Swagger.Enable()
	}
	return c
}

// Enable 启用网关
func (c *Gateway) Enable() *Gateway {
	c.Enabled = true
	return c
}

// Disable 禁用网关
func (c *Gateway) Disable() *Gateway {
	c.Enabled = false
	return c
}

// IsEnabled 检查是否启用
func (c *Gateway) IsEnabled() bool {
	return c.Enabled
}
