#!/usr/bin/env python3
import os
import re
import json

def analyze_go_file(filepath):
    """分析Go文件的结构体字段和Default函数"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # 提取主要结构体定义 (通常是模块名对应的结构体)
        struct_pattern = r'type\s+(\w+)\s+struct\s*\{([^}]+)\}'
        struct_matches = re.findall(struct_pattern, content, re.DOTALL)
        
        # 找到Default函数
        default_pattern = r'func\s+Default\(\)\s*\*?(\w+)\s*\{([^}]+)\}'
        default_match = re.search(default_pattern, content, re.DOTALL)
        
        result = {
            'file': filepath,
            'structs': [],
            'default_function': None,
            'issues': []
        }
        
        # 分析结构体
        for struct_name, struct_body in struct_matches:
            fields = []
            field_pattern = r'^\s*(\w+)\s+([^`\n]+)'
            for line in struct_body.split('\n'):
                line = line.strip()
                if line and not line.startswith('//'):
                    field_match = re.match(field_pattern, line)
                    if field_match:
                        field_name = field_match.group(1)
                        field_type = field_match.group(2).strip()
                        fields.append({
                            'name': field_name,
                            'type': field_type
                        })
            
            result['structs'].append({
                'name': struct_name,
                'fields': fields
            })
        
        # 分析Default函数
        if default_match:
            return_type = default_match.group(1)
            function_body = default_match.group(2)
            result['default_function'] = {
                'return_type': return_type,
                'body': function_body.strip()
            }
            
            # 检查Default函数是否为所有字段设置了值
            main_struct = None
            for struct in result['structs']:
                if struct['name'] == return_type:
                    main_struct = struct
                    break
            
            if main_struct:
                # 提取Default函数中设置的字段
                set_fields = set()
                field_assignment_pattern = r'(\w+):\s*'
                assignments = re.findall(field_assignment_pattern, function_body)
                set_fields.update(assignments)
                
                # 检查缺失的字段
                all_fields = {field['name'] for field in main_struct['fields']}
                missing_fields = all_fields - set_fields
                
                if missing_fields:
                    result['issues'].append(f"缺失字段默认值: {', '.join(missing_fields)}")
                
                # 检查空字符串字段
                empty_string_pattern = r'(\w+):\s*""'
                empty_strings = re.findall(empty_string_pattern, function_body)
                if empty_strings:
                    result['issues'].append(f"字段设置为空字符串: {', '.join(empty_strings)}")
        else:
            result['issues'].append("未找到Default()函数")
        
        return result
    
    except Exception as e:
        return {'file': filepath, 'error': str(e)}

def main():
    modules = [
        "access/access.go",
        "alerting/alerting.go", 
        "banner/banner.go",
        "breaker/breaker.go",
        "cache/cache.go",
        "captcha/captcha.go",
        "consul/consul.go",
        "cors/cors.go",
        "elasticsearch/elasticsearch.go",
        "email/email.go",
        "etcd/etcd.go",
        "ftp/ftp.go",
        "gateway/gateway.go",
        "grafana/grafana.go",
        "health/health.go",
        "i18n/i18n.go",
        "jaeger/jaeger.go",
        "jwt/jwt.go",
        "kafka/kafka.go",
        "logging/logging.go",
        "metrics/metrics.go",
        "middleware/middleware.go",
        "monitoring/monitoring.go",
        "pprof/pprof.go",
        "prometheus/prometheus.go",
        "queue/mqtt.go",
        "ratelimit/ratelimit.go",
        "recovery/recovery.go",
        "requestid/requestid.go",
        "restful/restful.go",
        "rpcclient/rpcclient.go",
        "rpcserver/rpcserver.go",
        "security/security.go",
        "signature/signature.go",
        "sms/aliyun.go",
        "smtp/smtp.go",
        "sts/aliyun.go",
        "swagger/swagger.go",
        "timeout/timeout.go",
        "tracing/tracing.go",
        "wsc/wsc.go",
        "youzan/youzan.go",
        "zap/zap.go"
    ]
    
    base_dir = "E:/WorkSpaces/IMProjects/go-config/pkg"
    results = []
    
    for module in modules:
        filepath = os.path.join(base_dir, module)
        if os.path.exists(filepath):
            result = analyze_go_file(filepath)
            results.append(result)
    
    # 输出分析报告
    print("# Go Config 模块 Default 函数分析报告\n")
    
    issues_count = 0
    for result in results:
        if 'error' in result:
            print(f"## {result['file']}")
            print(f"错误: {result['error']}\n")
            continue
            
        if result['issues']:
            issues_count += 1
            print(f"## {os.path.basename(result['file'])}")
            
            # 显示结构体信息
            if result['structs']:
                main_struct = result['structs'][0]  # 通常第一个是主结构体
                print(f"**结构体:** {main_struct['name']}")
                print(f"**字段数量:** {len(main_struct['fields'])}")
                print("**字段列表:**")
                for field in main_struct['fields'][:10]:  # 只显示前10个字段
                    print(f"  - {field['name']}: {field['type'][:50]}")
                if len(main_struct['fields']) > 10:
                    print(f"  - ... 还有 {len(main_struct['fields']) - 10} 个字段")
                print()
            
            # 显示问题
            print("**发现的问题:**")
            for issue in result['issues']:
                print(f"  - {issue}")
            print()
    
    print(f"\n## 总结")
    print(f"- 检查了 {len(results)} 个模块")
    print(f"- 发现 {issues_count} 个模块有问题")
    print(f"- 主要问题类型:")
    print(f"  - 缺失字段默认值")
    print(f"  - 字段设置为空字符串")
    print(f"  - 缺失Default()函数")

if __name__ == "__main__":
    main()