package controllers

import (
	"io"
	"net/http"
)

func DeliveryReport(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 1")
}
