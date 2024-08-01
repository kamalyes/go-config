/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-07-31 20:29:54
 * @FilePath: \go-config\global\global.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package global

import (
	"github.com/bwmarrin/snowflake"
	"github.com/casbin/casbin/v2"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v8"
	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-config/env"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 单节点应用常用全局变量
var (

	// ENV 设置环境
	ENV env.Environment

	// DB 数据库
	DB *gorm.DB

	// REDIS 默认客户端
	REDIS *redis.Client

	// MQTT 客户端
	MQTT *mqtt.Client

	// CONFIG 全局系统配置
	CONFIG *goconfig.Config

	// VP 通过 viper 读取的yaml配置文件
	VP *viper.Viper

	// LOG 全局日志
	LOG *zap.Logger

	// CSBEF casbin实施者
	CSBEF casbin.IEnforcer

	// 雪花ID节点
	Node *snowflake.Node

	// MinIO客户端
	MinIO *minio.Client
)
