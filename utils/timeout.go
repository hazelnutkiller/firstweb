package utils

import (
	"context"
	"firstweb/model"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Cleanup() {

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	if err := model.APIServer.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown: ", err)
	}

	log.Println("Server exiting")

}
