规范与相关约定
===

## 站点(site)命名规范
使用小写的DNS域名

如:
```
www.casicloud.com
ms.casicloud.com
```

## 日志命名规范
### acess log
{site-name}-access.log
### error log
{site-name}-err.log
### 全局error log
error.log

## 配置文件约定
PID路径:  ${workspace}/logs/nginx.pid

User: nobody

logs路径使用单独的${workspace}/logs