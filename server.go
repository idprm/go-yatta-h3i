package main

import (
	"log"
	"net/http"
	"time"

	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/controllers"
)

func init() {

}

func main() {
	// Set routing rules
	http.HandleFunc("/moh3i", controllers.MessageOriginated)
	http.HandleFunc("/drh3i", controllers.DeliveryReport)

	server := http.Server{
		Addr:           ":" + config.ViperEnv("APP_PORT"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe(), nil)
}
