adapter将reciver的webhook转成其他消息系统的webhook
===

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
