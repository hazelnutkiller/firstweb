package api

import (
	"encoding/json"
	"firstweb/model"
	"firstweb/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//-------------------------------------------------------------------------------------------------------------
func PlayerCreate(c *gin.Context) {
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	opPlayerID := c.PostForm("playerID")
	nickname := c.PostForm("nickname")
	appSecret := c.PostForm("appSecret")
	requestTime := utils.Time()
	fmt.Println(requestTime)

	//請求需求欄位
	values.Set("operatorID", operatorID)
	values.Set("playerID", opPlayerID)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)
	values.Set("nickname", nickname)

	// Step 1: Check the required parameters 驗證是否沒填
	if missing := utils.CheckPostFormData(c, "operatorID", "playerID", "nickname"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	// Step 2: Check length 驗證長度
	if len(operatorID) > 20 || len(opPlayerID) > 20 || len(nickname) > 200 {
		utils.ErrorResponse(c, 400, "Maximum request length exceeded", nil)
		return
	}

	// Step 3: Check request time 驗證請求時間
	rtErr := utils.CheckRequestTime(requestTime)
	if rtErr != nil {
		utils.ErrorResponse(c, 400, "Incorrect requestTime", rtErr)
		return
	}

	//簽名組成
	st := (appSecret + c.PostForm("nickname") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime)
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/create", strings.NewReader(values.Encode()))
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
	if r.StatusCode == 409 {
		utils.ErrorResponse(c, 409, "PlayerID already exists", err)
		return
	}
	if r.StatusCode != 200 {
		panic(r)
	}

	// //模型中serform結構體對象傳入
	// create := &model.Createdemo{}
	// //拿到用戶的請求流
	// utils.ParseRequestBody(c.Request, create)
	// fmt.Println(create)
	// //直接調用CreatePlayer去新增玩家
	// newCreate := create.CreatePlayer()
	// //把新創的玩家json化
	// json.Marshal(newCreate)
	//返回給調用端
	// setSameHeader(w)
	// w.WriteHeader(http.StatusOK)
	//c.Write(r)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	//读取整个响应体
	body, _ := ioutil.ReadAll(r.Body)
	//定義data使用的類別
	var data model.Userform

	json.Unmarshal(body, &data)
	//帶入設定的結構帶進兩個表格
	create := &model.Userform{
		PlayerID: data.PlayerID,
		Currency: data.Currency,
		Time:     data.Time,
	}

	addtrans := &model.Createdemo{
		PlayerID: data.PlayerID,
		Currency: data.Currency,
		Time:     data.Time,
	}

	Id := addtrans.Addplayer()
	fmt.Println(Id)
	id := create.CreatePlayer()
	msg := fmt.Sprint("insert successful ", create.ID)
	fmt.Println(id)
	c.JSON(http.StatusOK, gin.H{
		"Msg":      msg,
		"PlayerID": data.PlayerID,
		"Currency": data.Currency,
		"Time":     data.Time,
	})
}

//-----------------------------------------------------------------------------------------
func PlayerLogin(c *gin.Context) {
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	opPlayerID := c.PostForm("playerID")
	appSecret := c.PostForm("appSecret")
	requestTime := utils.Time()
	fmt.Println(requestTime)

	//請求需求欄位
	values.Set("operatorID", operatorID)
	values.Set("playerID", opPlayerID)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)

	// Step 1: 自動帶入簽名 (appSecret + operatorID + playerID + requestTime)

	st := (c.PostForm("appSecret") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime)
	md5Str := utils.GetSignature(st)
	//fmt.Println(md5Str)

	//发送JSON数据的post请求
	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/login", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", md5Str)

	//設置客戶請求5秒超時

	clt := &http.Client{
		Timeout: time.Second * 5,
	}
	r, _ := clt.Do(req)
	if err != nil {
		utils.ErrorResponse(c, 504, "Timeout", err)
		return
	}

	//客戶端完成之後要關閉請求

	defer r.Body.Close()
	//读取整个响应体
	body, _ := ioutil.ReadAll(r.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	c.JSON(200, data)
	//打印看返回的cjson是什麼
	fmt.Println("data json:", data)

}

//-------------------------------------------------------------------------------------------------------
func PlayerDeposit(c *gin.Context) {
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	opPlayerID := c.PostForm("playerID")
	amount := c.PostForm("amount")
	appSecret := c.PostForm("appSecret")

	//取得requestTime
	requestTime := utils.Time()
	//fmt.Println(requestTime)

	//自動取得uid

	uid := utils.Generate()
	//fmt.Println(uid)

	//請求需求欄位
	values.Set("operatorID", operatorID)
	values.Set("playerID", opPlayerID)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)
	values.Set("amount", amount)
	values.Set("uid", uid)

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "playerID", "appSecret", "amount"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}
	// Step 2: Check amount Int64 驗證金額格式

	gainBal, formatErr := utils.CheckAmount(amount)

	if formatErr != nil {
		utils.ErrorResponse(c, 400, "Incorrect amount format", formatErr)
		return
	}
	fmt.Println(gainBal)

	//存款簽名組成
	st := (c.PostForm("amount") + c.PostForm("appSecret") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime + uid)
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/deposit", strings.NewReader(values.Encode()))
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
	if r.StatusCode == 404 {
		utils.ErrorResponse(c, 404, "PlayerID not found", err)
		return
	}
	if r.StatusCode != 200 {
		panic(r)
	}

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	//读取整个响应体
	body, _ := ioutil.ReadAll(r.Body)

	var data model.Createdemo
	var updata model.Userform
	data.PlayerID = opPlayerID
	updata.PlayerID = opPlayerID

	json.Unmarshal(body, &data)
	json.Unmarshal(body, &updata)

	addtrans := &model.Createdemo{
		Deposit:  amount,
		Balance:  data.Balance,
		PlayerID: data.PlayerID,
		Currency: data.Currency,
		Time:     data.Time,
		RefID:    data.RefID,
	}
	//更新餘額
	err2 := model.UpdataBalance(&updata)
	if err2 != nil {
		fmt.Println(err2)
	}
	id := addtrans.Addtrans()
	msg := fmt.Sprint("insert successful ", addtrans.Balance)
	fmt.Println(msg)
	fmt.Println(id)

	c.JSON(http.StatusOK, gin.H{
		"Balance":  data.Balance,
		"Currency": data.Currency,
		"Time":     data.Time,
		"RefID":    data.RefID,
	})

}

//--------------------------------------------------------------------------------------------------------------
func PlayerWithdraw(c *gin.Context) {
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	opPlayerID := c.PostForm("playerID")
	amount := c.PostForm("amount")
	appSecret := c.PostForm("appSecret")

	//取得requestTime
	requestTime := utils.Time()
	//fmt.Println(requestTime)

	//自動取得uid

	uid := utils.Generate()

	//請求需求欄位
	values.Set("operatorID", operatorID)
	values.Set("playerID", opPlayerID)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)
	values.Set("amount", amount)
	values.Set("uid", uid)

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "playerID", "appSecret", "amount"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}
	// Step 2: Check amount Int64 驗證金額格式

	gainBal, formatErr := utils.CheckAmount(amount)

	if formatErr != nil {
		utils.ErrorResponse(c, 400, "Incorrect amount format", formatErr)
		return
	}
	fmt.Println(gainBal)

	//存款簽名組成
	st := (c.PostForm("amount") + c.PostForm("appSecret") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime + uid)
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/withdraw", strings.NewReader(values.Encode()))
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
	if r.StatusCode == 404 {
		utils.ErrorResponse(c, 404, "PlayerID not found", err)
		return
	}
	if r.StatusCode != 200 {
		panic(r)
	}

	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	//读取整个响应体
	body, _ := ioutil.ReadAll(r.Body)
	var data model.Createdemo
	var updata model.Userform
	data.PlayerID = opPlayerID
	updata.PlayerID = opPlayerID

	json.Unmarshal(body, &data)
	json.Unmarshal(body, &updata)

	addtrans := &model.Createdemo{
		Deposit:  amount,
		Balance:  data.Balance,
		PlayerID: data.PlayerID,
		Currency: data.Currency,
		Time:     data.Time,
		RefID:    data.RefID,
	}

	err2 := model.UpdataBalance(&updata)
	if err2 != nil {
		fmt.Println(err2)
	}
	id := addtrans.Addtrans()
	msg := fmt.Sprint("insert successful ", addtrans.Balance)
	fmt.Println(msg)
	fmt.Println(id)

	c.JSON(http.StatusOK, gin.H{
		"Balance":  data.Balance,
		"Currency": data.Currency,
		"Time":     data.Time,
		"RefID":    data.RefID,
	})
}

//-----------------------------------------------------------------------------------------------------------
func PlayerLogout(c *gin.Context) {
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	opPlayerID := c.PostForm("playerID")
	appSecret := c.PostForm("appSecret")
	errorResponseCode := c.PostForm("errorResponseCode")

	//取得requestTime
	requestTime := utils.Time()
	fmt.Println(requestTime)

	//請求需求欄位
	values.Set("operatorID", operatorID)
	values.Set("playerID", opPlayerID)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)
	values.Set("errorResponseCode", errorResponseCode)

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "playerID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}
	///
	//存款簽名組成
	st := (c.PostForm("appSecret") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime)
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/logout", strings.NewReader(values.Encode()))
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

	// Step 8: Add Transcation Record 增加交易記錄於db中
	//tranDate := time.Now().Unix()
	//	refID := xid.New().String()
	//rsp, err := srvclient.WalletClient.MWAddFundTran(context.TODO(), &wallet.MWAddFundTranRequest{
	//FundTran: &trn.FundTran{
	//	PlayerID:     data.PlayerProfile.PlayerID,
	//	OpPlayerID:   opPlayerID,
	//	OperatorID:   operatorID,
	//	BetID:        "", // ???
	//	PgFundTranID: refID,
	//	OpFundTranID: uid,
	//TranType:     trn.FundTranType_FUND_TO_PG,
	//TranDate:     tranDate,
	//	TranAmount:   gainBal,
	//	OpBalAmount:  0,
	// PgBalAmount:  currentBal,
	//	Reference: "", // ???
	//	},
	//	})
	//if err != nil {
	//utils.ErrorResponse(c, 500, "Transaction failed", err)
	//} else {
	// Check if balance is negative
	//playerBalance := int64(0)
	//if rsp.BalAmount > 0 {
	//	playerBalance = rsp.BalAmount
	//	}
	//	c.JSON(200, gin.H{
	//"balance":  playerBalance,
	//"currency": Currency.String(),
	//	"time":     tranDate,
	//	"refID":    refID,
	//	})
	// publish balance change message
	//go kafka.PublishBalanceChange("in", data.PlayerProfile.PlayerID, opPlayerID, operatorID, data.PlayerProfile.AgentID, rsp.BalAmount+gainBal, rsp.BalAmount)
	//	}
}
