package dao

import "fmt"

func getTestConfig() *Config {
	username := "root"     //账号
	password := "1234.asd" //密码
	host := "127.0.0.1"    //数据库地址，可以是Ip或者域名
	port := 3306           //数据库端口
	Dbname := "marco"      //数据库名
	timeout := "10s"       //连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	return &Config{
		Debug:        true,
		DBType:       "mysql",
		MaxIdleConns: 20480,
		DSN:          dsn,
	}
}
