// Package log wrap logrus
package log

import (
	"flag"
	"fmt"
	"io"

	"k8s.io/klog/v2"
)

var (
	// use level 2 as debug output
	DEBUG     klog.Level = 2
	klogFlags flag.FlagSet
)

func init() {
	klog.InitFlags(&klogFlags)
	klogFlags.Set("alsologtostderr", "true")
	klogFlags.Set("logtostderr", "false")
}

func SetLevel(level int) {
	klogFlags.Set("v", fmt.Sprintf("%d", level))
}

func SetOutputFile(file string) {
	klogFlags.Set("log_file", file)
}

func SetOutput(output io.Writer) {
	klog.SetOutput(output)
}

func Debug(args ...interface{}) {
	// klog 目前没有debug 级别的api 暂时用 info代替
	klog.V(DEBUG).Info(args...)
}

func Debugf(format string, args ...interface{}) {
	klog.V(DEBUG).Infof(format, args...)
}

func Debugln(args ...interface{}) {
	klog.V(DEBUG).Infoln(args...)
}

func Info(args ...interface{}) {
	// klog 目前没有debug 级别的api 暂时用 info代替
	klog.Info(args...)
}

func Infof(format string, args ...interface{}) {
	klog.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	klog.Infoln(args...)
}

func Error(args ...interface{}) {
	// klog 目前没有debug 级别的api 暂时用 info代替
	klog.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	klog.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	klog.Errorln(args...)
}

func Fatal(args ...interface{}) {
	// klog 目前没有debug 级别的api 暂时用 info代替
	klog.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	klog.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	klog.Fatalln(args...)
}
