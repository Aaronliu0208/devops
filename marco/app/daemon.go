package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync/atomic"
	"syscall"
	"time"

	"casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/pkg/logger"
	"github.com/sirupsen/logrus"
)

//启动http server管理本地nginx
// 通过NginxController启停Nginx
// 通过RevisionController 管理工作目录版本

// InitLogger 初始化日志模块
func InitLogger() (func(), error) {
	c := config.C.Log
	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(int(level))
	logger.SetFormatter(c.Format)

	// 设定日志输出
	var file *os.File
	if c.Output != "" {
		switch c.Output {
		case "stdout":
			logger.SetOutput(os.Stdout)
		case "stderr":
			logger.SetOutput(os.Stderr)
		case "file":
			if name := c.File; name != "" {
				_ = os.MkdirAll(filepath.Dir(name), 0777)

				f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					return nil, err
				}
				logger.SetOutput(f)
				file = f
			}
		}
	}

	return func() {
		if file != nil {
			file.Close()
		}
	}, nil
}

// Init 应用初始化
func Init(ctx context.Context) (func(), error) {

	// 初始化日志模块
	loggerCleanFunc, err := InitLogger()
	if err != nil {
		return nil, err
	}

	return func() {
		loggerCleanFunc()
	}, nil
}

// Run app entry point
func Run(ctx context.Context) error {
	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx)
	if err != nil {
		return err
	}
EXIT:
	for {
		sig := <-sc
		fmt.Printf("接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.CompareAndSwapInt32(&state, 1, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	fmt.Printf("服务退出")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
	return nil
}
