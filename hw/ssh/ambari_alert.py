#!/usr/bin/env python
# -*- coding: utf-8 -*-
# https://github.com/apache/ambari/blob/trunk/ambari-server/docs/api/v1/alerts.md

import datetime
import requests
import json

def get_alert_items():
    auth =('admin', 'Itn28gfJTmigt+5r/kDPQg==')
    url = 'http://10.153.51.115:8088/api/v1/clusters/HDP/alerts'
    params = b"fields=*&Alert/state.in(WARNING,CRITICAL)&sortBy=Alert/state"
    response = requests.get(url, params=params, auth=auth)
    response_json = response.json()
    return response_json['items']

def get_alert_message():
    alert_message=list()
    items = get_alert_items()
    for item in items:
        alert = item['Alert']
        original_timestamp = alert['original_timestamp']
        start_time = datetime.datetime.fromtimestamp(original_timestamp / 1e3).strftime("%Y-%m-%d %H:%M:%S")
        service_name = alert['service_name']
        host_name = alert['host_name']
        label = alert['label']
        text = alert['text']
        alert_message.append('=====门户Hadoop集群=====\r\n'
                             '服务名: %s\r\n'
                             '主机名: %s\r\n'
                             '告警类型: %s\r\n'
                             '告警内容: %s\r\n'
                             '告警时间: %s\r\n'% (service_name, host_name, label, text, start_time))
    return alert_message



def to_weixin():
    corpid = 'wx5df686548b138675'
    secret = 'QGmxXQU7tQc3FSvVUYWc9jBmecEWkCiuKDqHR6B4O6c'
    token_url = 'https://qyapi.weixin.qq.com/cgi-bin/gettoken'
    playload = {'corpid': corpid, 'corpsecret': secret}
    r = requests.get(token_url, params=playload)
    access_token = r.json()['access_token']
    msg = get_alert_message()
    for message in msg:
        data = {"touser": "@all", "msgtype": "text", "agentid": 1000009, "text": {"content": message}}
        send_message_url = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + access_token
        send_message = requests.post(send_message_url, data=json.dumps(data))


def main():
    to_weixin()

if __name__=="__main__":
    main()

