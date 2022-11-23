package logrus

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	//now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := "log.log"
	//日誌檔案
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	//寫入檔案
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//例項化
	logger := logrus.New()

	//設定輸出
	logger.Out = src

	//設定日誌級別
	logger.SetLevel(logrus.DebugLevel)

	//設定日誌格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func Logrus() gin.HandlerFunc {
	logger := Logger()
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// Complete log
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		c.Next()

		header := ""
		for k, v := range c.Request.Header {
			header += k + ": " + fmt.Sprint(v) + "\r\n"
		}

		if c.Request.URL.RawQuery != "" {
			path += "?" + c.Request.URL.RawQuery
		}
		// 開始時間
		startTime := time.Now()
		// 結束時間
		endTime := time.Now()
		// 執行時間
		latencyTime := endTime.Sub(startTime)

		errorMsg := c.Keys["ErrorMsg"]

		logger.WithField("HTTPResponse", LogResponse{
			Time:   latencyTime.String(),
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
