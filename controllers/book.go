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
// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	book := &models.Book{}
// 	//拿到用戶請求流調用實用類進行解析
// 	utils.ParseRequestBody(r, book)//book傳入後通知對象進行解析
// 	//把請求流中的json各式各樣的性質付給到結構體當中
// 	fmt.Println(book)
// //成功的話返回模型book結構體對象
// 	nerBook := book.CreateBook()//直接調用createbook去添加紀錄

// 	//把結構體對象傳入通知
// 	res, _ := json.Marshal(newBook)//把添加的新書json化
// 	//設置響應頭
// 	setSameHeader(w)//並且返回給調用端
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(ResponseResult{
// 		Result: "CreateBook",
// 	})
// }

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
