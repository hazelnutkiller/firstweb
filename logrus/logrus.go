package logrus

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var LogInstance = logrus.New()

func Logrus() gin.HandlerFunc {
	// 設置日誌記錄級別//設置日誌級別为TraceLevel，为了能看到Trace和Debug日志
	logrus.SetLevel(logrus.TraceLevel)

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		c.Next()

		header := ""
		for k, v := range c.Request.Header {
			header += k + ": " + fmt.Sprint(v) + "\r\n"
		}
		//寫進文件
		writer1 := &bytes.Buffer{}
		writer2 := os.Stdout
		writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 6379)
		if err != nil {
			log.Fatalf("create file log.txt failed: %v", err)
		}

		LogInstance.SetOutput(&lumberjack.Logger{

			MaxSize:    10, // 單文件最大容量,單位是MB
			MaxBackups: 30, // 最大保留過期文件個數
			MaxAge:     1,  // 保留過期文件的最大時間間隔,單位是天

		})
		// 同時將日誌寫入文件和控制檯
		logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))

		if c.Request.URL.RawQuery != "" {
			path += "?" + c.Request.URL.RawQuery
		}

		errorMsg := c.Keys["ErrorMsg"]

		logrus.WithField("HTTPResponse", LogResponse{
			Time:   time.Now().Format("2006/01/02 15:04:05"),
			IP:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   path,
			Header: header,
			Body:   string(body),
			Status: c.Writer.Status(),
			Error:  errorMsg,
		}).Println(c.Request.Method, path, c.Writer.Status())

	}
}

type LogResponse struct {
	Time   string
	IP     string
	Method string
	Path   string
	Header string
	Body   string
	Status int
	Params string
	Error  interface{} `json:"FailMsg,omitempty"`
}

// 	logrus.Info("just some info logging...")
// 	//設置日誌級別为TraceLevel，为了能看到Trace和Debug日志
// 	logrus.SetLevel(logrus.TraceLevel)
// 	// 設置日誌格式爲json格式
// 	//log.SetFormatter(&log.JSONFormatter{})

// 	logrus.Trace("trace msg")
// 	logrus.Debug("debug msg")
// 	logrus.Info("info msg")
// 	//logrus.Warn("warn msg")
// 	logrus.Error("error msg")
// 	//logrus.Fatal("fatal msg")
// 	//logrus.Panic("panic msg")

// 	//time：输出日志的时间
// 	//level：日志级别
// 	//msg：日志信息。

// 	//调用logrus.SetReportCaller(true)设置在输出日志中添加文件名和方法信息
// 	logrus.SetReportCaller(true)

// 	//有时候需要在输出中添加一些字段，可以通过调用logrus.WithField和logrus.WithFields实现。
// 	//logrus.WithFields接受一个logrus.Fields类型的参数，其底层实际上为map[string]interface{}：

// 	writer1 := &bytes.Buffer{}
// 	writer2 := os.Stdout
// 	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 6379)
// 	if err != nil {
// 		log.Fatalf("create file log.txt failed: %v", err)
// 	}
// 	//默认情况下，日志输出到io.Stderr。可以调用logrus.SetOutput传入一个io.Writer参数。
// 	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
// 	//传入一个io.MultiWriter， 同时将日志写到bytes.Buffer、标准输出和文件中

// 	logrus.Info("info msg")
// }

// }
// //type Hook interface {
// //Levels()方法返回感兴趣的日志级别，输出其他日志时不会触发钩子
// //Levels() []Level
// //Fire是日志输出前调用的钩子方法。
// //Fire(*Entry) error
// //}

// func init() {
// 	hookConfig := logredis.HookConfig{
// 		Host:     "localhost",
// 		Key:      "abc",
// 		Format:   "v1",
// 		App:      "firstweb",
// 		Port:     6379,
// 		Hostname: "my_app_hostname",
// 		//DB:       0, // optional
// 		TTL: 3600,
// 	}

// 	hook, err := logredis.NewHook(hookConfig)
// 	if err == nil {
// 		logrus.AddHook(hook)
// 	} else {
// 		logrus.Errorf("logredis error: %q", err)
// 	}
// }

// func Logrus() {

// 	//設置日誌級別为TraceLevel，为了能看到Trace和Debug日志
// 	logrus.SetLevel(logrus.TraceLevel)
// 	// 設置日誌格式爲json格式
// 	//log.SetFormatter(&log.JSONFormatter{})

// 	logrus.Trace("trace msg")
// 	//logrus.Debug("debug msg")
// 	logrus.Info("info msg")
// 	//logrus.Warn("warn msg")
// 	logrus.Error("error msg")
// 	//logrus.Fatal("fatal msg")
// 	//logrus.Panic("panic msg")

// 	//logrus.Panic("panic msg")
// 	//time：输出日志的时间
// 	//level：日志级别
// 	//msg：日志信息。

// 	//调用logrus.SetReportCaller(true)设置在输出日志中添加文件名和方法信息
// 	logrus.SetReportCaller(true)

// 	//有时候需要在输出中添加一些字段，可以通过调用logrus.WithField和logrus.WithFields实现。
// 	//logrus.WithFields接受一个logrus.Fields类型的参数，其底层实际上为map[string]interface{}：

// 	writer1 := &bytes.Buffer{}
// 	writer2 := os.Stdout
// 	writer3, err := os.OpenFile("/logrus"+".log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
// 	//os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 6379)
// 	//logfile, _ :=
// 	if err != nil {
// 		log.Fatalf("create file log.txt failed: %v", err)
// 	}
// 	//默认情况下，日志输出到io.Stderr。可以调用logrus.SetOutput传入一个io.Writer参数。
// 	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
// 	//传入一个io.MultiWriter， 同时将日志写到bytes.Buffer、标准输出和文件中

// 	logrus.Info("info msg")

// }
