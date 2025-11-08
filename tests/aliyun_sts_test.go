/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 13:10:06
 * @FilePath: \go-config\tests\aliyun_sts_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/sts"
	"github.com/kamalyes/go-toolbox/pkg/random"
	"github.com/stretchr/testify/assert"
)

// 生成随机的 AliyunSts 配置参数
func generateAliyunStsTestParams() *sts.AliyunSts {
	return &sts.AliyunSts{
		ModuleName:      random.RandString(10, random.CAPITAL),
		RegionID:        random.RandString(10, random.CAPITAL),
		AccessKeyID:     random.RandString(10, random.CAPITAL),
		AccessKeySecret: random.RandString(10, random.CAPITAL),
		RoleArn:         random.RandString(10, random.CAPITAL),
		RoleSessionName: random.RandString(10, random.CAPITAL),
	}
}

func TestAliyunStsClone(t *testing.T) {
	params := generateAliyunStsTestParams()
	original := sts.NewAliyunSts(params)
	cloned := original.Clone().(*sts.AliyunSts)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned) // 确保是不同的实例
}

func TestAliyunStsSet(t *testing.T) {
	oldParams := generateAliyunStsTestParams()
	newParams := generateAliyunStsTestParams()

	aliyunSts := sts.NewAliyunSts(oldParams)
	newConfig := sts.NewAliyunSts(newParams)

	aliyunSts.Set(newConfig)

	assert.Equal(t, newParams.ModuleName, aliyunSts.ModuleName)
	assert.Equal(t, newParams.RegionID, aliyunSts.RegionID)
	assert.Equal(t, newParams.AccessKeyID, aliyunSts.AccessKeyID)
	assert.Equal(t, newParams.AccessKeySecret, aliyunSts.AccessKeySecret)
	assert.Equal(t, newParams.RoleArn, aliyunSts.RoleArn)
	assert.Equal(t, newParams.RoleSessionName, aliyunSts.RoleSessionName)
}
