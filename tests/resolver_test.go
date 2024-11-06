/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-06 13:57:00
 * @FilePath: \go-config\tests\resolver_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */

package tests

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	goconfig "github.com/kamalyes/go-config"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// createConfigFile 创建配置文件并写入内容
func createConfigFile(filename string, content string) error {
	// 确保目录存在
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// 写入内容到文件
	return os.WriteFile(filename, []byte(content), 0644)
}

// 测试全局配置加载
func TestMultiConfig(t *testing.T) {
	// 测试获取全局配置
	ctx := context.Background()
	// 使用自定义值创建 ConfigManager
	model := &goconfig.MultiConfig{}

	resultModel, resultJson, _ := random.GenerateRandomModel(model)
	resultJson = strings.ReplaceAll(resultJson, "module_name", "modulename")
	resultJson = strings.ReplaceAll(resultJson, "_", "-")

	createConfigFile("./resources/dev_config.yaml", resultJson)

	customManager, err := goconfig.NewMultiConfigManager(ctx, nil)

	if err != nil {
		fmt.Println("Error:", err)
	}

	assert.NotNil(t, customManager)
	config := customManager.GetConfig()
	// 获取模块
	modelx := resultModel.(*goconfig.MultiConfig)
	searchKey := modelx.Server[0].ModuleName
	searchValue := modelx.Server[0].Addr

	if modelx.Server != nil {
		module, err := goconfig.GetSingleConfigByModuleName(*config, searchKey)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			assert.Equal(t, searchValue, module.Server.Addr)
			fmt.Printf("Found module: %+v\n", module)
		}
	}
}

// 测试全局配置加载
func TestSingleConfig(t *testing.T) {
	// 测试获取全局配置
	ctx := context.Background()
	// 使用自定义值创建 ConfigManager
	model := &goconfig.SingleConfig{}

	resultModel, resultJson, _ := random.GenerateRandomModel(model)
	resultJson = strings.ReplaceAll(resultJson, "module_name", "modulename")
	resultJson = strings.ReplaceAll(resultJson, "_", "-")

	createConfigFile("./resources/dev_config.yaml", resultJson)

	customManager, err := goconfig.NewSingleConfigManager(ctx, nil)

	if err != nil {
		fmt.Println("Error:", err)
	}

	assert.NotNil(t, customManager)
	config := customManager.GetConfig()
	// 获取模块
	modelx := resultModel.(*goconfig.SingleConfig)

	// Server
	assert.Equal(t, modelx.Server, config.Server)
	assert.Equal(t, modelx.Server.ModuleName, config.Server.ModuleName)
	assert.Equal(t, modelx.Server.Addr, config.Server.Addr)
	assert.Equal(t, modelx.Server.ServerName, config.Server.ServerName)
	assert.Equal(t, modelx.Server.DataDriver, config.Server.DataDriver)
	assert.Equal(t, modelx.Server.HandleMethodNotAllowed, config.Server.HandleMethodNotAllowed)
	assert.Equal(t, modelx.Server.ContextPath, config.Server.ContextPath)
	assert.Equal(t, modelx.Server.Language, config.Server.Language)

	// Cors
	assert.Equal(t, modelx.Cors, config.Cors)
	assert.Equal(t, modelx.Cors.ModuleName, config.Cors.ModuleName)
	assert.Equal(t, modelx.Cors.AllowedOrigins, config.Cors.AllowedOrigins)
	assert.Equal(t, modelx.Cors.AllowedMethods, config.Cors.AllowedMethods)
	assert.Equal(t, modelx.Cors.AllowedHeaders, config.Cors.AllowedHeaders)
	assert.Equal(t, modelx.Cors.MaxAge, config.Cors.MaxAge)
	assert.Equal(t, modelx.Cors.AllowedAllMethods, config.Cors.AllowedAllMethods)
	assert.Equal(t, modelx.Cors.ExposedHeaders, config.Cors.ExposedHeaders)
	assert.Equal(t, modelx.Cors.AllowCredentials, config.Cors.AllowCredentials)
	assert.Equal(t, modelx.Cors.OptionsResponseCode, config.Cors.OptionsResponseCode)

	// Consul
	assert.Equal(t, modelx.Consul, config.Consul)
	assert.Equal(t, modelx.Consul.ModuleName, config.Consul.ModuleName)
	assert.Equal(t, modelx.Consul.Addr, config.Consul.Addr)
	assert.Equal(t, modelx.Consul.RegisterInterval, config.Consul.RegisterInterval)

	// Captcha
	assert.Equal(t, modelx.Captcha, config.Captcha)
	assert.Equal(t, modelx.Captcha.ModuleName, config.Captcha.ModuleName)
	assert.Equal(t, modelx.Captcha.KeyLen, config.Captcha.KeyLen)
	assert.Equal(t, modelx.Captcha.ImgWidth, config.Captcha.ImgWidth)
	assert.Equal(t, modelx.Captcha.ImgHeight, config.Captcha.ImgHeight)
	assert.Equal(t, modelx.Captcha.MaxSkew, config.Captcha.MaxSkew)
	assert.Equal(t, modelx.Captcha.DotCount, config.Captcha.DotCount)

	// MySQL
	assert.Equal(t, modelx.MySQL, config.MySQL)
	assert.Equal(t, modelx.MySQL.ModuleName, config.MySQL.ModuleName)
	assert.Equal(t, modelx.MySQL.Host, config.MySQL.Host)
	assert.Equal(t, modelx.MySQL.Port, config.MySQL.Port)
	assert.Equal(t, modelx.MySQL.Config, config.MySQL.Config)
	assert.Equal(t, modelx.MySQL.LogLevel, config.MySQL.LogLevel)
	assert.Equal(t, modelx.MySQL.Dbname, config.MySQL.Dbname)
	assert.Equal(t, modelx.MySQL.Username, config.MySQL.Username)
	assert.Equal(t, modelx.MySQL.Password, config.MySQL.Password)
	assert.Equal(t, modelx.MySQL.MaxIdleConns, config.MySQL.MaxIdleConns)
	assert.Equal(t, modelx.MySQL.MaxOpenConns, config.MySQL.MaxOpenConns)
	assert.Equal(t, modelx.MySQL.ConnMaxIdleTime, config.MySQL.ConnMaxIdleTime)
	assert.Equal(t, modelx.MySQL.ConnMaxLifeTime, config.MySQL.ConnMaxLifeTime)

	// PostgreSQL
	assert.Equal(t, modelx.PostgreSQL, config.PostgreSQL)
	assert.Equal(t, modelx.PostgreSQL.ModuleName, config.PostgreSQL.ModuleName)
	assert.Equal(t, modelx.PostgreSQL.Host, config.PostgreSQL.Host)
	assert.Equal(t, modelx.PostgreSQL.Port, config.PostgreSQL.Port)
	assert.Equal(t, modelx.PostgreSQL.Config, config.PostgreSQL.Config)
	assert.Equal(t, modelx.PostgreSQL.LogLevel, config.PostgreSQL.LogLevel)
	assert.Equal(t, modelx.PostgreSQL.Dbname, config.PostgreSQL.Dbname)
	assert.Equal(t, modelx.PostgreSQL.Username, config.PostgreSQL.Username)
	assert.Equal(t, modelx.PostgreSQL.Password, config.PostgreSQL.Password)
	assert.Equal(t, modelx.PostgreSQL.MaxIdleConns, config.PostgreSQL.MaxIdleConns)
	assert.Equal(t, modelx.PostgreSQL.MaxOpenConns, config.PostgreSQL.MaxOpenConns)
	assert.Equal(t, modelx.PostgreSQL.ConnMaxIdleTime, config.PostgreSQL.ConnMaxIdleTime)
	assert.Equal(t, modelx.PostgreSQL.ConnMaxLifeTime, config.PostgreSQL.ConnMaxLifeTime)

	// SqlLite
	assert.Equal(t, modelx.SQLite, config.SQLite)
	assert.Equal(t, modelx.SQLite.ModuleName, config.SQLite.ModuleName)
	assert.Equal(t, modelx.SQLite.DbPath, config.SQLite.DbPath)
	assert.Equal(t, modelx.SQLite.MaxIdleConns, config.SQLite.MaxIdleConns)
	assert.Equal(t, modelx.SQLite.MaxOpenConns, config.SQLite.MaxOpenConns)
	assert.Equal(t, modelx.SQLite.LogLevel, config.SQLite.LogLevel)
	assert.Equal(t, modelx.SQLite.ConnMaxIdleTime, config.SQLite.ConnMaxIdleTime)
	assert.Equal(t, modelx.SQLite.ConnMaxLifeTime, config.SQLite.ConnMaxLifeTime)
	assert.Equal(t, modelx.SQLite.Vacuum, config.SQLite.Vacuum)

	// Redis
	assert.Equal(t, modelx.Redis, config.Redis)
	assert.Equal(t, modelx.Redis.ModuleName, config.Redis.ModuleName)
	assert.Equal(t, modelx.Redis.Addr, config.Redis.Addr)
	assert.Equal(t, modelx.Redis.DB, config.Redis.DB)
	assert.Equal(t, modelx.Redis.MaxRetries, config.Redis.MaxRetries)
	assert.Equal(t, modelx.Redis.MinIdleConns, config.Redis.MinIdleConns)
	assert.Equal(t, modelx.Redis.PoolSize, config.Redis.PoolSize)
	assert.Equal(t, modelx.Redis.Password, config.Redis.Password)

	// Email
	assert.Equal(t, modelx.Email, config.Email)
	assert.Equal(t, modelx.Email.ModuleName, config.Email.ModuleName)
	assert.Equal(t, modelx.Email.To, config.Email.To)
	assert.Equal(t, modelx.Email.From, config.Email.From)
	assert.Equal(t, modelx.Email.Host, config.Email.Host)
	assert.Equal(t, modelx.Email.Port, config.Email.Port)
	assert.Equal(t, modelx.Email.Secret, config.Email.Secret)
	assert.Equal(t, modelx.Email.IsSSL, config.Email.IsSSL)

	// Ftp
	assert.Equal(t, modelx.Ftp, config.Ftp)
	assert.Equal(t, modelx.Ftp.ModuleName, config.Ftp.ModuleName)
	assert.Equal(t, modelx.Ftp.Addr, config.Ftp.Addr)
	assert.Equal(t, modelx.Ftp.Username, config.Ftp.Username)
	assert.Equal(t, modelx.Ftp.Password, config.Ftp.Password)
	assert.Equal(t, modelx.Ftp.Cwd, config.Ftp.Cwd)

	// JWT
	assert.Equal(t, modelx.JWT, config.JWT)
	assert.Equal(t, modelx.JWT.ModuleName, config.JWT.ModuleName)
	assert.Equal(t, modelx.JWT.SigningKey, config.JWT.SigningKey)
	assert.Equal(t, modelx.JWT.ExpiresTime, config.JWT.ExpiresTime)
	assert.Equal(t, modelx.JWT.BufferTime, config.JWT.BufferTime)
	assert.Equal(t, modelx.JWT.UseMultipoint, config.JWT.UseMultipoint)

	// Minio
	assert.Equal(t, modelx.Minio, config.Minio)
	assert.Equal(t, modelx.Minio.ModuleName, config.Minio.ModuleName)
	assert.Equal(t, modelx.Minio.Host, config.Minio.Host)
	assert.Equal(t, modelx.Minio.Port, config.Minio.Port)
	assert.Equal(t, modelx.Minio.AccessKey, config.Minio.AccessKey)
	assert.Equal(t, modelx.Minio.SecretKey, config.Minio.SecretKey)

	// AliyunOss
	assert.Equal(t, modelx.AliyunOss, config.AliyunOss)
	assert.Equal(t, modelx.AliyunOss.ModuleName, config.AliyunOss.ModuleName)
	assert.Equal(t, modelx.AliyunOss.AccessKey, config.AliyunOss.AccessKey)
	assert.Equal(t, modelx.AliyunOss.SecretKey, config.AliyunOss.SecretKey)
	assert.Equal(t, modelx.AliyunOss.Endpoint, config.AliyunOss.Endpoint)
	assert.Equal(t, modelx.AliyunOss.Bucket, config.AliyunOss.Bucket)
	assert.Equal(t, modelx.AliyunOss.ReplaceOriginalHost, config.AliyunOss.ReplaceOriginalHost)
	assert.Equal(t, modelx.AliyunOss.ReplaceLaterHost, config.AliyunOss.ReplaceLaterHost)

	// Mqtt
	assert.Equal(t, modelx.Mqtt, config.Mqtt)
	assert.Equal(t, modelx.Mqtt.ModuleName, config.Mqtt.ModuleName)
	assert.Equal(t, modelx.Mqtt.Url, config.Mqtt.Url)
	assert.Equal(t, modelx.Mqtt.ProtocolVersion, config.Mqtt.ProtocolVersion)
	assert.Equal(t, modelx.Mqtt.KeepAlive, config.Mqtt.KeepAlive)
	assert.Equal(t, modelx.Mqtt.MaxReconnectInterval, config.Mqtt.MaxReconnectInterval)
	assert.Equal(t, modelx.Mqtt.PingTimeout, config.Mqtt.PingTimeout)
	assert.Equal(t, modelx.Mqtt.WriteTimeout, config.Mqtt.WriteTimeout)
	assert.Equal(t, modelx.Mqtt.ConnectTimeout, config.Mqtt.ConnectTimeout)
	assert.Equal(t, modelx.Mqtt.Username, config.Mqtt.Username)
	assert.Equal(t, modelx.Mqtt.Password, config.Mqtt.Password)
	assert.Equal(t, modelx.Mqtt.CleanSession, config.Mqtt.CleanSession)
	assert.Equal(t, modelx.Mqtt.AutoReconnect, config.Mqtt.AutoReconnect)
	assert.Equal(t, modelx.Mqtt.WillTopic, config.Mqtt.WillTopic)

	// Zap
	assert.Equal(t, modelx.Zap, config.Zap)
	assert.Equal(t, modelx.Zap.ModuleName, config.Zap.ModuleName)
	assert.Equal(t, modelx.Zap.Level, config.Zap.Level)
	assert.Equal(t, modelx.Zap.Format, config.Zap.Format)
	assert.Equal(t, modelx.Zap.Prefix, config.Zap.Prefix)
	assert.Equal(t, modelx.Zap.Director, config.Zap.Director)
	assert.Equal(t, modelx.Zap.MaxSize, config.Zap.MaxSize)
	assert.Equal(t, modelx.Zap.MaxAge, config.Zap.MaxAge)
	assert.Equal(t, modelx.Zap.MaxBackups, config.Zap.MaxBackups)
	assert.Equal(t, modelx.Zap.Compress, config.Zap.Compress)
	assert.Equal(t, modelx.Zap.LinkName, config.Zap.LinkName)
	assert.Equal(t, modelx.Zap.ShowLine, config.Zap.ShowLine)
	assert.Equal(t, modelx.Zap.EncodeLevel, config.Zap.EncodeLevel)
	assert.Equal(t, modelx.Zap.StacktraceKey, config.Zap.StacktraceKey)
	assert.Equal(t, modelx.Zap.LogInConsole, config.Zap.LogInConsole)

	// AliPay
	assert.Equal(t, modelx.AliPay, config.AliPay)
	assert.Equal(t, modelx.AliPay.ModuleName, config.AliPay.ModuleName)
	assert.Equal(t, modelx.AliPay.Pid, config.AliPay.Pid)
	assert.Equal(t, modelx.AliPay.AppId, config.AliPay.AppId)
	assert.Equal(t, modelx.AliPay.PriKey, config.AliPay.PriKey)
	assert.Equal(t, modelx.AliPay.PubKey, config.AliPay.PubKey)
	assert.Equal(t, modelx.AliPay.SignType, config.AliPay.SignType)
	assert.Equal(t, modelx.AliPay.NotifyUrl, config.AliPay.NotifyUrl)
	assert.Equal(t, modelx.AliPay.Subject, config.AliPay.Subject)

	// WechatPay
	assert.Equal(t, modelx.WechatPay, config.WechatPay)
	assert.Equal(t, modelx.WechatPay.ModuleName, config.WechatPay.ModuleName)
	assert.Equal(t, modelx.WechatPay.AppId, config.WechatPay.AppId)
	assert.Equal(t, modelx.WechatPay.ApiKey, config.WechatPay.ApiKey)
	assert.Equal(t, modelx.WechatPay.MchId, config.WechatPay.MchId)
	assert.Equal(t, modelx.WechatPay.NotifyUrl, config.WechatPay.NotifyUrl)
	assert.Equal(t, modelx.WechatPay.CertP12Path, config.WechatPay.CertP12Path)

	// AliyunSms
	assert.Equal(t, modelx.AliyunSms, config.AliyunSms)
	assert.Equal(t, modelx.AliyunSms.ModuleName, config.AliyunSms.ModuleName)
	assert.Equal(t, modelx.AliyunSms.SecretID, config.AliyunSms.SecretID)
	assert.Equal(t, modelx.AliyunSms.SecretKey, config.AliyunSms.SecretKey)
	assert.Equal(t, modelx.AliyunSms.Sign, config.AliyunSms.Sign)
	assert.Equal(t, modelx.AliyunSms.ResourceOwnerAccount, config.AliyunSms.ResourceOwnerAccount)
	assert.Equal(t, modelx.AliyunSms.ResourceOwnerID, config.AliyunSms.ResourceOwnerID)
	assert.Equal(t, modelx.AliyunSms.TemplateCodeVerify, config.AliyunSms.TemplateCodeVerify)
	assert.Equal(t, modelx.AliyunSms.Endpoint, config.AliyunSms.Endpoint)

	// AliyunSts
	assert.Equal(t, modelx.AliyunSts, config.AliyunSts)
	assert.Equal(t, modelx.AliyunSts.ModuleName, config.AliyunSts.ModuleName)
	assert.Equal(t, modelx.AliyunSts.RegionID, config.AliyunSts.RegionID)
	assert.Equal(t, modelx.AliyunSts.AccessKeyID, config.AliyunSts.AccessKeyID)
	assert.Equal(t, modelx.AliyunSts.AccessKeySecret, config.AliyunSts.AccessKeySecret)
	assert.Equal(t, modelx.AliyunSts.RoleArn, config.AliyunSts.RoleArn)
	assert.Equal(t, modelx.AliyunSts.RoleSessionName, config.AliyunSts.RoleSessionName)

	// Youzan
	assert.Equal(t, modelx.Youzan, config.Youzan)
	assert.Equal(t, modelx.Youzan.ModuleName, config.Youzan.ModuleName)
	assert.Equal(t, modelx.Youzan.Host, config.Youzan.Host)
	assert.Equal(t, modelx.Youzan.ClientID, config.Youzan.ClientID)
	assert.Equal(t, modelx.Youzan.ClientSecret, config.Youzan.ClientSecret)
	assert.Equal(t, modelx.Youzan.AuthorizeType, config.Youzan.AuthorizeType)
	assert.Equal(t, modelx.Youzan.GrantID, config.Youzan.GrantID)
	assert.Equal(t, modelx.Youzan.Refresh, config.Youzan.Refresh)

	// ZeroServer
	assert.Equal(t, modelx.ZeroServer, config.ZeroServer)
	assert.Equal(t, modelx.ZeroServer.ModuleName, config.ZeroServer.ModuleName)
	assert.Equal(t, modelx.ZeroServer.ListenOn, config.ZeroServer.ListenOn)
	assert.Equal(t, modelx.ZeroServer.Timeout, config.ZeroServer.Timeout)
	assert.Equal(t, modelx.ZeroServer.CpuThreshold, config.ZeroServer.CpuThreshold)
	assert.Equal(t, modelx.ZeroServer.Etcd, config.ZeroServer.Etcd)
	assert.Equal(t, modelx.ZeroServer.Auth, config.ZeroServer.Auth)
	assert.Equal(t, modelx.ZeroServer.StrictControl, config.ZeroServer.StrictControl)

	// ZeroClient
	assert.Equal(t, modelx.ZeroClient, config.ZeroClient)
	assert.Equal(t, modelx.ZeroClient.ModuleName, config.ZeroClient.ModuleName)
	assert.Equal(t, modelx.ZeroClient.Target, config.ZeroClient.Target)
	assert.Equal(t, modelx.ZeroClient.App, config.ZeroClient.App)
	assert.Equal(t, modelx.ZeroClient.Timeout, config.ZeroClient.Timeout)
	assert.Equal(t, modelx.ZeroClient.NonBlock, config.ZeroClient.NonBlock)

}
