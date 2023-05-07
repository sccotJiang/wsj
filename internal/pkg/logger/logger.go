//日志封装 方便后续切换log库
package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

var Logger *logrus.Logger

//LogConfig日志配置
type LogConfig struct {
	FileNamePrefix string //文件名前缀
	LogFilePath    string //储存路径
	Ip             string //机器ip
	Debug          bool   //是否开启debug日志
}

//自定义日志格式
type customFormatter struct {
	fields logrus.Fields    //自定义常驻字段
	lf     logrus.Formatter //常规formatter
}

//为了实现logrus.Formatter接口
func (f *customFormatter) Format(e *logrus.Entry) ([]byte, error) {
	//for k, v := range f.fields {
	//	e.Data[k] = v
	//}
	fmt.Println("实现接口用，TODO")
	return f.lf.Format(e)
}
func InitLogger(config LogConfig) {
	logger := logrus.New()
	if config.Debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.SetOutput(io.Discard)
	//设置日志格式
	format := customFormatter{
		fields: logrus.Fields{
			"ip": config.Ip,
		},
		lf: &logrus.JSONFormatter{
			TimestampFormat: "2006/01/02 15:04:05.000",
		},
	}
	logger.SetFormatter(&format)
	Logger = logger
}

//包装logrus的方法
func Errorf(format string, v ...interface{}) {
	Logger.Errorf(format, v...)
}

//包装Printf
func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}
