/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-16 13:59:50
 * @FilePath: \go-config\cors\cors.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package cors

// Cors 跨域配置结构体
type Cors struct {
	AllowedAllOrigins   bool     `mapstructure:"allowed-all-origins" json:"allowedAllOrigins" yaml:"allowed-all-origins"`       // 是否允许所有来源
	AllowedAllMethods   bool     `mapstructure:"allowed-all-methods" json:"allowedAllMethods" yaml:"allowed-all-methods"`       // 是否允许所有方法
	AllowedOrigins      []string `mapstructure:"allowed-origins" json:"allowedOrigins" yaml:"allowed-origins"`                  // 允许的来源
	AllowedMethods      []string `mapstructure:"allowed-methods" json:"allowedMethods" yaml:"allowed-methods"`                  // 允许的方法
	AllowedHeaders      []string `mapstructure:"allowed-headers" json:"allowedHeaders" yaml:"allowed-headers"`                  // 允许的头部
	MaxAge              string   `mapstructure:"max-age" json:"maxAge" yaml:"max-age"`                                          // 最大缓存时间
	ExposedHeaders      []string `mapstructure:"exposed-headers" json:"exposedHeaders" yaml:"exposed-headers"`                  // 暴露的头部
	AllowCredentials    bool     `mapstructure:"allow-credentials" json:"allowCredentials" yaml:"allow-credentials"`            // 允许凭证
	OptionsResponseCode int      `mapstructure:"options-response-code" json:"optionsResponseCode" yaml:"options-response-code"` // Options响应Code
}
