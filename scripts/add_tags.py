import os
import re

def process_file(filepath):
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # 匹配模式：`mapstructure:"xxx"` (后面可能有 validate 等其他属性)
    # 转换为：`mapstructure:"xxx" yaml:"xxx" json:"xxx"` (将 kebab-case 转为 snake_case)
    def add_tags(match):
        mapstructure_value = match.group(1)
        # 将 kebab-case 转为 snake_case
        yaml_json_value = mapstructure_value.replace('-', '_')
        rest = match.group(2) if match.group(2) else ''
        return f'`mapstructure:"{mapstructure_value}" yaml:"{yaml_json_value}" json:"{yaml_json_value}"{rest}`'
    
    # 匹配：`mapstructure:"xxx" (可选的其他属性)`，但排除已经有 yaml 或 json 的
    pattern = r'`mapstructure:"([^"]+)"(?!\s+yaml:)(?!\s+json:)(\s+[^`]*)?`'
    content = re.sub(pattern, add_tags, content)
    
    with open(filepath, 'w', encoding='utf-8') as f:
        f.write(content)
    print(f"处理完成: {filepath}")

# 遍历 pkg 目录下所有 .go 文件
for root, dirs, files in os.walk('pkg'):
    for file in files:
        if file.endswith('.go'):
            filepath = os.path.join(root, file)
            process_file(filepath)

print("所有文件处理完成!")
