package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/models"
)

type requestMO struct {
	MobileNo  string `json:"mobile_no"`
	ShortCode string `json:"short_code"`
	Message   string `json:"message"`
}

func MessageOriginated(c *gin.Context) {

	// http://35.247.131.49/moh3i?mobile_no=6289501845333&short_code=99879&message=reg+keren

	msisdn := c.Query("mobile_no")
	message := c.Query("message")
	code := c.Query("short_code")

	// CHECK IN BLACKLIST
	var blacklist models.Blacklist
	database.Database.Db.Where("msisdn", msisdn).First(&blacklist)

	// CHECK IN PERIODE
	var config models.Config
	database.Database.Db.Where("name", message).First(&config)

	// CHECK IN SERVICE
	var service models.Service
	database.Database.Db.Where("code", code).First(&service)

	// FILTER MESSAGE

	// SPLIT MESSAGE

	c.XML(http.StatusOK, gin.H{"status": "OK"})
}
