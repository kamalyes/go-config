#!/usr/bin/env python3
"""
修复Go项目中yaml标签，将kebab-case格式改为snake_case
只修改yaml标签部分，不动mapstructure和json标签
"""

import os
import re
import glob

def fix_yaml_tags_in_file(file_path):
    """修复单个文件中的yaml标签"""
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        
        # 查找并替换yaml标签中的连字符
        # 匹配模式：yaml:"xxx-yyy"
        pattern = r'yaml:"([^"]*-[^"]*)"'
        
        def replace_hyphens(match):
            yaml_value = match.group(1)
            # 将连字符替换为下划线
            snake_case = yaml_value.replace('-', '_')
            return f'yaml:"{snake_case}"'
        
        content = re.sub(pattern, replace_hyphens, content)
        
        # 如果内容发生了变化，写回文件
        if content != original_content:
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write(content)
            print(f"✓ Fixed: {file_path}")
            return True
        else:
            return False
            
    except Exception as e:
        print(f"✗ Error processing {file_path}: {e}")
        return False

def main():
    """主函数"""
    # 项目根目录
    project_root = r"E:\WorkSpaces\GoProjects\engine-im-service\go-config"
    
    # 查找所有Go文件
    go_files = []
    for root, dirs, files in os.walk(project_root):
        for file in files:
            if file.endswith('.go'):
                go_files.append(os.path.join(root, file))
    
    print(f"Found {len(go_files)} Go files")
    
    fixed_count = 0
    for file_path in go_files:
        if fix_yaml_tags_in_file(file_path):
            fixed_count += 1
    
    print(f"\n总结: 修复了 {fixed_count} 个文件")

if __name__ == "__main__":
    main()