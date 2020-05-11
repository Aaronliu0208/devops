云路公司站点监控部署方案
===

## 部署拓扑

### 内蒙机房:

10.153.51.181:9090 prometheus

10.153.51.181:9093 alertmanager(使用云路自己开发的)

10.153.51.181:8500 consul，部署在docker中

10.153.51.181:9115 blackbox-neimeng 云路内蒙站点监控采集点，部署在docker中

10.153.51.181:9116 blackbox-aliyun-bj 云路阿里云北京监控采集点 使用frps技术穿透到阿里云，部署在docker中

10.153.51.181:9117 blackbox-aliyun-hk 云路阿里云香港监控采集点 使用frps技术穿透到阿里云，部署在docker中

blackbox-k8s 在k8s中运行blackbox

106.74.152.39:9000 公网frps server 内网地址:10.153.51.193

其中需要在10.153.51.181上运行`frpc visitor`。参考[blackbox](blackbox)中的启动方法，分别启动container和visitor
### 阿里云-北京(112.126.96.91):
172.17.220.217:9115 blackbox-aliyun-bj 云路阿里云北京监控采集点 使用frps技术穿透到阿里云，部署在docker中
### 阿里云-香港(47.52.254.174):
172.31.205.207:9115 blackbox-aliyun-hk 云路阿里云香港监控采集点 使用frps技术穿透到阿里云，部署在docker中


## 部署方法
采用`docker-compose`部署
```bash
docker-compose -f docker-compose.yml
```
## 部署节点说明
### 检测点
部署在不同机房的检测服务，主要使用了[prometheus_blackbox](http://yunludev3.htyunlu.com/ams/adc/blackbox_exporter)作为检测进程，可以在不同机房部署多个位置测点
### visitor
frpc中的角色，每启动一个穿透代理server就要在对应的内网中起一个visitor。该visitor负责局域网内不其他机器访问这个代理服务

