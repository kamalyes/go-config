# 完整Gateway演示服务启动脚本 (PowerShell版本)

Write-Host "🚀 启动完整Gateway演示服务..." -ForegroundColor Cyan

# 设置环境变量
$env:APP_ENV = "development"
$env:CONFIG_PATH = ".\complete-gateway-config.yaml"

# 检查Go环境
try {
    $goVersion = go version 2>$null
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ Go环境检查: $goVersion" -ForegroundColor Green
    } else {
        throw "Go not found"
    }
} catch {
    Write-Host "❌ Go环境未安装，请先安装Go 1.19+" -ForegroundColor Red
    Read-Host "按任意键退出"
    exit 1
}

# 检查配置文件
if (Test-Path "complete-gateway-config.yaml") {
    Write-Host "✅ 配置文件: complete-gateway-config.yaml" -ForegroundColor Green
} else {
    Write-Host "❌ 配置文件 complete-gateway-config.yaml 不存在" -ForegroundColor Red
    Write-Host "请确保配置文件在当前目录下" -ForegroundColor Yellow
    Read-Host "按任意键退出"
    exit 1
}

Write-Host "✅ 环境检查完成" -ForegroundColor Green

# 显示服务信息
Write-Host ""
Write-Host "📋 服务信息:" -ForegroundColor Blue
Write-Host "   环境: $($env:APP_ENV)" -ForegroundColor White
Write-Host "   配置文件: $($env:CONFIG_PATH)" -ForegroundColor White
Write-Host "   服务地址: http://localhost:8080" -ForegroundColor White
Write-Host ""

# 显示可用接口
Write-Host "📋 主要接口:" -ForegroundColor Blue
Write-Host "   http://localhost:8080/         - 服务首页" -ForegroundColor White
Write-Host "   http://localhost:8080/config   - 配置信息" -ForegroundColor White
Write-Host "   http://localhost:8080/swagger/ - API文档" -ForegroundColor White
Write-Host "   http://localhost:8080/health   - 健康检查" -ForegroundColor White
Write-Host "   http://localhost:8080/metrics  - 监控指标" -ForegroundColor White
Write-Host ""

# 启动服务
Write-Host "🔥 启动服务..." -ForegroundColor Yellow

try {
    # 首先尝试编译检查
    Write-Host "检查代码编译..." -ForegroundColor Gray
    go build -v complete_gateway_demo_v2.go
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ 代码编译成功" -ForegroundColor Green
        
        # 清理编译产物
        if (Test-Path "complete_gateway_demo_v2.exe") {
            Remove-Item "complete_gateway_demo_v2.exe"
        }
        
        # 启动服务
        Write-Host ""
        Write-Host "🚀 正在启动服务，按 Ctrl+C 停止..." -ForegroundColor Cyan
        Write-Host ""
        
        go run complete_gateway_demo_v2.go $env:CONFIG_PATH
    } else {
        Write-Host "❌ 代码编译失败" -ForegroundColor Red
        Read-Host "按任意键退出"
        exit 1
    }
} catch {
    Write-Host "❌ 启动失败: $($_.Exception.Message)" -ForegroundColor Red
    Read-Host "按任意键退出"
    exit 1
}

Write-Host ""
Write-Host "服务已停止" -ForegroundColor Yellow
Read-Host "按任意键退出"