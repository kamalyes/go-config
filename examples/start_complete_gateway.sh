#!/bin/bash

# 完整Gateway演示服务启动脚本

echo "🚀 启动完整Gateway演示服务..."

# 设置环境变量
export APP_ENV=development
export CONFIG_PATH=./complete-gateway-config.yaml

# 检查Go环境
if ! command -v go &> /dev/null; then
    echo "❌ Go环境未安装，请先安装Go 1.19+"
    exit 1
fi

# 检查配置文件
if [ ! -f "./complete-gateway-config.yaml" ]; then
    echo "❌ 配置文件 complete-gateway-config.yaml 不存在"
    echo "请确保配置文件在当前目录下"
    exit 1
fi

echo "✅ 环境检查完成"

# 显示服务信息
echo ""
echo "📋 服务信息:"
echo "   环境: $APP_ENV"
echo "   配置文件: $CONFIG_PATH"
echo "   服务地址: http://localhost:8080"
echo ""

# 启动服务
echo "🔥 启动服务..."
go run complete_gateway_demo_v2.go "$CONFIG_PATH"