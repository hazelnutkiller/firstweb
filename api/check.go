package api

import (
	"encoding/json"
	"firstweb/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	urlstr := "https://uat-op-api.bpweg.com/check"
	values := url.Values{}
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")

	values.Set("operatorID", operatorID)
	values.Set("appSecret", appSecret)

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "appSecret"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	// Step 2: Check length
	if len(operatorID) > 20 {
		utils.ErrorResponse(c, 400, "Maximum request length exceeded", nil)
		return
	}

	r, err := http.PostForm(urlstr, values)
	if err != nil {
		log.Fatal(err)
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
