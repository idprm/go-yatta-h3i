package main

import (
	"github.com/gin-gonic/gin"
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/helpers"
	"waki.mobi/go-yatta-h3i/src/routers"
)

func init() {
	// Setup database
	database.Connect()
	// Setup logging
	helpers.WriteLog()
}

func main() {
	// Setup routing rules
	r := routers.SetupRouter()
	// Setup Trusted IP
	r.SetTrustedProxies([]string{"192.168.1.2"})
	// Logger
	r.Use(gin.Logger())

	r.Run(":8080")
}
