package util

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetDB() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config  app inited 。。。。")
	// fmt.Print()
}
