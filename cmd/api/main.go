package main

import (
	"fmt"

	"github.com/markgravity/golang-ic/bootstrap"
	"github.com/markgravity/golang-ic/helpers/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	bootstrap.Init()

	r := bootstrap.SetupRouter()

	err := r.Run(getAppPort())
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func getAppPort() string {
	if gin.Mode() == gin.ReleaseMode {
		return fmt.Sprint(":", viper.GetString("PORT"))
	}

	return fmt.Sprint(":", viper.GetString("app_port"))
}
