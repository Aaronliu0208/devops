使用prometheus+blackbox监控云路的微服务
===

## 技术要点

使用frps进行穿透保障多个监测点与prometheus正常通讯。

使用blackbox进行测点检测给prometheus提供metrics

使用consul进行站点的动态发现。通过[cli](cli)调用consul api将站点添加到服务中。

提供报警模板和prometheus配置模板

自定义[alertmanager](http://yunludev3.htyunlu.com/ams/adc/alertmanager) 添加云信和企业微信机器人报警通知方法

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

alertmanager webhook 格式定义
```json
{
  "version": "4",
  "groupKey": <string>,    // key identifying the group of alerts (e.g. to deduplicate)
  "status": "<resolved|firing>",
  "receiver": <string>,
  "groupLabels": <object>,
  "commonLabels": <object>,
  "commonAnnotations": <object>,
  "externalURL": <string>,  // backlink to the Alertmanager.
  "alerts": [
    {
      "status": "<resolved|firing>",
      "labels": <object>,
      "annotations": <object>,
      "startsAt": "<rfc3339>",
      "endsAt": "<rfc3339>",
      "generatorURL": <string> // identifies the entity that caused the alert
    },
    ...
  ]
}
```

# 企业微信机器人
https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=
```json
    {
        "msgtype": "text",
        "text": {
            "content": "广州今日天气：29度，大部分多云，降雨概率：60%",
            "mentioned_list":["wangqing","@all"],
            "mentioned_mobile_list":["13800001111","@all"]
        }
    }
```

## 功能设计
> *  支持alertmanager传过来的webhook内容
> * 支持企业微信机器人
> * 管理通知模板
## 接口设计
webhook接口
/api/tmplate/<tmplateid>/?<自定义的参数>

type: 代表adapter支持的类型（微信，企业微信，dingding)
不同的type同不同的


## 增加target的方法
参考[consul api](https://www.consul.io/api/agent/service.html)

使用curl
```bash
curl -X PUT -d '{
  "Name": "uc_monitor1",
  "Tags": [],
  "Address": "127.0.0.1",
  "Meta": {
    "address": "https://uc.ms.casicloud.com/1/subsystem/get?system_id=100&client_id=4g0ucoqrwtn92dxq&sign=ir834960bnjghze8343afajga"
  }
}' http://localhost:8500/v1/agent/service/register
```
同时在[cli](cli)中增加了命令行工具管理站点