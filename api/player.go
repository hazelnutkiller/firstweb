package api

import (
	"encoding/json"
	"firstweb/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

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

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "playerID", "nickname"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	// Step 2: Check length
	if len(operatorID) > 20 || len(opPlayerID) > 20 || len(nickname) > 200 {
		utils.ErrorResponse(c, 400, "Maximum request length exceeded", nil)
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
	if err != nil {
		panic(err)
	}
	//客戶端完成之後要關閉請求
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	c.JSON(200, string(body))

}

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

	// Step 4: Check Signature (appSecret + operatorID + playerID + requestTime)

	st := (c.PostForm("appSecret") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime)
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	//发送JSON数据的post请求
	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/login", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", md5Str)
	//設置超時
	clt := &http.Client{
		Timeout: time.Second * 1000,
	}
	r, _ := clt.Do(req)
	if err != nil {
		panic(err)
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

func PlayerDeposit(c *gin.Context) {
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	opPlayerID := c.PostForm("playerID")
	amount := c.PostForm("amount")
	appSecret := c.PostForm("appSecret")

	//取得requestTime
	requestTime := utils.Time()
	fmt.Println(requestTime)

	//自動取得uid

	uid := utils.Generate()
	fmt.Println(uid)

	//請求需求欄位
	values.Set("operatorID", operatorID)
	values.Set("playerID", opPlayerID)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)
	values.Set("amount", amount)
	values.Set("uid", uid)

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "playerID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	//存款簽名組成
	st := (c.PostForm("amount") + c.PostForm("appSecret") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime + uid)
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/deposit", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", md5Str)
	clt := &http.Client{}
	r, _ := clt.Do(req)
	if err != nil {
		panic(err)
	}
	//客戶端完成之後要關閉請求
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	c.JSON(200, string(body))

}

func PlayerWithdraw(c *gin.Context) {
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	opPlayerID := c.PostForm("playerID")
	amount := c.PostForm("amount")
	appSecret := c.PostForm("appSecret")

	//取得requestTime
	requestTime := utils.Time()
	fmt.Println(requestTime)

	//自動取得uid

	uid := utils.Generate()
	fmt.Println(uid)

	//請求需求欄位
	values.Set("operatorID", operatorID)
	values.Set("playerID", opPlayerID)
	values.Set("requestTime", requestTime)
	values.Set("appSecret", appSecret)
	values.Set("amount", amount)
	values.Set("uid", uid)

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "playerID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	//存款簽名組成
	st := (c.PostForm("amount") + c.PostForm("appSecret") + c.PostForm("operatorID") + c.PostForm("playerID") + requestTime + uid)
	md5Str := utils.GetSignature(st)
	fmt.Println(md5Str)

	req, err := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/withdraw", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", md5Str)
	clt := &http.Client{}
	r, _ := clt.Do(req)
	if err != nil {
		panic(err)
	}
	//客戶端完成之後要關閉請求
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	c.JSON(200, string(body))

}
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
	//客戶端完成之後要關閉請求
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	c.JSON(200, string(body))

}
