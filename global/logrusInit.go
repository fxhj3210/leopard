package global

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strconv"
)

func LogrusInit() {
	//标准输出
	Log.Out = os.Stdout

	//Json格式
	Log.Formatter = &logrus.JSONFormatter{}

	//保存日志且分割
	saveLog()

	//日志等级
	level, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		Log.Panic("日志等级获取失败", err)
	}
	Log.SetLevel(logrus.Level(level))
	Log.Info("日志等级为：" + strconv.Itoa(level))

	/*
		Log.WithFields(logrus.Fields{
			"animal": "walrus",
			"size":   100,
		}).Info("A group of walrus emerges from the ocean")
	*/

	//日志是否记录调用者信息
	Log.SetReportCaller(true)

}
func saveLog() {
	path, err := os.Getwd()
	path = path + "\\log\\leopard.log"
	if err != nil {
		Log.Panic("获取当前日志保存目录失败", err)
	}
	logger := &lumberjack.Logger{
		LocalTime:  true,
		Filename:   path,
		MaxSize:    20, // megabytes
		MaxBackups: 5,
		MaxAge:     30,    //days
		Compress:   false, // disabled by default
	}
	writers := []io.Writer{
		logger,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	Log.SetOutput(fileAndStdoutWriter)
	Log.Info("保存路径为：" + path)
}
