#!/bin/bash
# 限制内蒙古内网IP访问ssh端口的方法
D_PORT=22
ACCESS_IP=10.153.51.193
CHAIN=HW_YUNLU

iptables -t filter -F $CHAIN
iptables -t filter -X $CHAIN

# 创建Chain
iptables -t filter -N $CHAIN

#引用自定义的Chain
iptables -t filter -I INPUT -p tcp --dport $D_PORT -j $CHAIN

#默认保持state
iptables -A $CHAIN -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
#默认允许特定的IP的远程访问
iptables -A $CHAIN -s $ACCESS_IP -p tcp --dport $D_PORT -j ACCEPT
#默认阻止一切SSH
iptables -A $CHAIN -j DROP

iptables -S
#iptables-save > /etc/sysconfig/iptables