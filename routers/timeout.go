package routers

import (
	"net/http"
	"time"
)

func Timeout() http.Client {
	return http.Client{
		Timeout: time.Duration(5) * time.Millisecond, //超时时间
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   5,   //单个路由最大空闲连接数
			MaxConnsPerHost:       100, //单个路由最大连接数
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}
