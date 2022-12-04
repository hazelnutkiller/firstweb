package controllers

//controllers 控制器處理具體業務邏輯
//接收路由響應 讀取數據庫的地方 連接到數據庫模型

import (
	"encoding/json"
	"net/http"
)

//生成標準返回結構類型

type ResponseResult struct {
	Result string `json:"result"`
}

//設置共同頭部信息

func setSameHeader(w http.ResponseWriter) {
	w.Header().Set("content-Type", "application/json")
}

//定義各個處理函數
func CreateBook(w http.ResponseWriter, r *http.Request) {
	//設置響應頭
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "CreateBook",
	})
}

func ListBook(w http.ResponseWriter, r *http.Request) {
	//設置響應頭
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "ListBook",
	})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//設置響代碼
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseResult{
		//返回內容
		Result: "DeleteBook",
	})
}
