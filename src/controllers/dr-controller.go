package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeliveryReport(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"status": "OK"})
}
