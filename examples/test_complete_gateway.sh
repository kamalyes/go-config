#!/bin/bash

# 完整Gateway演示服务API测试脚本

SERVER_URL="http://localhost:8080"
SUCCESS_COUNT=0
FAIL_COUNT=0

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 测试函数
test_endpoint() {
    local method=$1
    local endpoint=$2
    local description=$3
    local expected_code=${4:-200}
    
    echo -n "测试 $method $endpoint - $description ... "
    
    if [ "$method" = "GET" ]; then
        response_code=$(curl -s -o /dev/null -w "%{http_code}" "$SERVER_URL$endpoint")
    else
        response_code=$(curl -s -o /dev/null -w "%{http_code}" -X "$method" "$SERVER_URL$endpoint")
    fi
    
    if [ "$response_code" = "$expected_code" ]; then
        echo -e "${GREEN}✅ 成功 ($response_code)${NC}"
        ((SUCCESS_COUNT++))
    else
        echo -e "${RED}❌ 失败 (期望: $expected_code, 实际: $response_code)${NC}"
        ((FAIL_COUNT++))
    fi
}

# 测试JSON响应
test_json_endpoint() {
    local endpoint=$1
    local description=$2
    
    echo -n "测试 GET $endpoint - $description ... "
    
    response=$(curl -s "$SERVER_URL$endpoint")
    if echo "$response" | jq empty > /dev/null 2>&1; then
        code=$(echo "$response" | jq -r '.code // 200')
        if [ "$code" = "200" ]; then
            echo -e "${GREEN}✅ 成功 (JSON格式正确)${NC}"
            ((SUCCESS_COUNT++))
        else
            echo -e "${YELLOW}⚠️  响应码异常 ($code)${NC}"
            ((FAIL_COUNT++))
        fi
    else
        echo -e "${RED}❌ 失败 (JSON格式错误)${NC}"
        ((FAIL_COUNT++))
    fi
}

echo -e "${BLUE}🚀 开始测试完整Gateway演示服务...${NC}"
echo -e "${BLUE}服务地址: $SERVER_URL${NC}"
echo ""

# 检查服务是否运行
echo -n "检查服务是否运行 ... "
if curl -s --connect-timeout 5 "$SERVER_URL" > /dev/null 2>&1; then
    echo -e "${GREEN}✅ 服务正在运行${NC}"
else
    echo -e "${RED}❌ 服务未运行或无法连接${NC}"
    echo "请先启动服务：./start_complete_gateway.sh"
    exit 1
fi

echo ""
echo -e "${BLUE}📋 开始API测试...${NC}"

# 基础接口测试
echo ""
echo "=== 基础接口测试 ==="
test_json_endpoint "/" "服务首页"
test_json_endpoint "/config" "配置信息"
test_json_endpoint "/status" "服务状态"

# 健康检查和监控测试
echo ""
echo "=== 健康检查和监控测试 ==="
test_json_endpoint "/health" "健康检查"
test_endpoint "GET" "/metrics" "监控指标"

# API文档测试
echo ""
echo "=== API文档测试 ==="
test_endpoint "GET" "/swagger/" "Swagger UI"
test_json_endpoint "/swagger/doc.json" "API文档JSON"

# 业务接口测试
echo ""
echo "=== 业务接口测试 ==="
test_json_endpoint "/api/users" "用户列表"
test_json_endpoint "/api/users/1" "用户详情"
test_json_endpoint "/api/cache/test" "缓存测试"
test_json_endpoint "/api/db/test" "数据库测试"

# 管理接口测试
echo ""
echo "=== 管理接口测试 ==="
test_json_endpoint "/admin/config/validate" "配置验证"
test_endpoint "POST" "/admin/config/reload" "配置重载"

echo ""
echo -e "${BLUE}🔥 测试热更新功能...${NC}"

# 备份原配置
echo -n "备份原配置文件 ... "
if cp complete-gateway-config.yaml complete-gateway-config.yaml.bak; then
    echo -e "${GREEN}✅ 成功${NC}"
else
    echo -e "${RED}❌ 失败${NC}"
fi

# 修改配置（切换debug模式）
echo -n "修改配置文件 ... "
if sed -i 's/debug: true/debug: false/g' complete-gateway-config.yaml 2>/dev/null || \
   sed -i '' 's/debug: true/debug: false/g' complete-gateway-config.yaml 2>/dev/null; then
    echo -e "${GREEN}✅ 成功${NC}"
else
    echo -e "${YELLOW}⚠️  sed命令可能不支持，请手动测试热更新${NC}"
fi

# 等待配置生效
echo -n "等待配置生效 ... "
sleep 2
echo -e "${GREEN}✅ 完成${NC}"

# 验证热更新
echo -n "验证热更新效果 ... "
response=$(curl -s "$SERVER_URL/config")
if echo "$response" | grep -q '"debug":false' || echo "$response" | grep -q '"debug": false'; then
    echo -e "${GREEN}✅ 热更新成功${NC}"
    ((SUCCESS_COUNT++))
else
    echo -e "${RED}❌ 热更新失败${NC}"
    ((FAIL_COUNT++))
fi

# 恢复原配置
echo -n "恢复原配置文件 ... "
if mv complete-gateway-config.yaml.bak complete-gateway-config.yaml; then
    echo -e "${GREEN}✅ 成功${NC}"
else
    echo -e "${RED}❌ 失败${NC}"
fi

# 手动重载配置
echo -n "手动重载配置 ... "
reload_response=$(curl -s -X POST "$SERVER_URL/admin/config/reload")
if echo "$reload_response" | grep -q '"success":true' || echo "$reload_response" | grep -q '"success": true'; then
    echo -e "${GREEN}✅ 成功${NC}"
    ((SUCCESS_COUNT++))
else
    echo -e "${RED}❌ 失败${NC}"
    ((FAIL_COUNT++))
fi

# 最终结果统计
echo ""
echo -e "${BLUE}📊 测试结果统计${NC}"
echo "=================================="
echo -e "✅ 成功: ${GREEN}$SUCCESS_COUNT${NC} 项"
echo -e "❌ 失败: ${RED}$FAIL_COUNT${NC} 项"
echo -e "📊 总计: $((SUCCESS_COUNT + FAIL_COUNT)) 项"

if [ $FAIL_COUNT -eq 0 ]; then
    echo ""
    echo -e "${GREEN}🎉 所有测试通过！完整Gateway演示服务运行正常！${NC}"
    exit 0
else
    echo ""
    echo -e "${YELLOW}⚠️  部分测试失败，请检查服务配置和运行状态${NC}"
    exit 1
fi