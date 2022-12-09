package controllers

//controllers 控制器處理具體業務邏輯
//接收路由響應 讀取數據庫的地方 連接到數據庫模型

import (
	"encoding/json"
	"firstweb/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//生成標準返回結構類型

type ResponseResult struct {
	Result string `json:"result"`
}

//設置共同頭部信息

func setSameHeader(w http.ResponseWriter) {
	w.Header().Set("content-Type", "application/x-www-form-urlencoded")
}

//定義各個處理函數
// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	book := &models.Book{}
// 	//拿到用戶請求流調用實用類進行解析
// 	utils.ParseRequestBody(r, book)//book傳入後通知對象進行解析
// 	//把請求流中的json各式各樣的性質付給到結構體當中
// 	fmt.Println(book)
// //成功的話返回模型book結構體對象
// 	nerBook := book.CreateBook()//直接調用createbook去添加紀錄

// 	//把結構體對象傳入通知
// 	res, _ := json.Marshal(newBook)//把添加的新書json化
// 	//設置響應頭
// 	setSameHeader(w)//並且返回給調用端
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(ResponseResult{
// 		Result: "CreateBook",
// 	})
// }

//------------------------get one data from mysql-------------------
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//首先拿到url裡的playerid
	createId := vars["id"]
	//類型轉換
	id, err := strconv.ParseInt(createId, 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		//沒發錯的話調用model中的方法
		create, _ := model.GetPlayer(id)
		res, err := json.Marshal(create)
		if err != nil {
			fmt.Println(err)
		}
		setSameHeader(w)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

//------------------------get ALL data from mysql-------------------
func Listplayers(w http.ResponseWriter, r *http.Request) {
	players := model.GetAllPlayers()
	res, err := json.Marshal(players)
	if err != nil {
		fmt.Println(err)
	}
	//設置響應頭
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//-------------------------delete data from mysql--------------------------------------------

func Deleteplayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//通過url傳遞過來的id
	createId := vars["id"]
	//進行數字解析
	id, err := strconv.ParseInt(createId, 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		//沒發錯的話調用model中的方法
		create := model.DeletePlayer(id)
		res, _ := json.Marshal(create)
		setSameHeader(w)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

//------------------------------update date from mysql----------------------
// func Updateplayer(w http.ResponseWriter, r *http.Request) {
// 	var updateplayer = &model.Createdemo{}
// 	//通過請求流接收到body把送過來的請求數據放到updateplayer這個結構體當中
// 	utils.ParseRequestBody(r, updateplayer)
// 	vars := mux.Vars(r)
// 	//拿到id
// 	updateId := vars["playerID"]
// 	//整形轉換
// 	playerID, err := strconv.ParseInt(updateId, 0, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		//從數據庫中讀到指定玩家
// 		update, db := model.UpdataPlayer(updateId)
// 		if updateplayer.PlayerID != "" {
// 			update.PlayerID = updateplayer.PlayerID
// 		}

// 		fmt.Println(playerID)
// 		db.Save(&update)
// 		res, _ := json.Marshal(update)
// 		setSameHeader(w)
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(res)
// 	}
// }

//----------------------------------get data from mysql 方法1-------------------------------
// localhost:9999/user/get/2
// var db *sql.DB

// func init() {
// 	log.Println(">>>> get database connection start <<<<")
// 	db = &sql.DB{}
// }

// func QueryById(context *gin.Context) {
// 	println(">>>> get user by id and name action start <<<<")

// 	// 獲取請求引數
// 	id := context.Param("id")

// 	// 查詢資料庫
// 	rows := db.QueryRow("select player_id,currency,time,id from createdemos where id = ? ", id)

// 	var user model.Createdemo
// 	//var Id uint16
// 	//var address string
// 	//var age uint8
// 	//var mobile string
// 	//var sex string
// 	err := rows.Scan(&user.PlayerID, &user.Currency, &user.Time, &user.Id)

// 	checkError(err)

// 	checkError(err)
// 	context.JSON(200, gin.H{
// 		"result": user,
// 	})
// }
// func checkError(e error) {
// 	if e != nil {
// 		log.Fatal(e)
// 	}
// }

//-----------------------------------get data from mysql 方法2---------------------------------
// func GetOne(c *gin.Context) {
// 	ids := c.Param("id")
// 	id, _ := strconv.Atoi(ids)
// 	p := model.Createdemo{
// 		Id: int64(id),
// 	}
// 	rs, _ := p.GetRow()
// 	c.JSON(http.StatusOK, gin.H{
// 		"result": rs,
// 	})
// }

//--------------------------------------------------------------------------------------------
