package main

import (
	"firstweb/routers"
)

func main() {
	routers.Router()
	go routers.RunRouter()
	routers.Cleanup()
}
