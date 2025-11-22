package wsc

import (
	"fmt"
	"github.com/kamalyes/go-config/pkg/cache"
	"testing"
)

// TestExampleWSCChainUsage 展示WSC链式调用的使用示例
func TestExampleWSCChainUsage(t *testing.T) {
	// === 基础配置示例 ===
	basicConfig := Default().
		WithNodeIP("192.168.1.100").
		WithNodePort(8080).
		WithHeartbeatInterval(30).
		WithClientTimeout(300)

	fmt.Printf("Basic Config: IP=%s, Port=%d\n",
		basicConfig.NodeIP, basicConfig.NodePort)

	// === VIP配置示例 ===
	vipConfig := Default().
		WithVIP(DefaultVIP().
			WithMaxLevel(8).
			WithPriorityMultiplier(2.5))

	fmt.Printf("VIP Config: MaxLevel=%d, Multiplier=%.1f\n",
		vipConfig.VIP.MaxLevel, vipConfig.VIP.PriorityMultiplier)

	// === 增强功能配置示例 ===
	enhancedConfig := Default().
		WithEnhancement(DefaultEnhancement().
			WithSmartRouting(true).
			WithLoadBalancing(true, "round_robin").
			WithSmartQueue(true, 5000).
			WithMonitoring(true).
			WithCircuitBreaker(true, 10, 5, 30))

	fmt.Printf("Enhancement Config: SmartRouting=%t, LoadBalance=%s\n",
		enhancedConfig.Enhancement.SmartRouting,
		enhancedConfig.Enhancement.LoadBalanceAlgorithm)

	// === 分布式配置示例 ===
	distributedConfig := Default().
		EnableDistributed().
		WithRedis(&cache.Redis{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

	fmt.Printf("Distributed Config: Enabled=%t\n",
		distributedConfig.Distributed.Enabled)

	// === 群组和工单配置示例 ===
	groupTicketConfig := Default().
		WithGroup(DefaultGroup().
			Enable().
			WithMaxSize(200).
			WithMessageRecord(true))

	fmt.Printf("Group Config: MaxSize=%d, MessageRecord=%t\n",
		groupTicketConfig.Group.MaxGroupSize,
		groupTicketConfig.Group.EnableMessageRecord)
}

// TestComplexChainExample 复杂链式配置测试示例
func TestComplexChainExample(t *testing.T) {
	// 企业级完整配置
	enterpriseConfig := Default().
		// 基础服务配置
		WithNodeIP("10.0.1.100").
		WithNodePort(443).
		WithHeartbeatInterval(15).
		WithClientTimeout(600).
		WithMessageBufferSize(2048).
		WithSSEMessageBuffer(1024).

		// 分布式配置
		WithDistributed(DefaultDistributed().
			Enable()).
		WithRedis(&cache.Redis{
			Addr:         "redis-cluster.internal:6379",
			Password:     "enterprise-redis-secret",
			DB:           2,
			PoolSize:     20,
			MinIdleConns: 5,
		}).

		// VIP系统配置
		WithVIP(DefaultVIP().
			WithMaxLevel(8).
			WithPriorityMultiplier(3.0)).

		// 增强功能配置
		WithEnhancement(DefaultEnhancement().
			WithSmartRouting(true).
			WithLoadBalancing(true, "weighted_round_robin").
			WithSmartQueue(true, 10000).
			WithMonitoring(true).
			WithClusterManagement(true).
			WithCircuitBreaker(true, 20, 10, 60)).

		// 群组管理配置
		WithGroup(DefaultGroup().
			Enable().
			WithMaxSize(1000).
			WithMessageRecord(true)).

		// 性能优化配置
		WithPerformance(&Performance{
			WriteBufferSize: 16384,
			ReadBufferSize:  16384,
		}).

		// 安全策略配置
		WithSecurity(&Security{
			EnableAuth:        true,
			EnableEncryption:  true,
			EnableRateLimit:   true,
			MaxMessageSize:    1024,
			AllowedUserTypes:  []string{"admin", "agent", "vip_customer"},
			EnableIPWhitelist: true,
			WhitelistIPs:      []string{"10.0.0.0/8", "172.16.0.0/12"},
			TokenExpiration:   3600,
			MaxLoginAttempts:  3,
			LoginLockDuration: 900,
		})

	// 验证配置
	if enterpriseConfig.NodePort != 443 {
		t.Errorf("Expected NodePort=443, got %d", enterpriseConfig.NodePort)
	}

	if enterpriseConfig.VIP.MaxLevel != 8 {
		t.Errorf("Expected VIP MaxLevel=8, got %d", enterpriseConfig.VIP.MaxLevel)
	}

	if enterpriseConfig.Enhancement.LoadBalanceAlgorithm != "weighted_round_robin" {
		t.Errorf("Expected LoadBalanceAlgorithm=weighted_round_robin, got %s",
			enterpriseConfig.Enhancement.LoadBalanceAlgorithm)
	}
}
