package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//處理用戶發送過來的請求流解析body對象
func ParseRequestBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		//如果body與我們定的對象一致就會用Unmarshal變成結構體
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
