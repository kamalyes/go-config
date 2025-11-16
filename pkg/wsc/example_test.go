package wsc

import (
	"fmt"
	"testing"

	"github.com/kamalyes/go-config/pkg/cache"
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
			WithMessageRecord(true)).
		WithTicket(DefaultTicket().
			Enable().
			WithAck(true, 5000, 3).
			WithMaxPerAgent(15))

	fmt.Printf("Group Config: MaxSize=%d, MessageRecord=%t\n", 
		groupTicketConfig.Group.MaxGroupSize,
		groupTicketConfig.Group.EnableMessageRecord)
	fmt.Printf("Ticket Config: MaxPerAgent=%d, AckTimeout=%dms\n", 
		groupTicketConfig.Ticket.MaxTicketsPerAgent,
		groupTicketConfig.Ticket.AckTimeoutMs)
}

// TestExampleWSCSafeAccess 展示WSC安全访问的使用示例
func TestExampleWSCSafeAccess(t *testing.T) {
	config := Default().
		WithNodePort(8080).
		WithVIP(DefaultVIP().
			WithMaxLevel(5)).
		WithEnhancement(DefaultEnhancement().
			WithCircuitBreaker(true, 15, 8, 45))

	// 创建安全访问器
	safe := config.Safe()

	// 安全访问基础配置
	nodePort := safe.NodePort(3000) // 默认值3000
	fmt.Printf("Node Port: %d\n", nodePort)

	// 安全访问VIP配置
	vipMaxLevel := safe.VIP().Field("MaxLevel").Int(0)
	fmt.Printf("VIP Max Level: %d\n", vipMaxLevel)

	// 安全访问增强功能配置
	circuitBreakerEnabled := safe.Enhancement().Field("CircuitBreaker").Bool(false)
	failureThreshold := safe.Enhancement().Field("FailureThreshold").Int(5)
	fmt.Printf("Circuit Breaker: %t, Failure Threshold: %d\n", 
		circuitBreakerEnabled, failureThreshold)
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
		
		// 工单系统配置
		WithTicket(DefaultTicket().
			Enable().
			WithAck(true, 3000, 5).
			WithMaxPerAgent(30)).
		
		// 性能优化配置
		WithPerformance(&Performance{
			WriteBufferSize:   16384,
			ReadBufferSize:    16384,
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

	if enterpriseConfig.Ticket.MaxTicketsPerAgent != 30 {
		t.Errorf("Expected MaxTicketsPerAgent=30, got %d", 
			enterpriseConfig.Ticket.MaxTicketsPerAgent)
	}

	// 测试安全访问
	safe := enterpriseConfig.Safe()

	// 验证安全访问各项配置
	nodeIP := safe.NodeIP("")
	if nodeIP != "10.0.1.100" {
		t.Errorf("Expected safe NodeIP=10.0.1.100, got %s", nodeIP)
	}

	vipEnabled := safe.VIP().Field("Enabled").Bool()
	if !vipEnabled {
		t.Error("Expected VIP enabled via safe access")
	}

	enhancementEnabled := safe.Enhancement().Field("Enabled").Bool()
	if !enhancementEnabled {
		t.Error("Expected Enhancement enabled via safe access")
	}

	writeBufferSize := safe.Performance().Field("WriteBufferSize").Int()
	if writeBufferSize != 16384 {
		t.Errorf("Expected WriteBufferSize=16384, got %d", writeBufferSize)
	}

	maxMessageSize := safe.Security().Field("MaxMessageSize").Int()
	if maxMessageSize != 1024 {
		t.Errorf("Expected MaxMessageSize=1024, got %d", maxMessageSize)
	}

	t.Log("✅ Complex enterprise configuration test passed successfully")
	t.Logf("   - Node: %s:%d", nodeIP, enterpriseConfig.NodePort)
	t.Logf("   - VIP Level: %d (×%.1f priority)", 
		enterpriseConfig.VIP.MaxLevel, enterpriseConfig.VIP.PriorityMultiplier)
	t.Logf("   - Load Balancing: %s", enterpriseConfig.Enhancement.LoadBalanceAlgorithm)
	t.Logf("   - Buffer Sizes: Write=%dKB, Read=%dKB", 
		writeBufferSize/1024, safe.Performance().Field("ReadBufferSize").Int()/1024)
}