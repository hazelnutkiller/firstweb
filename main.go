package main

import (
	//"firstweb/mysql"

	"firstweb/routers"
)

func main() {

	routers.Timeout()
	//----------------------------------------------------
	//實例化一個路由對象
	// router := mux.NewRouter()
	// // //註冊路由
	// routers.RegisterRoutes(router)
	// fmt.Printf("服務運行在端口:9999\n")
	// if err := http.ListenAndServe(":9999", router); err != nil {
	// 	log.Fatal(err)
	// }

	//-----------------------
	routers.Router()
	go routers.RunRouter()
	routers.Cleanup()
	//------------------------

	//mysql.Mysql()

}

// func main() {
// 	//實例化一個路由對象
// 	router := mux.NewRouter()
// 	//註冊路由
// router := mux.NewRouter()
// routers.RegisterRoutes(router)

// 	fmt.Printf("服務運行在端口:8080\n")
// 	if err := http.ListenAndServe(":8080", router); err!= nil {
// 		log.Fatal(err)
// 	}
// }
