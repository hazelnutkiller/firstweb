package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckPostFormData(c *gin.Context, vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(c.PostForm(v)) == "" {
			return v
		}
	}
	return ""
}

func CheckRequestTime(requestTime string) error {
	rtInt, rtErr := strconv.ParseInt(requestTime, 10, 64)
	if rtErr != nil {
		return fmt.Errorf("incorrect format:\r\nrequestTime :%s", requestTime)
	}

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

	return strconv.FormatInt(time.Now().Unix(), 10)
}
