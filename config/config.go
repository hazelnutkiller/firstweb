package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Config() {
	//要讀取的檔名
	viper.SetConfigName("app")
	//要讀取的附檔名
	viper.SetConfigType("yaml")
	//要讀取的路徑
	viper.AddConfigPath("./config")
	//設定參數預設值
	viper.SetDefault("application.port", 9999)
	//readinconfig 讀取設定檔
	err := viper.ReadInConfig()
	if err != nil {
		panic("讀取設定檔出現錯誤，原因為：" + err.Error())
	}
	fmt.Println("application port = " + viper.GetString("application.port"))
}
