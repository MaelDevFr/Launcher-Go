package main

import (
	"fmt"
	"log"

	"webapp/config"
	"webapp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	var appConfig config.Config
	appConfig.Port = "8080"
	router := gin.Default()
	router.LoadHTMLGlob("./public/*")
	router.GET("/", handlers.HomeFunc)
	log.Fatal(router.Run(fmt.Sprintf(":%s", appConfig.Port)))
}