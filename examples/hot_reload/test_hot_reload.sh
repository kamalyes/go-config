#!/bin/bash

# 测试热更新功能的脚本 - 循环修改配置文件

CONFIG_FILE="../resources/gateway-xl-dev.yaml"

# 配置参数
LOOP_COUNT=20       # 循环次数
SLEEP_INTERVAL=3    # 每次修改间隔（秒）

echo "🧪 开始测试热更新功能..."
echo "📝 将循环修改配置文件 $LOOP_COUNT 次，每次间隔 $SLEEP_INTERVAL 秒"
echo ""

# 循环修改配置
for ((i=1; i<=LOOP_COUNT; i++)); do
    # 第一次不等待，后续每次等待
    if [ $i -gt 1 ]; then
        sleep $SLEEP_INTERVAL
    fi
    
    # 动态生成测试数据
    NAME="Hot Reload Test v$i"
    VERSION="v3.$i.0"
    DEBUG=$( [ $((i % 2)) -eq 0 ] && echo "false" || echo "true" )
    PORT=$((8000 + i * 100))
    
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "🔄 第 $i 次修改配置..."
    echo "✏️  修改内容："
    echo "   - name: $NAME"
    echo "   - version: $VERSION"
    echo "   - debug: $DEBUG"
    echo "   - port: $PORT"
    
    # 读取配置文件
    content=$(cat "$CONFIG_FILE")
    
    # 修改多个字段（注意：只修改 http 部分的 port）
    content=$(echo "$content" | sed "s/name: .*/name: $NAME/")
    content=$(echo "$content" | sed "s/version: .*/version: $VERSION/")
    content=$(echo "$content" | sed "s/debug: .*/debug: $DEBUG/")
    content=$(echo "$content" | sed "/^http:/,/^[^ ]/ s/  port: [0-9]*/  port: $PORT/")
    
    # 写入文件
    echo "$content" > "$CONFIG_FILE"
    
    echo "✅ 配置文件已修改"
    echo ""
done

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ 测试完成！共修改 $LOOP_COUNT 次配置"
