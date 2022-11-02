package logrus

import (
	"github.com/sirupsen/logrus"
)

func Logrus() {

	//設置日誌級別为TraceLevel，为了能看到Trace和Debug日志
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	//logrus.Fatal会导致程序退出，下面的logrus.Panic不会执行到
	//logrus.Fatal("fatal msg")
	//logrus.Fatal会导致程序退出，下面的logrus.Panic不会执行到
	logrus.Panic("panic msg")
	//time：输出日志的时间；
	//level：日志级别；
	//msg：日志信息。
}
