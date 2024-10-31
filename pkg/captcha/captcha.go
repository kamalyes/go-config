/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 20:55:55
 * @FilePath: \go-config\pkg\captcha\captcha.go
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
	ModuleName string  `mapstructure:"MODULE_NAME" yaml:"modulename"` // 模块名称
	KeyLen     int     `mapstructure:"KEY_LEN" yaml:"key-len"`        // 数字或字符串长度
	ImgWidth   int     `mapstructure:"IMG_WIDTH" yaml:"img-width"`    // 验证码宽度
	ImgHeight  int     `mapstructure:"IMG_HEIGHT" yaml:"img-height"`  // 验证码高度
	MaxSkew    float64 `mapstructure:"MAX_SKEW" yaml:"max-skew"`      // 最大歪曲度 0.5-1.0
	DotCount   int     `mapstructure:"DOT_COUNT" yaml:"dot-count"`    // 分布的点的数量 推荐设置 100左右
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
	return nil
}
