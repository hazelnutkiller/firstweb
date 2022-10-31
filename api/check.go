package api

import (
	"firstweb/utils"
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
	b, _ := ioutil.ReadAll(r.Body)
	c.JSON(200, string(b))
}
