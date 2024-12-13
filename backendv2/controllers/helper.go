package controllers

import (
	"github.com/mmycin/backendv2/database"
	"github.com/mmycin/backendv2/models"
)

var colName string = database.ColName

func init() {
	database.Connect()
}

func reverseSlice(notices []models.Notice) []models.Notice {
	for i, j := 0, len(notices)-1; i < j; i, j = i+1, j-1 {
		notices[i], notices[j] = notices[j], notices[i]
	}
	return notices
}