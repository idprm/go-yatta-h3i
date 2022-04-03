package routers

import (
	"github.com/gin-gonic/gin"
	"waki.mobi/go-yatta-h3i/src/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/moh3i", controllers.MessageOriginated)
	r.GET("/drh3i", controllers.DeliveryReport)

	return r
}
