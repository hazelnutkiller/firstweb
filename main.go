package main

import (
	"firstweb/routers"
	//"firstweb/mysql"
)

func main() {
	//routers.Timeout()
	routers.Router()

	go routers.RunRouter()

	routers.Cleanup()
	//mysql.Mysql()

}
