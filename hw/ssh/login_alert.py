#!/usr/bin/env python
# -*- coding: utf-8 -*-
import datetime
import requests
import json
import sys

def to_weixin():
    corpid = 'wx5df686548b138675'
    secret = '2iH_YXtpHOj9A2ct69oOqPgdT7H8hTvrT0US_stMcyU'
    token_url = 'https://qyapi.weixin.qq.com/cgi-bin/gettoken'
    playload = {'corpid': corpid, 'corpsecret': secret}
    r = requests.get(token_url, params=playload)
    access_token = r.json()['access_token']
    msg = sys.stdin.read()
    data = {"toparty": "4", "msgtype": "text", "agentid": 1000008, "text": {"content": msg}}
    send_message_url = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + access_token
    send_message = requests.post(send_message_url, data=json.dumps(data))

def main():
    now_time = datetime.datetime.now()
    day = now_time.strftime('%Y-%m-%d')
    start_time = datetime.datetime.strptime(day + ' 08:30:00', '%Y-%m-%d %H:%M:%S')
    end_time = datetime.datetime.strptime(day + ' 18:00:00', '%Y-%m-%d %H:%M:%S')
    if now_time < start_time or now_time > end_time:
        to_weixin()
    else:
        pass
    #now_hour = int(datetime.datetime.now().strftime('%H'))
    #if now_hour < 8 or now_hour > 18:
    #    to_weixin()
    #else:
    #    pass

if __name__=="__main__":
    main()

