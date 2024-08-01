/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-08-01 14:19:51
 * @FilePath: \go-config\config.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"github.com/kamalyes/go-config/captcha"
	"github.com/kamalyes/go-config/consul"
	"github.com/kamalyes/go-config/cors"
	"github.com/kamalyes/go-config/database"
	"github.com/kamalyes/go-config/email"
	"github.com/kamalyes/go-config/ftp"
	"github.com/kamalyes/go-config/jwt"
	"github.com/kamalyes/go-config/minio"
	"github.com/kamalyes/go-config/mqtt"
	"github.com/kamalyes/go-config/pay"
	"github.com/kamalyes/go-config/redis"
	"github.com/kamalyes/go-config/server"
	"github.com/kamalyes/go-config/zap"
	"github.com/spf13/viper"
)

// Config 项目单节点统一配置
type Config struct {

	/** 服务实例配置 */
	Server server.Server `mapstructure:"server"             json:"server"            yaml:"server"`

	/** Cors 配置 */
	Cors cors.Cors `mapstructure:"cors" json:"cors" yaml:"cors"`

	/** 注册中心配置 */
	Consul consul.Consul `mapstructure:"consul"             json:"consul"            yaml:"consul"`

	/** 验证码配置 */
	Captcha captcha.Captcha `mapstructure:"captcha"            json:"captcha"           yaml:"captcha"`

	/** MySQL 数据库配置 */
	MySQL database.MySQL `mapstructure:"mysql"              json:"mysql"             yaml:"mysql"`

	/** Postgre 数据库配置 */
	PostgreSQL database.PostgreSQL `mapstructure:"postgre"            json:"postgre"           yaml:"postgre"`

	/** sqlite 数据库配置 */
	SQLite database.SQLite `mapstructure:"sqlite"             json:"sqlite"            yaml:"sqlite"`

	/** redis 缓存数据库配置 */
	Redis redis.Redis `mapstructure:"redis"              json:"redis"             yaml:"redis"`

	/** 邮件发送相关配置 */
	Email email.Email `mapstructure:"email"              json:"email"             yaml:"email"`

	/** 文件服务器配置 */
	Ftp ftp.Ftp `mapstructure:"ftp"                json:"local"             yaml:"ftp"`

	/** jwt token 相关配置 */
	JWT jwt.JWT `mapstructure:"jwt"                json:"jwt"               yaml:"jwt"`

	/** minio 相关配置 */
	Minio minio.Minio `mapstructure:"minio"               json:"minio"              yaml:"minio"`

	/** mqtt 代理服务器相关配置 */
	Mqtt mqtt.Mqtt `mapstructure:"mqtt"               json:"mqtt"              yaml:"mqtt"`

	/** 日志相关配置 */
	Zap zap.Zap `mapstructure:"zap"                json:"zap"               yaml:"zap"`

	/** 支付相关配置-支付宝 */
	AliPay pay.Alipay `mapstructure:"alipay"             json:"alipay"            yaml:"alipay"`

	/** 支付相关配置-微信 */
	Wechat pay.Wechat `mapstructure:"wechat"            json:"wechat"           yaml:"wechat"`

	/** 可自己取一些扩展配置 */
	Viper *viper.Viper
}
