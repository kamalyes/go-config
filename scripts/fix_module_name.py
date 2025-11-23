import os
import re

def process_file(filepath):
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # 1. 将 modulename 改为 module_name
    content = content.replace('modulename', 'module_name')
    
    # 2. 修改标签：mapstructure 和 json 用下划线，yaml 用横线
    def fix_tags(match):
        mapstructure_value = match.group(1)
        # mapstructure 和 json 使用下划线
        underscore_value = mapstructure_value.replace('-', '_')
        # yaml 使用横线
        kebab_value = mapstructure_value
        rest = match.group(2) if match.group(2) else ''
        return f'`mapstructure:"{underscore_value}" yaml:"{kebab_value}" json:"{underscore_value}"{rest}`'
    
    # 匹配现有的三标签格式
    pattern = r'`mapstructure:"([^"]+)" yaml:"[^"]+" json:"[^"]+"(\s+[^`]*)?`'
    content = re.sub(pattern, fix_tags, content)
    
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
