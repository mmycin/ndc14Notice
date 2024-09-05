package controller

import (
	"context"

	"github.com/mmycin/ndc14notice/database"
	"github.com/mmycin/ndc14notice/model"

	"github.com/mmycin/mongorm"
	"go.mongodb.org/mongo-driver/bson"
)

const colName string = database.ColName

func init() {
	database.Connect()
}

func FindAll() []model.Notice {
	var notices []model.Notice
	mongorm.ReadAll(context.Background(), colName, bson.M{}, &notices)
	
	// Reverse the slice
	for i, j := 0, len(notices)-1; i < j; i, j = i+1, j-1 {
		notices[i], notices[j] = notices[j], notices[i]
	}

	return notices
}
