package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mmycin/mongorm"
	"github.com/mmycin/mongorm/utils"
)

var URI string
var ColName string
var dbName string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	URI = os.Getenv("URI")
	ColName = os.Getenv("ColName")
	dbName = os.Getenv("dbName")
}

func GetURI() string {
	return URI
}

func GetColName() string {
	return ColName
}

func GetDBName() string {
	return dbName
}

func Connect() {
	fmt.Println("Hello", URI, dbName, ColName)
	_, err := mongorm.Initialize(URI, dbName)
	utils.HandleError(err)
}