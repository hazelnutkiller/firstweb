package api

import (
	"encoding/json"
	"firstweb/model"
	"firstweb/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var CardValues = map[string]string{
	"cluba": "club_a", "club2": "club_2", "club3": "club_3", "club4": "club_4", "club5": "club_5", "club6": "club_6", "club7": "club_7", "club8": "club_8", "club9": "club_9", "club10": "club_10", "clubj": "club_j", "clubq": "club_q", "clubk": "club_k", "hearta": "heart_a", "heart2": "heart_2", "heart3": "heart_3", "heart4": "heart_4", "heart5": "heart_5", "heart6": "heart_6", "heart7": "heart_7", "heart8": "heart_8", "heart9": "heart_9", "heart10": "heart_10", "heartj": "heart_j", "heartq": "heart_q", "heartk": "heart_k", "diamonda": "diamond_a", "diamond2": "diamond_2", "diamond3": "diamond_3", "diamond4": "diamond_4", "diamond5": "diamond_5", "diamond6": "diamond_6", "diamond7": "diamond_7", "diamond8": "diamond_8", "diamond9": "diamond_9", "diamond10": "diamond_10", "diamondj": "diamond_j", "diamondq": "diamond_q", "diamondk": "diamond_k", "spadea": "spade_a", "spade2": "spade_2", "spade3": "spade_3", "spade4": "spade_4", "spade5": "spade_5", "spade6": "spade_6", "spade7": "spade_7", "spade8": "spade_8", "spade9": "spade_9", "spade10": "spade_10", "spadej": "spade_j", "spadeq": "spade_q", "spadek": "spade_k",
}

type TransferInfo struct {
	OperatorID   string `json:"operatorID"`
	PlayerID     string `json:"playerID"`
	UID          string `json:"uid"`
	RefID        string `json:"refID"`
	TransferType string `json:"transferType"`
	TransferTime int64  `json:"transferTime"`
	TranAmount   int64  `json:"tranAmount"`
	Balance      int64  `json:"balance"`
}

type BetInfo struct {
	Id             int64  `json:"id" form:"id"`
	BetID          string `json:"betID" form:"betID"`
	OperatorID     string `json:"operatorID" form:"operatorID"`
	PlayerID       string `json:"playerID" form:"playerID"`
	WEPlayerID     string `json:"wEPlayerID" form:"wEPlayerID"`
	BetDateTime    int64  `json:"betDateTime" form:"betDateTime"`
	SettlementTime int64  `json:"settlementTime" form:"settlementTime"`
	BetStatus      string `json:"betStatus" form:"betStatus"`
	BetCode        string `json:"betCode" form:"betCode"`
	ValidBetAmount int64  `json:"validBetAmount" form:"validBetAmount"`
	GameResult     string `json:"gameResult" form:"gameResult"`
	//Device         string `json:"device" form:"device"`
	BetAmount     int64 `json:"betAmount" form:"betAmount"`
	WinlossAmount int64 `json:"winlossAmount" form:"winlossAmount"`
	//Category       string `json:"category" form:"category"`
	GameType    string `json:"gameType" form:"gameType"`
	GameRoundID string `json:"gameRoundID" form:"gameRoundID"`
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

// func ConvertToTransferInfo(val TransferInfo) TransferInfo {
// 	m := TransferInfo{
// 		OPERATORID:   val.OperatorID,
// 		PLAYERID:     val.PlayerID,
// 		UID:          val.UID,
// 		REFID:        val.RefID,
// 		TRANSFERTYPE: val.TransferType,
// 		TRANSFERTIME: val.TransferTime,
// 		TRANAMOUNT:   val.TranAmount,
// 		BALANCE:      val.Balance,
// 	}

// switch val.TransferType{
// case TransferInfo:
// 	m.TransferType := "deposit"
// case TransferInfo:
// 	m.TransferType := "withdraw"
// }

// 	return m
// }

func ConvertCardResult(in string) map[string]string {

	if in != "" {

		out := map[string]string{}

		err := json.Unmarshal([]byte(in), &out)
		if err == nil {
			for k, v := range out {
				out[k] = CardValues[v]
			}
			return out
		}

	}

	return nil
}

// var checkGameType map[string]bool = map[string]bool{
// 	"BAC":  true,
// 	"BAS":  true,
// 	"BAI":  true,
// 	"DT":   true,
// 	"BAM":  true,
// 	"BAB":  true,
// 	"DTB":  true,
// 	"BAMB": true,
// 	"BASB": true,
// 	"ZJH":  true,
// 	"OX":   true,
// 	"ZJHB": true,
// 	"OXB":  true,
// 	"BAA":  true,
// 	"DTS":  true,
// 	"BAL":  true,
// }

//------------------------------------------------------------------------------------------------------------type BetTranInfo struct {
type TranInfo struct {
	DataCount int            `json:"dataCount"`
	Data      []TransferInfo `json:"data"`
}

func HistoryTransfer(c *gin.Context) {
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	opPlayerID := c.PostForm("playerID")
	//uid := c.PostForm("uid")
	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	appSecret := c.PostForm("appSecret")
	limit := c.PostForm("limit")

	requestTime := utils.Time()
	fmt.Println(requestTime)

	//自動取得uid

	// uid := utils.Generate()
	// fmt.Println(uid)

	//請求需求欄位
	values.Set("operatorID", operatorID)
	values.Set("startTime", startTime)
	values.Set("endTime", endTime)
	values.Set("playerID", opPlayerID)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)
	//values.Set("uid", uid)

	sTime, eTime := 0, 0

	//資料筆數 預設:50
	if limit == "" {
		limit = "50"
	}

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	// if uid == "" && (startTime == "" || endTime == "") {
	// 	utils.ErrorResponse(c, 400, "Missing parameter: endTime|startTime", nil)
	// 	return
	// }

	rtErr := utils.CheckRequestTime(requestTime)
	if rtErr != nil {
		utils.ErrorResponse(c, 400, "Incorrect requestTime", rtErr)
		return
	}
	//資料限制筆數 最大:500
	iLimit, err := strconv.Atoi(limit)
	if err != nil {
		utils.ErrorResponse(c, 400, "Incorrect limit format: "+limit, nil)
		return
	} else if iLimit > 500 {
		iLimit = 500
	}

	if startTime != "" && endTime != "" {
		sTime, err = strconv.Atoi(startTime)
		if err != nil {
			utils.ErrorResponse(c, 400, "Incorrect startTime format: "+startTime, nil)
			return
		}

		eTime, err = strconv.Atoi(endTime)
		if err != nil {
			utils.ErrorResponse(c, 400, "Incorrect endTime format: "+endTime, nil)
			return
		}
		//結束與開始時間間隔不得大於一個月
		if (eTime - sTime) > 2592000 {
			utils.ErrorResponse(c, 400, "Incorrect time period", nil)
			return
		}
	}
	//轉帳記錄簽名組成
	st := (c.PostForm("appSecret") + c.PostForm("endTime") + c.PostForm("limit") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime + c.PostForm("startTime"))
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/history/transfer", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", md5Str)
	clt := &http.Client{}
	r, _ := clt.Do(req)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	//读取整个响应体
	body, _ := ioutil.ReadAll(r.Body)
	var history TranInfo
	json.Unmarshal(body, &history)

	c.JSON(200, gin.H{
		"DataCount": history.DataCount,
		"Data":      history.Data,
	})
}

//---------------------------------------------------------------------------------------------------------
type BetTranInfo struct {
	TotalCount int       `json:"totalCount"`
	DataCount  int       `json:"dataCount"`
	Limit      int       `json:"limit"`
	Data       []BetInfo `json:"data"`
}

func HistoryBet(c *gin.Context) {

	// d := ConvertCardResult(`{"A1":"spade5","A2":"spade9","A3":"heartj","B1":"spade4","B2":"spade2","B3":""}`)

	// c.JSON(200, gin.H{"data": d})
	// return
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	agentID := c.PostForm("agentID")
	playerID := c.PostForm("playerID")
	betID := c.PostForm("betID")
	category := c.PostForm("category")
	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	limit := c.PostForm("limit")
	offset := c.PostForm("offset")
	betstatus := c.PostForm("betstatus")
	isSettlementTime := c.PostForm("isSettlementTime")
	appSecret := c.PostForm("appSecret")

	requestTime := utils.Time()
	fmt.Println(requestTime)

	//請求需求欄位
	values.Set("agentID", agentID)
	values.Set("operatorID", operatorID)
	values.Set("startTime", startTime)
	values.Set("endTime", endTime)
	values.Set("playerID", playerID)
	values.Set("betID", betID)
	values.Set("category", category)
	values.Set("betstatus", betstatus)
	values.Set("limit", limit)
	values.Set("offset", offset)
	values.Set("requestTime", requestTime)
	values.Set("isSettlementTime", isSettlementTime)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)

	//預設資料筆數為50筆
	if limit == "" {
		limit = "50"
	}
	//略過筆數預設為0
	if offset == "" {
		offset = "0"
	}

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "startTime", "endTime"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	//資料限制筆數最大:500
	iLimit, err := strconv.Atoi(limit)
	if err != nil {
		utils.ErrorResponse(c, 400, "Incorrect limit format: "+limit, nil)
		return
	} else if iLimit > 500 {
		iLimit = 500
	}

	sTime, err := strconv.Atoi(startTime)
	if err != nil {
		utils.ErrorResponse(c, 400, "Incorrect startTime format: "+startTime, nil)
		return
	}

	eTime, err := strconv.Atoi(endTime)
	if err != nil {
		utils.ErrorResponse(c, 400, "Incorrect endTime format: "+endTime, nil)
		return
	}
	//結束與開始時間間隔不得大於一個月
	if (eTime - sTime) > 2592000 {
		utils.ErrorResponse(c, 400, "Incorrect time period", nil)
		return
	}

	//轉帳記錄簽名組成
	st := (c.PostForm("appSecret") + c.PostForm("betID") + c.PostForm("betstatus") + c.PostForm("category") + c.PostForm("endTime") + c.PostForm("isSettlementTime") + c.PostForm("limit") + c.PostForm("offset") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime + c.PostForm("startTime"))
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/history/bet", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", md5Str)
	clt := &http.Client{}
	r, _ := clt.Do(req)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	//读取整个响应体
	body, _ := ioutil.ReadAll(r.Body)
	var data BetTranInfo
	json.Unmarshal(body, &data)
	c.JSON(200, gin.H{
		"DATACOUNT":  data.DataCount,
		"TOTALCOUNT": data.TotalCount,
		"LIMIT":      data.Limit,
		"DATA":       data.Data,
	})

}

//------------------------------------------------------------------------------------------------------------
func HistorySummary(c *gin.Context) {

	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	appSecret := c.PostForm("appSecret")

	requestTime := utils.Time()
	fmt.Println(requestTime)

	//請求需求欄位

	values.Set("operatorID", operatorID)
	values.Set("startTime", startTime)
	values.Set("endTime", endTime)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "startTime", "endTime"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	sTime, err := strconv.Atoi(startTime)
	if err != nil {
		utils.ErrorResponse(c, 400, "Incorrect startTime format: "+startTime, nil)
		return
	}

	eTime, err := strconv.Atoi(endTime)
	if err != nil {
		utils.ErrorResponse(c, 400, "Incorrect endTime format: "+endTime, nil)
		return
	}
	//結束與開始時間間隔不得大於一個月
	if (eTime - sTime) > 2592000 {
		utils.ErrorResponse(c, 400, "Incorrect time period", nil)
		return
	}

	//投注記錄統計簽名組成
	st := (c.PostForm("appSecret") + c.PostForm("endTime") + c.PostForm("operatorID") + requestTime + c.PostForm("startTime"))
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/history/summary", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", md5Str)
	clt := &http.Client{}
	r, _ := clt.Do(req)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	//读取整个响应体
	body, _ := ioutil.ReadAll(r.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	c.JSON(200, data)
	//打印看返回的cjson是什麼
	fmt.Println("data json:", data)
}

//------------------------------------------------------------------------------------------------------------

type ReportInfo struct {
	TotalCount int       `json:"totalCount"`
	DataCount  int       `json:"dataCount"`
	Data       []BetInfo `json:"data"`
}

func ReportBet(c *gin.Context) {
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	agentID := c.PostForm("agentID")
	playerID := c.PostForm("playerID")
	betID := c.PostForm("betID")
	category := c.PostForm("category")
	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	limit := c.PostForm("limit")
	offset := c.PostForm("offset")
	betstatus := c.PostForm("betstatus")
	isSettlementTime := c.PostForm("isSettlementTime")
	appSecret := c.PostForm("appSecret")

	requestTime := utils.Time()
	fmt.Println(requestTime)

	//請求需求欄位
	values.Set("agentID", agentID)
	values.Set("operatorID", operatorID)
	values.Set("startTime", startTime)
	values.Set("endTime", endTime)
	values.Set("playerID", playerID)
	values.Set("betID", betID)
	values.Set("category", category)
	values.Set("betstatus", betstatus)
	values.Set("limit", limit)
	values.Set("offset", offset)
	values.Set("requestTime", requestTime)
	values.Set("isSettlementTime", isSettlementTime)
	values.Set("appSecret", appSecret)

	//預設資料限制筆數為50筆
	if limit == "" {
		limit = "50"
	}
	//略過筆數預設為0
	if offset == "" {
		offset = "0"
	}

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "startTime", "endTime"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	//資料限制筆數最大:500
	iLimit, err := strconv.Atoi(limit) //字串轉int：Atoi()
	if err != nil {
		utils.ErrorResponse(c, 400, "Incorrect limit format: "+limit, nil)
		return
	} else if iLimit > 500 {
		iLimit = 500
	}
	//驗證時間格式
	sTime, err := strconv.Atoi(startTime)
	if err != nil {
		utils.ErrorResponse(c, 400, "Incorrect startTime format: "+startTime, nil)
		return
	}
	//驗證時間格式
	eTime, err := strconv.Atoi(endTime)
	if err != nil {
		utils.ErrorResponse(c, 400, "Incorrect endTime format: "+endTime, nil)
		return
	}
	//結束與開始時間間隔不得大於一個月
	if (eTime - sTime) > 2592000 {
		utils.ErrorResponse(c, 400, "Incorrect time period", nil)
		return
	}

	//轉帳記錄簽名組成
	st := (c.PostForm("appSecret") + c.PostForm("endTime") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime + c.PostForm("startTime"))
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/report/bet", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", md5Str)
	clt := &http.Client{}
	r, _ := clt.Do(req)
	if r.StatusCode == 400 {
		utils.ErrorResponse(c, 400, "Incorrect operatorID", err)
		return
	}
	if r.StatusCode == 401 {
		utils.ErrorResponse(c, 401, "Incorrect signature", err)
		return
	}
	if r.StatusCode == 409 { //**需要調整**
		utils.ErrorResponse(c, 409, "PlayerID already exists", err)
		return
	}
	if r.StatusCode != 200 {
		panic(r)
	}

	defer r.Body.Close()
	//读取整个响应体
	body, _ := ioutil.ReadAll(r.Body)

	//定義data使用的類別
	var data model.BetInfo1
	var repo model.BetInfo
	//data.PlayerID = playerID
	//json.Unmarshal([]byte(body), &data)

	json.Unmarshal(body, &data)
	sis := string(data)
	fmt.Println(sis)
	json.Unmarshal(body, &repo)
	//帶入設定的結構帶進表格
	betrecord := &model.BetInfo1{
		WEPlayerID:     data.WEPlayerID,
		PlayerID:       data.PlayerID,
		OperatorID:     data.OperatorID,
		BetID:          data.BetID,
		BetDateTime:    data.BetDateTime,
		SettlementTime: data.SettlementTime,
		BetStatus:      data.BetStatus,
		ValidBetAmount: data.ValidBetAmount,
		WinlossAmount:  data.WinlossAmount,
		BetAmount:      data.BetAmount,
		GameRoundID:    data.GameRoundID,
		GameType:       data.GameType,
		GameResult:     data.GameResult,
		BetCode:        data.BetCode,
		//CardResult:     data.CardResult,
	}
	id := betrecord.BetRecord()
	fmt.Println(id)

	c.JSON(200, gin.H{
		"DATACOUNT":  repo.DataCount,
		"TOTALCOUNT": repo.TotalCount,
		"DATA":       repo.Data,
	})
}
