#!/bin/bash
# 创建管理自定义链的方法
# 本脚本主要是为了管理frpc暴露在外面的ssh登录端口。通过iptable来进行安全配置
D_PORT=6000
IPT_CHAIN=SSH_CHAIN_$D_PORT

# 清理现有Chain
# 这里要清除对应INPUT中的链引用
iptables -t filter -F $IPT_CHAIN
iptables -t filter -X $IPT_CHAIN

# 创建Chain
iptables -t filter -N $IPT_CHAIN

#引用自定义的Chain
iptables -t filter -I INPUT -p tcp --dport $D_PORT -j $IPT_CHAIN

#默认保持state
iptables -A $IPT_CHAIN -m conntrack --ctstate ESTABLISHED, RELATED -j ACCEPT

#默认阻止一切SSH
iptables -A $IPT_CHAIN -j DROP

#默认允许阿里云的远程访问
iptables -A $IPT_CHAIN -s 111.193.48.97 -p tcp --dport $D_PORT -j ACCEPT

iptables -S
iptables-save > /etc/sysconfig/iptables
