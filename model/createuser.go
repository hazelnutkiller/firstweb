package model

import (
	"firstweb/config"

	"log"

	"gorm.io/gorm"
)

//數據庫對象
var (
	db *gorm.DB
)

//定義數據模型
type Createdemo struct {
	Id       int64  `json:"id" form:"id"`
	PlayerID string `json:"playerID" form:"playerID"`
	Currency string `json:"currency" form:"currency"`
	Time     int    `json:"time" form:"time"`
}

func init() {
	//fmt.Println("modle.creat.init()")
	//調用它會連接app.go裡面
	config.Connect()
	//拿到全局數據庫對象
	db = config.GetDB()
	//往數據庫裡更新schema
	db.AutoMigrate(&Createdemo{})
}

func (create *Createdemo) CreatePlayer() *Createdemo {
	db.Create(&create)
	return create
}

//[查詢玩家]

func (p *Createdemo) GetRow() (create Createdemo, err error) {
	create = Createdemo{}
	err = config.SqlDB.QueryRow("Select id,player_id,currency,time from createdemos where id = ?", p.Id).Scan(&create.Id, &create.PlayerID, &create.Currency, &create.Time)
	return
}

//[創建玩家]
// func (create *Createdemo) CreatePlayer() int64 {

// 	rs, err := config.SqlDB.Exec("INSERT into createdemo (id, playerID, currency, time) value (?,?,?,?)", create.Id, create.PlayerID, create.Currency, create.Time)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	id, err := rs.LastInsertId()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return id
// }

//[查询所有记录]

func (create *Createdemo) GetRows() (creates []Createdemo, err error) {
	rows, err := config.SqlDB.Query("select id,playerID,currency,time from createdemo")
	for rows.Next() {
		create := Createdemo{}
		err := rows.Scan(&create.Id, &create.PlayerID, &create.Currency, &create.Time)
		if err != nil {
			log.Fatal(err)
		}
		creates = append(creates, create)
	}
	rows.Close()
	return
}

//-------------------------http方法------------------------------------------------

//func (create *Createdemo) CreatePlayer() *Createdemo {
// db.Create(&create)
// return create
//}

// func GetAllPlayers() []Createdemo {
// 	var Players []Createdemo
// 	db.Find(&Players)
// 	return Players
// }

// func DeletePlayers(ID int64) Createdemo {
// 	var create Createdemo
// 	db.Where("ID=?", ID).Delete(&create)
// 	return create
// }

func GetPlayer(id int64) (*Createdemo, *gorm.DB) {
	var create Createdemo
	//對象映射方式
	db.Where("ID=?", id).Find(&create)
	return &create, db
}

func GetAllPlayers() []Createdemo {
	var createdemos []Createdemo
	db.Find(&createdemos)
	return createdemos
}
