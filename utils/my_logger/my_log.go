package my_logger

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego/logs"
)

var logger *logs.BeeLogger

func Setup() {
	logger = logs.NewLogger()
	//logger.SetLevel(7)
	config := fmt.Sprintf(`{"filename":"runtime/%s.log","daily":true,"maxdays":10}`, "bee"+time.Now().Format("2006-01-02"))
	logger.SetLogger(logs.AdapterFile, config)
	logger.SetLogFuncCallDepth(2)
}

func Debug(format string, v ...interface{}) {
	logger.Debug(format, v)
}

func Info(format string, v ...interface{}) {
	logger.Info(format, v)
}

func Warn(format string, v ...interface{}) {
	logger.Warn(format, v)
}

func Error(format string, v ...interface{}) {
	logger.Error(format, v)
}

func Fatal(format string, v ...interface{}) {
	logger.Emergency(format, v)
	log.Fatalf(format, v)
}
