# 批量修复 go-config 中的标签格式
# mapstructure 和 yaml: 下划线 → 短横线 (module_name → module-name)
# json: 下划线 → 小驼峰 (module_name → moduleName)

function Convert-ToCamelCase {
    param([string]$str)
    
    # 将 snake_case 转换为 camelCase
    $parts = $str -split '_'
    $result = $parts[0].ToLower()
    for ($i = 1; $i -lt $parts.Length; $i++) {
        $result += $parts[$i].Substring(0,1).ToUpper() + $parts[$i].Substring(1).ToLower()
    }
    return $result
}

function Convert-ToKebabCase {
    param([string]$str)
    
    # 将 snake_case 转换为 kebab-case
    return $str -replace '_', '-'
}

# 获取所有需要修改的文件
$files = Get-ChildItem -Path "pkg" -Filter "*.go" -Recurse -File | 
    Where-Object { $_.Name -notlike "*_test.go" } |
    ForEach-Object { 
        if (Select-String -Path $_.FullName -Pattern 'mapstructure:"[a-z_]+_[a-z_]+"' -Quiet) { 
            $_.FullName 
        } 
    }

Write-Host "找到 $($files.Count) 个需要修改的文件" -ForegroundColor Green

$totalReplacements = 0

foreach ($file in $files) {
    Write-Host "`n处理文件: $file" -ForegroundColor Cyan
    
    $content = Get-Content $file -Raw -Encoding UTF8
    $originalContent = $content
    $fileReplacements = 0
    
    # 正则表达式匹配模式
    # 匹配: `mapstructure:"xxx_yyy" yaml:"xxx-yyy" json:"xxx_yyy"`
    $pattern = '`mapstructure:"([a-z_]+)" yaml:"([a-z-]+)" json:"([a-z_]+)"'
    
    $matches = [regex]::Matches($content, $pattern)
    
    foreach ($match in $matches) {
        $mapstructureValue = $match.Groups[1].Value
        $yamlValue = $match.Groups[2].Value
        $jsonValue = $match.Groups[3].Value
        
        # 检查是否包含下划线
        if ($mapstructureValue -match '_') {
            $newMapstructure = Convert-ToKebabCase $mapstructureValue
            $newYaml = Convert-ToKebabCase $mapstructureValue
            $newJson = Convert-ToCamelCase $mapstructureValue
            
            $oldTag = "`mapstructure:`"$mapstructureValue`" yaml:`"$yamlValue`" json:`"$jsonValue`""
            $newTag = "`mapstructure:`"$newMapstructure`" yaml:`"$newYaml`" json:`"$newJson`""
            
            if ($content -match [regex]::Escape($oldTag)) {
                $content = $content -replace [regex]::Escape($oldTag), $newTag
                $fileReplacements++
                Write-Host "  替换: $mapstructureValue → mapstructure:$newMapstructure yaml:$newYaml json:$newJson" -ForegroundColor Yellow
            }
        }
    }
    
    # 如果有修改，保存文件
    if ($content -ne $originalContent) {
        Set-Content -Path $file -Value $content -Encoding UTF8 -NoNewline
        Write-Host "  ✓ 完成 $fileReplacements 处替换" -ForegroundColor Green
        $totalReplacements += $fileReplacements
    } else {
        Write-Host "  - 无需修改" -ForegroundColor Gray
    }
}

Write-Host "`n========================================" -ForegroundColor Green
Write-Host "总计修改: $totalReplacements 处" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Green
