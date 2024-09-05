package database

import (
	"github.com/mmycin/mongorm"
	"github.com/mmycin/mongorm/utils"
)

const URI string = "mongodb+srv://Mycin:myc23084.fun@cluster0.yzel00n.mongodb.net/"
const ColName string = "notice"
const dbName string = "Notice"

func Connect() {
	err := mongorm.Initialize(URI, dbName)
	utils.HandleError(err)
}