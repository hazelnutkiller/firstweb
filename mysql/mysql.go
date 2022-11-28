package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//將連線的資訊設定為常數
const (
	USERNAME = "root"
	PASSWORD = "mindy123"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "firstweb"
)

// type Create_demo struct {
// 	C_id     int
// 	playerID string
// 	currency string
// 	time     int
// }

func Mysql() {
	//連線字串拼湊出來
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	//透過 sql.Open 方法進行連線
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err)
		return
	}
	//檢查資料庫是否連線正常
	if err := db.Ping(); err != nil {
		fmt.Println("資料庫連線錯誤，原因為：", err.Error())
		return
	}
	defer db.Close()
}
