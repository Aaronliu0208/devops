nginx 模块设计
===

此模块主要是为了将nginx config 抽象成对象进行管理。

## Block
是nginx配置中的大括号`{}`代表一组指令集合，如`http`, `server`, `location`等指令
## Option
具体的key,value配置项
参考golang json的源代码实现，添加序列化的方法，从`struct`序列化到`map[string]interface{}`