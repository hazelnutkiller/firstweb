package routers

//路由註冊

import (
	"firstweb/api"
	"firstweb/controllers"
	"firstweb/logrus"
	"firstweb/model"
	"fmt"

	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func Router() {

	//router := gin.New()
	router := gin.Default()

	//----------------------------------get data from mysql--------------------------------
	router.GET("/hello", func(context *gin.Context) {
		log.Println(">>>> hello gin start <<<<")
		context.JSON(200, gin.H{
			"code":    200,
			"success": true,
			"data":    "Hello LXR！",
		})
	})
	//---------------------------------------------------------------------------------------

	router.Use(logrus.Logrus())
	router.RedirectFixedPath = true

	// create limiter 一分鐘呼叫60次
	// store := memory.NewStore()
	//  rateLimit := viper.GetInt64("rate_limit")
	//  rate := limiter.Rate{
	//  	Period: time.Second * 60,
	//  	Limit:  rateLimit,
	// 	fmt.Println(rate)
	//  }
	// Alternatively, you can pass options to the limiter instance with several options.
	//limiterInstance := limiter.New(store, rate, limiter.WithTrustForwardHeader(true))

	//limitGroup := r.Group("/history", mgin.NewMiddleware(limiterInstance))

	// @Api:History@
	//limitGroup.POST("/transfer", api.HistoryTransfer)
	//limitGroup.POST("/bet", api.HistoryBet)
	//limitGroup.POST("/", api.PlayerCreate)
	// @end

	//跨來源資源共用
	router.Use(cors.New(cors.Config{
		//允許的HTTP Method
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		//允許的Header 信息
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "signature"},
		//允許的domain
		AllowAllOrigins: true,
		//允許請求包含驗證憑證
		AllowCredentials: false,
		//可被存取的時間
		MaxAge: 12 * time.Hour,
	}))

	//set Logger
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/Check", api.Check)
	router.POST("/player/login", api.PlayerLogin)
	router.POST("/player/create", api.PlayerCreate)
	router.POST("/player/deposit", api.PlayerDeposit)
	router.POST("/player/withdraw", api.PlayerWithdraw)
	router.POST("/player/logout", api.PlayerLogout)
	router.POST("/history/transfer", api.HistoryTransfer)
	router.POST("/history/bet", api.HistoryBet)
	router.POST("/history/summary", api.HistorySummary)
	router.POST("report/bet", api.ReportBet)

	router.HTMLRender = loadTemplates()

	router.NoRoute(func(c *gin.Context) {

		c.JSON(400, gin.H{"error": "Bad Request"})
	})

	model.APIServer = http.Server{
		Addr:    ":9999",
		Handler: router,
	}

}
func RunRouter() { //把listen的功能用go routine的方式丟到背景去運行
	//接到os強制關閉服務

	if err := model.APIServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func loadTemplates() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	includes, err := filepath.Glob("doc/page/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, include := range includes {
		r.AddFromFiles(filepath.Base(include), "doc/layout/base.html", include)
	}
	return r
}

//處理 /資源

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go world")
}

//指定一個路由函數 接收路由對象
var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", index).Methods("POST")
	router.HandleFunc("/listplayers", controllers.Listplayers).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPlayer).Methods("GET")

	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

}
