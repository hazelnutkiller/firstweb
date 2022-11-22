package routers

import (
	"firstweb/api"
	"firstweb/logrus"
	"firstweb/utils"
	"log"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.New()
	//router := gin.Default()
	router.Use(utils.Timeout())
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

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "signature"},
		AllowAllOrigins:  true,
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
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

	router.Run(":9999")
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
