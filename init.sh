cat <<EOF > main.go
//腳本中創立main.go 啟動腳本
package main

import "fmt"

func main() {
    fmt.Println("努力學習得第一")
}

EOF
//建出原文件夾
mkdir src 
//全局設置文件
mkdir src/config
touch src/config/app.go

//控制器
mkdir src/controllers
touch src/controllers/book.go

//數據庫模型
mkdir src/models
touch src/models/book.go

//路由器
mkdir src/routes
touch src/routes/routes.go

//實用程序函數
mkdir src/utils
touch src/utils/utils.go
