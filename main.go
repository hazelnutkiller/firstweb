package main

import (
	"firstweb/routers"
	"firstweb/utils"
)

func main() {
	routers.Router()
	go routers.RunRouter()
	utils.Cleanup()
}
