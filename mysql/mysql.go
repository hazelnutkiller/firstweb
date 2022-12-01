package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
//go get -u github.com/gorilla/mux 路由組件
//從docker下載mysql
//docker pull xxxx(docker名)/xxxx(檔案名):latest
//下載完成運行指令
//docker run --name xxxx(檔案名) -e MYSQL_ROOT_PASSWORD=mindy1234 -p 3306:3306 -d --restart=always xxxx(docker名)/xxxx(檔案名):latest
//訪問docker shell
//docker exec -it xxxx(檔案名) bash -p
//連線mysql指令 mysql -u root -p -h 127.0.0.1
//mysql -u root -test -p -h 127.0.0.1
//show databases;use PcschoolWeb;show tables;select * from employee;

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
//   C_id     int
//   playerID string
//   currency string
//   time     int
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
