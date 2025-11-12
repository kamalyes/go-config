@echo off
setlocal EnableDelayedExpansion

REM 完整Gateway演示服务启动脚本 (Windows版本)

echo 🚀 启动完整Gateway演示服务...

REM 设置环境变量
set APP_ENV=development
set CONFIG_PATH=.\complete-gateway-config.yaml

REM 检查Go环境
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Go环境未安装，请先安装Go 1.19+
    pause
    exit /b 1
)

REM 检查配置文件
if not exist "complete-gateway-config.yaml" (
    echo ❌ 配置文件 complete-gateway-config.yaml 不存在
    echo 请确保配置文件在当前目录下
    pause
    exit /b 1
)

echo ✅ 环境检查完成

REM 显示服务信息
echo.
echo 📋 服务信息:
echo    环境: %APP_ENV%
echo    配置文件: %CONFIG_PATH%
echo    服务地址: http://localhost:8080
echo.

REM 启动服务
echo 🔥 启动服务...
go run complete_gateway_demo_v2.go "%CONFIG_PATH%"

pause