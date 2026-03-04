#!/usr/bin/env pwsh
# Run all example programs

$ErrorActionPreference = "Continue"

# Get all directories containing main.go
$examples = Get-ChildItem -Directory | Where-Object { Test-Path "$($_.FullName)\main.go" }

if ($examples.Count -eq 0) {
    Write-Host "No example programs found" -ForegroundColor Yellow
    exit 1
}

$separator = "=" * 80

Write-Host "Found $($examples.Count) example programs" -ForegroundColor Cyan
Write-Host ""

foreach ($example in $examples) {
    Write-Host $separator -ForegroundColor Gray
    Write-Host "Running: $($example.Name)" -ForegroundColor Green
    Write-Host $separator -ForegroundColor Gray
    
    Push-Location $example.FullName
    
    try {
        go run main.go
        $exitCode = $LASTEXITCODE
        
        if ($exitCode -eq 0) {
            Write-Host ""
            Write-Host "[OK] $($example.Name) succeeded" -ForegroundColor Green
            Write-Host ""
        } else {
            Write-Host ""
            Write-Host "[ERROR] $($example.Name) failed (exit code: $exitCode)" -ForegroundColor Red
            Write-Host ""
        }
    }
    catch {
        Write-Host ""
        Write-Host "[ERROR] $($example.Name) error: $_" -ForegroundColor Red
        Write-Host ""
    }
    finally {
        Pop-Location
    }
    
    Start-Sleep -Milliseconds 500
}

Write-Host $separator -ForegroundColor Gray
Write-Host "All examples completed" -ForegroundColor Cyan
Write-Host $separator -ForegroundColor Gray
