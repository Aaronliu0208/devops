云路运维管理服务器端
===


参考[go-admin](https://github.com/LyricTian/gin-admin)与[gin-admin-react](https://github.com/LyricTian/gin-admin-react)进行前后端分离的运维管理服务


## features
### 人员角色管理

### 报警事件管理
> 管理`prometheus alertmanager`传送的报警信息。实现原理是自定义一个alertmanager的webhook记录报警事件。
> 
> 管理阿里云监控发送的报警事件。实现原理是[阿里云使用报警回调](https://help.aliyun.com/document_detail/60714.html)

> 实现基本的查询功能，按字段排序，分页查询。以及简单的聚合图表分析

### 运维事件管理
记录运维对线上进行的相关配置变更操作，升级维护操作以及故障处理等操作。主要提供创建，修改，条件查询以及聚合分析等操作