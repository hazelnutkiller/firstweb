package main

import (
	//"firstweb/mysql"
	"firstweb/routers"
)

func main() {
	//routers.Timeout()
	routers.Router()

	go routers.RunRouter()

	routers.Cleanup()
	//mysql.Mysql()

}
