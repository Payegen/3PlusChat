package main

import (
	"fmt"
	"gochat/router"
	"gochat/util"

	"github.com/spf13/viper"
)

func main() {
	util.GetDB()
	fmt.Print(viper.GetString("mysql.url"))
	r := router.Router()

	r.Run(":8086")
}
