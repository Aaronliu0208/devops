snapshot
===

使用[gogs](https://gogs.io)作为snapshot的远程git服务器，通过sshkey进行访问限制。
一个cluster对应一个snapshot的repo。 每创建一个snapshot就创建一个commit. 

## commit约定
描述: snapshot: <Snapinfo:ID>

Author: 配置文件中约定好的

Description: <Snapinfo:Description>

测试AccessToken: ea4ede54549411cb1a73e07248cc143265b4496d

