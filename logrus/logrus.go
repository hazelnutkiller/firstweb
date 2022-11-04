package logrus

import (
	"bytes"
	"io"
	"log"
	"os"

	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

//type Hook interface {
//Levels()方法返回感兴趣的日志级别，输出其他日志时不会触发钩子
//Levels() []Level
//Fire是日志输出前调用的钩子方法。
//Fire(*Entry) error
//}

func Init() {
	hookConfig := logredis.HookConfig{
		Host:     "localhost",
		Key:      "mykey",
		Format:   "v0",
		App:      "aweosome",
		Hostname: "localhost",
		TTL:      3600,
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}
}

func Logrus() {
	logrus.Info("just some info logging...")
	//設置日誌級別为TraceLevel，为了能看到Trace和Debug日志
	logrus.SetLevel(logrus.TraceLevel)
	// 設置日誌格式爲json格式
	//log.SetFormatter(&log.JSONFormatter{})

	//logrus.Trace("trace msg")
	//logrus.Debug("debug msg")
	logrus.Info("info msg")
	//logrus.Warn("warn msg")
	//logrus.Error("error msg")
	//logrus.Fatal("fatal msg")
	//logrus.Panic("panic msg")

	//time：输出日志的时间
	//level：日志级别
	//msg：日志信息。

	//调用logrus.SetReportCaller(true)设置在输出日志中添加文件名和方法信息
	logrus.SetReportCaller(true)

	//有时候需要在输出中添加一些字段，可以通过调用logrus.WithField和logrus.WithFields实现。
	//logrus.WithFields接受一个logrus.Fields类型的参数，其底层实际上为map[string]interface{}：

	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 6379)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	//默认情况下，日志输出到io.Stderr。可以调用logrus.SetOutput传入一个io.Writer参数。
	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	//传入一个io.MultiWriter， 同时将日志写到bytes.Buffer、标准输出和文件中

	logrus.Info("info msg")

	logrus.Info("This will only be sent to Redis")

}
