package controllers

import (
	"github.com/mmycin/mongorm"
	"github.com/mmycin/backendv2/models"
)


func FindAll() []models.Notice {
	var notices []models.Notice
	mongorm.ReadAll(colName, &notices)
	
	// Reverse the slice
	notices = reverseSlice(notices)

	return notices
}
