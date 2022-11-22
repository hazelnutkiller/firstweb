package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//amount進來的型別是字串，把字串轉為數字才能計算或比較大小

func CheckAmount(amount string) (int64, error) {
	gainBal, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return 0, err
	} else if gainBal < 1 {
		return 0, fmt.Errorf("amount < 1: %v", gainBal)
	}
	return gainBal, nil
}

//
func CheckPostFormData(c *gin.Context, vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(c.PostForm(v)) == "" {
			return v
		}
	}
	return ""
}

func CheckRequestTime(requestTime string) error {
	//驗證時間格式
	rtInt, rtErr := strconv.ParseInt(requestTime, 10, 64)
	if rtErr != nil {
		return fmt.Errorf("incorrect format:\r\nrequestTime :%s", requestTime)
	}
	//超過前後兩分鐘
	if rtInt-time.Now().Unix() > 120 || time.Now().Unix()-rtInt > 120 {
		return fmt.Errorf("expired:\r\nrequestTime :%s", requestTime)
	}
	return nil
}

func GetSignature(singSource string) string {
	data := []byte(singSource)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
	//if !strings.EqualFold(signature, singMD5) {
	//return fmt.Errorf("\r\nrequest :%s\r\ngenerate:%s", signature, singMD5)
	//}
	//return nil
}

func Time() string {

	return strconv.FormatInt(time.Now().Unix(), 10)

}

func Generate() string {

	uuid := uuid.New()
	uid := uuid.String()
	return uid
}
