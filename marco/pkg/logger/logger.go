// Package logger wrap logrus
package logger

import (
	"context"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var hostName = os.Getenv("HOSTNAME")

//Option log option
type Option struct {
	HostName string
	WithFunc bool
}

// Get logger by name and option
func Get(name string, option *Option) *Logger {
	opt := Option{
		HostName: hostName,
		WithFunc: true,
	}
	if option != nil {
		opt = *option
	}
	return &Logger{name: name, option: opt}
}

//Logger logger wrap logrus
type Logger struct {
	name   string
	option Option
}

// WithContext creates an entry from the standard logger and adds a context to it.
func (l Logger) WithContext(ctx context.Context) *logrus.Entry {
	return l.entry().WithContext(ctx)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (l Logger) WithField(key, value string) *logrus.Entry {
	return l.entry().WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (l Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.entry().WithFields(fields)
}

//WithError  AddHook adds a hook to the standard logger hooks.
func (l Logger) WithError(err error) *logrus.Entry {
	return l.entry().WithError(err)
}

// WithTime creats an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (l Logger) WithTime(time time.Time) *logrus.Entry {
	return l.entry().WithTime(time)
}

// Trace logs a message at level Trace on the standard logger.
func (l Logger) Trace(args ...interface{}) {
	l.entry().Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func (l Logger) Debug(args ...interface{}) {
	l.entry().Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func (l Logger) Print(args ...interface{}) {
	l.entry().Print(args...)
}

// Info logs a message at level Info on the standard logger.
func (l Logger) Info(args ...interface{}) {
	l.entry().Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l Logger) Warn(args ...interface{}) {
	l.entry().Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func (l Logger) Warning(args ...interface{}) {
	l.entry().Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func (l Logger) Error(args ...interface{}) {
	l.entry().Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func (l Logger) Panic(args ...interface{}) {
	l.entry().Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l Logger) Fatal(args ...interface{}) {
	l.entry().Fatal(args...)
}

// Traceln logs a message at level Trace on the standard logger.
func (l Logger) Traceln(args ...interface{}) {
	l.entry().Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func (l Logger) Debugln(args ...interface{}) {
	l.entry().Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func (l Logger) Println(args ...interface{}) {
	l.entry().Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func (l Logger) Infoln(args ...interface{}) {
	l.entry().Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func (l Logger) Warnln(args ...interface{}) {
	l.entry().Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func (l Logger) Warningln(args ...interface{}) {
	l.entry().Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func (l Logger) Errorln(args ...interface{}) {
	l.entry().Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func (l Logger) Panicln(args ...interface{}) {
	l.entry().Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l Logger) Fatalln(args ...interface{}) {
	l.entry().Fatalln(args...)
}

// Tracef logs a message at level Trace on the standard logger.
func (l Logger) Tracef(format string, args ...interface{}) {
	l.entry().Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (l Logger) Debugf(format string, args ...interface{}) {
	l.entry().Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func (l Logger) Printf(format string, args ...interface{}) {
	l.entry().Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func (l Logger) Infof(format string, args ...interface{}) {
	l.entry().Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l Logger) Warnf(format string, args ...interface{}) {
	l.entry().Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func (l Logger) Warningf(format string, args ...interface{}) {
	l.entry().Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l Logger) Errorf(format string, args ...interface{}) {
	l.entry().Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func (l Logger) Panicf(format string, args ...interface{}) {
	l.entry().Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l Logger) Fatalf(format string, args ...interface{}) {
	l.entry().Fatalf(format, args...)
}

func (l Logger) entry() *logrus.Entry {
	entry := logrus.WithField("hostname", l.option.HostName).WithField("name", l.name)
	if l.option.WithFunc {
		entry = entry.WithField("func", getFunc(3))
	}

	return entry
}

func getFunc(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}

	f := runtime.FuncForPC(pc)
	_, line := f.FileLine(pc)
	return f.Name() + ":" + strconv.Itoa(line)
}
