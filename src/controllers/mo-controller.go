package controllers

import (
	"io"
	"net/http"

	"waki.mobi/go-yatta-h3i/src/databases"
	"waki.mobi/go-yatta-h3i/src/models"
)

func MessageOriginated(w http.ResponseWriter, r *http.Request) {
	var blacklist models.Blacklist
	databases.Database.Db.Where("msisdn", r.Body).First(&blacklist)
	io.WriteString(w, "version 1")
}
