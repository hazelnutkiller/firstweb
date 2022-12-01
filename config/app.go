package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

//整個全局控制數據庫對象的文件

func Config() {
	//要讀取的檔名
	viper.SetConfigName("app")
	//要讀取的附檔名
	viper.SetConfigType("yaml")
	//要讀取的路徑
	viper.AddConfigPath("./config")
	//設定參數預設值
	viper.SetDefault("application.port", 9999)
	//readinconfig 讀取設定檔
	err := viper.ReadInConfig()
	if err != nil {
		panic("讀取設定檔出現錯誤，原因為：" + err.Error())
	}
	fmt.Println("application port = " + viper.GetString("application.port"))
}

//全局數據庫對象
// var (
// 	db *gorm.DB
// )

// func Connect() {
// 	//通過connect連接數據庫
// 	//用戶名：口令ip地址數據庫名
// 	dsn := "root:mindy123@tcp(127.0.0.1)/firstweb?charset=utf8&parseTime=True&loc=Local"
// 	_db, _err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if _err != nil {
// 		panic(_err)
// 	}
// 	db = _db
// }

// //把模版級的數據對象返回給要使用他的函數

// func GetDB() *gorm.DB {
// 	return db
// }

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:mindy123@tcp(127.0.0.1)/firstweb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
