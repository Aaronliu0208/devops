openresty配置管理工具
===
参考kong开发openresty的web应用程序，通过golang 进行配置文件管理

## 用go-bindata将resty安装包打入binary中
### 安装 
```bash
go get -u github.com/jteeuwen/go-bindata/...
```
### 运行
```bash
cd resource
go-bindata -pkg resource -o resty.go openresty-1.17.8.2.tar.gz
```

### 获取tar包
```golang
data, err := Asset("openresty-1.17.8.2.tar.gz")
if err != nil {
    // Asset was not found.
}
```