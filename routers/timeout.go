package routers

import (
	"context"
	"time"
)

// func Timeout() http.Client {
// 	return http.Client{
// 		Timeout: time.Duration(5) * time.Millisecond, //超时时间
// 		Transport: &http.Transport{
// 			MaxIdleConnsPerHost:   5,   //单个路由最大空闲连接数
// 			MaxConnsPerHost:       100, //单个路由最大连接数
// 			IdleConnTimeout:       90 * time.Second,
// 			TLSHandshakeTimeout:   10 * time.Second,
// 			ExpectContinueTimeout: 1 * time.Second,
// 		},
// 	}
// }

func (c *conn) readRequest(ctx context.Context) (w *response, err error) {

	if d := c.server.ReadTimeout; d != 0 {
		c.rwc.SetReadDeadline(time.Now().Add(d))
	}
	if d := c.server.WriteTimeout; d != 0 {
		defer func() {
			c.rwc.SetWriteDeadline(time.Now().Add(d))
		}()
	}

}
