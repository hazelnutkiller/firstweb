package main

import (
	"firstweb/logrus"
	"firstweb/routers"
	"fmt"
)

func main() {
	logrus.Init()
	logrus.Logrus()
	routers.Router()
	fmt.Println("can run")

}
