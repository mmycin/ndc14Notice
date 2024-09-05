package model

import "github.com/mmycin/mongorm/model"

type Notice struct {
	model.BaseModel
	Date    string `json:"date"`
	Title	string	`json:"title"`
	Content string `json:"content"`
}
