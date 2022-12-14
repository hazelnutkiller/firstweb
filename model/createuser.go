package model

import (
	"firstweb/config"
	"fmt"
	"time"

	"gorm.io/datatypes"
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
type BetData struct {
	PlayerID       string         `json:"playerID" gorm:"type:varchar(20);notnull"`
	GameRoundID    string         `json:"gameRoundID" gorm:"type:varchar(50);notnull"`
	BetID          string         `json:"betID" gorm:"type:varchar(50);notnull"`
	OperatorID     string         `json:"operatorID" gorm:"-"`
	WePlayerID     string         `json:"wePlayerID" gorm:"-"`
	GameType       string         `json:"gameType" gorm:"type:varchar(20);notnull"`
	BetCode        string         `json:"betCode" gorm:"type:varchar(400);notnull"`
	GameResult     string         `json:"gameResult" gorm:"type:varchar(100);notnull"`
	BetStatus      string         `json:"betStatus" gorm:""`
	BetDataTime    int            `json:"betDateTime" gorm:"-"`
	SettlementTime int            `json:"settlementTime" gorm:"-"`
	ValidBetAmount int            `json:"validBetAmount" gorm:"-"`
	Device         string         `json:"device" gorm:"-"`
	BetAmount      int            `json:"betAmount" gorm:"type:bigint;notnull"`
	WinlossAmount  int            `json:"winlossAmount" gorm:"type:bigint;notnull"`
	Category       string         `json:"category" gorm:"-"`
	Ip             string         `json:"ip" gorm:"-"`
	CardResult     datatypes.JSON `json:"cardresult" gorm:"-"`
	CreatedAt      int64          `json:"-"`
	UpdatedAt      int64          `json:"-"`
}

type BetInfo struct {
	TotalCount int       `json:"totalCount"`
	DataCount  int       `json:"dataCount"`
	Data       []BetData `json:"data" gorm:"foreignKey:playerID"`
}

// func (t *Data) Scan(value interface{}) error {
// 	bytesValue, _ := value.([]byte)
// 	return json.Unmarshal(bytesValue, t)
// }

// func (t Data) Value() (driver.Value, error) {
// 	return json.Marshal(t)
// }

//[投注記錄新增進db]
//創建投注資料
func SaveBetInfo(b *BetInfo) (err error) {
	bt := db.Migrator().HasTable(&BetData{})
	if bt == false {
		db.Migrator().CreateTable(&BetData{})
	}

	for _, i := range b.Data {
		newInfo := BetData{
			PlayerID:      i.PlayerID,
			GameRoundID:   i.GameRoundID,
			BetID:         i.BetID,
			GameType:      i.GameType,
			BetCode:       i.BetCode,
			GameResult:    i.GameResult,
			BetAmount:     i.BetAmount,
			WinlossAmount: i.WinlossAmount,

			// CardResult:    i.CardResult,
		}

		if err := db.Create(&newInfo).Error; err != nil {
			fmt.Println("failed to created info", err)
		}
	}
	return nil

}
