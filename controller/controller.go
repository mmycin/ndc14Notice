package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllNotice(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, FindAll())
}