/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-07 23:39:27
 * @FilePath: \go-config\pkg\captcha\captcha.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package captcha

import (
	"github.com/kamalyes/go-config/internal"
)

// Captcha 结构体
type Captcha struct {
	KeyLen     int     `mapstructure:"key-len"     yaml:"key-len"     json:"key_len"      validate:"required,min=1"`           // 数字或字符串长度
	ImgWidth   int     `mapstructure:"img-width"   yaml:"img-width"   json:"img_width"    validate:"required,min=1"`           // 验证码宽度
	ImgHeight  int     `mapstructure:"img-height"  yaml:"img-height"  json:"img_height"   validate:"required,min=1"`           // 验证码高度
	MaxSkew    float64 `mapstructure:"max-skew"    yaml:"max-skew"    json:"max_skew"     validate:"required,min=0.5,max=1.0"` // 最大歪曲度 0.5-1.0
	DotCount   int     `mapstructure:"dot-count"   yaml:"dot-count"   json:"dot_count"    validate:"required,min=1"`           // 分布的点的数量 推荐设置 100左右
	ModuleName string  `mapstructure:"modulename"  yaml:"modulename"  json:"module_name"  `                                    // 模块名称
}

// NewCaptcha 创建一个新的 Captcha 实例
func NewCaptcha(opt *Captcha) *Captcha {
	var captchaInstance *Captcha

	internal.LockFunc(func() {
		captchaInstance = opt
	})
	return captchaInstance
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
	return internal.ValidateStruct(c)
}

// DefaultCaptcha 返回默认Captcha配置
func DefaultCaptcha() Captcha {
	return Captcha{
		ModuleName: "captcha",
		KeyLen:     4,
		ImgWidth:   120,
		ImgHeight:  40,
		MaxSkew:    0.7,
		DotCount:   80,
	}
}

// Default 返回默认Captcha配置的指针，支持链式调用
func Default() *Captcha {
	config := DefaultCaptcha()
	return &config
}

// WithModuleName 设置模块名称
func (c *Captcha) WithModuleName(moduleName string) *Captcha {
	c.ModuleName = moduleName
	return c
}

// WithKeyLen 设置验证码长度
func (c *Captcha) WithKeyLen(keyLen int) *Captcha {
	c.KeyLen = keyLen
	return c
}

// WithImgWidth 设置验证码宽度
func (c *Captcha) WithImgWidth(imgWidth int) *Captcha {
	c.ImgWidth = imgWidth
	return c
}

// WithImgHeight 设置验证码高度
func (c *Captcha) WithImgHeight(imgHeight int) *Captcha {
	c.ImgHeight = imgHeight
	return c
}

// WithMaxSkew 设置最大歪曲度
func (c *Captcha) WithMaxSkew(maxSkew float64) *Captcha {
	c.MaxSkew = maxSkew
	return c
}

// WithDotCount 设置点的数量
func (c *Captcha) WithDotCount(dotCount int) *Captcha {
	c.DotCount = dotCount
	return c
}
