/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-11-03 20:54:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-11-03 21:08:17
 * @FilePath: \go-config\tests\database_common_test.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package tests

import (
	"testing"

	"github.com/kamalyes/go-config/pkg/database"
)

// TestMySQLConfig 测试 MySQL 配置
func TestMySQLConfig(t *testing.T) {
	mysqlConfig := database.MySQL{
		Host:            "localhost",
		Username:        "root",
		Password:        "password",
		Dbname:          "testdb",
		Port:            "3306",
		Config:          "charset=utf8",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxIdleTime: 30,
		ConnMaxLifeTime: 60,
		LogLevel:        "info",
	}

	commonConfig := mysqlConfig.GetCommonConfig()

	// 验证返回的配置是否与原始配置一致
	if commonConfig.Host != mysqlConfig.Host {
		t.Errorf("expected Host %s, got %s", mysqlConfig.Host, commonConfig.Host)
	}
	if commonConfig.Username != mysqlConfig.Username {
		t.Errorf("expected Username %s, got %s", mysqlConfig.Username, commonConfig.Username)
	}
	if commonConfig.Password != mysqlConfig.Password {
		t.Errorf("expected Password %s, got %s", mysqlConfig.Password, commonConfig.Password)
	}
	if commonConfig.Dbname != mysqlConfig.Dbname {
		t.Errorf("expected Dbname %s, got %s", mysqlConfig.Dbname, commonConfig.Dbname)
	}
}

// TestPostgreSQLConfig 测试 PostgreSQL 配置
func TestPostgreSQLConfig(t *testing.T) {
	postgresConfig := database.PostgreSQL{
		Host:            "localhost",
		Username:        "postgres",
		Password:        "password",
		Dbname:          "testdb",
		Port:            "5432",
		Config:          "sslmode=disable",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxIdleTime: 30,
		ConnMaxLifeTime: 60,
		LogLevel:        "info",
	}

	commonConfig := postgresConfig.GetCommonConfig()

	if commonConfig.Host != postgresConfig.Host {
		t.Errorf("expected Host %s, got %s", postgresConfig.Host, commonConfig.Host)
	}
	if commonConfig.Username != postgresConfig.Username {
		t.Errorf("expected Username %s, got %s", postgresConfig.Username, commonConfig.Username)
	}
	if commonConfig.Password != postgresConfig.Password {
		t.Errorf("expected Password %s, got %s", postgresConfig.Password, commonConfig.Password)
	}
	if commonConfig.Dbname != postgresConfig.Dbname {
		t.Errorf("expected Dbname %s, got %s", postgresConfig.Dbname, commonConfig.Dbname)
	}
}

// TestSQLiteConfig 测试 SQLite 配置
func TestSQLiteConfig(t *testing.T) {
	sqliteConfig := database.SQLite{
		DbPath:          "test.db",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxIdleTime: 30,
		ConnMaxLifeTime: 60,
		LogLevel:        "info",
		Vacuum:          true,
	}

	commonConfig := sqliteConfig.GetCommonConfig()

	if commonConfig.DbPath != sqliteConfig.DbPath {
		t.Errorf("expected DbPath %s, got %s", sqliteConfig.DbPath, commonConfig.DbPath)
	}
}

// TestNewDBConfig 测试 NewDBConfig 函数
func TestNewDBConfig(t *testing.T) {

	tests := []struct {
		name      string
		dbType    string
		config    database.DBConfigInterface
		expectErr bool
	}{
		{
			name:   "Valid MySQL Config",
			dbType: database.DBTypeMySQL,
			config: database.MySQL{
				Host:            "localhost",
				Username:        "root",
				Password:        "password",
				Dbname:          "testdb",
				Port:            "3306",
				Config:          "charset=utf8",
				MaxIdleConns:    10,
				MaxOpenConns:    100,
				ConnMaxIdleTime: 30,
				ConnMaxLifeTime: 60,
				LogLevel:        "info",
			},
			expectErr: false,
		},
		{
			name:   "Valid PostgreSQL Config",
			dbType: database.DBTypePostgres,
			config: database.PostgreSQL{
				Host:            "localhost",
				Username:        "postgres",
				Password:        "password",
				Dbname:          "testdb",
				Port:            "5432",
				Config:          "sslmode=disable",
				MaxIdleConns:    10,
				MaxOpenConns:    100,
				ConnMaxIdleTime: 30,
				ConnMaxLifeTime: 60,
				LogLevel:        "info",
			},
			expectErr: false,
		},
		{
			name:   "Valid SQLite Config",
			dbType: database.DBTypeSQLite,
			config: database.SQLite{
				DbPath:          "test.db",
				MaxIdleConns:    10,
				MaxOpenConns:    100,
				ConnMaxIdleTime: 30,
				ConnMaxLifeTime: 60,
				LogLevel:        "info",
				Vacuum:          true,
			},
			expectErr: false,
		},
		{
			name:   "Invalid DB Type",
			dbType: "invalid_db_type",
			config: database.MySQL{
				Host: "localhost",
			},
			expectErr: true,
		},
		{
			name:      "Nil Config",
			dbType:    database.DBTypeMySQL,
			config:    nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := database.NewDBConfig(tt.dbType, tt.config)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
				return // 继续下一个测试用例
			} else {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
			}

			// 检查返回的配置是否与输入配置一致
			switch tt.dbType {
			case database.DBTypeMySQL:
				mysqlConfig, ok := tt.config.(database.MySQL)
				if !ok {
					t.Fatalf("config should be of type MySQL")
				}
				if config.Host != mysqlConfig.Host {
					t.Errorf("expected Host %s, got %s", mysqlConfig.Host, config.Host)
				}
			case database.DBTypePostgres:
				postgresConfig, ok := tt.config.(database.PostgreSQL)
				if !ok {
					t.Fatalf("config should be of type PostgreSQL")
				}
				if config.Host != postgresConfig.Host {
					t.Errorf("expected Host %s, got %s", postgresConfig.Host, config.Host)
				}
			case database.DBTypeSQLite:
				sqliteConfig, ok := tt.config.(database.SQLite)
				if !ok {
					t.Fatalf("config should be of type SQLite")
				}
				if config.DbPath != sqliteConfig.DbPath {
					t.Errorf("expected DbPath %s, got %s", sqliteConfig.DbPath, config.DbPath)
				}
			}
		})
	}
}
