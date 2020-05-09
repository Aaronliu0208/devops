使用prometheus+blackbox监控云路的微服务
===

## 测点的prometheus配置文件定义

```yaml
  - job_name: 'blackbox-2'
    metrics_path: /probe
    params:
      module: [http_2xx]  # Look for a HTTP 200 response.
    consul_sd_configs:
      - server: 'consul:8500'
        scheme: 'http'
        datacenter: 'dc1'
        refresh_interval: '2s'    
    relabel_configs:
      - source_labels: [__meta_consul_service_metadata_address]
        target_label: __param_target
      - source_labels: [__meta_consul_service_metadata_address]
        target_label: site
      - source_labels: [__meta_consul_service_metadata_service_name]
        target_label: instance
      - target_label: __address__
        replacement: blackbox_exporter2:9115  # The blackbox exporter's real hostname:port.
      - source_labels: [__meta_consul_service]
        regex: "consul"
        action: drop
```

## blackbox监控规则
[https://awesome-prometheus-alerts.grep.to/rules#blackbox](https://awesome-prometheus-alerts.grep.to/rules#blackbox)

## 安装python的方法
```bash
 pip install -r requirements.txt -i https://mirrors.aliyun.com/pypi/simple/
```