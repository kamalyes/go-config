package wsc

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/cache"
	"github.com/stretchr/testify/assert"
)

// TestWSCChainMethods 测试所有链式调用方法
func TestWSCChainMethods(t *testing.T) {
	// 测试基础链式调用
	config := Default().
		WithNodeIP("192.168.1.100").
		WithNodePort(8080).
		EnableDistributed().
		EnableRedis().
		EnableGroup().
		EnableTicket()

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

	// 测试Ticket链式调用
	ticketConfig := DefaultTicket().
		WithAck(true, 5000, 3).
		WithMaxPerAgent(20)
	config = config.WithTicket(ticketConfig)

	assert.True(t, config.Ticket.EnableAck, "Ticket EnableAck should be true")
	assert.Equal(t, 5000, config.Ticket.AckTimeoutMs, "Ticket AckTimeoutMs should be 5000")
	assert.Equal(t, 3, config.Ticket.MaxRetry, "Ticket MaxRetry should be 3")
	assert.Equal(t, 20, config.Ticket.MaxTicketsPerAgent, "Ticket MaxTicketsPerAgent should be 20")
}

// TestWSCSafeAccess 测试安全访问模式
func TestWSCSafeAccess(t *testing.T) {
	config := Default().
		WithNodePort(8080).
		WithVIP(DefaultVIP().
			WithMaxLevel(5)).
		WithEnhancement(DefaultEnhancement().
			WithCircuitBreaker(true, 15, 8, 45))

	safeConfig := config.Safe()

	// 测试安全获取基础配置
	nodePort := safeConfig.Field("NodePort").Int()
	if nodePort != 8080 {
		t.Errorf("Expected SafeConfig NodePort=8080, got %d", nodePort)
	}

	// 测试安全获取布尔值
	vipEnabled := safeConfig.VIP().Field("Enabled").Bool()
	if !vipEnabled {
		t.Error("Expected SafeConfig VIP.Enabled=true")
	}

	enhancementEnabled := safeConfig.Enhancement().Field("Enabled").Bool()
	if !enhancementEnabled {
		t.Error("Expected SafeConfig Enhancement.Enabled=true")
	}

	circuitBreakerEnabled := safeConfig.Enhancement().Field("CircuitBreaker").Bool()
	if !circuitBreakerEnabled {
		t.Error("Expected SafeConfig Enhancement.CircuitBreaker=true")
	}
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
		WithTicket(DefaultTicket().
			Enable().
			WithAck(true, 3000, 5).
			WithMaxPerAgent(50)).
		WithPerformance(&Performance{
			WriteBufferSize:   8192,
			ReadBufferSize:    8192,
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

	if enterpriseConfig.Ticket.MaxTicketsPerAgent != 50 {
		t.Errorf("Expected Ticket.MaxTicketsPerAgent=50, got %d", enterpriseConfig.Ticket.MaxTicketsPerAgent)
	}

	// 测试安全访问
	safeConfig := enterpriseConfig.Safe()
	writeBufferSize := safeConfig.Performance().Field("WriteBufferSize").Int()
	if writeBufferSize != 8192 {
		t.Errorf("Expected Performance.WriteBufferSize=8192, got %d", writeBufferSize)
	}

	t.Logf("✅ Enterprise configuration test passed with %d write buffer size", writeBufferSize)
}