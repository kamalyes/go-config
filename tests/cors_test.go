/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2023-07-28 00:50:58
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-01 09:56:28
 * @FilePath: \go-config\tests\cors_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"reflect"
	"testing"

	"github.com/kamalyes/go-config/pkg/cors"
)

// 公共测试数据
var testCorsData = &cors.Cors{
	AllowedAllOrigins:   true,
	AllowedAllMethods:   false,
	AllowedOrigins:      []string{"https://example.com"},
	AllowedMethods:      []string{"GET", "POST"},
	AllowedHeaders:      []string{"Content-Type"},
	MaxAge:              "3600",
	ExposedHeaders:      []string{"X-Header"},
	AllowCredentials:    true,
	OptionsResponseCode: 200,
}

func TestNewCors(t *testing.T) {
	cors := cors.NewCors("testModule", testCorsData.AllowedAllOrigins, testCorsData.AllowedAllMethods, testCorsData.AllowedOrigins, testCorsData.AllowedMethods, testCorsData.AllowedHeaders, testCorsData.ExposedHeaders, testCorsData.MaxAge, testCorsData.AllowCredentials, testCorsData.OptionsResponseCode)

	if cors.ModuleName != "testModule" {
		t.Errorf("Expected ModuleName to be 'testModule', got '%s'", cors.ModuleName)
	}
	if cors.AllowedAllOrigins != testCorsData.AllowedAllOrigins {
		t.Error("Expected AllowedAllOrigins to be true")
	}
	if cors.AllowedAllMethods != testCorsData.AllowedAllMethods {
		t.Error("Expected AllowedAllMethods to be false")
	}
	if !reflect.DeepEqual(cors.AllowedOrigins, testCorsData.AllowedOrigins) {
		t.Errorf("Expected AllowedOrigins to be %v, got %v", testCorsData.AllowedOrigins, cors.AllowedOrigins)
	}
	if !reflect.DeepEqual(cors.AllowedMethods, testCorsData.AllowedMethods) {
		t.Errorf("Expected AllowedMethods to be %v, got %v", testCorsData.AllowedMethods, cors.AllowedMethods)
	}
	if !reflect.DeepEqual(cors.AllowedHeaders, testCorsData.AllowedHeaders) {
		t.Errorf("Expected AllowedHeaders to be %v, got %v", testCorsData.AllowedHeaders, cors.AllowedHeaders)
	}
	if cors.MaxAge != testCorsData.MaxAge {
		t.Errorf("Expected MaxAge to be '%s', got '%s'", testCorsData.MaxAge, cors.MaxAge)
	}
	if !reflect.DeepEqual(cors.ExposedHeaders, testCorsData.ExposedHeaders) {
		t.Errorf("Expected ExposedHeaders to be %v, got %v", testCorsData.ExposedHeaders, cors.ExposedHeaders)
	}
	if cors.AllowCredentials != testCorsData.AllowCredentials {
		t.Error("Expected AllowCredentials to be true")
	}
	if cors.OptionsResponseCode != testCorsData.OptionsResponseCode {
		t.Errorf("Expected OptionsResponseCode to be %d, got %d", testCorsData.OptionsResponseCode, cors.OptionsResponseCode)
	}
}

func TestCorsToMap(t *testing.T) {
	cors := cors.NewCors("testModule", testCorsData.AllowedAllOrigins, testCorsData.AllowedAllMethods, testCorsData.AllowedOrigins, testCorsData.AllowedMethods, testCorsData.AllowedHeaders, testCorsData.ExposedHeaders, testCorsData.MaxAge, testCorsData.AllowCredentials, testCorsData.OptionsResponseCode)
	result := cors.ToMap()

	expectedKeys := []string{"allowedAllOrigins", "allowedAllMethods", "allowedOrigins", "allowedMethods", "allowedHeaders", "maxAge", "exposedHeaders", "allowCredentials", "optionsResponseCode"}
	for _, key := range expectedKeys {
		if _, ok := result[key]; !ok {
			t.Errorf("Expected key '%s' not found in result map", key)
		}
	}
}

func TestCorsFromMap(t *testing.T) {
	cors := &cors.Cors{}
	data := map[string]interface{}{
		"allowedAllOrigins":   testCorsData.AllowedAllOrigins,
		"allowedAllMethods":   testCorsData.AllowedAllMethods,
		"allowedOrigins":      testCorsData.AllowedOrigins,
		"allowedMethods":      testCorsData.AllowedMethods,
		"allowedHeaders":      testCorsData.AllowedHeaders,
		"maxAge":              testCorsData.MaxAge,
		"exposedHeaders":      testCorsData.ExposedHeaders,
		"allowCredentials":    testCorsData.AllowCredentials,
		"optionsResponseCode": testCorsData.OptionsResponseCode,
	}

	cors.FromMap(data)

	if cors.AllowedAllOrigins != testCorsData.AllowedAllOrigins {
		t.Error("Expected AllowedAllOrigins to be true")
	}
	if cors.AllowedAllMethods != testCorsData.AllowedAllMethods {
		t.Error("Expected AllowedAllMethods to be false")
	}
	if !reflect.DeepEqual(cors.AllowedOrigins, testCorsData.AllowedOrigins) {
		t.Errorf("Expected AllowedOrigins to be %v, got %v", testCorsData.AllowedOrigins, cors.AllowedOrigins)
	}
}

func TestCorsClone(t *testing.T) {
	corsInit := cors.NewCors("testModule", testCorsData.AllowedAllOrigins, testCorsData.AllowedAllMethods, testCorsData.AllowedOrigins, testCorsData.AllowedMethods, testCorsData.AllowedHeaders, testCorsData.ExposedHeaders, testCorsData.MaxAge, testCorsData.AllowCredentials, testCorsData.OptionsResponseCode)
	clone := corsInit.Clone().(*cors.Cors)

	if !reflect.DeepEqual(corsInit, clone) {
		t.Error("Expected clone to be equal to original")
	}

	// Modify the clone and check that the original is unaffected
	clone.AllowedAllOrigins = false
	if corsInit.AllowedAllOrigins == clone.AllowedAllOrigins {
		t.Error("Expected original AllowedAllOrigins to be unaffected by clone modification")
	}
}
