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
package ftp

type Ftp struct {

	/** ftp 服务器Ip和端口 */
	Addr string `mapstructure:"addr"                json:"addr"              yaml:"addr"`

	/** 用户 */
	Username string `mapstructure:"username"            json:"username"          yaml:"username"`

	/** 密码 */
	Password string `mapstructure:"password"            json:"password"          yaml:"password"`

	/** 指定目录 */
	Cwd string `mapstructure:"cwd"                 json:"cwd"               yaml:"cwd"`
}
