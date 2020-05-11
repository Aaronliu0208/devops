不同地区对应不同监测点的相关配置
===
检测点通过自定义的`blackbox` `hub.htres.cn/pub/blackbox-exporter`运行在不同地方。通过`frpc`将本地的检测端口映射到`prometheus`对应的监控群集中

开发4个脚本运行container和visitor

启动container
```bash
./start_container.sh -n <blackbox name>
```

停止container
```bash
./stop_container.sh -n <blackbox name>
```

启动visitor
```bash
./start_visitor.sh -n <blackbox name> -p <bind local port>
```

停止visitor
```bash
./stop_visitor.sh -n <blackbox name>
```

bash开发相关参考:
[getopts_tutorial](https://wiki.bash-hackers.org/howto/getopts_tutorial)

[Bash getopts usage template](https://gist.github.com/magnetikonline/0e44ab972a7efa3ac138)