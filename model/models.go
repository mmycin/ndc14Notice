package model

import "github.com/mmycin/mongorm/model"

type Notice struct {
	model.BaseModel
	Date    string `json:"date"`
	Content string `json:"content"`
}
