package wsc

import (
	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestWSCChainMethods 测试所有链式调用方法
func TestWSCChainMethods(t *testing.T) {
	// 测试基础链式调用
	config := Default().
		WithNodeIP("192.168.1.100").
		WithNodePort(8080).
		EnableDistributed().
		EnableRedis().
		EnableGroup()

	// 验证基础配置
	assert.Equal(t, "192.168.1.100", config.NodeIP, "NodeIP should be 192.168.1.100")
	assert.Equal(t, 8080, config.NodePort, "NodePort should be 8080")

	// 测试VIP链式调用
	vipConfig := DefaultVIP().
		WithMaxLevel(8).
		WithPriorityMultiplier(2.5)
	config = config.WithVIP(vipConfig)

	assert.True(t, config.VIP.Enabled, "VIP should be enabled")
	assert.Equal(t, 8, config.VIP.MaxLevel, "VIP MaxLevel should be 8")
	assert.Equal(t, 2.5, config.VIP.PriorityMultiplier, "VIP PriorityMultiplier should be 2.5")

	// 测试Enhancement链式调用
	enhancementConfig := DefaultEnhancement().
		Enable().
		WithSmartRouting(true).
		WithLoadBalancing(true, "round_robin").
		WithSmartQueue(true, 10000).
		WithMonitoring(true).
		WithClusterManagement(true).
		WithCircuitBreaker(true, 10, 5, 30)
	config = config.WithEnhancement(enhancementConfig)

	assert.True(t, config.Enhancement.Enabled, "Enhancement should be enabled")
	assert.True(t, config.Enhancement.SmartRouting, "SmartRouting should be enabled")
	assert.True(t, config.Enhancement.LoadBalancing, "LoadBalancing should be enabled")
	assert.Equal(t, "round_robin", config.Enhancement.LoadBalanceAlgorithm, "LoadBalanceAlgorithm should be round_robin")
	assert.Equal(t, 10000, config.Enhancement.MaxQueueSize, "MaxQueueSize should be 10000")
	assert.Equal(t, 10, config.Enhancement.FailureThreshold, "FailureThreshold should be 10")

	// 测试Group链式调用
	groupConfig := DefaultGroup().
		WithMaxSize(500).
		WithMessageRecord(true)
	config = config.WithGroup(groupConfig)

	assert.Equal(t, 500, config.Group.MaxGroupSize, "Group MaxGroupSize should be 500")
	assert.True(t, config.Group.EnableMessageRecord, "Group EnableMessageRecord should be true")
}

// TestCompleteChainConfiguration 测试完整链式配置
func TestCompleteChainConfiguration(t *testing.T) {
	// 模拟企业级配置场景
	enterpriseConfig := Default().
		WithNodeIP("10.0.1.100").
		WithNodePort(443).
		WithDistributed(DefaultDistributed().Enable()).
		WithRedis(&cache.Redis{
			Addr:     "redis-cluster:6379",
			Password: "enterprise-secret",
			DB:       1,
		}).
		WithVIP(DefaultVIP().
			WithMaxLevel(8).
			WithPriorityMultiplier(3.0)).
		WithEnhancement(DefaultEnhancement().
			WithSmartRouting(true).
			WithLoadBalancing(true, "weighted_round_robin").
			WithSmartQueue(true, 50000).
			WithMonitoring(true).
			WithClusterManagement(true).
			WithCircuitBreaker(true, 20, 10, 60)).
		WithGroup(DefaultGroup().
			Enable().
			WithMaxSize(1000).
			WithMessageRecord(true)).
		WithPerformance(&Performance{
			WriteBufferSize: 8192,
			ReadBufferSize:  8192,
		}).
		WithSecurity(&Security{
			EnableAuth:       true,
			EnableEncryption: true,
		})

	// 验证完整配置
	if enterpriseConfig.NodePort != 443 {
		t.Errorf("Expected enterprise NodePort=443, got %d", enterpriseConfig.NodePort)
	}

	if enterpriseConfig.VIP.MaxLevel != 8 {
		t.Errorf("Expected enterprise VIP.MaxLevel=8, got %d", enterpriseConfig.VIP.MaxLevel)
	}

	if enterpriseConfig.Enhancement.LoadBalanceAlgorithm != "weighted_round_robin" {
		t.Errorf("Expected LoadBalanceAlgorithm=weighted_round_robin, got %s", enterpriseConfig.Enhancement.LoadBalanceAlgorithm)
	}

	if enterpriseConfig.Group.MaxGroupSize != 1000 {
		t.Errorf("Expected Group.MaxGroupSize=1000, got %d", enterpriseConfig.Group.MaxGroupSize)
	}
}
