package routers

import (
	"context"
	"firstweb/model"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//優雅的關閉伺服器
func Cleanup() {

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.

	//這邊使用os interrupt的功能
	//首先建立channel os.Signal channel buffer設為1，目的是不會block
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it

	//signal.Notify 是接收系統裡兩種kill-2,-9的方式
	//或者是下docker compose done都會接到signal.Notify
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit //把channel讀出來，下control+c就會從這邊結束，拿到channel裡的東西就結束
	log.Println("Shutdown Server ...")
	//系統層如果有處理超過5秒的也不管了，5秒一到關閉連線
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//http package裡面有個shutdown的函式
	//shutdown是丟個context進去，等所有連線結束後才會持續shutdown
	if err := model.APIServer.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown: ", err)
	}

	log.Println("Server exiting")

}
