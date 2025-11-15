/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-11-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-11-15 13:53:37
 * @FilePath: \go-config\config_safe_access_test.go
 * @Description: 配置安全访问辅助工具测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package goconfig

import (
	"fmt"
	"testing"
	"time"

	"github.com/kamalyes/go-toolbox/pkg/safe"
)

// TestSafeConfig_BasicAccess 测试基本安全访问
func TestSafeConfig_BasicAccessX(t *testing.T) {
	config := map[string]interface{}{
		"Health": map[string]interface{}{
			"Enabled": true,
			"Redis": map[string]interface{}{
				"Enabled": true,
				"Timeout": "30s",
			},
		},
	}

	safeConfig := safe.Safe(config)

	// 测试每一步的访问
	fmt.Println("原始配置:", config)
	fmt.Println("SafeAccess值:", safeConfig.Value())

	health := safeConfig.Field("Health")
	fmt.Println("Health字段:", health.Value())
	fmt.Println("Health是否有效:", health.IsValid())

	if health.IsValid() {
		redis := health.Field("Redis")
		fmt.Println("Health.Redis字段:", redis.Value())
		fmt.Println("Health.Redis是否有效:", redis.IsValid())

		if redis.IsValid() {
			timeout := redis.Field("Timeout")
			fmt.Println("Health.Redis.Timeout字段:", timeout.Value())
			fmt.Println("Health.Redis.Timeout是否有效:", timeout.IsValid())

			timeoutDuration := timeout.Duration(10 * time.Second)
			fmt.Println("解析后的Duration:", timeoutDuration)
		}
	}
}
