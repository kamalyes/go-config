/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-11-01 17:52:06
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-11-01 18:18:27
 * @FilePath: \go-config\pkg\gateway\request_extract_bench_test.go
 * @Description: 按真实默认配置量属性提取的成本
 *
 * 为什么基准放在 pkg/gateway 而不是 pkg/common：提取逻辑在 common，但真实的来源清单在
 * gateway 的 DefaultRequestContext()，而 gateway 导入 common，反向导入会成环
 * 拿真实默认配置来量，才不会量成一个虚构的工作负载
 *
 * 默认配置的规模：42 个 header 来源、49 个 query 来源、29 个 cookie 来源，
 * 分布在 28 个属性的来源清单里。网关每个请求要把这些全走一遍
 *
 * [EN] The benchmark lives in pkg/gateway rather than pkg/common because the extraction logic is
 * in common while the real source lists are in gateway's DefaultRequestContext, and gateway
 * already imports common so the reverse would be a cycle. Measuring against the real default
 * configuration keeps this from benchmarking a fictional workload: 42 header, 49 query and 29
 * cookie sources spread over 28 attribute lists, all walked once per request
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kamalyes/go-config/pkg/common"
)

// allDefaultSources 把默认配置里全部来源清单摊平成「每个属性一份清单」
//
// 这正是调用方的真实用法：网关逐个属性调 ExtractAttribute，而不是只调一次
// [EN] Flattens every source list in the default configuration into one list per attribute,
// which is how a caller actually uses it: one ExtractAttribute call per attribute
func allDefaultSources() [][]common.AttributeSource {
	rc := DefaultRequestContext()
	return [][]common.AttributeSource{
		rc.IDSources, rc.TraceIDSources, rc.RequestIDSources, rc.UserIDSources,
		rc.DomainSources, rc.RoleCodeSources, rc.TenantIDSources, rc.TenantCodeSources,
		rc.SessionIDSources, rc.TimezoneSources, rc.TimestampSources, rc.SignatureSources,
		rc.AuthorizationSources, rc.JtiSources, rc.FamilyIdSources, rc.AccessKeySources,
		rc.AppIDSources, rc.DeviceIDSources, rc.AppVersionSources, rc.PlatformIDSources,
		rc.PlatformCodeSources, rc.RegionIDSources, rc.RegionCodeSources, rc.AgentLineIDSources,
		rc.NonceSources, rc.PushTokenSources, rc.TokenSources, rc.ForwardedHostSources,
	}
}

// benchExtractShapes 覆盖四种来源命中形态
//
// 形态很重要：ExtractAttribute 是「按顺序取第一个非空」，所以属性从 header 命中就提前返回、
// 根本不碰 query；属性哪儿都没有才会把所有来源走穿
// 最坏情况恰恰是常态 —— 绝大多数请求不会带满 28 个自定义头
// [EN] Shape matters: the first non-empty source wins, so a header hit returns before the query
// is touched while an absent attribute walks every source. The worst case is the common one
func benchExtractShapes() []struct {
	name string
	req  *http.Request
} {
	plain := httptest.NewRequest(http.MethodGet, "/api/v1/order/create", nil)

	withHeaders := httptest.NewRequest(http.MethodGet, "/api/v1/order/create", nil)
	withHeaders.Header.Set("X-Trace-Id", "trace-0123456789abcdef")
	withHeaders.Header.Set("X-Request-Id", "req-0123456789abcdef")
	withHeaders.Header.Set("X-User-ID", "u-100286")
	withHeaders.Header.Set("X-Tenant-ID", "t-37")
	withHeaders.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.e30.x")

	withQuery := httptest.NewRequest(http.MethodGet,
		"/api/v1/order/create?trace_id=trace-0123456789abcdef&request_id=req-0123456789abcdef"+
			"&tenant_id=t-37&page=3&page_size=20&sort=-created_at", nil)

	withCookies := httptest.NewRequest(http.MethodGet, "/api/v1/order/create", nil)
	withCookies.Header.Set("Cookie",
		"trace_id=trace-0123456789abcdef; request_id=req-0123456789abcdef; "+
			"session_id=sess-abcdef0123456789; theme=dark")

	mixed := httptest.NewRequest(http.MethodGet,
		"/api/v1/order/create?page=3&page_size=20", nil)
	mixed.Header.Set("X-Trace-Id", "trace-0123456789abcdef")
	mixed.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.e30.x")
	mixed.Header.Set("Cookie", "session_id=sess-abcdef0123456789; theme=dark")

	return []struct {
		name string
		req  *http.Request
	}{
		{"absent", plain},
		{"headers", withHeaders},
		{"query", withQuery},
		{"cookies", withCookies},
		{"mixed", mixed},
	}
}

var benchSink string

// BenchmarkExtractAttribute_DefaultSources 按真实默认配置量「取齐 28 个属性」的成本
func BenchmarkExtractAttribute_DefaultSources(b *testing.B) {
	lists := allDefaultSources()
	for _, shape := range benchExtractShapes() {
		b.Run(shape.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, sources := range lists {
					benchSink = common.ExtractAttribute(shape.req, sources)
				}
			}
		})
	}
}

// BenchmarkExtractAttribute_SourceCounts 把默认配置的来源规模打出来，便于解读上面的数字
func BenchmarkExtractAttribute_SourceCounts(b *testing.B) {
	var hdr, qry, ck int
	for _, sources := range allDefaultSources() {
		for _, s := range sources {
			switch s.Type {
			case common.SourceTypeHeader:
				hdr++
			case common.SourceTypeQuery:
				qry++
			case common.SourceTypeCookie:
				ck++
			}
		}
	}
	b.Logf("默认配置来源规模：header=%d query=%d cookie=%d，分布在 %d 个属性清单里",
		hdr, qry, ck, len(allDefaultSources()))
	b.SkipNow()
}

// BenchmarkAttributeExtractor_Reused 一次请求建一个提取器，28 个属性共用
//
// 与 BenchmarkExtractAttribute_DefaultSources 对照着看：那个是旧用法「一个属性一次调用」，
// 查询串与 Cookie 头在每次调用里各解析一次；这个是把解析摊到整个请求上，各只解析一次
// 两者的差值就是调用方改用提取器能拿到的收益
//
// [EN] Read against BenchmarkExtractAttribute_DefaultSources: that one is the per-attribute call
// pattern where the query string is parsed once per call, while this one amortizes the parsing
// across the whole request. The gap is what a caller gains by switching to the extractor
func BenchmarkAttributeExtractor_Reused(b *testing.B) {
	lists := allDefaultSources()
	for _, shape := range benchExtractShapes() {
		b.Run(shape.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				e := common.NewAttributeExtractor(shape.req)
				for _, sources := range lists {
					benchSink = e.Extract(sources)
				}
			}
		})
	}
}

// TestNormalizeHeaderKeysOnValidate 用户在 YAML 里写非规范 key 时，Validate 应把它归一
//
// 这是上面那条零分配的前提：归一化没落地的话，每次 Header.Get 都要现算 key 并分配
// [EN] The zero-allocation result above depends on this: without normalization every
// Header.Get recomputes and allocates the canonical key
func TestNormalizeHeaderKeysOnValidate(t *testing.T) {
	rc := &RequestContext{
		UserIDSources: []common.AttributeSource{
			{Type: common.SourceTypeHeader, Key: "x-user-id"}, // 用户常写的小写
			{Type: common.SourceTypeHeader, Key: "X-USER-ID"}, // 全大写
			{Type: common.SourceTypeQuery, Key: "user_id"},    // query 是大小写敏感的，不该被动
			{Type: common.SourceTypeCookie, Key: "User_Id"},   // cookie 同理
		},
	}
	if err := rc.Validate(); err != nil {
		t.Fatalf("Validate 失败: %v", err)
	}
	want := []string{"X-User-Id", "X-User-Id", "user_id", "User_Id"}
	for i, w := range want {
		if got := rc.UserIDSources[i].Key; got != w {
			t.Errorf("第 %d 个来源(%s)：期望 Key=%q，实得 %q",
				i, rc.UserIDSources[i].Type, w, got)
		}
	}
}

// BenchmarkAttributeExtractor_UserWrittenKeys 用户写非规范 key、经 Validate 归一后的成本
//
// 与 Reused/headers 对照：两者应当同样是零分配，说明归一化把用户的拼法差异吸收掉了
// [EN] Read against Reused/headers: both should be allocation-free, showing that normalization
// absorbs however the user spelled the keys
func BenchmarkAttributeExtractor_UserWrittenKeys(b *testing.B) {
	rc := &RequestContext{
		TraceIDSources:  []common.AttributeSource{{Type: common.SourceTypeHeader, Key: "x-trace-id"}},
		UserIDSources:   []common.AttributeSource{{Type: common.SourceTypeHeader, Key: "X-USER-ID"}},
		TenantIDSources: []common.AttributeSource{{Type: common.SourceTypeHeader, Key: "x-tenant-ID"}},
	}
	lists := [][]common.AttributeSource{rc.TraceIDSources, rc.UserIDSources, rc.TenantIDSources}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/order/create", nil)
	req.Header.Set("X-Trace-Id", "trace-0123456789abcdef")
	req.Header.Set("X-User-Id", "u-100286")
	req.Header.Set("X-Tenant-Id", "t-37")

	b.Run("归一化前", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			e := common.NewAttributeExtractor(req)
			for _, s := range lists {
				benchSink = e.Extract(s)
			}
		}
	})

	if err := rc.Validate(); err != nil {
		b.Fatalf("Validate 失败: %v", err)
	}

	b.Run("归一化后", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			e := common.NewAttributeExtractor(req)
			for _, s := range lists {
				benchSink = e.Extract(s)
			}
		}
	})
}
