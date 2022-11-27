package main

import (
	"firstweb/routers"
)

func main() {
	routers.Timeout()
	routers.Router()

	go routers.RunRouter()

	routers.Cleanup()
}
