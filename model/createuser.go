package model

import (
	"database/sql/driver"
	"encoding/json"
	"firstweb/config"
	"time"

	"gorm.io/gorm"
)

//數據庫對象
var (
	db *gorm.DB
)

//定義數據模型
type Createdemo struct {
	Id        int64          `json:"id" form:"id"`
	PlayerID  string         `json:"playerID" form:"playerID"`
	Currency  string         `json:"currency" form:"currency"`
	Time      int            `json:"time" form:"time"`
	Balance   int64          `json:"balance" form:"balance"`
	RefID     string         `json:"refID" form:"refID"`
	Deposit   string         `json:"deposit" form:"deposit"`
	Withdraw  int64          `json:"withdraw" form:"withdraw"`
	CreatedAt time.Time      `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Userform struct {
	ID        int64          `json:"Id" form:"Id"`
	PlayerID  string         `json:"playerID" form:"playerID"`
	Currency  string         `json:"currency" form:"currency"`
	Time      int            `json:"time" form:"time"`
	Balance   int64          `json:"balance" form:"balance"`
	CreatedAt time.Time      `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func init() {
	//fmt.Println("modle.creat.init()")
	//調用它會連接app.go裡面
	config.Connect()
	//拿到全局數據庫對象
	db = config.GetDB()
	//往數據庫裡更新schema
	db.AutoMigrate(&Createdemo{})
	db.AutoMigrate(&Userform{})
	db.AutoMigrate(&BetInfo1{})

}

//-------------------------小馬gorm方法------------------------------------------------
//[創建玩家]
func (create *Userform) CreatePlayer() *Userform {
	db.Create(&create)
	return create
}
func (add *Createdemo) Addplayer() *Createdemo {
	db.Create(&add)
	return add
}

//[取得玩家]----------------------------------------------------
func GetPlayer(id int64) (*Createdemo, *gorm.DB) {
	var create Createdemo
	//對象映射方式
	db.Where("ID=?", id).Find(&create)
	return &create, db
}

//[增減款項]-----------------------------------------------
func (addtrans *Createdemo) Addtrans() *Createdemo {
	db.Create(&addtrans)
	return addtrans
}

//[取得所有玩家]----------------------------------------------------
func GetAllPlayers() []Createdemo {
	var createdemos []Createdemo
	db.Find(&createdemos)
	return createdemos
}

//[軟刪除表格Createdemo]----------------------------------------------------
func DeletePlayer(id int64) Createdemo {
	var create Createdemo
	//調用where這個方法傳入id並且把玩家的實例付給create結構體對象
	db.Where("ID=?", id).Delete(&create)
	//
	return create
}

//[更新餘額]-------------------------------------------------------
func UpdataBalance(w *Userform) (err error) {

	if err = db.Model(&w).Where("player_id=?", w.PlayerID).Update("balance", w.Balance).Error; err != nil {
		return err
	}
	return nil

}

// func Updata(playerID string) (*Userform, *gorm.DB) {
// 	var update Userform
// 	//調用where這個方法傳入id並且把玩家的實例付給create結構體對象
// 	db.Where("player_id=?", update.PlayerID).Update("balance", update.Balance)
// 	//
// 	return &update, db
// }

//修改
// func (person *Userform) Update() int64 {
// 	rs, err := db.Exec("update userform set balance = ? where player_id = ?", person.Balance, person.PlayerID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	rows, err := rs.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return rows
// }

//-----------------------------------gin方法-------------------------------------------------

//[查詢玩家]

// func (p *Createdemo) GetRow() (create Createdemo, err error) {
// 	create = Createdemo{}
// 	err = config.SqlDB.QueryRow("Select id,player_id,currency,time from createdemos where id = ?", p.Id).Scan(&create.Id, &create.PlayerID, &create.Currency, &create.Time)
// 	return
// }

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

// func (create *Createdemo) GetRows() (creates []Createdemo, err error) {
// 	rows, err := config.SqlDB.Query("select id,playerID,currency,time from createdemo")
// 	for rows.Next() {
// 		create := Createdemo{}
// 		err := rows.Scan(&create.Id, &create.PlayerID, &create.Currency, &create.Time)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		creates = append(creates, create)
// 	}
// 	rows.Close()
// 	return
// }

//-------------------------betrecord---------------------------------------------------..

//定義投注記錄數據模型

type BetInfo1 struct {
	Id             int64  `json:"id" form:"id"`
	BetID          string `gorm:"betID" json:"betID"`
	OperatorID     string `gorm:"operatorID" json:"operatorID" `
	PlayerID       string `gorm:"playerID" json:"playerID"`
	WEPlayerID     string `gorm:"wEPlayerID" json:"wEPlayerID"`
	BetDateTime    int64  `gorm:"betDateTime" json:"betDateTime"`
	SettlementTime int64  `gorm:"settlementTime" json:"settlementTime"`
	BetStatus      string `gorm:"betStatus" json:"betStatus"`
	BetCode        string `gorm:"betCode" json:"betCode"`
	ValidBetAmount int64  `gorm:"validBetAmount" json:"validBetAmount"`
	GameResult     string `gorm:"gameResult" json:"gameResult"`
	//Device         string `json:"device" form:"device"`
	BetAmount     int64 `gorm:"betAmount" json:"betAmount"`
	WinlossAmount int64 `gorm:"winlossAmount" json:"winlossAmount"`
	//Category       string `json:"category" form:"category"`
	GameType    string `gorm:"gameType" json:"gameType"`
	GameRoundID string `gorm:"gameRoundID" json:"gameRoundID"`
	//IP             string `json:"ip" form:"ip"`
	//UID   string `json:"uid" form:"uid"`
	//RefID string `json:"RefID"  form:"RefID"`
	//CardResult     map[string]string `json:"cardresult"`
	//TransferType   string            `json:"transferType"`
	//TransferTime   int64             `json:"transferTime"`
	//TranAmount     int64             `json:"tranAmount"`
	//Balance        int64             `json:"balance"`
	CreatedAt time.Time      `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Data []string

type BetInfo struct {
	Data       []BetInfo1 `json:"data" form:"data"`
	TotalCount int        `json:"totalCount" form:"totalCount"`
	DataCount  int        `json:"dataCount" form:"dataCount"`
	Id         int64      `json:"id" form:"id"`
	// BetID          string            `json:"betID" form:"betID"`
	// OperatorID     string            `json:"operatorID" form:"operatorID"`
	// PlayerID       string            `json:"playerID" form:"playerID"`
	// WEPlayerID     string            `json:"wEPlayerID" form:"wEPlayerID"`
	// SettlementTime int64             `json:"settlementTime" form:"settlementTime"`
	// BetStatus      string            `json:"betStatus" form:"betStatus"`
	// BetCode        string            `json:"betCode" form:"betCode"`
	// ValidBetAmount int64             `json:"validBetAmount" form:"validBetAmount"`
	// GameResult     string            `json:"gameResult" form:"gameResult"`
	// Device         string            `json:"device" form:"device"`
	// BetAmount      int64             `json:"betAmount" form:"betAmount"`
	// WinlossAmount  int64             `json:"winlossAmount" form:"winlossAmount"`
	// Category       string            `json:"category" form:"category"`
	// GameType       string            `json:"gameType" form:"gameType"`
	// GameRoundID    string            `json:"gameRoundID" form:"gameRoundID"`
	// IP             string            `json:"ip" form:"ip"`
	// UID            string            `json:"uid" form:"uid"`
	// RefID          string            `json:"RefID"  form:"RefID"`
	// CardResult     map[string]string `json:"cardresult"`
	// TransferType   string            `json:"transferType"`
	// TransferTime   int64             `json:"transferTime"`
	// TranAmount     int64             `json:"tranAmount"`
	// Balance        int64             `json:"balance"`
}

func (t *Data) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Data) Value() (driver.Value, error) {
	return json.Marshal(t)
}

//[投注記錄新增進db]
func (betrecord *BetInfo1) BetRecord() *BetInfo1 {
	db.Create(&betrecord)
	return betrecord
}
