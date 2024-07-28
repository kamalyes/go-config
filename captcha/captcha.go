/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-28 09:24:50
 * @FilePath: \go-config\captcha\captcha.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package captcha

// Captcha 验证码格式配置
type Captcha struct {

	/** 数字或字符串长度 */
	KeyLen int `mapstructure:"key-len"           json:"keyLen"         yaml:"key-len"`

	/** 验证码宽度 */
	ImgWidth int `mapstructure:"img-width"         json:"imgWidth"       yaml:"img-width"`

	/** 验证码高度 */
	ImgHeight int `mapstructure:"img-height"        json:"imgHeight"      yaml:"img-height"`

	/** 最大歪曲度 0.5-1.0 */
	MaxSkew float64 `mapstructure:"max-skew"          json:"maxSkew"        yaml:"max-skew"`

	/** 分布的点的数量  推荐设置 100左右 */
	DotCount int `mapstructure:"dot-count"         json:"dotCount"       yaml:"dot-count"`
}
