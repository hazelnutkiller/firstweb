package model

import (
	"firstweb/config"
	"fmt"
	"log"
)

//數據庫對象
// var (
// 	db *gorm.DB
// )

//定義數據模型

type Createdemo struct {
	Id       int    `json:"id" form:"id"`
	PlayerID string `json:"playerID" form:"playerID"`
	Currency string `json:"currency" form:"currency"`
	Time     int    `json:"time" form:"time"`
}

func (create *Createdemo) CreatePlayer() int64 {
	// db.Create(&create)
	// return create
	fmt.Println(create.Currency)
	rs, err := config.SqlDB.Exec("INSERT into createdemo (id,playerID, currency, time) value (?,?,?,?)", create.Id, create.PlayerID, create.Currency, create.Time)
	if err != nil {
		log.Fatal(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return id
}

//func init() {
//fmt.Println("modle.creat.init()")
//調用它會連接app.go裡面
//config.Connect()
//拿到全局數據庫對象
//db = config.GetDB()
//往數據庫裡更新schema
//db.AutoMigrate(&Createdemo{})
//}

// func GetAllPlayers() []Createdemo {
// 	var Players []Createdemo
// 	db.Find(&Players)
// 	return Players
// }

// func GetPlayers(ID int64) (*Createdemo, *gorm.DB) {
// 	var create Createdemo
// 	db.Where("ID=?", ID).Find(&create)
// 	return &create, db
// }

// func DeletePlayers(ID int64) Createdemo {
// 	var create Createdemo
// 	db.Where("ID=?", ID).Delete(&create)
// 	return create
// }
