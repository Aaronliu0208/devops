web monitor cli
===
使用cli控制添加删除监控站点


## 参考
[python argparse用法总结](https://www.jianshu.com/p/fef2d215b91d)

[Python argparse 子命令](https://www.jianshu.com/p/27ce67dab97e)

[argparse --- 命令行选项、参数和子命令解析器](https://docs.python.org/zh-cn/3/library/argparse.html)

[Python @staticmethod@classmethod用法](https://blog.csdn.net/sinat_34079973/article/details/53502348)

## vscode中'unresolved import'解决方法
修改`.vscode`下的`settings.json`添加env和interprerter相关配置。参考如下
```json
{
    "python.pythonPath": "/home/shanyou/.pyenv/shims/python",
    "python.envFile": "${workspaceRoot}/.env" 
}
```

在工作目录根目录下创建`.env`文件
```bash
PYTHONPATH=./prometheus_web_monitor/cli:${PYTHONPATH}
```