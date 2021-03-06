package controllers

import (
	"fmt"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
)

func Index(w http.ResponseWriter, request *http.Request) {
	db := models.GetDbSession()
	var boards []models.Board
	_, err := db.Select(&boards, "SELECT * FROM boards")
	if err != nil {
		fmt.Printf("[error] Could not get boards (%s)\n", err.Error())
	}

	utils.RenderTemplate(w, request, "index.html", map[string]interface{}{
		"boards": boards,
	})
}
