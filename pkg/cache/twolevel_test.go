/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-24 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-12-24 00:00:00
 * @FilePath: \go-config\pkg\cache\twolevel_test.go
 * @Description: TwoLevel缓存配置测试
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoLevel_Clone(t *testing.T) {
	twoLevel := &TwoLevel{
		ModuleName:       "twolevel",
		L1Type:           "memory",
		L2Type:           "redis",
		L1TTL:            300,
		L2TTL:            3600,
		SyncStrategy:     "write-through",
		L1Size:           1000,
		L2Size:           10000,
		PromoteThreshold: 3,
	}

	cloned := twoLevel.Clone()

	assert.NotNil(t, cloned)
	clonedTwoLevel, ok := cloned.(*TwoLevel)
	assert.True(t, ok)
	assert.Equal(t, twoLevel.L1Type, clonedTwoLevel.L1Type)
	assert.Equal(t, twoLevel.L2Type, clonedTwoLevel.L2Type)
	assert.Equal(t, twoLevel.L1TTL, clonedTwoLevel.L1TTL)
	assert.Equal(t, twoLevel.L2TTL, clonedTwoLevel.L2TTL)
	assert.Equal(t, twoLevel.SyncStrategy, clonedTwoLevel.SyncStrategy)
	assert.Equal(t, twoLevel.PromoteThreshold, clonedTwoLevel.PromoteThreshold)

	// 验证是独立副本
	clonedTwoLevel.L1TTL = 600
	assert.NotEqual(t, twoLevel.L1TTL, clonedTwoLevel.L1TTL)
}
