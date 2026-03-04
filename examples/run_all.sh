#!/bin/bash
# 运行所有示例程序

set -e

echo ""
echo "========================================"
echo "运行所有 Go 示例程序"
echo "========================================"
echo ""

count=0

for dir in */; do
    if [ -f "${dir}main.go" ]; then
        count=$((count + 1))
        echo "========================================"
        echo "运行示例: ${dir%/}"
        echo "========================================"
        
        cd "$dir"
        
        if go run main.go; then
            echo ""
            echo "✅ ${dir%/} 运行成功"
            echo ""
        else
            echo ""
            echo "❌ ${dir%/} 运行失败"
            echo ""
        fi
        
        cd ..
        
        # 等待一下
        sleep 0.5
    fi
done

if [ $count -eq 0 ]; then
    echo "未找到任何示例程序"
    exit 1
fi

echo ""
echo "========================================"
echo "共运行 $count 个示例程序"
echo "========================================"
