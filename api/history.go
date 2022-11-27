package api

import (
	"encoding/json"
	"firstweb/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/test/data"
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

type DataInfo struct {
	DataCount string         `json:"dataCount"`
	Data      []TransferInfo `json:"data"`
}

type BetInfo struct {
	BetID          string            `json:"betID"`
	OperatorID     string            `json:"operatorID"`
	PlayerID       string            `json:"playerID"`
	WEPlayerID     string            `json:"wePlayerID"`
	BetDateTime    int64             `json:"betDateTime"`
	SettlementTime int64             `json:"settlementTime"`
	BetStatus      string            `json:"betStatus"`
	BetCode        string            `json:"betCode"`
	ValidBetAmount int64             `json:"validBetAmount"`
	GameResult     string            `json:"gameResult"`
	Device         string            `json:"device"`
	BetAmount      int64             `json:"betAmount"`
	WinlossAmount  int64             `json:"winlossAmount"`
	Category       string            `json:"category"`
	GameType       string            `json:"gameType"`
	GameRoundID    string            `json:"gameRoundID"`
	IP             string            `json:"ip"`
	UID            string            `json:"uid"`
	RefID          string            `json:"RefID"`
	CardResult     map[string]string `json:"cardresult"`
	TransferType   string            `json:"transferType"`
	TransferTime   int64             `json:"transferTime"`
	TranAmount     int64             `json:"tranAmount"`
	Balance        int64             `json:"balance"`
}

func ConvertToTransferInfo(val BetInfo) TransferInfo {
	m := TransferInfo{
		OperatorID:   val.OperatorID,
		PlayerID:     val.PlayerID,
		UID:          val.UID,
		RefID:        val.RefID,
		TransferType: val.TransferType,
		TransferTime: val.TransferTime,
		TranAmount:   val.TranAmount,
		Balance:      val.Balance,
	}

	// switch val.TransferType {
	// case
	// 	m.TransferType = "deposit"
	// case fundtran.FundTranType_FUND_FROM_PG:
	// 	m.TransferType = "withdraw"
	// }

	return m
}

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

var checkGameType map[string]bool = map[string]bool{
	"BAC":  true,
	"BAS":  true,
	"BAI":  true,
	"DT":   true,
	"BAM":  true,
	"BAB":  true,
	"DTB":  true,
	"BAMB": true,
	"BASB": true,
	"ZJH":  true,
	"OX":   true,
	"ZJHB": true,
	"OXB":  true,
	"BAA":  true,
	"DTS":  true,
	"BAL":  true,
}

func ConvertToBetInfo(val map[string]string) BetInfo {
	betDateTime, _ := strconv.ParseInt(val["betdatetime"], 10, 64)
	settlementTime, _ := strconv.ParseInt(val["settlementtime"], 10, 64)
	betAmount, _ := strconv.ParseInt(val["betamount"], 10, 64)
	winlossAmount, _ := strconv.ParseInt(val["winlossamount"], 10, 64)
	validBetAmount, _ := strconv.ParseInt(val["validbetamount"], 10, 64)
	cardresult := ""
	if _, ok := checkGameType[val["gametype"]]; ok {
		cardresult = val["cardresult"]
	}
	m := BetInfo{
		BetID:          val["betid"],
		OperatorID:     val["operatorid"],
		PlayerID:       val["opplayerid"],
		WEPlayerID:     val["playerid"],
		BetDateTime:    betDateTime / 1000,
		SettlementTime: settlementTime / 1000,
		BetCode:        val["betcode"],
		ValidBetAmount: validBetAmount,
		GameResult:     val["gameresult"],
		Device:         val["device"],
		BetStatus:      val["betstatus"],
		BetAmount:      betAmount,
		WinlossAmount:  winlossAmount,
		Category:       val["category"],
		GameType:       val["gametype"],
		GameRoundID:    val["gameroundid"],
		IP:             val["ip"],
		CardResult:     ConvertCardResult(cardresult),
	}

	// {"A1":"spade5","A2":"spade9","A3":"heartj","B1":"spade4","B2":"spade2","B3":""}

	return m
}

//------------------------------------------------------------------------------------------------------------

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

	uid := utils.Generate()
	fmt.Println(uid)

	//請求需求欄位
	values.Set("operatorID", operatorID)
	values.Set("startTime", startTime)
	values.Set("endTime", endTime)
	values.Set("playerID", opPlayerID)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)
	values.Set("uid", uid)

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

	if uid == "" && (startTime == "" || endTime == "") {
		utils.ErrorResponse(c, 400, "Missing parameter: endTime|startTime", nil)
		return
	}

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
	st := (c.PostForm("appSecret") + c.PostForm("endTime") + c.PostForm("limit") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime + c.PostForm("startTime") + uid)
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
	//容器轉換
	count := 0
	res := []TransferInfo{}
	for _, val := range data.Data {
		res = append(res, ConvertToTransferInfo(val))
		count++
	}

	defer r.Body.Close()
	//读取整个响应体
	body, _ := ioutil.ReadAll(r.Body)
	var data DataInfo
	json.Unmarshal(body, &data)
	c.JSON(200, gin.H{"dataCount": count, "data": res})
	//打印看返回的cjson是什麼
	fmt.Println("data json:", data)

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
	//打印看返回的cjson是什麼
	fmt.Println("data json:", data)
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

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/report/bet", strings.NewReader(values.Encode()))
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
