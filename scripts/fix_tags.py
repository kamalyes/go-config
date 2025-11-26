import os
import re

def process_file(filepath):
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # 正则匹配：`mapstructure:"..." yaml:"..." json:"..."`
    # 替换为：`mapstructure:"..."`，并将值转为 kebab-case
    def replace_tags(match):
        mapstructure_value = match.group(1)
        # 将下划线改为横线
        kebab_value = mapstructure_value.replace('_', '-')
        # 保留 validate 等其他属性
        rest = match.group(2) if match.group(2) else ''
        return f'`mapstructure:"{kebab_value}"{rest}`'
    
    # 匹配模式：`mapstructure:"xxx" yaml:"xxx" json:"xxx" (其他可选属性)`
    pattern = r'`mapstructure:"([^"]+)"\s+yaml:"[^"]+"\s+json:"[^"]+"(\s+[^`]*)?`'
    content = re.sub(pattern, replace_tags, content)
    
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
