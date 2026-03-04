@echo off
REM 运行所有示例程序

setlocal enabledelayedexpansion

echo.
echo ========================================
echo 运行所有 Go 示例程序
echo ========================================
echo.

set count=0

for /d %%d in (*) do (
    if exist "%%d\main.go" (
        set /a count+=1
        echo ========================================
        echo 运行示例: %%d
        echo ========================================
        
        pushd "%%d"
        go run main.go
        
        if !errorlevel! equ 0 (
            echo.
            echo [OK] %%d 运行成功
            echo.
        ) else (
            echo.
            echo [ERROR] %%d 运行失败 ^(退出码: !errorlevel!^)
            echo.
        )
        
        popd
        
        REM 等待一下
        timeout /t 1 /nobreak >nul
    )
)

if !count! equ 0 (
    echo 未找到任何示例程序
    exit /b 1
)

echo.
echo ========================================
echo 共运行 !count! 个示例程序
echo ========================================

endlocal
