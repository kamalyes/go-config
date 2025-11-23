#!/usr/bin/env python3
import re
import os
import glob
from collections import defaultdict

def is_snake_case(tag):
    """检查是否符合snake_case规范"""
    # snake_case规则：全小写，用下划线分隔，不能有连字符或大写字母
    if not tag:
        return True  # 空标签认为是合法的
    
    # 特殊情况处理：包含逗号的标签（如 "distributed,omitempty"）
    if ',' in tag:
        # 分割并检查主要部分
        main_tag = tag.split(',')[0].strip()
        return is_snake_case(main_tag)
    
    # 检查是否包含大写字母或连字符
    if re.search(r'[A-Z-]', tag):
        return False
    
    # 检查是否全是小写字母、数字和下划线
    if re.match(r'^[a-z0-9_]+$', tag):
        return True
    
    return False

def categorize_violation(tag):
    """分类违反类型"""
    if ',' in tag:
        main_tag = tag.split(',')[0].strip()
        return categorize_violation(main_tag)
    
    if '-' in tag:
        return "使用连字符"
    elif re.search(r'[A-Z]', tag):
        return "包含大写字母"
    elif re.search(r'[^a-z0-9_]', tag):
        return "包含特殊字符"
    else:
        return "其他"

def suggest_snake_case(tag):
    """建议snake_case格式"""
    if ',' in tag:
        parts = tag.split(',')
        main_tag = parts[0].strip()
        other_parts = ','.join(parts[1:])
        suggested_main = suggest_snake_case(main_tag)
        return suggested_main + (',' + other_parts if other_parts else '')
    
    # 将连字符替换为下划线，转为小写
    suggested = tag.replace('-', '_').lower()
    
    # 处理驼峰命名法，在大写字母前加下划线
    suggested = re.sub(r'([a-z])([A-Z])', r'\1_\2', tag).lower()
    suggested = suggested.replace('-', '_')
    
    return suggested

def analyze_yaml_tags():
    """分析项目中的yaml标签"""
    yaml_pattern = re.compile(r'yaml:"([^"]*)"')
    
    all_tags = []
    violations = []
    violation_types = defaultdict(int)
    file_violations = defaultdict(list)
    
    print("🔍 开始搜索.go文件...")
    
    # 递归查找所有.go文件
    go_files = list(glob.glob('**/*.go', recursive=True))
    print(f"找到 {len(go_files)} 个.go文件")
    
    for go_file in go_files:
        if 'USAGE.md' in go_file:  # 排除USAGE.md
            continue
            
        try:
            with open(go_file, 'r', encoding='utf-8', errors='ignore') as f:
                lines = f.readlines()
                
            for line_num, line in enumerate(lines, 1):
                matches = yaml_pattern.findall(line)
                for match in matches:
                    all_tags.append(match)
                    
                    if not is_snake_case(match):
                        violation_type = categorize_violation(match)
                        violation_types[violation_type] += 1
                        
                        violation_info = {
                            'file': go_file,
                            'line': line_num,
                            'tag': match,
                            'suggested': suggest_snake_case(match),
                            'type': violation_type,
                            'context': line.strip()
                        }
                        violations.append(violation_info)
                        file_violations[go_file].append(violation_info)
                        
        except Exception as e:
            print(f"Error processing {go_file}: {e}")
    
    # 统计结果
    total_tags = len(all_tags)
    total_violations = len(violations)
    compliance_rate = ((total_tags - total_violations) / total_tags * 100) if total_tags > 0 else 100
    
    print("🔍 Go项目YAML标签snake_case规范分析报告")
    print("=" * 60)
    print(f"📊 总体统计：")
    print(f"   • 总YAML标签数量: {total_tags}")
    print(f"   • 不符合snake_case规范: {total_violations}")
    print(f"   • 符合规范: {total_tags - total_violations}")
    print(f"   • 合规率: {compliance_rate:.1f}%")
    print()
    
    if violation_types:
        print("📋 违规类型统计：")
        for violation_type, count in sorted(violation_types.items(), key=lambda x: x[1], reverse=True):
            percentage = (count / total_violations * 100) if total_violations > 0 else 0
            print(f"   • {violation_type}: {count} ({percentage:.1f}%)")
        print()
    
    # 按文件分组显示违规详情
    if violations:
        print("📝 详细违规情况：")
        print("-" * 60)
        
        for file_path, file_violations_list in sorted(file_violations.items()):
            print(f"\n📁 {file_path} ({len(file_violations_list)} 个问题)")
            
            for i, violation in enumerate(file_violations_list[:10], 1):  # 每个文件最多显示10个
                print(f"   {i:2d}. 行 {violation['line']:3d}: yaml:\"{violation['tag']}\"")
                print(f"       违规类型: {violation['type']}")
                print(f"       建议修改: yaml:\"{violation['suggested']}\"")
                if i < len(file_violations_list):
                    print()
            
            if len(file_violations_list) > 10:
                print(f"       ... 还有 {len(file_violations_list) - 10} 个问题")
            print()
    
    # 输出汇总建议
    if violations:
        print("💡 修改建议：")
        print("-" * 60)
        unique_violations = {}
        for v in violations:
            key = (v['tag'], v['suggested'])
            if key not in unique_violations:
                unique_violations[key] = []
            unique_violations[key].append(v['file'])
        
        for (original, suggested), files in sorted(unique_violations.items()):
            if original != suggested:
                print(f"   • \"{original}\" → \"{suggested}\" (出现在 {len(files)} 个文件中)")
    
    return {
        'total_tags': total_tags,
        'violations': total_violations,
        'compliance_rate': compliance_rate,
        'violation_types': dict(violation_types),
        'violations_detail': violations
    }

if __name__ == "__main__":
    results = analyze_yaml_tags()