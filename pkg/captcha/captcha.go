/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 10:05:18
 * @FilePath: \go-config\captcha\captcha.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package captcha

import (
	"github.com/kamalyes/go-config/internal"
)

// Captcha 验证码格式配置
type Captcha struct {

	/** 基础 **/
	ModuleName string `mapstructure:"modulename" json:"moduleName" yaml:"modulename"`

	// 数字或字符串长度
	KeyLen int `mapstructure:"key-len" json:"keyLen" yaml:"key-len"`

	// 验证码宽度
	ImgWidth int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`

	// 验证码高度
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`

	// 最大歪曲度 0.5-1.0
	MaxSkew float64 `mapstructure:"max-skew" json:"maxSkew" yaml:"max-skew"`

	// 分布的点的数量 推荐设置 100左右
	DotCount int `mapstructure:"dot-count" json:"dotCount" yaml:"dot-count"`
}

// NewCaptcha 创建一个新的 Captcha 实例
func NewCaptcha(moduleName string, keyLen, imgWidth, imgHeight int, maxSkew float64, dotCount int) *Captcha {
	var captchaInstance *Captcha

	internal.LockFunc(func() {
		captchaInstance = &Captcha{
			ModuleName: moduleName,
			KeyLen:     keyLen,
			ImgWidth:   imgWidth,
			ImgHeight:  imgHeight,
			MaxSkew:    maxSkew,
			DotCount:   dotCount,
		}
	})
	return captchaInstance
}

// ToMap 将配置转换为映射
func (c *Captcha) ToMap() map[string]interface{} {
	return internal.ToMap(c)
}

// FromMap 从映射中填充配置
func (c *Captcha) FromMap(data map[string]interface{}) {
	internal.FromMap(c, data)
}

// Clone 返回 Captcha 配置的副本
func (c *Captcha) Clone() internal.Configurable {
	return &Captcha{
		ModuleName: c.ModuleName,
		KeyLen:     c.KeyLen,
		ImgWidth:   c.ImgWidth,
		ImgHeight:  c.ImgHeight,
		MaxSkew:    c.MaxSkew,
		DotCount:   c.DotCount,
	}
}

// Get 返回 Captcha 配置的所有字段
func (c *Captcha) Get() interface{} {
	return c
}

// Set 更新 Captcha 配置的字段
func (c *Captcha) Set(data interface{}) {
	if configData, ok := data.(*Captcha); ok {
		c.ModuleName = configData.ModuleName
		c.KeyLen = configData.KeyLen
		c.ImgWidth = configData.ImgWidth
		c.ImgHeight = configData.ImgHeight
		c.MaxSkew = configData.MaxSkew
		c.DotCount = configData.DotCount
	}
}

// Validate 检查 Captcha 配置的有效性
func (c *Captcha) Validate() error {
	return nil
}
