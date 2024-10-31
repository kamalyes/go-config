/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 17:00:31
 * @FilePath: \go-config\tests\email_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/email"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 Email 配置参数
func generateEmailTestParams() *email.Email {
	return &email.Email{
		ModuleName: random.RandString(10, random.CAPITAL),                                                                                    // 随机生成模块名称
		To:         fmt.Sprintf("%s@example.com,%s@example.com", random.RandString(5, random.CAPITAL), random.RandString(5, random.CAPITAL)), // 随机生成多个收件人
		From:       fmt.Sprintf("%s@example.com", random.RandString(5, random.CAPITAL)),                                                      // 随机生成发件人
		Host:       fmt.Sprintf("smtp.%s.com", random.RandString(5, random.CAPITAL)),                                                         // 随机生成邮件服务器地址
		Port:       random.RandInt(1024, 65535),                                                                                              // 随机生成端口
		IsSSL:      random.FRandBool(),                                                                                                       // 随机生成是否使用 SSL
		Secret:     random.RandString(16, random.CAPITAL),                                                                                    // 随机生成密钥
	}
}

func TestEmailClone(t *testing.T) {
	params := generateEmailTestParams()
	original := email.NewEmail(params)
	cloned := original.Clone().(*email.Email) // 假设您有 Clone 方法

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestEmailSet(t *testing.T) {
	oldParams := generateEmailTestParams()
	newParams := generateEmailTestParams()

	emailInstance := email.NewEmail(oldParams)
	newConfig := email.NewEmail(newParams)

	emailInstance.Set(newConfig) // 假设您有 Set 方法

	assert.Equal(t, newParams.ModuleName, emailInstance.ModuleName)
	assert.Equal(t, newParams.To, emailInstance.To)
	assert.Equal(t, newParams.From, emailInstance.From)
	assert.Equal(t, newParams.Host, emailInstance.Host)
	assert.Equal(t, newParams.Port, emailInstance.Port)
	assert.Equal(t, newParams.IsSSL, emailInstance.IsSSL)
	assert.Equal(t, newParams.Secret, emailInstance.Secret)
}
