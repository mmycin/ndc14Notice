package models

import "github.com/mmycin/mongorm/model"

type Notice struct {
	model.BaseModel `json:"-"`
	Title           string `json:"title"`
	Date            string `json:"date"`
	Content         string `json:"content"`
}
